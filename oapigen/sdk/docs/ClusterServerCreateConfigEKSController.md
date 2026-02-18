# ClusterServerCreateConfigEKSController

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Role** | Pointer to [**ClusterServerCreateConfigEKSControllerRole**](ClusterServerCreateConfigEKSControllerRole.md) |  | [optional] 
**Network** | Pointer to [**ClusterServerCreateConfigEKSControllerNetwork**](ClusterServerCreateConfigEKSControllerNetwork.md) |  | [optional] 
**SecurityGroup** | Pointer to [**ClusterServerCreateConfigEKSControllerSecurityGroup**](ClusterServerCreateConfigEKSControllerSecurityGroup.md) |  | [optional] 

## Methods

### NewClusterServerCreateConfigEKSController

`func NewClusterServerCreateConfigEKSController() *ClusterServerCreateConfigEKSController`

NewClusterServerCreateConfigEKSController instantiates a new ClusterServerCreateConfigEKSController object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterServerCreateConfigEKSControllerWithDefaults

`func NewClusterServerCreateConfigEKSControllerWithDefaults() *ClusterServerCreateConfigEKSController`

NewClusterServerCreateConfigEKSControllerWithDefaults instantiates a new ClusterServerCreateConfigEKSController object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRole

`func (o *ClusterServerCreateConfigEKSController) GetRole() ClusterServerCreateConfigEKSControllerRole`

GetRole returns the Role field if non-nil, zero value otherwise.

### GetRoleOk

`func (o *ClusterServerCreateConfigEKSController) GetRoleOk() (*ClusterServerCreateConfigEKSControllerRole, bool)`

GetRoleOk returns a tuple with the Role field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRole

`func (o *ClusterServerCreateConfigEKSController) SetRole(v ClusterServerCreateConfigEKSControllerRole)`

SetRole sets Role field to given value.

### HasRole

`func (o *ClusterServerCreateConfigEKSController) HasRole() bool`

HasRole returns a boolean if a field has been set.

### GetNetwork

`func (o *ClusterServerCreateConfigEKSController) GetNetwork() ClusterServerCreateConfigEKSControllerNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ClusterServerCreateConfigEKSController) GetNetworkOk() (*ClusterServerCreateConfigEKSControllerNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ClusterServerCreateConfigEKSController) SetNetwork(v ClusterServerCreateConfigEKSControllerNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ClusterServerCreateConfigEKSController) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetSecurityGroup

`func (o *ClusterServerCreateConfigEKSController) GetSecurityGroup() ClusterServerCreateConfigEKSControllerSecurityGroup`

GetSecurityGroup returns the SecurityGroup field if non-nil, zero value otherwise.

### GetSecurityGroupOk

`func (o *ClusterServerCreateConfigEKSController) GetSecurityGroupOk() (*ClusterServerCreateConfigEKSControllerSecurityGroup, bool)`

GetSecurityGroupOk returns a tuple with the SecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroup

`func (o *ClusterServerCreateConfigEKSController) SetSecurityGroup(v ClusterServerCreateConfigEKSControllerSecurityGroup)`

SetSecurityGroup sets SecurityGroup field to given value.

### HasSecurityGroup

`func (o *ClusterServerCreateConfigEKSController) HasSecurityGroup() bool`

HasSecurityGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


