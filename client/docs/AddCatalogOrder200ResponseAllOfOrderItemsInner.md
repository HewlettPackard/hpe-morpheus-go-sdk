# AddCatalogOrder200ResponseAllOfOrderItemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Type** | Pointer to [**ListApps200ResponseAllOfAppsInnerBlueprint**](ListApps200ResponseAllOfAppsInnerBlueprint.md) |  | [optional] 
**Quantity** | Pointer to **int64** |  | [optional] 
**Price** | Pointer to **float32** |  | [optional] 
**Currency** | Pointer to **string** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**Valid** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAddCatalogOrder200ResponseAllOfOrderItemsInner

`func NewAddCatalogOrder200ResponseAllOfOrderItemsInner() *AddCatalogOrder200ResponseAllOfOrderItemsInner`

NewAddCatalogOrder200ResponseAllOfOrderItemsInner instantiates a new AddCatalogOrder200ResponseAllOfOrderItemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCatalogOrder200ResponseAllOfOrderItemsInnerWithDefaults

`func NewAddCatalogOrder200ResponseAllOfOrderItemsInnerWithDefaults() *AddCatalogOrder200ResponseAllOfOrderItemsInner`

NewAddCatalogOrder200ResponseAllOfOrderItemsInnerWithDefaults instantiates a new AddCatalogOrder200ResponseAllOfOrderItemsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) GetType() ListApps200ResponseAllOfAppsInnerBlueprint`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) GetTypeOk() (*ListApps200ResponseAllOfAppsInnerBlueprint, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) SetType(v ListApps200ResponseAllOfAppsInnerBlueprint)`

SetType sets Type field to given value.

### HasType

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetQuantity

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) GetQuantity() int64`

GetQuantity returns the Quantity field if non-nil, zero value otherwise.

### GetQuantityOk

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) GetQuantityOk() (*int64, bool)`

GetQuantityOk returns a tuple with the Quantity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuantity

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) SetQuantity(v int64)`

SetQuantity sets Quantity field to given value.

### HasQuantity

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) HasQuantity() bool`

HasQuantity returns a boolean if a field has been set.

### GetPrice

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) GetPrice() float32`

GetPrice returns the Price field if non-nil, zero value otherwise.

### GetPriceOk

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) GetPriceOk() (*float32, bool)`

GetPriceOk returns a tuple with the Price field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrice

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) SetPrice(v float32)`

SetPrice sets Price field to given value.

### HasPrice

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) HasPrice() bool`

HasPrice returns a boolean if a field has been set.

### GetCurrency

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) GetCurrency() string`

GetCurrency returns the Currency field if non-nil, zero value otherwise.

### GetCurrencyOk

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) GetCurrencyOk() (*string, bool)`

GetCurrencyOk returns a tuple with the Currency field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrency

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) SetCurrency(v string)`

SetCurrency sets Currency field to given value.

### HasCurrency

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) HasCurrency() bool`

HasCurrency returns a boolean if a field has been set.

### GetUnit

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetValid

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetStatus

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDateCreated

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AddCatalogOrder200ResponseAllOfOrderItemsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


