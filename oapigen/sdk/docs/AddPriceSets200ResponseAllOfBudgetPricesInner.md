# AddPriceSets200ResponseAllOfBudgetPricesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**PriceType** | Pointer to **string** |  | [optional] 
**PriceUnit** | Pointer to **string** |  | [optional] 
**AdditionalPriceUnit** | Pointer to **NullableString** |  | [optional] 
**Price** | Pointer to **NullableFloat32** |  | [optional] 
**CustomPrice** | Pointer to **NullableFloat32** |  | [optional] 
**MarkupType** | Pointer to **NullableString** |  | [optional] 
**Markup** | Pointer to **int64** |  | [optional] 
**MarkupPercent** | Pointer to **NullableFloat32** |  | [optional] 
**Cost** | Pointer to **NullableFloat32** |  | [optional] 
**Currency** | Pointer to **NullableString** |  | [optional] 
**IncurCharges** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **NullableString** |  | [optional] 
**Software** | Pointer to **NullableString** |  | [optional] 
**VolumeType** | Pointer to [**AddPriceSets200ResponseAllOfBudgetPricesInnerVolumeType**](AddPriceSets200ResponseAllOfBudgetPricesInnerVolumeType.md) |  | [optional] 
**Datastore** | Pointer to **NullableString** |  | [optional] 
**CrossCloudApply** | Pointer to **NullableBool** |  | [optional] 
**Account** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAddPriceSets200ResponseAllOfBudgetPricesInner

`func NewAddPriceSets200ResponseAllOfBudgetPricesInner() *AddPriceSets200ResponseAllOfBudgetPricesInner`

NewAddPriceSets200ResponseAllOfBudgetPricesInner instantiates a new AddPriceSets200ResponseAllOfBudgetPricesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetActive

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPriceType

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetPriceType() string`

GetPriceType returns the PriceType field if non-nil, zero value otherwise.

### GetPriceTypeOk

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetPriceTypeOk() (*string, bool)`

GetPriceTypeOk returns a tuple with the PriceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceType

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) SetPriceType(v string)`

SetPriceType sets PriceType field to given value.

### HasPriceType

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) HasPriceType() bool`

HasPriceType returns a boolean if a field has been set.

### GetPriceUnit

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### GetAdditionalPriceUnit

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetAdditionalPriceUnit() string`

GetAdditionalPriceUnit returns the AdditionalPriceUnit field if non-nil, zero value otherwise.

### GetAdditionalPriceUnitOk

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetAdditionalPriceUnitOk() (*string, bool)`

GetAdditionalPriceUnitOk returns a tuple with the AdditionalPriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalPriceUnit

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) SetAdditionalPriceUnit(v string)`

SetAdditionalPriceUnit sets AdditionalPriceUnit field to given value.

### HasAdditionalPriceUnit

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) HasAdditionalPriceUnit() bool`

HasAdditionalPriceUnit returns a boolean if a field has been set.

### SetAdditionalPriceUnitNil

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) SetAdditionalPriceUnitNil(b bool)`

 SetAdditionalPriceUnitNil sets the value for AdditionalPriceUnit to be an explicit nil

### UnsetAdditionalPriceUnit
`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) UnsetAdditionalPriceUnit()`

UnsetAdditionalPriceUnit ensures that no value is present for AdditionalPriceUnit, not even an explicit nil
### GetPrice

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetCustomPrice

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetCustomPrice() float32`

GetCustomPrice returns the CustomPrice field if non-nil, zero value otherwise.

### GetCustomPriceOk

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetCustomPriceOk() (*float32, bool)`

GetCustomPriceOk returns a tuple with the CustomPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPrice

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) SetCustomPrice(v float32)`

SetCustomPrice sets CustomPrice field to given value.

### HasCustomPrice

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) HasCustomPrice() bool`

HasCustomPrice returns a boolean if a field has been set.

### SetCustomPriceNil

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) SetCustomPriceNil(b bool)`

 SetCustomPriceNil sets the value for CustomPrice to be an explicit nil

### UnsetCustomPrice
`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) UnsetCustomPrice()`

UnsetCustomPrice ensures that no value is present for CustomPrice, not even an explicit nil
### GetMarkupType

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetMarkupType() string`

GetMarkupType returns the MarkupType field if non-nil, zero value otherwise.

### GetMarkupTypeOk

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetMarkupTypeOk() (*string, bool)`

GetMarkupTypeOk returns a tuple with the MarkupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkupType

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) SetMarkupType(v string)`

SetMarkupType sets MarkupType field to given value.

### HasMarkupType

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) HasMarkupType() bool`

HasMarkupType returns a boolean if a field has been set.

### SetMarkupTypeNil

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) SetMarkupTypeNil(b bool)`

 SetMarkupTypeNil sets the value for MarkupType to be an explicit nil

### UnsetMarkupType
`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) UnsetMarkupType()`

UnsetMarkupType ensures that no value is present for MarkupType, not even an explicit nil
### GetMarkup

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetMarkup() int64`

GetMarkup returns the Markup field if non-nil, zero value otherwise.

### GetMarkupOk

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetMarkupOk() (*int64, bool)`

GetMarkupOk returns a tuple with the Markup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkup

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) SetMarkup(v int64)`

SetMarkup sets Markup field to given value.

### HasMarkup

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) HasMarkup() bool`

HasMarkup returns a boolean if a field has been set.

### GetMarkupPercent

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetMarkupPercent() float32`

GetMarkupPercent returns the MarkupPercent field if non-nil, zero value otherwise.

### GetMarkupPercentOk

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetMarkupPercentOk() (*float32, bool)`

GetMarkupPercentOk returns a tuple with the MarkupPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkupPercent

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) SetMarkupPercent(v float32)`

SetMarkupPercent sets MarkupPercent field to given value.

### HasMarkupPercent

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) HasMarkupPercent() bool`

HasMarkupPercent returns a boolean if a field has been set.

### SetMarkupPercentNil

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) SetMarkupPercentNil(b bool)`

 SetMarkupPercentNil sets the value for MarkupPercent to be an explicit nil

### UnsetMarkupPercent
`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) UnsetMarkupPercent()`

UnsetMarkupPercent ensures that no value is present for MarkupPercent, not even an explicit nil
### GetCost

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) HasCost() bool`

HasCost returns a boolean if a field has been set.

### SetCostNil

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetCurrency

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### SetCurrencyNil

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) SetCurrencyNil(b bool)`

 SetCurrencyNil sets the value for Currency to be an explicit nil

### UnsetCurrency
`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) UnsetCurrency()`

UnsetCurrency ensures that no value is present for Currency, not even an explicit nil
### GetIncurCharges

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetIncurCharges() string`

GetIncurCharges returns the IncurCharges field if non-nil, zero value otherwise.

### GetIncurChargesOk

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetIncurChargesOk() (*string, bool)`

GetIncurChargesOk returns a tuple with the IncurCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncurCharges

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) SetIncurCharges(v string)`

SetIncurCharges sets IncurCharges field to given value.

### HasIncurCharges

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) HasIncurCharges() bool`

HasIncurCharges returns a boolean if a field has been set.

### GetPlatform

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetSoftware

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetSoftware() string`

GetSoftware returns the Software field if non-nil, zero value otherwise.

### GetSoftwareOk

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetSoftwareOk() (*string, bool)`

GetSoftwareOk returns a tuple with the Software field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftware

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) SetSoftware(v string)`

SetSoftware sets Software field to given value.

### HasSoftware

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) HasSoftware() bool`

HasSoftware returns a boolean if a field has been set.

### SetSoftwareNil

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) SetSoftwareNil(b bool)`

 SetSoftwareNil sets the value for Software to be an explicit nil

### UnsetSoftware
`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) UnsetSoftware()`

UnsetSoftware ensures that no value is present for Software, not even an explicit nil
### GetVolumeType

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetVolumeType() AddPriceSets200ResponseAllOfBudgetPricesInnerVolumeType`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetVolumeTypeOk() (*AddPriceSets200ResponseAllOfBudgetPricesInnerVolumeType, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) SetVolumeType(v AddPriceSets200ResponseAllOfBudgetPricesInnerVolumeType)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetDatastore

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetDatastore() string`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetDatastoreOk() (*string, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) SetDatastore(v string)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### SetDatastoreNil

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) SetDatastoreNil(b bool)`

 SetDatastoreNil sets the value for Datastore to be an explicit nil

### UnsetDatastore
`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) UnsetDatastore()`

UnsetDatastore ensures that no value is present for Datastore, not even an explicit nil
### GetCrossCloudApply

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetCrossCloudApply() bool`

GetCrossCloudApply returns the CrossCloudApply field if non-nil, zero value otherwise.

### GetCrossCloudApplyOk

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetCrossCloudApplyOk() (*bool, bool)`

GetCrossCloudApplyOk returns a tuple with the CrossCloudApply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossCloudApply

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) SetCrossCloudApply(v bool)`

SetCrossCloudApply sets CrossCloudApply field to given value.

### HasCrossCloudApply

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) HasCrossCloudApply() bool`

HasCrossCloudApply returns a boolean if a field has been set.

### SetCrossCloudApplyNil

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) SetCrossCloudApplyNil(b bool)`

 SetCrossCloudApplyNil sets the value for CrossCloudApply to be an explicit nil

### UnsetCrossCloudApply
`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) UnsetCrossCloudApply()`

UnsetCrossCloudApply ensures that no value is present for CrossCloudApply, not even an explicit nil
### GetAccount

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *AddPriceSets200ResponseAllOfBudgetPricesInner) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


