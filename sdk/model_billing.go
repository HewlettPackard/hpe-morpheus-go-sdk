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

// checks if the Billing type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Billing{}

// Billing struct for Billing
type Billing struct {
	AccountId            *int64                                                    `json:"accountId,omitempty"`
	AccountUUID          *string                                                   `json:"accountUUID,omitempty"`
	Name                 *string                                                   `json:"name,omitempty"`
	StartDate            *time.Time                                                `json:"startDate,omitempty"`
	EndDate              *time.Time                                                `json:"endDate,omitempty"`
	PriceUnit            *string                                                   `json:"priceUnit,omitempty"`
	Price                *float32                                                  `json:"price,omitempty"`
	Cost                 *float32                                                  `json:"cost,omitempty"`
	Zones                []ListBillingAccount200ResponseAllOfBillingInfoZonesInner `json:"zones,omitempty"`
	AdditionalProperties map[string]interface{}                                    `json:",remain"`
}

type _Billing Billing

// NewBilling instantiates a new Billing object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBilling() *Billing {
	this := Billing{}
	return &this
}

// NewBillingWithDefaults instantiates a new Billing object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBillingWithDefaults() *Billing {
	this := Billing{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *Billing) GetAccountId() int64 {
	if o == nil || IsNil(o.AccountId) {
		var ret int64
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Billing) GetAccountIdOk() (*int64, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// IsSetAccountId returns a boolean if a field has been set.
func (o *Billing) IsSetAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given int64 and assigns it to the AccountId field.
func (o *Billing) SetAccountId(v int64) {
	o.AccountId = &v
}

// GetAccountUUID returns the AccountUUID field value if set, zero value otherwise.
func (o *Billing) GetAccountUUID() string {
	if o == nil || IsNil(o.AccountUUID) {
		var ret string
		return ret
	}
	return *o.AccountUUID
}

// GetAccountUUIDOk returns a tuple with the AccountUUID field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Billing) GetAccountUUIDOk() (*string, bool) {
	if o == nil || IsNil(o.AccountUUID) {
		return nil, false
	}
	return o.AccountUUID, true
}

// IsSetAccountUUID returns a boolean if a field has been set.
func (o *Billing) IsSetAccountUUID() bool {
	if o != nil && !IsNil(o.AccountUUID) {
		return true
	}

	return false
}

// SetAccountUUID gets a reference to the given string and assigns it to the AccountUUID field.
func (o *Billing) SetAccountUUID(v string) {
	o.AccountUUID = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Billing) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Billing) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *Billing) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Billing) SetName(v string) {
	o.Name = &v
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *Billing) GetStartDate() time.Time {
	if o == nil || IsNil(o.StartDate) {
		var ret time.Time
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Billing) GetStartDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// IsSetStartDate returns a boolean if a field has been set.
func (o *Billing) IsSetStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given time.Time and assigns it to the StartDate field.
func (o *Billing) SetStartDate(v time.Time) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *Billing) GetEndDate() time.Time {
	if o == nil || IsNil(o.EndDate) {
		var ret time.Time
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Billing) GetEndDateOk() (*time.Time, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// IsSetEndDate returns a boolean if a field has been set.
func (o *Billing) IsSetEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given time.Time and assigns it to the EndDate field.
func (o *Billing) SetEndDate(v time.Time) {
	o.EndDate = &v
}

// GetPriceUnit returns the PriceUnit field value if set, zero value otherwise.
func (o *Billing) GetPriceUnit() string {
	if o == nil || IsNil(o.PriceUnit) {
		var ret string
		return ret
	}
	return *o.PriceUnit
}

// GetPriceUnitOk returns a tuple with the PriceUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Billing) GetPriceUnitOk() (*string, bool) {
	if o == nil || IsNil(o.PriceUnit) {
		return nil, false
	}
	return o.PriceUnit, true
}

// IsSetPriceUnit returns a boolean if a field has been set.
func (o *Billing) IsSetPriceUnit() bool {
	if o != nil && !IsNil(o.PriceUnit) {
		return true
	}

	return false
}

// SetPriceUnit gets a reference to the given string and assigns it to the PriceUnit field.
func (o *Billing) SetPriceUnit(v string) {
	o.PriceUnit = &v
}

// GetPrice returns the Price field value if set, zero value otherwise.
func (o *Billing) GetPrice() float32 {
	if o == nil || IsNil(o.Price) {
		var ret float32
		return ret
	}
	return *o.Price
}

// GetPriceOk returns a tuple with the Price field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Billing) GetPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.Price) {
		return nil, false
	}
	return o.Price, true
}

// IsSetPrice returns a boolean if a field has been set.
func (o *Billing) IsSetPrice() bool {
	if o != nil && !IsNil(o.Price) {
		return true
	}

	return false
}

// SetPrice gets a reference to the given float32 and assigns it to the Price field.
func (o *Billing) SetPrice(v float32) {
	o.Price = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *Billing) GetCost() float32 {
	if o == nil || IsNil(o.Cost) {
		var ret float32
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Billing) GetCostOk() (*float32, bool) {
	if o == nil || IsNil(o.Cost) {
		return nil, false
	}
	return o.Cost, true
}

// IsSetCost returns a boolean if a field has been set.
func (o *Billing) IsSetCost() bool {
	if o != nil && !IsNil(o.Cost) {
		return true
	}

	return false
}

// SetCost gets a reference to the given float32 and assigns it to the Cost field.
func (o *Billing) SetCost(v float32) {
	o.Cost = &v
}

// GetZones returns the Zones field value if set, zero value otherwise.
func (o *Billing) GetZones() []ListBillingAccount200ResponseAllOfBillingInfoZonesInner {
	if o == nil || IsNil(o.Zones) {
		var ret []ListBillingAccount200ResponseAllOfBillingInfoZonesInner
		return ret
	}
	return o.Zones
}

// GetZonesOk returns a tuple with the Zones field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Billing) GetZonesOk() ([]ListBillingAccount200ResponseAllOfBillingInfoZonesInner, bool) {
	if o == nil || IsNil(o.Zones) {
		return nil, false
	}
	return o.Zones, true
}

// IsSetZones returns a boolean if a field has been set.
func (o *Billing) IsSetZones() bool {
	if o != nil && !IsNil(o.Zones) {
		return true
	}

	return false
}

// SetZones gets a reference to the given []ListBillingAccount200ResponseAllOfBillingInfoZonesInner and assigns it to the Zones field.
func (o *Billing) SetZones(v []ListBillingAccount200ResponseAllOfBillingInfoZonesInner) {
	o.Zones = v
}

func (o Billing) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Billing) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["accountId"] = o.AccountId
	}
	if !IsNil(o.AccountUUID) {
		toSerialize["accountUUID"] = o.AccountUUID
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.PriceUnit) {
		toSerialize["priceUnit"] = o.PriceUnit
	}
	if !IsNil(o.Price) {
		toSerialize["price"] = o.Price
	}
	if !IsNil(o.Cost) {
		toSerialize["cost"] = o.Cost
	}
	if !IsNil(o.Zones) {
		toSerialize["zones"] = o.Zones
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *Billing) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache
