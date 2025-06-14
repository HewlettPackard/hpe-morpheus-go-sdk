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

// checks if the NetworkDomain type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NetworkDomain{}

// NetworkDomain struct for NetworkDomain
type NetworkDomain struct {
	Id                   *int64                                                                  `json:"id,omitempty"`
	Name                 *string                                                                 `json:"name,omitempty"`
	Active               *bool                                                                   `json:"active,omitempty"`
	Fqdn                 *string                                                                 `json:"fqdn,omitempty"`
	Description          *string                                                                 `json:"description,omitempty"`
	Visibility           *string                                                                 `json:"visibility,omitempty"`
	DomainController     *bool                                                                   `json:"domainController,omitempty"`
	PublicZone           *bool                                                                   `json:"publicZone,omitempty"`
	DomainUsername       *string                                                                 `json:"domainUsername,omitempty"`
	DomainPassword       *string                                                                 `json:"domainPassword,omitempty"`
	RefType              *string                                                                 `json:"refType,omitempty"`
	RefId                *int64                                                                  `json:"refId,omitempty"`
	RefSource            *string                                                                 `json:"refSource,omitempty"`
	InternalId           *string                                                                 `json:"internalId,omitempty"`
	OuPath               *string                                                                 `json:"ouPath,omitempty"`
	DcServer             *string                                                                 `json:"dcServer,omitempty"`
	ZoneType             *string                                                                 `json:"zoneType,omitempty"`
	Dnssec               *string                                                                 `json:"dnssec,omitempty"`
	DomainSerial         *string                                                                 `json:"domainSerial,omitempty"`
	Account              *ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner `json:"account,omitempty"`
	Owner                *ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner `json:"owner,omitempty"`
	AdditionalProperties map[string]interface{}                                                  `json:",remain"`
}

type _NetworkDomain NetworkDomain

// NewNetworkDomain instantiates a new NetworkDomain object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNetworkDomain() *NetworkDomain {
	this := NetworkDomain{}
	return &this
}

// NewNetworkDomainWithDefaults instantiates a new NetworkDomain object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNetworkDomainWithDefaults() *NetworkDomain {
	this := NetworkDomain{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *NetworkDomain) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDomain) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *NetworkDomain) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *NetworkDomain) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *NetworkDomain) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDomain) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *NetworkDomain) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *NetworkDomain) SetName(v string) {
	o.Name = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *NetworkDomain) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDomain) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// IsSetActive returns a boolean if a field has been set.
func (o *NetworkDomain) IsSetActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *NetworkDomain) SetActive(v bool) {
	o.Active = &v
}

// GetFqdn returns the Fqdn field value if set, zero value otherwise.
func (o *NetworkDomain) GetFqdn() string {
	if o == nil || IsNil(o.Fqdn) {
		var ret string
		return ret
	}
	return *o.Fqdn
}

// GetFqdnOk returns a tuple with the Fqdn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDomain) GetFqdnOk() (*string, bool) {
	if o == nil || IsNil(o.Fqdn) {
		return nil, false
	}
	return o.Fqdn, true
}

// IsSetFqdn returns a boolean if a field has been set.
func (o *NetworkDomain) IsSetFqdn() bool {
	if o != nil && !IsNil(o.Fqdn) {
		return true
	}

	return false
}

// SetFqdn gets a reference to the given string and assigns it to the Fqdn field.
func (o *NetworkDomain) SetFqdn(v string) {
	o.Fqdn = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *NetworkDomain) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDomain) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// IsSetDescription returns a boolean if a field has been set.
func (o *NetworkDomain) IsSetDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *NetworkDomain) SetDescription(v string) {
	o.Description = &v
}

// GetVisibility returns the Visibility field value if set, zero value otherwise.
func (o *NetworkDomain) GetVisibility() string {
	if o == nil || IsNil(o.Visibility) {
		var ret string
		return ret
	}
	return *o.Visibility
}

// GetVisibilityOk returns a tuple with the Visibility field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDomain) GetVisibilityOk() (*string, bool) {
	if o == nil || IsNil(o.Visibility) {
		return nil, false
	}
	return o.Visibility, true
}

// IsSetVisibility returns a boolean if a field has been set.
func (o *NetworkDomain) IsSetVisibility() bool {
	if o != nil && !IsNil(o.Visibility) {
		return true
	}

	return false
}

// SetVisibility gets a reference to the given string and assigns it to the Visibility field.
func (o *NetworkDomain) SetVisibility(v string) {
	o.Visibility = &v
}

// GetDomainController returns the DomainController field value if set, zero value otherwise.
func (o *NetworkDomain) GetDomainController() bool {
	if o == nil || IsNil(o.DomainController) {
		var ret bool
		return ret
	}
	return *o.DomainController
}

// GetDomainControllerOk returns a tuple with the DomainController field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDomain) GetDomainControllerOk() (*bool, bool) {
	if o == nil || IsNil(o.DomainController) {
		return nil, false
	}
	return o.DomainController, true
}

// IsSetDomainController returns a boolean if a field has been set.
func (o *NetworkDomain) IsSetDomainController() bool {
	if o != nil && !IsNil(o.DomainController) {
		return true
	}

	return false
}

// SetDomainController gets a reference to the given bool and assigns it to the DomainController field.
func (o *NetworkDomain) SetDomainController(v bool) {
	o.DomainController = &v
}

// GetPublicZone returns the PublicZone field value if set, zero value otherwise.
func (o *NetworkDomain) GetPublicZone() bool {
	if o == nil || IsNil(o.PublicZone) {
		var ret bool
		return ret
	}
	return *o.PublicZone
}

// GetPublicZoneOk returns a tuple with the PublicZone field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDomain) GetPublicZoneOk() (*bool, bool) {
	if o == nil || IsNil(o.PublicZone) {
		return nil, false
	}
	return o.PublicZone, true
}

// IsSetPublicZone returns a boolean if a field has been set.
func (o *NetworkDomain) IsSetPublicZone() bool {
	if o != nil && !IsNil(o.PublicZone) {
		return true
	}

	return false
}

// SetPublicZone gets a reference to the given bool and assigns it to the PublicZone field.
func (o *NetworkDomain) SetPublicZone(v bool) {
	o.PublicZone = &v
}

// GetDomainUsername returns the DomainUsername field value if set, zero value otherwise.
func (o *NetworkDomain) GetDomainUsername() string {
	if o == nil || IsNil(o.DomainUsername) {
		var ret string
		return ret
	}
	return *o.DomainUsername
}

// GetDomainUsernameOk returns a tuple with the DomainUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDomain) GetDomainUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.DomainUsername) {
		return nil, false
	}
	return o.DomainUsername, true
}

// IsSetDomainUsername returns a boolean if a field has been set.
func (o *NetworkDomain) IsSetDomainUsername() bool {
	if o != nil && !IsNil(o.DomainUsername) {
		return true
	}

	return false
}

// SetDomainUsername gets a reference to the given string and assigns it to the DomainUsername field.
func (o *NetworkDomain) SetDomainUsername(v string) {
	o.DomainUsername = &v
}

// GetDomainPassword returns the DomainPassword field value if set, zero value otherwise.
func (o *NetworkDomain) GetDomainPassword() string {
	if o == nil || IsNil(o.DomainPassword) {
		var ret string
		return ret
	}
	return *o.DomainPassword
}

// GetDomainPasswordOk returns a tuple with the DomainPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDomain) GetDomainPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.DomainPassword) {
		return nil, false
	}
	return o.DomainPassword, true
}

// IsSetDomainPassword returns a boolean if a field has been set.
func (o *NetworkDomain) IsSetDomainPassword() bool {
	if o != nil && !IsNil(o.DomainPassword) {
		return true
	}

	return false
}

// SetDomainPassword gets a reference to the given string and assigns it to the DomainPassword field.
func (o *NetworkDomain) SetDomainPassword(v string) {
	o.DomainPassword = &v
}

// GetRefType returns the RefType field value if set, zero value otherwise.
func (o *NetworkDomain) GetRefType() string {
	if o == nil || IsNil(o.RefType) {
		var ret string
		return ret
	}
	return *o.RefType
}

// GetRefTypeOk returns a tuple with the RefType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDomain) GetRefTypeOk() (*string, bool) {
	if o == nil || IsNil(o.RefType) {
		return nil, false
	}
	return o.RefType, true
}

// IsSetRefType returns a boolean if a field has been set.
func (o *NetworkDomain) IsSetRefType() bool {
	if o != nil && !IsNil(o.RefType) {
		return true
	}

	return false
}

// SetRefType gets a reference to the given string and assigns it to the RefType field.
func (o *NetworkDomain) SetRefType(v string) {
	o.RefType = &v
}

// GetRefId returns the RefId field value if set, zero value otherwise.
func (o *NetworkDomain) GetRefId() int64 {
	if o == nil || IsNil(o.RefId) {
		var ret int64
		return ret
	}
	return *o.RefId
}

// GetRefIdOk returns a tuple with the RefId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDomain) GetRefIdOk() (*int64, bool) {
	if o == nil || IsNil(o.RefId) {
		return nil, false
	}
	return o.RefId, true
}

// IsSetRefId returns a boolean if a field has been set.
func (o *NetworkDomain) IsSetRefId() bool {
	if o != nil && !IsNil(o.RefId) {
		return true
	}

	return false
}

// SetRefId gets a reference to the given int64 and assigns it to the RefId field.
func (o *NetworkDomain) SetRefId(v int64) {
	o.RefId = &v
}

// GetRefSource returns the RefSource field value if set, zero value otherwise.
func (o *NetworkDomain) GetRefSource() string {
	if o == nil || IsNil(o.RefSource) {
		var ret string
		return ret
	}
	return *o.RefSource
}

// GetRefSourceOk returns a tuple with the RefSource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDomain) GetRefSourceOk() (*string, bool) {
	if o == nil || IsNil(o.RefSource) {
		return nil, false
	}
	return o.RefSource, true
}

// IsSetRefSource returns a boolean if a field has been set.
func (o *NetworkDomain) IsSetRefSource() bool {
	if o != nil && !IsNil(o.RefSource) {
		return true
	}

	return false
}

// SetRefSource gets a reference to the given string and assigns it to the RefSource field.
func (o *NetworkDomain) SetRefSource(v string) {
	o.RefSource = &v
}

// GetInternalId returns the InternalId field value if set, zero value otherwise.
func (o *NetworkDomain) GetInternalId() string {
	if o == nil || IsNil(o.InternalId) {
		var ret string
		return ret
	}
	return *o.InternalId
}

// GetInternalIdOk returns a tuple with the InternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDomain) GetInternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.InternalId) {
		return nil, false
	}
	return o.InternalId, true
}

// IsSetInternalId returns a boolean if a field has been set.
func (o *NetworkDomain) IsSetInternalId() bool {
	if o != nil && !IsNil(o.InternalId) {
		return true
	}

	return false
}

// SetInternalId gets a reference to the given string and assigns it to the InternalId field.
func (o *NetworkDomain) SetInternalId(v string) {
	o.InternalId = &v
}

// GetOuPath returns the OuPath field value if set, zero value otherwise.
func (o *NetworkDomain) GetOuPath() string {
	if o == nil || IsNil(o.OuPath) {
		var ret string
		return ret
	}
	return *o.OuPath
}

// GetOuPathOk returns a tuple with the OuPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDomain) GetOuPathOk() (*string, bool) {
	if o == nil || IsNil(o.OuPath) {
		return nil, false
	}
	return o.OuPath, true
}

// IsSetOuPath returns a boolean if a field has been set.
func (o *NetworkDomain) IsSetOuPath() bool {
	if o != nil && !IsNil(o.OuPath) {
		return true
	}

	return false
}

// SetOuPath gets a reference to the given string and assigns it to the OuPath field.
func (o *NetworkDomain) SetOuPath(v string) {
	o.OuPath = &v
}

// GetDcServer returns the DcServer field value if set, zero value otherwise.
func (o *NetworkDomain) GetDcServer() string {
	if o == nil || IsNil(o.DcServer) {
		var ret string
		return ret
	}
	return *o.DcServer
}

// GetDcServerOk returns a tuple with the DcServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDomain) GetDcServerOk() (*string, bool) {
	if o == nil || IsNil(o.DcServer) {
		return nil, false
	}
	return o.DcServer, true
}

// IsSetDcServer returns a boolean if a field has been set.
func (o *NetworkDomain) IsSetDcServer() bool {
	if o != nil && !IsNil(o.DcServer) {
		return true
	}

	return false
}

// SetDcServer gets a reference to the given string and assigns it to the DcServer field.
func (o *NetworkDomain) SetDcServer(v string) {
	o.DcServer = &v
}

// GetZoneType returns the ZoneType field value if set, zero value otherwise.
func (o *NetworkDomain) GetZoneType() string {
	if o == nil || IsNil(o.ZoneType) {
		var ret string
		return ret
	}
	return *o.ZoneType
}

// GetZoneTypeOk returns a tuple with the ZoneType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDomain) GetZoneTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ZoneType) {
		return nil, false
	}
	return o.ZoneType, true
}

// IsSetZoneType returns a boolean if a field has been set.
func (o *NetworkDomain) IsSetZoneType() bool {
	if o != nil && !IsNil(o.ZoneType) {
		return true
	}

	return false
}

// SetZoneType gets a reference to the given string and assigns it to the ZoneType field.
func (o *NetworkDomain) SetZoneType(v string) {
	o.ZoneType = &v
}

// GetDnssec returns the Dnssec field value if set, zero value otherwise.
func (o *NetworkDomain) GetDnssec() string {
	if o == nil || IsNil(o.Dnssec) {
		var ret string
		return ret
	}
	return *o.Dnssec
}

// GetDnssecOk returns a tuple with the Dnssec field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDomain) GetDnssecOk() (*string, bool) {
	if o == nil || IsNil(o.Dnssec) {
		return nil, false
	}
	return o.Dnssec, true
}

// IsSetDnssec returns a boolean if a field has been set.
func (o *NetworkDomain) IsSetDnssec() bool {
	if o != nil && !IsNil(o.Dnssec) {
		return true
	}

	return false
}

// SetDnssec gets a reference to the given string and assigns it to the Dnssec field.
func (o *NetworkDomain) SetDnssec(v string) {
	o.Dnssec = &v
}

// GetDomainSerial returns the DomainSerial field value if set, zero value otherwise.
func (o *NetworkDomain) GetDomainSerial() string {
	if o == nil || IsNil(o.DomainSerial) {
		var ret string
		return ret
	}
	return *o.DomainSerial
}

// GetDomainSerialOk returns a tuple with the DomainSerial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDomain) GetDomainSerialOk() (*string, bool) {
	if o == nil || IsNil(o.DomainSerial) {
		return nil, false
	}
	return o.DomainSerial, true
}

// IsSetDomainSerial returns a boolean if a field has been set.
func (o *NetworkDomain) IsSetDomainSerial() bool {
	if o != nil && !IsNil(o.DomainSerial) {
		return true
	}

	return false
}

// SetDomainSerial gets a reference to the given string and assigns it to the DomainSerial field.
func (o *NetworkDomain) SetDomainSerial(v string) {
	o.DomainSerial = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *NetworkDomain) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner {
	if o == nil || IsNil(o.Account) {
		var ret ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDomain) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// IsSetAccount returns a boolean if a field has been set.
func (o *NetworkDomain) IsSetAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner and assigns it to the Account field.
func (o *NetworkDomain) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner) {
	o.Account = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *NetworkDomain) GetOwner() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner {
	if o == nil || IsNil(o.Owner) {
		var ret ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NetworkDomain) GetOwnerOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool) {
	if o == nil || IsNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// IsSetOwner returns a boolean if a field has been set.
func (o *NetworkDomain) IsSetOwner() bool {
	if o != nil && !IsNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner and assigns it to the Owner field.
func (o *NetworkDomain) SetOwner(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner) {
	o.Owner = &v
}

func (o NetworkDomain) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NetworkDomain) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Fqdn) {
		toSerialize["fqdn"] = o.Fqdn
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Visibility) {
		toSerialize["visibility"] = o.Visibility
	}
	if !IsNil(o.DomainController) {
		toSerialize["domainController"] = o.DomainController
	}
	if !IsNil(o.PublicZone) {
		toSerialize["publicZone"] = o.PublicZone
	}
	if !IsNil(o.DomainUsername) {
		toSerialize["domainUsername"] = o.DomainUsername
	}
	if !IsNil(o.DomainPassword) {
		toSerialize["domainPassword"] = o.DomainPassword
	}
	if !IsNil(o.RefType) {
		toSerialize["refType"] = o.RefType
	}
	if !IsNil(o.RefId) {
		toSerialize["refId"] = o.RefId
	}
	if !IsNil(o.RefSource) {
		toSerialize["refSource"] = o.RefSource
	}
	if !IsNil(o.InternalId) {
		toSerialize["internalId"] = o.InternalId
	}
	if !IsNil(o.OuPath) {
		toSerialize["ouPath"] = o.OuPath
	}
	if !IsNil(o.DcServer) {
		toSerialize["dcServer"] = o.DcServer
	}
	if !IsNil(o.ZoneType) {
		toSerialize["zoneType"] = o.ZoneType
	}
	if !IsNil(o.Dnssec) {
		toSerialize["dnssec"] = o.Dnssec
	}
	if !IsNil(o.DomainSerial) {
		toSerialize["domainSerial"] = o.DomainSerial
	}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *NetworkDomain) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache
