package command_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/rezadhah/shell/internal/command"
	"github.com/rezadhah/shell/internal/errors"
)

var _ = Describe("Test CD Command", Label("CD"), func() {
	var cdCommand command.Command

	BeforeEach(func() {
		cdCommand = command.NewCDCommand()
	})

	When("No argument passed to command", func() {
		It("should return error path required", func() {
			err := cdCommand.Execute("cd")
			Expect(err).To(MatchError(errors.ErrNoPath))
		})

		It("should run with no error if path was passed into arguments", func() {
			err := cdCommand.Execute("cd", "..")
			Expect(err).To(BeNil())
		})
	})
})
