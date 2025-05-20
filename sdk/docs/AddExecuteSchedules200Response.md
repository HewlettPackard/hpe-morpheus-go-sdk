# AddExecuteSchedules200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedule** | Pointer to [**ListExecuteSchedules200ResponseAllOfSchedulesInner**](ListExecuteSchedules200ResponseAllOfSchedulesInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddExecuteSchedules200Response

`func NewAddExecuteSchedules200Response() *AddExecuteSchedules200Response`

NewAddExecuteSchedules200Response instantiates a new AddExecuteSchedules200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddExecuteSchedules200ResponseWithDefaults

`func NewAddExecuteSchedules200ResponseWithDefaults() *AddExecuteSchedules200Response`

NewAddExecuteSchedules200ResponseWithDefaults instantiates a new AddExecuteSchedules200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedule

`func (o *AddExecuteSchedules200Response) GetSchedule() ListExecuteSchedules200ResponseAllOfSchedulesInner`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *AddExecuteSchedules200Response) GetScheduleOk() (*ListExecuteSchedules200ResponseAllOfSchedulesInner, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *AddExecuteSchedules200Response) SetSchedule(v ListExecuteSchedules200ResponseAllOfSchedulesInner)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *AddExecuteSchedules200Response) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetSuccess

`func (o *AddExecuteSchedules200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddExecuteSchedules200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddExecuteSchedules200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddExecuteSchedules200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


