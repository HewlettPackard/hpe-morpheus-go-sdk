# UpdateNodeTypeRequestContainerType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Node type name | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**ShortName** | Pointer to **string** | The short name is a name with no spaces used for display in your container list. | [optional] 
**Description** | Pointer to **string** | Node type description | [optional] 
**ContainerVersion** | Pointer to **string** | Version of the node type | [optional] 
**ProvisionTypeCode** | Pointer to **string** | Provision type code, eg. &#x60;amazon&#x60;, etc. | [optional] 
**Scripts** | Pointer to **[]int64** | Array of script IDs. | [optional] 
**Templates** | Pointer to **[]int64** | Array of file template IDs. | [optional] 
**VirtualImageId** | Pointer to **int64** | Virtual image ID | [optional] 
**OsTypeId** | Pointer to **int64** | OsType ID | [optional] 
**StatTypeCode** | Pointer to **string** | Stat type code.  Varies with node type, see Provision Types (customOptionTypes) for allowed values within selected type. | [optional] 
**LogTypeCode** | Pointer to **string** | Log type code.  Varies with node type, see Provision Types (customOptionTypes) for allowed values within selected type. | [optional] 
**ServerType** | Pointer to **string** | Server type.  Always pass \&quot;vm\&quot;. | [optional] 
**ContainerPorts** | Pointer to [**[]AddNodeTypeRequestContainerTypeContainerPortsInner**](AddNodeTypeRequestContainerTypeContainerPortsInner.md) | List of exposed port definitions in the format NAME&#x3D;PORT|PROTOCOL | [optional] 
**EnvironmentVariables** | Pointer to [**[]AddClusterLayoutsRequestLayoutEnvironmentVariablesInner**](AddClusterLayoutsRequestLayoutEnvironmentVariablesInner.md) | The environmentVariables parameter is array of env objects. | [optional] 
**Config** | Pointer to **map[string]interface{}** | Config object varies with node type.  If using docker, scvmm, ARM, hyperv, or cloudformation, look up provision type details (customOptionTypes) for information. | [optional] 

## Methods

### NewUpdateNodeTypeRequestContainerType

`func NewUpdateNodeTypeRequestContainerType() *UpdateNodeTypeRequestContainerType`

NewUpdateNodeTypeRequestContainerType instantiates a new UpdateNodeTypeRequestContainerType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNodeTypeRequestContainerTypeWithDefaults

`func NewUpdateNodeTypeRequestContainerTypeWithDefaults() *UpdateNodeTypeRequestContainerType`

NewUpdateNodeTypeRequestContainerTypeWithDefaults instantiates a new UpdateNodeTypeRequestContainerType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateNodeTypeRequestContainerType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNodeTypeRequestContainerType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNodeTypeRequestContainerType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateNodeTypeRequestContainerType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *UpdateNodeTypeRequestContainerType) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateNodeTypeRequestContainerType) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateNodeTypeRequestContainerType) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateNodeTypeRequestContainerType) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetShortName

`func (o *UpdateNodeTypeRequestContainerType) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *UpdateNodeTypeRequestContainerType) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *UpdateNodeTypeRequestContainerType) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *UpdateNodeTypeRequestContainerType) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateNodeTypeRequestContainerType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateNodeTypeRequestContainerType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateNodeTypeRequestContainerType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateNodeTypeRequestContainerType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetContainerVersion

`func (o *UpdateNodeTypeRequestContainerType) GetContainerVersion() string`

GetContainerVersion returns the ContainerVersion field if non-nil, zero value otherwise.

### GetContainerVersionOk

`func (o *UpdateNodeTypeRequestContainerType) GetContainerVersionOk() (*string, bool)`

GetContainerVersionOk returns a tuple with the ContainerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerVersion

`func (o *UpdateNodeTypeRequestContainerType) SetContainerVersion(v string)`

SetContainerVersion sets ContainerVersion field to given value.

### HasContainerVersion

`func (o *UpdateNodeTypeRequestContainerType) HasContainerVersion() bool`

HasContainerVersion returns a boolean if a field has been set.

### GetProvisionTypeCode

`func (o *UpdateNodeTypeRequestContainerType) GetProvisionTypeCode() string`

GetProvisionTypeCode returns the ProvisionTypeCode field if non-nil, zero value otherwise.

### GetProvisionTypeCodeOk

`func (o *UpdateNodeTypeRequestContainerType) GetProvisionTypeCodeOk() (*string, bool)`

GetProvisionTypeCodeOk returns a tuple with the ProvisionTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionTypeCode

`func (o *UpdateNodeTypeRequestContainerType) SetProvisionTypeCode(v string)`

SetProvisionTypeCode sets ProvisionTypeCode field to given value.

### HasProvisionTypeCode

`func (o *UpdateNodeTypeRequestContainerType) HasProvisionTypeCode() bool`

HasProvisionTypeCode returns a boolean if a field has been set.

### GetScripts

`func (o *UpdateNodeTypeRequestContainerType) GetScripts() []int64`

GetScripts returns the Scripts field if non-nil, zero value otherwise.

### GetScriptsOk

`func (o *UpdateNodeTypeRequestContainerType) GetScriptsOk() (*[]int64, bool)`

GetScriptsOk returns a tuple with the Scripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScripts

`func (o *UpdateNodeTypeRequestContainerType) SetScripts(v []int64)`

SetScripts sets Scripts field to given value.

### HasScripts

`func (o *UpdateNodeTypeRequestContainerType) HasScripts() bool`

HasScripts returns a boolean if a field has been set.

### GetTemplates

`func (o *UpdateNodeTypeRequestContainerType) GetTemplates() []int64`

GetTemplates returns the Templates field if non-nil, zero value otherwise.

### GetTemplatesOk

`func (o *UpdateNodeTypeRequestContainerType) GetTemplatesOk() (*[]int64, bool)`

GetTemplatesOk returns a tuple with the Templates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplates

`func (o *UpdateNodeTypeRequestContainerType) SetTemplates(v []int64)`

SetTemplates sets Templates field to given value.

### HasTemplates

`func (o *UpdateNodeTypeRequestContainerType) HasTemplates() bool`

HasTemplates returns a boolean if a field has been set.

### GetVirtualImageId

`func (o *UpdateNodeTypeRequestContainerType) GetVirtualImageId() int64`

GetVirtualImageId returns the VirtualImageId field if non-nil, zero value otherwise.

### GetVirtualImageIdOk

`func (o *UpdateNodeTypeRequestContainerType) GetVirtualImageIdOk() (*int64, bool)`

GetVirtualImageIdOk returns a tuple with the VirtualImageId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImageId

`func (o *UpdateNodeTypeRequestContainerType) SetVirtualImageId(v int64)`

SetVirtualImageId sets VirtualImageId field to given value.

### HasVirtualImageId

`func (o *UpdateNodeTypeRequestContainerType) HasVirtualImageId() bool`

HasVirtualImageId returns a boolean if a field has been set.

### GetOsTypeId

`func (o *UpdateNodeTypeRequestContainerType) GetOsTypeId() int64`

GetOsTypeId returns the OsTypeId field if non-nil, zero value otherwise.

### GetOsTypeIdOk

`func (o *UpdateNodeTypeRequestContainerType) GetOsTypeIdOk() (*int64, bool)`

GetOsTypeIdOk returns a tuple with the OsTypeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsTypeId

`func (o *UpdateNodeTypeRequestContainerType) SetOsTypeId(v int64)`

SetOsTypeId sets OsTypeId field to given value.

### HasOsTypeId

`func (o *UpdateNodeTypeRequestContainerType) HasOsTypeId() bool`

HasOsTypeId returns a boolean if a field has been set.

### GetStatTypeCode

`func (o *UpdateNodeTypeRequestContainerType) GetStatTypeCode() string`

GetStatTypeCode returns the StatTypeCode field if non-nil, zero value otherwise.

### GetStatTypeCodeOk

`func (o *UpdateNodeTypeRequestContainerType) GetStatTypeCodeOk() (*string, bool)`

GetStatTypeCodeOk returns a tuple with the StatTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatTypeCode

`func (o *UpdateNodeTypeRequestContainerType) SetStatTypeCode(v string)`

SetStatTypeCode sets StatTypeCode field to given value.

### HasStatTypeCode

`func (o *UpdateNodeTypeRequestContainerType) HasStatTypeCode() bool`

HasStatTypeCode returns a boolean if a field has been set.

### GetLogTypeCode

`func (o *UpdateNodeTypeRequestContainerType) GetLogTypeCode() string`

GetLogTypeCode returns the LogTypeCode field if non-nil, zero value otherwise.

### GetLogTypeCodeOk

`func (o *UpdateNodeTypeRequestContainerType) GetLogTypeCodeOk() (*string, bool)`

GetLogTypeCodeOk returns a tuple with the LogTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogTypeCode

`func (o *UpdateNodeTypeRequestContainerType) SetLogTypeCode(v string)`

SetLogTypeCode sets LogTypeCode field to given value.

### HasLogTypeCode

`func (o *UpdateNodeTypeRequestContainerType) HasLogTypeCode() bool`

HasLogTypeCode returns a boolean if a field has been set.

### GetServerType

`func (o *UpdateNodeTypeRequestContainerType) GetServerType() string`

GetServerType returns the ServerType field if non-nil, zero value otherwise.

### GetServerTypeOk

`func (o *UpdateNodeTypeRequestContainerType) GetServerTypeOk() (*string, bool)`

GetServerTypeOk returns a tuple with the ServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerType

`func (o *UpdateNodeTypeRequestContainerType) SetServerType(v string)`

SetServerType sets ServerType field to given value.

### HasServerType

`func (o *UpdateNodeTypeRequestContainerType) HasServerType() bool`

HasServerType returns a boolean if a field has been set.

### GetContainerPorts

`func (o *UpdateNodeTypeRequestContainerType) GetContainerPorts() []AddNodeTypeRequestContainerTypeContainerPortsInner`

GetContainerPorts returns the ContainerPorts field if non-nil, zero value otherwise.

### GetContainerPortsOk

`func (o *UpdateNodeTypeRequestContainerType) GetContainerPortsOk() (*[]AddNodeTypeRequestContainerTypeContainerPortsInner, bool)`

GetContainerPortsOk returns a tuple with the ContainerPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerPorts

`func (o *UpdateNodeTypeRequestContainerType) SetContainerPorts(v []AddNodeTypeRequestContainerTypeContainerPortsInner)`

SetContainerPorts sets ContainerPorts field to given value.

### HasContainerPorts

`func (o *UpdateNodeTypeRequestContainerType) HasContainerPorts() bool`

HasContainerPorts returns a boolean if a field has been set.

### GetEnvironmentVariables

`func (o *UpdateNodeTypeRequestContainerType) GetEnvironmentVariables() []AddClusterLayoutsRequestLayoutEnvironmentVariablesInner`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *UpdateNodeTypeRequestContainerType) GetEnvironmentVariablesOk() (*[]AddClusterLayoutsRequestLayoutEnvironmentVariablesInner, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *UpdateNodeTypeRequestContainerType) SetEnvironmentVariables(v []AddClusterLayoutsRequestLayoutEnvironmentVariablesInner)`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *UpdateNodeTypeRequestContainerType) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateNodeTypeRequestContainerType) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateNodeTypeRequestContainerType) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateNodeTypeRequestContainerType) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateNodeTypeRequestContainerType) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


