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
	if numArgs, _ := g.cmds[args[0]]; numArgs != len(args) {
		return errors.ErrInvalidArgs
	}

	return nil
}
