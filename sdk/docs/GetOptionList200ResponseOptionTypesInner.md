# GetOptionList200ResponseOptionTypesInner

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

### NewGetOptionList200ResponseOptionTypesInner

`func NewGetOptionList200ResponseOptionTypesInner() *GetOptionList200ResponseOptionTypesInner`

NewGetOptionList200ResponseOptionTypesInner instantiates a new GetOptionList200ResponseOptionTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetOptionList200ResponseOptionTypesInnerWithDefaults

`func NewGetOptionList200ResponseOptionTypesInnerWithDefaults() *GetOptionList200ResponseOptionTypesInner`

NewGetOptionList200ResponseOptionTypesInnerWithDefaults instantiates a new GetOptionList200ResponseOptionTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetOptionList200ResponseOptionTypesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetOptionList200ResponseOptionTypesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetOptionList200ResponseOptionTypesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetOptionList200ResponseOptionTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetOptionList200ResponseOptionTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetOptionList200ResponseOptionTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetOptionList200ResponseOptionTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetOptionList200ResponseOptionTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetOptionList200ResponseOptionTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetOptionList200ResponseOptionTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetOptionList200ResponseOptionTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetOptionList200ResponseOptionTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetOptionList200ResponseOptionTypesInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetOptionList200ResponseOptionTypesInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLabels

`func (o *GetOptionList200ResponseOptionTypesInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetOptionList200ResponseOptionTypesInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetOptionList200ResponseOptionTypesInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetOptionList200ResponseOptionTypesInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetType

`func (o *GetOptionList200ResponseOptionTypesInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetOptionList200ResponseOptionTypesInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetOptionList200ResponseOptionTypesInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetOptionList200ResponseOptionTypesInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetSourceUrl

`func (o *GetOptionList200ResponseOptionTypesInner) GetSourceUrl() string`

GetSourceUrl returns the SourceUrl field if non-nil, zero value otherwise.

### GetSourceUrlOk

`func (o *GetOptionList200ResponseOptionTypesInner) GetSourceUrlOk() (*string, bool)`

GetSourceUrlOk returns a tuple with the SourceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceUrl

`func (o *GetOptionList200ResponseOptionTypesInner) SetSourceUrl(v string)`

SetSourceUrl sets SourceUrl field to given value.

### HasSourceUrl

`func (o *GetOptionList200ResponseOptionTypesInner) HasSourceUrl() bool`

HasSourceUrl returns a boolean if a field has been set.

### GetSourceMethod

`func (o *GetOptionList200ResponseOptionTypesInner) GetSourceMethod() string`

GetSourceMethod returns the SourceMethod field if non-nil, zero value otherwise.

### GetSourceMethodOk

`func (o *GetOptionList200ResponseOptionTypesInner) GetSourceMethodOk() (*string, bool)`

GetSourceMethodOk returns a tuple with the SourceMethod field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSourceMethod

`func (o *GetOptionList200ResponseOptionTypesInner) SetSourceMethod(v string)`

SetSourceMethod sets SourceMethod field to given value.

### HasSourceMethod

`func (o *GetOptionList200ResponseOptionTypesInner) HasSourceMethod() bool`

HasSourceMethod returns a boolean if a field has been set.

### GetApiType

`func (o *GetOptionList200ResponseOptionTypesInner) GetApiType() string`

GetApiType returns the ApiType field if non-nil, zero value otherwise.

### GetApiTypeOk

`func (o *GetOptionList200ResponseOptionTypesInner) GetApiTypeOk() (*string, bool)`

GetApiTypeOk returns a tuple with the ApiType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiType

`func (o *GetOptionList200ResponseOptionTypesInner) SetApiType(v string)`

SetApiType sets ApiType field to given value.

### HasApiType

`func (o *GetOptionList200ResponseOptionTypesInner) HasApiType() bool`

HasApiType returns a boolean if a field has been set.

### SetApiTypeNil

`func (o *GetOptionList200ResponseOptionTypesInner) SetApiTypeNil(b bool)`

 SetApiTypeNil sets the value for ApiType to be an explicit nil

### UnsetApiType
`func (o *GetOptionList200ResponseOptionTypesInner) UnsetApiType()`

UnsetApiType ensures that no value is present for ApiType, not even an explicit nil
### GetIgnoreSSLErrors

`func (o *GetOptionList200ResponseOptionTypesInner) GetIgnoreSSLErrors() bool`

GetIgnoreSSLErrors returns the IgnoreSSLErrors field if non-nil, zero value otherwise.

### GetIgnoreSSLErrorsOk

`func (o *GetOptionList200ResponseOptionTypesInner) GetIgnoreSSLErrorsOk() (*bool, bool)`

GetIgnoreSSLErrorsOk returns a tuple with the IgnoreSSLErrors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSSLErrors

`func (o *GetOptionList200ResponseOptionTypesInner) SetIgnoreSSLErrors(v bool)`

SetIgnoreSSLErrors sets IgnoreSSLErrors field to given value.

### HasIgnoreSSLErrors

`func (o *GetOptionList200ResponseOptionTypesInner) HasIgnoreSSLErrors() bool`

HasIgnoreSSLErrors returns a boolean if a field has been set.

### GetRealTime

`func (o *GetOptionList200ResponseOptionTypesInner) GetRealTime() bool`

GetRealTime returns the RealTime field if non-nil, zero value otherwise.

### GetRealTimeOk

`func (o *GetOptionList200ResponseOptionTypesInner) GetRealTimeOk() (*bool, bool)`

GetRealTimeOk returns a tuple with the RealTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRealTime

`func (o *GetOptionList200ResponseOptionTypesInner) SetRealTime(v bool)`

SetRealTime sets RealTime field to given value.

### HasRealTime

`func (o *GetOptionList200ResponseOptionTypesInner) HasRealTime() bool`

HasRealTime returns a boolean if a field has been set.

### GetVisibility

`func (o *GetOptionList200ResponseOptionTypesInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetOptionList200ResponseOptionTypesInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetOptionList200ResponseOptionTypesInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetOptionList200ResponseOptionTypesInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetConfig

`func (o *GetOptionList200ResponseOptionTypesInner) GetConfig() ListOptionLists200ResponseAllOfOptionTypesInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetOptionList200ResponseOptionTypesInner) GetConfigOk() (*ListOptionLists200ResponseAllOfOptionTypesInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetOptionList200ResponseOptionTypesInner) SetConfig(v ListOptionLists200ResponseAllOfOptionTypesInnerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetOptionList200ResponseOptionTypesInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCredential

`func (o *GetOptionList200ResponseOptionTypesInner) GetCredential() ListClouds200ResponseAllOfZonesInnerCredentialAnyOf`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *GetOptionList200ResponseOptionTypesInner) GetCredentialOk() (*ListClouds200ResponseAllOfZonesInnerCredentialAnyOf, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *GetOptionList200ResponseOptionTypesInner) SetCredential(v ListClouds200ResponseAllOfZonesInnerCredentialAnyOf)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *GetOptionList200ResponseOptionTypesInner) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetServiceUsername

`func (o *GetOptionList200ResponseOptionTypesInner) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *GetOptionList200ResponseOptionTypesInner) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *GetOptionList200ResponseOptionTypesInner) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.

### HasServiceUsername

`func (o *GetOptionList200ResponseOptionTypesInner) HasServiceUsername() bool`

HasServiceUsername returns a boolean if a field has been set.

### SetServiceUsernameNil

`func (o *GetOptionList200ResponseOptionTypesInner) SetServiceUsernameNil(b bool)`

 SetServiceUsernameNil sets the value for ServiceUsername to be an explicit nil

### UnsetServiceUsername
`func (o *GetOptionList200ResponseOptionTypesInner) UnsetServiceUsername()`

UnsetServiceUsername ensures that no value is present for ServiceUsername, not even an explicit nil
### GetServicePassword

`func (o *GetOptionList200ResponseOptionTypesInner) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *GetOptionList200ResponseOptionTypesInner) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *GetOptionList200ResponseOptionTypesInner) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *GetOptionList200ResponseOptionTypesInner) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### SetServicePasswordNil

`func (o *GetOptionList200ResponseOptionTypesInner) SetServicePasswordNil(b bool)`

 SetServicePasswordNil sets the value for ServicePassword to be an explicit nil

### UnsetServicePassword
`func (o *GetOptionList200ResponseOptionTypesInner) UnsetServicePassword()`

UnsetServicePassword ensures that no value is present for ServicePassword, not even an explicit nil
### GetInitialDataset

`func (o *GetOptionList200ResponseOptionTypesInner) GetInitialDataset() string`

GetInitialDataset returns the InitialDataset field if non-nil, zero value otherwise.

### GetInitialDatasetOk

`func (o *GetOptionList200ResponseOptionTypesInner) GetInitialDatasetOk() (*string, bool)`

GetInitialDatasetOk returns a tuple with the InitialDataset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialDataset

`func (o *GetOptionList200ResponseOptionTypesInner) SetInitialDataset(v string)`

SetInitialDataset sets InitialDataset field to given value.

### HasInitialDataset

`func (o *GetOptionList200ResponseOptionTypesInner) HasInitialDataset() bool`

HasInitialDataset returns a boolean if a field has been set.

### SetInitialDatasetNil

`func (o *GetOptionList200ResponseOptionTypesInner) SetInitialDatasetNil(b bool)`

 SetInitialDatasetNil sets the value for InitialDataset to be an explicit nil

### UnsetInitialDataset
`func (o *GetOptionList200ResponseOptionTypesInner) UnsetInitialDataset()`

UnsetInitialDataset ensures that no value is present for InitialDataset, not even an explicit nil
### GetTranslationScript

`func (o *GetOptionList200ResponseOptionTypesInner) GetTranslationScript() string`

GetTranslationScript returns the TranslationScript field if non-nil, zero value otherwise.

### GetTranslationScriptOk

`func (o *GetOptionList200ResponseOptionTypesInner) GetTranslationScriptOk() (*string, bool)`

GetTranslationScriptOk returns a tuple with the TranslationScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTranslationScript

`func (o *GetOptionList200ResponseOptionTypesInner) SetTranslationScript(v string)`

SetTranslationScript sets TranslationScript field to given value.

### HasTranslationScript

`func (o *GetOptionList200ResponseOptionTypesInner) HasTranslationScript() bool`

HasTranslationScript returns a boolean if a field has been set.

### GetRequestScript

`func (o *GetOptionList200ResponseOptionTypesInner) GetRequestScript() string`

GetRequestScript returns the RequestScript field if non-nil, zero value otherwise.

### GetRequestScriptOk

`func (o *GetOptionList200ResponseOptionTypesInner) GetRequestScriptOk() (*string, bool)`

GetRequestScriptOk returns a tuple with the RequestScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestScript

`func (o *GetOptionList200ResponseOptionTypesInner) SetRequestScript(v string)`

SetRequestScript sets RequestScript field to given value.

### HasRequestScript

`func (o *GetOptionList200ResponseOptionTypesInner) HasRequestScript() bool`

HasRequestScript returns a boolean if a field has been set.

### SetRequestScriptNil

`func (o *GetOptionList200ResponseOptionTypesInner) SetRequestScriptNil(b bool)`

 SetRequestScriptNil sets the value for RequestScript to be an explicit nil

### UnsetRequestScript
`func (o *GetOptionList200ResponseOptionTypesInner) UnsetRequestScript()`

UnsetRequestScript ensures that no value is present for RequestScript, not even an explicit nil
### GetAccount

`func (o *GetOptionList200ResponseOptionTypesInner) GetAccount() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetOptionList200ResponseOptionTypesInner) GetAccountOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetOptionList200ResponseOptionTypesInner) SetAccount(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetOptionList200ResponseOptionTypesInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


