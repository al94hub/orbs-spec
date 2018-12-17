// AUTO GENERATED FILE (by membufc proto compiler v0.0.21)
package consensus

import (
	"bytes"
	"fmt"
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
)

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixBlockProof

// reader

type LeanHelixBlockProof struct {
	// BlockRef LeanHelixBlockRef
	// Nodes []LeanHelixSenderSignature
	// RandomSeedSignature primitives.Bls1Sig

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *LeanHelixBlockProof) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{BlockRef:%s,Nodes:%s,RandomSeedSignature:%s,}", x.StringBlockRef(), x.StringNodes(), x.StringRandomSeedSignature())
}

var _LeanHelixBlockProof_Scheme = []membuffers.FieldType{membuffers.TypeMessage, membuffers.TypeMessageArray, membuffers.TypeBytes}
var _LeanHelixBlockProof_Unions = [][]membuffers.FieldType{}

func LeanHelixBlockProofReader(buf []byte) *LeanHelixBlockProof {
	x := &LeanHelixBlockProof{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _LeanHelixBlockProof_Scheme, _LeanHelixBlockProof_Unions)
	return x
}

func (x *LeanHelixBlockProof) IsValid() bool {
	return x._message.IsValid()
}

func (x *LeanHelixBlockProof) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *LeanHelixBlockProof) Equal(y *LeanHelixBlockProof) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *LeanHelixBlockProof) BlockRef() *LeanHelixBlockRef {
	b, s := x._message.GetMessage(0)
	return LeanHelixBlockRefReader(b[:s])
}

func (x *LeanHelixBlockProof) RawBlockRef() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *LeanHelixBlockProof) RawBlockRefWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *LeanHelixBlockProof) StringBlockRef() string {
	return x.BlockRef().String()
}

func (x *LeanHelixBlockProof) NodesIterator() *LeanHelixBlockProofNodesIterator {
	return &LeanHelixBlockProofNodesIterator{iterator: x._message.GetMessageArrayIterator(1)}
}

type LeanHelixBlockProofNodesIterator struct {
	iterator *membuffers.Iterator
}

func (i *LeanHelixBlockProofNodesIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *LeanHelixBlockProofNodesIterator) NextNodes() *LeanHelixSenderSignature {
	b, s := i.iterator.NextMessage()
	return LeanHelixSenderSignatureReader(b[:s])
}

func (x *LeanHelixBlockProof) RawNodesArray() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *LeanHelixBlockProof) RawNodesArrayWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(1, 0)
}

func (x *LeanHelixBlockProof) StringNodes() (res string) {
	res = "["
	for i := x.NodesIterator(); i.HasNext(); {
		res += i.NextNodes().String() + ","
	}
	res += "]"
	return
}

func (x *LeanHelixBlockProof) RandomSeedSignature() primitives.Bls1Sig {
	return primitives.Bls1Sig(x._message.GetBytes(2))
}

func (x *LeanHelixBlockProof) RawRandomSeedSignature() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *LeanHelixBlockProof) RawRandomSeedSignatureWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(2, 0)
}

func (x *LeanHelixBlockProof) MutateRandomSeedSignature(v primitives.Bls1Sig) error {
	return x._message.SetBytes(2, []byte(v))
}

func (x *LeanHelixBlockProof) StringRandomSeedSignature() string {
	return fmt.Sprintf("%s", x.RandomSeedSignature())
}

// builder

type LeanHelixBlockProofBuilder struct {
	BlockRef            *LeanHelixBlockRefBuilder
	Nodes               []*LeanHelixSenderSignatureBuilder
	RandomSeedSignature primitives.Bls1Sig

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *LeanHelixBlockProofBuilder) arrayOfNodes() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.Nodes))
	for i, v := range w.Nodes {
		res[i] = v
	}
	return res
}

func (w *LeanHelixBlockProofBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	w._builder.NotifyBuildStart()
	defer w._builder.NotifyBuildEnd()
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	if w._overrideWithRawBuffer != nil {
		return w._builder.WriteOverrideWithRawBuffer(buf, w._overrideWithRawBuffer)
	}
	w._builder.Reset()
	err = w._builder.WriteMessage(buf, w.BlockRef)
	if err != nil {
		return
	}
	err = w._builder.WriteMessageArray(buf, w.arrayOfNodes())
	if err != nil {
		return
	}
	w._builder.WriteBytes(buf, []byte(w.RandomSeedSignature))
	return nil
}

func (w *LeanHelixBlockProofBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.HexDumpMessage(prefix, offsetFromStart, "LeanHelixBlockProof.BlockRef", w.BlockRef)
	if err != nil {
		return
	}
	err = w._builder.HexDumpMessageArray(prefix, offsetFromStart, "LeanHelixBlockProof.Nodes", w.arrayOfNodes())
	if err != nil {
		return
	}
	w._builder.HexDumpBytes(prefix, offsetFromStart, "LeanHelixBlockProof.RandomSeedSignature", []byte(w.RandomSeedSignature))
	return nil
}

func (w *LeanHelixBlockProofBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *LeanHelixBlockProofBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *LeanHelixBlockProofBuilder) Build() *LeanHelixBlockProof {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return LeanHelixBlockProofReader(buf)
}

func LeanHelixBlockProofBuilderFromRaw(raw []byte) *LeanHelixBlockProofBuilder {
	return &LeanHelixBlockProofBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixSenderSignature

// reader

type LeanHelixSenderSignature struct {
	// SenderNodeAddress primitives.NodeAddress
	// Signature primitives.EcdsaSecp256K1Sig

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *LeanHelixSenderSignature) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{SenderNodeAddress:%s,Signature:%s,}", x.StringSenderNodeAddress(), x.StringSignature())
}

var _LeanHelixSenderSignature_Scheme = []membuffers.FieldType{membuffers.TypeBytes, membuffers.TypeBytes}
var _LeanHelixSenderSignature_Unions = [][]membuffers.FieldType{}

func LeanHelixSenderSignatureReader(buf []byte) *LeanHelixSenderSignature {
	x := &LeanHelixSenderSignature{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _LeanHelixSenderSignature_Scheme, _LeanHelixSenderSignature_Unions)
	return x
}

func (x *LeanHelixSenderSignature) IsValid() bool {
	return x._message.IsValid()
}

func (x *LeanHelixSenderSignature) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *LeanHelixSenderSignature) Equal(y *LeanHelixSenderSignature) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *LeanHelixSenderSignature) SenderNodeAddress() primitives.NodeAddress {
	return primitives.NodeAddress(x._message.GetBytes(0))
}

func (x *LeanHelixSenderSignature) RawSenderNodeAddress() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *LeanHelixSenderSignature) RawSenderNodeAddressWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *LeanHelixSenderSignature) MutateSenderNodeAddress(v primitives.NodeAddress) error {
	return x._message.SetBytes(0, []byte(v))
}

func (x *LeanHelixSenderSignature) StringSenderNodeAddress() string {
	return fmt.Sprintf("%s", x.SenderNodeAddress())
}

func (x *LeanHelixSenderSignature) Signature() primitives.EcdsaSecp256K1Sig {
	return primitives.EcdsaSecp256K1Sig(x._message.GetBytes(1))
}

func (x *LeanHelixSenderSignature) RawSignature() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *LeanHelixSenderSignature) RawSignatureWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(1, 0)
}

func (x *LeanHelixSenderSignature) MutateSignature(v primitives.EcdsaSecp256K1Sig) error {
	return x._message.SetBytes(1, []byte(v))
}

func (x *LeanHelixSenderSignature) StringSignature() string {
	return fmt.Sprintf("%s", x.Signature())
}

// builder

type LeanHelixSenderSignatureBuilder struct {
	SenderNodeAddress primitives.NodeAddress
	Signature         primitives.EcdsaSecp256K1Sig

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *LeanHelixSenderSignatureBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	w._builder.NotifyBuildStart()
	defer w._builder.NotifyBuildEnd()
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	if w._overrideWithRawBuffer != nil {
		return w._builder.WriteOverrideWithRawBuffer(buf, w._overrideWithRawBuffer)
	}
	w._builder.Reset()
	w._builder.WriteBytes(buf, []byte(w.SenderNodeAddress))
	w._builder.WriteBytes(buf, []byte(w.Signature))
	return nil
}

func (w *LeanHelixSenderSignatureBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.HexDumpBytes(prefix, offsetFromStart, "LeanHelixSenderSignature.SenderNodeAddress", []byte(w.SenderNodeAddress))
	w._builder.HexDumpBytes(prefix, offsetFromStart, "LeanHelixSenderSignature.Signature", []byte(w.Signature))
	return nil
}

func (w *LeanHelixSenderSignatureBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *LeanHelixSenderSignatureBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *LeanHelixSenderSignatureBuilder) Build() *LeanHelixSenderSignature {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return LeanHelixSenderSignatureReader(buf)
}

func LeanHelixSenderSignatureBuilderFromRaw(raw []byte) *LeanHelixSenderSignatureBuilder {
	return &LeanHelixSenderSignatureBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message LeanHelixBlockRef

// reader

type LeanHelixBlockRef struct {
	// MessageType LeanHelixMessageType
	// BlockHeight primitives.BlockHeight
	// View uint32
	// BlockHash primitives.Uint256

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *LeanHelixBlockRef) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{MessageType:%s,BlockHeight:%s,View:%s,BlockHash:%s,}", x.StringMessageType(), x.StringBlockHeight(), x.StringView(), x.StringBlockHash())
}

var _LeanHelixBlockRef_Scheme = []membuffers.FieldType{membuffers.TypeUint16, membuffers.TypeUint64, membuffers.TypeUint32, membuffers.TypeBytes}
var _LeanHelixBlockRef_Unions = [][]membuffers.FieldType{}

func LeanHelixBlockRefReader(buf []byte) *LeanHelixBlockRef {
	x := &LeanHelixBlockRef{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _LeanHelixBlockRef_Scheme, _LeanHelixBlockRef_Unions)
	return x
}

func (x *LeanHelixBlockRef) IsValid() bool {
	return x._message.IsValid()
}

func (x *LeanHelixBlockRef) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *LeanHelixBlockRef) Equal(y *LeanHelixBlockRef) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *LeanHelixBlockRef) MessageType() LeanHelixMessageType {
	return LeanHelixMessageType(x._message.GetUint16(0))
}

func (x *LeanHelixBlockRef) RawMessageType() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *LeanHelixBlockRef) MutateMessageType(v LeanHelixMessageType) error {
	return x._message.SetUint16(0, uint16(v))
}

func (x *LeanHelixBlockRef) StringMessageType() string {
	return x.MessageType().String()
}

func (x *LeanHelixBlockRef) BlockHeight() primitives.BlockHeight {
	return primitives.BlockHeight(x._message.GetUint64(1))
}

func (x *LeanHelixBlockRef) RawBlockHeight() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *LeanHelixBlockRef) MutateBlockHeight(v primitives.BlockHeight) error {
	return x._message.SetUint64(1, uint64(v))
}

func (x *LeanHelixBlockRef) StringBlockHeight() string {
	return fmt.Sprintf("%s", x.BlockHeight())
}

func (x *LeanHelixBlockRef) View() uint32 {
	return x._message.GetUint32(2)
}

func (x *LeanHelixBlockRef) RawView() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *LeanHelixBlockRef) MutateView(v uint32) error {
	return x._message.SetUint32(2, v)
}

func (x *LeanHelixBlockRef) StringView() string {
	return fmt.Sprintf("%x", x.View())
}

func (x *LeanHelixBlockRef) BlockHash() primitives.Uint256 {
	return primitives.Uint256(x._message.GetBytes(3))
}

func (x *LeanHelixBlockRef) RawBlockHash() []byte {
	return x._message.RawBufferForField(3, 0)
}

func (x *LeanHelixBlockRef) RawBlockHashWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(3, 0)
}

func (x *LeanHelixBlockRef) MutateBlockHash(v primitives.Uint256) error {
	return x._message.SetBytes(3, []byte(v))
}

func (x *LeanHelixBlockRef) StringBlockHash() string {
	return fmt.Sprintf("%s", x.BlockHash())
}

// builder

type LeanHelixBlockRefBuilder struct {
	MessageType LeanHelixMessageType
	BlockHeight primitives.BlockHeight
	View        uint32
	BlockHash   primitives.Uint256

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *LeanHelixBlockRefBuilder) Write(buf []byte) (err error) {
	if w == nil {
		return
	}
	w._builder.NotifyBuildStart()
	defer w._builder.NotifyBuildEnd()
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	if w._overrideWithRawBuffer != nil {
		return w._builder.WriteOverrideWithRawBuffer(buf, w._overrideWithRawBuffer)
	}
	w._builder.Reset()
	w._builder.WriteUint16(buf, uint16(w.MessageType))
	w._builder.WriteUint64(buf, uint64(w.BlockHeight))
	w._builder.WriteUint32(buf, w.View)
	w._builder.WriteBytes(buf, []byte(w.BlockHash))
	return nil
}

func (w *LeanHelixBlockRefBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.HexDumpUint16(prefix, offsetFromStart, "LeanHelixBlockRef.MessageType", uint16(w.MessageType))
	w._builder.HexDumpUint64(prefix, offsetFromStart, "LeanHelixBlockRef.BlockHeight", uint64(w.BlockHeight))
	w._builder.HexDumpUint32(prefix, offsetFromStart, "LeanHelixBlockRef.View", w.View)
	w._builder.HexDumpBytes(prefix, offsetFromStart, "LeanHelixBlockRef.BlockHash", []byte(w.BlockHash))
	return nil
}

func (w *LeanHelixBlockRefBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *LeanHelixBlockRefBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *LeanHelixBlockRefBuilder) Build() *LeanHelixBlockRef {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return LeanHelixBlockRefReader(buf)
}

func LeanHelixBlockRefBuilderFromRaw(raw []byte) *LeanHelixBlockRefBuilder {
	return &LeanHelixBlockRefBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// enums

type LeanHelixMessageType uint16

const ()

func (n LeanHelixMessageType) String() string {
	switch n {
	}
	return "UNKNOWN"
}
