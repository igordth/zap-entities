package encoder

import "zap-cores/convtostr"

type EmptyPrimitiveArray struct{}

func (p EmptyPrimitiveArray) AppendBool(b bool)             {}
func (p EmptyPrimitiveArray) AppendByteString(bytes []byte) {}
func (p EmptyPrimitiveArray) AppendComplex128(c complex128) {}
func (p EmptyPrimitiveArray) AppendComplex64(c complex64)   {}
func (p EmptyPrimitiveArray) AppendFloat64(f float64)       {}
func (p EmptyPrimitiveArray) AppendFloat32(f float32)       {}
func (p EmptyPrimitiveArray) AppendInt(i int)               {}
func (p EmptyPrimitiveArray) AppendInt64(i int64)           {}
func (p EmptyPrimitiveArray) AppendInt32(i int32)           {}
func (p EmptyPrimitiveArray) AppendInt16(i int16)           {}
func (p EmptyPrimitiveArray) AppendInt8(i int8)             {}
func (p EmptyPrimitiveArray) AppendString(s string)         {}
func (p EmptyPrimitiveArray) AppendUint(u uint)             {}
func (p EmptyPrimitiveArray) AppendUint64(u uint64)         {}
func (p EmptyPrimitiveArray) AppendUint32(u uint32)         {}
func (p EmptyPrimitiveArray) AppendUint16(u uint16)         {}
func (p EmptyPrimitiveArray) AppendUint8(u uint8)           {}
func (p EmptyPrimitiveArray) AppendUintptr(u uintptr)       {}

type ValuePrimitiveArray struct{ Value string }

func (p *ValuePrimitiveArray) AppendBool(b bool)             { p.Value = convtostr.Bool(b) }
func (p *ValuePrimitiveArray) AppendByteString(bytes []byte) { p.Value = convtostr.Byte(bytes) }
func (p *ValuePrimitiveArray) AppendFloat64(f float64)       { p.Value = convtostr.Float64(f) }
func (p *ValuePrimitiveArray) AppendFloat32(f float32)       { p.Value = convtostr.Float32(f) }
func (p *ValuePrimitiveArray) AppendInt(i int)               { p.Value = convtostr.Int(i) }
func (p *ValuePrimitiveArray) AppendInt64(i int64)           { p.Value = convtostr.Int64(i) }
func (p *ValuePrimitiveArray) AppendInt32(i int32)           { p.Value = convtostr.Int32(i) }
func (p *ValuePrimitiveArray) AppendInt16(i int16)           { p.Value = convtostr.Int16(i) }
func (p *ValuePrimitiveArray) AppendInt8(i int8)             { p.Value = convtostr.Int8(i) }
func (p *ValuePrimitiveArray) AppendUint(u uint)             { p.Value = convtostr.Uint(u) }
func (p *ValuePrimitiveArray) AppendUint64(u uint64)         { p.Value = convtostr.Uint64(u) }
func (p *ValuePrimitiveArray) AppendUint32(u uint32)         { p.Value = convtostr.Uint32(u) }
func (p *ValuePrimitiveArray) AppendUint16(u uint16)         { p.Value = convtostr.Uint16(u) }
func (p *ValuePrimitiveArray) AppendUint8(u uint8)           { p.Value = convtostr.Uint8(u) }
func (p *ValuePrimitiveArray) AppendUintptr(u uintptr)       { p.Value = convtostr.Uintptr(u) }
func (p *ValuePrimitiveArray) AppendString(s string)         { p.Value = s }
func (p *ValuePrimitiveArray) AppendComplex64(c complex64)   { p.Value = convtostr.Complex64(c) }
func (p *ValuePrimitiveArray) AppendComplex128(c complex128) { p.Value = convtostr.Complex128(c) }
