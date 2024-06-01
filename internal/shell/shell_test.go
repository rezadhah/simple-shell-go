package shell_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/rezadhah/shell/internal/command"
	"github.com/rezadhah/shell/internal/shell"
)

var _ = Describe("Shell", func() {
	When("Constructed", func() {
		It("should has registered allowed commands", func() {
			sh := *shell.NewShell()
			Expect(len(sh.GetAllowedCommands())).To(Equal(len(command.AllowedCommands())))
		})
	})
})
