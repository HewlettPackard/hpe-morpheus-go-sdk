# CreateInstanceSchedule200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceSchedule** | Pointer to [**GetInstanceThreshold200ResponseInstanceSchedulesInner**](GetInstanceThreshold200ResponseInstanceSchedulesInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewCreateInstanceSchedule200Response

`func NewCreateInstanceSchedule200Response() *CreateInstanceSchedule200Response`

NewCreateInstanceSchedule200Response instantiates a new CreateInstanceSchedule200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateInstanceSchedule200ResponseWithDefaults

`func NewCreateInstanceSchedule200ResponseWithDefaults() *CreateInstanceSchedule200Response`

NewCreateInstanceSchedule200ResponseWithDefaults instantiates a new CreateInstanceSchedule200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceSchedule

`func (o *CreateInstanceSchedule200Response) GetInstanceSchedule() GetInstanceThreshold200ResponseInstanceSchedulesInner`

GetInstanceSchedule returns the InstanceSchedule field if non-nil, zero value otherwise.

### GetInstanceScheduleOk

`func (o *CreateInstanceSchedule200Response) GetInstanceScheduleOk() (*GetInstanceThreshold200ResponseInstanceSchedulesInner, bool)`

GetInstanceScheduleOk returns a tuple with the InstanceSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSchedule

`func (o *CreateInstanceSchedule200Response) SetInstanceSchedule(v GetInstanceThreshold200ResponseInstanceSchedulesInner)`

SetInstanceSchedule sets InstanceSchedule field to given value.

### HasInstanceSchedule

`func (o *CreateInstanceSchedule200Response) HasInstanceSchedule() bool`

HasInstanceSchedule returns a boolean if a field has been set.

### GetSuccess

`func (o *CreateInstanceSchedule200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CreateInstanceSchedule200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CreateInstanceSchedule200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CreateInstanceSchedule200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


