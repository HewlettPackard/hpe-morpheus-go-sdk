# GetArchiveBucket200ResponseArchiveFilesInner

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

### NewGetArchiveBucket200ResponseArchiveFilesInner

`func NewGetArchiveBucket200ResponseArchiveFilesInner() *GetArchiveBucket200ResponseArchiveFilesInner`

NewGetArchiveBucket200ResponseArchiveFilesInner instantiates a new GetArchiveBucket200ResponseArchiveFilesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetArchiveBucket200ResponseArchiveFilesInnerWithDefaults

`func NewGetArchiveBucket200ResponseArchiveFilesInnerWithDefaults() *GetArchiveBucket200ResponseArchiveFilesInner`

NewGetArchiveBucket200ResponseArchiveFilesInnerWithDefaults instantiates a new GetArchiveBucket200ResponseArchiveFilesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFilePath

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### GetArchiveBucket

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) GetArchiveBucket() GetArchiveBucket200ResponseArchiveFilesInnerArchiveBucket`

GetArchiveBucket returns the ArchiveBucket field if non-nil, zero value otherwise.

### GetArchiveBucketOk

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) GetArchiveBucketOk() (*GetArchiveBucket200ResponseArchiveFilesInnerArchiveBucket, bool)`

GetArchiveBucketOk returns a tuple with the ArchiveBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveBucket

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) SetArchiveBucket(v GetArchiveBucket200ResponseArchiveFilesInnerArchiveBucket)`

SetArchiveBucket sets ArchiveBucket field to given value.

### HasArchiveBucket

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) HasArchiveBucket() bool`

HasArchiveBucket returns a boolean if a field has been set.

### GetCreatedBy

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) GetCreatedBy() GetArchiveBucket200ResponseArchiveFilesInnerCreatedBy`

GetCreatedBy returns the CreatedBy field if non-nil, zero value otherwise.

### GetCreatedByOk

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) GetCreatedByOk() (*GetArchiveBucket200ResponseArchiveFilesInnerCreatedBy, bool)`

GetCreatedByOk returns a tuple with the CreatedBy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedBy

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) SetCreatedBy(v GetArchiveBucket200ResponseArchiveFilesInnerCreatedBy)`

SetCreatedBy sets CreatedBy field to given value.

### HasCreatedBy

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) HasCreatedBy() bool`

HasCreatedBy returns a boolean if a field has been set.

### GetIsDirectory

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) GetIsDirectory() bool`

GetIsDirectory returns the IsDirectory field if non-nil, zero value otherwise.

### GetIsDirectoryOk

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) GetIsDirectoryOk() (*bool, bool)`

GetIsDirectoryOk returns a tuple with the IsDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsDirectory

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) SetIsDirectory(v bool)`

SetIsDirectory sets IsDirectory field to given value.

### HasIsDirectory

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) HasIsDirectory() bool`

HasIsDirectory returns a boolean if a field has been set.

### GetStatus

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetRawSize

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) GetRawSize() int64`

GetRawSize returns the RawSize field if non-nil, zero value otherwise.

### GetRawSizeOk

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) GetRawSizeOk() (*int64, bool)`

GetRawSizeOk returns a tuple with the RawSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRawSize

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) SetRawSize(v int64)`

SetRawSize sets RawSize field to given value.

### HasRawSize

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) HasRawSize() bool`

HasRawSize returns a boolean if a field has been set.

### GetContentType

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### GetDownloadCount

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) GetDownloadCount() int64`

GetDownloadCount returns the DownloadCount field if non-nil, zero value otherwise.

### GetDownloadCountOk

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) GetDownloadCountOk() (*int64, bool)`

GetDownloadCountOk returns a tuple with the DownloadCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDownloadCount

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) SetDownloadCount(v int64)`

SetDownloadCount sets DownloadCount field to given value.

### HasDownloadCount

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) HasDownloadCount() bool`

HasDownloadCount returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetArchiveBucket200ResponseArchiveFilesInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


