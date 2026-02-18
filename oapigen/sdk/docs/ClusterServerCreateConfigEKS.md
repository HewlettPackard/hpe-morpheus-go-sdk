# ClusterServerCreateConfigEKS

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Controller** | Pointer to [**ClusterServerCreateConfigEKSController**](ClusterServerCreateConfigEKSController.md) |  | [optional] 
**Worker** | Pointer to [**ClusterServerCreateConfigEKSWorker**](ClusterServerCreateConfigEKSWorker.md) |  | [optional] 
**PublicIpType** | Pointer to **string** |  | [optional] 
**NodeCount** | Pointer to **int64** |  | [optional] 
**CreateUser** | Pointer to **bool** |  | [optional] 

## Methods

### NewClusterServerCreateConfigEKS

`func NewClusterServerCreateConfigEKS() *ClusterServerCreateConfigEKS`

NewClusterServerCreateConfigEKS instantiates a new ClusterServerCreateConfigEKS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterServerCreateConfigEKSWithDefaults

`func NewClusterServerCreateConfigEKSWithDefaults() *ClusterServerCreateConfigEKS`

NewClusterServerCreateConfigEKSWithDefaults instantiates a new ClusterServerCreateConfigEKS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetController

`func (o *ClusterServerCreateConfigEKS) GetController() ClusterServerCreateConfigEKSController`

GetController returns the Controller field if non-nil, zero value otherwise.

### GetControllerOk

`func (o *ClusterServerCreateConfigEKS) GetControllerOk() (*ClusterServerCreateConfigEKSController, bool)`

GetControllerOk returns a tuple with the Controller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetController

`func (o *ClusterServerCreateConfigEKS) SetController(v ClusterServerCreateConfigEKSController)`

SetController sets Controller field to given value.

### HasController

`func (o *ClusterServerCreateConfigEKS) HasController() bool`

HasController returns a boolean if a field has been set.

### GetWorker

`func (o *ClusterServerCreateConfigEKS) GetWorker() ClusterServerCreateConfigEKSWorker`

GetWorker returns the Worker field if non-nil, zero value otherwise.

### GetWorkerOk

`func (o *ClusterServerCreateConfigEKS) GetWorkerOk() (*ClusterServerCreateConfigEKSWorker, bool)`

GetWorkerOk returns a tuple with the Worker field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorker

`func (o *ClusterServerCreateConfigEKS) SetWorker(v ClusterServerCreateConfigEKSWorker)`

SetWorker sets Worker field to given value.

### HasWorker

`func (o *ClusterServerCreateConfigEKS) HasWorker() bool`

HasWorker returns a boolean if a field has been set.

### GetPublicIpType

`func (o *ClusterServerCreateConfigEKS) GetPublicIpType() string`

GetPublicIpType returns the PublicIpType field if non-nil, zero value otherwise.

### GetPublicIpTypeOk

`func (o *ClusterServerCreateConfigEKS) GetPublicIpTypeOk() (*string, bool)`

GetPublicIpTypeOk returns a tuple with the PublicIpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpType

`func (o *ClusterServerCreateConfigEKS) SetPublicIpType(v string)`

SetPublicIpType sets PublicIpType field to given value.

### HasPublicIpType

`func (o *ClusterServerCreateConfigEKS) HasPublicIpType() bool`

HasPublicIpType returns a boolean if a field has been set.

### GetNodeCount

`func (o *ClusterServerCreateConfigEKS) GetNodeCount() int64`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ClusterServerCreateConfigEKS) GetNodeCountOk() (*int64, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ClusterServerCreateConfigEKS) SetNodeCount(v int64)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *ClusterServerCreateConfigEKS) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetCreateUser

`func (o *ClusterServerCreateConfigEKS) GetCreateUser() bool`

GetCreateUser returns the CreateUser field if non-nil, zero value otherwise.

### GetCreateUserOk

`func (o *ClusterServerCreateConfigEKS) GetCreateUserOk() (*bool, bool)`

GetCreateUserOk returns a tuple with the CreateUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateUser

`func (o *ClusterServerCreateConfigEKS) SetCreateUser(v bool)`

SetCreateUser sets CreateUser field to given value.

### HasCreateUser

`func (o *ClusterServerCreateConfigEKS) HasCreateUser() bool`

HasCreateUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


