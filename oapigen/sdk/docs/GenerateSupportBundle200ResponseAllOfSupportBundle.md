# GenerateSupportBundle200ResponseAllOfSupportBundle

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
**StorageProvider** | Pointer to [**GenerateSupportBundle200ResponseAllOfSupportBundleStorageProvider**](GenerateSupportBundle200ResponseAllOfSupportBundleStorageProvider.md) |  | [optional] 

## Methods

### NewGenerateSupportBundle200ResponseAllOfSupportBundle

`func NewGenerateSupportBundle200ResponseAllOfSupportBundle() *GenerateSupportBundle200ResponseAllOfSupportBundle`

NewGenerateSupportBundle200ResponseAllOfSupportBundle instantiates a new GenerateSupportBundle200ResponseAllOfSupportBundle object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateSupportBundle200ResponseAllOfSupportBundleWithDefaults

`func NewGenerateSupportBundle200ResponseAllOfSupportBundleWithDefaults() *GenerateSupportBundle200ResponseAllOfSupportBundle`

NewGenerateSupportBundle200ResponseAllOfSupportBundleWithDefaults instantiates a new GenerateSupportBundle200ResponseAllOfSupportBundle object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetStatus

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetStartedAt

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetCompletedAt

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetFilePath

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### SetFilePathNil

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) SetFilePathNil(b bool)`

 SetFilePathNil sets the value for FilePath to be an explicit nil

### UnsetFilePath
`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) UnsetFilePath()`

UnsetFilePath ensures that no value is present for FilePath, not even an explicit nil
### GetContentLength

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) GetContentLength() int64`

GetContentLength returns the ContentLength field if non-nil, zero value otherwise.

### GetContentLengthOk

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) GetContentLengthOk() (*int64, bool)`

GetContentLengthOk returns a tuple with the ContentLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentLength

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) SetContentLength(v int64)`

SetContentLength sets ContentLength field to given value.

### HasContentLength

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) HasContentLength() bool`

HasContentLength returns a boolean if a field has been set.

### SetContentLengthNil

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) SetContentLengthNil(b bool)`

 SetContentLengthNil sets the value for ContentLength to be an explicit nil

### UnsetContentLength
`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) UnsetContentLength()`

UnsetContentLength ensures that no value is present for ContentLength, not even an explicit nil
### GetContentType

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetChecksum

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### SetChecksumNil

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) SetChecksumNil(b bool)`

 SetChecksumNil sets the value for Checksum to be an explicit nil

### UnsetChecksum
`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) UnsetChecksum()`

UnsetChecksum ensures that no value is present for Checksum, not even an explicit nil
### GetStorageProvider

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) GetStorageProvider() GenerateSupportBundle200ResponseAllOfSupportBundleStorageProvider`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) GetStorageProviderOk() (*GenerateSupportBundle200ResponseAllOfSupportBundleStorageProvider, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) SetStorageProvider(v GenerateSupportBundle200ResponseAllOfSupportBundleStorageProvider)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *GenerateSupportBundle200ResponseAllOfSupportBundle) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


