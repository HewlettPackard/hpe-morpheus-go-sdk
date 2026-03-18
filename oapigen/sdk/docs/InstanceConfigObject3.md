# InstanceConfigObject3

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | [**InstanceConfigObject1Group**](InstanceConfigObject1Group.md) |  | 
**Cloud** | [**InstanceConfigObject1Cloud**](InstanceConfigObject1Cloud.md) |  | 
**Type** | **string** | The type of instance by code we want to fetch. | 
**Name** | **string** | Name of the instance to be created. | 
**Config** | [**InstanceConfigObject1Config**](InstanceConfigObject1Config.md) |  | 
**Volumes** | [**[]InstanceConfigObject1VolumesInner**](InstanceConfigObject1VolumesInner.md) | The (optional) volumes parameter is for LV configuration, can create additional LVs at provision It should be passed as an array of | 
**HostName** | Pointer to **string** | Hostname of the instance to be created.  Can be the same as the instance name. | [optional] 
**Environment** | Pointer to **string** | Environment code | [optional] 
**Layout** | [**InstanceConfigObject1Layout**](InstanceConfigObject1Layout.md) |  | 
**Plan** | [**InstanceConfigObject1Plan**](InstanceConfigObject1Plan.md) |  | 
**Version** | Pointer to **string** | Version of the layout to create. | [optional] 
**Evars** | Pointer to [**[]InstanceConfigObject1EvarsInner**](InstanceConfigObject1EvarsInner.md) | Environment Variables, an array of objects that have name and value. | [optional] 
**ServicePlanOptions** | Pointer to [**InstanceConfigObject1ServicePlanOptions**](InstanceConfigObject1ServicePlanOptions.md) |  | [optional] 
**SecurityGroups** | Pointer to [**[]InstanceConfigObject1SecurityGroupsInner**](InstanceConfigObject1SecurityGroupsInner.md) | Key for security group configuration. It should be passed as an array of objects containing the id of the security group to assign the instance to. | [optional] 
**NetworkInterfaces** | Pointer to [**[]InstancesNetworkInterfaces1**](InstancesNetworkInterfaces1.md) | The networkInterfaces parameter is for network configuration.  The Options API &#x60;/api/options/zoneNetworkOptions?zoneId&#x3D;5&amp;provisionTypeId&#x3D;10&#x60; can be used to see which options are available.  | [optional] 
**Labels** | Pointer to **[]string** | Array of strings (keywords). | [optional] 
**Tags** | Pointer to [**[]InstanceConfigObject1TagsInner**](InstanceConfigObject1TagsInner.md) | Metadata tags, Array of objects having a name and value. | [optional] 
**Metadata** | Pointer to [**[]InstanceConfigObject1MetadataInner**](InstanceConfigObject1MetadataInner.md) | Alias for &#x60;tags&#x60;. | [optional] 
**Ports** | Pointer to [**[]InstanceConfigObject1PortsInner**](InstanceConfigObject1PortsInner.md) | The ports parameter is for port configuration.  The layout may have default ports, which are defined in node types, that are always configured. This parameter will be for additional custom ports to be opened.  | [optional] 
**TaskSetId** | Pointer to **int64** | The Workflow ID to execute. | [optional] 
**TaskSetName** | Pointer to **string** | The Workflow Name to execute. | [optional] 

## Methods

### NewInstanceConfigObject3

`func NewInstanceConfigObject3(group InstanceConfigObject1Group, cloud InstanceConfigObject1Cloud, type_ string, name string, config InstanceConfigObject1Config, volumes []InstanceConfigObject1VolumesInner, layout InstanceConfigObject1Layout, plan InstanceConfigObject1Plan, ) *InstanceConfigObject3`

NewInstanceConfigObject3 instantiates a new InstanceConfigObject3 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceConfigObject3WithDefaults

`func NewInstanceConfigObject3WithDefaults() *InstanceConfigObject3`

NewInstanceConfigObject3WithDefaults instantiates a new InstanceConfigObject3 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *InstanceConfigObject3) GetGroup() InstanceConfigObject1Group`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *InstanceConfigObject3) GetGroupOk() (*InstanceConfigObject1Group, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *InstanceConfigObject3) SetGroup(v InstanceConfigObject1Group)`

SetGroup sets Group field to given value.


### GetCloud

`func (o *InstanceConfigObject3) GetCloud() InstanceConfigObject1Cloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *InstanceConfigObject3) GetCloudOk() (*InstanceConfigObject1Cloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *InstanceConfigObject3) SetCloud(v InstanceConfigObject1Cloud)`

SetCloud sets Cloud field to given value.


### GetType

`func (o *InstanceConfigObject3) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstanceConfigObject3) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstanceConfigObject3) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *InstanceConfigObject3) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceConfigObject3) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceConfigObject3) SetName(v string)`

SetName sets Name field to given value.


### GetConfig

`func (o *InstanceConfigObject3) GetConfig() InstanceConfigObject1Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *InstanceConfigObject3) GetConfigOk() (*InstanceConfigObject1Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *InstanceConfigObject3) SetConfig(v InstanceConfigObject1Config)`

SetConfig sets Config field to given value.


### GetVolumes

`func (o *InstanceConfigObject3) GetVolumes() []InstanceConfigObject1VolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *InstanceConfigObject3) GetVolumesOk() (*[]InstanceConfigObject1VolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *InstanceConfigObject3) SetVolumes(v []InstanceConfigObject1VolumesInner)`

SetVolumes sets Volumes field to given value.


### GetHostName

`func (o *InstanceConfigObject3) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *InstanceConfigObject3) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *InstanceConfigObject3) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *InstanceConfigObject3) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetEnvironment

`func (o *InstanceConfigObject3) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *InstanceConfigObject3) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *InstanceConfigObject3) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *InstanceConfigObject3) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetLayout

`func (o *InstanceConfigObject3) GetLayout() InstanceConfigObject1Layout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *InstanceConfigObject3) GetLayoutOk() (*InstanceConfigObject1Layout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *InstanceConfigObject3) SetLayout(v InstanceConfigObject1Layout)`

SetLayout sets Layout field to given value.


### GetPlan

`func (o *InstanceConfigObject3) GetPlan() InstanceConfigObject1Plan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *InstanceConfigObject3) GetPlanOk() (*InstanceConfigObject1Plan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *InstanceConfigObject3) SetPlan(v InstanceConfigObject1Plan)`

SetPlan sets Plan field to given value.


### GetVersion

`func (o *InstanceConfigObject3) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InstanceConfigObject3) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InstanceConfigObject3) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InstanceConfigObject3) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetEvars

`func (o *InstanceConfigObject3) GetEvars() []InstanceConfigObject1EvarsInner`

GetEvars returns the Evars field if non-nil, zero value otherwise.

### GetEvarsOk

`func (o *InstanceConfigObject3) GetEvarsOk() (*[]InstanceConfigObject1EvarsInner, bool)`

GetEvarsOk returns a tuple with the Evars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvars

`func (o *InstanceConfigObject3) SetEvars(v []InstanceConfigObject1EvarsInner)`

SetEvars sets Evars field to given value.

### HasEvars

`func (o *InstanceConfigObject3) HasEvars() bool`

HasEvars returns a boolean if a field has been set.

### GetServicePlanOptions

`func (o *InstanceConfigObject3) GetServicePlanOptions() InstanceConfigObject1ServicePlanOptions`

GetServicePlanOptions returns the ServicePlanOptions field if non-nil, zero value otherwise.

### GetServicePlanOptionsOk

`func (o *InstanceConfigObject3) GetServicePlanOptionsOk() (*InstanceConfigObject1ServicePlanOptions, bool)`

GetServicePlanOptionsOk returns a tuple with the ServicePlanOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanOptions

`func (o *InstanceConfigObject3) SetServicePlanOptions(v InstanceConfigObject1ServicePlanOptions)`

SetServicePlanOptions sets ServicePlanOptions field to given value.

### HasServicePlanOptions

`func (o *InstanceConfigObject3) HasServicePlanOptions() bool`

HasServicePlanOptions returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *InstanceConfigObject3) GetSecurityGroups() []InstanceConfigObject1SecurityGroupsInner`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *InstanceConfigObject3) GetSecurityGroupsOk() (*[]InstanceConfigObject1SecurityGroupsInner, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *InstanceConfigObject3) SetSecurityGroups(v []InstanceConfigObject1SecurityGroupsInner)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *InstanceConfigObject3) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *InstanceConfigObject3) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *InstanceConfigObject3) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil
### GetNetworkInterfaces

`func (o *InstanceConfigObject3) GetNetworkInterfaces() []InstancesNetworkInterfaces1`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *InstanceConfigObject3) GetNetworkInterfacesOk() (*[]InstancesNetworkInterfaces1, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *InstanceConfigObject3) SetNetworkInterfaces(v []InstancesNetworkInterfaces1)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *InstanceConfigObject3) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetLabels

`func (o *InstanceConfigObject3) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *InstanceConfigObject3) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *InstanceConfigObject3) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *InstanceConfigObject3) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTags

`func (o *InstanceConfigObject3) GetTags() []InstanceConfigObject1TagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InstanceConfigObject3) GetTagsOk() (*[]InstanceConfigObject1TagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InstanceConfigObject3) SetTags(v []InstanceConfigObject1TagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InstanceConfigObject3) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMetadata

`func (o *InstanceConfigObject3) GetMetadata() []InstanceConfigObject1MetadataInner`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InstanceConfigObject3) GetMetadataOk() (*[]InstanceConfigObject1MetadataInner, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InstanceConfigObject3) SetMetadata(v []InstanceConfigObject1MetadataInner)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InstanceConfigObject3) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPorts

`func (o *InstanceConfigObject3) GetPorts() []InstanceConfigObject1PortsInner`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *InstanceConfigObject3) GetPortsOk() (*[]InstanceConfigObject1PortsInner, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *InstanceConfigObject3) SetPorts(v []InstanceConfigObject1PortsInner)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *InstanceConfigObject3) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetTaskSetId

`func (o *InstanceConfigObject3) GetTaskSetId() int64`

GetTaskSetId returns the TaskSetId field if non-nil, zero value otherwise.

### GetTaskSetIdOk

`func (o *InstanceConfigObject3) GetTaskSetIdOk() (*int64, bool)`

GetTaskSetIdOk returns a tuple with the TaskSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSetId

`func (o *InstanceConfigObject3) SetTaskSetId(v int64)`

SetTaskSetId sets TaskSetId field to given value.

### HasTaskSetId

`func (o *InstanceConfigObject3) HasTaskSetId() bool`

HasTaskSetId returns a boolean if a field has been set.

### GetTaskSetName

`func (o *InstanceConfigObject3) GetTaskSetName() string`

GetTaskSetName returns the TaskSetName field if non-nil, zero value otherwise.

### GetTaskSetNameOk

`func (o *InstanceConfigObject3) GetTaskSetNameOk() (*string, bool)`

GetTaskSetNameOk returns a tuple with the TaskSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSetName

`func (o *InstanceConfigObject3) SetTaskSetName(v string)`

SetTaskSetName sets TaskSetName field to given value.

### HasTaskSetName

`func (o *InstanceConfigObject3) HasTaskSetName() bool`

HasTaskSetName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


