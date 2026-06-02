# GetCatalogType200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CatalogItemTypes** | Pointer to [**[]GetCatalogType200ResponseAllOfCatalogItemTypesInner**](GetCatalogType200ResponseAllOfCatalogItemTypesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListApprovals200ResponseAllOfMeta**](ListApprovals200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewGetCatalogType200Response

`func NewGetCatalogType200Response() *GetCatalogType200Response`

NewGetCatalogType200Response instantiates a new GetCatalogType200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetCatalogItemTypes

`func (o *GetCatalogType200Response) GetCatalogItemTypes() []GetCatalogType200ResponseAllOfCatalogItemTypesInner`

GetCatalogItemTypes returns the CatalogItemTypes field if non-nil, zero value otherwise.

### GetCatalogItemTypesOk

`func (o *GetCatalogType200Response) GetCatalogItemTypesOk() (*[]GetCatalogType200ResponseAllOfCatalogItemTypesInner, bool)`

GetCatalogItemTypesOk returns a tuple with the CatalogItemTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCatalogItemTypes

`func (o *GetCatalogType200Response) SetCatalogItemTypes(v []GetCatalogType200ResponseAllOfCatalogItemTypesInner)`

SetCatalogItemTypes sets CatalogItemTypes field to given value.

### HasCatalogItemTypes

`func (o *GetCatalogType200Response) HasCatalogItemTypes() bool`

HasCatalogItemTypes returns a boolean if a field has been set.

### GetMeta

`func (o *GetCatalogType200Response) GetMeta() ListApprovals200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetCatalogType200Response) GetMetaOk() (*ListApprovals200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetCatalogType200Response) SetMeta(v ListApprovals200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetCatalogType200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


