package templater

type MainData struct {
	ProjectRoot string
}

type RootData struct {
	Methods           []string
	Command           Command
	CommandsHierarchy []Hierarchy
	Flags             []Flag
}

type DataWrapper struct {
	RootData
	Commands []Command
}

type Flag struct {
	Name       string
	Usage      string
	Type       string
	Shorthand  string
	Value      string
	Persistent bool
	VarCmd     string
}

type Command struct {
	CommandName string
	VarCmd      string
	Flags       []Flag

	Use   string
	Short string
	Long  string
	Run   bool
}

type Hierarchy struct {
	Parent string
	Child  string
}
