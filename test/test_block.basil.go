// Code generated by Basil. DO NOT EDIT.
package test

import (
	"fmt"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/variable"
)

type TestBlockInterpreter struct{}

// Create creates a new TestBlock block
func (i TestBlockInterpreter) CreateBlock(id basil.ID) basil.Block {
	return &TestBlock{
		IDField: id,
	}
}

// Params returns with the list of valid parameters
func (i TestBlockInterpreter) Params() map[basil.ID]basil.ParameterDescriptor {
	return map[basil.ID]basil.ParameterDescriptor{
		"value": {
			Type:       "interface{}",
			EvalStage:  basil.EvalStages["main"],
			IsRequired: false,
			IsOutput:   false,
		},
		"field_string": {
			Type:       "string",
			EvalStage:  basil.EvalStages["main"],
			IsRequired: false,
			IsOutput:   false,
		},
		"field_int": {
			Type:       "int64",
			EvalStage:  basil.EvalStages["main"],
			IsRequired: false,
			IsOutput:   false,
		},
		"field_float": {
			Type:       "float64",
			EvalStage:  basil.EvalStages["main"],
			IsRequired: false,
			IsOutput:   false,
		},
		"field_bool": {
			Type:       "bool",
			EvalStage:  basil.EvalStages["main"],
			IsRequired: false,
			IsOutput:   false,
		},
		"field_array": {
			Type:       "[]interface{}",
			EvalStage:  basil.EvalStages["main"],
			IsRequired: false,
			IsOutput:   false,
		},
		"field_map": {
			Type:       "map[string]interface{}",
			EvalStage:  basil.EvalStages["main"],
			IsRequired: false,
			IsOutput:   false,
		},
		"field_time_duration": {
			Type:       "time.Duration",
			EvalStage:  basil.EvalStages["main"],
			IsRequired: false,
			IsOutput:   false,
		},
		"custom_field": {
			Type:       "string",
			EvalStage:  basil.EvalStages["main"],
			IsRequired: false,
			IsOutput:   false,
		},
	}
}

// Blocks returns with the list of valid blocks
func (i TestBlockInterpreter) Blocks() map[basil.ID]basil.BlockDescriptor {
	return map[basil.ID]basil.BlockDescriptor{
		"testblock": {
			Type:       "*TestBlock",
			EvalStage:  basil.EvalStages["main"],
			IsRequired: false,
			IsOutput:   false,
			IsMany:     true,
		},
	}
}

// HasForeignID returns true if the block ID is referencing an other block id
func (i TestBlockInterpreter) HasForeignID() bool {
	return false
}

// HasShortFormat returns true if the block can be defined in the short block format
func (i TestBlockInterpreter) ValueParamName() basil.ID {
	return "value"
}

// ParseContext returns with the parse context for the block
func (i TestBlockInterpreter) ParseContext(ctx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *TestBlock
	if b, ok := basil.Block(nilBlock).(basil.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i TestBlockInterpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	case "id":
		return b.(*TestBlock).IDField
	case "value":
		return b.(*TestBlock).Value
	case "field_string":
		return b.(*TestBlock).FieldString
	case "field_int":
		return b.(*TestBlock).FieldInt
	case "field_float":
		return b.(*TestBlock).FieldFloat
	case "field_bool":
		return b.(*TestBlock).FieldBool
	case "field_array":
		return b.(*TestBlock).FieldArray
	case "field_map":
		return b.(*TestBlock).FieldMap
	case "field_time_duration":
		return b.(*TestBlock).FieldTimeDuration
	case "custom_field":
		return b.(*TestBlock).FieldCustomName
	default:
		panic(fmt.Errorf("unexpected parameter %q in TestBlock", name))
	}
}

func (i TestBlockInterpreter) SetParam(block basil.Block, name basil.ID, value interface{}) error {
	var err error
	b := block.(*TestBlock)
	switch name {
	case "id":
		b.IDField, err = variable.IdentifierValue(value)
	case "value":
		b.Value, err = variable.AnyValue(value)
	case "field_string":
		b.FieldString, err = variable.StringValue(value)
	case "field_int":
		b.FieldInt, err = variable.IntegerValue(value)
	case "field_float":
		b.FieldFloat, err = variable.FloatValue(value)
	case "field_bool":
		b.FieldBool, err = variable.BoolValue(value)
	case "field_array":
		b.FieldArray, err = variable.ArrayValue(value)
	case "field_map":
		b.FieldMap, err = variable.MapValue(value)
	case "field_time_duration":
		b.FieldTimeDuration, err = variable.TimeDurationValue(value)
	case "custom_field":
		b.FieldCustomName, err = variable.StringValue(value)
	}
	return err
}

func (i TestBlockInterpreter) SetBlock(block basil.Block, name basil.ID, value interface{}) error {
	b := block.(*TestBlock)
	switch name {
	case "testblock":
		b.TestBlock = append(b.TestBlock, value.(*TestBlock))
	}
	return nil
}

func (i TestBlockInterpreter) ProcessChannels(blockContainer basil.BlockContainer) {
}

func (i TestBlockInterpreter) CloseChannels(blockContainer basil.BlockContainer) {
}
