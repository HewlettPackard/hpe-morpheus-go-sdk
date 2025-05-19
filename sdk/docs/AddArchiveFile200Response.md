# AddArchiveFile200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveFile** | Pointer to [**GetArchiveBucket200ResponseArchiveFilesInner**](GetArchiveBucket200ResponseArchiveFilesInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddArchiveFile200Response

`func NewAddArchiveFile200Response() *AddArchiveFile200Response`

NewAddArchiveFile200Response instantiates a new AddArchiveFile200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddArchiveFile200ResponseWithDefaults

`func NewAddArchiveFile200ResponseWithDefaults() *AddArchiveFile200Response`

NewAddArchiveFile200ResponseWithDefaults instantiates a new AddArchiveFile200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveFile

`func (o *AddArchiveFile200Response) GetArchiveFile() GetArchiveBucket200ResponseArchiveFilesInner`

GetArchiveFile returns the ArchiveFile field if non-nil, zero value otherwise.

### GetArchiveFileOk

`func (o *AddArchiveFile200Response) GetArchiveFileOk() (*GetArchiveBucket200ResponseArchiveFilesInner, bool)`

GetArchiveFileOk returns a tuple with the ArchiveFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveFile

`func (o *AddArchiveFile200Response) SetArchiveFile(v GetArchiveBucket200ResponseArchiveFilesInner)`

SetArchiveFile sets ArchiveFile field to given value.

### HasArchiveFile

`func (o *AddArchiveFile200Response) HasArchiveFile() bool`

HasArchiveFile returns a boolean if a field has been set.

### GetSuccess

`func (o *AddArchiveFile200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddArchiveFile200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddArchiveFile200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddArchiveFile200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


