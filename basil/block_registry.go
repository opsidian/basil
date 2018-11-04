package basil

import "github.com/opsidian/parsley/parsley"

// BlockRegistry is an interface for retrieving and creating block factories
type BlockRegistry interface {
	BlockTypeExists(ID) bool
	CreateBlockFactory(
		ctx interface{},
		typeNode parsley.Node,
		idNode parsley.Node,
		paramNodes map[ID]parsley.Node,
		blockNodes []parsley.Node,
	) (BlockFactory, parsley.Error)
}

// BlockRegistryAware defines a function to access a block registry
type BlockRegistryAware interface {
	GetBlockRegistry() BlockRegistry
}
