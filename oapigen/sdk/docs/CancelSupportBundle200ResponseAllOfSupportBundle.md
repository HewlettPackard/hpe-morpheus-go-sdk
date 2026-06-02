# CancelSupportBundle200ResponseAllOfSupportBundle

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
**StorageProvider** | Pointer to [**CancelSupportBundle200ResponseAllOfSupportBundleStorageProvider**](CancelSupportBundle200ResponseAllOfSupportBundleStorageProvider.md) |  | [optional] 

## Methods

### NewCancelSupportBundle200ResponseAllOfSupportBundle

`func NewCancelSupportBundle200ResponseAllOfSupportBundle() *CancelSupportBundle200ResponseAllOfSupportBundle`

NewCancelSupportBundle200ResponseAllOfSupportBundle instantiates a new CancelSupportBundle200ResponseAllOfSupportBundle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetStatus

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetStartedAt

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetCompletedAt

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetFilePath

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### SetFilePathNil

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) SetFilePathNil(b bool)`

 SetFilePathNil sets the value for FilePath to be an explicit nil

### UnsetFilePath
`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) UnsetFilePath()`

UnsetFilePath ensures that no value is present for FilePath, not even an explicit nil
### GetContentLength

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) GetContentLength() int64`

GetContentLength returns the ContentLength field if non-nil, zero value otherwise.

### GetContentLengthOk

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) GetContentLengthOk() (*int64, bool)`

GetContentLengthOk returns a tuple with the ContentLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentLength

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) SetContentLength(v int64)`

SetContentLength sets ContentLength field to given value.

### HasContentLength

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) HasContentLength() bool`

HasContentLength returns a boolean if a field has been set.

### SetContentLengthNil

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) SetContentLengthNil(b bool)`

 SetContentLengthNil sets the value for ContentLength to be an explicit nil

### UnsetContentLength
`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) UnsetContentLength()`

UnsetContentLength ensures that no value is present for ContentLength, not even an explicit nil
### GetContentType

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetChecksum

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### SetChecksumNil

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) SetChecksumNil(b bool)`

 SetChecksumNil sets the value for Checksum to be an explicit nil

### UnsetChecksum
`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) UnsetChecksum()`

UnsetChecksum ensures that no value is present for Checksum, not even an explicit nil
### GetStorageProvider

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) GetStorageProvider() CancelSupportBundle200ResponseAllOfSupportBundleStorageProvider`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) GetStorageProviderOk() (*CancelSupportBundle200ResponseAllOfSupportBundleStorageProvider, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) SetStorageProvider(v CancelSupportBundle200ResponseAllOfSupportBundleStorageProvider)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *CancelSupportBundle200ResponseAllOfSupportBundle) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


