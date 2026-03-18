# GetIntegrations200ResponseAllOfIntegrationOneOf4

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**IntegrationType** | Pointer to [**GetIntegrations200ResponseAllOfIntegrationOneOf4IntegrationType**](GetIntegrations200ResponseAllOfIntegrationOneOf4IntegrationType.md) |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**PasswordHash** | Pointer to **string** |  | [optional] 
**IsPlugin** | Pointer to **bool** |  | [optional] 
**Config** | Pointer to [**GetIntegrations200ResponseAllOfIntegrationOneOf4Config**](GetIntegrations200ResponseAllOfIntegrationOneOf4Config.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusDate** | Pointer to **NullableTime** |  | [optional] 
**StatusMessage** | Pointer to **NullableString** |  | [optional] 
**LastSync** | Pointer to **NullableString** |  | [optional] 
**LastSyncDuration** | Pointer to **NullableString** |  | [optional] 
**Credential** | Pointer to [**GetIntegrations200ResponseAllOfIntegrationOneOf4Credential**](GetIntegrations200ResponseAllOfIntegrationOneOf4Credential.md) |  | [optional] 

## Methods

### NewGetIntegrations200ResponseAllOfIntegrationOneOf4

`func NewGetIntegrations200ResponseAllOfIntegrationOneOf4() *GetIntegrations200ResponseAllOfIntegrationOneOf4`

NewGetIntegrations200ResponseAllOfIntegrationOneOf4 instantiates a new GetIntegrations200ResponseAllOfIntegrationOneOf4 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIntegrations200ResponseAllOfIntegrationOneOf4WithDefaults

`func NewGetIntegrations200ResponseAllOfIntegrationOneOf4WithDefaults() *GetIntegrations200ResponseAllOfIntegrationOneOf4`

NewGetIntegrations200ResponseAllOfIntegrationOneOf4WithDefaults instantiates a new GetIntegrations200ResponseAllOfIntegrationOneOf4 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetType

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIntegrationType

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) GetIntegrationType() GetIntegrations200ResponseAllOfIntegrationOneOf4IntegrationType`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) GetIntegrationTypeOk() (*GetIntegrations200ResponseAllOfIntegrationOneOf4IntegrationType, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) SetIntegrationType(v GetIntegrations200ResponseAllOfIntegrationOneOf4IntegrationType)`

SetIntegrationType sets IntegrationType field to given value.

### HasIntegrationType

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) HasIntegrationType() bool`

HasIntegrationType returns a boolean if a field has been set.

### GetUrl

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsername

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordHash

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### GetIsPlugin

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) GetIsPlugin() bool`

GetIsPlugin returns the IsPlugin field if non-nil, zero value otherwise.

### GetIsPluginOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) GetIsPluginOk() (*bool, bool)`

GetIsPluginOk returns a tuple with the IsPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlugin

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) SetIsPlugin(v bool)`

SetIsPlugin sets IsPlugin field to given value.

### HasIsPlugin

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) HasIsPlugin() bool`

HasIsPlugin returns a boolean if a field has been set.

### GetConfig

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) GetConfig() GetIntegrations200ResponseAllOfIntegrationOneOf4Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) GetConfigOk() (*GetIntegrations200ResponseAllOfIntegrationOneOf4Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) SetConfig(v GetIntegrations200ResponseAllOfIntegrationOneOf4Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStatus

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDate

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### SetStatusDateNil

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) SetStatusDateNil(b bool)`

 SetStatusDateNil sets the value for StatusDate to be an explicit nil

### UnsetStatusDate
`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) UnsetStatusDate()`

UnsetStatusDate ensures that no value is present for StatusDate, not even an explicit nil
### GetStatusMessage

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetLastSync

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) GetLastSync() string`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) GetLastSyncOk() (*string, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) SetLastSync(v string)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### SetLastSyncNil

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) SetLastSyncNil(b bool)`

 SetLastSyncNil sets the value for LastSync to be an explicit nil

### UnsetLastSync
`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) UnsetLastSync()`

UnsetLastSync ensures that no value is present for LastSync, not even an explicit nil
### GetLastSyncDuration

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) GetLastSyncDuration() string`

GetLastSyncDuration returns the LastSyncDuration field if non-nil, zero value otherwise.

### GetLastSyncDurationOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) GetLastSyncDurationOk() (*string, bool)`

GetLastSyncDurationOk returns a tuple with the LastSyncDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDuration

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) SetLastSyncDuration(v string)`

SetLastSyncDuration sets LastSyncDuration field to given value.

### HasLastSyncDuration

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) HasLastSyncDuration() bool`

HasLastSyncDuration returns a boolean if a field has been set.

### SetLastSyncDurationNil

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) SetLastSyncDurationNil(b bool)`

 SetLastSyncDurationNil sets the value for LastSyncDuration to be an explicit nil

### UnsetLastSyncDuration
`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) UnsetLastSyncDuration()`

UnsetLastSyncDuration ensures that no value is present for LastSyncDuration, not even an explicit nil
### GetCredential

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) GetCredential() GetIntegrations200ResponseAllOfIntegrationOneOf4Credential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) GetCredentialOk() (*GetIntegrations200ResponseAllOfIntegrationOneOf4Credential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) SetCredential(v GetIntegrations200ResponseAllOfIntegrationOneOf4Credential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf4) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


