# ArchiveBucketFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**FilePath** | Pointer to **string** |  | [optional] 
**ArchiveBucket** | Pointer to [**GetArchiveBucket200ResponseArchiveFilesInnerArchiveBucket**](GetArchiveBucket200ResponseArchiveFilesInnerArchiveBucket.md) |  | [optional] 
**CreatedBy** | Pointer to [**GetArchiveBucket200ResponseArchiveFilesInnerCreatedBy**](GetArchiveBucket200ResponseArchiveFilesInnerCreatedBy.md) |  | [optional] 
**IsDirectory** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**RawSize** | Pointer to **int64** |  | [optional] 
**ContentType** | Pointer to **string** |  | [optional] 
**DownloadCount** | Pointer to **int64** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewArchiveBucketFile

`func NewArchiveBucketFile() *ArchiveBucketFile`

NewArchiveBucketFile instantiates a new ArchiveBucketFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewArchiveBucketFileWithDefaults

`func NewArchiveBucketFileWithDefaults() *ArchiveBucketFile`

NewArchiveBucketFileWithDefaults instantiates a new ArchiveBucketFile object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ArchiveBucketFile) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ArchiveBucketFile) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ArchiveBucketFile) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ArchiveBucketFile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ArchiveBucketFile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ArchiveBucketFile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ArchiveBucketFile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ArchiveBucketFile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFilePath

`func (o *ArchiveBucketFile) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *ArchiveBucketFile) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *ArchiveBucketFile) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *ArchiveBucketFile) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetArchiveBucket

`func (o *ArchiveBucketFile) GetArchiveBucket() GetArchiveBucket200ResponseArchiveFilesInnerArchiveBucket`

GetArchiveBucket returns the ArchiveBucket field if non-nil, zero value otherwise.

### GetArchiveBucketOk

`func (o *ArchiveBucketFile) GetArchiveBucketOk() (*GetArchiveBucket200ResponseArchiveFilesInnerArchiveBucket, bool)`

GetArchiveBucketOk returns a tuple with the ArchiveBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveBucket

`func (o *ArchiveBucketFile) SetArchiveBucket(v GetArchiveBucket200ResponseArchiveFilesInnerArchiveBucket)`

SetArchiveBucket sets ArchiveBucket field to given value.

### HasArchiveBucket

`func (o *ArchiveBucketFile) HasArchiveBucket() bool`

HasArchiveBucket returns a boolean if a field has been set.

### GetCreatedBy

`func (o *ArchiveBucketFile) GetCreatedBy() GetArchiveBucket200ResponseArchiveFilesInnerCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *ArchiveBucketFile) GetCreatedByOk() (*GetArchiveBucket200ResponseArchiveFilesInnerCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *ArchiveBucketFile) SetCreatedBy(v GetArchiveBucket200ResponseArchiveFilesInnerCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *ArchiveBucketFile) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetIsDirectory

`func (o *ArchiveBucketFile) GetIsDirectory() bool`

GetIsDirectory returns the IsDirectory field if non-nil, zero value otherwise.

### GetIsDirectoryOk

`func (o *ArchiveBucketFile) GetIsDirectoryOk() (*bool, bool)`

GetIsDirectoryOk returns a tuple with the IsDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDirectory

`func (o *ArchiveBucketFile) SetIsDirectory(v bool)`

SetIsDirectory sets IsDirectory field to given value.

### HasIsDirectory

`func (o *ArchiveBucketFile) HasIsDirectory() bool`

HasIsDirectory returns a boolean if a field has been set.

### GetStatus

`func (o *ArchiveBucketFile) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ArchiveBucketFile) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ArchiveBucketFile) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ArchiveBucketFile) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRawSize

`func (o *ArchiveBucketFile) GetRawSize() int64`

GetRawSize returns the RawSize field if non-nil, zero value otherwise.

### GetRawSizeOk

`func (o *ArchiveBucketFile) GetRawSizeOk() (*int64, bool)`

GetRawSizeOk returns a tuple with the RawSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawSize

`func (o *ArchiveBucketFile) SetRawSize(v int64)`

SetRawSize sets RawSize field to given value.

### HasRawSize

`func (o *ArchiveBucketFile) HasRawSize() bool`

HasRawSize returns a boolean if a field has been set.

### GetContentType

`func (o *ArchiveBucketFile) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ArchiveBucketFile) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ArchiveBucketFile) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *ArchiveBucketFile) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetDownloadCount

`func (o *ArchiveBucketFile) GetDownloadCount() int64`

GetDownloadCount returns the DownloadCount field if non-nil, zero value otherwise.

### GetDownloadCountOk

`func (o *ArchiveBucketFile) GetDownloadCountOk() (*int64, bool)`

GetDownloadCountOk returns a tuple with the DownloadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadCount

`func (o *ArchiveBucketFile) SetDownloadCount(v int64)`

SetDownloadCount sets DownloadCount field to given value.

### HasDownloadCount

`func (o *ArchiveBucketFile) HasDownloadCount() bool`

HasDownloadCount returns a boolean if a field has been set.

### GetDateCreated

`func (o *ArchiveBucketFile) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ArchiveBucketFile) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ArchiveBucketFile) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ArchiveBucketFile) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ArchiveBucketFile) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ArchiveBucketFile) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ArchiveBucketFile) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ArchiveBucketFile) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


