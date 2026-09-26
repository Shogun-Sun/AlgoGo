LOCAL_BIN:=$(CURDIR)/.bin/
BUILD_DIR:=$(CURDIR)/build/
BENCHSTAT:=$(LOCAL_BIN)benchstat

.PHONY: bench-test clean build

$(BENCHSTAT):
	GOBIN=$(LOCAL_BIN) go install golang.org/x/perf/cmd/benchstat

bench-test: $(BENCHSTAT)
	@go test -bench=. -benchmem -count=6 ./test | $(BENCHSTAT) -

bench-bubble: $(BENCHSTAT)
	@go test -bench=BenchmarkBubbleSort -benchmem -count=6 ./test | $(BENCHSTAT) -

bench-stupid: $(BENCHSTAT)
	@go test -bench=BenchmarkStupidSort -benchmem -count=6 ./test | $(BENCHSTAT) -

run:
	@go mod tidy && go run ./cmd/algogo

clean:
	rm -rf $(LOCAL_BIN) $(BUILD_DIR)

build:
	@go build -o $(BUILD_DIR)/algogo ./cmd/algogo
