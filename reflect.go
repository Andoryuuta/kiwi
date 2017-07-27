package kiwi

import (
	"fmt"
	"reflect"
	"unsafe"
)

func getDataAddr(v reflect.Value) uintptr {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Slice:
		sh := (*reflect.SliceHeader)(unsafe.Pointer(v.UnsafeAddr()))
		return sh.Data
	default:
		return reflect.Indirect(v).UnsafeAddr()
	}
}

func getDataSize(v reflect.Value) uintptr {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}

	switch v.Kind() {
	case reflect.Slice, reflect.Array:
		size := v.Type().Elem().Size()
		return size * uintptr(v.Len())

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64,
		reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64, reflect.Bool:
		size := v.Type().Size()
		return size
	default:
		panic(fmt.Sprintf("dataSize: Unsupported type: %s", reflect.TypeOf(v).String()))
	}
}
