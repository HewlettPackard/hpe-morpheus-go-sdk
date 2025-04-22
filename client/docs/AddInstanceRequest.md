# AddInstanceRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instance** | [**AddInstanceRequestInstance**](AddInstanceRequestInstance.md) |  | 
**ZoneId** | Pointer to **int64** | The Cloud ID to provision the instance onto. | [optional] 
**Evars** | Pointer to [**[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner**](AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner.md) | Environment Variables, an array of objects that have name and value. | [optional] 
**Copies** | Pointer to **int64** | Number of copies to provision. | [optional] [default to 1]
**LayoutSize** | Pointer to **int64** | Apply a multiply factor of containers/vms within the instance. | [optional] [default to 1]
**ServicePlanOptions** | Pointer to **map[string]interface{}** | Map of custom options depending on selected service plan. | [optional] 
**SecurityGroups** | Pointer to **[]map[string]interface{}** | Key for security group configuration. It should be passed as an array of objects containing the id of the security group to assign the instance to. | [optional] 
**Volumes** | Pointer to [**[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigVolumesInner**](AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigVolumesInner.md) | The (optional) volumes parameter is for LV configuration, can create additional LVs at provision It should be passed as an array of | [optional] 
**NetworkInterfaces** | Pointer to [**[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInner**](AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInner.md) | The networkInterfaces parameter is for network configuration.  The Options API &#x60;/api/options/zoneNetworkOptions?zoneId&#x3D;5&amp;provisionTypeId&#x3D;10&#x60; can be used to see which options are available.  | [optional] 
**Config** | [**AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfig**](AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfig.md) |  | 
**Labels** | Pointer to **[]string** | Array of strings (keywords). | [optional] 
**Tags** | Pointer to [**[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner**](AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner.md) | Metadata tags, Array of objects having a name and value. | [optional] 
**Metadata** | Pointer to [**[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner**](AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner.md) | Alias for &#x60;tags&#x60;. | [optional] 
**Ports** | Pointer to [**[]AddInstanceRequestPortsInner**](AddInstanceRequestPortsInner.md) | The ports parameter is for port configuration.  The layout may have default ports, which are defined in node types, that are always configured. This parameter will be for additional custom ports to be opened.  | [optional] 
**TaskSetId** | Pointer to **int64** | The Workflow ID to execute. | [optional] 
**TaskSetName** | Pointer to **string** | The Workflow Name to execute. | [optional] 

## Methods

### NewAddInstanceRequest

`func NewAddInstanceRequest(instance AddInstanceRequestInstance, config AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfig, ) *AddInstanceRequest`

NewAddInstanceRequest instantiates a new AddInstanceRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddInstanceRequestWithDefaults

`func NewAddInstanceRequestWithDefaults() *AddInstanceRequest`

NewAddInstanceRequestWithDefaults instantiates a new AddInstanceRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstance

`func (o *AddInstanceRequest) GetInstance() AddInstanceRequestInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *AddInstanceRequest) GetInstanceOk() (*AddInstanceRequestInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *AddInstanceRequest) SetInstance(v AddInstanceRequestInstance)`

SetInstance sets Instance field to given value.


### GetZoneId

`func (o *AddInstanceRequest) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *AddInstanceRequest) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *AddInstanceRequest) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *AddInstanceRequest) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetEvars

`func (o *AddInstanceRequest) GetEvars() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner`

GetEvars returns the Evars field if non-nil, zero value otherwise.

### GetEvarsOk

`func (o *AddInstanceRequest) GetEvarsOk() (*[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner, bool)`

GetEvarsOk returns a tuple with the Evars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvars

`func (o *AddInstanceRequest) SetEvars(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner)`

SetEvars sets Evars field to given value.

### HasEvars

`func (o *AddInstanceRequest) HasEvars() bool`

HasEvars returns a boolean if a field has been set.

### GetCopies

`func (o *AddInstanceRequest) GetCopies() int64`

GetCopies returns the Copies field if non-nil, zero value otherwise.

### GetCopiesOk

`func (o *AddInstanceRequest) GetCopiesOk() (*int64, bool)`

GetCopiesOk returns a tuple with the Copies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopies

`func (o *AddInstanceRequest) SetCopies(v int64)`

SetCopies sets Copies field to given value.

### HasCopies

`func (o *AddInstanceRequest) HasCopies() bool`

HasCopies returns a boolean if a field has been set.

### GetLayoutSize

`func (o *AddInstanceRequest) GetLayoutSize() int64`

GetLayoutSize returns the LayoutSize field if non-nil, zero value otherwise.

### GetLayoutSizeOk

`func (o *AddInstanceRequest) GetLayoutSizeOk() (*int64, bool)`

GetLayoutSizeOk returns a tuple with the LayoutSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutSize

`func (o *AddInstanceRequest) SetLayoutSize(v int64)`

SetLayoutSize sets LayoutSize field to given value.

### HasLayoutSize

`func (o *AddInstanceRequest) HasLayoutSize() bool`

HasLayoutSize returns a boolean if a field has been set.

### GetServicePlanOptions

`func (o *AddInstanceRequest) GetServicePlanOptions() map[string]interface{}`

GetServicePlanOptions returns the ServicePlanOptions field if non-nil, zero value otherwise.

### GetServicePlanOptionsOk

`func (o *AddInstanceRequest) GetServicePlanOptionsOk() (*map[string]interface{}, bool)`

GetServicePlanOptionsOk returns a tuple with the ServicePlanOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanOptions

`func (o *AddInstanceRequest) SetServicePlanOptions(v map[string]interface{})`

SetServicePlanOptions sets ServicePlanOptions field to given value.

### HasServicePlanOptions

`func (o *AddInstanceRequest) HasServicePlanOptions() bool`

HasServicePlanOptions returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *AddInstanceRequest) GetSecurityGroups() []map[string]interface{}`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *AddInstanceRequest) GetSecurityGroupsOk() (*[]map[string]interface{}, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *AddInstanceRequest) SetSecurityGroups(v []map[string]interface{})`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *AddInstanceRequest) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetVolumes

`func (o *AddInstanceRequest) GetVolumes() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *AddInstanceRequest) GetVolumesOk() (*[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *AddInstanceRequest) SetVolumes(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigVolumesInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *AddInstanceRequest) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *AddInstanceRequest) GetNetworkInterfaces() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInner`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *AddInstanceRequest) GetNetworkInterfacesOk() (*[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInner, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *AddInstanceRequest) SetNetworkInterfaces(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInner)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *AddInstanceRequest) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetConfig

`func (o *AddInstanceRequest) GetConfig() AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddInstanceRequest) GetConfigOk() (*AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddInstanceRequest) SetConfig(v AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfig)`

SetConfig sets Config field to given value.


### GetLabels

`func (o *AddInstanceRequest) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddInstanceRequest) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddInstanceRequest) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddInstanceRequest) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTags

`func (o *AddInstanceRequest) GetTags() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AddInstanceRequest) GetTagsOk() (*[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AddInstanceRequest) SetTags(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AddInstanceRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMetadata

`func (o *AddInstanceRequest) GetMetadata() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AddInstanceRequest) GetMetadataOk() (*[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AddInstanceRequest) SetMetadata(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AddInstanceRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPorts

`func (o *AddInstanceRequest) GetPorts() []AddInstanceRequestPortsInner`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *AddInstanceRequest) GetPortsOk() (*[]AddInstanceRequestPortsInner, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *AddInstanceRequest) SetPorts(v []AddInstanceRequestPortsInner)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *AddInstanceRequest) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetTaskSetId

`func (o *AddInstanceRequest) GetTaskSetId() int64`

GetTaskSetId returns the TaskSetId field if non-nil, zero value otherwise.

### GetTaskSetIdOk

`func (o *AddInstanceRequest) GetTaskSetIdOk() (*int64, bool)`

GetTaskSetIdOk returns a tuple with the TaskSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSetId

`func (o *AddInstanceRequest) SetTaskSetId(v int64)`

SetTaskSetId sets TaskSetId field to given value.

### HasTaskSetId

`func (o *AddInstanceRequest) HasTaskSetId() bool`

HasTaskSetId returns a boolean if a field has been set.

### GetTaskSetName

`func (o *AddInstanceRequest) GetTaskSetName() string`

GetTaskSetName returns the TaskSetName field if non-nil, zero value otherwise.

### GetTaskSetNameOk

`func (o *AddInstanceRequest) GetTaskSetNameOk() (*string, bool)`

GetTaskSetNameOk returns a tuple with the TaskSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSetName

`func (o *AddInstanceRequest) SetTaskSetName(v string)`

SetTaskSetName sets TaskSetName field to given value.

### HasTaskSetName

`func (o *AddInstanceRequest) HasTaskSetName() bool`

HasTaskSetName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


