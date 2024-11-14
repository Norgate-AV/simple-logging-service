TARGET=simple-logging-service

build:
	@go build -o bin/simple-logging-service

run: build
	@./bin/simple-logging-service

install: build
	@go install

clean:
	@rm -rf bin