# `blackieops/synonym`

This is a small Go microservice to handle returning Go module metadata and
redirect/alias the paths to Github.

## Usage

First you need a `config.yaml`. For example:

```yaml
---
hostname: "go.b8s.dev"
target_base_url: "github.com/blackieops"
default_branch_name: "main"
```

We recommend deploying synonym via Docker. A Dockerfile is provided to build
the container image:

```
$ docker build -t synonym .
```

Then you can just run it, forwarding the port and mounting your config file:

```
$ docker run --rm -p 6969:6969 -v `pwd`/config.yaml:/config.yaml synonym
```
