# ListStorageBuckets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageBuckets** | Pointer to [**[]ListStorageBuckets200ResponseAllOfStorageBucketsInner**](ListStorageBuckets200ResponseAllOfStorageBucketsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListStorageBuckets200Response

`func NewListStorageBuckets200Response() *ListStorageBuckets200Response`

NewListStorageBuckets200Response instantiates a new ListStorageBuckets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListStorageBuckets200ResponseWithDefaults

`func NewListStorageBuckets200ResponseWithDefaults() *ListStorageBuckets200Response`

NewListStorageBuckets200ResponseWithDefaults instantiates a new ListStorageBuckets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageBuckets

`func (o *ListStorageBuckets200Response) GetStorageBuckets() []ListStorageBuckets200ResponseAllOfStorageBucketsInner`

GetStorageBuckets returns the StorageBuckets field if non-nil, zero value otherwise.

### GetStorageBucketsOk

`func (o *ListStorageBuckets200Response) GetStorageBucketsOk() (*[]ListStorageBuckets200ResponseAllOfStorageBucketsInner, bool)`

GetStorageBucketsOk returns a tuple with the StorageBuckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageBuckets

`func (o *ListStorageBuckets200Response) SetStorageBuckets(v []ListStorageBuckets200ResponseAllOfStorageBucketsInner)`

SetStorageBuckets sets StorageBuckets field to given value.

### HasStorageBuckets

`func (o *ListStorageBuckets200Response) HasStorageBuckets() bool`

HasStorageBuckets returns a boolean if a field has been set.

### GetMeta

`func (o *ListStorageBuckets200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListStorageBuckets200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListStorageBuckets200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListStorageBuckets200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


