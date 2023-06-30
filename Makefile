# Go parameters
GOBUILD = go build
GOCLEAN = go clean
GOTEST = go test

# Build target
BINARY_NAME = receipt-processor-challenge.out

# Main build task
all: clean build

build:
	$(GOBUILD) -o $(BINARY_NAME) ./main.go

clean:
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

test:
	$(GOTEST) -cover -v ./...

run: clean build
	./$(BINARY_NAME)