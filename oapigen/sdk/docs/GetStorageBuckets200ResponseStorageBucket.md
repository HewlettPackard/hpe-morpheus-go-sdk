# GetStorageBuckets200ResponseStorageBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**ProviderType** | Pointer to **string** |  | [optional] 
**StorageServer** | Pointer to [**GetStorageBuckets200ResponseStorageBucketStorageServer**](GetStorageBuckets200ResponseStorageBucketStorageServer.md) |  | [optional] 
**Config** | Pointer to [**GetStorageBuckets200ResponseStorageBucketConfig**](GetStorageBuckets200ResponseStorageBucketConfig.md) |  | [optional] 
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

### NewGetStorageBuckets200ResponseStorageBucket

`func NewGetStorageBuckets200ResponseStorageBucket() *GetStorageBuckets200ResponseStorageBucket`

NewGetStorageBuckets200ResponseStorageBucket instantiates a new GetStorageBuckets200ResponseStorageBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStorageBuckets200ResponseStorageBucketWithDefaults

`func NewGetStorageBuckets200ResponseStorageBucketWithDefaults() *GetStorageBuckets200ResponseStorageBucket`

NewGetStorageBuckets200ResponseStorageBucketWithDefaults instantiates a new GetStorageBuckets200ResponseStorageBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetStorageBuckets200ResponseStorageBucket) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetStorageBuckets200ResponseStorageBucket) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetStorageBuckets200ResponseStorageBucket) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetStorageBuckets200ResponseStorageBucket) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetStorageBuckets200ResponseStorageBucket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetStorageBuckets200ResponseStorageBucket) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetStorageBuckets200ResponseStorageBucket) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetStorageBuckets200ResponseStorageBucket) HasName() bool`

HasName returns a boolean if a field has been set.

### GetActive

`func (o *GetStorageBuckets200ResponseStorageBucket) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetStorageBuckets200ResponseStorageBucket) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetStorageBuckets200ResponseStorageBucket) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetStorageBuckets200ResponseStorageBucket) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetAccountId

`func (o *GetStorageBuckets200ResponseStorageBucket) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetStorageBuckets200ResponseStorageBucket) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetStorageBuckets200ResponseStorageBucket) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetStorageBuckets200ResponseStorageBucket) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetProviderType

`func (o *GetStorageBuckets200ResponseStorageBucket) GetProviderType() string`

GetProviderType returns the ProviderType field if non-nil, zero value otherwise.

### GetProviderTypeOk

`func (o *GetStorageBuckets200ResponseStorageBucket) GetProviderTypeOk() (*string, bool)`

GetProviderTypeOk returns a tuple with the ProviderType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviderType

`func (o *GetStorageBuckets200ResponseStorageBucket) SetProviderType(v string)`

SetProviderType sets ProviderType field to given value.

### HasProviderType

`func (o *GetStorageBuckets200ResponseStorageBucket) HasProviderType() bool`

HasProviderType returns a boolean if a field has been set.

### GetStorageServer

`func (o *GetStorageBuckets200ResponseStorageBucket) GetStorageServer() GetStorageBuckets200ResponseStorageBucketStorageServer`

GetStorageServer returns the StorageServer field if non-nil, zero value otherwise.

### GetStorageServerOk

`func (o *GetStorageBuckets200ResponseStorageBucket) GetStorageServerOk() (*GetStorageBuckets200ResponseStorageBucketStorageServer, bool)`

GetStorageServerOk returns a tuple with the StorageServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageServer

`func (o *GetStorageBuckets200ResponseStorageBucket) SetStorageServer(v GetStorageBuckets200ResponseStorageBucketStorageServer)`

SetStorageServer sets StorageServer field to given value.

### HasStorageServer

`func (o *GetStorageBuckets200ResponseStorageBucket) HasStorageServer() bool`

HasStorageServer returns a boolean if a field has been set.

### GetConfig

`func (o *GetStorageBuckets200ResponseStorageBucket) GetConfig() GetStorageBuckets200ResponseStorageBucketConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetStorageBuckets200ResponseStorageBucket) GetConfigOk() (*GetStorageBuckets200ResponseStorageBucketConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetStorageBuckets200ResponseStorageBucket) SetConfig(v GetStorageBuckets200ResponseStorageBucketConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetStorageBuckets200ResponseStorageBucket) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetBucketName

`func (o *GetStorageBuckets200ResponseStorageBucket) GetBucketName() string`

GetBucketName returns the BucketName field if non-nil, zero value otherwise.

### GetBucketNameOk

`func (o *GetStorageBuckets200ResponseStorageBucket) GetBucketNameOk() (*string, bool)`

GetBucketNameOk returns a tuple with the BucketName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBucketName

`func (o *GetStorageBuckets200ResponseStorageBucket) SetBucketName(v string)`

SetBucketName sets BucketName field to given value.

### HasBucketName

`func (o *GetStorageBuckets200ResponseStorageBucket) HasBucketName() bool`

HasBucketName returns a boolean if a field has been set.

### GetReadOnly

`func (o *GetStorageBuckets200ResponseStorageBucket) GetReadOnly() bool`

GetReadOnly returns the ReadOnly field if non-nil, zero value otherwise.

### GetReadOnlyOk

`func (o *GetStorageBuckets200ResponseStorageBucket) GetReadOnlyOk() (*bool, bool)`

GetReadOnlyOk returns a tuple with the ReadOnly field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnly

`func (o *GetStorageBuckets200ResponseStorageBucket) SetReadOnly(v bool)`

SetReadOnly sets ReadOnly field to given value.

### HasReadOnly

`func (o *GetStorageBuckets200ResponseStorageBucket) HasReadOnly() bool`

HasReadOnly returns a boolean if a field has been set.

### GetDefaultBackupTarget

`func (o *GetStorageBuckets200ResponseStorageBucket) GetDefaultBackupTarget() bool`

GetDefaultBackupTarget returns the DefaultBackupTarget field if non-nil, zero value otherwise.

### GetDefaultBackupTargetOk

`func (o *GetStorageBuckets200ResponseStorageBucket) GetDefaultBackupTargetOk() (*bool, bool)`

GetDefaultBackupTargetOk returns a tuple with the DefaultBackupTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultBackupTarget

`func (o *GetStorageBuckets200ResponseStorageBucket) SetDefaultBackupTarget(v bool)`

SetDefaultBackupTarget sets DefaultBackupTarget field to given value.

### HasDefaultBackupTarget

`func (o *GetStorageBuckets200ResponseStorageBucket) HasDefaultBackupTarget() bool`

HasDefaultBackupTarget returns a boolean if a field has been set.

### GetDefaultDeploymentTarget

`func (o *GetStorageBuckets200ResponseStorageBucket) GetDefaultDeploymentTarget() bool`

GetDefaultDeploymentTarget returns the DefaultDeploymentTarget field if non-nil, zero value otherwise.

### GetDefaultDeploymentTargetOk

`func (o *GetStorageBuckets200ResponseStorageBucket) GetDefaultDeploymentTargetOk() (*bool, bool)`

GetDefaultDeploymentTargetOk returns a tuple with the DefaultDeploymentTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDeploymentTarget

`func (o *GetStorageBuckets200ResponseStorageBucket) SetDefaultDeploymentTarget(v bool)`

SetDefaultDeploymentTarget sets DefaultDeploymentTarget field to given value.

### HasDefaultDeploymentTarget

`func (o *GetStorageBuckets200ResponseStorageBucket) HasDefaultDeploymentTarget() bool`

HasDefaultDeploymentTarget returns a boolean if a field has been set.

### GetDefaultVirtualImageTarget

`func (o *GetStorageBuckets200ResponseStorageBucket) GetDefaultVirtualImageTarget() bool`

GetDefaultVirtualImageTarget returns the DefaultVirtualImageTarget field if non-nil, zero value otherwise.

### GetDefaultVirtualImageTargetOk

`func (o *GetStorageBuckets200ResponseStorageBucket) GetDefaultVirtualImageTargetOk() (*bool, bool)`

GetDefaultVirtualImageTargetOk returns a tuple with the DefaultVirtualImageTarget field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultVirtualImageTarget

`func (o *GetStorageBuckets200ResponseStorageBucket) SetDefaultVirtualImageTarget(v bool)`

SetDefaultVirtualImageTarget sets DefaultVirtualImageTarget field to given value.

### HasDefaultVirtualImageTarget

`func (o *GetStorageBuckets200ResponseStorageBucket) HasDefaultVirtualImageTarget() bool`

HasDefaultVirtualImageTarget returns a boolean if a field has been set.

### GetCopyToStore

`func (o *GetStorageBuckets200ResponseStorageBucket) GetCopyToStore() bool`

GetCopyToStore returns the CopyToStore field if non-nil, zero value otherwise.

### GetCopyToStoreOk

`func (o *GetStorageBuckets200ResponseStorageBucket) GetCopyToStoreOk() (*bool, bool)`

GetCopyToStoreOk returns a tuple with the CopyToStore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopyToStore

`func (o *GetStorageBuckets200ResponseStorageBucket) SetCopyToStore(v bool)`

SetCopyToStore sets CopyToStore field to given value.

### HasCopyToStore

`func (o *GetStorageBuckets200ResponseStorageBucket) HasCopyToStore() bool`

HasCopyToStore returns a boolean if a field has been set.

### GetRetentionPolicyType

`func (o *GetStorageBuckets200ResponseStorageBucket) GetRetentionPolicyType() string`

GetRetentionPolicyType returns the RetentionPolicyType field if non-nil, zero value otherwise.

### GetRetentionPolicyTypeOk

`func (o *GetStorageBuckets200ResponseStorageBucket) GetRetentionPolicyTypeOk() (*string, bool)`

GetRetentionPolicyTypeOk returns a tuple with the RetentionPolicyType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPolicyType

`func (o *GetStorageBuckets200ResponseStorageBucket) SetRetentionPolicyType(v string)`

SetRetentionPolicyType sets RetentionPolicyType field to given value.

### HasRetentionPolicyType

`func (o *GetStorageBuckets200ResponseStorageBucket) HasRetentionPolicyType() bool`

HasRetentionPolicyType returns a boolean if a field has been set.

### SetRetentionPolicyTypeNil

`func (o *GetStorageBuckets200ResponseStorageBucket) SetRetentionPolicyTypeNil(b bool)`

 SetRetentionPolicyTypeNil sets the value for RetentionPolicyType to be an explicit nil

### UnsetRetentionPolicyType
`func (o *GetStorageBuckets200ResponseStorageBucket) UnsetRetentionPolicyType()`

UnsetRetentionPolicyType ensures that no value is present for RetentionPolicyType, not even an explicit nil
### GetRetentionPolicyDays

`func (o *GetStorageBuckets200ResponseStorageBucket) GetRetentionPolicyDays() string`

GetRetentionPolicyDays returns the RetentionPolicyDays field if non-nil, zero value otherwise.

### GetRetentionPolicyDaysOk

`func (o *GetStorageBuckets200ResponseStorageBucket) GetRetentionPolicyDaysOk() (*string, bool)`

GetRetentionPolicyDaysOk returns a tuple with the RetentionPolicyDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionPolicyDays

`func (o *GetStorageBuckets200ResponseStorageBucket) SetRetentionPolicyDays(v string)`

SetRetentionPolicyDays sets RetentionPolicyDays field to given value.

### HasRetentionPolicyDays

`func (o *GetStorageBuckets200ResponseStorageBucket) HasRetentionPolicyDays() bool`

HasRetentionPolicyDays returns a boolean if a field has been set.

### SetRetentionPolicyDaysNil

`func (o *GetStorageBuckets200ResponseStorageBucket) SetRetentionPolicyDaysNil(b bool)`

 SetRetentionPolicyDaysNil sets the value for RetentionPolicyDays to be an explicit nil

### UnsetRetentionPolicyDays
`func (o *GetStorageBuckets200ResponseStorageBucket) UnsetRetentionPolicyDays()`

UnsetRetentionPolicyDays ensures that no value is present for RetentionPolicyDays, not even an explicit nil
### GetRetentionProvider

`func (o *GetStorageBuckets200ResponseStorageBucket) GetRetentionProvider() string`

GetRetentionProvider returns the RetentionProvider field if non-nil, zero value otherwise.

### GetRetentionProviderOk

`func (o *GetStorageBuckets200ResponseStorageBucket) GetRetentionProviderOk() (*string, bool)`

GetRetentionProviderOk returns a tuple with the RetentionProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetentionProvider

`func (o *GetStorageBuckets200ResponseStorageBucket) SetRetentionProvider(v string)`

SetRetentionProvider sets RetentionProvider field to given value.

### HasRetentionProvider

`func (o *GetStorageBuckets200ResponseStorageBucket) HasRetentionProvider() bool`

HasRetentionProvider returns a boolean if a field has been set.

### SetRetentionProviderNil

`func (o *GetStorageBuckets200ResponseStorageBucket) SetRetentionProviderNil(b bool)`

 SetRetentionProviderNil sets the value for RetentionProvider to be an explicit nil

### UnsetRetentionProvider
`func (o *GetStorageBuckets200ResponseStorageBucket) UnsetRetentionProvider()`

UnsetRetentionProvider ensures that no value is present for RetentionProvider, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


