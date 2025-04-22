# AddClusterRequestClusterServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | [**AddClusterRequestClusterServerConfig**](AddClusterRequestClusterServerConfig.md) |  | 
**ServerType** | Pointer to [**AddClusterRequestClusterServerServerType**](AddClusterRequestClusterServerServerType.md) |  | [optional] 
**Name** | **string** | Name to be used for host(s) created in the cluster | 
**Plan** | [**AddClusterRequestClusterServerPlan**](AddClusterRequestClusterServerPlan.md) |  | 
**ServicePlanOptions** | Pointer to [**AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigServicePlanOptions**](AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigServicePlanOptions.md) |  | [optional] 
**Volumes** | Pointer to [**[]AddClusterRequestClusterServerVolumesInner**](AddClusterRequestClusterServerVolumesInner.md) | The (optional) volumes parameter is for LV configuration, can create additional LVs at provision It should be passed as an array of Objects | [optional] 
**NetworkInterfaces** | Pointer to [**[]AddClusterRequestClusterServerNetworkInterfacesInner**](AddClusterRequestClusterServerNetworkInterfacesInner.md) | The networkInterfaces parameter is for network configuration.  The Options API /api/options/zoneNetworkOptions can be used to see which options are available.  It should be passed as an array of Objects with the following attributes  | [optional] 
**SecurityGroups** | Pointer to **[]string** | Key for security group configuration. | [optional] 
**Visibility** | Pointer to **string** | Visibility for server host | [optional] [default to "private"]
**UserGroup** | Pointer to [**AddClusterRequestClusterServerUserGroup**](AddClusterRequestClusterServerUserGroup.md) |  | [optional] 
**NetworkDomain** | Pointer to **string** | Network domain | [optional] 
**Hostname** | Pointer to **string** | Hostname for server host | [optional] 
**NodeCount** | Pointer to **int64** | Number of workers or hosts | [optional] 
**Tags** | Pointer to [**[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner**](AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner.md) | Metadata tags, Array of objects having a name and value. | [optional] 
**Labels** | Pointer to **[]string** | Array of strings (keywords). This will set labels on the server and also on the cluster as well by default. | [optional] 
**SshHosts** | Pointer to [**[]AddClusterRequestClusterServerSshHostsInner**](AddClusterRequestClusterServerSshHostsInner.md) | Array of Host IPs and Names. This is used in conjunction with sshUsername and sshPassword/sshKeyPair to add existing hosts such as with HPE VM clusters. | [optional] 
**SshUsername** | Pointer to **string** | SSH Username | [optional] 
**SshPassword** | Pointer to **string** | SSH Password | [optional] 
**SshKeyPair** | Pointer to [**AddClusterRequestClusterServerSshKeyPair**](AddClusterRequestClusterServerSshKeyPair.md) |  | [optional] 

## Methods

### NewAddClusterRequestClusterServer

`func NewAddClusterRequestClusterServer(config AddClusterRequestClusterServerConfig, name string, plan AddClusterRequestClusterServerPlan, ) *AddClusterRequestClusterServer`

NewAddClusterRequestClusterServer instantiates a new AddClusterRequestClusterServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddClusterRequestClusterServerWithDefaults

`func NewAddClusterRequestClusterServerWithDefaults() *AddClusterRequestClusterServer`

NewAddClusterRequestClusterServerWithDefaults instantiates a new AddClusterRequestClusterServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *AddClusterRequestClusterServer) GetConfig() AddClusterRequestClusterServerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddClusterRequestClusterServer) GetConfigOk() (*AddClusterRequestClusterServerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddClusterRequestClusterServer) SetConfig(v AddClusterRequestClusterServerConfig)`

SetConfig sets Config field to given value.


### GetServerType

`func (o *AddClusterRequestClusterServer) GetServerType() AddClusterRequestClusterServerServerType`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *AddClusterRequestClusterServer) GetServerTypeOk() (*AddClusterRequestClusterServerServerType, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *AddClusterRequestClusterServer) SetServerType(v AddClusterRequestClusterServerServerType)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *AddClusterRequestClusterServer) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetName

`func (o *AddClusterRequestClusterServer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddClusterRequestClusterServer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddClusterRequestClusterServer) SetName(v string)`

SetName sets Name field to given value.


### GetPlan

`func (o *AddClusterRequestClusterServer) GetPlan() AddClusterRequestClusterServerPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *AddClusterRequestClusterServer) GetPlanOk() (*AddClusterRequestClusterServerPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *AddClusterRequestClusterServer) SetPlan(v AddClusterRequestClusterServerPlan)`

SetPlan sets Plan field to given value.


### GetServicePlanOptions

`func (o *AddClusterRequestClusterServer) GetServicePlanOptions() AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigServicePlanOptions`

GetServicePlanOptions returns the ServicePlanOptions field if non-nil, zero value otherwise.

### GetServicePlanOptionsOk

`func (o *AddClusterRequestClusterServer) GetServicePlanOptionsOk() (*AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigServicePlanOptions, bool)`

GetServicePlanOptionsOk returns a tuple with the ServicePlanOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanOptions

`func (o *AddClusterRequestClusterServer) SetServicePlanOptions(v AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigServicePlanOptions)`

SetServicePlanOptions sets ServicePlanOptions field to given value.

### HasServicePlanOptions

`func (o *AddClusterRequestClusterServer) HasServicePlanOptions() bool`

HasServicePlanOptions returns a boolean if a field has been set.

### GetVolumes

`func (o *AddClusterRequestClusterServer) GetVolumes() []AddClusterRequestClusterServerVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *AddClusterRequestClusterServer) GetVolumesOk() (*[]AddClusterRequestClusterServerVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *AddClusterRequestClusterServer) SetVolumes(v []AddClusterRequestClusterServerVolumesInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *AddClusterRequestClusterServer) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *AddClusterRequestClusterServer) GetNetworkInterfaces() []AddClusterRequestClusterServerNetworkInterfacesInner`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *AddClusterRequestClusterServer) GetNetworkInterfacesOk() (*[]AddClusterRequestClusterServerNetworkInterfacesInner, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *AddClusterRequestClusterServer) SetNetworkInterfaces(v []AddClusterRequestClusterServerNetworkInterfacesInner)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *AddClusterRequestClusterServer) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *AddClusterRequestClusterServer) GetSecurityGroups() []string`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *AddClusterRequestClusterServer) GetSecurityGroupsOk() (*[]string, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *AddClusterRequestClusterServer) SetSecurityGroups(v []string)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *AddClusterRequestClusterServer) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetVisibility

`func (o *AddClusterRequestClusterServer) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AddClusterRequestClusterServer) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AddClusterRequestClusterServer) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AddClusterRequestClusterServer) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetUserGroup

`func (o *AddClusterRequestClusterServer) GetUserGroup() AddClusterRequestClusterServerUserGroup`

GetUserGroup returns the UserGroup field if non-nil, zero value otherwise.

### GetUserGroupOk

`func (o *AddClusterRequestClusterServer) GetUserGroupOk() (*AddClusterRequestClusterServerUserGroup, bool)`

GetUserGroupOk returns a tuple with the UserGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserGroup

`func (o *AddClusterRequestClusterServer) SetUserGroup(v AddClusterRequestClusterServerUserGroup)`

SetUserGroup sets UserGroup field to given value.

### HasUserGroup

`func (o *AddClusterRequestClusterServer) HasUserGroup() bool`

HasUserGroup returns a boolean if a field has been set.

### GetNetworkDomain

`func (o *AddClusterRequestClusterServer) GetNetworkDomain() string`

GetNetworkDomain returns the NetworkDomain field if non-nil, zero value otherwise.

### GetNetworkDomainOk

`func (o *AddClusterRequestClusterServer) GetNetworkDomainOk() (*string, bool)`

GetNetworkDomainOk returns a tuple with the NetworkDomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDomain

`func (o *AddClusterRequestClusterServer) SetNetworkDomain(v string)`

SetNetworkDomain sets NetworkDomain field to given value.

### HasNetworkDomain

`func (o *AddClusterRequestClusterServer) HasNetworkDomain() bool`

HasNetworkDomain returns a boolean if a field has been set.

### GetHostname

`func (o *AddClusterRequestClusterServer) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *AddClusterRequestClusterServer) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *AddClusterRequestClusterServer) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *AddClusterRequestClusterServer) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetNodeCount

`func (o *AddClusterRequestClusterServer) GetNodeCount() int64`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *AddClusterRequestClusterServer) GetNodeCountOk() (*int64, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *AddClusterRequestClusterServer) SetNodeCount(v int64)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *AddClusterRequestClusterServer) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetTags

`func (o *AddClusterRequestClusterServer) GetTags() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AddClusterRequestClusterServer) GetTagsOk() (*[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AddClusterRequestClusterServer) SetTags(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AddClusterRequestClusterServer) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetLabels

`func (o *AddClusterRequestClusterServer) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddClusterRequestClusterServer) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddClusterRequestClusterServer) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddClusterRequestClusterServer) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetSshHosts

`func (o *AddClusterRequestClusterServer) GetSshHosts() []AddClusterRequestClusterServerSshHostsInner`

GetSshHosts returns the SshHosts field if non-nil, zero value otherwise.

### GetSshHostsOk

`func (o *AddClusterRequestClusterServer) GetSshHostsOk() (*[]AddClusterRequestClusterServerSshHostsInner, bool)`

GetSshHostsOk returns a tuple with the SshHosts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshHosts

`func (o *AddClusterRequestClusterServer) SetSshHosts(v []AddClusterRequestClusterServerSshHostsInner)`

SetSshHosts sets SshHosts field to given value.

### HasSshHosts

`func (o *AddClusterRequestClusterServer) HasSshHosts() bool`

HasSshHosts returns a boolean if a field has been set.

### GetSshUsername

`func (o *AddClusterRequestClusterServer) GetSshUsername() string`

GetSshUsername returns the SshUsername field if non-nil, zero value otherwise.

### GetSshUsernameOk

`func (o *AddClusterRequestClusterServer) GetSshUsernameOk() (*string, bool)`

GetSshUsernameOk returns a tuple with the SshUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshUsername

`func (o *AddClusterRequestClusterServer) SetSshUsername(v string)`

SetSshUsername sets SshUsername field to given value.

### HasSshUsername

`func (o *AddClusterRequestClusterServer) HasSshUsername() bool`

HasSshUsername returns a boolean if a field has been set.

### GetSshPassword

`func (o *AddClusterRequestClusterServer) GetSshPassword() string`

GetSshPassword returns the SshPassword field if non-nil, zero value otherwise.

### GetSshPasswordOk

`func (o *AddClusterRequestClusterServer) GetSshPasswordOk() (*string, bool)`

GetSshPasswordOk returns a tuple with the SshPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshPassword

`func (o *AddClusterRequestClusterServer) SetSshPassword(v string)`

SetSshPassword sets SshPassword field to given value.

### HasSshPassword

`func (o *AddClusterRequestClusterServer) HasSshPassword() bool`

HasSshPassword returns a boolean if a field has been set.

### GetSshKeyPair

`func (o *AddClusterRequestClusterServer) GetSshKeyPair() AddClusterRequestClusterServerSshKeyPair`

GetSshKeyPair returns the SshKeyPair field if non-nil, zero value otherwise.

### GetSshKeyPairOk

`func (o *AddClusterRequestClusterServer) GetSshKeyPairOk() (*AddClusterRequestClusterServerSshKeyPair, bool)`

GetSshKeyPairOk returns a tuple with the SshKeyPair field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSshKeyPair

`func (o *AddClusterRequestClusterServer) SetSshKeyPair(v AddClusterRequestClusterServerSshKeyPair)`

SetSshKeyPair sets SshKeyPair field to given value.

### HasSshKeyPair

`func (o *AddClusterRequestClusterServer) HasSshKeyPair() bool`

HasSshKeyPair returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


