// Code generated by counterfeiter. DO NOT EDIT.
package basilfakes

import (
	"sync"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/parsley/parsley"
)

type FakeBlockNode struct {
	BlockTypeStub        func() basil.ID
	blockTypeMutex       sync.RWMutex
	blockTypeArgsForCall []struct {
	}
	blockTypeReturns struct {
		result1 basil.ID
	}
	blockTypeReturnsOnCall map[int]struct {
		result1 basil.ID
	}
	ChildrenStub        func() []basil.Node
	childrenMutex       sync.RWMutex
	childrenArgsForCall []struct {
	}
	childrenReturns struct {
		result1 []basil.Node
	}
	childrenReturnsOnCall map[int]struct {
		result1 []basil.Node
	}
	DependenciesStub        func() basil.Dependencies
	dependenciesMutex       sync.RWMutex
	dependenciesArgsForCall []struct {
	}
	dependenciesReturns struct {
		result1 basil.Dependencies
	}
	dependenciesReturnsOnCall map[int]struct {
		result1 basil.Dependencies
	}
	EvalStageStub        func() basil.EvalStage
	evalStageMutex       sync.RWMutex
	evalStageArgsForCall []struct {
	}
	evalStageReturns struct {
		result1 basil.EvalStage
	}
	evalStageReturnsOnCall map[int]struct {
		result1 basil.EvalStage
	}
	IDStub        func() basil.ID
	iDMutex       sync.RWMutex
	iDArgsForCall []struct {
	}
	iDReturns struct {
		result1 basil.ID
	}
	iDReturnsOnCall map[int]struct {
		result1 basil.ID
	}
	InterpreterStub        func() basil.BlockInterpreter
	interpreterMutex       sync.RWMutex
	interpreterArgsForCall []struct {
	}
	interpreterReturns struct {
		result1 basil.BlockInterpreter
	}
	interpreterReturnsOnCall map[int]struct {
		result1 basil.BlockInterpreter
	}
	ParamTypeStub        func(basil.ID) (string, bool)
	paramTypeMutex       sync.RWMutex
	paramTypeArgsForCall []struct {
		arg1 basil.ID
	}
	paramTypeReturns struct {
		result1 string
		result2 bool
	}
	paramTypeReturnsOnCall map[int]struct {
		result1 string
		result2 bool
	}
	PosStub        func() parsley.Pos
	posMutex       sync.RWMutex
	posArgsForCall []struct {
	}
	posReturns struct {
		result1 parsley.Pos
	}
	posReturnsOnCall map[int]struct {
		result1 parsley.Pos
	}
	ProvidesStub        func() []basil.ID
	providesMutex       sync.RWMutex
	providesArgsForCall []struct {
	}
	providesReturns struct {
		result1 []basil.ID
	}
	providesReturnsOnCall map[int]struct {
		result1 []basil.ID
	}
	ReaderPosStub        func() parsley.Pos
	readerPosMutex       sync.RWMutex
	readerPosArgsForCall []struct {
	}
	readerPosReturns struct {
		result1 parsley.Pos
	}
	readerPosReturnsOnCall map[int]struct {
		result1 parsley.Pos
	}
	TokenStub        func() string
	tokenMutex       sync.RWMutex
	tokenArgsForCall []struct {
	}
	tokenReturns struct {
		result1 string
	}
	tokenReturnsOnCall map[int]struct {
		result1 string
	}
	TypeStub        func() string
	typeMutex       sync.RWMutex
	typeArgsForCall []struct {
	}
	typeReturns struct {
		result1 string
	}
	typeReturnsOnCall map[int]struct {
		result1 string
	}
	ValueStub        func(interface{}) (interface{}, parsley.Error)
	valueMutex       sync.RWMutex
	valueArgsForCall []struct {
		arg1 interface{}
	}
	valueReturns struct {
		result1 interface{}
		result2 parsley.Error
	}
	valueReturnsOnCall map[int]struct {
		result1 interface{}
		result2 parsley.Error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeBlockNode) BlockType() basil.ID {
	fake.blockTypeMutex.Lock()
	ret, specificReturn := fake.blockTypeReturnsOnCall[len(fake.blockTypeArgsForCall)]
	fake.blockTypeArgsForCall = append(fake.blockTypeArgsForCall, struct {
	}{})
	fake.recordInvocation("BlockType", []interface{}{})
	fake.blockTypeMutex.Unlock()
	if fake.BlockTypeStub != nil {
		return fake.BlockTypeStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.blockTypeReturns
	return fakeReturns.result1
}

func (fake *FakeBlockNode) BlockTypeCallCount() int {
	fake.blockTypeMutex.RLock()
	defer fake.blockTypeMutex.RUnlock()
	return len(fake.blockTypeArgsForCall)
}

func (fake *FakeBlockNode) BlockTypeCalls(stub func() basil.ID) {
	fake.blockTypeMutex.Lock()
	defer fake.blockTypeMutex.Unlock()
	fake.BlockTypeStub = stub
}

func (fake *FakeBlockNode) BlockTypeReturns(result1 basil.ID) {
	fake.blockTypeMutex.Lock()
	defer fake.blockTypeMutex.Unlock()
	fake.BlockTypeStub = nil
	fake.blockTypeReturns = struct {
		result1 basil.ID
	}{result1}
}

func (fake *FakeBlockNode) BlockTypeReturnsOnCall(i int, result1 basil.ID) {
	fake.blockTypeMutex.Lock()
	defer fake.blockTypeMutex.Unlock()
	fake.BlockTypeStub = nil
	if fake.blockTypeReturnsOnCall == nil {
		fake.blockTypeReturnsOnCall = make(map[int]struct {
			result1 basil.ID
		})
	}
	fake.blockTypeReturnsOnCall[i] = struct {
		result1 basil.ID
	}{result1}
}

func (fake *FakeBlockNode) Children() []basil.Node {
	fake.childrenMutex.Lock()
	ret, specificReturn := fake.childrenReturnsOnCall[len(fake.childrenArgsForCall)]
	fake.childrenArgsForCall = append(fake.childrenArgsForCall, struct {
	}{})
	fake.recordInvocation("Children", []interface{}{})
	fake.childrenMutex.Unlock()
	if fake.ChildrenStub != nil {
		return fake.ChildrenStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.childrenReturns
	return fakeReturns.result1
}

func (fake *FakeBlockNode) ChildrenCallCount() int {
	fake.childrenMutex.RLock()
	defer fake.childrenMutex.RUnlock()
	return len(fake.childrenArgsForCall)
}

func (fake *FakeBlockNode) ChildrenCalls(stub func() []basil.Node) {
	fake.childrenMutex.Lock()
	defer fake.childrenMutex.Unlock()
	fake.ChildrenStub = stub
}

func (fake *FakeBlockNode) ChildrenReturns(result1 []basil.Node) {
	fake.childrenMutex.Lock()
	defer fake.childrenMutex.Unlock()
	fake.ChildrenStub = nil
	fake.childrenReturns = struct {
		result1 []basil.Node
	}{result1}
}

func (fake *FakeBlockNode) ChildrenReturnsOnCall(i int, result1 []basil.Node) {
	fake.childrenMutex.Lock()
	defer fake.childrenMutex.Unlock()
	fake.ChildrenStub = nil
	if fake.childrenReturnsOnCall == nil {
		fake.childrenReturnsOnCall = make(map[int]struct {
			result1 []basil.Node
		})
	}
	fake.childrenReturnsOnCall[i] = struct {
		result1 []basil.Node
	}{result1}
}

func (fake *FakeBlockNode) Dependencies() basil.Dependencies {
	fake.dependenciesMutex.Lock()
	ret, specificReturn := fake.dependenciesReturnsOnCall[len(fake.dependenciesArgsForCall)]
	fake.dependenciesArgsForCall = append(fake.dependenciesArgsForCall, struct {
	}{})
	fake.recordInvocation("Dependencies", []interface{}{})
	fake.dependenciesMutex.Unlock()
	if fake.DependenciesStub != nil {
		return fake.DependenciesStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.dependenciesReturns
	return fakeReturns.result1
}

func (fake *FakeBlockNode) DependenciesCallCount() int {
	fake.dependenciesMutex.RLock()
	defer fake.dependenciesMutex.RUnlock()
	return len(fake.dependenciesArgsForCall)
}

func (fake *FakeBlockNode) DependenciesCalls(stub func() basil.Dependencies) {
	fake.dependenciesMutex.Lock()
	defer fake.dependenciesMutex.Unlock()
	fake.DependenciesStub = stub
}

func (fake *FakeBlockNode) DependenciesReturns(result1 basil.Dependencies) {
	fake.dependenciesMutex.Lock()
	defer fake.dependenciesMutex.Unlock()
	fake.DependenciesStub = nil
	fake.dependenciesReturns = struct {
		result1 basil.Dependencies
	}{result1}
}

func (fake *FakeBlockNode) DependenciesReturnsOnCall(i int, result1 basil.Dependencies) {
	fake.dependenciesMutex.Lock()
	defer fake.dependenciesMutex.Unlock()
	fake.DependenciesStub = nil
	if fake.dependenciesReturnsOnCall == nil {
		fake.dependenciesReturnsOnCall = make(map[int]struct {
			result1 basil.Dependencies
		})
	}
	fake.dependenciesReturnsOnCall[i] = struct {
		result1 basil.Dependencies
	}{result1}
}

func (fake *FakeBlockNode) EvalStage() basil.EvalStage {
	fake.evalStageMutex.Lock()
	ret, specificReturn := fake.evalStageReturnsOnCall[len(fake.evalStageArgsForCall)]
	fake.evalStageArgsForCall = append(fake.evalStageArgsForCall, struct {
	}{})
	fake.recordInvocation("EvalStage", []interface{}{})
	fake.evalStageMutex.Unlock()
	if fake.EvalStageStub != nil {
		return fake.EvalStageStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.evalStageReturns
	return fakeReturns.result1
}

func (fake *FakeBlockNode) EvalStageCallCount() int {
	fake.evalStageMutex.RLock()
	defer fake.evalStageMutex.RUnlock()
	return len(fake.evalStageArgsForCall)
}

func (fake *FakeBlockNode) EvalStageCalls(stub func() basil.EvalStage) {
	fake.evalStageMutex.Lock()
	defer fake.evalStageMutex.Unlock()
	fake.EvalStageStub = stub
}

func (fake *FakeBlockNode) EvalStageReturns(result1 basil.EvalStage) {
	fake.evalStageMutex.Lock()
	defer fake.evalStageMutex.Unlock()
	fake.EvalStageStub = nil
	fake.evalStageReturns = struct {
		result1 basil.EvalStage
	}{result1}
}

func (fake *FakeBlockNode) EvalStageReturnsOnCall(i int, result1 basil.EvalStage) {
	fake.evalStageMutex.Lock()
	defer fake.evalStageMutex.Unlock()
	fake.EvalStageStub = nil
	if fake.evalStageReturnsOnCall == nil {
		fake.evalStageReturnsOnCall = make(map[int]struct {
			result1 basil.EvalStage
		})
	}
	fake.evalStageReturnsOnCall[i] = struct {
		result1 basil.EvalStage
	}{result1}
}

func (fake *FakeBlockNode) ID() basil.ID {
	fake.iDMutex.Lock()
	ret, specificReturn := fake.iDReturnsOnCall[len(fake.iDArgsForCall)]
	fake.iDArgsForCall = append(fake.iDArgsForCall, struct {
	}{})
	fake.recordInvocation("ID", []interface{}{})
	fake.iDMutex.Unlock()
	if fake.IDStub != nil {
		return fake.IDStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.iDReturns
	return fakeReturns.result1
}

func (fake *FakeBlockNode) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeBlockNode) IDCalls(stub func() basil.ID) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = stub
}

func (fake *FakeBlockNode) IDReturns(result1 basil.ID) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 basil.ID
	}{result1}
}

func (fake *FakeBlockNode) IDReturnsOnCall(i int, result1 basil.ID) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	if fake.iDReturnsOnCall == nil {
		fake.iDReturnsOnCall = make(map[int]struct {
			result1 basil.ID
		})
	}
	fake.iDReturnsOnCall[i] = struct {
		result1 basil.ID
	}{result1}
}

func (fake *FakeBlockNode) Interpreter() basil.BlockInterpreter {
	fake.interpreterMutex.Lock()
	ret, specificReturn := fake.interpreterReturnsOnCall[len(fake.interpreterArgsForCall)]
	fake.interpreterArgsForCall = append(fake.interpreterArgsForCall, struct {
	}{})
	fake.recordInvocation("Interpreter", []interface{}{})
	fake.interpreterMutex.Unlock()
	if fake.InterpreterStub != nil {
		return fake.InterpreterStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.interpreterReturns
	return fakeReturns.result1
}

func (fake *FakeBlockNode) InterpreterCallCount() int {
	fake.interpreterMutex.RLock()
	defer fake.interpreterMutex.RUnlock()
	return len(fake.interpreterArgsForCall)
}

func (fake *FakeBlockNode) InterpreterCalls(stub func() basil.BlockInterpreter) {
	fake.interpreterMutex.Lock()
	defer fake.interpreterMutex.Unlock()
	fake.InterpreterStub = stub
}

func (fake *FakeBlockNode) InterpreterReturns(result1 basil.BlockInterpreter) {
	fake.interpreterMutex.Lock()
	defer fake.interpreterMutex.Unlock()
	fake.InterpreterStub = nil
	fake.interpreterReturns = struct {
		result1 basil.BlockInterpreter
	}{result1}
}

func (fake *FakeBlockNode) InterpreterReturnsOnCall(i int, result1 basil.BlockInterpreter) {
	fake.interpreterMutex.Lock()
	defer fake.interpreterMutex.Unlock()
	fake.InterpreterStub = nil
	if fake.interpreterReturnsOnCall == nil {
		fake.interpreterReturnsOnCall = make(map[int]struct {
			result1 basil.BlockInterpreter
		})
	}
	fake.interpreterReturnsOnCall[i] = struct {
		result1 basil.BlockInterpreter
	}{result1}
}

func (fake *FakeBlockNode) ParamType(arg1 basil.ID) (string, bool) {
	fake.paramTypeMutex.Lock()
	ret, specificReturn := fake.paramTypeReturnsOnCall[len(fake.paramTypeArgsForCall)]
	fake.paramTypeArgsForCall = append(fake.paramTypeArgsForCall, struct {
		arg1 basil.ID
	}{arg1})
	fake.recordInvocation("ParamType", []interface{}{arg1})
	fake.paramTypeMutex.Unlock()
	if fake.ParamTypeStub != nil {
		return fake.ParamTypeStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.paramTypeReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBlockNode) ParamTypeCallCount() int {
	fake.paramTypeMutex.RLock()
	defer fake.paramTypeMutex.RUnlock()
	return len(fake.paramTypeArgsForCall)
}

func (fake *FakeBlockNode) ParamTypeCalls(stub func(basil.ID) (string, bool)) {
	fake.paramTypeMutex.Lock()
	defer fake.paramTypeMutex.Unlock()
	fake.ParamTypeStub = stub
}

func (fake *FakeBlockNode) ParamTypeArgsForCall(i int) basil.ID {
	fake.paramTypeMutex.RLock()
	defer fake.paramTypeMutex.RUnlock()
	argsForCall := fake.paramTypeArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBlockNode) ParamTypeReturns(result1 string, result2 bool) {
	fake.paramTypeMutex.Lock()
	defer fake.paramTypeMutex.Unlock()
	fake.ParamTypeStub = nil
	fake.paramTypeReturns = struct {
		result1 string
		result2 bool
	}{result1, result2}
}

func (fake *FakeBlockNode) ParamTypeReturnsOnCall(i int, result1 string, result2 bool) {
	fake.paramTypeMutex.Lock()
	defer fake.paramTypeMutex.Unlock()
	fake.ParamTypeStub = nil
	if fake.paramTypeReturnsOnCall == nil {
		fake.paramTypeReturnsOnCall = make(map[int]struct {
			result1 string
			result2 bool
		})
	}
	fake.paramTypeReturnsOnCall[i] = struct {
		result1 string
		result2 bool
	}{result1, result2}
}

func (fake *FakeBlockNode) Pos() parsley.Pos {
	fake.posMutex.Lock()
	ret, specificReturn := fake.posReturnsOnCall[len(fake.posArgsForCall)]
	fake.posArgsForCall = append(fake.posArgsForCall, struct {
	}{})
	fake.recordInvocation("Pos", []interface{}{})
	fake.posMutex.Unlock()
	if fake.PosStub != nil {
		return fake.PosStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.posReturns
	return fakeReturns.result1
}

func (fake *FakeBlockNode) PosCallCount() int {
	fake.posMutex.RLock()
	defer fake.posMutex.RUnlock()
	return len(fake.posArgsForCall)
}

func (fake *FakeBlockNode) PosCalls(stub func() parsley.Pos) {
	fake.posMutex.Lock()
	defer fake.posMutex.Unlock()
	fake.PosStub = stub
}

func (fake *FakeBlockNode) PosReturns(result1 parsley.Pos) {
	fake.posMutex.Lock()
	defer fake.posMutex.Unlock()
	fake.PosStub = nil
	fake.posReturns = struct {
		result1 parsley.Pos
	}{result1}
}

func (fake *FakeBlockNode) PosReturnsOnCall(i int, result1 parsley.Pos) {
	fake.posMutex.Lock()
	defer fake.posMutex.Unlock()
	fake.PosStub = nil
	if fake.posReturnsOnCall == nil {
		fake.posReturnsOnCall = make(map[int]struct {
			result1 parsley.Pos
		})
	}
	fake.posReturnsOnCall[i] = struct {
		result1 parsley.Pos
	}{result1}
}

func (fake *FakeBlockNode) Provides() []basil.ID {
	fake.providesMutex.Lock()
	ret, specificReturn := fake.providesReturnsOnCall[len(fake.providesArgsForCall)]
	fake.providesArgsForCall = append(fake.providesArgsForCall, struct {
	}{})
	fake.recordInvocation("Provides", []interface{}{})
	fake.providesMutex.Unlock()
	if fake.ProvidesStub != nil {
		return fake.ProvidesStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.providesReturns
	return fakeReturns.result1
}

func (fake *FakeBlockNode) ProvidesCallCount() int {
	fake.providesMutex.RLock()
	defer fake.providesMutex.RUnlock()
	return len(fake.providesArgsForCall)
}

func (fake *FakeBlockNode) ProvidesCalls(stub func() []basil.ID) {
	fake.providesMutex.Lock()
	defer fake.providesMutex.Unlock()
	fake.ProvidesStub = stub
}

func (fake *FakeBlockNode) ProvidesReturns(result1 []basil.ID) {
	fake.providesMutex.Lock()
	defer fake.providesMutex.Unlock()
	fake.ProvidesStub = nil
	fake.providesReturns = struct {
		result1 []basil.ID
	}{result1}
}

func (fake *FakeBlockNode) ProvidesReturnsOnCall(i int, result1 []basil.ID) {
	fake.providesMutex.Lock()
	defer fake.providesMutex.Unlock()
	fake.ProvidesStub = nil
	if fake.providesReturnsOnCall == nil {
		fake.providesReturnsOnCall = make(map[int]struct {
			result1 []basil.ID
		})
	}
	fake.providesReturnsOnCall[i] = struct {
		result1 []basil.ID
	}{result1}
}

func (fake *FakeBlockNode) ReaderPos() parsley.Pos {
	fake.readerPosMutex.Lock()
	ret, specificReturn := fake.readerPosReturnsOnCall[len(fake.readerPosArgsForCall)]
	fake.readerPosArgsForCall = append(fake.readerPosArgsForCall, struct {
	}{})
	fake.recordInvocation("ReaderPos", []interface{}{})
	fake.readerPosMutex.Unlock()
	if fake.ReaderPosStub != nil {
		return fake.ReaderPosStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.readerPosReturns
	return fakeReturns.result1
}

func (fake *FakeBlockNode) ReaderPosCallCount() int {
	fake.readerPosMutex.RLock()
	defer fake.readerPosMutex.RUnlock()
	return len(fake.readerPosArgsForCall)
}

func (fake *FakeBlockNode) ReaderPosCalls(stub func() parsley.Pos) {
	fake.readerPosMutex.Lock()
	defer fake.readerPosMutex.Unlock()
	fake.ReaderPosStub = stub
}

func (fake *FakeBlockNode) ReaderPosReturns(result1 parsley.Pos) {
	fake.readerPosMutex.Lock()
	defer fake.readerPosMutex.Unlock()
	fake.ReaderPosStub = nil
	fake.readerPosReturns = struct {
		result1 parsley.Pos
	}{result1}
}

func (fake *FakeBlockNode) ReaderPosReturnsOnCall(i int, result1 parsley.Pos) {
	fake.readerPosMutex.Lock()
	defer fake.readerPosMutex.Unlock()
	fake.ReaderPosStub = nil
	if fake.readerPosReturnsOnCall == nil {
		fake.readerPosReturnsOnCall = make(map[int]struct {
			result1 parsley.Pos
		})
	}
	fake.readerPosReturnsOnCall[i] = struct {
		result1 parsley.Pos
	}{result1}
}

func (fake *FakeBlockNode) Token() string {
	fake.tokenMutex.Lock()
	ret, specificReturn := fake.tokenReturnsOnCall[len(fake.tokenArgsForCall)]
	fake.tokenArgsForCall = append(fake.tokenArgsForCall, struct {
	}{})
	fake.recordInvocation("Token", []interface{}{})
	fake.tokenMutex.Unlock()
	if fake.TokenStub != nil {
		return fake.TokenStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.tokenReturns
	return fakeReturns.result1
}

func (fake *FakeBlockNode) TokenCallCount() int {
	fake.tokenMutex.RLock()
	defer fake.tokenMutex.RUnlock()
	return len(fake.tokenArgsForCall)
}

func (fake *FakeBlockNode) TokenCalls(stub func() string) {
	fake.tokenMutex.Lock()
	defer fake.tokenMutex.Unlock()
	fake.TokenStub = stub
}

func (fake *FakeBlockNode) TokenReturns(result1 string) {
	fake.tokenMutex.Lock()
	defer fake.tokenMutex.Unlock()
	fake.TokenStub = nil
	fake.tokenReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeBlockNode) TokenReturnsOnCall(i int, result1 string) {
	fake.tokenMutex.Lock()
	defer fake.tokenMutex.Unlock()
	fake.TokenStub = nil
	if fake.tokenReturnsOnCall == nil {
		fake.tokenReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.tokenReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeBlockNode) Type() string {
	fake.typeMutex.Lock()
	ret, specificReturn := fake.typeReturnsOnCall[len(fake.typeArgsForCall)]
	fake.typeArgsForCall = append(fake.typeArgsForCall, struct {
	}{})
	fake.recordInvocation("Type", []interface{}{})
	fake.typeMutex.Unlock()
	if fake.TypeStub != nil {
		return fake.TypeStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.typeReturns
	return fakeReturns.result1
}

func (fake *FakeBlockNode) TypeCallCount() int {
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	return len(fake.typeArgsForCall)
}

func (fake *FakeBlockNode) TypeCalls(stub func() string) {
	fake.typeMutex.Lock()
	defer fake.typeMutex.Unlock()
	fake.TypeStub = stub
}

func (fake *FakeBlockNode) TypeReturns(result1 string) {
	fake.typeMutex.Lock()
	defer fake.typeMutex.Unlock()
	fake.TypeStub = nil
	fake.typeReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeBlockNode) TypeReturnsOnCall(i int, result1 string) {
	fake.typeMutex.Lock()
	defer fake.typeMutex.Unlock()
	fake.TypeStub = nil
	if fake.typeReturnsOnCall == nil {
		fake.typeReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.typeReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeBlockNode) Value(arg1 interface{}) (interface{}, parsley.Error) {
	fake.valueMutex.Lock()
	ret, specificReturn := fake.valueReturnsOnCall[len(fake.valueArgsForCall)]
	fake.valueArgsForCall = append(fake.valueArgsForCall, struct {
		arg1 interface{}
	}{arg1})
	fake.recordInvocation("Value", []interface{}{arg1})
	fake.valueMutex.Unlock()
	if fake.ValueStub != nil {
		return fake.ValueStub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.valueReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeBlockNode) ValueCallCount() int {
	fake.valueMutex.RLock()
	defer fake.valueMutex.RUnlock()
	return len(fake.valueArgsForCall)
}

func (fake *FakeBlockNode) ValueCalls(stub func(interface{}) (interface{}, parsley.Error)) {
	fake.valueMutex.Lock()
	defer fake.valueMutex.Unlock()
	fake.ValueStub = stub
}

func (fake *FakeBlockNode) ValueArgsForCall(i int) interface{} {
	fake.valueMutex.RLock()
	defer fake.valueMutex.RUnlock()
	argsForCall := fake.valueArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeBlockNode) ValueReturns(result1 interface{}, result2 parsley.Error) {
	fake.valueMutex.Lock()
	defer fake.valueMutex.Unlock()
	fake.ValueStub = nil
	fake.valueReturns = struct {
		result1 interface{}
		result2 parsley.Error
	}{result1, result2}
}

func (fake *FakeBlockNode) ValueReturnsOnCall(i int, result1 interface{}, result2 parsley.Error) {
	fake.valueMutex.Lock()
	defer fake.valueMutex.Unlock()
	fake.ValueStub = nil
	if fake.valueReturnsOnCall == nil {
		fake.valueReturnsOnCall = make(map[int]struct {
			result1 interface{}
			result2 parsley.Error
		})
	}
	fake.valueReturnsOnCall[i] = struct {
		result1 interface{}
		result2 parsley.Error
	}{result1, result2}
}

func (fake *FakeBlockNode) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.blockTypeMutex.RLock()
	defer fake.blockTypeMutex.RUnlock()
	fake.childrenMutex.RLock()
	defer fake.childrenMutex.RUnlock()
	fake.dependenciesMutex.RLock()
	defer fake.dependenciesMutex.RUnlock()
	fake.evalStageMutex.RLock()
	defer fake.evalStageMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.interpreterMutex.RLock()
	defer fake.interpreterMutex.RUnlock()
	fake.paramTypeMutex.RLock()
	defer fake.paramTypeMutex.RUnlock()
	fake.posMutex.RLock()
	defer fake.posMutex.RUnlock()
	fake.providesMutex.RLock()
	defer fake.providesMutex.RUnlock()
	fake.readerPosMutex.RLock()
	defer fake.readerPosMutex.RUnlock()
	fake.tokenMutex.RLock()
	defer fake.tokenMutex.RUnlock()
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	fake.valueMutex.RLock()
	defer fake.valueMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeBlockNode) recordInvocation(key string, args []interface{}) {
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

var _ basil.BlockNode = new(FakeBlockNode)
