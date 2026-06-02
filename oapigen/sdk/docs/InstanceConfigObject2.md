# InstanceConfigObject2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | [**InstanceConfigObject2Group**](InstanceConfigObject2Group.md) |  | 
**Cloud** | [**InstanceConfigObject2Cloud**](InstanceConfigObject2Cloud.md) |  | 
**Type** | **string** | The type of instance by code we want to fetch. | 
**Name** | **string** | Name of the instance to be created. | 
**Config** | [**InstanceConfigObject2Config**](InstanceConfigObject2Config.md) |  | 
**Volumes** | [**[]InstanceConfigObject2VolumesInner**](InstanceConfigObject2VolumesInner.md) | The (optional) volumes parameter is for LV configuration, can create additional LVs at provision It should be passed as an array of | 
**HostName** | Pointer to **string** | Hostname of the instance to be created.  Can be the same as the instance name. | [optional] 
**Environment** | Pointer to **string** | Environment code | [optional] 
**Layout** | [**InstanceConfigObject2Layout**](InstanceConfigObject2Layout.md) |  | 
**Plan** | [**InstanceConfigObject2Plan**](InstanceConfigObject2Plan.md) |  | 
**Version** | Pointer to **string** | Version of the layout to create. | [optional] 
**Evars** | Pointer to [**[]InstanceConfigObject2EvarsInner**](InstanceConfigObject2EvarsInner.md) | Environment Variables, an array of objects that have name and value. | [optional] 
**ServicePlanOptions** | Pointer to [**InstanceConfigObject2ServicePlanOptions**](InstanceConfigObject2ServicePlanOptions.md) |  | [optional] 
**SecurityGroups** | Pointer to [**[]InstanceConfigObject2SecurityGroupsInner**](InstanceConfigObject2SecurityGroupsInner.md) | Key for security group configuration. It should be passed as an array of objects containing the id of the security group to assign the instance to. | [optional] 
**NetworkInterfaces** | Pointer to [**[]InstancesNetworkInterfaces5**](InstancesNetworkInterfaces5.md) | The networkInterfaces parameter is for network configuration.  The Options API &#x60;/api/options/zoneNetworkOptions?zoneId&#x3D;5&amp;provisionTypeId&#x3D;10&#x60; can be used to see which options are available.  | [optional] 
**Labels** | Pointer to **[]string** | Array of strings (keywords). | [optional] 
**Tags** | Pointer to [**[]InstanceConfigObject2TagsInner**](InstanceConfigObject2TagsInner.md) | Metadata tags, Array of objects having a name and value. | [optional] 
**Metadata** | Pointer to [**[]InstanceConfigObject2MetadataInner**](InstanceConfigObject2MetadataInner.md) | Alias for &#x60;tags&#x60;. | [optional] 
**Ports** | Pointer to [**[]InstanceConfigObject2PortsInner**](InstanceConfigObject2PortsInner.md) | The ports parameter is for port configuration.  The layout may have default ports, which are defined in node types, that are always configured. This parameter will be for additional custom ports to be opened.  | [optional] 
**TaskSetId** | Pointer to **int64** | The Workflow ID to execute. | [optional] 
**TaskSetName** | Pointer to **string** | The Workflow Name to execute. | [optional] 

## Methods

### NewInstanceConfigObject2

`func NewInstanceConfigObject2(group InstanceConfigObject2Group, cloud InstanceConfigObject2Cloud, type_ string, name string, config InstanceConfigObject2Config, volumes []InstanceConfigObject2VolumesInner, layout InstanceConfigObject2Layout, plan InstanceConfigObject2Plan, ) *InstanceConfigObject2`

NewInstanceConfigObject2 instantiates a new InstanceConfigObject2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetGroup

`func (o *InstanceConfigObject2) GetGroup() InstanceConfigObject2Group`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *InstanceConfigObject2) GetGroupOk() (*InstanceConfigObject2Group, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *InstanceConfigObject2) SetGroup(v InstanceConfigObject2Group)`

SetGroup sets Group field to given value.


### GetCloud

`func (o *InstanceConfigObject2) GetCloud() InstanceConfigObject2Cloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *InstanceConfigObject2) GetCloudOk() (*InstanceConfigObject2Cloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *InstanceConfigObject2) SetCloud(v InstanceConfigObject2Cloud)`

SetCloud sets Cloud field to given value.


### GetType

`func (o *InstanceConfigObject2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstanceConfigObject2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstanceConfigObject2) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *InstanceConfigObject2) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceConfigObject2) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceConfigObject2) SetName(v string)`

SetName sets Name field to given value.


### GetConfig

`func (o *InstanceConfigObject2) GetConfig() InstanceConfigObject2Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *InstanceConfigObject2) GetConfigOk() (*InstanceConfigObject2Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *InstanceConfigObject2) SetConfig(v InstanceConfigObject2Config)`

SetConfig sets Config field to given value.


### GetVolumes

`func (o *InstanceConfigObject2) GetVolumes() []InstanceConfigObject2VolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *InstanceConfigObject2) GetVolumesOk() (*[]InstanceConfigObject2VolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *InstanceConfigObject2) SetVolumes(v []InstanceConfigObject2VolumesInner)`

SetVolumes sets Volumes field to given value.


### GetHostName

`func (o *InstanceConfigObject2) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *InstanceConfigObject2) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *InstanceConfigObject2) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *InstanceConfigObject2) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetEnvironment

`func (o *InstanceConfigObject2) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *InstanceConfigObject2) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *InstanceConfigObject2) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *InstanceConfigObject2) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetLayout

`func (o *InstanceConfigObject2) GetLayout() InstanceConfigObject2Layout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *InstanceConfigObject2) GetLayoutOk() (*InstanceConfigObject2Layout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *InstanceConfigObject2) SetLayout(v InstanceConfigObject2Layout)`

SetLayout sets Layout field to given value.


### GetPlan

`func (o *InstanceConfigObject2) GetPlan() InstanceConfigObject2Plan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *InstanceConfigObject2) GetPlanOk() (*InstanceConfigObject2Plan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *InstanceConfigObject2) SetPlan(v InstanceConfigObject2Plan)`

SetPlan sets Plan field to given value.


### GetVersion

`func (o *InstanceConfigObject2) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InstanceConfigObject2) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InstanceConfigObject2) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InstanceConfigObject2) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetEvars

`func (o *InstanceConfigObject2) GetEvars() []InstanceConfigObject2EvarsInner`

GetEvars returns the Evars field if non-nil, zero value otherwise.

### GetEvarsOk

`func (o *InstanceConfigObject2) GetEvarsOk() (*[]InstanceConfigObject2EvarsInner, bool)`

GetEvarsOk returns a tuple with the Evars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvars

`func (o *InstanceConfigObject2) SetEvars(v []InstanceConfigObject2EvarsInner)`

SetEvars sets Evars field to given value.

### HasEvars

`func (o *InstanceConfigObject2) HasEvars() bool`

HasEvars returns a boolean if a field has been set.

### GetServicePlanOptions

`func (o *InstanceConfigObject2) GetServicePlanOptions() InstanceConfigObject2ServicePlanOptions`

GetServicePlanOptions returns the ServicePlanOptions field if non-nil, zero value otherwise.

### GetServicePlanOptionsOk

`func (o *InstanceConfigObject2) GetServicePlanOptionsOk() (*InstanceConfigObject2ServicePlanOptions, bool)`

GetServicePlanOptionsOk returns a tuple with the ServicePlanOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanOptions

`func (o *InstanceConfigObject2) SetServicePlanOptions(v InstanceConfigObject2ServicePlanOptions)`

SetServicePlanOptions sets ServicePlanOptions field to given value.

### HasServicePlanOptions

`func (o *InstanceConfigObject2) HasServicePlanOptions() bool`

HasServicePlanOptions returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *InstanceConfigObject2) GetSecurityGroups() []InstanceConfigObject2SecurityGroupsInner`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *InstanceConfigObject2) GetSecurityGroupsOk() (*[]InstanceConfigObject2SecurityGroupsInner, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *InstanceConfigObject2) SetSecurityGroups(v []InstanceConfigObject2SecurityGroupsInner)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *InstanceConfigObject2) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *InstanceConfigObject2) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *InstanceConfigObject2) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil
### GetNetworkInterfaces

`func (o *InstanceConfigObject2) GetNetworkInterfaces() []InstancesNetworkInterfaces5`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *InstanceConfigObject2) GetNetworkInterfacesOk() (*[]InstancesNetworkInterfaces5, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *InstanceConfigObject2) SetNetworkInterfaces(v []InstancesNetworkInterfaces5)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *InstanceConfigObject2) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetLabels

`func (o *InstanceConfigObject2) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *InstanceConfigObject2) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *InstanceConfigObject2) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *InstanceConfigObject2) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTags

`func (o *InstanceConfigObject2) GetTags() []InstanceConfigObject2TagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InstanceConfigObject2) GetTagsOk() (*[]InstanceConfigObject2TagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InstanceConfigObject2) SetTags(v []InstanceConfigObject2TagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InstanceConfigObject2) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMetadata

`func (o *InstanceConfigObject2) GetMetadata() []InstanceConfigObject2MetadataInner`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InstanceConfigObject2) GetMetadataOk() (*[]InstanceConfigObject2MetadataInner, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InstanceConfigObject2) SetMetadata(v []InstanceConfigObject2MetadataInner)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InstanceConfigObject2) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPorts

`func (o *InstanceConfigObject2) GetPorts() []InstanceConfigObject2PortsInner`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *InstanceConfigObject2) GetPortsOk() (*[]InstanceConfigObject2PortsInner, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *InstanceConfigObject2) SetPorts(v []InstanceConfigObject2PortsInner)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *InstanceConfigObject2) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetTaskSetId

`func (o *InstanceConfigObject2) GetTaskSetId() int64`

GetTaskSetId returns the TaskSetId field if non-nil, zero value otherwise.

### GetTaskSetIdOk

`func (o *InstanceConfigObject2) GetTaskSetIdOk() (*int64, bool)`

GetTaskSetIdOk returns a tuple with the TaskSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSetId

`func (o *InstanceConfigObject2) SetTaskSetId(v int64)`

SetTaskSetId sets TaskSetId field to given value.

### HasTaskSetId

`func (o *InstanceConfigObject2) HasTaskSetId() bool`

HasTaskSetId returns a boolean if a field has been set.

### GetTaskSetName

`func (o *InstanceConfigObject2) GetTaskSetName() string`

GetTaskSetName returns the TaskSetName field if non-nil, zero value otherwise.

### GetTaskSetNameOk

`func (o *InstanceConfigObject2) GetTaskSetNameOk() (*string, bool)`

GetTaskSetNameOk returns a tuple with the TaskSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSetName

`func (o *InstanceConfigObject2) SetTaskSetName(v string)`

SetTaskSetName sets TaskSetName field to given value.

### HasTaskSetName

`func (o *InstanceConfigObject2) HasTaskSetName() bool`

HasTaskSetName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


