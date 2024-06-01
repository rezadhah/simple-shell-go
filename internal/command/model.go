package command

const (
	COMMAND_CD   = "cd"
	COMMAND_EXIT = "exit"
)

func AllowedCommands() []string {
	return []string{COMMAND_CD, COMMAND_EXIT}
}

func AllowedArgumentsOnCommand() map[string]int {
	return map[string]int{
		COMMAND_CD:   2,
		COMMAND_EXIT: 1,
	}
}
