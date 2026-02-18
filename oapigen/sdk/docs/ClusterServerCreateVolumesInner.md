# ClusterServerCreateVolumesInner

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

### NewClusterServerCreateVolumesInner

`func NewClusterServerCreateVolumesInner(name string, datastoreId NullableString, ) *ClusterServerCreateVolumesInner`

NewClusterServerCreateVolumesInner instantiates a new ClusterServerCreateVolumesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterServerCreateVolumesInnerWithDefaults

`func NewClusterServerCreateVolumesInnerWithDefaults() *ClusterServerCreateVolumesInner`

NewClusterServerCreateVolumesInnerWithDefaults instantiates a new ClusterServerCreateVolumesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ClusterServerCreateVolumesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ClusterServerCreateVolumesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ClusterServerCreateVolumesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ClusterServerCreateVolumesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetRootVolume

`func (o *ClusterServerCreateVolumesInner) GetRootVolume() bool`

GetRootVolume returns the RootVolume field if non-nil, zero value otherwise.

### GetRootVolumeOk

`func (o *ClusterServerCreateVolumesInner) GetRootVolumeOk() (*bool, bool)`

GetRootVolumeOk returns a tuple with the RootVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolume

`func (o *ClusterServerCreateVolumesInner) SetRootVolume(v bool)`

SetRootVolume sets RootVolume field to given value.

### HasRootVolume

`func (o *ClusterServerCreateVolumesInner) HasRootVolume() bool`

HasRootVolume returns a boolean if a field has been set.

### GetName

`func (o *ClusterServerCreateVolumesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterServerCreateVolumesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterServerCreateVolumesInner) SetName(v string)`

SetName sets Name field to given value.


### GetSize

`func (o *ClusterServerCreateVolumesInner) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ClusterServerCreateVolumesInner) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ClusterServerCreateVolumesInner) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *ClusterServerCreateVolumesInner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetSizeId

`func (o *ClusterServerCreateVolumesInner) GetSizeId() string`

GetSizeId returns the SizeId field if non-nil, zero value otherwise.

### GetSizeIdOk

`func (o *ClusterServerCreateVolumesInner) GetSizeIdOk() (*string, bool)`

GetSizeIdOk returns a tuple with the SizeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSizeId

`func (o *ClusterServerCreateVolumesInner) SetSizeId(v string)`

SetSizeId sets SizeId field to given value.

### HasSizeId

`func (o *ClusterServerCreateVolumesInner) HasSizeId() bool`

HasSizeId returns a boolean if a field has been set.

### SetSizeIdNil

`func (o *ClusterServerCreateVolumesInner) SetSizeIdNil(b bool)`

 SetSizeIdNil sets the value for SizeId to be an explicit nil

### UnsetSizeId
`func (o *ClusterServerCreateVolumesInner) UnsetSizeId()`

UnsetSizeId ensures that no value is present for SizeId, not even an explicit nil
### GetStorageType

`func (o *ClusterServerCreateVolumesInner) GetStorageType() int64`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *ClusterServerCreateVolumesInner) GetStorageTypeOk() (*int64, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *ClusterServerCreateVolumesInner) SetStorageType(v int64)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *ClusterServerCreateVolumesInner) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetDatastoreId

`func (o *ClusterServerCreateVolumesInner) GetDatastoreId() string`

GetDatastoreId returns the DatastoreId field if non-nil, zero value otherwise.

### GetDatastoreIdOk

`func (o *ClusterServerCreateVolumesInner) GetDatastoreIdOk() (*string, bool)`

GetDatastoreIdOk returns a tuple with the DatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreId

`func (o *ClusterServerCreateVolumesInner) SetDatastoreId(v string)`

SetDatastoreId sets DatastoreId field to given value.


### SetDatastoreIdNil

`func (o *ClusterServerCreateVolumesInner) SetDatastoreIdNil(b bool)`

 SetDatastoreIdNil sets the value for DatastoreId to be an explicit nil

### UnsetDatastoreId
`func (o *ClusterServerCreateVolumesInner) UnsetDatastoreId()`

UnsetDatastoreId ensures that no value is present for DatastoreId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


