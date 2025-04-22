# InfobloxNetworkPoolServerUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Name | [optional] 
**Enabled** | Pointer to **bool** | Can be used to enable / disable the network pool server. | [optional] [default to true]
**ServiceUrl** | Pointer to **string** | URL | [optional] 
**ServiceUsername** | Pointer to **string** | Username | [optional] 
**ServicePassword** | Pointer to **string** | Password | [optional] 
**ServiceThrottleRate** | Pointer to **int64** | Throttle Rate | [optional] [default to 0]
**IgnoreSsl** | Pointer to **bool** | Disable SSL SNI Verification | [optional] 
**NetworkFilter** | Pointer to **string** | Network Filter | [optional] 
**ZoneFilter** | Pointer to **string** | Zone Filter | [optional] 
**TenantMatch** | Pointer to **string** | Tenant Match | [optional] 
**ServiceMode** | Pointer to **string** | IP Mode | [optional] [default to "static"]
**Config** | Pointer to [**InfobloxNetworkPoolServerConfig**](InfobloxNetworkPoolServerConfig.md) |  | [optional] 
**Credential** | Pointer to [**NSXNetworkServerCredential**](NSXNetworkServerCredential.md) |  | [optional] 

## Methods

### NewInfobloxNetworkPoolServerUpdate

`func NewInfobloxNetworkPoolServerUpdate() *InfobloxNetworkPoolServerUpdate`

NewInfobloxNetworkPoolServerUpdate instantiates a new InfobloxNetworkPoolServerUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfobloxNetworkPoolServerUpdateWithDefaults

`func NewInfobloxNetworkPoolServerUpdateWithDefaults() *InfobloxNetworkPoolServerUpdate`

NewInfobloxNetworkPoolServerUpdateWithDefaults instantiates a new InfobloxNetworkPoolServerUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InfobloxNetworkPoolServerUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InfobloxNetworkPoolServerUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InfobloxNetworkPoolServerUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InfobloxNetworkPoolServerUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *InfobloxNetworkPoolServerUpdate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InfobloxNetworkPoolServerUpdate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InfobloxNetworkPoolServerUpdate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InfobloxNetworkPoolServerUpdate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetServiceUrl

`func (o *InfobloxNetworkPoolServerUpdate) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *InfobloxNetworkPoolServerUpdate) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *InfobloxNetworkPoolServerUpdate) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.

### HasServiceUrl

`func (o *InfobloxNetworkPoolServerUpdate) HasServiceUrl() bool`

HasServiceUrl returns a boolean if a field has been set.

### GetServiceUsername

`func (o *InfobloxNetworkPoolServerUpdate) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *InfobloxNetworkPoolServerUpdate) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *InfobloxNetworkPoolServerUpdate) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.

### HasServiceUsername

`func (o *InfobloxNetworkPoolServerUpdate) HasServiceUsername() bool`

HasServiceUsername returns a boolean if a field has been set.

### GetServicePassword

`func (o *InfobloxNetworkPoolServerUpdate) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *InfobloxNetworkPoolServerUpdate) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *InfobloxNetworkPoolServerUpdate) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *InfobloxNetworkPoolServerUpdate) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### GetServiceThrottleRate

`func (o *InfobloxNetworkPoolServerUpdate) GetServiceThrottleRate() int64`

GetServiceThrottleRate returns the ServiceThrottleRate field if non-nil, zero value otherwise.

### GetServiceThrottleRateOk

`func (o *InfobloxNetworkPoolServerUpdate) GetServiceThrottleRateOk() (*int64, bool)`

GetServiceThrottleRateOk returns a tuple with the ServiceThrottleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceThrottleRate

`func (o *InfobloxNetworkPoolServerUpdate) SetServiceThrottleRate(v int64)`

SetServiceThrottleRate sets ServiceThrottleRate field to given value.

### HasServiceThrottleRate

`func (o *InfobloxNetworkPoolServerUpdate) HasServiceThrottleRate() bool`

HasServiceThrottleRate returns a boolean if a field has been set.

### GetIgnoreSsl

`func (o *InfobloxNetworkPoolServerUpdate) GetIgnoreSsl() bool`

GetIgnoreSsl returns the IgnoreSsl field if non-nil, zero value otherwise.

### GetIgnoreSslOk

`func (o *InfobloxNetworkPoolServerUpdate) GetIgnoreSslOk() (*bool, bool)`

GetIgnoreSslOk returns a tuple with the IgnoreSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSsl

`func (o *InfobloxNetworkPoolServerUpdate) SetIgnoreSsl(v bool)`

SetIgnoreSsl sets IgnoreSsl field to given value.

### HasIgnoreSsl

`func (o *InfobloxNetworkPoolServerUpdate) HasIgnoreSsl() bool`

HasIgnoreSsl returns a boolean if a field has been set.

### GetNetworkFilter

`func (o *InfobloxNetworkPoolServerUpdate) GetNetworkFilter() string`

GetNetworkFilter returns the NetworkFilter field if non-nil, zero value otherwise.

### GetNetworkFilterOk

`func (o *InfobloxNetworkPoolServerUpdate) GetNetworkFilterOk() (*string, bool)`

GetNetworkFilterOk returns a tuple with the NetworkFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFilter

`func (o *InfobloxNetworkPoolServerUpdate) SetNetworkFilter(v string)`

SetNetworkFilter sets NetworkFilter field to given value.

### HasNetworkFilter

`func (o *InfobloxNetworkPoolServerUpdate) HasNetworkFilter() bool`

HasNetworkFilter returns a boolean if a field has been set.

### GetZoneFilter

`func (o *InfobloxNetworkPoolServerUpdate) GetZoneFilter() string`

GetZoneFilter returns the ZoneFilter field if non-nil, zero value otherwise.

### GetZoneFilterOk

`func (o *InfobloxNetworkPoolServerUpdate) GetZoneFilterOk() (*string, bool)`

GetZoneFilterOk returns a tuple with the ZoneFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneFilter

`func (o *InfobloxNetworkPoolServerUpdate) SetZoneFilter(v string)`

SetZoneFilter sets ZoneFilter field to given value.

### HasZoneFilter

`func (o *InfobloxNetworkPoolServerUpdate) HasZoneFilter() bool`

HasZoneFilter returns a boolean if a field has been set.

### GetTenantMatch

`func (o *InfobloxNetworkPoolServerUpdate) GetTenantMatch() string`

GetTenantMatch returns the TenantMatch field if non-nil, zero value otherwise.

### GetTenantMatchOk

`func (o *InfobloxNetworkPoolServerUpdate) GetTenantMatchOk() (*string, bool)`

GetTenantMatchOk returns a tuple with the TenantMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantMatch

`func (o *InfobloxNetworkPoolServerUpdate) SetTenantMatch(v string)`

SetTenantMatch sets TenantMatch field to given value.

### HasTenantMatch

`func (o *InfobloxNetworkPoolServerUpdate) HasTenantMatch() bool`

HasTenantMatch returns a boolean if a field has been set.

### GetServiceMode

`func (o *InfobloxNetworkPoolServerUpdate) GetServiceMode() string`

GetServiceMode returns the ServiceMode field if non-nil, zero value otherwise.

### GetServiceModeOk

`func (o *InfobloxNetworkPoolServerUpdate) GetServiceModeOk() (*string, bool)`

GetServiceModeOk returns a tuple with the ServiceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMode

`func (o *InfobloxNetworkPoolServerUpdate) SetServiceMode(v string)`

SetServiceMode sets ServiceMode field to given value.

### HasServiceMode

`func (o *InfobloxNetworkPoolServerUpdate) HasServiceMode() bool`

HasServiceMode returns a boolean if a field has been set.

### GetConfig

`func (o *InfobloxNetworkPoolServerUpdate) GetConfig() InfobloxNetworkPoolServerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *InfobloxNetworkPoolServerUpdate) GetConfigOk() (*InfobloxNetworkPoolServerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *InfobloxNetworkPoolServerUpdate) SetConfig(v InfobloxNetworkPoolServerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *InfobloxNetworkPoolServerUpdate) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCredential

`func (o *InfobloxNetworkPoolServerUpdate) GetCredential() NSXNetworkServerCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *InfobloxNetworkPoolServerUpdate) GetCredentialOk() (*NSXNetworkServerCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *InfobloxNetworkPoolServerUpdate) SetCredential(v NSXNetworkServerCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *InfobloxNetworkPoolServerUpdate) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


