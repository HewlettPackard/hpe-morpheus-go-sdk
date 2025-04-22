# UpdateHostManagedRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server** | Pointer to [**UpdateHostManagedRequestServer**](UpdateHostManagedRequestServer.md) |  | [optional] 
**InstallAgent** | Pointer to **bool** | Install agent. Set to false to manually install agent instead. | [optional] [default to true]
**InstanceTypeId** | Pointer to **int64** | Instance Type ID for the new Instance | [optional] 
**Layout** | Pointer to **int64** | Layout ID for the new Instance | [optional] 

## Methods

### NewUpdateHostManagedRequest

`func NewUpdateHostManagedRequest() *UpdateHostManagedRequest`

NewUpdateHostManagedRequest instantiates a new UpdateHostManagedRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateHostManagedRequestWithDefaults

`func NewUpdateHostManagedRequestWithDefaults() *UpdateHostManagedRequest`

NewUpdateHostManagedRequestWithDefaults instantiates a new UpdateHostManagedRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer

`func (o *UpdateHostManagedRequest) GetServer() UpdateHostManagedRequestServer`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *UpdateHostManagedRequest) GetServerOk() (*UpdateHostManagedRequestServer, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *UpdateHostManagedRequest) SetServer(v UpdateHostManagedRequestServer)`

SetServer sets Server field to given value.

### HasServer

`func (o *UpdateHostManagedRequest) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetInstallAgent

`func (o *UpdateHostManagedRequest) GetInstallAgent() bool`

GetInstallAgent returns the InstallAgent field if non-nil, zero value otherwise.

### GetInstallAgentOk

`func (o *UpdateHostManagedRequest) GetInstallAgentOk() (*bool, bool)`

GetInstallAgentOk returns a tuple with the InstallAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallAgent

`func (o *UpdateHostManagedRequest) SetInstallAgent(v bool)`

SetInstallAgent sets InstallAgent field to given value.

### HasInstallAgent

`func (o *UpdateHostManagedRequest) HasInstallAgent() bool`

HasInstallAgent returns a boolean if a field has been set.

### GetInstanceTypeId

`func (o *UpdateHostManagedRequest) GetInstanceTypeId() int64`

GetInstanceTypeId returns the InstanceTypeId field if non-nil, zero value otherwise.

### GetInstanceTypeIdOk

`func (o *UpdateHostManagedRequest) GetInstanceTypeIdOk() (*int64, bool)`

GetInstanceTypeIdOk returns a tuple with the InstanceTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceTypeId

`func (o *UpdateHostManagedRequest) SetInstanceTypeId(v int64)`

SetInstanceTypeId sets InstanceTypeId field to given value.

### HasInstanceTypeId

`func (o *UpdateHostManagedRequest) HasInstanceTypeId() bool`

HasInstanceTypeId returns a boolean if a field has been set.

### GetLayout

`func (o *UpdateHostManagedRequest) GetLayout() int64`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *UpdateHostManagedRequest) GetLayoutOk() (*int64, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *UpdateHostManagedRequest) SetLayout(v int64)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *UpdateHostManagedRequest) HasLayout() bool`

HasLayout returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


