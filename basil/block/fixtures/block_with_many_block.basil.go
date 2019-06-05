// Code generated by Basil. DO NOT EDIT.
package fixtures

import (
	"fmt"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/variable"
)

type BlockWithManyBlockInterpreter struct{}

// Create creates a new BlockWithManyBlock block
func (i BlockWithManyBlockInterpreter) CreateBlock(id basil.ID) basil.Block {
	return &BlockWithManyBlock{
		IDField: id,
	}
}

// Params returns with the list of valid parameters
func (i BlockWithManyBlockInterpreter) Params() map[basil.ID]basil.ParameterDescriptor {
	return nil
}

// Blocks returns with the list of valid blocks
func (i BlockWithManyBlockInterpreter) Blocks() map[basil.ID]basil.BlockDescriptor {
	return map[basil.ID]basil.BlockDescriptor{
		"block_simple": {
			EvalStage:  basil.EvalStages["main"],
			IsRequired: false,
			IsOutput:   false,
			IsMany:     true,
		},
	}
}

// HasForeignID returns true if the block ID is referencing an other block id
func (i BlockWithManyBlockInterpreter) HasForeignID() bool {
	return false
}

// HasShortFormat returns true if the block can be defined in the short block format
func (i BlockWithManyBlockInterpreter) ValueParamName() basil.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i BlockWithManyBlockInterpreter) ParseContext(ctx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *BlockWithManyBlock
	if b, ok := basil.Block(nilBlock).(basil.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i BlockWithManyBlockInterpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	case "id":
		return b.(*BlockWithManyBlock).IDField
	default:
		panic(fmt.Errorf("unexpected parameter %q in BlockWithManyBlock", name))
	}
}

func (i BlockWithManyBlockInterpreter) SetParam(block basil.Block, name basil.ID, value interface{}) error {
	var err error
	b := block.(*BlockWithManyBlock)
	switch name {
	case "id":
		b.IDField, err = variable.IdentifierValue(value)
	}
	return err
}

func (i BlockWithManyBlockInterpreter) SetBlock(block basil.Block, name basil.ID, value interface{}) error {
	b := block.(*BlockWithManyBlock)
	switch name {
	case "block_simple":
		b.BlockSimple = append(b.BlockSimple, value.(*BlockSimple))
	}
	return nil
}

func (i BlockWithManyBlockInterpreter) ProcessChannels(blockContainer basil.BlockContainer) {
}

func (i BlockWithManyBlockInterpreter) CloseChannels(blockContainer basil.BlockContainer) {
}