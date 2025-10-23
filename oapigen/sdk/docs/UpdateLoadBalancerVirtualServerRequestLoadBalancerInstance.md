# UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VipName** | Pointer to **string** | VIP Name | [optional] 
**Description** | Pointer to **string** | Description | [optional] 
**VipAddress** | Pointer to **string** | VIP Address | [optional] 
**VipPort** | Pointer to **int64** | VIP Port | [optional] 
**VipProtocol** | Pointer to **string** | VIP Protocol | [optional] 
**VipHostname** | Pointer to **string** | VIP Hostname | [optional] 
**Pool** | Pointer to **int64** |  | [optional] 
**SslCert** | Pointer to **int64** | SSL Client Certificate ID | [optional] 
**SslServerCert** | Pointer to **int64** | SSL Server Certificate ID | [optional] 
**Config** | Pointer to [**CreateLoadBalancerVirtualServerRequestLoadBalancerInstanceConfig**](CreateLoadBalancerVirtualServerRequestLoadBalancerInstanceConfig.md) |  | [optional] 

## Methods

### NewUpdateLoadBalancerVirtualServerRequestLoadBalancerInstance

`func NewUpdateLoadBalancerVirtualServerRequestLoadBalancerInstance() *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance`

NewUpdateLoadBalancerVirtualServerRequestLoadBalancerInstance instantiates a new UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLoadBalancerVirtualServerRequestLoadBalancerInstanceWithDefaults

`func NewUpdateLoadBalancerVirtualServerRequestLoadBalancerInstanceWithDefaults() *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance`

NewUpdateLoadBalancerVirtualServerRequestLoadBalancerInstanceWithDefaults instantiates a new UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVipName

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) GetVipName() string`

GetVipName returns the VipName field if non-nil, zero value otherwise.

### GetVipNameOk

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) GetVipNameOk() (*string, bool)`

GetVipNameOk returns a tuple with the VipName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipName

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) SetVipName(v string)`

SetVipName sets VipName field to given value.

### HasVipName

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) HasVipName() bool`

HasVipName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVipAddress

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) GetVipAddress() string`

GetVipAddress returns the VipAddress field if non-nil, zero value otherwise.

### GetVipAddressOk

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) GetVipAddressOk() (*string, bool)`

GetVipAddressOk returns a tuple with the VipAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipAddress

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) SetVipAddress(v string)`

SetVipAddress sets VipAddress field to given value.

### HasVipAddress

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) HasVipAddress() bool`

HasVipAddress returns a boolean if a field has been set.

### GetVipPort

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) GetVipPort() int64`

GetVipPort returns the VipPort field if non-nil, zero value otherwise.

### GetVipPortOk

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) GetVipPortOk() (*int64, bool)`

GetVipPortOk returns a tuple with the VipPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipPort

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) SetVipPort(v int64)`

SetVipPort sets VipPort field to given value.

### HasVipPort

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) HasVipPort() bool`

HasVipPort returns a boolean if a field has been set.

### GetVipProtocol

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) GetVipProtocol() string`

GetVipProtocol returns the VipProtocol field if non-nil, zero value otherwise.

### GetVipProtocolOk

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) GetVipProtocolOk() (*string, bool)`

GetVipProtocolOk returns a tuple with the VipProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipProtocol

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) SetVipProtocol(v string)`

SetVipProtocol sets VipProtocol field to given value.

### HasVipProtocol

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) HasVipProtocol() bool`

HasVipProtocol returns a boolean if a field has been set.

### GetVipHostname

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) GetVipHostname() string`

GetVipHostname returns the VipHostname field if non-nil, zero value otherwise.

### GetVipHostnameOk

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) GetVipHostnameOk() (*string, bool)`

GetVipHostnameOk returns a tuple with the VipHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipHostname

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) SetVipHostname(v string)`

SetVipHostname sets VipHostname field to given value.

### HasVipHostname

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) HasVipHostname() bool`

HasVipHostname returns a boolean if a field has been set.

### GetPool

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) GetPool() int64`

GetPool returns the Pool field if non-nil, zero value otherwise.

### GetPoolOk

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) GetPoolOk() (*int64, bool)`

GetPoolOk returns a tuple with the Pool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPool

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) SetPool(v int64)`

SetPool sets Pool field to given value.

### HasPool

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) HasPool() bool`

HasPool returns a boolean if a field has been set.

### GetSslCert

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) GetSslCert() int64`

GetSslCert returns the SslCert field if non-nil, zero value otherwise.

### GetSslCertOk

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) GetSslCertOk() (*int64, bool)`

GetSslCertOk returns a tuple with the SslCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCert

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) SetSslCert(v int64)`

SetSslCert sets SslCert field to given value.

### HasSslCert

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) HasSslCert() bool`

HasSslCert returns a boolean if a field has been set.

### GetSslServerCert

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) GetSslServerCert() int64`

GetSslServerCert returns the SslServerCert field if non-nil, zero value otherwise.

### GetSslServerCertOk

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) GetSslServerCertOk() (*int64, bool)`

GetSslServerCertOk returns a tuple with the SslServerCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslServerCert

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) SetSslServerCert(v int64)`

SetSslServerCert sets SslServerCert field to given value.

### HasSslServerCert

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) HasSslServerCert() bool`

HasSslServerCert returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) GetConfig() CreateLoadBalancerVirtualServerRequestLoadBalancerInstanceConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) GetConfigOk() (*CreateLoadBalancerVirtualServerRequestLoadBalancerInstanceConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) SetConfig(v CreateLoadBalancerVirtualServerRequestLoadBalancerInstanceConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateLoadBalancerVirtualServerRequestLoadBalancerInstance) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


