# UpdateRoleWorkflowAccessRequestOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TaskSetId** | **int32** | &#x60;id&#x60; of the workflow (taskSet) | 
**Access** | **string** | The new access level. | 

## Methods

### NewUpdateRoleWorkflowAccessRequestOneOf

`func NewUpdateRoleWorkflowAccessRequestOneOf(taskSetId int32, access string, ) *UpdateRoleWorkflowAccessRequestOneOf`

NewUpdateRoleWorkflowAccessRequestOneOf instantiates a new UpdateRoleWorkflowAccessRequestOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateRoleWorkflowAccessRequestOneOfWithDefaults

`func NewUpdateRoleWorkflowAccessRequestOneOfWithDefaults() *UpdateRoleWorkflowAccessRequestOneOf`

NewUpdateRoleWorkflowAccessRequestOneOfWithDefaults instantiates a new UpdateRoleWorkflowAccessRequestOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTaskSetId

`func (o *UpdateRoleWorkflowAccessRequestOneOf) GetTaskSetId() int32`

GetTaskSetId returns the TaskSetId field if non-nil, zero value otherwise.

### GetTaskSetIdOk

`func (o *UpdateRoleWorkflowAccessRequestOneOf) GetTaskSetIdOk() (*int32, bool)`

GetTaskSetIdOk returns a tuple with the TaskSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSetId

`func (o *UpdateRoleWorkflowAccessRequestOneOf) SetTaskSetId(v int32)`

SetTaskSetId sets TaskSetId field to given value.


### GetAccess

`func (o *UpdateRoleWorkflowAccessRequestOneOf) GetAccess() string`

GetAccess returns the Access field if non-nil, zero value otherwise.

### GetAccessOk

`func (o *UpdateRoleWorkflowAccessRequestOneOf) GetAccessOk() (*string, bool)`

GetAccessOk returns a tuple with the Access field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccess

`func (o *UpdateRoleWorkflowAccessRequestOneOf) SetAccess(v string)`

SetAccess sets Access field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


