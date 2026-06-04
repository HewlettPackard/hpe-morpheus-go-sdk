# NSXTLoadBalancerPoolConfigObject1MemberGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Path** | Pointer to **string** | NSX-T member group path. | [optional] 
**IpRevisionFilter** | Pointer to **string** | IP revision filter for the member group. | [optional] 
**MaxIpListSize** | Pointer to **NullableInt64** | Maximum IP list size for the member group. | [optional] 
**Port** | Pointer to **NullableInt64** | Port number for the member group. | [optional] 

## Methods

### NewNSXTLoadBalancerPoolConfigObject1MemberGroup

`func NewNSXTLoadBalancerPoolConfigObject1MemberGroup() *NSXTLoadBalancerPoolConfigObject1MemberGroup`

NewNSXTLoadBalancerPoolConfigObject1MemberGroup instantiates a new NSXTLoadBalancerPoolConfigObject1MemberGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNSXTLoadBalancerPoolConfigObject1MemberGroupWithDefaults

`func NewNSXTLoadBalancerPoolConfigObject1MemberGroupWithDefaults() *NSXTLoadBalancerPoolConfigObject1MemberGroup`

NewNSXTLoadBalancerPoolConfigObject1MemberGroupWithDefaults instantiates a new NSXTLoadBalancerPoolConfigObject1MemberGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPath

`func (o *NSXTLoadBalancerPoolConfigObject1MemberGroup) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *NSXTLoadBalancerPoolConfigObject1MemberGroup) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *NSXTLoadBalancerPoolConfigObject1MemberGroup) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *NSXTLoadBalancerPoolConfigObject1MemberGroup) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetIpRevisionFilter

`func (o *NSXTLoadBalancerPoolConfigObject1MemberGroup) GetIpRevisionFilter() string`

GetIpRevisionFilter returns the IpRevisionFilter field if non-nil, zero value otherwise.

### GetIpRevisionFilterOk

`func (o *NSXTLoadBalancerPoolConfigObject1MemberGroup) GetIpRevisionFilterOk() (*string, bool)`

GetIpRevisionFilterOk returns a tuple with the IpRevisionFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIpRevisionFilter

`func (o *NSXTLoadBalancerPoolConfigObject1MemberGroup) SetIpRevisionFilter(v string)`

SetIpRevisionFilter sets IpRevisionFilter field to given value.

### HasIpRevisionFilter

`func (o *NSXTLoadBalancerPoolConfigObject1MemberGroup) HasIpRevisionFilter() bool`

HasIpRevisionFilter returns a boolean if a field has been set.

### GetMaxIpListSize

`func (o *NSXTLoadBalancerPoolConfigObject1MemberGroup) GetMaxIpListSize() int64`

GetMaxIpListSize returns the MaxIpListSize field if non-nil, zero value otherwise.

### GetMaxIpListSizeOk

`func (o *NSXTLoadBalancerPoolConfigObject1MemberGroup) GetMaxIpListSizeOk() (*int64, bool)`

GetMaxIpListSizeOk returns a tuple with the MaxIpListSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxIpListSize

`func (o *NSXTLoadBalancerPoolConfigObject1MemberGroup) SetMaxIpListSize(v int64)`

SetMaxIpListSize sets MaxIpListSize field to given value.

### HasMaxIpListSize

`func (o *NSXTLoadBalancerPoolConfigObject1MemberGroup) HasMaxIpListSize() bool`

HasMaxIpListSize returns a boolean if a field has been set.

### SetMaxIpListSizeNil

`func (o *NSXTLoadBalancerPoolConfigObject1MemberGroup) SetMaxIpListSizeNil(b bool)`

 SetMaxIpListSizeNil sets the value for MaxIpListSize to be an explicit nil

### UnsetMaxIpListSize
`func (o *NSXTLoadBalancerPoolConfigObject1MemberGroup) UnsetMaxIpListSize()`

UnsetMaxIpListSize ensures that no value is present for MaxIpListSize, not even an explicit nil
### GetPort

`func (o *NSXTLoadBalancerPoolConfigObject1MemberGroup) GetPort() int64`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *NSXTLoadBalancerPoolConfigObject1MemberGroup) GetPortOk() (*int64, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *NSXTLoadBalancerPoolConfigObject1MemberGroup) SetPort(v int64)`

SetPort sets Port field to given value.

### HasPort

`func (o *NSXTLoadBalancerPoolConfigObject1MemberGroup) HasPort() bool`

HasPort returns a boolean if a field has been set.

### SetPortNil

`func (o *NSXTLoadBalancerPoolConfigObject1MemberGroup) SetPortNil(b bool)`

 SetPortNil sets the value for Port to be an explicit nil

### UnsetPort
`func (o *NSXTLoadBalancerPoolConfigObject1MemberGroup) UnsetPort()`

UnsetPort ensures that no value is present for Port, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


