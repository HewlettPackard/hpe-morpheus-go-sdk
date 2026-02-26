# InstanceConfigObjectCloud

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | [**InstanceConfigObjectCloudId**](InstanceConfigObjectCloudId.md) |  | 
**Name** | Pointer to **string** | The cloud name | [optional] 

## Methods

### NewInstanceConfigObjectCloud

`func NewInstanceConfigObjectCloud(id InstanceConfigObjectCloudId, ) *InstanceConfigObjectCloud`

NewInstanceConfigObjectCloud instantiates a new InstanceConfigObjectCloud object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceConfigObjectCloudWithDefaults

`func NewInstanceConfigObjectCloudWithDefaults() *InstanceConfigObjectCloud`

NewInstanceConfigObjectCloudWithDefaults instantiates a new InstanceConfigObjectCloud object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceConfigObjectCloud) GetId() InstanceConfigObjectCloudId`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceConfigObjectCloud) GetIdOk() (*InstanceConfigObjectCloudId, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceConfigObjectCloud) SetId(v InstanceConfigObjectCloudId)`

SetId sets Id field to given value.


### GetName

`func (o *InstanceConfigObjectCloud) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceConfigObjectCloud) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceConfigObjectCloud) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceConfigObjectCloud) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


