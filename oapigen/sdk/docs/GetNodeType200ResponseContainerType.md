# GetNodeType200ResponseContainerType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Account** | Pointer to [**GetNodeType200ResponseContainerTypeAccount**](GetNodeType200ResponseContainerTypeAccount.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**ShortName** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**ContainerVersion** | Pointer to **string** |  | [optional] 
**ProvisionType** | Pointer to [**GetNodeType200ResponseContainerTypeProvisionType**](GetNodeType200ResponseContainerTypeProvisionType.md) |  | [optional] 
**VirtualImage** | Pointer to [**GetNodeType200ResponseContainerTypeVirtualImage**](GetNodeType200ResponseContainerTypeVirtualImage.md) |  | [optional] 
**OsType** | Pointer to [**GetNodeType200ResponseContainerTypeOsType**](GetNodeType200ResponseContainerTypeOsType.md) |  | [optional] 
**Category** | Pointer to **NullableString** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**ContainerPorts** | Pointer to [**[]GetNodeType200ResponseContainerTypeContainerPortsInner**](GetNodeType200ResponseContainerTypeContainerPortsInner.md) |  | [optional] 
**ContainerScripts** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ContainerTemplates** | Pointer to **[]map[string]interface{}** |  | [optional] 
**EnvironmentVariables** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewGetNodeType200ResponseContainerType

`func NewGetNodeType200ResponseContainerType() *GetNodeType200ResponseContainerType`

NewGetNodeType200ResponseContainerType instantiates a new GetNodeType200ResponseContainerType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetNodeType200ResponseContainerType) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNodeType200ResponseContainerType) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNodeType200ResponseContainerType) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetNodeType200ResponseContainerType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *GetNodeType200ResponseContainerType) GetAccount() GetNodeType200ResponseContainerTypeAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetNodeType200ResponseContainerType) GetAccountOk() (*GetNodeType200ResponseContainerTypeAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetNodeType200ResponseContainerType) SetAccount(v GetNodeType200ResponseContainerTypeAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetNodeType200ResponseContainerType) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetName

`func (o *GetNodeType200ResponseContainerType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNodeType200ResponseContainerType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNodeType200ResponseContainerType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNodeType200ResponseContainerType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *GetNodeType200ResponseContainerType) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetNodeType200ResponseContainerType) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetNodeType200ResponseContainerType) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetNodeType200ResponseContainerType) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetShortName

`func (o *GetNodeType200ResponseContainerType) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *GetNodeType200ResponseContainerType) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *GetNodeType200ResponseContainerType) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *GetNodeType200ResponseContainerType) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetCode

`func (o *GetNodeType200ResponseContainerType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetNodeType200ResponseContainerType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetNodeType200ResponseContainerType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetNodeType200ResponseContainerType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetContainerVersion

`func (o *GetNodeType200ResponseContainerType) GetContainerVersion() string`

GetContainerVersion returns the ContainerVersion field if non-nil, zero value otherwise.

### GetContainerVersionOk

`func (o *GetNodeType200ResponseContainerType) GetContainerVersionOk() (*string, bool)`

GetContainerVersionOk returns a tuple with the ContainerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerVersion

`func (o *GetNodeType200ResponseContainerType) SetContainerVersion(v string)`

SetContainerVersion sets ContainerVersion field to given value.

### HasContainerVersion

`func (o *GetNodeType200ResponseContainerType) HasContainerVersion() bool`

HasContainerVersion returns a boolean if a field has been set.

### GetProvisionType

`func (o *GetNodeType200ResponseContainerType) GetProvisionType() GetNodeType200ResponseContainerTypeProvisionType`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *GetNodeType200ResponseContainerType) GetProvisionTypeOk() (*GetNodeType200ResponseContainerTypeProvisionType, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *GetNodeType200ResponseContainerType) SetProvisionType(v GetNodeType200ResponseContainerTypeProvisionType)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *GetNodeType200ResponseContainerType) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetVirtualImage

`func (o *GetNodeType200ResponseContainerType) GetVirtualImage() GetNodeType200ResponseContainerTypeVirtualImage`

GetVirtualImage returns the VirtualImage field if non-nil, zero value otherwise.

### GetVirtualImageOk

`func (o *GetNodeType200ResponseContainerType) GetVirtualImageOk() (*GetNodeType200ResponseContainerTypeVirtualImage, bool)`

GetVirtualImageOk returns a tuple with the VirtualImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImage

`func (o *GetNodeType200ResponseContainerType) SetVirtualImage(v GetNodeType200ResponseContainerTypeVirtualImage)`

SetVirtualImage sets VirtualImage field to given value.

### HasVirtualImage

`func (o *GetNodeType200ResponseContainerType) HasVirtualImage() bool`

HasVirtualImage returns a boolean if a field has been set.

### GetOsType

`func (o *GetNodeType200ResponseContainerType) GetOsType() GetNodeType200ResponseContainerTypeOsType`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *GetNodeType200ResponseContainerType) GetOsTypeOk() (*GetNodeType200ResponseContainerTypeOsType, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *GetNodeType200ResponseContainerType) SetOsType(v GetNodeType200ResponseContainerTypeOsType)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *GetNodeType200ResponseContainerType) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetCategory

`func (o *GetNodeType200ResponseContainerType) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetNodeType200ResponseContainerType) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetNodeType200ResponseContainerType) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetNodeType200ResponseContainerType) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *GetNodeType200ResponseContainerType) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *GetNodeType200ResponseContainerType) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetConfig

`func (o *GetNodeType200ResponseContainerType) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetNodeType200ResponseContainerType) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetNodeType200ResponseContainerType) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetNodeType200ResponseContainerType) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *GetNodeType200ResponseContainerType) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *GetNodeType200ResponseContainerType) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetContainerPorts

`func (o *GetNodeType200ResponseContainerType) GetContainerPorts() []GetNodeType200ResponseContainerTypeContainerPortsInner`

GetContainerPorts returns the ContainerPorts field if non-nil, zero value otherwise.

### GetContainerPortsOk

`func (o *GetNodeType200ResponseContainerType) GetContainerPortsOk() (*[]GetNodeType200ResponseContainerTypeContainerPortsInner, bool)`

GetContainerPortsOk returns a tuple with the ContainerPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerPorts

`func (o *GetNodeType200ResponseContainerType) SetContainerPorts(v []GetNodeType200ResponseContainerTypeContainerPortsInner)`

SetContainerPorts sets ContainerPorts field to given value.

### HasContainerPorts

`func (o *GetNodeType200ResponseContainerType) HasContainerPorts() bool`

HasContainerPorts returns a boolean if a field has been set.

### SetContainerPortsNil

`func (o *GetNodeType200ResponseContainerType) SetContainerPortsNil(b bool)`

 SetContainerPortsNil sets the value for ContainerPorts to be an explicit nil

### UnsetContainerPorts
`func (o *GetNodeType200ResponseContainerType) UnsetContainerPorts()`

UnsetContainerPorts ensures that no value is present for ContainerPorts, not even an explicit nil
### GetContainerScripts

`func (o *GetNodeType200ResponseContainerType) GetContainerScripts() []map[string]interface{}`

GetContainerScripts returns the ContainerScripts field if non-nil, zero value otherwise.

### GetContainerScriptsOk

`func (o *GetNodeType200ResponseContainerType) GetContainerScriptsOk() (*[]map[string]interface{}, bool)`

GetContainerScriptsOk returns a tuple with the ContainerScripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerScripts

`func (o *GetNodeType200ResponseContainerType) SetContainerScripts(v []map[string]interface{})`

SetContainerScripts sets ContainerScripts field to given value.

### HasContainerScripts

`func (o *GetNodeType200ResponseContainerType) HasContainerScripts() bool`

HasContainerScripts returns a boolean if a field has been set.

### SetContainerScriptsNil

`func (o *GetNodeType200ResponseContainerType) SetContainerScriptsNil(b bool)`

 SetContainerScriptsNil sets the value for ContainerScripts to be an explicit nil

### UnsetContainerScripts
`func (o *GetNodeType200ResponseContainerType) UnsetContainerScripts()`

UnsetContainerScripts ensures that no value is present for ContainerScripts, not even an explicit nil
### GetContainerTemplates

`func (o *GetNodeType200ResponseContainerType) GetContainerTemplates() []map[string]interface{}`

GetContainerTemplates returns the ContainerTemplates field if non-nil, zero value otherwise.

### GetContainerTemplatesOk

`func (o *GetNodeType200ResponseContainerType) GetContainerTemplatesOk() (*[]map[string]interface{}, bool)`

GetContainerTemplatesOk returns a tuple with the ContainerTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTemplates

`func (o *GetNodeType200ResponseContainerType) SetContainerTemplates(v []map[string]interface{})`

SetContainerTemplates sets ContainerTemplates field to given value.

### HasContainerTemplates

`func (o *GetNodeType200ResponseContainerType) HasContainerTemplates() bool`

HasContainerTemplates returns a boolean if a field has been set.

### SetContainerTemplatesNil

`func (o *GetNodeType200ResponseContainerType) SetContainerTemplatesNil(b bool)`

 SetContainerTemplatesNil sets the value for ContainerTemplates to be an explicit nil

### UnsetContainerTemplates
`func (o *GetNodeType200ResponseContainerType) UnsetContainerTemplates()`

UnsetContainerTemplates ensures that no value is present for ContainerTemplates, not even an explicit nil
### GetEnvironmentVariables

`func (o *GetNodeType200ResponseContainerType) GetEnvironmentVariables() []map[string]interface{}`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *GetNodeType200ResponseContainerType) GetEnvironmentVariablesOk() (*[]map[string]interface{}, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *GetNodeType200ResponseContainerType) SetEnvironmentVariables(v []map[string]interface{})`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *GetNodeType200ResponseContainerType) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### SetEnvironmentVariablesNil

`func (o *GetNodeType200ResponseContainerType) SetEnvironmentVariablesNil(b bool)`

 SetEnvironmentVariablesNil sets the value for EnvironmentVariables to be an explicit nil

### UnsetEnvironmentVariables
`func (o *GetNodeType200ResponseContainerType) UnsetEnvironmentVariables()`

UnsetEnvironmentVariables ensures that no value is present for EnvironmentVariables, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


