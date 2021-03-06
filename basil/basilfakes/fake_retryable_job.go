// Code generated by counterfeiter. DO NOT EDIT.
package basilfakes

import (
	"sync"
	"time"

	"github.com/opsidian/basil/basil"
)

type FakeRetryableJob struct {
	JobIDStub        func() int
	jobIDMutex       sync.RWMutex
	jobIDArgsForCall []struct {
	}
	jobIDReturns struct {
		result1 int
	}
	jobIDReturnsOnCall map[int]struct {
		result1 int
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
	RetryCountStub        func() int
	retryCountMutex       sync.RWMutex
	retryCountArgsForCall []struct {
	}
	retryCountReturns struct {
		result1 int
	}
	retryCountReturnsOnCall map[int]struct {
		result1 int
	}
	RetryDelayStub        func(int) time.Duration
	retryDelayMutex       sync.RWMutex
	retryDelayArgsForCall []struct {
		arg1 int
	}
	retryDelayReturns struct {
		result1 time.Duration
	}
	retryDelayReturnsOnCall map[int]struct {
		result1 time.Duration
	}
	RunStub        func()
	runMutex       sync.RWMutex
	runArgsForCall []struct {
	}
	SetJobIDStub        func(int)
	setJobIDMutex       sync.RWMutex
	setJobIDArgsForCall []struct {
		arg1 int
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRetryableJob) JobID() int {
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

func (fake *FakeRetryableJob) JobIDCallCount() int {
	fake.jobIDMutex.RLock()
	defer fake.jobIDMutex.RUnlock()
	return len(fake.jobIDArgsForCall)
}

func (fake *FakeRetryableJob) JobIDCalls(stub func() int) {
	fake.jobIDMutex.Lock()
	defer fake.jobIDMutex.Unlock()
	fake.JobIDStub = stub
}

func (fake *FakeRetryableJob) JobIDReturns(result1 int) {
	fake.jobIDMutex.Lock()
	defer fake.jobIDMutex.Unlock()
	fake.JobIDStub = nil
	fake.jobIDReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeRetryableJob) JobIDReturnsOnCall(i int, result1 int) {
	fake.jobIDMutex.Lock()
	defer fake.jobIDMutex.Unlock()
	fake.JobIDStub = nil
	if fake.jobIDReturnsOnCall == nil {
		fake.jobIDReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.jobIDReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeRetryableJob) JobName() basil.ID {
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

func (fake *FakeRetryableJob) JobNameCallCount() int {
	fake.jobNameMutex.RLock()
	defer fake.jobNameMutex.RUnlock()
	return len(fake.jobNameArgsForCall)
}

func (fake *FakeRetryableJob) JobNameCalls(stub func() basil.ID) {
	fake.jobNameMutex.Lock()
	defer fake.jobNameMutex.Unlock()
	fake.JobNameStub = stub
}

func (fake *FakeRetryableJob) JobNameReturns(result1 basil.ID) {
	fake.jobNameMutex.Lock()
	defer fake.jobNameMutex.Unlock()
	fake.JobNameStub = nil
	fake.jobNameReturns = struct {
		result1 basil.ID
	}{result1}
}

func (fake *FakeRetryableJob) JobNameReturnsOnCall(i int, result1 basil.ID) {
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

func (fake *FakeRetryableJob) Lightweight() bool {
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

func (fake *FakeRetryableJob) LightweightCallCount() int {
	fake.lightweightMutex.RLock()
	defer fake.lightweightMutex.RUnlock()
	return len(fake.lightweightArgsForCall)
}

func (fake *FakeRetryableJob) LightweightCalls(stub func() bool) {
	fake.lightweightMutex.Lock()
	defer fake.lightweightMutex.Unlock()
	fake.LightweightStub = stub
}

func (fake *FakeRetryableJob) LightweightReturns(result1 bool) {
	fake.lightweightMutex.Lock()
	defer fake.lightweightMutex.Unlock()
	fake.LightweightStub = nil
	fake.lightweightReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeRetryableJob) LightweightReturnsOnCall(i int, result1 bool) {
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

func (fake *FakeRetryableJob) RetryCount() int {
	fake.retryCountMutex.Lock()
	ret, specificReturn := fake.retryCountReturnsOnCall[len(fake.retryCountArgsForCall)]
	fake.retryCountArgsForCall = append(fake.retryCountArgsForCall, struct {
	}{})
	fake.recordInvocation("RetryCount", []interface{}{})
	fake.retryCountMutex.Unlock()
	if fake.RetryCountStub != nil {
		return fake.RetryCountStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.retryCountReturns
	return fakeReturns.result1
}

func (fake *FakeRetryableJob) RetryCountCallCount() int {
	fake.retryCountMutex.RLock()
	defer fake.retryCountMutex.RUnlock()
	return len(fake.retryCountArgsForCall)
}

func (fake *FakeRetryableJob) RetryCountCalls(stub func() int) {
	fake.retryCountMutex.Lock()
	defer fake.retryCountMutex.Unlock()
	fake.RetryCountStub = stub
}

func (fake *FakeRetryableJob) RetryCountReturns(result1 int) {
	fake.retryCountMutex.Lock()
	defer fake.retryCountMutex.Unlock()
	fake.RetryCountStub = nil
	fake.retryCountReturns = struct {
		result1 int
	}{result1}
}

func (fake *FakeRetryableJob) RetryCountReturnsOnCall(i int, result1 int) {
	fake.retryCountMutex.Lock()
	defer fake.retryCountMutex.Unlock()
	fake.RetryCountStub = nil
	if fake.retryCountReturnsOnCall == nil {
		fake.retryCountReturnsOnCall = make(map[int]struct {
			result1 int
		})
	}
	fake.retryCountReturnsOnCall[i] = struct {
		result1 int
	}{result1}
}

func (fake *FakeRetryableJob) RetryDelay(arg1 int) time.Duration {
	fake.retryDelayMutex.Lock()
	ret, specificReturn := fake.retryDelayReturnsOnCall[len(fake.retryDelayArgsForCall)]
	fake.retryDelayArgsForCall = append(fake.retryDelayArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("RetryDelay", []interface{}{arg1})
	fake.retryDelayMutex.Unlock()
	if fake.RetryDelayStub != nil {
		return fake.RetryDelayStub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.retryDelayReturns
	return fakeReturns.result1
}

func (fake *FakeRetryableJob) RetryDelayCallCount() int {
	fake.retryDelayMutex.RLock()
	defer fake.retryDelayMutex.RUnlock()
	return len(fake.retryDelayArgsForCall)
}

func (fake *FakeRetryableJob) RetryDelayCalls(stub func(int) time.Duration) {
	fake.retryDelayMutex.Lock()
	defer fake.retryDelayMutex.Unlock()
	fake.RetryDelayStub = stub
}

func (fake *FakeRetryableJob) RetryDelayArgsForCall(i int) int {
	fake.retryDelayMutex.RLock()
	defer fake.retryDelayMutex.RUnlock()
	argsForCall := fake.retryDelayArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRetryableJob) RetryDelayReturns(result1 time.Duration) {
	fake.retryDelayMutex.Lock()
	defer fake.retryDelayMutex.Unlock()
	fake.RetryDelayStub = nil
	fake.retryDelayReturns = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeRetryableJob) RetryDelayReturnsOnCall(i int, result1 time.Duration) {
	fake.retryDelayMutex.Lock()
	defer fake.retryDelayMutex.Unlock()
	fake.RetryDelayStub = nil
	if fake.retryDelayReturnsOnCall == nil {
		fake.retryDelayReturnsOnCall = make(map[int]struct {
			result1 time.Duration
		})
	}
	fake.retryDelayReturnsOnCall[i] = struct {
		result1 time.Duration
	}{result1}
}

func (fake *FakeRetryableJob) Run() {
	fake.runMutex.Lock()
	fake.runArgsForCall = append(fake.runArgsForCall, struct {
	}{})
	fake.recordInvocation("Run", []interface{}{})
	fake.runMutex.Unlock()
	if fake.RunStub != nil {
		fake.RunStub()
	}
}

func (fake *FakeRetryableJob) RunCallCount() int {
	fake.runMutex.RLock()
	defer fake.runMutex.RUnlock()
	return len(fake.runArgsForCall)
}

func (fake *FakeRetryableJob) RunCalls(stub func()) {
	fake.runMutex.Lock()
	defer fake.runMutex.Unlock()
	fake.RunStub = stub
}

func (fake *FakeRetryableJob) SetJobID(arg1 int) {
	fake.setJobIDMutex.Lock()
	fake.setJobIDArgsForCall = append(fake.setJobIDArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("SetJobID", []interface{}{arg1})
	fake.setJobIDMutex.Unlock()
	if fake.SetJobIDStub != nil {
		fake.SetJobIDStub(arg1)
	}
}

func (fake *FakeRetryableJob) SetJobIDCallCount() int {
	fake.setJobIDMutex.RLock()
	defer fake.setJobIDMutex.RUnlock()
	return len(fake.setJobIDArgsForCall)
}

func (fake *FakeRetryableJob) SetJobIDCalls(stub func(int)) {
	fake.setJobIDMutex.Lock()
	defer fake.setJobIDMutex.Unlock()
	fake.SetJobIDStub = stub
}

func (fake *FakeRetryableJob) SetJobIDArgsForCall(i int) int {
	fake.setJobIDMutex.RLock()
	defer fake.setJobIDMutex.RUnlock()
	argsForCall := fake.setJobIDArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeRetryableJob) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.jobIDMutex.RLock()
	defer fake.jobIDMutex.RUnlock()
	fake.jobNameMutex.RLock()
	defer fake.jobNameMutex.RUnlock()
	fake.lightweightMutex.RLock()
	defer fake.lightweightMutex.RUnlock()
	fake.retryCountMutex.RLock()
	defer fake.retryCountMutex.RUnlock()
	fake.retryDelayMutex.RLock()
	defer fake.retryDelayMutex.RUnlock()
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

func (fake *FakeRetryableJob) recordInvocation(key string, args []interface{}) {
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

var _ basil.RetryableJob = new(FakeRetryableJob)
