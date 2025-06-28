package parser

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type TextParser struct {
}

type OrderLine struct {
	UserId    string    `json:"user_id"`
	Name      string    `json:"name"`
	OrderId   string    `json:"order_id"`
	ProductId string    `json:"product_id"`
	Price     float64   `json:"total"`
	OrderDate time.Time `json:"date"`
}

type Order struct {
	OrderId   string     `json:"order_id"`
	Total     float64    `json:"total"`
	OrderDate time.Time  `json:"date"`
	Products  []Products `json:"products"`
}

type Products struct {
	ProductId string  `json:"product_id"`
	Value     float64 `json:"value"`
}

type UserWithOrders struct {
	UserId    string  `json:"user_id"`
	Name      string  `json:"name"`
	Orders []Order `json:"orders"`
}

func NewTextParser() *TextParser {
	return &TextParser{}
}

func (t *TextParser) Parse(content []byte) ([]byte, error) {
	scanner := bufio.NewScanner(bytes.NewReader(content))
	var ordersMap = make(map[string][]OrderLine)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		orderLine, err := buildOrder(line)
		if err != nil {
			return nil, err
		}
		addOrderToMap(orderLine, ordersMap)
	}

	userWithOrdersList := processUserOrders(ordersMap)

	result, err := json.Marshal(userWithOrdersList)
	if err != nil {
		fmt.Errorf("erro ao converter text para JSON: %w", err)
	}

	return result, nil
}

func processUserOrders(ordersMap map[string][]OrderLine) []UserWithOrders {
	var userOrderList []UserWithOrders

	for orderId := range ordersMap {
		orders := ordersMap[orderId]
		products := make([]Products, len(orders))
		for _, orderLine := range orders {
			product := Products{
				ProductId: orderLine.ProductId,
				Value: orderLine.Price,
			}
			products = append(products, product)
		}
		userOrders := UserWithOrders{
			UserId:    orders[0].UserId,
			Name:      orders[0].Name,
			Orders: 
		}
		userOrderList = append(userOrderList, userOrders)
	}

	return userOrderList
}

func addOrderToMap(line OrderLine, ordersMap map[string][]OrderLine) {

	if orders, exists := ordersMap[line.OrderId]; exists {
		ordersMap[line.OrderId] = append(orders, line)
	} else {
		ordersMap[line.OrderId] = []OrderLine{line}
	}
}

func buildOrder(line string) (OrderLine, error) {
	obj := OrderLine{
		UserId:    strings.TrimSpace(line[0:10]),
		Name:      strings.TrimSpace(line[10:55]),
		OrderId:   strings.TrimSpace(line[55:65]),
		ProductId: strings.TrimSpace(line[65:75]),
	}
	valorStr := strings.TrimSpace(line[75:87])
	valor, err := strconv.ParseFloat(valorStr, 64)
	if err != nil {
		return OrderLine{}, fmt.Errorf("erro ao converter valor do produto: %w", err)
	}
	obj.Price = valor

	dateStr := strings.TrimSpace(line[87:95])
	date, err := time.Parse("20060102", dateStr)
	if err != nil {
		return OrderLine{}, fmt.Errorf("erro ao converter data: %w", err)
	}
	obj.OrderDate = date

	return obj, nil
}
