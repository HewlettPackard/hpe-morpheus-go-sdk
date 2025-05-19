# ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**IntegrationType** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**ServiceKey** | Pointer to [**ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**IsPlugin** | Pointer to **bool** |  | [optional] 
**Config** | Pointer to [**ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfConfig**](ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfConfig.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusDate** | Pointer to **time.Time** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**LastSync** | Pointer to **string** |  | [optional] 
**LastSyncDuration** | Pointer to **int64** |  | [optional] 
**Credential** | Pointer to [**ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential**](ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential.md) |  | [optional] 

## Methods

### NewListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf

`func NewListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf() *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf`

NewListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf instantiates a new ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfWithDefaults

`func NewListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfWithDefaults() *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf`

NewListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfWithDefaults instantiates a new ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetType

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIntegrationType

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) GetIntegrationType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) GetIntegrationTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) SetIntegrationType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetIntegrationType sets IntegrationType field to given value.

### HasIntegrationType

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) HasIntegrationType() bool`

HasIntegrationType returns a boolean if a field has been set.

### GetUrl

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetServiceKey

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) GetServiceKey() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) GetServiceKeyOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) SetServiceKey(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetServiceKey sets ServiceKey field to given value.

### HasServiceKey

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) HasServiceKey() bool`

HasServiceKey returns a boolean if a field has been set.

### GetIsPlugin

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) GetIsPlugin() bool`

GetIsPlugin returns the IsPlugin field if non-nil, zero value otherwise.

### GetIsPluginOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) GetIsPluginOk() (*bool, bool)`

GetIsPluginOk returns a tuple with the IsPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlugin

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) SetIsPlugin(v bool)`

SetIsPlugin sets IsPlugin field to given value.

### HasIsPlugin

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) HasIsPlugin() bool`

HasIsPlugin returns a boolean if a field has been set.

### GetConfig

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) GetConfig() ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) GetConfigOk() (*ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) SetConfig(v ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStatus

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDate

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetLastSync

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) GetLastSync() string`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) GetLastSyncOk() (*string, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) SetLastSync(v string)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### GetLastSyncDuration

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) GetLastSyncDuration() int64`

GetLastSyncDuration returns the LastSyncDuration field if non-nil, zero value otherwise.

### GetLastSyncDurationOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) GetLastSyncDurationOk() (*int64, bool)`

GetLastSyncDurationOk returns a tuple with the LastSyncDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDuration

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) SetLastSyncDuration(v int64)`

SetLastSyncDuration sets LastSyncDuration field to given value.

### HasLastSyncDuration

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) HasLastSyncDuration() bool`

HasLastSyncDuration returns a boolean if a field has been set.

### GetCredential

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) GetCredential() ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) GetCredentialOk() (*ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) SetCredential(v ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


