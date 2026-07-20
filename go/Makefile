# The Dafny-extracted Chat reducer is large enough that the Go compiler's inliner
# OOMs on it (`compile: signal: killed`) under default settings, even on a 128 GB
# machine. Serial compilation with inlining disabled builds it reliably.
GOFLAGS_VERIFIED := -p=1 -gcflags=-l

.PHONY: build test clean
build:
	go build $(GOFLAGS_VERIFIED) ./...
test:
	go test $(GOFLAGS_VERIFIED) ./...
clean:
	go clean -cache
