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

// checks if the Deployment type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Deployment{}

// Deployment struct for Deployment
type Deployment struct {
	Id                   *int64                                            `json:"id,omitempty"`
	Name                 *string                                           `json:"name,omitempty"`
	Description          *string                                           `json:"description,omitempty"`
	AccountId            *int64                                            `json:"accountId,omitempty"`
	ExternalId           *string                                           `json:"externalId,omitempty"`
	DateCreated          *time.Time                                        `json:"dateCreated,omitempty"`
	LastUpdated          *time.Time                                        `json:"lastUpdated,omitempty"`
	VersionCount         *int64                                            `json:"versionCount,omitempty"`
	Versions             []GetDeployment200ResponseDeploymentVersionsInner `json:"versions,omitempty"`
	AdditionalProperties map[string]interface{}                            `json:",remain"`
}

type _Deployment Deployment

// NewDeployment instantiates a new Deployment object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeployment() *Deployment {
	this := Deployment{}
	return &this
}

// NewDeploymentWithDefaults instantiates a new Deployment object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploymentWithDefaults() *Deployment {
	this := Deployment{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Deployment) GetId() int64 {
	if o == nil || IsNil(o.Id) {
		var ret int64
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetIdOk() (*int64, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *Deployment) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int64 and assigns it to the Id field.
func (o *Deployment) SetId(v int64) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Deployment) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *Deployment) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Deployment) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Deployment) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// IsSetDescription returns a boolean if a field has been set.
func (o *Deployment) IsSetDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Deployment) SetDescription(v string) {
	o.Description = &v
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *Deployment) GetAccountId() int64 {
	if o == nil || IsNil(o.AccountId) {
		var ret int64
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetAccountIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// IsSetAccountId returns a boolean if a field has been set.
func (o *Deployment) IsSetAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given int64 and assigns it to the AccountId field.
func (o *Deployment) SetAccountId(v int64) {
	o.AccountId = &v
}

// GetExternalId returns the ExternalId field value if set, zero value otherwise.
func (o *Deployment) GetExternalId() string {
	if o == nil || IsNil(o.ExternalId) {
		var ret string
		return ret
	}
	return *o.ExternalId
}

// GetExternalIdOk returns a tuple with the ExternalId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetExternalIdOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalId) {
		return nil, false
	}
	return o.ExternalId, true
}

// IsSetExternalId returns a boolean if a field has been set.
func (o *Deployment) IsSetExternalId() bool {
	if o != nil && !IsNil(o.ExternalId) {
		return true
	}

	return false
}

// SetExternalId gets a reference to the given string and assigns it to the ExternalId field.
func (o *Deployment) SetExternalId(v string) {
	o.ExternalId = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *Deployment) GetDateCreated() time.Time {
	if o == nil || IsNil(o.DateCreated) {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.DateCreated) {
		return nil, false
	}
	return o.DateCreated, true
}

// IsSetDateCreated returns a boolean if a field has been set.
func (o *Deployment) IsSetDateCreated() bool {
	if o != nil && !IsNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *Deployment) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *Deployment) GetLastUpdated() time.Time {
	if o == nil || IsNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || IsNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// IsSetLastUpdated returns a boolean if a field has been set.
func (o *Deployment) IsSetLastUpdated() bool {
	if o != nil && !IsNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *Deployment) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetVersionCount returns the VersionCount field value if set, zero value otherwise.
func (o *Deployment) GetVersionCount() int64 {
	if o == nil || IsNil(o.VersionCount) {
		var ret int64
		return ret
	}
	return *o.VersionCount
}

// GetVersionCountOk returns a tuple with the VersionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetVersionCountOk() (*int64, bool) {
	if o == nil || IsNil(o.VersionCount) {
		return nil, false
	}
	return o.VersionCount, true
}

// IsSetVersionCount returns a boolean if a field has been set.
func (o *Deployment) IsSetVersionCount() bool {
	if o != nil && !IsNil(o.VersionCount) {
		return true
	}

	return false
}

// SetVersionCount gets a reference to the given int64 and assigns it to the VersionCount field.
func (o *Deployment) SetVersionCount(v int64) {
	o.VersionCount = &v
}

// GetVersions returns the Versions field value if set, zero value otherwise.
func (o *Deployment) GetVersions() []GetDeployment200ResponseDeploymentVersionsInner {
	if o == nil || IsNil(o.Versions) {
		var ret []GetDeployment200ResponseDeploymentVersionsInner
		return ret
	}
	return o.Versions
}

// GetVersionsOk returns a tuple with the Versions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Deployment) GetVersionsOk() ([]GetDeployment200ResponseDeploymentVersionsInner, bool) {
	if o == nil || IsNil(o.Versions) {
		return nil, false
	}
	return o.Versions, true
}

// IsSetVersions returns a boolean if a field has been set.
func (o *Deployment) IsSetVersions() bool {
	if o != nil && !IsNil(o.Versions) {
		return true
	}

	return false
}

// SetVersions gets a reference to the given []GetDeployment200ResponseDeploymentVersionsInner and assigns it to the Versions field.
func (o *Deployment) SetVersions(v []GetDeployment200ResponseDeploymentVersionsInner) {
	o.Versions = v
}

func (o Deployment) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Deployment) ToMap() (map[string]interface{}, error) {
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
	if !IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !IsNil(o.ExternalId) {
		toSerialize["externalId"] = o.ExternalId
	}
	if !IsNil(o.DateCreated) {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if !IsNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !IsNil(o.VersionCount) {
		toSerialize["versionCount"] = o.VersionCount
	}
	if !IsNil(o.Versions) {
		toSerialize["versions"] = o.Versions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *Deployment) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache
