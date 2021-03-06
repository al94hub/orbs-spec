// AUTO GENERATED FILE (by membufc proto compiler v0.4.0)
package services

import (
	"context"
	"fmt"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"github.com/orbs-network/orbs-spec/types/go/protocol"
	"github.com/orbs-network/orbs-spec/types/go/services/handlers"
)

/////////////////////////////////////////////////////////////////////////////
// service Processor

type Processor interface {
	ProcessCall(ctx context.Context, input *ProcessCallInput) (*ProcessCallOutput, error)
	GetContractInfo(ctx context.Context, input *GetContractInfoInput) (*GetContractInfoOutput, error)
	RegisterContractSdkCallHandler(handler handlers.ContractSdkCallHandler)
}

/////////////////////////////////////////////////////////////////////////////
// message ProcessCallInput (non serializable)

type ProcessCallInput struct {
	ContextId              primitives.ExecutionContextId
	ContractName           primitives.ContractName
	MethodName             primitives.MethodName
	InputArgumentArray     *protocol.ArgumentArray
	AccessScope            protocol.ExecutionAccessScope
	CallingPermissionScope protocol.ExecutionPermissionScope
}

func (x *ProcessCallInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ContextId:%s,ContractName:%s,MethodName:%s,InputArgumentArray:%s,AccessScope:%s,CallingPermissionScope:%s,}", x.StringContextId(), x.StringContractName(), x.StringMethodName(), x.StringInputArgumentArray(), x.StringAccessScope(), x.StringCallingPermissionScope())
}

func (x *ProcessCallInput) StringContextId() (res string) {
	res = fmt.Sprintf("%s", x.ContextId)
	return
}

func (x *ProcessCallInput) StringContractName() (res string) {
	res = fmt.Sprintf("%s", x.ContractName)
	return
}

func (x *ProcessCallInput) StringMethodName() (res string) {
	res = fmt.Sprintf("%s", x.MethodName)
	return
}

func (x *ProcessCallInput) StringInputArgumentArray() (res string) {
	res = x.InputArgumentArray.String()
	return
}

func (x *ProcessCallInput) StringAccessScope() (res string) {
	res = fmt.Sprintf("%x", x.AccessScope)
	return
}

func (x *ProcessCallInput) StringCallingPermissionScope() (res string) {
	res = fmt.Sprintf("%x", x.CallingPermissionScope)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message ProcessCallOutput (non serializable)

type ProcessCallOutput struct {
	OutputArgumentArray *protocol.ArgumentArray
	CallResult          protocol.ExecutionResult
}

func (x *ProcessCallOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{OutputArgumentArray:%s,CallResult:%s,}", x.StringOutputArgumentArray(), x.StringCallResult())
}

func (x *ProcessCallOutput) StringOutputArgumentArray() (res string) {
	res = x.OutputArgumentArray.String()
	return
}

func (x *ProcessCallOutput) StringCallResult() (res string) {
	res = fmt.Sprintf("%x", x.CallResult)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetContractInfoInput (non serializable)

type GetContractInfoInput struct {
	ContextId    primitives.ExecutionContextId
	ContractName primitives.ContractName
}

func (x *GetContractInfoInput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ContextId:%s,ContractName:%s,}", x.StringContextId(), x.StringContractName())
}

func (x *GetContractInfoInput) StringContextId() (res string) {
	res = fmt.Sprintf("%s", x.ContextId)
	return
}

func (x *GetContractInfoInput) StringContractName() (res string) {
	res = fmt.Sprintf("%s", x.ContractName)
	return
}

/////////////////////////////////////////////////////////////////////////////
// message GetContractInfoOutput (non serializable)

type GetContractInfoOutput struct {
	PermissionScope protocol.ExecutionPermissionScope
}

func (x *GetContractInfoOutput) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{PermissionScope:%s,}", x.StringPermissionScope())
}

func (x *GetContractInfoOutput) StringPermissionScope() (res string) {
	res = fmt.Sprintf("%x", x.PermissionScope)
	return
}

/////////////////////////////////////////////////////////////////////////////
// enums
