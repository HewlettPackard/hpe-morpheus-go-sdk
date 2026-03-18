# UpdatePowerSchedules200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedule** | Pointer to [**UpdatePowerSchedules200ResponseAllOfSchedule**](UpdatePowerSchedules200ResponseAllOfSchedule.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdatePowerSchedules200Response

`func NewUpdatePowerSchedules200Response() *UpdatePowerSchedules200Response`

NewUpdatePowerSchedules200Response instantiates a new UpdatePowerSchedules200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePowerSchedules200ResponseWithDefaults

`func NewUpdatePowerSchedules200ResponseWithDefaults() *UpdatePowerSchedules200Response`

NewUpdatePowerSchedules200ResponseWithDefaults instantiates a new UpdatePowerSchedules200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedule

`func (o *UpdatePowerSchedules200Response) GetSchedule() UpdatePowerSchedules200ResponseAllOfSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *UpdatePowerSchedules200Response) GetScheduleOk() (*UpdatePowerSchedules200ResponseAllOfSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *UpdatePowerSchedules200Response) SetSchedule(v UpdatePowerSchedules200ResponseAllOfSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *UpdatePowerSchedules200Response) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdatePowerSchedules200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdatePowerSchedules200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdatePowerSchedules200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdatePowerSchedules200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


