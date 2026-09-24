.PHONY: bench-test

bench-test:
	@go test -bench=. -benchmem ./test