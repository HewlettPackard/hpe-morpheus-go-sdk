# InstanceConfigObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | [**InstanceConfigObjectGroup**](InstanceConfigObjectGroup.md) |  | 
**Cloud** | [**InstanceConfigObjectCloud**](InstanceConfigObjectCloud.md) |  | 
**Type** | **string** | The type of instance by code we want to fetch. | 
**Name** | **string** | Name of the instance to be created. | 
**Config** | [**InstanceConfigObjectConfig**](InstanceConfigObjectConfig.md) |  | 
**Volumes** | [**[]InstanceConfigObjectVolumesInner**](InstanceConfigObjectVolumesInner.md) | The (optional) volumes parameter is for LV configuration, can create additional LVs at provision It should be passed as an array of | 
**HostName** | Pointer to **string** | Hostname of the instance to be created.  Can be the same as the instance name. | [optional] 
**Environment** | Pointer to **string** | Environment code | [optional] 
**Layout** | [**InstanceConfigObjectLayout**](InstanceConfigObjectLayout.md) |  | 
**Plan** | [**InstanceConfigObjectPlan**](InstanceConfigObjectPlan.md) |  | 
**Version** | Pointer to **string** | Version of the layout to create. | [optional] 
**Evars** | Pointer to [**[]InstanceConfigObjectEvarsInner**](InstanceConfigObjectEvarsInner.md) | Environment Variables, an array of objects that have name and value. | [optional] 
**ServicePlanOptions** | Pointer to [**InstanceConfigObjectServicePlanOptions**](InstanceConfigObjectServicePlanOptions.md) |  | [optional] 
**SecurityGroups** | Pointer to [**[]InstanceConfigObjectSecurityGroupsInner**](InstanceConfigObjectSecurityGroupsInner.md) | Key for security group configuration. It should be passed as an array of objects containing the id of the security group to assign the instance to. | [optional] 
**NetworkInterfaces** | Pointer to [**[]InstancesNetworkInterfaces**](InstancesNetworkInterfaces.md) | The networkInterfaces parameter is for network configuration.  The Options API &#x60;/api/options/zoneNetworkOptions?zoneId&#x3D;5&amp;provisionTypeId&#x3D;10&#x60; can be used to see which options are available.  | [optional] 
**Labels** | Pointer to **[]string** | Array of strings (keywords). | [optional] 
**Tags** | Pointer to [**[]InstanceConfigObjectTagsInner**](InstanceConfigObjectTagsInner.md) | Metadata tags, Array of objects having a name and value. | [optional] 
**Metadata** | Pointer to [**[]InstanceConfigObjectMetadataInner**](InstanceConfigObjectMetadataInner.md) | Alias for &#x60;tags&#x60;. | [optional] 
**Ports** | Pointer to [**[]InstanceConfigObjectPortsInner**](InstanceConfigObjectPortsInner.md) | The ports parameter is for port configuration.  The layout may have default ports, which are defined in node types, that are always configured. This parameter will be for additional custom ports to be opened.  | [optional] 
**TaskSetId** | Pointer to **int64** | The Workflow ID to execute. | [optional] 
**TaskSetName** | Pointer to **string** | The Workflow Name to execute. | [optional] 

## Methods

### NewInstanceConfigObject

`func NewInstanceConfigObject(group InstanceConfigObjectGroup, cloud InstanceConfigObjectCloud, type_ string, name string, config InstanceConfigObjectConfig, volumes []InstanceConfigObjectVolumesInner, layout InstanceConfigObjectLayout, plan InstanceConfigObjectPlan, ) *InstanceConfigObject`

NewInstanceConfigObject instantiates a new InstanceConfigObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceConfigObjectWithDefaults

`func NewInstanceConfigObjectWithDefaults() *InstanceConfigObject`

NewInstanceConfigObjectWithDefaults instantiates a new InstanceConfigObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *InstanceConfigObject) GetGroup() InstanceConfigObjectGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *InstanceConfigObject) GetGroupOk() (*InstanceConfigObjectGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *InstanceConfigObject) SetGroup(v InstanceConfigObjectGroup)`

SetGroup sets Group field to given value.


### GetCloud

`func (o *InstanceConfigObject) GetCloud() InstanceConfigObjectCloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *InstanceConfigObject) GetCloudOk() (*InstanceConfigObjectCloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *InstanceConfigObject) SetCloud(v InstanceConfigObjectCloud)`

SetCloud sets Cloud field to given value.


### GetType

`func (o *InstanceConfigObject) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstanceConfigObject) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstanceConfigObject) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *InstanceConfigObject) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceConfigObject) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceConfigObject) SetName(v string)`

SetName sets Name field to given value.


### GetConfig

`func (o *InstanceConfigObject) GetConfig() InstanceConfigObjectConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *InstanceConfigObject) GetConfigOk() (*InstanceConfigObjectConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *InstanceConfigObject) SetConfig(v InstanceConfigObjectConfig)`

SetConfig sets Config field to given value.


### GetVolumes

`func (o *InstanceConfigObject) GetVolumes() []InstanceConfigObjectVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *InstanceConfigObject) GetVolumesOk() (*[]InstanceConfigObjectVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *InstanceConfigObject) SetVolumes(v []InstanceConfigObjectVolumesInner)`

SetVolumes sets Volumes field to given value.


### GetHostName

`func (o *InstanceConfigObject) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *InstanceConfigObject) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *InstanceConfigObject) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *InstanceConfigObject) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetEnvironment

`func (o *InstanceConfigObject) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *InstanceConfigObject) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *InstanceConfigObject) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *InstanceConfigObject) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetLayout

`func (o *InstanceConfigObject) GetLayout() InstanceConfigObjectLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *InstanceConfigObject) GetLayoutOk() (*InstanceConfigObjectLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *InstanceConfigObject) SetLayout(v InstanceConfigObjectLayout)`

SetLayout sets Layout field to given value.


### GetPlan

`func (o *InstanceConfigObject) GetPlan() InstanceConfigObjectPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *InstanceConfigObject) GetPlanOk() (*InstanceConfigObjectPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *InstanceConfigObject) SetPlan(v InstanceConfigObjectPlan)`

SetPlan sets Plan field to given value.


### GetVersion

`func (o *InstanceConfigObject) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InstanceConfigObject) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InstanceConfigObject) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InstanceConfigObject) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetEvars

`func (o *InstanceConfigObject) GetEvars() []InstanceConfigObjectEvarsInner`

GetEvars returns the Evars field if non-nil, zero value otherwise.

### GetEvarsOk

`func (o *InstanceConfigObject) GetEvarsOk() (*[]InstanceConfigObjectEvarsInner, bool)`

GetEvarsOk returns a tuple with the Evars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvars

`func (o *InstanceConfigObject) SetEvars(v []InstanceConfigObjectEvarsInner)`

SetEvars sets Evars field to given value.

### HasEvars

`func (o *InstanceConfigObject) HasEvars() bool`

HasEvars returns a boolean if a field has been set.

### GetServicePlanOptions

`func (o *InstanceConfigObject) GetServicePlanOptions() InstanceConfigObjectServicePlanOptions`

GetServicePlanOptions returns the ServicePlanOptions field if non-nil, zero value otherwise.

### GetServicePlanOptionsOk

`func (o *InstanceConfigObject) GetServicePlanOptionsOk() (*InstanceConfigObjectServicePlanOptions, bool)`

GetServicePlanOptionsOk returns a tuple with the ServicePlanOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanOptions

`func (o *InstanceConfigObject) SetServicePlanOptions(v InstanceConfigObjectServicePlanOptions)`

SetServicePlanOptions sets ServicePlanOptions field to given value.

### HasServicePlanOptions

`func (o *InstanceConfigObject) HasServicePlanOptions() bool`

HasServicePlanOptions returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *InstanceConfigObject) GetSecurityGroups() []InstanceConfigObjectSecurityGroupsInner`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *InstanceConfigObject) GetSecurityGroupsOk() (*[]InstanceConfigObjectSecurityGroupsInner, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *InstanceConfigObject) SetSecurityGroups(v []InstanceConfigObjectSecurityGroupsInner)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *InstanceConfigObject) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### SetSecurityGroupsNil

`func (o *InstanceConfigObject) SetSecurityGroupsNil(b bool)`

 SetSecurityGroupsNil sets the value for SecurityGroups to be an explicit nil

### UnsetSecurityGroups
`func (o *InstanceConfigObject) UnsetSecurityGroups()`

UnsetSecurityGroups ensures that no value is present for SecurityGroups, not even an explicit nil
### GetNetworkInterfaces

`func (o *InstanceConfigObject) GetNetworkInterfaces() []InstancesNetworkInterfaces`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *InstanceConfigObject) GetNetworkInterfacesOk() (*[]InstancesNetworkInterfaces, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *InstanceConfigObject) SetNetworkInterfaces(v []InstancesNetworkInterfaces)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *InstanceConfigObject) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetLabels

`func (o *InstanceConfigObject) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *InstanceConfigObject) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *InstanceConfigObject) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *InstanceConfigObject) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTags

`func (o *InstanceConfigObject) GetTags() []InstanceConfigObjectTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InstanceConfigObject) GetTagsOk() (*[]InstanceConfigObjectTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InstanceConfigObject) SetTags(v []InstanceConfigObjectTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InstanceConfigObject) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMetadata

`func (o *InstanceConfigObject) GetMetadata() []InstanceConfigObjectMetadataInner`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InstanceConfigObject) GetMetadataOk() (*[]InstanceConfigObjectMetadataInner, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InstanceConfigObject) SetMetadata(v []InstanceConfigObjectMetadataInner)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InstanceConfigObject) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPorts

`func (o *InstanceConfigObject) GetPorts() []InstanceConfigObjectPortsInner`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *InstanceConfigObject) GetPortsOk() (*[]InstanceConfigObjectPortsInner, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *InstanceConfigObject) SetPorts(v []InstanceConfigObjectPortsInner)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *InstanceConfigObject) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetTaskSetId

`func (o *InstanceConfigObject) GetTaskSetId() int64`

GetTaskSetId returns the TaskSetId field if non-nil, zero value otherwise.

### GetTaskSetIdOk

`func (o *InstanceConfigObject) GetTaskSetIdOk() (*int64, bool)`

GetTaskSetIdOk returns a tuple with the TaskSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSetId

`func (o *InstanceConfigObject) SetTaskSetId(v int64)`

SetTaskSetId sets TaskSetId field to given value.

### HasTaskSetId

`func (o *InstanceConfigObject) HasTaskSetId() bool`

HasTaskSetId returns a boolean if a field has been set.

### GetTaskSetName

`func (o *InstanceConfigObject) GetTaskSetName() string`

GetTaskSetName returns the TaskSetName field if non-nil, zero value otherwise.

### GetTaskSetNameOk

`func (o *InstanceConfigObject) GetTaskSetNameOk() (*string, bool)`

GetTaskSetNameOk returns a tuple with the TaskSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSetName

`func (o *InstanceConfigObject) SetTaskSetName(v string)`

SetTaskSetName sets TaskSetName field to given value.

### HasTaskSetName

`func (o *InstanceConfigObject) HasTaskSetName() bool`

HasTaskSetName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


