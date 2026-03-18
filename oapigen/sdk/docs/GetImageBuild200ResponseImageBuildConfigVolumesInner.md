# GetImageBuild200ResponseImageBuildConfigVolumesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Size** | Pointer to **int64** |  | [optional] 
**MaxIOPS** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**RootVolume** | Pointer to **bool** |  | [optional] 
**StorageType** | Pointer to **int64** |  | [optional] 
**DatastoreId** | Pointer to **string** |  | [optional] 

## Methods

### NewGetImageBuild200ResponseImageBuildConfigVolumesInner

`func NewGetImageBuild200ResponseImageBuildConfigVolumesInner() *GetImageBuild200ResponseImageBuildConfigVolumesInner`

NewGetImageBuild200ResponseImageBuildConfigVolumesInner instantiates a new GetImageBuild200ResponseImageBuildConfigVolumesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetImageBuild200ResponseImageBuildConfigVolumesInnerWithDefaults

`func NewGetImageBuild200ResponseImageBuildConfigVolumesInnerWithDefaults() *GetImageBuild200ResponseImageBuildConfigVolumesInner`

NewGetImageBuild200ResponseImageBuildConfigVolumesInnerWithDefaults instantiates a new GetImageBuild200ResponseImageBuildConfigVolumesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetImageBuild200ResponseImageBuildConfigVolumesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetImageBuild200ResponseImageBuildConfigVolumesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetImageBuild200ResponseImageBuildConfigVolumesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetImageBuild200ResponseImageBuildConfigVolumesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSize

`func (o *GetImageBuild200ResponseImageBuildConfigVolumesInner) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *GetImageBuild200ResponseImageBuildConfigVolumesInner) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *GetImageBuild200ResponseImageBuildConfigVolumesInner) SetSize(v int64)`

SetSize sets Size field to given value.

### HasSize

`func (o *GetImageBuild200ResponseImageBuildConfigVolumesInner) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetMaxIOPS

`func (o *GetImageBuild200ResponseImageBuildConfigVolumesInner) GetMaxIOPS() string`

GetMaxIOPS returns the MaxIOPS field if non-nil, zero value otherwise.

### GetMaxIOPSOk

`func (o *GetImageBuild200ResponseImageBuildConfigVolumesInner) GetMaxIOPSOk() (*string, bool)`

GetMaxIOPSOk returns a tuple with the MaxIOPS field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIOPS

`func (o *GetImageBuild200ResponseImageBuildConfigVolumesInner) SetMaxIOPS(v string)`

SetMaxIOPS sets MaxIOPS field to given value.

### HasMaxIOPS

`func (o *GetImageBuild200ResponseImageBuildConfigVolumesInner) HasMaxIOPS() bool`

HasMaxIOPS returns a boolean if a field has been set.

### SetMaxIOPSNil

`func (o *GetImageBuild200ResponseImageBuildConfigVolumesInner) SetMaxIOPSNil(b bool)`

 SetMaxIOPSNil sets the value for MaxIOPS to be an explicit nil

### UnsetMaxIOPS
`func (o *GetImageBuild200ResponseImageBuildConfigVolumesInner) UnsetMaxIOPS()`

UnsetMaxIOPS ensures that no value is present for MaxIOPS, not even an explicit nil
### GetName

`func (o *GetImageBuild200ResponseImageBuildConfigVolumesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetImageBuild200ResponseImageBuildConfigVolumesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetImageBuild200ResponseImageBuildConfigVolumesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetImageBuild200ResponseImageBuildConfigVolumesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRootVolume

`func (o *GetImageBuild200ResponseImageBuildConfigVolumesInner) GetRootVolume() bool`

GetRootVolume returns the RootVolume field if non-nil, zero value otherwise.

### GetRootVolumeOk

`func (o *GetImageBuild200ResponseImageBuildConfigVolumesInner) GetRootVolumeOk() (*bool, bool)`

GetRootVolumeOk returns a tuple with the RootVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRootVolume

`func (o *GetImageBuild200ResponseImageBuildConfigVolumesInner) SetRootVolume(v bool)`

SetRootVolume sets RootVolume field to given value.

### HasRootVolume

`func (o *GetImageBuild200ResponseImageBuildConfigVolumesInner) HasRootVolume() bool`

HasRootVolume returns a boolean if a field has been set.

### GetStorageType

`func (o *GetImageBuild200ResponseImageBuildConfigVolumesInner) GetStorageType() int64`

GetStorageType returns the StorageType field if non-nil, zero value otherwise.

### GetStorageTypeOk

`func (o *GetImageBuild200ResponseImageBuildConfigVolumesInner) GetStorageTypeOk() (*int64, bool)`

GetStorageTypeOk returns a tuple with the StorageType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageType

`func (o *GetImageBuild200ResponseImageBuildConfigVolumesInner) SetStorageType(v int64)`

SetStorageType sets StorageType field to given value.

### HasStorageType

`func (o *GetImageBuild200ResponseImageBuildConfigVolumesInner) HasStorageType() bool`

HasStorageType returns a boolean if a field has been set.

### GetDatastoreId

`func (o *GetImageBuild200ResponseImageBuildConfigVolumesInner) GetDatastoreId() string`

GetDatastoreId returns the DatastoreId field if non-nil, zero value otherwise.

### GetDatastoreIdOk

`func (o *GetImageBuild200ResponseImageBuildConfigVolumesInner) GetDatastoreIdOk() (*string, bool)`

GetDatastoreIdOk returns a tuple with the DatastoreId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastoreId

`func (o *GetImageBuild200ResponseImageBuildConfigVolumesInner) SetDatastoreId(v string)`

SetDatastoreId sets DatastoreId field to given value.

### HasDatastoreId

`func (o *GetImageBuild200ResponseImageBuildConfigVolumesInner) HasDatastoreId() bool`

HasDatastoreId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


