# InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Account** | Pointer to [**InstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerAccount**](InstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerAccount.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**ShortName** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**ContainerVersion** | Pointer to **string** |  | [optional] 
**ProvisionType** | Pointer to [**InstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerProvisionType**](InstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerProvisionType.md) |  | [optional] 
**VirtualImage** | Pointer to [**InstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerVirtualImage**](InstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerVirtualImage.md) |  | [optional] 
**OsType** | Pointer to [**InstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerOsType**](InstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerOsType.md) |  | [optional] 
**Category** | Pointer to **NullableString** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**ContainerPorts** | Pointer to [**[]InstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerContainerPortsInner**](InstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerContainerPortsInner.md) |  | [optional] 
**ContainerScripts** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ContainerTemplates** | Pointer to **[]map[string]interface{}** |  | [optional] 
**EnvironmentVariables** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner

`func NewInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner() *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner`

NewInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner instantiates a new InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerWithDefaults

`func NewInstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerWithDefaults() *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner`

NewInstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerWithDefaults instantiates a new InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetAccount() InstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetAccountOk() (*InstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetAccount(v InstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetName

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetShortName

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetCode

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetContainerVersion

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetContainerVersion() string`

GetContainerVersion returns the ContainerVersion field if non-nil, zero value otherwise.

### GetContainerVersionOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetContainerVersionOk() (*string, bool)`

GetContainerVersionOk returns a tuple with the ContainerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerVersion

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetContainerVersion(v string)`

SetContainerVersion sets ContainerVersion field to given value.

### HasContainerVersion

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) HasContainerVersion() bool`

HasContainerVersion returns a boolean if a field has been set.

### GetProvisionType

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetProvisionType() InstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerProvisionType`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetProvisionTypeOk() (*InstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerProvisionType, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetProvisionType(v InstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerProvisionType)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetVirtualImage

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetVirtualImage() InstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerVirtualImage`

GetVirtualImage returns the VirtualImage field if non-nil, zero value otherwise.

### GetVirtualImageOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetVirtualImageOk() (*InstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerVirtualImage, bool)`

GetVirtualImageOk returns a tuple with the VirtualImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImage

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetVirtualImage(v InstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerVirtualImage)`

SetVirtualImage sets VirtualImage field to given value.

### HasVirtualImage

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) HasVirtualImage() bool`

HasVirtualImage returns a boolean if a field has been set.

### GetOsType

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetOsType() InstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerOsType`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetOsTypeOk() (*InstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerOsType, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetOsType(v InstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerOsType)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetCategory

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetConfig

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetContainerPorts

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetContainerPorts() []InstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerContainerPortsInner`

GetContainerPorts returns the ContainerPorts field if non-nil, zero value otherwise.

### GetContainerPortsOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetContainerPortsOk() (*[]InstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerContainerPortsInner, bool)`

GetContainerPortsOk returns a tuple with the ContainerPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerPorts

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetContainerPorts(v []InstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerContainerPortsInner)`

SetContainerPorts sets ContainerPorts field to given value.

### HasContainerPorts

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) HasContainerPorts() bool`

HasContainerPorts returns a boolean if a field has been set.

### SetContainerPortsNil

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetContainerPortsNil(b bool)`

 SetContainerPortsNil sets the value for ContainerPorts to be an explicit nil

### UnsetContainerPorts
`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) UnsetContainerPorts()`

UnsetContainerPorts ensures that no value is present for ContainerPorts, not even an explicit nil
### GetContainerScripts

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetContainerScripts() []map[string]interface{}`

GetContainerScripts returns the ContainerScripts field if non-nil, zero value otherwise.

### GetContainerScriptsOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetContainerScriptsOk() (*[]map[string]interface{}, bool)`

GetContainerScriptsOk returns a tuple with the ContainerScripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerScripts

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetContainerScripts(v []map[string]interface{})`

SetContainerScripts sets ContainerScripts field to given value.

### HasContainerScripts

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) HasContainerScripts() bool`

HasContainerScripts returns a boolean if a field has been set.

### SetContainerScriptsNil

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetContainerScriptsNil(b bool)`

 SetContainerScriptsNil sets the value for ContainerScripts to be an explicit nil

### UnsetContainerScripts
`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) UnsetContainerScripts()`

UnsetContainerScripts ensures that no value is present for ContainerScripts, not even an explicit nil
### GetContainerTemplates

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetContainerTemplates() []map[string]interface{}`

GetContainerTemplates returns the ContainerTemplates field if non-nil, zero value otherwise.

### GetContainerTemplatesOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetContainerTemplatesOk() (*[]map[string]interface{}, bool)`

GetContainerTemplatesOk returns a tuple with the ContainerTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTemplates

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetContainerTemplates(v []map[string]interface{})`

SetContainerTemplates sets ContainerTemplates field to given value.

### HasContainerTemplates

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) HasContainerTemplates() bool`

HasContainerTemplates returns a boolean if a field has been set.

### SetContainerTemplatesNil

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetContainerTemplatesNil(b bool)`

 SetContainerTemplatesNil sets the value for ContainerTemplates to be an explicit nil

### UnsetContainerTemplates
`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) UnsetContainerTemplates()`

UnsetContainerTemplates ensures that no value is present for ContainerTemplates, not even an explicit nil
### GetEnvironmentVariables

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetEnvironmentVariables() []map[string]interface{}`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetEnvironmentVariablesOk() (*[]map[string]interface{}, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetEnvironmentVariables(v []map[string]interface{})`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### SetEnvironmentVariablesNil

`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetEnvironmentVariablesNil(b bool)`

 SetEnvironmentVariablesNil sets the value for EnvironmentVariables to be an explicit nil

### UnsetEnvironmentVariables
`func (o *InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) UnsetEnvironmentVariables()`

UnsetEnvironmentVariables ensures that no value is present for EnvironmentVariables, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


