package validation_test

import (
	"github.com/eirini-forks/validus"
	"github.com/eirini-forks/validus/validation"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type x struct {
	Foo string
}

type y struct {
	Bar x
}

var _ = Describe("Field", func() {
	requireFoo := validation.Field("Foo", "foo", validation.Required())

	It("succeeds if the inner validation succeeds", func() {
		Expect(validus.Validate(x{Foo: "foo"}, requireFoo)).To(Succeed())
	})

	It("fails if the inner validation fails", func() {
		Expect(validus.Validate(x{}, requireFoo)).To(MatchError(SatisfyAll(
			ContainSubstring("foo"),
			ContainSubstring("required"),
		)))
	})
})
