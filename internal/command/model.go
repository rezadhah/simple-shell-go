package command

const (
	COMMAND_CD   = "cd"
	COMMAND_EXIT = "exit"
)

func AllowedCommands() []string {
	return []string{COMMAND_CD, COMMAND_EXIT}
}
