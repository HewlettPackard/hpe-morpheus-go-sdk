# ListVDIPools200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VdiPools** | Pointer to [**[]ListVDIPools200ResponseAllOfVdiPoolsInner**](ListVDIPools200ResponseAllOfVdiPoolsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListVDIPools200Response

`func NewListVDIPools200Response() *ListVDIPools200Response`

NewListVDIPools200Response instantiates a new ListVDIPools200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListVDIPools200ResponseWithDefaults

`func NewListVDIPools200ResponseWithDefaults() *ListVDIPools200Response`

NewListVDIPools200ResponseWithDefaults instantiates a new ListVDIPools200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVdiPools

`func (o *ListVDIPools200Response) GetVdiPools() []ListVDIPools200ResponseAllOfVdiPoolsInner`

GetVdiPools returns the VdiPools field if non-nil, zero value otherwise.

### GetVdiPoolsOk

`func (o *ListVDIPools200Response) GetVdiPoolsOk() (*[]ListVDIPools200ResponseAllOfVdiPoolsInner, bool)`

GetVdiPoolsOk returns a tuple with the VdiPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdiPools

`func (o *ListVDIPools200Response) SetVdiPools(v []ListVDIPools200ResponseAllOfVdiPoolsInner)`

SetVdiPools sets VdiPools field to given value.

### HasVdiPools

`func (o *ListVDIPools200Response) HasVdiPools() bool`

HasVdiPools returns a boolean if a field has been set.

### GetMeta

`func (o *ListVDIPools200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListVDIPools200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListVDIPools200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListVDIPools200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


