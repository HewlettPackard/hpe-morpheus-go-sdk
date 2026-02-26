# InstanceConfigObjectGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**InstanceConfigObjectGroupId**](InstanceConfigObjectGroupId.md) |  | 
**Name** | Pointer to **string** | The group name | [optional] 

## Methods

### NewInstanceConfigObjectGroup

`func NewInstanceConfigObjectGroup(id InstanceConfigObjectGroupId, ) *InstanceConfigObjectGroup`

NewInstanceConfigObjectGroup instantiates a new InstanceConfigObjectGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceConfigObjectGroupWithDefaults

`func NewInstanceConfigObjectGroupWithDefaults() *InstanceConfigObjectGroup`

NewInstanceConfigObjectGroupWithDefaults instantiates a new InstanceConfigObjectGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceConfigObjectGroup) GetId() InstanceConfigObjectGroupId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceConfigObjectGroup) GetIdOk() (*InstanceConfigObjectGroupId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceConfigObjectGroup) SetId(v InstanceConfigObjectGroupId)`

SetId sets Id field to given value.


### GetName

`func (o *InstanceConfigObjectGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceConfigObjectGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceConfigObjectGroup) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceConfigObjectGroup) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


