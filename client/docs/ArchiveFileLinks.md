# ArchiveFileLinks

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**SecretAccessKey** | Pointer to **string** |  | [optional] 
**ArchiveFile** | Pointer to [**GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInnerArchiveFile**](GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInnerArchiveFile.md) |  | [optional] 
**CreatedBy** | Pointer to [**GetArchiveBucket200ResponseArchiveFilesInnerCreatedBy**](GetArchiveBucket200ResponseArchiveFilesInnerCreatedBy.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**LastAccessDate** | Pointer to **time.Time** |  | [optional] 
**ExpirationDate** | Pointer to **time.Time** |  | [optional] 
**DownloadCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewArchiveFileLinks

`func NewArchiveFileLinks() *ArchiveFileLinks`

NewArchiveFileLinks instantiates a new ArchiveFileLinks object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchiveFileLinksWithDefaults

`func NewArchiveFileLinksWithDefaults() *ArchiveFileLinks`

NewArchiveFileLinksWithDefaults instantiates a new ArchiveFileLinks object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ArchiveFileLinks) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArchiveFileLinks) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArchiveFileLinks) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ArchiveFileLinks) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSecretAccessKey

`func (o *ArchiveFileLinks) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *ArchiveFileLinks) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *ArchiveFileLinks) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.

### HasSecretAccessKey

`func (o *ArchiveFileLinks) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.

### GetArchiveFile

`func (o *ArchiveFileLinks) GetArchiveFile() GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInnerArchiveFile`

GetArchiveFile returns the ArchiveFile field if non-nil, zero value otherwise.

### GetArchiveFileOk

`func (o *ArchiveFileLinks) GetArchiveFileOk() (*GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInnerArchiveFile, bool)`

GetArchiveFileOk returns a tuple with the ArchiveFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveFile

`func (o *ArchiveFileLinks) SetArchiveFile(v GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInnerArchiveFile)`

SetArchiveFile sets ArchiveFile field to given value.

### HasArchiveFile

`func (o *ArchiveFileLinks) HasArchiveFile() bool`

HasArchiveFile returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ArchiveFileLinks) GetCreatedBy() GetArchiveBucket200ResponseArchiveFilesInnerCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ArchiveFileLinks) GetCreatedByOk() (*GetArchiveBucket200ResponseArchiveFilesInnerCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ArchiveFileLinks) SetCreatedBy(v GetArchiveBucket200ResponseArchiveFilesInnerCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ArchiveFileLinks) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *ArchiveFileLinks) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ArchiveFileLinks) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ArchiveFileLinks) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ArchiveFileLinks) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ArchiveFileLinks) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ArchiveFileLinks) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ArchiveFileLinks) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ArchiveFileLinks) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastAccessDate

`func (o *ArchiveFileLinks) GetLastAccessDate() time.Time`

GetLastAccessDate returns the LastAccessDate field if non-nil, zero value otherwise.

### GetLastAccessDateOk

`func (o *ArchiveFileLinks) GetLastAccessDateOk() (*time.Time, bool)`

GetLastAccessDateOk returns a tuple with the LastAccessDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessDate

`func (o *ArchiveFileLinks) SetLastAccessDate(v time.Time)`

SetLastAccessDate sets LastAccessDate field to given value.

### HasLastAccessDate

`func (o *ArchiveFileLinks) HasLastAccessDate() bool`

HasLastAccessDate returns a boolean if a field has been set.

### GetExpirationDate

`func (o *ArchiveFileLinks) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *ArchiveFileLinks) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *ArchiveFileLinks) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *ArchiveFileLinks) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetDownloadCount

`func (o *ArchiveFileLinks) GetDownloadCount() int64`

GetDownloadCount returns the DownloadCount field if non-nil, zero value otherwise.

### GetDownloadCountOk

`func (o *ArchiveFileLinks) GetDownloadCountOk() (*int64, bool)`

GetDownloadCountOk returns a tuple with the DownloadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadCount

`func (o *ArchiveFileLinks) SetDownloadCount(v int64)`

SetDownloadCount sets DownloadCount field to given value.

### HasDownloadCount

`func (o *ArchiveFileLinks) HasDownloadCount() bool`

HasDownloadCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


