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

// checks if the Tenant type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Tenant{}

// Tenant struct for Tenant
type Tenant struct {
	Id                   *int64                                         `json:"id,omitempty"`
	Name                 *string                                        `json:"name,omitempty"`
	Description          *string                                        `json:"description,omitempty"`
	Subdomain            *string                                        `json:"subdomain,omitempty"`
	Currency             *string                                        `json:"currency,omitempty"`
	ExternalId           *string                                        `json:"externalId,omitempty"`
	CustomerNumber       *string                                        `json:"customerNumber,omitempty"`
	AccountNumber        *string                                        `json:"accountNumber,omitempty"`
	AccountName          *string                                        `json:"accountName,omitempty"`
	Active               *bool                                          `json:"active,omitempty"`
	Master               *bool                                          `json:"master,omitempty"`
	Role                 *ListTenants200ResponseAllOfAccountsInnerRole  `json:"role,omitempty"`
	Stats                *ListTenants200ResponseAllOfAccountsInnerStats `json:"stats,omitempty"`
	DateCreated          *time.Time                                     `json:"dateCreated,omitempty"`
	LastUpdated          *time.Time                                     `json:"lastUpdated,omitempty"`
	AdditionalProperties map[string]interface{}                         `json:",remain"`
}

type _Tenant Tenant

// NewTenant instantiates a new Tenant object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenant() *Tenant {
	this := Tenant{}
	return &this
}

// NewTenantWithDefaults instantiates a new Tenant object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantWithDefaults() *Tenant {
	this := Tenant{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Tenant) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *Tenant) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Tenant) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Tenant) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *Tenant) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Tenant) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Tenant) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// IsSetDescription returns a boolean if a field has been set.
func (o *Tenant) IsSetDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Tenant) SetDescription(v string) {
	o.Description = &v
}

// GetSubdomain returns the Subdomain field value if set, zero value otherwise.
func (o *Tenant) GetSubdomain() string {
	if o == nil || IsNil(o.Subdomain) {
		var ret string
		return ret
	}
	return *o.Subdomain
}

// GetSubdomainOk returns a tuple with the Subdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetSubdomainOk() (*string, bool) {
	if o == nil || IsNil(o.Subdomain) {
		return nil, false
	}
	return o.Subdomain, true
}

// IsSetSubdomain returns a boolean if a field has been set.
func (o *Tenant) IsSetSubdomain() bool {
	if o != nil && !IsNil(o.Subdomain) {
		return true
	}

	return false
}

// SetSubdomain gets a reference to the given string and assigns it to the Subdomain field.
func (o *Tenant) SetSubdomain(v string) {
	o.Subdomain = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *Tenant) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// IsSetCurrency returns a boolean if a field has been set.
func (o *Tenant) IsSetCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *Tenant) SetCurrency(v string) {
	o.Currency = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *Tenant) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// IsSetExternalId returns a boolean if a field has been set.
func (o *Tenant) IsSetExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *Tenant) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetCustomerNumber returns the CustomerNumber field value if set, zero value otherwise.
func (o *Tenant) GetCustomerNumber() string {
	if o == nil || IsNil(o.CustomerNumber) {
		var ret string
		return ret
	}
	return *o.CustomerNumber
}

// GetCustomerNumberOk returns a tuple with the CustomerNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetCustomerNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CustomerNumber) {
		return nil, false
	}
	return o.CustomerNumber, true
}

// IsSetCustomerNumber returns a boolean if a field has been set.
func (o *Tenant) IsSetCustomerNumber() bool {
	if o != nil && !IsNil(o.CustomerNumber) {
		return true
	}

	return false
}

// SetCustomerNumber gets a reference to the given string and assigns it to the CustomerNumber field.
func (o *Tenant) SetCustomerNumber(v string) {
	o.CustomerNumber = &v
}

// GetAccountNumber returns the AccountNumber field value if set, zero value otherwise.
func (o *Tenant) GetAccountNumber() string {
	if o == nil || IsNil(o.AccountNumber) {
		var ret string
		return ret
	}
	return *o.AccountNumber
}

// GetAccountNumberOk returns a tuple with the AccountNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetAccountNumberOk() (*string, bool) {
	if o == nil || IsNil(o.AccountNumber) {
		return nil, false
	}
	return o.AccountNumber, true
}

// IsSetAccountNumber returns a boolean if a field has been set.
func (o *Tenant) IsSetAccountNumber() bool {
	if o != nil && !IsNil(o.AccountNumber) {
		return true
	}

	return false
}

// SetAccountNumber gets a reference to the given string and assigns it to the AccountNumber field.
func (o *Tenant) SetAccountNumber(v string) {
	o.AccountNumber = &v
}

// GetAccountName returns the AccountName field value if set, zero value otherwise.
func (o *Tenant) GetAccountName() string {
	if o == nil || IsNil(o.AccountName) {
		var ret string
		return ret
	}
	return *o.AccountName
}

// GetAccountNameOk returns a tuple with the AccountName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetAccountNameOk() (*string, bool) {
	if o == nil || IsNil(o.AccountName) {
		return nil, false
	}
	return o.AccountName, true
}

// IsSetAccountName returns a boolean if a field has been set.
func (o *Tenant) IsSetAccountName() bool {
	if o != nil && !IsNil(o.AccountName) {
		return true
	}

	return false
}

// SetAccountName gets a reference to the given string and assigns it to the AccountName field.
func (o *Tenant) SetAccountName(v string) {
	o.AccountName = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *Tenant) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// IsSetActive returns a boolean if a field has been set.
func (o *Tenant) IsSetActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *Tenant) SetActive(v bool) {
	o.Active = &v
}

// GetMaster returns the Master field value if set, zero value otherwise.
func (o *Tenant) GetMaster() bool {
	if o == nil || IsNil(o.Master) {
		var ret bool
		return ret
	}
	return *o.Master
}

// GetMasterOk returns a tuple with the Master field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetMasterOk() (*bool, bool) {
	if o == nil || IsNil(o.Master) {
		return nil, false
	}
	return o.Master, true
}

// IsSetMaster returns a boolean if a field has been set.
func (o *Tenant) IsSetMaster() bool {
	if o != nil && !IsNil(o.Master) {
		return true
	}

	return false
}

// SetMaster gets a reference to the given bool and assigns it to the Master field.
func (o *Tenant) SetMaster(v bool) {
	o.Master = &v
}

// GetRole returns the Role field value if set, zero value otherwise.
func (o *Tenant) GetRole() ListTenants200ResponseAllOfAccountsInnerRole {
	if o == nil || IsNil(o.Role) {
		var ret ListTenants200ResponseAllOfAccountsInnerRole
		return ret
	}
	return *o.Role
}

// GetRoleOk returns a tuple with the Role field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetRoleOk() (*ListTenants200ResponseAllOfAccountsInnerRole, bool) {
	if o == nil || IsNil(o.Role) {
		return nil, false
	}
	return o.Role, true
}

// IsSetRole returns a boolean if a field has been set.
func (o *Tenant) IsSetRole() bool {
	if o != nil && !IsNil(o.Role) {
		return true
	}

	return false
}

// SetRole gets a reference to the given ListTenants200ResponseAllOfAccountsInnerRole and assigns it to the Role field.
func (o *Tenant) SetRole(v ListTenants200ResponseAllOfAccountsInnerRole) {
	o.Role = &v
}

// GetStats returns the Stats field value if set, zero value otherwise.
func (o *Tenant) GetStats() ListTenants200ResponseAllOfAccountsInnerStats {
	if o == nil || IsNil(o.Stats) {
		var ret ListTenants200ResponseAllOfAccountsInnerStats
		return ret
	}
	return *o.Stats
}

// GetStatsOk returns a tuple with the Stats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetStatsOk() (*ListTenants200ResponseAllOfAccountsInnerStats, bool) {
	if o == nil || IsNil(o.Stats) {
		return nil, false
	}
	return o.Stats, true
}

// IsSetStats returns a boolean if a field has been set.
func (o *Tenant) IsSetStats() bool {
	if o != nil && !IsNil(o.Stats) {
		return true
	}

	return false
}

// SetStats gets a reference to the given ListTenants200ResponseAllOfAccountsInnerStats and assigns it to the Stats field.
func (o *Tenant) SetStats(v ListTenants200ResponseAllOfAccountsInnerStats) {
	o.Stats = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *Tenant) GetDateCreated() time.Time {
	if o == nil || IsNil(o.DateCreated) {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateCreated) {
		return nil, false
	}
	return o.DateCreated, true
}

// IsSetDateCreated returns a boolean if a field has been set.
func (o *Tenant) IsSetDateCreated() bool {
	if o != nil && !IsNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *Tenant) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *Tenant) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tenant) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// IsSetLastUpdated returns a boolean if a field has been set.
func (o *Tenant) IsSetLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *Tenant) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

func (o Tenant) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Tenant) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Subdomain) {
		toSerialize["subdomain"] = o.Subdomain
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !IsNil(o.CustomerNumber) {
		toSerialize["customerNumber"] = o.CustomerNumber
	}
	if !IsNil(o.AccountNumber) {
		toSerialize["accountNumber"] = o.AccountNumber
	}
	if !IsNil(o.AccountName) {
		toSerialize["accountName"] = o.AccountName
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Master) {
		toSerialize["master"] = o.Master
	}
	if !IsNil(o.Role) {
		toSerialize["role"] = o.Role
	}
	if !IsNil(o.Stats) {
		toSerialize["stats"] = o.Stats
	}
	if !IsNil(o.DateCreated) {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *Tenant) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache
