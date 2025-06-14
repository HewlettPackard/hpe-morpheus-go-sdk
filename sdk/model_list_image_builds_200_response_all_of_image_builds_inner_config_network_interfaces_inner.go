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

// checks if the ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner{}

// ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner struct for ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner
type ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner struct {
	IpMode                     *string                                                                             `json:"ipMode,omitempty"`
	PrimaryInterface           *bool                                                                               `json:"primaryInterface,omitempty"`
	ShowNetworkPoolLabel       *bool                                                                               `json:"showNetworkPoolLabel,omitempty"`
	ShowNetworkDhcpLabel       *bool                                                                               `json:"showNetworkDhcpLabel,omitempty"`
	Network                    *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork `json:"network,omitempty"`
	NetworkInterfaceTypeId     *int64                                                                              `json:"networkInterfaceTypeId,omitempty"`
	NetworkInterfaceTypeIdName *string                                                                             `json:"networkInterfaceTypeIdName,omitempty"`
	AdditionalProperties       map[string]interface{}                                                              `json:",remain"`
}

type _ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner

// NewListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner instantiates a new ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner() *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner {
	this := ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner{}
	return &this
}

// NewListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerWithDefaults instantiates a new ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerWithDefaults() *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner {
	this := ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner{}
	return &this
}

// GetIpMode returns the IpMode field value if set, zero value otherwise.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) GetIpMode() string {
	if o == nil || IsNil(o.IpMode) {
		var ret string
		return ret
	}
	return *o.IpMode
}

// GetIpModeOk returns a tuple with the IpMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) GetIpModeOk() (*string, bool) {
	if o == nil || IsNil(o.IpMode) {
		return nil, false
	}
	return o.IpMode, true
}

// IsSetIpMode returns a boolean if a field has been set.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) IsSetIpMode() bool {
	if o != nil && !IsNil(o.IpMode) {
		return true
	}

	return false
}

// SetIpMode gets a reference to the given string and assigns it to the IpMode field.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) SetIpMode(v string) {
	o.IpMode = &v
}

// GetPrimaryInterface returns the PrimaryInterface field value if set, zero value otherwise.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) GetPrimaryInterface() bool {
	if o == nil || IsNil(o.PrimaryInterface) {
		var ret bool
		return ret
	}
	return *o.PrimaryInterface
}

// GetPrimaryInterfaceOk returns a tuple with the PrimaryInterface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) GetPrimaryInterfaceOk() (*bool, bool) {
	if o == nil || IsNil(o.PrimaryInterface) {
		return nil, false
	}
	return o.PrimaryInterface, true
}

// IsSetPrimaryInterface returns a boolean if a field has been set.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) IsSetPrimaryInterface() bool {
	if o != nil && !IsNil(o.PrimaryInterface) {
		return true
	}

	return false
}

// SetPrimaryInterface gets a reference to the given bool and assigns it to the PrimaryInterface field.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) SetPrimaryInterface(v bool) {
	o.PrimaryInterface = &v
}

// GetShowNetworkPoolLabel returns the ShowNetworkPoolLabel field value if set, zero value otherwise.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) GetShowNetworkPoolLabel() bool {
	if o == nil || IsNil(o.ShowNetworkPoolLabel) {
		var ret bool
		return ret
	}
	return *o.ShowNetworkPoolLabel
}

// GetShowNetworkPoolLabelOk returns a tuple with the ShowNetworkPoolLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) GetShowNetworkPoolLabelOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowNetworkPoolLabel) {
		return nil, false
	}
	return o.ShowNetworkPoolLabel, true
}

// IsSetShowNetworkPoolLabel returns a boolean if a field has been set.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) IsSetShowNetworkPoolLabel() bool {
	if o != nil && !IsNil(o.ShowNetworkPoolLabel) {
		return true
	}

	return false
}

// SetShowNetworkPoolLabel gets a reference to the given bool and assigns it to the ShowNetworkPoolLabel field.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) SetShowNetworkPoolLabel(v bool) {
	o.ShowNetworkPoolLabel = &v
}

// GetShowNetworkDhcpLabel returns the ShowNetworkDhcpLabel field value if set, zero value otherwise.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) GetShowNetworkDhcpLabel() bool {
	if o == nil || IsNil(o.ShowNetworkDhcpLabel) {
		var ret bool
		return ret
	}
	return *o.ShowNetworkDhcpLabel
}

// GetShowNetworkDhcpLabelOk returns a tuple with the ShowNetworkDhcpLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) GetShowNetworkDhcpLabelOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowNetworkDhcpLabel) {
		return nil, false
	}
	return o.ShowNetworkDhcpLabel, true
}

// IsSetShowNetworkDhcpLabel returns a boolean if a field has been set.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) IsSetShowNetworkDhcpLabel() bool {
	if o != nil && !IsNil(o.ShowNetworkDhcpLabel) {
		return true
	}

	return false
}

// SetShowNetworkDhcpLabel gets a reference to the given bool and assigns it to the ShowNetworkDhcpLabel field.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) SetShowNetworkDhcpLabel(v bool) {
	o.ShowNetworkDhcpLabel = &v
}

// GetNetwork returns the Network field value if set, zero value otherwise.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) GetNetwork() ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork {
	if o == nil || IsNil(o.Network) {
		var ret ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork
		return ret
	}
	return *o.Network
}

// GetNetworkOk returns a tuple with the Network field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) GetNetworkOk() (*ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork, bool) {
	if o == nil || IsNil(o.Network) {
		return nil, false
	}
	return o.Network, true
}

// IsSetNetwork returns a boolean if a field has been set.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) IsSetNetwork() bool {
	if o != nil && !IsNil(o.Network) {
		return true
	}

	return false
}

// SetNetwork gets a reference to the given ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork and assigns it to the Network field.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) SetNetwork(v ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInnerNetwork) {
	o.Network = &v
}

// GetNetworkInterfaceTypeId returns the NetworkInterfaceTypeId field value if set, zero value otherwise.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) GetNetworkInterfaceTypeId() int64 {
	if o == nil || IsNil(o.NetworkInterfaceTypeId) {
		var ret int64
		return ret
	}
	return *o.NetworkInterfaceTypeId
}

// GetNetworkInterfaceTypeIdOk returns a tuple with the NetworkInterfaceTypeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) GetNetworkInterfaceTypeIdOk() (*int64, bool) {
	if o == nil || IsNil(o.NetworkInterfaceTypeId) {
		return nil, false
	}
	return o.NetworkInterfaceTypeId, true
}

// IsSetNetworkInterfaceTypeId returns a boolean if a field has been set.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) IsSetNetworkInterfaceTypeId() bool {
	if o != nil && !IsNil(o.NetworkInterfaceTypeId) {
		return true
	}

	return false
}

// SetNetworkInterfaceTypeId gets a reference to the given int64 and assigns it to the NetworkInterfaceTypeId field.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) SetNetworkInterfaceTypeId(v int64) {
	o.NetworkInterfaceTypeId = &v
}

// GetNetworkInterfaceTypeIdName returns the NetworkInterfaceTypeIdName field value if set, zero value otherwise.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) GetNetworkInterfaceTypeIdName() string {
	if o == nil || IsNil(o.NetworkInterfaceTypeIdName) {
		var ret string
		return ret
	}
	return *o.NetworkInterfaceTypeIdName
}

// GetNetworkInterfaceTypeIdNameOk returns a tuple with the NetworkInterfaceTypeIdName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) GetNetworkInterfaceTypeIdNameOk() (*string, bool) {
	if o == nil || IsNil(o.NetworkInterfaceTypeIdName) {
		return nil, false
	}
	return o.NetworkInterfaceTypeIdName, true
}

// IsSetNetworkInterfaceTypeIdName returns a boolean if a field has been set.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) IsSetNetworkInterfaceTypeIdName() bool {
	if o != nil && !IsNil(o.NetworkInterfaceTypeIdName) {
		return true
	}

	return false
}

// SetNetworkInterfaceTypeIdName gets a reference to the given string and assigns it to the NetworkInterfaceTypeIdName field.
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) SetNetworkInterfaceTypeIdName(v string) {
	o.NetworkInterfaceTypeIdName = &v
}

func (o ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IpMode) {
		toSerialize["ipMode"] = o.IpMode
	}
	if !IsNil(o.PrimaryInterface) {
		toSerialize["primaryInterface"] = o.PrimaryInterface
	}
	if !IsNil(o.ShowNetworkPoolLabel) {
		toSerialize["showNetworkPoolLabel"] = o.ShowNetworkPoolLabel
	}
	if !IsNil(o.ShowNetworkDhcpLabel) {
		toSerialize["showNetworkDhcpLabel"] = o.ShowNetworkDhcpLabel
	}
	if !IsNil(o.Network) {
		toSerialize["network"] = o.Network
	}
	if !IsNil(o.NetworkInterfaceTypeId) {
		toSerialize["networkInterfaceTypeId"] = o.NetworkInterfaceTypeId
	}
	if !IsNil(o.NetworkInterfaceTypeIdName) {
		toSerialize["networkInterfaceTypeIdName"] = o.NetworkInterfaceTypeIdName
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListImageBuilds200ResponseAllOfImageBuildsInnerConfigNetworkInterfacesInner) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache
