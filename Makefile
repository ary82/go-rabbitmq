.PHONY: run
run: build
	@echo "running..."
	@./bin/go-rabbitmq

.PHONY: build
build:
	@echo "building..."
	@go build -o ./bin/go-rabbitmq ./main.go

.PHONY: watch
watch:
	@${HOME}/go/bin/air

.PHONY: test
test:
	go test -v ./...

.PHONY: clean
clean:
	@echo "removing bin/ files"
	@rm ./bin/*
	@echo "removing tmp/ files"
	@rm ./tmp/*
