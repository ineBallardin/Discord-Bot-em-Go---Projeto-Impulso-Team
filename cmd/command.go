package cmd

type Command interface {
	Execute(interaction map[string]interface{}) error
}

type CommandInfo struct {
	Name        string
	Description string
	Options     map[string]interface{}
	Command     Command
}
