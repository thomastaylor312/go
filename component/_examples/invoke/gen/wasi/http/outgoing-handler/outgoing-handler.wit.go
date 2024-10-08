// Code generated by wit-bindgen-go. DO NOT EDIT.

// Package outgoinghandler represents the imported interface "wasi:http/outgoing-handler@0.2.0".
package outgoinghandler

import (
	"github.com/bytecodealliance/wasm-tools-go/cm"
	"github.com/wasmCloud/component-sdk-go/_examples/invoke/gen/wasi/http/types"
)

// Handle represents the imported function "handle".
//
//	handle: func(request: outgoing-request, options: option<request-options>) -> result<future-incoming-response,
//	error-code>
//
//go:nosplit
func Handle(request types.OutgoingRequest, options cm.Option[types.RequestOptions]) (result cm.Result[ErrorCodeShape, types.FutureIncomingResponse, types.ErrorCode]) {
	request0 := cm.Reinterpret[uint32](request)
	options0, options1 := lower_OptionRequestOptions(options)
	wasmimport_Handle((uint32)(request0), (uint32)(options0), (uint32)(options1), &result)
	return
}

//go:wasmimport wasi:http/outgoing-handler@0.2.0 handle
//go:noescape
func wasmimport_Handle(request0 uint32, options0 uint32, options1 uint32, result *cm.Result[ErrorCodeShape, types.FutureIncomingResponse, types.ErrorCode])
