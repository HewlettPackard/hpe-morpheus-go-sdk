# AddIntegrations200ResponseAllOfIntegrationOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**IntegrationType** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**ServiceKey** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**IsPlugin** | Pointer to **bool** |  | [optional] 
**Config** | Pointer to [**ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfConfig**](ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfConfig.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusDate** | Pointer to **time.Time** |  | [optional] 
**StatusMessage** | Pointer to **NullableString** |  | [optional] 
**LastSync** | Pointer to **NullableString** |  | [optional] 
**LastSyncDuration** | Pointer to **NullableInt64** |  | [optional] 
**Credential** | Pointer to [**ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential**](ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential.md) |  | [optional] 

## Methods

### NewAddIntegrations200ResponseAllOfIntegrationOneOf

`func NewAddIntegrations200ResponseAllOfIntegrationOneOf() *AddIntegrations200ResponseAllOfIntegrationOneOf`

NewAddIntegrations200ResponseAllOfIntegrationOneOf instantiates a new AddIntegrations200ResponseAllOfIntegrationOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddIntegrations200ResponseAllOfIntegrationOneOfWithDefaults

`func NewAddIntegrations200ResponseAllOfIntegrationOneOfWithDefaults() *AddIntegrations200ResponseAllOfIntegrationOneOf`

NewAddIntegrations200ResponseAllOfIntegrationOneOfWithDefaults instantiates a new AddIntegrations200ResponseAllOfIntegrationOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetType

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIntegrationType

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) GetIntegrationType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) GetIntegrationTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) SetIntegrationType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetIntegrationType sets IntegrationType field to given value.

### HasIntegrationType

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) HasIntegrationType() bool`

HasIntegrationType returns a boolean if a field has been set.

### GetUrl

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetServiceKey

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) GetServiceKey() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) GetServiceKeyOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) SetServiceKey(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetServiceKey sets ServiceKey field to given value.

### HasServiceKey

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) HasServiceKey() bool`

HasServiceKey returns a boolean if a field has been set.

### GetIsPlugin

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) GetIsPlugin() bool`

GetIsPlugin returns the IsPlugin field if non-nil, zero value otherwise.

### GetIsPluginOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) GetIsPluginOk() (*bool, bool)`

GetIsPluginOk returns a tuple with the IsPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlugin

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) SetIsPlugin(v bool)`

SetIsPlugin sets IsPlugin field to given value.

### HasIsPlugin

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) HasIsPlugin() bool`

HasIsPlugin returns a boolean if a field has been set.

### GetConfig

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) GetConfig() ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) GetConfigOk() (*ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) SetConfig(v ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStatus

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDate

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetStatusMessage

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetLastSync

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) GetLastSync() string`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) GetLastSyncOk() (*string, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) SetLastSync(v string)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### SetLastSyncNil

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) SetLastSyncNil(b bool)`

 SetLastSyncNil sets the value for LastSync to be an explicit nil

### UnsetLastSync
`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) UnsetLastSync()`

UnsetLastSync ensures that no value is present for LastSync, not even an explicit nil
### GetLastSyncDuration

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) GetLastSyncDuration() int64`

GetLastSyncDuration returns the LastSyncDuration field if non-nil, zero value otherwise.

### GetLastSyncDurationOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) GetLastSyncDurationOk() (*int64, bool)`

GetLastSyncDurationOk returns a tuple with the LastSyncDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDuration

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) SetLastSyncDuration(v int64)`

SetLastSyncDuration sets LastSyncDuration field to given value.

### HasLastSyncDuration

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) HasLastSyncDuration() bool`

HasLastSyncDuration returns a boolean if a field has been set.

### SetLastSyncDurationNil

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) SetLastSyncDurationNil(b bool)`

 SetLastSyncDurationNil sets the value for LastSyncDuration to be an explicit nil

### UnsetLastSyncDuration
`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) UnsetLastSyncDuration()`

UnsetLastSyncDuration ensures that no value is present for LastSyncDuration, not even an explicit nil
### GetCredential

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) GetCredential() ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) GetCredentialOk() (*ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) SetCredential(v ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


