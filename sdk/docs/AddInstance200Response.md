# AddInstance200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Errors** | Pointer to **map[string]interface{}** |  | [optional] 
**Instance** | [**AddInstanceRequestInstance**](AddInstanceRequestInstance.md) |  | 
**ZoneId** | Pointer to **int64** | The Cloud ID to provision the instance onto. | [optional] 
**Evars** | Pointer to [**[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner**](AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner.md) | Environment Variables, an array of objects that have name and value. | [optional] 
**Copies** | Pointer to **int64** | Number of copies to provision. | [optional] [default to 1]
**LayoutSize** | Pointer to **int64** | Apply a multiply factor of containers/vms within the instance. | [optional] [default to 1]
**ServicePlanOptions** | Pointer to **interface{}** | Map of custom options depending on selected service plan. | [optional] 
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

### NewAddInstance200Response

`func NewAddInstance200Response(instance AddInstanceRequestInstance, config AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfig, ) *AddInstance200Response`

NewAddInstance200Response instantiates a new AddInstance200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddInstance200ResponseWithDefaults

`func NewAddInstance200ResponseWithDefaults() *AddInstance200Response`

NewAddInstance200ResponseWithDefaults instantiates a new AddInstance200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *AddInstance200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddInstance200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddInstance200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddInstance200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetErrors

`func (o *AddInstance200Response) GetErrors() map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *AddInstance200Response) GetErrorsOk() (*map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *AddInstance200Response) SetErrors(v map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *AddInstance200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### GetInstance

`func (o *AddInstance200Response) GetInstance() AddInstanceRequestInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *AddInstance200Response) GetInstanceOk() (*AddInstanceRequestInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *AddInstance200Response) SetInstance(v AddInstanceRequestInstance)`

SetInstance sets Instance field to given value.


### GetZoneId

`func (o *AddInstance200Response) GetZoneId() int64`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *AddInstance200Response) GetZoneIdOk() (*int64, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *AddInstance200Response) SetZoneId(v int64)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *AddInstance200Response) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.

### GetEvars

`func (o *AddInstance200Response) GetEvars() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner`

GetEvars returns the Evars field if non-nil, zero value otherwise.

### GetEvarsOk

`func (o *AddInstance200Response) GetEvarsOk() (*[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner, bool)`

GetEvarsOk returns a tuple with the Evars field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvars

`func (o *AddInstance200Response) SetEvars(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner)`

SetEvars sets Evars field to given value.

### HasEvars

`func (o *AddInstance200Response) HasEvars() bool`

HasEvars returns a boolean if a field has been set.

### GetCopies

`func (o *AddInstance200Response) GetCopies() int64`

GetCopies returns the Copies field if non-nil, zero value otherwise.

### GetCopiesOk

`func (o *AddInstance200Response) GetCopiesOk() (*int64, bool)`

GetCopiesOk returns a tuple with the Copies field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCopies

`func (o *AddInstance200Response) SetCopies(v int64)`

SetCopies sets Copies field to given value.

### HasCopies

`func (o *AddInstance200Response) HasCopies() bool`

HasCopies returns a boolean if a field has been set.

### GetLayoutSize

`func (o *AddInstance200Response) GetLayoutSize() int64`

GetLayoutSize returns the LayoutSize field if non-nil, zero value otherwise.

### GetLayoutSizeOk

`func (o *AddInstance200Response) GetLayoutSizeOk() (*int64, bool)`

GetLayoutSizeOk returns a tuple with the LayoutSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayoutSize

`func (o *AddInstance200Response) SetLayoutSize(v int64)`

SetLayoutSize sets LayoutSize field to given value.

### HasLayoutSize

`func (o *AddInstance200Response) HasLayoutSize() bool`

HasLayoutSize returns a boolean if a field has been set.

### GetServicePlanOptions

`func (o *AddInstance200Response) GetServicePlanOptions() interface{}`

GetServicePlanOptions returns the ServicePlanOptions field if non-nil, zero value otherwise.

### GetServicePlanOptionsOk

`func (o *AddInstance200Response) GetServicePlanOptionsOk() (*interface{}, bool)`

GetServicePlanOptionsOk returns a tuple with the ServicePlanOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlanOptions

`func (o *AddInstance200Response) SetServicePlanOptions(v interface{})`

SetServicePlanOptions sets ServicePlanOptions field to given value.

### HasServicePlanOptions

`func (o *AddInstance200Response) HasServicePlanOptions() bool`

HasServicePlanOptions returns a boolean if a field has been set.

### SetServicePlanOptionsNil

`func (o *AddInstance200Response) SetServicePlanOptionsNil(b bool)`

 SetServicePlanOptionsNil sets the value for ServicePlanOptions to be an explicit nil

### UnsetServicePlanOptions
`func (o *AddInstance200Response) UnsetServicePlanOptions()`

UnsetServicePlanOptions ensures that no value is present for ServicePlanOptions, not even an explicit nil
### GetSecurityGroups

`func (o *AddInstance200Response) GetSecurityGroups() []map[string]interface{}`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *AddInstance200Response) GetSecurityGroupsOk() (*[]map[string]interface{}, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *AddInstance200Response) SetSecurityGroups(v []map[string]interface{})`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *AddInstance200Response) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetVolumes

`func (o *AddInstance200Response) GetVolumes() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *AddInstance200Response) GetVolumesOk() (*[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *AddInstance200Response) SetVolumes(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigVolumesInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *AddInstance200Response) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetNetworkInterfaces

`func (o *AddInstance200Response) GetNetworkInterfaces() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInner`

GetNetworkInterfaces returns the NetworkInterfaces field if non-nil, zero value otherwise.

### GetNetworkInterfacesOk

`func (o *AddInstance200Response) GetNetworkInterfacesOk() (*[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInner, bool)`

GetNetworkInterfacesOk returns a tuple with the NetworkInterfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkInterfaces

`func (o *AddInstance200Response) SetNetworkInterfaces(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigNetworkInterfacesInner)`

SetNetworkInterfaces sets NetworkInterfaces field to given value.

### HasNetworkInterfaces

`func (o *AddInstance200Response) HasNetworkInterfaces() bool`

HasNetworkInterfaces returns a boolean if a field has been set.

### GetConfig

`func (o *AddInstance200Response) GetConfig() AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddInstance200Response) GetConfigOk() (*AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddInstance200Response) SetConfig(v AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigConfig)`

SetConfig sets Config field to given value.


### GetLabels

`func (o *AddInstance200Response) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddInstance200Response) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddInstance200Response) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddInstance200Response) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetTags

`func (o *AddInstance200Response) GetTags() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AddInstance200Response) GetTagsOk() (*[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AddInstance200Response) SetTags(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AddInstance200Response) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetMetadata

`func (o *AddInstance200Response) GetMetadata() []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *AddInstance200Response) GetMetadataOk() (*[]AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *AddInstance200Response) SetMetadata(v []AddCatalogItemTypeRequestCatalogItemTypeOneOfConfigEvarsInner)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *AddInstance200Response) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetPorts

`func (o *AddInstance200Response) GetPorts() []AddInstanceRequestPortsInner`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *AddInstance200Response) GetPortsOk() (*[]AddInstanceRequestPortsInner, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *AddInstance200Response) SetPorts(v []AddInstanceRequestPortsInner)`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *AddInstance200Response) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### GetTaskSetId

`func (o *AddInstance200Response) GetTaskSetId() int64`

GetTaskSetId returns the TaskSetId field if non-nil, zero value otherwise.

### GetTaskSetIdOk

`func (o *AddInstance200Response) GetTaskSetIdOk() (*int64, bool)`

GetTaskSetIdOk returns a tuple with the TaskSetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSetId

`func (o *AddInstance200Response) SetTaskSetId(v int64)`

SetTaskSetId sets TaskSetId field to given value.

### HasTaskSetId

`func (o *AddInstance200Response) HasTaskSetId() bool`

HasTaskSetId returns a boolean if a field has been set.

### GetTaskSetName

`func (o *AddInstance200Response) GetTaskSetName() string`

GetTaskSetName returns the TaskSetName field if non-nil, zero value otherwise.

### GetTaskSetNameOk

`func (o *AddInstance200Response) GetTaskSetNameOk() (*string, bool)`

GetTaskSetNameOk returns a tuple with the TaskSetName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSetName

`func (o *AddInstance200Response) SetTaskSetName(v string)`

SetTaskSetName sets TaskSetName field to given value.

### HasTaskSetName

`func (o *AddInstance200Response) HasTaskSetName() bool`

HasTaskSetName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


