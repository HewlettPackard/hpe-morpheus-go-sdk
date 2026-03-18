# ListArchiveFiles200ResponseAllOfArchiveFilesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**FilePath** | Pointer to **string** |  | [optional] 
**ArchiveBucket** | Pointer to [**ListArchiveFiles200ResponseAllOfArchiveFilesInnerArchiveBucket**](ListArchiveFiles200ResponseAllOfArchiveFilesInnerArchiveBucket.md) |  | [optional] 
**CreatedBy** | Pointer to [**ListArchiveFiles200ResponseAllOfArchiveFilesInnerCreatedBy**](ListArchiveFiles200ResponseAllOfArchiveFilesInnerCreatedBy.md) |  | [optional] 
**IsDirectory** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**RawSize** | Pointer to **int64** |  | [optional] 
**ContentType** | Pointer to **NullableString** |  | [optional] 
**DownloadCount** | Pointer to **int64** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListArchiveFiles200ResponseAllOfArchiveFilesInner

`func NewListArchiveFiles200ResponseAllOfArchiveFilesInner() *ListArchiveFiles200ResponseAllOfArchiveFilesInner`

NewListArchiveFiles200ResponseAllOfArchiveFilesInner instantiates a new ListArchiveFiles200ResponseAllOfArchiveFilesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListArchiveFiles200ResponseAllOfArchiveFilesInnerWithDefaults

`func NewListArchiveFiles200ResponseAllOfArchiveFilesInnerWithDefaults() *ListArchiveFiles200ResponseAllOfArchiveFilesInner`

NewListArchiveFiles200ResponseAllOfArchiveFilesInnerWithDefaults instantiates a new ListArchiveFiles200ResponseAllOfArchiveFilesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFilePath

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetArchiveBucket

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) GetArchiveBucket() ListArchiveFiles200ResponseAllOfArchiveFilesInnerArchiveBucket`

GetArchiveBucket returns the ArchiveBucket field if non-nil, zero value otherwise.

### GetArchiveBucketOk

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) GetArchiveBucketOk() (*ListArchiveFiles200ResponseAllOfArchiveFilesInnerArchiveBucket, bool)`

GetArchiveBucketOk returns a tuple with the ArchiveBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveBucket

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) SetArchiveBucket(v ListArchiveFiles200ResponseAllOfArchiveFilesInnerArchiveBucket)`

SetArchiveBucket sets ArchiveBucket field to given value.

### HasArchiveBucket

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) HasArchiveBucket() bool`

HasArchiveBucket returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) GetCreatedBy() ListArchiveFiles200ResponseAllOfArchiveFilesInnerCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) GetCreatedByOk() (*ListArchiveFiles200ResponseAllOfArchiveFilesInnerCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) SetCreatedBy(v ListArchiveFiles200ResponseAllOfArchiveFilesInnerCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetIsDirectory

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) GetIsDirectory() bool`

GetIsDirectory returns the IsDirectory field if non-nil, zero value otherwise.

### GetIsDirectoryOk

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) GetIsDirectoryOk() (*bool, bool)`

GetIsDirectoryOk returns a tuple with the IsDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDirectory

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) SetIsDirectory(v bool)`

SetIsDirectory sets IsDirectory field to given value.

### HasIsDirectory

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) HasIsDirectory() bool`

HasIsDirectory returns a boolean if a field has been set.

### GetStatus

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRawSize

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) GetRawSize() int64`

GetRawSize returns the RawSize field if non-nil, zero value otherwise.

### GetRawSizeOk

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) GetRawSizeOk() (*int64, bool)`

GetRawSizeOk returns a tuple with the RawSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawSize

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) SetRawSize(v int64)`

SetRawSize sets RawSize field to given value.

### HasRawSize

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) HasRawSize() bool`

HasRawSize returns a boolean if a field has been set.

### GetContentType

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetDownloadCount

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) GetDownloadCount() int64`

GetDownloadCount returns the DownloadCount field if non-nil, zero value otherwise.

### GetDownloadCountOk

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) GetDownloadCountOk() (*int64, bool)`

GetDownloadCountOk returns a tuple with the DownloadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadCount

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) SetDownloadCount(v int64)`

SetDownloadCount sets DownloadCount field to given value.

### HasDownloadCount

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) HasDownloadCount() bool`

HasDownloadCount returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListArchiveFiles200ResponseAllOfArchiveFilesInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


