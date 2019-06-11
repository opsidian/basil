// Code generated by counterfeiter. DO NOT EDIT.
package basilfakes

import (
	"sync"

	"github.com/opsidian/basil/basil"
	"github.com/opsidian/parsley/parsley"
)

type FakeParameterNode struct {
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
	GeneratedStub        func() bool
	generatedMutex       sync.RWMutex
	generatedArgsForCall []struct {
	}
	generatedReturns struct {
		result1 bool
	}
	generatedReturnsOnCall map[int]struct {
		result1 bool
	}
	GeneratesStub        func() []basil.ID
	generatesMutex       sync.RWMutex
	generatesArgsForCall []struct {
	}
	generatesReturns struct {
		result1 []basil.ID
	}
	generatesReturnsOnCall map[int]struct {
		result1 []basil.ID
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
	IsDeclarationStub        func() bool
	isDeclarationMutex       sync.RWMutex
	isDeclarationArgsForCall []struct {
	}
	isDeclarationReturns struct {
		result1 bool
	}
	isDeclarationReturnsOnCall map[int]struct {
		result1 bool
	}
	NameStub        func() basil.ID
	nameMutex       sync.RWMutex
	nameArgsForCall []struct {
	}
	nameReturns struct {
		result1 basil.ID
	}
	nameReturnsOnCall map[int]struct {
		result1 basil.ID
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
	SetDescriptorStub        func(basil.ParameterDescriptor)
	setDescriptorMutex       sync.RWMutex
	setDescriptorArgsForCall []struct {
		arg1 basil.ParameterDescriptor
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
	ValueNodeStub        func() parsley.Node
	valueNodeMutex       sync.RWMutex
	valueNodeArgsForCall []struct {
	}
	valueNodeReturns struct {
		result1 parsley.Node
	}
	valueNodeReturnsOnCall map[int]struct {
		result1 parsley.Node
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeParameterNode) Dependencies() basil.Dependencies {
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

func (fake *FakeParameterNode) DependenciesCallCount() int {
	fake.dependenciesMutex.RLock()
	defer fake.dependenciesMutex.RUnlock()
	return len(fake.dependenciesArgsForCall)
}

func (fake *FakeParameterNode) DependenciesCalls(stub func() basil.Dependencies) {
	fake.dependenciesMutex.Lock()
	defer fake.dependenciesMutex.Unlock()
	fake.DependenciesStub = stub
}

func (fake *FakeParameterNode) DependenciesReturns(result1 basil.Dependencies) {
	fake.dependenciesMutex.Lock()
	defer fake.dependenciesMutex.Unlock()
	fake.DependenciesStub = nil
	fake.dependenciesReturns = struct {
		result1 basil.Dependencies
	}{result1}
}

func (fake *FakeParameterNode) DependenciesReturnsOnCall(i int, result1 basil.Dependencies) {
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

func (fake *FakeParameterNode) EvalStage() basil.EvalStage {
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

func (fake *FakeParameterNode) EvalStageCallCount() int {
	fake.evalStageMutex.RLock()
	defer fake.evalStageMutex.RUnlock()
	return len(fake.evalStageArgsForCall)
}

func (fake *FakeParameterNode) EvalStageCalls(stub func() basil.EvalStage) {
	fake.evalStageMutex.Lock()
	defer fake.evalStageMutex.Unlock()
	fake.EvalStageStub = stub
}

func (fake *FakeParameterNode) EvalStageReturns(result1 basil.EvalStage) {
	fake.evalStageMutex.Lock()
	defer fake.evalStageMutex.Unlock()
	fake.EvalStageStub = nil
	fake.evalStageReturns = struct {
		result1 basil.EvalStage
	}{result1}
}

func (fake *FakeParameterNode) EvalStageReturnsOnCall(i int, result1 basil.EvalStage) {
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

func (fake *FakeParameterNode) Generated() bool {
	fake.generatedMutex.Lock()
	ret, specificReturn := fake.generatedReturnsOnCall[len(fake.generatedArgsForCall)]
	fake.generatedArgsForCall = append(fake.generatedArgsForCall, struct {
	}{})
	fake.recordInvocation("Generated", []interface{}{})
	fake.generatedMutex.Unlock()
	if fake.GeneratedStub != nil {
		return fake.GeneratedStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.generatedReturns
	return fakeReturns.result1
}

func (fake *FakeParameterNode) GeneratedCallCount() int {
	fake.generatedMutex.RLock()
	defer fake.generatedMutex.RUnlock()
	return len(fake.generatedArgsForCall)
}

func (fake *FakeParameterNode) GeneratedCalls(stub func() bool) {
	fake.generatedMutex.Lock()
	defer fake.generatedMutex.Unlock()
	fake.GeneratedStub = stub
}

func (fake *FakeParameterNode) GeneratedReturns(result1 bool) {
	fake.generatedMutex.Lock()
	defer fake.generatedMutex.Unlock()
	fake.GeneratedStub = nil
	fake.generatedReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeParameterNode) GeneratedReturnsOnCall(i int, result1 bool) {
	fake.generatedMutex.Lock()
	defer fake.generatedMutex.Unlock()
	fake.GeneratedStub = nil
	if fake.generatedReturnsOnCall == nil {
		fake.generatedReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.generatedReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeParameterNode) Generates() []basil.ID {
	fake.generatesMutex.Lock()
	ret, specificReturn := fake.generatesReturnsOnCall[len(fake.generatesArgsForCall)]
	fake.generatesArgsForCall = append(fake.generatesArgsForCall, struct {
	}{})
	fake.recordInvocation("Generates", []interface{}{})
	fake.generatesMutex.Unlock()
	if fake.GeneratesStub != nil {
		return fake.GeneratesStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.generatesReturns
	return fakeReturns.result1
}

func (fake *FakeParameterNode) GeneratesCallCount() int {
	fake.generatesMutex.RLock()
	defer fake.generatesMutex.RUnlock()
	return len(fake.generatesArgsForCall)
}

func (fake *FakeParameterNode) GeneratesCalls(stub func() []basil.ID) {
	fake.generatesMutex.Lock()
	defer fake.generatesMutex.Unlock()
	fake.GeneratesStub = stub
}

func (fake *FakeParameterNode) GeneratesReturns(result1 []basil.ID) {
	fake.generatesMutex.Lock()
	defer fake.generatesMutex.Unlock()
	fake.GeneratesStub = nil
	fake.generatesReturns = struct {
		result1 []basil.ID
	}{result1}
}

func (fake *FakeParameterNode) GeneratesReturnsOnCall(i int, result1 []basil.ID) {
	fake.generatesMutex.Lock()
	defer fake.generatesMutex.Unlock()
	fake.GeneratesStub = nil
	if fake.generatesReturnsOnCall == nil {
		fake.generatesReturnsOnCall = make(map[int]struct {
			result1 []basil.ID
		})
	}
	fake.generatesReturnsOnCall[i] = struct {
		result1 []basil.ID
	}{result1}
}

func (fake *FakeParameterNode) ID() basil.ID {
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

func (fake *FakeParameterNode) IDCallCount() int {
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	return len(fake.iDArgsForCall)
}

func (fake *FakeParameterNode) IDCalls(stub func() basil.ID) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = stub
}

func (fake *FakeParameterNode) IDReturns(result1 basil.ID) {
	fake.iDMutex.Lock()
	defer fake.iDMutex.Unlock()
	fake.IDStub = nil
	fake.iDReturns = struct {
		result1 basil.ID
	}{result1}
}

func (fake *FakeParameterNode) IDReturnsOnCall(i int, result1 basil.ID) {
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

func (fake *FakeParameterNode) IsDeclaration() bool {
	fake.isDeclarationMutex.Lock()
	ret, specificReturn := fake.isDeclarationReturnsOnCall[len(fake.isDeclarationArgsForCall)]
	fake.isDeclarationArgsForCall = append(fake.isDeclarationArgsForCall, struct {
	}{})
	fake.recordInvocation("IsDeclaration", []interface{}{})
	fake.isDeclarationMutex.Unlock()
	if fake.IsDeclarationStub != nil {
		return fake.IsDeclarationStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.isDeclarationReturns
	return fakeReturns.result1
}

func (fake *FakeParameterNode) IsDeclarationCallCount() int {
	fake.isDeclarationMutex.RLock()
	defer fake.isDeclarationMutex.RUnlock()
	return len(fake.isDeclarationArgsForCall)
}

func (fake *FakeParameterNode) IsDeclarationCalls(stub func() bool) {
	fake.isDeclarationMutex.Lock()
	defer fake.isDeclarationMutex.Unlock()
	fake.IsDeclarationStub = stub
}

func (fake *FakeParameterNode) IsDeclarationReturns(result1 bool) {
	fake.isDeclarationMutex.Lock()
	defer fake.isDeclarationMutex.Unlock()
	fake.IsDeclarationStub = nil
	fake.isDeclarationReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeParameterNode) IsDeclarationReturnsOnCall(i int, result1 bool) {
	fake.isDeclarationMutex.Lock()
	defer fake.isDeclarationMutex.Unlock()
	fake.IsDeclarationStub = nil
	if fake.isDeclarationReturnsOnCall == nil {
		fake.isDeclarationReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isDeclarationReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeParameterNode) Name() basil.ID {
	fake.nameMutex.Lock()
	ret, specificReturn := fake.nameReturnsOnCall[len(fake.nameArgsForCall)]
	fake.nameArgsForCall = append(fake.nameArgsForCall, struct {
	}{})
	fake.recordInvocation("Name", []interface{}{})
	fake.nameMutex.Unlock()
	if fake.NameStub != nil {
		return fake.NameStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.nameReturns
	return fakeReturns.result1
}

func (fake *FakeParameterNode) NameCallCount() int {
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	return len(fake.nameArgsForCall)
}

func (fake *FakeParameterNode) NameCalls(stub func() basil.ID) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = stub
}

func (fake *FakeParameterNode) NameReturns(result1 basil.ID) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	fake.nameReturns = struct {
		result1 basil.ID
	}{result1}
}

func (fake *FakeParameterNode) NameReturnsOnCall(i int, result1 basil.ID) {
	fake.nameMutex.Lock()
	defer fake.nameMutex.Unlock()
	fake.NameStub = nil
	if fake.nameReturnsOnCall == nil {
		fake.nameReturnsOnCall = make(map[int]struct {
			result1 basil.ID
		})
	}
	fake.nameReturnsOnCall[i] = struct {
		result1 basil.ID
	}{result1}
}

func (fake *FakeParameterNode) Pos() parsley.Pos {
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

func (fake *FakeParameterNode) PosCallCount() int {
	fake.posMutex.RLock()
	defer fake.posMutex.RUnlock()
	return len(fake.posArgsForCall)
}

func (fake *FakeParameterNode) PosCalls(stub func() parsley.Pos) {
	fake.posMutex.Lock()
	defer fake.posMutex.Unlock()
	fake.PosStub = stub
}

func (fake *FakeParameterNode) PosReturns(result1 parsley.Pos) {
	fake.posMutex.Lock()
	defer fake.posMutex.Unlock()
	fake.PosStub = nil
	fake.posReturns = struct {
		result1 parsley.Pos
	}{result1}
}

func (fake *FakeParameterNode) PosReturnsOnCall(i int, result1 parsley.Pos) {
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

func (fake *FakeParameterNode) Provides() []basil.ID {
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

func (fake *FakeParameterNode) ProvidesCallCount() int {
	fake.providesMutex.RLock()
	defer fake.providesMutex.RUnlock()
	return len(fake.providesArgsForCall)
}

func (fake *FakeParameterNode) ProvidesCalls(stub func() []basil.ID) {
	fake.providesMutex.Lock()
	defer fake.providesMutex.Unlock()
	fake.ProvidesStub = stub
}

func (fake *FakeParameterNode) ProvidesReturns(result1 []basil.ID) {
	fake.providesMutex.Lock()
	defer fake.providesMutex.Unlock()
	fake.ProvidesStub = nil
	fake.providesReturns = struct {
		result1 []basil.ID
	}{result1}
}

func (fake *FakeParameterNode) ProvidesReturnsOnCall(i int, result1 []basil.ID) {
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

func (fake *FakeParameterNode) ReaderPos() parsley.Pos {
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

func (fake *FakeParameterNode) ReaderPosCallCount() int {
	fake.readerPosMutex.RLock()
	defer fake.readerPosMutex.RUnlock()
	return len(fake.readerPosArgsForCall)
}

func (fake *FakeParameterNode) ReaderPosCalls(stub func() parsley.Pos) {
	fake.readerPosMutex.Lock()
	defer fake.readerPosMutex.Unlock()
	fake.ReaderPosStub = stub
}

func (fake *FakeParameterNode) ReaderPosReturns(result1 parsley.Pos) {
	fake.readerPosMutex.Lock()
	defer fake.readerPosMutex.Unlock()
	fake.ReaderPosStub = nil
	fake.readerPosReturns = struct {
		result1 parsley.Pos
	}{result1}
}

func (fake *FakeParameterNode) ReaderPosReturnsOnCall(i int, result1 parsley.Pos) {
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

func (fake *FakeParameterNode) SetDescriptor(arg1 basil.ParameterDescriptor) {
	fake.setDescriptorMutex.Lock()
	fake.setDescriptorArgsForCall = append(fake.setDescriptorArgsForCall, struct {
		arg1 basil.ParameterDescriptor
	}{arg1})
	fake.recordInvocation("SetDescriptor", []interface{}{arg1})
	fake.setDescriptorMutex.Unlock()
	if fake.SetDescriptorStub != nil {
		fake.SetDescriptorStub(arg1)
	}
}

func (fake *FakeParameterNode) SetDescriptorCallCount() int {
	fake.setDescriptorMutex.RLock()
	defer fake.setDescriptorMutex.RUnlock()
	return len(fake.setDescriptorArgsForCall)
}

func (fake *FakeParameterNode) SetDescriptorCalls(stub func(basil.ParameterDescriptor)) {
	fake.setDescriptorMutex.Lock()
	defer fake.setDescriptorMutex.Unlock()
	fake.SetDescriptorStub = stub
}

func (fake *FakeParameterNode) SetDescriptorArgsForCall(i int) basil.ParameterDescriptor {
	fake.setDescriptorMutex.RLock()
	defer fake.setDescriptorMutex.RUnlock()
	argsForCall := fake.setDescriptorArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeParameterNode) Token() string {
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

func (fake *FakeParameterNode) TokenCallCount() int {
	fake.tokenMutex.RLock()
	defer fake.tokenMutex.RUnlock()
	return len(fake.tokenArgsForCall)
}

func (fake *FakeParameterNode) TokenCalls(stub func() string) {
	fake.tokenMutex.Lock()
	defer fake.tokenMutex.Unlock()
	fake.TokenStub = stub
}

func (fake *FakeParameterNode) TokenReturns(result1 string) {
	fake.tokenMutex.Lock()
	defer fake.tokenMutex.Unlock()
	fake.TokenStub = nil
	fake.tokenReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeParameterNode) TokenReturnsOnCall(i int, result1 string) {
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

func (fake *FakeParameterNode) Type() string {
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

func (fake *FakeParameterNode) TypeCallCount() int {
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	return len(fake.typeArgsForCall)
}

func (fake *FakeParameterNode) TypeCalls(stub func() string) {
	fake.typeMutex.Lock()
	defer fake.typeMutex.Unlock()
	fake.TypeStub = stub
}

func (fake *FakeParameterNode) TypeReturns(result1 string) {
	fake.typeMutex.Lock()
	defer fake.typeMutex.Unlock()
	fake.TypeStub = nil
	fake.typeReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeParameterNode) TypeReturnsOnCall(i int, result1 string) {
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

func (fake *FakeParameterNode) Value(arg1 interface{}) (interface{}, parsley.Error) {
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

func (fake *FakeParameterNode) ValueCallCount() int {
	fake.valueMutex.RLock()
	defer fake.valueMutex.RUnlock()
	return len(fake.valueArgsForCall)
}

func (fake *FakeParameterNode) ValueCalls(stub func(interface{}) (interface{}, parsley.Error)) {
	fake.valueMutex.Lock()
	defer fake.valueMutex.Unlock()
	fake.ValueStub = stub
}

func (fake *FakeParameterNode) ValueArgsForCall(i int) interface{} {
	fake.valueMutex.RLock()
	defer fake.valueMutex.RUnlock()
	argsForCall := fake.valueArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeParameterNode) ValueReturns(result1 interface{}, result2 parsley.Error) {
	fake.valueMutex.Lock()
	defer fake.valueMutex.Unlock()
	fake.ValueStub = nil
	fake.valueReturns = struct {
		result1 interface{}
		result2 parsley.Error
	}{result1, result2}
}

func (fake *FakeParameterNode) ValueReturnsOnCall(i int, result1 interface{}, result2 parsley.Error) {
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

func (fake *FakeParameterNode) ValueNode() parsley.Node {
	fake.valueNodeMutex.Lock()
	ret, specificReturn := fake.valueNodeReturnsOnCall[len(fake.valueNodeArgsForCall)]
	fake.valueNodeArgsForCall = append(fake.valueNodeArgsForCall, struct {
	}{})
	fake.recordInvocation("ValueNode", []interface{}{})
	fake.valueNodeMutex.Unlock()
	if fake.ValueNodeStub != nil {
		return fake.ValueNodeStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.valueNodeReturns
	return fakeReturns.result1
}

func (fake *FakeParameterNode) ValueNodeCallCount() int {
	fake.valueNodeMutex.RLock()
	defer fake.valueNodeMutex.RUnlock()
	return len(fake.valueNodeArgsForCall)
}

func (fake *FakeParameterNode) ValueNodeCalls(stub func() parsley.Node) {
	fake.valueNodeMutex.Lock()
	defer fake.valueNodeMutex.Unlock()
	fake.ValueNodeStub = stub
}

func (fake *FakeParameterNode) ValueNodeReturns(result1 parsley.Node) {
	fake.valueNodeMutex.Lock()
	defer fake.valueNodeMutex.Unlock()
	fake.ValueNodeStub = nil
	fake.valueNodeReturns = struct {
		result1 parsley.Node
	}{result1}
}

func (fake *FakeParameterNode) ValueNodeReturnsOnCall(i int, result1 parsley.Node) {
	fake.valueNodeMutex.Lock()
	defer fake.valueNodeMutex.Unlock()
	fake.ValueNodeStub = nil
	if fake.valueNodeReturnsOnCall == nil {
		fake.valueNodeReturnsOnCall = make(map[int]struct {
			result1 parsley.Node
		})
	}
	fake.valueNodeReturnsOnCall[i] = struct {
		result1 parsley.Node
	}{result1}
}

func (fake *FakeParameterNode) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.dependenciesMutex.RLock()
	defer fake.dependenciesMutex.RUnlock()
	fake.evalStageMutex.RLock()
	defer fake.evalStageMutex.RUnlock()
	fake.generatedMutex.RLock()
	defer fake.generatedMutex.RUnlock()
	fake.generatesMutex.RLock()
	defer fake.generatesMutex.RUnlock()
	fake.iDMutex.RLock()
	defer fake.iDMutex.RUnlock()
	fake.isDeclarationMutex.RLock()
	defer fake.isDeclarationMutex.RUnlock()
	fake.nameMutex.RLock()
	defer fake.nameMutex.RUnlock()
	fake.posMutex.RLock()
	defer fake.posMutex.RUnlock()
	fake.providesMutex.RLock()
	defer fake.providesMutex.RUnlock()
	fake.readerPosMutex.RLock()
	defer fake.readerPosMutex.RUnlock()
	fake.setDescriptorMutex.RLock()
	defer fake.setDescriptorMutex.RUnlock()
	fake.tokenMutex.RLock()
	defer fake.tokenMutex.RUnlock()
	fake.typeMutex.RLock()
	defer fake.typeMutex.RUnlock()
	fake.valueMutex.RLock()
	defer fake.valueMutex.RUnlock()
	fake.valueNodeMutex.RLock()
	defer fake.valueNodeMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeParameterNode) recordInvocation(key string, args []interface{}) {
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

var _ basil.ParameterNode = new(FakeParameterNode)
