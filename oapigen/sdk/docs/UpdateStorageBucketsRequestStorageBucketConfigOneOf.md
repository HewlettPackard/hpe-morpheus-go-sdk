# UpdateStorageBucketsRequestStorageBucketConfigOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** | Access Key | [optional] 
**SecretKey** | Pointer to **string** | Secret Key | [optional] 
**Region** | Pointer to **string** | Optional Amazon region if creating a new bucket | [optional] 
**Endpoint** | Pointer to **string** | Optional endpoint URL if pointing to an object store other than amazon that mimics the Amazon S3 APIs. | [optional] 

## Methods

### NewUpdateStorageBucketsRequestStorageBucketConfigOneOf

`func NewUpdateStorageBucketsRequestStorageBucketConfigOneOf() *UpdateStorageBucketsRequestStorageBucketConfigOneOf`

NewUpdateStorageBucketsRequestStorageBucketConfigOneOf instantiates a new UpdateStorageBucketsRequestStorageBucketConfigOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStorageBucketsRequestStorageBucketConfigOneOfWithDefaults

`func NewUpdateStorageBucketsRequestStorageBucketConfigOneOfWithDefaults() *UpdateStorageBucketsRequestStorageBucketConfigOneOf`

NewUpdateStorageBucketsRequestStorageBucketConfigOneOfWithDefaults instantiates a new UpdateStorageBucketsRequestStorageBucketConfigOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *UpdateStorageBucketsRequestStorageBucketConfigOneOf) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *UpdateStorageBucketsRequestStorageBucketConfigOneOf) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *UpdateStorageBucketsRequestStorageBucketConfigOneOf) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *UpdateStorageBucketsRequestStorageBucketConfigOneOf) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *UpdateStorageBucketsRequestStorageBucketConfigOneOf) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *UpdateStorageBucketsRequestStorageBucketConfigOneOf) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *UpdateStorageBucketsRequestStorageBucketConfigOneOf) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *UpdateStorageBucketsRequestStorageBucketConfigOneOf) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetRegion

`func (o *UpdateStorageBucketsRequestStorageBucketConfigOneOf) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *UpdateStorageBucketsRequestStorageBucketConfigOneOf) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *UpdateStorageBucketsRequestStorageBucketConfigOneOf) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *UpdateStorageBucketsRequestStorageBucketConfigOneOf) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetEndpoint

`func (o *UpdateStorageBucketsRequestStorageBucketConfigOneOf) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *UpdateStorageBucketsRequestStorageBucketConfigOneOf) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *UpdateStorageBucketsRequestStorageBucketConfigOneOf) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *UpdateStorageBucketsRequestStorageBucketConfigOneOf) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


