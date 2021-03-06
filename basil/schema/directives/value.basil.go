// Code generated by Basil. DO NOT EDIT.

package directives

import (
	"fmt"
	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/schema"
)

// ValueInterpreter is the basil interpreter for the Value block
type ValueInterpreter struct {
	s schema.Schema
}

func (i ValueInterpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = &schema.Object{
			Name: "Value",
			Properties: map[string]schema.Schema{
				"id": &schema.String{
					Metadata: schema.Metadata{
						Annotations: map[string]string{"id": "true"},
						ReadOnly:    true,
					},
					Format: "basil.ID",
				},
			},
		}
	}
	return i.s
}

// Create creates a new Value block
func (i ValueInterpreter) CreateBlock(id basil.ID) basil.Block {
	return &Value{
		id: id,
	}
}

// ValueParamName returns the name of the parameter marked as value field, if there is one set
func (i ValueInterpreter) ValueParamName() basil.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i ValueInterpreter) ParseContext(ctx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *Value
	if b, ok := basil.Block(nilBlock).(basil.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i ValueInterpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	case "id":
		return b.(*Value).id
	default:
		panic(fmt.Errorf("unexpected parameter %q in Value", name))
	}
}

func (i ValueInterpreter) SetParam(block basil.Block, name basil.ID, value interface{}) error {
	return nil
}

func (i ValueInterpreter) SetBlock(block basil.Block, name basil.ID, value interface{}) error {
	return nil
}
