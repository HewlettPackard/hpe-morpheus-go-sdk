# ClusterCreateServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**ClusterCreateServerConfig**](ClusterCreateServerConfig.md) |  | 
**ServerType** | Pointer to [**ClusterCreateServerServerType**](ClusterCreateServerServerType.md) |  | [optional] 
**Name** | **string** | Name to be used for host(s) created in the cluster | 
**Plan** | [**ClusterCreateServerPlan**](ClusterCreateServerPlan.md) |  | 
**ServicePlanOptions** | Pointer to [**ClusterCreateServerServicePlanOptions**](ClusterCreateServerServicePlanOptions.md) |  | [optional] 
**Volumes** | Pointer to [**[]ClusterCreateServerVolumesInner**](ClusterCreateServerVolumesInner.md) | The (optional) volumes parameter is for LV configuration, can create additional LVs at provision It should be passed as an array of Objects | [optional] 
**NetworkInterfaces** | Pointer to [**[]ClusterCreateServerNetworkInterfacesInner**](ClusterCreateServerNetworkInterfacesInner.md) | The networkInterfaces parameter is for network configuration.  The Options API /api/options/zoneNetworkOptions can be used to see which options are available.  It should be passed as an array of Objects with the following attributes  | [optional] 
**SecurityGroups** | Pointer to **[]string** | Key for security group configuration. | [optional] 
**Visibility** | Pointer to **string** | Visibility for server host | [optional] [default to "private"]
**UserGroup** | Pointer to [**ClusterCreateServerUserGroup**](ClusterCreateServerUserGroup.md) |  | [optional] 
**NetworkDomain** | Pointer to **NullableString** | Network domain | [optional] 
**Hostname** | Pointer to **NullableString** | Hostname for server host | [optional] 
**NodeCount** | Pointer to **int64** | Number of workers or hosts | [optional] 
**Tags** | Pointer to [**[]ClusterCreateServerTagsInner**](ClusterCreateServerTagsInner.md) | Metadata tags, Array of objects having a name and value. | [optional] 
**Labels** | Pointer to **[]string** | Array of strings (keywords). This will set labels on the server and also on the cluster as well by default. | [optional] 
**SshHosts** | Pointer to [**[]ClusterCreateServerSshHostsInner**](ClusterCreateServerSshHostsInner.md) | Array of Host IPs and Names. This is used in conjunction with sshUsername and sshPassword/sshKeyPair to add existing hosts such as with HPE VM clusters. | [optional] 
**SshUsername** | Pointer to **string** | SSH Username | [optional] 
**SshPassword** | Pointer to **NullableString** | SSH Password | [optional] 
**SshKeyPair** | Pointer to [**ClusterCreateServerSshKeyPair**](ClusterCreateServerSshKeyPair.md) |  | [optional] 

## Methods

### NewClusterCreateServer

`func NewClusterCreateServer(config ClusterCreateServerConfig, name string, plan ClusterCreateServerPlan, ) *ClusterCreateServer`

NewClusterCreateServer instantiates a new ClusterCreateServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterCreateServerWithDefaults

`func NewClusterCreateServerWithDefaults() *ClusterCreateServer`

NewClusterCreateServerWithDefaults instantiates a new ClusterCreateServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *ClusterCreateServer) GetConfig() ClusterCreateServerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ClusterCreateServer) GetConfigOk() (*ClusterCreateServerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ClusterCreateServer) SetConfig(v ClusterCreateServerConfig)`

SetConfig sets Config field to given value.


### GetServerType

`func (o *ClusterCreateServer) GetServerType() ClusterCreateServerServerType`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *ClusterCreateServer) GetServerTypeOk() (*ClusterCreateServerServerType, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *ClusterCreateServer) SetServerType(v ClusterCreateServerServerType)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *ClusterCreateServer) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetName

`func (o *ClusterCreateServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterCreateServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterCreateServer) SetName(v string)`

SetName sets Name field to given value.


### GetPlan

`func (o *ClusterCreateServer) GetPlan() ClusterCreateServerPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *ClusterCreateServer) GetPlanOk() (*ClusterCreateServerPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *ClusterCreateServer) SetPlan(v ClusterCreateServerPlan)`

SetPlan sets Plan field to given value.


### GetServicePlanOptions

`func (o *ClusterCreateServer) GetServicePlanOptions() ClusterCreateServerServicePlanOptions`

GetServicePlanOptions returns the ServicePlanOptions field if non-nil, zero value otherwise.

### GetServicePlanOptionsOk

`func (o *ClusterCreateServer) GetServicePlanOptionsOk() (*ClusterCreateServerServicePlanOptions, bool)`

GetServicePlanOptionsOk returns a tuple with the ServicePlanOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanOptions

`func (o *ClusterCreateServer) SetServicePlanOptions(v ClusterCreateServerServicePlanOptions)`

SetServicePlanOptions sets ServicePlanOptions field to given value.

### HasServicePlanOptions

`func (o *ClusterCreateServer) HasServicePlanOptions() bool`

HasServicePlanOptions returns a boolean if a field has been set.

### GetVolumes

`func (o *ClusterCreateServer) GetVolumes() []ClusterCreateServerVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *ClusterCreateServer) GetVolumesOk() (*[]ClusterCreateServerVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *ClusterCreateServer) SetVolumes(v []ClusterCreateServerVolumesInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *ClusterCreateServer) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *ClusterCreateServer) GetNetworkInterfaces() []ClusterCreateServerNetworkInterfacesInner`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *ClusterCreateServer) GetNetworkInterfacesOk() (*[]ClusterCreateServerNetworkInterfacesInner, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *ClusterCreateServer) SetNetworkInterfaces(v []ClusterCreateServerNetworkInterfacesInner)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *ClusterCreateServer) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *ClusterCreateServer) GetSecurityGroups() []string`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *ClusterCreateServer) GetSecurityGroupsOk() (*[]string, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *ClusterCreateServer) SetSecurityGroups(v []string)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *ClusterCreateServer) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetVisibility

`func (o *ClusterCreateServer) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ClusterCreateServer) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ClusterCreateServer) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ClusterCreateServer) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetUserGroup

`func (o *ClusterCreateServer) GetUserGroup() ClusterCreateServerUserGroup`

GetUserGroup returns the UserGroup field if non-nil, zero value otherwise.

### GetUserGroupOk

`func (o *ClusterCreateServer) GetUserGroupOk() (*ClusterCreateServerUserGroup, bool)`

GetUserGroupOk returns a tuple with the UserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroup

`func (o *ClusterCreateServer) SetUserGroup(v ClusterCreateServerUserGroup)`

SetUserGroup sets UserGroup field to given value.

### HasUserGroup

`func (o *ClusterCreateServer) HasUserGroup() bool`

HasUserGroup returns a boolean if a field has been set.

### GetNetworkDomain

`func (o *ClusterCreateServer) GetNetworkDomain() string`

GetNetworkDomain returns the NetworkDomain field if non-nil, zero value otherwise.

### GetNetworkDomainOk

`func (o *ClusterCreateServer) GetNetworkDomainOk() (*string, bool)`

GetNetworkDomainOk returns a tuple with the NetworkDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomain

`func (o *ClusterCreateServer) SetNetworkDomain(v string)`

SetNetworkDomain sets NetworkDomain field to given value.

### HasNetworkDomain

`func (o *ClusterCreateServer) HasNetworkDomain() bool`

HasNetworkDomain returns a boolean if a field has been set.

### SetNetworkDomainNil

`func (o *ClusterCreateServer) SetNetworkDomainNil(b bool)`

 SetNetworkDomainNil sets the value for NetworkDomain to be an explicit nil

### UnsetNetworkDomain
`func (o *ClusterCreateServer) UnsetNetworkDomain()`

UnsetNetworkDomain ensures that no value is present for NetworkDomain, not even an explicit nil
### GetHostname

`func (o *ClusterCreateServer) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ClusterCreateServer) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ClusterCreateServer) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ClusterCreateServer) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### SetHostnameNil

`func (o *ClusterCreateServer) SetHostnameNil(b bool)`

 SetHostnameNil sets the value for Hostname to be an explicit nil

### UnsetHostname
`func (o *ClusterCreateServer) UnsetHostname()`

UnsetHostname ensures that no value is present for Hostname, not even an explicit nil
### GetNodeCount

`func (o *ClusterCreateServer) GetNodeCount() int64`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ClusterCreateServer) GetNodeCountOk() (*int64, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ClusterCreateServer) SetNodeCount(v int64)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *ClusterCreateServer) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetTags

`func (o *ClusterCreateServer) GetTags() []ClusterCreateServerTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ClusterCreateServer) GetTagsOk() (*[]ClusterCreateServerTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ClusterCreateServer) SetTags(v []ClusterCreateServerTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ClusterCreateServer) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLabels

`func (o *ClusterCreateServer) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ClusterCreateServer) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ClusterCreateServer) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ClusterCreateServer) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetSshHosts

`func (o *ClusterCreateServer) GetSshHosts() []ClusterCreateServerSshHostsInner`

GetSshHosts returns the SshHosts field if non-nil, zero value otherwise.

### GetSshHostsOk

`func (o *ClusterCreateServer) GetSshHostsOk() (*[]ClusterCreateServerSshHostsInner, bool)`

GetSshHostsOk returns a tuple with the SshHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshHosts

`func (o *ClusterCreateServer) SetSshHosts(v []ClusterCreateServerSshHostsInner)`

SetSshHosts sets SshHosts field to given value.

### HasSshHosts

`func (o *ClusterCreateServer) HasSshHosts() bool`

HasSshHosts returns a boolean if a field has been set.

### GetSshUsername

`func (o *ClusterCreateServer) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *ClusterCreateServer) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *ClusterCreateServer) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *ClusterCreateServer) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetSshPassword

`func (o *ClusterCreateServer) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *ClusterCreateServer) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *ClusterCreateServer) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *ClusterCreateServer) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### SetSshPasswordNil

`func (o *ClusterCreateServer) SetSshPasswordNil(b bool)`

 SetSshPasswordNil sets the value for SshPassword to be an explicit nil

### UnsetSshPassword
`func (o *ClusterCreateServer) UnsetSshPassword()`

UnsetSshPassword ensures that no value is present for SshPassword, not even an explicit nil
### GetSshKeyPair

`func (o *ClusterCreateServer) GetSshKeyPair() ClusterCreateServerSshKeyPair`

GetSshKeyPair returns the SshKeyPair field if non-nil, zero value otherwise.

### GetSshKeyPairOk

`func (o *ClusterCreateServer) GetSshKeyPairOk() (*ClusterCreateServerSshKeyPair, bool)`

GetSshKeyPairOk returns a tuple with the SshKeyPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyPair

`func (o *ClusterCreateServer) SetSshKeyPair(v ClusterCreateServerSshKeyPair)`

SetSshKeyPair sets SshKeyPair field to given value.

### HasSshKeyPair

`func (o *ClusterCreateServer) HasSshKeyPair() bool`

HasSshKeyPair returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


