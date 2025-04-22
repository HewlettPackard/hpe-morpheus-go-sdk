# AddIntegrations200ResponseAllOfIntegration

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
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusDate** | Pointer to **time.Time** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**LastSync** | Pointer to **string** |  | [optional] 
**LastSyncDuration** | Pointer to **string** |  | [optional] 
**Credential** | Pointer to [**ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential**](ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential.md) |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Host** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**PasswordHash** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**TokenHash** | Pointer to **string** |  | [optional] 
**ServiceFlag** | Pointer to **bool** |  | [optional] 
**Port** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**WindowsVersion** | Pointer to **string** |  | [optional] 
**RepoUrl** | Pointer to **string** |  | [optional] 
**ServiceMode** | Pointer to **string** |  | [optional] 
**AuthType** | Pointer to **string** |  | [optional] 
**AuthId** | Pointer to **string** |  | [optional] 

## Methods

### NewAddIntegrations200ResponseAllOfIntegration

`func NewAddIntegrations200ResponseAllOfIntegration() *AddIntegrations200ResponseAllOfIntegration`

NewAddIntegrations200ResponseAllOfIntegration instantiates a new AddIntegrations200ResponseAllOfIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddIntegrations200ResponseAllOfIntegrationWithDefaults

`func NewAddIntegrations200ResponseAllOfIntegrationWithDefaults() *AddIntegrations200ResponseAllOfIntegration`

NewAddIntegrations200ResponseAllOfIntegrationWithDefaults instantiates a new AddIntegrations200ResponseAllOfIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddIntegrations200ResponseAllOfIntegration) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddIntegrations200ResponseAllOfIntegration) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddIntegrations200ResponseAllOfIntegration) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddIntegrations200ResponseAllOfIntegration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddIntegrations200ResponseAllOfIntegration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddIntegrations200ResponseAllOfIntegration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddIntegrations200ResponseAllOfIntegration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddIntegrations200ResponseAllOfIntegration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *AddIntegrations200ResponseAllOfIntegration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddIntegrations200ResponseAllOfIntegration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddIntegrations200ResponseAllOfIntegration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddIntegrations200ResponseAllOfIntegration) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetType

`func (o *AddIntegrations200ResponseAllOfIntegration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddIntegrations200ResponseAllOfIntegration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddIntegrations200ResponseAllOfIntegration) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddIntegrations200ResponseAllOfIntegration) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIntegrationType

`func (o *AddIntegrations200ResponseAllOfIntegration) GetIntegrationType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *AddIntegrations200ResponseAllOfIntegration) GetIntegrationTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *AddIntegrations200ResponseAllOfIntegration) SetIntegrationType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetIntegrationType sets IntegrationType field to given value.

### HasIntegrationType

`func (o *AddIntegrations200ResponseAllOfIntegration) HasIntegrationType() bool`

HasIntegrationType returns a boolean if a field has been set.

### GetUrl

`func (o *AddIntegrations200ResponseAllOfIntegration) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AddIntegrations200ResponseAllOfIntegration) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AddIntegrations200ResponseAllOfIntegration) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AddIntegrations200ResponseAllOfIntegration) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetServiceKey

`func (o *AddIntegrations200ResponseAllOfIntegration) GetServiceKey() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *AddIntegrations200ResponseAllOfIntegration) GetServiceKeyOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *AddIntegrations200ResponseAllOfIntegration) SetServiceKey(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetServiceKey sets ServiceKey field to given value.

### HasServiceKey

`func (o *AddIntegrations200ResponseAllOfIntegration) HasServiceKey() bool`

HasServiceKey returns a boolean if a field has been set.

### GetIsPlugin

`func (o *AddIntegrations200ResponseAllOfIntegration) GetIsPlugin() bool`

GetIsPlugin returns the IsPlugin field if non-nil, zero value otherwise.

### GetIsPluginOk

`func (o *AddIntegrations200ResponseAllOfIntegration) GetIsPluginOk() (*bool, bool)`

GetIsPluginOk returns a tuple with the IsPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlugin

`func (o *AddIntegrations200ResponseAllOfIntegration) SetIsPlugin(v bool)`

SetIsPlugin sets IsPlugin field to given value.

### HasIsPlugin

`func (o *AddIntegrations200ResponseAllOfIntegration) HasIsPlugin() bool`

HasIsPlugin returns a boolean if a field has been set.

### GetConfig

`func (o *AddIntegrations200ResponseAllOfIntegration) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddIntegrations200ResponseAllOfIntegration) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddIntegrations200ResponseAllOfIntegration) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddIntegrations200ResponseAllOfIntegration) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStatus

`func (o *AddIntegrations200ResponseAllOfIntegration) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddIntegrations200ResponseAllOfIntegration) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddIntegrations200ResponseAllOfIntegration) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AddIntegrations200ResponseAllOfIntegration) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDate

`func (o *AddIntegrations200ResponseAllOfIntegration) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *AddIntegrations200ResponseAllOfIntegration) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *AddIntegrations200ResponseAllOfIntegration) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *AddIntegrations200ResponseAllOfIntegration) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetStatusMessage

`func (o *AddIntegrations200ResponseAllOfIntegration) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *AddIntegrations200ResponseAllOfIntegration) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *AddIntegrations200ResponseAllOfIntegration) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *AddIntegrations200ResponseAllOfIntegration) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetLastSync

`func (o *AddIntegrations200ResponseAllOfIntegration) GetLastSync() string`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *AddIntegrations200ResponseAllOfIntegration) GetLastSyncOk() (*string, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *AddIntegrations200ResponseAllOfIntegration) SetLastSync(v string)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *AddIntegrations200ResponseAllOfIntegration) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### GetLastSyncDuration

`func (o *AddIntegrations200ResponseAllOfIntegration) GetLastSyncDuration() string`

GetLastSyncDuration returns the LastSyncDuration field if non-nil, zero value otherwise.

### GetLastSyncDurationOk

`func (o *AddIntegrations200ResponseAllOfIntegration) GetLastSyncDurationOk() (*string, bool)`

GetLastSyncDurationOk returns a tuple with the LastSyncDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDuration

`func (o *AddIntegrations200ResponseAllOfIntegration) SetLastSyncDuration(v string)`

SetLastSyncDuration sets LastSyncDuration field to given value.

### HasLastSyncDuration

`func (o *AddIntegrations200ResponseAllOfIntegration) HasLastSyncDuration() bool`

HasLastSyncDuration returns a boolean if a field has been set.

### GetCredential

`func (o *AddIntegrations200ResponseAllOfIntegration) GetCredential() ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *AddIntegrations200ResponseAllOfIntegration) GetCredentialOk() (*ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *AddIntegrations200ResponseAllOfIntegration) SetCredential(v ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *AddIntegrations200ResponseAllOfIntegration) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetVersion

`func (o *AddIntegrations200ResponseAllOfIntegration) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AddIntegrations200ResponseAllOfIntegration) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AddIntegrations200ResponseAllOfIntegration) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AddIntegrations200ResponseAllOfIntegration) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetHost

`func (o *AddIntegrations200ResponseAllOfIntegration) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *AddIntegrations200ResponseAllOfIntegration) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *AddIntegrations200ResponseAllOfIntegration) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *AddIntegrations200ResponseAllOfIntegration) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetUsername

`func (o *AddIntegrations200ResponseAllOfIntegration) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AddIntegrations200ResponseAllOfIntegration) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AddIntegrations200ResponseAllOfIntegration) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AddIntegrations200ResponseAllOfIntegration) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *AddIntegrations200ResponseAllOfIntegration) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AddIntegrations200ResponseAllOfIntegration) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AddIntegrations200ResponseAllOfIntegration) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AddIntegrations200ResponseAllOfIntegration) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordHash

`func (o *AddIntegrations200ResponseAllOfIntegration) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *AddIntegrations200ResponseAllOfIntegration) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *AddIntegrations200ResponseAllOfIntegration) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *AddIntegrations200ResponseAllOfIntegration) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### GetToken

`func (o *AddIntegrations200ResponseAllOfIntegration) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AddIntegrations200ResponseAllOfIntegration) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AddIntegrations200ResponseAllOfIntegration) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AddIntegrations200ResponseAllOfIntegration) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTokenHash

`func (o *AddIntegrations200ResponseAllOfIntegration) GetTokenHash() string`

GetTokenHash returns the TokenHash field if non-nil, zero value otherwise.

### GetTokenHashOk

`func (o *AddIntegrations200ResponseAllOfIntegration) GetTokenHashOk() (*string, bool)`

GetTokenHashOk returns a tuple with the TokenHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenHash

`func (o *AddIntegrations200ResponseAllOfIntegration) SetTokenHash(v string)`

SetTokenHash sets TokenHash field to given value.

### HasTokenHash

`func (o *AddIntegrations200ResponseAllOfIntegration) HasTokenHash() bool`

HasTokenHash returns a boolean if a field has been set.

### GetServiceFlag

`func (o *AddIntegrations200ResponseAllOfIntegration) GetServiceFlag() bool`

GetServiceFlag returns the ServiceFlag field if non-nil, zero value otherwise.

### GetServiceFlagOk

`func (o *AddIntegrations200ResponseAllOfIntegration) GetServiceFlagOk() (*bool, bool)`

GetServiceFlagOk returns a tuple with the ServiceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceFlag

`func (o *AddIntegrations200ResponseAllOfIntegration) SetServiceFlag(v bool)`

SetServiceFlag sets ServiceFlag field to given value.

### HasServiceFlag

`func (o *AddIntegrations200ResponseAllOfIntegration) HasServiceFlag() bool`

HasServiceFlag returns a boolean if a field has been set.

### GetPort

`func (o *AddIntegrations200ResponseAllOfIntegration) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *AddIntegrations200ResponseAllOfIntegration) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *AddIntegrations200ResponseAllOfIntegration) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *AddIntegrations200ResponseAllOfIntegration) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPath

`func (o *AddIntegrations200ResponseAllOfIntegration) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AddIntegrations200ResponseAllOfIntegration) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AddIntegrations200ResponseAllOfIntegration) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *AddIntegrations200ResponseAllOfIntegration) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetWindowsVersion

`func (o *AddIntegrations200ResponseAllOfIntegration) GetWindowsVersion() string`

GetWindowsVersion returns the WindowsVersion field if non-nil, zero value otherwise.

### GetWindowsVersionOk

`func (o *AddIntegrations200ResponseAllOfIntegration) GetWindowsVersionOk() (*string, bool)`

GetWindowsVersionOk returns a tuple with the WindowsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsVersion

`func (o *AddIntegrations200ResponseAllOfIntegration) SetWindowsVersion(v string)`

SetWindowsVersion sets WindowsVersion field to given value.

### HasWindowsVersion

`func (o *AddIntegrations200ResponseAllOfIntegration) HasWindowsVersion() bool`

HasWindowsVersion returns a boolean if a field has been set.

### GetRepoUrl

`func (o *AddIntegrations200ResponseAllOfIntegration) GetRepoUrl() string`

GetRepoUrl returns the RepoUrl field if non-nil, zero value otherwise.

### GetRepoUrlOk

`func (o *AddIntegrations200ResponseAllOfIntegration) GetRepoUrlOk() (*string, bool)`

GetRepoUrlOk returns a tuple with the RepoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoUrl

`func (o *AddIntegrations200ResponseAllOfIntegration) SetRepoUrl(v string)`

SetRepoUrl sets RepoUrl field to given value.

### HasRepoUrl

`func (o *AddIntegrations200ResponseAllOfIntegration) HasRepoUrl() bool`

HasRepoUrl returns a boolean if a field has been set.

### GetServiceMode

`func (o *AddIntegrations200ResponseAllOfIntegration) GetServiceMode() string`

GetServiceMode returns the ServiceMode field if non-nil, zero value otherwise.

### GetServiceModeOk

`func (o *AddIntegrations200ResponseAllOfIntegration) GetServiceModeOk() (*string, bool)`

GetServiceModeOk returns a tuple with the ServiceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMode

`func (o *AddIntegrations200ResponseAllOfIntegration) SetServiceMode(v string)`

SetServiceMode sets ServiceMode field to given value.

### HasServiceMode

`func (o *AddIntegrations200ResponseAllOfIntegration) HasServiceMode() bool`

HasServiceMode returns a boolean if a field has been set.

### GetAuthType

`func (o *AddIntegrations200ResponseAllOfIntegration) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *AddIntegrations200ResponseAllOfIntegration) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *AddIntegrations200ResponseAllOfIntegration) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *AddIntegrations200ResponseAllOfIntegration) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetAuthId

`func (o *AddIntegrations200ResponseAllOfIntegration) GetAuthId() string`

GetAuthId returns the AuthId field if non-nil, zero value otherwise.

### GetAuthIdOk

`func (o *AddIntegrations200ResponseAllOfIntegration) GetAuthIdOk() (*string, bool)`

GetAuthIdOk returns a tuple with the AuthId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthId

`func (o *AddIntegrations200ResponseAllOfIntegration) SetAuthId(v string)`

SetAuthId sets AuthId field to given value.

### HasAuthId

`func (o *AddIntegrations200ResponseAllOfIntegration) HasAuthId() bool`

HasAuthId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


