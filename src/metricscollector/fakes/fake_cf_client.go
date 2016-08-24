// This file was generated by counterfeiter
package fakes

import (
	"cf"
	"sync"
)

type FakeCfClient struct {
	LoginStub        func() error
	loginMutex       sync.RWMutex
	loginArgsForCall []struct{}
	loginReturns     struct {
		result1 error
	}
	RefreshAuthTokenStub        func() (string, error)
	refreshAuthTokenMutex       sync.RWMutex
	refreshAuthTokenArgsForCall []struct{}
	refreshAuthTokenReturns     struct {
		result1 string
		result2 error
	}
	GetTokensStub        func() cf.Tokens
	getTokensMutex       sync.RWMutex
	getTokensArgsForCall []struct{}
	getTokensReturns     struct {
		result1 cf.Tokens
	}
	GetTokensWithRefreshStub        func() cf.Tokens
	getTokensWithRefreshMutex       sync.RWMutex
	getTokensWithRefreshArgsForCall []struct{}
	getTokensWithRefreshReturns     struct {
		result1 cf.Tokens
	}
	GetEndpointsStub        func() cf.Endpoints
	getEndpointsMutex       sync.RWMutex
	getEndpointsArgsForCall []struct{}
	getEndpointsReturns     struct {
		result1 cf.Endpoints
	}
	GetAppInstancesStub        func(string) (int, error)
	getAppInstancesMutex       sync.RWMutex
	getAppInstancesArgsForCall []struct {
		arg1 string
	}
	getAppInstancesReturns struct {
		result1 int
		result2 error
	}
	SetAppInstancesStub        func(string, int) error
	setAppInstancesMutex       sync.RWMutex
	setAppInstancesArgsForCall []struct {
		arg1 string
		arg2 int
	}
	setAppInstancesReturns struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeCfClient) Login() error {
	fake.loginMutex.Lock()
	fake.loginArgsForCall = append(fake.loginArgsForCall, struct{}{})
	fake.recordInvocation("Login", []interface{}{})
	fake.loginMutex.Unlock()
	if fake.LoginStub != nil {
		return fake.LoginStub()
	} else {
		return fake.loginReturns.result1
	}
}

func (fake *FakeCfClient) LoginCallCount() int {
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	return len(fake.loginArgsForCall)
}

func (fake *FakeCfClient) LoginReturns(result1 error) {
	fake.LoginStub = nil
	fake.loginReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCfClient) RefreshAuthToken() (string, error) {
	fake.refreshAuthTokenMutex.Lock()
	fake.refreshAuthTokenArgsForCall = append(fake.refreshAuthTokenArgsForCall, struct{}{})
	fake.recordInvocation("RefreshAuthToken", []interface{}{})
	fake.refreshAuthTokenMutex.Unlock()
	if fake.RefreshAuthTokenStub != nil {
		return fake.RefreshAuthTokenStub()
	} else {
		return fake.refreshAuthTokenReturns.result1, fake.refreshAuthTokenReturns.result2
	}
}

func (fake *FakeCfClient) RefreshAuthTokenCallCount() int {
	fake.refreshAuthTokenMutex.RLock()
	defer fake.refreshAuthTokenMutex.RUnlock()
	return len(fake.refreshAuthTokenArgsForCall)
}

func (fake *FakeCfClient) RefreshAuthTokenReturns(result1 string, result2 error) {
	fake.RefreshAuthTokenStub = nil
	fake.refreshAuthTokenReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeCfClient) GetTokens() cf.Tokens {
	fake.getTokensMutex.Lock()
	fake.getTokensArgsForCall = append(fake.getTokensArgsForCall, struct{}{})
	fake.recordInvocation("GetTokens", []interface{}{})
	fake.getTokensMutex.Unlock()
	if fake.GetTokensStub != nil {
		return fake.GetTokensStub()
	} else {
		return fake.getTokensReturns.result1
	}
}

func (fake *FakeCfClient) GetTokensCallCount() int {
	fake.getTokensMutex.RLock()
	defer fake.getTokensMutex.RUnlock()
	return len(fake.getTokensArgsForCall)
}

func (fake *FakeCfClient) GetTokensReturns(result1 cf.Tokens) {
	fake.GetTokensStub = nil
	fake.getTokensReturns = struct {
		result1 cf.Tokens
	}{result1}
}

func (fake *FakeCfClient) GetTokensWithRefresh() cf.Tokens {
	fake.getTokensWithRefreshMutex.Lock()
	fake.getTokensWithRefreshArgsForCall = append(fake.getTokensWithRefreshArgsForCall, struct{}{})
	fake.recordInvocation("GetTokensWithRefresh", []interface{}{})
	fake.getTokensWithRefreshMutex.Unlock()
	if fake.GetTokensWithRefreshStub != nil {
		return fake.GetTokensWithRefreshStub()
	} else {
		return fake.getTokensWithRefreshReturns.result1
	}
}

func (fake *FakeCfClient) GetTokensWithRefreshCallCount() int {
	fake.getTokensWithRefreshMutex.RLock()
	defer fake.getTokensWithRefreshMutex.RUnlock()
	return len(fake.getTokensWithRefreshArgsForCall)
}

func (fake *FakeCfClient) GetTokensWithRefreshReturns(result1 cf.Tokens) {
	fake.GetTokensWithRefreshStub = nil
	fake.getTokensWithRefreshReturns = struct {
		result1 cf.Tokens
	}{result1}
}

func (fake *FakeCfClient) GetEndpoints() cf.Endpoints {
	fake.getEndpointsMutex.Lock()
	fake.getEndpointsArgsForCall = append(fake.getEndpointsArgsForCall, struct{}{})
	fake.recordInvocation("GetEndpoints", []interface{}{})
	fake.getEndpointsMutex.Unlock()
	if fake.GetEndpointsStub != nil {
		return fake.GetEndpointsStub()
	} else {
		return fake.getEndpointsReturns.result1
	}
}

func (fake *FakeCfClient) GetEndpointsCallCount() int {
	fake.getEndpointsMutex.RLock()
	defer fake.getEndpointsMutex.RUnlock()
	return len(fake.getEndpointsArgsForCall)
}

func (fake *FakeCfClient) GetEndpointsReturns(result1 cf.Endpoints) {
	fake.GetEndpointsStub = nil
	fake.getEndpointsReturns = struct {
		result1 cf.Endpoints
	}{result1}
}

func (fake *FakeCfClient) GetAppInstances(arg1 string) (int, error) {
	fake.getAppInstancesMutex.Lock()
	fake.getAppInstancesArgsForCall = append(fake.getAppInstancesArgsForCall, struct {
		arg1 string
	}{arg1})
	fake.recordInvocation("GetAppInstances", []interface{}{arg1})
	fake.getAppInstancesMutex.Unlock()
	if fake.GetAppInstancesStub != nil {
		return fake.GetAppInstancesStub(arg1)
	} else {
		return fake.getAppInstancesReturns.result1, fake.getAppInstancesReturns.result2
	}
}

func (fake *FakeCfClient) GetAppInstancesCallCount() int {
	fake.getAppInstancesMutex.RLock()
	defer fake.getAppInstancesMutex.RUnlock()
	return len(fake.getAppInstancesArgsForCall)
}

func (fake *FakeCfClient) GetAppInstancesArgsForCall(i int) string {
	fake.getAppInstancesMutex.RLock()
	defer fake.getAppInstancesMutex.RUnlock()
	return fake.getAppInstancesArgsForCall[i].arg1
}

func (fake *FakeCfClient) GetAppInstancesReturns(result1 int, result2 error) {
	fake.GetAppInstancesStub = nil
	fake.getAppInstancesReturns = struct {
		result1 int
		result2 error
	}{result1, result2}
}

func (fake *FakeCfClient) SetAppInstances(arg1 string, arg2 int) error {
	fake.setAppInstancesMutex.Lock()
	fake.setAppInstancesArgsForCall = append(fake.setAppInstancesArgsForCall, struct {
		arg1 string
		arg2 int
	}{arg1, arg2})
	fake.recordInvocation("SetAppInstances", []interface{}{arg1, arg2})
	fake.setAppInstancesMutex.Unlock()
	if fake.SetAppInstancesStub != nil {
		return fake.SetAppInstancesStub(arg1, arg2)
	} else {
		return fake.setAppInstancesReturns.result1
	}
}

func (fake *FakeCfClient) SetAppInstancesCallCount() int {
	fake.setAppInstancesMutex.RLock()
	defer fake.setAppInstancesMutex.RUnlock()
	return len(fake.setAppInstancesArgsForCall)
}

func (fake *FakeCfClient) SetAppInstancesArgsForCall(i int) (string, int) {
	fake.setAppInstancesMutex.RLock()
	defer fake.setAppInstancesMutex.RUnlock()
	return fake.setAppInstancesArgsForCall[i].arg1, fake.setAppInstancesArgsForCall[i].arg2
}

func (fake *FakeCfClient) SetAppInstancesReturns(result1 error) {
	fake.SetAppInstancesStub = nil
	fake.setAppInstancesReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeCfClient) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.loginMutex.RLock()
	defer fake.loginMutex.RUnlock()
	fake.refreshAuthTokenMutex.RLock()
	defer fake.refreshAuthTokenMutex.RUnlock()
	fake.getTokensMutex.RLock()
	defer fake.getTokensMutex.RUnlock()
	fake.getTokensWithRefreshMutex.RLock()
	defer fake.getTokensWithRefreshMutex.RUnlock()
	fake.getEndpointsMutex.RLock()
	defer fake.getEndpointsMutex.RUnlock()
	fake.getAppInstancesMutex.RLock()
	defer fake.getAppInstancesMutex.RUnlock()
	fake.setAppInstancesMutex.RLock()
	defer fake.setAppInstancesMutex.RUnlock()
	return fake.invocations
}

func (fake *FakeCfClient) recordInvocation(key string, args []interface{}) {
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

var _ cf.CfClient = new(FakeCfClient)
