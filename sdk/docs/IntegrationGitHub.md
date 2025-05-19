# IntegrationGitHub

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**IntegrationType** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**TokenHash** | Pointer to **string** |  | [optional] 
**ServiceKey** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**IsPlugin** | Pointer to **bool** |  | [optional] 
**Config** | Pointer to [**ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf7Config**](ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf7Config.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusDate** | Pointer to **time.Time** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**LastSync** | Pointer to **string** |  | [optional] 
**LastSyncDuration** | Pointer to **string** |  | [optional] 
**Credential** | Pointer to [**ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential**](ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential.md) |  | [optional] 

## Methods

### NewIntegrationGitHub

`func NewIntegrationGitHub() *IntegrationGitHub`

NewIntegrationGitHub instantiates a new IntegrationGitHub object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIntegrationGitHubWithDefaults

`func NewIntegrationGitHubWithDefaults() *IntegrationGitHub`

NewIntegrationGitHubWithDefaults instantiates a new IntegrationGitHub object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *IntegrationGitHub) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *IntegrationGitHub) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *IntegrationGitHub) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *IntegrationGitHub) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *IntegrationGitHub) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *IntegrationGitHub) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *IntegrationGitHub) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *IntegrationGitHub) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *IntegrationGitHub) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *IntegrationGitHub) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *IntegrationGitHub) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *IntegrationGitHub) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetType

`func (o *IntegrationGitHub) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *IntegrationGitHub) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *IntegrationGitHub) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *IntegrationGitHub) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIntegrationType

`func (o *IntegrationGitHub) GetIntegrationType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *IntegrationGitHub) GetIntegrationTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *IntegrationGitHub) SetIntegrationType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetIntegrationType sets IntegrationType field to given value.

### HasIntegrationType

`func (o *IntegrationGitHub) HasIntegrationType() bool`

HasIntegrationType returns a boolean if a field has been set.

### GetUsername

`func (o *IntegrationGitHub) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *IntegrationGitHub) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *IntegrationGitHub) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *IntegrationGitHub) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetToken

`func (o *IntegrationGitHub) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *IntegrationGitHub) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *IntegrationGitHub) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *IntegrationGitHub) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTokenHash

`func (o *IntegrationGitHub) GetTokenHash() string`

GetTokenHash returns the TokenHash field if non-nil, zero value otherwise.

### GetTokenHashOk

`func (o *IntegrationGitHub) GetTokenHashOk() (*string, bool)`

GetTokenHashOk returns a tuple with the TokenHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenHash

`func (o *IntegrationGitHub) SetTokenHash(v string)`

SetTokenHash sets TokenHash field to given value.

### HasTokenHash

`func (o *IntegrationGitHub) HasTokenHash() bool`

HasTokenHash returns a boolean if a field has been set.

### GetServiceKey

`func (o *IntegrationGitHub) GetServiceKey() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *IntegrationGitHub) GetServiceKeyOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *IntegrationGitHub) SetServiceKey(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetServiceKey sets ServiceKey field to given value.

### HasServiceKey

`func (o *IntegrationGitHub) HasServiceKey() bool`

HasServiceKey returns a boolean if a field has been set.

### GetIsPlugin

`func (o *IntegrationGitHub) GetIsPlugin() bool`

GetIsPlugin returns the IsPlugin field if non-nil, zero value otherwise.

### GetIsPluginOk

`func (o *IntegrationGitHub) GetIsPluginOk() (*bool, bool)`

GetIsPluginOk returns a tuple with the IsPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlugin

`func (o *IntegrationGitHub) SetIsPlugin(v bool)`

SetIsPlugin sets IsPlugin field to given value.

### HasIsPlugin

`func (o *IntegrationGitHub) HasIsPlugin() bool`

HasIsPlugin returns a boolean if a field has been set.

### GetConfig

`func (o *IntegrationGitHub) GetConfig() ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf7Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *IntegrationGitHub) GetConfigOk() (*ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf7Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *IntegrationGitHub) SetConfig(v ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf7Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *IntegrationGitHub) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStatus

`func (o *IntegrationGitHub) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *IntegrationGitHub) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *IntegrationGitHub) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *IntegrationGitHub) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDate

`func (o *IntegrationGitHub) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *IntegrationGitHub) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *IntegrationGitHub) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *IntegrationGitHub) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetStatusMessage

`func (o *IntegrationGitHub) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *IntegrationGitHub) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *IntegrationGitHub) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *IntegrationGitHub) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetLastSync

`func (o *IntegrationGitHub) GetLastSync() string`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *IntegrationGitHub) GetLastSyncOk() (*string, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *IntegrationGitHub) SetLastSync(v string)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *IntegrationGitHub) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### GetLastSyncDuration

`func (o *IntegrationGitHub) GetLastSyncDuration() string`

GetLastSyncDuration returns the LastSyncDuration field if non-nil, zero value otherwise.

### GetLastSyncDurationOk

`func (o *IntegrationGitHub) GetLastSyncDurationOk() (*string, bool)`

GetLastSyncDurationOk returns a tuple with the LastSyncDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDuration

`func (o *IntegrationGitHub) SetLastSyncDuration(v string)`

SetLastSyncDuration sets LastSyncDuration field to given value.

### HasLastSyncDuration

`func (o *IntegrationGitHub) HasLastSyncDuration() bool`

HasLastSyncDuration returns a boolean if a field has been set.

### GetCredential

`func (o *IntegrationGitHub) GetCredential() ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *IntegrationGitHub) GetCredentialOk() (*ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *IntegrationGitHub) SetCredential(v ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *IntegrationGitHub) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


