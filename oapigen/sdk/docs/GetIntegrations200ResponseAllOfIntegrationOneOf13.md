# GetIntegrations200ResponseAllOfIntegrationOneOf13

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**IntegrationType** | Pointer to [**GetIntegrations200ResponseAllOfIntegrationOneOf13IntegrationType**](GetIntegrations200ResponseAllOfIntegrationOneOf13IntegrationType.md) |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**IsPlugin** | Pointer to **bool** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusDate** | Pointer to **time.Time** |  | [optional] 
**StatusMessage** | Pointer to **NullableString** |  | [optional] 
**LastSync** | Pointer to **NullableString** |  | [optional] 
**LastSyncDuration** | Pointer to **NullableString** |  | [optional] 
**Credential** | Pointer to [**GetIntegrations200ResponseAllOfIntegrationOneOf13Credential**](GetIntegrations200ResponseAllOfIntegrationOneOf13Credential.md) |  | [optional] 

## Methods

### NewGetIntegrations200ResponseAllOfIntegrationOneOf13

`func NewGetIntegrations200ResponseAllOfIntegrationOneOf13() *GetIntegrations200ResponseAllOfIntegrationOneOf13`

NewGetIntegrations200ResponseAllOfIntegrationOneOf13 instantiates a new GetIntegrations200ResponseAllOfIntegrationOneOf13 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIntegrations200ResponseAllOfIntegrationOneOf13WithDefaults

`func NewGetIntegrations200ResponseAllOfIntegrationOneOf13WithDefaults() *GetIntegrations200ResponseAllOfIntegrationOneOf13`

NewGetIntegrations200ResponseAllOfIntegrationOneOf13WithDefaults instantiates a new GetIntegrations200ResponseAllOfIntegrationOneOf13 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetType

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIntegrationType

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) GetIntegrationType() GetIntegrations200ResponseAllOfIntegrationOneOf13IntegrationType`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) GetIntegrationTypeOk() (*GetIntegrations200ResponseAllOfIntegrationOneOf13IntegrationType, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) SetIntegrationType(v GetIntegrations200ResponseAllOfIntegrationOneOf13IntegrationType)`

SetIntegrationType sets IntegrationType field to given value.

### HasIntegrationType

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) HasIntegrationType() bool`

HasIntegrationType returns a boolean if a field has been set.

### GetUrl

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetIsPlugin

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) GetIsPlugin() bool`

GetIsPlugin returns the IsPlugin field if non-nil, zero value otherwise.

### GetIsPluginOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) GetIsPluginOk() (*bool, bool)`

GetIsPluginOk returns a tuple with the IsPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlugin

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) SetIsPlugin(v bool)`

SetIsPlugin sets IsPlugin field to given value.

### HasIsPlugin

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) HasIsPlugin() bool`

HasIsPlugin returns a boolean if a field has been set.

### GetConfig

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetStatus

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDate

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetStatusMessage

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetLastSync

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) GetLastSync() string`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) GetLastSyncOk() (*string, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) SetLastSync(v string)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### SetLastSyncNil

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) SetLastSyncNil(b bool)`

 SetLastSyncNil sets the value for LastSync to be an explicit nil

### UnsetLastSync
`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) UnsetLastSync()`

UnsetLastSync ensures that no value is present for LastSync, not even an explicit nil
### GetLastSyncDuration

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) GetLastSyncDuration() string`

GetLastSyncDuration returns the LastSyncDuration field if non-nil, zero value otherwise.

### GetLastSyncDurationOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) GetLastSyncDurationOk() (*string, bool)`

GetLastSyncDurationOk returns a tuple with the LastSyncDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDuration

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) SetLastSyncDuration(v string)`

SetLastSyncDuration sets LastSyncDuration field to given value.

### HasLastSyncDuration

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) HasLastSyncDuration() bool`

HasLastSyncDuration returns a boolean if a field has been set.

### SetLastSyncDurationNil

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) SetLastSyncDurationNil(b bool)`

 SetLastSyncDurationNil sets the value for LastSyncDuration to be an explicit nil

### UnsetLastSyncDuration
`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) UnsetLastSyncDuration()`

UnsetLastSyncDuration ensures that no value is present for LastSyncDuration, not even an explicit nil
### GetCredential

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) GetCredential() GetIntegrations200ResponseAllOfIntegrationOneOf13Credential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) GetCredentialOk() (*GetIntegrations200ResponseAllOfIntegrationOneOf13Credential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) SetCredential(v GetIntegrations200ResponseAllOfIntegrationOneOf13Credential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf13) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


