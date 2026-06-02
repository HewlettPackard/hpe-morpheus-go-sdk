# ListStorageVolumes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageVolumes** | Pointer to [**[]ListStorageVolumes200ResponseAllOfStorageVolumesInner**](ListStorageVolumes200ResponseAllOfStorageVolumesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListApprovals200ResponseAllOfMeta**](ListApprovals200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListStorageVolumes200Response

`func NewListStorageVolumes200Response() *ListStorageVolumes200Response`

NewListStorageVolumes200Response instantiates a new ListStorageVolumes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetStorageVolumes

`func (o *ListStorageVolumes200Response) GetStorageVolumes() []ListStorageVolumes200ResponseAllOfStorageVolumesInner`

GetStorageVolumes returns the StorageVolumes field if non-nil, zero value otherwise.

### GetStorageVolumesOk

`func (o *ListStorageVolumes200Response) GetStorageVolumesOk() (*[]ListStorageVolumes200ResponseAllOfStorageVolumesInner, bool)`

GetStorageVolumesOk returns a tuple with the StorageVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageVolumes

`func (o *ListStorageVolumes200Response) SetStorageVolumes(v []ListStorageVolumes200ResponseAllOfStorageVolumesInner)`

SetStorageVolumes sets StorageVolumes field to given value.

### HasStorageVolumes

`func (o *ListStorageVolumes200Response) HasStorageVolumes() bool`

HasStorageVolumes returns a boolean if a field has been set.

### GetMeta

`func (o *ListStorageVolumes200Response) GetMeta() ListApprovals200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListStorageVolumes200Response) GetMetaOk() (*ListApprovals200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListStorageVolumes200Response) SetMeta(v ListApprovals200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListStorageVolumes200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


