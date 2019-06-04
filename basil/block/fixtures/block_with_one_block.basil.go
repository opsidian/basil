// Code generated by Basil. DO NOT EDIT.
package fixtures

import (
	"fmt"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/variable"
)

type BlockWithOneBlockInterpreter struct{}

// Create creates a new BlockWithOneBlock block
func (i BlockWithOneBlockInterpreter) CreateBlock(id basil.ID) basil.Block {
	return &BlockWithOneBlock{
		IDField: id,
	}
}

// Params returns with the list of valid parameters
func (i BlockWithOneBlockInterpreter) Params() map[basil.ID]basil.ParameterDescriptor {
	return nil
}

// Blocks returns with the list of valid blocks
func (i BlockWithOneBlockInterpreter) Blocks() map[basil.ID]basil.BlockDescriptor {
	return map[basil.ID]basil.BlockDescriptor{
		"block_simple": {
			Type:       "*BlockSimple",
			EvalStage:  basil.EvalStages["main"],
			IsRequired: false,
			IsOutput:   false,
			IsMany:     false,
		},
	}
}

// HasForeignID returns true if the block ID is referencing an other block id
func (i BlockWithOneBlockInterpreter) HasForeignID() bool {
	return false
}

// HasShortFormat returns true if the block can be defined in the short block format
func (i BlockWithOneBlockInterpreter) ValueParamName() basil.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i BlockWithOneBlockInterpreter) ParseContext(ctx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *BlockWithOneBlock
	if b, ok := basil.Block(nilBlock).(basil.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i BlockWithOneBlockInterpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	case "id":
		return b.(*BlockWithOneBlock).IDField
	default:
		panic(fmt.Errorf("unexpected parameter %q in BlockWithOneBlock", name))
	}
}

func (i BlockWithOneBlockInterpreter) SetParam(block basil.Block, name basil.ID, value interface{}) error {
	var err error
	b := block.(*BlockWithOneBlock)
	switch name {
	case "id":
		b.IDField, err = variable.IdentifierValue(value)
	}
	return err
}

func (i BlockWithOneBlockInterpreter) SetBlock(block basil.Block, name basil.ID, value interface{}) error {
	b := block.(*BlockWithOneBlock)
	switch name {
	case "block_simple":
		b.BlockSimple = value.(*BlockSimple)
	}
	return nil
}

func (i BlockWithOneBlockInterpreter) ProcessChannels(blockContainer basil.BlockContainer) {
}

func (i BlockWithOneBlockInterpreter) CloseChannels(blockContainer basil.BlockContainer) {
}
