# ClusterServerCreateConfigAKS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeCount** | Pointer to **int64** | Number of workers or hosts | [optional] 
**CreateUser** | Pointer to **bool** |  | [optional] 

## Methods

### NewClusterServerCreateConfigAKS

`func NewClusterServerCreateConfigAKS() *ClusterServerCreateConfigAKS`

NewClusterServerCreateConfigAKS instantiates a new ClusterServerCreateConfigAKS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterServerCreateConfigAKSWithDefaults

`func NewClusterServerCreateConfigAKSWithDefaults() *ClusterServerCreateConfigAKS`

NewClusterServerCreateConfigAKSWithDefaults instantiates a new ClusterServerCreateConfigAKS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeCount

`func (o *ClusterServerCreateConfigAKS) GetNodeCount() int64`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ClusterServerCreateConfigAKS) GetNodeCountOk() (*int64, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ClusterServerCreateConfigAKS) SetNodeCount(v int64)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *ClusterServerCreateConfigAKS) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetCreateUser

`func (o *ClusterServerCreateConfigAKS) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *ClusterServerCreateConfigAKS) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *ClusterServerCreateConfigAKS) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *ClusterServerCreateConfigAKS) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


