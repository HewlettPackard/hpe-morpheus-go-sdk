# UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Network Pool Server ID | [optional] 
**Type** | Pointer to [**UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServerType**](UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServerType.md) |  | [optional] 
**Name** | Pointer to **string** | Name | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**ServiceUrl** | Pointer to **NullableString** | Service URL | [optional] 
**ServiceHost** | Pointer to **NullableString** | Service Host | [optional] 
**ServicePort** | Pointer to **NullableInt32** | Service Port | [optional] 
**ServiceMode** | Pointer to **NullableString** | Service Mode | [optional] 
**ServiceUsername** | Pointer to **NullableString** | Service Username | [optional] 
**ServicePassword** | Pointer to **NullableString** | Service Password | [optional] 
**ServicePasswordHash** | Pointer to **string** |  | [optional] 
**ServiceThrottleRate** | Pointer to **NullableInt64** | Throttle Rate | [optional] [default to 0]
**IgnoreSsl** | Pointer to **NullableBool** | Disable SSL SNI Verification | [optional] [default to true]
**Status** | Pointer to **string** | Status | [optional] 
**StatusMessage** | Pointer to **NullableString** | Status Message | [optional] 
**StatusDate** | Pointer to **NullableTime** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** | Config object varies with pool server type. | [optional] 
**NetworkFilter** | Pointer to **NullableString** | Network Filter | [optional] 
**ZoneFilter** | Pointer to **NullableString** | Zone Filter | [optional] 
**TenantMatch** | Pointer to **NullableString** | Tenant Match | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Account** | Pointer to [**UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServerAccount**](UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServerAccount.md) |  | [optional] 
**Integration** | Pointer to [**UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServerIntegration**](UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServerIntegration.md) |  | [optional] 
**Pools** | Pointer to [**[]UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServerPoolsInner**](UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServerPoolsInner.md) |  | [optional] 
**Credential** | Pointer to [**UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServerCredential**](UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServerCredential.md) |  | [optional] 

## Methods

### NewUpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer

`func NewUpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer() *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer`

NewUpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer instantiates a new UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNetworkPoolServer200ResponseAllOfNetworkPoolServerWithDefaults

`func NewUpdateNetworkPoolServer200ResponseAllOfNetworkPoolServerWithDefaults() *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer`

NewUpdateNetworkPoolServer200ResponseAllOfNetworkPoolServerWithDefaults instantiates a new UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetType() UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetTypeOk() (*UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetType(v UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServerType)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetServiceUrl

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.

### HasServiceUrl

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) HasServiceUrl() bool`

HasServiceUrl returns a boolean if a field has been set.

### SetServiceUrlNil

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetServiceUrlNil(b bool)`

 SetServiceUrlNil sets the value for ServiceUrl to be an explicit nil

### UnsetServiceUrl
`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) UnsetServiceUrl()`

UnsetServiceUrl ensures that no value is present for ServiceUrl, not even an explicit nil
### GetServiceHost

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetServiceHost() string`

GetServiceHost returns the ServiceHost field if non-nil, zero value otherwise.

### GetServiceHostOk

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetServiceHostOk() (*string, bool)`

GetServiceHostOk returns a tuple with the ServiceHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHost

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetServiceHost(v string)`

SetServiceHost sets ServiceHost field to given value.

### HasServiceHost

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) HasServiceHost() bool`

HasServiceHost returns a boolean if a field has been set.

### SetServiceHostNil

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetServiceHostNil(b bool)`

 SetServiceHostNil sets the value for ServiceHost to be an explicit nil

### UnsetServiceHost
`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) UnsetServiceHost()`

UnsetServiceHost ensures that no value is present for ServiceHost, not even an explicit nil
### GetServicePort

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetServicePort() int32`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetServicePortOk() (*int32, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetServicePort(v int32)`

SetServicePort sets ServicePort field to given value.

### HasServicePort

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.

### SetServicePortNil

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetServicePortNil(b bool)`

 SetServicePortNil sets the value for ServicePort to be an explicit nil

### UnsetServicePort
`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) UnsetServicePort()`

UnsetServicePort ensures that no value is present for ServicePort, not even an explicit nil
### GetServiceMode

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetServiceMode() string`

GetServiceMode returns the ServiceMode field if non-nil, zero value otherwise.

### GetServiceModeOk

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetServiceModeOk() (*string, bool)`

GetServiceModeOk returns a tuple with the ServiceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMode

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetServiceMode(v string)`

SetServiceMode sets ServiceMode field to given value.

### HasServiceMode

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) HasServiceMode() bool`

HasServiceMode returns a boolean if a field has been set.

### SetServiceModeNil

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetServiceModeNil(b bool)`

 SetServiceModeNil sets the value for ServiceMode to be an explicit nil

### UnsetServiceMode
`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) UnsetServiceMode()`

UnsetServiceMode ensures that no value is present for ServiceMode, not even an explicit nil
### GetServiceUsername

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.

### HasServiceUsername

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) HasServiceUsername() bool`

HasServiceUsername returns a boolean if a field has been set.

### SetServiceUsernameNil

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetServiceUsernameNil(b bool)`

 SetServiceUsernameNil sets the value for ServiceUsername to be an explicit nil

### UnsetServiceUsername
`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) UnsetServiceUsername()`

UnsetServiceUsername ensures that no value is present for ServiceUsername, not even an explicit nil
### GetServicePassword

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### SetServicePasswordNil

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetServicePasswordNil(b bool)`

 SetServicePasswordNil sets the value for ServicePassword to be an explicit nil

### UnsetServicePassword
`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) UnsetServicePassword()`

UnsetServicePassword ensures that no value is present for ServicePassword, not even an explicit nil
### GetServicePasswordHash

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetServicePasswordHash() string`

GetServicePasswordHash returns the ServicePasswordHash field if non-nil, zero value otherwise.

### GetServicePasswordHashOk

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetServicePasswordHashOk() (*string, bool)`

GetServicePasswordHashOk returns a tuple with the ServicePasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePasswordHash

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetServicePasswordHash(v string)`

SetServicePasswordHash sets ServicePasswordHash field to given value.

### HasServicePasswordHash

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) HasServicePasswordHash() bool`

HasServicePasswordHash returns a boolean if a field has been set.

### GetServiceThrottleRate

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetServiceThrottleRate() int64`

GetServiceThrottleRate returns the ServiceThrottleRate field if non-nil, zero value otherwise.

### GetServiceThrottleRateOk

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetServiceThrottleRateOk() (*int64, bool)`

GetServiceThrottleRateOk returns a tuple with the ServiceThrottleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceThrottleRate

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetServiceThrottleRate(v int64)`

SetServiceThrottleRate sets ServiceThrottleRate field to given value.

### HasServiceThrottleRate

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) HasServiceThrottleRate() bool`

HasServiceThrottleRate returns a boolean if a field has been set.

### SetServiceThrottleRateNil

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetServiceThrottleRateNil(b bool)`

 SetServiceThrottleRateNil sets the value for ServiceThrottleRate to be an explicit nil

### UnsetServiceThrottleRate
`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) UnsetServiceThrottleRate()`

UnsetServiceThrottleRate ensures that no value is present for ServiceThrottleRate, not even an explicit nil
### GetIgnoreSsl

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetIgnoreSsl() bool`

GetIgnoreSsl returns the IgnoreSsl field if non-nil, zero value otherwise.

### GetIgnoreSslOk

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetIgnoreSslOk() (*bool, bool)`

GetIgnoreSslOk returns a tuple with the IgnoreSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSsl

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetIgnoreSsl(v bool)`

SetIgnoreSsl sets IgnoreSsl field to given value.

### HasIgnoreSsl

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) HasIgnoreSsl() bool`

HasIgnoreSsl returns a boolean if a field has been set.

### SetIgnoreSslNil

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetIgnoreSslNil(b bool)`

 SetIgnoreSslNil sets the value for IgnoreSsl to be an explicit nil

### UnsetIgnoreSsl
`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) UnsetIgnoreSsl()`

UnsetIgnoreSsl ensures that no value is present for IgnoreSsl, not even an explicit nil
### GetStatus

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetStatusDate

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### SetStatusDateNil

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetStatusDateNil(b bool)`

 SetStatusDateNil sets the value for StatusDate to be an explicit nil

### UnsetStatusDate
`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) UnsetStatusDate()`

UnsetStatusDate ensures that no value is present for StatusDate, not even an explicit nil
### GetConfig

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetNetworkFilter

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetNetworkFilter() string`

GetNetworkFilter returns the NetworkFilter field if non-nil, zero value otherwise.

### GetNetworkFilterOk

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetNetworkFilterOk() (*string, bool)`

GetNetworkFilterOk returns a tuple with the NetworkFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFilter

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetNetworkFilter(v string)`

SetNetworkFilter sets NetworkFilter field to given value.

### HasNetworkFilter

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) HasNetworkFilter() bool`

HasNetworkFilter returns a boolean if a field has been set.

### SetNetworkFilterNil

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetNetworkFilterNil(b bool)`

 SetNetworkFilterNil sets the value for NetworkFilter to be an explicit nil

### UnsetNetworkFilter
`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) UnsetNetworkFilter()`

UnsetNetworkFilter ensures that no value is present for NetworkFilter, not even an explicit nil
### GetZoneFilter

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetZoneFilter() string`

GetZoneFilter returns the ZoneFilter field if non-nil, zero value otherwise.

### GetZoneFilterOk

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetZoneFilterOk() (*string, bool)`

GetZoneFilterOk returns a tuple with the ZoneFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneFilter

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetZoneFilter(v string)`

SetZoneFilter sets ZoneFilter field to given value.

### HasZoneFilter

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) HasZoneFilter() bool`

HasZoneFilter returns a boolean if a field has been set.

### SetZoneFilterNil

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetZoneFilterNil(b bool)`

 SetZoneFilterNil sets the value for ZoneFilter to be an explicit nil

### UnsetZoneFilter
`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) UnsetZoneFilter()`

UnsetZoneFilter ensures that no value is present for ZoneFilter, not even an explicit nil
### GetTenantMatch

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetTenantMatch() string`

GetTenantMatch returns the TenantMatch field if non-nil, zero value otherwise.

### GetTenantMatchOk

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetTenantMatchOk() (*string, bool)`

GetTenantMatchOk returns a tuple with the TenantMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantMatch

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetTenantMatch(v string)`

SetTenantMatch sets TenantMatch field to given value.

### HasTenantMatch

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) HasTenantMatch() bool`

HasTenantMatch returns a boolean if a field has been set.

### SetTenantMatchNil

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetTenantMatchNil(b bool)`

 SetTenantMatchNil sets the value for TenantMatch to be an explicit nil

### UnsetTenantMatch
`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) UnsetTenantMatch()`

UnsetTenantMatch ensures that no value is present for TenantMatch, not even an explicit nil
### GetDateCreated

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetAccount

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetAccount() UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetAccountOk() (*UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetAccount(v UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetIntegration

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetIntegration() UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServerIntegration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetIntegrationOk() (*UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServerIntegration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetIntegration(v UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServerIntegration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetPools

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetPools() []UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServerPoolsInner`

GetPools returns the Pools field if non-nil, zero value otherwise.

### GetPoolsOk

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetPoolsOk() (*[]UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServerPoolsInner, bool)`

GetPoolsOk returns a tuple with the Pools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPools

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetPools(v []UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServerPoolsInner)`

SetPools sets Pools field to given value.

### HasPools

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) HasPools() bool`

HasPools returns a boolean if a field has been set.

### GetCredential

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetCredential() UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServerCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) GetCredentialOk() (*UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServerCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) SetCredential(v UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServerCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *UpdateNetworkPoolServer200ResponseAllOfNetworkPoolServer) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


