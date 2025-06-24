# OptionTypeList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**SourceUrl** | Pointer to **string** |  | [optional] 
**SourceMethod** | Pointer to **string** |  | [optional] 
**ApiType** | Pointer to **NullableString** |  | [optional] 
**IgnoreSSLErrors** | Pointer to **bool** |  | [optional] 
**RealTime** | Pointer to **bool** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Config** | Pointer to [**ListOptionLists200ResponseAllOfOptionTypesInnerConfig**](ListOptionLists200ResponseAllOfOptionTypesInnerConfig.md) |  | [optional] 
**Credential** | Pointer to [**ListClouds200ResponseAllOfZonesInnerCredentialAnyOf**](ListClouds200ResponseAllOfZonesInnerCredentialAnyOf.md) |  | [optional] 
**ServiceUsername** | Pointer to **NullableString** |  | [optional] 
**ServicePassword** | Pointer to **NullableString** |  | [optional] 
**InitialDataset** | Pointer to **NullableString** |  | [optional] 
**TranslationScript** | Pointer to **string** |  | [optional] 
**RequestScript** | Pointer to **NullableString** |  | [optional] 
**Account** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 

## Methods

### NewOptionTypeList

`func NewOptionTypeList() *OptionTypeList`

NewOptionTypeList instantiates a new OptionTypeList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionTypeListWithDefaults

`func NewOptionTypeListWithDefaults() *OptionTypeList`

NewOptionTypeListWithDefaults instantiates a new OptionTypeList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OptionTypeList) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OptionTypeList) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OptionTypeList) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OptionTypeList) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OptionTypeList) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OptionTypeList) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OptionTypeList) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OptionTypeList) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *OptionTypeList) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OptionTypeList) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OptionTypeList) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OptionTypeList) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *OptionTypeList) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *OptionTypeList) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLabels

`func (o *OptionTypeList) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *OptionTypeList) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *OptionTypeList) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *OptionTypeList) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetType

`func (o *OptionTypeList) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *OptionTypeList) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *OptionTypeList) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *OptionTypeList) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSourceUrl

`func (o *OptionTypeList) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *OptionTypeList) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *OptionTypeList) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *OptionTypeList) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### GetSourceMethod

`func (o *OptionTypeList) GetSourceMethod() string`

GetSourceMethod returns the SourceMethod field if non-nil, zero value otherwise.

### GetSourceMethodOk

`func (o *OptionTypeList) GetSourceMethodOk() (*string, bool)`

GetSourceMethodOk returns a tuple with the SourceMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMethod

`func (o *OptionTypeList) SetSourceMethod(v string)`

SetSourceMethod sets SourceMethod field to given value.

### HasSourceMethod

`func (o *OptionTypeList) HasSourceMethod() bool`

HasSourceMethod returns a boolean if a field has been set.

### GetApiType

`func (o *OptionTypeList) GetApiType() string`

GetApiType returns the ApiType field if non-nil, zero value otherwise.

### GetApiTypeOk

`func (o *OptionTypeList) GetApiTypeOk() (*string, bool)`

GetApiTypeOk returns a tuple with the ApiType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiType

`func (o *OptionTypeList) SetApiType(v string)`

SetApiType sets ApiType field to given value.

### HasApiType

`func (o *OptionTypeList) HasApiType() bool`

HasApiType returns a boolean if a field has been set.

### SetApiTypeNil

`func (o *OptionTypeList) SetApiTypeNil(b bool)`

 SetApiTypeNil sets the value for ApiType to be an explicit nil

### UnsetApiType
`func (o *OptionTypeList) UnsetApiType()`

UnsetApiType ensures that no value is present for ApiType, not even an explicit nil
### GetIgnoreSSLErrors

`func (o *OptionTypeList) GetIgnoreSSLErrors() bool`

GetIgnoreSSLErrors returns the IgnoreSSLErrors field if non-nil, zero value otherwise.

### GetIgnoreSSLErrorsOk

`func (o *OptionTypeList) GetIgnoreSSLErrorsOk() (*bool, bool)`

GetIgnoreSSLErrorsOk returns a tuple with the IgnoreSSLErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSSLErrors

`func (o *OptionTypeList) SetIgnoreSSLErrors(v bool)`

SetIgnoreSSLErrors sets IgnoreSSLErrors field to given value.

### HasIgnoreSSLErrors

`func (o *OptionTypeList) HasIgnoreSSLErrors() bool`

HasIgnoreSSLErrors returns a boolean if a field has been set.

### GetRealTime

`func (o *OptionTypeList) GetRealTime() bool`

GetRealTime returns the RealTime field if non-nil, zero value otherwise.

### GetRealTimeOk

`func (o *OptionTypeList) GetRealTimeOk() (*bool, bool)`

GetRealTimeOk returns a tuple with the RealTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealTime

`func (o *OptionTypeList) SetRealTime(v bool)`

SetRealTime sets RealTime field to given value.

### HasRealTime

`func (o *OptionTypeList) HasRealTime() bool`

HasRealTime returns a boolean if a field has been set.

### GetVisibility

`func (o *OptionTypeList) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *OptionTypeList) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *OptionTypeList) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *OptionTypeList) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetConfig

`func (o *OptionTypeList) GetConfig() ListOptionLists200ResponseAllOfOptionTypesInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *OptionTypeList) GetConfigOk() (*ListOptionLists200ResponseAllOfOptionTypesInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *OptionTypeList) SetConfig(v ListOptionLists200ResponseAllOfOptionTypesInnerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *OptionTypeList) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCredential

`func (o *OptionTypeList) GetCredential() ListClouds200ResponseAllOfZonesInnerCredentialAnyOf`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *OptionTypeList) GetCredentialOk() (*ListClouds200ResponseAllOfZonesInnerCredentialAnyOf, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *OptionTypeList) SetCredential(v ListClouds200ResponseAllOfZonesInnerCredentialAnyOf)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *OptionTypeList) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetServiceUsername

`func (o *OptionTypeList) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *OptionTypeList) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *OptionTypeList) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.

### HasServiceUsername

`func (o *OptionTypeList) HasServiceUsername() bool`

HasServiceUsername returns a boolean if a field has been set.

### SetServiceUsernameNil

`func (o *OptionTypeList) SetServiceUsernameNil(b bool)`

 SetServiceUsernameNil sets the value for ServiceUsername to be an explicit nil

### UnsetServiceUsername
`func (o *OptionTypeList) UnsetServiceUsername()`

UnsetServiceUsername ensures that no value is present for ServiceUsername, not even an explicit nil
### GetServicePassword

`func (o *OptionTypeList) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *OptionTypeList) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *OptionTypeList) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *OptionTypeList) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### SetServicePasswordNil

`func (o *OptionTypeList) SetServicePasswordNil(b bool)`

 SetServicePasswordNil sets the value for ServicePassword to be an explicit nil

### UnsetServicePassword
`func (o *OptionTypeList) UnsetServicePassword()`

UnsetServicePassword ensures that no value is present for ServicePassword, not even an explicit nil
### GetInitialDataset

`func (o *OptionTypeList) GetInitialDataset() string`

GetInitialDataset returns the InitialDataset field if non-nil, zero value otherwise.

### GetInitialDatasetOk

`func (o *OptionTypeList) GetInitialDatasetOk() (*string, bool)`

GetInitialDatasetOk returns a tuple with the InitialDataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialDataset

`func (o *OptionTypeList) SetInitialDataset(v string)`

SetInitialDataset sets InitialDataset field to given value.

### HasInitialDataset

`func (o *OptionTypeList) HasInitialDataset() bool`

HasInitialDataset returns a boolean if a field has been set.

### SetInitialDatasetNil

`func (o *OptionTypeList) SetInitialDatasetNil(b bool)`

 SetInitialDatasetNil sets the value for InitialDataset to be an explicit nil

### UnsetInitialDataset
`func (o *OptionTypeList) UnsetInitialDataset()`

UnsetInitialDataset ensures that no value is present for InitialDataset, not even an explicit nil
### GetTranslationScript

`func (o *OptionTypeList) GetTranslationScript() string`

GetTranslationScript returns the TranslationScript field if non-nil, zero value otherwise.

### GetTranslationScriptOk

`func (o *OptionTypeList) GetTranslationScriptOk() (*string, bool)`

GetTranslationScriptOk returns a tuple with the TranslationScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslationScript

`func (o *OptionTypeList) SetTranslationScript(v string)`

SetTranslationScript sets TranslationScript field to given value.

### HasTranslationScript

`func (o *OptionTypeList) HasTranslationScript() bool`

HasTranslationScript returns a boolean if a field has been set.

### GetRequestScript

`func (o *OptionTypeList) GetRequestScript() string`

GetRequestScript returns the RequestScript field if non-nil, zero value otherwise.

### GetRequestScriptOk

`func (o *OptionTypeList) GetRequestScriptOk() (*string, bool)`

GetRequestScriptOk returns a tuple with the RequestScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestScript

`func (o *OptionTypeList) SetRequestScript(v string)`

SetRequestScript sets RequestScript field to given value.

### HasRequestScript

`func (o *OptionTypeList) HasRequestScript() bool`

HasRequestScript returns a boolean if a field has been set.

### SetRequestScriptNil

`func (o *OptionTypeList) SetRequestScriptNil(b bool)`

 SetRequestScriptNil sets the value for RequestScript to be an explicit nil

### UnsetRequestScript
`func (o *OptionTypeList) UnsetRequestScript()`

UnsetRequestScript ensures that no value is present for RequestScript, not even an explicit nil
### GetAccount

`func (o *OptionTypeList) GetAccount() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *OptionTypeList) GetAccountOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *OptionTypeList) SetAccount(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *OptionTypeList) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


