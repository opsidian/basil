// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package blocks

import (
	"bufio"
	"fmt"
	"io"

	"github.com/opsidian/basil/basil"
)

// Println will write a string followed by a new line to the standard output
// @block
type Println struct {
	// @id
	id basil.ID
	// @value
	// @required
	value interface{}
}

func (p *Println) ID() basil.ID {
	return p.id
}

func (p *Println) Main(ctx basil.BlockContext) error {
	switch v := p.value.(type) {
	case io.ReadCloser:
		scanner := bufio.NewScanner(v)
		for scanner.Scan() {
			if _, err := fmt.Fprintln(ctx.Stdout(), scanner.Text()); err != nil {
				return err
			}
		}
		return nil
	default:
		if _, err := fmt.Fprintln(ctx.Stdout(), p.value); err != nil {
			return err
		}
	}
	return nil
}
