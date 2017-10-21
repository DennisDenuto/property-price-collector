// This file was generated by counterfeiter
package propertypricehistorycomfakes

import (
	"sync"

	"github.com/DennisDenuto/property-price-collector/site/propertypricehistorycom"
)

type FakePostcodeSuburbLookup struct {
	LoadStub        func() error
	loadMutex       sync.RWMutex
	loadArgsForCall []struct{}
	loadReturns     struct {
		result1 error
	}
	GetSuburbStub        func(int) ([]string, bool)
	getSuburbMutex       sync.RWMutex
	getSuburbArgsForCall []struct {
		arg1 int
	}
	getSuburbReturns struct {
		result1 []string
		result2 bool
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakePostcodeSuburbLookup) Load() error {
	fake.loadMutex.Lock()
	fake.loadArgsForCall = append(fake.loadArgsForCall, struct{}{})
	fake.recordInvocation("Load", []interface{}{})
	fake.loadMutex.Unlock()
	if fake.LoadStub != nil {
		return fake.LoadStub()
	} else {
		return fake.loadReturns.result1
	}
}

func (fake *FakePostcodeSuburbLookup) LoadCallCount() int {
	fake.loadMutex.RLock()
	defer fake.loadMutex.RUnlock()
	return len(fake.loadArgsForCall)
}

func (fake *FakePostcodeSuburbLookup) LoadReturns(result1 error) {
	fake.LoadStub = nil
	fake.loadReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakePostcodeSuburbLookup) GetSuburb(arg1 int) ([]string, bool) {
	fake.getSuburbMutex.Lock()
	fake.getSuburbArgsForCall = append(fake.getSuburbArgsForCall, struct {
		arg1 int
	}{arg1})
	fake.recordInvocation("GetSuburb", []interface{}{arg1})
	fake.getSuburbMutex.Unlock()
	if fake.GetSuburbStub != nil {
		return fake.GetSuburbStub(arg1)
	} else {
		return fake.getSuburbReturns.result1, fake.getSuburbReturns.result2
	}
}

func (fake *FakePostcodeSuburbLookup) GetSuburbCallCount() int {
	fake.getSuburbMutex.RLock()
	defer fake.getSuburbMutex.RUnlock()
	return len(fake.getSuburbArgsForCall)
}

func (fake *FakePostcodeSuburbLookup) GetSuburbArgsForCall(i int) int {
	fake.getSuburbMutex.RLock()
	defer fake.getSuburbMutex.RUnlock()
	return fake.getSuburbArgsForCall[i].arg1
}

func (fake *FakePostcodeSuburbLookup) GetSuburbReturns(result1 []string, result2 bool) {
	fake.GetSuburbStub = nil
	fake.getSuburbReturns = struct {
		result1 []string
		result2 bool
	}{result1, result2}
}

func (fake *FakePostcodeSuburbLookup) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.loadMutex.RLock()
	defer fake.loadMutex.RUnlock()
	fake.getSuburbMutex.RLock()
	defer fake.getSuburbMutex.RUnlock()
	return fake.invocations
}

func (fake *FakePostcodeSuburbLookup) recordInvocation(key string, args []interface{}) {
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

var _ propertypricehistorycom.PostcodeSuburbLookup = new(FakePostcodeSuburbLookup)