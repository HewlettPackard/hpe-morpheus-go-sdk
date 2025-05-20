# ListPlugins200ResponseAllOfPluginsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Author** | Pointer to **string** |  | [optional] 
**WebsiteUrl** | Pointer to **string** |  | [optional] 
**SourceCodeLocationUrl** | Pointer to **string** |  | [optional] 
**IssueTrackerUrl** | Pointer to **string** |  | [optional] 
**Valid** | Pointer to **bool** |  | [optional] 
**HasValidUpdate** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**Providers** | Pointer to [**[]GetAppState200ResponseAllOfSpecsInnerTemplate**](GetAppState200ResponseAllOfSpecsInnerTemplate.md) |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**OptionTypes** | Pointer to [**[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner**](ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner.md) |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewListPlugins200ResponseAllOfPluginsInner

`func NewListPlugins200ResponseAllOfPluginsInner() *ListPlugins200ResponseAllOfPluginsInner`

NewListPlugins200ResponseAllOfPluginsInner instantiates a new ListPlugins200ResponseAllOfPluginsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListPlugins200ResponseAllOfPluginsInnerWithDefaults

`func NewListPlugins200ResponseAllOfPluginsInnerWithDefaults() *ListPlugins200ResponseAllOfPluginsInner`

NewListPlugins200ResponseAllOfPluginsInnerWithDefaults instantiates a new ListPlugins200ResponseAllOfPluginsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListPlugins200ResponseAllOfPluginsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListPlugins200ResponseAllOfPluginsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListPlugins200ResponseAllOfPluginsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListPlugins200ResponseAllOfPluginsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListPlugins200ResponseAllOfPluginsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListPlugins200ResponseAllOfPluginsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListPlugins200ResponseAllOfPluginsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListPlugins200ResponseAllOfPluginsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *ListPlugins200ResponseAllOfPluginsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListPlugins200ResponseAllOfPluginsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListPlugins200ResponseAllOfPluginsInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListPlugins200ResponseAllOfPluginsInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *ListPlugins200ResponseAllOfPluginsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListPlugins200ResponseAllOfPluginsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListPlugins200ResponseAllOfPluginsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListPlugins200ResponseAllOfPluginsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersion

`func (o *ListPlugins200ResponseAllOfPluginsInner) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListPlugins200ResponseAllOfPluginsInner) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListPlugins200ResponseAllOfPluginsInner) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ListPlugins200ResponseAllOfPluginsInner) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetEnabled

`func (o *ListPlugins200ResponseAllOfPluginsInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListPlugins200ResponseAllOfPluginsInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListPlugins200ResponseAllOfPluginsInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListPlugins200ResponseAllOfPluginsInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetAuthor

`func (o *ListPlugins200ResponseAllOfPluginsInner) GetAuthor() string`

GetAuthor returns the Author field if non-nil, zero value otherwise.

### GetAuthorOk

`func (o *ListPlugins200ResponseAllOfPluginsInner) GetAuthorOk() (*string, bool)`

GetAuthorOk returns a tuple with the Author field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthor

`func (o *ListPlugins200ResponseAllOfPluginsInner) SetAuthor(v string)`

SetAuthor sets Author field to given value.

### HasAuthor

`func (o *ListPlugins200ResponseAllOfPluginsInner) HasAuthor() bool`

HasAuthor returns a boolean if a field has been set.

### GetWebsiteUrl

`func (o *ListPlugins200ResponseAllOfPluginsInner) GetWebsiteUrl() string`

GetWebsiteUrl returns the WebsiteUrl field if non-nil, zero value otherwise.

### GetWebsiteUrlOk

`func (o *ListPlugins200ResponseAllOfPluginsInner) GetWebsiteUrlOk() (*string, bool)`

GetWebsiteUrlOk returns a tuple with the WebsiteUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebsiteUrl

`func (o *ListPlugins200ResponseAllOfPluginsInner) SetWebsiteUrl(v string)`

SetWebsiteUrl sets WebsiteUrl field to given value.

### HasWebsiteUrl

`func (o *ListPlugins200ResponseAllOfPluginsInner) HasWebsiteUrl() bool`

HasWebsiteUrl returns a boolean if a field has been set.

### GetSourceCodeLocationUrl

`func (o *ListPlugins200ResponseAllOfPluginsInner) GetSourceCodeLocationUrl() string`

GetSourceCodeLocationUrl returns the SourceCodeLocationUrl field if non-nil, zero value otherwise.

### GetSourceCodeLocationUrlOk

`func (o *ListPlugins200ResponseAllOfPluginsInner) GetSourceCodeLocationUrlOk() (*string, bool)`

GetSourceCodeLocationUrlOk returns a tuple with the SourceCodeLocationUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceCodeLocationUrl

`func (o *ListPlugins200ResponseAllOfPluginsInner) SetSourceCodeLocationUrl(v string)`

SetSourceCodeLocationUrl sets SourceCodeLocationUrl field to given value.

### HasSourceCodeLocationUrl

`func (o *ListPlugins200ResponseAllOfPluginsInner) HasSourceCodeLocationUrl() bool`

HasSourceCodeLocationUrl returns a boolean if a field has been set.

### GetIssueTrackerUrl

`func (o *ListPlugins200ResponseAllOfPluginsInner) GetIssueTrackerUrl() string`

GetIssueTrackerUrl returns the IssueTrackerUrl field if non-nil, zero value otherwise.

### GetIssueTrackerUrlOk

`func (o *ListPlugins200ResponseAllOfPluginsInner) GetIssueTrackerUrlOk() (*string, bool)`

GetIssueTrackerUrlOk returns a tuple with the IssueTrackerUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIssueTrackerUrl

`func (o *ListPlugins200ResponseAllOfPluginsInner) SetIssueTrackerUrl(v string)`

SetIssueTrackerUrl sets IssueTrackerUrl field to given value.

### HasIssueTrackerUrl

`func (o *ListPlugins200ResponseAllOfPluginsInner) HasIssueTrackerUrl() bool`

HasIssueTrackerUrl returns a boolean if a field has been set.

### GetValid

`func (o *ListPlugins200ResponseAllOfPluginsInner) GetValid() bool`

GetValid returns the Valid field if non-nil, zero value otherwise.

### GetValidOk

`func (o *ListPlugins200ResponseAllOfPluginsInner) GetValidOk() (*bool, bool)`

GetValidOk returns a tuple with the Valid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValid

`func (o *ListPlugins200ResponseAllOfPluginsInner) SetValid(v bool)`

SetValid sets Valid field to given value.

### HasValid

`func (o *ListPlugins200ResponseAllOfPluginsInner) HasValid() bool`

HasValid returns a boolean if a field has been set.

### GetHasValidUpdate

`func (o *ListPlugins200ResponseAllOfPluginsInner) GetHasValidUpdate() bool`

GetHasValidUpdate returns the HasValidUpdate field if non-nil, zero value otherwise.

### GetHasValidUpdateOk

`func (o *ListPlugins200ResponseAllOfPluginsInner) GetHasValidUpdateOk() (*bool, bool)`

GetHasValidUpdateOk returns a tuple with the HasValidUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasValidUpdate

`func (o *ListPlugins200ResponseAllOfPluginsInner) SetHasValidUpdate(v bool)`

SetHasValidUpdate sets HasValidUpdate field to given value.

### HasHasValidUpdate

`func (o *ListPlugins200ResponseAllOfPluginsInner) HasHasValidUpdate() bool`

HasHasValidUpdate returns a boolean if a field has been set.

### GetStatus

`func (o *ListPlugins200ResponseAllOfPluginsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListPlugins200ResponseAllOfPluginsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListPlugins200ResponseAllOfPluginsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListPlugins200ResponseAllOfPluginsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ListPlugins200ResponseAllOfPluginsInner) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ListPlugins200ResponseAllOfPluginsInner) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ListPlugins200ResponseAllOfPluginsInner) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ListPlugins200ResponseAllOfPluginsInner) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetProviders

`func (o *ListPlugins200ResponseAllOfPluginsInner) GetProviders() []GetAppState200ResponseAllOfSpecsInnerTemplate`

GetProviders returns the Providers field if non-nil, zero value otherwise.

### GetProvidersOk

`func (o *ListPlugins200ResponseAllOfPluginsInner) GetProvidersOk() (*[]GetAppState200ResponseAllOfSpecsInnerTemplate, bool)`

GetProvidersOk returns a tuple with the Providers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProviders

`func (o *ListPlugins200ResponseAllOfPluginsInner) SetProviders(v []GetAppState200ResponseAllOfSpecsInnerTemplate)`

SetProviders sets Providers field to given value.

### HasProviders

`func (o *ListPlugins200ResponseAllOfPluginsInner) HasProviders() bool`

HasProviders returns a boolean if a field has been set.

### GetConfig

`func (o *ListPlugins200ResponseAllOfPluginsInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListPlugins200ResponseAllOfPluginsInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListPlugins200ResponseAllOfPluginsInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListPlugins200ResponseAllOfPluginsInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetOptionTypes

`func (o *ListPlugins200ResponseAllOfPluginsInner) GetOptionTypes() []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *ListPlugins200ResponseAllOfPluginsInner) GetOptionTypesOk() (*[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *ListPlugins200ResponseAllOfPluginsInner) SetOptionTypes(v []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *ListPlugins200ResponseAllOfPluginsInner) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListPlugins200ResponseAllOfPluginsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListPlugins200ResponseAllOfPluginsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListPlugins200ResponseAllOfPluginsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListPlugins200ResponseAllOfPluginsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListPlugins200ResponseAllOfPluginsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListPlugins200ResponseAllOfPluginsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListPlugins200ResponseAllOfPluginsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListPlugins200ResponseAllOfPluginsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


