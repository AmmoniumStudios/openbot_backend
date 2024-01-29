BINARY_NAME=openbot_backend.out

build:
	go build -o bin/$(BINARY_NAME) src/main.go
run:
	go build -o bin/$(BINARY_NAME) src/main.go
	./bin/$(BINARY_NAME)
clean:
	rm -rf ./bin

