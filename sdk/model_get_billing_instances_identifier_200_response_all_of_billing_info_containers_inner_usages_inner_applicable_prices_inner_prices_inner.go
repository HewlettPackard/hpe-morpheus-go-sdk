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

// checks if the GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner{}

// GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner struct for GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner
type GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner struct {
	Type                 *string                `json:"type,omitempty"`
	PricePerUnit         *float32               `json:"pricePerUnit,omitempty"`
	CostPerUnit          *float32               `json:"costPerUnit,omitempty"`
	Cost                 *float32               `json:"cost,omitempty"`
	Price                *float32               `json:"price,omitempty"`
	Quantity             *int64                 `json:"quantity,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner

// NewGetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner instantiates a new GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner() *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner {
	this := GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner{}
	return &this
}

// NewGetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInnerWithDefaults instantiates a new GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInnerWithDefaults() *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner {
	this := GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// IsSetType returns a boolean if a field has been set.
func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner) IsSetType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner) SetType(v string) {
	o.Type = &v
}

// GetPricePerUnit returns the PricePerUnit field value if set, zero value otherwise.
func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner) GetPricePerUnit() float32 {
	if o == nil || IsNil(o.PricePerUnit) {
		var ret float32
		return ret
	}
	return *o.PricePerUnit
}

// GetPricePerUnitOk returns a tuple with the PricePerUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner) GetPricePerUnitOk() (*float32, bool) {
	if o == nil || IsNil(o.PricePerUnit) {
		return nil, false
	}
	return o.PricePerUnit, true
}

// IsSetPricePerUnit returns a boolean if a field has been set.
func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner) IsSetPricePerUnit() bool {
	if o != nil && !IsNil(o.PricePerUnit) {
		return true
	}

	return false
}

// SetPricePerUnit gets a reference to the given float32 and assigns it to the PricePerUnit field.
func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner) SetPricePerUnit(v float32) {
	o.PricePerUnit = &v
}

// GetCostPerUnit returns the CostPerUnit field value if set, zero value otherwise.
func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner) GetCostPerUnit() float32 {
	if o == nil || IsNil(o.CostPerUnit) {
		var ret float32
		return ret
	}
	return *o.CostPerUnit
}

// GetCostPerUnitOk returns a tuple with the CostPerUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner) GetCostPerUnitOk() (*float32, bool) {
	if o == nil || IsNil(o.CostPerUnit) {
		return nil, false
	}
	return o.CostPerUnit, true
}

// IsSetCostPerUnit returns a boolean if a field has been set.
func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner) IsSetCostPerUnit() bool {
	if o != nil && !IsNil(o.CostPerUnit) {
		return true
	}

	return false
}

// SetCostPerUnit gets a reference to the given float32 and assigns it to the CostPerUnit field.
func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner) SetCostPerUnit(v float32) {
	o.CostPerUnit = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner) GetCost() float32 {
	if o == nil || IsNil(o.Cost) {
		var ret float32
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner) GetCostOk() (*float32, bool) {
	if o == nil || IsNil(o.Cost) {
		return nil, false
	}
	return o.Cost, true
}

// IsSetCost returns a boolean if a field has been set.
func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner) IsSetCost() bool {
	if o != nil && !IsNil(o.Cost) {
		return true
	}

	return false
}

// SetCost gets a reference to the given float32 and assigns it to the Cost field.
func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner) SetCost(v float32) {
	o.Cost = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner) GetPrice() float32 {
	if o == nil || IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner) GetPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// IsSetPrice returns a boolean if a field has been set.
func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner) IsSetPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner) SetPrice(v float32) {
	o.Price = &v
}

// GetQuantity returns the Quantity field value if set, zero value otherwise.
func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner) GetQuantity() int64 {
	if o == nil || IsNil(o.Quantity) {
		var ret int64
		return ret
	}
	return *o.Quantity
}

// GetQuantityOk returns a tuple with the Quantity field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner) GetQuantityOk() (*int64, bool) {
	if o == nil || IsNil(o.Quantity) {
		return nil, false
	}
	return o.Quantity, true
}

// IsSetQuantity returns a boolean if a field has been set.
func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner) IsSetQuantity() bool {
	if o != nil && !IsNil(o.Quantity) {
		return true
	}

	return false
}

// SetQuantity gets a reference to the given int64 and assigns it to the Quantity field.
func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner) SetQuantity(v int64) {
	o.Quantity = &v
}

func (o GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.PricePerUnit) {
		toSerialize["pricePerUnit"] = o.PricePerUnit
	}
	if !IsNil(o.CostPerUnit) {
		toSerialize["costPerUnit"] = o.CostPerUnit
	}
	if !IsNil(o.Cost) {
		toSerialize["cost"] = o.Cost
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.Quantity) {
		toSerialize["quantity"] = o.Quantity
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *GetBillingInstancesIdentifier200ResponseAllOfBillingInfoContainersInnerUsagesInnerApplicablePricesInnerPricesInner) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache
