# ListNetworkPoolServers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkPoolServers** | Pointer to [**[]ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner**](ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListNetworkPoolServers200Response

`func NewListNetworkPoolServers200Response() *ListNetworkPoolServers200Response`

NewListNetworkPoolServers200Response instantiates a new ListNetworkPoolServers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNetworkPoolServers200ResponseWithDefaults

`func NewListNetworkPoolServers200ResponseWithDefaults() *ListNetworkPoolServers200Response`

NewListNetworkPoolServers200ResponseWithDefaults instantiates a new ListNetworkPoolServers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkPoolServers

`func (o *ListNetworkPoolServers200Response) GetNetworkPoolServers() []ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner`

GetNetworkPoolServers returns the NetworkPoolServers field if non-nil, zero value otherwise.

### GetNetworkPoolServersOk

`func (o *ListNetworkPoolServers200Response) GetNetworkPoolServersOk() (*[]ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner, bool)`

GetNetworkPoolServersOk returns a tuple with the NetworkPoolServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkPoolServers

`func (o *ListNetworkPoolServers200Response) SetNetworkPoolServers(v []ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner)`

SetNetworkPoolServers sets NetworkPoolServers field to given value.

### HasNetworkPoolServers

`func (o *ListNetworkPoolServers200Response) HasNetworkPoolServers() bool`

HasNetworkPoolServers returns a boolean if a field has been set.

### GetMeta

`func (o *ListNetworkPoolServers200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListNetworkPoolServers200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListNetworkPoolServers200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListNetworkPoolServers200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


