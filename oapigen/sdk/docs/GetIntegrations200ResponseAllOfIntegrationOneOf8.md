# GetIntegrations200ResponseAllOfIntegrationOneOf8

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**IntegrationType** | Pointer to [**GetIntegrations200ResponseAllOfIntegrationOneOf8IntegrationType**](GetIntegrations200ResponseAllOfIntegrationOneOf8IntegrationType.md) |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Token** | Pointer to **string** |  | [optional] 
**TokenHash** | Pointer to **string** |  | [optional] 
**ServiceKey** | Pointer to [**GetIntegrations200ResponseAllOfIntegrationOneOf8ServiceKey**](GetIntegrations200ResponseAllOfIntegrationOneOf8ServiceKey.md) |  | [optional] 
**IsPlugin** | Pointer to **bool** |  | [optional] 
**Config** | Pointer to [**GetIntegrations200ResponseAllOfIntegrationOneOf8Config**](GetIntegrations200ResponseAllOfIntegrationOneOf8Config.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusDate** | Pointer to **time.Time** |  | [optional] 
**StatusMessage** | Pointer to **NullableString** |  | [optional] 
**LastSync** | Pointer to **NullableString** |  | [optional] 
**LastSyncDuration** | Pointer to **NullableString** |  | [optional] 
**Credential** | Pointer to [**GetIntegrations200ResponseAllOfIntegrationOneOf8Credential**](GetIntegrations200ResponseAllOfIntegrationOneOf8Credential.md) |  | [optional] 

## Methods

### NewGetIntegrations200ResponseAllOfIntegrationOneOf8

`func NewGetIntegrations200ResponseAllOfIntegrationOneOf8() *GetIntegrations200ResponseAllOfIntegrationOneOf8`

NewGetIntegrations200ResponseAllOfIntegrationOneOf8 instantiates a new GetIntegrations200ResponseAllOfIntegrationOneOf8 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetType

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIntegrationType

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) GetIntegrationType() GetIntegrations200ResponseAllOfIntegrationOneOf8IntegrationType`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) GetIntegrationTypeOk() (*GetIntegrations200ResponseAllOfIntegrationOneOf8IntegrationType, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) SetIntegrationType(v GetIntegrations200ResponseAllOfIntegrationOneOf8IntegrationType)`

SetIntegrationType sets IntegrationType field to given value.

### HasIntegrationType

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) HasIntegrationType() bool`

HasIntegrationType returns a boolean if a field has been set.

### GetUsername

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetToken

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) SetToken(v string)`

SetToken sets Token field to given value.

### HasToken

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) HasToken() bool`

HasToken returns a boolean if a field has been set.

### GetTokenHash

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) GetTokenHash() string`

GetTokenHash returns the TokenHash field if non-nil, zero value otherwise.

### GetTokenHashOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) GetTokenHashOk() (*string, bool)`

GetTokenHashOk returns a tuple with the TokenHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenHash

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) SetTokenHash(v string)`

SetTokenHash sets TokenHash field to given value.

### HasTokenHash

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) HasTokenHash() bool`

HasTokenHash returns a boolean if a field has been set.

### GetServiceKey

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) GetServiceKey() GetIntegrations200ResponseAllOfIntegrationOneOf8ServiceKey`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) GetServiceKeyOk() (*GetIntegrations200ResponseAllOfIntegrationOneOf8ServiceKey, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) SetServiceKey(v GetIntegrations200ResponseAllOfIntegrationOneOf8ServiceKey)`

SetServiceKey sets ServiceKey field to given value.

### HasServiceKey

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) HasServiceKey() bool`

HasServiceKey returns a boolean if a field has been set.

### GetIsPlugin

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) GetIsPlugin() bool`

GetIsPlugin returns the IsPlugin field if non-nil, zero value otherwise.

### GetIsPluginOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) GetIsPluginOk() (*bool, bool)`

GetIsPluginOk returns a tuple with the IsPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlugin

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) SetIsPlugin(v bool)`

SetIsPlugin sets IsPlugin field to given value.

### HasIsPlugin

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) HasIsPlugin() bool`

HasIsPlugin returns a boolean if a field has been set.

### GetConfig

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) GetConfig() GetIntegrations200ResponseAllOfIntegrationOneOf8Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) GetConfigOk() (*GetIntegrations200ResponseAllOfIntegrationOneOf8Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) SetConfig(v GetIntegrations200ResponseAllOfIntegrationOneOf8Config)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStatus

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDate

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetStatusMessage

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetLastSync

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) GetLastSync() string`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) GetLastSyncOk() (*string, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) SetLastSync(v string)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### SetLastSyncNil

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) SetLastSyncNil(b bool)`

 SetLastSyncNil sets the value for LastSync to be an explicit nil

### UnsetLastSync
`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) UnsetLastSync()`

UnsetLastSync ensures that no value is present for LastSync, not even an explicit nil
### GetLastSyncDuration

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) GetLastSyncDuration() string`

GetLastSyncDuration returns the LastSyncDuration field if non-nil, zero value otherwise.

### GetLastSyncDurationOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) GetLastSyncDurationOk() (*string, bool)`

GetLastSyncDurationOk returns a tuple with the LastSyncDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDuration

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) SetLastSyncDuration(v string)`

SetLastSyncDuration sets LastSyncDuration field to given value.

### HasLastSyncDuration

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) HasLastSyncDuration() bool`

HasLastSyncDuration returns a boolean if a field has been set.

### SetLastSyncDurationNil

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) SetLastSyncDurationNil(b bool)`

 SetLastSyncDurationNil sets the value for LastSyncDuration to be an explicit nil

### UnsetLastSyncDuration
`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) UnsetLastSyncDuration()`

UnsetLastSyncDuration ensures that no value is present for LastSyncDuration, not even an explicit nil
### GetCredential

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) GetCredential() GetIntegrations200ResponseAllOfIntegrationOneOf8Credential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) GetCredentialOk() (*GetIntegrations200ResponseAllOfIntegrationOneOf8Credential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) SetCredential(v GetIntegrations200ResponseAllOfIntegrationOneOf8Credential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf8) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


