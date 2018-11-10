// Code generated by Basil. DO NOT EDIT.
package function

import (
	"fmt"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/variable"
	"github.com/opsidian/parsley/parsley"
)

// ReplaceInterpreter is an AST node interpreter for the Replace function
type ReplaceInterpreter struct{}

// StaticCheck runs a static analysis on the function parameters
func (i ReplaceInterpreter) StaticCheck(ctx interface{}, node basil.FunctionNode) (string, parsley.Error) {
	if len(node.ArgumentNodes()) != 3 {
		return "", parsley.NewError(node.Pos(), fmt.Errorf("%s expects 3 arguments", node.Name()))
	}

	arguments := node.ArgumentNodes()

	if err := variable.CheckNodeType(arguments[0], "string"); err != nil {
		return "", err
	}

	if err := variable.CheckNodeType(arguments[1], "string"); err != nil {
		return "", err
	}

	if err := variable.CheckNodeType(arguments[2], "string"); err != nil {
		return "", err
	}

	return "string", nil

}

// Eval returns with the result of the function
func (i ReplaceInterpreter) Eval(ctx interface{}, node basil.FunctionNode) (interface{}, parsley.Error) {
	arguments := node.ArgumentNodes()

	arg0, err := variable.NodeStringValue(arguments[0], ctx)
	if err != nil {
		return nil, err
	}

	arg1, err := variable.NodeStringValue(arguments[1], ctx)
	if err != nil {
		return nil, err
	}

	arg2, err := variable.NodeStringValue(arguments[2], ctx)
	if err != nil {
		return nil, err
	}

	return Replace(
		arg0,
		arg1,
		arg2,
	), nil

}
