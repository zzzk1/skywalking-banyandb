// Code generated by the FlatBuffers compiler. DO NOT EDIT.

package v1

import (
	"strconv"

	flatbuffers "github.com/google/flatbuffers/go"
)

type ValueType byte

const (
	ValueTypeNONE        ValueType = 0
	ValueTypeString      ValueType = 1
	ValueTypeStringArray ValueType = 2
	ValueTypeInt         ValueType = 3
	ValueTypeIntArray    ValueType = 4
)

var EnumNamesValueType = map[ValueType]string{
	ValueTypeNONE:        "NONE",
	ValueTypeString:      "String",
	ValueTypeStringArray: "StringArray",
	ValueTypeInt:         "Int",
	ValueTypeIntArray:    "IntArray",
}

var EnumValuesValueType = map[string]ValueType{
	"NONE":        ValueTypeNONE,
	"String":      ValueTypeString,
	"StringArray": ValueTypeStringArray,
	"Int":         ValueTypeInt,
	"IntArray":    ValueTypeIntArray,
}

func (v ValueType) String() string {
	if s, ok := EnumNamesValueType[v]; ok {
		return s
	}
	return "ValueType(" + strconv.FormatInt(int64(v), 10) + ")"
}

type String struct {
	_tab flatbuffers.Table
}

func GetRootAsString(buf []byte, offset flatbuffers.UOffsetT) *String {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &String{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsString(buf []byte, offset flatbuffers.UOffsetT) *String {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &String{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *String) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *String) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *String) Value() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func StringStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func StringAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(value), 0)
}
func StringEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

type Int struct {
	_tab flatbuffers.Table
}

func GetRootAsInt(buf []byte, offset flatbuffers.UOffsetT) *Int {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Int{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsInt(buf []byte, offset flatbuffers.UOffsetT) *Int {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Int{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Int) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Int) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Int) Value() int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.GetInt64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *Int) MutateValue(n int64) bool {
	return rcv._tab.MutateInt64Slot(4, n)
}

func IntStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func IntAddValue(builder *flatbuffers.Builder, value int64) {
	builder.PrependInt64Slot(0, value, 0)
}
func IntEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

type StringArray struct {
	_tab flatbuffers.Table
}

func GetRootAsStringArray(buf []byte, offset flatbuffers.UOffsetT) *StringArray {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &StringArray{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsStringArray(buf []byte, offset flatbuffers.UOffsetT) *StringArray {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &StringArray{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *StringArray) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *StringArray) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *StringArray) Value(j int) []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.ByteVector(a + flatbuffers.UOffsetT(j*4))
	}
	return nil
}

func (rcv *StringArray) ValueLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func StringArrayStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func StringArrayAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(value), 0)
}
func StringArrayStartValueVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func StringArrayEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

type IntArray struct {
	_tab flatbuffers.Table
}

func GetRootAsIntArray(buf []byte, offset flatbuffers.UOffsetT) *IntArray {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &IntArray{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsIntArray(buf []byte, offset flatbuffers.UOffsetT) *IntArray {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &IntArray{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *IntArray) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *IntArray) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *IntArray) Value(j int) int64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetInt64(a + flatbuffers.UOffsetT(j*8))
	}
	return 0
}

func (rcv *IntArray) ValueLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *IntArray) MutateValue(j int, n int64) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateInt64(a+flatbuffers.UOffsetT(j*8), n)
	}
	return false
}

func IntArrayStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func IntArrayAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(value), 0)
}
func IntArrayStartValueVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(8, numElems, 8)
}
func IntArrayEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

type FieldValue struct {
	_tab flatbuffers.Table
}

func GetRootAsFieldValue(buf []byte, offset flatbuffers.UOffsetT) *FieldValue {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &FieldValue{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsFieldValue(buf []byte, offset flatbuffers.UOffsetT) *FieldValue {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &FieldValue{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *FieldValue) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *FieldValue) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *FieldValue) ValueType() ValueType {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return ValueType(rcv._tab.GetByte(o + rcv._tab.Pos))
	}
	return 0
}

func (rcv *FieldValue) MutateValueType(n ValueType) bool {
	return rcv._tab.MutateByteSlot(4, byte(n))
}

func (rcv *FieldValue) Value(obj *flatbuffers.Table) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		rcv._tab.Union(obj, o)
		return true
	}
	return false
}

func FieldValueStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func FieldValueAddValueType(builder *flatbuffers.Builder, valueType ValueType) {
	builder.PrependByteSlot(0, byte(valueType), 0)
}
func FieldValueAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(value), 0)
}
func FieldValueEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

type Field struct {
	_tab flatbuffers.Table
}

func GetRootAsField(buf []byte, offset flatbuffers.UOffsetT) *Field {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &Field{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsField(buf []byte, offset flatbuffers.UOffsetT) *Field {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &Field{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *Field) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *Field) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *Field) Value(obj *FieldValue, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *Field) ValueLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func FieldStart(builder *flatbuffers.Builder) {
	builder.StartObject(1)
}
func FieldAddValue(builder *flatbuffers.Builder, value flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(value), 0)
}
func FieldStartValueVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func FieldEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

type entityValue struct {
	_tab flatbuffers.Table
}

func GetRootAsentityValue(buf []byte, offset flatbuffers.UOffsetT) *entityValue {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &entityValue{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsentityValue(buf []byte, offset flatbuffers.UOffsetT) *entityValue {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &entityValue{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *entityValue) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *entityValue) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *entityValue) EntityId() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *entityValue) TimestampNanoseconds() uint64 {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		return rcv._tab.GetUint64(o + rcv._tab.Pos)
	}
	return 0
}

func (rcv *entityValue) MutateTimestampNanoseconds(n uint64) bool {
	return rcv._tab.MutateUint64Slot(6, n)
}

func (rcv *entityValue) DataBinary(j int) byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.GetByte(a + flatbuffers.UOffsetT(j*1))
	}
	return 0
}

func (rcv *entityValue) DataBinaryLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func (rcv *entityValue) DataBinaryBytes() []byte {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		return rcv._tab.ByteVector(o + rcv._tab.Pos)
	}
	return nil
}

func (rcv *entityValue) MutateDataBinary(j int, n byte) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(8))
	if o != 0 {
		a := rcv._tab.Vector(o)
		return rcv._tab.MutateByte(a+flatbuffers.UOffsetT(j*1), n)
	}
	return false
}

func (rcv *entityValue) Fields(obj *Field, j int) bool {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		x := rcv._tab.Vector(o)
		x += flatbuffers.UOffsetT(j) * 4
		x = rcv._tab.Indirect(x)
		obj.Init(rcv._tab.Bytes, x)
		return true
	}
	return false
}

func (rcv *entityValue) FieldsLength() int {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(10))
	if o != 0 {
		return rcv._tab.VectorLen(o)
	}
	return 0
}

func entityValueStart(builder *flatbuffers.Builder) {
	builder.StartObject(4)
}
func entityValueAddEntityId(builder *flatbuffers.Builder, entityId flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(entityId), 0)
}
func entityValueAddTimestampNanoseconds(builder *flatbuffers.Builder, timestampNanoseconds uint64) {
	builder.PrependUint64Slot(1, timestampNanoseconds, 0)
}
func entityValueAddDataBinary(builder *flatbuffers.Builder, dataBinary flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(2, flatbuffers.UOffsetT(dataBinary), 0)
}
func entityValueStartDataBinaryVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(1, numElems, 1)
}
func entityValueAddFields(builder *flatbuffers.Builder, fields flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(3, flatbuffers.UOffsetT(fields), 0)
}
func entityValueStartFieldsVector(builder *flatbuffers.Builder, numElems int) flatbuffers.UOffsetT {
	return builder.StartVector(4, numElems, 4)
}
func entityValueEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

type WriteEntity struct {
	_tab flatbuffers.Table
}

func GetRootAsWriteEntity(buf []byte, offset flatbuffers.UOffsetT) *WriteEntity {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &WriteEntity{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsWriteEntity(buf []byte, offset flatbuffers.UOffsetT) *WriteEntity {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &WriteEntity{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *WriteEntity) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *WriteEntity) Table() flatbuffers.Table {
	return rcv._tab
}

func (rcv *WriteEntity) MetaData(obj *Metadata) *Metadata {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(4))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(Metadata)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func (rcv *WriteEntity) Entity(obj *entityValue) *entityValue {
	o := flatbuffers.UOffsetT(rcv._tab.Offset(6))
	if o != 0 {
		x := rcv._tab.Indirect(o + rcv._tab.Pos)
		if obj == nil {
			obj = new(entityValue)
		}
		obj.Init(rcv._tab.Bytes, x)
		return obj
	}
	return nil
}

func WriteEntityStart(builder *flatbuffers.Builder) {
	builder.StartObject(2)
}
func WriteEntityAddMetaData(builder *flatbuffers.Builder, metaData flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(0, flatbuffers.UOffsetT(metaData), 0)
}
func WriteEntityAddEntity(builder *flatbuffers.Builder, entity flatbuffers.UOffsetT) {
	builder.PrependUOffsetTSlot(1, flatbuffers.UOffsetT(entity), 0)
}
func WriteEntityEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}

type WriteResponse struct {
	_tab flatbuffers.Table
}

func GetRootAsWriteResponse(buf []byte, offset flatbuffers.UOffsetT) *WriteResponse {
	n := flatbuffers.GetUOffsetT(buf[offset:])
	x := &WriteResponse{}
	x.Init(buf, n+offset)
	return x
}

func GetSizePrefixedRootAsWriteResponse(buf []byte, offset flatbuffers.UOffsetT) *WriteResponse {
	n := flatbuffers.GetUOffsetT(buf[offset+flatbuffers.SizeUint32:])
	x := &WriteResponse{}
	x.Init(buf, n+offset+flatbuffers.SizeUint32)
	return x
}

func (rcv *WriteResponse) Init(buf []byte, i flatbuffers.UOffsetT) {
	rcv._tab.Bytes = buf
	rcv._tab.Pos = i
}

func (rcv *WriteResponse) Table() flatbuffers.Table {
	return rcv._tab
}

func WriteResponseStart(builder *flatbuffers.Builder) {
	builder.StartObject(0)
}
func WriteResponseEnd(builder *flatbuffers.Builder) flatbuffers.UOffsetT {
	return builder.EndObject()
}
