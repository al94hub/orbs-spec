// AUTO GENERATED FILE (by membufc proto compiler v0.0.20)
package services

import (
	"context"
	"fmt"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
)

/////////////////////////////////////////////////////////////////////////////
// service CrosschainConnector

type CrosschainConnector interface {
	EthereumCallContract(ctx context.Context, input *EthereumCallContractInput) (*EthereumCallContractOutput, error)
	EthereumGetTransactionLogs(ctx context.Context, input *EthereumGetTransactionLogsInput) (*EthereumGetTransactionLogsOutput, error)
}

/////////////////////////////////////////////////////////////////////////////
// message EthereumCallContractInput (non serializable)

type EthereumCallContractInput struct {
	ReferenceTimestamp           primitives.TimestampNano
	EthereumContractAddress      string
	EthereumFunctionName         string
	EthereumAbi                  string
	EthereumPackedInputArguments []byte
}

func (x *EthereumCallContractInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ReferenceTimestamp:%s,EthereumContractAddress:%s,EthereumFunctionName:%s,EthereumAbi:%s,EthereumPackedInputArguments:%s,}", x.StringReferenceTimestamp(), x.StringEthereumContractAddress(), x.StringEthereumFunctionName(), x.StringEthereumAbi(), x.StringEthereumPackedInputArguments())
}

func (x *EthereumCallContractInput) StringReferenceTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.ReferenceTimestamp)
	return
}

func (x *EthereumCallContractInput) StringEthereumContractAddress() (res string) {
	res = fmt.Sprintf(x.EthereumContractAddress)
	return
}

func (x *EthereumCallContractInput) StringEthereumFunctionName() (res string) {
	res = fmt.Sprintf(x.EthereumFunctionName)
	return
}

func (x *EthereumCallContractInput) StringEthereumAbi() (res string) {
	res = fmt.Sprintf(x.EthereumAbi)
	return
}

func (x *EthereumCallContractInput) StringEthereumPackedInputArguments() (res string) {
	res = fmt.Sprintf("%x", x.EthereumPackedInputArguments)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message EthereumCallContractOutput (non serializable)

type EthereumCallContractOutput struct {
	EthereumPackedOutput []byte
}

func (x *EthereumCallContractOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{EthereumPackedOutput:%s,}", x.StringEthereumPackedOutput())
}

func (x *EthereumCallContractOutput) StringEthereumPackedOutput() (res string) {
	res = fmt.Sprintf("%x", x.EthereumPackedOutput)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message EthereumGetTransactionLogsInput (non serializable)

type EthereumGetTransactionLogsInput struct {
	ReferenceTimestamp      primitives.TimestampNano
	EthereumTxhash          primitives.Uint256
	EthereumContractAddress string
	EventSignature          string
}

func (x *EthereumGetTransactionLogsInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ReferenceTimestamp:%s,EthereumTxhash:%s,EthereumContractAddress:%s,EventSignature:%s,}", x.StringReferenceTimestamp(), x.StringEthereumTxhash(), x.StringEthereumContractAddress(), x.StringEventSignature())
}

func (x *EthereumGetTransactionLogsInput) StringReferenceTimestamp() (res string) {
	res = fmt.Sprintf("%s", x.ReferenceTimestamp)
	return
}

func (x *EthereumGetTransactionLogsInput) StringEthereumTxhash() (res string) {
	res = fmt.Sprintf("%s", x.EthereumTxhash)
	return
}

func (x *EthereumGetTransactionLogsInput) StringEthereumContractAddress() (res string) {
	res = fmt.Sprintf(x.EthereumContractAddress)
	return
}

func (x *EthereumGetTransactionLogsInput) StringEventSignature() (res string) {
	res = fmt.Sprintf(x.EventSignature)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message EthereumGetTransactionLogsOutput (non serializable)

type EthereumGetTransactionLogsOutput struct {
	EthereumPackedEventTopics [][]byte
	EthereumPackedEventData   []byte
}

func (x *EthereumGetTransactionLogsOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{EthereumPackedEventTopics:%s,EthereumPackedEventData:%s,}", x.StringEthereumPackedEventTopics(), x.StringEthereumPackedEventData())
}

func (x *EthereumGetTransactionLogsOutput) StringEthereumPackedEventTopics() (res string) {
	res = "["
	for _, v := range x.EthereumPackedEventTopics {
		res += fmt.Sprintf("%x", v) + ","
	}
	res += "]"
	return
}

func (x *EthereumGetTransactionLogsOutput) StringEthereumPackedEventData() (res string) {
	res = fmt.Sprintf("%x", x.EthereumPackedEventData)
	return
}
