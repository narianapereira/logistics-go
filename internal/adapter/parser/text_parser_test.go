package parser

import (
	"encoding/json"
	"testing"
	"time"
)

func TestTextParser_Parse(t *testing.T) {
	parser := NewTextParser()

	input := []byte(`
0000000077                         Mrs. Stephen Trantow00000008440000000005     1288.7720211127
0000000061                           Dimple Bergstrom I00000006710000000004       43.3620211104
0000000077                         Mrs. Stephen Trantow00000008320000000006      961.3720210513
`)

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

	if result[0].UserId == "0000000077" && len(result[0].OrderList) != 2 {
		t.Errorf("esperado 2 pedidos para usuário Stephen, recebido %d", len(result[0].OrderList))
	}

	expectedDate, _ := time.Parse("2006-01-02", "2021-11-27")
	gotDate := result[0].OrderList[0].OrderDate

	if !gotDate.Equal(expectedDate) {
		t.Errorf("data esperada: %v, recebida: %v", expectedDate, gotDate)
	}
}
