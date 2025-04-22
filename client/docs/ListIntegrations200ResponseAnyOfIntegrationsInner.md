# ListIntegrations200ResponseAnyOfIntegrationsInner

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

### NewListIntegrations200ResponseAnyOfIntegrationsInner

`func NewListIntegrations200ResponseAnyOfIntegrationsInner() *ListIntegrations200ResponseAnyOfIntegrationsInner`

NewListIntegrations200ResponseAnyOfIntegrationsInner instantiates a new ListIntegrations200ResponseAnyOfIntegrationsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListIntegrations200ResponseAnyOfIntegrationsInnerWithDefaults

`func NewListIntegrations200ResponseAnyOfIntegrationsInnerWithDefaults() *ListIntegrations200ResponseAnyOfIntegrationsInner`

NewListIntegrations200ResponseAnyOfIntegrationsInnerWithDefaults instantiates a new ListIntegrations200ResponseAnyOfIntegrationsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetType

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIntegrationType

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetIntegrationType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetIntegrationTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) SetIntegrationType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetIntegrationType sets IntegrationType field to given value.

### HasIntegrationType

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) HasIntegrationType() bool`

HasIntegrationType returns a boolean if a field has been set.

### GetUrl

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetServiceKey

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetServiceKey() ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetServiceKeyOk() (*ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) SetServiceKey(v ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetServiceKey sets ServiceKey field to given value.

### HasServiceKey

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) HasServiceKey() bool`

HasServiceKey returns a boolean if a field has been set.

### GetIsPlugin

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetIsPlugin() bool`

GetIsPlugin returns the IsPlugin field if non-nil, zero value otherwise.

### GetIsPluginOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetIsPluginOk() (*bool, bool)`

GetIsPluginOk returns a tuple with the IsPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlugin

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) SetIsPlugin(v bool)`

SetIsPlugin sets IsPlugin field to given value.

### HasIsPlugin

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) HasIsPlugin() bool`

HasIsPlugin returns a boolean if a field has been set.

### GetConfig

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStatus

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDate

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetLastSync

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetLastSync() string`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetLastSyncOk() (*string, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) SetLastSync(v string)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### GetLastSyncDuration

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetLastSyncDuration() string`

GetLastSyncDuration returns the LastSyncDuration field if non-nil, zero value otherwise.

### GetLastSyncDurationOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetLastSyncDurationOk() (*string, bool)`

GetLastSyncDurationOk returns a tuple with the LastSyncDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDuration

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) SetLastSyncDuration(v string)`

SetLastSyncDuration sets LastSyncDuration field to given value.

### HasLastSyncDuration

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) HasLastSyncDuration() bool`

HasLastSyncDuration returns a boolean if a field has been set.

### GetCredential

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetCredential() ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetCredentialOk() (*ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) SetCredential(v ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetVersion

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetHost

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetHost() string`

GetHost returns the Host field if non-nil, zero value otherwise.

### GetHostOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetHostOk() (*string, bool)`

GetHostOk returns a tuple with the Host field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHost

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) SetHost(v string)`

SetHost sets Host field to given value.

### HasHost

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) HasHost() bool`

HasHost returns a boolean if a field has been set.

### GetUsername

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordHash

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### GetToken

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTokenHash

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetTokenHash() string`

GetTokenHash returns the TokenHash field if non-nil, zero value otherwise.

### GetTokenHashOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetTokenHashOk() (*string, bool)`

GetTokenHashOk returns a tuple with the TokenHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenHash

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) SetTokenHash(v string)`

SetTokenHash sets TokenHash field to given value.

### HasTokenHash

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) HasTokenHash() bool`

HasTokenHash returns a boolean if a field has been set.

### GetServiceFlag

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetServiceFlag() bool`

GetServiceFlag returns the ServiceFlag field if non-nil, zero value otherwise.

### GetServiceFlagOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetServiceFlagOk() (*bool, bool)`

GetServiceFlagOk returns a tuple with the ServiceFlag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceFlag

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) SetServiceFlag(v bool)`

SetServiceFlag sets ServiceFlag field to given value.

### HasServiceFlag

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) HasServiceFlag() bool`

HasServiceFlag returns a boolean if a field has been set.

### GetPort

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetPath

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetWindowsVersion

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetWindowsVersion() string`

GetWindowsVersion returns the WindowsVersion field if non-nil, zero value otherwise.

### GetWindowsVersionOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetWindowsVersionOk() (*string, bool)`

GetWindowsVersionOk returns a tuple with the WindowsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsVersion

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) SetWindowsVersion(v string)`

SetWindowsVersion sets WindowsVersion field to given value.

### HasWindowsVersion

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) HasWindowsVersion() bool`

HasWindowsVersion returns a boolean if a field has been set.

### GetRepoUrl

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetRepoUrl() string`

GetRepoUrl returns the RepoUrl field if non-nil, zero value otherwise.

### GetRepoUrlOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetRepoUrlOk() (*string, bool)`

GetRepoUrlOk returns a tuple with the RepoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoUrl

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) SetRepoUrl(v string)`

SetRepoUrl sets RepoUrl field to given value.

### HasRepoUrl

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) HasRepoUrl() bool`

HasRepoUrl returns a boolean if a field has been set.

### GetServiceMode

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetServiceMode() string`

GetServiceMode returns the ServiceMode field if non-nil, zero value otherwise.

### GetServiceModeOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetServiceModeOk() (*string, bool)`

GetServiceModeOk returns a tuple with the ServiceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMode

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) SetServiceMode(v string)`

SetServiceMode sets ServiceMode field to given value.

### HasServiceMode

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) HasServiceMode() bool`

HasServiceMode returns a boolean if a field has been set.

### GetAuthType

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetAuthType() string`

GetAuthType returns the AuthType field if non-nil, zero value otherwise.

### GetAuthTypeOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetAuthTypeOk() (*string, bool)`

GetAuthTypeOk returns a tuple with the AuthType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthType

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) SetAuthType(v string)`

SetAuthType sets AuthType field to given value.

### HasAuthType

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) HasAuthType() bool`

HasAuthType returns a boolean if a field has been set.

### GetAuthId

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetAuthId() string`

GetAuthId returns the AuthId field if non-nil, zero value otherwise.

### GetAuthIdOk

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) GetAuthIdOk() (*string, bool)`

GetAuthIdOk returns a tuple with the AuthId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthId

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) SetAuthId(v string)`

SetAuthId sets AuthId field to given value.

### HasAuthId

`func (o *ListIntegrations200ResponseAnyOfIntegrationsInner) HasAuthId() bool`

HasAuthId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


