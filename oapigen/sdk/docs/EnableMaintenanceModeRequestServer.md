# EnableMaintenanceModeRequestServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IgnoreDaemonsets** | Pointer to **bool** | option relevant to kubernetes nodes | [optional] 
**Force** | Pointer to **bool** | option relevant to kubernetes nodes | [optional] 
**DeleteEmptyDir** | Pointer to **bool** | option relevant to kubernetes nodes | [optional] 
**DeleteLocalData** | Pointer to **bool** | option relevant to kubernetes nodes | [optional] 

## Methods

### NewEnableMaintenanceModeRequestServer

`func NewEnableMaintenanceModeRequestServer() *EnableMaintenanceModeRequestServer`

NewEnableMaintenanceModeRequestServer instantiates a new EnableMaintenanceModeRequestServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEnableMaintenanceModeRequestServerWithDefaults

`func NewEnableMaintenanceModeRequestServerWithDefaults() *EnableMaintenanceModeRequestServer`

NewEnableMaintenanceModeRequestServerWithDefaults instantiates a new EnableMaintenanceModeRequestServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIgnoreDaemonsets

`func (o *EnableMaintenanceModeRequestServer) GetIgnoreDaemonsets() bool`

GetIgnoreDaemonsets returns the IgnoreDaemonsets field if non-nil, zero value otherwise.

### GetIgnoreDaemonsetsOk

`func (o *EnableMaintenanceModeRequestServer) GetIgnoreDaemonsetsOk() (*bool, bool)`

GetIgnoreDaemonsetsOk returns a tuple with the IgnoreDaemonsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreDaemonsets

`func (o *EnableMaintenanceModeRequestServer) SetIgnoreDaemonsets(v bool)`

SetIgnoreDaemonsets sets IgnoreDaemonsets field to given value.

### HasIgnoreDaemonsets

`func (o *EnableMaintenanceModeRequestServer) HasIgnoreDaemonsets() bool`

HasIgnoreDaemonsets returns a boolean if a field has been set.

### GetForce

`func (o *EnableMaintenanceModeRequestServer) GetForce() bool`

GetForce returns the Force field if non-nil, zero value otherwise.

### GetForceOk

`func (o *EnableMaintenanceModeRequestServer) GetForceOk() (*bool, bool)`

GetForceOk returns a tuple with the Force field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForce

`func (o *EnableMaintenanceModeRequestServer) SetForce(v bool)`

SetForce sets Force field to given value.

### HasForce

`func (o *EnableMaintenanceModeRequestServer) HasForce() bool`

HasForce returns a boolean if a field has been set.

### GetDeleteEmptyDir

`func (o *EnableMaintenanceModeRequestServer) GetDeleteEmptyDir() bool`

GetDeleteEmptyDir returns the DeleteEmptyDir field if non-nil, zero value otherwise.

### GetDeleteEmptyDirOk

`func (o *EnableMaintenanceModeRequestServer) GetDeleteEmptyDirOk() (*bool, bool)`

GetDeleteEmptyDirOk returns a tuple with the DeleteEmptyDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteEmptyDir

`func (o *EnableMaintenanceModeRequestServer) SetDeleteEmptyDir(v bool)`

SetDeleteEmptyDir sets DeleteEmptyDir field to given value.

### HasDeleteEmptyDir

`func (o *EnableMaintenanceModeRequestServer) HasDeleteEmptyDir() bool`

HasDeleteEmptyDir returns a boolean if a field has been set.

### GetDeleteLocalData

`func (o *EnableMaintenanceModeRequestServer) GetDeleteLocalData() bool`

GetDeleteLocalData returns the DeleteLocalData field if non-nil, zero value otherwise.

### GetDeleteLocalDataOk

`func (o *EnableMaintenanceModeRequestServer) GetDeleteLocalDataOk() (*bool, bool)`

GetDeleteLocalDataOk returns a tuple with the DeleteLocalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleteLocalData

`func (o *EnableMaintenanceModeRequestServer) SetDeleteLocalData(v bool)`

SetDeleteLocalData sets DeleteLocalData field to given value.

### HasDeleteLocalData

`func (o *EnableMaintenanceModeRequestServer) HasDeleteLocalData() bool`

HasDeleteLocalData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


