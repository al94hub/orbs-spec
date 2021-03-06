// AUTO GENERATED FILE (by membufc proto compiler v0.4.0)
package protocol

import (
	"bytes"
	"fmt"
	"github.com/orbs-network/membuffers/go"
	"github.com/orbs-network/orbs-spec/types/go/primitives"
	"math/big"
)

/////////////////////////////////////////////////////////////////////////////
// message Argument

// reader

type Argument struct {
	// Type ArgumentType

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *Argument) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Type:%s,}", x.StringType())
}

var _Argument_Scheme = []membuffers.FieldType{membuffers.TypeUnion}
var _Argument_Unions = [][]membuffers.FieldType{{membuffers.TypeUint32, membuffers.TypeUint64, membuffers.TypeString, membuffers.TypeBytes, membuffers.TypeBool, membuffers.TypeUint256, membuffers.TypeBytes20, membuffers.TypeBytes32, membuffers.TypeUint32Array, membuffers.TypeUint64Array, membuffers.TypeStringArray, membuffers.TypeBytesArray, membuffers.TypeBoolArray, membuffers.TypeUint256Array, membuffers.TypeBytes20Array, membuffers.TypeBytes32Array}}

func ArgumentReader(buf []byte) *Argument {
	x := &Argument{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _Argument_Scheme, _Argument_Unions)
	return x
}

func (x *Argument) IsValid() bool {
	return x._message.IsValid()
}

func (x *Argument) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *Argument) Equal(y *Argument) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

type ArgumentType uint16

const (
	ARGUMENT_TYPE_UINT_32_VALUE        ArgumentType = 0
	ARGUMENT_TYPE_UINT_64_VALUE        ArgumentType = 1
	ARGUMENT_TYPE_STRING_VALUE         ArgumentType = 2
	ARGUMENT_TYPE_BYTES_VALUE          ArgumentType = 3
	ARGUMENT_TYPE_BOOL_VALUE           ArgumentType = 4
	ARGUMENT_TYPE_UINT_256_VALUE       ArgumentType = 5
	ARGUMENT_TYPE_BYTES_20_VALUE       ArgumentType = 6
	ARGUMENT_TYPE_BYTES_32_VALUE       ArgumentType = 7
	ARGUMENT_TYPE_UINT_32_ARRAY_VALUE  ArgumentType = 8
	ARGUMENT_TYPE_UINT_64_ARRAY_VALUE  ArgumentType = 9
	ARGUMENT_TYPE_STRING_ARRAY_VALUE   ArgumentType = 10
	ARGUMENT_TYPE_BYTES_ARRAY_VALUE    ArgumentType = 11
	ARGUMENT_TYPE_BOOL_ARRAY_VALUE     ArgumentType = 12
	ARGUMENT_TYPE_UINT_256_ARRAY_VALUE ArgumentType = 13
	ARGUMENT_TYPE_BYTES_20_ARRAY_VALUE ArgumentType = 14
	ARGUMENT_TYPE_BYTES_32_ARRAY_VALUE ArgumentType = 15
)

func (x *Argument) Type() ArgumentType {
	return ArgumentType(x._message.GetUnionIndex(0, 0))
}

func (x *Argument) IsTypeUint32Value() bool {
	is, _ := x._message.IsUnionIndex(0, 0, 0)
	return is
}

func (x *Argument) Uint32Value() uint32 {
	is, off := x._message.IsUnionIndex(0, 0, 0)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	return x._message.GetUint32InOffset(off)
}

func (x *Argument) StringUint32Value() string {
	return fmt.Sprintf("%v", x.Uint32Value())
}

func (x *Argument) MutateUint32Value(v uint32) error {
	is, off := x._message.IsUnionIndex(0, 0, 0)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetUint32InOffset(off, v)
	return nil
}

func (x *Argument) IsTypeUint64Value() bool {
	is, _ := x._message.IsUnionIndex(0, 0, 1)
	return is
}

func (x *Argument) Uint64Value() uint64 {
	is, off := x._message.IsUnionIndex(0, 0, 1)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	return x._message.GetUint64InOffset(off)
}

func (x *Argument) StringUint64Value() string {
	return fmt.Sprintf("%v", x.Uint64Value())
}

func (x *Argument) MutateUint64Value(v uint64) error {
	is, off := x._message.IsUnionIndex(0, 0, 1)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetUint64InOffset(off, v)
	return nil
}

func (x *Argument) IsTypeStringValue() bool {
	is, _ := x._message.IsUnionIndex(0, 0, 2)
	return is
}

func (x *Argument) StringValue() string {
	is, off := x._message.IsUnionIndex(0, 0, 2)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	return x._message.GetStringInOffset(off)
}

func (x *Argument) StringStringValue() string {
	return fmt.Sprintf("%s", x.StringValue())
}

func (x *Argument) MutateStringValue(v string) error {
	is, off := x._message.IsUnionIndex(0, 0, 2)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetStringInOffset(off, v)
	return nil
}

func (x *Argument) IsTypeBytesValue() bool {
	is, _ := x._message.IsUnionIndex(0, 0, 3)
	return is
}

func (x *Argument) BytesValue() []byte {
	is, off := x._message.IsUnionIndex(0, 0, 3)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	return x._message.GetBytesInOffset(off)
}

func (x *Argument) StringBytesValue() string {
	return fmt.Sprintf("%x", x.BytesValue())
}

func (x *Argument) MutateBytesValue(v []byte) error {
	is, off := x._message.IsUnionIndex(0, 0, 3)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetBytesInOffset(off, v)
	return nil
}

func (x *Argument) IsTypeBoolValue() bool {
	is, _ := x._message.IsUnionIndex(0, 0, 4)
	return is
}

func (x *Argument) BoolValue() bool {
	is, off := x._message.IsUnionIndex(0, 0, 4)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	return x._message.GetBoolInOffset(off)
}

func (x *Argument) StringBoolValue() string {
	return fmt.Sprintf("%v", x.BoolValue())
}

func (x *Argument) MutateBoolValue(v bool) error {
	is, off := x._message.IsUnionIndex(0, 0, 4)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetBoolInOffset(off, v)
	return nil
}

func (x *Argument) IsTypeUint256Value() bool {
	is, _ := x._message.IsUnionIndex(0, 0, 5)
	return is
}

func (x *Argument) Uint256Value() *big.Int {
	is, off := x._message.IsUnionIndex(0, 0, 5)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	return x._message.GetUint256InOffset(off)
}

func (x *Argument) StringUint256Value() string {
	return fmt.Sprintf("%v", x.Uint256Value())
}

func (x *Argument) MutateUint256Value(v *big.Int) error {
	is, off := x._message.IsUnionIndex(0, 0, 5)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetUint256InOffset(off, v)
	return nil
}

func (x *Argument) IsTypeBytes20Value() bool {
	is, _ := x._message.IsUnionIndex(0, 0, 6)
	return is
}

func (x *Argument) Bytes20Value() [20]byte {
	is, off := x._message.IsUnionIndex(0, 0, 6)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	return x._message.GetBytes20InOffset(off)
}

func (x *Argument) StringBytes20Value() string {
	return fmt.Sprintf("%x", x.Bytes20Value())
}

func (x *Argument) MutateBytes20Value(v [20]byte) error {
	is, off := x._message.IsUnionIndex(0, 0, 6)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetBytes20InOffset(off, v)
	return nil
}

func (x *Argument) IsTypeBytes32Value() bool {
	is, _ := x._message.IsUnionIndex(0, 0, 7)
	return is
}

func (x *Argument) Bytes32Value() [32]byte {
	is, off := x._message.IsUnionIndex(0, 0, 7)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	return x._message.GetBytes32InOffset(off)
}

func (x *Argument) StringBytes32Value() string {
	return fmt.Sprintf("%x", x.Bytes32Value())
}

func (x *Argument) MutateBytes32Value(v [32]byte) error {
	is, off := x._message.IsUnionIndex(0, 0, 7)
	if !is {
		return &membuffers.ErrInvalidField{}
	}
	x._message.SetBytes32InOffset(off, v)
	return nil
}

func (x *Argument) IsTypeUint32ArrayValue() bool {
	is, _ := x._message.IsUnionIndex(0, 0, 8)
	return is
}

func (x *Argument) Uint32ArrayValueCopiedToNative() []uint32 {
	is, off := x._message.IsUnionIndex(0, 0, 8)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	var res []uint32
	itr := x._message.GetUint32ArrayIteratorInOffset(off)
	for itr.HasNext() {
		res = append(res, itr.NextUint32())
	}
	return res
}

func (x *Argument) Uint32ArrayValueIterator() *membuffers.Iterator {
	is, off := x._message.IsUnionIndex(0, 0, 8)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	return x._message.GetUint32ArrayIteratorInOffset(off)
}

func (x *Argument) StringUint32ArrayValue() string {
	res := "["
	itr := x.Uint32ArrayValueIterator()
	for itr.HasNext() {
		res += fmt.Sprintf("%v", itr.NextUint32()) + ","
	}
	res += "]"
	return res
}

func (x *Argument) IsTypeUint64ArrayValue() bool {
	is, _ := x._message.IsUnionIndex(0, 0, 9)
	return is
}

func (x *Argument) Uint64ArrayValueCopiedToNative() []uint64 {
	is, off := x._message.IsUnionIndex(0, 0, 9)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	var res []uint64
	itr := x._message.GetUint64ArrayIteratorInOffset(off)
	for itr.HasNext() {
		res = append(res, itr.NextUint64())
	}
	return res
}

func (x *Argument) Uint64ArrayValueIterator() *membuffers.Iterator {
	is, off := x._message.IsUnionIndex(0, 0, 9)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	return x._message.GetUint64ArrayIteratorInOffset(off)
}

func (x *Argument) StringUint64ArrayValue() string {
	res := "["
	itr := x.Uint64ArrayValueIterator()
	for itr.HasNext() {
		res += fmt.Sprintf("%v", itr.NextUint64()) + ","
	}
	res += "]"
	return res
}

func (x *Argument) IsTypeStringArrayValue() bool {
	is, _ := x._message.IsUnionIndex(0, 0, 10)
	return is
}

func (x *Argument) StringArrayValueCopiedToNative() []string {
	is, off := x._message.IsUnionIndex(0, 0, 10)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	var res []string
	itr := x._message.GetStringArrayIteratorInOffset(off)
	for itr.HasNext() {
		res = append(res, itr.NextString())
	}
	return res
}

func (x *Argument) StringArrayValueIterator() *membuffers.Iterator {
	is, off := x._message.IsUnionIndex(0, 0, 10)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	return x._message.GetStringArrayIteratorInOffset(off)
}

func (x *Argument) StringStringArrayValue() string {
	res := "["
	itr := x.StringArrayValueIterator()
	for itr.HasNext() {
		res += fmt.Sprintf("%s", itr.NextString()) + ","
	}
	res += "]"
	return res
}

func (x *Argument) IsTypeBytesArrayValue() bool {
	is, _ := x._message.IsUnionIndex(0, 0, 11)
	return is
}

func (x *Argument) BytesArrayValueCopiedToNative() [][]byte {
	is, off := x._message.IsUnionIndex(0, 0, 11)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	var res [][]byte
	itr := x._message.GetBytesArrayIteratorInOffset(off)
	for itr.HasNext() {
		res = append(res, itr.NextBytes())
	}
	return res
}

func (x *Argument) BytesArrayValueIterator() *membuffers.Iterator {
	is, off := x._message.IsUnionIndex(0, 0, 11)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	return x._message.GetBytesArrayIteratorInOffset(off)
}

func (x *Argument) StringBytesArrayValue() string {
	res := "["
	itr := x.BytesArrayValueIterator()
	for itr.HasNext() {
		res += fmt.Sprintf("%x", itr.NextBytes()) + ","
	}
	res += "]"
	return res
}

func (x *Argument) IsTypeBoolArrayValue() bool {
	is, _ := x._message.IsUnionIndex(0, 0, 12)
	return is
}

func (x *Argument) BoolArrayValueCopiedToNative() []bool {
	is, off := x._message.IsUnionIndex(0, 0, 12)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	var res []bool
	itr := x._message.GetBoolArrayIteratorInOffset(off)
	for itr.HasNext() {
		res = append(res, itr.NextBool())
	}
	return res
}

func (x *Argument) BoolArrayValueIterator() *membuffers.Iterator {
	is, off := x._message.IsUnionIndex(0, 0, 12)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	return x._message.GetBoolArrayIteratorInOffset(off)
}

func (x *Argument) StringBoolArrayValue() string {
	res := "["
	itr := x.BoolArrayValueIterator()
	for itr.HasNext() {
		res += fmt.Sprintf("%v", itr.NextBool()) + ","
	}
	res += "]"
	return res
}

func (x *Argument) IsTypeUint256ArrayValue() bool {
	is, _ := x._message.IsUnionIndex(0, 0, 13)
	return is
}

func (x *Argument) Uint256ArrayValueCopiedToNative() []*big.Int {
	is, off := x._message.IsUnionIndex(0, 0, 13)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	var res []*big.Int
	itr := x._message.GetUint256ArrayIteratorInOffset(off)
	for itr.HasNext() {
		res = append(res, itr.NextUint256())
	}
	return res
}

func (x *Argument) Uint256ArrayValueIterator() *membuffers.Iterator {
	is, off := x._message.IsUnionIndex(0, 0, 13)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	return x._message.GetUint256ArrayIteratorInOffset(off)
}

func (x *Argument) StringUint256ArrayValue() string {
	res := "["
	itr := x.Uint256ArrayValueIterator()
	for itr.HasNext() {
		res += fmt.Sprintf("%v", itr.NextUint256()) + ","
	}
	res += "]"
	return res
}

func (x *Argument) IsTypeBytes20ArrayValue() bool {
	is, _ := x._message.IsUnionIndex(0, 0, 14)
	return is
}

func (x *Argument) Bytes20ArrayValueCopiedToNative() [][20]byte {
	is, off := x._message.IsUnionIndex(0, 0, 14)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	var res [][20]byte
	itr := x._message.GetBytes20ArrayIteratorInOffset(off)
	for itr.HasNext() {
		res = append(res, itr.NextBytes20())
	}
	return res
}

func (x *Argument) Bytes20ArrayValueIterator() *membuffers.Iterator {
	is, off := x._message.IsUnionIndex(0, 0, 14)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	return x._message.GetBytes20ArrayIteratorInOffset(off)
}

func (x *Argument) StringBytes20ArrayValue() string {
	res := "["
	itr := x.Bytes20ArrayValueIterator()
	for itr.HasNext() {
		res += fmt.Sprintf("%x", itr.NextBytes20()) + ","
	}
	res += "]"
	return res
}

func (x *Argument) IsTypeBytes32ArrayValue() bool {
	is, _ := x._message.IsUnionIndex(0, 0, 15)
	return is
}

func (x *Argument) Bytes32ArrayValueCopiedToNative() [][32]byte {
	is, off := x._message.IsUnionIndex(0, 0, 15)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	var res [][32]byte
	itr := x._message.GetBytes32ArrayIteratorInOffset(off)
	for itr.HasNext() {
		res = append(res, itr.NextBytes32())
	}
	return res
}

func (x *Argument) Bytes32ArrayValueIterator() *membuffers.Iterator {
	is, off := x._message.IsUnionIndex(0, 0, 15)
	if !is {
		panic("Accessed union field of incorrect type, did you check which union type it is first?")
	}
	return x._message.GetBytes32ArrayIteratorInOffset(off)
}

func (x *Argument) StringBytes32ArrayValue() string {
	res := "["
	itr := x.Bytes32ArrayValueIterator()
	for itr.HasNext() {
		res += fmt.Sprintf("%x", itr.NextBytes32()) + ","
	}
	res += "]"
	return res
}

func (x *Argument) RawType() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *Argument) RawTypeWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *Argument) StringType() string {
	switch x.Type() {
	case ARGUMENT_TYPE_UINT_32_VALUE:
		return "(Uint32Value)" + x.StringUint32Value()
	case ARGUMENT_TYPE_UINT_64_VALUE:
		return "(Uint64Value)" + x.StringUint64Value()
	case ARGUMENT_TYPE_STRING_VALUE:
		return "(StringValue)" + x.StringStringValue()
	case ARGUMENT_TYPE_BYTES_VALUE:
		return "(BytesValue)" + x.StringBytesValue()
	case ARGUMENT_TYPE_BOOL_VALUE:
		return "(BoolValue)" + x.StringBoolValue()
	case ARGUMENT_TYPE_UINT_256_VALUE:
		return "(Uint256Value)" + x.StringUint256Value()
	case ARGUMENT_TYPE_BYTES_20_VALUE:
		return "(Bytes20Value)" + x.StringBytes20Value()
	case ARGUMENT_TYPE_BYTES_32_VALUE:
		return "(Bytes32Value)" + x.StringBytes32Value()
	case ARGUMENT_TYPE_UINT_32_ARRAY_VALUE:
		return "(Uint32ArrayValue)" + x.StringUint32ArrayValue()
	case ARGUMENT_TYPE_UINT_64_ARRAY_VALUE:
		return "(Uint64ArrayValue)" + x.StringUint64ArrayValue()
	case ARGUMENT_TYPE_STRING_ARRAY_VALUE:
		return "(StringArrayValue)" + x.StringStringArrayValue()
	case ARGUMENT_TYPE_BYTES_ARRAY_VALUE:
		return "(BytesArrayValue)" + x.StringBytesArrayValue()
	case ARGUMENT_TYPE_BOOL_ARRAY_VALUE:
		return "(BoolArrayValue)" + x.StringBoolArrayValue()
	case ARGUMENT_TYPE_UINT_256_ARRAY_VALUE:
		return "(Uint256ArrayValue)" + x.StringUint256ArrayValue()
	case ARGUMENT_TYPE_BYTES_20_ARRAY_VALUE:
		return "(Bytes20ArrayValue)" + x.StringBytes20ArrayValue()
	case ARGUMENT_TYPE_BYTES_32_ARRAY_VALUE:
		return "(Bytes32ArrayValue)" + x.StringBytes32ArrayValue()
	}
	return "(Unknown)"
}

// builder

type ArgumentBuilder struct {
	Type              ArgumentType
	Uint32Value       uint32
	Uint64Value       uint64
	StringValue       string
	BytesValue        []byte
	BoolValue         bool
	Uint256Value      *big.Int
	Bytes20Value      [20]byte
	Bytes32Value      [32]byte
	Uint32ArrayValue  []uint32
	Uint64ArrayValue  []uint64
	StringArrayValue  []string
	BytesArrayValue   [][]byte
	BoolArrayValue    []bool
	Uint256ArrayValue []*big.Int
	Bytes20ArrayValue [][20]byte
	Bytes32ArrayValue [][32]byte

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *ArgumentBuilder) Write(buf []byte) (err error) {
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
	w._builder.WriteUnionIndex(buf, uint16(w.Type))
	switch w.Type {
	case ARGUMENT_TYPE_UINT_32_VALUE:
		w._builder.WriteUint32(buf, w.Uint32Value)
	case ARGUMENT_TYPE_UINT_64_VALUE:
		w._builder.WriteUint64(buf, w.Uint64Value)
	case ARGUMENT_TYPE_STRING_VALUE:
		w._builder.WriteString(buf, w.StringValue)
	case ARGUMENT_TYPE_BYTES_VALUE:
		w._builder.WriteBytes(buf, w.BytesValue)
	case ARGUMENT_TYPE_BOOL_VALUE:
		w._builder.WriteBool(buf, w.BoolValue)
	case ARGUMENT_TYPE_UINT_256_VALUE:
		w._builder.WriteUint256(buf, w.Uint256Value)
	case ARGUMENT_TYPE_BYTES_20_VALUE:
		w._builder.WriteBytes20(buf, w.Bytes20Value)
	case ARGUMENT_TYPE_BYTES_32_VALUE:
		w._builder.WriteBytes32(buf, w.Bytes32Value)
	case ARGUMENT_TYPE_UINT_32_ARRAY_VALUE:
		w._builder.WriteUint32Array(buf, w.Uint32ArrayValue)
	case ARGUMENT_TYPE_UINT_64_ARRAY_VALUE:
		w._builder.WriteUint64Array(buf, w.Uint64ArrayValue)
	case ARGUMENT_TYPE_STRING_ARRAY_VALUE:
		w._builder.WriteStringArray(buf, w.StringArrayValue)
	case ARGUMENT_TYPE_BYTES_ARRAY_VALUE:
		w._builder.WriteBytesArray(buf, w.BytesArrayValue)
	case ARGUMENT_TYPE_BOOL_ARRAY_VALUE:
		w._builder.WriteBoolArray(buf, w.BoolArrayValue)
	case ARGUMENT_TYPE_UINT_256_ARRAY_VALUE:
		w._builder.WriteUint256Array(buf, w.Uint256ArrayValue)
	case ARGUMENT_TYPE_BYTES_20_ARRAY_VALUE:
		w._builder.WriteBytes20Array(buf, w.Bytes20ArrayValue)
	case ARGUMENT_TYPE_BYTES_32_ARRAY_VALUE:
		w._builder.WriteBytes32Array(buf, w.Bytes32ArrayValue)
	}
	return nil
}

func (w *ArgumentBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.HexDumpUnionIndex(prefix, offsetFromStart, "Argument.Type", uint16(w.Type))
	switch w.Type {
	case ARGUMENT_TYPE_UINT_32_VALUE:
		w._builder.HexDumpUint32(prefix, offsetFromStart, "Argument.Uint32Value", w.Uint32Value)
	case ARGUMENT_TYPE_UINT_64_VALUE:
		w._builder.HexDumpUint64(prefix, offsetFromStart, "Argument.Uint64Value", w.Uint64Value)
	case ARGUMENT_TYPE_STRING_VALUE:
		w._builder.HexDumpString(prefix, offsetFromStart, "Argument.StringValue", w.StringValue)
	case ARGUMENT_TYPE_BYTES_VALUE:
		w._builder.HexDumpBytes(prefix, offsetFromStart, "Argument.BytesValue", w.BytesValue)
	case ARGUMENT_TYPE_BOOL_VALUE:
		w._builder.HexDumpBool(prefix, offsetFromStart, "Argument.BoolValue", w.BoolValue)
	case ARGUMENT_TYPE_UINT_256_VALUE:
		w._builder.HexDumpUint256(prefix, offsetFromStart, "Argument.Uint256Value", w.Uint256Value)
	case ARGUMENT_TYPE_BYTES_20_VALUE:
		w._builder.HexDumpBytes20(prefix, offsetFromStart, "Argument.Bytes20Value", w.Bytes20Value)
	case ARGUMENT_TYPE_BYTES_32_VALUE:
		w._builder.HexDumpBytes32(prefix, offsetFromStart, "Argument.Bytes32Value", w.Bytes32Value)
	case ARGUMENT_TYPE_UINT_32_ARRAY_VALUE:
		w._builder.HexDumpUint32Array(prefix, offsetFromStart, "Argument.Uint32ArrayValue", w.Uint32ArrayValue)
	case ARGUMENT_TYPE_UINT_64_ARRAY_VALUE:
		w._builder.HexDumpUint64Array(prefix, offsetFromStart, "Argument.Uint64ArrayValue", w.Uint64ArrayValue)
	case ARGUMENT_TYPE_STRING_ARRAY_VALUE:
		w._builder.HexDumpStringArray(prefix, offsetFromStart, "Argument.StringArrayValue", w.StringArrayValue)
	case ARGUMENT_TYPE_BYTES_ARRAY_VALUE:
		w._builder.HexDumpBytesArray(prefix, offsetFromStart, "Argument.BytesArrayValue", w.BytesArrayValue)
	case ARGUMENT_TYPE_BOOL_ARRAY_VALUE:
		w._builder.HexDumpBoolArray(prefix, offsetFromStart, "Argument.BoolArrayValue", w.BoolArrayValue)
	case ARGUMENT_TYPE_UINT_256_ARRAY_VALUE:
		w._builder.HexDumpUint256Array(prefix, offsetFromStart, "Argument.Uint256ArrayValue", w.Uint256ArrayValue)
	case ARGUMENT_TYPE_BYTES_20_ARRAY_VALUE:
		w._builder.HexDumpBytes20Array(prefix, offsetFromStart, "Argument.Bytes20ArrayValue", w.Bytes20ArrayValue)
	case ARGUMENT_TYPE_BYTES_32_ARRAY_VALUE:
		w._builder.HexDumpBytes32Array(prefix, offsetFromStart, "Argument.Bytes32ArrayValue", w.Bytes32ArrayValue)
	}
	return nil
}

func (w *ArgumentBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *ArgumentBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *ArgumentBuilder) Build() *Argument {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ArgumentReader(buf)
}

func ArgumentBuilderFromRaw(raw []byte) *ArgumentBuilder {
	return &ArgumentBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message ArgumentArray

// reader

type ArgumentArray struct {
	// Arguments []Argument

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *ArgumentArray) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Arguments:%s,}", x.StringArguments())
}

var _ArgumentArray_Scheme = []membuffers.FieldType{membuffers.TypeMessageArray}
var _ArgumentArray_Unions = [][]membuffers.FieldType{}

func ArgumentArrayReader(buf []byte) *ArgumentArray {
	x := &ArgumentArray{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _ArgumentArray_Scheme, _ArgumentArray_Unions)
	return x
}

func (x *ArgumentArray) IsValid() bool {
	return x._message.IsValid()
}

func (x *ArgumentArray) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *ArgumentArray) Equal(y *ArgumentArray) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *ArgumentArray) ArgumentsIterator() *ArgumentArrayArgumentsIterator {
	return &ArgumentArrayArgumentsIterator{iterator: x._message.GetMessageArrayIterator(0)}
}

type ArgumentArrayArgumentsIterator struct {
	iterator *membuffers.Iterator
}

func (i *ArgumentArrayArgumentsIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *ArgumentArrayArgumentsIterator) NextArguments() *Argument {
	b, s := i.iterator.NextMessage()
	return ArgumentReader(b[:s])
}

func (x *ArgumentArray) RawArgumentsArray() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *ArgumentArray) RawArgumentsArrayWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *ArgumentArray) StringArguments() (res string) {
	res = "["
	for i := x.ArgumentsIterator(); i.HasNext(); {
		res += i.NextArguments().String() + ","
	}
	res += "]"
	return
}

// builder

type ArgumentArrayBuilder struct {
	Arguments []*ArgumentBuilder

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *ArgumentArrayBuilder) arrayOfArguments() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.Arguments))
	for i, v := range w.Arguments {
		res[i] = v
	}
	return res
}

func (w *ArgumentArrayBuilder) Write(buf []byte) (err error) {
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
	err = w._builder.WriteMessageArray(buf, w.arrayOfArguments())
	if err != nil {
		return
	}
	return nil
}

func (w *ArgumentArrayBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.HexDumpMessageArray(prefix, offsetFromStart, "ArgumentArray.Arguments", w.arrayOfArguments())
	if err != nil {
		return
	}
	return nil
}

func (w *ArgumentArrayBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *ArgumentArrayBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *ArgumentArrayBuilder) Build() *ArgumentArray {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ArgumentArrayReader(buf)
}

func ArgumentArrayBuilderFromRaw(raw []byte) *ArgumentArrayBuilder {
	return &ArgumentArrayBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message StateRecord

// reader

type StateRecord struct {
	// Key []byte
	// Value []byte

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *StateRecord) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Key:%s,Value:%s,}", x.StringKey(), x.StringValue())
}

var _StateRecord_Scheme = []membuffers.FieldType{membuffers.TypeBytes, membuffers.TypeBytes}
var _StateRecord_Unions = [][]membuffers.FieldType{}

func StateRecordReader(buf []byte) *StateRecord {
	x := &StateRecord{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _StateRecord_Scheme, _StateRecord_Unions)
	return x
}

func (x *StateRecord) IsValid() bool {
	return x._message.IsValid()
}

func (x *StateRecord) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *StateRecord) Equal(y *StateRecord) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *StateRecord) Key() []byte {
	return x._message.GetBytes(0)
}

func (x *StateRecord) RawKey() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *StateRecord) RawKeyWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *StateRecord) MutateKey(v []byte) error {
	return x._message.SetBytes(0, v)
}

func (x *StateRecord) StringKey() string {
	return fmt.Sprintf("%x", x.Key())
}

func (x *StateRecord) Value() []byte {
	return x._message.GetBytes(1)
}

func (x *StateRecord) RawValue() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *StateRecord) RawValueWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(1, 0)
}

func (x *StateRecord) MutateValue(v []byte) error {
	return x._message.SetBytes(1, v)
}

func (x *StateRecord) StringValue() string {
	return fmt.Sprintf("%x", x.Value())
}

// builder

type StateRecordBuilder struct {
	Key   []byte
	Value []byte

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *StateRecordBuilder) Write(buf []byte) (err error) {
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
	w._builder.WriteBytes(buf, w.Key)
	w._builder.WriteBytes(buf, w.Value)
	return nil
}

func (w *StateRecordBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.HexDumpBytes(prefix, offsetFromStart, "StateRecord.Key", w.Key)
	w._builder.HexDumpBytes(prefix, offsetFromStart, "StateRecord.Value", w.Value)
	return nil
}

func (w *StateRecordBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *StateRecordBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *StateRecordBuilder) Build() *StateRecord {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return StateRecordReader(buf)
}

func StateRecordBuilderFromRaw(raw []byte) *StateRecordBuilder {
	return &StateRecordBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message ContractStateDiff

// reader

type ContractStateDiff struct {
	// ContractName primitives.ContractName
	// StateDiffs []StateRecord

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *ContractStateDiff) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ContractName:%s,StateDiffs:%s,}", x.StringContractName(), x.StringStateDiffs())
}

var _ContractStateDiff_Scheme = []membuffers.FieldType{membuffers.TypeString, membuffers.TypeMessageArray}
var _ContractStateDiff_Unions = [][]membuffers.FieldType{}

func ContractStateDiffReader(buf []byte) *ContractStateDiff {
	x := &ContractStateDiff{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _ContractStateDiff_Scheme, _ContractStateDiff_Unions)
	return x
}

func (x *ContractStateDiff) IsValid() bool {
	return x._message.IsValid()
}

func (x *ContractStateDiff) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *ContractStateDiff) Equal(y *ContractStateDiff) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *ContractStateDiff) ContractName() primitives.ContractName {
	return primitives.ContractName(x._message.GetString(0))
}

func (x *ContractStateDiff) RawContractName() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *ContractStateDiff) RawContractNameWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *ContractStateDiff) MutateContractName(v primitives.ContractName) error {
	return x._message.SetString(0, string(v))
}

func (x *ContractStateDiff) StringContractName() string {
	return fmt.Sprintf("%s", x.ContractName())
}

func (x *ContractStateDiff) StateDiffsIterator() *ContractStateDiffStateDiffsIterator {
	return &ContractStateDiffStateDiffsIterator{iterator: x._message.GetMessageArrayIterator(1)}
}

type ContractStateDiffStateDiffsIterator struct {
	iterator *membuffers.Iterator
}

func (i *ContractStateDiffStateDiffsIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *ContractStateDiffStateDiffsIterator) NextStateDiffs() *StateRecord {
	b, s := i.iterator.NextMessage()
	return StateRecordReader(b[:s])
}

func (x *ContractStateDiff) RawStateDiffsArray() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *ContractStateDiff) RawStateDiffsArrayWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(1, 0)
}

func (x *ContractStateDiff) StringStateDiffs() (res string) {
	res = "["
	for i := x.StateDiffsIterator(); i.HasNext(); {
		res += i.NextStateDiffs().String() + ","
	}
	res += "]"
	return
}

// builder

type ContractStateDiffBuilder struct {
	ContractName primitives.ContractName
	StateDiffs   []*StateRecordBuilder

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *ContractStateDiffBuilder) arrayOfStateDiffs() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.StateDiffs))
	for i, v := range w.StateDiffs {
		res[i] = v
	}
	return res
}

func (w *ContractStateDiffBuilder) Write(buf []byte) (err error) {
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
	w._builder.WriteString(buf, string(w.ContractName))
	err = w._builder.WriteMessageArray(buf, w.arrayOfStateDiffs())
	if err != nil {
		return
	}
	return nil
}

func (w *ContractStateDiffBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.HexDumpString(prefix, offsetFromStart, "ContractStateDiff.ContractName", string(w.ContractName))
	err = w._builder.HexDumpMessageArray(prefix, offsetFromStart, "ContractStateDiff.StateDiffs", w.arrayOfStateDiffs())
	if err != nil {
		return
	}
	return nil
}

func (w *ContractStateDiffBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *ContractStateDiffBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *ContractStateDiffBuilder) Build() *ContractStateDiff {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return ContractStateDiffReader(buf)
}

func ContractStateDiffBuilderFromRaw(raw []byte) *ContractStateDiffBuilder {
	return &ContractStateDiffBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message Event

// reader

type Event struct {
	// ContractName primitives.ContractName
	// EventName primitives.EventName
	// OutputArgumentArray primitives.PackedArgumentArray

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *Event) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{ContractName:%s,EventName:%s,OutputArgumentArray:%s,}", x.StringContractName(), x.StringEventName(), x.StringOutputArgumentArray())
}

var _Event_Scheme = []membuffers.FieldType{membuffers.TypeString, membuffers.TypeString, membuffers.TypeBytes}
var _Event_Unions = [][]membuffers.FieldType{}

func EventReader(buf []byte) *Event {
	x := &Event{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _Event_Scheme, _Event_Unions)
	return x
}

func (x *Event) IsValid() bool {
	return x._message.IsValid()
}

func (x *Event) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *Event) Equal(y *Event) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *Event) ContractName() primitives.ContractName {
	return primitives.ContractName(x._message.GetString(0))
}

func (x *Event) RawContractName() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *Event) RawContractNameWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *Event) MutateContractName(v primitives.ContractName) error {
	return x._message.SetString(0, string(v))
}

func (x *Event) StringContractName() string {
	return fmt.Sprintf("%s", x.ContractName())
}

func (x *Event) EventName() primitives.EventName {
	return primitives.EventName(x._message.GetString(1))
}

func (x *Event) RawEventName() []byte {
	return x._message.RawBufferForField(1, 0)
}

func (x *Event) RawEventNameWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(1, 0)
}

func (x *Event) MutateEventName(v primitives.EventName) error {
	return x._message.SetString(1, string(v))
}

func (x *Event) StringEventName() string {
	return fmt.Sprintf("%s", x.EventName())
}

func (x *Event) OutputArgumentArray() primitives.PackedArgumentArray {
	return primitives.PackedArgumentArray(x._message.GetBytes(2))
}

func (x *Event) RawOutputArgumentArray() []byte {
	return x._message.RawBufferForField(2, 0)
}

func (x *Event) RawOutputArgumentArrayWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(2, 0)
}

func (x *Event) MutateOutputArgumentArray(v primitives.PackedArgumentArray) error {
	return x._message.SetBytes(2, []byte(v))
}

func (x *Event) StringOutputArgumentArray() string {
	return fmt.Sprintf("%s", x.OutputArgumentArray())
}

// builder

type EventBuilder struct {
	ContractName        primitives.ContractName
	EventName           primitives.EventName
	OutputArgumentArray primitives.PackedArgumentArray

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *EventBuilder) Write(buf []byte) (err error) {
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
	w._builder.WriteString(buf, string(w.ContractName))
	w._builder.WriteString(buf, string(w.EventName))
	w._builder.WriteBytes(buf, []byte(w.OutputArgumentArray))
	return nil
}

func (w *EventBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	w._builder.HexDumpString(prefix, offsetFromStart, "Event.ContractName", string(w.ContractName))
	w._builder.HexDumpString(prefix, offsetFromStart, "Event.EventName", string(w.EventName))
	w._builder.HexDumpBytes(prefix, offsetFromStart, "Event.OutputArgumentArray", []byte(w.OutputArgumentArray))
	return nil
}

func (w *EventBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *EventBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *EventBuilder) Build() *Event {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return EventReader(buf)
}

func EventBuilderFromRaw(raw []byte) *EventBuilder {
	return &EventBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// message EventsArray

// reader

type EventsArray struct {
	// Events []Event

	// internal
	// implements membuffers.Message
	_message membuffers.InternalMessage
}

func (x *EventsArray) String() string {
	if x == nil {
		return "<nil>"
	}
	return fmt.Sprintf("{Events:%s,}", x.StringEvents())
}

var _EventsArray_Scheme = []membuffers.FieldType{membuffers.TypeMessageArray}
var _EventsArray_Unions = [][]membuffers.FieldType{}

func EventsArrayReader(buf []byte) *EventsArray {
	x := &EventsArray{}
	x._message.Init(buf, membuffers.Offset(len(buf)), _EventsArray_Scheme, _EventsArray_Unions)
	return x
}

func (x *EventsArray) IsValid() bool {
	return x._message.IsValid()
}

func (x *EventsArray) Raw() []byte {
	return x._message.RawBuffer()
}

func (x *EventsArray) Equal(y *EventsArray) bool {
	if x == nil && y == nil {
		return true
	}
	if x == nil || y == nil {
		return false
	}
	return bytes.Equal(x.Raw(), y.Raw())
}

func (x *EventsArray) EventsIterator() *EventsArrayEventsIterator {
	return &EventsArrayEventsIterator{iterator: x._message.GetMessageArrayIterator(0)}
}

type EventsArrayEventsIterator struct {
	iterator *membuffers.Iterator
}

func (i *EventsArrayEventsIterator) HasNext() bool {
	return i.iterator.HasNext()
}

func (i *EventsArrayEventsIterator) NextEvents() *Event {
	b, s := i.iterator.NextMessage()
	return EventReader(b[:s])
}

func (x *EventsArray) RawEventsArray() []byte {
	return x._message.RawBufferForField(0, 0)
}

func (x *EventsArray) RawEventsArrayWithHeader() []byte {
	return x._message.RawBufferWithHeaderForField(0, 0)
}

func (x *EventsArray) StringEvents() (res string) {
	res = "["
	for i := x.EventsIterator(); i.HasNext(); {
		res += i.NextEvents().String() + ","
	}
	res += "]"
	return
}

// builder

type EventsArrayBuilder struct {
	Events []*EventBuilder

	// internal
	// implements membuffers.Builder
	_builder               membuffers.InternalBuilder
	_overrideWithRawBuffer []byte
}

func (w *EventsArrayBuilder) arrayOfEvents() []membuffers.MessageWriter {
	res := make([]membuffers.MessageWriter, len(w.Events))
	for i, v := range w.Events {
		res[i] = v
	}
	return res
}

func (w *EventsArrayBuilder) Write(buf []byte) (err error) {
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
	err = w._builder.WriteMessageArray(buf, w.arrayOfEvents())
	if err != nil {
		return
	}
	return nil
}

func (w *EventsArrayBuilder) HexDump(prefix string, offsetFromStart membuffers.Offset) (err error) {
	if w == nil {
		return
	}
	defer func() {
		if r := recover(); r != nil {
			err = &membuffers.ErrBufferOverrun{}
		}
	}()
	w._builder.Reset()
	err = w._builder.HexDumpMessageArray(prefix, offsetFromStart, "EventsArray.Events", w.arrayOfEvents())
	if err != nil {
		return
	}
	return nil
}

func (w *EventsArrayBuilder) GetSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	return w._builder.GetSize()
}

func (w *EventsArrayBuilder) CalcRequiredSize() membuffers.Offset {
	if w == nil {
		return 0
	}
	w.Write(nil)
	return w._builder.GetSize()
}

func (w *EventsArrayBuilder) Build() *EventsArray {
	buf := make([]byte, w.CalcRequiredSize())
	if w.Write(buf) != nil {
		return nil
	}
	return EventsArrayReader(buf)
}

func EventsArrayBuilderFromRaw(raw []byte) *EventsArrayBuilder {
	return &EventsArrayBuilder{_overrideWithRawBuffer: raw}
}

/////////////////////////////////////////////////////////////////////////////
// enums

type ExecutionAccessScope uint16

const (
	ACCESS_SCOPE_RESERVED   ExecutionAccessScope = 0
	ACCESS_SCOPE_READ_ONLY  ExecutionAccessScope = 1
	ACCESS_SCOPE_READ_WRITE ExecutionAccessScope = 2
)

func (n ExecutionAccessScope) String() string {
	switch n {
	case ACCESS_SCOPE_RESERVED:
		return "ACCESS_SCOPE_RESERVED"
	case ACCESS_SCOPE_READ_ONLY:
		return "ACCESS_SCOPE_READ_ONLY"
	case ACCESS_SCOPE_READ_WRITE:
		return "ACCESS_SCOPE_READ_WRITE"
	}
	return "UNKNOWN"
}

type ExecutionPermissionScope uint16

const (
	PERMISSION_SCOPE_RESERVED ExecutionPermissionScope = 0
	PERMISSION_SCOPE_SYSTEM   ExecutionPermissionScope = 1
	PERMISSION_SCOPE_SERVICE  ExecutionPermissionScope = 2
)

func (n ExecutionPermissionScope) String() string {
	switch n {
	case PERMISSION_SCOPE_RESERVED:
		return "PERMISSION_SCOPE_RESERVED"
	case PERMISSION_SCOPE_SYSTEM:
		return "PERMISSION_SCOPE_SYSTEM"
	case PERMISSION_SCOPE_SERVICE:
		return "PERMISSION_SCOPE_SERVICE"
	}
	return "UNKNOWN"
}

type ProcessorType uint16

const (
	PROCESSOR_TYPE_RESERVED   ProcessorType = 0
	PROCESSOR_TYPE_NATIVE     ProcessorType = 1
	PROCESSOR_TYPE_JAVASCRIPT ProcessorType = 2
)

func (n ProcessorType) String() string {
	switch n {
	case PROCESSOR_TYPE_RESERVED:
		return "PROCESSOR_TYPE_RESERVED"
	case PROCESSOR_TYPE_NATIVE:
		return "PROCESSOR_TYPE_NATIVE"
	case PROCESSOR_TYPE_JAVASCRIPT:
		return "PROCESSOR_TYPE_JAVASCRIPT"
	}
	return "UNKNOWN"
}

type CrosschainConnectorType uint16

const (
	CROSSCHAIN_CONNECTOR_TYPE_RESERVED CrosschainConnectorType = 0
	CROSSCHAIN_CONNECTOR_TYPE_ETHEREUM CrosschainConnectorType = 1
)

func (n CrosschainConnectorType) String() string {
	switch n {
	case CROSSCHAIN_CONNECTOR_TYPE_RESERVED:
		return "CROSSCHAIN_CONNECTOR_TYPE_RESERVED"
	case CROSSCHAIN_CONNECTOR_TYPE_ETHEREUM:
		return "CROSSCHAIN_CONNECTOR_TYPE_ETHEREUM"
	}
	return "UNKNOWN"
}
