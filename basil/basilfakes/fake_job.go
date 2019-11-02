// Code generated by counterfeiter. DO NOT EDIT.
package basilfakes

import (
	"sync"

	"github.com/opsidian/basil/basil"
)

type FakeJob struct {
	JobIDStub        func() basil.ID
	jobIDMutex       sync.RWMutex
	jobIDArgsForCall []struct {
	}
	jobIDReturns struct {
		result1 basil.ID
	}
	jobIDReturnsOnCall map[int]struct {
		result1 basil.ID
	}
	JobNameStub        func() basil.ID
	jobNameMutex       sync.RWMutex
	jobNameArgsForCall []struct {
	}
	jobNameReturns struct {
		result1 basil.ID
	}
	jobNameReturnsOnCall map[int]struct {
		result1 basil.ID
	}
	LightweightStub        func() bool
	lightweightMutex       sync.RWMutex
	lightweightArgsForCall []struct {
	}
	lightweightReturns struct {
		result1 bool
	}
	lightweightReturnsOnCall map[int]struct {
		result1 bool
	}
	RunStub        func()
	runMutex       sync.RWMutex
	runArgsForCall []struct {
	}
	SetJobIDStub        func(basil.ID)
	setJobIDMutex       sync.RWMutex
	setJobIDArgsForCall []struct {
		arg1 basil.ID
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeJob) JobID() basil.ID {
	fake.jobIDMutex.Lock()
	ret, specificReturn := fake.jobIDReturnsOnCall[len(fake.jobIDArgsForCall)]
	fake.jobIDArgsForCall = append(fake.jobIDArgsForCall, struct {
	}{})
	fake.recordInvocation("JobID", []interface{}{})
	fake.jobIDMutex.Unlock()
	if fake.JobIDStub != nil {
		return fake.JobIDStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.jobIDReturns
	return fakeReturns.result1
}

func (fake *FakeJob) JobIDCallCount() int {
	fake.jobIDMutex.RLock()
	defer fake.jobIDMutex.RUnlock()
	return len(fake.jobIDArgsForCall)
}

func (fake *FakeJob) JobIDCalls(stub func() basil.ID) {
	fake.jobIDMutex.Lock()
	defer fake.jobIDMutex.Unlock()
	fake.JobIDStub = stub
}

func (fake *FakeJob) JobIDReturns(result1 basil.ID) {
	fake.jobIDMutex.Lock()
	defer fake.jobIDMutex.Unlock()
	fake.JobIDStub = nil
	fake.jobIDReturns = struct {
		result1 basil.ID
	}{result1}
}

func (fake *FakeJob) JobIDReturnsOnCall(i int, result1 basil.ID) {
	fake.jobIDMutex.Lock()
	defer fake.jobIDMutex.Unlock()
	fake.JobIDStub = nil
	if fake.jobIDReturnsOnCall == nil {
		fake.jobIDReturnsOnCall = make(map[int]struct {
			result1 basil.ID
		})
	}
	fake.jobIDReturnsOnCall[i] = struct {
		result1 basil.ID
	}{result1}
}

func (fake *FakeJob) JobName() basil.ID {
	fake.jobNameMutex.Lock()
	ret, specificReturn := fake.jobNameReturnsOnCall[len(fake.jobNameArgsForCall)]
	fake.jobNameArgsForCall = append(fake.jobNameArgsForCall, struct {
	}{})
	fake.recordInvocation("JobName", []interface{}{})
	fake.jobNameMutex.Unlock()
	if fake.JobNameStub != nil {
		return fake.JobNameStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.jobNameReturns
	return fakeReturns.result1
}

func (fake *FakeJob) JobNameCallCount() int {
	fake.jobNameMutex.RLock()
	defer fake.jobNameMutex.RUnlock()
	return len(fake.jobNameArgsForCall)
}

func (fake *FakeJob) JobNameCalls(stub func() basil.ID) {
	fake.jobNameMutex.Lock()
	defer fake.jobNameMutex.Unlock()
	fake.JobNameStub = stub
}

func (fake *FakeJob) JobNameReturns(result1 basil.ID) {
	fake.jobNameMutex.Lock()
	defer fake.jobNameMutex.Unlock()
	fake.JobNameStub = nil
	fake.jobNameReturns = struct {
		result1 basil.ID
	}{result1}
}

func (fake *FakeJob) JobNameReturnsOnCall(i int, result1 basil.ID) {
	fake.jobNameMutex.Lock()
	defer fake.jobNameMutex.Unlock()
	fake.JobNameStub = nil
	if fake.jobNameReturnsOnCall == nil {
		fake.jobNameReturnsOnCall = make(map[int]struct {
			result1 basil.ID
		})
	}
	fake.jobNameReturnsOnCall[i] = struct {
		result1 basil.ID
	}{result1}
}

func (fake *FakeJob) Lightweight() bool {
	fake.lightweightMutex.Lock()
	ret, specificReturn := fake.lightweightReturnsOnCall[len(fake.lightweightArgsForCall)]
	fake.lightweightArgsForCall = append(fake.lightweightArgsForCall, struct {
	}{})
	fake.recordInvocation("Lightweight", []interface{}{})
	fake.lightweightMutex.Unlock()
	if fake.LightweightStub != nil {
		return fake.LightweightStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.lightweightReturns
	return fakeReturns.result1
}

func (fake *FakeJob) LightweightCallCount() int {
	fake.lightweightMutex.RLock()
	defer fake.lightweightMutex.RUnlock()
	return len(fake.lightweightArgsForCall)
}

func (fake *FakeJob) LightweightCalls(stub func() bool) {
	fake.lightweightMutex.Lock()
	defer fake.lightweightMutex.Unlock()
	fake.LightweightStub = stub
}

func (fake *FakeJob) LightweightReturns(result1 bool) {
	fake.lightweightMutex.Lock()
	defer fake.lightweightMutex.Unlock()
	fake.LightweightStub = nil
	fake.lightweightReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeJob) LightweightReturnsOnCall(i int, result1 bool) {
	fake.lightweightMutex.Lock()
	defer fake.lightweightMutex.Unlock()
	fake.LightweightStub = nil
	if fake.lightweightReturnsOnCall == nil {
		fake.lightweightReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.lightweightReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeJob) Run() {
	fake.runMutex.Lock()
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
	}{})
	fake.recordInvocation("Run", []interface{}{})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		fake.RunStub()
	}
}

func (fake *FakeJob) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeJob) RunCalls(stub func()) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = stub
}

func (fake *FakeJob) SetJobID(arg1 basil.ID) {
	fake.setJobIDMutex.Lock()
	fake.setJobIDArgsForCall = append(fake.setJobIDArgsForCall, struct {
		arg1 basil.ID
	}{arg1})
	fake.recordInvocation("SetJobID", []interface{}{arg1})
	fake.setJobIDMutex.Unlock()
	if fake.SetJobIDStub != nil {
		fake.SetJobIDStub(arg1)
	}
}

func (fake *FakeJob) SetJobIDCallCount() int {
	fake.setJobIDMutex.RLock()
	defer fake.setJobIDMutex.RUnlock()
	return len(fake.setJobIDArgsForCall)
}

func (fake *FakeJob) SetJobIDCalls(stub func(basil.ID)) {
	fake.setJobIDMutex.Lock()
	defer fake.setJobIDMutex.Unlock()
	fake.SetJobIDStub = stub
}

func (fake *FakeJob) SetJobIDArgsForCall(i int) basil.ID {
	fake.setJobIDMutex.RLock()
	defer fake.setJobIDMutex.RUnlock()
	argsForCall := fake.setJobIDArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeJob) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.jobIDMutex.RLock()
	defer fake.jobIDMutex.RUnlock()
	fake.jobNameMutex.RLock()
	defer fake.jobNameMutex.RUnlock()
	fake.lightweightMutex.RLock()
	defer fake.lightweightMutex.RUnlock()
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	fake.setJobIDMutex.RLock()
	defer fake.setJobIDMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeJob) recordInvocation(key string, args []interface{}) {
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

var _ basil.Job = new(FakeJob)
