package pkg

import (
	"github.com/najeal/clipper/internal/specs"
	"github.com/najeal/clipper/internal/templater"
)

func transform(packageName string, in specs.Root) templater.Root {
	return templater.Root{
		PackageName: packageName,
		Methods:     generateMethodsFromRoot(in),
		App:         generateAppFromRoot(in),
	}
}

func generateAppFromRoot(in specs.Root) templater.App {
	return templater.App{
		Name:     in.Name,
		Usage:    in.Usage,
		Action:   in.Action,
		Flags:    generateFlagsFromFlags(in.Flags),
		Commands: generateCommandsFromCommands(in.Commands),
	}
}

func generateFlagsFromFlags(flags []specs.Flag) []templater.Flag {
	tflags := make([]templater.Flag, 0, len(flags))
	for _, flag := range flags {
		tflags = append(tflags, generateFlagFromFlag(flag))
	}
	return tflags
}

func generateFlagFromFlag(flag specs.Flag) templater.Flag {
	return templater.Flag{
		Name:    flag.Name,
		Type:    flag.Type,
		Value:   flag.Value,
		Usage:   flag.Usage,
		Aliases: flag.Aliases,
		EnvVars: flag.EnvVars,
	}
}

func generateCommandsFromCommands(commands []specs.Command) []templater.Command {
	tcommands := make([]templater.Command, 0, len(commands))
	for _, command := range commands {
		tcommands = append(tcommands, generateCommandFromCommand(command))
	}
	return tcommands
}

func generateCommandFromCommand(command specs.Command) templater.Command {
	return templater.Command{
		Name:        command.Name,
		Usage:       command.Usage,
		Action:      command.Action,
		Flags:       generateFlagsFromFlags(command.Flags),
		SubCommands: generateCommandsFromCommands(command.SubCommands),
	}
}

func generateMethodsFromRoot(in specs.Root) []string {
	methods := generateMethodsFromCommands(in.Commands)
	if in.Action != "" {
		methods = append([]string{in.Action}, methods...)
	}
	return methods
}

func generateMethodsFromCommands(commands []specs.Command) []string {
	methods := make([]string, 0, len(commands))
	for _, command := range commands {
		if command.Action != "" {
			methods = append(methods, command.Action)
		}
		methods = append(methods, generateMethodsFromCommands(command.SubCommands)...)
	}
	return methods
}
