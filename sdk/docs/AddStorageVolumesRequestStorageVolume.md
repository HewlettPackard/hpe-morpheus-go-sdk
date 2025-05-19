# AddStorageVolumesRequestStorageVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A unique name scoped to your account for the storage volume | 
**Type** | **string** | Storage Type Code or ID | 
**Config** | Pointer to **map[string]interface{}** | Configuration object with parameters that vary by &#x60;type&#x60;. | [optional] 
**StorageServer** | [**AddClusterLayoutsRequestLayoutMastersInnerContainerType**](AddClusterLayoutsRequestLayoutMastersInnerContainerType.md) |  | 
**StorageGroup** | [**AddClusterLayoutsRequestLayoutMastersInnerContainerType**](AddClusterLayoutsRequestLayoutMastersInnerContainerType.md) |  | 

## Methods

### NewAddStorageVolumesRequestStorageVolume

`func NewAddStorageVolumesRequestStorageVolume(name string, type_ string, storageServer AddClusterLayoutsRequestLayoutMastersInnerContainerType, storageGroup AddClusterLayoutsRequestLayoutMastersInnerContainerType, ) *AddStorageVolumesRequestStorageVolume`

NewAddStorageVolumesRequestStorageVolume instantiates a new AddStorageVolumesRequestStorageVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddStorageVolumesRequestStorageVolumeWithDefaults

`func NewAddStorageVolumesRequestStorageVolumeWithDefaults() *AddStorageVolumesRequestStorageVolume`

NewAddStorageVolumesRequestStorageVolumeWithDefaults instantiates a new AddStorageVolumesRequestStorageVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddStorageVolumesRequestStorageVolume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddStorageVolumesRequestStorageVolume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddStorageVolumesRequestStorageVolume) SetName(v string)`

SetName sets Name field to given value.


### GetType

`func (o *AddStorageVolumesRequestStorageVolume) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddStorageVolumesRequestStorageVolume) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddStorageVolumesRequestStorageVolume) SetType(v string)`

SetType sets Type field to given value.


### GetConfig

`func (o *AddStorageVolumesRequestStorageVolume) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddStorageVolumesRequestStorageVolume) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddStorageVolumesRequestStorageVolume) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddStorageVolumesRequestStorageVolume) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStorageServer

`func (o *AddStorageVolumesRequestStorageVolume) GetStorageServer() AddClusterLayoutsRequestLayoutMastersInnerContainerType`

GetStorageServer returns the StorageServer field if non-nil, zero value otherwise.

### GetStorageServerOk

`func (o *AddStorageVolumesRequestStorageVolume) GetStorageServerOk() (*AddClusterLayoutsRequestLayoutMastersInnerContainerType, bool)`

GetStorageServerOk returns a tuple with the StorageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServer

`func (o *AddStorageVolumesRequestStorageVolume) SetStorageServer(v AddClusterLayoutsRequestLayoutMastersInnerContainerType)`

SetStorageServer sets StorageServer field to given value.


### GetStorageGroup

`func (o *AddStorageVolumesRequestStorageVolume) GetStorageGroup() AddClusterLayoutsRequestLayoutMastersInnerContainerType`

GetStorageGroup returns the StorageGroup field if non-nil, zero value otherwise.

### GetStorageGroupOk

`func (o *AddStorageVolumesRequestStorageVolume) GetStorageGroupOk() (*AddClusterLayoutsRequestLayoutMastersInnerContainerType, bool)`

GetStorageGroupOk returns a tuple with the StorageGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageGroup

`func (o *AddStorageVolumesRequestStorageVolume) SetStorageGroup(v AddClusterLayoutsRequestLayoutMastersInnerContainerType)`

SetStorageGroup sets StorageGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


