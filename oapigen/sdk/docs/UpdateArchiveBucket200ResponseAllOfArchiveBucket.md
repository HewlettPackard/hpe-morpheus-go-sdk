# UpdateArchiveBucket200ResponseAllOfArchiveBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**StorageProvider** | Pointer to [**UpdateArchiveBucket200ResponseAllOfArchiveBucketStorageProvider**](UpdateArchiveBucket200ResponseAllOfArchiveBucketStorageProvider.md) |  | [optional] 
**Owner** | Pointer to [**UpdateArchiveBucket200ResponseAllOfArchiveBucketOwner**](UpdateArchiveBucket200ResponseAllOfArchiveBucketOwner.md) |  | [optional] 
**CreatedBy** | Pointer to [**UpdateArchiveBucket200ResponseAllOfArchiveBucketCreatedBy**](UpdateArchiveBucket200ResponseAllOfArchiveBucketCreatedBy.md) |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**FilePath** | Pointer to **string** |  | [optional] 
**RawSize** | Pointer to **NullableInt64** |  | [optional] 
**FileCount** | Pointer to **int64** |  | [optional] 
**Accounts** | Pointer to [**[]UpdateArchiveBucket200ResponseAllOfArchiveBucketAccountsInner**](UpdateArchiveBucket200ResponseAllOfArchiveBucketAccountsInner.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUpdateArchiveBucket200ResponseAllOfArchiveBucket

`func NewUpdateArchiveBucket200ResponseAllOfArchiveBucket() *UpdateArchiveBucket200ResponseAllOfArchiveBucket`

NewUpdateArchiveBucket200ResponseAllOfArchiveBucket instantiates a new UpdateArchiveBucket200ResponseAllOfArchiveBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateArchiveBucket200ResponseAllOfArchiveBucketWithDefaults

`func NewUpdateArchiveBucket200ResponseAllOfArchiveBucketWithDefaults() *UpdateArchiveBucket200ResponseAllOfArchiveBucket`

NewUpdateArchiveBucket200ResponseAllOfArchiveBucketWithDefaults instantiates a new UpdateArchiveBucket200ResponseAllOfArchiveBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStorageProvider

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) GetStorageProvider() UpdateArchiveBucket200ResponseAllOfArchiveBucketStorageProvider`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) GetStorageProviderOk() (*UpdateArchiveBucket200ResponseAllOfArchiveBucketStorageProvider, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) SetStorageProvider(v UpdateArchiveBucket200ResponseAllOfArchiveBucketStorageProvider)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.

### GetOwner

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) GetOwner() UpdateArchiveBucket200ResponseAllOfArchiveBucketOwner`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) GetOwnerOk() (*UpdateArchiveBucket200ResponseAllOfArchiveBucketOwner, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) SetOwner(v UpdateArchiveBucket200ResponseAllOfArchiveBucketOwner)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetCreatedBy

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) GetCreatedBy() UpdateArchiveBucket200ResponseAllOfArchiveBucketCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) GetCreatedByOk() (*UpdateArchiveBucket200ResponseAllOfArchiveBucketCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) SetCreatedBy(v UpdateArchiveBucket200ResponseAllOfArchiveBucketCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetIsPublic

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetCode

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetFilePath

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetRawSize

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) GetRawSize() int64`

GetRawSize returns the RawSize field if non-nil, zero value otherwise.

### GetRawSizeOk

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) GetRawSizeOk() (*int64, bool)`

GetRawSizeOk returns a tuple with the RawSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawSize

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) SetRawSize(v int64)`

SetRawSize sets RawSize field to given value.

### HasRawSize

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) HasRawSize() bool`

HasRawSize returns a boolean if a field has been set.

### SetRawSizeNil

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) SetRawSizeNil(b bool)`

 SetRawSizeNil sets the value for RawSize to be an explicit nil

### UnsetRawSize
`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) UnsetRawSize()`

UnsetRawSize ensures that no value is present for RawSize, not even an explicit nil
### GetFileCount

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) GetFileCount() int64`

GetFileCount returns the FileCount field if non-nil, zero value otherwise.

### GetFileCountOk

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) GetFileCountOk() (*int64, bool)`

GetFileCountOk returns a tuple with the FileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileCount

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) SetFileCount(v int64)`

SetFileCount sets FileCount field to given value.

### HasFileCount

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) HasFileCount() bool`

HasFileCount returns a boolean if a field has been set.

### GetAccounts

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) GetAccounts() []UpdateArchiveBucket200ResponseAllOfArchiveBucketAccountsInner`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) GetAccountsOk() (*[]UpdateArchiveBucket200ResponseAllOfArchiveBucketAccountsInner, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) SetAccounts(v []UpdateArchiveBucket200ResponseAllOfArchiveBucketAccountsInner)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetDateCreated

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateArchiveBucket200ResponseAllOfArchiveBucket) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


