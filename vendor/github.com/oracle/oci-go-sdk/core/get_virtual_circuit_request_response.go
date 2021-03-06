// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetVirtualCircuitRequest wrapper for the GetVirtualCircuit operation
type GetVirtualCircuitRequest struct {

	// The OCID of the virtual circuit.
	VirtualCircuitId *string `mandatory:"true" contributesTo:"path" name:"virtualCircuitId"`
}

func (request GetVirtualCircuitRequest) String() string {
	return common.PointerString(request)
}

// GetVirtualCircuitResponse wrapper for the GetVirtualCircuit operation
type GetVirtualCircuitResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The VirtualCircuit instance
	VirtualCircuit `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetVirtualCircuitResponse) String() string {
	return common.PointerString(response)
}
