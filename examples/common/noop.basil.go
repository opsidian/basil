// Code generated by Basil. DO NOT EDIT.
package common

import (
	"fmt"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/variable"
)

type NoopInterpreter struct{}

// Create creates a new Noop block
func (i NoopInterpreter) CreateBlock(id basil.ID) basil.Block {
	return &Noop{
		id: id,
	}
}

// Params returns with the list of valid parameters
func (i NoopInterpreter) Params() map[basil.ID]basil.ParameterDescriptor {
	return nil
}

// Blocks returns with the list of valid blocks
func (i NoopInterpreter) Blocks() map[basil.ID]basil.BlockDescriptor {
	return nil
}

// HasForeignID returns true if the block ID is referencing an other block id
func (i NoopInterpreter) HasForeignID() bool {
	return false
}

// HasShortFormat returns true if the block can be defined in the short block format
func (i NoopInterpreter) ValueParamName() basil.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i NoopInterpreter) ParseContext(ctx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *Noop
	if b, ok := basil.Block(nilBlock).(basil.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i NoopInterpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	case "id":
		return b.(*Noop).id
	default:
		panic(fmt.Errorf("unexpected parameter %q in Noop", name))
	}
}

func (i NoopInterpreter) SetParam(block basil.Block, name basil.ID, value interface{}) error {
	var err error
	b := block.(*Noop)
	switch name {
	case "id":
		b.id, err = variable.IdentifierValue(value)
	}
	return err
}

func (i NoopInterpreter) SetBlock(block basil.Block, name basil.ID, value interface{}) error {
	return nil
}
