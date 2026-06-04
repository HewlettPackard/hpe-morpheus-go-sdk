# UpdateLoadBalancerPoolRequestLoadBalancerPoolConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveMonitorPaths** | Pointer to **int64** | The ID of the active health monitor (NetworkLoadBalancerMonitor). The Options API &#x60;/api/options/nsxt/nsxtLBPoolActiveMonitor?loadBalancerId&#x3D;{id}&#x60; can be used to see which options are available. | [optional] 
**PassiveMonitorPath** | Pointer to **int64** | The ID of the passive health monitor (NetworkLoadBalancerMonitor). The Options API &#x60;/api/options/nsxt/nsxtLBPoolPassiveMonitor?loadBalancerId&#x3D;{id}&#x60; can be used to see which options are available. | [optional] 
**SnatTranslationType** | Pointer to **string** | SNAT translation type. Determines how source NAT is applied to pool traffic. | [optional] 
**SnatIpAddresses** | Pointer to **[]string** | List of SNAT IP addresses. Required when snatTranslationType is LBSnatIpPool. | [optional] 
**TcpMultiplexing** | Pointer to **bool** | Whether TCP multiplexing is enabled for the pool. | [optional] 
**TcpMultiplexingNumber** | Pointer to **int64** | Maximum number of TCP multiplexing connections. Defaults to 6. | [optional] 
**MemberGroup** | Pointer to [**NSXTLoadBalancerPoolConfigObject1MemberGroup**](NSXTLoadBalancerPoolConfigObject1MemberGroup.md) |  | [optional] 

## Methods

### NewUpdateLoadBalancerPoolRequestLoadBalancerPoolConfig

`func NewUpdateLoadBalancerPoolRequestLoadBalancerPoolConfig() *UpdateLoadBalancerPoolRequestLoadBalancerPoolConfig`

NewUpdateLoadBalancerPoolRequestLoadBalancerPoolConfig instantiates a new UpdateLoadBalancerPoolRequestLoadBalancerPoolConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLoadBalancerPoolRequestLoadBalancerPoolConfigWithDefaults

`func NewUpdateLoadBalancerPoolRequestLoadBalancerPoolConfigWithDefaults() *UpdateLoadBalancerPoolRequestLoadBalancerPoolConfig`

NewUpdateLoadBalancerPoolRequestLoadBalancerPoolConfigWithDefaults instantiates a new UpdateLoadBalancerPoolRequestLoadBalancerPoolConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveMonitorPaths

`func (o *UpdateLoadBalancerPoolRequestLoadBalancerPoolConfig) GetActiveMonitorPaths() int64`

GetActiveMonitorPaths returns the ActiveMonitorPaths field if non-nil, zero value otherwise.

### GetActiveMonitorPathsOk

`func (o *UpdateLoadBalancerPoolRequestLoadBalancerPoolConfig) GetActiveMonitorPathsOk() (*int64, bool)`

GetActiveMonitorPathsOk returns a tuple with the ActiveMonitorPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveMonitorPaths

`func (o *UpdateLoadBalancerPoolRequestLoadBalancerPoolConfig) SetActiveMonitorPaths(v int64)`

SetActiveMonitorPaths sets ActiveMonitorPaths field to given value.

### HasActiveMonitorPaths

`func (o *UpdateLoadBalancerPoolRequestLoadBalancerPoolConfig) HasActiveMonitorPaths() bool`

HasActiveMonitorPaths returns a boolean if a field has been set.

### GetPassiveMonitorPath

`func (o *UpdateLoadBalancerPoolRequestLoadBalancerPoolConfig) GetPassiveMonitorPath() int64`

GetPassiveMonitorPath returns the PassiveMonitorPath field if non-nil, zero value otherwise.

### GetPassiveMonitorPathOk

`func (o *UpdateLoadBalancerPoolRequestLoadBalancerPoolConfig) GetPassiveMonitorPathOk() (*int64, bool)`

GetPassiveMonitorPathOk returns a tuple with the PassiveMonitorPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassiveMonitorPath

`func (o *UpdateLoadBalancerPoolRequestLoadBalancerPoolConfig) SetPassiveMonitorPath(v int64)`

SetPassiveMonitorPath sets PassiveMonitorPath field to given value.

### HasPassiveMonitorPath

`func (o *UpdateLoadBalancerPoolRequestLoadBalancerPoolConfig) HasPassiveMonitorPath() bool`

HasPassiveMonitorPath returns a boolean if a field has been set.

### GetSnatTranslationType

`func (o *UpdateLoadBalancerPoolRequestLoadBalancerPoolConfig) GetSnatTranslationType() string`

GetSnatTranslationType returns the SnatTranslationType field if non-nil, zero value otherwise.

### GetSnatTranslationTypeOk

`func (o *UpdateLoadBalancerPoolRequestLoadBalancerPoolConfig) GetSnatTranslationTypeOk() (*string, bool)`

GetSnatTranslationTypeOk returns a tuple with the SnatTranslationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnatTranslationType

`func (o *UpdateLoadBalancerPoolRequestLoadBalancerPoolConfig) SetSnatTranslationType(v string)`

SetSnatTranslationType sets SnatTranslationType field to given value.

### HasSnatTranslationType

`func (o *UpdateLoadBalancerPoolRequestLoadBalancerPoolConfig) HasSnatTranslationType() bool`

HasSnatTranslationType returns a boolean if a field has been set.

### GetSnatIpAddresses

`func (o *UpdateLoadBalancerPoolRequestLoadBalancerPoolConfig) GetSnatIpAddresses() []string`

GetSnatIpAddresses returns the SnatIpAddresses field if non-nil, zero value otherwise.

### GetSnatIpAddressesOk

`func (o *UpdateLoadBalancerPoolRequestLoadBalancerPoolConfig) GetSnatIpAddressesOk() (*[]string, bool)`

GetSnatIpAddressesOk returns a tuple with the SnatIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnatIpAddresses

`func (o *UpdateLoadBalancerPoolRequestLoadBalancerPoolConfig) SetSnatIpAddresses(v []string)`

SetSnatIpAddresses sets SnatIpAddresses field to given value.

### HasSnatIpAddresses

`func (o *UpdateLoadBalancerPoolRequestLoadBalancerPoolConfig) HasSnatIpAddresses() bool`

HasSnatIpAddresses returns a boolean if a field has been set.

### GetTcpMultiplexing

`func (o *UpdateLoadBalancerPoolRequestLoadBalancerPoolConfig) GetTcpMultiplexing() bool`

GetTcpMultiplexing returns the TcpMultiplexing field if non-nil, zero value otherwise.

### GetTcpMultiplexingOk

`func (o *UpdateLoadBalancerPoolRequestLoadBalancerPoolConfig) GetTcpMultiplexingOk() (*bool, bool)`

GetTcpMultiplexingOk returns a tuple with the TcpMultiplexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMultiplexing

`func (o *UpdateLoadBalancerPoolRequestLoadBalancerPoolConfig) SetTcpMultiplexing(v bool)`

SetTcpMultiplexing sets TcpMultiplexing field to given value.

### HasTcpMultiplexing

`func (o *UpdateLoadBalancerPoolRequestLoadBalancerPoolConfig) HasTcpMultiplexing() bool`

HasTcpMultiplexing returns a boolean if a field has been set.

### GetTcpMultiplexingNumber

`func (o *UpdateLoadBalancerPoolRequestLoadBalancerPoolConfig) GetTcpMultiplexingNumber() int64`

GetTcpMultiplexingNumber returns the TcpMultiplexingNumber field if non-nil, zero value otherwise.

### GetTcpMultiplexingNumberOk

`func (o *UpdateLoadBalancerPoolRequestLoadBalancerPoolConfig) GetTcpMultiplexingNumberOk() (*int64, bool)`

GetTcpMultiplexingNumberOk returns a tuple with the TcpMultiplexingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMultiplexingNumber

`func (o *UpdateLoadBalancerPoolRequestLoadBalancerPoolConfig) SetTcpMultiplexingNumber(v int64)`

SetTcpMultiplexingNumber sets TcpMultiplexingNumber field to given value.

### HasTcpMultiplexingNumber

`func (o *UpdateLoadBalancerPoolRequestLoadBalancerPoolConfig) HasTcpMultiplexingNumber() bool`

HasTcpMultiplexingNumber returns a boolean if a field has been set.

### GetMemberGroup

`func (o *UpdateLoadBalancerPoolRequestLoadBalancerPoolConfig) GetMemberGroup() NSXTLoadBalancerPoolConfigObject1MemberGroup`

GetMemberGroup returns the MemberGroup field if non-nil, zero value otherwise.

### GetMemberGroupOk

`func (o *UpdateLoadBalancerPoolRequestLoadBalancerPoolConfig) GetMemberGroupOk() (*NSXTLoadBalancerPoolConfigObject1MemberGroup, bool)`

GetMemberGroupOk returns a tuple with the MemberGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberGroup

`func (o *UpdateLoadBalancerPoolRequestLoadBalancerPoolConfig) SetMemberGroup(v NSXTLoadBalancerPoolConfigObject1MemberGroup)`

SetMemberGroup sets MemberGroup field to given value.

### HasMemberGroup

`func (o *UpdateLoadBalancerPoolRequestLoadBalancerPoolConfig) HasMemberGroup() bool`

HasMemberGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


