# GetArchiveBucket200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveBucket** | Pointer to [**ListArchiveBuckets200ResponseAllOfArchiveBucketsInner**](ListArchiveBuckets200ResponseAllOfArchiveBucketsInner.md) |  | [optional] 
**IsOwner** | Pointer to **bool** |  | [optional] 
**ParentDirectory** | Pointer to **string** |  | [optional] 
**ArchiveFiles** | Pointer to [**[]GetArchiveBucket200ResponseArchiveFilesInner**](GetArchiveBucket200ResponseArchiveFilesInner.md) |  | [optional] 
**ArchiveFileCount** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetArchiveBucket200Response

`func NewGetArchiveBucket200Response() *GetArchiveBucket200Response`

NewGetArchiveBucket200Response instantiates a new GetArchiveBucket200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetArchiveBucket200ResponseWithDefaults

`func NewGetArchiveBucket200ResponseWithDefaults() *GetArchiveBucket200Response`

NewGetArchiveBucket200ResponseWithDefaults instantiates a new GetArchiveBucket200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveBucket

`func (o *GetArchiveBucket200Response) GetArchiveBucket() ListArchiveBuckets200ResponseAllOfArchiveBucketsInner`

GetArchiveBucket returns the ArchiveBucket field if non-nil, zero value otherwise.

### GetArchiveBucketOk

`func (o *GetArchiveBucket200Response) GetArchiveBucketOk() (*ListArchiveBuckets200ResponseAllOfArchiveBucketsInner, bool)`

GetArchiveBucketOk returns a tuple with the ArchiveBucket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveBucket

`func (o *GetArchiveBucket200Response) SetArchiveBucket(v ListArchiveBuckets200ResponseAllOfArchiveBucketsInner)`

SetArchiveBucket sets ArchiveBucket field to given value.

### HasArchiveBucket

`func (o *GetArchiveBucket200Response) HasArchiveBucket() bool`

HasArchiveBucket returns a boolean if a field has been set.

### GetIsOwner

`func (o *GetArchiveBucket200Response) GetIsOwner() bool`

GetIsOwner returns the IsOwner field if non-nil, zero value otherwise.

### GetIsOwnerOk

`func (o *GetArchiveBucket200Response) GetIsOwnerOk() (*bool, bool)`

GetIsOwnerOk returns a tuple with the IsOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOwner

`func (o *GetArchiveBucket200Response) SetIsOwner(v bool)`

SetIsOwner sets IsOwner field to given value.

### HasIsOwner

`func (o *GetArchiveBucket200Response) HasIsOwner() bool`

HasIsOwner returns a boolean if a field has been set.

### GetParentDirectory

`func (o *GetArchiveBucket200Response) GetParentDirectory() string`

GetParentDirectory returns the ParentDirectory field if non-nil, zero value otherwise.

### GetParentDirectoryOk

`func (o *GetArchiveBucket200Response) GetParentDirectoryOk() (*string, bool)`

GetParentDirectoryOk returns a tuple with the ParentDirectory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentDirectory

`func (o *GetArchiveBucket200Response) SetParentDirectory(v string)`

SetParentDirectory sets ParentDirectory field to given value.

### HasParentDirectory

`func (o *GetArchiveBucket200Response) HasParentDirectory() bool`

HasParentDirectory returns a boolean if a field has been set.

### GetArchiveFiles

`func (o *GetArchiveBucket200Response) GetArchiveFiles() []GetArchiveBucket200ResponseArchiveFilesInner`

GetArchiveFiles returns the ArchiveFiles field if non-nil, zero value otherwise.

### GetArchiveFilesOk

`func (o *GetArchiveBucket200Response) GetArchiveFilesOk() (*[]GetArchiveBucket200ResponseArchiveFilesInner, bool)`

GetArchiveFilesOk returns a tuple with the ArchiveFiles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveFiles

`func (o *GetArchiveBucket200Response) SetArchiveFiles(v []GetArchiveBucket200ResponseArchiveFilesInner)`

SetArchiveFiles sets ArchiveFiles field to given value.

### HasArchiveFiles

`func (o *GetArchiveBucket200Response) HasArchiveFiles() bool`

HasArchiveFiles returns a boolean if a field has been set.

### GetArchiveFileCount

`func (o *GetArchiveBucket200Response) GetArchiveFileCount() int64`

GetArchiveFileCount returns the ArchiveFileCount field if non-nil, zero value otherwise.

### GetArchiveFileCountOk

`func (o *GetArchiveBucket200Response) GetArchiveFileCountOk() (*int64, bool)`

GetArchiveFileCountOk returns a tuple with the ArchiveFileCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveFileCount

`func (o *GetArchiveBucket200Response) SetArchiveFileCount(v int64)`

SetArchiveFileCount sets ArchiveFileCount field to given value.

### HasArchiveFileCount

`func (o *GetArchiveBucket200Response) HasArchiveFileCount() bool`

HasArchiveFileCount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


