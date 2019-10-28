// AUTO GENERATED FILE (by membufc proto compiler v0.0.32)
package services

import (
	"context"
	"github.com/orbs-network/go-mock"
)

/////////////////////////////////////////////////////////////////////////////
// service EthereumConnector

type MockEthereumConnector struct {
	mock.Mock
}

func (s *MockEthereumConnector) EthereumCallContract(ctx context.Context, input *EthereumCallContractInput) (*EthereumCallContractOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*EthereumCallContractOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockEthereumConnector) EthereumGetTransactionLogs(ctx context.Context, input *EthereumGetTransactionLogsInput) (*EthereumGetTransactionLogsOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*EthereumGetTransactionLogsOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockEthereumConnector) EthereumGetBlockNumber(ctx context.Context, input *EthereumGetBlockNumberInput) (*EthereumGetBlockNumberOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*EthereumGetBlockNumberOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockEthereumConnector) EthereumGetBlockTime(ctx context.Context, input *EthereumGetBlockTimeInput) (*EthereumGetBlockTimeOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*EthereumGetBlockTimeOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockEthereumConnector) EthereumGetBlockTimeByNumber(ctx context.Context, input *EthereumGetBlockTimeByNumberInput) (*EthereumGetBlockTimeByNumberOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*EthereumGetBlockTimeByNumberOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}

func (s *MockEthereumConnector) EthereumGetBlockNumberByTime(ctx context.Context, input *EthereumGetBlockNumberByTimeInput) (*EthereumGetBlockNumberByTimeOutput, error) {
	ret := s.Called(ctx, input)
	if out := ret.Get(0); out != nil {
		return out.(*EthereumGetBlockNumberByTimeOutput), ret.Error(1)
	} else {
		return nil, ret.Error(1)
	}
}
