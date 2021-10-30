BENCH_RUN ?= .

.PHONY: test test-cover test-race-condition bench

test:
	go test ./...

test-cover:
	go test ./... -coverprofile cover.out
	go tool cover -html=cover.out

test-race-condition:
	go test ./... -cpu=1,9 -race -count=50 -failfast

bench:
	cd benchmark && go test -v -benchmem -run=^$$ -bench=$(BENCH_RUN) ./...

bench-mem:
	cd benchmark && go test -v -memprofile mem.out -benchmem -run=^$$ -bench=$(BENCH_RUN) ./...
	go tool pprof ./benchmark/mem.out