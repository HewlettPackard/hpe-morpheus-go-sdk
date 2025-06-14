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

// checks if the AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf{}

// AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf struct for AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf
type AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf struct {
	// id of the resource group to be used, can be prefixed with `pool-`. A resource pool group can be specified instead by prefixing its ID with `poolGroup-`.
	ResourcePoolId *string `json:"resourcePoolId,omitempty"`
	// Availability Options
	AvailabilityOptions *string `json:"availabilityOptions,omitempty"`
	// Availability Set
	AvailabilitySet *string `json:"availabilitySet,omitempty"`
	// Availability Zone
	AvailabilityZone *int64 `json:"availabilityZone,omitempty"`
	// Assign Public IP
	AzurefloatingIp *string `json:"azurefloatingIp,omitempty"`
	// Boot Diagnostics
	BootDiagnostics *string `json:"bootDiagnostics,omitempty"`
	// OS Guest Diagnostics
	OsGuestDiagnostics *string `json:"osGuestDiagnostics,omitempty"`
	// Diagnostics Storage Account
	DiagnosticsStorageAccount *string `json:"diagnosticsStorageAccount,omitempty"`
	// Create User
	CreateUser           *bool                  `json:"createUser,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf

// NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf instantiates a new AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf() *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf {
	this := AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf{}
	var createUser bool = true
	this.CreateUser = &createUser
	return &this
}

// NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOfWithDefaults instantiates a new AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOfWithDefaults() *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf {
	this := AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf{}
	var createUser bool = true
	this.CreateUser = &createUser
	return &this
}

// GetResourcePoolId returns the ResourcePoolId field value if set, zero value otherwise.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) GetResourcePoolId() string {
	if o == nil || IsNil(o.ResourcePoolId) {
		var ret string
		return ret
	}
	return *o.ResourcePoolId
}

// GetResourcePoolIdOk returns a tuple with the ResourcePoolId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) GetResourcePoolIdOk() (*string, bool) {
	if o == nil || IsNil(o.ResourcePoolId) {
		return nil, false
	}
	return o.ResourcePoolId, true
}

// IsSetResourcePoolId returns a boolean if a field has been set.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) IsSetResourcePoolId() bool {
	if o != nil && !IsNil(o.ResourcePoolId) {
		return true
	}

	return false
}

// SetResourcePoolId gets a reference to the given string and assigns it to the ResourcePoolId field.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) SetResourcePoolId(v string) {
	o.ResourcePoolId = &v
}

// GetAvailabilityOptions returns the AvailabilityOptions field value if set, zero value otherwise.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) GetAvailabilityOptions() string {
	if o == nil || IsNil(o.AvailabilityOptions) {
		var ret string
		return ret
	}
	return *o.AvailabilityOptions
}

// GetAvailabilityOptionsOk returns a tuple with the AvailabilityOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) GetAvailabilityOptionsOk() (*string, bool) {
	if o == nil || IsNil(o.AvailabilityOptions) {
		return nil, false
	}
	return o.AvailabilityOptions, true
}

// IsSetAvailabilityOptions returns a boolean if a field has been set.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) IsSetAvailabilityOptions() bool {
	if o != nil && !IsNil(o.AvailabilityOptions) {
		return true
	}

	return false
}

// SetAvailabilityOptions gets a reference to the given string and assigns it to the AvailabilityOptions field.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) SetAvailabilityOptions(v string) {
	o.AvailabilityOptions = &v
}

// GetAvailabilitySet returns the AvailabilitySet field value if set, zero value otherwise.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) GetAvailabilitySet() string {
	if o == nil || IsNil(o.AvailabilitySet) {
		var ret string
		return ret
	}
	return *o.AvailabilitySet
}

// GetAvailabilitySetOk returns a tuple with the AvailabilitySet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) GetAvailabilitySetOk() (*string, bool) {
	if o == nil || IsNil(o.AvailabilitySet) {
		return nil, false
	}
	return o.AvailabilitySet, true
}

// IsSetAvailabilitySet returns a boolean if a field has been set.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) IsSetAvailabilitySet() bool {
	if o != nil && !IsNil(o.AvailabilitySet) {
		return true
	}

	return false
}

// SetAvailabilitySet gets a reference to the given string and assigns it to the AvailabilitySet field.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) SetAvailabilitySet(v string) {
	o.AvailabilitySet = &v
}

// GetAvailabilityZone returns the AvailabilityZone field value if set, zero value otherwise.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) GetAvailabilityZone() int64 {
	if o == nil || IsNil(o.AvailabilityZone) {
		var ret int64
		return ret
	}
	return *o.AvailabilityZone
}

// GetAvailabilityZoneOk returns a tuple with the AvailabilityZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) GetAvailabilityZoneOk() (*int64, bool) {
	if o == nil || IsNil(o.AvailabilityZone) {
		return nil, false
	}
	return o.AvailabilityZone, true
}

// IsSetAvailabilityZone returns a boolean if a field has been set.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) IsSetAvailabilityZone() bool {
	if o != nil && !IsNil(o.AvailabilityZone) {
		return true
	}

	return false
}

// SetAvailabilityZone gets a reference to the given int64 and assigns it to the AvailabilityZone field.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) SetAvailabilityZone(v int64) {
	o.AvailabilityZone = &v
}

// GetAzurefloatingIp returns the AzurefloatingIp field value if set, zero value otherwise.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) GetAzurefloatingIp() string {
	if o == nil || IsNil(o.AzurefloatingIp) {
		var ret string
		return ret
	}
	return *o.AzurefloatingIp
}

// GetAzurefloatingIpOk returns a tuple with the AzurefloatingIp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) GetAzurefloatingIpOk() (*string, bool) {
	if o == nil || IsNil(o.AzurefloatingIp) {
		return nil, false
	}
	return o.AzurefloatingIp, true
}

// IsSetAzurefloatingIp returns a boolean if a field has been set.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) IsSetAzurefloatingIp() bool {
	if o != nil && !IsNil(o.AzurefloatingIp) {
		return true
	}

	return false
}

// SetAzurefloatingIp gets a reference to the given string and assigns it to the AzurefloatingIp field.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) SetAzurefloatingIp(v string) {
	o.AzurefloatingIp = &v
}

// GetBootDiagnostics returns the BootDiagnostics field value if set, zero value otherwise.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) GetBootDiagnostics() string {
	if o == nil || IsNil(o.BootDiagnostics) {
		var ret string
		return ret
	}
	return *o.BootDiagnostics
}

// GetBootDiagnosticsOk returns a tuple with the BootDiagnostics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) GetBootDiagnosticsOk() (*string, bool) {
	if o == nil || IsNil(o.BootDiagnostics) {
		return nil, false
	}
	return o.BootDiagnostics, true
}

// IsSetBootDiagnostics returns a boolean if a field has been set.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) IsSetBootDiagnostics() bool {
	if o != nil && !IsNil(o.BootDiagnostics) {
		return true
	}

	return false
}

// SetBootDiagnostics gets a reference to the given string and assigns it to the BootDiagnostics field.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) SetBootDiagnostics(v string) {
	o.BootDiagnostics = &v
}

// GetOsGuestDiagnostics returns the OsGuestDiagnostics field value if set, zero value otherwise.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) GetOsGuestDiagnostics() string {
	if o == nil || IsNil(o.OsGuestDiagnostics) {
		var ret string
		return ret
	}
	return *o.OsGuestDiagnostics
}

// GetOsGuestDiagnosticsOk returns a tuple with the OsGuestDiagnostics field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) GetOsGuestDiagnosticsOk() (*string, bool) {
	if o == nil || IsNil(o.OsGuestDiagnostics) {
		return nil, false
	}
	return o.OsGuestDiagnostics, true
}

// IsSetOsGuestDiagnostics returns a boolean if a field has been set.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) IsSetOsGuestDiagnostics() bool {
	if o != nil && !IsNil(o.OsGuestDiagnostics) {
		return true
	}

	return false
}

// SetOsGuestDiagnostics gets a reference to the given string and assigns it to the OsGuestDiagnostics field.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) SetOsGuestDiagnostics(v string) {
	o.OsGuestDiagnostics = &v
}

// GetDiagnosticsStorageAccount returns the DiagnosticsStorageAccount field value if set, zero value otherwise.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) GetDiagnosticsStorageAccount() string {
	if o == nil || IsNil(o.DiagnosticsStorageAccount) {
		var ret string
		return ret
	}
	return *o.DiagnosticsStorageAccount
}

// GetDiagnosticsStorageAccountOk returns a tuple with the DiagnosticsStorageAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) GetDiagnosticsStorageAccountOk() (*string, bool) {
	if o == nil || IsNil(o.DiagnosticsStorageAccount) {
		return nil, false
	}
	return o.DiagnosticsStorageAccount, true
}

// IsSetDiagnosticsStorageAccount returns a boolean if a field has been set.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) IsSetDiagnosticsStorageAccount() bool {
	if o != nil && !IsNil(o.DiagnosticsStorageAccount) {
		return true
	}

	return false
}

// SetDiagnosticsStorageAccount gets a reference to the given string and assigns it to the DiagnosticsStorageAccount field.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) SetDiagnosticsStorageAccount(v string) {
	o.DiagnosticsStorageAccount = &v
}

// GetCreateUser returns the CreateUser field value if set, zero value otherwise.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) GetCreateUser() bool {
	if o == nil || IsNil(o.CreateUser) {
		var ret bool
		return ret
	}
	return *o.CreateUser
}

// GetCreateUserOk returns a tuple with the CreateUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) GetCreateUserOk() (*bool, bool) {
	if o == nil || IsNil(o.CreateUser) {
		return nil, false
	}
	return o.CreateUser, true
}

// IsSetCreateUser returns a boolean if a field has been set.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) IsSetCreateUser() bool {
	if o != nil && !IsNil(o.CreateUser) {
		return true
	}

	return false
}

// SetCreateUser gets a reference to the given bool and assigns it to the CreateUser field.
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) SetCreateUser(v bool) {
	o.CreateUser = &v
}

func (o AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ResourcePoolId) {
		toSerialize["resourcePoolId"] = o.ResourcePoolId
	}
	if !IsNil(o.AvailabilityOptions) {
		toSerialize["availabilityOptions"] = o.AvailabilityOptions
	}
	if !IsNil(o.AvailabilitySet) {
		toSerialize["availabilitySet"] = o.AvailabilitySet
	}
	if !IsNil(o.AvailabilityZone) {
		toSerialize["availabilityZone"] = o.AvailabilityZone
	}
	if !IsNil(o.AzurefloatingIp) {
		toSerialize["azurefloatingIp"] = o.AzurefloatingIp
	}
	if !IsNil(o.BootDiagnostics) {
		toSerialize["bootDiagnostics"] = o.BootDiagnostics
	}
	if !IsNil(o.OsGuestDiagnostics) {
		toSerialize["osGuestDiagnostics"] = o.OsGuestDiagnostics
	}
	if !IsNil(o.DiagnosticsStorageAccount) {
		toSerialize["diagnosticsStorageAccount"] = o.DiagnosticsStorageAccount
	}
	if !IsNil(o.CreateUser) {
		toSerialize["createUser"] = o.CreateUser
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfigAnyOf) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache
