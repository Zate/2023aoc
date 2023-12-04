# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=aoc

.PHONY: test clean tidy run day

test: 
	@$(GOTEST) -v ./...
clean: 
	@$(GOCMD) clean
	@rm -f $(BINARY_NAME)
tidy:
	@$(GOCMD) mod tidy
run:
	@$(GOBUILD) -o $(BINARY_NAME) -v .
	@./$(BINARY_NAME)

day%: tidy
	@$(GOTEST) -v ./cmd/root.go ./cmd/day$*.go ./cmd/day$*_test.go
	@$(GOBUILD) -o $(BINARY_NAME) -v .
	@./$(BINARY_NAME) day$*