# NetworkServer

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

### NewNetworkServer

`func NewNetworkServer() *NetworkServer`

NewNetworkServer instantiates a new NetworkServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkServerWithDefaults

`func NewNetworkServerWithDefaults() *NetworkServer`

NewNetworkServerWithDefaults instantiates a new NetworkServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkServer) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkServer) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkServer) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *NetworkServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *NetworkServer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *NetworkServer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *NetworkServer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *NetworkServer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *NetworkServer) GetType() ListNetworkServers200ResponseAllOfNetworkServersInnerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkServer) GetTypeOk() (*ListNetworkServers200ResponseAllOfNetworkServersInnerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkServer) SetType(v ListNetworkServers200ResponseAllOfNetworkServersInnerType)`

SetType sets Type field to given value.

### HasType

`func (o *NetworkServer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIntegration

`func (o *NetworkServer) GetIntegration() ListNetworkServers200ResponseAllOfNetworkServersInnerIntegration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *NetworkServer) GetIntegrationOk() (*ListNetworkServers200ResponseAllOfNetworkServersInnerIntegration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *NetworkServer) SetIntegration(v ListNetworkServers200ResponseAllOfNetworkServersInnerIntegration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *NetworkServer) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetAccount

`func (o *NetworkServer) GetAccount() ListNetworkServers200ResponseAllOfNetworkServersInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *NetworkServer) GetAccountOk() (*ListNetworkServers200ResponseAllOfNetworkServersInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *NetworkServer) SetAccount(v ListNetworkServers200ResponseAllOfNetworkServersInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *NetworkServer) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetVisibility

`func (o *NetworkServer) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *NetworkServer) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *NetworkServer) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *NetworkServer) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetInternalId

`func (o *NetworkServer) GetInternalId() string`

GetInternalId returns the InternalId field if non-nil, zero value otherwise.

### GetInternalIdOk

`func (o *NetworkServer) GetInternalIdOk() (*string, bool)`

GetInternalIdOk returns a tuple with the InternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalId

`func (o *NetworkServer) SetInternalId(v string)`

SetInternalId sets InternalId field to given value.

### HasInternalId

`func (o *NetworkServer) HasInternalId() bool`

HasInternalId returns a boolean if a field has been set.

### GetExternalId

`func (o *NetworkServer) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *NetworkServer) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *NetworkServer) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *NetworkServer) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetServiceUrl

`func (o *NetworkServer) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *NetworkServer) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *NetworkServer) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.

### HasServiceUrl

`func (o *NetworkServer) HasServiceUrl() bool`

HasServiceUrl returns a boolean if a field has been set.

### GetServiceHost

`func (o *NetworkServer) GetServiceHost() string`

GetServiceHost returns the ServiceHost field if non-nil, zero value otherwise.

### GetServiceHostOk

`func (o *NetworkServer) GetServiceHostOk() (*string, bool)`

GetServiceHostOk returns a tuple with the ServiceHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHost

`func (o *NetworkServer) SetServiceHost(v string)`

SetServiceHost sets ServiceHost field to given value.

### HasServiceHost

`func (o *NetworkServer) HasServiceHost() bool`

HasServiceHost returns a boolean if a field has been set.

### GetServicePort

`func (o *NetworkServer) GetServicePort() int32`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *NetworkServer) GetServicePortOk() (*int32, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *NetworkServer) SetServicePort(v int32)`

SetServicePort sets ServicePort field to given value.

### HasServicePort

`func (o *NetworkServer) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.

### GetServiceMode

`func (o *NetworkServer) GetServiceMode() string`

GetServiceMode returns the ServiceMode field if non-nil, zero value otherwise.

### GetServiceModeOk

`func (o *NetworkServer) GetServiceModeOk() (*string, bool)`

GetServiceModeOk returns a tuple with the ServiceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMode

`func (o *NetworkServer) SetServiceMode(v string)`

SetServiceMode sets ServiceMode field to given value.

### HasServiceMode

`func (o *NetworkServer) HasServiceMode() bool`

HasServiceMode returns a boolean if a field has been set.

### GetServicePath

`func (o *NetworkServer) GetServicePath() string`

GetServicePath returns the ServicePath field if non-nil, zero value otherwise.

### GetServicePathOk

`func (o *NetworkServer) GetServicePathOk() (*string, bool)`

GetServicePathOk returns a tuple with the ServicePath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePath

`func (o *NetworkServer) SetServicePath(v string)`

SetServicePath sets ServicePath field to given value.

### HasServicePath

`func (o *NetworkServer) HasServicePath() bool`

HasServicePath returns a boolean if a field has been set.

### GetServiceUsername

`func (o *NetworkServer) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *NetworkServer) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *NetworkServer) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.

### HasServiceUsername

`func (o *NetworkServer) HasServiceUsername() bool`

HasServiceUsername returns a boolean if a field has been set.

### GetServicePassword

`func (o *NetworkServer) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *NetworkServer) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *NetworkServer) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *NetworkServer) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### GetServicePasswordHash

`func (o *NetworkServer) GetServicePasswordHash() string`

GetServicePasswordHash returns the ServicePasswordHash field if non-nil, zero value otherwise.

### GetServicePasswordHashOk

`func (o *NetworkServer) GetServicePasswordHashOk() (*string, bool)`

GetServicePasswordHashOk returns a tuple with the ServicePasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePasswordHash

`func (o *NetworkServer) SetServicePasswordHash(v string)`

SetServicePasswordHash sets ServicePasswordHash field to given value.

### HasServicePasswordHash

`func (o *NetworkServer) HasServicePasswordHash() bool`

HasServicePasswordHash returns a boolean if a field has been set.

### GetServiceToken

`func (o *NetworkServer) GetServiceToken() string`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *NetworkServer) GetServiceTokenOk() (*string, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *NetworkServer) SetServiceToken(v string)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *NetworkServer) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.

### GetServiceTokenHash

`func (o *NetworkServer) GetServiceTokenHash() string`

GetServiceTokenHash returns the ServiceTokenHash field if non-nil, zero value otherwise.

### GetServiceTokenHashOk

`func (o *NetworkServer) GetServiceTokenHashOk() (*string, bool)`

GetServiceTokenHashOk returns a tuple with the ServiceTokenHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceTokenHash

`func (o *NetworkServer) SetServiceTokenHash(v string)`

SetServiceTokenHash sets ServiceTokenHash field to given value.

### HasServiceTokenHash

`func (o *NetworkServer) HasServiceTokenHash() bool`

HasServiceTokenHash returns a boolean if a field has been set.

### GetApiPort

`func (o *NetworkServer) GetApiPort() int32`

GetApiPort returns the ApiPort field if non-nil, zero value otherwise.

### GetApiPortOk

`func (o *NetworkServer) GetApiPortOk() (*int32, bool)`

GetApiPortOk returns a tuple with the ApiPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiPort

`func (o *NetworkServer) SetApiPort(v int32)`

SetApiPort sets ApiPort field to given value.

### HasApiPort

`func (o *NetworkServer) HasApiPort() bool`

HasApiPort returns a boolean if a field has been set.

### GetAdminPort

`func (o *NetworkServer) GetAdminPort() int32`

GetAdminPort returns the AdminPort field if non-nil, zero value otherwise.

### GetAdminPortOk

`func (o *NetworkServer) GetAdminPortOk() (*int32, bool)`

GetAdminPortOk returns a tuple with the AdminPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPort

`func (o *NetworkServer) SetAdminPort(v int32)`

SetAdminPort sets AdminPort field to given value.

### HasAdminPort

`func (o *NetworkServer) HasAdminPort() bool`

HasAdminPort returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkServer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkServer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkServer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkServer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *NetworkServer) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *NetworkServer) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *NetworkServer) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *NetworkServer) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetStatusDate

`func (o *NetworkServer) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *NetworkServer) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *NetworkServer) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *NetworkServer) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetLastSync

`func (o *NetworkServer) GetLastSync() time.Time`

GetLastSync returns the LastSync field if non-nil, zero value otherwise.

### GetLastSyncOk

`func (o *NetworkServer) GetLastSyncOk() (*time.Time, bool)`

GetLastSyncOk returns a tuple with the LastSync field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSync

`func (o *NetworkServer) SetLastSync(v time.Time)`

SetLastSync sets LastSync field to given value.

### HasLastSync

`func (o *NetworkServer) HasLastSync() bool`

HasLastSync returns a boolean if a field has been set.

### GetNextRunDate

`func (o *NetworkServer) GetNextRunDate() time.Time`

GetNextRunDate returns the NextRunDate field if non-nil, zero value otherwise.

### GetNextRunDateOk

`func (o *NetworkServer) GetNextRunDateOk() (*time.Time, bool)`

GetNextRunDateOk returns a tuple with the NextRunDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextRunDate

`func (o *NetworkServer) SetNextRunDate(v time.Time)`

SetNextRunDate sets NextRunDate field to given value.

### HasNextRunDate

`func (o *NetworkServer) HasNextRunDate() bool`

HasNextRunDate returns a boolean if a field has been set.

### GetLastSyncDuration

`func (o *NetworkServer) GetLastSyncDuration() int64`

GetLastSyncDuration returns the LastSyncDuration field if non-nil, zero value otherwise.

### GetLastSyncDurationOk

`func (o *NetworkServer) GetLastSyncDurationOk() (*int64, bool)`

GetLastSyncDurationOk returns a tuple with the LastSyncDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSyncDuration

`func (o *NetworkServer) SetLastSyncDuration(v int64)`

SetLastSyncDuration sets LastSyncDuration field to given value.

### HasLastSyncDuration

`func (o *NetworkServer) HasLastSyncDuration() bool`

HasLastSyncDuration returns a boolean if a field has been set.

### GetConfig

`func (o *NetworkServer) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *NetworkServer) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *NetworkServer) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *NetworkServer) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetNetworkFilter

`func (o *NetworkServer) GetNetworkFilter() string`

GetNetworkFilter returns the NetworkFilter field if non-nil, zero value otherwise.

### GetNetworkFilterOk

`func (o *NetworkServer) GetNetworkFilterOk() (*string, bool)`

GetNetworkFilterOk returns a tuple with the NetworkFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFilter

`func (o *NetworkServer) SetNetworkFilter(v string)`

SetNetworkFilter sets NetworkFilter field to given value.

### HasNetworkFilter

`func (o *NetworkServer) HasNetworkFilter() bool`

HasNetworkFilter returns a boolean if a field has been set.

### GetTenantMatch

`func (o *NetworkServer) GetTenantMatch() string`

GetTenantMatch returns the TenantMatch field if non-nil, zero value otherwise.

### GetTenantMatchOk

`func (o *NetworkServer) GetTenantMatchOk() (*string, bool)`

GetTenantMatchOk returns a tuple with the TenantMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantMatch

`func (o *NetworkServer) SetTenantMatch(v string)`

SetTenantMatch sets TenantMatch field to given value.

### HasTenantMatch

`func (o *NetworkServer) HasTenantMatch() bool`

HasTenantMatch returns a boolean if a field has been set.

### GetZoneId

`func (o *NetworkServer) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *NetworkServer) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *NetworkServer) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *NetworkServer) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetDateCreated

`func (o *NetworkServer) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *NetworkServer) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *NetworkServer) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *NetworkServer) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *NetworkServer) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *NetworkServer) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *NetworkServer) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *NetworkServer) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetEnabled

`func (o *NetworkServer) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworkServer) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworkServer) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NetworkServer) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetVisible

`func (o *NetworkServer) GetVisible() bool`

GetVisible returns the Visible field if non-nil, zero value otherwise.

### GetVisibleOk

`func (o *NetworkServer) GetVisibleOk() (*bool, bool)`

GetVisibleOk returns a tuple with the Visible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisible

`func (o *NetworkServer) SetVisible(v bool)`

SetVisible sets Visible field to given value.

### HasVisible

`func (o *NetworkServer) HasVisible() bool`

HasVisible returns a boolean if a field has been set.

### GetCredential

`func (o *NetworkServer) GetCredential() ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *NetworkServer) GetCredentialOk() (*ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *NetworkServer) SetCredential(v ListIntegrations200ResponseAnyOfIntegrationsInnerAnyOfCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *NetworkServer) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetTenants

`func (o *NetworkServer) GetTenants() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *NetworkServer) GetTenantsOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *NetworkServer) SetTenants(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *NetworkServer) HasTenants() bool`

HasTenants returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


