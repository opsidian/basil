package generator

type InterpreterTemplateParams struct {
	Package            string
	Type               string
	Name               string
	Stages             []string
	IDField            *Field
	HasForeignID       bool
	ValueField         *Field
	Params             []*Field
	InputParams        []*Field
	Blocks             []*Field
	ValueFunctionNames map[string]string
}

const interpreterTemplate = `
// Code generated by Basil. DO NOT EDIT.
package {{.Package}}

import (
	"fmt"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/util"
	"github.com/opsidian/basil/basil/variable"
	"github.com/opsidian/parsley/parsley"
)

{{ $root := .}}

type {{.Name}}Interpreter struct {}

// Create creates a new {{.Name}} block
func (i {{.Name}}Interpreter) CreateBlock(id basil.ID) basil.Block {
	return &{{.Name}}{
		{{.IDField.Name}}: id,
	}
}

// Params returns with the list of valid parameters
func (i {{.Name}}Interpreter) Params() map[basil.ID]basil.ParameterDescriptor {
	{{ if .Params -}}
	return map[basil.ID]basil.ParameterDescriptor{
		{{ range .Params -}}
		"{{.ParamName}}": {Type: "{{.Type}}", IsRequired: {{.IsRequired}}, IsOutput: {{.IsOutput}}},
		{{ end -}}
	}
	{{ else -}}
	return nil
	{{ end -}}
}

// HasForeignID returns true if the block ID is referencing an other block id
func (i {{.Name}}Interpreter) HasForeignID() bool {
	return {{.HasForeignID}}
}

// HasShortFormat returns true if the block can be defined in the short block format
func (i {{.Name}}Interpreter) ValueParamName() basil.ID {
	return {{ if .ValueField }}"{{.ValueField.ParamName}}"{{ else }}""{{ end }}
}

// ParseContext returns with the parse context for the block
func (i {{.Name}}Interpreter) ParseContext(ctx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *{{.Name}}
	if b, ok := basil.Block(nilBlock).(basil.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i {{.Name}}Interpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	case "id":
		return b.(*{{$root.Name}}).{{.IDField.Name}}
	{{ range .Params -}}
	case "{{.ParamName}}":
		return b.(*{{$root.Name}}).{{.Name}}
	{{ end -}}
	default:
		panic(fmt.Errorf("unexpected parameter %q in {{.Name}}", name))
	}
}

func (i {{.Name}}Interpreter) SetParam(b basil.Block, name basil.ID, value interface{}) error {
	{{ if .InputParams -}}
	var err error
	switch name {
	{{ range .InputParams -}}
	case "{{.ParamName}}":
		b.(*{{$root.Name}}).{{.Name}}, err = variable.{{index $root.ValueFunctionNames .Type}}(value)
	{{ end -}}
	}
	return err
	{{ else -}}
	return nil
	{{ end -}}
}

func (i {{.Name}}Interpreter) SetBlock(b basil.Block, name basil.ID, value interface{}) error {
	{{ if .Blocks -}}
	switch name {
	{{ range .Blocks -}}
	case "{{.ParamName}}":
		b.(*{{$root.Name}}).{{.Name}} = append(b.(*{{$root.Name}}).{{.Name}}, value.({{trimPrefix .Type "[]" }}))
	{{ end -}}
	}
	{{ end -}}
	return nil
}

`
