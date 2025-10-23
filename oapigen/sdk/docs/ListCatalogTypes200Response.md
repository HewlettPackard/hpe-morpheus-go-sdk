# ListCatalogTypes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatalogItemTypes** | Pointer to [**[]ListCatalogTypes200ResponseAllOfCatalogItemTypesInner**](ListCatalogTypes200ResponseAllOfCatalogItemTypesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListCatalogTypes200Response

`func NewListCatalogTypes200Response() *ListCatalogTypes200Response`

NewListCatalogTypes200Response instantiates a new ListCatalogTypes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCatalogTypes200ResponseWithDefaults

`func NewListCatalogTypes200ResponseWithDefaults() *ListCatalogTypes200Response`

NewListCatalogTypes200ResponseWithDefaults instantiates a new ListCatalogTypes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCatalogItemTypes

`func (o *ListCatalogTypes200Response) GetCatalogItemTypes() []ListCatalogTypes200ResponseAllOfCatalogItemTypesInner`

GetCatalogItemTypes returns the CatalogItemTypes field if non-nil, zero value otherwise.

### GetCatalogItemTypesOk

`func (o *ListCatalogTypes200Response) GetCatalogItemTypesOk() (*[]ListCatalogTypes200ResponseAllOfCatalogItemTypesInner, bool)`

GetCatalogItemTypesOk returns a tuple with the CatalogItemTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogItemTypes

`func (o *ListCatalogTypes200Response) SetCatalogItemTypes(v []ListCatalogTypes200ResponseAllOfCatalogItemTypesInner)`

SetCatalogItemTypes sets CatalogItemTypes field to given value.

### HasCatalogItemTypes

`func (o *ListCatalogTypes200Response) HasCatalogItemTypes() bool`

HasCatalogItemTypes returns a boolean if a field has been set.

### GetMeta

`func (o *ListCatalogTypes200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListCatalogTypes200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListCatalogTypes200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListCatalogTypes200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


