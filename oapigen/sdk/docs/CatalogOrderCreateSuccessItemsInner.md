# CatalogOrderCreateSuccessItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to [**CatalogOrderCreateSuccessItemsInnerType**](CatalogOrderCreateSuccessItemsInnerType.md) |  | [optional] 
**Quantity** | Pointer to **int64** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**Valid** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewCatalogOrderCreateSuccessItemsInner

`func NewCatalogOrderCreateSuccessItemsInner() *CatalogOrderCreateSuccessItemsInner`

NewCatalogOrderCreateSuccessItemsInner instantiates a new CatalogOrderCreateSuccessItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *CatalogOrderCreateSuccessItemsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogOrderCreateSuccessItemsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogOrderCreateSuccessItemsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CatalogOrderCreateSuccessItemsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *CatalogOrderCreateSuccessItemsInner) GetType() CatalogOrderCreateSuccessItemsInnerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogOrderCreateSuccessItemsInner) GetTypeOk() (*CatalogOrderCreateSuccessItemsInnerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogOrderCreateSuccessItemsInner) SetType(v CatalogOrderCreateSuccessItemsInnerType)`

SetType sets Type field to given value.

### HasType

`func (o *CatalogOrderCreateSuccessItemsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetQuantity

`func (o *CatalogOrderCreateSuccessItemsInner) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *CatalogOrderCreateSuccessItemsInner) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *CatalogOrderCreateSuccessItemsInner) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *CatalogOrderCreateSuccessItemsInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetPrice

`func (o *CatalogOrderCreateSuccessItemsInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *CatalogOrderCreateSuccessItemsInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *CatalogOrderCreateSuccessItemsInner) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *CatalogOrderCreateSuccessItemsInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCurrency

`func (o *CatalogOrderCreateSuccessItemsInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *CatalogOrderCreateSuccessItemsInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *CatalogOrderCreateSuccessItemsInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *CatalogOrderCreateSuccessItemsInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetUnit

`func (o *CatalogOrderCreateSuccessItemsInner) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *CatalogOrderCreateSuccessItemsInner) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *CatalogOrderCreateSuccessItemsInner) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *CatalogOrderCreateSuccessItemsInner) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetValid

`func (o *CatalogOrderCreateSuccessItemsInner) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *CatalogOrderCreateSuccessItemsInner) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *CatalogOrderCreateSuccessItemsInner) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *CatalogOrderCreateSuccessItemsInner) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetStatus

`func (o *CatalogOrderCreateSuccessItemsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CatalogOrderCreateSuccessItemsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CatalogOrderCreateSuccessItemsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CatalogOrderCreateSuccessItemsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDateCreated

`func (o *CatalogOrderCreateSuccessItemsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *CatalogOrderCreateSuccessItemsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *CatalogOrderCreateSuccessItemsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *CatalogOrderCreateSuccessItemsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *CatalogOrderCreateSuccessItemsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *CatalogOrderCreateSuccessItemsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *CatalogOrderCreateSuccessItemsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *CatalogOrderCreateSuccessItemsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


