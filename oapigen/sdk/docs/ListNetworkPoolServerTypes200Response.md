# ListNetworkPoolServerTypes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkPoolServerTypes** | Pointer to [**[]ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInner**](ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListNetworkPoolServerTypes200Response

`func NewListNetworkPoolServerTypes200Response() *ListNetworkPoolServerTypes200Response`

NewListNetworkPoolServerTypes200Response instantiates a new ListNetworkPoolServerTypes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNetworkPoolServerTypes200ResponseWithDefaults

`func NewListNetworkPoolServerTypes200ResponseWithDefaults() *ListNetworkPoolServerTypes200Response`

NewListNetworkPoolServerTypes200ResponseWithDefaults instantiates a new ListNetworkPoolServerTypes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkPoolServerTypes

`func (o *ListNetworkPoolServerTypes200Response) GetNetworkPoolServerTypes() []ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInner`

GetNetworkPoolServerTypes returns the NetworkPoolServerTypes field if non-nil, zero value otherwise.

### GetNetworkPoolServerTypesOk

`func (o *ListNetworkPoolServerTypes200Response) GetNetworkPoolServerTypesOk() (*[]ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInner, bool)`

GetNetworkPoolServerTypesOk returns a tuple with the NetworkPoolServerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPoolServerTypes

`func (o *ListNetworkPoolServerTypes200Response) SetNetworkPoolServerTypes(v []ListNetworkPoolServerTypes200ResponseAllOfNetworkPoolServerTypesInner)`

SetNetworkPoolServerTypes sets NetworkPoolServerTypes field to given value.

### HasNetworkPoolServerTypes

`func (o *ListNetworkPoolServerTypes200Response) HasNetworkPoolServerTypes() bool`

HasNetworkPoolServerTypes returns a boolean if a field has been set.

### GetMeta

`func (o *ListNetworkPoolServerTypes200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListNetworkPoolServerTypes200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListNetworkPoolServerTypes200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListNetworkPoolServerTypes200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


