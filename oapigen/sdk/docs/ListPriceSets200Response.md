# ListPriceSets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PriceSets** | Pointer to [**[]ListPriceSets200ResponseAllOfPriceSetsInner**](ListPriceSets200ResponseAllOfPriceSetsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListPriceSets200Response

`func NewListPriceSets200Response() *ListPriceSets200Response`

NewListPriceSets200Response instantiates a new ListPriceSets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPriceSets200ResponseWithDefaults

`func NewListPriceSets200ResponseWithDefaults() *ListPriceSets200Response`

NewListPriceSets200ResponseWithDefaults instantiates a new ListPriceSets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPriceSets

`func (o *ListPriceSets200Response) GetPriceSets() []ListPriceSets200ResponseAllOfPriceSetsInner`

GetPriceSets returns the PriceSets field if non-nil, zero value otherwise.

### GetPriceSetsOk

`func (o *ListPriceSets200Response) GetPriceSetsOk() (*[]ListPriceSets200ResponseAllOfPriceSetsInner, bool)`

GetPriceSetsOk returns a tuple with the PriceSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSets

`func (o *ListPriceSets200Response) SetPriceSets(v []ListPriceSets200ResponseAllOfPriceSetsInner)`

SetPriceSets sets PriceSets field to given value.

### HasPriceSets

`func (o *ListPriceSets200Response) HasPriceSets() bool`

HasPriceSets returns a boolean if a field has been set.

### GetMeta

`func (o *ListPriceSets200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListPriceSets200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListPriceSets200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListPriceSets200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


