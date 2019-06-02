package generator_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/opsidian/basil/basil/block/generator"
	"github.com/opsidian/basil/basil/variable"
)

var _ = Describe("Field", func() {
	var f *generator.Field

	BeforeEach(func() {
		f = &generator.Field{
			Name:        "foo",
			ParamName:   "param_foo",
			Type:        "string",
			IsRequired:  false,
			Stage:       "default",
			IsID:        false,
			IsValue:     false,
			IsReference: false,
			IsBlock:     false,
			IsNode:      false,
		}
	})

	Describe("Validate", func() {

		It("returns no error for a valid field", func() {
			Expect(f.Validate()).To(BeNil())
		})

		It("allows the reference tag on an id field", func() {
			f.IsID = true
			f.IsReference = true
			f.Type = variable.TypeIdentifier
			Expect(f.Validate()).To(BeNil())
		})

		It("allows an array type with block", func() {
			f.IsBlock = true
			f.Type = "[]SomeBlock"
			Expect(f.Validate()).To(BeNil())
		})

		It("allows an array type with node", func() {
			f.IsNode = true
			f.Type = "[]SomeNode"
			Expect(f.Validate()).To(BeNil())
		})

		It("returns error if id and value are both set", func() {
			f.IsID = true
			f.IsValue = true
			Expect(f.Validate()).To(MatchError("field \"foo\" must only have one tag of: id, value, block or node"))
		})

		It("returns error if block and node are both set", func() {
			f.IsBlock = true
			f.IsNode = true
			Expect(f.Validate()).To(MatchError("field \"foo\" must only have one tag of: id, value, block or node"))
		})

		It("returns error if reference is on a non-id field", func() {
			f.IsReference = true
			Expect(f.Validate()).To(MatchError("the \"reference\" tag can only be set on the id field"))
		})

		It("returns error if param name is invalid", func() {
			f.ParamName = "invalid name"
			Expect(f.Validate()).To(MatchError("\"name\" tag is invalid on field \"foo\", it must be a valid identifier"))
		})

		It("returns error if block is not on an array field", func() {
			f.IsBlock = true
			Expect(f.Validate()).To(MatchError("field \"foo\" must be an array"))
		})

		It("returns error if node is not on an array field", func() {
			f.IsNode = true
			Expect(f.Validate()).To(MatchError("field \"foo\" must be an array"))
		})

		It("returns an error for an invalid type", func() {
			f.Type = "invalidtype"
			Expect(f.Validate()).To(MatchError("invalid field type on field \"foo\", use valid type or use ignore tag"))
		})

		It("returns an error for an empty stage", func() {
			f.Stage = ""
			Expect(f.Validate()).To(MatchError("\"stage\" can not be empty on field \"foo\""))
		})

		It("returns an error for a non-string id field", func() {
			f.IsID = true
			f.Type = "int64"
			Expect(f.Validate()).To(MatchError("field \"foo\" must be defined as basil.ID"))
		})

	})
})