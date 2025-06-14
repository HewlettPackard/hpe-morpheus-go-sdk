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

// checks if the UpdateCheckGroupsRequestCheckGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateCheckGroupsRequestCheckGroup{}

// UpdateCheckGroupsRequestCheckGroup Payload for creating a new monitoring check group
type UpdateCheckGroupsRequestCheckGroup struct {
	// Unique name scoped to your account for the check group
	Name *string `json:"name,omitempty"`
	// Optional description field
	Description *string `json:"description,omitempty"`
	// This specifies the minimum number of checks within the group that must be happy to keep the group from becoming unhealthy.
	MinHappy *int32 `json:"minHappy,omitempty"`
	// Used to determine if check should affect account wide availability calculations
	InUptime *bool `json:"inUptime,omitempty"`
	// Determines the maximum severity level this group can incur on an incident when failing
	Severity *string `json:"severity,omitempty"`
	// Used to determine if check group is active
	Active               *bool                  `json:"active,omitempty"`
	Checks               []int32                `json:"checks,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _UpdateCheckGroupsRequestCheckGroup UpdateCheckGroupsRequestCheckGroup

// NewUpdateCheckGroupsRequestCheckGroup instantiates a new UpdateCheckGroupsRequestCheckGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateCheckGroupsRequestCheckGroup() *UpdateCheckGroupsRequestCheckGroup {
	this := UpdateCheckGroupsRequestCheckGroup{}
	var minHappy int32 = 1
	this.MinHappy = &minHappy
	var inUptime bool = true
	this.InUptime = &inUptime
	var severity string = "critical"
	this.Severity = &severity
	var active bool = true
	this.Active = &active
	return &this
}

// NewUpdateCheckGroupsRequestCheckGroupWithDefaults instantiates a new UpdateCheckGroupsRequestCheckGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateCheckGroupsRequestCheckGroupWithDefaults() *UpdateCheckGroupsRequestCheckGroup {
	this := UpdateCheckGroupsRequestCheckGroup{}
	var minHappy int32 = 1
	this.MinHappy = &minHappy
	var inUptime bool = true
	this.InUptime = &inUptime
	var severity string = "critical"
	this.Severity = &severity
	var active bool = true
	this.Active = &active
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdateCheckGroupsRequestCheckGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCheckGroupsRequestCheckGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *UpdateCheckGroupsRequestCheckGroup) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdateCheckGroupsRequestCheckGroup) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *UpdateCheckGroupsRequestCheckGroup) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCheckGroupsRequestCheckGroup) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// IsSetDescription returns a boolean if a field has been set.
func (o *UpdateCheckGroupsRequestCheckGroup) IsSetDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *UpdateCheckGroupsRequestCheckGroup) SetDescription(v string) {
	o.Description = &v
}

// GetMinHappy returns the MinHappy field value if set, zero value otherwise.
func (o *UpdateCheckGroupsRequestCheckGroup) GetMinHappy() int32 {
	if o == nil || IsNil(o.MinHappy) {
		var ret int32
		return ret
	}
	return *o.MinHappy
}

// GetMinHappyOk returns a tuple with the MinHappy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCheckGroupsRequestCheckGroup) GetMinHappyOk() (*int32, bool) {
	if o == nil || IsNil(o.MinHappy) {
		return nil, false
	}
	return o.MinHappy, true
}

// IsSetMinHappy returns a boolean if a field has been set.
func (o *UpdateCheckGroupsRequestCheckGroup) IsSetMinHappy() bool {
	if o != nil && !IsNil(o.MinHappy) {
		return true
	}

	return false
}

// SetMinHappy gets a reference to the given int32 and assigns it to the MinHappy field.
func (o *UpdateCheckGroupsRequestCheckGroup) SetMinHappy(v int32) {
	o.MinHappy = &v
}

// GetInUptime returns the InUptime field value if set, zero value otherwise.
func (o *UpdateCheckGroupsRequestCheckGroup) GetInUptime() bool {
	if o == nil || IsNil(o.InUptime) {
		var ret bool
		return ret
	}
	return *o.InUptime
}

// GetInUptimeOk returns a tuple with the InUptime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCheckGroupsRequestCheckGroup) GetInUptimeOk() (*bool, bool) {
	if o == nil || IsNil(o.InUptime) {
		return nil, false
	}
	return o.InUptime, true
}

// IsSetInUptime returns a boolean if a field has been set.
func (o *UpdateCheckGroupsRequestCheckGroup) IsSetInUptime() bool {
	if o != nil && !IsNil(o.InUptime) {
		return true
	}

	return false
}

// SetInUptime gets a reference to the given bool and assigns it to the InUptime field.
func (o *UpdateCheckGroupsRequestCheckGroup) SetInUptime(v bool) {
	o.InUptime = &v
}

// GetSeverity returns the Severity field value if set, zero value otherwise.
func (o *UpdateCheckGroupsRequestCheckGroup) GetSeverity() string {
	if o == nil || IsNil(o.Severity) {
		var ret string
		return ret
	}
	return *o.Severity
}

// GetSeverityOk returns a tuple with the Severity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCheckGroupsRequestCheckGroup) GetSeverityOk() (*string, bool) {
	if o == nil || IsNil(o.Severity) {
		return nil, false
	}
	return o.Severity, true
}

// IsSetSeverity returns a boolean if a field has been set.
func (o *UpdateCheckGroupsRequestCheckGroup) IsSetSeverity() bool {
	if o != nil && !IsNil(o.Severity) {
		return true
	}

	return false
}

// SetSeverity gets a reference to the given string and assigns it to the Severity field.
func (o *UpdateCheckGroupsRequestCheckGroup) SetSeverity(v string) {
	o.Severity = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *UpdateCheckGroupsRequestCheckGroup) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCheckGroupsRequestCheckGroup) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// IsSetActive returns a boolean if a field has been set.
func (o *UpdateCheckGroupsRequestCheckGroup) IsSetActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *UpdateCheckGroupsRequestCheckGroup) SetActive(v bool) {
	o.Active = &v
}

// GetChecks returns the Checks field value if set, zero value otherwise.
func (o *UpdateCheckGroupsRequestCheckGroup) GetChecks() []int32 {
	if o == nil || IsNil(o.Checks) {
		var ret []int32
		return ret
	}
	return o.Checks
}

// GetChecksOk returns a tuple with the Checks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateCheckGroupsRequestCheckGroup) GetChecksOk() ([]int32, bool) {
	if o == nil || IsNil(o.Checks) {
		return nil, false
	}
	return o.Checks, true
}

// IsSetChecks returns a boolean if a field has been set.
func (o *UpdateCheckGroupsRequestCheckGroup) IsSetChecks() bool {
	if o != nil && !IsNil(o.Checks) {
		return true
	}

	return false
}

// SetChecks gets a reference to the given []int32 and assigns it to the Checks field.
func (o *UpdateCheckGroupsRequestCheckGroup) SetChecks(v []int32) {
	o.Checks = v
}

func (o UpdateCheckGroupsRequestCheckGroup) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateCheckGroupsRequestCheckGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.MinHappy) {
		toSerialize["minHappy"] = o.MinHappy
	}
	if !IsNil(o.InUptime) {
		toSerialize["inUptime"] = o.InUptime
	}
	if !IsNil(o.Severity) {
		toSerialize["severity"] = o.Severity
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	if !IsNil(o.Checks) {
		toSerialize["checks"] = o.Checks
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *UpdateCheckGroupsRequestCheckGroup) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache
