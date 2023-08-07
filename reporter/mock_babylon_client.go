// Code generated by MockGen. DO NOT EDIT.
// Source: reporter/expected_babylon_client.go

// Package reporter is a generated GoMock package.
package reporter

import (
	reflect "reflect"

	types "github.com/babylonchain/babylon/x/btccheckpoint/types"
	types0 "github.com/babylonchain/babylon/x/btclightclient/types"
	config "github.com/babylonchain/rpc-client/config"
	chainhash "github.com/btcsuite/btcd/chaincfg/chainhash"
	types1 "github.com/cosmos/cosmos-sdk/types"
	gomock "github.com/golang/mock/gomock"
)

// MockBabylonClient is a mock of BabylonClient interface.
type MockBabylonClient struct {
	ctrl     *gomock.Controller
	recorder *MockBabylonClientMockRecorder
}

// MockBabylonClientMockRecorder is the mock recorder for MockBabylonClient.
type MockBabylonClientMockRecorder struct {
	mock *MockBabylonClient
}

// NewMockBabylonClient creates a new mock instance.
func NewMockBabylonClient(ctrl *gomock.Controller) *MockBabylonClient {
	mock := &MockBabylonClient{ctrl: ctrl}
	mock.recorder = &MockBabylonClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBabylonClient) EXPECT() *MockBabylonClientMockRecorder {
	return m.recorder
}

// BTCBaseHeader mocks base method.
func (m *MockBabylonClient) BTCBaseHeader() (*types0.QueryBaseHeaderResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BTCBaseHeader")
	ret0, _ := ret[0].(*types0.QueryBaseHeaderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BTCBaseHeader indicates an expected call of BTCBaseHeader.
func (mr *MockBabylonClientMockRecorder) BTCBaseHeader() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BTCBaseHeader", reflect.TypeOf((*MockBabylonClient)(nil).BTCBaseHeader))
}

// BTCCheckpointParams mocks base method.
func (m *MockBabylonClient) BTCCheckpointParams() (*types.QueryParamsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BTCCheckpointParams")
	ret0, _ := ret[0].(*types.QueryParamsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BTCCheckpointParams indicates an expected call of BTCCheckpointParams.
func (mr *MockBabylonClientMockRecorder) BTCCheckpointParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BTCCheckpointParams", reflect.TypeOf((*MockBabylonClient)(nil).BTCCheckpointParams))
}

// BTCHeaderChainTip mocks base method.
func (m *MockBabylonClient) BTCHeaderChainTip() (*types0.QueryTipResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BTCHeaderChainTip")
	ret0, _ := ret[0].(*types0.QueryTipResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BTCHeaderChainTip indicates an expected call of BTCHeaderChainTip.
func (mr *MockBabylonClientMockRecorder) BTCHeaderChainTip() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BTCHeaderChainTip", reflect.TypeOf((*MockBabylonClient)(nil).BTCHeaderChainTip))
}

// ContainsBTCBlock mocks base method.
func (m *MockBabylonClient) ContainsBTCBlock(blockHash *chainhash.Hash) (*types0.QueryContainsBytesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainsBTCBlock", blockHash)
	ret0, _ := ret[0].(*types0.QueryContainsBytesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainsBTCBlock indicates an expected call of ContainsBTCBlock.
func (mr *MockBabylonClientMockRecorder) ContainsBTCBlock(blockHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainsBTCBlock", reflect.TypeOf((*MockBabylonClient)(nil).ContainsBTCBlock), blockHash)
}

// GetConfig mocks base method.
func (m *MockBabylonClient) GetConfig() *config.BabylonConfig {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConfig")
	ret0, _ := ret[0].(*config.BabylonConfig)
	return ret0
}

// GetConfig indicates an expected call of GetConfig.
func (mr *MockBabylonClientMockRecorder) GetConfig() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConfig", reflect.TypeOf((*MockBabylonClient)(nil).GetConfig))
}

// InsertBTCSpvProof mocks base method.
func (m *MockBabylonClient) InsertBTCSpvProof(msg *types.MsgInsertBTCSpvProof) (*types1.TxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertBTCSpvProof", msg)
	ret0, _ := ret[0].(*types1.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertBTCSpvProof indicates an expected call of InsertBTCSpvProof.
func (mr *MockBabylonClientMockRecorder) InsertBTCSpvProof(msg interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertBTCSpvProof", reflect.TypeOf((*MockBabylonClient)(nil).InsertBTCSpvProof), msg)
}

// InsertHeaders mocks base method.
func (m *MockBabylonClient) InsertHeaders(msgs []*types0.MsgInsertHeader) (*types1.TxResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertHeaders", msgs)
	ret0, _ := ret[0].(*types1.TxResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InsertHeaders indicates an expected call of InsertHeaders.
func (mr *MockBabylonClientMockRecorder) InsertHeaders(msgs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertHeaders", reflect.TypeOf((*MockBabylonClient)(nil).InsertHeaders), msgs)
}

// MustGetAddr mocks base method.
func (m *MockBabylonClient) MustGetAddr() types1.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MustGetAddr")
	ret0, _ := ret[0].(types1.AccAddress)
	return ret0
}

// MustGetAddr indicates an expected call of MustGetAddr.
func (mr *MockBabylonClientMockRecorder) MustGetAddr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MustGetAddr", reflect.TypeOf((*MockBabylonClient)(nil).MustGetAddr))
}

// Stop mocks base method.
func (m *MockBabylonClient) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockBabylonClientMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockBabylonClient)(nil).Stop))
}
