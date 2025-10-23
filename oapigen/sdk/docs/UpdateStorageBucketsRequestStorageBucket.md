# UpdateStorageBucketsRequestStorageBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A unique name scoped to your account for the storage bucket | [optional] 
**ProviderType** | Pointer to **string** | The type of storage bucket | [optional] 
**DefaultBackupTarget** | Pointer to **bool** | Default Backup Target | [optional] [default to false]
**CopyToStore** | Pointer to **bool** | Archive Snapshots | [optional] 
**DefaultDeploymentTarget** | Pointer to **bool** | Default Deployment Target | [optional] [default to false]
**DefaultVirtualImageTarget** | Pointer to **bool** | Default Virtual Image Store | [optional] [default to false]
**RetentionPolicyType** | Pointer to **string** | Cleanup mode. &#x60;backup&#x60; - Move old files to a backup provider. &#x60;delete&#x60; - Delete old files. &#x60;none&#x60; - Keep all files. | [optional] [default to "none"]
**RetentionPolicyDays** | Pointer to **int64** | The number of days old a file must be before it is deleted. | [optional] 
**RetentionProvider** | Pointer to **string** | The backup Storage Bucket where old files are moved to. | [optional] 
**BucketName** | Pointer to **string** | The name of the bucket. Only applies to &#x60;Amazon&#x60;, &#x60;Azure&#x60;, &#x60;CIFS&#x60;, &#x60;NFSv3&#x60;, &#x60;Openstack Swift&#x60;, and &#x60;Rackspace CDN&#x60;. | [optional] 
**CreateBucket** | Pointer to **bool** | Create the bucket if it does not exist. Only applies to &#x60;Amazon&#x60;, &#x60;Azure&#x60;, &#x60;Openstack Swift&#x60;, and &#x60;Rackspace CDN&#x60;. | [optional] [default to false]
**Config** | Pointer to [**AddStorageBucketsRequestStorageBucketConfig**](AddStorageBucketsRequestStorageBucketConfig.md) |  | [optional] 

## Methods

### NewUpdateStorageBucketsRequestStorageBucket

`func NewUpdateStorageBucketsRequestStorageBucket() *UpdateStorageBucketsRequestStorageBucket`

NewUpdateStorageBucketsRequestStorageBucket instantiates a new UpdateStorageBucketsRequestStorageBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStorageBucketsRequestStorageBucketWithDefaults

`func NewUpdateStorageBucketsRequestStorageBucketWithDefaults() *UpdateStorageBucketsRequestStorageBucket`

NewUpdateStorageBucketsRequestStorageBucketWithDefaults instantiates a new UpdateStorageBucketsRequestStorageBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateStorageBucketsRequestStorageBucket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateStorageBucketsRequestStorageBucket) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateStorageBucketsRequestStorageBucket) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateStorageBucketsRequestStorageBucket) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProviderType

`func (o *UpdateStorageBucketsRequestStorageBucket) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *UpdateStorageBucketsRequestStorageBucket) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *UpdateStorageBucketsRequestStorageBucket) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.

### HasProviderType

`func (o *UpdateStorageBucketsRequestStorageBucket) HasProviderType() bool`

HasProviderType returns a boolean if a field has been set.

### GetDefaultBackupTarget

`func (o *UpdateStorageBucketsRequestStorageBucket) GetDefaultBackupTarget() bool`

GetDefaultBackupTarget returns the DefaultBackupTarget field if non-nil, zero value otherwise.

### GetDefaultBackupTargetOk

`func (o *UpdateStorageBucketsRequestStorageBucket) GetDefaultBackupTargetOk() (*bool, bool)`

GetDefaultBackupTargetOk returns a tuple with the DefaultBackupTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBackupTarget

`func (o *UpdateStorageBucketsRequestStorageBucket) SetDefaultBackupTarget(v bool)`

SetDefaultBackupTarget sets DefaultBackupTarget field to given value.

### HasDefaultBackupTarget

`func (o *UpdateStorageBucketsRequestStorageBucket) HasDefaultBackupTarget() bool`

HasDefaultBackupTarget returns a boolean if a field has been set.

### GetCopyToStore

`func (o *UpdateStorageBucketsRequestStorageBucket) GetCopyToStore() bool`

GetCopyToStore returns the CopyToStore field if non-nil, zero value otherwise.

### GetCopyToStoreOk

`func (o *UpdateStorageBucketsRequestStorageBucket) GetCopyToStoreOk() (*bool, bool)`

GetCopyToStoreOk returns a tuple with the CopyToStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyToStore

`func (o *UpdateStorageBucketsRequestStorageBucket) SetCopyToStore(v bool)`

SetCopyToStore sets CopyToStore field to given value.

### HasCopyToStore

`func (o *UpdateStorageBucketsRequestStorageBucket) HasCopyToStore() bool`

HasCopyToStore returns a boolean if a field has been set.

### GetDefaultDeploymentTarget

`func (o *UpdateStorageBucketsRequestStorageBucket) GetDefaultDeploymentTarget() bool`

GetDefaultDeploymentTarget returns the DefaultDeploymentTarget field if non-nil, zero value otherwise.

### GetDefaultDeploymentTargetOk

`func (o *UpdateStorageBucketsRequestStorageBucket) GetDefaultDeploymentTargetOk() (*bool, bool)`

GetDefaultDeploymentTargetOk returns a tuple with the DefaultDeploymentTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDeploymentTarget

`func (o *UpdateStorageBucketsRequestStorageBucket) SetDefaultDeploymentTarget(v bool)`

SetDefaultDeploymentTarget sets DefaultDeploymentTarget field to given value.

### HasDefaultDeploymentTarget

`func (o *UpdateStorageBucketsRequestStorageBucket) HasDefaultDeploymentTarget() bool`

HasDefaultDeploymentTarget returns a boolean if a field has been set.

### GetDefaultVirtualImageTarget

`func (o *UpdateStorageBucketsRequestStorageBucket) GetDefaultVirtualImageTarget() bool`

GetDefaultVirtualImageTarget returns the DefaultVirtualImageTarget field if non-nil, zero value otherwise.

### GetDefaultVirtualImageTargetOk

`func (o *UpdateStorageBucketsRequestStorageBucket) GetDefaultVirtualImageTargetOk() (*bool, bool)`

GetDefaultVirtualImageTargetOk returns a tuple with the DefaultVirtualImageTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVirtualImageTarget

`func (o *UpdateStorageBucketsRequestStorageBucket) SetDefaultVirtualImageTarget(v bool)`

SetDefaultVirtualImageTarget sets DefaultVirtualImageTarget field to given value.

### HasDefaultVirtualImageTarget

`func (o *UpdateStorageBucketsRequestStorageBucket) HasDefaultVirtualImageTarget() bool`

HasDefaultVirtualImageTarget returns a boolean if a field has been set.

### GetRetentionPolicyType

`func (o *UpdateStorageBucketsRequestStorageBucket) GetRetentionPolicyType() string`

GetRetentionPolicyType returns the RetentionPolicyType field if non-nil, zero value otherwise.

### GetRetentionPolicyTypeOk

`func (o *UpdateStorageBucketsRequestStorageBucket) GetRetentionPolicyTypeOk() (*string, bool)`

GetRetentionPolicyTypeOk returns a tuple with the RetentionPolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPolicyType

`func (o *UpdateStorageBucketsRequestStorageBucket) SetRetentionPolicyType(v string)`

SetRetentionPolicyType sets RetentionPolicyType field to given value.

### HasRetentionPolicyType

`func (o *UpdateStorageBucketsRequestStorageBucket) HasRetentionPolicyType() bool`

HasRetentionPolicyType returns a boolean if a field has been set.

### GetRetentionPolicyDays

`func (o *UpdateStorageBucketsRequestStorageBucket) GetRetentionPolicyDays() int64`

GetRetentionPolicyDays returns the RetentionPolicyDays field if non-nil, zero value otherwise.

### GetRetentionPolicyDaysOk

`func (o *UpdateStorageBucketsRequestStorageBucket) GetRetentionPolicyDaysOk() (*int64, bool)`

GetRetentionPolicyDaysOk returns a tuple with the RetentionPolicyDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPolicyDays

`func (o *UpdateStorageBucketsRequestStorageBucket) SetRetentionPolicyDays(v int64)`

SetRetentionPolicyDays sets RetentionPolicyDays field to given value.

### HasRetentionPolicyDays

`func (o *UpdateStorageBucketsRequestStorageBucket) HasRetentionPolicyDays() bool`

HasRetentionPolicyDays returns a boolean if a field has been set.

### GetRetentionProvider

`func (o *UpdateStorageBucketsRequestStorageBucket) GetRetentionProvider() string`

GetRetentionProvider returns the RetentionProvider field if non-nil, zero value otherwise.

### GetRetentionProviderOk

`func (o *UpdateStorageBucketsRequestStorageBucket) GetRetentionProviderOk() (*string, bool)`

GetRetentionProviderOk returns a tuple with the RetentionProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionProvider

`func (o *UpdateStorageBucketsRequestStorageBucket) SetRetentionProvider(v string)`

SetRetentionProvider sets RetentionProvider field to given value.

### HasRetentionProvider

`func (o *UpdateStorageBucketsRequestStorageBucket) HasRetentionProvider() bool`

HasRetentionProvider returns a boolean if a field has been set.

### GetBucketName

`func (o *UpdateStorageBucketsRequestStorageBucket) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *UpdateStorageBucketsRequestStorageBucket) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *UpdateStorageBucketsRequestStorageBucket) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *UpdateStorageBucketsRequestStorageBucket) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetCreateBucket

`func (o *UpdateStorageBucketsRequestStorageBucket) GetCreateBucket() bool`

GetCreateBucket returns the CreateBucket field if non-nil, zero value otherwise.

### GetCreateBucketOk

`func (o *UpdateStorageBucketsRequestStorageBucket) GetCreateBucketOk() (*bool, bool)`

GetCreateBucketOk returns a tuple with the CreateBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateBucket

`func (o *UpdateStorageBucketsRequestStorageBucket) SetCreateBucket(v bool)`

SetCreateBucket sets CreateBucket field to given value.

### HasCreateBucket

`func (o *UpdateStorageBucketsRequestStorageBucket) HasCreateBucket() bool`

HasCreateBucket returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateStorageBucketsRequestStorageBucket) GetConfig() AddStorageBucketsRequestStorageBucketConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateStorageBucketsRequestStorageBucket) GetConfigOk() (*AddStorageBucketsRequestStorageBucketConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateStorageBucketsRequestStorageBucket) SetConfig(v AddStorageBucketsRequestStorageBucketConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateStorageBucketsRequestStorageBucket) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


