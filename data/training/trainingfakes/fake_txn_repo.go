// This file was generated by counterfeiter
package trainingfakes

import (
	"sync"

	"github.com/DennisDenuto/property-price-collector/data/training"
)

type FakeTxnRepo struct {
	CreateStub        func() error
	createMutex       sync.RWMutex
	createArgsForCall []struct{}
	createReturns     struct {
		result1 error
	}
	AddStub        func(interface{}) error
	addMutex       sync.RWMutex
	addArgsForCall []struct {
		arg1 interface{}
	}
	addReturns struct {
		result1 error
	}
	StartTxnStub        func() (string, error)
	startTxnMutex       sync.RWMutex
	startTxnArgsForCall []struct{}
	startTxnReturns     struct {
		result1 string
		result2 error
	}
	CommitStub        func(commitId string) error
	commitMutex       sync.RWMutex
	commitArgsForCall []struct {
		commitId string
	}
	commitReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeTxnRepo) Create() error {
	fake.createMutex.Lock()
	fake.createArgsForCall = append(fake.createArgsForCall, struct{}{})
	fake.recordInvocation("Create", []interface{}{})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub()
	} else {
		return fake.createReturns.result1
	}
}

func (fake *FakeTxnRepo) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeTxnRepo) CreateReturns(result1 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTxnRepo) Add(arg1 interface{}) error {
	fake.addMutex.Lock()
	fake.addArgsForCall = append(fake.addArgsForCall, struct {
		arg1 interface{}
	}{arg1})
	fake.recordInvocation("Add", []interface{}{arg1})
	fake.addMutex.Unlock()
	if fake.AddStub != nil {
		return fake.AddStub(arg1)
	} else {
		return fake.addReturns.result1
	}
}

func (fake *FakeTxnRepo) AddCallCount() int {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return len(fake.addArgsForCall)
}

func (fake *FakeTxnRepo) AddArgsForCall(i int) interface{} {
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	return fake.addArgsForCall[i].arg1
}

func (fake *FakeTxnRepo) AddReturns(result1 error) {
	fake.AddStub = nil
	fake.addReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTxnRepo) StartTxn() (string, error) {
	fake.startTxnMutex.Lock()
	fake.startTxnArgsForCall = append(fake.startTxnArgsForCall, struct{}{})
	fake.recordInvocation("StartTxn", []interface{}{})
	fake.startTxnMutex.Unlock()
	if fake.StartTxnStub != nil {
		return fake.StartTxnStub()
	} else {
		return fake.startTxnReturns.result1, fake.startTxnReturns.result2
	}
}

func (fake *FakeTxnRepo) StartTxnCallCount() int {
	fake.startTxnMutex.RLock()
	defer fake.startTxnMutex.RUnlock()
	return len(fake.startTxnArgsForCall)
}

func (fake *FakeTxnRepo) StartTxnReturns(result1 string, result2 error) {
	fake.StartTxnStub = nil
	fake.startTxnReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeTxnRepo) Commit(commitId string) error {
	fake.commitMutex.Lock()
	fake.commitArgsForCall = append(fake.commitArgsForCall, struct {
		commitId string
	}{commitId})
	fake.recordInvocation("Commit", []interface{}{commitId})
	fake.commitMutex.Unlock()
	if fake.CommitStub != nil {
		return fake.CommitStub(commitId)
	} else {
		return fake.commitReturns.result1
	}
}

func (fake *FakeTxnRepo) CommitCallCount() int {
	fake.commitMutex.RLock()
	defer fake.commitMutex.RUnlock()
	return len(fake.commitArgsForCall)
}

func (fake *FakeTxnRepo) CommitArgsForCall(i int) string {
	fake.commitMutex.RLock()
	defer fake.commitMutex.RUnlock()
	return fake.commitArgsForCall[i].commitId
}

func (fake *FakeTxnRepo) CommitReturns(result1 error) {
	fake.CommitStub = nil
	fake.commitReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeTxnRepo) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.addMutex.RLock()
	defer fake.addMutex.RUnlock()
	fake.startTxnMutex.RLock()
	defer fake.startTxnMutex.RUnlock()
	fake.commitMutex.RLock()
	defer fake.commitMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeTxnRepo) recordInvocation(key string, args []interface{}) {
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

var _ training.TxnRepo = new(FakeTxnRepo)
