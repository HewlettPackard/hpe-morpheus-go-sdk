# UpdateInstanceSchedule200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceSchedule** | Pointer to [**UpdateInstanceSchedule200ResponseAllOfInstanceSchedule**](UpdateInstanceSchedule200ResponseAllOfInstanceSchedule.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateInstanceSchedule200Response

`func NewUpdateInstanceSchedule200Response() *UpdateInstanceSchedule200Response`

NewUpdateInstanceSchedule200Response instantiates a new UpdateInstanceSchedule200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInstanceSchedule200ResponseWithDefaults

`func NewUpdateInstanceSchedule200ResponseWithDefaults() *UpdateInstanceSchedule200Response`

NewUpdateInstanceSchedule200ResponseWithDefaults instantiates a new UpdateInstanceSchedule200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceSchedule

`func (o *UpdateInstanceSchedule200Response) GetInstanceSchedule() UpdateInstanceSchedule200ResponseAllOfInstanceSchedule`

GetInstanceSchedule returns the InstanceSchedule field if non-nil, zero value otherwise.

### GetInstanceScheduleOk

`func (o *UpdateInstanceSchedule200Response) GetInstanceScheduleOk() (*UpdateInstanceSchedule200ResponseAllOfInstanceSchedule, bool)`

GetInstanceScheduleOk returns a tuple with the InstanceSchedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSchedule

`func (o *UpdateInstanceSchedule200Response) SetInstanceSchedule(v UpdateInstanceSchedule200ResponseAllOfInstanceSchedule)`

SetInstanceSchedule sets InstanceSchedule field to given value.

### HasInstanceSchedule

`func (o *UpdateInstanceSchedule200Response) HasInstanceSchedule() bool`

HasInstanceSchedule returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateInstanceSchedule200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateInstanceSchedule200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateInstanceSchedule200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateInstanceSchedule200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


