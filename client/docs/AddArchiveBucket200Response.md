# AddArchiveBucket200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveBucket** | Pointer to [**ListArchiveBuckets200ResponseAllOfArchiveBucketsInner**](ListArchiveBuckets200ResponseAllOfArchiveBucketsInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddArchiveBucket200Response

`func NewAddArchiveBucket200Response() *AddArchiveBucket200Response`

NewAddArchiveBucket200Response instantiates a new AddArchiveBucket200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddArchiveBucket200ResponseWithDefaults

`func NewAddArchiveBucket200ResponseWithDefaults() *AddArchiveBucket200Response`

NewAddArchiveBucket200ResponseWithDefaults instantiates a new AddArchiveBucket200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveBucket

`func (o *AddArchiveBucket200Response) GetArchiveBucket() ListArchiveBuckets200ResponseAllOfArchiveBucketsInner`

GetArchiveBucket returns the ArchiveBucket field if non-nil, zero value otherwise.

### GetArchiveBucketOk

`func (o *AddArchiveBucket200Response) GetArchiveBucketOk() (*ListArchiveBuckets200ResponseAllOfArchiveBucketsInner, bool)`

GetArchiveBucketOk returns a tuple with the ArchiveBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveBucket

`func (o *AddArchiveBucket200Response) SetArchiveBucket(v ListArchiveBuckets200ResponseAllOfArchiveBucketsInner)`

SetArchiveBucket sets ArchiveBucket field to given value.

### HasArchiveBucket

`func (o *AddArchiveBucket200Response) HasArchiveBucket() bool`

HasArchiveBucket returns a boolean if a field has been set.

### GetSuccess

`func (o *AddArchiveBucket200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddArchiveBucket200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddArchiveBucket200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddArchiveBucket200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


