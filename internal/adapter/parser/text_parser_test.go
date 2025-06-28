package parser

import (
	"encoding/json"
	"testing"
)

func TestTextParser_Parse(t *testing.T) {
	parser := NewTextParser()

	input := []byte(
		`0000000077                         Mrs. Stephen Trantow00000008440000000005     1288.7720211127
0000000061                           Dimple Bergstrom I00000006710000000004       43.3620211104
0000000077                         Mrs. Stephen Trantow00000008320000000006      961.3720210513`)

	output, err := parser.Parse(input)
	if err != nil {
		t.Fatalf("erro ao fazer parse: %v", err)
	}

	var result []UserWithOrders
	if err := json.Unmarshal(output, &result); err != nil {
		t.Fatalf("erro ao fazer unmarshal do resultado: %v", err)
	}

	if len(result) != 2 {
		t.Errorf("esperado 2 usuários, recebido %d", len(result))
	}

	var user77 *UserWithOrders
	for i := range result {
		if result[i].UserId == "77" {
			user77 = &result[i]
			break
		}
	}
	if user77 == nil {
		t.Fatalf("usuário com ID 77 não encontrado no resultado")
	}

	if len(user77.Orders) != 2 {
		t.Errorf("esperado 2 pedidos para usuário 77, recebido %d", len(user77.Orders))
	}

	expectedDate := "2021-11-27"
	gotDate := user77.Orders[0].Date

	if gotDate != expectedDate {
		t.Errorf("data esperada: %v, recebida: %v", expectedDate, gotDate)
	}

	if user77.Orders[0].Total != "1288.77" {
		t.Errorf("total esperado: %v, recebido: %v", "1288.77", user77.Orders[0].Total)
	}

	var user61 *UserWithOrders
	for i := range result {
		if result[i].UserId == "61" {
			user61 = &result[i]
			break
		}
	}
	if user61 == nil {
		t.Fatalf("usuário com ID 61 não encontrado no resultado")
	}

	if len(user61.Orders) != 1 {
		t.Errorf("esperado 1 pedido para usuário 61, recebido %d", len(user61.Orders))
	}

	if user61.Orders[0].Total != "43.36" {
		t.Errorf("total esperado: %v, recebido: %v", "43.36", user61.Orders[0].Total)
	}
}
