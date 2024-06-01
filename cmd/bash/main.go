package main

import (
	"log"

	"github.com/rezadhah/shell/internal/shell"
)

func main() {
	sh := shell.NewShell()
	if err := sh.InputMode(); err != nil {
		log.Fatalln(err)
	}
}
