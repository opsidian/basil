// Code generated by Basil. DO NOT EDIT.
package common

import (
	"fmt"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/variable"
)

type ItInterpreter struct{}

// Create creates a new It block
func (i ItInterpreter) CreateBlock(id basil.ID) basil.Block {
	return &It{
		id: id,
	}
}

// Params returns with the list of valid parameters
func (i ItInterpreter) Params() map[basil.ID]basil.ParameterDescriptor {
	return map[basil.ID]basil.ParameterDescriptor{
		"value": {
			Type:       "int64",
			EvalStage:  basil.EvalStages["close"],
			IsRequired: false,
			IsOutput:   true,
		},
	}
}

// Blocks returns with the list of valid blocks
func (i ItInterpreter) Blocks() map[basil.ID]basil.BlockDescriptor {
	return nil
}

// HasForeignID returns true if the block ID is referencing an other block id
func (i ItInterpreter) HasForeignID() bool {
	return false
}

// HasShortFormat returns true if the block can be defined in the short block format
func (i ItInterpreter) ValueParamName() basil.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i ItInterpreter) ParseContext(ctx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *It
	if b, ok := basil.Block(nilBlock).(basil.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i ItInterpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	case "id":
		return b.(*It).id
	case "value":
		return b.(*It).value
	default:
		panic(fmt.Errorf("unexpected parameter %q in It", name))
	}
}

func (i ItInterpreter) SetParam(block basil.Block, name basil.ID, value interface{}) error {
	var err error
	b := block.(*It)
	switch name {
	case "id":
		b.id, err = variable.IdentifierValue(value)
	}
	return err
}

func (i ItInterpreter) SetBlock(block basil.Block, name basil.ID, value interface{}) error {
	return nil
}