package templater

type MainData struct {
	ProjectRoot string
}

type RootData struct {
	Methods           []string
	Command           Command
	CommandsHierarchy []Hierarchy
}

type Command struct {
	CommandName string
	VarCmd      string

	Use   string
	Short string
	Long  string
	Run   bool
}

type Hierarchy struct {
	Parent string
	Child  string
}
