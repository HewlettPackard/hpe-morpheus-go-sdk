# ListOptionLists200ResponseAllOfOptionTypesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**SourceUrl** | Pointer to **string** |  | [optional] 
**SourceMethod** | Pointer to **string** |  | [optional] 
**ApiType** | Pointer to **string** |  | [optional] 
**IgnoreSSLErrors** | Pointer to **bool** |  | [optional] 
**RealTime** | Pointer to **bool** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Config** | Pointer to [**ListOptionLists200ResponseAllOfOptionTypesInnerConfig**](ListOptionLists200ResponseAllOfOptionTypesInnerConfig.md) |  | [optional] 
**Credential** | Pointer to [**ListClouds200ResponseAllOfZonesInnerCredentialAnyOf**](ListClouds200ResponseAllOfZonesInnerCredentialAnyOf.md) |  | [optional] 
**ServiceUsername** | Pointer to **string** |  | [optional] 
**ServicePassword** | Pointer to **string** |  | [optional] 
**InitialDataset** | Pointer to **string** |  | [optional] 
**TranslationScript** | Pointer to **string** |  | [optional] 
**RequestScript** | Pointer to **string** |  | [optional] 
**Account** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 

## Methods

### NewListOptionLists200ResponseAllOfOptionTypesInner

`func NewListOptionLists200ResponseAllOfOptionTypesInner() *ListOptionLists200ResponseAllOfOptionTypesInner`

NewListOptionLists200ResponseAllOfOptionTypesInner instantiates a new ListOptionLists200ResponseAllOfOptionTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListOptionLists200ResponseAllOfOptionTypesInnerWithDefaults

`func NewListOptionLists200ResponseAllOfOptionTypesInnerWithDefaults() *ListOptionLists200ResponseAllOfOptionTypesInner`

NewListOptionLists200ResponseAllOfOptionTypesInnerWithDefaults instantiates a new ListOptionLists200ResponseAllOfOptionTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetType

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSourceUrl

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### GetSourceMethod

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) GetSourceMethod() string`

GetSourceMethod returns the SourceMethod field if non-nil, zero value otherwise.

### GetSourceMethodOk

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) GetSourceMethodOk() (*string, bool)`

GetSourceMethodOk returns a tuple with the SourceMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMethod

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) SetSourceMethod(v string)`

SetSourceMethod sets SourceMethod field to given value.

### HasSourceMethod

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) HasSourceMethod() bool`

HasSourceMethod returns a boolean if a field has been set.

### GetApiType

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) GetApiType() string`

GetApiType returns the ApiType field if non-nil, zero value otherwise.

### GetApiTypeOk

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) GetApiTypeOk() (*string, bool)`

GetApiTypeOk returns a tuple with the ApiType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiType

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) SetApiType(v string)`

SetApiType sets ApiType field to given value.

### HasApiType

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) HasApiType() bool`

HasApiType returns a boolean if a field has been set.

### GetIgnoreSSLErrors

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) GetIgnoreSSLErrors() bool`

GetIgnoreSSLErrors returns the IgnoreSSLErrors field if non-nil, zero value otherwise.

### GetIgnoreSSLErrorsOk

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) GetIgnoreSSLErrorsOk() (*bool, bool)`

GetIgnoreSSLErrorsOk returns a tuple with the IgnoreSSLErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSSLErrors

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) SetIgnoreSSLErrors(v bool)`

SetIgnoreSSLErrors sets IgnoreSSLErrors field to given value.

### HasIgnoreSSLErrors

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) HasIgnoreSSLErrors() bool`

HasIgnoreSSLErrors returns a boolean if a field has been set.

### GetRealTime

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) GetRealTime() bool`

GetRealTime returns the RealTime field if non-nil, zero value otherwise.

### GetRealTimeOk

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) GetRealTimeOk() (*bool, bool)`

GetRealTimeOk returns a tuple with the RealTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealTime

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) SetRealTime(v bool)`

SetRealTime sets RealTime field to given value.

### HasRealTime

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) HasRealTime() bool`

HasRealTime returns a boolean if a field has been set.

### GetVisibility

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetConfig

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) GetConfig() ListOptionLists200ResponseAllOfOptionTypesInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) GetConfigOk() (*ListOptionLists200ResponseAllOfOptionTypesInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) SetConfig(v ListOptionLists200ResponseAllOfOptionTypesInnerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCredential

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) GetCredential() ListClouds200ResponseAllOfZonesInnerCredentialAnyOf`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) GetCredentialOk() (*ListClouds200ResponseAllOfZonesInnerCredentialAnyOf, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) SetCredential(v ListClouds200ResponseAllOfZonesInnerCredentialAnyOf)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetServiceUsername

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.

### HasServiceUsername

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) HasServiceUsername() bool`

HasServiceUsername returns a boolean if a field has been set.

### GetServicePassword

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### GetInitialDataset

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) GetInitialDataset() string`

GetInitialDataset returns the InitialDataset field if non-nil, zero value otherwise.

### GetInitialDatasetOk

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) GetInitialDatasetOk() (*string, bool)`

GetInitialDatasetOk returns a tuple with the InitialDataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialDataset

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) SetInitialDataset(v string)`

SetInitialDataset sets InitialDataset field to given value.

### HasInitialDataset

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) HasInitialDataset() bool`

HasInitialDataset returns a boolean if a field has been set.

### GetTranslationScript

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) GetTranslationScript() string`

GetTranslationScript returns the TranslationScript field if non-nil, zero value otherwise.

### GetTranslationScriptOk

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) GetTranslationScriptOk() (*string, bool)`

GetTranslationScriptOk returns a tuple with the TranslationScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslationScript

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) SetTranslationScript(v string)`

SetTranslationScript sets TranslationScript field to given value.

### HasTranslationScript

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) HasTranslationScript() bool`

HasTranslationScript returns a boolean if a field has been set.

### GetRequestScript

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) GetRequestScript() string`

GetRequestScript returns the RequestScript field if non-nil, zero value otherwise.

### GetRequestScriptOk

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) GetRequestScriptOk() (*string, bool)`

GetRequestScriptOk returns a tuple with the RequestScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestScript

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) SetRequestScript(v string)`

SetRequestScript sets RequestScript field to given value.

### HasRequestScript

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) HasRequestScript() bool`

HasRequestScript returns a boolean if a field has been set.

### GetAccount

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) GetAccount() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) GetAccountOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) SetAccount(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListOptionLists200ResponseAllOfOptionTypesInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


