// Code generated by Basil. DO NOT EDIT.
package function

import (
	"fmt"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/variable"
	"github.com/opsidian/parsley/parsley"
)

// TitleInterpreter is an AST node interpreter for the Title function
type TitleInterpreter struct{}

// StaticCheck runs a static analysis on the function parameters
func (i TitleInterpreter) StaticCheck(ctx interface{}, node basil.FunctionNode) (string, parsley.Error) {
	if len(node.ArgumentNodes()) != 1 {
		return "", parsley.NewError(node.Pos(), fmt.Errorf("%s expects 1 arguments", node.Name()))
	}

	arguments := node.ArgumentNodes()

	if err := variable.CheckNodeType(arguments[0], "string"); err != nil {
		return "", err
	}

	return "string", nil

}

// Eval returns with the result of the function
func (i TitleInterpreter) Eval(ctx interface{}, node basil.FunctionNode) (interface{}, parsley.Error) {
	arguments := node.ArgumentNodes()

	arg0, err := variable.NodeStringValue(arguments[0], ctx)
	if err != nil {
		return nil, err
	}

	return Title(
		arg0,
	), nil

}
