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
	"time"
)

// checks if the GetNetworkDhcpRelay200ResponseNetworkDhcpRelay type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetNetworkDhcpRelay200ResponseNetworkDhcpRelay{}

// GetNetworkDhcpRelay200ResponseNetworkDhcpRelay struct for GetNetworkDhcpRelay200ResponseNetworkDhcpRelay
type GetNetworkDhcpRelay200ResponseNetworkDhcpRelay struct {
	Id                   *int64                                       `json:"id,omitempty"`
	DateCreated          *time.Time                                   `json:"dateCreated,omitempty"`
	ProviderId           *string                                      `json:"providerId,omitempty"`
	LastUpdated          *time.Time                                   `json:"lastUpdated,omitempty"`
	Name                 *string                                      `json:"name,omitempty"`
	ExternalId           *string                                      `json:"externalId,omitempty"`
	ServerIpAddresses    []string                                     `json:"serverIpAddresses,omitempty"`
	Owner                *GetAlerts200ResponseAllOfChecksInnerAccount `json:"owner,omitempty"`
	NetworkServer        *GetAlerts200ResponseAllOfChecksInnerAccount `json:"networkServer,omitempty"`
	AdditionalProperties map[string]interface{}                       `json:",remain"`
}

type _GetNetworkDhcpRelay200ResponseNetworkDhcpRelay GetNetworkDhcpRelay200ResponseNetworkDhcpRelay

// NewGetNetworkDhcpRelay200ResponseNetworkDhcpRelay instantiates a new GetNetworkDhcpRelay200ResponseNetworkDhcpRelay object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetNetworkDhcpRelay200ResponseNetworkDhcpRelay() *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay {
	this := GetNetworkDhcpRelay200ResponseNetworkDhcpRelay{}
	return &this
}

// NewGetNetworkDhcpRelay200ResponseNetworkDhcpRelayWithDefaults instantiates a new GetNetworkDhcpRelay200ResponseNetworkDhcpRelay object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetNetworkDhcpRelay200ResponseNetworkDhcpRelayWithDefaults() *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay {
	this := GetNetworkDhcpRelay200ResponseNetworkDhcpRelay{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) SetId(v int64) {
	o.Id = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) GetDateCreated() time.Time {
	if o == nil || IsNil(o.DateCreated) {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateCreated) {
		return nil, false
	}
	return o.DateCreated, true
}

// IsSetDateCreated returns a boolean if a field has been set.
func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) IsSetDateCreated() bool {
	if o != nil && !IsNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetProviderId returns the ProviderId field value if set, zero value otherwise.
func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) GetProviderId() string {
	if o == nil || IsNil(o.ProviderId) {
		var ret string
		return ret
	}
	return *o.ProviderId
}

// GetProviderIdOk returns a tuple with the ProviderId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) GetProviderIdOk() (*string, bool) {
	if o == nil || IsNil(o.ProviderId) {
		return nil, false
	}
	return o.ProviderId, true
}

// IsSetProviderId returns a boolean if a field has been set.
func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) IsSetProviderId() bool {
	if o != nil && !IsNil(o.ProviderId) {
		return true
	}

	return false
}

// SetProviderId gets a reference to the given string and assigns it to the ProviderId field.
func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) SetProviderId(v string) {
	o.ProviderId = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// IsSetLastUpdated returns a boolean if a field has been set.
func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) IsSetLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) SetName(v string) {
	o.Name = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// IsSetExternalId returns a boolean if a field has been set.
func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) IsSetExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetServerIpAddresses returns the ServerIpAddresses field value if set, zero value otherwise.
func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) GetServerIpAddresses() []string {
	if o == nil || IsNil(o.ServerIpAddresses) {
		var ret []string
		return ret
	}
	return o.ServerIpAddresses
}

// GetServerIpAddressesOk returns a tuple with the ServerIpAddresses field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) GetServerIpAddressesOk() ([]string, bool) {
	if o == nil || IsNil(o.ServerIpAddresses) {
		return nil, false
	}
	return o.ServerIpAddresses, true
}

// IsSetServerIpAddresses returns a boolean if a field has been set.
func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) IsSetServerIpAddresses() bool {
	if o != nil && !IsNil(o.ServerIpAddresses) {
		return true
	}

	return false
}

// SetServerIpAddresses gets a reference to the given []string and assigns it to the ServerIpAddresses field.
func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) SetServerIpAddresses(v []string) {
	o.ServerIpAddresses = v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) GetOwner() GetAlerts200ResponseAllOfChecksInnerAccount {
	if o == nil || IsNil(o.Owner) {
		var ret GetAlerts200ResponseAllOfChecksInnerAccount
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) GetOwnerOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// IsSetOwner returns a boolean if a field has been set.
func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) IsSetOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given GetAlerts200ResponseAllOfChecksInnerAccount and assigns it to the Owner field.
func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) SetOwner(v GetAlerts200ResponseAllOfChecksInnerAccount) {
	o.Owner = &v
}

// GetNetworkServer returns the NetworkServer field value if set, zero value otherwise.
func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) GetNetworkServer() GetAlerts200ResponseAllOfChecksInnerAccount {
	if o == nil || IsNil(o.NetworkServer) {
		var ret GetAlerts200ResponseAllOfChecksInnerAccount
		return ret
	}
	return *o.NetworkServer
}

// GetNetworkServerOk returns a tuple with the NetworkServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) GetNetworkServerOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool) {
	if o == nil || IsNil(o.NetworkServer) {
		return nil, false
	}
	return o.NetworkServer, true
}

// IsSetNetworkServer returns a boolean if a field has been set.
func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) IsSetNetworkServer() bool {
	if o != nil && !IsNil(o.NetworkServer) {
		return true
	}

	return false
}

// SetNetworkServer gets a reference to the given GetAlerts200ResponseAllOfChecksInnerAccount and assigns it to the NetworkServer field.
func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) SetNetworkServer(v GetAlerts200ResponseAllOfChecksInnerAccount) {
	o.NetworkServer = &v
}

func (o GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.DateCreated) {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if !IsNil(o.ProviderId) {
		toSerialize["providerId"] = o.ProviderId
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !IsNil(o.ServerIpAddresses) {
		toSerialize["serverIpAddresses"] = o.ServerIpAddresses
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}
	if !IsNil(o.NetworkServer) {
		toSerialize["networkServer"] = o.NetworkServer
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *GetNetworkDhcpRelay200ResponseNetworkDhcpRelay) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache
