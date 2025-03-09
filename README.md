# `blackieops/synonym`

This is a small Go microservice to handle returning Go module metadata and
redirect/alias the paths to Github.

## Usage

First you need a `config.yaml`. For example:

```yaml
---
port: 8080
hostname: "go.b8s.dev"
target_base_url: "github.com/blackieops"
default_branch_name: "main"
```

We recommend deploying synonym via Docker. Just forward the port and mount your
config file:

```
$ docker run --rm -p 6969:6969 -v `pwd`/config.yaml:/config.yaml ghcr.io/blackieops/synonym:main
```

## Custom Mappings

By default, all requests get rewritten to the configured `target_base_url`, eg.,

```
go.example.com/somepkg -> github.com/example/somepkg
```

But this may not suit all repos. Sometimes you may need some custom overrides
or aliasing to point to different providers, user namespaces, etc.

You can configure `custom_mappings` in the config file to accomplish this:

```yaml
---
# [...]

custom_mappings:
  - path: /example
    target: other.example.com/some/other/path/to/repo
```

## Development

A `Makefile` is included to automate common build and development commands.

To build the binary:

```
$ make build
```

To run the test suite:

```
$ make test
```

To run formatters and linters:

```
$ make fmt
$ make lint
```

To tidy the deps:

```
$ make tidy
```
