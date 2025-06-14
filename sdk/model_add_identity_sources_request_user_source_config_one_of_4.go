/*
Morpheus API

Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.

API version: 8.0.7
Contact: dev@morpheusdata.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
)

// checks if the AddIdentitySourcesRequestUserSourceConfigOneOf4 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddIdentitySourcesRequestUserSourceConfigOneOf4{}

// AddIdentitySourcesRequestUserSourceConfigOneOf4 OneLogin Configuration
type AddIdentitySourcesRequestUserSourceConfigOneOf4 struct {
	// OneLogin Subdomain
	Subdomain *string `json:"subdomain,omitempty"`
	// OneLogin Region
	Region *string `json:"region,omitempty"`
	// API Client Secret
	ClientSecret *string `json:"clientSecret,omitempty"`
	// API Client ID
	ClientId *string `json:"clientId,omitempty"`
	// Required Role
	RequiredRole         *string                `json:"requiredRole,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _AddIdentitySourcesRequestUserSourceConfigOneOf4 AddIdentitySourcesRequestUserSourceConfigOneOf4

// NewAddIdentitySourcesRequestUserSourceConfigOneOf4 instantiates a new AddIdentitySourcesRequestUserSourceConfigOneOf4 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddIdentitySourcesRequestUserSourceConfigOneOf4() *AddIdentitySourcesRequestUserSourceConfigOneOf4 {
	this := AddIdentitySourcesRequestUserSourceConfigOneOf4{}
	return &this
}

// NewAddIdentitySourcesRequestUserSourceConfigOneOf4WithDefaults instantiates a new AddIdentitySourcesRequestUserSourceConfigOneOf4 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddIdentitySourcesRequestUserSourceConfigOneOf4WithDefaults() *AddIdentitySourcesRequestUserSourceConfigOneOf4 {
	this := AddIdentitySourcesRequestUserSourceConfigOneOf4{}
	return &this
}

// GetSubdomain returns the Subdomain field value if set, zero value otherwise.
func (o *AddIdentitySourcesRequestUserSourceConfigOneOf4) GetSubdomain() string {
	if o == nil || IsNil(o.Subdomain) {
		var ret string
		return ret
	}
	return *o.Subdomain
}

// GetSubdomainOk returns a tuple with the Subdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIdentitySourcesRequestUserSourceConfigOneOf4) GetSubdomainOk() (*string, bool) {
	if o == nil || IsNil(o.Subdomain) {
		return nil, false
	}
	return o.Subdomain, true
}

// IsSetSubdomain returns a boolean if a field has been set.
func (o *AddIdentitySourcesRequestUserSourceConfigOneOf4) IsSetSubdomain() bool {
	if o != nil && !IsNil(o.Subdomain) {
		return true
	}

	return false
}

// SetSubdomain gets a reference to the given string and assigns it to the Subdomain field.
func (o *AddIdentitySourcesRequestUserSourceConfigOneOf4) SetSubdomain(v string) {
	o.Subdomain = &v
}

// GetRegion returns the Region field value if set, zero value otherwise.
func (o *AddIdentitySourcesRequestUserSourceConfigOneOf4) GetRegion() string {
	if o == nil || IsNil(o.Region) {
		var ret string
		return ret
	}
	return *o.Region
}

// GetRegionOk returns a tuple with the Region field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIdentitySourcesRequestUserSourceConfigOneOf4) GetRegionOk() (*string, bool) {
	if o == nil || IsNil(o.Region) {
		return nil, false
	}
	return o.Region, true
}

// IsSetRegion returns a boolean if a field has been set.
func (o *AddIdentitySourcesRequestUserSourceConfigOneOf4) IsSetRegion() bool {
	if o != nil && !IsNil(o.Region) {
		return true
	}

	return false
}

// SetRegion gets a reference to the given string and assigns it to the Region field.
func (o *AddIdentitySourcesRequestUserSourceConfigOneOf4) SetRegion(v string) {
	o.Region = &v
}

// GetClientSecret returns the ClientSecret field value if set, zero value otherwise.
func (o *AddIdentitySourcesRequestUserSourceConfigOneOf4) GetClientSecret() string {
	if o == nil || IsNil(o.ClientSecret) {
		var ret string
		return ret
	}
	return *o.ClientSecret
}

// GetClientSecretOk returns a tuple with the ClientSecret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIdentitySourcesRequestUserSourceConfigOneOf4) GetClientSecretOk() (*string, bool) {
	if o == nil || IsNil(o.ClientSecret) {
		return nil, false
	}
	return o.ClientSecret, true
}

// IsSetClientSecret returns a boolean if a field has been set.
func (o *AddIdentitySourcesRequestUserSourceConfigOneOf4) IsSetClientSecret() bool {
	if o != nil && !IsNil(o.ClientSecret) {
		return true
	}

	return false
}

// SetClientSecret gets a reference to the given string and assigns it to the ClientSecret field.
func (o *AddIdentitySourcesRequestUserSourceConfigOneOf4) SetClientSecret(v string) {
	o.ClientSecret = &v
}

// GetClientId returns the ClientId field value if set, zero value otherwise.
func (o *AddIdentitySourcesRequestUserSourceConfigOneOf4) GetClientId() string {
	if o == nil || IsNil(o.ClientId) {
		var ret string
		return ret
	}
	return *o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIdentitySourcesRequestUserSourceConfigOneOf4) GetClientIdOk() (*string, bool) {
	if o == nil || IsNil(o.ClientId) {
		return nil, false
	}
	return o.ClientId, true
}

// IsSetClientId returns a boolean if a field has been set.
func (o *AddIdentitySourcesRequestUserSourceConfigOneOf4) IsSetClientId() bool {
	if o != nil && !IsNil(o.ClientId) {
		return true
	}

	return false
}

// SetClientId gets a reference to the given string and assigns it to the ClientId field.
func (o *AddIdentitySourcesRequestUserSourceConfigOneOf4) SetClientId(v string) {
	o.ClientId = &v
}

// GetRequiredRole returns the RequiredRole field value if set, zero value otherwise.
func (o *AddIdentitySourcesRequestUserSourceConfigOneOf4) GetRequiredRole() string {
	if o == nil || IsNil(o.RequiredRole) {
		var ret string
		return ret
	}
	return *o.RequiredRole
}

// GetRequiredRoleOk returns a tuple with the RequiredRole field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddIdentitySourcesRequestUserSourceConfigOneOf4) GetRequiredRoleOk() (*string, bool) {
	if o == nil || IsNil(o.RequiredRole) {
		return nil, false
	}
	return o.RequiredRole, true
}

// IsSetRequiredRole returns a boolean if a field has been set.
func (o *AddIdentitySourcesRequestUserSourceConfigOneOf4) IsSetRequiredRole() bool {
	if o != nil && !IsNil(o.RequiredRole) {
		return true
	}

	return false
}

// SetRequiredRole gets a reference to the given string and assigns it to the RequiredRole field.
func (o *AddIdentitySourcesRequestUserSourceConfigOneOf4) SetRequiredRole(v string) {
	o.RequiredRole = &v
}

func (o AddIdentitySourcesRequestUserSourceConfigOneOf4) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddIdentitySourcesRequestUserSourceConfigOneOf4) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Subdomain) {
		toSerialize["subdomain"] = o.Subdomain
	}
	if !IsNil(o.Region) {
		toSerialize["region"] = o.Region
	}
	if !IsNil(o.ClientSecret) {
		toSerialize["clientSecret"] = o.ClientSecret
	}
	if !IsNil(o.ClientId) {
		toSerialize["clientId"] = o.ClientId
	}
	if !IsNil(o.RequiredRole) {
		toSerialize["requiredRole"] = o.RequiredRole
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *AddIdentitySourcesRequestUserSourceConfigOneOf4) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache
