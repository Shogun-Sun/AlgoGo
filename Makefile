LOCAL_BIN:=$(CURDIR)/.bin/
BENCHSTAT:=$(LOCAL_BIN)benchstat

.PHONY: bench-test run

$(BENCHSTAT):
	GOBIN=$(LOCAL_BIN) go install golang.org/x/perf/cmd/benchstat

bench-test: $(BENCHSTAT)
	@go test -bench=. -benchmem -count=6 ./test | $(BENCHSTAT) -

run:
	@go run .

clean:
	rm -rf $(LOCAL_BIN)
