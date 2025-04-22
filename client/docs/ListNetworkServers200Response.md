# ListNetworkServers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NetworkServers** | Pointer to [**[]ListNetworkServers200ResponseAllOfNetworkServersInner**](ListNetworkServers200ResponseAllOfNetworkServersInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListNetworkServers200Response

`func NewListNetworkServers200Response() *ListNetworkServers200Response`

NewListNetworkServers200Response instantiates a new ListNetworkServers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNetworkServers200ResponseWithDefaults

`func NewListNetworkServers200ResponseWithDefaults() *ListNetworkServers200Response`

NewListNetworkServers200ResponseWithDefaults instantiates a new ListNetworkServers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetworkServers

`func (o *ListNetworkServers200Response) GetNetworkServers() []ListNetworkServers200ResponseAllOfNetworkServersInner`

GetNetworkServers returns the NetworkServers field if non-nil, zero value otherwise.

### GetNetworkServersOk

`func (o *ListNetworkServers200Response) GetNetworkServersOk() (*[]ListNetworkServers200ResponseAllOfNetworkServersInner, bool)`

GetNetworkServersOk returns a tuple with the NetworkServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkServers

`func (o *ListNetworkServers200Response) SetNetworkServers(v []ListNetworkServers200ResponseAllOfNetworkServersInner)`

SetNetworkServers sets NetworkServers field to given value.

### HasNetworkServers

`func (o *ListNetworkServers200Response) HasNetworkServers() bool`

HasNetworkServers returns a boolean if a field has been set.

### GetMeta

`func (o *ListNetworkServers200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListNetworkServers200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListNetworkServers200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListNetworkServers200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


