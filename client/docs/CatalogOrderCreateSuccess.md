# CatalogOrderCreateSuccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]AddCatalogOrder200ResponseAllOfOrderItemsInner**](AddCatalogOrder200ResponseAllOfOrderItemsInner.md) |  | [optional] 
**Stats** | Pointer to [**ListCatalogCart200ResponseCartStats**](ListCatalogCart200ResponseCartStats.md) |  | [optional] 

## Methods

### NewCatalogOrderCreateSuccess

`func NewCatalogOrderCreateSuccess() *CatalogOrderCreateSuccess`

NewCatalogOrderCreateSuccess instantiates a new CatalogOrderCreateSuccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogOrderCreateSuccessWithDefaults

`func NewCatalogOrderCreateSuccessWithDefaults() *CatalogOrderCreateSuccess`

NewCatalogOrderCreateSuccessWithDefaults instantiates a new CatalogOrderCreateSuccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CatalogOrderCreateSuccess) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CatalogOrderCreateSuccess) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CatalogOrderCreateSuccess) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CatalogOrderCreateSuccess) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CatalogOrderCreateSuccess) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogOrderCreateSuccess) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogOrderCreateSuccess) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CatalogOrderCreateSuccess) HasName() bool`

HasName returns a boolean if a field has been set.

### GetItems

`func (o *CatalogOrderCreateSuccess) GetItems() []AddCatalogOrder200ResponseAllOfOrderItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *CatalogOrderCreateSuccess) GetItemsOk() (*[]AddCatalogOrder200ResponseAllOfOrderItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *CatalogOrderCreateSuccess) SetItems(v []AddCatalogOrder200ResponseAllOfOrderItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *CatalogOrderCreateSuccess) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetStats

`func (o *CatalogOrderCreateSuccess) GetStats() ListCatalogCart200ResponseCartStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *CatalogOrderCreateSuccess) GetStatsOk() (*ListCatalogCart200ResponseCartStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *CatalogOrderCreateSuccess) SetStats(v ListCatalogCart200ResponseCartStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *CatalogOrderCreateSuccess) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


