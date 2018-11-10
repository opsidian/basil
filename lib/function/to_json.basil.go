// Code generated by Basil. DO NOT EDIT.
package function

import (
	"fmt"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/function"
	"github.com/opsidian/basil/variable"
	"github.com/opsidian/parsley/parsley"
)

// ToJSONInterpreter is an AST node interpreter for the ToJSON function
type ToJSONInterpreter struct{}

// StaticCheck runs a static analysis on the function parameters
func (i ToJSONInterpreter) StaticCheck(ctx interface{}, node basil.FunctionNode) (string, parsley.Error) {
	if len(node.ArgumentNodes()) != 1 {
		return "", parsley.NewError(node.Pos(), fmt.Errorf("%s expects 1 arguments", node.Name()))
	}

	arguments := node.ArgumentNodes()

	if err := variable.CheckNodeType(arguments[0], "interface{}"); err != nil {
		return "", err
	}

	return "string", nil

}

// Eval returns with the result of the function
func (i ToJSONInterpreter) Eval(ctx interface{}, node basil.FunctionNode) (interface{}, parsley.Error) {
	arguments := node.ArgumentNodes()

	arg0, err := variable.NodeAnyValue(arguments[0], ctx)
	if err != nil {
		return nil, err
	}

	res, resErr := ToJSON(arg0)
	if resErr != nil {
		if funcErr, ok := resErr.(*function.Error); ok {
			return nil, parsley.NewError(arguments[funcErr.ArgIndex].Pos(), funcErr.Err)
		}
		return nil, parsley.NewError(node.Pos(), resErr)
	}

	return res, nil

}
