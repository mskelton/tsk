.PHONY: test docs

test:
	go test ./...

docs:
	mdbook serve docs --open
