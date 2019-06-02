// Code generated by Basil. DO NOT EDIT.
package function

import (
	"fmt"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/variable"
	"github.com/opsidian/parsley/parsley"
)

// TrimPrefixInterpreter is an AST node interpreter for the TrimPrefix function
type TrimPrefixInterpreter struct{}

// StaticCheck runs a static analysis on the function parameters
func (i TrimPrefixInterpreter) StaticCheck(ctx interface{}, node basil.FunctionNode) (string, parsley.Error) {
	if len(node.ArgumentNodes()) != 2 {
		return "", parsley.NewError(node.Pos(), fmt.Errorf("%s expects 2 arguments", node.Name()))
	}

	arguments := node.ArgumentNodes()

	if err := variable.CheckNodeType(arguments[0], "string"); err != nil {
		return "", err
	}

	if err := variable.CheckNodeType(arguments[1], "string"); err != nil {
		return "", err
	}

	return "string", nil

}

// Eval returns with the result of the function
func (i TrimPrefixInterpreter) Eval(ctx interface{}, node basil.FunctionNode) (interface{}, parsley.Error) {
	arguments := node.ArgumentNodes()

	arg0, evalErr := arguments[0].Value(ctx)
	if evalErr != nil {
		return nil, evalErr
	}

	val0, convertErr := variable.StringValue(arg0)
	if convertErr != nil {
		return nil, parsley.NewError(arguments[0].Pos(), convertErr)
	}

	arg1, evalErr := arguments[1].Value(ctx)
	if evalErr != nil {
		return nil, evalErr
	}

	val1, convertErr := variable.StringValue(arg1)
	if convertErr != nil {
		return nil, parsley.NewError(arguments[1].Pos(), convertErr)
	}

	return TrimPrefix(
		val0,
		val1,
	), nil

}
