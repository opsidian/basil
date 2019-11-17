// Code generated by Basil. DO NOT EDIT.
package main

import (
	"fmt"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/basil/basil/variable"
	"github.com/opsidian/basil/blocks"
)

type ExecInterpreter struct{}

// Create creates a new Exec block
func (i ExecInterpreter) CreateBlock(id basil.ID) basil.Block {
	return &Exec{
		id: id,
	}
}

// Params returns with the list of valid parameters
func (i ExecInterpreter) Params() map[basil.ID]basil.ParameterDescriptor {
	return map[basil.ID]basil.ParameterDescriptor{
		"cmd": {
			Type:       "string",
			EvalStage:  basil.EvalStages["main"],
			IsRequired: true,
			IsOutput:   false,
		},
		"params": {
			Type:       "[]string",
			EvalStage:  basil.EvalStages["main"],
			IsRequired: false,
			IsOutput:   false,
		},
		"dir": {
			Type:       "string",
			EvalStage:  basil.EvalStages["main"],
			IsRequired: false,
			IsOutput:   false,
		},
		"env": {
			Type:       "[]string",
			EvalStage:  basil.EvalStages["main"],
			IsRequired: false,
			IsOutput:   false,
		},
		"exit_code": {
			Type:       "int64",
			EvalStage:  basil.EvalStages["close"],
			IsRequired: false,
			IsOutput:   true,
		},
	}
}

// Blocks returns with the list of valid blocks
func (i ExecInterpreter) Blocks() map[basil.ID]basil.BlockDescriptor {
	return map[basil.ID]basil.BlockDescriptor{
		"stdout": {
			EvalStage:   basil.EvalStages["init"],
			IsGenerated: true,
			IsRequired:  true,
			IsMany:      false,
		},
		"stderr": {
			EvalStage:   basil.EvalStages["init"],
			IsGenerated: true,
			IsRequired:  true,
			IsMany:      false,
		},
	}
}

// HasForeignID returns true if the block ID is referencing an other block id
func (i ExecInterpreter) HasForeignID() bool {
	return false
}

// HasShortFormat returns true if the block can be defined in the short block format
func (i ExecInterpreter) ValueParamName() basil.ID {
	return ""
}

// ParseContext returns with the parse context for the block
func (i ExecInterpreter) ParseContext(ctx *basil.ParseContext) *basil.ParseContext {
	var nilBlock *Exec
	if b, ok := basil.Block(nilBlock).(basil.ParseContextOverrider); ok {
		return ctx.New(b.ParseContextOverride())
	}

	return ctx
}

func (i ExecInterpreter) Param(b basil.Block, name basil.ID) interface{} {
	switch name {
	case "id":
		return b.(*Exec).id
	case "cmd":
		return b.(*Exec).cmd
	case "params":
		return b.(*Exec).params
	case "dir":
		return b.(*Exec).dir
	case "env":
		return b.(*Exec).env
	case "exit_code":
		return b.(*Exec).exitCode
	default:
		panic(fmt.Errorf("unexpected parameter %q in Exec", name))
	}
}

func (i ExecInterpreter) SetParam(block basil.Block, name basil.ID, value interface{}) error {
	var err error
	b := block.(*Exec)
	switch name {
	case "id":
		b.id, err = variable.IdentifierValue(value)
	case "cmd":
		b.cmd, err = variable.StringValue(value)
	case "params":
		b.params, err = variable.StringArrayValue(value)
	case "dir":
		b.dir, err = variable.StringValue(value)
	case "env":
		b.env, err = variable.StringArrayValue(value)
	}
	return err
}

func (i ExecInterpreter) SetBlock(block basil.Block, name basil.ID, value interface{}) error {
	b := block.(*Exec)
	switch name {
	case "stdout":
		b.stdout = value.(*blocks.Stream)
	case "stderr":
		b.stderr = value.(*blocks.Stream)
	}
	return nil
}

func (i ExecInterpreter) EvalStage() basil.EvalStage {
	var nilBlock *Exec
	if b, ok := basil.Block(nilBlock).(basil.EvalStageAware); ok {
		return b.EvalStage()
	}

	return basil.EvalStageUndefined
}
