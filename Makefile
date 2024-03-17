.PHONY: build
build:
	go build -o server ./cmd/main.go

run:
	./server --config ./config.yaml

clean:
	rm ./server

.DEFAULT_GOAL = build