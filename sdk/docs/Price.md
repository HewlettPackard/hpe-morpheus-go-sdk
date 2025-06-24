# Price

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
**VolumeType** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**Datastore** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**CrossCloudApply** | Pointer to **NullableBool** |  | [optional] 
**Account** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewPrice

`func NewPrice() *Price`

NewPrice instantiates a new Price object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPriceWithDefaults

`func NewPriceWithDefaults() *Price`

NewPriceWithDefaults instantiates a new Price object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Price) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Price) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Price) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *Price) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Price) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Price) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Price) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Price) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *Price) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *Price) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *Price) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *Price) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetActive

`func (o *Price) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Price) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Price) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *Price) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPriceType

`func (o *Price) GetPriceType() string`

GetPriceType returns the PriceType field if non-nil, zero value otherwise.

### GetPriceTypeOk

`func (o *Price) GetPriceTypeOk() (*string, bool)`

GetPriceTypeOk returns a tuple with the PriceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceType

`func (o *Price) SetPriceType(v string)`

SetPriceType sets PriceType field to given value.

### HasPriceType

`func (o *Price) HasPriceType() bool`

HasPriceType returns a boolean if a field has been set.

### GetPriceUnit

`func (o *Price) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *Price) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *Price) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *Price) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### GetAdditionalPriceUnit

`func (o *Price) GetAdditionalPriceUnit() string`

GetAdditionalPriceUnit returns the AdditionalPriceUnit field if non-nil, zero value otherwise.

### GetAdditionalPriceUnitOk

`func (o *Price) GetAdditionalPriceUnitOk() (*string, bool)`

GetAdditionalPriceUnitOk returns a tuple with the AdditionalPriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalPriceUnit

`func (o *Price) SetAdditionalPriceUnit(v string)`

SetAdditionalPriceUnit sets AdditionalPriceUnit field to given value.

### HasAdditionalPriceUnit

`func (o *Price) HasAdditionalPriceUnit() bool`

HasAdditionalPriceUnit returns a boolean if a field has been set.

### SetAdditionalPriceUnitNil

`func (o *Price) SetAdditionalPriceUnitNil(b bool)`

 SetAdditionalPriceUnitNil sets the value for AdditionalPriceUnit to be an explicit nil

### UnsetAdditionalPriceUnit
`func (o *Price) UnsetAdditionalPriceUnit()`

UnsetAdditionalPriceUnit ensures that no value is present for AdditionalPriceUnit, not even an explicit nil
### GetPrice

`func (o *Price) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *Price) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *Price) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *Price) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### SetPriceNil

`func (o *Price) SetPriceNil(b bool)`

 SetPriceNil sets the value for Price to be an explicit nil

### UnsetPrice
`func (o *Price) UnsetPrice()`

UnsetPrice ensures that no value is present for Price, not even an explicit nil
### GetCustomPrice

`func (o *Price) GetCustomPrice() float32`

GetCustomPrice returns the CustomPrice field if non-nil, zero value otherwise.

### GetCustomPriceOk

`func (o *Price) GetCustomPriceOk() (*float32, bool)`

GetCustomPriceOk returns a tuple with the CustomPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPrice

`func (o *Price) SetCustomPrice(v float32)`

SetCustomPrice sets CustomPrice field to given value.

### HasCustomPrice

`func (o *Price) HasCustomPrice() bool`

HasCustomPrice returns a boolean if a field has been set.

### SetCustomPriceNil

`func (o *Price) SetCustomPriceNil(b bool)`

 SetCustomPriceNil sets the value for CustomPrice to be an explicit nil

### UnsetCustomPrice
`func (o *Price) UnsetCustomPrice()`

UnsetCustomPrice ensures that no value is present for CustomPrice, not even an explicit nil
### GetMarkupType

`func (o *Price) GetMarkupType() string`

GetMarkupType returns the MarkupType field if non-nil, zero value otherwise.

### GetMarkupTypeOk

`func (o *Price) GetMarkupTypeOk() (*string, bool)`

GetMarkupTypeOk returns a tuple with the MarkupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkupType

`func (o *Price) SetMarkupType(v string)`

SetMarkupType sets MarkupType field to given value.

### HasMarkupType

`func (o *Price) HasMarkupType() bool`

HasMarkupType returns a boolean if a field has been set.

### SetMarkupTypeNil

`func (o *Price) SetMarkupTypeNil(b bool)`

 SetMarkupTypeNil sets the value for MarkupType to be an explicit nil

### UnsetMarkupType
`func (o *Price) UnsetMarkupType()`

UnsetMarkupType ensures that no value is present for MarkupType, not even an explicit nil
### GetMarkup

`func (o *Price) GetMarkup() float32`

GetMarkup returns the Markup field if non-nil, zero value otherwise.

### GetMarkupOk

`func (o *Price) GetMarkupOk() (*float32, bool)`

GetMarkupOk returns a tuple with the Markup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkup

`func (o *Price) SetMarkup(v float32)`

SetMarkup sets Markup field to given value.

### HasMarkup

`func (o *Price) HasMarkup() bool`

HasMarkup returns a boolean if a field has been set.

### SetMarkupNil

`func (o *Price) SetMarkupNil(b bool)`

 SetMarkupNil sets the value for Markup to be an explicit nil

### UnsetMarkup
`func (o *Price) UnsetMarkup()`

UnsetMarkup ensures that no value is present for Markup, not even an explicit nil
### GetMarkupPercent

`func (o *Price) GetMarkupPercent() float32`

GetMarkupPercent returns the MarkupPercent field if non-nil, zero value otherwise.

### GetMarkupPercentOk

`func (o *Price) GetMarkupPercentOk() (*float32, bool)`

GetMarkupPercentOk returns a tuple with the MarkupPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkupPercent

`func (o *Price) SetMarkupPercent(v float32)`

SetMarkupPercent sets MarkupPercent field to given value.

### HasMarkupPercent

`func (o *Price) HasMarkupPercent() bool`

HasMarkupPercent returns a boolean if a field has been set.

### SetMarkupPercentNil

`func (o *Price) SetMarkupPercentNil(b bool)`

 SetMarkupPercentNil sets the value for MarkupPercent to be an explicit nil

### UnsetMarkupPercent
`func (o *Price) UnsetMarkupPercent()`

UnsetMarkupPercent ensures that no value is present for MarkupPercent, not even an explicit nil
### GetCost

`func (o *Price) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *Price) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *Price) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *Price) HasCost() bool`

HasCost returns a boolean if a field has been set.

### SetCostNil

`func (o *Price) SetCostNil(b bool)`

 SetCostNil sets the value for Cost to be an explicit nil

### UnsetCost
`func (o *Price) UnsetCost()`

UnsetCost ensures that no value is present for Cost, not even an explicit nil
### GetCurrency

`func (o *Price) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *Price) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *Price) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *Price) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetIncurCharges

`func (o *Price) GetIncurCharges() string`

GetIncurCharges returns the IncurCharges field if non-nil, zero value otherwise.

### GetIncurChargesOk

`func (o *Price) GetIncurChargesOk() (*string, bool)`

GetIncurChargesOk returns a tuple with the IncurCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncurCharges

`func (o *Price) SetIncurCharges(v string)`

SetIncurCharges sets IncurCharges field to given value.

### HasIncurCharges

`func (o *Price) HasIncurCharges() bool`

HasIncurCharges returns a boolean if a field has been set.

### GetPlatform

`func (o *Price) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *Price) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *Price) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *Price) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### SetPlatformNil

`func (o *Price) SetPlatformNil(b bool)`

 SetPlatformNil sets the value for Platform to be an explicit nil

### UnsetPlatform
`func (o *Price) UnsetPlatform()`

UnsetPlatform ensures that no value is present for Platform, not even an explicit nil
### GetSoftware

`func (o *Price) GetSoftware() string`

GetSoftware returns the Software field if non-nil, zero value otherwise.

### GetSoftwareOk

`func (o *Price) GetSoftwareOk() (*string, bool)`

GetSoftwareOk returns a tuple with the Software field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftware

`func (o *Price) SetSoftware(v string)`

SetSoftware sets Software field to given value.

### HasSoftware

`func (o *Price) HasSoftware() bool`

HasSoftware returns a boolean if a field has been set.

### SetSoftwareNil

`func (o *Price) SetSoftwareNil(b bool)`

 SetSoftwareNil sets the value for Software to be an explicit nil

### UnsetSoftware
`func (o *Price) UnsetSoftware()`

UnsetSoftware ensures that no value is present for Software, not even an explicit nil
### GetVolumeType

`func (o *Price) GetVolumeType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *Price) GetVolumeTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *Price) SetVolumeType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *Price) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetDatastore

`func (o *Price) GetDatastore() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *Price) GetDatastoreOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *Price) SetDatastore(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *Price) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetCrossCloudApply

`func (o *Price) GetCrossCloudApply() bool`

GetCrossCloudApply returns the CrossCloudApply field if non-nil, zero value otherwise.

### GetCrossCloudApplyOk

`func (o *Price) GetCrossCloudApplyOk() (*bool, bool)`

GetCrossCloudApplyOk returns a tuple with the CrossCloudApply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossCloudApply

`func (o *Price) SetCrossCloudApply(v bool)`

SetCrossCloudApply sets CrossCloudApply field to given value.

### HasCrossCloudApply

`func (o *Price) HasCrossCloudApply() bool`

HasCrossCloudApply returns a boolean if a field has been set.

### SetCrossCloudApplyNil

`func (o *Price) SetCrossCloudApplyNil(b bool)`

 SetCrossCloudApplyNil sets the value for CrossCloudApply to be an explicit nil

### UnsetCrossCloudApply
`func (o *Price) UnsetCrossCloudApply()`

UnsetCrossCloudApply ensures that no value is present for CrossCloudApply, not even an explicit nil
### GetAccount

`func (o *Price) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *Price) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *Price) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *Price) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *Price) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *Price) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


