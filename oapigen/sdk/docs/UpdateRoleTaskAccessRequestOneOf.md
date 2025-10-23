# UpdateRoleTaskAccessRequestOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskId** | **int32** | &#x60;id&#x60; of the task | 
**Access** | **string** | The new access level. | 

## Methods

### NewUpdateRoleTaskAccessRequestOneOf

`func NewUpdateRoleTaskAccessRequestOneOf(taskId int32, access string, ) *UpdateRoleTaskAccessRequestOneOf`

NewUpdateRoleTaskAccessRequestOneOf instantiates a new UpdateRoleTaskAccessRequestOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRoleTaskAccessRequestOneOfWithDefaults

`func NewUpdateRoleTaskAccessRequestOneOfWithDefaults() *UpdateRoleTaskAccessRequestOneOf`

NewUpdateRoleTaskAccessRequestOneOfWithDefaults instantiates a new UpdateRoleTaskAccessRequestOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskId

`func (o *UpdateRoleTaskAccessRequestOneOf) GetTaskId() int32`

GetTaskId returns the TaskId field if non-nil, zero value otherwise.

### GetTaskIdOk

`func (o *UpdateRoleTaskAccessRequestOneOf) GetTaskIdOk() (*int32, bool)`

GetTaskIdOk returns a tuple with the TaskId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskId

`func (o *UpdateRoleTaskAccessRequestOneOf) SetTaskId(v int32)`

SetTaskId sets TaskId field to given value.


### GetAccess

`func (o *UpdateRoleTaskAccessRequestOneOf) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *UpdateRoleTaskAccessRequestOneOf) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *UpdateRoleTaskAccessRequestOneOf) SetAccess(v string)`

SetAccess sets Access field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


