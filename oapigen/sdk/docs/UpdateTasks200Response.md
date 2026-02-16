# UpdateTasks200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Task** | Pointer to [**UpdateTasks200ResponseAllOfTask**](UpdateTasks200ResponseAllOfTask.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateTasks200Response

`func NewUpdateTasks200Response() *UpdateTasks200Response`

NewUpdateTasks200Response instantiates a new UpdateTasks200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateTasks200ResponseWithDefaults

`func NewUpdateTasks200ResponseWithDefaults() *UpdateTasks200Response`

NewUpdateTasks200ResponseWithDefaults instantiates a new UpdateTasks200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTask

`func (o *UpdateTasks200Response) GetTask() UpdateTasks200ResponseAllOfTask`

GetTask returns the Task field if non-nil, zero value otherwise.

### GetTaskOk

`func (o *UpdateTasks200Response) GetTaskOk() (*UpdateTasks200ResponseAllOfTask, bool)`

GetTaskOk returns a tuple with the Task field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTask

`func (o *UpdateTasks200Response) SetTask(v UpdateTasks200ResponseAllOfTask)`

SetTask sets Task field to given value.

### HasTask

`func (o *UpdateTasks200Response) HasTask() bool`

HasTask returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateTasks200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateTasks200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateTasks200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateTasks200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


