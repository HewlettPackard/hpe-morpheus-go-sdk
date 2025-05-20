# GetPowerSchedules200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instances** | Pointer to [**[]GetPowerSchedules200ResponseAllOfInstancesInner**](GetPowerSchedules200ResponseAllOfInstancesInner.md) |  | [optional] 
**Servers** | Pointer to [**[]GetPowerSchedules200ResponseAllOfInstancesInner**](GetPowerSchedules200ResponseAllOfInstancesInner.md) |  | [optional] 
**Schedule** | Pointer to [**ListPowerSchedules200ResponseAllOfSchedulesInner**](ListPowerSchedules200ResponseAllOfSchedulesInner.md) |  | [optional] 

## Methods

### NewGetPowerSchedules200Response

`func NewGetPowerSchedules200Response() *GetPowerSchedules200Response`

NewGetPowerSchedules200Response instantiates a new GetPowerSchedules200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPowerSchedules200ResponseWithDefaults

`func NewGetPowerSchedules200ResponseWithDefaults() *GetPowerSchedules200Response`

NewGetPowerSchedules200ResponseWithDefaults instantiates a new GetPowerSchedules200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

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

### GetServers

`func (o *GetPowerSchedules200Response) GetServers() []GetPowerSchedules200ResponseAllOfInstancesInner`

GetServers returns the Servers field if non-nil, zero value otherwise.

### GetServersOk

`func (o *GetPowerSchedules200Response) GetServersOk() (*[]GetPowerSchedules200ResponseAllOfInstancesInner, bool)`

GetServersOk returns a tuple with the Servers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServers

`func (o *GetPowerSchedules200Response) SetServers(v []GetPowerSchedules200ResponseAllOfInstancesInner)`

SetServers sets Servers field to given value.

### HasServers

`func (o *GetPowerSchedules200Response) HasServers() bool`

HasServers returns a boolean if a field has been set.

### GetSchedule

`func (o *GetPowerSchedules200Response) GetSchedule() ListPowerSchedules200ResponseAllOfSchedulesInner`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *GetPowerSchedules200Response) GetScheduleOk() (*ListPowerSchedules200ResponseAllOfSchedulesInner, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *GetPowerSchedules200Response) SetSchedule(v ListPowerSchedules200ResponseAllOfSchedulesInner)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *GetPowerSchedules200Response) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


