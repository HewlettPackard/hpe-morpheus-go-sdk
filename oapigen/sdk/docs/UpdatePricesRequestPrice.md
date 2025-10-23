# UpdatePricesRequestPrice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Price name | [optional] 
**Code** | Pointer to **string** | Price code, must be unique | [optional] 
**Account** | Pointer to [**AddPricesRequestPriceAccount**](AddPricesRequestPriceAccount.md) |  | [optional] 
**PriceType** | Pointer to **string** | Restricts query to only load only prices with specified priceType. * &#x60;fixed&#x60; - Everything * &#x60;compute&#x60; - Memory + CPU * &#x60;memory&#x60; - Memory * &#x60;cores&#x60; - Cores * &#x60;storage&#x60; - Storage * &#x60;datastore&#x60; - Datastore * &#x60;platform&#x60; - Platform * &#x60;software&#x60; - Software * &#x60;load_balancer&#x60; - Load Balancer * &#x60;load_balancer_virtual_server&#x60; - Load Balancer Virtual Server  | [optional] 
**PriceUnit** | Pointer to **string** | The unit of pricing | [optional] 
**IncurCharges** | Pointer to **string** | Indicates when to incur charge | [optional] 
**Currency** | Pointer to **string** | ISO Currency code | [optional] 
**Cost** | Pointer to **float32** | Cost | [optional] 
**MarkupType** | Pointer to **string** | Price adjustment type | [optional] 
**Markup** | Pointer to **float32** | Amount for &#x60;fixed&#x60; price adjustment type | [optional] 
**MarkupPercent** | Pointer to **float32** | Percent for &#x60;percent&#x60; price adjustment type | [optional] 
**CustomPrice** | Pointer to **float32** | Custom price for &#x60;custom&#x60; price adjustment type | [optional] 
**Platform** | Pointer to **string** | Platform.  Required for &#x60;platform&#x60; price type | [optional] 
**Software** | Pointer to **string** | Software.  Required for software price type | [optional] 
**VolumeType** | Pointer to [**AddPricesRequestPriceVolumeType**](AddPricesRequestPriceVolumeType.md) |  | [optional] 
**Datastore** | Pointer to [**AddPricesRequestPriceDatastore**](AddPricesRequestPriceDatastore.md) |  | [optional] 
**CrossCloudApply** | Pointer to **bool** | Apply price across clouds, optional true/false flag for datastore price type | [optional] 

## Methods

### NewUpdatePricesRequestPrice

`func NewUpdatePricesRequestPrice() *UpdatePricesRequestPrice`

NewUpdatePricesRequestPrice instantiates a new UpdatePricesRequestPrice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePricesRequestPriceWithDefaults

`func NewUpdatePricesRequestPriceWithDefaults() *UpdatePricesRequestPrice`

NewUpdatePricesRequestPriceWithDefaults instantiates a new UpdatePricesRequestPrice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdatePricesRequestPrice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatePricesRequestPrice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatePricesRequestPrice) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdatePricesRequestPrice) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *UpdatePricesRequestPrice) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdatePricesRequestPrice) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdatePricesRequestPrice) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdatePricesRequestPrice) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetAccount

`func (o *UpdatePricesRequestPrice) GetAccount() AddPricesRequestPriceAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UpdatePricesRequestPrice) GetAccountOk() (*AddPricesRequestPriceAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UpdatePricesRequestPrice) SetAccount(v AddPricesRequestPriceAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UpdatePricesRequestPrice) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetPriceType

`func (o *UpdatePricesRequestPrice) GetPriceType() string`

GetPriceType returns the PriceType field if non-nil, zero value otherwise.

### GetPriceTypeOk

`func (o *UpdatePricesRequestPrice) GetPriceTypeOk() (*string, bool)`

GetPriceTypeOk returns a tuple with the PriceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceType

`func (o *UpdatePricesRequestPrice) SetPriceType(v string)`

SetPriceType sets PriceType field to given value.

### HasPriceType

`func (o *UpdatePricesRequestPrice) HasPriceType() bool`

HasPriceType returns a boolean if a field has been set.

### GetPriceUnit

`func (o *UpdatePricesRequestPrice) GetPriceUnit() string`

GetPriceUnit returns the PriceUnit field if non-nil, zero value otherwise.

### GetPriceUnitOk

`func (o *UpdatePricesRequestPrice) GetPriceUnitOk() (*string, bool)`

GetPriceUnitOk returns a tuple with the PriceUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceUnit

`func (o *UpdatePricesRequestPrice) SetPriceUnit(v string)`

SetPriceUnit sets PriceUnit field to given value.

### HasPriceUnit

`func (o *UpdatePricesRequestPrice) HasPriceUnit() bool`

HasPriceUnit returns a boolean if a field has been set.

### GetIncurCharges

`func (o *UpdatePricesRequestPrice) GetIncurCharges() string`

GetIncurCharges returns the IncurCharges field if non-nil, zero value otherwise.

### GetIncurChargesOk

`func (o *UpdatePricesRequestPrice) GetIncurChargesOk() (*string, bool)`

GetIncurChargesOk returns a tuple with the IncurCharges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncurCharges

`func (o *UpdatePricesRequestPrice) SetIncurCharges(v string)`

SetIncurCharges sets IncurCharges field to given value.

### HasIncurCharges

`func (o *UpdatePricesRequestPrice) HasIncurCharges() bool`

HasIncurCharges returns a boolean if a field has been set.

### GetCurrency

`func (o *UpdatePricesRequestPrice) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *UpdatePricesRequestPrice) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *UpdatePricesRequestPrice) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *UpdatePricesRequestPrice) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetCost

`func (o *UpdatePricesRequestPrice) GetCost() float32`

GetCost returns the Cost field if non-nil, zero value otherwise.

### GetCostOk

`func (o *UpdatePricesRequestPrice) GetCostOk() (*float32, bool)`

GetCostOk returns a tuple with the Cost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCost

`func (o *UpdatePricesRequestPrice) SetCost(v float32)`

SetCost sets Cost field to given value.

### HasCost

`func (o *UpdatePricesRequestPrice) HasCost() bool`

HasCost returns a boolean if a field has been set.

### GetMarkupType

`func (o *UpdatePricesRequestPrice) GetMarkupType() string`

GetMarkupType returns the MarkupType field if non-nil, zero value otherwise.

### GetMarkupTypeOk

`func (o *UpdatePricesRequestPrice) GetMarkupTypeOk() (*string, bool)`

GetMarkupTypeOk returns a tuple with the MarkupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkupType

`func (o *UpdatePricesRequestPrice) SetMarkupType(v string)`

SetMarkupType sets MarkupType field to given value.

### HasMarkupType

`func (o *UpdatePricesRequestPrice) HasMarkupType() bool`

HasMarkupType returns a boolean if a field has been set.

### GetMarkup

`func (o *UpdatePricesRequestPrice) GetMarkup() float32`

GetMarkup returns the Markup field if non-nil, zero value otherwise.

### GetMarkupOk

`func (o *UpdatePricesRequestPrice) GetMarkupOk() (*float32, bool)`

GetMarkupOk returns a tuple with the Markup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkup

`func (o *UpdatePricesRequestPrice) SetMarkup(v float32)`

SetMarkup sets Markup field to given value.

### HasMarkup

`func (o *UpdatePricesRequestPrice) HasMarkup() bool`

HasMarkup returns a boolean if a field has been set.

### GetMarkupPercent

`func (o *UpdatePricesRequestPrice) GetMarkupPercent() float32`

GetMarkupPercent returns the MarkupPercent field if non-nil, zero value otherwise.

### GetMarkupPercentOk

`func (o *UpdatePricesRequestPrice) GetMarkupPercentOk() (*float32, bool)`

GetMarkupPercentOk returns a tuple with the MarkupPercent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMarkupPercent

`func (o *UpdatePricesRequestPrice) SetMarkupPercent(v float32)`

SetMarkupPercent sets MarkupPercent field to given value.

### HasMarkupPercent

`func (o *UpdatePricesRequestPrice) HasMarkupPercent() bool`

HasMarkupPercent returns a boolean if a field has been set.

### GetCustomPrice

`func (o *UpdatePricesRequestPrice) GetCustomPrice() float32`

GetCustomPrice returns the CustomPrice field if non-nil, zero value otherwise.

### GetCustomPriceOk

`func (o *UpdatePricesRequestPrice) GetCustomPriceOk() (*float32, bool)`

GetCustomPriceOk returns a tuple with the CustomPrice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomPrice

`func (o *UpdatePricesRequestPrice) SetCustomPrice(v float32)`

SetCustomPrice sets CustomPrice field to given value.

### HasCustomPrice

`func (o *UpdatePricesRequestPrice) HasCustomPrice() bool`

HasCustomPrice returns a boolean if a field has been set.

### GetPlatform

`func (o *UpdatePricesRequestPrice) GetPlatform() string`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *UpdatePricesRequestPrice) GetPlatformOk() (*string, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *UpdatePricesRequestPrice) SetPlatform(v string)`

SetPlatform sets Platform field to given value.

### HasPlatform

`func (o *UpdatePricesRequestPrice) HasPlatform() bool`

HasPlatform returns a boolean if a field has been set.

### GetSoftware

`func (o *UpdatePricesRequestPrice) GetSoftware() string`

GetSoftware returns the Software field if non-nil, zero value otherwise.

### GetSoftwareOk

`func (o *UpdatePricesRequestPrice) GetSoftwareOk() (*string, bool)`

GetSoftwareOk returns a tuple with the Software field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSoftware

`func (o *UpdatePricesRequestPrice) SetSoftware(v string)`

SetSoftware sets Software field to given value.

### HasSoftware

`func (o *UpdatePricesRequestPrice) HasSoftware() bool`

HasSoftware returns a boolean if a field has been set.

### GetVolumeType

`func (o *UpdatePricesRequestPrice) GetVolumeType() AddPricesRequestPriceVolumeType`

GetVolumeType returns the VolumeType field if non-nil, zero value otherwise.

### GetVolumeTypeOk

`func (o *UpdatePricesRequestPrice) GetVolumeTypeOk() (*AddPricesRequestPriceVolumeType, bool)`

GetVolumeTypeOk returns a tuple with the VolumeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeType

`func (o *UpdatePricesRequestPrice) SetVolumeType(v AddPricesRequestPriceVolumeType)`

SetVolumeType sets VolumeType field to given value.

### HasVolumeType

`func (o *UpdatePricesRequestPrice) HasVolumeType() bool`

HasVolumeType returns a boolean if a field has been set.

### GetDatastore

`func (o *UpdatePricesRequestPrice) GetDatastore() AddPricesRequestPriceDatastore`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *UpdatePricesRequestPrice) GetDatastoreOk() (*AddPricesRequestPriceDatastore, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *UpdatePricesRequestPrice) SetDatastore(v AddPricesRequestPriceDatastore)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *UpdatePricesRequestPrice) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetCrossCloudApply

`func (o *UpdatePricesRequestPrice) GetCrossCloudApply() bool`

GetCrossCloudApply returns the CrossCloudApply field if non-nil, zero value otherwise.

### GetCrossCloudApplyOk

`func (o *UpdatePricesRequestPrice) GetCrossCloudApplyOk() (*bool, bool)`

GetCrossCloudApplyOk returns a tuple with the CrossCloudApply field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrossCloudApply

`func (o *UpdatePricesRequestPrice) SetCrossCloudApply(v bool)`

SetCrossCloudApply sets CrossCloudApply field to given value.

### HasCrossCloudApply

`func (o *UpdatePricesRequestPrice) HasCrossCloudApply() bool`

HasCrossCloudApply returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


