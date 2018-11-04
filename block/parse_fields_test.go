package block_test

import (
	"go/ast"
	"go/parser"
	"go/token"

	"github.com/opsidian/basil/block"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

const q = "`"

var _ = Describe("ParseFields", func() {

	var source string
	var fields []*block.Field
	var parseErr error

	JustBeforeEach(func() {
		file, err := parser.ParseFile(token.NewFileSet(), "testfile", source, parser.AllErrors)
		Expect(err).ToNot(HaveOccurred())

		str := file.Decls[0].(*ast.GenDecl).Specs[0].(*ast.TypeSpec).Type.(*ast.StructType)
		fields, parseErr = block.ParseFields(str, file)
	})

	Context("when the tags are valid", func() {
		BeforeEach(func() {
			source = `
				package foo
				type Foo struct {
					id string ` + q + `basil:"id"` + q + `
				}`
		})

		It("should return with the parsed fields", func() {
			Expect(parseErr).ToNot(HaveOccurred())
			Expect(fields).To(Equal([]*block.Field{
				{
					Name:      "id",
					ParamName: "id",
					Type:      "string",
					IsID:      true,
					Stage:     "default",
				},
			}))
		})
	})

	Context("when there are fields with valid types", func() {
		BeforeEach(func() {
			source = `
				package foo
				type Foo struct {
					id string ` + q + `basil:"id"` + q + `
					field_string string
					field_integer int64
					field_float float64
					field_bool bool
					field_array []interface{}
					field_map map[string]interface{}
					field_time_duration time.Duration
				}`
		})

		It("should return with the parsed fields", func() {
			Expect(parseErr).ToNot(HaveOccurred())
			Expect(fields).To(Equal([]*block.Field{
				{
					Name:      "id",
					ParamName: "id",
					Type:      "string",
					IsID:      true,
					Stage:     "default",
				},
				{
					Name:      "field_string",
					ParamName: "field_string",
					Type:      "string",
					Stage:     "default",
				},
				{
					Name:      "field_integer",
					ParamName: "field_integer",
					Type:      "int64",
					Stage:     "default",
				},
				{
					Name:      "field_float",
					ParamName: "field_float",
					Type:      "float64",
					Stage:     "default",
				},
				{
					Name:      "field_bool",
					ParamName: "field_bool",
					Type:      "bool",
					Stage:     "default",
				},
				{
					Name:      "field_array",
					ParamName: "field_array",
					Type:      "[]interface{}",
					Stage:     "default",
				},
				{
					Name:      "field_map",
					ParamName: "field_map",
					Type:      "map[string]interface{}",
					Stage:     "default",
				},
				{
					Name:      "field_time_duration",
					ParamName: "field_time_duration",
					Type:      "time.Duration",
					Stage:     "default",
				},
			}))
		})
	})

	Context("when there is no id field", func() {
		BeforeEach(func() {
			source = `
				package foo
				type Foo struct {
					foo string
				}`
		})

		It("should return with error", func() {
			Expect(parseErr).To(MatchError("no fields were found with the \"id\" tag"))
		})
	})

	Context("when there are multiple id fields", func() {
		BeforeEach(func() {
			source = `
				package foo
				type Foo struct {
					id1 string ` + q + `basil:"id"` + q + `
					id2 string ` + q + `basil:"id"` + q + `
				}`
		})

		It("should return with error", func() {
			Expect(parseErr).To(MatchError("multiple id fields were found: id1, id2"))
		})
	})

	Context("when there are multiple value fields", func() {
		BeforeEach(func() {
			source = `
				package foo
				type Foo struct {
					id string ` + q + `basil:"id"` + q + `
					value1 string ` + q + `basil:"value"` + q + `
					value2 string ` + q + `basil:"value"` + q + `
				}`
		})

		It("should return with error", func() {
			Expect(parseErr).To(MatchError("multiple value fields were found: value1, value2"))
		})
	})

	Context("when there are required fields other than the value field", func() {
		BeforeEach(func() {
			source = `
				package foo
				type Foo struct {
					id string ` + q + `basil:"id"` + q + `
					value string ` + q + `basil:"value"` + q + `
					foo string ` + q + `basil:"required"` + q + `
				}`
		})

		It("should return with error", func() {
			Expect(parseErr).To(MatchError("when setting a value field then no other fields can be required"))
		})
	})

	Context("when there is an unknown tag", func() {
		BeforeEach(func() {
			source = `
				package foo
				type Foo struct {
					id string ` + q + `basil:"id"` + q + `
					foo string ` + q + `basil:"nonexisting"` + q + `
				}`
		})

		It("should return with error", func() {
			Expect(parseErr).To(MatchError("invalid tag \"nonexisting\" on field \"foo\""))
		})
	})

})