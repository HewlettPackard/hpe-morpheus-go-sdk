# AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | [**AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigGroup**](AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigGroup.md) |  | 
**Cloud** | [**AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigCloud**](AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigCloud.md) |  | 
**Type** | **string** | The type of instance by code we want to fetch. | 
**Name** | **string** | Name of the instance to be created. | 
**Config** | [**AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfig**](AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfig.md) |  | 
**Volumes** | [**[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigVolumesInner**](AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigVolumesInner.md) | The (optional) volumes parameter is for LV configuration, can create additional LVs at provision It should be passed as an array of | 
**HostName** | Pointer to **string** | Hostname of the instance to be created.  Can be the same as the instance name. | [optional] 
**Environment** | Pointer to **string** | Environment code | [optional] 
**Layout** | [**AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigLayout**](AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigLayout.md) |  | 
**Plan** | [**AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan**](AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan.md) |  | 
**Version** | Pointer to **string** | Version of the layout to create. | [optional] 
**Evars** | Pointer to [**[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner**](AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner.md) | Environment Variables, an array of objects that have name and value. | [optional] 
**ServicePlanOptions** | Pointer to [**AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigServicePlanOptions**](AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigServicePlanOptions.md) |  | [optional] 
**SecurityGroups** | Pointer to [**[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigSecurityGroupsInner**](AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigSecurityGroupsInner.md) | Key for security group configuration. It should be passed as an array of objects containing the id of the security group to assign the instance to. | [optional] 
**NetworkInterfaces** | Pointer to [**[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInner**](AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInner.md) | The networkInterfaces parameter is for network configuration.  The Options API &#x60;/api/options/zoneNetworkOptions?zoneId&#x3D;5&amp;provisionTypeId&#x3D;10&#x60; can be used to see which options are available.  | [optional] 
**Labels** | Pointer to **[]string** | Array of strings (keywords). | [optional] 
**Tags** | Pointer to [**[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner**](AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner.md) | Metadata tags, Array of objects having a name and value. | [optional] 
**Metadata** | Pointer to [**[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner**](AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner.md) | Alias for &#x60;tags&#x60;. | [optional] 
**Ports** | Pointer to [**[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPortsInner**](AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPortsInner.md) | The ports parameter is for port configuration.  The layout may have default ports, which are defined in node types, that are always configured. This parameter will be for additional custom ports to be opened.  | [optional] 
**TaskSetId** | Pointer to **int64** | The Workflow ID to execute. | [optional] 
**TaskSetName** | Pointer to **string** | The Workflow Name to execute. | [optional] 

## Methods

### NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfig

`func NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfig(group AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigGroup, cloud AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigCloud, type_ string, name string, config AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfig, volumes []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigVolumesInner, layout AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigLayout, plan AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan, ) *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig`

NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfig instantiates a new AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigWithDefaults

`func NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigWithDefaults() *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig`

NewAddCatalogItemTypeRequestCatalogItemTypeOneOfConfigWithDefaults instantiates a new AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetGroup() AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetGroupOk() (*AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) SetGroup(v AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigGroup)`

SetGroup sets Group field to given value.


### GetCloud

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetCloud() AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigCloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetCloudOk() (*AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigCloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) SetCloud(v AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigCloud)`

SetCloud sets Cloud field to given value.


### GetType

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) SetName(v string)`

SetName sets Name field to given value.


### GetConfig

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetConfig() AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetConfigOk() (*AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) SetConfig(v AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfig)`

SetConfig sets Config field to given value.


### GetVolumes

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetVolumes() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetVolumesOk() (*[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) SetVolumes(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigVolumesInner)`

SetVolumes sets Volumes field to given value.


### GetHostName

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetEnvironment

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetLayout

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetLayout() AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetLayoutOk() (*AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) SetLayout(v AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigLayout)`

SetLayout sets Layout field to given value.


### GetPlan

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetPlan() AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetPlanOk() (*AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) SetPlan(v AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan)`

SetPlan sets Plan field to given value.


### GetVersion

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetEvars

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetEvars() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner`

GetEvars returns the Evars field if non-nil, zero value otherwise.

### GetEvarsOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetEvarsOk() (*[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner, bool)`

GetEvarsOk returns a tuple with the Evars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvars

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) SetEvars(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner)`

SetEvars sets Evars field to given value.

### HasEvars

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) HasEvars() bool`

HasEvars returns a boolean if a field has been set.

### GetServicePlanOptions

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetServicePlanOptions() AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigServicePlanOptions`

GetServicePlanOptions returns the ServicePlanOptions field if non-nil, zero value otherwise.

### GetServicePlanOptionsOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetServicePlanOptionsOk() (*AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigServicePlanOptions, bool)`

GetServicePlanOptionsOk returns a tuple with the ServicePlanOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanOptions

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) SetServicePlanOptions(v AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigServicePlanOptions)`

SetServicePlanOptions sets ServicePlanOptions field to given value.

### HasServicePlanOptions

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) HasServicePlanOptions() bool`

HasServicePlanOptions returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetSecurityGroups() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigSecurityGroupsInner`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetSecurityGroupsOk() (*[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigSecurityGroupsInner, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) SetSecurityGroups(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigSecurityGroupsInner)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetNetworkInterfaces() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInner`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetNetworkInterfacesOk() (*[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInner, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) SetNetworkInterfaces(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInner)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetLabels

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTags

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetTags() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetTagsOk() (*[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) SetTags(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMetadata

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetMetadata() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetMetadataOk() (*[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) SetMetadata(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPorts

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetPorts() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPortsInner`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetPortsOk() (*[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPortsInner, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) SetPorts(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPortsInner)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetTaskSetId

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetTaskSetId() int64`

GetTaskSetId returns the TaskSetId field if non-nil, zero value otherwise.

### GetTaskSetIdOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetTaskSetIdOk() (*int64, bool)`

GetTaskSetIdOk returns a tuple with the TaskSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSetId

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) SetTaskSetId(v int64)`

SetTaskSetId sets TaskSetId field to given value.

### HasTaskSetId

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) HasTaskSetId() bool`

HasTaskSetId returns a boolean if a field has been set.

### GetTaskSetName

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetTaskSetName() string`

GetTaskSetName returns the TaskSetName field if non-nil, zero value otherwise.

### GetTaskSetNameOk

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) GetTaskSetNameOk() (*string, bool)`

GetTaskSetNameOk returns a tuple with the TaskSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSetName

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) SetTaskSetName(v string)`

SetTaskSetName sets TaskSetName field to given value.

### HasTaskSetName

`func (o *AddCatalogItemTypeRequestCatalogItemTypeOneOfConfig) HasTaskSetName() bool`

HasTaskSetName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


