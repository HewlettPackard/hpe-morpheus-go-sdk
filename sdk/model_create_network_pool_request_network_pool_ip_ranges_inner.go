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

// checks if the CreateNetworkPoolRequestNetworkPoolIpRangesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateNetworkPoolRequestNetworkPoolIpRangesInner{}

// CreateNetworkPoolRequestNetworkPoolIpRangesInner struct for CreateNetworkPoolRequestNetworkPoolIpRangesInner
type CreateNetworkPoolRequestNetworkPoolIpRangesInner struct {
	// Starting IP Address
	StartAddress *string `json:"startAddress,omitempty"`
	// Ending IP Address
	EndAddress *string `json:"endAddress,omitempty"`
	// IPv6 Network CIDR
	CidrIPv6             *string                `json:"cidrIPv6,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _CreateNetworkPoolRequestNetworkPoolIpRangesInner CreateNetworkPoolRequestNetworkPoolIpRangesInner

// NewCreateNetworkPoolRequestNetworkPoolIpRangesInner instantiates a new CreateNetworkPoolRequestNetworkPoolIpRangesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateNetworkPoolRequestNetworkPoolIpRangesInner() *CreateNetworkPoolRequestNetworkPoolIpRangesInner {
	this := CreateNetworkPoolRequestNetworkPoolIpRangesInner{}
	return &this
}

// NewCreateNetworkPoolRequestNetworkPoolIpRangesInnerWithDefaults instantiates a new CreateNetworkPoolRequestNetworkPoolIpRangesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateNetworkPoolRequestNetworkPoolIpRangesInnerWithDefaults() *CreateNetworkPoolRequestNetworkPoolIpRangesInner {
	this := CreateNetworkPoolRequestNetworkPoolIpRangesInner{}
	return &this
}

// GetStartAddress returns the StartAddress field value if set, zero value otherwise.
func (o *CreateNetworkPoolRequestNetworkPoolIpRangesInner) GetStartAddress() string {
	if o == nil || IsNil(o.StartAddress) {
		var ret string
		return ret
	}
	return *o.StartAddress
}

// GetStartAddressOk returns a tuple with the StartAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkPoolRequestNetworkPoolIpRangesInner) GetStartAddressOk() (*string, bool) {
	if o == nil || IsNil(o.StartAddress) {
		return nil, false
	}
	return o.StartAddress, true
}

// IsSetStartAddress returns a boolean if a field has been set.
func (o *CreateNetworkPoolRequestNetworkPoolIpRangesInner) IsSetStartAddress() bool {
	if o != nil && !IsNil(o.StartAddress) {
		return true
	}

	return false
}

// SetStartAddress gets a reference to the given string and assigns it to the StartAddress field.
func (o *CreateNetworkPoolRequestNetworkPoolIpRangesInner) SetStartAddress(v string) {
	o.StartAddress = &v
}

// GetEndAddress returns the EndAddress field value if set, zero value otherwise.
func (o *CreateNetworkPoolRequestNetworkPoolIpRangesInner) GetEndAddress() string {
	if o == nil || IsNil(o.EndAddress) {
		var ret string
		return ret
	}
	return *o.EndAddress
}

// GetEndAddressOk returns a tuple with the EndAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkPoolRequestNetworkPoolIpRangesInner) GetEndAddressOk() (*string, bool) {
	if o == nil || IsNil(o.EndAddress) {
		return nil, false
	}
	return o.EndAddress, true
}

// IsSetEndAddress returns a boolean if a field has been set.
func (o *CreateNetworkPoolRequestNetworkPoolIpRangesInner) IsSetEndAddress() bool {
	if o != nil && !IsNil(o.EndAddress) {
		return true
	}

	return false
}

// SetEndAddress gets a reference to the given string and assigns it to the EndAddress field.
func (o *CreateNetworkPoolRequestNetworkPoolIpRangesInner) SetEndAddress(v string) {
	o.EndAddress = &v
}

// GetCidrIPv6 returns the CidrIPv6 field value if set, zero value otherwise.
func (o *CreateNetworkPoolRequestNetworkPoolIpRangesInner) GetCidrIPv6() string {
	if o == nil || IsNil(o.CidrIPv6) {
		var ret string
		return ret
	}
	return *o.CidrIPv6
}

// GetCidrIPv6Ok returns a tuple with the CidrIPv6 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateNetworkPoolRequestNetworkPoolIpRangesInner) GetCidrIPv6Ok() (*string, bool) {
	if o == nil || IsNil(o.CidrIPv6) {
		return nil, false
	}
	return o.CidrIPv6, true
}

// IsSetCidrIPv6 returns a boolean if a field has been set.
func (o *CreateNetworkPoolRequestNetworkPoolIpRangesInner) IsSetCidrIPv6() bool {
	if o != nil && !IsNil(o.CidrIPv6) {
		return true
	}

	return false
}

// SetCidrIPv6 gets a reference to the given string and assigns it to the CidrIPv6 field.
func (o *CreateNetworkPoolRequestNetworkPoolIpRangesInner) SetCidrIPv6(v string) {
	o.CidrIPv6 = &v
}

func (o CreateNetworkPoolRequestNetworkPoolIpRangesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateNetworkPoolRequestNetworkPoolIpRangesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartAddress) {
		toSerialize["startAddress"] = o.StartAddress
	}
	if !IsNil(o.EndAddress) {
		toSerialize["endAddress"] = o.EndAddress
	}
	if !IsNil(o.CidrIPv6) {
		toSerialize["cidrIPv6"] = o.CidrIPv6
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *CreateNetworkPoolRequestNetworkPoolIpRangesInner) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache
