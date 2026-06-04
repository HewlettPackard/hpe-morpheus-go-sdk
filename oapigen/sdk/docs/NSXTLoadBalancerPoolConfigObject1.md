# NSXTLoadBalancerPoolConfigObject1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActiveMonitorPaths** | Pointer to **NullableInt64** | The ID of the active health monitor (NetworkLoadBalancerMonitor). The Options API &#x60;/api/options/nsxt/nsxtLBPoolActiveMonitor?loadBalancerId&#x3D;{id}&#x60; can be used to see which options are available. | [optional] 
**PassiveMonitorPath** | Pointer to **NullableInt64** | The ID of the passive health monitor (NetworkLoadBalancerMonitor). The Options API &#x60;/api/options/nsxt/nsxtLBPoolPassiveMonitor?loadBalancerId&#x3D;{id}&#x60; can be used to see which options are available. | [optional] 
**SnatTranslationType** | Pointer to **string** | SNAT translation type. Determines how source NAT is applied to pool traffic. | [optional] 
**SnatIpAddresses** | Pointer to **[]string** | List of SNAT IP addresses. Required when snatTranslationType is LBSnatIpPool. | [optional] 
**TcpMultiplexing** | Pointer to **bool** | Whether TCP multiplexing is enabled for the pool. | [optional] 
**TcpMultiplexingNumber** | Pointer to **NullableInt64** | Maximum number of TCP multiplexing connections. Defaults to 6. | [optional] 
**MemberGroup** | Pointer to [**NSXTLoadBalancerPoolConfigObject1MemberGroup**](NSXTLoadBalancerPoolConfigObject1MemberGroup.md) |  | [optional] 

## Methods

### NewNSXTLoadBalancerPoolConfigObject1

`func NewNSXTLoadBalancerPoolConfigObject1() *NSXTLoadBalancerPoolConfigObject1`

NewNSXTLoadBalancerPoolConfigObject1 instantiates a new NSXTLoadBalancerPoolConfigObject1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNSXTLoadBalancerPoolConfigObject1WithDefaults

`func NewNSXTLoadBalancerPoolConfigObject1WithDefaults() *NSXTLoadBalancerPoolConfigObject1`

NewNSXTLoadBalancerPoolConfigObject1WithDefaults instantiates a new NSXTLoadBalancerPoolConfigObject1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActiveMonitorPaths

`func (o *NSXTLoadBalancerPoolConfigObject1) GetActiveMonitorPaths() int64`

GetActiveMonitorPaths returns the ActiveMonitorPaths field if non-nil, zero value otherwise.

### GetActiveMonitorPathsOk

`func (o *NSXTLoadBalancerPoolConfigObject1) GetActiveMonitorPathsOk() (*int64, bool)`

GetActiveMonitorPathsOk returns a tuple with the ActiveMonitorPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActiveMonitorPaths

`func (o *NSXTLoadBalancerPoolConfigObject1) SetActiveMonitorPaths(v int64)`

SetActiveMonitorPaths sets ActiveMonitorPaths field to given value.

### HasActiveMonitorPaths

`func (o *NSXTLoadBalancerPoolConfigObject1) HasActiveMonitorPaths() bool`

HasActiveMonitorPaths returns a boolean if a field has been set.

### SetActiveMonitorPathsNil

`func (o *NSXTLoadBalancerPoolConfigObject1) SetActiveMonitorPathsNil(b bool)`

 SetActiveMonitorPathsNil sets the value for ActiveMonitorPaths to be an explicit nil

### UnsetActiveMonitorPaths
`func (o *NSXTLoadBalancerPoolConfigObject1) UnsetActiveMonitorPaths()`

UnsetActiveMonitorPaths ensures that no value is present for ActiveMonitorPaths, not even an explicit nil
### GetPassiveMonitorPath

`func (o *NSXTLoadBalancerPoolConfigObject1) GetPassiveMonitorPath() int64`

GetPassiveMonitorPath returns the PassiveMonitorPath field if non-nil, zero value otherwise.

### GetPassiveMonitorPathOk

`func (o *NSXTLoadBalancerPoolConfigObject1) GetPassiveMonitorPathOk() (*int64, bool)`

GetPassiveMonitorPathOk returns a tuple with the PassiveMonitorPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassiveMonitorPath

`func (o *NSXTLoadBalancerPoolConfigObject1) SetPassiveMonitorPath(v int64)`

SetPassiveMonitorPath sets PassiveMonitorPath field to given value.

### HasPassiveMonitorPath

`func (o *NSXTLoadBalancerPoolConfigObject1) HasPassiveMonitorPath() bool`

HasPassiveMonitorPath returns a boolean if a field has been set.

### SetPassiveMonitorPathNil

`func (o *NSXTLoadBalancerPoolConfigObject1) SetPassiveMonitorPathNil(b bool)`

 SetPassiveMonitorPathNil sets the value for PassiveMonitorPath to be an explicit nil

### UnsetPassiveMonitorPath
`func (o *NSXTLoadBalancerPoolConfigObject1) UnsetPassiveMonitorPath()`

UnsetPassiveMonitorPath ensures that no value is present for PassiveMonitorPath, not even an explicit nil
### GetSnatTranslationType

`func (o *NSXTLoadBalancerPoolConfigObject1) GetSnatTranslationType() string`

GetSnatTranslationType returns the SnatTranslationType field if non-nil, zero value otherwise.

### GetSnatTranslationTypeOk

`func (o *NSXTLoadBalancerPoolConfigObject1) GetSnatTranslationTypeOk() (*string, bool)`

GetSnatTranslationTypeOk returns a tuple with the SnatTranslationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnatTranslationType

`func (o *NSXTLoadBalancerPoolConfigObject1) SetSnatTranslationType(v string)`

SetSnatTranslationType sets SnatTranslationType field to given value.

### HasSnatTranslationType

`func (o *NSXTLoadBalancerPoolConfigObject1) HasSnatTranslationType() bool`

HasSnatTranslationType returns a boolean if a field has been set.

### GetSnatIpAddresses

`func (o *NSXTLoadBalancerPoolConfigObject1) GetSnatIpAddresses() []string`

GetSnatIpAddresses returns the SnatIpAddresses field if non-nil, zero value otherwise.

### GetSnatIpAddressesOk

`func (o *NSXTLoadBalancerPoolConfigObject1) GetSnatIpAddressesOk() (*[]string, bool)`

GetSnatIpAddressesOk returns a tuple with the SnatIpAddresses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnatIpAddresses

`func (o *NSXTLoadBalancerPoolConfigObject1) SetSnatIpAddresses(v []string)`

SetSnatIpAddresses sets SnatIpAddresses field to given value.

### HasSnatIpAddresses

`func (o *NSXTLoadBalancerPoolConfigObject1) HasSnatIpAddresses() bool`

HasSnatIpAddresses returns a boolean if a field has been set.

### GetTcpMultiplexing

`func (o *NSXTLoadBalancerPoolConfigObject1) GetTcpMultiplexing() bool`

GetTcpMultiplexing returns the TcpMultiplexing field if non-nil, zero value otherwise.

### GetTcpMultiplexingOk

`func (o *NSXTLoadBalancerPoolConfigObject1) GetTcpMultiplexingOk() (*bool, bool)`

GetTcpMultiplexingOk returns a tuple with the TcpMultiplexing field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMultiplexing

`func (o *NSXTLoadBalancerPoolConfigObject1) SetTcpMultiplexing(v bool)`

SetTcpMultiplexing sets TcpMultiplexing field to given value.

### HasTcpMultiplexing

`func (o *NSXTLoadBalancerPoolConfigObject1) HasTcpMultiplexing() bool`

HasTcpMultiplexing returns a boolean if a field has been set.

### GetTcpMultiplexingNumber

`func (o *NSXTLoadBalancerPoolConfigObject1) GetTcpMultiplexingNumber() int64`

GetTcpMultiplexingNumber returns the TcpMultiplexingNumber field if non-nil, zero value otherwise.

### GetTcpMultiplexingNumberOk

`func (o *NSXTLoadBalancerPoolConfigObject1) GetTcpMultiplexingNumberOk() (*int64, bool)`

GetTcpMultiplexingNumberOk returns a tuple with the TcpMultiplexingNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTcpMultiplexingNumber

`func (o *NSXTLoadBalancerPoolConfigObject1) SetTcpMultiplexingNumber(v int64)`

SetTcpMultiplexingNumber sets TcpMultiplexingNumber field to given value.

### HasTcpMultiplexingNumber

`func (o *NSXTLoadBalancerPoolConfigObject1) HasTcpMultiplexingNumber() bool`

HasTcpMultiplexingNumber returns a boolean if a field has been set.

### SetTcpMultiplexingNumberNil

`func (o *NSXTLoadBalancerPoolConfigObject1) SetTcpMultiplexingNumberNil(b bool)`

 SetTcpMultiplexingNumberNil sets the value for TcpMultiplexingNumber to be an explicit nil

### UnsetTcpMultiplexingNumber
`func (o *NSXTLoadBalancerPoolConfigObject1) UnsetTcpMultiplexingNumber()`

UnsetTcpMultiplexingNumber ensures that no value is present for TcpMultiplexingNumber, not even an explicit nil
### GetMemberGroup

`func (o *NSXTLoadBalancerPoolConfigObject1) GetMemberGroup() NSXTLoadBalancerPoolConfigObject1MemberGroup`

GetMemberGroup returns the MemberGroup field if non-nil, zero value otherwise.

### GetMemberGroupOk

`func (o *NSXTLoadBalancerPoolConfigObject1) GetMemberGroupOk() (*NSXTLoadBalancerPoolConfigObject1MemberGroup, bool)`

GetMemberGroupOk returns a tuple with the MemberGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberGroup

`func (o *NSXTLoadBalancerPoolConfigObject1) SetMemberGroup(v NSXTLoadBalancerPoolConfigObject1MemberGroup)`

SetMemberGroup sets MemberGroup field to given value.

### HasMemberGroup

`func (o *NSXTLoadBalancerPoolConfigObject1) HasMemberGroup() bool`

HasMemberGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


