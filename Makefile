VERSION = latest-$(shell git log --pretty=format:'%h' -n 1)

build:
	go build -o ./bin/visual-expenses -ldflags "-X main.version=$(VERSION)"

.PHONY: test
test:
	go test ./... -v -cover

.PHONY: clean
clean:
	rm -rf bin/*
