# ListPrices200ResponseAllOfPricesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**PriceType** | Pointer to **string** |  | [optional] 
**PriceUnit** | Pointer to **string** |  | [optional] 
**AdditionalPriceUnit** | Pointer to **string** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**CustomPrice** | Pointer to **float32** |  | [optional] 
**MarkupType** | Pointer to **string** |  | [optional] 
**Markup** | Pointer to **float32** |  | [optional] 
**MarkupPercent** | Pointer to **float32** |  | [optional] 
**Cost** | Pointer to **float32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**IncurCharges** | Pointer to **string** |  | [optional] 
**Platform** | Pointer to **string** |  | [optional] 
**Software** | Pointer to **string** |  | [optional] 
**VolumeType** | Pointer to [**ListPriceSets200ResponseAllOfPriceSetsInnerPricesInnerVolumeType**](ListPriceSets200ResponseAllOfPriceSetsInnerPricesInnerVolumeType.md) |  | [optional] 
**Datastore** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**CrossCloudApply** | Pointer to **bool** |  | [optional] 
**Account** | Pointer to **string** |  | [optional] 

## Methods

### NewListPrices200ResponseAllOfPricesInner

`func NewListPrices200ResponseAllOfPricesInner() *ListPrices200ResponseAllOfPricesInner`

NewListPrices200ResponseAllOfPricesInner instantiates a new ListPrices200ResponseAllOfPricesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPrices200ResponseAllOfPricesInnerWithDefaults

`func NewListPrices200ResponseAllOfPricesInnerWithDefaults() *ListPrices200ResponseAllOfPricesInner`

NewListPrices200ResponseAllOfPricesInnerWithDefaults instantiates a new ListPrices200ResponseAllOfPricesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListPrices200ResponseAllOfPricesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListPrices200ResponseAllOfPricesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListPrices200ResponseAllOfPricesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListPrices200ResponseAllOfPricesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListPrices200ResponseAllOfPricesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListPrices200ResponseAllOfPricesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListPrices200ResponseAllOfPricesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListPrices200ResponseAllOfPricesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *ListPrices200ResponseAllOfPricesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListPrices200ResponseAllOfPricesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListPrices200ResponseAllOfPricesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListPrices200ResponseAllOfPricesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetActive

`func (o *ListPrices200ResponseAllOfPricesInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ListPrices200ResponseAllOfPricesInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ListPrices200ResponseAllOfPricesInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ListPrices200ResponseAllOfPricesInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetPriceType

`func (o *ListPrices200ResponseAllOfPricesInner) GetPriceType() string`

GetPriceType returns the PriceType field if non-nil, zero value otherwise.

### GetPriceTypeOk

`func (o *ListPrices200ResponseAllOfPricesInner) GetPriceTypeOk() (*string, bool)`

GetPriceTypeOk returns a tuple with the PriceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceType

`func (o *ListPrices200ResponseAllOfPricesInner) SetPriceType(v string)`

SetPriceType sets PriceType field to given value.

### HasPriceType

`func (o *ListPrices200ResponseAllOfPricesInner) HasPriceType() bool`

HasPriceType returns a boolean if a field has been set.

### GetPriceUnit

`func (o *ListPrices200ResponseAllOfPricesInner) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *ListPrices200ResponseAllOfPricesInner) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *ListPrices200ResponseAllOfPricesInner) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *ListPrices200ResponseAllOfPricesInner) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### GetAdditionalPriceUnit

`func (o *ListPrices200ResponseAllOfPricesInner) GetAdditionalPriceUnit() string`

GetAdditionalPriceUnit returns the AdditionalPriceUnit field if non-nil, zero value otherwise.

### GetAdditionalPriceUnitOk

`func (o *ListPrices200ResponseAllOfPricesInner) GetAdditionalPriceUnitOk() (*string, bool)`

GetAdditionalPriceUnitOk returns a tuple with the AdditionalPriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdditionalPriceUnit

`func (o *ListPrices200ResponseAllOfPricesInner) SetAdditionalPriceUnit(v string)`

SetAdditionalPriceUnit sets AdditionalPriceUnit field to given value.

### HasAdditionalPriceUnit

`func (o *ListPrices200ResponseAllOfPricesInner) HasAdditionalPriceUnit() bool`

HasAdditionalPriceUnit returns a boolean if a field has been set.

### GetPrice

`func (o *ListPrices200ResponseAllOfPricesInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *ListPrices200ResponseAllOfPricesInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *ListPrices200ResponseAllOfPricesInner) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *ListPrices200ResponseAllOfPricesInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCustomPrice

`func (o *ListPrices200ResponseAllOfPricesInner) GetCustomPrice() float32`

GetCustomPrice returns the CustomPrice field if non-nil, zero value otherwise.

### GetCustomPriceOk

`func (o *ListPrices200ResponseAllOfPricesInner) GetCustomPriceOk() (*float32, bool)`

GetCustomPriceOk returns a tuple with the CustomPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPrice

`func (o *ListPrices200ResponseAllOfPricesInner) SetCustomPrice(v float32)`

SetCustomPrice sets CustomPrice field to given value.

### HasCustomPrice

`func (o *ListPrices200ResponseAllOfPricesInner) HasCustomPrice() bool`

HasCustomPrice returns a boolean if a field has been set.

### GetMarkupType

`func (o *ListPrices200ResponseAllOfPricesInner) GetMarkupType() string`

GetMarkupType returns the MarkupType field if non-nil, zero value otherwise.

### GetMarkupTypeOk

`func (o *ListPrices200ResponseAllOfPricesInner) GetMarkupTypeOk() (*string, bool)`

GetMarkupTypeOk returns a tuple with the MarkupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkupType

`func (o *ListPrices200ResponseAllOfPricesInner) SetMarkupType(v string)`

SetMarkupType sets MarkupType field to given value.

### HasMarkupType

`func (o *ListPrices200ResponseAllOfPricesInner) HasMarkupType() bool`

HasMarkupType returns a boolean if a field has been set.

### GetMarkup

`func (o *ListPrices200ResponseAllOfPricesInner) GetMarkup() float32`

GetMarkup returns the Markup field if non-nil, zero value otherwise.

### GetMarkupOk

`func (o *ListPrices200ResponseAllOfPricesInner) GetMarkupOk() (*float32, bool)`

GetMarkupOk returns a tuple with the Markup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkup

`func (o *ListPrices200ResponseAllOfPricesInner) SetMarkup(v float32)`

SetMarkup sets Markup field to given value.

### HasMarkup

`func (o *ListPrices200ResponseAllOfPricesInner) HasMarkup() bool`

HasMarkup returns a boolean if a field has been set.

### GetMarkupPercent

`func (o *ListPrices200ResponseAllOfPricesInner) GetMarkupPercent() float32`

GetMarkupPercent returns the MarkupPercent field if non-nil, zero value otherwise.

### GetMarkupPercentOk

`func (o *ListPrices200ResponseAllOfPricesInner) GetMarkupPercentOk() (*float32, bool)`

GetMarkupPercentOk returns a tuple with the MarkupPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkupPercent

`func (o *ListPrices200ResponseAllOfPricesInner) SetMarkupPercent(v float32)`

SetMarkupPercent sets MarkupPercent field to given value.

### HasMarkupPercent

`func (o *ListPrices200ResponseAllOfPricesInner) HasMarkupPercent() bool`

HasMarkupPercent returns a boolean if a field has been set.

### GetCost

`func (o *ListPrices200ResponseAllOfPricesInner) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *ListPrices200ResponseAllOfPricesInner) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *ListPrices200ResponseAllOfPricesInner) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *ListPrices200ResponseAllOfPricesInner) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetCurrency

`func (o *ListPrices200ResponseAllOfPricesInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *ListPrices200ResponseAllOfPricesInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *ListPrices200ResponseAllOfPricesInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *ListPrices200ResponseAllOfPricesInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetIncurCharges

`func (o *ListPrices200ResponseAllOfPricesInner) GetIncurCharges() string`

GetIncurCharges returns the IncurCharges field if non-nil, zero value otherwise.

### GetIncurChargesOk

`func (o *ListPrices200ResponseAllOfPricesInner) GetIncurChargesOk() (*string, bool)`

GetIncurChargesOk returns a tuple with the IncurCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncurCharges

`func (o *ListPrices200ResponseAllOfPricesInner) SetIncurCharges(v string)`

SetIncurCharges sets IncurCharges field to given value.

### HasIncurCharges

`func (o *ListPrices200ResponseAllOfPricesInner) HasIncurCharges() bool`

HasIncurCharges returns a boolean if a field has been set.

### GetPlatform

`func (o *ListPrices200ResponseAllOfPricesInner) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *ListPrices200ResponseAllOfPricesInner) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *ListPrices200ResponseAllOfPricesInner) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *ListPrices200ResponseAllOfPricesInner) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSoftware

`func (o *ListPrices200ResponseAllOfPricesInner) GetSoftware() string`

GetSoftware returns the Software field if non-nil, zero value otherwise.

### GetSoftwareOk

`func (o *ListPrices200ResponseAllOfPricesInner) GetSoftwareOk() (*string, bool)`

GetSoftwareOk returns a tuple with the Software field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftware

`func (o *ListPrices200ResponseAllOfPricesInner) SetSoftware(v string)`

SetSoftware sets Software field to given value.

### HasSoftware

`func (o *ListPrices200ResponseAllOfPricesInner) HasSoftware() bool`

HasSoftware returns a boolean if a field has been set.

### GetVolumeType

`func (o *ListPrices200ResponseAllOfPricesInner) GetVolumeType() ListPriceSets200ResponseAllOfPriceSetsInnerPricesInnerVolumeType`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *ListPrices200ResponseAllOfPricesInner) GetVolumeTypeOk() (*ListPriceSets200ResponseAllOfPriceSetsInnerPricesInnerVolumeType, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *ListPrices200ResponseAllOfPricesInner) SetVolumeType(v ListPriceSets200ResponseAllOfPriceSetsInnerPricesInnerVolumeType)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *ListPrices200ResponseAllOfPricesInner) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetDatastore

`func (o *ListPrices200ResponseAllOfPricesInner) GetDatastore() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *ListPrices200ResponseAllOfPricesInner) GetDatastoreOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *ListPrices200ResponseAllOfPricesInner) SetDatastore(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *ListPrices200ResponseAllOfPricesInner) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetCrossCloudApply

`func (o *ListPrices200ResponseAllOfPricesInner) GetCrossCloudApply() bool`

GetCrossCloudApply returns the CrossCloudApply field if non-nil, zero value otherwise.

### GetCrossCloudApplyOk

`func (o *ListPrices200ResponseAllOfPricesInner) GetCrossCloudApplyOk() (*bool, bool)`

GetCrossCloudApplyOk returns a tuple with the CrossCloudApply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossCloudApply

`func (o *ListPrices200ResponseAllOfPricesInner) SetCrossCloudApply(v bool)`

SetCrossCloudApply sets CrossCloudApply field to given value.

### HasCrossCloudApply

`func (o *ListPrices200ResponseAllOfPricesInner) HasCrossCloudApply() bool`

HasCrossCloudApply returns a boolean if a field has been set.

### GetAccount

`func (o *ListPrices200ResponseAllOfPricesInner) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListPrices200ResponseAllOfPricesInner) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListPrices200ResponseAllOfPricesInner) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListPrices200ResponseAllOfPricesInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


