build:
	go build -i -o ./build/main ./cmd/naivecoin/main.go -r ace

clean:
	rm -rf ./build

test:
	go test
	 