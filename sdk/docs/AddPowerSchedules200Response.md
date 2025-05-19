# AddPowerSchedules200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Schedule** | Pointer to [**AddPowerSchedules200ResponseAllOfSchedule**](AddPowerSchedules200ResponseAllOfSchedule.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddPowerSchedules200Response

`func NewAddPowerSchedules200Response() *AddPowerSchedules200Response`

NewAddPowerSchedules200Response instantiates a new AddPowerSchedules200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPowerSchedules200ResponseWithDefaults

`func NewAddPowerSchedules200ResponseWithDefaults() *AddPowerSchedules200Response`

NewAddPowerSchedules200ResponseWithDefaults instantiates a new AddPowerSchedules200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSchedule

`func (o *AddPowerSchedules200Response) GetSchedule() AddPowerSchedules200ResponseAllOfSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *AddPowerSchedules200Response) GetScheduleOk() (*AddPowerSchedules200ResponseAllOfSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *AddPowerSchedules200Response) SetSchedule(v AddPowerSchedules200ResponseAllOfSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *AddPowerSchedules200Response) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetSuccess

`func (o *AddPowerSchedules200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddPowerSchedules200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddPowerSchedules200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddPowerSchedules200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


