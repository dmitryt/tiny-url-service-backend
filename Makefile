# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GODEV=air
GOLINT=golangci-lint run
GOGET=$(GOCMD) get
BINARY_NAME=tiny-url-service
BINARY_DIR=tmp

lint:
	$(GOLINT) ./...

clean:
	$(GOCLEAN)
	rm -r $(BINARY_DIR)

dev:
	$(GODEV)

run:
	mkdir -p $(BINARY_DIR) && $(GOBUILD) -o $(BINARY_DIR) ./...
	./$(BINARY_DIR)/$(BINARY_NAME)

build:
	mkdir -p $(BINARY_DIR) && $(GOBUILD) -o $(BINARY_DIR) ./...
