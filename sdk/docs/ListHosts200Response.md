# ListHosts200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Servers** | Pointer to [**[]ListHosts200ResponseAllOfServersInner**](ListHosts200ResponseAllOfServersInner.md) |  | [optional] 
**Stats** | Pointer to **map[string]interface{}** |  | [optional] 
**Multitenant** | Pointer to **bool** |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListHosts200Response

`func NewListHosts200Response() *ListHosts200Response`

NewListHosts200Response instantiates a new ListHosts200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListHosts200ResponseWithDefaults

`func NewListHosts200ResponseWithDefaults() *ListHosts200Response`

NewListHosts200ResponseWithDefaults instantiates a new ListHosts200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServers

`func (o *ListHosts200Response) GetServers() []ListHosts200ResponseAllOfServersInner`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *ListHosts200Response) GetServersOk() (*[]ListHosts200ResponseAllOfServersInner, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *ListHosts200Response) SetServers(v []ListHosts200ResponseAllOfServersInner)`

SetServers sets Servers field to given value.

### HasServers

`func (o *ListHosts200Response) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetStats

`func (o *ListHosts200Response) GetStats() map[string]interface{}`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ListHosts200Response) GetStatsOk() (*map[string]interface{}, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ListHosts200Response) SetStats(v map[string]interface{})`

SetStats sets Stats field to given value.

### HasStats

`func (o *ListHosts200Response) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetMultitenant

`func (o *ListHosts200Response) GetMultitenant() bool`

GetMultitenant returns the Multitenant field if non-nil, zero value otherwise.

### GetMultitenantOk

`func (o *ListHosts200Response) GetMultitenantOk() (*bool, bool)`

GetMultitenantOk returns a tuple with the Multitenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMultitenant

`func (o *ListHosts200Response) SetMultitenant(v bool)`

SetMultitenant sets Multitenant field to given value.

### HasMultitenant

`func (o *ListHosts200Response) HasMultitenant() bool`

HasMultitenant returns a boolean if a field has been set.

### GetMeta

`func (o *ListHosts200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListHosts200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListHosts200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListHosts200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


