# GetPowerSchedules200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instances** | Pointer to [**[]GetPowerSchedules200ResponseAllOfInstancesInner**](GetPowerSchedules200ResponseAllOfInstancesInner.md) |  | [optional] 
**Servers** | Pointer to [**[]GetPowerSchedules200ResponseAllOfServersInner**](GetPowerSchedules200ResponseAllOfServersInner.md) |  | [optional] 
**Schedule** | Pointer to [**GetPowerSchedules200ResponseAllOfSchedule**](GetPowerSchedules200ResponseAllOfSchedule.md) |  | [optional] 

## Methods

### NewGetPowerSchedules200Response

`func NewGetPowerSchedules200Response() *GetPowerSchedules200Response`

NewGetPowerSchedules200Response instantiates a new GetPowerSchedules200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetInstances

`func (o *GetPowerSchedules200Response) GetInstances() []GetPowerSchedules200ResponseAllOfInstancesInner`

GetInstances returns the Instances field if non-nil, zero value otherwise.

### GetInstancesOk

`func (o *GetPowerSchedules200Response) GetInstancesOk() (*[]GetPowerSchedules200ResponseAllOfInstancesInner, bool)`

GetInstancesOk returns a tuple with the Instances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstances

`func (o *GetPowerSchedules200Response) SetInstances(v []GetPowerSchedules200ResponseAllOfInstancesInner)`

SetInstances sets Instances field to given value.

### HasInstances

`func (o *GetPowerSchedules200Response) HasInstances() bool`

HasInstances returns a boolean if a field has been set.

### SetInstancesNil

`func (o *GetPowerSchedules200Response) SetInstancesNil(b bool)`

 SetInstancesNil sets the value for Instances to be an explicit nil

### UnsetInstances
`func (o *GetPowerSchedules200Response) UnsetInstances()`

UnsetInstances ensures that no value is present for Instances, not even an explicit nil
### GetServers

`func (o *GetPowerSchedules200Response) GetServers() []GetPowerSchedules200ResponseAllOfServersInner`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *GetPowerSchedules200Response) GetServersOk() (*[]GetPowerSchedules200ResponseAllOfServersInner, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *GetPowerSchedules200Response) SetServers(v []GetPowerSchedules200ResponseAllOfServersInner)`

SetServers sets Servers field to given value.

### HasServers

`func (o *GetPowerSchedules200Response) HasServers() bool`

HasServers returns a boolean if a field has been set.

### SetServersNil

`func (o *GetPowerSchedules200Response) SetServersNil(b bool)`

 SetServersNil sets the value for Servers to be an explicit nil

### UnsetServers
`func (o *GetPowerSchedules200Response) UnsetServers()`

UnsetServers ensures that no value is present for Servers, not even an explicit nil
### GetSchedule

`func (o *GetPowerSchedules200Response) GetSchedule() GetPowerSchedules200ResponseAllOfSchedule`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *GetPowerSchedules200Response) GetScheduleOk() (*GetPowerSchedules200ResponseAllOfSchedule, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *GetPowerSchedules200Response) SetSchedule(v GetPowerSchedules200ResponseAllOfSchedule)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *GetPowerSchedules200Response) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


