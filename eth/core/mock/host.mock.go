// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"pkg.berachain.dev/stargazer/eth/core"
	"pkg.berachain.dev/stargazer/eth/core/precompile"
	"pkg.berachain.dev/stargazer/eth/core/state"
	"sync"
)

// Ensure, that StargazerHostChainMock does implement core.StargazerHostChain.
// If this is not the case, regenerate this file with moq.
var _ core.StargazerHostChain = &StargazerHostChainMock{}

// StargazerHostChainMock is a mock implementation of core.StargazerHostChain.
//
//	func TestSomethingThatUsesStargazerHostChain(t *testing.T) {
//
//		// make and configure a mocked core.StargazerHostChain
//		mockedStargazerHostChain := &StargazerHostChainMock{
//			GetBlockPluginFunc: func() core.BlockPlugin {
//				panic("mock out the GetBlockPlugin method")
//			},
//			GetConfigurationPluginFunc: func() core.ConfigurationPlugin {
//				panic("mock out the GetConfigurationPlugin method")
//			},
//			GetGasPluginFunc: func() core.GasPlugin {
//				panic("mock out the GetGasPlugin method")
//			},
//			GetPrecompilePluginFunc: func() precompile.Plugin {
//				panic("mock out the GetPrecompilePlugin method")
//			},
//			GetStatePluginFunc: func() state.Plugin {
//				panic("mock out the GetStatePlugin method")
//			},
//		}
//
//		// use mockedStargazerHostChain in code that requires core.StargazerHostChain
//		// and then make assertions.
//
//	}
type StargazerHostChainMock struct {
	// GetBlockPluginFunc mocks the GetBlockPlugin method.
	GetBlockPluginFunc func() core.BlockPlugin

	// GetConfigurationPluginFunc mocks the GetConfigurationPlugin method.
	GetConfigurationPluginFunc func() core.ConfigurationPlugin

	// GetGasPluginFunc mocks the GetGasPlugin method.
	GetGasPluginFunc func() core.GasPlugin

	// GetPrecompilePluginFunc mocks the GetPrecompilePlugin method.
	GetPrecompilePluginFunc func() precompile.Plugin

	// GetStatePluginFunc mocks the GetStatePlugin method.
	GetStatePluginFunc func() state.Plugin

	// calls tracks calls to the methods.
	calls struct {
		// GetBlockPlugin holds details about calls to the GetBlockPlugin method.
		GetBlockPlugin []struct {
		}
		// GetConfigurationPlugin holds details about calls to the GetConfigurationPlugin method.
		GetConfigurationPlugin []struct {
		}
		// GetGasPlugin holds details about calls to the GetGasPlugin method.
		GetGasPlugin []struct {
		}
		// GetPrecompilePlugin holds details about calls to the GetPrecompilePlugin method.
		GetPrecompilePlugin []struct {
		}
		// GetStatePlugin holds details about calls to the GetStatePlugin method.
		GetStatePlugin []struct {
		}
	}
	lockGetBlockPlugin         sync.RWMutex
	lockGetConfigurationPlugin sync.RWMutex
	lockGetGasPlugin           sync.RWMutex
	lockGetPrecompilePlugin    sync.RWMutex
	lockGetStatePlugin         sync.RWMutex
}

// GetBlockPlugin calls GetBlockPluginFunc.
func (mock *StargazerHostChainMock) GetBlockPlugin() core.BlockPlugin {
	if mock.GetBlockPluginFunc == nil {
		panic("StargazerHostChainMock.GetBlockPluginFunc: method is nil but StargazerHostChain.GetBlockPlugin was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetBlockPlugin.Lock()
	mock.calls.GetBlockPlugin = append(mock.calls.GetBlockPlugin, callInfo)
	mock.lockGetBlockPlugin.Unlock()
	return mock.GetBlockPluginFunc()
}

// GetBlockPluginCalls gets all the calls that were made to GetBlockPlugin.
// Check the length with:
//
//	len(mockedStargazerHostChain.GetBlockPluginCalls())
func (mock *StargazerHostChainMock) GetBlockPluginCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetBlockPlugin.RLock()
	calls = mock.calls.GetBlockPlugin
	mock.lockGetBlockPlugin.RUnlock()
	return calls
}

// GetConfigurationPlugin calls GetConfigurationPluginFunc.
func (mock *StargazerHostChainMock) GetConfigurationPlugin() core.ConfigurationPlugin {
	if mock.GetConfigurationPluginFunc == nil {
		panic("StargazerHostChainMock.GetConfigurationPluginFunc: method is nil but StargazerHostChain.GetConfigurationPlugin was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetConfigurationPlugin.Lock()
	mock.calls.GetConfigurationPlugin = append(mock.calls.GetConfigurationPlugin, callInfo)
	mock.lockGetConfigurationPlugin.Unlock()
	return mock.GetConfigurationPluginFunc()
}

// GetConfigurationPluginCalls gets all the calls that were made to GetConfigurationPlugin.
// Check the length with:
//
//	len(mockedStargazerHostChain.GetConfigurationPluginCalls())
func (mock *StargazerHostChainMock) GetConfigurationPluginCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetConfigurationPlugin.RLock()
	calls = mock.calls.GetConfigurationPlugin
	mock.lockGetConfigurationPlugin.RUnlock()
	return calls
}

// GetGasPlugin calls GetGasPluginFunc.
func (mock *StargazerHostChainMock) GetGasPlugin() core.GasPlugin {
	if mock.GetGasPluginFunc == nil {
		panic("StargazerHostChainMock.GetGasPluginFunc: method is nil but StargazerHostChain.GetGasPlugin was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetGasPlugin.Lock()
	mock.calls.GetGasPlugin = append(mock.calls.GetGasPlugin, callInfo)
	mock.lockGetGasPlugin.Unlock()
	return mock.GetGasPluginFunc()
}

// GetGasPluginCalls gets all the calls that were made to GetGasPlugin.
// Check the length with:
//
//	len(mockedStargazerHostChain.GetGasPluginCalls())
func (mock *StargazerHostChainMock) GetGasPluginCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetGasPlugin.RLock()
	calls = mock.calls.GetGasPlugin
	mock.lockGetGasPlugin.RUnlock()
	return calls
}

// GetPrecompilePlugin calls GetPrecompilePluginFunc.
func (mock *StargazerHostChainMock) GetPrecompilePlugin() precompile.Plugin {
	if mock.GetPrecompilePluginFunc == nil {
		panic("StargazerHostChainMock.GetPrecompilePluginFunc: method is nil but StargazerHostChain.GetPrecompilePlugin was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetPrecompilePlugin.Lock()
	mock.calls.GetPrecompilePlugin = append(mock.calls.GetPrecompilePlugin, callInfo)
	mock.lockGetPrecompilePlugin.Unlock()
	return mock.GetPrecompilePluginFunc()
}

// GetPrecompilePluginCalls gets all the calls that were made to GetPrecompilePlugin.
// Check the length with:
//
//	len(mockedStargazerHostChain.GetPrecompilePluginCalls())
func (mock *StargazerHostChainMock) GetPrecompilePluginCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetPrecompilePlugin.RLock()
	calls = mock.calls.GetPrecompilePlugin
	mock.lockGetPrecompilePlugin.RUnlock()
	return calls
}

// GetStatePlugin calls GetStatePluginFunc.
func (mock *StargazerHostChainMock) GetStatePlugin() state.Plugin {
	if mock.GetStatePluginFunc == nil {
		panic("StargazerHostChainMock.GetStatePluginFunc: method is nil but StargazerHostChain.GetStatePlugin was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetStatePlugin.Lock()
	mock.calls.GetStatePlugin = append(mock.calls.GetStatePlugin, callInfo)
	mock.lockGetStatePlugin.Unlock()
	return mock.GetStatePluginFunc()
}

// GetStatePluginCalls gets all the calls that were made to GetStatePlugin.
// Check the length with:
//
//	len(mockedStargazerHostChain.GetStatePluginCalls())
func (mock *StargazerHostChainMock) GetStatePluginCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetStatePlugin.RLock()
	calls = mock.calls.GetStatePlugin
	mock.lockGetStatePlugin.RUnlock()
	return calls
}
