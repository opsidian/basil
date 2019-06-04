// Code generated by Basil. DO NOT EDIT.
package main

import (
	"fmt"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/variable"
)

type HelloInterpreter struct{}

// Create creates a new Hello block
func (i HelloInterpreter) CreateBlock(id basil.ID) basil.Block {
	return &Hello{
		id: id,
	}
}

// Params returns with the list of valid parameters
func (i HelloInterpreter) Params() map[basil.ID]basil.ParameterDescriptor {
	return map[basil.ID]basil.ParameterDescriptor{
		"to": {
			Type:       "string",
			EvalStage:  basil.EvalStages["main"],
			IsRequired: true,
			IsOutput:   false,
		},
		"greeting": {
			Type:       "string",
			EvalStage:  basil.EvalStages["close"],
			IsRequired: false,
			IsOutput:   true,
		},
	}
}

// Blocks returns with the list of valid blocks
func (i HelloInterpreter) Blocks() map[basil.ID]basil.BlockDescriptor {
	return nil
}

// HasForeignID returns true if the block ID is referencing an other block id
func (i HelloInterpreter) HasForeignID() bool {
	return false
}

// HasShortFormat returns true if the block can be defined in the short block format
func (i HelloInterpreter) ValueParamName() basil.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i HelloInterpreter) ParseContext(ctx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *Hello
	if b, ok := basil.Block(nilBlock).(basil.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i HelloInterpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	case "id":
		return b.(*Hello).id
	case "to":
		return b.(*Hello).to
	case "greeting":
		return b.(*Hello).greeting
	default:
		panic(fmt.Errorf("unexpected parameter %q in Hello", name))
	}
}

func (i HelloInterpreter) SetParam(block basil.Block, name basil.ID, value interface{}) error {
	var err error
	b := block.(*Hello)
	switch name {
	case "id":
		b.id, err = variable.IdentifierValue(value)
	case "to":
		b.to, err = variable.StringValue(value)
	}
	return err
}

func (i HelloInterpreter) SetBlock(block basil.Block, name basil.ID, value interface{}) error {
	return nil
}

func (i HelloInterpreter) ProcessChannels(blockContainer basil.BlockContainer) {
}

func (i HelloInterpreter) CloseChannels(blockContainer basil.BlockContainer) {
}
