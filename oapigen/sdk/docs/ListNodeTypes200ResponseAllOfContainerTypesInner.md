# ListNodeTypes200ResponseAllOfContainerTypesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Account** | Pointer to [**ListNodeTypes200ResponseAllOfContainerTypesInnerAccount**](ListNodeTypes200ResponseAllOfContainerTypesInnerAccount.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**ShortName** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**ContainerVersion** | Pointer to **string** |  | [optional] 
**ProvisionType** | Pointer to [**ListNodeTypes200ResponseAllOfContainerTypesInnerProvisionType**](ListNodeTypes200ResponseAllOfContainerTypesInnerProvisionType.md) |  | [optional] 
**VirtualImage** | Pointer to [**ListNodeTypes200ResponseAllOfContainerTypesInnerVirtualImage**](ListNodeTypes200ResponseAllOfContainerTypesInnerVirtualImage.md) |  | [optional] 
**OsType** | Pointer to [**ListNodeTypes200ResponseAllOfContainerTypesInnerOsType**](ListNodeTypes200ResponseAllOfContainerTypesInnerOsType.md) |  | [optional] 
**Category** | Pointer to **NullableString** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**ContainerPorts** | Pointer to [**[]ListNodeTypes200ResponseAllOfContainerTypesInnerContainerPortsInner**](ListNodeTypes200ResponseAllOfContainerTypesInnerContainerPortsInner.md) |  | [optional] 
**ContainerScripts** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ContainerTemplates** | Pointer to **[]map[string]interface{}** |  | [optional] 
**EnvironmentVariables** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewListNodeTypes200ResponseAllOfContainerTypesInner

`func NewListNodeTypes200ResponseAllOfContainerTypesInner() *ListNodeTypes200ResponseAllOfContainerTypesInner`

NewListNodeTypes200ResponseAllOfContainerTypesInner instantiates a new ListNodeTypes200ResponseAllOfContainerTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) GetAccount() ListNodeTypes200ResponseAllOfContainerTypesInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) GetAccountOk() (*ListNodeTypes200ResponseAllOfContainerTypesInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) SetAccount(v ListNodeTypes200ResponseAllOfContainerTypesInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetName

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetShortName

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetCode

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetContainerVersion

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) GetContainerVersion() string`

GetContainerVersion returns the ContainerVersion field if non-nil, zero value otherwise.

### GetContainerVersionOk

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) GetContainerVersionOk() (*string, bool)`

GetContainerVersionOk returns a tuple with the ContainerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerVersion

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) SetContainerVersion(v string)`

SetContainerVersion sets ContainerVersion field to given value.

### HasContainerVersion

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) HasContainerVersion() bool`

HasContainerVersion returns a boolean if a field has been set.

### GetProvisionType

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) GetProvisionType() ListNodeTypes200ResponseAllOfContainerTypesInnerProvisionType`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) GetProvisionTypeOk() (*ListNodeTypes200ResponseAllOfContainerTypesInnerProvisionType, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) SetProvisionType(v ListNodeTypes200ResponseAllOfContainerTypesInnerProvisionType)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetVirtualImage

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) GetVirtualImage() ListNodeTypes200ResponseAllOfContainerTypesInnerVirtualImage`

GetVirtualImage returns the VirtualImage field if non-nil, zero value otherwise.

### GetVirtualImageOk

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) GetVirtualImageOk() (*ListNodeTypes200ResponseAllOfContainerTypesInnerVirtualImage, bool)`

GetVirtualImageOk returns a tuple with the VirtualImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImage

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) SetVirtualImage(v ListNodeTypes200ResponseAllOfContainerTypesInnerVirtualImage)`

SetVirtualImage sets VirtualImage field to given value.

### HasVirtualImage

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) HasVirtualImage() bool`

HasVirtualImage returns a boolean if a field has been set.

### GetOsType

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) GetOsType() ListNodeTypes200ResponseAllOfContainerTypesInnerOsType`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) GetOsTypeOk() (*ListNodeTypes200ResponseAllOfContainerTypesInnerOsType, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) SetOsType(v ListNodeTypes200ResponseAllOfContainerTypesInnerOsType)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetCategory

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetConfig

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetContainerPorts

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) GetContainerPorts() []ListNodeTypes200ResponseAllOfContainerTypesInnerContainerPortsInner`

GetContainerPorts returns the ContainerPorts field if non-nil, zero value otherwise.

### GetContainerPortsOk

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) GetContainerPortsOk() (*[]ListNodeTypes200ResponseAllOfContainerTypesInnerContainerPortsInner, bool)`

GetContainerPortsOk returns a tuple with the ContainerPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerPorts

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) SetContainerPorts(v []ListNodeTypes200ResponseAllOfContainerTypesInnerContainerPortsInner)`

SetContainerPorts sets ContainerPorts field to given value.

### HasContainerPorts

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) HasContainerPorts() bool`

HasContainerPorts returns a boolean if a field has been set.

### SetContainerPortsNil

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) SetContainerPortsNil(b bool)`

 SetContainerPortsNil sets the value for ContainerPorts to be an explicit nil

### UnsetContainerPorts
`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) UnsetContainerPorts()`

UnsetContainerPorts ensures that no value is present for ContainerPorts, not even an explicit nil
### GetContainerScripts

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) GetContainerScripts() []map[string]interface{}`

GetContainerScripts returns the ContainerScripts field if non-nil, zero value otherwise.

### GetContainerScriptsOk

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) GetContainerScriptsOk() (*[]map[string]interface{}, bool)`

GetContainerScriptsOk returns a tuple with the ContainerScripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerScripts

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) SetContainerScripts(v []map[string]interface{})`

SetContainerScripts sets ContainerScripts field to given value.

### HasContainerScripts

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) HasContainerScripts() bool`

HasContainerScripts returns a boolean if a field has been set.

### SetContainerScriptsNil

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) SetContainerScriptsNil(b bool)`

 SetContainerScriptsNil sets the value for ContainerScripts to be an explicit nil

### UnsetContainerScripts
`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) UnsetContainerScripts()`

UnsetContainerScripts ensures that no value is present for ContainerScripts, not even an explicit nil
### GetContainerTemplates

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) GetContainerTemplates() []map[string]interface{}`

GetContainerTemplates returns the ContainerTemplates field if non-nil, zero value otherwise.

### GetContainerTemplatesOk

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) GetContainerTemplatesOk() (*[]map[string]interface{}, bool)`

GetContainerTemplatesOk returns a tuple with the ContainerTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTemplates

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) SetContainerTemplates(v []map[string]interface{})`

SetContainerTemplates sets ContainerTemplates field to given value.

### HasContainerTemplates

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) HasContainerTemplates() bool`

HasContainerTemplates returns a boolean if a field has been set.

### SetContainerTemplatesNil

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) SetContainerTemplatesNil(b bool)`

 SetContainerTemplatesNil sets the value for ContainerTemplates to be an explicit nil

### UnsetContainerTemplates
`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) UnsetContainerTemplates()`

UnsetContainerTemplates ensures that no value is present for ContainerTemplates, not even an explicit nil
### GetEnvironmentVariables

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) GetEnvironmentVariables() []map[string]interface{}`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) GetEnvironmentVariablesOk() (*[]map[string]interface{}, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) SetEnvironmentVariables(v []map[string]interface{})`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### SetEnvironmentVariablesNil

`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) SetEnvironmentVariablesNil(b bool)`

 SetEnvironmentVariablesNil sets the value for EnvironmentVariables to be an explicit nil

### UnsetEnvironmentVariables
`func (o *ListNodeTypes200ResponseAllOfContainerTypesInner) UnsetEnvironmentVariables()`

UnsetEnvironmentVariables ensures that no value is present for EnvironmentVariables, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


