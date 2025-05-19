# NetworkPoolServerUpdateBluecat

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
**Config** | Pointer to [**BluecatNetworkPoolServerConfig**](BluecatNetworkPoolServerConfig.md) |  | [optional] 
**Credential** | Pointer to [**NSXNetworkServerCredential**](NSXNetworkServerCredential.md) |  | [optional] 

## Methods

### NewNetworkPoolServerUpdateBluecat

`func NewNetworkPoolServerUpdateBluecat() *NetworkPoolServerUpdateBluecat`

NewNetworkPoolServerUpdateBluecat instantiates a new NetworkPoolServerUpdateBluecat object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkPoolServerUpdateBluecatWithDefaults

`func NewNetworkPoolServerUpdateBluecatWithDefaults() *NetworkPoolServerUpdateBluecat`

NewNetworkPoolServerUpdateBluecatWithDefaults instantiates a new NetworkPoolServerUpdateBluecat object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NetworkPoolServerUpdateBluecat) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkPoolServerUpdateBluecat) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkPoolServerUpdateBluecat) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkPoolServerUpdateBluecat) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *NetworkPoolServerUpdateBluecat) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworkPoolServerUpdateBluecat) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworkPoolServerUpdateBluecat) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NetworkPoolServerUpdateBluecat) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetServiceUrl

`func (o *NetworkPoolServerUpdateBluecat) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *NetworkPoolServerUpdateBluecat) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *NetworkPoolServerUpdateBluecat) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.

### HasServiceUrl

`func (o *NetworkPoolServerUpdateBluecat) HasServiceUrl() bool`

HasServiceUrl returns a boolean if a field has been set.

### GetServiceUsername

`func (o *NetworkPoolServerUpdateBluecat) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *NetworkPoolServerUpdateBluecat) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *NetworkPoolServerUpdateBluecat) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.

### HasServiceUsername

`func (o *NetworkPoolServerUpdateBluecat) HasServiceUsername() bool`

HasServiceUsername returns a boolean if a field has been set.

### GetServicePassword

`func (o *NetworkPoolServerUpdateBluecat) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *NetworkPoolServerUpdateBluecat) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *NetworkPoolServerUpdateBluecat) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *NetworkPoolServerUpdateBluecat) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### GetServiceThrottleRate

`func (o *NetworkPoolServerUpdateBluecat) GetServiceThrottleRate() int64`

GetServiceThrottleRate returns the ServiceThrottleRate field if non-nil, zero value otherwise.

### GetServiceThrottleRateOk

`func (o *NetworkPoolServerUpdateBluecat) GetServiceThrottleRateOk() (*int64, bool)`

GetServiceThrottleRateOk returns a tuple with the ServiceThrottleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceThrottleRate

`func (o *NetworkPoolServerUpdateBluecat) SetServiceThrottleRate(v int64)`

SetServiceThrottleRate sets ServiceThrottleRate field to given value.

### HasServiceThrottleRate

`func (o *NetworkPoolServerUpdateBluecat) HasServiceThrottleRate() bool`

HasServiceThrottleRate returns a boolean if a field has been set.

### GetIgnoreSsl

`func (o *NetworkPoolServerUpdateBluecat) GetIgnoreSsl() bool`

GetIgnoreSsl returns the IgnoreSsl field if non-nil, zero value otherwise.

### GetIgnoreSslOk

`func (o *NetworkPoolServerUpdateBluecat) GetIgnoreSslOk() (*bool, bool)`

GetIgnoreSslOk returns a tuple with the IgnoreSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSsl

`func (o *NetworkPoolServerUpdateBluecat) SetIgnoreSsl(v bool)`

SetIgnoreSsl sets IgnoreSsl field to given value.

### HasIgnoreSsl

`func (o *NetworkPoolServerUpdateBluecat) HasIgnoreSsl() bool`

HasIgnoreSsl returns a boolean if a field has been set.

### GetNetworkFilter

`func (o *NetworkPoolServerUpdateBluecat) GetNetworkFilter() string`

GetNetworkFilter returns the NetworkFilter field if non-nil, zero value otherwise.

### GetNetworkFilterOk

`func (o *NetworkPoolServerUpdateBluecat) GetNetworkFilterOk() (*string, bool)`

GetNetworkFilterOk returns a tuple with the NetworkFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFilter

`func (o *NetworkPoolServerUpdateBluecat) SetNetworkFilter(v string)`

SetNetworkFilter sets NetworkFilter field to given value.

### HasNetworkFilter

`func (o *NetworkPoolServerUpdateBluecat) HasNetworkFilter() bool`

HasNetworkFilter returns a boolean if a field has been set.

### GetConfig

`func (o *NetworkPoolServerUpdateBluecat) GetConfig() BluecatNetworkPoolServerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *NetworkPoolServerUpdateBluecat) GetConfigOk() (*BluecatNetworkPoolServerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *NetworkPoolServerUpdateBluecat) SetConfig(v BluecatNetworkPoolServerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *NetworkPoolServerUpdateBluecat) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCredential

`func (o *NetworkPoolServerUpdateBluecat) GetCredential() NSXNetworkServerCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *NetworkPoolServerUpdateBluecat) GetCredentialOk() (*NSXNetworkServerCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *NetworkPoolServerUpdateBluecat) SetCredential(v NSXNetworkServerCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *NetworkPoolServerUpdateBluecat) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


