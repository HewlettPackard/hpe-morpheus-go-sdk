# AddStorageServers200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageServer** | Pointer to [**ListStorageServers200ResponseAllOfStorageServersInner**](ListStorageServers200ResponseAllOfStorageServersInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddStorageServers200Response

`func NewAddStorageServers200Response() *AddStorageServers200Response`

NewAddStorageServers200Response instantiates a new AddStorageServers200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddStorageServers200ResponseWithDefaults

`func NewAddStorageServers200ResponseWithDefaults() *AddStorageServers200Response`

NewAddStorageServers200ResponseWithDefaults instantiates a new AddStorageServers200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageServer

`func (o *AddStorageServers200Response) GetStorageServer() ListStorageServers200ResponseAllOfStorageServersInner`

GetStorageServer returns the StorageServer field if non-nil, zero value otherwise.

### GetStorageServerOk

`func (o *AddStorageServers200Response) GetStorageServerOk() (*ListStorageServers200ResponseAllOfStorageServersInner, bool)`

GetStorageServerOk returns a tuple with the StorageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServer

`func (o *AddStorageServers200Response) SetStorageServer(v ListStorageServers200ResponseAllOfStorageServersInner)`

SetStorageServer sets StorageServer field to given value.

### HasStorageServer

`func (o *AddStorageServers200Response) HasStorageServer() bool`

HasStorageServer returns a boolean if a field has been set.

### GetSuccess

`func (o *AddStorageServers200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddStorageServers200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddStorageServers200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddStorageServers200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


