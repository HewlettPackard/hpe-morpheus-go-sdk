# ListStorageServerTypes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageServerTypes** | Pointer to [**[]ListStorageServerTypes200ResponseAllOfStorageServerTypesInner**](ListStorageServerTypes200ResponseAllOfStorageServerTypesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListStorageServerTypes200Response

`func NewListStorageServerTypes200Response() *ListStorageServerTypes200Response`

NewListStorageServerTypes200Response instantiates a new ListStorageServerTypes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListStorageServerTypes200ResponseWithDefaults

`func NewListStorageServerTypes200ResponseWithDefaults() *ListStorageServerTypes200Response`

NewListStorageServerTypes200ResponseWithDefaults instantiates a new ListStorageServerTypes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageServerTypes

`func (o *ListStorageServerTypes200Response) GetStorageServerTypes() []ListStorageServerTypes200ResponseAllOfStorageServerTypesInner`

GetStorageServerTypes returns the StorageServerTypes field if non-nil, zero value otherwise.

### GetStorageServerTypesOk

`func (o *ListStorageServerTypes200Response) GetStorageServerTypesOk() (*[]ListStorageServerTypes200ResponseAllOfStorageServerTypesInner, bool)`

GetStorageServerTypesOk returns a tuple with the StorageServerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServerTypes

`func (o *ListStorageServerTypes200Response) SetStorageServerTypes(v []ListStorageServerTypes200ResponseAllOfStorageServerTypesInner)`

SetStorageServerTypes sets StorageServerTypes field to given value.

### HasStorageServerTypes

`func (o *ListStorageServerTypes200Response) HasStorageServerTypes() bool`

HasStorageServerTypes returns a boolean if a field has been set.

### GetMeta

`func (o *ListStorageServerTypes200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListStorageServerTypes200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListStorageServerTypes200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListStorageServerTypes200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


