.PHONY: test test-cover test-race-condition

test:
	go test ./...

test-cover:
	go test ./... -coverprofile cover.out
	go tool cover -html=cover.out

test-race-condition:
	go test ./... -cpu=1,9 -race -count=50 -failfast