# ClusterServerCreateConfigEKSWorker

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Network** | Pointer to [**ClusterServerCreateConfigEKSWorkerNetwork**](ClusterServerCreateConfigEKSWorkerNetwork.md) |  | [optional] 
**SecurityGroup** | Pointer to [**ClusterServerCreateConfigEKSWorkerSecurityGroup**](ClusterServerCreateConfigEKSWorkerSecurityGroup.md) |  | [optional] 
**Plan** | Pointer to [**ClusterServerCreateConfigEKSWorkerPlan**](ClusterServerCreateConfigEKSWorkerPlan.md) |  | [optional] 
**Role** | Pointer to [**ClusterServerCreateConfigEKSWorkerRole**](ClusterServerCreateConfigEKSWorkerRole.md) |  | [optional] 

## Methods

### NewClusterServerCreateConfigEKSWorker

`func NewClusterServerCreateConfigEKSWorker() *ClusterServerCreateConfigEKSWorker`

NewClusterServerCreateConfigEKSWorker instantiates a new ClusterServerCreateConfigEKSWorker object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterServerCreateConfigEKSWorkerWithDefaults

`func NewClusterServerCreateConfigEKSWorkerWithDefaults() *ClusterServerCreateConfigEKSWorker`

NewClusterServerCreateConfigEKSWorkerWithDefaults instantiates a new ClusterServerCreateConfigEKSWorker object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNetwork

`func (o *ClusterServerCreateConfigEKSWorker) GetNetwork() ClusterServerCreateConfigEKSWorkerNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ClusterServerCreateConfigEKSWorker) GetNetworkOk() (*ClusterServerCreateConfigEKSWorkerNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ClusterServerCreateConfigEKSWorker) SetNetwork(v ClusterServerCreateConfigEKSWorkerNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ClusterServerCreateConfigEKSWorker) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetSecurityGroup

`func (o *ClusterServerCreateConfigEKSWorker) GetSecurityGroup() ClusterServerCreateConfigEKSWorkerSecurityGroup`

GetSecurityGroup returns the SecurityGroup field if non-nil, zero value otherwise.

### GetSecurityGroupOk

`func (o *ClusterServerCreateConfigEKSWorker) GetSecurityGroupOk() (*ClusterServerCreateConfigEKSWorkerSecurityGroup, bool)`

GetSecurityGroupOk returns a tuple with the SecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroup

`func (o *ClusterServerCreateConfigEKSWorker) SetSecurityGroup(v ClusterServerCreateConfigEKSWorkerSecurityGroup)`

SetSecurityGroup sets SecurityGroup field to given value.

### HasSecurityGroup

`func (o *ClusterServerCreateConfigEKSWorker) HasSecurityGroup() bool`

HasSecurityGroup returns a boolean if a field has been set.

### GetPlan

`func (o *ClusterServerCreateConfigEKSWorker) GetPlan() ClusterServerCreateConfigEKSWorkerPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *ClusterServerCreateConfigEKSWorker) GetPlanOk() (*ClusterServerCreateConfigEKSWorkerPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *ClusterServerCreateConfigEKSWorker) SetPlan(v ClusterServerCreateConfigEKSWorkerPlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *ClusterServerCreateConfigEKSWorker) HasPlan() bool`

HasPlan returns a boolean if a field has been set.

### GetRole

`func (o *ClusterServerCreateConfigEKSWorker) GetRole() ClusterServerCreateConfigEKSWorkerRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ClusterServerCreateConfigEKSWorker) GetRoleOk() (*ClusterServerCreateConfigEKSWorkerRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ClusterServerCreateConfigEKSWorker) SetRole(v ClusterServerCreateConfigEKSWorkerRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *ClusterServerCreateConfigEKSWorker) HasRole() bool`

HasRole returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


