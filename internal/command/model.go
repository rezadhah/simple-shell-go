package command

const (
	COMMAND_EMPTY = ""
	COMMAND_CD    = "cd"
	COMMAND_EXIT  = "exit"
)

func AllowedCommands() []string {
	return []string{COMMAND_CD, COMMAND_EXIT, COMMAND_EMPTY}
}

func AllowedArgumentsOnCommand() map[string]int {
	return map[string]int{
		COMMAND_EMPTY: 0,
		COMMAND_CD:    2,
		COMMAND_EXIT:  1,
	}
}
