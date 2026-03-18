# AddIntegrations200ResponseAllOfIntegrationOneOf6

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**IntegrationType** | Pointer to [**AddIntegrations200ResponseAllOfIntegrationOneOf6IntegrationType**](AddIntegrations200ResponseAllOfIntegrationOneOf6IntegrationType.md) |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**PasswordHash** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**TokenHash** | Pointer to **string** |  | [optional] 
**IsPlugin** | Pointer to **bool** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | Pointer to **NullableString** |  | [optional] 
**StatusDate** | Pointer to **NullableTime** |  | [optional] 
**StatusMessage** | Pointer to **NullableString** |  | [optional] 
**LastSync** | Pointer to **NullableString** |  | [optional] 
**LastSyncDuration** | Pointer to **NullableString** |  | [optional] 
**Credential** | Pointer to [**AddIntegrations200ResponseAllOfIntegrationOneOf6Credential**](AddIntegrations200ResponseAllOfIntegrationOneOf6Credential.md) |  | [optional] 

## Methods

### NewAddIntegrations200ResponseAllOfIntegrationOneOf6

`func NewAddIntegrations200ResponseAllOfIntegrationOneOf6() *AddIntegrations200ResponseAllOfIntegrationOneOf6`

NewAddIntegrations200ResponseAllOfIntegrationOneOf6 instantiates a new AddIntegrations200ResponseAllOfIntegrationOneOf6 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddIntegrations200ResponseAllOfIntegrationOneOf6WithDefaults

`func NewAddIntegrations200ResponseAllOfIntegrationOneOf6WithDefaults() *AddIntegrations200ResponseAllOfIntegrationOneOf6`

NewAddIntegrations200ResponseAllOfIntegrationOneOf6WithDefaults instantiates a new AddIntegrations200ResponseAllOfIntegrationOneOf6 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetType

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIntegrationType

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) GetIntegrationType() AddIntegrations200ResponseAllOfIntegrationOneOf6IntegrationType`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) GetIntegrationTypeOk() (*AddIntegrations200ResponseAllOfIntegrationOneOf6IntegrationType, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) SetIntegrationType(v AddIntegrations200ResponseAllOfIntegrationOneOf6IntegrationType)`

SetIntegrationType sets IntegrationType field to given value.

### HasIntegrationType

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) HasIntegrationType() bool`

HasIntegrationType returns a boolean if a field has been set.

### GetUrl

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetUsername

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPasswordHash

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) GetPasswordHash() string`

GetPasswordHash returns the PasswordHash field if non-nil, zero value otherwise.

### GetPasswordHashOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) GetPasswordHashOk() (*string, bool)`

GetPasswordHashOk returns a tuple with the PasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPasswordHash

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) SetPasswordHash(v string)`

SetPasswordHash sets PasswordHash field to given value.

### HasPasswordHash

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) HasPasswordHash() bool`

HasPasswordHash returns a boolean if a field has been set.

### GetToken

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTokenHash

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) GetTokenHash() string`

GetTokenHash returns the TokenHash field if non-nil, zero value otherwise.

### GetTokenHashOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) GetTokenHashOk() (*string, bool)`

GetTokenHashOk returns a tuple with the TokenHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenHash

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) SetTokenHash(v string)`

SetTokenHash sets TokenHash field to given value.

### HasTokenHash

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) HasTokenHash() bool`

HasTokenHash returns a boolean if a field has been set.

### GetIsPlugin

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) GetIsPlugin() bool`

GetIsPlugin returns the IsPlugin field if non-nil, zero value otherwise.

### GetIsPluginOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) GetIsPluginOk() (*bool, bool)`

GetIsPluginOk returns a tuple with the IsPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlugin

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) SetIsPlugin(v bool)`

SetIsPlugin sets IsPlugin field to given value.

### HasIsPlugin

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) HasIsPlugin() bool`

HasIsPlugin returns a boolean if a field has been set.

### GetConfig

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetStatus

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### SetStatusNil

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) SetStatusNil(b bool)`

 SetStatusNil sets the value for Status to be an explicit nil

### UnsetStatus
`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) UnsetStatus()`

UnsetStatus ensures that no value is present for Status, not even an explicit nil
### GetStatusDate

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### SetStatusDateNil

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) SetStatusDateNil(b bool)`

 SetStatusDateNil sets the value for StatusDate to be an explicit nil

### UnsetStatusDate
`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) UnsetStatusDate()`

UnsetStatusDate ensures that no value is present for StatusDate, not even an explicit nil
### GetStatusMessage

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetLastSync

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) GetLastSync() string`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) GetLastSyncOk() (*string, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) SetLastSync(v string)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### SetLastSyncNil

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) SetLastSyncNil(b bool)`

 SetLastSyncNil sets the value for LastSync to be an explicit nil

### UnsetLastSync
`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) UnsetLastSync()`

UnsetLastSync ensures that no value is present for LastSync, not even an explicit nil
### GetLastSyncDuration

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) GetLastSyncDuration() string`

GetLastSyncDuration returns the LastSyncDuration field if non-nil, zero value otherwise.

### GetLastSyncDurationOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) GetLastSyncDurationOk() (*string, bool)`

GetLastSyncDurationOk returns a tuple with the LastSyncDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDuration

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) SetLastSyncDuration(v string)`

SetLastSyncDuration sets LastSyncDuration field to given value.

### HasLastSyncDuration

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) HasLastSyncDuration() bool`

HasLastSyncDuration returns a boolean if a field has been set.

### SetLastSyncDurationNil

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) SetLastSyncDurationNil(b bool)`

 SetLastSyncDurationNil sets the value for LastSyncDuration to be an explicit nil

### UnsetLastSyncDuration
`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) UnsetLastSyncDuration()`

UnsetLastSyncDuration ensures that no value is present for LastSyncDuration, not even an explicit nil
### GetCredential

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) GetCredential() AddIntegrations200ResponseAllOfIntegrationOneOf6Credential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) GetCredentialOk() (*AddIntegrations200ResponseAllOfIntegrationOneOf6Credential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) SetCredential(v AddIntegrations200ResponseAllOfIntegrationOneOf6Credential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *AddIntegrations200ResponseAllOfIntegrationOneOf6) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


