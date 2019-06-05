// Code generated by Basil. DO NOT EDIT.
package main

import (
	"fmt"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/variable"
)

type MainInterpreter struct{}

// Create creates a new Main block
func (i MainInterpreter) CreateBlock(id basil.ID) basil.Block {
	return &Main{
		id: id,
	}
}

// Params returns with the list of valid parameters
func (i MainInterpreter) Params() map[basil.ID]basil.ParameterDescriptor {
	return nil
}

// Blocks returns with the list of valid blocks
func (i MainInterpreter) Blocks() map[basil.ID]basil.BlockDescriptor {
	return nil
}

// HasForeignID returns true if the block ID is referencing an other block id
func (i MainInterpreter) HasForeignID() bool {
	return false
}

// HasShortFormat returns true if the block can be defined in the short block format
func (i MainInterpreter) ValueParamName() basil.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i MainInterpreter) ParseContext(ctx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *Main
	if b, ok := basil.Block(nilBlock).(basil.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i MainInterpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	case "id":
		return b.(*Main).id
	default:
		panic(fmt.Errorf("unexpected parameter %q in Main", name))
	}
}

func (i MainInterpreter) SetParam(block basil.Block, name basil.ID, value interface{}) error {
	var err error
	b := block.(*Main)
	switch name {
	case "id":
		b.id, err = variable.IdentifierValue(value)
	}
	return err
}

func (i MainInterpreter) SetBlock(block basil.Block, name basil.ID, value interface{}) error {
	return nil
}

func (i MainInterpreter) ProcessChannels(blockContainer basil.BlockContainer) {
}

func (i MainInterpreter) CloseChannels(blockContainer basil.BlockContainer) {
}