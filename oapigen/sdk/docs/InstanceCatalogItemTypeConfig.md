# InstanceCatalogItemTypeConfig

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

### NewInstanceCatalogItemTypeConfig

`func NewInstanceCatalogItemTypeConfig(group InstanceConfigObjectGroup, cloud InstanceConfigObjectCloud, type_ string, name string, config InstanceConfigObjectConfig, volumes []InstanceConfigObjectVolumesInner, layout InstanceConfigObjectLayout, plan InstanceConfigObjectPlan, ) *InstanceCatalogItemTypeConfig`

NewInstanceCatalogItemTypeConfig instantiates a new InstanceCatalogItemTypeConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceCatalogItemTypeConfigWithDefaults

`func NewInstanceCatalogItemTypeConfigWithDefaults() *InstanceCatalogItemTypeConfig`

NewInstanceCatalogItemTypeConfigWithDefaults instantiates a new InstanceCatalogItemTypeConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *InstanceCatalogItemTypeConfig) GetGroup() InstanceConfigObjectGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *InstanceCatalogItemTypeConfig) GetGroupOk() (*InstanceConfigObjectGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *InstanceCatalogItemTypeConfig) SetGroup(v InstanceConfigObjectGroup)`

SetGroup sets Group field to given value.


### GetCloud

`func (o *InstanceCatalogItemTypeConfig) GetCloud() InstanceConfigObjectCloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *InstanceCatalogItemTypeConfig) GetCloudOk() (*InstanceConfigObjectCloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *InstanceCatalogItemTypeConfig) SetCloud(v InstanceConfigObjectCloud)`

SetCloud sets Cloud field to given value.


### GetType

`func (o *InstanceCatalogItemTypeConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InstanceCatalogItemTypeConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InstanceCatalogItemTypeConfig) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *InstanceCatalogItemTypeConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceCatalogItemTypeConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceCatalogItemTypeConfig) SetName(v string)`

SetName sets Name field to given value.


### GetConfig

`func (o *InstanceCatalogItemTypeConfig) GetConfig() InstanceConfigObjectConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *InstanceCatalogItemTypeConfig) GetConfigOk() (*InstanceConfigObjectConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *InstanceCatalogItemTypeConfig) SetConfig(v InstanceConfigObjectConfig)`

SetConfig sets Config field to given value.


### GetVolumes

`func (o *InstanceCatalogItemTypeConfig) GetVolumes() []InstanceConfigObjectVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *InstanceCatalogItemTypeConfig) GetVolumesOk() (*[]InstanceConfigObjectVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *InstanceCatalogItemTypeConfig) SetVolumes(v []InstanceConfigObjectVolumesInner)`

SetVolumes sets Volumes field to given value.


### GetHostName

`func (o *InstanceCatalogItemTypeConfig) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *InstanceCatalogItemTypeConfig) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *InstanceCatalogItemTypeConfig) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *InstanceCatalogItemTypeConfig) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetEnvironment

`func (o *InstanceCatalogItemTypeConfig) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *InstanceCatalogItemTypeConfig) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *InstanceCatalogItemTypeConfig) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *InstanceCatalogItemTypeConfig) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetLayout

`func (o *InstanceCatalogItemTypeConfig) GetLayout() InstanceConfigObjectLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *InstanceCatalogItemTypeConfig) GetLayoutOk() (*InstanceConfigObjectLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *InstanceCatalogItemTypeConfig) SetLayout(v InstanceConfigObjectLayout)`

SetLayout sets Layout field to given value.


### GetPlan

`func (o *InstanceCatalogItemTypeConfig) GetPlan() InstanceConfigObjectPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *InstanceCatalogItemTypeConfig) GetPlanOk() (*InstanceConfigObjectPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *InstanceCatalogItemTypeConfig) SetPlan(v InstanceConfigObjectPlan)`

SetPlan sets Plan field to given value.


### GetVersion

`func (o *InstanceCatalogItemTypeConfig) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InstanceCatalogItemTypeConfig) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InstanceCatalogItemTypeConfig) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InstanceCatalogItemTypeConfig) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetEvars

`func (o *InstanceCatalogItemTypeConfig) GetEvars() []InstanceConfigObjectEvarsInner`

GetEvars returns the Evars field if non-nil, zero value otherwise.

### GetEvarsOk

`func (o *InstanceCatalogItemTypeConfig) GetEvarsOk() (*[]InstanceConfigObjectEvarsInner, bool)`

GetEvarsOk returns a tuple with the Evars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvars

`func (o *InstanceCatalogItemTypeConfig) SetEvars(v []InstanceConfigObjectEvarsInner)`

SetEvars sets Evars field to given value.

### HasEvars

`func (o *InstanceCatalogItemTypeConfig) HasEvars() bool`

HasEvars returns a boolean if a field has been set.

### GetServicePlanOptions

`func (o *InstanceCatalogItemTypeConfig) GetServicePlanOptions() InstanceConfigObjectServicePlanOptions`

GetServicePlanOptions returns the ServicePlanOptions field if non-nil, zero value otherwise.

### GetServicePlanOptionsOk

`func (o *InstanceCatalogItemTypeConfig) GetServicePlanOptionsOk() (*InstanceConfigObjectServicePlanOptions, bool)`

GetServicePlanOptionsOk returns a tuple with the ServicePlanOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanOptions

`func (o *InstanceCatalogItemTypeConfig) SetServicePlanOptions(v InstanceConfigObjectServicePlanOptions)`

SetServicePlanOptions sets ServicePlanOptions field to given value.

### HasServicePlanOptions

`func (o *InstanceCatalogItemTypeConfig) HasServicePlanOptions() bool`

HasServicePlanOptions returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *InstanceCatalogItemTypeConfig) GetSecurityGroups() []InstanceConfigObjectSecurityGroupsInner`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *InstanceCatalogItemTypeConfig) GetSecurityGroupsOk() (*[]InstanceConfigObjectSecurityGroupsInner, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *InstanceCatalogItemTypeConfig) SetSecurityGroups(v []InstanceConfigObjectSecurityGroupsInner)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *InstanceCatalogItemTypeConfig) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *InstanceCatalogItemTypeConfig) GetNetworkInterfaces() []InstancesNetworkInterfaces`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *InstanceCatalogItemTypeConfig) GetNetworkInterfacesOk() (*[]InstancesNetworkInterfaces, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *InstanceCatalogItemTypeConfig) SetNetworkInterfaces(v []InstancesNetworkInterfaces)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *InstanceCatalogItemTypeConfig) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetLabels

`func (o *InstanceCatalogItemTypeConfig) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *InstanceCatalogItemTypeConfig) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *InstanceCatalogItemTypeConfig) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *InstanceCatalogItemTypeConfig) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTags

`func (o *InstanceCatalogItemTypeConfig) GetTags() []InstanceConfigObjectTagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *InstanceCatalogItemTypeConfig) GetTagsOk() (*[]InstanceConfigObjectTagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *InstanceCatalogItemTypeConfig) SetTags(v []InstanceConfigObjectTagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *InstanceCatalogItemTypeConfig) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMetadata

`func (o *InstanceCatalogItemTypeConfig) GetMetadata() []InstanceConfigObjectMetadataInner`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *InstanceCatalogItemTypeConfig) GetMetadataOk() (*[]InstanceConfigObjectMetadataInner, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *InstanceCatalogItemTypeConfig) SetMetadata(v []InstanceConfigObjectMetadataInner)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *InstanceCatalogItemTypeConfig) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPorts

`func (o *InstanceCatalogItemTypeConfig) GetPorts() []InstanceConfigObjectPortsInner`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *InstanceCatalogItemTypeConfig) GetPortsOk() (*[]InstanceConfigObjectPortsInner, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *InstanceCatalogItemTypeConfig) SetPorts(v []InstanceConfigObjectPortsInner)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *InstanceCatalogItemTypeConfig) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetTaskSetId

`func (o *InstanceCatalogItemTypeConfig) GetTaskSetId() int64`

GetTaskSetId returns the TaskSetId field if non-nil, zero value otherwise.

### GetTaskSetIdOk

`func (o *InstanceCatalogItemTypeConfig) GetTaskSetIdOk() (*int64, bool)`

GetTaskSetIdOk returns a tuple with the TaskSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSetId

`func (o *InstanceCatalogItemTypeConfig) SetTaskSetId(v int64)`

SetTaskSetId sets TaskSetId field to given value.

### HasTaskSetId

`func (o *InstanceCatalogItemTypeConfig) HasTaskSetId() bool`

HasTaskSetId returns a boolean if a field has been set.

### GetTaskSetName

`func (o *InstanceCatalogItemTypeConfig) GetTaskSetName() string`

GetTaskSetName returns the TaskSetName field if non-nil, zero value otherwise.

### GetTaskSetNameOk

`func (o *InstanceCatalogItemTypeConfig) GetTaskSetNameOk() (*string, bool)`

GetTaskSetNameOk returns a tuple with the TaskSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSetName

`func (o *InstanceCatalogItemTypeConfig) SetTaskSetName(v string)`

SetTaskSetName sets TaskSetName field to given value.

### HasTaskSetName

`func (o *InstanceCatalogItemTypeConfig) HasTaskSetName() bool`

HasTaskSetName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


