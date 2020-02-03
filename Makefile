# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOINSTALL=$(GOCMD) install
BINARY_NAME=main
    
all: install test build
build: 
	$(GOBUILD) ./cmd/alien-invasion/main.go
test: 
	$(GOTEST) -v ./pkg/Util
	$(GOTEST) -v ./pkg/Destroy
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
run:
	./$(BINARY_NAME)
install:
	$(GOINSTALL) ./pkg/Structure
	$(GOINSTALL) ./pkg/Util
	$(GOINSTALL) ./pkg/Destroy
	$(GOINSTALL) ./pkg/Come
	$(GOINSTALL) ./pkg/Move
