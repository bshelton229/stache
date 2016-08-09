dist:
	./dist.sh
.PHONY: dist

test:
	go test ./...
.DEFAULT_GOAL := test
.PHONY: test
