# 🚚 logistics-go

Aplicação escrita em Go para processar arquivos legados de pedidos e expor os dados normalizados via API REST.

---

## ✨ Funcionalidades

- 🔍 Lê arquivos de texto com pedidos desnormalizados
- 🛒 Agrupa por usuário e por pedido
- 📦 Retorna JSON estruturado com usuários, pedidos e produtos
- 🐳 Deploy via Docker

---

### Exemplo de arquivo de entrada

[Baixe o arquivo input.txt](./examples/input.txt)

---

## 📤 Exemplo de saída (JSON normalizado)
```json
[
  {
    "user_id": 1,
    "name": "Zarelli",
    "orders": [
      {
        "order_id": 123,
        "total": "1024.48",
        "date": "2021-12-01",
        "products": [
          { "product_id": 111, "value": "512.24" },
          { "product_id": 122, "value": "512.24" }
        ]
      }
    ]
  }
]
```
---

## 🛠 Como subir a aplicação usando o Makefile
O projeto inclui um Makefile para facilitar comandos comuns no desenvolvimento.

🔧 Comandos disponíveis
- `make build` ➔ Compila a aplicação e gera o binário
- `make run` ➔ Executa a aplicação localmente (go run ./cmd)
- `make test`	➔ Roda todos os testes com saída verbosa (go test -v)
- `make docker`	➔ Sobe a aplicação com docker, expondo na porta 80:80

---

## 👩‍💻 Desenvolvido por
Nariana Pereira

