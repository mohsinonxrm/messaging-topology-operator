// Code generated by counterfeiter. DO NOT EDIT.
package internalfakes

import (
	"sync"

	"github.com/hashicorp/vault/api"
	"github.com/rabbitmq/messaging-topology-operator/internal"
)

type FakeVaultDelegate struct {
	AuthenticateStub        func(map[string]interface{}) (*api.Secret, error)
	authenticateMutex       sync.RWMutex
	authenticateArgsForCall []struct {
		arg1 map[string]interface{}
	}
	authenticateReturns struct {
		result1 *api.Secret
		result2 error
	}
	authenticateReturnsOnCall map[int]struct {
		result1 *api.Secret
		result2 error
	}
	ManageTokenLifecycleStub        func(*api.Secret) error
	manageTokenLifecycleMutex       sync.RWMutex
	manageTokenLifecycleArgsForCall []struct {
		arg1 *api.Secret
	}
	manageTokenLifecycleReturns struct {
		result1 error
	}
	manageTokenLifecycleReturnsOnCall map[int]struct {
		result1 error
	}
	ReadSecretStub        func(string) (*api.Secret, error)
	readSecretMutex       sync.RWMutex
	readSecretArgsForCall []struct {
		arg1 string
	}
	readSecretReturns struct {
		result1 *api.Secret
		result2 error
	}
	readSecretReturnsOnCall map[int]struct {
		result1 *api.Secret
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVaultDelegate) Authenticate(arg1 map[string]interface{}) (*api.Secret, error) {
	fake.authenticateMutex.Lock()
	ret, specificReturn := fake.authenticateReturnsOnCall[len(fake.authenticateArgsForCall)]
	fake.authenticateArgsForCall = append(fake.authenticateArgsForCall, struct {
		arg1 map[string]interface{}
	}{arg1})
	stub := fake.AuthenticateStub
	fakeReturns := fake.authenticateReturns
	fake.recordInvocation("Authenticate", []interface{}{arg1})
	fake.authenticateMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVaultDelegate) AuthenticateCallCount() int {
	fake.authenticateMutex.RLock()
	defer fake.authenticateMutex.RUnlock()
	return len(fake.authenticateArgsForCall)
}

func (fake *FakeVaultDelegate) AuthenticateCalls(stub func(map[string]interface{}) (*api.Secret, error)) {
	fake.authenticateMutex.Lock()
	defer fake.authenticateMutex.Unlock()
	fake.AuthenticateStub = stub
}

func (fake *FakeVaultDelegate) AuthenticateArgsForCall(i int) map[string]interface{} {
	fake.authenticateMutex.RLock()
	defer fake.authenticateMutex.RUnlock()
	argsForCall := fake.authenticateArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeVaultDelegate) AuthenticateReturns(result1 *api.Secret, result2 error) {
	fake.authenticateMutex.Lock()
	defer fake.authenticateMutex.Unlock()
	fake.AuthenticateStub = nil
	fake.authenticateReturns = struct {
		result1 *api.Secret
		result2 error
	}{result1, result2}
}

func (fake *FakeVaultDelegate) AuthenticateReturnsOnCall(i int, result1 *api.Secret, result2 error) {
	fake.authenticateMutex.Lock()
	defer fake.authenticateMutex.Unlock()
	fake.AuthenticateStub = nil
	if fake.authenticateReturnsOnCall == nil {
		fake.authenticateReturnsOnCall = make(map[int]struct {
			result1 *api.Secret
			result2 error
		})
	}
	fake.authenticateReturnsOnCall[i] = struct {
		result1 *api.Secret
		result2 error
	}{result1, result2}
}

func (fake *FakeVaultDelegate) ManageTokenLifecycle(arg1 *api.Secret) error {
	fake.manageTokenLifecycleMutex.Lock()
	ret, specificReturn := fake.manageTokenLifecycleReturnsOnCall[len(fake.manageTokenLifecycleArgsForCall)]
	fake.manageTokenLifecycleArgsForCall = append(fake.manageTokenLifecycleArgsForCall, struct {
		arg1 *api.Secret
	}{arg1})
	stub := fake.ManageTokenLifecycleStub
	fakeReturns := fake.manageTokenLifecycleReturns
	fake.recordInvocation("ManageTokenLifecycle", []interface{}{arg1})
	fake.manageTokenLifecycleMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeVaultDelegate) ManageTokenLifecycleCallCount() int {
	fake.manageTokenLifecycleMutex.RLock()
	defer fake.manageTokenLifecycleMutex.RUnlock()
	return len(fake.manageTokenLifecycleArgsForCall)
}

func (fake *FakeVaultDelegate) ManageTokenLifecycleCalls(stub func(*api.Secret) error) {
	fake.manageTokenLifecycleMutex.Lock()
	defer fake.manageTokenLifecycleMutex.Unlock()
	fake.ManageTokenLifecycleStub = stub
}

func (fake *FakeVaultDelegate) ManageTokenLifecycleArgsForCall(i int) *api.Secret {
	fake.manageTokenLifecycleMutex.RLock()
	defer fake.manageTokenLifecycleMutex.RUnlock()
	argsForCall := fake.manageTokenLifecycleArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeVaultDelegate) ManageTokenLifecycleReturns(result1 error) {
	fake.manageTokenLifecycleMutex.Lock()
	defer fake.manageTokenLifecycleMutex.Unlock()
	fake.ManageTokenLifecycleStub = nil
	fake.manageTokenLifecycleReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeVaultDelegate) ManageTokenLifecycleReturnsOnCall(i int, result1 error) {
	fake.manageTokenLifecycleMutex.Lock()
	defer fake.manageTokenLifecycleMutex.Unlock()
	fake.ManageTokenLifecycleStub = nil
	if fake.manageTokenLifecycleReturnsOnCall == nil {
		fake.manageTokenLifecycleReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.manageTokenLifecycleReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeVaultDelegate) ReadSecret(arg1 string) (*api.Secret, error) {
	fake.readSecretMutex.Lock()
	ret, specificReturn := fake.readSecretReturnsOnCall[len(fake.readSecretArgsForCall)]
	fake.readSecretArgsForCall = append(fake.readSecretArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.ReadSecretStub
	fakeReturns := fake.readSecretReturns
	fake.recordInvocation("ReadSecret", []interface{}{arg1})
	fake.readSecretMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVaultDelegate) ReadSecretCallCount() int {
	fake.readSecretMutex.RLock()
	defer fake.readSecretMutex.RUnlock()
	return len(fake.readSecretArgsForCall)
}

func (fake *FakeVaultDelegate) ReadSecretCalls(stub func(string) (*api.Secret, error)) {
	fake.readSecretMutex.Lock()
	defer fake.readSecretMutex.Unlock()
	fake.ReadSecretStub = stub
}

func (fake *FakeVaultDelegate) ReadSecretArgsForCall(i int) string {
	fake.readSecretMutex.RLock()
	defer fake.readSecretMutex.RUnlock()
	argsForCall := fake.readSecretArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeVaultDelegate) ReadSecretReturns(result1 *api.Secret, result2 error) {
	fake.readSecretMutex.Lock()
	defer fake.readSecretMutex.Unlock()
	fake.ReadSecretStub = nil
	fake.readSecretReturns = struct {
		result1 *api.Secret
		result2 error
	}{result1, result2}
}

func (fake *FakeVaultDelegate) ReadSecretReturnsOnCall(i int, result1 *api.Secret, result2 error) {
	fake.readSecretMutex.Lock()
	defer fake.readSecretMutex.Unlock()
	fake.ReadSecretStub = nil
	if fake.readSecretReturnsOnCall == nil {
		fake.readSecretReturnsOnCall = make(map[int]struct {
			result1 *api.Secret
			result2 error
		})
	}
	fake.readSecretReturnsOnCall[i] = struct {
		result1 *api.Secret
		result2 error
	}{result1, result2}
}

func (fake *FakeVaultDelegate) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.authenticateMutex.RLock()
	defer fake.authenticateMutex.RUnlock()
	fake.manageTokenLifecycleMutex.RLock()
	defer fake.manageTokenLifecycleMutex.RUnlock()
	fake.readSecretMutex.RLock()
	defer fake.readSecretMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeVaultDelegate) recordInvocation(key string, args []interface{}) {
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

var _ internal.VaultDelegate = new(FakeVaultDelegate)