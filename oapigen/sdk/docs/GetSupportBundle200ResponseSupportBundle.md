# GetSupportBundle200ResponseSupportBundle

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** | The current status of the support bundle generation | [optional] 
**StatusMessage** | Pointer to **NullableString** |  | [optional] 
**StartedAt** | Pointer to **NullableTime** |  | [optional] 
**CompletedAt** | Pointer to **NullableTime** |  | [optional] 
**FilePath** | Pointer to **NullableString** | The file path / filename of the generated bundle archive | [optional] 
**ContentLength** | Pointer to **NullableInt64** | File size in bytes | [optional] 
**ContentType** | Pointer to **NullableString** | MIME type of the bundle file | [optional] 
**Checksum** | Pointer to **NullableString** |  | [optional] 
**StorageProvider** | Pointer to [**GetSupportBundle200ResponseSupportBundleStorageProvider**](GetSupportBundle200ResponseSupportBundleStorageProvider.md) |  | [optional] 

## Methods

### NewGetSupportBundle200ResponseSupportBundle

`func NewGetSupportBundle200ResponseSupportBundle() *GetSupportBundle200ResponseSupportBundle`

NewGetSupportBundle200ResponseSupportBundle instantiates a new GetSupportBundle200ResponseSupportBundle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetSupportBundle200ResponseSupportBundle) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSupportBundle200ResponseSupportBundle) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSupportBundle200ResponseSupportBundle) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetSupportBundle200ResponseSupportBundle) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetSupportBundle200ResponseSupportBundle) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetSupportBundle200ResponseSupportBundle) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetSupportBundle200ResponseSupportBundle) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetSupportBundle200ResponseSupportBundle) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *GetSupportBundle200ResponseSupportBundle) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GetSupportBundle200ResponseSupportBundle) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GetSupportBundle200ResponseSupportBundle) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GetSupportBundle200ResponseSupportBundle) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetStatus

`func (o *GetSupportBundle200ResponseSupportBundle) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSupportBundle200ResponseSupportBundle) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSupportBundle200ResponseSupportBundle) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetSupportBundle200ResponseSupportBundle) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *GetSupportBundle200ResponseSupportBundle) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GetSupportBundle200ResponseSupportBundle) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GetSupportBundle200ResponseSupportBundle) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *GetSupportBundle200ResponseSupportBundle) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *GetSupportBundle200ResponseSupportBundle) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *GetSupportBundle200ResponseSupportBundle) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetStartedAt

`func (o *GetSupportBundle200ResponseSupportBundle) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *GetSupportBundle200ResponseSupportBundle) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *GetSupportBundle200ResponseSupportBundle) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *GetSupportBundle200ResponseSupportBundle) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *GetSupportBundle200ResponseSupportBundle) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *GetSupportBundle200ResponseSupportBundle) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetCompletedAt

`func (o *GetSupportBundle200ResponseSupportBundle) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *GetSupportBundle200ResponseSupportBundle) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *GetSupportBundle200ResponseSupportBundle) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *GetSupportBundle200ResponseSupportBundle) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *GetSupportBundle200ResponseSupportBundle) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *GetSupportBundle200ResponseSupportBundle) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetFilePath

`func (o *GetSupportBundle200ResponseSupportBundle) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *GetSupportBundle200ResponseSupportBundle) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *GetSupportBundle200ResponseSupportBundle) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *GetSupportBundle200ResponseSupportBundle) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### SetFilePathNil

`func (o *GetSupportBundle200ResponseSupportBundle) SetFilePathNil(b bool)`

 SetFilePathNil sets the value for FilePath to be an explicit nil

### UnsetFilePath
`func (o *GetSupportBundle200ResponseSupportBundle) UnsetFilePath()`

UnsetFilePath ensures that no value is present for FilePath, not even an explicit nil
### GetContentLength

`func (o *GetSupportBundle200ResponseSupportBundle) GetContentLength() int64`

GetContentLength returns the ContentLength field if non-nil, zero value otherwise.

### GetContentLengthOk

`func (o *GetSupportBundle200ResponseSupportBundle) GetContentLengthOk() (*int64, bool)`

GetContentLengthOk returns a tuple with the ContentLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentLength

`func (o *GetSupportBundle200ResponseSupportBundle) SetContentLength(v int64)`

SetContentLength sets ContentLength field to given value.

### HasContentLength

`func (o *GetSupportBundle200ResponseSupportBundle) HasContentLength() bool`

HasContentLength returns a boolean if a field has been set.

### SetContentLengthNil

`func (o *GetSupportBundle200ResponseSupportBundle) SetContentLengthNil(b bool)`

 SetContentLengthNil sets the value for ContentLength to be an explicit nil

### UnsetContentLength
`func (o *GetSupportBundle200ResponseSupportBundle) UnsetContentLength()`

UnsetContentLength ensures that no value is present for ContentLength, not even an explicit nil
### GetContentType

`func (o *GetSupportBundle200ResponseSupportBundle) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *GetSupportBundle200ResponseSupportBundle) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *GetSupportBundle200ResponseSupportBundle) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *GetSupportBundle200ResponseSupportBundle) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *GetSupportBundle200ResponseSupportBundle) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *GetSupportBundle200ResponseSupportBundle) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetChecksum

`func (o *GetSupportBundle200ResponseSupportBundle) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *GetSupportBundle200ResponseSupportBundle) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *GetSupportBundle200ResponseSupportBundle) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *GetSupportBundle200ResponseSupportBundle) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### SetChecksumNil

`func (o *GetSupportBundle200ResponseSupportBundle) SetChecksumNil(b bool)`

 SetChecksumNil sets the value for Checksum to be an explicit nil

### UnsetChecksum
`func (o *GetSupportBundle200ResponseSupportBundle) UnsetChecksum()`

UnsetChecksum ensures that no value is present for Checksum, not even an explicit nil
### GetStorageProvider

`func (o *GetSupportBundle200ResponseSupportBundle) GetStorageProvider() GetSupportBundle200ResponseSupportBundleStorageProvider`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *GetSupportBundle200ResponseSupportBundle) GetStorageProviderOk() (*GetSupportBundle200ResponseSupportBundleStorageProvider, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *GetSupportBundle200ResponseSupportBundle) SetStorageProvider(v GetSupportBundle200ResponseSupportBundleStorageProvider)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *GetSupportBundle200ResponseSupportBundle) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


