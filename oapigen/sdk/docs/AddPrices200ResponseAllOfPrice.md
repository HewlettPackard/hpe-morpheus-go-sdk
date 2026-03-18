# AddPrices200ResponseAllOfPrice

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
**VolumeType** | Pointer to [**AddPrices200ResponseAllOfPriceVolumeType**](AddPrices200ResponseAllOfPriceVolumeType.md) |  | [optional] 
**Datastore** | Pointer to [**AddPrices200ResponseAllOfPriceDatastore**](AddPrices200ResponseAllOfPriceDatastore.md) |  | [optional] 
**CrossCloudApply** | Pointer to **NullableBool** |  | [optional] 
**Account** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAddPrices200ResponseAllOfPrice

`func NewAddPrices200ResponseAllOfPrice() *AddPrices200ResponseAllOfPrice`

NewAddPrices200ResponseAllOfPrice instantiates a new AddPrices200ResponseAllOfPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPrices200ResponseAllOfPriceWithDefaults

`func NewAddPrices200ResponseAllOfPriceWithDefaults() *AddPrices200ResponseAllOfPrice`

NewAddPrices200ResponseAllOfPriceWithDefaults instantiates a new AddPrices200ResponseAllOfPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddPrices200ResponseAllOfPrice) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddPrices200ResponseAllOfPrice) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddPrices200ResponseAllOfPrice) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddPrices200ResponseAllOfPrice) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddPrices200ResponseAllOfPrice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddPrices200ResponseAllOfPrice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddPrices200ResponseAllOfPrice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddPrices200ResponseAllOfPrice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *AddPrices200ResponseAllOfPrice) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AddPrices200ResponseAllOfPrice) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AddPrices200ResponseAllOfPrice) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AddPrices200ResponseAllOfPrice) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetActive

`func (o *AddPrices200ResponseAllOfPrice) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AddPrices200ResponseAllOfPrice) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AddPrices200ResponseAllOfPrice) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AddPrices200ResponseAllOfPrice) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPriceType

`func (o *AddPrices200ResponseAllOfPrice) GetPriceType() string`

GetPriceType returns the PriceType field if non-nil, zero value otherwise.

### GetPriceTypeOk

`func (o *AddPrices200ResponseAllOfPrice) GetPriceTypeOk() (*string, bool)`

GetPriceTypeOk returns a tuple with the PriceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceType

`func (o *AddPrices200ResponseAllOfPrice) SetPriceType(v string)`

SetPriceType sets PriceType field to given value.

### HasPriceType

`func (o *AddPrices200ResponseAllOfPrice) HasPriceType() bool`

HasPriceType returns a boolean if a field has been set.

### GetPriceUnit

`func (o *AddPrices200ResponseAllOfPrice) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *AddPrices200ResponseAllOfPrice) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *AddPrices200ResponseAllOfPrice) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *AddPrices200ResponseAllOfPrice) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### GetAdditionalPriceUnit

`func (o *AddPrices200ResponseAllOfPrice) GetAdditionalPriceUnit() string`

GetAdditionalPriceUnit returns the AdditionalPriceUnit field if non-nil, zero value otherwise.

### GetAdditionalPriceUnitOk

`func (o *AddPrices200ResponseAllOfPrice) GetAdditionalPriceUnitOk() (*string, bool)`

GetAdditionalPriceUnitOk returns a tuple with the AdditionalPriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalPriceUnit

`func (o *AddPrices200ResponseAllOfPrice) SetAdditionalPriceUnit(v string)`

SetAdditionalPriceUnit sets AdditionalPriceUnit field to given value.

### HasAdditionalPriceUnit

`func (o *AddPrices200ResponseAllOfPrice) HasAdditionalPriceUnit() bool`

HasAdditionalPriceUnit returns a boolean if a field has been set.

### SetAdditionalPriceUnitNil

`func (o *AddPrices200ResponseAllOfPrice) SetAdditionalPriceUnitNil(b bool)`

 SetAdditionalPriceUnitNil sets the value for AdditionalPriceUnit to be an explicit nil

### UnsetAdditionalPriceUnit
`func (o *AddPrices200ResponseAllOfPrice) UnsetAdditionalPriceUnit()`

UnsetAdditionalPriceUnit ensures that no value is present for AdditionalPriceUnit, not even an explicit nil
### GetPrice

`func (o *AddPrices200ResponseAllOfPrice) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *AddPrices200ResponseAllOfPrice) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *AddPrices200ResponseAllOfPrice) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *AddPrices200ResponseAllOfPrice) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *AddPrices200ResponseAllOfPrice) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *AddPrices200ResponseAllOfPrice) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetCustomPrice

`func (o *AddPrices200ResponseAllOfPrice) GetCustomPrice() float32`

GetCustomPrice returns the CustomPrice field if non-nil, zero value otherwise.

### GetCustomPriceOk

`func (o *AddPrices200ResponseAllOfPrice) GetCustomPriceOk() (*float32, bool)`

GetCustomPriceOk returns a tuple with the CustomPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPrice

`func (o *AddPrices200ResponseAllOfPrice) SetCustomPrice(v float32)`

SetCustomPrice sets CustomPrice field to given value.

### HasCustomPrice

`func (o *AddPrices200ResponseAllOfPrice) HasCustomPrice() bool`

HasCustomPrice returns a boolean if a field has been set.

### SetCustomPriceNil

`func (o *AddPrices200ResponseAllOfPrice) SetCustomPriceNil(b bool)`

 SetCustomPriceNil sets the value for CustomPrice to be an explicit nil

### UnsetCustomPrice
`func (o *AddPrices200ResponseAllOfPrice) UnsetCustomPrice()`

UnsetCustomPrice ensures that no value is present for CustomPrice, not even an explicit nil
### GetMarkupType

`func (o *AddPrices200ResponseAllOfPrice) GetMarkupType() string`

GetMarkupType returns the MarkupType field if non-nil, zero value otherwise.

### GetMarkupTypeOk

`func (o *AddPrices200ResponseAllOfPrice) GetMarkupTypeOk() (*string, bool)`

GetMarkupTypeOk returns a tuple with the MarkupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkupType

`func (o *AddPrices200ResponseAllOfPrice) SetMarkupType(v string)`

SetMarkupType sets MarkupType field to given value.

### HasMarkupType

`func (o *AddPrices200ResponseAllOfPrice) HasMarkupType() bool`

HasMarkupType returns a boolean if a field has been set.

### SetMarkupTypeNil

`func (o *AddPrices200ResponseAllOfPrice) SetMarkupTypeNil(b bool)`

 SetMarkupTypeNil sets the value for MarkupType to be an explicit nil

### UnsetMarkupType
`func (o *AddPrices200ResponseAllOfPrice) UnsetMarkupType()`

UnsetMarkupType ensures that no value is present for MarkupType, not even an explicit nil
### GetMarkup

`func (o *AddPrices200ResponseAllOfPrice) GetMarkup() float32`

GetMarkup returns the Markup field if non-nil, zero value otherwise.

### GetMarkupOk

`func (o *AddPrices200ResponseAllOfPrice) GetMarkupOk() (*float32, bool)`

GetMarkupOk returns a tuple with the Markup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkup

`func (o *AddPrices200ResponseAllOfPrice) SetMarkup(v float32)`

SetMarkup sets Markup field to given value.

### HasMarkup

`func (o *AddPrices200ResponseAllOfPrice) HasMarkup() bool`

HasMarkup returns a boolean if a field has been set.

### SetMarkupNil

`func (o *AddPrices200ResponseAllOfPrice) SetMarkupNil(b bool)`

 SetMarkupNil sets the value for Markup to be an explicit nil

### UnsetMarkup
`func (o *AddPrices200ResponseAllOfPrice) UnsetMarkup()`

UnsetMarkup ensures that no value is present for Markup, not even an explicit nil
### GetMarkupPercent

`func (o *AddPrices200ResponseAllOfPrice) GetMarkupPercent() float32`

GetMarkupPercent returns the MarkupPercent field if non-nil, zero value otherwise.

### GetMarkupPercentOk

`func (o *AddPrices200ResponseAllOfPrice) GetMarkupPercentOk() (*float32, bool)`

GetMarkupPercentOk returns a tuple with the MarkupPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkupPercent

`func (o *AddPrices200ResponseAllOfPrice) SetMarkupPercent(v float32)`

SetMarkupPercent sets MarkupPercent field to given value.

### HasMarkupPercent

`func (o *AddPrices200ResponseAllOfPrice) HasMarkupPercent() bool`

HasMarkupPercent returns a boolean if a field has been set.

### SetMarkupPercentNil

`func (o *AddPrices200ResponseAllOfPrice) SetMarkupPercentNil(b bool)`

 SetMarkupPercentNil sets the value for MarkupPercent to be an explicit nil

### UnsetMarkupPercent
`func (o *AddPrices200ResponseAllOfPrice) UnsetMarkupPercent()`

UnsetMarkupPercent ensures that no value is present for MarkupPercent, not even an explicit nil
### GetCost

`func (o *AddPrices200ResponseAllOfPrice) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *AddPrices200ResponseAllOfPrice) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *AddPrices200ResponseAllOfPrice) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *AddPrices200ResponseAllOfPrice) HasCost() bool`

HasCost returns a boolean if a field has been set.

### SetCostNil

`func (o *AddPrices200ResponseAllOfPrice) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *AddPrices200ResponseAllOfPrice) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetCurrency

`func (o *AddPrices200ResponseAllOfPrice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AddPrices200ResponseAllOfPrice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AddPrices200ResponseAllOfPrice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *AddPrices200ResponseAllOfPrice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetIncurCharges

`func (o *AddPrices200ResponseAllOfPrice) GetIncurCharges() string`

GetIncurCharges returns the IncurCharges field if non-nil, zero value otherwise.

### GetIncurChargesOk

`func (o *AddPrices200ResponseAllOfPrice) GetIncurChargesOk() (*string, bool)`

GetIncurChargesOk returns a tuple with the IncurCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncurCharges

`func (o *AddPrices200ResponseAllOfPrice) SetIncurCharges(v string)`

SetIncurCharges sets IncurCharges field to given value.

### HasIncurCharges

`func (o *AddPrices200ResponseAllOfPrice) HasIncurCharges() bool`

HasIncurCharges returns a boolean if a field has been set.

### GetPlatform

`func (o *AddPrices200ResponseAllOfPrice) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *AddPrices200ResponseAllOfPrice) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *AddPrices200ResponseAllOfPrice) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *AddPrices200ResponseAllOfPrice) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *AddPrices200ResponseAllOfPrice) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *AddPrices200ResponseAllOfPrice) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetSoftware

`func (o *AddPrices200ResponseAllOfPrice) GetSoftware() string`

GetSoftware returns the Software field if non-nil, zero value otherwise.

### GetSoftwareOk

`func (o *AddPrices200ResponseAllOfPrice) GetSoftwareOk() (*string, bool)`

GetSoftwareOk returns a tuple with the Software field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftware

`func (o *AddPrices200ResponseAllOfPrice) SetSoftware(v string)`

SetSoftware sets Software field to given value.

### HasSoftware

`func (o *AddPrices200ResponseAllOfPrice) HasSoftware() bool`

HasSoftware returns a boolean if a field has been set.

### SetSoftwareNil

`func (o *AddPrices200ResponseAllOfPrice) SetSoftwareNil(b bool)`

 SetSoftwareNil sets the value for Software to be an explicit nil

### UnsetSoftware
`func (o *AddPrices200ResponseAllOfPrice) UnsetSoftware()`

UnsetSoftware ensures that no value is present for Software, not even an explicit nil
### GetVolumeType

`func (o *AddPrices200ResponseAllOfPrice) GetVolumeType() AddPrices200ResponseAllOfPriceVolumeType`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *AddPrices200ResponseAllOfPrice) GetVolumeTypeOk() (*AddPrices200ResponseAllOfPriceVolumeType, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *AddPrices200ResponseAllOfPrice) SetVolumeType(v AddPrices200ResponseAllOfPriceVolumeType)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *AddPrices200ResponseAllOfPrice) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetDatastore

`func (o *AddPrices200ResponseAllOfPrice) GetDatastore() AddPrices200ResponseAllOfPriceDatastore`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *AddPrices200ResponseAllOfPrice) GetDatastoreOk() (*AddPrices200ResponseAllOfPriceDatastore, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *AddPrices200ResponseAllOfPrice) SetDatastore(v AddPrices200ResponseAllOfPriceDatastore)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *AddPrices200ResponseAllOfPrice) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetCrossCloudApply

`func (o *AddPrices200ResponseAllOfPrice) GetCrossCloudApply() bool`

GetCrossCloudApply returns the CrossCloudApply field if non-nil, zero value otherwise.

### GetCrossCloudApplyOk

`func (o *AddPrices200ResponseAllOfPrice) GetCrossCloudApplyOk() (*bool, bool)`

GetCrossCloudApplyOk returns a tuple with the CrossCloudApply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossCloudApply

`func (o *AddPrices200ResponseAllOfPrice) SetCrossCloudApply(v bool)`

SetCrossCloudApply sets CrossCloudApply field to given value.

### HasCrossCloudApply

`func (o *AddPrices200ResponseAllOfPrice) HasCrossCloudApply() bool`

HasCrossCloudApply returns a boolean if a field has been set.

### SetCrossCloudApplyNil

`func (o *AddPrices200ResponseAllOfPrice) SetCrossCloudApplyNil(b bool)`

 SetCrossCloudApplyNil sets the value for CrossCloudApply to be an explicit nil

### UnsetCrossCloudApply
`func (o *AddPrices200ResponseAllOfPrice) UnsetCrossCloudApply()`

UnsetCrossCloudApply ensures that no value is present for CrossCloudApply, not even an explicit nil
### GetAccount

`func (o *AddPrices200ResponseAllOfPrice) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AddPrices200ResponseAllOfPrice) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AddPrices200ResponseAllOfPrice) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AddPrices200ResponseAllOfPrice) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *AddPrices200ResponseAllOfPrice) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *AddPrices200ResponseAllOfPrice) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


