# UploadPlugin200ResponseAllOfPlugin

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Author** | Pointer to **NullableString** |  | [optional] 
**WebsiteUrl** | Pointer to **NullableString** |  | [optional] 
**SourceCodeLocationUrl** | Pointer to **NullableString** |  | [optional] 
**IssueTrackerUrl** | Pointer to **NullableString** |  | [optional] 
**Valid** | Pointer to **bool** |  | [optional] 
**HasValidUpdate** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **NullableString** |  | [optional] 
**Providers** | Pointer to [**[]UploadPlugin200ResponseAllOfPluginProvidersInner**](UploadPlugin200ResponseAllOfPluginProvidersInner.md) |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**OptionTypes** | Pointer to [**[]UploadPlugin200ResponseAllOfPluginOptionTypesInner**](UploadPlugin200ResponseAllOfPluginOptionTypesInner.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUploadPlugin200ResponseAllOfPlugin

`func NewUploadPlugin200ResponseAllOfPlugin() *UploadPlugin200ResponseAllOfPlugin`

NewUploadPlugin200ResponseAllOfPlugin instantiates a new UploadPlugin200ResponseAllOfPlugin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *UploadPlugin200ResponseAllOfPlugin) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UploadPlugin200ResponseAllOfPlugin) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UploadPlugin200ResponseAllOfPlugin) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UploadPlugin200ResponseAllOfPlugin) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UploadPlugin200ResponseAllOfPlugin) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UploadPlugin200ResponseAllOfPlugin) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UploadPlugin200ResponseAllOfPlugin) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UploadPlugin200ResponseAllOfPlugin) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *UploadPlugin200ResponseAllOfPlugin) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UploadPlugin200ResponseAllOfPlugin) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UploadPlugin200ResponseAllOfPlugin) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UploadPlugin200ResponseAllOfPlugin) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *UploadPlugin200ResponseAllOfPlugin) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UploadPlugin200ResponseAllOfPlugin) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UploadPlugin200ResponseAllOfPlugin) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UploadPlugin200ResponseAllOfPlugin) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersion

`func (o *UploadPlugin200ResponseAllOfPlugin) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UploadPlugin200ResponseAllOfPlugin) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UploadPlugin200ResponseAllOfPlugin) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UploadPlugin200ResponseAllOfPlugin) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetEnabled

`func (o *UploadPlugin200ResponseAllOfPlugin) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UploadPlugin200ResponseAllOfPlugin) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UploadPlugin200ResponseAllOfPlugin) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UploadPlugin200ResponseAllOfPlugin) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAuthor

`func (o *UploadPlugin200ResponseAllOfPlugin) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *UploadPlugin200ResponseAllOfPlugin) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *UploadPlugin200ResponseAllOfPlugin) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *UploadPlugin200ResponseAllOfPlugin) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### SetAuthorNil

`func (o *UploadPlugin200ResponseAllOfPlugin) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *UploadPlugin200ResponseAllOfPlugin) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetWebsiteUrl

`func (o *UploadPlugin200ResponseAllOfPlugin) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *UploadPlugin200ResponseAllOfPlugin) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *UploadPlugin200ResponseAllOfPlugin) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *UploadPlugin200ResponseAllOfPlugin) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### SetWebsiteUrlNil

`func (o *UploadPlugin200ResponseAllOfPlugin) SetWebsiteUrlNil(b bool)`

 SetWebsiteUrlNil sets the value for WebsiteUrl to be an explicit nil

### UnsetWebsiteUrl
`func (o *UploadPlugin200ResponseAllOfPlugin) UnsetWebsiteUrl()`

UnsetWebsiteUrl ensures that no value is present for WebsiteUrl, not even an explicit nil
### GetSourceCodeLocationUrl

`func (o *UploadPlugin200ResponseAllOfPlugin) GetSourceCodeLocationUrl() string`

GetSourceCodeLocationUrl returns the SourceCodeLocationUrl field if non-nil, zero value otherwise.

### GetSourceCodeLocationUrlOk

`func (o *UploadPlugin200ResponseAllOfPlugin) GetSourceCodeLocationUrlOk() (*string, bool)`

GetSourceCodeLocationUrlOk returns a tuple with the SourceCodeLocationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCodeLocationUrl

`func (o *UploadPlugin200ResponseAllOfPlugin) SetSourceCodeLocationUrl(v string)`

SetSourceCodeLocationUrl sets SourceCodeLocationUrl field to given value.

### HasSourceCodeLocationUrl

`func (o *UploadPlugin200ResponseAllOfPlugin) HasSourceCodeLocationUrl() bool`

HasSourceCodeLocationUrl returns a boolean if a field has been set.

### SetSourceCodeLocationUrlNil

`func (o *UploadPlugin200ResponseAllOfPlugin) SetSourceCodeLocationUrlNil(b bool)`

 SetSourceCodeLocationUrlNil sets the value for SourceCodeLocationUrl to be an explicit nil

### UnsetSourceCodeLocationUrl
`func (o *UploadPlugin200ResponseAllOfPlugin) UnsetSourceCodeLocationUrl()`

UnsetSourceCodeLocationUrl ensures that no value is present for SourceCodeLocationUrl, not even an explicit nil
### GetIssueTrackerUrl

`func (o *UploadPlugin200ResponseAllOfPlugin) GetIssueTrackerUrl() string`

GetIssueTrackerUrl returns the IssueTrackerUrl field if non-nil, zero value otherwise.

### GetIssueTrackerUrlOk

`func (o *UploadPlugin200ResponseAllOfPlugin) GetIssueTrackerUrlOk() (*string, bool)`

GetIssueTrackerUrlOk returns a tuple with the IssueTrackerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTrackerUrl

`func (o *UploadPlugin200ResponseAllOfPlugin) SetIssueTrackerUrl(v string)`

SetIssueTrackerUrl sets IssueTrackerUrl field to given value.

### HasIssueTrackerUrl

`func (o *UploadPlugin200ResponseAllOfPlugin) HasIssueTrackerUrl() bool`

HasIssueTrackerUrl returns a boolean if a field has been set.

### SetIssueTrackerUrlNil

`func (o *UploadPlugin200ResponseAllOfPlugin) SetIssueTrackerUrlNil(b bool)`

 SetIssueTrackerUrlNil sets the value for IssueTrackerUrl to be an explicit nil

### UnsetIssueTrackerUrl
`func (o *UploadPlugin200ResponseAllOfPlugin) UnsetIssueTrackerUrl()`

UnsetIssueTrackerUrl ensures that no value is present for IssueTrackerUrl, not even an explicit nil
### GetValid

`func (o *UploadPlugin200ResponseAllOfPlugin) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *UploadPlugin200ResponseAllOfPlugin) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *UploadPlugin200ResponseAllOfPlugin) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *UploadPlugin200ResponseAllOfPlugin) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetHasValidUpdate

`func (o *UploadPlugin200ResponseAllOfPlugin) GetHasValidUpdate() bool`

GetHasValidUpdate returns the HasValidUpdate field if non-nil, zero value otherwise.

### GetHasValidUpdateOk

`func (o *UploadPlugin200ResponseAllOfPlugin) GetHasValidUpdateOk() (*bool, bool)`

GetHasValidUpdateOk returns a tuple with the HasValidUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasValidUpdate

`func (o *UploadPlugin200ResponseAllOfPlugin) SetHasValidUpdate(v bool)`

SetHasValidUpdate sets HasValidUpdate field to given value.

### HasHasValidUpdate

`func (o *UploadPlugin200ResponseAllOfPlugin) HasHasValidUpdate() bool`

HasHasValidUpdate returns a boolean if a field has been set.

### GetStatus

`func (o *UploadPlugin200ResponseAllOfPlugin) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UploadPlugin200ResponseAllOfPlugin) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UploadPlugin200ResponseAllOfPlugin) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UploadPlugin200ResponseAllOfPlugin) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *UploadPlugin200ResponseAllOfPlugin) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *UploadPlugin200ResponseAllOfPlugin) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *UploadPlugin200ResponseAllOfPlugin) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *UploadPlugin200ResponseAllOfPlugin) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *UploadPlugin200ResponseAllOfPlugin) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *UploadPlugin200ResponseAllOfPlugin) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetProviders

`func (o *UploadPlugin200ResponseAllOfPlugin) GetProviders() []UploadPlugin200ResponseAllOfPluginProvidersInner`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *UploadPlugin200ResponseAllOfPlugin) GetProvidersOk() (*[]UploadPlugin200ResponseAllOfPluginProvidersInner, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *UploadPlugin200ResponseAllOfPlugin) SetProviders(v []UploadPlugin200ResponseAllOfPluginProvidersInner)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *UploadPlugin200ResponseAllOfPlugin) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetConfig

`func (o *UploadPlugin200ResponseAllOfPlugin) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UploadPlugin200ResponseAllOfPlugin) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UploadPlugin200ResponseAllOfPlugin) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UploadPlugin200ResponseAllOfPlugin) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetOptionTypes

`func (o *UploadPlugin200ResponseAllOfPlugin) GetOptionTypes() []UploadPlugin200ResponseAllOfPluginOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *UploadPlugin200ResponseAllOfPlugin) GetOptionTypesOk() (*[]UploadPlugin200ResponseAllOfPluginOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *UploadPlugin200ResponseAllOfPlugin) SetOptionTypes(v []UploadPlugin200ResponseAllOfPluginOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *UploadPlugin200ResponseAllOfPlugin) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetDateCreated

`func (o *UploadPlugin200ResponseAllOfPlugin) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UploadPlugin200ResponseAllOfPlugin) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UploadPlugin200ResponseAllOfPlugin) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UploadPlugin200ResponseAllOfPlugin) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UploadPlugin200ResponseAllOfPlugin) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UploadPlugin200ResponseAllOfPlugin) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UploadPlugin200ResponseAllOfPlugin) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UploadPlugin200ResponseAllOfPlugin) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


