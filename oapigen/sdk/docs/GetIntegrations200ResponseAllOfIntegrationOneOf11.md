# GetIntegrations200ResponseAllOfIntegrationOneOf11

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**IntegrationType** | Pointer to [**GetIntegrations200ResponseAllOfIntegrationOneOf11IntegrationType**](GetIntegrations200ResponseAllOfIntegrationOneOf11IntegrationType.md) |  | [optional] 
**IsPlugin** | Pointer to **bool** |  | [optional] 
**Config** | Pointer to [**GetIntegrations200ResponseAllOfIntegrationOneOf11Config**](GetIntegrations200ResponseAllOfIntegrationOneOf11Config.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusDate** | Pointer to **NullableTime** |  | [optional] 
**StatusMessage** | Pointer to **NullableString** |  | [optional] 
**LastSync** | Pointer to **NullableString** |  | [optional] 
**LastSyncDuration** | Pointer to **NullableString** |  | [optional] 
**Credential** | Pointer to [**GetIntegrations200ResponseAllOfIntegrationOneOf11Credential**](GetIntegrations200ResponseAllOfIntegrationOneOf11Credential.md) |  | [optional] 

## Methods

### NewGetIntegrations200ResponseAllOfIntegrationOneOf11

`func NewGetIntegrations200ResponseAllOfIntegrationOneOf11() *GetIntegrations200ResponseAllOfIntegrationOneOf11`

NewGetIntegrations200ResponseAllOfIntegrationOneOf11 instantiates a new GetIntegrations200ResponseAllOfIntegrationOneOf11 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetType

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIntegrationType

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) GetIntegrationType() GetIntegrations200ResponseAllOfIntegrationOneOf11IntegrationType`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) GetIntegrationTypeOk() (*GetIntegrations200ResponseAllOfIntegrationOneOf11IntegrationType, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) SetIntegrationType(v GetIntegrations200ResponseAllOfIntegrationOneOf11IntegrationType)`

SetIntegrationType sets IntegrationType field to given value.

### HasIntegrationType

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) HasIntegrationType() bool`

HasIntegrationType returns a boolean if a field has been set.

### GetIsPlugin

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) GetIsPlugin() bool`

GetIsPlugin returns the IsPlugin field if non-nil, zero value otherwise.

### GetIsPluginOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) GetIsPluginOk() (*bool, bool)`

GetIsPluginOk returns a tuple with the IsPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlugin

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) SetIsPlugin(v bool)`

SetIsPlugin sets IsPlugin field to given value.

### HasIsPlugin

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) HasIsPlugin() bool`

HasIsPlugin returns a boolean if a field has been set.

### GetConfig

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) GetConfig() GetIntegrations200ResponseAllOfIntegrationOneOf11Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) GetConfigOk() (*GetIntegrations200ResponseAllOfIntegrationOneOf11Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) SetConfig(v GetIntegrations200ResponseAllOfIntegrationOneOf11Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStatus

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDate

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### SetStatusDateNil

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) SetStatusDateNil(b bool)`

 SetStatusDateNil sets the value for StatusDate to be an explicit nil

### UnsetStatusDate
`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) UnsetStatusDate()`

UnsetStatusDate ensures that no value is present for StatusDate, not even an explicit nil
### GetStatusMessage

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetLastSync

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) GetLastSync() string`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) GetLastSyncOk() (*string, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) SetLastSync(v string)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### SetLastSyncNil

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) SetLastSyncNil(b bool)`

 SetLastSyncNil sets the value for LastSync to be an explicit nil

### UnsetLastSync
`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) UnsetLastSync()`

UnsetLastSync ensures that no value is present for LastSync, not even an explicit nil
### GetLastSyncDuration

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) GetLastSyncDuration() string`

GetLastSyncDuration returns the LastSyncDuration field if non-nil, zero value otherwise.

### GetLastSyncDurationOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) GetLastSyncDurationOk() (*string, bool)`

GetLastSyncDurationOk returns a tuple with the LastSyncDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDuration

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) SetLastSyncDuration(v string)`

SetLastSyncDuration sets LastSyncDuration field to given value.

### HasLastSyncDuration

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) HasLastSyncDuration() bool`

HasLastSyncDuration returns a boolean if a field has been set.

### SetLastSyncDurationNil

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) SetLastSyncDurationNil(b bool)`

 SetLastSyncDurationNil sets the value for LastSyncDuration to be an explicit nil

### UnsetLastSyncDuration
`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) UnsetLastSyncDuration()`

UnsetLastSyncDuration ensures that no value is present for LastSyncDuration, not even an explicit nil
### GetCredential

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) GetCredential() GetIntegrations200ResponseAllOfIntegrationOneOf11Credential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) GetCredentialOk() (*GetIntegrations200ResponseAllOfIntegrationOneOf11Credential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) SetCredential(v GetIntegrations200ResponseAllOfIntegrationOneOf11Credential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf11) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


