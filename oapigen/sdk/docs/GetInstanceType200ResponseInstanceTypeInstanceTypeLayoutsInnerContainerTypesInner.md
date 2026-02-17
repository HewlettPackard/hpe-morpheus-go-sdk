# GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Account** | Pointer to [**GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerAccount**](GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerAccount.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**ShortName** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**ContainerVersion** | Pointer to **string** |  | [optional] 
**ProvisionType** | Pointer to [**GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerProvisionType**](GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerProvisionType.md) |  | [optional] 
**VirtualImage** | Pointer to [**GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerVirtualImage**](GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerVirtualImage.md) |  | [optional] 
**OsType** | Pointer to [**GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerOsType**](GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerOsType.md) |  | [optional] 
**Category** | Pointer to **NullableString** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**ContainerPorts** | Pointer to [**[]GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerContainerPortsInner**](GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerContainerPortsInner.md) |  | [optional] 
**ContainerScripts** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ContainerTemplates** | Pointer to **[]map[string]interface{}** |  | [optional] 
**EnvironmentVariables** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewGetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner

`func NewGetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner() *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner`

NewGetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner instantiates a new GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerWithDefaults

`func NewGetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerWithDefaults() *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner`

NewGetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerWithDefaults instantiates a new GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetAccount() GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetAccountOk() (*GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetAccount(v GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetName

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetShortName

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetCode

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetContainerVersion

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetContainerVersion() string`

GetContainerVersion returns the ContainerVersion field if non-nil, zero value otherwise.

### GetContainerVersionOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetContainerVersionOk() (*string, bool)`

GetContainerVersionOk returns a tuple with the ContainerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerVersion

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetContainerVersion(v string)`

SetContainerVersion sets ContainerVersion field to given value.

### HasContainerVersion

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) HasContainerVersion() bool`

HasContainerVersion returns a boolean if a field has been set.

### GetProvisionType

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetProvisionType() GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerProvisionType`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetProvisionTypeOk() (*GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerProvisionType, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetProvisionType(v GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerProvisionType)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetVirtualImage

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetVirtualImage() GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerVirtualImage`

GetVirtualImage returns the VirtualImage field if non-nil, zero value otherwise.

### GetVirtualImageOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetVirtualImageOk() (*GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerVirtualImage, bool)`

GetVirtualImageOk returns a tuple with the VirtualImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImage

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetVirtualImage(v GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerVirtualImage)`

SetVirtualImage sets VirtualImage field to given value.

### HasVirtualImage

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) HasVirtualImage() bool`

HasVirtualImage returns a boolean if a field has been set.

### GetOsType

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetOsType() GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerOsType`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetOsTypeOk() (*GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerOsType, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetOsType(v GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerOsType)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetCategory

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetConfig

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetContainerPorts

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetContainerPorts() []GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerContainerPortsInner`

GetContainerPorts returns the ContainerPorts field if non-nil, zero value otherwise.

### GetContainerPortsOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetContainerPortsOk() (*[]GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerContainerPortsInner, bool)`

GetContainerPortsOk returns a tuple with the ContainerPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerPorts

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetContainerPorts(v []GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInnerContainerPortsInner)`

SetContainerPorts sets ContainerPorts field to given value.

### HasContainerPorts

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) HasContainerPorts() bool`

HasContainerPorts returns a boolean if a field has been set.

### SetContainerPortsNil

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetContainerPortsNil(b bool)`

 SetContainerPortsNil sets the value for ContainerPorts to be an explicit nil

### UnsetContainerPorts
`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) UnsetContainerPorts()`

UnsetContainerPorts ensures that no value is present for ContainerPorts, not even an explicit nil
### GetContainerScripts

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetContainerScripts() []map[string]interface{}`

GetContainerScripts returns the ContainerScripts field if non-nil, zero value otherwise.

### GetContainerScriptsOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetContainerScriptsOk() (*[]map[string]interface{}, bool)`

GetContainerScriptsOk returns a tuple with the ContainerScripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerScripts

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetContainerScripts(v []map[string]interface{})`

SetContainerScripts sets ContainerScripts field to given value.

### HasContainerScripts

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) HasContainerScripts() bool`

HasContainerScripts returns a boolean if a field has been set.

### SetContainerScriptsNil

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetContainerScriptsNil(b bool)`

 SetContainerScriptsNil sets the value for ContainerScripts to be an explicit nil

### UnsetContainerScripts
`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) UnsetContainerScripts()`

UnsetContainerScripts ensures that no value is present for ContainerScripts, not even an explicit nil
### GetContainerTemplates

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetContainerTemplates() []map[string]interface{}`

GetContainerTemplates returns the ContainerTemplates field if non-nil, zero value otherwise.

### GetContainerTemplatesOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetContainerTemplatesOk() (*[]map[string]interface{}, bool)`

GetContainerTemplatesOk returns a tuple with the ContainerTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTemplates

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetContainerTemplates(v []map[string]interface{})`

SetContainerTemplates sets ContainerTemplates field to given value.

### HasContainerTemplates

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) HasContainerTemplates() bool`

HasContainerTemplates returns a boolean if a field has been set.

### SetContainerTemplatesNil

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetContainerTemplatesNil(b bool)`

 SetContainerTemplatesNil sets the value for ContainerTemplates to be an explicit nil

### UnsetContainerTemplates
`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) UnsetContainerTemplates()`

UnsetContainerTemplates ensures that no value is present for ContainerTemplates, not even an explicit nil
### GetEnvironmentVariables

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetEnvironmentVariables() []map[string]interface{}`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) GetEnvironmentVariablesOk() (*[]map[string]interface{}, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetEnvironmentVariables(v []map[string]interface{})`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### SetEnvironmentVariablesNil

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) SetEnvironmentVariablesNil(b bool)`

 SetEnvironmentVariablesNil sets the value for EnvironmentVariables to be an explicit nil

### UnsetEnvironmentVariables
`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner) UnsetEnvironmentVariables()`

UnsetEnvironmentVariables ensures that no value is present for EnvironmentVariables, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


