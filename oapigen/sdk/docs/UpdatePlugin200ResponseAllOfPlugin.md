# UpdatePlugin200ResponseAllOfPlugin

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
**Providers** | Pointer to [**[]UpdatePlugin200ResponseAllOfPluginProvidersInner**](UpdatePlugin200ResponseAllOfPluginProvidersInner.md) |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**OptionTypes** | Pointer to [**[]UpdatePlugin200ResponseAllOfPluginOptionTypesInner**](UpdatePlugin200ResponseAllOfPluginOptionTypesInner.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewUpdatePlugin200ResponseAllOfPlugin

`func NewUpdatePlugin200ResponseAllOfPlugin() *UpdatePlugin200ResponseAllOfPlugin`

NewUpdatePlugin200ResponseAllOfPlugin instantiates a new UpdatePlugin200ResponseAllOfPlugin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *UpdatePlugin200ResponseAllOfPlugin) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdatePlugin200ResponseAllOfPlugin) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdatePlugin200ResponseAllOfPlugin) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdatePlugin200ResponseAllOfPlugin) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdatePlugin200ResponseAllOfPlugin) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdatePlugin200ResponseAllOfPlugin) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdatePlugin200ResponseAllOfPlugin) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdatePlugin200ResponseAllOfPlugin) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *UpdatePlugin200ResponseAllOfPlugin) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdatePlugin200ResponseAllOfPlugin) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdatePlugin200ResponseAllOfPlugin) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdatePlugin200ResponseAllOfPlugin) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *UpdatePlugin200ResponseAllOfPlugin) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdatePlugin200ResponseAllOfPlugin) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdatePlugin200ResponseAllOfPlugin) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdatePlugin200ResponseAllOfPlugin) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersion

`func (o *UpdatePlugin200ResponseAllOfPlugin) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpdatePlugin200ResponseAllOfPlugin) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpdatePlugin200ResponseAllOfPlugin) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UpdatePlugin200ResponseAllOfPlugin) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdatePlugin200ResponseAllOfPlugin) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdatePlugin200ResponseAllOfPlugin) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdatePlugin200ResponseAllOfPlugin) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdatePlugin200ResponseAllOfPlugin) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAuthor

`func (o *UpdatePlugin200ResponseAllOfPlugin) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *UpdatePlugin200ResponseAllOfPlugin) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *UpdatePlugin200ResponseAllOfPlugin) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *UpdatePlugin200ResponseAllOfPlugin) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### SetAuthorNil

`func (o *UpdatePlugin200ResponseAllOfPlugin) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *UpdatePlugin200ResponseAllOfPlugin) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetWebsiteUrl

`func (o *UpdatePlugin200ResponseAllOfPlugin) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *UpdatePlugin200ResponseAllOfPlugin) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *UpdatePlugin200ResponseAllOfPlugin) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *UpdatePlugin200ResponseAllOfPlugin) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### SetWebsiteUrlNil

`func (o *UpdatePlugin200ResponseAllOfPlugin) SetWebsiteUrlNil(b bool)`

 SetWebsiteUrlNil sets the value for WebsiteUrl to be an explicit nil

### UnsetWebsiteUrl
`func (o *UpdatePlugin200ResponseAllOfPlugin) UnsetWebsiteUrl()`

UnsetWebsiteUrl ensures that no value is present for WebsiteUrl, not even an explicit nil
### GetSourceCodeLocationUrl

`func (o *UpdatePlugin200ResponseAllOfPlugin) GetSourceCodeLocationUrl() string`

GetSourceCodeLocationUrl returns the SourceCodeLocationUrl field if non-nil, zero value otherwise.

### GetSourceCodeLocationUrlOk

`func (o *UpdatePlugin200ResponseAllOfPlugin) GetSourceCodeLocationUrlOk() (*string, bool)`

GetSourceCodeLocationUrlOk returns a tuple with the SourceCodeLocationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCodeLocationUrl

`func (o *UpdatePlugin200ResponseAllOfPlugin) SetSourceCodeLocationUrl(v string)`

SetSourceCodeLocationUrl sets SourceCodeLocationUrl field to given value.

### HasSourceCodeLocationUrl

`func (o *UpdatePlugin200ResponseAllOfPlugin) HasSourceCodeLocationUrl() bool`

HasSourceCodeLocationUrl returns a boolean if a field has been set.

### SetSourceCodeLocationUrlNil

`func (o *UpdatePlugin200ResponseAllOfPlugin) SetSourceCodeLocationUrlNil(b bool)`

 SetSourceCodeLocationUrlNil sets the value for SourceCodeLocationUrl to be an explicit nil

### UnsetSourceCodeLocationUrl
`func (o *UpdatePlugin200ResponseAllOfPlugin) UnsetSourceCodeLocationUrl()`

UnsetSourceCodeLocationUrl ensures that no value is present for SourceCodeLocationUrl, not even an explicit nil
### GetIssueTrackerUrl

`func (o *UpdatePlugin200ResponseAllOfPlugin) GetIssueTrackerUrl() string`

GetIssueTrackerUrl returns the IssueTrackerUrl field if non-nil, zero value otherwise.

### GetIssueTrackerUrlOk

`func (o *UpdatePlugin200ResponseAllOfPlugin) GetIssueTrackerUrlOk() (*string, bool)`

GetIssueTrackerUrlOk returns a tuple with the IssueTrackerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTrackerUrl

`func (o *UpdatePlugin200ResponseAllOfPlugin) SetIssueTrackerUrl(v string)`

SetIssueTrackerUrl sets IssueTrackerUrl field to given value.

### HasIssueTrackerUrl

`func (o *UpdatePlugin200ResponseAllOfPlugin) HasIssueTrackerUrl() bool`

HasIssueTrackerUrl returns a boolean if a field has been set.

### SetIssueTrackerUrlNil

`func (o *UpdatePlugin200ResponseAllOfPlugin) SetIssueTrackerUrlNil(b bool)`

 SetIssueTrackerUrlNil sets the value for IssueTrackerUrl to be an explicit nil

### UnsetIssueTrackerUrl
`func (o *UpdatePlugin200ResponseAllOfPlugin) UnsetIssueTrackerUrl()`

UnsetIssueTrackerUrl ensures that no value is present for IssueTrackerUrl, not even an explicit nil
### GetValid

`func (o *UpdatePlugin200ResponseAllOfPlugin) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *UpdatePlugin200ResponseAllOfPlugin) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *UpdatePlugin200ResponseAllOfPlugin) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *UpdatePlugin200ResponseAllOfPlugin) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetHasValidUpdate

`func (o *UpdatePlugin200ResponseAllOfPlugin) GetHasValidUpdate() bool`

GetHasValidUpdate returns the HasValidUpdate field if non-nil, zero value otherwise.

### GetHasValidUpdateOk

`func (o *UpdatePlugin200ResponseAllOfPlugin) GetHasValidUpdateOk() (*bool, bool)`

GetHasValidUpdateOk returns a tuple with the HasValidUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasValidUpdate

`func (o *UpdatePlugin200ResponseAllOfPlugin) SetHasValidUpdate(v bool)`

SetHasValidUpdate sets HasValidUpdate field to given value.

### HasHasValidUpdate

`func (o *UpdatePlugin200ResponseAllOfPlugin) HasHasValidUpdate() bool`

HasHasValidUpdate returns a boolean if a field has been set.

### GetStatus

`func (o *UpdatePlugin200ResponseAllOfPlugin) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdatePlugin200ResponseAllOfPlugin) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdatePlugin200ResponseAllOfPlugin) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdatePlugin200ResponseAllOfPlugin) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *UpdatePlugin200ResponseAllOfPlugin) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *UpdatePlugin200ResponseAllOfPlugin) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *UpdatePlugin200ResponseAllOfPlugin) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *UpdatePlugin200ResponseAllOfPlugin) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *UpdatePlugin200ResponseAllOfPlugin) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *UpdatePlugin200ResponseAllOfPlugin) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetProviders

`func (o *UpdatePlugin200ResponseAllOfPlugin) GetProviders() []UpdatePlugin200ResponseAllOfPluginProvidersInner`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *UpdatePlugin200ResponseAllOfPlugin) GetProvidersOk() (*[]UpdatePlugin200ResponseAllOfPluginProvidersInner, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *UpdatePlugin200ResponseAllOfPlugin) SetProviders(v []UpdatePlugin200ResponseAllOfPluginProvidersInner)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *UpdatePlugin200ResponseAllOfPlugin) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetConfig

`func (o *UpdatePlugin200ResponseAllOfPlugin) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdatePlugin200ResponseAllOfPlugin) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdatePlugin200ResponseAllOfPlugin) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdatePlugin200ResponseAllOfPlugin) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetOptionTypes

`func (o *UpdatePlugin200ResponseAllOfPlugin) GetOptionTypes() []UpdatePlugin200ResponseAllOfPluginOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *UpdatePlugin200ResponseAllOfPlugin) GetOptionTypesOk() (*[]UpdatePlugin200ResponseAllOfPluginOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *UpdatePlugin200ResponseAllOfPlugin) SetOptionTypes(v []UpdatePlugin200ResponseAllOfPluginOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *UpdatePlugin200ResponseAllOfPlugin) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetDateCreated

`func (o *UpdatePlugin200ResponseAllOfPlugin) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdatePlugin200ResponseAllOfPlugin) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdatePlugin200ResponseAllOfPlugin) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdatePlugin200ResponseAllOfPlugin) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdatePlugin200ResponseAllOfPlugin) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdatePlugin200ResponseAllOfPlugin) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdatePlugin200ResponseAllOfPlugin) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdatePlugin200ResponseAllOfPlugin) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


