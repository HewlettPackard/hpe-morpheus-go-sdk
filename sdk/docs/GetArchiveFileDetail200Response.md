# GetArchiveFileDetail200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArchiveFile** | Pointer to [**GetArchiveBucket200ResponseArchiveFilesInner**](GetArchiveBucket200ResponseArchiveFilesInner.md) |  | [optional] 
**IsOwner** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetArchiveFileDetail200Response

`func NewGetArchiveFileDetail200Response() *GetArchiveFileDetail200Response`

NewGetArchiveFileDetail200Response instantiates a new GetArchiveFileDetail200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetArchiveFileDetail200ResponseWithDefaults

`func NewGetArchiveFileDetail200ResponseWithDefaults() *GetArchiveFileDetail200Response`

NewGetArchiveFileDetail200ResponseWithDefaults instantiates a new GetArchiveFileDetail200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArchiveFile

`func (o *GetArchiveFileDetail200Response) GetArchiveFile() GetArchiveBucket200ResponseArchiveFilesInner`

GetArchiveFile returns the ArchiveFile field if non-nil, zero value otherwise.

### GetArchiveFileOk

`func (o *GetArchiveFileDetail200Response) GetArchiveFileOk() (*GetArchiveBucket200ResponseArchiveFilesInner, bool)`

GetArchiveFileOk returns a tuple with the ArchiveFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArchiveFile

`func (o *GetArchiveFileDetail200Response) SetArchiveFile(v GetArchiveBucket200ResponseArchiveFilesInner)`

SetArchiveFile sets ArchiveFile field to given value.

### HasArchiveFile

`func (o *GetArchiveFileDetail200Response) HasArchiveFile() bool`

HasArchiveFile returns a boolean if a field has been set.

### GetIsOwner

`func (o *GetArchiveFileDetail200Response) GetIsOwner() bool`

GetIsOwner returns the IsOwner field if non-nil, zero value otherwise.

### GetIsOwnerOk

`func (o *GetArchiveFileDetail200Response) GetIsOwnerOk() (*bool, bool)`

GetIsOwnerOk returns a tuple with the IsOwner field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsOwner

`func (o *GetArchiveFileDetail200Response) SetIsOwner(v bool)`

SetIsOwner sets IsOwner field to given value.

### HasIsOwner

`func (o *GetArchiveFileDetail200Response) HasIsOwner() bool`

HasIsOwner returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


