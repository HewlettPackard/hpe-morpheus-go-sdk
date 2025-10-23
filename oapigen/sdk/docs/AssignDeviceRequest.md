# AssignDeviceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TargetServerId** | **int64** | Target Server (VM) ID | 

## Methods

### NewAssignDeviceRequest

`func NewAssignDeviceRequest(targetServerId int64, ) *AssignDeviceRequest`

NewAssignDeviceRequest instantiates a new AssignDeviceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssignDeviceRequestWithDefaults

`func NewAssignDeviceRequestWithDefaults() *AssignDeviceRequest`

NewAssignDeviceRequestWithDefaults instantiates a new AssignDeviceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTargetServerId

`func (o *AssignDeviceRequest) GetTargetServerId() int64`

GetTargetServerId returns the TargetServerId field if non-nil, zero value otherwise.

### GetTargetServerIdOk

`func (o *AssignDeviceRequest) GetTargetServerIdOk() (*int64, bool)`

GetTargetServerIdOk returns a tuple with the TargetServerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetServerId

`func (o *AssignDeviceRequest) SetTargetServerId(v int64)`

SetTargetServerId sets TargetServerId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


