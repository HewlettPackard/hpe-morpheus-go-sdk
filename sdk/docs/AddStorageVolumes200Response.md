# AddStorageVolumes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageVolume** | Pointer to [**ListClusterVolumes200ResponseAllOfVolumesInner**](ListClusterVolumes200ResponseAllOfVolumesInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddStorageVolumes200Response

`func NewAddStorageVolumes200Response() *AddStorageVolumes200Response`

NewAddStorageVolumes200Response instantiates a new AddStorageVolumes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddStorageVolumes200ResponseWithDefaults

`func NewAddStorageVolumes200ResponseWithDefaults() *AddStorageVolumes200Response`

NewAddStorageVolumes200ResponseWithDefaults instantiates a new AddStorageVolumes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageVolume

`func (o *AddStorageVolumes200Response) GetStorageVolume() ListClusterVolumes200ResponseAllOfVolumesInner`

GetStorageVolume returns the StorageVolume field if non-nil, zero value otherwise.

### GetStorageVolumeOk

`func (o *AddStorageVolumes200Response) GetStorageVolumeOk() (*ListClusterVolumes200ResponseAllOfVolumesInner, bool)`

GetStorageVolumeOk returns a tuple with the StorageVolume field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageVolume

`func (o *AddStorageVolumes200Response) SetStorageVolume(v ListClusterVolumes200ResponseAllOfVolumesInner)`

SetStorageVolume sets StorageVolume field to given value.

### HasStorageVolume

`func (o *AddStorageVolumes200Response) HasStorageVolume() bool`

HasStorageVolume returns a boolean if a field has been set.

### GetSuccess

`func (o *AddStorageVolumes200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddStorageVolumes200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddStorageVolumes200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddStorageVolumes200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


