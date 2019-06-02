// Code generated by counterfeiter. DO NOT EDIT.
package basilfakes

import (
	"sync"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/parsley/parsley"
)

type FakeFunctionInterpreter struct {
	EvalStub        func(interface{}, basil.FunctionNode) (interface{}, parsley.Error)
	evalMutex       sync.RWMutex
	evalArgsForCall []struct {
		arg1 interface{}
		arg2 basil.FunctionNode
	}
	evalReturns struct {
		result1 interface{}
		result2 parsley.Error
	}
	evalReturnsOnCall map[int]struct {
		result1 interface{}
		result2 parsley.Error
	}
	StaticCheckStub        func(interface{}, basil.FunctionNode) (string, parsley.Error)
	staticCheckMutex       sync.RWMutex
	staticCheckArgsForCall []struct {
		arg1 interface{}
		arg2 basil.FunctionNode
	}
	staticCheckReturns struct {
		result1 string
		result2 parsley.Error
	}
	staticCheckReturnsOnCall map[int]struct {
		result1 string
		result2 parsley.Error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeFunctionInterpreter) Eval(arg1 interface{}, arg2 basil.FunctionNode) (interface{}, parsley.Error) {
	fake.evalMutex.Lock()
	ret, specificReturn := fake.evalReturnsOnCall[len(fake.evalArgsForCall)]
	fake.evalArgsForCall = append(fake.evalArgsForCall, struct {
		arg1 interface{}
		arg2 basil.FunctionNode
	}{arg1, arg2})
	fake.recordInvocation("Eval", []interface{}{arg1, arg2})
	fake.evalMutex.Unlock()
	if fake.EvalStub != nil {
		return fake.EvalStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.evalReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFunctionInterpreter) EvalCallCount() int {
	fake.evalMutex.RLock()
	defer fake.evalMutex.RUnlock()
	return len(fake.evalArgsForCall)
}

func (fake *FakeFunctionInterpreter) EvalCalls(stub func(interface{}, basil.FunctionNode) (interface{}, parsley.Error)) {
	fake.evalMutex.Lock()
	defer fake.evalMutex.Unlock()
	fake.EvalStub = stub
}

func (fake *FakeFunctionInterpreter) EvalArgsForCall(i int) (interface{}, basil.FunctionNode) {
	fake.evalMutex.RLock()
	defer fake.evalMutex.RUnlock()
	argsForCall := fake.evalArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeFunctionInterpreter) EvalReturns(result1 interface{}, result2 parsley.Error) {
	fake.evalMutex.Lock()
	defer fake.evalMutex.Unlock()
	fake.EvalStub = nil
	fake.evalReturns = struct {
		result1 interface{}
		result2 parsley.Error
	}{result1, result2}
}

func (fake *FakeFunctionInterpreter) EvalReturnsOnCall(i int, result1 interface{}, result2 parsley.Error) {
	fake.evalMutex.Lock()
	defer fake.evalMutex.Unlock()
	fake.EvalStub = nil
	if fake.evalReturnsOnCall == nil {
		fake.evalReturnsOnCall = make(map[int]struct {
			result1 interface{}
			result2 parsley.Error
		})
	}
	fake.evalReturnsOnCall[i] = struct {
		result1 interface{}
		result2 parsley.Error
	}{result1, result2}
}

func (fake *FakeFunctionInterpreter) StaticCheck(arg1 interface{}, arg2 basil.FunctionNode) (string, parsley.Error) {
	fake.staticCheckMutex.Lock()
	ret, specificReturn := fake.staticCheckReturnsOnCall[len(fake.staticCheckArgsForCall)]
	fake.staticCheckArgsForCall = append(fake.staticCheckArgsForCall, struct {
		arg1 interface{}
		arg2 basil.FunctionNode
	}{arg1, arg2})
	fake.recordInvocation("StaticCheck", []interface{}{arg1, arg2})
	fake.staticCheckMutex.Unlock()
	if fake.StaticCheckStub != nil {
		return fake.StaticCheckStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.staticCheckReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeFunctionInterpreter) StaticCheckCallCount() int {
	fake.staticCheckMutex.RLock()
	defer fake.staticCheckMutex.RUnlock()
	return len(fake.staticCheckArgsForCall)
}

func (fake *FakeFunctionInterpreter) StaticCheckCalls(stub func(interface{}, basil.FunctionNode) (string, parsley.Error)) {
	fake.staticCheckMutex.Lock()
	defer fake.staticCheckMutex.Unlock()
	fake.StaticCheckStub = stub
}

func (fake *FakeFunctionInterpreter) StaticCheckArgsForCall(i int) (interface{}, basil.FunctionNode) {
	fake.staticCheckMutex.RLock()
	defer fake.staticCheckMutex.RUnlock()
	argsForCall := fake.staticCheckArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeFunctionInterpreter) StaticCheckReturns(result1 string, result2 parsley.Error) {
	fake.staticCheckMutex.Lock()
	defer fake.staticCheckMutex.Unlock()
	fake.StaticCheckStub = nil
	fake.staticCheckReturns = struct {
		result1 string
		result2 parsley.Error
	}{result1, result2}
}

func (fake *FakeFunctionInterpreter) StaticCheckReturnsOnCall(i int, result1 string, result2 parsley.Error) {
	fake.staticCheckMutex.Lock()
	defer fake.staticCheckMutex.Unlock()
	fake.StaticCheckStub = nil
	if fake.staticCheckReturnsOnCall == nil {
		fake.staticCheckReturnsOnCall = make(map[int]struct {
			result1 string
			result2 parsley.Error
		})
	}
	fake.staticCheckReturnsOnCall[i] = struct {
		result1 string
		result2 parsley.Error
	}{result1, result2}
}

func (fake *FakeFunctionInterpreter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.evalMutex.RLock()
	defer fake.evalMutex.RUnlock()
	fake.staticCheckMutex.RLock()
	defer fake.staticCheckMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeFunctionInterpreter) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ basil.FunctionInterpreter = new(FakeFunctionInterpreter)