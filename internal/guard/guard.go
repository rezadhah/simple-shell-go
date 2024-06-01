package guard

import (
	"github.com/rezadhah/shell/internal/command"
	"github.com/rezadhah/shell/internal/errors"
)

type Guardian struct {
	cmds map[string]int
}

func NewGuard() *Guardian {
	return &Guardian{cmds: command.AllowedArgumentsOnCommand()}
}

func (g *Guardian) Guard(args ...string) error {
	if len(args) == 1 && args[0] == command.COMMAND_EMPTY {
		return nil
	}

	if numArgs, ok := g.cmds[args[0]]; ok && numArgs != len(args) {
		return errors.ErrInvalidArgs
	}

	return nil
}
