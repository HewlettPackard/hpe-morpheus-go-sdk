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

// checks if the ListInstances200ResponseAllOfInstancesInnerInterfacesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListInstances200ResponseAllOfInstancesInnerInterfacesInner{}

// ListInstances200ResponseAllOfInstancesInnerInterfacesInner struct for ListInstances200ResponseAllOfInstancesInnerInterfacesInner
type ListInstances200ResponseAllOfInstancesInnerInterfacesInner struct {
	Id                     *string                                                            `json:"id,omitempty"`
	Network                *ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork `json:"network,omitempty"`
	IpAddress              *string                                                            `json:"ipAddress,omitempty"`
	NetworkInterfaceTypeId *int64                                                             `json:"networkInterfaceTypeId,omitempty"`
	IpMode                 *string                                                            `json:"ipMode,omitempty"`
	AdditionalProperties   map[string]interface{}                                             `json:",remain"`
}

type _ListInstances200ResponseAllOfInstancesInnerInterfacesInner ListInstances200ResponseAllOfInstancesInnerInterfacesInner

// NewListInstances200ResponseAllOfInstancesInnerInterfacesInner instantiates a new ListInstances200ResponseAllOfInstancesInnerInterfacesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListInstances200ResponseAllOfInstancesInnerInterfacesInner() *ListInstances200ResponseAllOfInstancesInnerInterfacesInner {
	this := ListInstances200ResponseAllOfInstancesInnerInterfacesInner{}
	return &this
}

// NewListInstances200ResponseAllOfInstancesInnerInterfacesInnerWithDefaults instantiates a new ListInstances200ResponseAllOfInstancesInnerInterfacesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListInstances200ResponseAllOfInstancesInnerInterfacesInnerWithDefaults() *ListInstances200ResponseAllOfInstancesInnerInterfacesInner {
	this := ListInstances200ResponseAllOfInstancesInnerInterfacesInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInner) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInner) SetId(v string) {
	o.Id = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInner) GetNetwork() ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork {
	if o == nil || IsNil(o.Network) {
		var ret ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInner) GetNetworkOk() (*ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// IsSetNetwork returns a boolean if a field has been set.
func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInner) IsSetNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork and assigns it to the Network field.
func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInner) SetNetwork(v ListInstances200ResponseAllOfInstancesInnerInterfacesInnerNetwork) {
	o.Network = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInner) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInner) GetIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// IsSetIpAddress returns a boolean if a field has been set.
func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInner) IsSetIpAddress() bool {
	if o != nil && !IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInner) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetNetworkInterfaceTypeId returns the NetworkInterfaceTypeId field value if set, zero value otherwise.
func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInner) GetNetworkInterfaceTypeId() int64 {
	if o == nil || IsNil(o.NetworkInterfaceTypeId) {
		var ret int64
		return ret
	}
	return *o.NetworkInterfaceTypeId
}

// GetNetworkInterfaceTypeIdOk returns a tuple with the NetworkInterfaceTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInner) GetNetworkInterfaceTypeIdOk() (*int64, bool) {
	if o == nil || IsNil(o.NetworkInterfaceTypeId) {
		return nil, false
	}
	return o.NetworkInterfaceTypeId, true
}

// IsSetNetworkInterfaceTypeId returns a boolean if a field has been set.
func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInner) IsSetNetworkInterfaceTypeId() bool {
	if o != nil && !IsNil(o.NetworkInterfaceTypeId) {
		return true
	}

	return false
}

// SetNetworkInterfaceTypeId gets a reference to the given int64 and assigns it to the NetworkInterfaceTypeId field.
func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInner) SetNetworkInterfaceTypeId(v int64) {
	o.NetworkInterfaceTypeId = &v
}

// GetIpMode returns the IpMode field value if set, zero value otherwise.
func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInner) GetIpMode() string {
	if o == nil || IsNil(o.IpMode) {
		var ret string
		return ret
	}
	return *o.IpMode
}

// GetIpModeOk returns a tuple with the IpMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInner) GetIpModeOk() (*string, bool) {
	if o == nil || IsNil(o.IpMode) {
		return nil, false
	}
	return o.IpMode, true
}

// IsSetIpMode returns a boolean if a field has been set.
func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInner) IsSetIpMode() bool {
	if o != nil && !IsNil(o.IpMode) {
		return true
	}

	return false
}

// SetIpMode gets a reference to the given string and assigns it to the IpMode field.
func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInner) SetIpMode(v string) {
	o.IpMode = &v
}

func (o ListInstances200ResponseAllOfInstancesInnerInterfacesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListInstances200ResponseAllOfInstancesInnerInterfacesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.IpAddress) {
		toSerialize["ipAddress"] = o.IpAddress
	}
	if !IsNil(o.NetworkInterfaceTypeId) {
		toSerialize["networkInterfaceTypeId"] = o.NetworkInterfaceTypeId
	}
	if !IsNil(o.IpMode) {
		toSerialize["ipMode"] = o.IpMode
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListInstances200ResponseAllOfInstancesInnerInterfacesInner) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache
