# GetIntegrations200ResponseAllOfIntegrationOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**IntegrationType** | Pointer to [**GetIntegrations200ResponseAllOfIntegrationOneOfIntegrationType**](GetIntegrations200ResponseAllOfIntegrationOneOfIntegrationType.md) |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**ServiceKey** | Pointer to [**GetIntegrations200ResponseAllOfIntegrationOneOfServiceKey**](GetIntegrations200ResponseAllOfIntegrationOneOfServiceKey.md) |  | [optional] 
**IsPlugin** | Pointer to **bool** |  | [optional] 
**Config** | Pointer to [**GetIntegrations200ResponseAllOfIntegrationOneOfConfig**](GetIntegrations200ResponseAllOfIntegrationOneOfConfig.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusDate** | Pointer to **time.Time** |  | [optional] 
**StatusMessage** | Pointer to **NullableString** |  | [optional] 
**LastSync** | Pointer to **NullableString** |  | [optional] 
**LastSyncDuration** | Pointer to **NullableInt64** |  | [optional] 
**Credential** | Pointer to [**GetIntegrations200ResponseAllOfIntegrationOneOfCredential**](GetIntegrations200ResponseAllOfIntegrationOneOfCredential.md) |  | [optional] 

## Methods

### NewGetIntegrations200ResponseAllOfIntegrationOneOf

`func NewGetIntegrations200ResponseAllOfIntegrationOneOf() *GetIntegrations200ResponseAllOfIntegrationOneOf`

NewGetIntegrations200ResponseAllOfIntegrationOneOf instantiates a new GetIntegrations200ResponseAllOfIntegrationOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetIntegrations200ResponseAllOfIntegrationOneOfWithDefaults

`func NewGetIntegrations200ResponseAllOfIntegrationOneOfWithDefaults() *GetIntegrations200ResponseAllOfIntegrationOneOf`

NewGetIntegrations200ResponseAllOfIntegrationOneOfWithDefaults instantiates a new GetIntegrations200ResponseAllOfIntegrationOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetType

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIntegrationType

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) GetIntegrationType() GetIntegrations200ResponseAllOfIntegrationOneOfIntegrationType`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) GetIntegrationTypeOk() (*GetIntegrations200ResponseAllOfIntegrationOneOfIntegrationType, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) SetIntegrationType(v GetIntegrations200ResponseAllOfIntegrationOneOfIntegrationType)`

SetIntegrationType sets IntegrationType field to given value.

### HasIntegrationType

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) HasIntegrationType() bool`

HasIntegrationType returns a boolean if a field has been set.

### GetUrl

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetServiceKey

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) GetServiceKey() GetIntegrations200ResponseAllOfIntegrationOneOfServiceKey`

GetServiceKey returns the ServiceKey field if non-nil, zero value otherwise.

### GetServiceKeyOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) GetServiceKeyOk() (*GetIntegrations200ResponseAllOfIntegrationOneOfServiceKey, bool)`

GetServiceKeyOk returns a tuple with the ServiceKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceKey

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) SetServiceKey(v GetIntegrations200ResponseAllOfIntegrationOneOfServiceKey)`

SetServiceKey sets ServiceKey field to given value.

### HasServiceKey

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) HasServiceKey() bool`

HasServiceKey returns a boolean if a field has been set.

### GetIsPlugin

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) GetIsPlugin() bool`

GetIsPlugin returns the IsPlugin field if non-nil, zero value otherwise.

### GetIsPluginOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) GetIsPluginOk() (*bool, bool)`

GetIsPluginOk returns a tuple with the IsPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlugin

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) SetIsPlugin(v bool)`

SetIsPlugin sets IsPlugin field to given value.

### HasIsPlugin

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) HasIsPlugin() bool`

HasIsPlugin returns a boolean if a field has been set.

### GetConfig

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) GetConfig() GetIntegrations200ResponseAllOfIntegrationOneOfConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) GetConfigOk() (*GetIntegrations200ResponseAllOfIntegrationOneOfConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) SetConfig(v GetIntegrations200ResponseAllOfIntegrationOneOfConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStatus

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDate

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetStatusMessage

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetLastSync

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) GetLastSync() string`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) GetLastSyncOk() (*string, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) SetLastSync(v string)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### SetLastSyncNil

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) SetLastSyncNil(b bool)`

 SetLastSyncNil sets the value for LastSync to be an explicit nil

### UnsetLastSync
`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) UnsetLastSync()`

UnsetLastSync ensures that no value is present for LastSync, not even an explicit nil
### GetLastSyncDuration

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) GetLastSyncDuration() int64`

GetLastSyncDuration returns the LastSyncDuration field if non-nil, zero value otherwise.

### GetLastSyncDurationOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) GetLastSyncDurationOk() (*int64, bool)`

GetLastSyncDurationOk returns a tuple with the LastSyncDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDuration

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) SetLastSyncDuration(v int64)`

SetLastSyncDuration sets LastSyncDuration field to given value.

### HasLastSyncDuration

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) HasLastSyncDuration() bool`

HasLastSyncDuration returns a boolean if a field has been set.

### SetLastSyncDurationNil

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) SetLastSyncDurationNil(b bool)`

 SetLastSyncDurationNil sets the value for LastSyncDuration to be an explicit nil

### UnsetLastSyncDuration
`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) UnsetLastSyncDuration()`

UnsetLastSyncDuration ensures that no value is present for LastSyncDuration, not even an explicit nil
### GetCredential

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) GetCredential() GetIntegrations200ResponseAllOfIntegrationOneOfCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) GetCredentialOk() (*GetIntegrations200ResponseAllOfIntegrationOneOfCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) SetCredential(v GetIntegrations200ResponseAllOfIntegrationOneOfCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *GetIntegrations200ResponseAllOfIntegrationOneOf) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


