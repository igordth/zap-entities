package encoder

import (
	"github.com/igordth/tostring"
)

type EmptyPrimitiveArray struct{}

func (p EmptyPrimitiveArray) AppendBool(bool)             {}
func (p EmptyPrimitiveArray) AppendByteString([]byte)     {}
func (p EmptyPrimitiveArray) AppendComplex128(complex128) {}
func (p EmptyPrimitiveArray) AppendComplex64(complex64)   {}
func (p EmptyPrimitiveArray) AppendFloat64(float64)       {}
func (p EmptyPrimitiveArray) AppendFloat32(float32)       {}
func (p EmptyPrimitiveArray) AppendInt(int)               {}
func (p EmptyPrimitiveArray) AppendInt64(int64)           {}
func (p EmptyPrimitiveArray) AppendInt32(int32)           {}
func (p EmptyPrimitiveArray) AppendInt16(int16)           {}
func (p EmptyPrimitiveArray) AppendInt8(int8)             {}
func (p EmptyPrimitiveArray) AppendString(string)         {}
func (p EmptyPrimitiveArray) AppendUint(uint)             {}
func (p EmptyPrimitiveArray) AppendUint64(uint64)         {}
func (p EmptyPrimitiveArray) AppendUint32(uint32)         {}
func (p EmptyPrimitiveArray) AppendUint16(uint16)         {}
func (p EmptyPrimitiveArray) AppendUint8(uint8)           {}
func (p EmptyPrimitiveArray) AppendUintptr(uintptr)       {}

type ValuePrimitiveArray struct{ Value string }

func (p *ValuePrimitiveArray) AppendBool(b bool)             { p.Value = tostring.Bool(b) }
func (p *ValuePrimitiveArray) AppendByteString(bytes []byte) { p.Value = tostring.Byte(bytes) }
func (p *ValuePrimitiveArray) AppendFloat64(f float64)       { p.Value = tostring.Float64(f) }
func (p *ValuePrimitiveArray) AppendFloat32(f float32)       { p.Value = tostring.Float32(f) }
func (p *ValuePrimitiveArray) AppendInt(i int)               { p.Value = tostring.Int(i) }
func (p *ValuePrimitiveArray) AppendInt64(i int64)           { p.Value = tostring.Int64(i) }
func (p *ValuePrimitiveArray) AppendInt32(i int32)           { p.Value = tostring.Int32(i) }
func (p *ValuePrimitiveArray) AppendInt16(i int16)           { p.Value = tostring.Int16(i) }
func (p *ValuePrimitiveArray) AppendInt8(i int8)             { p.Value = tostring.Int8(i) }
func (p *ValuePrimitiveArray) AppendUint(u uint)             { p.Value = tostring.Uint(u) }
func (p *ValuePrimitiveArray) AppendUint64(u uint64)         { p.Value = tostring.Uint64(u) }
func (p *ValuePrimitiveArray) AppendUint32(u uint32)         { p.Value = tostring.Uint32(u) }
func (p *ValuePrimitiveArray) AppendUint16(u uint16)         { p.Value = tostring.Uint16(u) }
func (p *ValuePrimitiveArray) AppendUint8(u uint8)           { p.Value = tostring.Uint8(u) }
func (p *ValuePrimitiveArray) AppendUintptr(u uintptr)       { p.Value = tostring.Uintptr(u) }
func (p *ValuePrimitiveArray) AppendString(s string)         { p.Value = s }
func (p *ValuePrimitiveArray) AppendComplex64(c complex64)   { p.Value = tostring.Complex64(c) }
func (p *ValuePrimitiveArray) AppendComplex128(c complex128) { p.Value = tostring.Complex128(c) }
