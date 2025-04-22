# AddStorageBucketsRequestStorageBucketConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccessKey** | Pointer to **string** | Access Key | [optional] 
**SecretKey** | Pointer to **string** | Secret Key | [optional] 
**Region** | Pointer to **string** | Region | [optional] 
**Endpoint** | Pointer to **string** | Optional endpoint URL if pointing to an object store other than amazon that mimics the Amazon S3 APIs. | [optional] 
**StorageAccount** | Pointer to **string** | Storage Account | [optional] 
**StorageKey** | Pointer to **string** | Storage Key | [optional] 
**Host** | Pointer to **string** | Host | [optional] 
**Username** | Pointer to **string** | username | [optional] 
**Password** | Pointer to **string** | password | [optional] 
**BasePath** | Pointer to **string** | Storage Path | [optional] 
**ExportFolder** | Pointer to **string** | Export Folder | [optional] 
**ApiKey** | Pointer to **string** | API Key | [optional] 
**IdentityUrl** | Pointer to **string** | Identity URL | [optional] 

## Methods

### NewAddStorageBucketsRequestStorageBucketConfig

`func NewAddStorageBucketsRequestStorageBucketConfig() *AddStorageBucketsRequestStorageBucketConfig`

NewAddStorageBucketsRequestStorageBucketConfig instantiates a new AddStorageBucketsRequestStorageBucketConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddStorageBucketsRequestStorageBucketConfigWithDefaults

`func NewAddStorageBucketsRequestStorageBucketConfigWithDefaults() *AddStorageBucketsRequestStorageBucketConfig`

NewAddStorageBucketsRequestStorageBucketConfigWithDefaults instantiates a new AddStorageBucketsRequestStorageBucketConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccessKey

`func (o *AddStorageBucketsRequestStorageBucketConfig) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *AddStorageBucketsRequestStorageBucketConfig) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *AddStorageBucketsRequestStorageBucketConfig) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *AddStorageBucketsRequestStorageBucketConfig) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *AddStorageBucketsRequestStorageBucketConfig) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *AddStorageBucketsRequestStorageBucketConfig) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *AddStorageBucketsRequestStorageBucketConfig) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *AddStorageBucketsRequestStorageBucketConfig) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetRegion

`func (o *AddStorageBucketsRequestStorageBucketConfig) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AddStorageBucketsRequestStorageBucketConfig) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AddStorageBucketsRequestStorageBucketConfig) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AddStorageBucketsRequestStorageBucketConfig) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetEndpoint

`func (o *AddStorageBucketsRequestStorageBucketConfig) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *AddStorageBucketsRequestStorageBucketConfig) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *AddStorageBucketsRequestStorageBucketConfig) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *AddStorageBucketsRequestStorageBucketConfig) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetStorageAccount

`func (o *AddStorageBucketsRequestStorageBucketConfig) GetStorageAccount() string`

GetStorageAccount returns the StorageAccount field if non-nil, zero value otherwise.

### GetStorageAccountOk

`func (o *AddStorageBucketsRequestStorageBucketConfig) GetStorageAccountOk() (*string, bool)`

GetStorageAccountOk returns a tuple with the StorageAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageAccount

`func (o *AddStorageBucketsRequestStorageBucketConfig) SetStorageAccount(v string)`

SetStorageAccount sets StorageAccount field to given value.

### HasStorageAccount

`func (o *AddStorageBucketsRequestStorageBucketConfig) HasStorageAccount() bool`

HasStorageAccount returns a boolean if a field has been set.

### GetStorageKey

`func (o *AddStorageBucketsRequestStorageBucketConfig) GetStorageKey() string`

GetStorageKey returns the StorageKey field if non-nil, zero value otherwise.

### GetStorageKeyOk

`func (o *AddStorageBucketsRequestStorageBucketConfig) GetStorageKeyOk() (*string, bool)`

GetStorageKeyOk returns a tuple with the StorageKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageKey

`func (o *AddStorageBucketsRequestStorageBucketConfig) SetStorageKey(v string)`

SetStorageKey sets StorageKey field to given value.

### HasStorageKey

`func (o *AddStorageBucketsRequestStorageBucketConfig) HasStorageKey() bool`

HasStorageKey returns a boolean if a field has been set.

### GetHost

`func (o *AddStorageBucketsRequestStorageBucketConfig) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *AddStorageBucketsRequestStorageBucketConfig) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *AddStorageBucketsRequestStorageBucketConfig) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *AddStorageBucketsRequestStorageBucketConfig) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetUsername

`func (o *AddStorageBucketsRequestStorageBucketConfig) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AddStorageBucketsRequestStorageBucketConfig) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AddStorageBucketsRequestStorageBucketConfig) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AddStorageBucketsRequestStorageBucketConfig) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *AddStorageBucketsRequestStorageBucketConfig) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AddStorageBucketsRequestStorageBucketConfig) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AddStorageBucketsRequestStorageBucketConfig) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AddStorageBucketsRequestStorageBucketConfig) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetBasePath

`func (o *AddStorageBucketsRequestStorageBucketConfig) GetBasePath() string`

GetBasePath returns the BasePath field if non-nil, zero value otherwise.

### GetBasePathOk

`func (o *AddStorageBucketsRequestStorageBucketConfig) GetBasePathOk() (*string, bool)`

GetBasePathOk returns a tuple with the BasePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBasePath

`func (o *AddStorageBucketsRequestStorageBucketConfig) SetBasePath(v string)`

SetBasePath sets BasePath field to given value.

### HasBasePath

`func (o *AddStorageBucketsRequestStorageBucketConfig) HasBasePath() bool`

HasBasePath returns a boolean if a field has been set.

### GetExportFolder

`func (o *AddStorageBucketsRequestStorageBucketConfig) GetExportFolder() string`

GetExportFolder returns the ExportFolder field if non-nil, zero value otherwise.

### GetExportFolderOk

`func (o *AddStorageBucketsRequestStorageBucketConfig) GetExportFolderOk() (*string, bool)`

GetExportFolderOk returns a tuple with the ExportFolder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExportFolder

`func (o *AddStorageBucketsRequestStorageBucketConfig) SetExportFolder(v string)`

SetExportFolder sets ExportFolder field to given value.

### HasExportFolder

`func (o *AddStorageBucketsRequestStorageBucketConfig) HasExportFolder() bool`

HasExportFolder returns a boolean if a field has been set.

### GetApiKey

`func (o *AddStorageBucketsRequestStorageBucketConfig) GetApiKey() string`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *AddStorageBucketsRequestStorageBucketConfig) GetApiKeyOk() (*string, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *AddStorageBucketsRequestStorageBucketConfig) SetApiKey(v string)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *AddStorageBucketsRequestStorageBucketConfig) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.

### GetIdentityUrl

`func (o *AddStorageBucketsRequestStorageBucketConfig) GetIdentityUrl() string`

GetIdentityUrl returns the IdentityUrl field if non-nil, zero value otherwise.

### GetIdentityUrlOk

`func (o *AddStorageBucketsRequestStorageBucketConfig) GetIdentityUrlOk() (*string, bool)`

GetIdentityUrlOk returns a tuple with the IdentityUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityUrl

`func (o *AddStorageBucketsRequestStorageBucketConfig) SetIdentityUrl(v string)`

SetIdentityUrl sets IdentityUrl field to given value.

### HasIdentityUrl

`func (o *AddStorageBucketsRequestStorageBucketConfig) HasIdentityUrl() bool`

HasIdentityUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


