// Code generated by MockGen. DO NOT EDIT.
// Source: monitor/expected_babylon_client.go

// Package monitor is a generated GoMock package.
package monitor

import (
	reflect "reflect"

	types "github.com/babylonchain/babylon/x/btccheckpoint/types"
	types0 "github.com/babylonchain/babylon/x/btclightclient/types"
	types1 "github.com/babylonchain/babylon/x/btcstaking/types"
	types2 "github.com/babylonchain/babylon/x/checkpointing/types"
	types3 "github.com/babylonchain/babylon/x/epoching/types"
	types4 "github.com/babylonchain/babylon/x/monitor/types"
	chainhash "github.com/btcsuite/btcd/chaincfg/chainhash"
	coretypes "github.com/cometbft/cometbft/rpc/core/types"
	query "github.com/cosmos/cosmos-sdk/types/query"
	gomock "github.com/golang/mock/gomock"
)

// MockBabylonQueryClient is a mock of BabylonQueryClient interface.
type MockBabylonQueryClient struct {
	ctrl     *gomock.Controller
	recorder *MockBabylonQueryClientMockRecorder
}

// MockBabylonQueryClientMockRecorder is the mock recorder for MockBabylonQueryClient.
type MockBabylonQueryClientMockRecorder struct {
	mock *MockBabylonQueryClient
}

// NewMockBabylonQueryClient creates a new mock instance.
func NewMockBabylonQueryClient(ctrl *gomock.Controller) *MockBabylonQueryClient {
	mock := &MockBabylonQueryClient{ctrl: ctrl}
	mock.recorder = &MockBabylonQueryClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBabylonQueryClient) EXPECT() *MockBabylonQueryClientMockRecorder {
	return m.recorder
}

// BTCCheckpointParams mocks base method.
func (m *MockBabylonQueryClient) BTCCheckpointParams() (*types.QueryParamsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BTCCheckpointParams")
	ret0, _ := ret[0].(*types.QueryParamsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BTCCheckpointParams indicates an expected call of BTCCheckpointParams.
func (mr *MockBabylonQueryClientMockRecorder) BTCCheckpointParams() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BTCCheckpointParams", reflect.TypeOf((*MockBabylonQueryClient)(nil).BTCCheckpointParams))
}

// BTCHeaderChainTip mocks base method.
func (m *MockBabylonQueryClient) BTCHeaderChainTip() (*types0.QueryTipResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BTCHeaderChainTip")
	ret0, _ := ret[0].(*types0.QueryTipResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BTCHeaderChainTip indicates an expected call of BTCHeaderChainTip.
func (mr *MockBabylonQueryClientMockRecorder) BTCHeaderChainTip() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BTCHeaderChainTip", reflect.TypeOf((*MockBabylonQueryClient)(nil).BTCHeaderChainTip))
}

// BTCValidatorDelegations mocks base method.
func (m *MockBabylonQueryClient) BTCValidatorDelegations(valBtcPkHex string, pagination *query.PageRequest) (*types1.QueryBTCValidatorDelegationsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BTCValidatorDelegations", valBtcPkHex, pagination)
	ret0, _ := ret[0].(*types1.QueryBTCValidatorDelegationsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BTCValidatorDelegations indicates an expected call of BTCValidatorDelegations.
func (mr *MockBabylonQueryClientMockRecorder) BTCValidatorDelegations(valBtcPkHex, pagination interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BTCValidatorDelegations", reflect.TypeOf((*MockBabylonQueryClient)(nil).BTCValidatorDelegations), valBtcPkHex, pagination)
}

// BlsPublicKeyList mocks base method.
func (m *MockBabylonQueryClient) BlsPublicKeyList(epochNumber uint64, pagination *query.PageRequest) (*types2.QueryBlsPublicKeyListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BlsPublicKeyList", epochNumber, pagination)
	ret0, _ := ret[0].(*types2.QueryBlsPublicKeyListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BlsPublicKeyList indicates an expected call of BlsPublicKeyList.
func (mr *MockBabylonQueryClientMockRecorder) BlsPublicKeyList(epochNumber, pagination interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BlsPublicKeyList", reflect.TypeOf((*MockBabylonQueryClient)(nil).BlsPublicKeyList), epochNumber, pagination)
}

// ContainsBTCBlock mocks base method.
func (m *MockBabylonQueryClient) ContainsBTCBlock(blockHash *chainhash.Hash) (*types0.QueryContainsBytesResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ContainsBTCBlock", blockHash)
	ret0, _ := ret[0].(*types0.QueryContainsBytesResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainsBTCBlock indicates an expected call of ContainsBTCBlock.
func (mr *MockBabylonQueryClientMockRecorder) ContainsBTCBlock(blockHash interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainsBTCBlock", reflect.TypeOf((*MockBabylonQueryClient)(nil).ContainsBTCBlock), blockHash)
}

// CurrentEpoch mocks base method.
func (m *MockBabylonQueryClient) CurrentEpoch() (*types3.QueryCurrentEpochResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CurrentEpoch")
	ret0, _ := ret[0].(*types3.QueryCurrentEpochResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CurrentEpoch indicates an expected call of CurrentEpoch.
func (mr *MockBabylonQueryClientMockRecorder) CurrentEpoch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CurrentEpoch", reflect.TypeOf((*MockBabylonQueryClient)(nil).CurrentEpoch))
}

// EndedEpochBTCHeight mocks base method.
func (m *MockBabylonQueryClient) EndedEpochBTCHeight(epochNum uint64) (*types4.QueryEndedEpochBtcHeightResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EndedEpochBTCHeight", epochNum)
	ret0, _ := ret[0].(*types4.QueryEndedEpochBtcHeightResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EndedEpochBTCHeight indicates an expected call of EndedEpochBTCHeight.
func (mr *MockBabylonQueryClientMockRecorder) EndedEpochBTCHeight(epochNum interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndedEpochBTCHeight", reflect.TypeOf((*MockBabylonQueryClient)(nil).EndedEpochBTCHeight), epochNum)
}

// RawCheckpoint mocks base method.
func (m *MockBabylonQueryClient) RawCheckpoint(epochNumber uint64) (*types2.QueryRawCheckpointResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RawCheckpoint", epochNumber)
	ret0, _ := ret[0].(*types2.QueryRawCheckpointResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RawCheckpoint indicates an expected call of RawCheckpoint.
func (mr *MockBabylonQueryClientMockRecorder) RawCheckpoint(epochNumber interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RawCheckpoint", reflect.TypeOf((*MockBabylonQueryClient)(nil).RawCheckpoint), epochNumber)
}

// ReportedCheckpointBTCHeight mocks base method.
func (m *MockBabylonQueryClient) ReportedCheckpointBTCHeight(hashStr string) (*types4.QueryReportedCheckpointBtcHeightResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReportedCheckpointBTCHeight", hashStr)
	ret0, _ := ret[0].(*types4.QueryReportedCheckpointBtcHeightResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ReportedCheckpointBTCHeight indicates an expected call of ReportedCheckpointBTCHeight.
func (mr *MockBabylonQueryClientMockRecorder) ReportedCheckpointBTCHeight(hashStr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReportedCheckpointBTCHeight", reflect.TypeOf((*MockBabylonQueryClient)(nil).ReportedCheckpointBTCHeight), hashStr)
}

// Subscribe mocks base method.
func (m *MockBabylonQueryClient) Subscribe(subscriber, query string, outCapacity ...int) (<-chan coretypes.ResultEvent, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{subscriber, query}
	for _, a := range outCapacity {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Subscribe", varargs...)
	ret0, _ := ret[0].(<-chan coretypes.ResultEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockBabylonQueryClientMockRecorder) Subscribe(subscriber, query interface{}, outCapacity ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{subscriber, query}, outCapacity...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockBabylonQueryClient)(nil).Subscribe), varargs...)
}

// Unsubscribe mocks base method.
func (m *MockBabylonQueryClient) Unsubscribe(subscriber, query string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unsubscribe", subscriber, query)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unsubscribe indicates an expected call of Unsubscribe.
func (mr *MockBabylonQueryClientMockRecorder) Unsubscribe(subscriber, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockBabylonQueryClient)(nil).Unsubscribe), subscriber, query)
}
