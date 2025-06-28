
build:
	go build -o app ./cmd

run:
	go run ./cmd

test:
	go test ./... -v


docker:
	docker run -p 8080:8080 app-go