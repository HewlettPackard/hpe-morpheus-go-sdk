# AddClusterWorkerRequestServerVolumesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The id for the LV configuration being created | [optional] [default to -1]
**RootVolume** | Pointer to **bool** | If set to false then a non-root LV will be created | [optional] [default to true]
**Name** | **string** | Name/type of the LV being created | [default to "root"]
**Size** | Pointer to **int64** | Size of the LV to be created in GBs  Default is from the service plan  | [optional] 
**SizeId** | Pointer to **NullableString** | Can be used to select pre-existing LV choices from Morpheus | [optional] 
**StorageType** | Pointer to **int64** | Identifier for LV type | [optional] 
**DatastoreId** | **NullableString** | The ID of the specific datastore. Auto selection can be specified as auto or &#x60;autoCluster&#x60; (for clusters). | 

## Methods

### NewAddClusterWorkerRequestServerVolumesInner

`func NewAddClusterWorkerRequestServerVolumesInner(name string, datastoreId NullableString, ) *AddClusterWorkerRequestServerVolumesInner`

NewAddClusterWorkerRequestServerVolumesInner instantiates a new AddClusterWorkerRequestServerVolumesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *AddClusterWorkerRequestServerVolumesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddClusterWorkerRequestServerVolumesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddClusterWorkerRequestServerVolumesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddClusterWorkerRequestServerVolumesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRootVolume

`func (o *AddClusterWorkerRequestServerVolumesInner) GetRootVolume() bool`

GetRootVolume returns the RootVolume field if non-nil, zero value otherwise.

### GetRootVolumeOk

`func (o *AddClusterWorkerRequestServerVolumesInner) GetRootVolumeOk() (*bool, bool)`

GetRootVolumeOk returns a tuple with the RootVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolume

`func (o *AddClusterWorkerRequestServerVolumesInner) SetRootVolume(v bool)`

SetRootVolume sets RootVolume field to given value.

### HasRootVolume

`func (o *AddClusterWorkerRequestServerVolumesInner) HasRootVolume() bool`

HasRootVolume returns a boolean if a field has been set.

### GetName

`func (o *AddClusterWorkerRequestServerVolumesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddClusterWorkerRequestServerVolumesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddClusterWorkerRequestServerVolumesInner) SetName(v string)`

SetName sets Name field to given value.


### GetSize

`func (o *AddClusterWorkerRequestServerVolumesInner) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *AddClusterWorkerRequestServerVolumesInner) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *AddClusterWorkerRequestServerVolumesInner) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *AddClusterWorkerRequestServerVolumesInner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSizeId

`func (o *AddClusterWorkerRequestServerVolumesInner) GetSizeId() string`

GetSizeId returns the SizeId field if non-nil, zero value otherwise.

### GetSizeIdOk

`func (o *AddClusterWorkerRequestServerVolumesInner) GetSizeIdOk() (*string, bool)`

GetSizeIdOk returns a tuple with the SizeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeId

`func (o *AddClusterWorkerRequestServerVolumesInner) SetSizeId(v string)`

SetSizeId sets SizeId field to given value.

### HasSizeId

`func (o *AddClusterWorkerRequestServerVolumesInner) HasSizeId() bool`

HasSizeId returns a boolean if a field has been set.

### SetSizeIdNil

`func (o *AddClusterWorkerRequestServerVolumesInner) SetSizeIdNil(b bool)`

 SetSizeIdNil sets the value for SizeId to be an explicit nil

### UnsetSizeId
`func (o *AddClusterWorkerRequestServerVolumesInner) UnsetSizeId()`

UnsetSizeId ensures that no value is present for SizeId, not even an explicit nil
### GetStorageType

`func (o *AddClusterWorkerRequestServerVolumesInner) GetStorageType() int64`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *AddClusterWorkerRequestServerVolumesInner) GetStorageTypeOk() (*int64, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *AddClusterWorkerRequestServerVolumesInner) SetStorageType(v int64)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *AddClusterWorkerRequestServerVolumesInner) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetDatastoreId

`func (o *AddClusterWorkerRequestServerVolumesInner) GetDatastoreId() string`

GetDatastoreId returns the DatastoreId field if non-nil, zero value otherwise.

### GetDatastoreIdOk

`func (o *AddClusterWorkerRequestServerVolumesInner) GetDatastoreIdOk() (*string, bool)`

GetDatastoreIdOk returns a tuple with the DatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreId

`func (o *AddClusterWorkerRequestServerVolumesInner) SetDatastoreId(v string)`

SetDatastoreId sets DatastoreId field to given value.


### SetDatastoreIdNil

`func (o *AddClusterWorkerRequestServerVolumesInner) SetDatastoreIdNil(b bool)`

 SetDatastoreIdNil sets the value for DatastoreId to be an explicit nil

### UnsetDatastoreId
`func (o *AddClusterWorkerRequestServerVolumesInner) UnsetDatastoreId()`

UnsetDatastoreId ensures that no value is present for DatastoreId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


