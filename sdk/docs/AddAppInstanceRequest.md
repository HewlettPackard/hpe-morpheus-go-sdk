# AddAppInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceId** | **int64** | The ID of the instance being added | 
**TierName** | **string** | The Name of the Tier | 

## Methods

### NewAddAppInstanceRequest

`func NewAddAppInstanceRequest(instanceId int64, tierName string, ) *AddAppInstanceRequest`

NewAddAppInstanceRequest instantiates a new AddAppInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddAppInstanceRequestWithDefaults

`func NewAddAppInstanceRequestWithDefaults() *AddAppInstanceRequest`

NewAddAppInstanceRequestWithDefaults instantiates a new AddAppInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceId

`func (o *AddAppInstanceRequest) GetInstanceId() int64`

GetInstanceId returns the InstanceId field if non-nil, zero value otherwise.

### GetInstanceIdOk

`func (o *AddAppInstanceRequest) GetInstanceIdOk() (*int64, bool)`

GetInstanceIdOk returns a tuple with the InstanceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceId

`func (o *AddAppInstanceRequest) SetInstanceId(v int64)`

SetInstanceId sets InstanceId field to given value.


### GetTierName

`func (o *AddAppInstanceRequest) GetTierName() string`

GetTierName returns the TierName field if non-nil, zero value otherwise.

### GetTierNameOk

`func (o *AddAppInstanceRequest) GetTierNameOk() (*string, bool)`

GetTierNameOk returns a tuple with the TierName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTierName

`func (o *AddAppInstanceRequest) SetTierName(v string)`

SetTierName sets TierName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


