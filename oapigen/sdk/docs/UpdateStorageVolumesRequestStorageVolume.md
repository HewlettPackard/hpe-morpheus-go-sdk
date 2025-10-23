# UpdateStorageVolumesRequestStorageVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A unique name scoped to your account for the storage volume | [optional] 
**Type** | Pointer to **string** | Storage Type Code or ID | [optional] 
**Config** | Pointer to **map[string]interface{}** | Configuration object with parameters that vary by &#x60;type&#x60;. | [optional] 
**StorageServer** | Pointer to [**AddClusterLayoutsRequestLayoutMastersInnerContainerType**](AddClusterLayoutsRequestLayoutMastersInnerContainerType.md) |  | [optional] 
**StorageGroup** | Pointer to [**AddClusterLayoutsRequestLayoutMastersInnerContainerType**](AddClusterLayoutsRequestLayoutMastersInnerContainerType.md) |  | [optional] 

## Methods

### NewUpdateStorageVolumesRequestStorageVolume

`func NewUpdateStorageVolumesRequestStorageVolume() *UpdateStorageVolumesRequestStorageVolume`

NewUpdateStorageVolumesRequestStorageVolume instantiates a new UpdateStorageVolumesRequestStorageVolume object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStorageVolumesRequestStorageVolumeWithDefaults

`func NewUpdateStorageVolumesRequestStorageVolumeWithDefaults() *UpdateStorageVolumesRequestStorageVolume`

NewUpdateStorageVolumesRequestStorageVolumeWithDefaults instantiates a new UpdateStorageVolumesRequestStorageVolume object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateStorageVolumesRequestStorageVolume) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateStorageVolumesRequestStorageVolume) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateStorageVolumesRequestStorageVolume) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateStorageVolumesRequestStorageVolume) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *UpdateStorageVolumesRequestStorageVolume) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateStorageVolumesRequestStorageVolume) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateStorageVolumesRequestStorageVolume) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateStorageVolumesRequestStorageVolume) HasType() bool`

HasType returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateStorageVolumesRequestStorageVolume) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateStorageVolumesRequestStorageVolume) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateStorageVolumesRequestStorageVolume) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateStorageVolumesRequestStorageVolume) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStorageServer

`func (o *UpdateStorageVolumesRequestStorageVolume) GetStorageServer() AddClusterLayoutsRequestLayoutMastersInnerContainerType`

GetStorageServer returns the StorageServer field if non-nil, zero value otherwise.

### GetStorageServerOk

`func (o *UpdateStorageVolumesRequestStorageVolume) GetStorageServerOk() (*AddClusterLayoutsRequestLayoutMastersInnerContainerType, bool)`

GetStorageServerOk returns a tuple with the StorageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServer

`func (o *UpdateStorageVolumesRequestStorageVolume) SetStorageServer(v AddClusterLayoutsRequestLayoutMastersInnerContainerType)`

SetStorageServer sets StorageServer field to given value.

### HasStorageServer

`func (o *UpdateStorageVolumesRequestStorageVolume) HasStorageServer() bool`

HasStorageServer returns a boolean if a field has been set.

### GetStorageGroup

`func (o *UpdateStorageVolumesRequestStorageVolume) GetStorageGroup() AddClusterLayoutsRequestLayoutMastersInnerContainerType`

GetStorageGroup returns the StorageGroup field if non-nil, zero value otherwise.

### GetStorageGroupOk

`func (o *UpdateStorageVolumesRequestStorageVolume) GetStorageGroupOk() (*AddClusterLayoutsRequestLayoutMastersInnerContainerType, bool)`

GetStorageGroupOk returns a tuple with the StorageGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageGroup

`func (o *UpdateStorageVolumesRequestStorageVolume) SetStorageGroup(v AddClusterLayoutsRequestLayoutMastersInnerContainerType)`

SetStorageGroup sets StorageGroup field to given value.

### HasStorageGroup

`func (o *UpdateStorageVolumesRequestStorageVolume) HasStorageGroup() bool`

HasStorageGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


