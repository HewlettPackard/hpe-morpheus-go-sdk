# UpdateExecuteSchedules200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedule** | Pointer to [**UpdateExecuteSchedules200ResponseAllOfSchedule**](UpdateExecuteSchedules200ResponseAllOfSchedule.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateExecuteSchedules200Response

`func NewUpdateExecuteSchedules200Response() *UpdateExecuteSchedules200Response`

NewUpdateExecuteSchedules200Response instantiates a new UpdateExecuteSchedules200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetSchedule

`func (o *UpdateExecuteSchedules200Response) GetSchedule() UpdateExecuteSchedules200ResponseAllOfSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *UpdateExecuteSchedules200Response) GetScheduleOk() (*UpdateExecuteSchedules200ResponseAllOfSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *UpdateExecuteSchedules200Response) SetSchedule(v UpdateExecuteSchedules200ResponseAllOfSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *UpdateExecuteSchedules200Response) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateExecuteSchedules200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateExecuteSchedules200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateExecuteSchedules200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateExecuteSchedules200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


