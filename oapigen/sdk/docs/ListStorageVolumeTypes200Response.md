# ListStorageVolumeTypes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageVolumeTypes** | Pointer to [**[]ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner**](ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListStorageVolumeTypes200Response

`func NewListStorageVolumeTypes200Response() *ListStorageVolumeTypes200Response`

NewListStorageVolumeTypes200Response instantiates a new ListStorageVolumeTypes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListStorageVolumeTypes200ResponseWithDefaults

`func NewListStorageVolumeTypes200ResponseWithDefaults() *ListStorageVolumeTypes200Response`

NewListStorageVolumeTypes200ResponseWithDefaults instantiates a new ListStorageVolumeTypes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageVolumeTypes

`func (o *ListStorageVolumeTypes200Response) GetStorageVolumeTypes() []ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner`

GetStorageVolumeTypes returns the StorageVolumeTypes field if non-nil, zero value otherwise.

### GetStorageVolumeTypesOk

`func (o *ListStorageVolumeTypes200Response) GetStorageVolumeTypesOk() (*[]ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner, bool)`

GetStorageVolumeTypesOk returns a tuple with the StorageVolumeTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageVolumeTypes

`func (o *ListStorageVolumeTypes200Response) SetStorageVolumeTypes(v []ListStorageVolumeTypes200ResponseAllOfStorageVolumeTypesInner)`

SetStorageVolumeTypes sets StorageVolumeTypes field to given value.

### HasStorageVolumeTypes

`func (o *ListStorageVolumeTypes200Response) HasStorageVolumeTypes() bool`

HasStorageVolumeTypes returns a boolean if a field has been set.

### GetMeta

`func (o *ListStorageVolumeTypes200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListStorageVolumeTypes200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListStorageVolumeTypes200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListStorageVolumeTypes200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


