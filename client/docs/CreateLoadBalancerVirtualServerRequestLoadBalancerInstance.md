# CreateLoadBalancerVirtualServerRequestLoadBalancerInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VipName** | Pointer to **string** | VIP Name | [optional] 
**Description** | Pointer to **string** | Description | [optional] 
**VipAddress** | Pointer to **string** | VIP Address | [optional] 
**VipPort** | Pointer to **int64** | VIP Port | [optional] 
**VipProtocol** | Pointer to **string** | VIP Protocol | [optional] 
**VipHostname** | Pointer to **string** | VIP Hostname | [optional] 
**SslCert** | Pointer to **int64** | SSL Client Certificate ID | [optional] 
**SslServerCert** | Pointer to **int64** | SSL Server Certificate ID | [optional] 
**Config** | Pointer to [**CreateLoadBalancerVirtualServerRequestLoadBalancerInstanceConfig**](CreateLoadBalancerVirtualServerRequestLoadBalancerInstanceConfig.md) |  | [optional] 

## Methods

### NewCreateLoadBalancerVirtualServerRequestLoadBalancerInstance

`func NewCreateLoadBalancerVirtualServerRequestLoadBalancerInstance() *CreateLoadBalancerVirtualServerRequestLoadBalancerInstance`

NewCreateLoadBalancerVirtualServerRequestLoadBalancerInstance instantiates a new CreateLoadBalancerVirtualServerRequestLoadBalancerInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLoadBalancerVirtualServerRequestLoadBalancerInstanceWithDefaults

`func NewCreateLoadBalancerVirtualServerRequestLoadBalancerInstanceWithDefaults() *CreateLoadBalancerVirtualServerRequestLoadBalancerInstance`

NewCreateLoadBalancerVirtualServerRequestLoadBalancerInstanceWithDefaults instantiates a new CreateLoadBalancerVirtualServerRequestLoadBalancerInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVipName

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstance) GetVipName() string`

GetVipName returns the VipName field if non-nil, zero value otherwise.

### GetVipNameOk

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstance) GetVipNameOk() (*string, bool)`

GetVipNameOk returns a tuple with the VipName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipName

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstance) SetVipName(v string)`

SetVipName sets VipName field to given value.

### HasVipName

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstance) HasVipName() bool`

HasVipName returns a boolean if a field has been set.

### GetDescription

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstance) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstance) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstance) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstance) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVipAddress

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstance) GetVipAddress() string`

GetVipAddress returns the VipAddress field if non-nil, zero value otherwise.

### GetVipAddressOk

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstance) GetVipAddressOk() (*string, bool)`

GetVipAddressOk returns a tuple with the VipAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipAddress

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstance) SetVipAddress(v string)`

SetVipAddress sets VipAddress field to given value.

### HasVipAddress

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstance) HasVipAddress() bool`

HasVipAddress returns a boolean if a field has been set.

### GetVipPort

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstance) GetVipPort() int64`

GetVipPort returns the VipPort field if non-nil, zero value otherwise.

### GetVipPortOk

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstance) GetVipPortOk() (*int64, bool)`

GetVipPortOk returns a tuple with the VipPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipPort

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstance) SetVipPort(v int64)`

SetVipPort sets VipPort field to given value.

### HasVipPort

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstance) HasVipPort() bool`

HasVipPort returns a boolean if a field has been set.

### GetVipProtocol

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstance) GetVipProtocol() string`

GetVipProtocol returns the VipProtocol field if non-nil, zero value otherwise.

### GetVipProtocolOk

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstance) GetVipProtocolOk() (*string, bool)`

GetVipProtocolOk returns a tuple with the VipProtocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipProtocol

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstance) SetVipProtocol(v string)`

SetVipProtocol sets VipProtocol field to given value.

### HasVipProtocol

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstance) HasVipProtocol() bool`

HasVipProtocol returns a boolean if a field has been set.

### GetVipHostname

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstance) GetVipHostname() string`

GetVipHostname returns the VipHostname field if non-nil, zero value otherwise.

### GetVipHostnameOk

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstance) GetVipHostnameOk() (*string, bool)`

GetVipHostnameOk returns a tuple with the VipHostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipHostname

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstance) SetVipHostname(v string)`

SetVipHostname sets VipHostname field to given value.

### HasVipHostname

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstance) HasVipHostname() bool`

HasVipHostname returns a boolean if a field has been set.

### GetSslCert

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstance) GetSslCert() int64`

GetSslCert returns the SslCert field if non-nil, zero value otherwise.

### GetSslCertOk

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstance) GetSslCertOk() (*int64, bool)`

GetSslCertOk returns a tuple with the SslCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslCert

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstance) SetSslCert(v int64)`

SetSslCert sets SslCert field to given value.

### HasSslCert

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstance) HasSslCert() bool`

HasSslCert returns a boolean if a field has been set.

### GetSslServerCert

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstance) GetSslServerCert() int64`

GetSslServerCert returns the SslServerCert field if non-nil, zero value otherwise.

### GetSslServerCertOk

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstance) GetSslServerCertOk() (*int64, bool)`

GetSslServerCertOk returns a tuple with the SslServerCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSslServerCert

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstance) SetSslServerCert(v int64)`

SetSslServerCert sets SslServerCert field to given value.

### HasSslServerCert

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstance) HasSslServerCert() bool`

HasSslServerCert returns a boolean if a field has been set.

### GetConfig

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstance) GetConfig() CreateLoadBalancerVirtualServerRequestLoadBalancerInstanceConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstance) GetConfigOk() (*CreateLoadBalancerVirtualServerRequestLoadBalancerInstanceConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstance) SetConfig(v CreateLoadBalancerVirtualServerRequestLoadBalancerInstanceConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *CreateLoadBalancerVirtualServerRequestLoadBalancerInstance) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


