# AddArchiveBucketRequestArchiveBucket

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

### NewAddArchiveBucketRequestArchiveBucket

`func NewAddArchiveBucketRequestArchiveBucket() *AddArchiveBucketRequestArchiveBucket`

NewAddArchiveBucketRequestArchiveBucket instantiates a new AddArchiveBucketRequestArchiveBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddArchiveBucketRequestArchiveBucketWithDefaults

`func NewAddArchiveBucketRequestArchiveBucketWithDefaults() *AddArchiveBucketRequestArchiveBucket`

NewAddArchiveBucketRequestArchiveBucketWithDefaults instantiates a new AddArchiveBucketRequestArchiveBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddArchiveBucketRequestArchiveBucket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddArchiveBucketRequestArchiveBucket) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddArchiveBucketRequestArchiveBucket) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddArchiveBucketRequestArchiveBucket) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AddArchiveBucketRequestArchiveBucket) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddArchiveBucketRequestArchiveBucket) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddArchiveBucketRequestArchiveBucket) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddArchiveBucketRequestArchiveBucket) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetStorageProvider

`func (o *AddArchiveBucketRequestArchiveBucket) GetStorageProvider() AddArchiveBucketRequestArchiveBucketStorageProvider`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *AddArchiveBucketRequestArchiveBucket) GetStorageProviderOk() (*AddArchiveBucketRequestArchiveBucketStorageProvider, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *AddArchiveBucketRequestArchiveBucket) SetStorageProvider(v AddArchiveBucketRequestArchiveBucketStorageProvider)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *AddArchiveBucketRequestArchiveBucket) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.

### GetVisibility

`func (o *AddArchiveBucketRequestArchiveBucket) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AddArchiveBucketRequestArchiveBucket) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AddArchiveBucketRequestArchiveBucket) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AddArchiveBucketRequestArchiveBucket) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetIsPublic

`func (o *AddArchiveBucketRequestArchiveBucket) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *AddArchiveBucketRequestArchiveBucket) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *AddArchiveBucketRequestArchiveBucket) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *AddArchiveBucketRequestArchiveBucket) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetAccounts

`func (o *AddArchiveBucketRequestArchiveBucket) GetAccounts() GetAlerts200ResponseAllOfChecksInnerAccount`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *AddArchiveBucketRequestArchiveBucket) GetAccountsOk() (*GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *AddArchiveBucketRequestArchiveBucket) SetAccounts(v GetAlerts200ResponseAllOfChecksInnerAccount)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *AddArchiveBucketRequestArchiveBucket) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


