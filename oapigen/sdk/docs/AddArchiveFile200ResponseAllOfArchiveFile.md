# AddArchiveFile200ResponseAllOfArchiveFile

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**FilePath** | Pointer to **string** |  | [optional] 
**ArchiveBucket** | Pointer to [**AddArchiveFile200ResponseAllOfArchiveFileArchiveBucket**](AddArchiveFile200ResponseAllOfArchiveFileArchiveBucket.md) |  | [optional] 
**CreatedBy** | Pointer to [**AddArchiveFile200ResponseAllOfArchiveFileCreatedBy**](AddArchiveFile200ResponseAllOfArchiveFileCreatedBy.md) |  | [optional] 
**IsDirectory** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**RawSize** | Pointer to **int64** |  | [optional] 
**ContentType** | Pointer to **NullableString** |  | [optional] 
**DownloadCount** | Pointer to **int64** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewAddArchiveFile200ResponseAllOfArchiveFile

`func NewAddArchiveFile200ResponseAllOfArchiveFile() *AddArchiveFile200ResponseAllOfArchiveFile`

NewAddArchiveFile200ResponseAllOfArchiveFile instantiates a new AddArchiveFile200ResponseAllOfArchiveFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFilePath

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetArchiveBucket

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) GetArchiveBucket() AddArchiveFile200ResponseAllOfArchiveFileArchiveBucket`

GetArchiveBucket returns the ArchiveBucket field if non-nil, zero value otherwise.

### GetArchiveBucketOk

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) GetArchiveBucketOk() (*AddArchiveFile200ResponseAllOfArchiveFileArchiveBucket, bool)`

GetArchiveBucketOk returns a tuple with the ArchiveBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveBucket

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) SetArchiveBucket(v AddArchiveFile200ResponseAllOfArchiveFileArchiveBucket)`

SetArchiveBucket sets ArchiveBucket field to given value.

### HasArchiveBucket

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) HasArchiveBucket() bool`

HasArchiveBucket returns a boolean if a field has been set.

### GetCreatedBy

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) GetCreatedBy() AddArchiveFile200ResponseAllOfArchiveFileCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) GetCreatedByOk() (*AddArchiveFile200ResponseAllOfArchiveFileCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) SetCreatedBy(v AddArchiveFile200ResponseAllOfArchiveFileCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetIsDirectory

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) GetIsDirectory() bool`

GetIsDirectory returns the IsDirectory field if non-nil, zero value otherwise.

### GetIsDirectoryOk

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) GetIsDirectoryOk() (*bool, bool)`

GetIsDirectoryOk returns a tuple with the IsDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDirectory

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) SetIsDirectory(v bool)`

SetIsDirectory sets IsDirectory field to given value.

### HasIsDirectory

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) HasIsDirectory() bool`

HasIsDirectory returns a boolean if a field has been set.

### GetStatus

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRawSize

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) GetRawSize() int64`

GetRawSize returns the RawSize field if non-nil, zero value otherwise.

### GetRawSizeOk

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) GetRawSizeOk() (*int64, bool)`

GetRawSizeOk returns a tuple with the RawSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawSize

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) SetRawSize(v int64)`

SetRawSize sets RawSize field to given value.

### HasRawSize

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) HasRawSize() bool`

HasRawSize returns a boolean if a field has been set.

### GetContentType

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *AddArchiveFile200ResponseAllOfArchiveFile) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetDownloadCount

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) GetDownloadCount() int64`

GetDownloadCount returns the DownloadCount field if non-nil, zero value otherwise.

### GetDownloadCountOk

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) GetDownloadCountOk() (*int64, bool)`

GetDownloadCountOk returns a tuple with the DownloadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadCount

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) SetDownloadCount(v int64)`

SetDownloadCount sets DownloadCount field to given value.

### HasDownloadCount

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) HasDownloadCount() bool`

HasDownloadCount returns a boolean if a field has been set.

### GetDateCreated

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AddArchiveFile200ResponseAllOfArchiveFile) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


