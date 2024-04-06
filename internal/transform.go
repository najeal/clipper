package pkg

import (
	"github.com/najeal/go-clipper/internal/specs"
	"github.com/najeal/go-clipper/internal/templater"
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
		Commands: generateCommandsFromCommands(in.Commands),
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
