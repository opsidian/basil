package fixtures

import (
	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/block"
)

//go:generate basil generate
type BlockWithManyBlock struct {
	IDField     basil.ID       `basil:"id"`
	BlockSimple []*BlockSimple `basil:"block"`
}

func (b *BlockWithManyBlock) ParseContextOverride() basil.ParseContextOverride {
	return basil.ParseContextOverride{
		BlockTransformerRegistry: block.InterpreterRegistry{
			"block_simple": BlockSimpleInterpreter{},
		},
	}
}