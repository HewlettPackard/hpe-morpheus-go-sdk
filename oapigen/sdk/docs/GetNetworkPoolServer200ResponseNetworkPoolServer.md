# GetNetworkPoolServer200ResponseNetworkPoolServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Network Pool Server ID | [optional] 
**Type** | Pointer to [**GetNetworkPoolServer200ResponseNetworkPoolServerType**](GetNetworkPoolServer200ResponseNetworkPoolServerType.md) |  | [optional] 
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
**Account** | Pointer to [**GetNetworkPoolServer200ResponseNetworkPoolServerAccount**](GetNetworkPoolServer200ResponseNetworkPoolServerAccount.md) |  | [optional] 
**Integration** | Pointer to [**GetNetworkPoolServer200ResponseNetworkPoolServerIntegration**](GetNetworkPoolServer200ResponseNetworkPoolServerIntegration.md) |  | [optional] 
**Pools** | Pointer to [**[]GetNetworkPoolServer200ResponseNetworkPoolServerPoolsInner**](GetNetworkPoolServer200ResponseNetworkPoolServerPoolsInner.md) |  | [optional] 
**Credential** | Pointer to [**GetNetworkPoolServer200ResponseNetworkPoolServerCredential**](GetNetworkPoolServer200ResponseNetworkPoolServerCredential.md) |  | [optional] 

## Methods

### NewGetNetworkPoolServer200ResponseNetworkPoolServer

`func NewGetNetworkPoolServer200ResponseNetworkPoolServer() *GetNetworkPoolServer200ResponseNetworkPoolServer`

NewGetNetworkPoolServer200ResponseNetworkPoolServer instantiates a new GetNetworkPoolServer200ResponseNetworkPoolServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkPoolServer200ResponseNetworkPoolServerWithDefaults

`func NewGetNetworkPoolServer200ResponseNetworkPoolServerWithDefaults() *GetNetworkPoolServer200ResponseNetworkPoolServer`

NewGetNetworkPoolServer200ResponseNetworkPoolServerWithDefaults instantiates a new GetNetworkPoolServer200ResponseNetworkPoolServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetType

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetType() GetNetworkPoolServer200ResponseNetworkPoolServerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetTypeOk() (*GetNetworkPoolServer200ResponseNetworkPoolServerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetType(v GetNetworkPoolServer200ResponseNetworkPoolServerType)`

SetType sets Type field to given value.

### HasType

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetServiceUrl

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.

### HasServiceUrl

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) HasServiceUrl() bool`

HasServiceUrl returns a boolean if a field has been set.

### SetServiceUrlNil

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetServiceUrlNil(b bool)`

 SetServiceUrlNil sets the value for ServiceUrl to be an explicit nil

### UnsetServiceUrl
`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) UnsetServiceUrl()`

UnsetServiceUrl ensures that no value is present for ServiceUrl, not even an explicit nil
### GetServiceHost

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetServiceHost() string`

GetServiceHost returns the ServiceHost field if non-nil, zero value otherwise.

### GetServiceHostOk

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetServiceHostOk() (*string, bool)`

GetServiceHostOk returns a tuple with the ServiceHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceHost

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetServiceHost(v string)`

SetServiceHost sets ServiceHost field to given value.

### HasServiceHost

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) HasServiceHost() bool`

HasServiceHost returns a boolean if a field has been set.

### SetServiceHostNil

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetServiceHostNil(b bool)`

 SetServiceHostNil sets the value for ServiceHost to be an explicit nil

### UnsetServiceHost
`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) UnsetServiceHost()`

UnsetServiceHost ensures that no value is present for ServiceHost, not even an explicit nil
### GetServicePort

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetServicePort() int32`

GetServicePort returns the ServicePort field if non-nil, zero value otherwise.

### GetServicePortOk

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetServicePortOk() (*int32, bool)`

GetServicePortOk returns a tuple with the ServicePort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePort

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetServicePort(v int32)`

SetServicePort sets ServicePort field to given value.

### HasServicePort

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) HasServicePort() bool`

HasServicePort returns a boolean if a field has been set.

### SetServicePortNil

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetServicePortNil(b bool)`

 SetServicePortNil sets the value for ServicePort to be an explicit nil

### UnsetServicePort
`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) UnsetServicePort()`

UnsetServicePort ensures that no value is present for ServicePort, not even an explicit nil
### GetServiceMode

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetServiceMode() string`

GetServiceMode returns the ServiceMode field if non-nil, zero value otherwise.

### GetServiceModeOk

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetServiceModeOk() (*string, bool)`

GetServiceModeOk returns a tuple with the ServiceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMode

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetServiceMode(v string)`

SetServiceMode sets ServiceMode field to given value.

### HasServiceMode

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) HasServiceMode() bool`

HasServiceMode returns a boolean if a field has been set.

### SetServiceModeNil

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetServiceModeNil(b bool)`

 SetServiceModeNil sets the value for ServiceMode to be an explicit nil

### UnsetServiceMode
`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) UnsetServiceMode()`

UnsetServiceMode ensures that no value is present for ServiceMode, not even an explicit nil
### GetServiceUsername

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.

### HasServiceUsername

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) HasServiceUsername() bool`

HasServiceUsername returns a boolean if a field has been set.

### SetServiceUsernameNil

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetServiceUsernameNil(b bool)`

 SetServiceUsernameNil sets the value for ServiceUsername to be an explicit nil

### UnsetServiceUsername
`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) UnsetServiceUsername()`

UnsetServiceUsername ensures that no value is present for ServiceUsername, not even an explicit nil
### GetServicePassword

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### SetServicePasswordNil

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetServicePasswordNil(b bool)`

 SetServicePasswordNil sets the value for ServicePassword to be an explicit nil

### UnsetServicePassword
`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) UnsetServicePassword()`

UnsetServicePassword ensures that no value is present for ServicePassword, not even an explicit nil
### GetServicePasswordHash

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetServicePasswordHash() string`

GetServicePasswordHash returns the ServicePasswordHash field if non-nil, zero value otherwise.

### GetServicePasswordHashOk

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetServicePasswordHashOk() (*string, bool)`

GetServicePasswordHashOk returns a tuple with the ServicePasswordHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePasswordHash

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetServicePasswordHash(v string)`

SetServicePasswordHash sets ServicePasswordHash field to given value.

### HasServicePasswordHash

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) HasServicePasswordHash() bool`

HasServicePasswordHash returns a boolean if a field has been set.

### GetServiceThrottleRate

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetServiceThrottleRate() int64`

GetServiceThrottleRate returns the ServiceThrottleRate field if non-nil, zero value otherwise.

### GetServiceThrottleRateOk

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetServiceThrottleRateOk() (*int64, bool)`

GetServiceThrottleRateOk returns a tuple with the ServiceThrottleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceThrottleRate

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetServiceThrottleRate(v int64)`

SetServiceThrottleRate sets ServiceThrottleRate field to given value.

### HasServiceThrottleRate

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) HasServiceThrottleRate() bool`

HasServiceThrottleRate returns a boolean if a field has been set.

### SetServiceThrottleRateNil

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetServiceThrottleRateNil(b bool)`

 SetServiceThrottleRateNil sets the value for ServiceThrottleRate to be an explicit nil

### UnsetServiceThrottleRate
`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) UnsetServiceThrottleRate()`

UnsetServiceThrottleRate ensures that no value is present for ServiceThrottleRate, not even an explicit nil
### GetIgnoreSsl

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetIgnoreSsl() bool`

GetIgnoreSsl returns the IgnoreSsl field if non-nil, zero value otherwise.

### GetIgnoreSslOk

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetIgnoreSslOk() (*bool, bool)`

GetIgnoreSslOk returns a tuple with the IgnoreSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSsl

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetIgnoreSsl(v bool)`

SetIgnoreSsl sets IgnoreSsl field to given value.

### HasIgnoreSsl

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) HasIgnoreSsl() bool`

HasIgnoreSsl returns a boolean if a field has been set.

### SetIgnoreSslNil

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetIgnoreSslNil(b bool)`

 SetIgnoreSslNil sets the value for IgnoreSsl to be an explicit nil

### UnsetIgnoreSsl
`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) UnsetIgnoreSsl()`

UnsetIgnoreSsl ensures that no value is present for IgnoreSsl, not even an explicit nil
### GetStatus

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetStatusDate

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetStatusDate() time.Time`

GetStatusDate returns the StatusDate field if non-nil, zero value otherwise.

### GetStatusDateOk

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetStatusDateOk() (*time.Time, bool)`

GetStatusDateOk returns a tuple with the StatusDate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusDate

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetStatusDate(v time.Time)`

SetStatusDate sets StatusDate field to given value.

### HasStatusDate

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) HasStatusDate() bool`

HasStatusDate returns a boolean if a field has been set.

### SetStatusDateNil

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetStatusDateNil(b bool)`

 SetStatusDateNil sets the value for StatusDate to be an explicit nil

### UnsetStatusDate
`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) UnsetStatusDate()`

UnsetStatusDate ensures that no value is present for StatusDate, not even an explicit nil
### GetConfig

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetNetworkFilter

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetNetworkFilter() string`

GetNetworkFilter returns the NetworkFilter field if non-nil, zero value otherwise.

### GetNetworkFilterOk

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetNetworkFilterOk() (*string, bool)`

GetNetworkFilterOk returns a tuple with the NetworkFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFilter

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetNetworkFilter(v string)`

SetNetworkFilter sets NetworkFilter field to given value.

### HasNetworkFilter

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) HasNetworkFilter() bool`

HasNetworkFilter returns a boolean if a field has been set.

### SetNetworkFilterNil

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetNetworkFilterNil(b bool)`

 SetNetworkFilterNil sets the value for NetworkFilter to be an explicit nil

### UnsetNetworkFilter
`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) UnsetNetworkFilter()`

UnsetNetworkFilter ensures that no value is present for NetworkFilter, not even an explicit nil
### GetZoneFilter

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetZoneFilter() string`

GetZoneFilter returns the ZoneFilter field if non-nil, zero value otherwise.

### GetZoneFilterOk

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetZoneFilterOk() (*string, bool)`

GetZoneFilterOk returns a tuple with the ZoneFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneFilter

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetZoneFilter(v string)`

SetZoneFilter sets ZoneFilter field to given value.

### HasZoneFilter

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) HasZoneFilter() bool`

HasZoneFilter returns a boolean if a field has been set.

### SetZoneFilterNil

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetZoneFilterNil(b bool)`

 SetZoneFilterNil sets the value for ZoneFilter to be an explicit nil

### UnsetZoneFilter
`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) UnsetZoneFilter()`

UnsetZoneFilter ensures that no value is present for ZoneFilter, not even an explicit nil
### GetTenantMatch

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetTenantMatch() string`

GetTenantMatch returns the TenantMatch field if non-nil, zero value otherwise.

### GetTenantMatchOk

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetTenantMatchOk() (*string, bool)`

GetTenantMatchOk returns a tuple with the TenantMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantMatch

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetTenantMatch(v string)`

SetTenantMatch sets TenantMatch field to given value.

### HasTenantMatch

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) HasTenantMatch() bool`

HasTenantMatch returns a boolean if a field has been set.

### SetTenantMatchNil

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetTenantMatchNil(b bool)`

 SetTenantMatchNil sets the value for TenantMatch to be an explicit nil

### UnsetTenantMatch
`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) UnsetTenantMatch()`

UnsetTenantMatch ensures that no value is present for TenantMatch, not even an explicit nil
### GetDateCreated

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetAccount

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetAccount() GetNetworkPoolServer200ResponseNetworkPoolServerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetAccountOk() (*GetNetworkPoolServer200ResponseNetworkPoolServerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetAccount(v GetNetworkPoolServer200ResponseNetworkPoolServerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetIntegration

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetIntegration() GetNetworkPoolServer200ResponseNetworkPoolServerIntegration`

GetIntegration returns the Integration field if non-nil, zero value otherwise.

### GetIntegrationOk

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetIntegrationOk() (*GetNetworkPoolServer200ResponseNetworkPoolServerIntegration, bool)`

GetIntegrationOk returns a tuple with the Integration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegration

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetIntegration(v GetNetworkPoolServer200ResponseNetworkPoolServerIntegration)`

SetIntegration sets Integration field to given value.

### HasIntegration

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) HasIntegration() bool`

HasIntegration returns a boolean if a field has been set.

### GetPools

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetPools() []GetNetworkPoolServer200ResponseNetworkPoolServerPoolsInner`

GetPools returns the Pools field if non-nil, zero value otherwise.

### GetPoolsOk

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetPoolsOk() (*[]GetNetworkPoolServer200ResponseNetworkPoolServerPoolsInner, bool)`

GetPoolsOk returns a tuple with the Pools field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPools

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetPools(v []GetNetworkPoolServer200ResponseNetworkPoolServerPoolsInner)`

SetPools sets Pools field to given value.

### HasPools

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) HasPools() bool`

HasPools returns a boolean if a field has been set.

### GetCredential

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetCredential() GetNetworkPoolServer200ResponseNetworkPoolServerCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) GetCredentialOk() (*GetNetworkPoolServer200ResponseNetworkPoolServerCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) SetCredential(v GetNetworkPoolServer200ResponseNetworkPoolServerCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *GetNetworkPoolServer200ResponseNetworkPoolServer) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


