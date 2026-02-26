# CatalogItemTypeInstanceCreateConfig

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
**NetworkInterfaces** | Pointer to [**[]InstancesNetworkInterfaces4**](InstancesNetworkInterfaces4.md) | The networkInterfaces parameter is for network configuration.  The Options API &#x60;/api/options/zoneNetworkOptions?zoneId&#x3D;5&amp;provisionTypeId&#x3D;10&#x60; can be used to see which options are available.  | [optional] 
**Labels** | Pointer to **[]string** | Array of strings (keywords). | [optional] 
**Tags** | Pointer to [**[]InstanceConfigObject1TagsInner**](InstanceConfigObject1TagsInner.md) | Metadata tags, Array of objects having a name and value. | [optional] 
**Metadata** | Pointer to [**[]InstanceConfigObject1MetadataInner**](InstanceConfigObject1MetadataInner.md) | Alias for &#x60;tags&#x60;. | [optional] 
**Ports** | Pointer to [**[]InstanceConfigObject1PortsInner**](InstanceConfigObject1PortsInner.md) | The ports parameter is for port configuration.  The layout may have default ports, which are defined in node types, that are always configured. This parameter will be for additional custom ports to be opened.  | [optional] 
**TaskSetId** | Pointer to **int64** | The Workflow ID to execute. | [optional] 
**TaskSetName** | Pointer to **string** | The Workflow Name to execute. | [optional] 

## Methods

### NewCatalogItemTypeInstanceCreateConfig

`func NewCatalogItemTypeInstanceCreateConfig(group InstanceConfigObject1Group, cloud InstanceConfigObject1Cloud, type_ string, name string, config InstanceConfigObject1Config, volumes []InstanceConfigObject1VolumesInner, layout InstanceConfigObject1Layout, plan InstanceConfigObject1Plan, ) *CatalogItemTypeInstanceCreateConfig`

NewCatalogItemTypeInstanceCreateConfig instantiates a new CatalogItemTypeInstanceCreateConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogItemTypeInstanceCreateConfigWithDefaults

`func NewCatalogItemTypeInstanceCreateConfigWithDefaults() *CatalogItemTypeInstanceCreateConfig`

NewCatalogItemTypeInstanceCreateConfigWithDefaults instantiates a new CatalogItemTypeInstanceCreateConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *CatalogItemTypeInstanceCreateConfig) GetGroup() InstanceConfigObject1Group`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *CatalogItemTypeInstanceCreateConfig) GetGroupOk() (*InstanceConfigObject1Group, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *CatalogItemTypeInstanceCreateConfig) SetGroup(v InstanceConfigObject1Group)`

SetGroup sets Group field to given value.


### GetCloud

`func (o *CatalogItemTypeInstanceCreateConfig) GetCloud() InstanceConfigObject1Cloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *CatalogItemTypeInstanceCreateConfig) GetCloudOk() (*InstanceConfigObject1Cloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *CatalogItemTypeInstanceCreateConfig) SetCloud(v InstanceConfigObject1Cloud)`

SetCloud sets Cloud field to given value.


### GetType

`func (o *CatalogItemTypeInstanceCreateConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogItemTypeInstanceCreateConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogItemTypeInstanceCreateConfig) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *CatalogItemTypeInstanceCreateConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogItemTypeInstanceCreateConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogItemTypeInstanceCreateConfig) SetName(v string)`

SetName sets Name field to given value.


### GetConfig

`func (o *CatalogItemTypeInstanceCreateConfig) GetConfig() InstanceConfigObject1Config`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CatalogItemTypeInstanceCreateConfig) GetConfigOk() (*InstanceConfigObject1Config, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CatalogItemTypeInstanceCreateConfig) SetConfig(v InstanceConfigObject1Config)`

SetConfig sets Config field to given value.


### GetVolumes

`func (o *CatalogItemTypeInstanceCreateConfig) GetVolumes() []InstanceConfigObject1VolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *CatalogItemTypeInstanceCreateConfig) GetVolumesOk() (*[]InstanceConfigObject1VolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *CatalogItemTypeInstanceCreateConfig) SetVolumes(v []InstanceConfigObject1VolumesInner)`

SetVolumes sets Volumes field to given value.


### GetHostName

`func (o *CatalogItemTypeInstanceCreateConfig) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *CatalogItemTypeInstanceCreateConfig) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *CatalogItemTypeInstanceCreateConfig) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *CatalogItemTypeInstanceCreateConfig) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetEnvironment

`func (o *CatalogItemTypeInstanceCreateConfig) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CatalogItemTypeInstanceCreateConfig) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CatalogItemTypeInstanceCreateConfig) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CatalogItemTypeInstanceCreateConfig) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetLayout

`func (o *CatalogItemTypeInstanceCreateConfig) GetLayout() InstanceConfigObject1Layout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *CatalogItemTypeInstanceCreateConfig) GetLayoutOk() (*InstanceConfigObject1Layout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *CatalogItemTypeInstanceCreateConfig) SetLayout(v InstanceConfigObject1Layout)`

SetLayout sets Layout field to given value.


### GetPlan

`func (o *CatalogItemTypeInstanceCreateConfig) GetPlan() InstanceConfigObject1Plan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *CatalogItemTypeInstanceCreateConfig) GetPlanOk() (*InstanceConfigObject1Plan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *CatalogItemTypeInstanceCreateConfig) SetPlan(v InstanceConfigObject1Plan)`

SetPlan sets Plan field to given value.


### GetVersion

`func (o *CatalogItemTypeInstanceCreateConfig) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CatalogItemTypeInstanceCreateConfig) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CatalogItemTypeInstanceCreateConfig) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CatalogItemTypeInstanceCreateConfig) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetEvars

`func (o *CatalogItemTypeInstanceCreateConfig) GetEvars() []InstanceConfigObject1EvarsInner`

GetEvars returns the Evars field if non-nil, zero value otherwise.

### GetEvarsOk

`func (o *CatalogItemTypeInstanceCreateConfig) GetEvarsOk() (*[]InstanceConfigObject1EvarsInner, bool)`

GetEvarsOk returns a tuple with the Evars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvars

`func (o *CatalogItemTypeInstanceCreateConfig) SetEvars(v []InstanceConfigObject1EvarsInner)`

SetEvars sets Evars field to given value.

### HasEvars

`func (o *CatalogItemTypeInstanceCreateConfig) HasEvars() bool`

HasEvars returns a boolean if a field has been set.

### GetServicePlanOptions

`func (o *CatalogItemTypeInstanceCreateConfig) GetServicePlanOptions() InstanceConfigObject1ServicePlanOptions`

GetServicePlanOptions returns the ServicePlanOptions field if non-nil, zero value otherwise.

### GetServicePlanOptionsOk

`func (o *CatalogItemTypeInstanceCreateConfig) GetServicePlanOptionsOk() (*InstanceConfigObject1ServicePlanOptions, bool)`

GetServicePlanOptionsOk returns a tuple with the ServicePlanOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanOptions

`func (o *CatalogItemTypeInstanceCreateConfig) SetServicePlanOptions(v InstanceConfigObject1ServicePlanOptions)`

SetServicePlanOptions sets ServicePlanOptions field to given value.

### HasServicePlanOptions

`func (o *CatalogItemTypeInstanceCreateConfig) HasServicePlanOptions() bool`

HasServicePlanOptions returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *CatalogItemTypeInstanceCreateConfig) GetSecurityGroups() []InstanceConfigObject1SecurityGroupsInner`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *CatalogItemTypeInstanceCreateConfig) GetSecurityGroupsOk() (*[]InstanceConfigObject1SecurityGroupsInner, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *CatalogItemTypeInstanceCreateConfig) SetSecurityGroups(v []InstanceConfigObject1SecurityGroupsInner)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *CatalogItemTypeInstanceCreateConfig) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *CatalogItemTypeInstanceCreateConfig) GetNetworkInterfaces() []InstancesNetworkInterfaces4`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *CatalogItemTypeInstanceCreateConfig) GetNetworkInterfacesOk() (*[]InstancesNetworkInterfaces4, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *CatalogItemTypeInstanceCreateConfig) SetNetworkInterfaces(v []InstancesNetworkInterfaces4)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *CatalogItemTypeInstanceCreateConfig) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetLabels

`func (o *CatalogItemTypeInstanceCreateConfig) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CatalogItemTypeInstanceCreateConfig) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CatalogItemTypeInstanceCreateConfig) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CatalogItemTypeInstanceCreateConfig) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTags

`func (o *CatalogItemTypeInstanceCreateConfig) GetTags() []InstanceConfigObject1TagsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CatalogItemTypeInstanceCreateConfig) GetTagsOk() (*[]InstanceConfigObject1TagsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CatalogItemTypeInstanceCreateConfig) SetTags(v []InstanceConfigObject1TagsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CatalogItemTypeInstanceCreateConfig) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMetadata

`func (o *CatalogItemTypeInstanceCreateConfig) GetMetadata() []InstanceConfigObject1MetadataInner`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CatalogItemTypeInstanceCreateConfig) GetMetadataOk() (*[]InstanceConfigObject1MetadataInner, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CatalogItemTypeInstanceCreateConfig) SetMetadata(v []InstanceConfigObject1MetadataInner)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CatalogItemTypeInstanceCreateConfig) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPorts

`func (o *CatalogItemTypeInstanceCreateConfig) GetPorts() []InstanceConfigObject1PortsInner`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *CatalogItemTypeInstanceCreateConfig) GetPortsOk() (*[]InstanceConfigObject1PortsInner, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *CatalogItemTypeInstanceCreateConfig) SetPorts(v []InstanceConfigObject1PortsInner)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *CatalogItemTypeInstanceCreateConfig) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetTaskSetId

`func (o *CatalogItemTypeInstanceCreateConfig) GetTaskSetId() int64`

GetTaskSetId returns the TaskSetId field if non-nil, zero value otherwise.

### GetTaskSetIdOk

`func (o *CatalogItemTypeInstanceCreateConfig) GetTaskSetIdOk() (*int64, bool)`

GetTaskSetIdOk returns a tuple with the TaskSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSetId

`func (o *CatalogItemTypeInstanceCreateConfig) SetTaskSetId(v int64)`

SetTaskSetId sets TaskSetId field to given value.

### HasTaskSetId

`func (o *CatalogItemTypeInstanceCreateConfig) HasTaskSetId() bool`

HasTaskSetId returns a boolean if a field has been set.

### GetTaskSetName

`func (o *CatalogItemTypeInstanceCreateConfig) GetTaskSetName() string`

GetTaskSetName returns the TaskSetName field if non-nil, zero value otherwise.

### GetTaskSetNameOk

`func (o *CatalogItemTypeInstanceCreateConfig) GetTaskSetNameOk() (*string, bool)`

GetTaskSetNameOk returns a tuple with the TaskSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSetName

`func (o *CatalogItemTypeInstanceCreateConfig) SetTaskSetName(v string)`

SetTaskSetName sets TaskSetName field to given value.

### HasTaskSetName

`func (o *CatalogItemTypeInstanceCreateConfig) HasTaskSetName() bool`

HasTaskSetName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


