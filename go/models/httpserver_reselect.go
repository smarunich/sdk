package models

// This file is auto-generated.
// Please contact avi-sdk@avinetworks.com for any change requests.

// HttpserverReselect httpserver reselect
// swagger:model HTTPServerReselect
type HttpserverReselect struct {

	// Enable HTTP request reselect when server responds with specific response codes.
	// Required: true
	Enabled bool `json:"enabled"`

	// Number of times to retry an HTTP request when server responds with configured status codes.
	NumRetries int32 `json:"num_retries,omitempty"`

	// Allow retry of non-idempotent HTTP requests.
	RetryNonidempotent bool `json:"retry_nonidempotent,omitempty"`

	// Server response codes which will trigger an HTTP request retry.
	SvrRespCode *HTTPReselectRespCode `json:"svr_resp_code,omitempty"`
}
