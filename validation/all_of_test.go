package validation_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/eirini-forks/validus"
	"github.com/eirini-forks/validus/validation"
)

var _ = Describe("AllOf", func() {
	It("checks for all validations, in order", func() {
		type x struct {
			Foo string
			Bar string
		}

		vs := validation.AllOf(
			validation.Required(),
			validation.Field("Foo", "foo", validation.Required()),
			validation.Field("Bar", "bar", validation.Required()),
		)

		Expect(validus.Validate(x{}, vs)).To(MatchError(SatisfyAll(
			ContainSubstring("required"),
			Not(ContainSubstring("foo")),
			Not(ContainSubstring("bar")),
		)))
		Expect(validus.Validate(x{Foo: "", Bar: "bar"}, vs)).To(MatchError(SatisfyAll(
			ContainSubstring("foo"),
			Not(ContainSubstring("bar")),
		)))
		Expect(validus.Validate(x{Foo: "foo", Bar: ""}, vs)).To(MatchError(SatisfyAll(
			Not(ContainSubstring("foo")),
			ContainSubstring("bar"),
		)))
	})
})
