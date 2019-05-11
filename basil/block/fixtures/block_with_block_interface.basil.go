// Code generated by Basil. DO NOT EDIT.
package fixtures

import (
	"fmt"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/parsley/parsley"
)

type BlockWithBlockInterfaceInterpreter struct{}

// Create creates a new BlockWithBlockInterface block
func (i BlockWithBlockInterfaceInterpreter) Create(ctx *basil.EvalContext, node basil.BlockNode) basil.Block {
	return &BlockWithBlockInterface{
		IDField: node.ID(),
	}
}

// Params returns with the list of valid parameters
func (i BlockWithBlockInterfaceInterpreter) Params() map[basil.ID]basil.ParameterDescriptor {
	return nil
}

// HasForeignID returns true if the block ID is referencing an other block id
func (i BlockWithBlockInterfaceInterpreter) HasForeignID() bool {
	return false
}

// HasShortFormat returns true if the block can be defined in the short block format
func (i BlockWithBlockInterfaceInterpreter) ValueParamName() basil.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i BlockWithBlockInterfaceInterpreter) ParseContext(parentCtx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *BlockWithBlockInterface
	if b, ok := basil.Block(nilBlock).(basil.ParseContextAware); ok {
		return b.ParseContext(parentCtx)
	}

	return parentCtx
}

func (i BlockWithBlockInterfaceInterpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	case "id":
		return b.(*BlockWithBlockInterface).IDField
	default:
		panic(fmt.Errorf("unexpected parameter %q in BlockWithBlockInterface", name))
	}
}

func (i BlockWithBlockInterfaceInterpreter) SetParam(ctx *basil.EvalContext, b basil.Block, name basil.ID, node basil.ParameterNode) parsley.Error {
	return nil
}

func (i BlockWithBlockInterfaceInterpreter) SetBlock(ctx *basil.EvalContext, b basil.Block, name basil.ID, value interface{}) parsley.Error {
	switch name {
	case "block_simple":
		b.(*BlockWithBlockInterface).Blocks = append(b.(*BlockWithBlockInterface).Blocks, value.(BlockInterface))
	}
	return nil
}
