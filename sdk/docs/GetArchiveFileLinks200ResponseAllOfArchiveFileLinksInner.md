# GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner

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

### NewGetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner

`func NewGetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner() *GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner`

NewGetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner instantiates a new GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetArchiveFileLinks200ResponseAllOfArchiveFileLinksInnerWithDefaults

`func NewGetArchiveFileLinks200ResponseAllOfArchiveFileLinksInnerWithDefaults() *GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner`

NewGetArchiveFileLinks200ResponseAllOfArchiveFileLinksInnerWithDefaults instantiates a new GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetSecretAccessKey

`func (o *GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner) GetSecretAccessKey() string`

GetSecretAccessKey returns the SecretAccessKey field if non-nil, zero value otherwise.

### GetSecretAccessKeyOk

`func (o *GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner) GetSecretAccessKeyOk() (*string, bool)`

GetSecretAccessKeyOk returns a tuple with the SecretAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretAccessKey

`func (o *GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner) SetSecretAccessKey(v string)`

SetSecretAccessKey sets SecretAccessKey field to given value.

### HasSecretAccessKey

`func (o *GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner) HasSecretAccessKey() bool`

HasSecretAccessKey returns a boolean if a field has been set.

### GetArchiveFile

`func (o *GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner) GetArchiveFile() GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInnerArchiveFile`

GetArchiveFile returns the ArchiveFile field if non-nil, zero value otherwise.

### GetArchiveFileOk

`func (o *GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner) GetArchiveFileOk() (*GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInnerArchiveFile, bool)`

GetArchiveFileOk returns a tuple with the ArchiveFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveFile

`func (o *GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner) SetArchiveFile(v GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInnerArchiveFile)`

SetArchiveFile sets ArchiveFile field to given value.

### HasArchiveFile

`func (o *GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner) HasArchiveFile() bool`

HasArchiveFile returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner) GetCreatedBy() GetArchiveBucket200ResponseArchiveFilesInnerCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner) GetCreatedByOk() (*GetArchiveBucket200ResponseArchiveFilesInnerCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner) SetCreatedBy(v GetArchiveBucket200ResponseArchiveFilesInnerCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetLastAccessDate

`func (o *GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner) GetLastAccessDate() time.Time`

GetLastAccessDate returns the LastAccessDate field if non-nil, zero value otherwise.

### GetLastAccessDateOk

`func (o *GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner) GetLastAccessDateOk() (*time.Time, bool)`

GetLastAccessDateOk returns a tuple with the LastAccessDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAccessDate

`func (o *GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner) SetLastAccessDate(v time.Time)`

SetLastAccessDate sets LastAccessDate field to given value.

### HasLastAccessDate

`func (o *GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner) HasLastAccessDate() bool`

HasLastAccessDate returns a boolean if a field has been set.

### GetExpirationDate

`func (o *GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner) GetExpirationDate() time.Time`

GetExpirationDate returns the ExpirationDate field if non-nil, zero value otherwise.

### GetExpirationDateOk

`func (o *GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner) GetExpirationDateOk() (*time.Time, bool)`

GetExpirationDateOk returns a tuple with the ExpirationDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationDate

`func (o *GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner) SetExpirationDate(v time.Time)`

SetExpirationDate sets ExpirationDate field to given value.

### HasExpirationDate

`func (o *GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner) HasExpirationDate() bool`

HasExpirationDate returns a boolean if a field has been set.

### GetDownloadCount

`func (o *GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner) GetDownloadCount() int64`

GetDownloadCount returns the DownloadCount field if non-nil, zero value otherwise.

### GetDownloadCountOk

`func (o *GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner) GetDownloadCountOk() (*int64, bool)`

GetDownloadCountOk returns a tuple with the DownloadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadCount

`func (o *GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner) SetDownloadCount(v int64)`

SetDownloadCount sets DownloadCount field to given value.

### HasDownloadCount

`func (o *GetArchiveFileLinks200ResponseAllOfArchiveFileLinksInner) HasDownloadCount() bool`

HasDownloadCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


