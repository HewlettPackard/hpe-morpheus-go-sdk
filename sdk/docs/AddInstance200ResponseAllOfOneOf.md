# AddInstance200ResponseAllOfOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instance** | [**AddInstance200ResponseAllOfOneOfInstance**](AddInstance200ResponseAllOfOneOfInstance.md) |  | 
**ZoneId** | **int64** | The Cloud ID to provision the instance onto. | 

## Methods

### NewAddInstance200ResponseAllOfOneOf

`func NewAddInstance200ResponseAllOfOneOf(instance AddInstance200ResponseAllOfOneOfInstance, zoneId int64, ) *AddInstance200ResponseAllOfOneOf`

NewAddInstance200ResponseAllOfOneOf instantiates a new AddInstance200ResponseAllOfOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddInstance200ResponseAllOfOneOfWithDefaults

`func NewAddInstance200ResponseAllOfOneOfWithDefaults() *AddInstance200ResponseAllOfOneOf`

NewAddInstance200ResponseAllOfOneOfWithDefaults instantiates a new AddInstance200ResponseAllOfOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstance

`func (o *AddInstance200ResponseAllOfOneOf) GetInstance() AddInstance200ResponseAllOfOneOfInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *AddInstance200ResponseAllOfOneOf) GetInstanceOk() (*AddInstance200ResponseAllOfOneOfInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *AddInstance200ResponseAllOfOneOf) SetInstance(v AddInstance200ResponseAllOfOneOfInstance)`

SetInstance sets Instance field to given value.


### GetZoneId

`func (o *AddInstance200ResponseAllOfOneOf) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *AddInstance200ResponseAllOfOneOf) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *AddInstance200ResponseAllOfOneOf) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


