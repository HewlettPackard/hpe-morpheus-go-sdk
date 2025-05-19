# ListArchiveBuckets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveBuckets** | Pointer to [**[]ListArchiveBuckets200ResponseAllOfArchiveBucketsInner**](ListArchiveBuckets200ResponseAllOfArchiveBucketsInner.md) |  | [optional] 
**ArchiveBucketCount** | Pointer to **int64** |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListArchiveBuckets200Response

`func NewListArchiveBuckets200Response() *ListArchiveBuckets200Response`

NewListArchiveBuckets200Response instantiates a new ListArchiveBuckets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListArchiveBuckets200ResponseWithDefaults

`func NewListArchiveBuckets200ResponseWithDefaults() *ListArchiveBuckets200Response`

NewListArchiveBuckets200ResponseWithDefaults instantiates a new ListArchiveBuckets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveBuckets

`func (o *ListArchiveBuckets200Response) GetArchiveBuckets() []ListArchiveBuckets200ResponseAllOfArchiveBucketsInner`

GetArchiveBuckets returns the ArchiveBuckets field if non-nil, zero value otherwise.

### GetArchiveBucketsOk

`func (o *ListArchiveBuckets200Response) GetArchiveBucketsOk() (*[]ListArchiveBuckets200ResponseAllOfArchiveBucketsInner, bool)`

GetArchiveBucketsOk returns a tuple with the ArchiveBuckets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveBuckets

`func (o *ListArchiveBuckets200Response) SetArchiveBuckets(v []ListArchiveBuckets200ResponseAllOfArchiveBucketsInner)`

SetArchiveBuckets sets ArchiveBuckets field to given value.

### HasArchiveBuckets

`func (o *ListArchiveBuckets200Response) HasArchiveBuckets() bool`

HasArchiveBuckets returns a boolean if a field has been set.

### GetArchiveBucketCount

`func (o *ListArchiveBuckets200Response) GetArchiveBucketCount() int64`

GetArchiveBucketCount returns the ArchiveBucketCount field if non-nil, zero value otherwise.

### GetArchiveBucketCountOk

`func (o *ListArchiveBuckets200Response) GetArchiveBucketCountOk() (*int64, bool)`

GetArchiveBucketCountOk returns a tuple with the ArchiveBucketCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveBucketCount

`func (o *ListArchiveBuckets200Response) SetArchiveBucketCount(v int64)`

SetArchiveBucketCount sets ArchiveBucketCount field to given value.

### HasArchiveBucketCount

`func (o *ListArchiveBuckets200Response) HasArchiveBucketCount() bool`

HasArchiveBucketCount returns a boolean if a field has been set.

### GetMeta

`func (o *ListArchiveBuckets200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListArchiveBuckets200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListArchiveBuckets200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListArchiveBuckets200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


