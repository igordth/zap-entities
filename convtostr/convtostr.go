package convtostr

// todo move to pkg

import (
	"fmt"
	"strconv"
)

func AnyStringer(a any) string {
	switch a.(type) {
	case string:
		return a.(string)
	}
	if str, ok := a.(fmt.Stringer); ok {
		return str.String()
	}
	return ""
}

func Bool(b bool) string {
	return strconv.FormatBool(b)
}

func Byte(bytes []byte) string {
	return string(bytes)
}

func Complex128(c complex128) string {
	return strconv.FormatComplex(c, 'f', -1, 128)
}

func Complex64(c complex64) string {
	return strconv.FormatComplex(complex128(c), 'f', -1, 64)
}

func Float64(f float64) string {
	return strconv.FormatFloat(f, 'f', -1, 64)
}

func Float32(f float32) string {
	return strconv.FormatFloat(float64(f), 'f', -1, 64)
}

func Int(i int) string {
	return strconv.Itoa(i)
}

func Int64(i int64) string {
	return strconv.FormatInt(i, 10)
}

func Int32(i int32) string {
	return strconv.FormatInt(int64(i), 10)
}

func Int16(i int16) string {
	return strconv.FormatInt(int64(i), 10)
}

func Int8(i int8) string {
	return strconv.FormatInt(int64(i), 10)
}

func Uint(u uint) string {
	return strconv.FormatUint(uint64(u), 10)
}

func Uint64(u uint64) string {
	return strconv.FormatUint(u, 10)
}

func Uint32(u uint32) string {
	return strconv.FormatUint(uint64(u), 10)
}

func Uint16(u uint16) string {
	return strconv.FormatUint(uint64(u), 10)
}

func Uint8(u uint8) string {
	return strconv.FormatUint(uint64(u), 10)
}

func Uintptr(u uintptr) string {
	return strconv.FormatUint(uint64(u), 10)
}
