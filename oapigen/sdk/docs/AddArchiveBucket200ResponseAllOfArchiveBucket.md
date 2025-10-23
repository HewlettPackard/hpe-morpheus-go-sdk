# AddArchiveBucket200ResponseAllOfArchiveBucket

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**StorageProvider** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Owner** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**CreatedBy** | Pointer to [**ListArchiveBuckets200ResponseAllOfArchiveBucketsInnerCreatedBy**](ListArchiveBuckets200ResponseAllOfArchiveBucketsInnerCreatedBy.md) |  | [optional] 
**IsPublic** | Pointer to **bool** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**FilePath** | Pointer to **string** |  | [optional] 
**RawSize** | Pointer to **NullableInt64** |  | [optional] 
**FileCount** | Pointer to **int64** |  | [optional] 
**Accounts** | Pointer to [**[]GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAddArchiveBucket200ResponseAllOfArchiveBucket

`func NewAddArchiveBucket200ResponseAllOfArchiveBucket() *AddArchiveBucket200ResponseAllOfArchiveBucket`

NewAddArchiveBucket200ResponseAllOfArchiveBucket instantiates a new AddArchiveBucket200ResponseAllOfArchiveBucket object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddArchiveBucket200ResponseAllOfArchiveBucketWithDefaults

`func NewAddArchiveBucket200ResponseAllOfArchiveBucketWithDefaults() *AddArchiveBucket200ResponseAllOfArchiveBucket`

NewAddArchiveBucket200ResponseAllOfArchiveBucketWithDefaults instantiates a new AddArchiveBucket200ResponseAllOfArchiveBucket object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetStorageProvider

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) GetStorageProvider() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) GetStorageProviderOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) SetStorageProvider(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.

### GetOwner

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) GetOwner() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetOwner returns the Owner field if non-nil, zero value otherwise.

### GetOwnerOk

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) GetOwnerOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetOwnerOk returns a tuple with the Owner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOwner

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) SetOwner(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetOwner sets Owner field to given value.

### HasOwner

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) HasOwner() bool`

HasOwner returns a boolean if a field has been set.

### GetCreatedBy

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) GetCreatedBy() ListArchiveBuckets200ResponseAllOfArchiveBucketsInnerCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) GetCreatedByOk() (*ListArchiveBuckets200ResponseAllOfArchiveBucketsInnerCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) SetCreatedBy(v ListArchiveBuckets200ResponseAllOfArchiveBucketsInnerCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetIsPublic

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) GetIsPublic() bool`

GetIsPublic returns the IsPublic field if non-nil, zero value otherwise.

### GetIsPublicOk

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) GetIsPublicOk() (*bool, bool)`

GetIsPublicOk returns a tuple with the IsPublic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPublic

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) SetIsPublic(v bool)`

SetIsPublic sets IsPublic field to given value.

### HasIsPublic

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) HasIsPublic() bool`

HasIsPublic returns a boolean if a field has been set.

### GetVisibility

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetCode

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetFilePath

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetRawSize

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) GetRawSize() int64`

GetRawSize returns the RawSize field if non-nil, zero value otherwise.

### GetRawSizeOk

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) GetRawSizeOk() (*int64, bool)`

GetRawSizeOk returns a tuple with the RawSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawSize

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) SetRawSize(v int64)`

SetRawSize sets RawSize field to given value.

### HasRawSize

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) HasRawSize() bool`

HasRawSize returns a boolean if a field has been set.

### SetRawSizeNil

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) SetRawSizeNil(b bool)`

 SetRawSizeNil sets the value for RawSize to be an explicit nil

### UnsetRawSize
`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) UnsetRawSize()`

UnsetRawSize ensures that no value is present for RawSize, not even an explicit nil
### GetFileCount

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) GetFileCount() int64`

GetFileCount returns the FileCount field if non-nil, zero value otherwise.

### GetFileCountOk

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) GetFileCountOk() (*int64, bool)`

GetFileCountOk returns a tuple with the FileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFileCount

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) SetFileCount(v int64)`

SetFileCount sets FileCount field to given value.

### HasFileCount

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) HasFileCount() bool`

HasFileCount returns a boolean if a field has been set.

### GetAccounts

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) GetAccounts() []GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) GetAccountsOk() (*[]GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) SetAccounts(v []GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetDateCreated

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AddArchiveBucket200ResponseAllOfArchiveBucket) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


