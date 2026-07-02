# kernloom-adapter-ziti

`kernloom-adapter-ziti` is the OpenZiti adapter repository. It implements the Kernloom adapter protocol out of process and must not directly apply production configuration.

## Build

```sh
make build
```

## Test

```sh
make test
```

## Run

```sh
./bin/kernloom-adapter-ziti serve --addr 127.0.0.1:18081
```

For local inspection, running the binary without arguments prints the adapter descriptor as JSON.

## Release

Release pipelines must run protocol compatibility checks, contract tests, unit tests, lab integration tests, packaging and signing.

## Dependencies

The adapter uses `go 1.26` with `toolchain go1.26.4` and imports `github.com/kernloom/kernloom-protocol`. Local development uses a `replace` directive to the sibling protocol repo.

## Related Repos

Forge and KLIQ live in `kernloom-core`. Protocol definitions live in `kernloom-protocol`.
