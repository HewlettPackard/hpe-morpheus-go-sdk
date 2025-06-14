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

// checks if the UpdatePricesRequestPrice type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdatePricesRequestPrice{}

// UpdatePricesRequestPrice struct for UpdatePricesRequestPrice
type UpdatePricesRequestPrice struct {
	// Price name
	Name *string `json:"name,omitempty"`
	// Price code, must be unique
	Code    *string                       `json:"code,omitempty"`
	Account *AddPricesRequestPriceAccount `json:"account,omitempty"`
	// Restricts query to only load only prices with specified priceType. * `fixed` - Everything * `compute` - Memory + CPU * `memory` - Memory * `cores` - Cores * `storage` - Storage * `datastore` - Datastore * `platform` - Platform * `software` - Software * `load_balancer` - Load Balancer * `load_balancer_virtual_server` - Load Balancer Virtual Server
	PriceType *string `json:"priceType,omitempty"`
	// The unit of pricing
	PriceUnit *string `json:"priceUnit,omitempty"`
	// Indicates when to incur charge
	IncurCharges *string `json:"incurCharges,omitempty"`
	// ISO Currency code
	Currency *string `json:"currency,omitempty"`
	// Cost
	Cost *float32 `json:"cost,omitempty"`
	// Price adjustment type
	MarkupType *string `json:"markupType,omitempty"`
	// Amount for `fixed` price adjustment type
	Markup *float32 `json:"markup,omitempty"`
	// Percent for `percent` price adjustment type
	MarkupPercent *float32 `json:"markupPercent,omitempty"`
	// Custom price for `custom` price adjustment type
	CustomPrice *float32 `json:"customPrice,omitempty"`
	// Platform.  Required for `platform` price type
	Platform *string `json:"platform,omitempty"`
	// Software.  Required for software price type
	Software   *string                          `json:"software,omitempty"`
	VolumeType *AddPricesRequestPriceVolumeType `json:"volumeType,omitempty"`
	Datastore  *AddPricesRequestPriceDatastore  `json:"datastore,omitempty"`
	// Apply price across clouds, optional true/false flag for datastore price type
	CrossCloudApply      *bool                  `json:"crossCloudApply,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _UpdatePricesRequestPrice UpdatePricesRequestPrice

// NewUpdatePricesRequestPrice instantiates a new UpdatePricesRequestPrice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdatePricesRequestPrice() *UpdatePricesRequestPrice {
	this := UpdatePricesRequestPrice{}
	return &this
}

// NewUpdatePricesRequestPriceWithDefaults instantiates a new UpdatePricesRequestPrice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdatePricesRequestPriceWithDefaults() *UpdatePricesRequestPrice {
	this := UpdatePricesRequestPrice{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *UpdatePricesRequestPrice) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePricesRequestPrice) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// IsSetName returns a boolean if a field has been set.
func (o *UpdatePricesRequestPrice) IsSetName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *UpdatePricesRequestPrice) SetName(v string) {
	o.Name = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *UpdatePricesRequestPrice) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePricesRequestPrice) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// IsSetCode returns a boolean if a field has been set.
func (o *UpdatePricesRequestPrice) IsSetCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *UpdatePricesRequestPrice) SetCode(v string) {
	o.Code = &v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *UpdatePricesRequestPrice) GetAccount() AddPricesRequestPriceAccount {
	if o == nil || IsNil(o.Account) {
		var ret AddPricesRequestPriceAccount
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePricesRequestPrice) GetAccountOk() (*AddPricesRequestPriceAccount, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// IsSetAccount returns a boolean if a field has been set.
func (o *UpdatePricesRequestPrice) IsSetAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given AddPricesRequestPriceAccount and assigns it to the Account field.
func (o *UpdatePricesRequestPrice) SetAccount(v AddPricesRequestPriceAccount) {
	o.Account = &v
}

// GetPriceType returns the PriceType field value if set, zero value otherwise.
func (o *UpdatePricesRequestPrice) GetPriceType() string {
	if o == nil || IsNil(o.PriceType) {
		var ret string
		return ret
	}
	return *o.PriceType
}

// GetPriceTypeOk returns a tuple with the PriceType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePricesRequestPrice) GetPriceTypeOk() (*string, bool) {
	if o == nil || IsNil(o.PriceType) {
		return nil, false
	}
	return o.PriceType, true
}

// IsSetPriceType returns a boolean if a field has been set.
func (o *UpdatePricesRequestPrice) IsSetPriceType() bool {
	if o != nil && !IsNil(o.PriceType) {
		return true
	}

	return false
}

// SetPriceType gets a reference to the given string and assigns it to the PriceType field.
func (o *UpdatePricesRequestPrice) SetPriceType(v string) {
	o.PriceType = &v
}

// GetPriceUnit returns the PriceUnit field value if set, zero value otherwise.
func (o *UpdatePricesRequestPrice) GetPriceUnit() string {
	if o == nil || IsNil(o.PriceUnit) {
		var ret string
		return ret
	}
	return *o.PriceUnit
}

// GetPriceUnitOk returns a tuple with the PriceUnit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePricesRequestPrice) GetPriceUnitOk() (*string, bool) {
	if o == nil || IsNil(o.PriceUnit) {
		return nil, false
	}
	return o.PriceUnit, true
}

// IsSetPriceUnit returns a boolean if a field has been set.
func (o *UpdatePricesRequestPrice) IsSetPriceUnit() bool {
	if o != nil && !IsNil(o.PriceUnit) {
		return true
	}

	return false
}

// SetPriceUnit gets a reference to the given string and assigns it to the PriceUnit field.
func (o *UpdatePricesRequestPrice) SetPriceUnit(v string) {
	o.PriceUnit = &v
}

// GetIncurCharges returns the IncurCharges field value if set, zero value otherwise.
func (o *UpdatePricesRequestPrice) GetIncurCharges() string {
	if o == nil || IsNil(o.IncurCharges) {
		var ret string
		return ret
	}
	return *o.IncurCharges
}

// GetIncurChargesOk returns a tuple with the IncurCharges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePricesRequestPrice) GetIncurChargesOk() (*string, bool) {
	if o == nil || IsNil(o.IncurCharges) {
		return nil, false
	}
	return o.IncurCharges, true
}

// IsSetIncurCharges returns a boolean if a field has been set.
func (o *UpdatePricesRequestPrice) IsSetIncurCharges() bool {
	if o != nil && !IsNil(o.IncurCharges) {
		return true
	}

	return false
}

// SetIncurCharges gets a reference to the given string and assigns it to the IncurCharges field.
func (o *UpdatePricesRequestPrice) SetIncurCharges(v string) {
	o.IncurCharges = &v
}

// GetCurrency returns the Currency field value if set, zero value otherwise.
func (o *UpdatePricesRequestPrice) GetCurrency() string {
	if o == nil || IsNil(o.Currency) {
		var ret string
		return ret
	}
	return *o.Currency
}

// GetCurrencyOk returns a tuple with the Currency field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePricesRequestPrice) GetCurrencyOk() (*string, bool) {
	if o == nil || IsNil(o.Currency) {
		return nil, false
	}
	return o.Currency, true
}

// IsSetCurrency returns a boolean if a field has been set.
func (o *UpdatePricesRequestPrice) IsSetCurrency() bool {
	if o != nil && !IsNil(o.Currency) {
		return true
	}

	return false
}

// SetCurrency gets a reference to the given string and assigns it to the Currency field.
func (o *UpdatePricesRequestPrice) SetCurrency(v string) {
	o.Currency = &v
}

// GetCost returns the Cost field value if set, zero value otherwise.
func (o *UpdatePricesRequestPrice) GetCost() float32 {
	if o == nil || IsNil(o.Cost) {
		var ret float32
		return ret
	}
	return *o.Cost
}

// GetCostOk returns a tuple with the Cost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePricesRequestPrice) GetCostOk() (*float32, bool) {
	if o == nil || IsNil(o.Cost) {
		return nil, false
	}
	return o.Cost, true
}

// IsSetCost returns a boolean if a field has been set.
func (o *UpdatePricesRequestPrice) IsSetCost() bool {
	if o != nil && !IsNil(o.Cost) {
		return true
	}

	return false
}

// SetCost gets a reference to the given float32 and assigns it to the Cost field.
func (o *UpdatePricesRequestPrice) SetCost(v float32) {
	o.Cost = &v
}

// GetMarkupType returns the MarkupType field value if set, zero value otherwise.
func (o *UpdatePricesRequestPrice) GetMarkupType() string {
	if o == nil || IsNil(o.MarkupType) {
		var ret string
		return ret
	}
	return *o.MarkupType
}

// GetMarkupTypeOk returns a tuple with the MarkupType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePricesRequestPrice) GetMarkupTypeOk() (*string, bool) {
	if o == nil || IsNil(o.MarkupType) {
		return nil, false
	}
	return o.MarkupType, true
}

// IsSetMarkupType returns a boolean if a field has been set.
func (o *UpdatePricesRequestPrice) IsSetMarkupType() bool {
	if o != nil && !IsNil(o.MarkupType) {
		return true
	}

	return false
}

// SetMarkupType gets a reference to the given string and assigns it to the MarkupType field.
func (o *UpdatePricesRequestPrice) SetMarkupType(v string) {
	o.MarkupType = &v
}

// GetMarkup returns the Markup field value if set, zero value otherwise.
func (o *UpdatePricesRequestPrice) GetMarkup() float32 {
	if o == nil || IsNil(o.Markup) {
		var ret float32
		return ret
	}
	return *o.Markup
}

// GetMarkupOk returns a tuple with the Markup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePricesRequestPrice) GetMarkupOk() (*float32, bool) {
	if o == nil || IsNil(o.Markup) {
		return nil, false
	}
	return o.Markup, true
}

// IsSetMarkup returns a boolean if a field has been set.
func (o *UpdatePricesRequestPrice) IsSetMarkup() bool {
	if o != nil && !IsNil(o.Markup) {
		return true
	}

	return false
}

// SetMarkup gets a reference to the given float32 and assigns it to the Markup field.
func (o *UpdatePricesRequestPrice) SetMarkup(v float32) {
	o.Markup = &v
}

// GetMarkupPercent returns the MarkupPercent field value if set, zero value otherwise.
func (o *UpdatePricesRequestPrice) GetMarkupPercent() float32 {
	if o == nil || IsNil(o.MarkupPercent) {
		var ret float32
		return ret
	}
	return *o.MarkupPercent
}

// GetMarkupPercentOk returns a tuple with the MarkupPercent field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePricesRequestPrice) GetMarkupPercentOk() (*float32, bool) {
	if o == nil || IsNil(o.MarkupPercent) {
		return nil, false
	}
	return o.MarkupPercent, true
}

// IsSetMarkupPercent returns a boolean if a field has been set.
func (o *UpdatePricesRequestPrice) IsSetMarkupPercent() bool {
	if o != nil && !IsNil(o.MarkupPercent) {
		return true
	}

	return false
}

// SetMarkupPercent gets a reference to the given float32 and assigns it to the MarkupPercent field.
func (o *UpdatePricesRequestPrice) SetMarkupPercent(v float32) {
	o.MarkupPercent = &v
}

// GetCustomPrice returns the CustomPrice field value if set, zero value otherwise.
func (o *UpdatePricesRequestPrice) GetCustomPrice() float32 {
	if o == nil || IsNil(o.CustomPrice) {
		var ret float32
		return ret
	}
	return *o.CustomPrice
}

// GetCustomPriceOk returns a tuple with the CustomPrice field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePricesRequestPrice) GetCustomPriceOk() (*float32, bool) {
	if o == nil || IsNil(o.CustomPrice) {
		return nil, false
	}
	return o.CustomPrice, true
}

// IsSetCustomPrice returns a boolean if a field has been set.
func (o *UpdatePricesRequestPrice) IsSetCustomPrice() bool {
	if o != nil && !IsNil(o.CustomPrice) {
		return true
	}

	return false
}

// SetCustomPrice gets a reference to the given float32 and assigns it to the CustomPrice field.
func (o *UpdatePricesRequestPrice) SetCustomPrice(v float32) {
	o.CustomPrice = &v
}

// GetPlatform returns the Platform field value if set, zero value otherwise.
func (o *UpdatePricesRequestPrice) GetPlatform() string {
	if o == nil || IsNil(o.Platform) {
		var ret string
		return ret
	}
	return *o.Platform
}

// GetPlatformOk returns a tuple with the Platform field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePricesRequestPrice) GetPlatformOk() (*string, bool) {
	if o == nil || IsNil(o.Platform) {
		return nil, false
	}
	return o.Platform, true
}

// IsSetPlatform returns a boolean if a field has been set.
func (o *UpdatePricesRequestPrice) IsSetPlatform() bool {
	if o != nil && !IsNil(o.Platform) {
		return true
	}

	return false
}

// SetPlatform gets a reference to the given string and assigns it to the Platform field.
func (o *UpdatePricesRequestPrice) SetPlatform(v string) {
	o.Platform = &v
}

// GetSoftware returns the Software field value if set, zero value otherwise.
func (o *UpdatePricesRequestPrice) GetSoftware() string {
	if o == nil || IsNil(o.Software) {
		var ret string
		return ret
	}
	return *o.Software
}

// GetSoftwareOk returns a tuple with the Software field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePricesRequestPrice) GetSoftwareOk() (*string, bool) {
	if o == nil || IsNil(o.Software) {
		return nil, false
	}
	return o.Software, true
}

// IsSetSoftware returns a boolean if a field has been set.
func (o *UpdatePricesRequestPrice) IsSetSoftware() bool {
	if o != nil && !IsNil(o.Software) {
		return true
	}

	return false
}

// SetSoftware gets a reference to the given string and assigns it to the Software field.
func (o *UpdatePricesRequestPrice) SetSoftware(v string) {
	o.Software = &v
}

// GetVolumeType returns the VolumeType field value if set, zero value otherwise.
func (o *UpdatePricesRequestPrice) GetVolumeType() AddPricesRequestPriceVolumeType {
	if o == nil || IsNil(o.VolumeType) {
		var ret AddPricesRequestPriceVolumeType
		return ret
	}
	return *o.VolumeType
}

// GetVolumeTypeOk returns a tuple with the VolumeType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePricesRequestPrice) GetVolumeTypeOk() (*AddPricesRequestPriceVolumeType, bool) {
	if o == nil || IsNil(o.VolumeType) {
		return nil, false
	}
	return o.VolumeType, true
}

// IsSetVolumeType returns a boolean if a field has been set.
func (o *UpdatePricesRequestPrice) IsSetVolumeType() bool {
	if o != nil && !IsNil(o.VolumeType) {
		return true
	}

	return false
}

// SetVolumeType gets a reference to the given AddPricesRequestPriceVolumeType and assigns it to the VolumeType field.
func (o *UpdatePricesRequestPrice) SetVolumeType(v AddPricesRequestPriceVolumeType) {
	o.VolumeType = &v
}

// GetDatastore returns the Datastore field value if set, zero value otherwise.
func (o *UpdatePricesRequestPrice) GetDatastore() AddPricesRequestPriceDatastore {
	if o == nil || IsNil(o.Datastore) {
		var ret AddPricesRequestPriceDatastore
		return ret
	}
	return *o.Datastore
}

// GetDatastoreOk returns a tuple with the Datastore field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePricesRequestPrice) GetDatastoreOk() (*AddPricesRequestPriceDatastore, bool) {
	if o == nil || IsNil(o.Datastore) {
		return nil, false
	}
	return o.Datastore, true
}

// IsSetDatastore returns a boolean if a field has been set.
func (o *UpdatePricesRequestPrice) IsSetDatastore() bool {
	if o != nil && !IsNil(o.Datastore) {
		return true
	}

	return false
}

// SetDatastore gets a reference to the given AddPricesRequestPriceDatastore and assigns it to the Datastore field.
func (o *UpdatePricesRequestPrice) SetDatastore(v AddPricesRequestPriceDatastore) {
	o.Datastore = &v
}

// GetCrossCloudApply returns the CrossCloudApply field value if set, zero value otherwise.
func (o *UpdatePricesRequestPrice) GetCrossCloudApply() bool {
	if o == nil || IsNil(o.CrossCloudApply) {
		var ret bool
		return ret
	}
	return *o.CrossCloudApply
}

// GetCrossCloudApplyOk returns a tuple with the CrossCloudApply field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdatePricesRequestPrice) GetCrossCloudApplyOk() (*bool, bool) {
	if o == nil || IsNil(o.CrossCloudApply) {
		return nil, false
	}
	return o.CrossCloudApply, true
}

// IsSetCrossCloudApply returns a boolean if a field has been set.
func (o *UpdatePricesRequestPrice) IsSetCrossCloudApply() bool {
	if o != nil && !IsNil(o.CrossCloudApply) {
		return true
	}

	return false
}

// SetCrossCloudApply gets a reference to the given bool and assigns it to the CrossCloudApply field.
func (o *UpdatePricesRequestPrice) SetCrossCloudApply(v bool) {
	o.CrossCloudApply = &v
}

func (o UpdatePricesRequestPrice) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdatePricesRequestPrice) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.PriceType) {
		toSerialize["priceType"] = o.PriceType
	}
	if !IsNil(o.PriceUnit) {
		toSerialize["priceUnit"] = o.PriceUnit
	}
	if !IsNil(o.IncurCharges) {
		toSerialize["incurCharges"] = o.IncurCharges
	}
	if !IsNil(o.Currency) {
		toSerialize["currency"] = o.Currency
	}
	if !IsNil(o.Cost) {
		toSerialize["cost"] = o.Cost
	}
	if !IsNil(o.MarkupType) {
		toSerialize["markupType"] = o.MarkupType
	}
	if !IsNil(o.Markup) {
		toSerialize["markup"] = o.Markup
	}
	if !IsNil(o.MarkupPercent) {
		toSerialize["markupPercent"] = o.MarkupPercent
	}
	if !IsNil(o.CustomPrice) {
		toSerialize["customPrice"] = o.CustomPrice
	}
	if !IsNil(o.Platform) {
		toSerialize["platform"] = o.Platform
	}
	if !IsNil(o.Software) {
		toSerialize["software"] = o.Software
	}
	if !IsNil(o.VolumeType) {
		toSerialize["volumeType"] = o.VolumeType
	}
	if !IsNil(o.Datastore) {
		toSerialize["datastore"] = o.Datastore
	}
	if !IsNil(o.CrossCloudApply) {
		toSerialize["crossCloudApply"] = o.CrossCloudApply
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *UpdatePricesRequestPrice) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache
