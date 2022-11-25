package validation_test

import (
	"github.com/eirini-forks/validus"
	"github.com/eirini-forks/validus/validation"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

type j struct {
	Foo string `json:"foo"`
}

var _ = Describe("TaggedField", func() {
	requireFoo := validation.JSONField("Foo", validation.Required())

	It("succeeds if the inner validation succeeds", func() {
		Expect(validus.Validate(j{Foo: "foo"}, requireFoo)).To(Succeed())
	})

	It("fails if the inner validation fails", func() {
		Expect(validus.Validate(j{}, requireFoo)).To(MatchError(SatisfyAll(
			ContainSubstring("foo"),
			ContainSubstring("required"),
		)))
	})
})
