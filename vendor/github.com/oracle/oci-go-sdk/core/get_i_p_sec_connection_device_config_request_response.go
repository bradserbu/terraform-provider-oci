// Copyright (c) 2016, 2017, Oracle and/or its affiliates. All rights reserved.
// Code generated. DO NOT EDIT.

package core

import (
	"github.com/oracle/oci-go-sdk/common"
	"net/http"
)

// GetIPSecConnectionDeviceConfigRequest wrapper for the GetIPSecConnectionDeviceConfig operation
type GetIPSecConnectionDeviceConfigRequest struct {

	// The OCID of the IPSec connection.
	IpscId *string `mandatory:"true" contributesTo:"path" name:"ipscId"`
}

func (request GetIPSecConnectionDeviceConfigRequest) String() string {
	return common.PointerString(request)
}

// GetIPSecConnectionDeviceConfigResponse wrapper for the GetIPSecConnectionDeviceConfig operation
type GetIPSecConnectionDeviceConfigResponse struct {

	// The underlying http response
	RawResponse *http.Response

	// The IpSecConnectionDeviceConfig instance
	IpSecConnectionDeviceConfig `presentIn:"body"`

	// For optimistic concurrency control. See `if-match`.
	Etag *string `presentIn:"header" name:"etag"`

	// Unique Oracle-assigned identifier for the request. If you need to contact Oracle about
	// a particular request, please provide the request ID.
	OpcRequestId *string `presentIn:"header" name:"opc-request-id"`
}

func (response GetIPSecConnectionDeviceConfigResponse) String() string {
	return common.PointerString(response)
}
