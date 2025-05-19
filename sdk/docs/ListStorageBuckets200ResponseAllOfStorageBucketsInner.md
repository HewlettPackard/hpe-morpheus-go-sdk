# ListStorageBuckets200ResponseAllOfStorageBucketsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**ProviderType** | Pointer to **string** |  | [optional] 
**Config** | Pointer to [**ListStorageBuckets200ResponseAllOfStorageBucketsInnerConfig**](ListStorageBuckets200ResponseAllOfStorageBucketsInnerConfig.md) |  | [optional] 
**BucketName** | Pointer to **string** |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**DefaultBackupTarget** | Pointer to **bool** |  | [optional] 
**DefaultDeploymentTarget** | Pointer to **bool** |  | [optional] 
**DefaultVirtualImageTarget** | Pointer to **bool** |  | [optional] 
**CopyToStore** | Pointer to **bool** |  | [optional] 
**RetentionPolicyType** | Pointer to **string** |  | [optional] 
**RetentionPolicyDays** | Pointer to **string** |  | [optional] 
**RetentionProvider** | Pointer to **string** |  | [optional] 

## Methods

### NewListStorageBuckets200ResponseAllOfStorageBucketsInner

`func NewListStorageBuckets200ResponseAllOfStorageBucketsInner() *ListStorageBuckets200ResponseAllOfStorageBucketsInner`

NewListStorageBuckets200ResponseAllOfStorageBucketsInner instantiates a new ListStorageBuckets200ResponseAllOfStorageBucketsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListStorageBuckets200ResponseAllOfStorageBucketsInnerWithDefaults

`func NewListStorageBuckets200ResponseAllOfStorageBucketsInnerWithDefaults() *ListStorageBuckets200ResponseAllOfStorageBucketsInner`

NewListStorageBuckets200ResponseAllOfStorageBucketsInnerWithDefaults instantiates a new ListStorageBuckets200ResponseAllOfStorageBucketsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAccountId

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetProviderType

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.

### HasProviderType

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) HasProviderType() bool`

HasProviderType returns a boolean if a field has been set.

### GetConfig

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) GetConfig() ListStorageBuckets200ResponseAllOfStorageBucketsInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) GetConfigOk() (*ListStorageBuckets200ResponseAllOfStorageBucketsInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) SetConfig(v ListStorageBuckets200ResponseAllOfStorageBucketsInnerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetBucketName

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetReadOnly

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetDefaultBackupTarget

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) GetDefaultBackupTarget() bool`

GetDefaultBackupTarget returns the DefaultBackupTarget field if non-nil, zero value otherwise.

### GetDefaultBackupTargetOk

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) GetDefaultBackupTargetOk() (*bool, bool)`

GetDefaultBackupTargetOk returns a tuple with the DefaultBackupTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBackupTarget

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) SetDefaultBackupTarget(v bool)`

SetDefaultBackupTarget sets DefaultBackupTarget field to given value.

### HasDefaultBackupTarget

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) HasDefaultBackupTarget() bool`

HasDefaultBackupTarget returns a boolean if a field has been set.

### GetDefaultDeploymentTarget

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) GetDefaultDeploymentTarget() bool`

GetDefaultDeploymentTarget returns the DefaultDeploymentTarget field if non-nil, zero value otherwise.

### GetDefaultDeploymentTargetOk

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) GetDefaultDeploymentTargetOk() (*bool, bool)`

GetDefaultDeploymentTargetOk returns a tuple with the DefaultDeploymentTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDeploymentTarget

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) SetDefaultDeploymentTarget(v bool)`

SetDefaultDeploymentTarget sets DefaultDeploymentTarget field to given value.

### HasDefaultDeploymentTarget

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) HasDefaultDeploymentTarget() bool`

HasDefaultDeploymentTarget returns a boolean if a field has been set.

### GetDefaultVirtualImageTarget

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) GetDefaultVirtualImageTarget() bool`

GetDefaultVirtualImageTarget returns the DefaultVirtualImageTarget field if non-nil, zero value otherwise.

### GetDefaultVirtualImageTargetOk

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) GetDefaultVirtualImageTargetOk() (*bool, bool)`

GetDefaultVirtualImageTargetOk returns a tuple with the DefaultVirtualImageTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVirtualImageTarget

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) SetDefaultVirtualImageTarget(v bool)`

SetDefaultVirtualImageTarget sets DefaultVirtualImageTarget field to given value.

### HasDefaultVirtualImageTarget

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) HasDefaultVirtualImageTarget() bool`

HasDefaultVirtualImageTarget returns a boolean if a field has been set.

### GetCopyToStore

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) GetCopyToStore() bool`

GetCopyToStore returns the CopyToStore field if non-nil, zero value otherwise.

### GetCopyToStoreOk

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) GetCopyToStoreOk() (*bool, bool)`

GetCopyToStoreOk returns a tuple with the CopyToStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyToStore

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) SetCopyToStore(v bool)`

SetCopyToStore sets CopyToStore field to given value.

### HasCopyToStore

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) HasCopyToStore() bool`

HasCopyToStore returns a boolean if a field has been set.

### GetRetentionPolicyType

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) GetRetentionPolicyType() string`

GetRetentionPolicyType returns the RetentionPolicyType field if non-nil, zero value otherwise.

### GetRetentionPolicyTypeOk

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) GetRetentionPolicyTypeOk() (*string, bool)`

GetRetentionPolicyTypeOk returns a tuple with the RetentionPolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPolicyType

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) SetRetentionPolicyType(v string)`

SetRetentionPolicyType sets RetentionPolicyType field to given value.

### HasRetentionPolicyType

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) HasRetentionPolicyType() bool`

HasRetentionPolicyType returns a boolean if a field has been set.

### GetRetentionPolicyDays

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) GetRetentionPolicyDays() string`

GetRetentionPolicyDays returns the RetentionPolicyDays field if non-nil, zero value otherwise.

### GetRetentionPolicyDaysOk

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) GetRetentionPolicyDaysOk() (*string, bool)`

GetRetentionPolicyDaysOk returns a tuple with the RetentionPolicyDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPolicyDays

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) SetRetentionPolicyDays(v string)`

SetRetentionPolicyDays sets RetentionPolicyDays field to given value.

### HasRetentionPolicyDays

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) HasRetentionPolicyDays() bool`

HasRetentionPolicyDays returns a boolean if a field has been set.

### GetRetentionProvider

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) GetRetentionProvider() string`

GetRetentionProvider returns the RetentionProvider field if non-nil, zero value otherwise.

### GetRetentionProviderOk

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) GetRetentionProviderOk() (*string, bool)`

GetRetentionProviderOk returns a tuple with the RetentionProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionProvider

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) SetRetentionProvider(v string)`

SetRetentionProvider sets RetentionProvider field to given value.

### HasRetentionProvider

`func (o *ListStorageBuckets200ResponseAllOfStorageBucketsInner) HasRetentionProvider() bool`

HasRetentionProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


