# ListNetworkServers200ResponseAllOfNetworkServersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Network Server ID | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**Description** | Pointer to **string** | Description | [optional] 
**Type** | Pointer to [**ListNetworkServers200ResponseAllOfNetworkServersInnerType**](ListNetworkServers200ResponseAllOfNetworkServersInnerType.md) |  | [optional] 
**Integration** | Pointer to [**ListNetworkServers200ResponseAllOfNetworkServersInnerIntegration**](ListNetworkServers200ResponseAllOfNetworkServersInnerIntegration.md) |  | [optional] 
**Account** | Pointer to [**ListNetworkServers200ResponseAllOfNetworkServersInnerAccount**](ListNetworkServers200ResponseAllOfNetworkServersInnerAccount.md) |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**InternalId** | Pointer to **string** | Internal ID | [optional] 
**ExternalId** | Pointer to **string** | External ID | [optional] 
**ServiceUrl** | Pointer to **string** | Service URL | [optional] 
**ServiceHost** | Pointer to **string** | Service Host | [optional] 
**ServicePort** | Pointer to **int32** | Service Port | [optional] 
**ServiceMode** | Pointer to **string** | Service Mode | [optional] 
**ServicePath** | Pointer to **string** | Service Path | [optional] 
**ServiceUsername** | Pointer to **string** | Service Username | [optional] 
**ServicePassword** | Pointer to **string** | Service Password | [optional] 
**ServicePasswordHash** | Pointer to **string** |  | [optional] 
**ServiceToken** | Pointer to **string** | Service Token | [optional] 
**ServiceTokenHash** | Pointer to **string** |  | [optional] 
**ApiPort** | Pointer to **int32** |  | [optional] 
**AdminPort** | Pointer to **int32** |  | [optional] 
**Status** | Pointer to **string** | Status | [optional] 
**StatusMessage** | Pointer to **string** | Status Message | [optional] 
**StatusDate** | Pointer to **time.Time** |  | [optional] 
**LastSync** | Pointer to **time.Time** | Last Sync Date | [optional] 
**NextRunDate** | Pointer to **time.Time** | Next Run Date | [optional] 
**LastSyncDuration** | Pointer to **int64** | Last Sync Duration in milliseconds | [optional] 
**Config** | Pointer to **map[string]interface{}** | Config object varies with network server type. | [optional] 
**NetworkFilter** | Pointer to **string** | Network Filter | [optional] 
**TenantMatch** | Pointer to **string** | Tenant Match | [optional] 
**ZoneId** | Pointer to **int64** | Cloud ID | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Visible** | Pointer to **bool** |  | [optional] 
**Credential** | Pointer to [**ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential**](ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential.md) |  | [optional] 
**Tenants** | Pointer to [**[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 

## Methods

### NewListNetworkServers200ResponseAllOfNetworkServersInner

`func NewListNetworkServers200ResponseAllOfNetworkServersInner() *ListNetworkServers200ResponseAllOfNetworkServersInner`

NewListNetworkServers200ResponseAllOfNetworkServersInner instantiates a new ListNetworkServers200ResponseAllOfNetworkServersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNetworkServers200ResponseAllOfNetworkServersInnerWithDefaults

`func NewListNetworkServers200ResponseAllOfNetworkServersInnerWithDefaults() *ListNetworkServers200ResponseAllOfNetworkServersInner`

NewListNetworkServers200ResponseAllOfNetworkServersInnerWithDefaults instantiates a new ListNetworkServers200ResponseAllOfNetworkServersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetType() ListNetworkServers200ResponseAllOfNetworkServersInnerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetTypeOk() (*ListNetworkServers200ResponseAllOfNetworkServersInnerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) SetType(v ListNetworkServers200ResponseAllOfNetworkServersInnerType)`

SetType sets Type field to given value.

### HasType

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIntegration

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetIntegration() ListNetworkServers200ResponseAllOfNetworkServersInnerIntegration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetIntegrationOk() (*ListNetworkServers200ResponseAllOfNetworkServersInnerIntegration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) SetIntegration(v ListNetworkServers200ResponseAllOfNetworkServersInnerIntegration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetAccount

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetAccount() ListNetworkServers200ResponseAllOfNetworkServersInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetAccountOk() (*ListNetworkServers200ResponseAllOfNetworkServersInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) SetAccount(v ListNetworkServers200ResponseAllOfNetworkServersInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetVisibility

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetInternalId

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalId

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetServiceUrl

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.

### HasServiceUrl

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) HasServiceUrl() bool`

HasServiceUrl returns a boolean if a field has been set.

### GetServiceHost

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetServiceHost() string`

GetServiceHost returns the ServiceHost field if non-nil, zero value otherwise.

### GetServiceHostOk

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetServiceHostOk() (*string, bool)`

GetServiceHostOk returns a tuple with the ServiceHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHost

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) SetServiceHost(v string)`

SetServiceHost sets ServiceHost field to given value.

### HasServiceHost

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) HasServiceHost() bool`

HasServiceHost returns a boolean if a field has been set.

### GetServicePort

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetServicePort() int32`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetServicePortOk() (*int32, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) SetServicePort(v int32)`

SetServicePort sets ServicePort field to given value.

### HasServicePort

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.

### GetServiceMode

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetServiceMode() string`

GetServiceMode returns the ServiceMode field if non-nil, zero value otherwise.

### GetServiceModeOk

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetServiceModeOk() (*string, bool)`

GetServiceModeOk returns a tuple with the ServiceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMode

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) SetServiceMode(v string)`

SetServiceMode sets ServiceMode field to given value.

### HasServiceMode

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) HasServiceMode() bool`

HasServiceMode returns a boolean if a field has been set.

### GetServicePath

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetServicePath() string`

GetServicePath returns the ServicePath field if non-nil, zero value otherwise.

### GetServicePathOk

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetServicePathOk() (*string, bool)`

GetServicePathOk returns a tuple with the ServicePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePath

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) SetServicePath(v string)`

SetServicePath sets ServicePath field to given value.

### HasServicePath

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) HasServicePath() bool`

HasServicePath returns a boolean if a field has been set.

### GetServiceUsername

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.

### HasServiceUsername

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) HasServiceUsername() bool`

HasServiceUsername returns a boolean if a field has been set.

### GetServicePassword

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### GetServicePasswordHash

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetServicePasswordHash() string`

GetServicePasswordHash returns the ServicePasswordHash field if non-nil, zero value otherwise.

### GetServicePasswordHashOk

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetServicePasswordHashOk() (*string, bool)`

GetServicePasswordHashOk returns a tuple with the ServicePasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePasswordHash

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) SetServicePasswordHash(v string)`

SetServicePasswordHash sets ServicePasswordHash field to given value.

### HasServicePasswordHash

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) HasServicePasswordHash() bool`

HasServicePasswordHash returns a boolean if a field has been set.

### GetServiceToken

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetServiceToken() string`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetServiceTokenOk() (*string, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) SetServiceToken(v string)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.

### GetServiceTokenHash

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetServiceTokenHash() string`

GetServiceTokenHash returns the ServiceTokenHash field if non-nil, zero value otherwise.

### GetServiceTokenHashOk

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetServiceTokenHashOk() (*string, bool)`

GetServiceTokenHashOk returns a tuple with the ServiceTokenHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTokenHash

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) SetServiceTokenHash(v string)`

SetServiceTokenHash sets ServiceTokenHash field to given value.

### HasServiceTokenHash

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) HasServiceTokenHash() bool`

HasServiceTokenHash returns a boolean if a field has been set.

### GetApiPort

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetApiPort() int32`

GetApiPort returns the ApiPort field if non-nil, zero value otherwise.

### GetApiPortOk

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetApiPortOk() (*int32, bool)`

GetApiPortOk returns a tuple with the ApiPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPort

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) SetApiPort(v int32)`

SetApiPort sets ApiPort field to given value.

### HasApiPort

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) HasApiPort() bool`

HasApiPort returns a boolean if a field has been set.

### GetAdminPort

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetAdminPort() int32`

GetAdminPort returns the AdminPort field if non-nil, zero value otherwise.

### GetAdminPortOk

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetAdminPortOk() (*int32, bool)`

GetAdminPortOk returns a tuple with the AdminPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPort

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) SetAdminPort(v int32)`

SetAdminPort sets AdminPort field to given value.

### HasAdminPort

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) HasAdminPort() bool`

HasAdminPort returns a boolean if a field has been set.

### GetStatus

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetStatusDate

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetLastSync

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetLastSync() time.Time`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetLastSyncOk() (*time.Time, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) SetLastSync(v time.Time)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### GetNextRunDate

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetNextRunDate() time.Time`

GetNextRunDate returns the NextRunDate field if non-nil, zero value otherwise.

### GetNextRunDateOk

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetNextRunDateOk() (*time.Time, bool)`

GetNextRunDateOk returns a tuple with the NextRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunDate

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) SetNextRunDate(v time.Time)`

SetNextRunDate sets NextRunDate field to given value.

### HasNextRunDate

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) HasNextRunDate() bool`

HasNextRunDate returns a boolean if a field has been set.

### GetLastSyncDuration

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetLastSyncDuration() int64`

GetLastSyncDuration returns the LastSyncDuration field if non-nil, zero value otherwise.

### GetLastSyncDurationOk

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetLastSyncDurationOk() (*int64, bool)`

GetLastSyncDurationOk returns a tuple with the LastSyncDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDuration

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) SetLastSyncDuration(v int64)`

SetLastSyncDuration sets LastSyncDuration field to given value.

### HasLastSyncDuration

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) HasLastSyncDuration() bool`

HasLastSyncDuration returns a boolean if a field has been set.

### GetConfig

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetNetworkFilter

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetNetworkFilter() string`

GetNetworkFilter returns the NetworkFilter field if non-nil, zero value otherwise.

### GetNetworkFilterOk

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetNetworkFilterOk() (*string, bool)`

GetNetworkFilterOk returns a tuple with the NetworkFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFilter

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) SetNetworkFilter(v string)`

SetNetworkFilter sets NetworkFilter field to given value.

### HasNetworkFilter

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) HasNetworkFilter() bool`

HasNetworkFilter returns a boolean if a field has been set.

### GetTenantMatch

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetTenantMatch() string`

GetTenantMatch returns the TenantMatch field if non-nil, zero value otherwise.

### GetTenantMatchOk

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetTenantMatchOk() (*string, bool)`

GetTenantMatchOk returns a tuple with the TenantMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantMatch

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) SetTenantMatch(v string)`

SetTenantMatch sets TenantMatch field to given value.

### HasTenantMatch

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) HasTenantMatch() bool`

HasTenantMatch returns a boolean if a field has been set.

### GetZoneId

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetEnabled

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetVisible

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetCredential

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetCredential() ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetCredentialOk() (*ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) SetCredential(v ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetTenants

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetTenants() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) GetTenantsOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) SetTenants(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ListNetworkServers200ResponseAllOfNetworkServersInner) HasTenants() bool`

HasTenants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


