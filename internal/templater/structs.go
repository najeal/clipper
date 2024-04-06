package templater

type Root struct {
	PackageName string
	Methods     []string
	App         App
}

type App struct {
	Name  string
	Usage string
	// Action is the service method to call
	// if empty, calling the command will display the helper
	Action   string
	Commands []Command
}

type Command struct {
	Name  string
	Usage string
	// Action is the service method to call
	// if empty, calling the command will display the helper
	Action      string
	SubCommands []Command
}
