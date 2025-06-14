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

// checks if the BillingInstances type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BillingInstances{}

// BillingInstances struct for BillingInstances
type BillingInstances struct {
	Price                *float32                                                        `json:"price,omitempty"`
	Cost                 *float32                                                        `json:"cost,omitempty"`
	StartDate            *time.Time                                                      `json:"startDate,omitempty"`
	EndDate              *time.Time                                                      `json:"endDate,omitempty"`
	Instances            []ListBillingInstances200ResponseAllOfBillingInfoInstancesInner `json:"instances,omitempty"`
	AdditionalProperties map[string]interface{}                                          `json:",remain"`
}

type _BillingInstances BillingInstances

// NewBillingInstances instantiates a new BillingInstances object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBillingInstances() *BillingInstances {
	this := BillingInstances{}
	return &this
}

// NewBillingInstancesWithDefaults instantiates a new BillingInstances object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingInstancesWithDefaults() *BillingInstances {
	this := BillingInstances{}
	return &this
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *BillingInstances) GetPrice() float32 {
	if o == nil || IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstances) GetPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// IsSetPrice returns a boolean if a field has been set.
func (o *BillingInstances) IsSetPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *BillingInstances) SetPrice(v float32) {
	o.Price = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *BillingInstances) GetCost() float32 {
	if o == nil || IsNil(o.Cost) {
		var ret float32
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstances) GetCostOk() (*float32, bool) {
	if o == nil || IsNil(o.Cost) {
		return nil, false
	}
	return o.Cost, true
}

// IsSetCost returns a boolean if a field has been set.
func (o *BillingInstances) IsSetCost() bool {
	if o != nil && !IsNil(o.Cost) {
		return true
	}

	return false
}

// SetCost gets a reference to the given float32 and assigns it to the Cost field.
func (o *BillingInstances) SetCost(v float32) {
	o.Cost = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *BillingInstances) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstances) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// IsSetStartDate returns a boolean if a field has been set.
func (o *BillingInstances) IsSetStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *BillingInstances) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *BillingInstances) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstances) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// IsSetEndDate returns a boolean if a field has been set.
func (o *BillingInstances) IsSetEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *BillingInstances) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetInstances returns the Instances field value if set, zero value otherwise.
func (o *BillingInstances) GetInstances() []ListBillingInstances200ResponseAllOfBillingInfoInstancesInner {
	if o == nil || IsNil(o.Instances) {
		var ret []ListBillingInstances200ResponseAllOfBillingInfoInstancesInner
		return ret
	}
	return o.Instances
}

// GetInstancesOk returns a tuple with the Instances field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BillingInstances) GetInstancesOk() ([]ListBillingInstances200ResponseAllOfBillingInfoInstancesInner, bool) {
	if o == nil || IsNil(o.Instances) {
		return nil, false
	}
	return o.Instances, true
}

// IsSetInstances returns a boolean if a field has been set.
func (o *BillingInstances) IsSetInstances() bool {
	if o != nil && !IsNil(o.Instances) {
		return true
	}

	return false
}

// SetInstances gets a reference to the given []ListBillingInstances200ResponseAllOfBillingInfoInstancesInner and assigns it to the Instances field.
func (o *BillingInstances) SetInstances(v []ListBillingInstances200ResponseAllOfBillingInfoInstancesInner) {
	o.Instances = v
}

func (o BillingInstances) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BillingInstances) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.Cost) {
		toSerialize["cost"] = o.Cost
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.Instances) {
		toSerialize["instances"] = o.Instances
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *BillingInstances) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache
