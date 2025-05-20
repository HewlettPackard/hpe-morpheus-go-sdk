# ListVDIAllocations200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VdiAllocations** | Pointer to [**[]ListVDIAllocations200ResponseAllOfVdiAllocationsInner**](ListVDIAllocations200ResponseAllOfVdiAllocationsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListVDIAllocations200Response

`func NewListVDIAllocations200Response() *ListVDIAllocations200Response`

NewListVDIAllocations200Response instantiates a new ListVDIAllocations200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVDIAllocations200ResponseWithDefaults

`func NewListVDIAllocations200ResponseWithDefaults() *ListVDIAllocations200Response`

NewListVDIAllocations200ResponseWithDefaults instantiates a new ListVDIAllocations200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVdiAllocations

`func (o *ListVDIAllocations200Response) GetVdiAllocations() []ListVDIAllocations200ResponseAllOfVdiAllocationsInner`

GetVdiAllocations returns the VdiAllocations field if non-nil, zero value otherwise.

### GetVdiAllocationsOk

`func (o *ListVDIAllocations200Response) GetVdiAllocationsOk() (*[]ListVDIAllocations200ResponseAllOfVdiAllocationsInner, bool)`

GetVdiAllocationsOk returns a tuple with the VdiAllocations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdiAllocations

`func (o *ListVDIAllocations200Response) SetVdiAllocations(v []ListVDIAllocations200ResponseAllOfVdiAllocationsInner)`

SetVdiAllocations sets VdiAllocations field to given value.

### HasVdiAllocations

`func (o *ListVDIAllocations200Response) HasVdiAllocations() bool`

HasVdiAllocations returns a boolean if a field has been set.

### GetMeta

`func (o *ListVDIAllocations200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListVDIAllocations200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListVDIAllocations200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListVDIAllocations200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


