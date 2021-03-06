// Code generated by Basil. DO NOT EDIT.

package functions

import (
	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/function"
	"github.com/opsidian/basil/basil/schema"
	"github.com/opsidian/parsley/parsley"
)

// ArrayContainsInterpreter is the basil interpreter for the ArrayContains function
type ArrayContainsInterpreter struct {
	s schema.Schema
}

func (i ArrayContainsInterpreter) Schema() schema.Schema {
	if i.s == nil {
		i.s = &schema.Function{
			Metadata: schema.Metadata{
				Description: "It returns true if the array contains the given element",
			},
			Parameters: schema.Parameters{
				schema.NamedSchema{
					Name: "arr",
					Schema: &schema.Array{
						Items: &schema.Untyped{},
					},
				},
				schema.NamedSchema{
					Name:   "elem",
					Schema: &schema.Untyped{},
				},
			},
			Result: &schema.Boolean{},
		}
	}
	return i.s
}

// Eval returns with the result of the function
func (i ArrayContainsInterpreter) Eval(ctx interface{}, node basil.FunctionNode) (interface{}, parsley.Error) {
	parameters := i.Schema().(*schema.Function).GetParameters()
	arguments := node.ArgumentNodes()

	arg0, evalErr := parsley.EvaluateNode(ctx, arguments[0])
	if evalErr != nil {
		return nil, evalErr
	}
	if err := parameters[0].Schema.ValidateValue(arg0); err != nil {
		return nil, parsley.NewError(arguments[0].Pos(), err)
	}
	var val0 = arg0.([]interface{})

	arg1, evalErr := parsley.EvaluateNode(ctx, arguments[1])
	if evalErr != nil {
		return nil, evalErr
	}
	if err := parameters[1].Schema.ValidateValue(arg1); err != nil {
		return nil, parsley.NewError(arguments[1].Pos(), err)
	}
	var val1 = arg1

	res, resErr := ArrayContains(val0, val1)
	if resErr != nil {
		if funcErr, ok := resErr.(*function.Error); ok {
			return nil, parsley.NewError(arguments[funcErr.ArgIndex].Pos(), funcErr.Err)
		}
		return nil, parsley.NewError(node.Pos(), resErr)
	}

	return res, nil
}
