# UpdateStorageBuckets200ResponseAllOfStorageBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**ProviderType** | Pointer to **string** |  | [optional] 
**StorageServer** | Pointer to [**UpdateStorageBuckets200ResponseAllOfStorageBucketStorageServer**](UpdateStorageBuckets200ResponseAllOfStorageBucketStorageServer.md) |  | [optional] 
**Config** | Pointer to [**UpdateStorageBuckets200ResponseAllOfStorageBucketConfig**](UpdateStorageBuckets200ResponseAllOfStorageBucketConfig.md) |  | [optional] 
**BucketName** | Pointer to **string** |  | [optional] 
**ReadOnly** | Pointer to **bool** |  | [optional] 
**DefaultBackupTarget** | Pointer to **bool** |  | [optional] 
**DefaultDeploymentTarget** | Pointer to **bool** |  | [optional] 
**DefaultVirtualImageTarget** | Pointer to **bool** |  | [optional] 
**CopyToStore** | Pointer to **bool** |  | [optional] 
**RetentionPolicyType** | Pointer to **NullableString** |  | [optional] 
**RetentionPolicyDays** | Pointer to **NullableString** |  | [optional] 
**RetentionProvider** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateStorageBuckets200ResponseAllOfStorageBucket

`func NewUpdateStorageBuckets200ResponseAllOfStorageBucket() *UpdateStorageBuckets200ResponseAllOfStorageBucket`

NewUpdateStorageBuckets200ResponseAllOfStorageBucket instantiates a new UpdateStorageBuckets200ResponseAllOfStorageBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateStorageBuckets200ResponseAllOfStorageBucketWithDefaults

`func NewUpdateStorageBuckets200ResponseAllOfStorageBucketWithDefaults() *UpdateStorageBuckets200ResponseAllOfStorageBucket`

NewUpdateStorageBuckets200ResponseAllOfStorageBucketWithDefaults instantiates a new UpdateStorageBuckets200ResponseAllOfStorageBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAccountId

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetProviderType

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.

### HasProviderType

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) HasProviderType() bool`

HasProviderType returns a boolean if a field has been set.

### GetStorageServer

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) GetStorageServer() UpdateStorageBuckets200ResponseAllOfStorageBucketStorageServer`

GetStorageServer returns the StorageServer field if non-nil, zero value otherwise.

### GetStorageServerOk

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) GetStorageServerOk() (*UpdateStorageBuckets200ResponseAllOfStorageBucketStorageServer, bool)`

GetStorageServerOk returns a tuple with the StorageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServer

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) SetStorageServer(v UpdateStorageBuckets200ResponseAllOfStorageBucketStorageServer)`

SetStorageServer sets StorageServer field to given value.

### HasStorageServer

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) HasStorageServer() bool`

HasStorageServer returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) GetConfig() UpdateStorageBuckets200ResponseAllOfStorageBucketConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) GetConfigOk() (*UpdateStorageBuckets200ResponseAllOfStorageBucketConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) SetConfig(v UpdateStorageBuckets200ResponseAllOfStorageBucketConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetBucketName

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetReadOnly

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetDefaultBackupTarget

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) GetDefaultBackupTarget() bool`

GetDefaultBackupTarget returns the DefaultBackupTarget field if non-nil, zero value otherwise.

### GetDefaultBackupTargetOk

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) GetDefaultBackupTargetOk() (*bool, bool)`

GetDefaultBackupTargetOk returns a tuple with the DefaultBackupTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBackupTarget

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) SetDefaultBackupTarget(v bool)`

SetDefaultBackupTarget sets DefaultBackupTarget field to given value.

### HasDefaultBackupTarget

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) HasDefaultBackupTarget() bool`

HasDefaultBackupTarget returns a boolean if a field has been set.

### GetDefaultDeploymentTarget

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) GetDefaultDeploymentTarget() bool`

GetDefaultDeploymentTarget returns the DefaultDeploymentTarget field if non-nil, zero value otherwise.

### GetDefaultDeploymentTargetOk

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) GetDefaultDeploymentTargetOk() (*bool, bool)`

GetDefaultDeploymentTargetOk returns a tuple with the DefaultDeploymentTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDeploymentTarget

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) SetDefaultDeploymentTarget(v bool)`

SetDefaultDeploymentTarget sets DefaultDeploymentTarget field to given value.

### HasDefaultDeploymentTarget

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) HasDefaultDeploymentTarget() bool`

HasDefaultDeploymentTarget returns a boolean if a field has been set.

### GetDefaultVirtualImageTarget

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) GetDefaultVirtualImageTarget() bool`

GetDefaultVirtualImageTarget returns the DefaultVirtualImageTarget field if non-nil, zero value otherwise.

### GetDefaultVirtualImageTargetOk

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) GetDefaultVirtualImageTargetOk() (*bool, bool)`

GetDefaultVirtualImageTargetOk returns a tuple with the DefaultVirtualImageTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVirtualImageTarget

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) SetDefaultVirtualImageTarget(v bool)`

SetDefaultVirtualImageTarget sets DefaultVirtualImageTarget field to given value.

### HasDefaultVirtualImageTarget

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) HasDefaultVirtualImageTarget() bool`

HasDefaultVirtualImageTarget returns a boolean if a field has been set.

### GetCopyToStore

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) GetCopyToStore() bool`

GetCopyToStore returns the CopyToStore field if non-nil, zero value otherwise.

### GetCopyToStoreOk

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) GetCopyToStoreOk() (*bool, bool)`

GetCopyToStoreOk returns a tuple with the CopyToStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyToStore

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) SetCopyToStore(v bool)`

SetCopyToStore sets CopyToStore field to given value.

### HasCopyToStore

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) HasCopyToStore() bool`

HasCopyToStore returns a boolean if a field has been set.

### GetRetentionPolicyType

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) GetRetentionPolicyType() string`

GetRetentionPolicyType returns the RetentionPolicyType field if non-nil, zero value otherwise.

### GetRetentionPolicyTypeOk

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) GetRetentionPolicyTypeOk() (*string, bool)`

GetRetentionPolicyTypeOk returns a tuple with the RetentionPolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPolicyType

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) SetRetentionPolicyType(v string)`

SetRetentionPolicyType sets RetentionPolicyType field to given value.

### HasRetentionPolicyType

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) HasRetentionPolicyType() bool`

HasRetentionPolicyType returns a boolean if a field has been set.

### SetRetentionPolicyTypeNil

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) SetRetentionPolicyTypeNil(b bool)`

 SetRetentionPolicyTypeNil sets the value for RetentionPolicyType to be an explicit nil

### UnsetRetentionPolicyType
`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) UnsetRetentionPolicyType()`

UnsetRetentionPolicyType ensures that no value is present for RetentionPolicyType, not even an explicit nil
### GetRetentionPolicyDays

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) GetRetentionPolicyDays() string`

GetRetentionPolicyDays returns the RetentionPolicyDays field if non-nil, zero value otherwise.

### GetRetentionPolicyDaysOk

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) GetRetentionPolicyDaysOk() (*string, bool)`

GetRetentionPolicyDaysOk returns a tuple with the RetentionPolicyDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPolicyDays

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) SetRetentionPolicyDays(v string)`

SetRetentionPolicyDays sets RetentionPolicyDays field to given value.

### HasRetentionPolicyDays

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) HasRetentionPolicyDays() bool`

HasRetentionPolicyDays returns a boolean if a field has been set.

### SetRetentionPolicyDaysNil

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) SetRetentionPolicyDaysNil(b bool)`

 SetRetentionPolicyDaysNil sets the value for RetentionPolicyDays to be an explicit nil

### UnsetRetentionPolicyDays
`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) UnsetRetentionPolicyDays()`

UnsetRetentionPolicyDays ensures that no value is present for RetentionPolicyDays, not even an explicit nil
### GetRetentionProvider

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) GetRetentionProvider() string`

GetRetentionProvider returns the RetentionProvider field if non-nil, zero value otherwise.

### GetRetentionProviderOk

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) GetRetentionProviderOk() (*string, bool)`

GetRetentionProviderOk returns a tuple with the RetentionProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionProvider

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) SetRetentionProvider(v string)`

SetRetentionProvider sets RetentionProvider field to given value.

### HasRetentionProvider

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) HasRetentionProvider() bool`

HasRetentionProvider returns a boolean if a field has been set.

### SetRetentionProviderNil

`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) SetRetentionProviderNil(b bool)`

 SetRetentionProviderNil sets the value for RetentionProvider to be an explicit nil

### UnsetRetentionProvider
`func (o *UpdateStorageBuckets200ResponseAllOfStorageBucket) UnsetRetentionProvider()`

UnsetRetentionProvider ensures that no value is present for RetentionProvider, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


