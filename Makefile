# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GORUN=$(GOCMD) run
GOTEST=$(GOCMD) test
GOCLEAN=$(GOCMD) clean
GOGET=$(GOCMD) get
BINARY_NAME=open-user

all: build

build: 
	$(GOBUILD) -o $(BINARY_NAME) -v .

run: 
	$(GORUN) .

test: 
	$(GOTEST) -v ./...

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)

deps: 
	$(GOGET) -v ./...
