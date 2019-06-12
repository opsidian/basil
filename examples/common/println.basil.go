// Code generated by Basil. DO NOT EDIT.
package common

import (
	"fmt"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/variable"
)

type PrintlnInterpreter struct{}

// Create creates a new Println block
func (i PrintlnInterpreter) CreateBlock(id basil.ID) basil.Block {
	return &Println{
		id: id,
	}
}

// Params returns with the list of valid parameters
func (i PrintlnInterpreter) Params() map[basil.ID]basil.ParameterDescriptor {
	return map[basil.ID]basil.ParameterDescriptor{
		"value": {
			Type:       "interface{}",
			EvalStage:  basil.EvalStages["main"],
			IsRequired: true,
			IsOutput:   false,
		},
	}
}

// Blocks returns with the list of valid blocks
func (i PrintlnInterpreter) Blocks() map[basil.ID]basil.BlockDescriptor {
	return nil
}

// HasForeignID returns true if the block ID is referencing an other block id
func (i PrintlnInterpreter) HasForeignID() bool {
	return false
}

// HasShortFormat returns true if the block can be defined in the short block format
func (i PrintlnInterpreter) ValueParamName() basil.ID {
	return "value"
}

// ParseContext returns with the parse context for the block
func (i PrintlnInterpreter) ParseContext(ctx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *Println
	if b, ok := basil.Block(nilBlock).(basil.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i PrintlnInterpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	case "id":
		return b.(*Println).id
	case "value":
		return b.(*Println).value
	default:
		panic(fmt.Errorf("unexpected parameter %q in Println", name))
	}
}

func (i PrintlnInterpreter) SetParam(block basil.Block, name basil.ID, value interface{}) error {
	var err error
	b := block.(*Println)
	switch name {
	case "id":
		b.id, err = variable.IdentifierValue(value)
	case "value":
		b.value, err = variable.AnyValue(value)
	}
	return err
}

func (i PrintlnInterpreter) SetBlock(block basil.Block, name basil.ID, value interface{}) error {
	return nil
}
