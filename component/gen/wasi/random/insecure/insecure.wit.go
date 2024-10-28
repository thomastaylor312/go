// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package insecure represents the imported interface "wasi:random/insecure@0.2.0".
package insecure

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
)

// GetInsecureRandomBytes represents the imported function "get-insecure-random-bytes".
//
//	get-insecure-random-bytes: func(len: u64) -> list<u8>
//
//go:nosplit
func GetInsecureRandomBytes(len_ uint64) (result cm.List[uint8]) {
	len0 := (uint64)(len_)
	wasmimport_GetInsecureRandomBytes((uint64)(len0), &result)
	return
}

// GetInsecureRandomU64 represents the imported function "get-insecure-random-u64".
//
//	get-insecure-random-u64: func() -> u64
//
//go:nosplit
func GetInsecureRandomU64() (result uint64) {
	result0 := wasmimport_GetInsecureRandomU64()
	result = (uint64)((uint64)(result0))
	return
}
