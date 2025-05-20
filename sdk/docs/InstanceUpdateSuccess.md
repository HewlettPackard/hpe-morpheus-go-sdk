# InstanceUpdateSuccess

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instance** | [**UpdateInstance200ResponseAllOfOneOfInstance**](UpdateInstance200ResponseAllOfOneOfInstance.md) |  | 
**ZoneId** | **int64** | The Cloud ID to provision the instance onto. | 

## Methods

### NewInstanceUpdateSuccess

`func NewInstanceUpdateSuccess(instance UpdateInstance200ResponseAllOfOneOfInstance, zoneId int64, ) *InstanceUpdateSuccess`

NewInstanceUpdateSuccess instantiates a new InstanceUpdateSuccess object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceUpdateSuccessWithDefaults

`func NewInstanceUpdateSuccessWithDefaults() *InstanceUpdateSuccess`

NewInstanceUpdateSuccessWithDefaults instantiates a new InstanceUpdateSuccess object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstance

`func (o *InstanceUpdateSuccess) GetInstance() UpdateInstance200ResponseAllOfOneOfInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *InstanceUpdateSuccess) GetInstanceOk() (*UpdateInstance200ResponseAllOfOneOfInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *InstanceUpdateSuccess) SetInstance(v UpdateInstance200ResponseAllOfOneOfInstance)`

SetInstance sets Instance field to given value.


### GetZoneId

`func (o *InstanceUpdateSuccess) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *InstanceUpdateSuccess) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *InstanceUpdateSuccess) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


