# Building

```sh
make build   # or: go build -p=1 -gcflags=-l ./...
make test    # or: go test  -p=1 -gcflags=-l ./...
```

## Why the flags are required

The extracted `Chat` reducer is a single large generated package. Under the Go
compiler's default settings the inliner exhausts memory on it and the build dies
with `compile: signal: killed` — reproducibly, and independently of available RAM
(observed on a 128 GB machine). Compiling serially with inlining disabled
(`-p=1 -gcflags=-l`) builds and tests clean.

This affects compilation only. The built artifact is unaffected.
