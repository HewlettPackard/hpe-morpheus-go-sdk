# LoadBalancerPoolCreateConfigNSXTMemberGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | NSX-T member group path. | [optional] 
**IpRevisionFilter** | Pointer to **string** | IP revision filter for the member group. | [optional] 
**MaxIpListSize** | Pointer to **NullableInt64** | Maximum IP list size for the member group. | [optional] 
**Port** | Pointer to **NullableInt64** | Port number for the member group. | [optional] 

## Methods

### NewLoadBalancerPoolCreateConfigNSXTMemberGroup

`func NewLoadBalancerPoolCreateConfigNSXTMemberGroup() *LoadBalancerPoolCreateConfigNSXTMemberGroup`

NewLoadBalancerPoolCreateConfigNSXTMemberGroup instantiates a new LoadBalancerPoolCreateConfigNSXTMemberGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoadBalancerPoolCreateConfigNSXTMemberGroupWithDefaults

`func NewLoadBalancerPoolCreateConfigNSXTMemberGroupWithDefaults() *LoadBalancerPoolCreateConfigNSXTMemberGroup`

NewLoadBalancerPoolCreateConfigNSXTMemberGroupWithDefaults instantiates a new LoadBalancerPoolCreateConfigNSXTMemberGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *LoadBalancerPoolCreateConfigNSXTMemberGroup) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *LoadBalancerPoolCreateConfigNSXTMemberGroup) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *LoadBalancerPoolCreateConfigNSXTMemberGroup) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *LoadBalancerPoolCreateConfigNSXTMemberGroup) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetIpRevisionFilter

`func (o *LoadBalancerPoolCreateConfigNSXTMemberGroup) GetIpRevisionFilter() string`

GetIpRevisionFilter returns the IpRevisionFilter field if non-nil, zero value otherwise.

### GetIpRevisionFilterOk

`func (o *LoadBalancerPoolCreateConfigNSXTMemberGroup) GetIpRevisionFilterOk() (*string, bool)`

GetIpRevisionFilterOk returns a tuple with the IpRevisionFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRevisionFilter

`func (o *LoadBalancerPoolCreateConfigNSXTMemberGroup) SetIpRevisionFilter(v string)`

SetIpRevisionFilter sets IpRevisionFilter field to given value.

### HasIpRevisionFilter

`func (o *LoadBalancerPoolCreateConfigNSXTMemberGroup) HasIpRevisionFilter() bool`

HasIpRevisionFilter returns a boolean if a field has been set.

### GetMaxIpListSize

`func (o *LoadBalancerPoolCreateConfigNSXTMemberGroup) GetMaxIpListSize() int64`

GetMaxIpListSize returns the MaxIpListSize field if non-nil, zero value otherwise.

### GetMaxIpListSizeOk

`func (o *LoadBalancerPoolCreateConfigNSXTMemberGroup) GetMaxIpListSizeOk() (*int64, bool)`

GetMaxIpListSizeOk returns a tuple with the MaxIpListSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIpListSize

`func (o *LoadBalancerPoolCreateConfigNSXTMemberGroup) SetMaxIpListSize(v int64)`

SetMaxIpListSize sets MaxIpListSize field to given value.

### HasMaxIpListSize

`func (o *LoadBalancerPoolCreateConfigNSXTMemberGroup) HasMaxIpListSize() bool`

HasMaxIpListSize returns a boolean if a field has been set.

### SetMaxIpListSizeNil

`func (o *LoadBalancerPoolCreateConfigNSXTMemberGroup) SetMaxIpListSizeNil(b bool)`

 SetMaxIpListSizeNil sets the value for MaxIpListSize to be an explicit nil

### UnsetMaxIpListSize
`func (o *LoadBalancerPoolCreateConfigNSXTMemberGroup) UnsetMaxIpListSize()`

UnsetMaxIpListSize ensures that no value is present for MaxIpListSize, not even an explicit nil
### GetPort

`func (o *LoadBalancerPoolCreateConfigNSXTMemberGroup) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *LoadBalancerPoolCreateConfigNSXTMemberGroup) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *LoadBalancerPoolCreateConfigNSXTMemberGroup) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *LoadBalancerPoolCreateConfigNSXTMemberGroup) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *LoadBalancerPoolCreateConfigNSXTMemberGroup) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *LoadBalancerPoolCreateConfigNSXTMemberGroup) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


