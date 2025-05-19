# AddStorageBucketsRequestStorageBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A unique name scoped to your account for the storage bucket | 
**ProviderType** | **string** | The type of storage bucket | 
**DefaultBackupTarget** | Pointer to **bool** | Default Backup Target | [optional] [default to false]
**CopyToStore** | Pointer to **bool** | Archive Snapshots | [optional] 
**DefaultDeploymentTarget** | Pointer to **bool** | Default Deployment Target | [optional] [default to false]
**DefaultVirtualImageTarget** | Pointer to **bool** | Default Virtual Image Store | [optional] [default to false]
**RetentionPolicyType** | Pointer to **string** | Cleanup mode. &#x60;backup&#x60; - Move old files to a backup provider. &#x60;delete&#x60; - Delete old files. &#x60;none&#x60; - Keep all files. | [optional] [default to "none"]
**RetentionPolicyDays** | Pointer to **int64** | The number of days old a file must be before it is deleted. | [optional] 
**RetentionProvider** | Pointer to **string** | The backup Storage Bucket where old files are moved to. | [optional] 
**BucketName** | **string** | The name of the bucket. Only applies to &#x60;Amazon&#x60;, &#x60;Azure&#x60;, &#x60;CIFS&#x60;, &#x60;NFSv3&#x60;, &#x60;Openstack Swift&#x60;, and &#x60;Rackspace CDN&#x60;. | 
**CreateBucket** | Pointer to **bool** | Create the bucket if it does not exist. Only applies to &#x60;Amazon&#x60;, &#x60;Azure&#x60;, &#x60;Openstack Swift&#x60;, and &#x60;Rackspace CDN&#x60;. | [optional] [default to false]
**Config** | [**AddStorageBucketsRequestStorageBucketConfig**](AddStorageBucketsRequestStorageBucketConfig.md) |  | 

## Methods

### NewAddStorageBucketsRequestStorageBucket

`func NewAddStorageBucketsRequestStorageBucket(name string, providerType string, bucketName string, config AddStorageBucketsRequestStorageBucketConfig, ) *AddStorageBucketsRequestStorageBucket`

NewAddStorageBucketsRequestStorageBucket instantiates a new AddStorageBucketsRequestStorageBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddStorageBucketsRequestStorageBucketWithDefaults

`func NewAddStorageBucketsRequestStorageBucketWithDefaults() *AddStorageBucketsRequestStorageBucket`

NewAddStorageBucketsRequestStorageBucketWithDefaults instantiates a new AddStorageBucketsRequestStorageBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddStorageBucketsRequestStorageBucket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddStorageBucketsRequestStorageBucket) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddStorageBucketsRequestStorageBucket) SetName(v string)`

SetName sets Name field to given value.


### GetProviderType

`func (o *AddStorageBucketsRequestStorageBucket) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *AddStorageBucketsRequestStorageBucket) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *AddStorageBucketsRequestStorageBucket) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.


### GetDefaultBackupTarget

`func (o *AddStorageBucketsRequestStorageBucket) GetDefaultBackupTarget() bool`

GetDefaultBackupTarget returns the DefaultBackupTarget field if non-nil, zero value otherwise.

### GetDefaultBackupTargetOk

`func (o *AddStorageBucketsRequestStorageBucket) GetDefaultBackupTargetOk() (*bool, bool)`

GetDefaultBackupTargetOk returns a tuple with the DefaultBackupTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBackupTarget

`func (o *AddStorageBucketsRequestStorageBucket) SetDefaultBackupTarget(v bool)`

SetDefaultBackupTarget sets DefaultBackupTarget field to given value.

### HasDefaultBackupTarget

`func (o *AddStorageBucketsRequestStorageBucket) HasDefaultBackupTarget() bool`

HasDefaultBackupTarget returns a boolean if a field has been set.

### GetCopyToStore

`func (o *AddStorageBucketsRequestStorageBucket) GetCopyToStore() bool`

GetCopyToStore returns the CopyToStore field if non-nil, zero value otherwise.

### GetCopyToStoreOk

`func (o *AddStorageBucketsRequestStorageBucket) GetCopyToStoreOk() (*bool, bool)`

GetCopyToStoreOk returns a tuple with the CopyToStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyToStore

`func (o *AddStorageBucketsRequestStorageBucket) SetCopyToStore(v bool)`

SetCopyToStore sets CopyToStore field to given value.

### HasCopyToStore

`func (o *AddStorageBucketsRequestStorageBucket) HasCopyToStore() bool`

HasCopyToStore returns a boolean if a field has been set.

### GetDefaultDeploymentTarget

`func (o *AddStorageBucketsRequestStorageBucket) GetDefaultDeploymentTarget() bool`

GetDefaultDeploymentTarget returns the DefaultDeploymentTarget field if non-nil, zero value otherwise.

### GetDefaultDeploymentTargetOk

`func (o *AddStorageBucketsRequestStorageBucket) GetDefaultDeploymentTargetOk() (*bool, bool)`

GetDefaultDeploymentTargetOk returns a tuple with the DefaultDeploymentTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDeploymentTarget

`func (o *AddStorageBucketsRequestStorageBucket) SetDefaultDeploymentTarget(v bool)`

SetDefaultDeploymentTarget sets DefaultDeploymentTarget field to given value.

### HasDefaultDeploymentTarget

`func (o *AddStorageBucketsRequestStorageBucket) HasDefaultDeploymentTarget() bool`

HasDefaultDeploymentTarget returns a boolean if a field has been set.

### GetDefaultVirtualImageTarget

`func (o *AddStorageBucketsRequestStorageBucket) GetDefaultVirtualImageTarget() bool`

GetDefaultVirtualImageTarget returns the DefaultVirtualImageTarget field if non-nil, zero value otherwise.

### GetDefaultVirtualImageTargetOk

`func (o *AddStorageBucketsRequestStorageBucket) GetDefaultVirtualImageTargetOk() (*bool, bool)`

GetDefaultVirtualImageTargetOk returns a tuple with the DefaultVirtualImageTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVirtualImageTarget

`func (o *AddStorageBucketsRequestStorageBucket) SetDefaultVirtualImageTarget(v bool)`

SetDefaultVirtualImageTarget sets DefaultVirtualImageTarget field to given value.

### HasDefaultVirtualImageTarget

`func (o *AddStorageBucketsRequestStorageBucket) HasDefaultVirtualImageTarget() bool`

HasDefaultVirtualImageTarget returns a boolean if a field has been set.

### GetRetentionPolicyType

`func (o *AddStorageBucketsRequestStorageBucket) GetRetentionPolicyType() string`

GetRetentionPolicyType returns the RetentionPolicyType field if non-nil, zero value otherwise.

### GetRetentionPolicyTypeOk

`func (o *AddStorageBucketsRequestStorageBucket) GetRetentionPolicyTypeOk() (*string, bool)`

GetRetentionPolicyTypeOk returns a tuple with the RetentionPolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPolicyType

`func (o *AddStorageBucketsRequestStorageBucket) SetRetentionPolicyType(v string)`

SetRetentionPolicyType sets RetentionPolicyType field to given value.

### HasRetentionPolicyType

`func (o *AddStorageBucketsRequestStorageBucket) HasRetentionPolicyType() bool`

HasRetentionPolicyType returns a boolean if a field has been set.

### GetRetentionPolicyDays

`func (o *AddStorageBucketsRequestStorageBucket) GetRetentionPolicyDays() int64`

GetRetentionPolicyDays returns the RetentionPolicyDays field if non-nil, zero value otherwise.

### GetRetentionPolicyDaysOk

`func (o *AddStorageBucketsRequestStorageBucket) GetRetentionPolicyDaysOk() (*int64, bool)`

GetRetentionPolicyDaysOk returns a tuple with the RetentionPolicyDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPolicyDays

`func (o *AddStorageBucketsRequestStorageBucket) SetRetentionPolicyDays(v int64)`

SetRetentionPolicyDays sets RetentionPolicyDays field to given value.

### HasRetentionPolicyDays

`func (o *AddStorageBucketsRequestStorageBucket) HasRetentionPolicyDays() bool`

HasRetentionPolicyDays returns a boolean if a field has been set.

### GetRetentionProvider

`func (o *AddStorageBucketsRequestStorageBucket) GetRetentionProvider() string`

GetRetentionProvider returns the RetentionProvider field if non-nil, zero value otherwise.

### GetRetentionProviderOk

`func (o *AddStorageBucketsRequestStorageBucket) GetRetentionProviderOk() (*string, bool)`

GetRetentionProviderOk returns a tuple with the RetentionProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionProvider

`func (o *AddStorageBucketsRequestStorageBucket) SetRetentionProvider(v string)`

SetRetentionProvider sets RetentionProvider field to given value.

### HasRetentionProvider

`func (o *AddStorageBucketsRequestStorageBucket) HasRetentionProvider() bool`

HasRetentionProvider returns a boolean if a field has been set.

### GetBucketName

`func (o *AddStorageBucketsRequestStorageBucket) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *AddStorageBucketsRequestStorageBucket) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *AddStorageBucketsRequestStorageBucket) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.


### GetCreateBucket

`func (o *AddStorageBucketsRequestStorageBucket) GetCreateBucket() bool`

GetCreateBucket returns the CreateBucket field if non-nil, zero value otherwise.

### GetCreateBucketOk

`func (o *AddStorageBucketsRequestStorageBucket) GetCreateBucketOk() (*bool, bool)`

GetCreateBucketOk returns a tuple with the CreateBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBucket

`func (o *AddStorageBucketsRequestStorageBucket) SetCreateBucket(v bool)`

SetCreateBucket sets CreateBucket field to given value.

### HasCreateBucket

`func (o *AddStorageBucketsRequestStorageBucket) HasCreateBucket() bool`

HasCreateBucket returns a boolean if a field has been set.

### GetConfig

`func (o *AddStorageBucketsRequestStorageBucket) GetConfig() AddStorageBucketsRequestStorageBucketConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddStorageBucketsRequestStorageBucket) GetConfigOk() (*AddStorageBucketsRequestStorageBucketConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddStorageBucketsRequestStorageBucket) SetConfig(v AddStorageBucketsRequestStorageBucketConfig)`

SetConfig sets Config field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


