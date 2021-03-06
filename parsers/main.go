// Copyright (c) 2017 Opsidian Ltd.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package parsers

import (
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/opsidian/parsley/ast"
	"github.com/opsidian/parsley/combinator"
	"github.com/opsidian/parsley/data"
	"github.com/opsidian/parsley/parsley"
	"github.com/opsidian/parsley/text"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/block"
)

// NewMain returns a parser for parsing a main block (a block body)
//   S     -> (PARAM|BLOCK)*
//   ID    -> /[a-z][a-z0-9]*(?:_[a-z0-9]+)*/
//   PARAM -> ID ("="|":=") P
//   VALUE -> EXPRESSION
//         -> ARRAY
//         -> MAP
func NewMain(id basil.ID, interpreter basil.BlockInterpreter) *Main {
	m := &Main{
		id:          id,
		interpreter: interpreter,
	}

	expr := Expression()

	paramOrBlock := combinator.Choice(
		Parameter(expr, true, true),
		Block(expr),
	).Name("parameter or block definition")

	m.p = text.Trim(
		combinator.Seq(
			block.TokenBlockBody,
			func(i int) parsley.Parser {
				if i == 0 {
					return text.LeftTrim(paramOrBlock, text.WsSpacesNl)
				}
				return text.LeftTrim(paramOrBlock, text.WsSpacesForceNl)
			},
			func(int) bool {
				return true
			},
		).Bind(m),
	)

	return m
}

// Main is the main block parser
// It will parse a block body (list of params and blocks)
// and will return with a block with the given id and the type "main"
type Main struct {
	id          basil.ID
	interpreter basil.BlockInterpreter
	p           parsley.Parser
}

// Parse will parse the input into a block
func (m *Main) Parse(ctx *parsley.Context, leftRecCtx data.IntMap, pos parsley.Pos) (parsley.Node, data.IntSet, parsley.Error) {
	return m.p.Parse(ctx, leftRecCtx, pos)
}

// ParseText parses the string input as a main block
func (m *Main) ParseText(ctx *basil.ParseContext, input string) error {
	_, err := basil.ParseText(ctx, m.p, input)
	return err
}

// ParseFile parses the given file as a main block
func (m *Main) ParseFile(ctx *basil.ParseContext, path string) error {
	_, err := basil.ParseFile(ctx, m.p, path)
	return err
}

func (m *Main) ParseDir(ctx *basil.ParseContext, dir string) error {
	info, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("path %q does not exist", dir)
		}
	}
	if !info.IsDir() {
		return fmt.Errorf("path %q is not a directory", dir)
	}

	paths, err := filepath.Glob(path.Join(dir, "*.basil"))
	if err != nil {
		return err
	}

	return m.ParseFiles(ctx, paths...)
}

// ParseFiles parses multiple files as one block
func (m *Main) ParseFiles(ctx *basil.ParseContext, paths ...string) error {
	nodeBuilder := func(nodes []parsley.Node) parsley.Node {
		return ast.NewNonTerminalNode("BLOCK_BODY", nodes, m)
	}
	return basil.ParseFiles(ctx, m.p, nodeBuilder, paths)
}

// Eval will panic as it should not be called on a raw block node
func (m *Main) Eval(userCtx interface{}, node parsley.NonTerminalNode) (interface{}, parsley.Error) {
	panic("Eval should not be called on a raw block node")
}

// TransformNode will transform the parsley node into a basil block node
func (m *Main) TransformNode(userCtx interface{}, node parsley.Node) (parsley.Node, parsley.Error) {
	return block.TransformMainNode(userCtx, node, m.id, m.interpreter)
}
