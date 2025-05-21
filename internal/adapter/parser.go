package adapter

import (
	"bufio"
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type TextParser struct {
}

type Order struct {
	UserId    string
	Name      string
	OrderId   string
	ProductId string
	Price     float64
	OrderDate time.Time
}

type UserOrderList struct {
	UserId    string
	Name      string
	OrderList []Order
}

func NewTextParser() *TextParser {
	return &TextParser{}
}

func (t *TextParser) Parse(content []byte) (interface{}, error) {
	scanner := bufio.NewScanner(bytes.NewReader(content))
	var userOrdersMap = make(map[string][]Order)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		result, err := parseLine(line)
		if err != nil {
			return nil, err
		}
		addOrderToUser(result, userOrdersMap)
	}

	processUserOrders(userOrdersMap)
	return nil, nil
}

func processUserOrders(ordersMap map[string][]Order) UserOrderList {
	//continuar processamento das orders para transformar em JSON

	return UserOrderList{}
}

func addOrderToUser(result Order, userOrdersMap map[string][]Order) {
	if orders, exists := userOrdersMap[result.UserId]; exists {
		userOrdersMap[result.UserId] = append(orders, result)
	} else {
		userOrdersMap[result.UserId] = []Order{result}
	}
}

func parseLine(line string) (Order, error) {
	obj := Order{
		UserId:    strings.TrimSpace(line[0:10]),
		Name:      strings.TrimSpace(line[10:55]),
		OrderId:   strings.TrimSpace(line[55:65]),
		ProductId: strings.TrimSpace(line[65:75]),
	}
	valorStr := strings.TrimSpace(line[75:87])
	valor, err := strconv.ParseFloat(valorStr, 64)
	if err != nil {
		return Order{}, fmt.Errorf("erro ao converter valor do produto: %w", err)
	}
	obj.Price = valor

	dateStr := strings.TrimSpace(line[87:95])
	date, err := time.Parse("20060102", dateStr)
	if err != nil {
		return Order{}, fmt.Errorf("erro ao converter data: %w", err)
	}
	obj.OrderDate = date

	return obj, nil
}
