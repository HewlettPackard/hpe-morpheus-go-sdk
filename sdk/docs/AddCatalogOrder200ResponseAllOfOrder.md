# AddCatalogOrder200ResponseAllOfOrder

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]AddCatalogOrder200ResponseAllOfOrderItemsInner**](AddCatalogOrder200ResponseAllOfOrderItemsInner.md) |  | [optional] 
**Stats** | Pointer to [**ListCatalogCart200ResponseCartStats**](ListCatalogCart200ResponseCartStats.md) |  | [optional] 

## Methods

### NewAddCatalogOrder200ResponseAllOfOrder

`func NewAddCatalogOrder200ResponseAllOfOrder() *AddCatalogOrder200ResponseAllOfOrder`

NewAddCatalogOrder200ResponseAllOfOrder instantiates a new AddCatalogOrder200ResponseAllOfOrder object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCatalogOrder200ResponseAllOfOrderWithDefaults

`func NewAddCatalogOrder200ResponseAllOfOrderWithDefaults() *AddCatalogOrder200ResponseAllOfOrder`

NewAddCatalogOrder200ResponseAllOfOrderWithDefaults instantiates a new AddCatalogOrder200ResponseAllOfOrder object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddCatalogOrder200ResponseAllOfOrder) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddCatalogOrder200ResponseAllOfOrder) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddCatalogOrder200ResponseAllOfOrder) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddCatalogOrder200ResponseAllOfOrder) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddCatalogOrder200ResponseAllOfOrder) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddCatalogOrder200ResponseAllOfOrder) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddCatalogOrder200ResponseAllOfOrder) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddCatalogOrder200ResponseAllOfOrder) HasName() bool`

HasName returns a boolean if a field has been set.

### GetItems

`func (o *AddCatalogOrder200ResponseAllOfOrder) GetItems() []AddCatalogOrder200ResponseAllOfOrderItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AddCatalogOrder200ResponseAllOfOrder) GetItemsOk() (*[]AddCatalogOrder200ResponseAllOfOrderItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AddCatalogOrder200ResponseAllOfOrder) SetItems(v []AddCatalogOrder200ResponseAllOfOrderItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *AddCatalogOrder200ResponseAllOfOrder) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetStats

`func (o *AddCatalogOrder200ResponseAllOfOrder) GetStats() ListCatalogCart200ResponseCartStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *AddCatalogOrder200ResponseAllOfOrder) GetStatsOk() (*ListCatalogCart200ResponseCartStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *AddCatalogOrder200ResponseAllOfOrder) SetStats(v ListCatalogCart200ResponseCartStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *AddCatalogOrder200ResponseAllOfOrder) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


