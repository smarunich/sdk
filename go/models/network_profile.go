package models

// This file is auto-generated.
// Please contact avi-sdk@avinetworks.com for any change requests.

// NetworkProfile network profile
// swagger:model NetworkProfile
type NetworkProfile struct {

	// User defined description for the object.
	Description string `json:"description,omitempty"`

	// The name of the network profile.
	// Required: true
	Name string `json:"name"`

	// Placeholder for description of property profile of obj type NetworkProfile field type str  type object
	// Required: true
	Profile *NetworkProfileUnion `json:"profile"`

	//  It is a reference to an object of type Tenant.
	TenantRef string `json:"tenant_ref,omitempty"`

	// url
	// Read Only: true
	URL string `json:"url,omitempty"`

	// UUID of the network profile.
	UUID string `json:"uuid,omitempty"`
}
