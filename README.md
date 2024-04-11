# Welcome to Clipper

Clipper works on top of [urfave/cli](https://github.com/urfave/cli).<br>
- You declare cli into a yaml file just like specification.
- Clipper loads the file then initializes the urfave `*cli.App` and declares one interface.
- You implement the generated interface.

```
name: greet
version: v0.1.0
versionFlag:
  name: version-print
  usage: printing the version
  aliases:
   - p
usage: fight the loneliness!
usageText: greet
action: Greet
commands:
  - name: check
    usage: check flag value
    action: Check
    flags:
      - name: force
        type: bool
        usage: flag checked by commmand
exitCodes:
  missingForce:
    code: 4
    message: force flag is not set

```

```
func main() {
	app := gen.NewApp(&Service{})
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

type Service struct{}

func (*Service) Greet(ctx *cli.Context) error {
	fmt.Fprintln(os.Stdout, "greeting !")
	return nil
}

func (*Service) Check(ctx *cli.Context) error {
	if !ctx.Bool("force") {
		return gen.ExitMissingForce
	}
	fmt.Fprintln(os.Stdout, "force flag detected")
	return nil
}
```

# Getting started

## Install Clipper:
```go install github.com/najeal/clipper/cmd/clipper@latest```

## Take a look to the examples
- [greet](examples/greet/README.md)
- [subcommands](examples/subcommands/README.md)
- [flags](examples/flags/README.md)
- [exitcodes](examples/exitcodes/README.md)
- [version](examples/version/README.md)
