# ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner

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

### NewListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner

`func NewListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner() *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner`

NewListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner instantiates a new ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNetworkPoolServers200ResponseAllOfNetworkPoolServersInnerWithDefaults

`func NewListNetworkPoolServers200ResponseAllOfNetworkPoolServersInnerWithDefaults() *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner`

NewListNetworkPoolServers200ResponseAllOfNetworkPoolServersInnerWithDefaults instantiates a new ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetType() ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInnerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetTypeOk() (*ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInnerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) SetType(v ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInnerType)`

SetType sets Type field to given value.

### HasType

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetServiceUrl

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.

### HasServiceUrl

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) HasServiceUrl() bool`

HasServiceUrl returns a boolean if a field has been set.

### GetServiceHost

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetServiceHost() string`

GetServiceHost returns the ServiceHost field if non-nil, zero value otherwise.

### GetServiceHostOk

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetServiceHostOk() (*string, bool)`

GetServiceHostOk returns a tuple with the ServiceHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHost

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) SetServiceHost(v string)`

SetServiceHost sets ServiceHost field to given value.

### HasServiceHost

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) HasServiceHost() bool`

HasServiceHost returns a boolean if a field has been set.

### GetServicePort

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetServicePort() int32`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetServicePortOk() (*int32, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) SetServicePort(v int32)`

SetServicePort sets ServicePort field to given value.

### HasServicePort

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.

### GetServiceMode

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetServiceMode() string`

GetServiceMode returns the ServiceMode field if non-nil, zero value otherwise.

### GetServiceModeOk

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetServiceModeOk() (*string, bool)`

GetServiceModeOk returns a tuple with the ServiceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMode

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) SetServiceMode(v string)`

SetServiceMode sets ServiceMode field to given value.

### HasServiceMode

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) HasServiceMode() bool`

HasServiceMode returns a boolean if a field has been set.

### GetServiceUsername

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.

### HasServiceUsername

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) HasServiceUsername() bool`

HasServiceUsername returns a boolean if a field has been set.

### GetServicePassword

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### GetServicePasswordHash

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetServicePasswordHash() string`

GetServicePasswordHash returns the ServicePasswordHash field if non-nil, zero value otherwise.

### GetServicePasswordHashOk

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetServicePasswordHashOk() (*string, bool)`

GetServicePasswordHashOk returns a tuple with the ServicePasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePasswordHash

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) SetServicePasswordHash(v string)`

SetServicePasswordHash sets ServicePasswordHash field to given value.

### HasServicePasswordHash

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) HasServicePasswordHash() bool`

HasServicePasswordHash returns a boolean if a field has been set.

### GetServiceThrottleRate

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetServiceThrottleRate() int64`

GetServiceThrottleRate returns the ServiceThrottleRate field if non-nil, zero value otherwise.

### GetServiceThrottleRateOk

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetServiceThrottleRateOk() (*int64, bool)`

GetServiceThrottleRateOk returns a tuple with the ServiceThrottleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceThrottleRate

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) SetServiceThrottleRate(v int64)`

SetServiceThrottleRate sets ServiceThrottleRate field to given value.

### HasServiceThrottleRate

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) HasServiceThrottleRate() bool`

HasServiceThrottleRate returns a boolean if a field has been set.

### GetIgnoreSsl

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetIgnoreSsl() bool`

GetIgnoreSsl returns the IgnoreSsl field if non-nil, zero value otherwise.

### GetIgnoreSslOk

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetIgnoreSslOk() (*bool, bool)`

GetIgnoreSslOk returns a tuple with the IgnoreSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSsl

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) SetIgnoreSsl(v bool)`

SetIgnoreSsl sets IgnoreSsl field to given value.

### HasIgnoreSsl

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) HasIgnoreSsl() bool`

HasIgnoreSsl returns a boolean if a field has been set.

### GetStatus

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetStatusDate

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### GetConfig

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetNetworkFilter

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetNetworkFilter() string`

GetNetworkFilter returns the NetworkFilter field if non-nil, zero value otherwise.

### GetNetworkFilterOk

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetNetworkFilterOk() (*string, bool)`

GetNetworkFilterOk returns a tuple with the NetworkFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFilter

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) SetNetworkFilter(v string)`

SetNetworkFilter sets NetworkFilter field to given value.

### HasNetworkFilter

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) HasNetworkFilter() bool`

HasNetworkFilter returns a boolean if a field has been set.

### GetZoneFilter

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetZoneFilter() string`

GetZoneFilter returns the ZoneFilter field if non-nil, zero value otherwise.

### GetZoneFilterOk

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetZoneFilterOk() (*string, bool)`

GetZoneFilterOk returns a tuple with the ZoneFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneFilter

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) SetZoneFilter(v string)`

SetZoneFilter sets ZoneFilter field to given value.

### HasZoneFilter

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) HasZoneFilter() bool`

HasZoneFilter returns a boolean if a field has been set.

### GetTenantMatch

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetTenantMatch() string`

GetTenantMatch returns the TenantMatch field if non-nil, zero value otherwise.

### GetTenantMatchOk

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetTenantMatchOk() (*string, bool)`

GetTenantMatchOk returns a tuple with the TenantMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantMatch

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) SetTenantMatch(v string)`

SetTenantMatch sets TenantMatch field to given value.

### HasTenantMatch

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) HasTenantMatch() bool`

HasTenantMatch returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetAccount

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetAccount() ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetAccountOk() (*ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) SetAccount(v ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetIntegration

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetIntegration() ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInnerIntegration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetIntegrationOk() (*ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInnerIntegration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) SetIntegration(v ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInnerIntegration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetPools

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetPools() []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner`

GetPools returns the Pools field if non-nil, zero value otherwise.

### GetPoolsOk

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetPoolsOk() (*[]ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner, bool)`

GetPoolsOk returns a tuple with the Pools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPools

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) SetPools(v []ListApplianceSettings200ResponseApplianceSettingsEnabledZoneTypesInner)`

SetPools sets Pools field to given value.

### HasPools

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) HasPools() bool`

HasPools returns a boolean if a field has been set.

### GetCredential

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetCredential() ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInnerCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) GetCredentialOk() (*ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInnerCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) SetCredential(v ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInnerCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *ListNetworkPoolServers200ResponseAllOfNetworkPoolServersInner) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


