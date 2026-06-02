# GetArchiveFileDetail200ResponseArchiveFile

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

### NewGetArchiveFileDetail200ResponseArchiveFile

`func NewGetArchiveFileDetail200ResponseArchiveFile() *GetArchiveFileDetail200ResponseArchiveFile`

NewGetArchiveFileDetail200ResponseArchiveFile instantiates a new GetArchiveFileDetail200ResponseArchiveFile object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetArchiveFileDetail200ResponseArchiveFile) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetArchiveFileDetail200ResponseArchiveFile) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetArchiveFileDetail200ResponseArchiveFile) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetArchiveFileDetail200ResponseArchiveFile) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetArchiveFileDetail200ResponseArchiveFile) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetArchiveFileDetail200ResponseArchiveFile) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetArchiveFileDetail200ResponseArchiveFile) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetArchiveFileDetail200ResponseArchiveFile) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFilePath

`func (o *GetArchiveFileDetail200ResponseArchiveFile) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *GetArchiveFileDetail200ResponseArchiveFile) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *GetArchiveFileDetail200ResponseArchiveFile) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *GetArchiveFileDetail200ResponseArchiveFile) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetArchiveBucket

`func (o *GetArchiveFileDetail200ResponseArchiveFile) GetArchiveBucket() AddArchiveFile200ResponseAllOfArchiveFileArchiveBucket`

GetArchiveBucket returns the ArchiveBucket field if non-nil, zero value otherwise.

### GetArchiveBucketOk

`func (o *GetArchiveFileDetail200ResponseArchiveFile) GetArchiveBucketOk() (*AddArchiveFile200ResponseAllOfArchiveFileArchiveBucket, bool)`

GetArchiveBucketOk returns a tuple with the ArchiveBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveBucket

`func (o *GetArchiveFileDetail200ResponseArchiveFile) SetArchiveBucket(v AddArchiveFile200ResponseAllOfArchiveFileArchiveBucket)`

SetArchiveBucket sets ArchiveBucket field to given value.

### HasArchiveBucket

`func (o *GetArchiveFileDetail200ResponseArchiveFile) HasArchiveBucket() bool`

HasArchiveBucket returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetArchiveFileDetail200ResponseArchiveFile) GetCreatedBy() AddArchiveFile200ResponseAllOfArchiveFileCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetArchiveFileDetail200ResponseArchiveFile) GetCreatedByOk() (*AddArchiveFile200ResponseAllOfArchiveFileCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetArchiveFileDetail200ResponseArchiveFile) SetCreatedBy(v AddArchiveFile200ResponseAllOfArchiveFileCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetArchiveFileDetail200ResponseArchiveFile) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetIsDirectory

`func (o *GetArchiveFileDetail200ResponseArchiveFile) GetIsDirectory() bool`

GetIsDirectory returns the IsDirectory field if non-nil, zero value otherwise.

### GetIsDirectoryOk

`func (o *GetArchiveFileDetail200ResponseArchiveFile) GetIsDirectoryOk() (*bool, bool)`

GetIsDirectoryOk returns a tuple with the IsDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDirectory

`func (o *GetArchiveFileDetail200ResponseArchiveFile) SetIsDirectory(v bool)`

SetIsDirectory sets IsDirectory field to given value.

### HasIsDirectory

`func (o *GetArchiveFileDetail200ResponseArchiveFile) HasIsDirectory() bool`

HasIsDirectory returns a boolean if a field has been set.

### GetStatus

`func (o *GetArchiveFileDetail200ResponseArchiveFile) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetArchiveFileDetail200ResponseArchiveFile) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetArchiveFileDetail200ResponseArchiveFile) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetArchiveFileDetail200ResponseArchiveFile) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRawSize

`func (o *GetArchiveFileDetail200ResponseArchiveFile) GetRawSize() int64`

GetRawSize returns the RawSize field if non-nil, zero value otherwise.

### GetRawSizeOk

`func (o *GetArchiveFileDetail200ResponseArchiveFile) GetRawSizeOk() (*int64, bool)`

GetRawSizeOk returns a tuple with the RawSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawSize

`func (o *GetArchiveFileDetail200ResponseArchiveFile) SetRawSize(v int64)`

SetRawSize sets RawSize field to given value.

### HasRawSize

`func (o *GetArchiveFileDetail200ResponseArchiveFile) HasRawSize() bool`

HasRawSize returns a boolean if a field has been set.

### GetContentType

`func (o *GetArchiveFileDetail200ResponseArchiveFile) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *GetArchiveFileDetail200ResponseArchiveFile) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *GetArchiveFileDetail200ResponseArchiveFile) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *GetArchiveFileDetail200ResponseArchiveFile) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *GetArchiveFileDetail200ResponseArchiveFile) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *GetArchiveFileDetail200ResponseArchiveFile) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetDownloadCount

`func (o *GetArchiveFileDetail200ResponseArchiveFile) GetDownloadCount() int64`

GetDownloadCount returns the DownloadCount field if non-nil, zero value otherwise.

### GetDownloadCountOk

`func (o *GetArchiveFileDetail200ResponseArchiveFile) GetDownloadCountOk() (*int64, bool)`

GetDownloadCountOk returns a tuple with the DownloadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadCount

`func (o *GetArchiveFileDetail200ResponseArchiveFile) SetDownloadCount(v int64)`

SetDownloadCount sets DownloadCount field to given value.

### HasDownloadCount

`func (o *GetArchiveFileDetail200ResponseArchiveFile) HasDownloadCount() bool`

HasDownloadCount returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetArchiveFileDetail200ResponseArchiveFile) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetArchiveFileDetail200ResponseArchiveFile) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetArchiveFileDetail200ResponseArchiveFile) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetArchiveFileDetail200ResponseArchiveFile) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetArchiveFileDetail200ResponseArchiveFile) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetArchiveFileDetail200ResponseArchiveFile) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetArchiveFileDetail200ResponseArchiveFile) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetArchiveFileDetail200ResponseArchiveFile) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


