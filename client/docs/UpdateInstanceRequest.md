# UpdateInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instance** | Pointer to [**UpdateInstanceRequestInstance**](UpdateInstanceRequestInstance.md) |  | [optional] 
**Config** | Pointer to [**UpdateInstanceRequestConfig**](UpdateInstanceRequestConfig.md) |  | [optional] 

## Methods

### NewUpdateInstanceRequest

`func NewUpdateInstanceRequest() *UpdateInstanceRequest`

NewUpdateInstanceRequest instantiates a new UpdateInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInstanceRequestWithDefaults

`func NewUpdateInstanceRequestWithDefaults() *UpdateInstanceRequest`

NewUpdateInstanceRequestWithDefaults instantiates a new UpdateInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstance

`func (o *UpdateInstanceRequest) GetInstance() UpdateInstanceRequestInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *UpdateInstanceRequest) GetInstanceOk() (*UpdateInstanceRequestInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *UpdateInstanceRequest) SetInstance(v UpdateInstanceRequestInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *UpdateInstanceRequest) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateInstanceRequest) GetConfig() UpdateInstanceRequestConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateInstanceRequest) GetConfigOk() (*UpdateInstanceRequestConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateInstanceRequest) SetConfig(v UpdateInstanceRequestConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateInstanceRequest) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


