.PHONY: test build 

test:
	go test -v ./...

build:
	go build  -o ./dist/rayt *.go 

clean:
	rm -rf ./dist 

all: clean build 

	
