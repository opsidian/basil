// Code generated by Basil. DO NOT EDIT.
package function

import (
	"fmt"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/variable"
	"github.com/opsidian/parsley/parsley"
)

// FloorInterpreter is an AST node interpreter for the Floor function
type FloorInterpreter struct{}

// StaticCheck runs a static analysis on the function parameters
func (i FloorInterpreter) StaticCheck(ctx interface{}, node basil.FunctionNode) (string, parsley.Error) {
	if len(node.ArgumentNodes()) != 1 {
		return "", parsley.NewError(node.Pos(), fmt.Errorf("%s expects 1 arguments", node.Name()))
	}

	arguments := node.ArgumentNodes()

	if err := variable.CheckNodeType(arguments[0], "*variable.Number"); err != nil {
		return "", err
	}

	return "int64", nil

}

// Eval returns with the result of the function
func (i FloorInterpreter) Eval(ctx interface{}, node basil.FunctionNode) (interface{}, parsley.Error) {
	arguments := node.ArgumentNodes()

	arg0, evalErr := arguments[0].Value(ctx)
	if evalErr != nil {
		return nil, evalErr
	}

	val0, convertErr := variable.NumberValue(arg0)
	if convertErr != nil {
		return nil, parsley.NewError(arguments[0].Pos(), convertErr)
	}

	return Floor(
		val0,
	), nil

}
