# ListCatalogCart200ResponseCart

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Items** | Pointer to [**[]ListCatalogCart200ResponseCartItemsInner**](ListCatalogCart200ResponseCartItemsInner.md) |  | [optional] 
**Stats** | Pointer to [**ListCatalogCart200ResponseCartStats**](ListCatalogCart200ResponseCartStats.md) |  | [optional] 

## Methods

### NewListCatalogCart200ResponseCart

`func NewListCatalogCart200ResponseCart() *ListCatalogCart200ResponseCart`

NewListCatalogCart200ResponseCart instantiates a new ListCatalogCart200ResponseCart object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCatalogCart200ResponseCartWithDefaults

`func NewListCatalogCart200ResponseCartWithDefaults() *ListCatalogCart200ResponseCart`

NewListCatalogCart200ResponseCartWithDefaults instantiates a new ListCatalogCart200ResponseCart object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListCatalogCart200ResponseCart) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListCatalogCart200ResponseCart) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListCatalogCart200ResponseCart) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListCatalogCart200ResponseCart) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListCatalogCart200ResponseCart) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListCatalogCart200ResponseCart) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListCatalogCart200ResponseCart) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListCatalogCart200ResponseCart) HasName() bool`

HasName returns a boolean if a field has been set.

### GetItems

`func (o *ListCatalogCart200ResponseCart) GetItems() []ListCatalogCart200ResponseCartItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListCatalogCart200ResponseCart) GetItemsOk() (*[]ListCatalogCart200ResponseCartItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListCatalogCart200ResponseCart) SetItems(v []ListCatalogCart200ResponseCartItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListCatalogCart200ResponseCart) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetStats

`func (o *ListCatalogCart200ResponseCart) GetStats() ListCatalogCart200ResponseCartStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ListCatalogCart200ResponseCart) GetStatsOk() (*ListCatalogCart200ResponseCartStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ListCatalogCart200ResponseCart) SetStats(v ListCatalogCart200ResponseCartStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ListCatalogCart200ResponseCart) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


