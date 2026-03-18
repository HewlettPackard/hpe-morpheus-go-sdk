# ListCatalogItems200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]ListCatalogItems200ResponseAllOfItemsInner**](ListCatalogItems200ResponseAllOfItemsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListAlerts200ResponseAllOfMeta**](ListAlerts200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListCatalogItems200Response

`func NewListCatalogItems200Response() *ListCatalogItems200Response`

NewListCatalogItems200Response instantiates a new ListCatalogItems200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCatalogItems200ResponseWithDefaults

`func NewListCatalogItems200ResponseWithDefaults() *ListCatalogItems200Response`

NewListCatalogItems200ResponseWithDefaults instantiates a new ListCatalogItems200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ListCatalogItems200Response) GetItems() []ListCatalogItems200ResponseAllOfItemsInner`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ListCatalogItems200Response) GetItemsOk() (*[]ListCatalogItems200ResponseAllOfItemsInner, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ListCatalogItems200Response) SetItems(v []ListCatalogItems200ResponseAllOfItemsInner)`

SetItems sets Items field to given value.

### HasItems

`func (o *ListCatalogItems200Response) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetMeta

`func (o *ListCatalogItems200Response) GetMeta() ListAlerts200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListCatalogItems200Response) GetMetaOk() (*ListAlerts200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListCatalogItems200Response) SetMeta(v ListAlerts200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListCatalogItems200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


