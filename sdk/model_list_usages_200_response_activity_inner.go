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

// checks if the ListUsages200ResponseActivityInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListUsages200ResponseActivityInner{}

// ListUsages200ResponseActivityInner struct for ListUsages200ResponseActivityInner
type ListUsages200ResponseActivityInner struct {
	Id                   *string                                        `json:"_id,omitempty"`
	Success              *bool                                          `json:"success,omitempty"`
	ActivityType         *string                                        `json:"activityType,omitempty"`
	Name                 *string                                        `json:"name,omitempty"`
	Message              *string                                        `json:"message,omitempty"`
	ObjectType           *string                                        `json:"objectType,omitempty"`
	ObjectId             *int64                                         `json:"objectId,omitempty"`
	User                 *GetAlerts200ResponseAllOfChecksInnerCreatedBy `json:"user,omitempty"`
	Ts                   *time.Time                                     `json:"ts,omitempty"`
	AdditionalProperties map[string]interface{}                         `json:",remain"`
}

type _ListUsages200ResponseActivityInner ListUsages200ResponseActivityInner

// NewListUsages200ResponseActivityInner instantiates a new ListUsages200ResponseActivityInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListUsages200ResponseActivityInner() *ListUsages200ResponseActivityInner {
	this := ListUsages200ResponseActivityInner{}
	return &this
}

// NewListUsages200ResponseActivityInnerWithDefaults instantiates a new ListUsages200ResponseActivityInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListUsages200ResponseActivityInnerWithDefaults() *ListUsages200ResponseActivityInner {
	this := ListUsages200ResponseActivityInner{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ListUsages200ResponseActivityInner) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUsages200ResponseActivityInner) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// IsSetId returns a boolean if a field has been set.
func (o *ListUsages200ResponseActivityInner) IsSetId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *ListUsages200ResponseActivityInner) SetId(v string) {
	o.Id = &v
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ListUsages200ResponseActivityInner) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUsages200ResponseActivityInner) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// IsSetSuccess returns a boolean if a field has been set.
func (o *ListUsages200ResponseActivityInner) IsSetSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ListUsages200ResponseActivityInner) SetSuccess(v bool) {
	o.Success = &v
}

// GetActivityType returns the ActivityType field value if set, zero value otherwise.
func (o *ListUsages200ResponseActivityInner) GetActivityType() string {
	if o == nil || IsNil(o.ActivityType) {
		var ret string
		return ret
	}
	return *o.ActivityType
}

// GetActivityTypeOk returns a tuple with the ActivityType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUsages200ResponseActivityInner) GetActivityTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ActivityType) {
		return nil, false
	}
	return o.ActivityType, true
}

// IsSetActivityType returns a boolean if a field has been set.
func (o *ListUsages200ResponseActivityInner) IsSetActivityType() bool {
	if o != nil && !IsNil(o.ActivityType) {
		return true
	}

	return false
}

// SetActivityType gets a reference to the given string and assigns it to the ActivityType field.
func (o *ListUsages200ResponseActivityInner) SetActivityType(v string) {
	o.ActivityType = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ListUsages200ResponseActivityInner) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUsages200ResponseActivityInner) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *ListUsages200ResponseActivityInner) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *ListUsages200ResponseActivityInner) SetName(v string) {
	o.Name = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *ListUsages200ResponseActivityInner) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUsages200ResponseActivityInner) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// IsSetMessage returns a boolean if a field has been set.
func (o *ListUsages200ResponseActivityInner) IsSetMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *ListUsages200ResponseActivityInner) SetMessage(v string) {
	o.Message = &v
}

// GetObjectType returns the ObjectType field value if set, zero value otherwise.
func (o *ListUsages200ResponseActivityInner) GetObjectType() string {
	if o == nil || IsNil(o.ObjectType) {
		var ret string
		return ret
	}
	return *o.ObjectType
}

// GetObjectTypeOk returns a tuple with the ObjectType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUsages200ResponseActivityInner) GetObjectTypeOk() (*string, bool) {
	if o == nil || IsNil(o.ObjectType) {
		return nil, false
	}
	return o.ObjectType, true
}

// IsSetObjectType returns a boolean if a field has been set.
func (o *ListUsages200ResponseActivityInner) IsSetObjectType() bool {
	if o != nil && !IsNil(o.ObjectType) {
		return true
	}

	return false
}

// SetObjectType gets a reference to the given string and assigns it to the ObjectType field.
func (o *ListUsages200ResponseActivityInner) SetObjectType(v string) {
	o.ObjectType = &v
}

// GetObjectId returns the ObjectId field value if set, zero value otherwise.
func (o *ListUsages200ResponseActivityInner) GetObjectId() int64 {
	if o == nil || IsNil(o.ObjectId) {
		var ret int64
		return ret
	}
	return *o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUsages200ResponseActivityInner) GetObjectIdOk() (*int64, bool) {
	if o == nil || IsNil(o.ObjectId) {
		return nil, false
	}
	return o.ObjectId, true
}

// IsSetObjectId returns a boolean if a field has been set.
func (o *ListUsages200ResponseActivityInner) IsSetObjectId() bool {
	if o != nil && !IsNil(o.ObjectId) {
		return true
	}

	return false
}

// SetObjectId gets a reference to the given int64 and assigns it to the ObjectId field.
func (o *ListUsages200ResponseActivityInner) SetObjectId(v int64) {
	o.ObjectId = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *ListUsages200ResponseActivityInner) GetUser() GetAlerts200ResponseAllOfChecksInnerCreatedBy {
	if o == nil || IsNil(o.User) {
		var ret GetAlerts200ResponseAllOfChecksInnerCreatedBy
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUsages200ResponseActivityInner) GetUserOk() (*GetAlerts200ResponseAllOfChecksInnerCreatedBy, bool) {
	if o == nil || IsNil(o.User) {
		return nil, false
	}
	return o.User, true
}

// IsSetUser returns a boolean if a field has been set.
func (o *ListUsages200ResponseActivityInner) IsSetUser() bool {
	if o != nil && !IsNil(o.User) {
		return true
	}

	return false
}

// SetUser gets a reference to the given GetAlerts200ResponseAllOfChecksInnerCreatedBy and assigns it to the User field.
func (o *ListUsages200ResponseActivityInner) SetUser(v GetAlerts200ResponseAllOfChecksInnerCreatedBy) {
	o.User = &v
}

// GetTs returns the Ts field value if set, zero value otherwise.
func (o *ListUsages200ResponseActivityInner) GetTs() time.Time {
	if o == nil || IsNil(o.Ts) {
		var ret time.Time
		return ret
	}
	return *o.Ts
}

// GetTsOk returns a tuple with the Ts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListUsages200ResponseActivityInner) GetTsOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Ts) {
		return nil, false
	}
	return o.Ts, true
}

// IsSetTs returns a boolean if a field has been set.
func (o *ListUsages200ResponseActivityInner) IsSetTs() bool {
	if o != nil && !IsNil(o.Ts) {
		return true
	}

	return false
}

// SetTs gets a reference to the given time.Time and assigns it to the Ts field.
func (o *ListUsages200ResponseActivityInner) SetTs(v time.Time) {
	o.Ts = &v
}

func (o ListUsages200ResponseActivityInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListUsages200ResponseActivityInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["_id"] = o.Id
	}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.ActivityType) {
		toSerialize["activityType"] = o.ActivityType
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.ObjectType) {
		toSerialize["objectType"] = o.ObjectType
	}
	if !IsNil(o.ObjectId) {
		toSerialize["objectId"] = o.ObjectId
	}
	if !IsNil(o.User) {
		toSerialize["user"] = o.User
	}
	if !IsNil(o.Ts) {
		toSerialize["ts"] = o.Ts
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListUsages200ResponseActivityInner) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache
