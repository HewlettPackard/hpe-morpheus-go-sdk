# NetworkPoolServerUpdatePhpIpam

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
**Config** | Pointer to [**PhpIPAMNetworkPoolServerUpdateConfig**](PhpIPAMNetworkPoolServerUpdateConfig.md) |  | [optional] 
**Credential** | Pointer to [**NSXNetworkServerCredential**](NSXNetworkServerCredential.md) |  | [optional] 

## Methods

### NewNetworkPoolServerUpdatePhpIpam

`func NewNetworkPoolServerUpdatePhpIpam() *NetworkPoolServerUpdatePhpIpam`

NewNetworkPoolServerUpdatePhpIpam instantiates a new NetworkPoolServerUpdatePhpIpam object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNetworkPoolServerUpdatePhpIpamWithDefaults

`func NewNetworkPoolServerUpdatePhpIpamWithDefaults() *NetworkPoolServerUpdatePhpIpam`

NewNetworkPoolServerUpdatePhpIpamWithDefaults instantiates a new NetworkPoolServerUpdatePhpIpam object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *NetworkPoolServerUpdatePhpIpam) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *NetworkPoolServerUpdatePhpIpam) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *NetworkPoolServerUpdatePhpIpam) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *NetworkPoolServerUpdatePhpIpam) HasName() bool`

HasName returns a boolean if a field has been set.

### GetEnabled

`func (o *NetworkPoolServerUpdatePhpIpam) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NetworkPoolServerUpdatePhpIpam) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NetworkPoolServerUpdatePhpIpam) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NetworkPoolServerUpdatePhpIpam) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetServiceUrl

`func (o *NetworkPoolServerUpdatePhpIpam) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *NetworkPoolServerUpdatePhpIpam) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *NetworkPoolServerUpdatePhpIpam) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.

### HasServiceUrl

`func (o *NetworkPoolServerUpdatePhpIpam) HasServiceUrl() bool`

HasServiceUrl returns a boolean if a field has been set.

### GetServiceUsername

`func (o *NetworkPoolServerUpdatePhpIpam) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *NetworkPoolServerUpdatePhpIpam) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *NetworkPoolServerUpdatePhpIpam) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.

### HasServiceUsername

`func (o *NetworkPoolServerUpdatePhpIpam) HasServiceUsername() bool`

HasServiceUsername returns a boolean if a field has been set.

### GetServicePassword

`func (o *NetworkPoolServerUpdatePhpIpam) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *NetworkPoolServerUpdatePhpIpam) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *NetworkPoolServerUpdatePhpIpam) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *NetworkPoolServerUpdatePhpIpam) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### GetServiceThrottleRate

`func (o *NetworkPoolServerUpdatePhpIpam) GetServiceThrottleRate() int64`

GetServiceThrottleRate returns the ServiceThrottleRate field if non-nil, zero value otherwise.

### GetServiceThrottleRateOk

`func (o *NetworkPoolServerUpdatePhpIpam) GetServiceThrottleRateOk() (*int64, bool)`

GetServiceThrottleRateOk returns a tuple with the ServiceThrottleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceThrottleRate

`func (o *NetworkPoolServerUpdatePhpIpam) SetServiceThrottleRate(v int64)`

SetServiceThrottleRate sets ServiceThrottleRate field to given value.

### HasServiceThrottleRate

`func (o *NetworkPoolServerUpdatePhpIpam) HasServiceThrottleRate() bool`

HasServiceThrottleRate returns a boolean if a field has been set.

### GetIgnoreSsl

`func (o *NetworkPoolServerUpdatePhpIpam) GetIgnoreSsl() bool`

GetIgnoreSsl returns the IgnoreSsl field if non-nil, zero value otherwise.

### GetIgnoreSslOk

`func (o *NetworkPoolServerUpdatePhpIpam) GetIgnoreSslOk() (*bool, bool)`

GetIgnoreSslOk returns a tuple with the IgnoreSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSsl

`func (o *NetworkPoolServerUpdatePhpIpam) SetIgnoreSsl(v bool)`

SetIgnoreSsl sets IgnoreSsl field to given value.

### HasIgnoreSsl

`func (o *NetworkPoolServerUpdatePhpIpam) HasIgnoreSsl() bool`

HasIgnoreSsl returns a boolean if a field has been set.

### GetNetworkFilter

`func (o *NetworkPoolServerUpdatePhpIpam) GetNetworkFilter() string`

GetNetworkFilter returns the NetworkFilter field if non-nil, zero value otherwise.

### GetNetworkFilterOk

`func (o *NetworkPoolServerUpdatePhpIpam) GetNetworkFilterOk() (*string, bool)`

GetNetworkFilterOk returns a tuple with the NetworkFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFilter

`func (o *NetworkPoolServerUpdatePhpIpam) SetNetworkFilter(v string)`

SetNetworkFilter sets NetworkFilter field to given value.

### HasNetworkFilter

`func (o *NetworkPoolServerUpdatePhpIpam) HasNetworkFilter() bool`

HasNetworkFilter returns a boolean if a field has been set.

### GetConfig

`func (o *NetworkPoolServerUpdatePhpIpam) GetConfig() PhpIPAMNetworkPoolServerUpdateConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *NetworkPoolServerUpdatePhpIpam) GetConfigOk() (*PhpIPAMNetworkPoolServerUpdateConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *NetworkPoolServerUpdatePhpIpam) SetConfig(v PhpIPAMNetworkPoolServerUpdateConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *NetworkPoolServerUpdatePhpIpam) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCredential

`func (o *NetworkPoolServerUpdatePhpIpam) GetCredential() NSXNetworkServerCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *NetworkPoolServerUpdatePhpIpam) GetCredentialOk() (*NSXNetworkServerCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *NetworkPoolServerUpdatePhpIpam) SetCredential(v NSXNetworkServerCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *NetworkPoolServerUpdatePhpIpam) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


