# ArchiveBucketCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A name for the archive bucket. Must be globally unique. | [optional] 
**Description** | Pointer to **string** | A description for the archive bucket | [optional] 
**StorageProvider** | Pointer to [**AddArchiveBucketRequestArchiveBucketStorageProvider**](AddArchiveBucketRequestArchiveBucketStorageProvider.md) |  | [optional] 
**Visibility** | Pointer to **string** | Visibility - Set to public to allow all tenants | [optional] [default to "private"]
**IsPublic** | Pointer to **bool** | Public URL - Set to true to allow anonymous access | [optional] [default to false]
**Accounts** | Pointer to [**GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) |  | [optional] 

## Methods

### NewArchiveBucketCreate

`func NewArchiveBucketCreate() *ArchiveBucketCreate`

NewArchiveBucketCreate instantiates a new ArchiveBucketCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchiveBucketCreateWithDefaults

`func NewArchiveBucketCreateWithDefaults() *ArchiveBucketCreate`

NewArchiveBucketCreateWithDefaults instantiates a new ArchiveBucketCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ArchiveBucketCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArchiveBucketCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArchiveBucketCreate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ArchiveBucketCreate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ArchiveBucketCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ArchiveBucketCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ArchiveBucketCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ArchiveBucketCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStorageProvider

`func (o *ArchiveBucketCreate) GetStorageProvider() AddArchiveBucketRequestArchiveBucketStorageProvider`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *ArchiveBucketCreate) GetStorageProviderOk() (*AddArchiveBucketRequestArchiveBucketStorageProvider, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *ArchiveBucketCreate) SetStorageProvider(v AddArchiveBucketRequestArchiveBucketStorageProvider)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *ArchiveBucketCreate) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.

### GetVisibility

`func (o *ArchiveBucketCreate) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ArchiveBucketCreate) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ArchiveBucketCreate) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ArchiveBucketCreate) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetIsPublic

`func (o *ArchiveBucketCreate) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *ArchiveBucketCreate) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *ArchiveBucketCreate) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *ArchiveBucketCreate) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetAccounts

`func (o *ArchiveBucketCreate) GetAccounts() GetAlerts200ResponseAllOfChecksInnerAccount`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *ArchiveBucketCreate) GetAccountsOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *ArchiveBucketCreate) SetAccounts(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *ArchiveBucketCreate) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


