// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package directives

import "github.com/opsidian/basil/basil"

//go:generate basil generate
type Todo struct {
	id          basil.ID `basil:"id"`
	description string   `basil:"value,required"`
}

func (t *Todo) ID() basil.ID {
	return t.id
}

func (t *Todo) RuntimeConfig() basil.RuntimeConfig {
	return basil.RuntimeConfig{}
}

func (t *Todo) EvalStage() basil.EvalStage {
	return basil.EvalStageIgnore
}