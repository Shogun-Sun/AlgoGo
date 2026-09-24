LOCAL_BIN:=$(CURDIR)/.bin/
BENCHSTAT:=$(LOCAL_BIN)benchstat
.PHONY: bench-test run

bin-deps:
	@GOBIN=$(LOCAL_BIN) go install golang.org/x/perf/cmd/benchstat

bench-test: bin-deps
	@go test -bench=. -benchmem -count=6 ./test | $(BENCHSTAT) -

run:
	@go run .

clean:
	rm -rf .bin/