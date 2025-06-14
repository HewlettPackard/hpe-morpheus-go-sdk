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

// checks if the ListHealth200ResponseAllOfHealthRabbit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ListHealth200ResponseAllOfHealthRabbit{}

// ListHealth200ResponseAllOfHealthRabbit struct for ListHealth200ResponseAllOfHealthRabbit
type ListHealth200ResponseAllOfHealthRabbit struct {
	Success              *bool                                               `json:"success,omitempty"`
	BusyQueues           []map[string]interface{}                            `json:"busyQueues,omitempty"`
	ErrorQueues          []map[string]interface{}                            `json:"errorQueues,omitempty"`
	Status               *string                                             `json:"status,omitempty"`
	Queues               []ListHealth200ResponseAllOfHealthRabbitQueuesInner `json:"queues,omitempty"`
	AdditionalProperties map[string]interface{}                              `json:",remain"`
}

type _ListHealth200ResponseAllOfHealthRabbit ListHealth200ResponseAllOfHealthRabbit

// NewListHealth200ResponseAllOfHealthRabbit instantiates a new ListHealth200ResponseAllOfHealthRabbit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewListHealth200ResponseAllOfHealthRabbit() *ListHealth200ResponseAllOfHealthRabbit {
	this := ListHealth200ResponseAllOfHealthRabbit{}
	return &this
}

// NewListHealth200ResponseAllOfHealthRabbitWithDefaults instantiates a new ListHealth200ResponseAllOfHealthRabbit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewListHealth200ResponseAllOfHealthRabbitWithDefaults() *ListHealth200ResponseAllOfHealthRabbit {
	this := ListHealth200ResponseAllOfHealthRabbit{}
	return &this
}

// GetSuccess returns the Success field value if set, zero value otherwise.
func (o *ListHealth200ResponseAllOfHealthRabbit) GetSuccess() bool {
	if o == nil || IsNil(o.Success) {
		var ret bool
		return ret
	}
	return *o.Success
}

// GetSuccessOk returns a tuple with the Success field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHealth200ResponseAllOfHealthRabbit) GetSuccessOk() (*bool, bool) {
	if o == nil || IsNil(o.Success) {
		return nil, false
	}
	return o.Success, true
}

// IsSetSuccess returns a boolean if a field has been set.
func (o *ListHealth200ResponseAllOfHealthRabbit) IsSetSuccess() bool {
	if o != nil && !IsNil(o.Success) {
		return true
	}

	return false
}

// SetSuccess gets a reference to the given bool and assigns it to the Success field.
func (o *ListHealth200ResponseAllOfHealthRabbit) SetSuccess(v bool) {
	o.Success = &v
}

// GetBusyQueues returns the BusyQueues field value if set, zero value otherwise.
func (o *ListHealth200ResponseAllOfHealthRabbit) GetBusyQueues() []map[string]interface{} {
	if o == nil || IsNil(o.BusyQueues) {
		var ret []map[string]interface{}
		return ret
	}
	return o.BusyQueues
}

// GetBusyQueuesOk returns a tuple with the BusyQueues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHealth200ResponseAllOfHealthRabbit) GetBusyQueuesOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.BusyQueues) {
		return nil, false
	}
	return o.BusyQueues, true
}

// IsSetBusyQueues returns a boolean if a field has been set.
func (o *ListHealth200ResponseAllOfHealthRabbit) IsSetBusyQueues() bool {
	if o != nil && !IsNil(o.BusyQueues) {
		return true
	}

	return false
}

// SetBusyQueues gets a reference to the given []map[string]interface{} and assigns it to the BusyQueues field.
func (o *ListHealth200ResponseAllOfHealthRabbit) SetBusyQueues(v []map[string]interface{}) {
	o.BusyQueues = v
}

// GetErrorQueues returns the ErrorQueues field value if set, zero value otherwise.
func (o *ListHealth200ResponseAllOfHealthRabbit) GetErrorQueues() []map[string]interface{} {
	if o == nil || IsNil(o.ErrorQueues) {
		var ret []map[string]interface{}
		return ret
	}
	return o.ErrorQueues
}

// GetErrorQueuesOk returns a tuple with the ErrorQueues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHealth200ResponseAllOfHealthRabbit) GetErrorQueuesOk() ([]map[string]interface{}, bool) {
	if o == nil || IsNil(o.ErrorQueues) {
		return nil, false
	}
	return o.ErrorQueues, true
}

// IsSetErrorQueues returns a boolean if a field has been set.
func (o *ListHealth200ResponseAllOfHealthRabbit) IsSetErrorQueues() bool {
	if o != nil && !IsNil(o.ErrorQueues) {
		return true
	}

	return false
}

// SetErrorQueues gets a reference to the given []map[string]interface{} and assigns it to the ErrorQueues field.
func (o *ListHealth200ResponseAllOfHealthRabbit) SetErrorQueues(v []map[string]interface{}) {
	o.ErrorQueues = v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ListHealth200ResponseAllOfHealthRabbit) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHealth200ResponseAllOfHealthRabbit) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// IsSetStatus returns a boolean if a field has been set.
func (o *ListHealth200ResponseAllOfHealthRabbit) IsSetStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *ListHealth200ResponseAllOfHealthRabbit) SetStatus(v string) {
	o.Status = &v
}

// GetQueues returns the Queues field value if set, zero value otherwise.
func (o *ListHealth200ResponseAllOfHealthRabbit) GetQueues() []ListHealth200ResponseAllOfHealthRabbitQueuesInner {
	if o == nil || IsNil(o.Queues) {
		var ret []ListHealth200ResponseAllOfHealthRabbitQueuesInner
		return ret
	}
	return o.Queues
}

// GetQueuesOk returns a tuple with the Queues field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ListHealth200ResponseAllOfHealthRabbit) GetQueuesOk() ([]ListHealth200ResponseAllOfHealthRabbitQueuesInner, bool) {
	if o == nil || IsNil(o.Queues) {
		return nil, false
	}
	return o.Queues, true
}

// IsSetQueues returns a boolean if a field has been set.
func (o *ListHealth200ResponseAllOfHealthRabbit) IsSetQueues() bool {
	if o != nil && !IsNil(o.Queues) {
		return true
	}

	return false
}

// SetQueues gets a reference to the given []ListHealth200ResponseAllOfHealthRabbitQueuesInner and assigns it to the Queues field.
func (o *ListHealth200ResponseAllOfHealthRabbit) SetQueues(v []ListHealth200ResponseAllOfHealthRabbitQueuesInner) {
	o.Queues = v
}

func (o ListHealth200ResponseAllOfHealthRabbit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ListHealth200ResponseAllOfHealthRabbit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Success) {
		toSerialize["success"] = o.Success
	}
	if !IsNil(o.BusyQueues) {
		toSerialize["busyQueues"] = o.BusyQueues
	}
	if !IsNil(o.ErrorQueues) {
		toSerialize["errorQueues"] = o.ErrorQueues
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Queues) {
		toSerialize["queues"] = o.Queues
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *ListHealth200ResponseAllOfHealthRabbit) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache
