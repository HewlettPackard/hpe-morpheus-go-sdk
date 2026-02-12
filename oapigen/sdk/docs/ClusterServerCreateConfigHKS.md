# ClusterServerCreateConfigHKS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PodCidr** | Pointer to **string** |  | [optional] 
**ServiceCidr** | Pointer to **string** |  | [optional] 
**NodeCount** | Pointer to **int64** |  | [optional] 
**CreateUser** | Pointer to **bool** |  | [optional] 

## Methods

### NewClusterServerCreateConfigHKS

`func NewClusterServerCreateConfigHKS() *ClusterServerCreateConfigHKS`

NewClusterServerCreateConfigHKS instantiates a new ClusterServerCreateConfigHKS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterServerCreateConfigHKSWithDefaults

`func NewClusterServerCreateConfigHKSWithDefaults() *ClusterServerCreateConfigHKS`

NewClusterServerCreateConfigHKSWithDefaults instantiates a new ClusterServerCreateConfigHKS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPodCidr

`func (o *ClusterServerCreateConfigHKS) GetPodCidr() string`

GetPodCidr returns the PodCidr field if non-nil, zero value otherwise.

### GetPodCidrOk

`func (o *ClusterServerCreateConfigHKS) GetPodCidrOk() (*string, bool)`

GetPodCidrOk returns a tuple with the PodCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPodCidr

`func (o *ClusterServerCreateConfigHKS) SetPodCidr(v string)`

SetPodCidr sets PodCidr field to given value.

### HasPodCidr

`func (o *ClusterServerCreateConfigHKS) HasPodCidr() bool`

HasPodCidr returns a boolean if a field has been set.

### GetServiceCidr

`func (o *ClusterServerCreateConfigHKS) GetServiceCidr() string`

GetServiceCidr returns the ServiceCidr field if non-nil, zero value otherwise.

### GetServiceCidrOk

`func (o *ClusterServerCreateConfigHKS) GetServiceCidrOk() (*string, bool)`

GetServiceCidrOk returns a tuple with the ServiceCidr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceCidr

`func (o *ClusterServerCreateConfigHKS) SetServiceCidr(v string)`

SetServiceCidr sets ServiceCidr field to given value.

### HasServiceCidr

`func (o *ClusterServerCreateConfigHKS) HasServiceCidr() bool`

HasServiceCidr returns a boolean if a field has been set.

### GetNodeCount

`func (o *ClusterServerCreateConfigHKS) GetNodeCount() int64`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ClusterServerCreateConfigHKS) GetNodeCountOk() (*int64, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ClusterServerCreateConfigHKS) SetNodeCount(v int64)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *ClusterServerCreateConfigHKS) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetCreateUser

`func (o *ClusterServerCreateConfigHKS) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *ClusterServerCreateConfigHKS) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *ClusterServerCreateConfigHKS) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *ClusterServerCreateConfigHKS) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


