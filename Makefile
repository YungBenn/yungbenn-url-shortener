run:
	go run main.go

build:
	go build -o .

test-shortener:
	go test ./internal/shortener_test.go