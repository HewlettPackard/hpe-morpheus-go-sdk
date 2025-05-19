# PhpIPAMNetworkPoolServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type Code (phpIPAM) | 
**Name** | **string** | Name | 
**Enabled** | Pointer to **bool** | Can be used to enable / disable the network pool server. | [optional] [default to true]
**ServiceUrl** | **string** | URL | 
**ServiceUsername** | Pointer to **string** | Username | [optional] 
**ServicePassword** | Pointer to **string** | Password | [optional] 
**ServiceThrottleRate** | Pointer to **int64** | Throttle Rate | [optional] [default to 0]
**IgnoreSsl** | Pointer to **bool** | Disable SSL SNI Verification | [optional] [default to true]
**NetworkFilter** | Pointer to **string** | Network Filter | [optional] 
**Config** | [**PhpIPAMNetworkPoolServerConfig**](PhpIPAMNetworkPoolServerConfig.md) |  | 
**Credential** | Pointer to [**NSXNetworkServerCredential**](NSXNetworkServerCredential.md) |  | [optional] 

## Methods

### NewPhpIPAMNetworkPoolServer

`func NewPhpIPAMNetworkPoolServer(type_ string, name string, serviceUrl string, config PhpIPAMNetworkPoolServerConfig, ) *PhpIPAMNetworkPoolServer`

NewPhpIPAMNetworkPoolServer instantiates a new PhpIPAMNetworkPoolServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhpIPAMNetworkPoolServerWithDefaults

`func NewPhpIPAMNetworkPoolServerWithDefaults() *PhpIPAMNetworkPoolServer`

NewPhpIPAMNetworkPoolServerWithDefaults instantiates a new PhpIPAMNetworkPoolServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *PhpIPAMNetworkPoolServer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PhpIPAMNetworkPoolServer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PhpIPAMNetworkPoolServer) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *PhpIPAMNetworkPoolServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PhpIPAMNetworkPoolServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PhpIPAMNetworkPoolServer) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *PhpIPAMNetworkPoolServer) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PhpIPAMNetworkPoolServer) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PhpIPAMNetworkPoolServer) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *PhpIPAMNetworkPoolServer) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetServiceUrl

`func (o *PhpIPAMNetworkPoolServer) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *PhpIPAMNetworkPoolServer) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *PhpIPAMNetworkPoolServer) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.


### GetServiceUsername

`func (o *PhpIPAMNetworkPoolServer) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *PhpIPAMNetworkPoolServer) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *PhpIPAMNetworkPoolServer) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.

### HasServiceUsername

`func (o *PhpIPAMNetworkPoolServer) HasServiceUsername() bool`

HasServiceUsername returns a boolean if a field has been set.

### GetServicePassword

`func (o *PhpIPAMNetworkPoolServer) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *PhpIPAMNetworkPoolServer) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *PhpIPAMNetworkPoolServer) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *PhpIPAMNetworkPoolServer) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### GetServiceThrottleRate

`func (o *PhpIPAMNetworkPoolServer) GetServiceThrottleRate() int64`

GetServiceThrottleRate returns the ServiceThrottleRate field if non-nil, zero value otherwise.

### GetServiceThrottleRateOk

`func (o *PhpIPAMNetworkPoolServer) GetServiceThrottleRateOk() (*int64, bool)`

GetServiceThrottleRateOk returns a tuple with the ServiceThrottleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceThrottleRate

`func (o *PhpIPAMNetworkPoolServer) SetServiceThrottleRate(v int64)`

SetServiceThrottleRate sets ServiceThrottleRate field to given value.

### HasServiceThrottleRate

`func (o *PhpIPAMNetworkPoolServer) HasServiceThrottleRate() bool`

HasServiceThrottleRate returns a boolean if a field has been set.

### GetIgnoreSsl

`func (o *PhpIPAMNetworkPoolServer) GetIgnoreSsl() bool`

GetIgnoreSsl returns the IgnoreSsl field if non-nil, zero value otherwise.

### GetIgnoreSslOk

`func (o *PhpIPAMNetworkPoolServer) GetIgnoreSslOk() (*bool, bool)`

GetIgnoreSslOk returns a tuple with the IgnoreSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSsl

`func (o *PhpIPAMNetworkPoolServer) SetIgnoreSsl(v bool)`

SetIgnoreSsl sets IgnoreSsl field to given value.

### HasIgnoreSsl

`func (o *PhpIPAMNetworkPoolServer) HasIgnoreSsl() bool`

HasIgnoreSsl returns a boolean if a field has been set.

### GetNetworkFilter

`func (o *PhpIPAMNetworkPoolServer) GetNetworkFilter() string`

GetNetworkFilter returns the NetworkFilter field if non-nil, zero value otherwise.

### GetNetworkFilterOk

`func (o *PhpIPAMNetworkPoolServer) GetNetworkFilterOk() (*string, bool)`

GetNetworkFilterOk returns a tuple with the NetworkFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFilter

`func (o *PhpIPAMNetworkPoolServer) SetNetworkFilter(v string)`

SetNetworkFilter sets NetworkFilter field to given value.

### HasNetworkFilter

`func (o *PhpIPAMNetworkPoolServer) HasNetworkFilter() bool`

HasNetworkFilter returns a boolean if a field has been set.

### GetConfig

`func (o *PhpIPAMNetworkPoolServer) GetConfig() PhpIPAMNetworkPoolServerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *PhpIPAMNetworkPoolServer) GetConfigOk() (*PhpIPAMNetworkPoolServerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *PhpIPAMNetworkPoolServer) SetConfig(v PhpIPAMNetworkPoolServerConfig)`

SetConfig sets Config field to given value.


### GetCredential

`func (o *PhpIPAMNetworkPoolServer) GetCredential() NSXNetworkServerCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *PhpIPAMNetworkPoolServer) GetCredentialOk() (*NSXNetworkServerCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *PhpIPAMNetworkPoolServer) SetCredential(v NSXNetworkServerCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *PhpIPAMNetworkPoolServer) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


