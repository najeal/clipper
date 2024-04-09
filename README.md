# Welcome to Clipper

Clipper works on top of [urfave/cli](https://github.com/urfave/cli).<br>
- You declare cli into a yaml file just like specification.
- Clipper loads the file then initializes the urfave `*cli.App` and declares one interface.
- You implement the generated interface.

# Getting started

## Install Clipper:
```go install github.com/najeal/clipper/cmd/clipper@latest```

## Take a look to the examples
- [greet](examples/greet/README.md)
- [subcommands](examples/subcommands/README.md)
- [flags](examples/flags/README.md)
