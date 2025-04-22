# InfobloxNetworkPoolServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | Type Code (Infoblox) | 
**Name** | **string** | Name | 
**Enabled** | Pointer to **bool** | Can be used to enable / disable the network pool server. | [optional] [default to true]
**ServiceUrl** | **string** | URL | 
**ServiceUsername** | Pointer to **string** | Username | [optional] 
**ServicePassword** | Pointer to **string** | Password | [optional] 
**ServiceThrottleRate** | Pointer to **int64** | Throttle Rate | [optional] [default to 0]
**IgnoreSsl** | Pointer to **bool** | Disable SSL SNI Verification | [optional] [default to true]
**NetworkFilter** | Pointer to **string** | Network Filter | [optional] 
**ZoneFilter** | Pointer to **string** | Zone Filter | [optional] 
**TenantMatch** | Pointer to **string** | Tenant Match | [optional] 
**ServiceMode** | Pointer to **string** | IP Mode | [optional] [default to "static"]
**Config** | Pointer to [**InfobloxNetworkPoolServerConfig**](InfobloxNetworkPoolServerConfig.md) |  | [optional] 
**Credential** | Pointer to [**NSXNetworkServerCredential**](NSXNetworkServerCredential.md) |  | [optional] 

## Methods

### NewInfobloxNetworkPoolServer

`func NewInfobloxNetworkPoolServer(type_ string, name string, serviceUrl string, ) *InfobloxNetworkPoolServer`

NewInfobloxNetworkPoolServer instantiates a new InfobloxNetworkPoolServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInfobloxNetworkPoolServerWithDefaults

`func NewInfobloxNetworkPoolServerWithDefaults() *InfobloxNetworkPoolServer`

NewInfobloxNetworkPoolServerWithDefaults instantiates a new InfobloxNetworkPoolServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InfobloxNetworkPoolServer) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InfobloxNetworkPoolServer) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InfobloxNetworkPoolServer) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *InfobloxNetworkPoolServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InfobloxNetworkPoolServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InfobloxNetworkPoolServer) SetName(v string)`

SetName sets Name field to given value.


### GetEnabled

`func (o *InfobloxNetworkPoolServer) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *InfobloxNetworkPoolServer) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *InfobloxNetworkPoolServer) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *InfobloxNetworkPoolServer) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetServiceUrl

`func (o *InfobloxNetworkPoolServer) GetServiceUrl() string`

GetServiceUrl returns the ServiceUrl field if non-nil, zero value otherwise.

### GetServiceUrlOk

`func (o *InfobloxNetworkPoolServer) GetServiceUrlOk() (*string, bool)`

GetServiceUrlOk returns a tuple with the ServiceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUrl

`func (o *InfobloxNetworkPoolServer) SetServiceUrl(v string)`

SetServiceUrl sets ServiceUrl field to given value.


### GetServiceUsername

`func (o *InfobloxNetworkPoolServer) GetServiceUsername() string`

GetServiceUsername returns the ServiceUsername field if non-nil, zero value otherwise.

### GetServiceUsernameOk

`func (o *InfobloxNetworkPoolServer) GetServiceUsernameOk() (*string, bool)`

GetServiceUsernameOk returns a tuple with the ServiceUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceUsername

`func (o *InfobloxNetworkPoolServer) SetServiceUsername(v string)`

SetServiceUsername sets ServiceUsername field to given value.

### HasServiceUsername

`func (o *InfobloxNetworkPoolServer) HasServiceUsername() bool`

HasServiceUsername returns a boolean if a field has been set.

### GetServicePassword

`func (o *InfobloxNetworkPoolServer) GetServicePassword() string`

GetServicePassword returns the ServicePassword field if non-nil, zero value otherwise.

### GetServicePasswordOk

`func (o *InfobloxNetworkPoolServer) GetServicePasswordOk() (*string, bool)`

GetServicePasswordOk returns a tuple with the ServicePassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePassword

`func (o *InfobloxNetworkPoolServer) SetServicePassword(v string)`

SetServicePassword sets ServicePassword field to given value.

### HasServicePassword

`func (o *InfobloxNetworkPoolServer) HasServicePassword() bool`

HasServicePassword returns a boolean if a field has been set.

### GetServiceThrottleRate

`func (o *InfobloxNetworkPoolServer) GetServiceThrottleRate() int64`

GetServiceThrottleRate returns the ServiceThrottleRate field if non-nil, zero value otherwise.

### GetServiceThrottleRateOk

`func (o *InfobloxNetworkPoolServer) GetServiceThrottleRateOk() (*int64, bool)`

GetServiceThrottleRateOk returns a tuple with the ServiceThrottleRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceThrottleRate

`func (o *InfobloxNetworkPoolServer) SetServiceThrottleRate(v int64)`

SetServiceThrottleRate sets ServiceThrottleRate field to given value.

### HasServiceThrottleRate

`func (o *InfobloxNetworkPoolServer) HasServiceThrottleRate() bool`

HasServiceThrottleRate returns a boolean if a field has been set.

### GetIgnoreSsl

`func (o *InfobloxNetworkPoolServer) GetIgnoreSsl() bool`

GetIgnoreSsl returns the IgnoreSsl field if non-nil, zero value otherwise.

### GetIgnoreSslOk

`func (o *InfobloxNetworkPoolServer) GetIgnoreSslOk() (*bool, bool)`

GetIgnoreSslOk returns a tuple with the IgnoreSsl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIgnoreSsl

`func (o *InfobloxNetworkPoolServer) SetIgnoreSsl(v bool)`

SetIgnoreSsl sets IgnoreSsl field to given value.

### HasIgnoreSsl

`func (o *InfobloxNetworkPoolServer) HasIgnoreSsl() bool`

HasIgnoreSsl returns a boolean if a field has been set.

### GetNetworkFilter

`func (o *InfobloxNetworkPoolServer) GetNetworkFilter() string`

GetNetworkFilter returns the NetworkFilter field if non-nil, zero value otherwise.

### GetNetworkFilterOk

`func (o *InfobloxNetworkPoolServer) GetNetworkFilterOk() (*string, bool)`

GetNetworkFilterOk returns a tuple with the NetworkFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkFilter

`func (o *InfobloxNetworkPoolServer) SetNetworkFilter(v string)`

SetNetworkFilter sets NetworkFilter field to given value.

### HasNetworkFilter

`func (o *InfobloxNetworkPoolServer) HasNetworkFilter() bool`

HasNetworkFilter returns a boolean if a field has been set.

### GetZoneFilter

`func (o *InfobloxNetworkPoolServer) GetZoneFilter() string`

GetZoneFilter returns the ZoneFilter field if non-nil, zero value otherwise.

### GetZoneFilterOk

`func (o *InfobloxNetworkPoolServer) GetZoneFilterOk() (*string, bool)`

GetZoneFilterOk returns a tuple with the ZoneFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneFilter

`func (o *InfobloxNetworkPoolServer) SetZoneFilter(v string)`

SetZoneFilter sets ZoneFilter field to given value.

### HasZoneFilter

`func (o *InfobloxNetworkPoolServer) HasZoneFilter() bool`

HasZoneFilter returns a boolean if a field has been set.

### GetTenantMatch

`func (o *InfobloxNetworkPoolServer) GetTenantMatch() string`

GetTenantMatch returns the TenantMatch field if non-nil, zero value otherwise.

### GetTenantMatchOk

`func (o *InfobloxNetworkPoolServer) GetTenantMatchOk() (*string, bool)`

GetTenantMatchOk returns a tuple with the TenantMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantMatch

`func (o *InfobloxNetworkPoolServer) SetTenantMatch(v string)`

SetTenantMatch sets TenantMatch field to given value.

### HasTenantMatch

`func (o *InfobloxNetworkPoolServer) HasTenantMatch() bool`

HasTenantMatch returns a boolean if a field has been set.

### GetServiceMode

`func (o *InfobloxNetworkPoolServer) GetServiceMode() string`

GetServiceMode returns the ServiceMode field if non-nil, zero value otherwise.

### GetServiceModeOk

`func (o *InfobloxNetworkPoolServer) GetServiceModeOk() (*string, bool)`

GetServiceModeOk returns a tuple with the ServiceMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceMode

`func (o *InfobloxNetworkPoolServer) SetServiceMode(v string)`

SetServiceMode sets ServiceMode field to given value.

### HasServiceMode

`func (o *InfobloxNetworkPoolServer) HasServiceMode() bool`

HasServiceMode returns a boolean if a field has been set.

### GetConfig

`func (o *InfobloxNetworkPoolServer) GetConfig() InfobloxNetworkPoolServerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *InfobloxNetworkPoolServer) GetConfigOk() (*InfobloxNetworkPoolServerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *InfobloxNetworkPoolServer) SetConfig(v InfobloxNetworkPoolServerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *InfobloxNetworkPoolServer) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetCredential

`func (o *InfobloxNetworkPoolServer) GetCredential() NSXNetworkServerCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *InfobloxNetworkPoolServer) GetCredentialOk() (*NSXNetworkServerCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *InfobloxNetworkPoolServer) SetCredential(v NSXNetworkServerCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *InfobloxNetworkPoolServer) HasCredential() bool`

HasCredential returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


