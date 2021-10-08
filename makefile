.PHONY: test test-cover

test:
	go test ./...

test-cover:
	go test ./... -coverprofile cover.out
	go tool cover -html=cover.out