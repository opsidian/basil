// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package directives

import (
	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/schema"
)

// @block
type EvalStage struct {
	// @id
	id basil.ID
	// @value
	// @required
	// @enum ["init", "main", "close"]
	value string
}

func (s *EvalStage) ID() basil.ID {
	return s.id
}

func (s *EvalStage) ApplyToSchema(sch schema.Schema) error {
	return schema.UpdateMetadata(sch, func(meta schema.MetadataAccessor) error {
		meta.SetAnnotation("eval_stage", s.value)
		return nil
	})
}
