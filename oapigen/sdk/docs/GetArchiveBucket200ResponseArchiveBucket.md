# GetArchiveBucket200ResponseArchiveBucket

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

### NewGetArchiveBucket200ResponseArchiveBucket

`func NewGetArchiveBucket200ResponseArchiveBucket() *GetArchiveBucket200ResponseArchiveBucket`

NewGetArchiveBucket200ResponseArchiveBucket instantiates a new GetArchiveBucket200ResponseArchiveBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetArchiveBucket200ResponseArchiveBucketWithDefaults

`func NewGetArchiveBucket200ResponseArchiveBucketWithDefaults() *GetArchiveBucket200ResponseArchiveBucket`

NewGetArchiveBucket200ResponseArchiveBucketWithDefaults instantiates a new GetArchiveBucket200ResponseArchiveBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetArchiveBucket200ResponseArchiveBucket) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetArchiveBucket200ResponseArchiveBucket) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetArchiveBucket200ResponseArchiveBucket) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetArchiveBucket200ResponseArchiveBucket) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetArchiveBucket200ResponseArchiveBucket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetArchiveBucket200ResponseArchiveBucket) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetArchiveBucket200ResponseArchiveBucket) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetArchiveBucket200ResponseArchiveBucket) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetArchiveBucket200ResponseArchiveBucket) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetArchiveBucket200ResponseArchiveBucket) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetArchiveBucket200ResponseArchiveBucket) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetArchiveBucket200ResponseArchiveBucket) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetArchiveBucket200ResponseArchiveBucket) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetArchiveBucket200ResponseArchiveBucket) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStorageProvider

`func (o *GetArchiveBucket200ResponseArchiveBucket) GetStorageProvider() AddArchiveBucket200ResponseAllOfArchiveBucketStorageProvider`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *GetArchiveBucket200ResponseArchiveBucket) GetStorageProviderOk() (*AddArchiveBucket200ResponseAllOfArchiveBucketStorageProvider, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *GetArchiveBucket200ResponseArchiveBucket) SetStorageProvider(v AddArchiveBucket200ResponseAllOfArchiveBucketStorageProvider)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *GetArchiveBucket200ResponseArchiveBucket) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.

### GetOwner

`func (o *GetArchiveBucket200ResponseArchiveBucket) GetOwner() AddArchiveBucket200ResponseAllOfArchiveBucketOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *GetArchiveBucket200ResponseArchiveBucket) GetOwnerOk() (*AddArchiveBucket200ResponseAllOfArchiveBucketOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *GetArchiveBucket200ResponseArchiveBucket) SetOwner(v AddArchiveBucket200ResponseAllOfArchiveBucketOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *GetArchiveBucket200ResponseArchiveBucket) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetArchiveBucket200ResponseArchiveBucket) GetCreatedBy() AddArchiveBucket200ResponseAllOfArchiveBucketCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetArchiveBucket200ResponseArchiveBucket) GetCreatedByOk() (*AddArchiveBucket200ResponseAllOfArchiveBucketCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetArchiveBucket200ResponseArchiveBucket) SetCreatedBy(v AddArchiveBucket200ResponseAllOfArchiveBucketCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetArchiveBucket200ResponseArchiveBucket) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetIsPublic

`func (o *GetArchiveBucket200ResponseArchiveBucket) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *GetArchiveBucket200ResponseArchiveBucket) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *GetArchiveBucket200ResponseArchiveBucket) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *GetArchiveBucket200ResponseArchiveBucket) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetVisibility

`func (o *GetArchiveBucket200ResponseArchiveBucket) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetArchiveBucket200ResponseArchiveBucket) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetArchiveBucket200ResponseArchiveBucket) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetArchiveBucket200ResponseArchiveBucket) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetCode

`func (o *GetArchiveBucket200ResponseArchiveBucket) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetArchiveBucket200ResponseArchiveBucket) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetArchiveBucket200ResponseArchiveBucket) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetArchiveBucket200ResponseArchiveBucket) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetFilePath

`func (o *GetArchiveBucket200ResponseArchiveBucket) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *GetArchiveBucket200ResponseArchiveBucket) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *GetArchiveBucket200ResponseArchiveBucket) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *GetArchiveBucket200ResponseArchiveBucket) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetRawSize

`func (o *GetArchiveBucket200ResponseArchiveBucket) GetRawSize() int64`

GetRawSize returns the RawSize field if non-nil, zero value otherwise.

### GetRawSizeOk

`func (o *GetArchiveBucket200ResponseArchiveBucket) GetRawSizeOk() (*int64, bool)`

GetRawSizeOk returns a tuple with the RawSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawSize

`func (o *GetArchiveBucket200ResponseArchiveBucket) SetRawSize(v int64)`

SetRawSize sets RawSize field to given value.

### HasRawSize

`func (o *GetArchiveBucket200ResponseArchiveBucket) HasRawSize() bool`

HasRawSize returns a boolean if a field has been set.

### SetRawSizeNil

`func (o *GetArchiveBucket200ResponseArchiveBucket) SetRawSizeNil(b bool)`

 SetRawSizeNil sets the value for RawSize to be an explicit nil

### UnsetRawSize
`func (o *GetArchiveBucket200ResponseArchiveBucket) UnsetRawSize()`

UnsetRawSize ensures that no value is present for RawSize, not even an explicit nil
### GetFileCount

`func (o *GetArchiveBucket200ResponseArchiveBucket) GetFileCount() int64`

GetFileCount returns the FileCount field if non-nil, zero value otherwise.

### GetFileCountOk

`func (o *GetArchiveBucket200ResponseArchiveBucket) GetFileCountOk() (*int64, bool)`

GetFileCountOk returns a tuple with the FileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileCount

`func (o *GetArchiveBucket200ResponseArchiveBucket) SetFileCount(v int64)`

SetFileCount sets FileCount field to given value.

### HasFileCount

`func (o *GetArchiveBucket200ResponseArchiveBucket) HasFileCount() bool`

HasFileCount returns a boolean if a field has been set.

### GetAccounts

`func (o *GetArchiveBucket200ResponseArchiveBucket) GetAccounts() []AddArchiveBucket200ResponseAllOfArchiveBucketAccountsInner`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *GetArchiveBucket200ResponseArchiveBucket) GetAccountsOk() (*[]AddArchiveBucket200ResponseAllOfArchiveBucketAccountsInner, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *GetArchiveBucket200ResponseArchiveBucket) SetAccounts(v []AddArchiveBucket200ResponseAllOfArchiveBucketAccountsInner)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *GetArchiveBucket200ResponseArchiveBucket) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetArchiveBucket200ResponseArchiveBucket) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetArchiveBucket200ResponseArchiveBucket) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetArchiveBucket200ResponseArchiveBucket) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetArchiveBucket200ResponseArchiveBucket) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetArchiveBucket200ResponseArchiveBucket) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetArchiveBucket200ResponseArchiveBucket) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetArchiveBucket200ResponseArchiveBucket) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetArchiveBucket200ResponseArchiveBucket) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


