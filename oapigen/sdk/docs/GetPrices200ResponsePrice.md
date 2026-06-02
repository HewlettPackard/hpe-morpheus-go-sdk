# GetPrices200ResponsePrice

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
**Markup** | Pointer to **NullableFloat32** |  | [optional] 
**MarkupPercent** | Pointer to **NullableFloat32** |  | [optional] 
**Cost** | Pointer to **NullableFloat32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**IncurCharges** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **NullableString** |  | [optional] 
**Software** | Pointer to **NullableString** |  | [optional] 
**VolumeType** | Pointer to [**GetPrices200ResponsePriceVolumeType**](GetPrices200ResponsePriceVolumeType.md) |  | [optional] 
**Datastore** | Pointer to [**GetPrices200ResponsePriceDatastore**](GetPrices200ResponsePriceDatastore.md) |  | [optional] 
**CrossCloudApply** | Pointer to **NullableBool** |  | [optional] 
**Account** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGetPrices200ResponsePrice

`func NewGetPrices200ResponsePrice() *GetPrices200ResponsePrice`

NewGetPrices200ResponsePrice instantiates a new GetPrices200ResponsePrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetPrices200ResponsePrice) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetPrices200ResponsePrice) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetPrices200ResponsePrice) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetPrices200ResponsePrice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetPrices200ResponsePrice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetPrices200ResponsePrice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetPrices200ResponsePrice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetPrices200ResponsePrice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *GetPrices200ResponsePrice) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetPrices200ResponsePrice) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetPrices200ResponsePrice) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetPrices200ResponsePrice) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetActive

`func (o *GetPrices200ResponsePrice) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetPrices200ResponsePrice) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetPrices200ResponsePrice) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetPrices200ResponsePrice) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPriceType

`func (o *GetPrices200ResponsePrice) GetPriceType() string`

GetPriceType returns the PriceType field if non-nil, zero value otherwise.

### GetPriceTypeOk

`func (o *GetPrices200ResponsePrice) GetPriceTypeOk() (*string, bool)`

GetPriceTypeOk returns a tuple with the PriceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceType

`func (o *GetPrices200ResponsePrice) SetPriceType(v string)`

SetPriceType sets PriceType field to given value.

### HasPriceType

`func (o *GetPrices200ResponsePrice) HasPriceType() bool`

HasPriceType returns a boolean if a field has been set.

### GetPriceUnit

`func (o *GetPrices200ResponsePrice) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *GetPrices200ResponsePrice) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *GetPrices200ResponsePrice) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *GetPrices200ResponsePrice) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### GetAdditionalPriceUnit

`func (o *GetPrices200ResponsePrice) GetAdditionalPriceUnit() string`

GetAdditionalPriceUnit returns the AdditionalPriceUnit field if non-nil, zero value otherwise.

### GetAdditionalPriceUnitOk

`func (o *GetPrices200ResponsePrice) GetAdditionalPriceUnitOk() (*string, bool)`

GetAdditionalPriceUnitOk returns a tuple with the AdditionalPriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalPriceUnit

`func (o *GetPrices200ResponsePrice) SetAdditionalPriceUnit(v string)`

SetAdditionalPriceUnit sets AdditionalPriceUnit field to given value.

### HasAdditionalPriceUnit

`func (o *GetPrices200ResponsePrice) HasAdditionalPriceUnit() bool`

HasAdditionalPriceUnit returns a boolean if a field has been set.

### SetAdditionalPriceUnitNil

`func (o *GetPrices200ResponsePrice) SetAdditionalPriceUnitNil(b bool)`

 SetAdditionalPriceUnitNil sets the value for AdditionalPriceUnit to be an explicit nil

### UnsetAdditionalPriceUnit
`func (o *GetPrices200ResponsePrice) UnsetAdditionalPriceUnit()`

UnsetAdditionalPriceUnit ensures that no value is present for AdditionalPriceUnit, not even an explicit nil
### GetPrice

`func (o *GetPrices200ResponsePrice) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *GetPrices200ResponsePrice) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *GetPrices200ResponsePrice) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *GetPrices200ResponsePrice) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *GetPrices200ResponsePrice) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *GetPrices200ResponsePrice) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetCustomPrice

`func (o *GetPrices200ResponsePrice) GetCustomPrice() float32`

GetCustomPrice returns the CustomPrice field if non-nil, zero value otherwise.

### GetCustomPriceOk

`func (o *GetPrices200ResponsePrice) GetCustomPriceOk() (*float32, bool)`

GetCustomPriceOk returns a tuple with the CustomPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPrice

`func (o *GetPrices200ResponsePrice) SetCustomPrice(v float32)`

SetCustomPrice sets CustomPrice field to given value.

### HasCustomPrice

`func (o *GetPrices200ResponsePrice) HasCustomPrice() bool`

HasCustomPrice returns a boolean if a field has been set.

### SetCustomPriceNil

`func (o *GetPrices200ResponsePrice) SetCustomPriceNil(b bool)`

 SetCustomPriceNil sets the value for CustomPrice to be an explicit nil

### UnsetCustomPrice
`func (o *GetPrices200ResponsePrice) UnsetCustomPrice()`

UnsetCustomPrice ensures that no value is present for CustomPrice, not even an explicit nil
### GetMarkupType

`func (o *GetPrices200ResponsePrice) GetMarkupType() string`

GetMarkupType returns the MarkupType field if non-nil, zero value otherwise.

### GetMarkupTypeOk

`func (o *GetPrices200ResponsePrice) GetMarkupTypeOk() (*string, bool)`

GetMarkupTypeOk returns a tuple with the MarkupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkupType

`func (o *GetPrices200ResponsePrice) SetMarkupType(v string)`

SetMarkupType sets MarkupType field to given value.

### HasMarkupType

`func (o *GetPrices200ResponsePrice) HasMarkupType() bool`

HasMarkupType returns a boolean if a field has been set.

### SetMarkupTypeNil

`func (o *GetPrices200ResponsePrice) SetMarkupTypeNil(b bool)`

 SetMarkupTypeNil sets the value for MarkupType to be an explicit nil

### UnsetMarkupType
`func (o *GetPrices200ResponsePrice) UnsetMarkupType()`

UnsetMarkupType ensures that no value is present for MarkupType, not even an explicit nil
### GetMarkup

`func (o *GetPrices200ResponsePrice) GetMarkup() float32`

GetMarkup returns the Markup field if non-nil, zero value otherwise.

### GetMarkupOk

`func (o *GetPrices200ResponsePrice) GetMarkupOk() (*float32, bool)`

GetMarkupOk returns a tuple with the Markup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkup

`func (o *GetPrices200ResponsePrice) SetMarkup(v float32)`

SetMarkup sets Markup field to given value.

### HasMarkup

`func (o *GetPrices200ResponsePrice) HasMarkup() bool`

HasMarkup returns a boolean if a field has been set.

### SetMarkupNil

`func (o *GetPrices200ResponsePrice) SetMarkupNil(b bool)`

 SetMarkupNil sets the value for Markup to be an explicit nil

### UnsetMarkup
`func (o *GetPrices200ResponsePrice) UnsetMarkup()`

UnsetMarkup ensures that no value is present for Markup, not even an explicit nil
### GetMarkupPercent

`func (o *GetPrices200ResponsePrice) GetMarkupPercent() float32`

GetMarkupPercent returns the MarkupPercent field if non-nil, zero value otherwise.

### GetMarkupPercentOk

`func (o *GetPrices200ResponsePrice) GetMarkupPercentOk() (*float32, bool)`

GetMarkupPercentOk returns a tuple with the MarkupPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkupPercent

`func (o *GetPrices200ResponsePrice) SetMarkupPercent(v float32)`

SetMarkupPercent sets MarkupPercent field to given value.

### HasMarkupPercent

`func (o *GetPrices200ResponsePrice) HasMarkupPercent() bool`

HasMarkupPercent returns a boolean if a field has been set.

### SetMarkupPercentNil

`func (o *GetPrices200ResponsePrice) SetMarkupPercentNil(b bool)`

 SetMarkupPercentNil sets the value for MarkupPercent to be an explicit nil

### UnsetMarkupPercent
`func (o *GetPrices200ResponsePrice) UnsetMarkupPercent()`

UnsetMarkupPercent ensures that no value is present for MarkupPercent, not even an explicit nil
### GetCost

`func (o *GetPrices200ResponsePrice) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *GetPrices200ResponsePrice) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *GetPrices200ResponsePrice) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *GetPrices200ResponsePrice) HasCost() bool`

HasCost returns a boolean if a field has been set.

### SetCostNil

`func (o *GetPrices200ResponsePrice) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *GetPrices200ResponsePrice) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetCurrency

`func (o *GetPrices200ResponsePrice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *GetPrices200ResponsePrice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *GetPrices200ResponsePrice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *GetPrices200ResponsePrice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetIncurCharges

`func (o *GetPrices200ResponsePrice) GetIncurCharges() string`

GetIncurCharges returns the IncurCharges field if non-nil, zero value otherwise.

### GetIncurChargesOk

`func (o *GetPrices200ResponsePrice) GetIncurChargesOk() (*string, bool)`

GetIncurChargesOk returns a tuple with the IncurCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncurCharges

`func (o *GetPrices200ResponsePrice) SetIncurCharges(v string)`

SetIncurCharges sets IncurCharges field to given value.

### HasIncurCharges

`func (o *GetPrices200ResponsePrice) HasIncurCharges() bool`

HasIncurCharges returns a boolean if a field has been set.

### GetPlatform

`func (o *GetPrices200ResponsePrice) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GetPrices200ResponsePrice) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GetPrices200ResponsePrice) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *GetPrices200ResponsePrice) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *GetPrices200ResponsePrice) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *GetPrices200ResponsePrice) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetSoftware

`func (o *GetPrices200ResponsePrice) GetSoftware() string`

GetSoftware returns the Software field if non-nil, zero value otherwise.

### GetSoftwareOk

`func (o *GetPrices200ResponsePrice) GetSoftwareOk() (*string, bool)`

GetSoftwareOk returns a tuple with the Software field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftware

`func (o *GetPrices200ResponsePrice) SetSoftware(v string)`

SetSoftware sets Software field to given value.

### HasSoftware

`func (o *GetPrices200ResponsePrice) HasSoftware() bool`

HasSoftware returns a boolean if a field has been set.

### SetSoftwareNil

`func (o *GetPrices200ResponsePrice) SetSoftwareNil(b bool)`

 SetSoftwareNil sets the value for Software to be an explicit nil

### UnsetSoftware
`func (o *GetPrices200ResponsePrice) UnsetSoftware()`

UnsetSoftware ensures that no value is present for Software, not even an explicit nil
### GetVolumeType

`func (o *GetPrices200ResponsePrice) GetVolumeType() GetPrices200ResponsePriceVolumeType`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *GetPrices200ResponsePrice) GetVolumeTypeOk() (*GetPrices200ResponsePriceVolumeType, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *GetPrices200ResponsePrice) SetVolumeType(v GetPrices200ResponsePriceVolumeType)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *GetPrices200ResponsePrice) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetDatastore

`func (o *GetPrices200ResponsePrice) GetDatastore() GetPrices200ResponsePriceDatastore`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *GetPrices200ResponsePrice) GetDatastoreOk() (*GetPrices200ResponsePriceDatastore, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *GetPrices200ResponsePrice) SetDatastore(v GetPrices200ResponsePriceDatastore)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *GetPrices200ResponsePrice) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetCrossCloudApply

`func (o *GetPrices200ResponsePrice) GetCrossCloudApply() bool`

GetCrossCloudApply returns the CrossCloudApply field if non-nil, zero value otherwise.

### GetCrossCloudApplyOk

`func (o *GetPrices200ResponsePrice) GetCrossCloudApplyOk() (*bool, bool)`

GetCrossCloudApplyOk returns a tuple with the CrossCloudApply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossCloudApply

`func (o *GetPrices200ResponsePrice) SetCrossCloudApply(v bool)`

SetCrossCloudApply sets CrossCloudApply field to given value.

### HasCrossCloudApply

`func (o *GetPrices200ResponsePrice) HasCrossCloudApply() bool`

HasCrossCloudApply returns a boolean if a field has been set.

### SetCrossCloudApplyNil

`func (o *GetPrices200ResponsePrice) SetCrossCloudApplyNil(b bool)`

 SetCrossCloudApplyNil sets the value for CrossCloudApply to be an explicit nil

### UnsetCrossCloudApply
`func (o *GetPrices200ResponsePrice) UnsetCrossCloudApply()`

UnsetCrossCloudApply ensures that no value is present for CrossCloudApply, not even an explicit nil
### GetAccount

`func (o *GetPrices200ResponsePrice) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetPrices200ResponsePrice) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetPrices200ResponsePrice) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetPrices200ResponsePrice) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *GetPrices200ResponsePrice) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *GetPrices200ResponsePrice) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


