# GetPlugin200ResponseAllOfPlugin

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
**Providers** | Pointer to [**[]GetPlugin200ResponseAllOfPluginProvidersInner**](GetPlugin200ResponseAllOfPluginProvidersInner.md) |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**OptionTypes** | Pointer to [**[]GetPlugin200ResponseAllOfPluginOptionTypesInner**](GetPlugin200ResponseAllOfPluginOptionTypesInner.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetPlugin200ResponseAllOfPlugin

`func NewGetPlugin200ResponseAllOfPlugin() *GetPlugin200ResponseAllOfPlugin`

NewGetPlugin200ResponseAllOfPlugin instantiates a new GetPlugin200ResponseAllOfPlugin object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPlugin200ResponseAllOfPluginWithDefaults

`func NewGetPlugin200ResponseAllOfPluginWithDefaults() *GetPlugin200ResponseAllOfPlugin`

NewGetPlugin200ResponseAllOfPluginWithDefaults instantiates a new GetPlugin200ResponseAllOfPlugin object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetPlugin200ResponseAllOfPlugin) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetPlugin200ResponseAllOfPlugin) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetPlugin200ResponseAllOfPlugin) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetPlugin200ResponseAllOfPlugin) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetPlugin200ResponseAllOfPlugin) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetPlugin200ResponseAllOfPlugin) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetPlugin200ResponseAllOfPlugin) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetPlugin200ResponseAllOfPlugin) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *GetPlugin200ResponseAllOfPlugin) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetPlugin200ResponseAllOfPlugin) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetPlugin200ResponseAllOfPlugin) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetPlugin200ResponseAllOfPlugin) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *GetPlugin200ResponseAllOfPlugin) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetPlugin200ResponseAllOfPlugin) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetPlugin200ResponseAllOfPlugin) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetPlugin200ResponseAllOfPlugin) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersion

`func (o *GetPlugin200ResponseAllOfPlugin) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetPlugin200ResponseAllOfPlugin) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetPlugin200ResponseAllOfPlugin) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetPlugin200ResponseAllOfPlugin) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetEnabled

`func (o *GetPlugin200ResponseAllOfPlugin) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetPlugin200ResponseAllOfPlugin) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetPlugin200ResponseAllOfPlugin) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetPlugin200ResponseAllOfPlugin) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAuthor

`func (o *GetPlugin200ResponseAllOfPlugin) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *GetPlugin200ResponseAllOfPlugin) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *GetPlugin200ResponseAllOfPlugin) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *GetPlugin200ResponseAllOfPlugin) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### SetAuthorNil

`func (o *GetPlugin200ResponseAllOfPlugin) SetAuthorNil(b bool)`

 SetAuthorNil sets the value for Author to be an explicit nil

### UnsetAuthor
`func (o *GetPlugin200ResponseAllOfPlugin) UnsetAuthor()`

UnsetAuthor ensures that no value is present for Author, not even an explicit nil
### GetWebsiteUrl

`func (o *GetPlugin200ResponseAllOfPlugin) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *GetPlugin200ResponseAllOfPlugin) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *GetPlugin200ResponseAllOfPlugin) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *GetPlugin200ResponseAllOfPlugin) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### SetWebsiteUrlNil

`func (o *GetPlugin200ResponseAllOfPlugin) SetWebsiteUrlNil(b bool)`

 SetWebsiteUrlNil sets the value for WebsiteUrl to be an explicit nil

### UnsetWebsiteUrl
`func (o *GetPlugin200ResponseAllOfPlugin) UnsetWebsiteUrl()`

UnsetWebsiteUrl ensures that no value is present for WebsiteUrl, not even an explicit nil
### GetSourceCodeLocationUrl

`func (o *GetPlugin200ResponseAllOfPlugin) GetSourceCodeLocationUrl() string`

GetSourceCodeLocationUrl returns the SourceCodeLocationUrl field if non-nil, zero value otherwise.

### GetSourceCodeLocationUrlOk

`func (o *GetPlugin200ResponseAllOfPlugin) GetSourceCodeLocationUrlOk() (*string, bool)`

GetSourceCodeLocationUrlOk returns a tuple with the SourceCodeLocationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCodeLocationUrl

`func (o *GetPlugin200ResponseAllOfPlugin) SetSourceCodeLocationUrl(v string)`

SetSourceCodeLocationUrl sets SourceCodeLocationUrl field to given value.

### HasSourceCodeLocationUrl

`func (o *GetPlugin200ResponseAllOfPlugin) HasSourceCodeLocationUrl() bool`

HasSourceCodeLocationUrl returns a boolean if a field has been set.

### SetSourceCodeLocationUrlNil

`func (o *GetPlugin200ResponseAllOfPlugin) SetSourceCodeLocationUrlNil(b bool)`

 SetSourceCodeLocationUrlNil sets the value for SourceCodeLocationUrl to be an explicit nil

### UnsetSourceCodeLocationUrl
`func (o *GetPlugin200ResponseAllOfPlugin) UnsetSourceCodeLocationUrl()`

UnsetSourceCodeLocationUrl ensures that no value is present for SourceCodeLocationUrl, not even an explicit nil
### GetIssueTrackerUrl

`func (o *GetPlugin200ResponseAllOfPlugin) GetIssueTrackerUrl() string`

GetIssueTrackerUrl returns the IssueTrackerUrl field if non-nil, zero value otherwise.

### GetIssueTrackerUrlOk

`func (o *GetPlugin200ResponseAllOfPlugin) GetIssueTrackerUrlOk() (*string, bool)`

GetIssueTrackerUrlOk returns a tuple with the IssueTrackerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTrackerUrl

`func (o *GetPlugin200ResponseAllOfPlugin) SetIssueTrackerUrl(v string)`

SetIssueTrackerUrl sets IssueTrackerUrl field to given value.

### HasIssueTrackerUrl

`func (o *GetPlugin200ResponseAllOfPlugin) HasIssueTrackerUrl() bool`

HasIssueTrackerUrl returns a boolean if a field has been set.

### SetIssueTrackerUrlNil

`func (o *GetPlugin200ResponseAllOfPlugin) SetIssueTrackerUrlNil(b bool)`

 SetIssueTrackerUrlNil sets the value for IssueTrackerUrl to be an explicit nil

### UnsetIssueTrackerUrl
`func (o *GetPlugin200ResponseAllOfPlugin) UnsetIssueTrackerUrl()`

UnsetIssueTrackerUrl ensures that no value is present for IssueTrackerUrl, not even an explicit nil
### GetValid

`func (o *GetPlugin200ResponseAllOfPlugin) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *GetPlugin200ResponseAllOfPlugin) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *GetPlugin200ResponseAllOfPlugin) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *GetPlugin200ResponseAllOfPlugin) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetHasValidUpdate

`func (o *GetPlugin200ResponseAllOfPlugin) GetHasValidUpdate() bool`

GetHasValidUpdate returns the HasValidUpdate field if non-nil, zero value otherwise.

### GetHasValidUpdateOk

`func (o *GetPlugin200ResponseAllOfPlugin) GetHasValidUpdateOk() (*bool, bool)`

GetHasValidUpdateOk returns a tuple with the HasValidUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasValidUpdate

`func (o *GetPlugin200ResponseAllOfPlugin) SetHasValidUpdate(v bool)`

SetHasValidUpdate sets HasValidUpdate field to given value.

### HasHasValidUpdate

`func (o *GetPlugin200ResponseAllOfPlugin) HasHasValidUpdate() bool`

HasHasValidUpdate returns a boolean if a field has been set.

### GetStatus

`func (o *GetPlugin200ResponseAllOfPlugin) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetPlugin200ResponseAllOfPlugin) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetPlugin200ResponseAllOfPlugin) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetPlugin200ResponseAllOfPlugin) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *GetPlugin200ResponseAllOfPlugin) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GetPlugin200ResponseAllOfPlugin) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GetPlugin200ResponseAllOfPlugin) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *GetPlugin200ResponseAllOfPlugin) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *GetPlugin200ResponseAllOfPlugin) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *GetPlugin200ResponseAllOfPlugin) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetProviders

`func (o *GetPlugin200ResponseAllOfPlugin) GetProviders() []GetPlugin200ResponseAllOfPluginProvidersInner`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *GetPlugin200ResponseAllOfPlugin) GetProvidersOk() (*[]GetPlugin200ResponseAllOfPluginProvidersInner, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *GetPlugin200ResponseAllOfPlugin) SetProviders(v []GetPlugin200ResponseAllOfPluginProvidersInner)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *GetPlugin200ResponseAllOfPlugin) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetConfig

`func (o *GetPlugin200ResponseAllOfPlugin) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetPlugin200ResponseAllOfPlugin) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetPlugin200ResponseAllOfPlugin) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetPlugin200ResponseAllOfPlugin) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetOptionTypes

`func (o *GetPlugin200ResponseAllOfPlugin) GetOptionTypes() []GetPlugin200ResponseAllOfPluginOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *GetPlugin200ResponseAllOfPlugin) GetOptionTypesOk() (*[]GetPlugin200ResponseAllOfPluginOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *GetPlugin200ResponseAllOfPlugin) SetOptionTypes(v []GetPlugin200ResponseAllOfPluginOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *GetPlugin200ResponseAllOfPlugin) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetPlugin200ResponseAllOfPlugin) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetPlugin200ResponseAllOfPlugin) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetPlugin200ResponseAllOfPlugin) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetPlugin200ResponseAllOfPlugin) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetPlugin200ResponseAllOfPlugin) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetPlugin200ResponseAllOfPlugin) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetPlugin200ResponseAllOfPlugin) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetPlugin200ResponseAllOfPlugin) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


