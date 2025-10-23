# ListStorageServers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageServers** | Pointer to [**[]ListStorageServers200ResponseAllOfStorageServersInner**](ListStorageServers200ResponseAllOfStorageServersInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListStorageServers200Response

`func NewListStorageServers200Response() *ListStorageServers200Response`

NewListStorageServers200Response instantiates a new ListStorageServers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListStorageServers200ResponseWithDefaults

`func NewListStorageServers200ResponseWithDefaults() *ListStorageServers200Response`

NewListStorageServers200ResponseWithDefaults instantiates a new ListStorageServers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageServers

`func (o *ListStorageServers200Response) GetStorageServers() []ListStorageServers200ResponseAllOfStorageServersInner`

GetStorageServers returns the StorageServers field if non-nil, zero value otherwise.

### GetStorageServersOk

`func (o *ListStorageServers200Response) GetStorageServersOk() (*[]ListStorageServers200ResponseAllOfStorageServersInner, bool)`

GetStorageServersOk returns a tuple with the StorageServers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServers

`func (o *ListStorageServers200Response) SetStorageServers(v []ListStorageServers200ResponseAllOfStorageServersInner)`

SetStorageServers sets StorageServers field to given value.

### HasStorageServers

`func (o *ListStorageServers200Response) HasStorageServers() bool`

HasStorageServers returns a boolean if a field has been set.

### GetMeta

`func (o *ListStorageServers200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListStorageServers200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListStorageServers200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListStorageServers200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


