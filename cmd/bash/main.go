package main

import (
	"log"

	"github.com/rezadhah/shell/internal/shell"
)

const (
	CD_COMMAND   = "cd"
	EXIT_COMMAND = "exit"
)

func main() {
	sh := shell.NewShell()
	if err := sh.InputMode(); err != nil {
		log.Fatalln(err)
	}
}
