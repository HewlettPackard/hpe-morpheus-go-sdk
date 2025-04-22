# GetNetworkPools200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkPools** | Pointer to **interface{}** |  | [optional] 
**NetworkPoolCount** | Pointer to **int64** |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewGetNetworkPools200Response

`func NewGetNetworkPools200Response() *GetNetworkPools200Response`

NewGetNetworkPools200Response instantiates a new GetNetworkPools200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkPools200ResponseWithDefaults

`func NewGetNetworkPools200ResponseWithDefaults() *GetNetworkPools200Response`

NewGetNetworkPools200ResponseWithDefaults instantiates a new GetNetworkPools200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkPools

`func (o *GetNetworkPools200Response) GetNetworkPools() interface{}`

GetNetworkPools returns the NetworkPools field if non-nil, zero value otherwise.

### GetNetworkPoolsOk

`func (o *GetNetworkPools200Response) GetNetworkPoolsOk() (*interface{}, bool)`

GetNetworkPoolsOk returns a tuple with the NetworkPools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPools

`func (o *GetNetworkPools200Response) SetNetworkPools(v interface{})`

SetNetworkPools sets NetworkPools field to given value.

### HasNetworkPools

`func (o *GetNetworkPools200Response) HasNetworkPools() bool`

HasNetworkPools returns a boolean if a field has been set.

### SetNetworkPoolsNil

`func (o *GetNetworkPools200Response) SetNetworkPoolsNil(b bool)`

 SetNetworkPoolsNil sets the value for NetworkPools to be an explicit nil

### UnsetNetworkPools
`func (o *GetNetworkPools200Response) UnsetNetworkPools()`

UnsetNetworkPools ensures that no value is present for NetworkPools, not even an explicit nil
### GetNetworkPoolCount

`func (o *GetNetworkPools200Response) GetNetworkPoolCount() int64`

GetNetworkPoolCount returns the NetworkPoolCount field if non-nil, zero value otherwise.

### GetNetworkPoolCountOk

`func (o *GetNetworkPools200Response) GetNetworkPoolCountOk() (*int64, bool)`

GetNetworkPoolCountOk returns a tuple with the NetworkPoolCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPoolCount

`func (o *GetNetworkPools200Response) SetNetworkPoolCount(v int64)`

SetNetworkPoolCount sets NetworkPoolCount field to given value.

### HasNetworkPoolCount

`func (o *GetNetworkPools200Response) HasNetworkPoolCount() bool`

HasNetworkPoolCount returns a boolean if a field has been set.

### GetMeta

`func (o *GetNetworkPools200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetNetworkPools200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetNetworkPools200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetNetworkPools200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


