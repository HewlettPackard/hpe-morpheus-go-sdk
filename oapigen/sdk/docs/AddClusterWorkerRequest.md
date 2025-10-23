# AddClusterWorkerRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server** | Pointer to [**AddClusterRequestClusterServer**](AddClusterRequestClusterServer.md) |  | [optional] 

## Methods

### NewAddClusterWorkerRequest

`func NewAddClusterWorkerRequest() *AddClusterWorkerRequest`

NewAddClusterWorkerRequest instantiates a new AddClusterWorkerRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddClusterWorkerRequestWithDefaults

`func NewAddClusterWorkerRequestWithDefaults() *AddClusterWorkerRequest`

NewAddClusterWorkerRequestWithDefaults instantiates a new AddClusterWorkerRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer

`func (o *AddClusterWorkerRequest) GetServer() AddClusterRequestClusterServer`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *AddClusterWorkerRequest) GetServerOk() (*AddClusterRequestClusterServer, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *AddClusterWorkerRequest) SetServer(v AddClusterRequestClusterServer)`

SetServer sets Server field to given value.

### HasServer

`func (o *AddClusterWorkerRequest) HasServer() bool`

HasServer returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


