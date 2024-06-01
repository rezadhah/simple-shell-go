package guard_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/rezadhah/shell/internal/command"
	"github.com/rezadhah/shell/internal/errors"
	"github.com/rezadhah/shell/internal/guard"
)

var _ = Describe("Guard", func() {
	var g *guard.Guardian
	BeforeEach(func() {
		g = guard.NewGuard()
	})

	When("Given wrong number of args", func() {
		It("Should return ErrInvalidArgs", func() {
			err := g.Guard(command.COMMAND_CD, "1", "")
			Expect(err).To(MatchError(errors.ErrInvalidArgs))
		})
	})

	When("Given correct number of args", func() {
		It("Should return nil error", func() {
			err := g.Guard(command.COMMAND_CD, ".")
			Expect(err).To(BeNil())
		})
	})
})
