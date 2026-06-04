# CreateLoadBalancerPoolRequestLoadBalancerPoolConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveMonitorPaths** | Pointer to **int64** | The ID of the active health monitor (NetworkLoadBalancerMonitor). The Options API &#x60;/api/options/nsxt/nsxtLBPoolActiveMonitor?loadBalancerId&#x3D;{id}&#x60; can be used to see which options are available. | [optional] 
**PassiveMonitorPath** | Pointer to **int64** | The ID of the passive health monitor (NetworkLoadBalancerMonitor). The Options API &#x60;/api/options/nsxt/nsxtLBPoolPassiveMonitor?loadBalancerId&#x3D;{id}&#x60; can be used to see which options are available. | [optional] 
**SnatTranslationType** | Pointer to **string** | SNAT translation type. Determines how source NAT is applied to pool traffic. | [optional] 
**SnatIpAddresses** | Pointer to **[]string** | List of SNAT IP addresses. Required when snatTranslationType is LBSnatIpPool. | [optional] 
**TcpMultiplexing** | Pointer to **bool** | Whether TCP multiplexing is enabled for the pool. | [optional] 
**TcpMultiplexingNumber** | Pointer to **int64** | Maximum number of TCP multiplexing connections. Defaults to 6. | [optional] 
**MemberGroup** | Pointer to [**NSXTLoadBalancerPoolConfigObjectMemberGroup**](NSXTLoadBalancerPoolConfigObjectMemberGroup.md) |  | [optional] 

## Methods

### NewCreateLoadBalancerPoolRequestLoadBalancerPoolConfig

`func NewCreateLoadBalancerPoolRequestLoadBalancerPoolConfig() *CreateLoadBalancerPoolRequestLoadBalancerPoolConfig`

NewCreateLoadBalancerPoolRequestLoadBalancerPoolConfig instantiates a new CreateLoadBalancerPoolRequestLoadBalancerPoolConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateLoadBalancerPoolRequestLoadBalancerPoolConfigWithDefaults

`func NewCreateLoadBalancerPoolRequestLoadBalancerPoolConfigWithDefaults() *CreateLoadBalancerPoolRequestLoadBalancerPoolConfig`

NewCreateLoadBalancerPoolRequestLoadBalancerPoolConfigWithDefaults instantiates a new CreateLoadBalancerPoolRequestLoadBalancerPoolConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveMonitorPaths

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPoolConfig) GetActiveMonitorPaths() int64`

GetActiveMonitorPaths returns the ActiveMonitorPaths field if non-nil, zero value otherwise.

### GetActiveMonitorPathsOk

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPoolConfig) GetActiveMonitorPathsOk() (*int64, bool)`

GetActiveMonitorPathsOk returns a tuple with the ActiveMonitorPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveMonitorPaths

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPoolConfig) SetActiveMonitorPaths(v int64)`

SetActiveMonitorPaths sets ActiveMonitorPaths field to given value.

### HasActiveMonitorPaths

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPoolConfig) HasActiveMonitorPaths() bool`

HasActiveMonitorPaths returns a boolean if a field has been set.

### GetPassiveMonitorPath

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPoolConfig) GetPassiveMonitorPath() int64`

GetPassiveMonitorPath returns the PassiveMonitorPath field if non-nil, zero value otherwise.

### GetPassiveMonitorPathOk

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPoolConfig) GetPassiveMonitorPathOk() (*int64, bool)`

GetPassiveMonitorPathOk returns a tuple with the PassiveMonitorPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassiveMonitorPath

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPoolConfig) SetPassiveMonitorPath(v int64)`

SetPassiveMonitorPath sets PassiveMonitorPath field to given value.

### HasPassiveMonitorPath

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPoolConfig) HasPassiveMonitorPath() bool`

HasPassiveMonitorPath returns a boolean if a field has been set.

### GetSnatTranslationType

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPoolConfig) GetSnatTranslationType() string`

GetSnatTranslationType returns the SnatTranslationType field if non-nil, zero value otherwise.

### GetSnatTranslationTypeOk

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPoolConfig) GetSnatTranslationTypeOk() (*string, bool)`

GetSnatTranslationTypeOk returns a tuple with the SnatTranslationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnatTranslationType

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPoolConfig) SetSnatTranslationType(v string)`

SetSnatTranslationType sets SnatTranslationType field to given value.

### HasSnatTranslationType

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPoolConfig) HasSnatTranslationType() bool`

HasSnatTranslationType returns a boolean if a field has been set.

### GetSnatIpAddresses

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPoolConfig) GetSnatIpAddresses() []string`

GetSnatIpAddresses returns the SnatIpAddresses field if non-nil, zero value otherwise.

### GetSnatIpAddressesOk

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPoolConfig) GetSnatIpAddressesOk() (*[]string, bool)`

GetSnatIpAddressesOk returns a tuple with the SnatIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnatIpAddresses

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPoolConfig) SetSnatIpAddresses(v []string)`

SetSnatIpAddresses sets SnatIpAddresses field to given value.

### HasSnatIpAddresses

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPoolConfig) HasSnatIpAddresses() bool`

HasSnatIpAddresses returns a boolean if a field has been set.

### GetTcpMultiplexing

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPoolConfig) GetTcpMultiplexing() bool`

GetTcpMultiplexing returns the TcpMultiplexing field if non-nil, zero value otherwise.

### GetTcpMultiplexingOk

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPoolConfig) GetTcpMultiplexingOk() (*bool, bool)`

GetTcpMultiplexingOk returns a tuple with the TcpMultiplexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMultiplexing

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPoolConfig) SetTcpMultiplexing(v bool)`

SetTcpMultiplexing sets TcpMultiplexing field to given value.

### HasTcpMultiplexing

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPoolConfig) HasTcpMultiplexing() bool`

HasTcpMultiplexing returns a boolean if a field has been set.

### GetTcpMultiplexingNumber

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPoolConfig) GetTcpMultiplexingNumber() int64`

GetTcpMultiplexingNumber returns the TcpMultiplexingNumber field if non-nil, zero value otherwise.

### GetTcpMultiplexingNumberOk

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPoolConfig) GetTcpMultiplexingNumberOk() (*int64, bool)`

GetTcpMultiplexingNumberOk returns a tuple with the TcpMultiplexingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMultiplexingNumber

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPoolConfig) SetTcpMultiplexingNumber(v int64)`

SetTcpMultiplexingNumber sets TcpMultiplexingNumber field to given value.

### HasTcpMultiplexingNumber

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPoolConfig) HasTcpMultiplexingNumber() bool`

HasTcpMultiplexingNumber returns a boolean if a field has been set.

### GetMemberGroup

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPoolConfig) GetMemberGroup() NSXTLoadBalancerPoolConfigObjectMemberGroup`

GetMemberGroup returns the MemberGroup field if non-nil, zero value otherwise.

### GetMemberGroupOk

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPoolConfig) GetMemberGroupOk() (*NSXTLoadBalancerPoolConfigObjectMemberGroup, bool)`

GetMemberGroupOk returns a tuple with the MemberGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberGroup

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPoolConfig) SetMemberGroup(v NSXTLoadBalancerPoolConfigObjectMemberGroup)`

SetMemberGroup sets MemberGroup field to given value.

### HasMemberGroup

`func (o *CreateLoadBalancerPoolRequestLoadBalancerPoolConfig) HasMemberGroup() bool`

HasMemberGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


