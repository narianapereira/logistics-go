package parser

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
	"time"
)

type TextParser struct {
	logger *slog.Logger
}

type Product struct {
	ProductId string `json:"product_id"`
	Value     string `json:"value"`
}

type Order struct {
	OrderId  string    `json:"order_id"`
	Total    string    `json:"total"`
	Date     string    `json:"date"`
	Products []Product `json:"products"`
}

type UserWithOrders struct {
	UserId string  `json:"user_id"`
	Name   string  `json:"name"`
	Orders []Order `json:"orders"`
}

type orderLineRaw struct {
	UserId    string
	Name      string
	OrderId   string
	ProductId string
	Value     float64
	Date      time.Time
}

func NewTextParser(logger *slog.Logger) *TextParser {
	return &TextParser{
		logger: logger,
	}
}

func (t *TextParser) Parse(content []byte) ([]byte, error) {
	scanner := bufio.NewScanner(bytes.NewReader(content))

	var lines []orderLineRaw
	t.logger.Info("Stating document scanning")
	for scanner.Scan() {
		line := scanner.Text()
		if strings.TrimSpace(line) == "" {
			continue
		}

		orderLine, err := parseLine(line)
		if err != nil {
			t.logger.Error("Error parsing line")
			return nil, err
		}
		lines = append(lines, orderLine)
	}

	usersMap := make(map[string]*UserWithOrders)

	for _, l := range lines {
		user, exists := usersMap[l.UserId]
		if !exists {
			user = &UserWithOrders{
				UserId: l.UserId,
				Name:   l.Name,
				Orders: []Order{},
			}
			usersMap[l.UserId] = user
		}

		var order *Order
		for i := range user.Orders {
			if user.Orders[i].OrderId == l.OrderId {
				order = &user.Orders[i]
				break
			}
		}

		if order == nil {
			newOrder := Order{
				OrderId:  l.OrderId,
				Date:     l.Date.Format("2006-01-02"),
				Products: []Product{},
				Total:    "0.00",
			}
			user.Orders = append(user.Orders, newOrder)
			order = &user.Orders[len(user.Orders)-1]
		}

		order.Products = append(order.Products, Product{
			ProductId: l.ProductId,
			Value:     fmt.Sprintf("%.2f", l.Value),
		})

		totalFloat, _ := strconv.ParseFloat(order.Total, 64)
		totalFloat += l.Value
		order.Total = fmt.Sprintf("%.2f", totalFloat)
	}

	var result []UserWithOrders
	for _, u := range usersMap {
		result = append(result, *u)
	}

	jsonResult, err := json.MarshalIndent(result, "", "  ")
	if err != nil {
		t.logger.Error("Error converting text to JSON")
		return nil, fmt.Errorf("erro ao converter text para JSON: %w", err)
	}

	return jsonResult, nil
}

func parseLine(line string) (orderLineRaw, error) {
	if len(line) < 95 {
		return orderLineRaw{}, fmt.Errorf("linha com tamanho inválido")
	}

	userId := strings.TrimLeft(line[0:10], "0")
	name := strings.TrimSpace(line[10:55])

	orderId := strings.TrimLeft(line[55:65], "0")
	productId := strings.TrimLeft(line[65:75], "0")

	valueStr := strings.TrimSpace(line[75:87])
	value, err := strconv.ParseFloat(valueStr, 64)
	if err != nil {
		return orderLineRaw{}, fmt.Errorf("erro ao converter valor do produto: %w", err)
	}

	dateStr := strings.TrimSpace(line[87:95])
	date, err := time.Parse("20060102", dateStr)
	if err != nil {
		return orderLineRaw{}, fmt.Errorf("erro ao converter data: %w", err)
	}

	return orderLineRaw{
		UserId:    userId,
		Name:      name,
		OrderId:   orderId,
		ProductId: productId,
		Value:     value,
		Date:      date,
	}, nil
}
