// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package parser

import (
	"github.com/opsidian/parsley/ast"
	"github.com/opsidian/parsley/combinator"
	"github.com/opsidian/parsley/parser"
	"github.com/opsidian/parsley/parsley"
	"github.com/opsidian/parsley/text"
	"github.com/opsidian/parsley/text/terminal"
)

// And will match a logical and expression defined by the following rule, where P is the input parser:
//   S -> P ("&&" P)*
func And(p parsley.Parser) parser.Func {
	return combinator.Single(
		combinator.SepBy1(
			text.LeftTrim(p, text.WsSpaces),
			text.LeftTrim(terminal.Op("&&"), text.WsSpaces),
		).Bind(ast.InterpreterFunc(evalAnd)),
	)
}

func evalAnd(ctx interface{}, nodes []parsley.Node) (interface{}, parsley.Error) {
	res := true
	errorPos := nodes[0].Pos()
	expectsOp := false
	for _, node := range nodes {
		v, err := node.Value(ctx)
		if err != nil {
			return nil, err
		}
		if expectsOp {
			errorPos = node.Pos()
		} else {
			switch vt := v.(type) {
			case bool:
				res = res && vt
			default:
				return nil, parsley.NewErrorf(errorPos, "unsupported && operation on %T", v)
			}
		}
		expectsOp = !expectsOp
	}
	return res, nil
}