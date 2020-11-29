build:
	go build -i -o ./build/main -race ./cmd/naivecoin-go/main.go 

clean:
	rm -rf ./build

test:
	go test ./...
	 