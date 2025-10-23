# GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**IntegrationType** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**Url** | Pointer to **string** |  | [optional] 
**Port** | Pointer to **string** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 
**RefType** | Pointer to **string** |  | [optional] 
**RefId** | Pointer to **string** |  | [optional] 
**IsPlugin** | Pointer to **bool** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**StatusDate** | Pointer to **time.Time** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**LastSync** | Pointer to **time.Time** |  | [optional] 
**LastSyncDuration** | Pointer to **int64** |  | [optional] 

## Methods

### NewGetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration

`func NewGetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration() *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration`

NewGetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration instantiates a new GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkRouter200ResponseNetworkRouterNetworkServerIntegrationWithDefaults

`func NewGetNetworkRouter200ResponseNetworkRouterNetworkServerIntegrationWithDefaults() *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration`

NewGetNetworkRouter200ResponseNetworkRouterNetworkServerIntegrationWithDefaults instantiates a new GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetType

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIntegrationType

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetIntegrationType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetIntegrationType returns the IntegrationType field if non-nil, zero value otherwise.

### GetIntegrationTypeOk

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetIntegrationTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetIntegrationTypeOk returns a tuple with the IntegrationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationType

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) SetIntegrationType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetIntegrationType sets IntegrationType field to given value.

### HasIntegrationType

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) HasIntegrationType() bool`

HasIntegrationType returns a boolean if a field has been set.

### GetUrl

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetPort

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetPort() string`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetPortOk() (*string, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) SetPort(v string)`

SetPort sets Port field to given value.

### HasPort

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) HasPort() bool`

HasPort returns a boolean if a field has been set.

### GetUsername

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetPassword

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRefType

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetRefType() string`

GetRefType returns the RefType field if non-nil, zero value otherwise.

### GetRefTypeOk

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetRefTypeOk() (*string, bool)`

GetRefTypeOk returns a tuple with the RefType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefType

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) SetRefType(v string)`

SetRefType sets RefType field to given value.

### HasRefType

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) HasRefType() bool`

HasRefType returns a boolean if a field has been set.

### GetRefId

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetRefId() string`

GetRefId returns the RefId field if non-nil, zero value otherwise.

### GetRefIdOk

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetRefIdOk() (*string, bool)`

GetRefIdOk returns a tuple with the RefId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRefId

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) SetRefId(v string)`

SetRefId sets RefId field to given value.

### HasRefId

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) HasRefId() bool`

HasRefId returns a boolean if a field has been set.

### GetIsPlugin

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetIsPlugin() bool`

GetIsPlugin returns the IsPlugin field if non-nil, zero value otherwise.

### GetIsPluginOk

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetIsPluginOk() (*bool, bool)`

GetIsPluginOk returns a tuple with the IsPlugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsPlugin

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) SetIsPlugin(v bool)`

SetIsPlugin sets IsPlugin field to given value.

### HasIsPlugin

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) HasIsPlugin() bool`

HasIsPlugin returns a boolean if a field has been set.

### GetConfig

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStatus

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusDate

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetStatusMessage

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetLastSync

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetLastSync() time.Time`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetLastSyncOk() (*time.Time, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) SetLastSync(v time.Time)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### GetLastSyncDuration

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetLastSyncDuration() int64`

GetLastSyncDuration returns the LastSyncDuration field if non-nil, zero value otherwise.

### GetLastSyncDurationOk

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) GetLastSyncDurationOk() (*int64, bool)`

GetLastSyncDurationOk returns a tuple with the LastSyncDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDuration

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) SetLastSyncDuration(v int64)`

SetLastSyncDuration sets LastSyncDuration field to given value.

### HasLastSyncDuration

`func (o *GetNetworkRouter200ResponseNetworkRouterNetworkServerIntegration) HasLastSyncDuration() bool`

HasLastSyncDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


