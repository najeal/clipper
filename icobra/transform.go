package icobra

import (
	"strings"

	"github.com/najeal/clipper/icobra/specs"
	"github.com/najeal/clipper/icobra/templater"
)

func transform(packageName string, in specs.RootApp) templater.DataWrapper {
	varCmd := getVarCmd(in.Name)
	methods := []string{}
	if in.Runnable {
		methods = append(methods, in.Name)
	}
	rootData := templater.RootData{
		Methods: append(methods, getCommandsMethods(in.Commands)...),
		Command: templater.Command{
			Use:         in.Use,
			Short:       in.Short,
			Long:        in.Long,
			CommandName: in.Name,
			VarCmd:      varCmd,
			Run:         in.Runnable,
			Flags:       getFlags(varCmd, in.Flags),
		},
		CommandsHierarchy: getCommandsHierarchies(varCmd, in.Commands),
	}

	return templater.DataWrapper{
		RootData: rootData,
		Commands: getCommands(in.Commands),
	}
}

func getFlags(varCmd string, flags []specs.Flag) []templater.Flag {
	res := []templater.Flag{}
	for _, flag := range flags {
		res = append(res, templater.Flag{
			Name:       flag.Name,
			Usage:      flag.Usage,
			Type:       flag.Type,
			Shorthand:  flag.Shorthand,
			Value:      flag.Value,
			Persistent: flag.Persistent,
			VarCmd:     varCmd,
		})
	}
	return res
}

func getCommands(cmds []specs.Command) []templater.Command {
	res := []templater.Command{}
	for _, cmd := range cmds {
		res = append(res, templater.Command{
			Use:         cmd.Use,
			Short:       cmd.Short,
			Long:        cmd.Long,
			CommandName: cmd.CommandName,
			VarCmd:      getVarCmd(cmd.CommandName),
			Run:         cmd.Runnable,
			Flags:       getFlags(getVarCmd(cmd.CommandName), cmd.Flags),
		})
		res = append(res, getCommands(cmd.SubCommands)...)
	}
	return res
}

func getCommandsHierarchies(parentVarCmd string, cmds []specs.Command) []templater.Hierarchy {
	hierarchies := []templater.Hierarchy{}
	for _, cmd := range cmds {
		hierarchies = append(hierarchies, templater.Hierarchy{
			Parent: parentVarCmd,
			Child:  cmd.CommandName,
		})
		hierarchies = append(hierarchies, getCommandsHierarchies(getVarCmd(cmd.CommandName), cmd.SubCommands)...)
	}
	return hierarchies
}

func getCommandsMethods(cmds []specs.Command) []string {
	res := []string{}
	for _, cmd := range cmds {
		if cmd.Runnable {
			res = append(res, cmd.CommandName)
		}
		res = append(res, getCommandsMethods(cmd.SubCommands)...)
	}
	return res
}

func getVarCmd(name string) string {
	return strings.ToLower(string(name[0])) + name[1:]
}
