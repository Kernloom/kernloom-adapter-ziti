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

## Release

Release pipelines must run protocol compatibility checks, contract tests, unit tests, lab integration tests, packaging and signing.

## Dependencies

The adapter uses Go 1.26.4 and imports `github.com/kernloom/kernloom-protocol`. Local development uses a `replace` directive to the sibling protocol repo.

## Related Repos

Forge and KLIQ live in `kernloom-core`. Protocol definitions live in `kernloom-protocol`.
