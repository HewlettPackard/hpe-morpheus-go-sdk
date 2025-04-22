# ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**IntegrationType** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**IsPlugin** | Pointer to **bool** |  | [optional] 
**Config** | Pointer to [**ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config**](ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusDate** | Pointer to **time.Time** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**LastSync** | Pointer to **string** |  | [optional] 
**LastSyncDuration** | Pointer to **string** |  | [optional] 
**Credential** | Pointer to [**ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential**](ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential.md) |  | [optional] 

## Methods

### NewListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3

`func NewListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3() *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3`

NewListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3 instantiates a new ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3WithDefaults

`func NewListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3WithDefaults() *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3`

NewListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3WithDefaults instantiates a new ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetType

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIntegrationType

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) GetIntegrationType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) GetIntegrationTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) SetIntegrationType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetIntegrationType sets IntegrationType field to given value.

### HasIntegrationType

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) HasIntegrationType() bool`

HasIntegrationType returns a boolean if a field has been set.

### GetUrl

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetIsPlugin

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) GetIsPlugin() bool`

GetIsPlugin returns the IsPlugin field if non-nil, zero value otherwise.

### GetIsPluginOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) GetIsPluginOk() (*bool, bool)`

GetIsPluginOk returns a tuple with the IsPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlugin

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) SetIsPlugin(v bool)`

SetIsPlugin sets IsPlugin field to given value.

### HasIsPlugin

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) HasIsPlugin() bool`

HasIsPlugin returns a boolean if a field has been set.

### GetConfig

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) GetConfig() ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) GetConfigOk() (*ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) SetConfig(v ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStatus

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDate

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetLastSync

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) GetLastSync() string`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) GetLastSyncOk() (*string, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) SetLastSync(v string)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### GetLastSyncDuration

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) GetLastSyncDuration() string`

GetLastSyncDuration returns the LastSyncDuration field if non-nil, zero value otherwise.

### GetLastSyncDurationOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) GetLastSyncDurationOk() (*string, bool)`

GetLastSyncDurationOk returns a tuple with the LastSyncDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDuration

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) SetLastSyncDuration(v string)`

SetLastSyncDuration sets LastSyncDuration field to given value.

### HasLastSyncDuration

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) HasLastSyncDuration() bool`

HasLastSyncDuration returns a boolean if a field has been set.

### GetCredential

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) GetCredential() ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) GetCredentialOk() (*ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) SetCredential(v ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOf3) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


