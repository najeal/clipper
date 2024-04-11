package templater

type Root struct {
	PackageName string
	Methods     []string
	Version     string
	VersionFlag Flag
	App         App
	Exits       []Exit
}

type App struct {
	Name    string
	Version bool
	Usage   string
	Flags   []Flag
	// Action is the service method to call
	// if empty, calling the command will display the helper
	Action   string
	Commands []Command
}

type Command struct {
	Name  string
	Usage string
	Flags []Flag
	// Action is the service method to call
	// if empty, calling the command will display the helper
	Action      string
	SubCommands []Command
}

type Flag struct {
	Name    string
	Value   string
	Type    string
	Usage   string
	Aliases []string
	EnvVars []string
}

type Exit struct {
	VarName string
	Message string
	Code    int
}
