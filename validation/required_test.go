package validation_test

import (
	"github.com/eirini-forks/validus"
	"github.com/eirini-forks/validus/validation"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Required", func() {
	It("succeeds for non-empty values", func() {
		Expect(validus.Validate("foo", validation.Required())).To(Succeed())
	})

	It("fails for empty values", func() {
		Expect(validus.Validate("", validation.Required())).NotTo(Succeed())
	})
})
