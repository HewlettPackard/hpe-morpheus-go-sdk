# UpdateInstanceThreshold200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceThreshold** | Pointer to [**UpdateInstanceThreshold200ResponseAllOfInstanceThreshold**](UpdateInstanceThreshold200ResponseAllOfInstanceThreshold.md) |  | [optional] 
**InstanceSchedules** | Pointer to [**[]UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner**](UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateInstanceThreshold200Response

`func NewUpdateInstanceThreshold200Response() *UpdateInstanceThreshold200Response`

NewUpdateInstanceThreshold200Response instantiates a new UpdateInstanceThreshold200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetInstanceThreshold

`func (o *UpdateInstanceThreshold200Response) GetInstanceThreshold() UpdateInstanceThreshold200ResponseAllOfInstanceThreshold`

GetInstanceThreshold returns the InstanceThreshold field if non-nil, zero value otherwise.

### GetInstanceThresholdOk

`func (o *UpdateInstanceThreshold200Response) GetInstanceThresholdOk() (*UpdateInstanceThreshold200ResponseAllOfInstanceThreshold, bool)`

GetInstanceThresholdOk returns a tuple with the InstanceThreshold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceThreshold

`func (o *UpdateInstanceThreshold200Response) SetInstanceThreshold(v UpdateInstanceThreshold200ResponseAllOfInstanceThreshold)`

SetInstanceThreshold sets InstanceThreshold field to given value.

### HasInstanceThreshold

`func (o *UpdateInstanceThreshold200Response) HasInstanceThreshold() bool`

HasInstanceThreshold returns a boolean if a field has been set.

### GetInstanceSchedules

`func (o *UpdateInstanceThreshold200Response) GetInstanceSchedules() []UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner`

GetInstanceSchedules returns the InstanceSchedules field if non-nil, zero value otherwise.

### GetInstanceSchedulesOk

`func (o *UpdateInstanceThreshold200Response) GetInstanceSchedulesOk() (*[]UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner, bool)`

GetInstanceSchedulesOk returns a tuple with the InstanceSchedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceSchedules

`func (o *UpdateInstanceThreshold200Response) SetInstanceSchedules(v []UpdateInstanceThreshold200ResponseAllOfInstanceSchedulesInner)`

SetInstanceSchedules sets InstanceSchedules field to given value.

### HasInstanceSchedules

`func (o *UpdateInstanceThreshold200Response) HasInstanceSchedules() bool`

HasInstanceSchedules returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateInstanceThreshold200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateInstanceThreshold200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateInstanceThreshold200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateInstanceThreshold200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


