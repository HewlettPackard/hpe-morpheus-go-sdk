# ListSupportBundles200ResponseAllOfSupportBundlesInner

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
**StorageProvider** | Pointer to [**ListSupportBundles200ResponseAllOfSupportBundlesInnerStorageProvider**](ListSupportBundles200ResponseAllOfSupportBundlesInnerStorageProvider.md) |  | [optional] 

## Methods

### NewListSupportBundles200ResponseAllOfSupportBundlesInner

`func NewListSupportBundles200ResponseAllOfSupportBundlesInner() *ListSupportBundles200ResponseAllOfSupportBundlesInner`

NewListSupportBundles200ResponseAllOfSupportBundlesInner instantiates a new ListSupportBundles200ResponseAllOfSupportBundlesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSupportBundles200ResponseAllOfSupportBundlesInnerWithDefaults

`func NewListSupportBundles200ResponseAllOfSupportBundlesInnerWithDefaults() *ListSupportBundles200ResponseAllOfSupportBundlesInner`

NewListSupportBundles200ResponseAllOfSupportBundlesInnerWithDefaults instantiates a new ListSupportBundles200ResponseAllOfSupportBundlesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetStatus

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetStartedAt

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) GetStartedAt() time.Time`

GetStartedAt returns the StartedAt field if non-nil, zero value otherwise.

### GetStartedAtOk

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) GetStartedAtOk() (*time.Time, bool)`

GetStartedAtOk returns a tuple with the StartedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartedAt

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) SetStartedAt(v time.Time)`

SetStartedAt sets StartedAt field to given value.

### HasStartedAt

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) HasStartedAt() bool`

HasStartedAt returns a boolean if a field has been set.

### SetStartedAtNil

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) SetStartedAtNil(b bool)`

 SetStartedAtNil sets the value for StartedAt to be an explicit nil

### UnsetStartedAt
`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) UnsetStartedAt()`

UnsetStartedAt ensures that no value is present for StartedAt, not even an explicit nil
### GetCompletedAt

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) GetCompletedAt() time.Time`

GetCompletedAt returns the CompletedAt field if non-nil, zero value otherwise.

### GetCompletedAtOk

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) GetCompletedAtOk() (*time.Time, bool)`

GetCompletedAtOk returns a tuple with the CompletedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletedAt

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) SetCompletedAt(v time.Time)`

SetCompletedAt sets CompletedAt field to given value.

### HasCompletedAt

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) HasCompletedAt() bool`

HasCompletedAt returns a boolean if a field has been set.

### SetCompletedAtNil

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) SetCompletedAtNil(b bool)`

 SetCompletedAtNil sets the value for CompletedAt to be an explicit nil

### UnsetCompletedAt
`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) UnsetCompletedAt()`

UnsetCompletedAt ensures that no value is present for CompletedAt, not even an explicit nil
### GetFilePath

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) GetFilePath() string`

GetFilePath returns the FilePath field if non-nil, zero value otherwise.

### GetFilePathOk

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) GetFilePathOk() (*string, bool)`

GetFilePathOk returns a tuple with the FilePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilePath

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) SetFilePath(v string)`

SetFilePath sets FilePath field to given value.

### HasFilePath

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) HasFilePath() bool`

HasFilePath returns a boolean if a field has been set.

### SetFilePathNil

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) SetFilePathNil(b bool)`

 SetFilePathNil sets the value for FilePath to be an explicit nil

### UnsetFilePath
`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) UnsetFilePath()`

UnsetFilePath ensures that no value is present for FilePath, not even an explicit nil
### GetContentLength

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) GetContentLength() int64`

GetContentLength returns the ContentLength field if non-nil, zero value otherwise.

### GetContentLengthOk

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) GetContentLengthOk() (*int64, bool)`

GetContentLengthOk returns a tuple with the ContentLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentLength

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) SetContentLength(v int64)`

SetContentLength sets ContentLength field to given value.

### HasContentLength

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) HasContentLength() bool`

HasContentLength returns a boolean if a field has been set.

### SetContentLengthNil

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) SetContentLengthNil(b bool)`

 SetContentLengthNil sets the value for ContentLength to be an explicit nil

### UnsetContentLength
`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) UnsetContentLength()`

UnsetContentLength ensures that no value is present for ContentLength, not even an explicit nil
### GetContentType

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) GetContentType() string`

GetContentType returns the ContentType field if non-nil, zero value otherwise.

### GetContentTypeOk

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) GetContentTypeOk() (*string, bool)`

GetContentTypeOk returns a tuple with the ContentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentType

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) SetContentType(v string)`

SetContentType sets ContentType field to given value.

### HasContentType

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) HasContentType() bool`

HasContentType returns a boolean if a field has been set.

### SetContentTypeNil

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) SetContentTypeNil(b bool)`

 SetContentTypeNil sets the value for ContentType to be an explicit nil

### UnsetContentType
`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) UnsetContentType()`

UnsetContentType ensures that no value is present for ContentType, not even an explicit nil
### GetChecksum

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.

### HasChecksum

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) HasChecksum() bool`

HasChecksum returns a boolean if a field has been set.

### SetChecksumNil

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) SetChecksumNil(b bool)`

 SetChecksumNil sets the value for Checksum to be an explicit nil

### UnsetChecksum
`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) UnsetChecksum()`

UnsetChecksum ensures that no value is present for Checksum, not even an explicit nil
### GetStorageProvider

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) GetStorageProvider() ListSupportBundles200ResponseAllOfSupportBundlesInnerStorageProvider`

GetStorageProvider returns the StorageProvider field if non-nil, zero value otherwise.

### GetStorageProviderOk

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) GetStorageProviderOk() (*ListSupportBundles200ResponseAllOfSupportBundlesInnerStorageProvider, bool)`

GetStorageProviderOk returns a tuple with the StorageProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageProvider

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) SetStorageProvider(v ListSupportBundles200ResponseAllOfSupportBundlesInnerStorageProvider)`

SetStorageProvider sets StorageProvider field to given value.

### HasStorageProvider

`func (o *ListSupportBundles200ResponseAllOfSupportBundlesInner) HasStorageProvider() bool`

HasStorageProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


