package fixtures

import "github.com/opsidian/basil/basil"

//go:generate basil generate
type BlockWithReference struct {
	IDField basil.ID `basil:"id,reference"`
}
