.PHONY: help clean build run 
.PHONY: lint test coverage

# set default goal to help
.DEFAULT_GOAL := help

# set default binary output
BINARYOUTPUT = app.out

### * make help							Print this help
help: Makefile
	@sed -n 's/^###//p' $<

### * make clean							Remove all artifacts
clean:
	- rm $(BINARYOUTPUT)

### * make build							Build the application
build:
	- go build -o $(BINARYOUTPUT) 

### * make test							Run unit-tests for the application
test:
	- go test -v

### * make up							Start the application
up: build
	- ./$(BINARYOUTPUT)

### * make run							Clean and build the application
run: clean build up

