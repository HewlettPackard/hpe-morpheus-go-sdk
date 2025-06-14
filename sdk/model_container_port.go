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

// checks if the ContainerPort type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContainerPort{}

// ContainerPort struct for ContainerPort
type ContainerPort struct {
	Id                   *int64                 `json:"id,omitempty"`
	Index                *int64                 `json:"index,omitempty"`
	External             *int64                 `json:"external,omitempty"`
	Internal             *int64                 `json:"internal,omitempty"`
	DisplayName          *string                `json:"displayName,omitempty"`
	PrimaryPort          *bool                  `json:"primaryPort,omitempty"`
	Export               *bool                  `json:"export,omitempty"`
	Visible              *bool                  `json:"visible,omitempty"`
	ExportName           *string                `json:"exportName,omitempty"`
	LoadBalanceProtocol  *string                `json:"loadBalanceProtocol,omitempty"`
	LoadBalance          *bool                  `json:"loadBalance,omitempty"`
	Protocol             *string                `json:"protocol,omitempty"`
	Link                 *bool                  `json:"link,omitempty"`
	ExternalIp           *string                `json:"externalIp,omitempty"`
	InternalIp           *string                `json:"internalIp,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _ContainerPort ContainerPort

// NewContainerPort instantiates a new ContainerPort object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContainerPort() *ContainerPort {
	this := ContainerPort{}
	return &this
}

// NewContainerPortWithDefaults instantiates a new ContainerPort object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContainerPortWithDefaults() *ContainerPort {
	this := ContainerPort{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ContainerPort) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerPort) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *ContainerPort) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *ContainerPort) SetId(v int64) {
	o.Id = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *ContainerPort) GetIndex() int64 {
	if o == nil || IsNil(o.Index) {
		var ret int64
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerPort) GetIndexOk() (*int64, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// IsSetIndex returns a boolean if a field has been set.
func (o *ContainerPort) IsSetIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int64 and assigns it to the Index field.
func (o *ContainerPort) SetIndex(v int64) {
	o.Index = &v
}

// GetExternal returns the External field value if set, zero value otherwise.
func (o *ContainerPort) GetExternal() int64 {
	if o == nil || IsNil(o.External) {
		var ret int64
		return ret
	}
	return *o.External
}

// GetExternalOk returns a tuple with the External field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerPort) GetExternalOk() (*int64, bool) {
	if o == nil || IsNil(o.External) {
		return nil, false
	}
	return o.External, true
}

// IsSetExternal returns a boolean if a field has been set.
func (o *ContainerPort) IsSetExternal() bool {
	if o != nil && !IsNil(o.External) {
		return true
	}

	return false
}

// SetExternal gets a reference to the given int64 and assigns it to the External field.
func (o *ContainerPort) SetExternal(v int64) {
	o.External = &v
}

// GetInternal returns the Internal field value if set, zero value otherwise.
func (o *ContainerPort) GetInternal() int64 {
	if o == nil || IsNil(o.Internal) {
		var ret int64
		return ret
	}
	return *o.Internal
}

// GetInternalOk returns a tuple with the Internal field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerPort) GetInternalOk() (*int64, bool) {
	if o == nil || IsNil(o.Internal) {
		return nil, false
	}
	return o.Internal, true
}

// IsSetInternal returns a boolean if a field has been set.
func (o *ContainerPort) IsSetInternal() bool {
	if o != nil && !IsNil(o.Internal) {
		return true
	}

	return false
}

// SetInternal gets a reference to the given int64 and assigns it to the Internal field.
func (o *ContainerPort) SetInternal(v int64) {
	o.Internal = &v
}

// GetDisplayName returns the DisplayName field value if set, zero value otherwise.
func (o *ContainerPort) GetDisplayName() string {
	if o == nil || IsNil(o.DisplayName) {
		var ret string
		return ret
	}
	return *o.DisplayName
}

// GetDisplayNameOk returns a tuple with the DisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerPort) GetDisplayNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayName) {
		return nil, false
	}
	return o.DisplayName, true
}

// IsSetDisplayName returns a boolean if a field has been set.
func (o *ContainerPort) IsSetDisplayName() bool {
	if o != nil && !IsNil(o.DisplayName) {
		return true
	}

	return false
}

// SetDisplayName gets a reference to the given string and assigns it to the DisplayName field.
func (o *ContainerPort) SetDisplayName(v string) {
	o.DisplayName = &v
}

// GetPrimaryPort returns the PrimaryPort field value if set, zero value otherwise.
func (o *ContainerPort) GetPrimaryPort() bool {
	if o == nil || IsNil(o.PrimaryPort) {
		var ret bool
		return ret
	}
	return *o.PrimaryPort
}

// GetPrimaryPortOk returns a tuple with the PrimaryPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerPort) GetPrimaryPortOk() (*bool, bool) {
	if o == nil || IsNil(o.PrimaryPort) {
		return nil, false
	}
	return o.PrimaryPort, true
}

// IsSetPrimaryPort returns a boolean if a field has been set.
func (o *ContainerPort) IsSetPrimaryPort() bool {
	if o != nil && !IsNil(o.PrimaryPort) {
		return true
	}

	return false
}

// SetPrimaryPort gets a reference to the given bool and assigns it to the PrimaryPort field.
func (o *ContainerPort) SetPrimaryPort(v bool) {
	o.PrimaryPort = &v
}

// GetExport returns the Export field value if set, zero value otherwise.
func (o *ContainerPort) GetExport() bool {
	if o == nil || IsNil(o.Export) {
		var ret bool
		return ret
	}
	return *o.Export
}

// GetExportOk returns a tuple with the Export field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerPort) GetExportOk() (*bool, bool) {
	if o == nil || IsNil(o.Export) {
		return nil, false
	}
	return o.Export, true
}

// IsSetExport returns a boolean if a field has been set.
func (o *ContainerPort) IsSetExport() bool {
	if o != nil && !IsNil(o.Export) {
		return true
	}

	return false
}

// SetExport gets a reference to the given bool and assigns it to the Export field.
func (o *ContainerPort) SetExport(v bool) {
	o.Export = &v
}

// GetVisible returns the Visible field value if set, zero value otherwise.
func (o *ContainerPort) GetVisible() bool {
	if o == nil || IsNil(o.Visible) {
		var ret bool
		return ret
	}
	return *o.Visible
}

// GetVisibleOk returns a tuple with the Visible field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerPort) GetVisibleOk() (*bool, bool) {
	if o == nil || IsNil(o.Visible) {
		return nil, false
	}
	return o.Visible, true
}

// IsSetVisible returns a boolean if a field has been set.
func (o *ContainerPort) IsSetVisible() bool {
	if o != nil && !IsNil(o.Visible) {
		return true
	}

	return false
}

// SetVisible gets a reference to the given bool and assigns it to the Visible field.
func (o *ContainerPort) SetVisible(v bool) {
	o.Visible = &v
}

// GetExportName returns the ExportName field value if set, zero value otherwise.
func (o *ContainerPort) GetExportName() string {
	if o == nil || IsNil(o.ExportName) {
		var ret string
		return ret
	}
	return *o.ExportName
}

// GetExportNameOk returns a tuple with the ExportName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerPort) GetExportNameOk() (*string, bool) {
	if o == nil || IsNil(o.ExportName) {
		return nil, false
	}
	return o.ExportName, true
}

// IsSetExportName returns a boolean if a field has been set.
func (o *ContainerPort) IsSetExportName() bool {
	if o != nil && !IsNil(o.ExportName) {
		return true
	}

	return false
}

// SetExportName gets a reference to the given string and assigns it to the ExportName field.
func (o *ContainerPort) SetExportName(v string) {
	o.ExportName = &v
}

// GetLoadBalanceProtocol returns the LoadBalanceProtocol field value if set, zero value otherwise.
func (o *ContainerPort) GetLoadBalanceProtocol() string {
	if o == nil || IsNil(o.LoadBalanceProtocol) {
		var ret string
		return ret
	}
	return *o.LoadBalanceProtocol
}

// GetLoadBalanceProtocolOk returns a tuple with the LoadBalanceProtocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerPort) GetLoadBalanceProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.LoadBalanceProtocol) {
		return nil, false
	}
	return o.LoadBalanceProtocol, true
}

// IsSetLoadBalanceProtocol returns a boolean if a field has been set.
func (o *ContainerPort) IsSetLoadBalanceProtocol() bool {
	if o != nil && !IsNil(o.LoadBalanceProtocol) {
		return true
	}

	return false
}

// SetLoadBalanceProtocol gets a reference to the given string and assigns it to the LoadBalanceProtocol field.
func (o *ContainerPort) SetLoadBalanceProtocol(v string) {
	o.LoadBalanceProtocol = &v
}

// GetLoadBalance returns the LoadBalance field value if set, zero value otherwise.
func (o *ContainerPort) GetLoadBalance() bool {
	if o == nil || IsNil(o.LoadBalance) {
		var ret bool
		return ret
	}
	return *o.LoadBalance
}

// GetLoadBalanceOk returns a tuple with the LoadBalance field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerPort) GetLoadBalanceOk() (*bool, bool) {
	if o == nil || IsNil(o.LoadBalance) {
		return nil, false
	}
	return o.LoadBalance, true
}

// IsSetLoadBalance returns a boolean if a field has been set.
func (o *ContainerPort) IsSetLoadBalance() bool {
	if o != nil && !IsNil(o.LoadBalance) {
		return true
	}

	return false
}

// SetLoadBalance gets a reference to the given bool and assigns it to the LoadBalance field.
func (o *ContainerPort) SetLoadBalance(v bool) {
	o.LoadBalance = &v
}

// GetProtocol returns the Protocol field value if set, zero value otherwise.
func (o *ContainerPort) GetProtocol() string {
	if o == nil || IsNil(o.Protocol) {
		var ret string
		return ret
	}
	return *o.Protocol
}

// GetProtocolOk returns a tuple with the Protocol field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerPort) GetProtocolOk() (*string, bool) {
	if o == nil || IsNil(o.Protocol) {
		return nil, false
	}
	return o.Protocol, true
}

// IsSetProtocol returns a boolean if a field has been set.
func (o *ContainerPort) IsSetProtocol() bool {
	if o != nil && !IsNil(o.Protocol) {
		return true
	}

	return false
}

// SetProtocol gets a reference to the given string and assigns it to the Protocol field.
func (o *ContainerPort) SetProtocol(v string) {
	o.Protocol = &v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *ContainerPort) GetLink() bool {
	if o == nil || IsNil(o.Link) {
		var ret bool
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerPort) GetLinkOk() (*bool, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// IsSetLink returns a boolean if a field has been set.
func (o *ContainerPort) IsSetLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given bool and assigns it to the Link field.
func (o *ContainerPort) SetLink(v bool) {
	o.Link = &v
}

// GetExternalIp returns the ExternalIp field value if set, zero value otherwise.
func (o *ContainerPort) GetExternalIp() string {
	if o == nil || IsNil(o.ExternalIp) {
		var ret string
		return ret
	}
	return *o.ExternalIp
}

// GetExternalIpOk returns a tuple with the ExternalIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerPort) GetExternalIpOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalIp) {
		return nil, false
	}
	return o.ExternalIp, true
}

// IsSetExternalIp returns a boolean if a field has been set.
func (o *ContainerPort) IsSetExternalIp() bool {
	if o != nil && !IsNil(o.ExternalIp) {
		return true
	}

	return false
}

// SetExternalIp gets a reference to the given string and assigns it to the ExternalIp field.
func (o *ContainerPort) SetExternalIp(v string) {
	o.ExternalIp = &v
}

// GetInternalIp returns the InternalIp field value if set, zero value otherwise.
func (o *ContainerPort) GetInternalIp() string {
	if o == nil || IsNil(o.InternalIp) {
		var ret string
		return ret
	}
	return *o.InternalIp
}

// GetInternalIpOk returns a tuple with the InternalIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContainerPort) GetInternalIpOk() (*string, bool) {
	if o == nil || IsNil(o.InternalIp) {
		return nil, false
	}
	return o.InternalIp, true
}

// IsSetInternalIp returns a boolean if a field has been set.
func (o *ContainerPort) IsSetInternalIp() bool {
	if o != nil && !IsNil(o.InternalIp) {
		return true
	}

	return false
}

// SetInternalIp gets a reference to the given string and assigns it to the InternalIp field.
func (o *ContainerPort) SetInternalIp(v string) {
	o.InternalIp = &v
}

func (o ContainerPort) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContainerPort) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	if !IsNil(o.External) {
		toSerialize["external"] = o.External
	}
	if !IsNil(o.Internal) {
		toSerialize["internal"] = o.Internal
	}
	if !IsNil(o.DisplayName) {
		toSerialize["displayName"] = o.DisplayName
	}
	if !IsNil(o.PrimaryPort) {
		toSerialize["primaryPort"] = o.PrimaryPort
	}
	if !IsNil(o.Export) {
		toSerialize["export"] = o.Export
	}
	if !IsNil(o.Visible) {
		toSerialize["visible"] = o.Visible
	}
	if !IsNil(o.ExportName) {
		toSerialize["exportName"] = o.ExportName
	}
	if !IsNil(o.LoadBalanceProtocol) {
		toSerialize["loadBalanceProtocol"] = o.LoadBalanceProtocol
	}
	if !IsNil(o.LoadBalance) {
		toSerialize["loadBalance"] = o.LoadBalance
	}
	if !IsNil(o.Protocol) {
		toSerialize["protocol"] = o.Protocol
	}
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	if !IsNil(o.ExternalIp) {
		toSerialize["externalIp"] = o.ExternalIp
	}
	if !IsNil(o.InternalIp) {
		toSerialize["internalIp"] = o.InternalIp
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ContainerPort) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache
