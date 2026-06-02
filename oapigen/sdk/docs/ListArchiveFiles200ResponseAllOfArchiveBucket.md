# ListArchiveFiles200ResponseAllOfArchiveBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**StorageProvider** | Pointer to [**AddArchiveBucket200ResponseAllOfArchiveBucketStorageProvider**](AddArchiveBucket200ResponseAllOfArchiveBucketStorageProvider.md) |  | [optional] 
**Owner** | Pointer to [**AddArchiveBucket200ResponseAllOfArchiveBucketOwner**](AddArchiveBucket200ResponseAllOfArchiveBucketOwner.md) |  | [optional] 
**CreatedBy** | Pointer to [**AddArchiveBucket200ResponseAllOfArchiveBucketCreatedBy**](AddArchiveBucket200ResponseAllOfArchiveBucketCreatedBy.md) |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**FilePath** | Pointer to **string** |  | [optional] 
**RawSize** | Pointer to **NullableInt64** |  | [optional] 
**FileCount** | Pointer to **int64** |  | [optional] 
**Accounts** | Pointer to [**[]AddArchiveBucket200ResponseAllOfArchiveBucketAccountsInner**](AddArchiveBucket200ResponseAllOfArchiveBucketAccountsInner.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListArchiveFiles200ResponseAllOfArchiveBucket

`func NewListArchiveFiles200ResponseAllOfArchiveBucket() *ListArchiveFiles200ResponseAllOfArchiveBucket`

NewListArchiveFiles200ResponseAllOfArchiveBucket instantiates a new ListArchiveFiles200ResponseAllOfArchiveBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStorageProvider

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) GetStorageProvider() AddArchiveBucket200ResponseAllOfArchiveBucketStorageProvider`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) GetStorageProviderOk() (*AddArchiveBucket200ResponseAllOfArchiveBucketStorageProvider, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) SetStorageProvider(v AddArchiveBucket200ResponseAllOfArchiveBucketStorageProvider)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.

### GetOwner

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) GetOwner() AddArchiveBucket200ResponseAllOfArchiveBucketOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) GetOwnerOk() (*AddArchiveBucket200ResponseAllOfArchiveBucketOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) SetOwner(v AddArchiveBucket200ResponseAllOfArchiveBucketOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) GetCreatedBy() AddArchiveBucket200ResponseAllOfArchiveBucketCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) GetCreatedByOk() (*AddArchiveBucket200ResponseAllOfArchiveBucketCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) SetCreatedBy(v AddArchiveBucket200ResponseAllOfArchiveBucketCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetIsPublic

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetVisibility

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetCode

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetFilePath

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetRawSize

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) GetRawSize() int64`

GetRawSize returns the RawSize field if non-nil, zero value otherwise.

### GetRawSizeOk

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) GetRawSizeOk() (*int64, bool)`

GetRawSizeOk returns a tuple with the RawSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawSize

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) SetRawSize(v int64)`

SetRawSize sets RawSize field to given value.

### HasRawSize

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) HasRawSize() bool`

HasRawSize returns a boolean if a field has been set.

### SetRawSizeNil

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) SetRawSizeNil(b bool)`

 SetRawSizeNil sets the value for RawSize to be an explicit nil

### UnsetRawSize
`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) UnsetRawSize()`

UnsetRawSize ensures that no value is present for RawSize, not even an explicit nil
### GetFileCount

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) GetFileCount() int64`

GetFileCount returns the FileCount field if non-nil, zero value otherwise.

### GetFileCountOk

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) GetFileCountOk() (*int64, bool)`

GetFileCountOk returns a tuple with the FileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileCount

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) SetFileCount(v int64)`

SetFileCount sets FileCount field to given value.

### HasFileCount

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) HasFileCount() bool`

HasFileCount returns a boolean if a field has been set.

### GetAccounts

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) GetAccounts() []AddArchiveBucket200ResponseAllOfArchiveBucketAccountsInner`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) GetAccountsOk() (*[]AddArchiveBucket200ResponseAllOfArchiveBucketAccountsInner, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) SetAccounts(v []AddArchiveBucket200ResponseAllOfArchiveBucketAccountsInner)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListArchiveFiles200ResponseAllOfArchiveBucket) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


