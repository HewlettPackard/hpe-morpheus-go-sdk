# NetworkPoolServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Network Pool Server ID | [optional] 
**Type** | Pointer to [**ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInnerType**](ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInnerType.md) |  | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ServiceUrl** | Pointer to **string** | Service URL | [optional] 
**ServiceHost** | Pointer to **string** | Service Host | [optional] 
**ServicePort** | Pointer to **int32** | Service Port | [optional] 
**ServiceMode** | Pointer to **string** | Service Mode | [optional] 
**ServiceUsername** | Pointer to **string** | Service Username | [optional] 
**ServicePassword** | Pointer to **string** | Service Password | [optional] 
**ServicePasswordHash** | Pointer to **string** |  | [optional] 
**ServiceThrottleRate** | Pointer to **int64** | Throttle Rate | [optional] [default to 0]
**IgnoreSsl** | Pointer to **bool** | Disable SSL SNI Verification | [optional] [default to true]
**Status** | Pointer to **string** | Status | [optional] 
**StatusMessage** | Pointer to **string** | Status Message | [optional] 
**StatusDate** | Pointer to **time.Time** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** | Config object varies with pool server type. | [optional] 
**NetworkFilter** | Pointer to **string** | Network Filter | [optional] 
**ZoneFilter** | Pointer to **string** | Zone Filter | [optional] 
**TenantMatch** | Pointer to **string** | Tenant Match | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Account** | Pointer to [**ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInnerAccount**](ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInnerAccount.md) |  | [optional] 
**Integration** | Pointer to [**ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInnerIntegration**](ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInnerIntegration.md) |  | [optional] 
**Pools** | Pointer to [**[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner**](ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner.md) |  | [optional] 
**Credential** | Pointer to [**ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInnerCredential**](ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInnerCredential.md) |  | [optional] 

## Methods

### NewNetworkPoolServer

`func NewNetworkPoolServer() *NetworkPoolServer`

NewNetworkPoolServer instantiates a new NetworkPoolServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkPoolServerWithDefaults

`func NewNetworkPoolServerWithDefaults() *NetworkPoolServer`

NewNetworkPoolServerWithDefaults instantiates a new NetworkPoolServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *NetworkPoolServer) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *NetworkPoolServer) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *NetworkPoolServer) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *NetworkPoolServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *NetworkPoolServer) GetType() ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInnerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NetworkPoolServer) GetTypeOk() (*ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInnerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NetworkPoolServer) SetType(v ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInnerType)`

SetType sets Type field to given value.

### HasType

`func (o *NetworkPoolServer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *NetworkPoolServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkPoolServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkPoolServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkPoolServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *NetworkPoolServer) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworkPoolServer) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworkPoolServer) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NetworkPoolServer) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetServiceUrl

`func (o *NetworkPoolServer) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *NetworkPoolServer) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *NetworkPoolServer) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.

### HasServiceUrl

`func (o *NetworkPoolServer) HasServiceUrl() bool`

HasServiceUrl returns a boolean if a field has been set.

### GetServiceHost

`func (o *NetworkPoolServer) GetServiceHost() string`

GetServiceHost returns the ServiceHost field if non-nil, zero value otherwise.

### GetServiceHostOk

`func (o *NetworkPoolServer) GetServiceHostOk() (*string, bool)`

GetServiceHostOk returns a tuple with the ServiceHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHost

`func (o *NetworkPoolServer) SetServiceHost(v string)`

SetServiceHost sets ServiceHost field to given value.

### HasServiceHost

`func (o *NetworkPoolServer) HasServiceHost() bool`

HasServiceHost returns a boolean if a field has been set.

### GetServicePort

`func (o *NetworkPoolServer) GetServicePort() int32`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *NetworkPoolServer) GetServicePortOk() (*int32, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *NetworkPoolServer) SetServicePort(v int32)`

SetServicePort sets ServicePort field to given value.

### HasServicePort

`func (o *NetworkPoolServer) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.

### GetServiceMode

`func (o *NetworkPoolServer) GetServiceMode() string`

GetServiceMode returns the ServiceMode field if non-nil, zero value otherwise.

### GetServiceModeOk

`func (o *NetworkPoolServer) GetServiceModeOk() (*string, bool)`

GetServiceModeOk returns a tuple with the ServiceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMode

`func (o *NetworkPoolServer) SetServiceMode(v string)`

SetServiceMode sets ServiceMode field to given value.

### HasServiceMode

`func (o *NetworkPoolServer) HasServiceMode() bool`

HasServiceMode returns a boolean if a field has been set.

### GetServiceUsername

`func (o *NetworkPoolServer) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *NetworkPoolServer) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *NetworkPoolServer) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.

### HasServiceUsername

`func (o *NetworkPoolServer) HasServiceUsername() bool`

HasServiceUsername returns a boolean if a field has been set.

### GetServicePassword

`func (o *NetworkPoolServer) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *NetworkPoolServer) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *NetworkPoolServer) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *NetworkPoolServer) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### GetServicePasswordHash

`func (o *NetworkPoolServer) GetServicePasswordHash() string`

GetServicePasswordHash returns the ServicePasswordHash field if non-nil, zero value otherwise.

### GetServicePasswordHashOk

`func (o *NetworkPoolServer) GetServicePasswordHashOk() (*string, bool)`

GetServicePasswordHashOk returns a tuple with the ServicePasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePasswordHash

`func (o *NetworkPoolServer) SetServicePasswordHash(v string)`

SetServicePasswordHash sets ServicePasswordHash field to given value.

### HasServicePasswordHash

`func (o *NetworkPoolServer) HasServicePasswordHash() bool`

HasServicePasswordHash returns a boolean if a field has been set.

### GetServiceThrottleRate

`func (o *NetworkPoolServer) GetServiceThrottleRate() int64`

GetServiceThrottleRate returns the ServiceThrottleRate field if non-nil, zero value otherwise.

### GetServiceThrottleRateOk

`func (o *NetworkPoolServer) GetServiceThrottleRateOk() (*int64, bool)`

GetServiceThrottleRateOk returns a tuple with the ServiceThrottleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceThrottleRate

`func (o *NetworkPoolServer) SetServiceThrottleRate(v int64)`

SetServiceThrottleRate sets ServiceThrottleRate field to given value.

### HasServiceThrottleRate

`func (o *NetworkPoolServer) HasServiceThrottleRate() bool`

HasServiceThrottleRate returns a boolean if a field has been set.

### GetIgnoreSsl

`func (o *NetworkPoolServer) GetIgnoreSsl() bool`

GetIgnoreSsl returns the IgnoreSsl field if non-nil, zero value otherwise.

### GetIgnoreSslOk

`func (o *NetworkPoolServer) GetIgnoreSslOk() (*bool, bool)`

GetIgnoreSslOk returns a tuple with the IgnoreSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSsl

`func (o *NetworkPoolServer) SetIgnoreSsl(v bool)`

SetIgnoreSsl sets IgnoreSsl field to given value.

### HasIgnoreSsl

`func (o *NetworkPoolServer) HasIgnoreSsl() bool`

HasIgnoreSsl returns a boolean if a field has been set.

### GetStatus

`func (o *NetworkPoolServer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *NetworkPoolServer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *NetworkPoolServer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *NetworkPoolServer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *NetworkPoolServer) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *NetworkPoolServer) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *NetworkPoolServer) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *NetworkPoolServer) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetStatusDate

`func (o *NetworkPoolServer) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *NetworkPoolServer) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *NetworkPoolServer) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *NetworkPoolServer) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetConfig

`func (o *NetworkPoolServer) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *NetworkPoolServer) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *NetworkPoolServer) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *NetworkPoolServer) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetNetworkFilter

`func (o *NetworkPoolServer) GetNetworkFilter() string`

GetNetworkFilter returns the NetworkFilter field if non-nil, zero value otherwise.

### GetNetworkFilterOk

`func (o *NetworkPoolServer) GetNetworkFilterOk() (*string, bool)`

GetNetworkFilterOk returns a tuple with the NetworkFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFilter

`func (o *NetworkPoolServer) SetNetworkFilter(v string)`

SetNetworkFilter sets NetworkFilter field to given value.

### HasNetworkFilter

`func (o *NetworkPoolServer) HasNetworkFilter() bool`

HasNetworkFilter returns a boolean if a field has been set.

### GetZoneFilter

`func (o *NetworkPoolServer) GetZoneFilter() string`

GetZoneFilter returns the ZoneFilter field if non-nil, zero value otherwise.

### GetZoneFilterOk

`func (o *NetworkPoolServer) GetZoneFilterOk() (*string, bool)`

GetZoneFilterOk returns a tuple with the ZoneFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneFilter

`func (o *NetworkPoolServer) SetZoneFilter(v string)`

SetZoneFilter sets ZoneFilter field to given value.

### HasZoneFilter

`func (o *NetworkPoolServer) HasZoneFilter() bool`

HasZoneFilter returns a boolean if a field has been set.

### GetTenantMatch

`func (o *NetworkPoolServer) GetTenantMatch() string`

GetTenantMatch returns the TenantMatch field if non-nil, zero value otherwise.

### GetTenantMatchOk

`func (o *NetworkPoolServer) GetTenantMatchOk() (*string, bool)`

GetTenantMatchOk returns a tuple with the TenantMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantMatch

`func (o *NetworkPoolServer) SetTenantMatch(v string)`

SetTenantMatch sets TenantMatch field to given value.

### HasTenantMatch

`func (o *NetworkPoolServer) HasTenantMatch() bool`

HasTenantMatch returns a boolean if a field has been set.

### GetDateCreated

`func (o *NetworkPoolServer) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *NetworkPoolServer) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *NetworkPoolServer) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *NetworkPoolServer) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *NetworkPoolServer) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *NetworkPoolServer) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *NetworkPoolServer) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *NetworkPoolServer) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetAccount

`func (o *NetworkPoolServer) GetAccount() ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *NetworkPoolServer) GetAccountOk() (*ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *NetworkPoolServer) SetAccount(v ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *NetworkPoolServer) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetIntegration

`func (o *NetworkPoolServer) GetIntegration() ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInnerIntegration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *NetworkPoolServer) GetIntegrationOk() (*ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInnerIntegration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *NetworkPoolServer) SetIntegration(v ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInnerIntegration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *NetworkPoolServer) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetPools

`func (o *NetworkPoolServer) GetPools() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetPools returns the Pools field if non-nil, zero value otherwise.

### GetPoolsOk

`func (o *NetworkPoolServer) GetPoolsOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetPoolsOk returns a tuple with the Pools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPools

`func (o *NetworkPoolServer) SetPools(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetPools sets Pools field to given value.

### HasPools

`func (o *NetworkPoolServer) HasPools() bool`

HasPools returns a boolean if a field has been set.

### GetCredential

`func (o *NetworkPoolServer) GetCredential() ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInnerCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *NetworkPoolServer) GetCredentialOk() (*ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInnerCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *NetworkPoolServer) SetCredential(v ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInnerCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *NetworkPoolServer) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


