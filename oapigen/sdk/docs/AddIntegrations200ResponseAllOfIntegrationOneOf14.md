# AddIntegrations200ResponseAllOfIntegrationOneOf14

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**IntegrationType** | Pointer to [**AddIntegrations200ResponseAllOfIntegrationOneOf14IntegrationType**](AddIntegrations200ResponseAllOfIntegrationOneOf14IntegrationType.md) |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**PasswordHash** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**WindowsVersion** | Pointer to **string** |  | [optional] 
**RepoUrl** | Pointer to **string** |  | [optional] 
**ServiceMode** | Pointer to **string** |  | [optional] 
**IsPlugin** | Pointer to **bool** |  | [optional] 
**Config** | Pointer to [**AddIntegrations200ResponseAllOfIntegrationOneOf14Config**](AddIntegrations200ResponseAllOfIntegrationOneOf14Config.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusDate** | Pointer to **time.Time** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**LastSync** | Pointer to **NullableString** |  | [optional] 
**LastSyncDuration** | Pointer to **NullableString** |  | [optional] 
**Credential** | Pointer to [**AddIntegrations200ResponseAllOfIntegrationOneOf14Credential**](AddIntegrations200ResponseAllOfIntegrationOneOf14Credential.md) |  | [optional] 

## Methods

### NewAddIntegrations200ResponseAllOfIntegrationOneOf14

`func NewAddIntegrations200ResponseAllOfIntegrationOneOf14() *AddIntegrations200ResponseAllOfIntegrationOneOf14`

NewAddIntegrations200ResponseAllOfIntegrationOneOf14 instantiates a new AddIntegrations200ResponseAllOfIntegrationOneOf14 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddIntegrations200ResponseAllOfIntegrationOneOf14WithDefaults

`func NewAddIntegrations200ResponseAllOfIntegrationOneOf14WithDefaults() *AddIntegrations200ResponseAllOfIntegrationOneOf14`

NewAddIntegrations200ResponseAllOfIntegrationOneOf14WithDefaults instantiates a new AddIntegrations200ResponseAllOfIntegrationOneOf14 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetType

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIntegrationType

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetIntegrationType() AddIntegrations200ResponseAllOfIntegrationOneOf14IntegrationType`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetIntegrationTypeOk() (*AddIntegrations200ResponseAllOfIntegrationOneOf14IntegrationType, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) SetIntegrationType(v AddIntegrations200ResponseAllOfIntegrationOneOf14IntegrationType)`

SetIntegrationType sets IntegrationType field to given value.

### HasIntegrationType

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) HasIntegrationType() bool`

HasIntegrationType returns a boolean if a field has been set.

### GetUrl

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetPort

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUsername

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordHash

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### GetPath

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetVersion

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetWindowsVersion

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetWindowsVersion() string`

GetWindowsVersion returns the WindowsVersion field if non-nil, zero value otherwise.

### GetWindowsVersionOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetWindowsVersionOk() (*string, bool)`

GetWindowsVersionOk returns a tuple with the WindowsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsVersion

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) SetWindowsVersion(v string)`

SetWindowsVersion sets WindowsVersion field to given value.

### HasWindowsVersion

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) HasWindowsVersion() bool`

HasWindowsVersion returns a boolean if a field has been set.

### GetRepoUrl

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetRepoUrl() string`

GetRepoUrl returns the RepoUrl field if non-nil, zero value otherwise.

### GetRepoUrlOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetRepoUrlOk() (*string, bool)`

GetRepoUrlOk returns a tuple with the RepoUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepoUrl

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) SetRepoUrl(v string)`

SetRepoUrl sets RepoUrl field to given value.

### HasRepoUrl

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) HasRepoUrl() bool`

HasRepoUrl returns a boolean if a field has been set.

### GetServiceMode

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetServiceMode() string`

GetServiceMode returns the ServiceMode field if non-nil, zero value otherwise.

### GetServiceModeOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetServiceModeOk() (*string, bool)`

GetServiceModeOk returns a tuple with the ServiceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMode

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) SetServiceMode(v string)`

SetServiceMode sets ServiceMode field to given value.

### HasServiceMode

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) HasServiceMode() bool`

HasServiceMode returns a boolean if a field has been set.

### GetIsPlugin

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetIsPlugin() bool`

GetIsPlugin returns the IsPlugin field if non-nil, zero value otherwise.

### GetIsPluginOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetIsPluginOk() (*bool, bool)`

GetIsPluginOk returns a tuple with the IsPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlugin

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) SetIsPlugin(v bool)`

SetIsPlugin sets IsPlugin field to given value.

### HasIsPlugin

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) HasIsPlugin() bool`

HasIsPlugin returns a boolean if a field has been set.

### GetConfig

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetConfig() AddIntegrations200ResponseAllOfIntegrationOneOf14Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetConfigOk() (*AddIntegrations200ResponseAllOfIntegrationOneOf14Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) SetConfig(v AddIntegrations200ResponseAllOfIntegrationOneOf14Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStatus

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDate

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetStatusMessage

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetLastSync

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetLastSync() string`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetLastSyncOk() (*string, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) SetLastSync(v string)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### SetLastSyncNil

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) SetLastSyncNil(b bool)`

 SetLastSyncNil sets the value for LastSync to be an explicit nil

### UnsetLastSync
`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) UnsetLastSync()`

UnsetLastSync ensures that no value is present for LastSync, not even an explicit nil
### GetLastSyncDuration

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetLastSyncDuration() string`

GetLastSyncDuration returns the LastSyncDuration field if non-nil, zero value otherwise.

### GetLastSyncDurationOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetLastSyncDurationOk() (*string, bool)`

GetLastSyncDurationOk returns a tuple with the LastSyncDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDuration

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) SetLastSyncDuration(v string)`

SetLastSyncDuration sets LastSyncDuration field to given value.

### HasLastSyncDuration

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) HasLastSyncDuration() bool`

HasLastSyncDuration returns a boolean if a field has been set.

### SetLastSyncDurationNil

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) SetLastSyncDurationNil(b bool)`

 SetLastSyncDurationNil sets the value for LastSyncDuration to be an explicit nil

### UnsetLastSyncDuration
`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) UnsetLastSyncDuration()`

UnsetLastSyncDuration ensures that no value is present for LastSyncDuration, not even an explicit nil
### GetCredential

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetCredential() AddIntegrations200ResponseAllOfIntegrationOneOf14Credential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) GetCredentialOk() (*AddIntegrations200ResponseAllOfIntegrationOneOf14Credential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) SetCredential(v AddIntegrations200ResponseAllOfIntegrationOneOf14Credential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf14) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


