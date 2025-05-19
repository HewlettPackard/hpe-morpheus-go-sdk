# AddStorageBuckets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageBucket** | Pointer to [**ListStorageBuckets200ResponseAllOfStorageBucketsInner**](ListStorageBuckets200ResponseAllOfStorageBucketsInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddStorageBuckets200Response

`func NewAddStorageBuckets200Response() *AddStorageBuckets200Response`

NewAddStorageBuckets200Response instantiates a new AddStorageBuckets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddStorageBuckets200ResponseWithDefaults

`func NewAddStorageBuckets200ResponseWithDefaults() *AddStorageBuckets200Response`

NewAddStorageBuckets200ResponseWithDefaults instantiates a new AddStorageBuckets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageBucket

`func (o *AddStorageBuckets200Response) GetStorageBucket() ListStorageBuckets200ResponseAllOfStorageBucketsInner`

GetStorageBucket returns the StorageBucket field if non-nil, zero value otherwise.

### GetStorageBucketOk

`func (o *AddStorageBuckets200Response) GetStorageBucketOk() (*ListStorageBuckets200ResponseAllOfStorageBucketsInner, bool)`

GetStorageBucketOk returns a tuple with the StorageBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageBucket

`func (o *AddStorageBuckets200Response) SetStorageBucket(v ListStorageBuckets200ResponseAllOfStorageBucketsInner)`

SetStorageBucket sets StorageBucket field to given value.

### HasStorageBucket

`func (o *AddStorageBuckets200Response) HasStorageBucket() bool`

HasStorageBucket returns a boolean if a field has been set.

### GetSuccess

`func (o *AddStorageBuckets200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddStorageBuckets200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddStorageBuckets200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddStorageBuckets200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


