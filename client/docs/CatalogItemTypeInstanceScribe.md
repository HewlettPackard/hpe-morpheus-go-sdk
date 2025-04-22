# CatalogItemTypeInstanceScribe

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

### NewCatalogItemTypeInstanceScribe

`func NewCatalogItemTypeInstanceScribe(group AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigGroup, cloud AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigCloud, type_ string, name string, config AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfig, volumes []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigVolumesInner, layout AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigLayout, plan AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan, ) *CatalogItemTypeInstanceScribe`

NewCatalogItemTypeInstanceScribe instantiates a new CatalogItemTypeInstanceScribe object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogItemTypeInstanceScribeWithDefaults

`func NewCatalogItemTypeInstanceScribeWithDefaults() *CatalogItemTypeInstanceScribe`

NewCatalogItemTypeInstanceScribeWithDefaults instantiates a new CatalogItemTypeInstanceScribe object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *CatalogItemTypeInstanceScribe) GetGroup() AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *CatalogItemTypeInstanceScribe) GetGroupOk() (*AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *CatalogItemTypeInstanceScribe) SetGroup(v AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigGroup)`

SetGroup sets Group field to given value.


### GetCloud

`func (o *CatalogItemTypeInstanceScribe) GetCloud() AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigCloud`

GetCloud returns the Cloud field if non-nil, zero value otherwise.

### GetCloudOk

`func (o *CatalogItemTypeInstanceScribe) GetCloudOk() (*AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigCloud, bool)`

GetCloudOk returns a tuple with the Cloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCloud

`func (o *CatalogItemTypeInstanceScribe) SetCloud(v AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigCloud)`

SetCloud sets Cloud field to given value.


### GetType

`func (o *CatalogItemTypeInstanceScribe) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CatalogItemTypeInstanceScribe) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CatalogItemTypeInstanceScribe) SetType(v string)`

SetType sets Type field to given value.


### GetName

`func (o *CatalogItemTypeInstanceScribe) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CatalogItemTypeInstanceScribe) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CatalogItemTypeInstanceScribe) SetName(v string)`

SetName sets Name field to given value.


### GetConfig

`func (o *CatalogItemTypeInstanceScribe) GetConfig() AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *CatalogItemTypeInstanceScribe) GetConfigOk() (*AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *CatalogItemTypeInstanceScribe) SetConfig(v AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfig)`

SetConfig sets Config field to given value.


### GetVolumes

`func (o *CatalogItemTypeInstanceScribe) GetVolumes() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *CatalogItemTypeInstanceScribe) GetVolumesOk() (*[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *CatalogItemTypeInstanceScribe) SetVolumes(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigVolumesInner)`

SetVolumes sets Volumes field to given value.


### GetHostName

`func (o *CatalogItemTypeInstanceScribe) GetHostName() string`

GetHostName returns the HostName field if non-nil, zero value otherwise.

### GetHostNameOk

`func (o *CatalogItemTypeInstanceScribe) GetHostNameOk() (*string, bool)`

GetHostNameOk returns a tuple with the HostName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostName

`func (o *CatalogItemTypeInstanceScribe) SetHostName(v string)`

SetHostName sets HostName field to given value.

### HasHostName

`func (o *CatalogItemTypeInstanceScribe) HasHostName() bool`

HasHostName returns a boolean if a field has been set.

### GetEnvironment

`func (o *CatalogItemTypeInstanceScribe) GetEnvironment() string`

GetEnvironment returns the Environment field if non-nil, zero value otherwise.

### GetEnvironmentOk

`func (o *CatalogItemTypeInstanceScribe) GetEnvironmentOk() (*string, bool)`

GetEnvironmentOk returns a tuple with the Environment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironment

`func (o *CatalogItemTypeInstanceScribe) SetEnvironment(v string)`

SetEnvironment sets Environment field to given value.

### HasEnvironment

`func (o *CatalogItemTypeInstanceScribe) HasEnvironment() bool`

HasEnvironment returns a boolean if a field has been set.

### GetLayout

`func (o *CatalogItemTypeInstanceScribe) GetLayout() AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *CatalogItemTypeInstanceScribe) GetLayoutOk() (*AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *CatalogItemTypeInstanceScribe) SetLayout(v AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigLayout)`

SetLayout sets Layout field to given value.


### GetPlan

`func (o *CatalogItemTypeInstanceScribe) GetPlan() AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *CatalogItemTypeInstanceScribe) GetPlanOk() (*AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *CatalogItemTypeInstanceScribe) SetPlan(v AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPlan)`

SetPlan sets Plan field to given value.


### GetVersion

`func (o *CatalogItemTypeInstanceScribe) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CatalogItemTypeInstanceScribe) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CatalogItemTypeInstanceScribe) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CatalogItemTypeInstanceScribe) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetEvars

`func (o *CatalogItemTypeInstanceScribe) GetEvars() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner`

GetEvars returns the Evars field if non-nil, zero value otherwise.

### GetEvarsOk

`func (o *CatalogItemTypeInstanceScribe) GetEvarsOk() (*[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner, bool)`

GetEvarsOk returns a tuple with the Evars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvars

`func (o *CatalogItemTypeInstanceScribe) SetEvars(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner)`

SetEvars sets Evars field to given value.

### HasEvars

`func (o *CatalogItemTypeInstanceScribe) HasEvars() bool`

HasEvars returns a boolean if a field has been set.

### GetServicePlanOptions

`func (o *CatalogItemTypeInstanceScribe) GetServicePlanOptions() AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigServicePlanOptions`

GetServicePlanOptions returns the ServicePlanOptions field if non-nil, zero value otherwise.

### GetServicePlanOptionsOk

`func (o *CatalogItemTypeInstanceScribe) GetServicePlanOptionsOk() (*AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigServicePlanOptions, bool)`

GetServicePlanOptionsOk returns a tuple with the ServicePlanOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanOptions

`func (o *CatalogItemTypeInstanceScribe) SetServicePlanOptions(v AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigServicePlanOptions)`

SetServicePlanOptions sets ServicePlanOptions field to given value.

### HasServicePlanOptions

`func (o *CatalogItemTypeInstanceScribe) HasServicePlanOptions() bool`

HasServicePlanOptions returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *CatalogItemTypeInstanceScribe) GetSecurityGroups() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigSecurityGroupsInner`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *CatalogItemTypeInstanceScribe) GetSecurityGroupsOk() (*[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigSecurityGroupsInner, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *CatalogItemTypeInstanceScribe) SetSecurityGroups(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigSecurityGroupsInner)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *CatalogItemTypeInstanceScribe) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *CatalogItemTypeInstanceScribe) GetNetworkInterfaces() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInner`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *CatalogItemTypeInstanceScribe) GetNetworkInterfacesOk() (*[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInner, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *CatalogItemTypeInstanceScribe) SetNetworkInterfaces(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInner)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *CatalogItemTypeInstanceScribe) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetLabels

`func (o *CatalogItemTypeInstanceScribe) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *CatalogItemTypeInstanceScribe) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *CatalogItemTypeInstanceScribe) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *CatalogItemTypeInstanceScribe) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTags

`func (o *CatalogItemTypeInstanceScribe) GetTags() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CatalogItemTypeInstanceScribe) GetTagsOk() (*[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CatalogItemTypeInstanceScribe) SetTags(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CatalogItemTypeInstanceScribe) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMetadata

`func (o *CatalogItemTypeInstanceScribe) GetMetadata() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CatalogItemTypeInstanceScribe) GetMetadataOk() (*[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CatalogItemTypeInstanceScribe) SetMetadata(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CatalogItemTypeInstanceScribe) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPorts

`func (o *CatalogItemTypeInstanceScribe) GetPorts() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPortsInner`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *CatalogItemTypeInstanceScribe) GetPortsOk() (*[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPortsInner, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *CatalogItemTypeInstanceScribe) SetPorts(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigPortsInner)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *CatalogItemTypeInstanceScribe) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetTaskSetId

`func (o *CatalogItemTypeInstanceScribe) GetTaskSetId() int64`

GetTaskSetId returns the TaskSetId field if non-nil, zero value otherwise.

### GetTaskSetIdOk

`func (o *CatalogItemTypeInstanceScribe) GetTaskSetIdOk() (*int64, bool)`

GetTaskSetIdOk returns a tuple with the TaskSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSetId

`func (o *CatalogItemTypeInstanceScribe) SetTaskSetId(v int64)`

SetTaskSetId sets TaskSetId field to given value.

### HasTaskSetId

`func (o *CatalogItemTypeInstanceScribe) HasTaskSetId() bool`

HasTaskSetId returns a boolean if a field has been set.

### GetTaskSetName

`func (o *CatalogItemTypeInstanceScribe) GetTaskSetName() string`

GetTaskSetName returns the TaskSetName field if non-nil, zero value otherwise.

### GetTaskSetNameOk

`func (o *CatalogItemTypeInstanceScribe) GetTaskSetNameOk() (*string, bool)`

GetTaskSetNameOk returns a tuple with the TaskSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSetName

`func (o *CatalogItemTypeInstanceScribe) SetTaskSetName(v string)`

SetTaskSetName sets TaskSetName field to given value.

### HasTaskSetName

`func (o *CatalogItemTypeInstanceScribe) HasTaskSetName() bool`

HasTaskSetName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


