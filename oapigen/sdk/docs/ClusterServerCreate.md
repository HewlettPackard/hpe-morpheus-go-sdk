# ClusterServerCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**AddClusterWorkerRequestServerConfig**](AddClusterWorkerRequestServerConfig.md) |  | 
**ServerType** | Pointer to [**AddClusterWorkerRequestServerServerType**](AddClusterWorkerRequestServerServerType.md) |  | [optional] 
**Name** | **string** | Name to be used for host(s) created in the cluster | 
**Plan** | [**AddClusterWorkerRequestServerPlan**](AddClusterWorkerRequestServerPlan.md) |  | 
**ServicePlanOptions** | Pointer to [**AddClusterWorkerRequestServerServicePlanOptions**](AddClusterWorkerRequestServerServicePlanOptions.md) |  | [optional] 
**Volumes** | Pointer to [**[]AddClusterWorkerRequestServerVolumesInner**](AddClusterWorkerRequestServerVolumesInner.md) | The (optional) volumes parameter is for LV configuration, can create additional LVs at provision It should be passed as an array of Objects | [optional] 
**Network** | Pointer to [**AddClusterWorkerRequestServerNetwork**](AddClusterWorkerRequestServerNetwork.md) |  | [optional] 
**NetworkInterfaces** | Pointer to [**[]AddClusterWorkerRequestServerNetworkInterfacesInner**](AddClusterWorkerRequestServerNetworkInterfacesInner.md) | The networkInterfaces parameter is for network configuration.  The Options API /api/options/zoneNetworkOptions can be used to see which options are available.  It should be passed as an array of Objects with the following attributes  | [optional] 
**SecurityGroups** | Pointer to **[]string** | Key for security group configuration. | [optional] 
**Visibility** | Pointer to **string** | Visibility for server host | [optional] [default to "private"]
**UserGroup** | Pointer to [**AddClusterWorkerRequestServerUserGroup**](AddClusterWorkerRequestServerUserGroup.md) |  | [optional] 
**NetworkDomain** | Pointer to **NullableString** | Network domain | [optional] 
**Hostname** | Pointer to **NullableString** | Hostname for server host | [optional] 
**NodeCount** | Pointer to **int64** | Number of workers or hosts | [optional] 
**Tags** | Pointer to [**[]AddClusterWorkerRequestServerTagsInner**](AddClusterWorkerRequestServerTagsInner.md) | Metadata tags, Array of objects having a name and value. | [optional] 
**Labels** | Pointer to **[]string** | Array of strings (keywords). This will set labels on the server and also on the cluster as well by default. | [optional] 
**SshHosts** | Pointer to [**[]AddClusterWorkerRequestServerSshHostsInner**](AddClusterWorkerRequestServerSshHostsInner.md) | Array of Host IPs and Names. This is used in conjunction with sshUsername and sshPassword/sshKeyPair to add existing hosts such as with HPE VM clusters. | [optional] 
**SshMasterHosts** | Pointer to **string** | A string consisting of comma-separated master host IP addresses. | [optional] 
**SshWorkerHosts** | Pointer to **string** | A string consisting of comma-separated worker host IP addresses. | [optional] 
**SshPort** | Pointer to **int64** | The port which the worker&#39;s SSH server is listening on. | [optional] 
**SshUsername** | Pointer to **string** | SSH Username | [optional] 
**SshPassword** | Pointer to **NullableString** | SSH Password | [optional] 
**SshKeyPair** | Pointer to [**AddClusterWorkerRequestServerSshKeyPair**](AddClusterWorkerRequestServerSshKeyPair.md) |  | [optional] 
**DataDevice** | Pointer to **string** |  | [optional] 
**LvmEnabled** | Pointer to **bool** |  | [optional] 

## Methods

### NewClusterServerCreate

`func NewClusterServerCreate(config AddClusterWorkerRequestServerConfig, name string, plan AddClusterWorkerRequestServerPlan, ) *ClusterServerCreate`

NewClusterServerCreate instantiates a new ClusterServerCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterServerCreateWithDefaults

`func NewClusterServerCreateWithDefaults() *ClusterServerCreate`

NewClusterServerCreateWithDefaults instantiates a new ClusterServerCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *ClusterServerCreate) GetConfig() AddClusterWorkerRequestServerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ClusterServerCreate) GetConfigOk() (*AddClusterWorkerRequestServerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ClusterServerCreate) SetConfig(v AddClusterWorkerRequestServerConfig)`

SetConfig sets Config field to given value.


### GetServerType

`func (o *ClusterServerCreate) GetServerType() AddClusterWorkerRequestServerServerType`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *ClusterServerCreate) GetServerTypeOk() (*AddClusterWorkerRequestServerServerType, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *ClusterServerCreate) SetServerType(v AddClusterWorkerRequestServerServerType)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *ClusterServerCreate) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetName

`func (o *ClusterServerCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterServerCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterServerCreate) SetName(v string)`

SetName sets Name field to given value.


### GetPlan

`func (o *ClusterServerCreate) GetPlan() AddClusterWorkerRequestServerPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *ClusterServerCreate) GetPlanOk() (*AddClusterWorkerRequestServerPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *ClusterServerCreate) SetPlan(v AddClusterWorkerRequestServerPlan)`

SetPlan sets Plan field to given value.


### GetServicePlanOptions

`func (o *ClusterServerCreate) GetServicePlanOptions() AddClusterWorkerRequestServerServicePlanOptions`

GetServicePlanOptions returns the ServicePlanOptions field if non-nil, zero value otherwise.

### GetServicePlanOptionsOk

`func (o *ClusterServerCreate) GetServicePlanOptionsOk() (*AddClusterWorkerRequestServerServicePlanOptions, bool)`

GetServicePlanOptionsOk returns a tuple with the ServicePlanOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanOptions

`func (o *ClusterServerCreate) SetServicePlanOptions(v AddClusterWorkerRequestServerServicePlanOptions)`

SetServicePlanOptions sets ServicePlanOptions field to given value.

### HasServicePlanOptions

`func (o *ClusterServerCreate) HasServicePlanOptions() bool`

HasServicePlanOptions returns a boolean if a field has been set.

### GetVolumes

`func (o *ClusterServerCreate) GetVolumes() []AddClusterWorkerRequestServerVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *ClusterServerCreate) GetVolumesOk() (*[]AddClusterWorkerRequestServerVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *ClusterServerCreate) SetVolumes(v []AddClusterWorkerRequestServerVolumesInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *ClusterServerCreate) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetNetwork

`func (o *ClusterServerCreate) GetNetwork() AddClusterWorkerRequestServerNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ClusterServerCreate) GetNetworkOk() (*AddClusterWorkerRequestServerNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ClusterServerCreate) SetNetwork(v AddClusterWorkerRequestServerNetwork)`

SetNetwork sets Network field to given value.

### HasNetwork

`func (o *ClusterServerCreate) HasNetwork() bool`

HasNetwork returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *ClusterServerCreate) GetNetworkInterfaces() []AddClusterWorkerRequestServerNetworkInterfacesInner`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *ClusterServerCreate) GetNetworkInterfacesOk() (*[]AddClusterWorkerRequestServerNetworkInterfacesInner, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *ClusterServerCreate) SetNetworkInterfaces(v []AddClusterWorkerRequestServerNetworkInterfacesInner)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *ClusterServerCreate) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *ClusterServerCreate) GetSecurityGroups() []string`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *ClusterServerCreate) GetSecurityGroupsOk() (*[]string, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *ClusterServerCreate) SetSecurityGroups(v []string)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *ClusterServerCreate) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetVisibility

`func (o *ClusterServerCreate) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ClusterServerCreate) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ClusterServerCreate) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ClusterServerCreate) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetUserGroup

`func (o *ClusterServerCreate) GetUserGroup() AddClusterWorkerRequestServerUserGroup`

GetUserGroup returns the UserGroup field if non-nil, zero value otherwise.

### GetUserGroupOk

`func (o *ClusterServerCreate) GetUserGroupOk() (*AddClusterWorkerRequestServerUserGroup, bool)`

GetUserGroupOk returns a tuple with the UserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroup

`func (o *ClusterServerCreate) SetUserGroup(v AddClusterWorkerRequestServerUserGroup)`

SetUserGroup sets UserGroup field to given value.

### HasUserGroup

`func (o *ClusterServerCreate) HasUserGroup() bool`

HasUserGroup returns a boolean if a field has been set.

### GetNetworkDomain

`func (o *ClusterServerCreate) GetNetworkDomain() string`

GetNetworkDomain returns the NetworkDomain field if non-nil, zero value otherwise.

### GetNetworkDomainOk

`func (o *ClusterServerCreate) GetNetworkDomainOk() (*string, bool)`

GetNetworkDomainOk returns a tuple with the NetworkDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomain

`func (o *ClusterServerCreate) SetNetworkDomain(v string)`

SetNetworkDomain sets NetworkDomain field to given value.

### HasNetworkDomain

`func (o *ClusterServerCreate) HasNetworkDomain() bool`

HasNetworkDomain returns a boolean if a field has been set.

### SetNetworkDomainNil

`func (o *ClusterServerCreate) SetNetworkDomainNil(b bool)`

 SetNetworkDomainNil sets the value for NetworkDomain to be an explicit nil

### UnsetNetworkDomain
`func (o *ClusterServerCreate) UnsetNetworkDomain()`

UnsetNetworkDomain ensures that no value is present for NetworkDomain, not even an explicit nil
### GetHostname

`func (o *ClusterServerCreate) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *ClusterServerCreate) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *ClusterServerCreate) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *ClusterServerCreate) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### SetHostnameNil

`func (o *ClusterServerCreate) SetHostnameNil(b bool)`

 SetHostnameNil sets the value for Hostname to be an explicit nil

### UnsetHostname
`func (o *ClusterServerCreate) UnsetHostname()`

UnsetHostname ensures that no value is present for Hostname, not even an explicit nil
### GetNodeCount

`func (o *ClusterServerCreate) GetNodeCount() int64`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ClusterServerCreate) GetNodeCountOk() (*int64, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ClusterServerCreate) SetNodeCount(v int64)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *ClusterServerCreate) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetTags

`func (o *ClusterServerCreate) GetTags() []AddClusterWorkerRequestServerTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ClusterServerCreate) GetTagsOk() (*[]AddClusterWorkerRequestServerTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ClusterServerCreate) SetTags(v []AddClusterWorkerRequestServerTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ClusterServerCreate) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLabels

`func (o *ClusterServerCreate) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ClusterServerCreate) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ClusterServerCreate) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ClusterServerCreate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetSshHosts

`func (o *ClusterServerCreate) GetSshHosts() []AddClusterWorkerRequestServerSshHostsInner`

GetSshHosts returns the SshHosts field if non-nil, zero value otherwise.

### GetSshHostsOk

`func (o *ClusterServerCreate) GetSshHostsOk() (*[]AddClusterWorkerRequestServerSshHostsInner, bool)`

GetSshHostsOk returns a tuple with the SshHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshHosts

`func (o *ClusterServerCreate) SetSshHosts(v []AddClusterWorkerRequestServerSshHostsInner)`

SetSshHosts sets SshHosts field to given value.

### HasSshHosts

`func (o *ClusterServerCreate) HasSshHosts() bool`

HasSshHosts returns a boolean if a field has been set.

### GetSshMasterHosts

`func (o *ClusterServerCreate) GetSshMasterHosts() string`

GetSshMasterHosts returns the SshMasterHosts field if non-nil, zero value otherwise.

### GetSshMasterHostsOk

`func (o *ClusterServerCreate) GetSshMasterHostsOk() (*string, bool)`

GetSshMasterHostsOk returns a tuple with the SshMasterHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshMasterHosts

`func (o *ClusterServerCreate) SetSshMasterHosts(v string)`

SetSshMasterHosts sets SshMasterHosts field to given value.

### HasSshMasterHosts

`func (o *ClusterServerCreate) HasSshMasterHosts() bool`

HasSshMasterHosts returns a boolean if a field has been set.

### GetSshWorkerHosts

`func (o *ClusterServerCreate) GetSshWorkerHosts() string`

GetSshWorkerHosts returns the SshWorkerHosts field if non-nil, zero value otherwise.

### GetSshWorkerHostsOk

`func (o *ClusterServerCreate) GetSshWorkerHostsOk() (*string, bool)`

GetSshWorkerHostsOk returns a tuple with the SshWorkerHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshWorkerHosts

`func (o *ClusterServerCreate) SetSshWorkerHosts(v string)`

SetSshWorkerHosts sets SshWorkerHosts field to given value.

### HasSshWorkerHosts

`func (o *ClusterServerCreate) HasSshWorkerHosts() bool`

HasSshWorkerHosts returns a boolean if a field has been set.

### GetSshPort

`func (o *ClusterServerCreate) GetSshPort() int64`

GetSshPort returns the SshPort field if non-nil, zero value otherwise.

### GetSshPortOk

`func (o *ClusterServerCreate) GetSshPortOk() (*int64, bool)`

GetSshPortOk returns a tuple with the SshPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPort

`func (o *ClusterServerCreate) SetSshPort(v int64)`

SetSshPort sets SshPort field to given value.

### HasSshPort

`func (o *ClusterServerCreate) HasSshPort() bool`

HasSshPort returns a boolean if a field has been set.

### GetSshUsername

`func (o *ClusterServerCreate) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *ClusterServerCreate) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *ClusterServerCreate) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *ClusterServerCreate) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetSshPassword

`func (o *ClusterServerCreate) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *ClusterServerCreate) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *ClusterServerCreate) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *ClusterServerCreate) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### SetSshPasswordNil

`func (o *ClusterServerCreate) SetSshPasswordNil(b bool)`

 SetSshPasswordNil sets the value for SshPassword to be an explicit nil

### UnsetSshPassword
`func (o *ClusterServerCreate) UnsetSshPassword()`

UnsetSshPassword ensures that no value is present for SshPassword, not even an explicit nil
### GetSshKeyPair

`func (o *ClusterServerCreate) GetSshKeyPair() AddClusterWorkerRequestServerSshKeyPair`

GetSshKeyPair returns the SshKeyPair field if non-nil, zero value otherwise.

### GetSshKeyPairOk

`func (o *ClusterServerCreate) GetSshKeyPairOk() (*AddClusterWorkerRequestServerSshKeyPair, bool)`

GetSshKeyPairOk returns a tuple with the SshKeyPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyPair

`func (o *ClusterServerCreate) SetSshKeyPair(v AddClusterWorkerRequestServerSshKeyPair)`

SetSshKeyPair sets SshKeyPair field to given value.

### HasSshKeyPair

`func (o *ClusterServerCreate) HasSshKeyPair() bool`

HasSshKeyPair returns a boolean if a field has been set.

### GetDataDevice

`func (o *ClusterServerCreate) GetDataDevice() string`

GetDataDevice returns the DataDevice field if non-nil, zero value otherwise.

### GetDataDeviceOk

`func (o *ClusterServerCreate) GetDataDeviceOk() (*string, bool)`

GetDataDeviceOk returns a tuple with the DataDevice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataDevice

`func (o *ClusterServerCreate) SetDataDevice(v string)`

SetDataDevice sets DataDevice field to given value.

### HasDataDevice

`func (o *ClusterServerCreate) HasDataDevice() bool`

HasDataDevice returns a boolean if a field has been set.

### GetLvmEnabled

`func (o *ClusterServerCreate) GetLvmEnabled() bool`

GetLvmEnabled returns the LvmEnabled field if non-nil, zero value otherwise.

### GetLvmEnabledOk

`func (o *ClusterServerCreate) GetLvmEnabledOk() (*bool, bool)`

GetLvmEnabledOk returns a tuple with the LvmEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLvmEnabled

`func (o *ClusterServerCreate) SetLvmEnabled(v bool)`

SetLvmEnabled sets LvmEnabled field to given value.

### HasLvmEnabled

`func (o *ClusterServerCreate) HasLvmEnabled() bool`

HasLvmEnabled returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


