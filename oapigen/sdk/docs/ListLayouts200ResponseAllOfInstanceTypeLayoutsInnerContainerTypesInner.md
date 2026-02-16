# ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Account** | Pointer to [**ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerAccount**](ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerAccount.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**ShortName** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**ContainerVersion** | Pointer to **string** |  | [optional] 
**ProvisionType** | Pointer to [**ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerProvisionType**](ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerProvisionType.md) |  | [optional] 
**VirtualImage** | Pointer to [**ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerVirtualImage**](ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerVirtualImage.md) |  | [optional] 
**OsType** | Pointer to [**ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerOsType**](ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerOsType.md) |  | [optional] 
**Category** | Pointer to **NullableString** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**ContainerPorts** | Pointer to [**[]ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerContainerPortsInner**](ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerContainerPortsInner.md) |  | [optional] 
**ContainerScripts** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ContainerTemplates** | Pointer to **[]map[string]interface{}** |  | [optional] 
**EnvironmentVariables** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner

`func NewListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner() *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner`

NewListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner instantiates a new ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerWithDefaults

`func NewListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerWithDefaults() *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner`

NewListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerWithDefaults instantiates a new ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetAccount() ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetAccountOk() (*ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetAccount(v ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetName

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetShortName

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetCode

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetContainerVersion

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetContainerVersion() string`

GetContainerVersion returns the ContainerVersion field if non-nil, zero value otherwise.

### GetContainerVersionOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetContainerVersionOk() (*string, bool)`

GetContainerVersionOk returns a tuple with the ContainerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerVersion

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetContainerVersion(v string)`

SetContainerVersion sets ContainerVersion field to given value.

### HasContainerVersion

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) HasContainerVersion() bool`

HasContainerVersion returns a boolean if a field has been set.

### GetProvisionType

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetProvisionType() ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerProvisionType`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetProvisionTypeOk() (*ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerProvisionType, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetProvisionType(v ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerProvisionType)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetVirtualImage

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetVirtualImage() ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerVirtualImage`

GetVirtualImage returns the VirtualImage field if non-nil, zero value otherwise.

### GetVirtualImageOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetVirtualImageOk() (*ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerVirtualImage, bool)`

GetVirtualImageOk returns a tuple with the VirtualImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImage

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetVirtualImage(v ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerVirtualImage)`

SetVirtualImage sets VirtualImage field to given value.

### HasVirtualImage

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) HasVirtualImage() bool`

HasVirtualImage returns a boolean if a field has been set.

### GetOsType

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetOsType() ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerOsType`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetOsTypeOk() (*ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerOsType, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetOsType(v ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerOsType)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetCategory

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetConfig

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetContainerPorts

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetContainerPorts() []ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerContainerPortsInner`

GetContainerPorts returns the ContainerPorts field if non-nil, zero value otherwise.

### GetContainerPortsOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetContainerPortsOk() (*[]ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerContainerPortsInner, bool)`

GetContainerPortsOk returns a tuple with the ContainerPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerPorts

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetContainerPorts(v []ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerContainerPortsInner)`

SetContainerPorts sets ContainerPorts field to given value.

### HasContainerPorts

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) HasContainerPorts() bool`

HasContainerPorts returns a boolean if a field has been set.

### SetContainerPortsNil

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetContainerPortsNil(b bool)`

 SetContainerPortsNil sets the value for ContainerPorts to be an explicit nil

### UnsetContainerPorts
`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) UnsetContainerPorts()`

UnsetContainerPorts ensures that no value is present for ContainerPorts, not even an explicit nil
### GetContainerScripts

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetContainerScripts() []map[string]interface{}`

GetContainerScripts returns the ContainerScripts field if non-nil, zero value otherwise.

### GetContainerScriptsOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetContainerScriptsOk() (*[]map[string]interface{}, bool)`

GetContainerScriptsOk returns a tuple with the ContainerScripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerScripts

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetContainerScripts(v []map[string]interface{})`

SetContainerScripts sets ContainerScripts field to given value.

### HasContainerScripts

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) HasContainerScripts() bool`

HasContainerScripts returns a boolean if a field has been set.

### SetContainerScriptsNil

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetContainerScriptsNil(b bool)`

 SetContainerScriptsNil sets the value for ContainerScripts to be an explicit nil

### UnsetContainerScripts
`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) UnsetContainerScripts()`

UnsetContainerScripts ensures that no value is present for ContainerScripts, not even an explicit nil
### GetContainerTemplates

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetContainerTemplates() []map[string]interface{}`

GetContainerTemplates returns the ContainerTemplates field if non-nil, zero value otherwise.

### GetContainerTemplatesOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetContainerTemplatesOk() (*[]map[string]interface{}, bool)`

GetContainerTemplatesOk returns a tuple with the ContainerTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTemplates

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetContainerTemplates(v []map[string]interface{})`

SetContainerTemplates sets ContainerTemplates field to given value.

### HasContainerTemplates

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) HasContainerTemplates() bool`

HasContainerTemplates returns a boolean if a field has been set.

### SetContainerTemplatesNil

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetContainerTemplatesNil(b bool)`

 SetContainerTemplatesNil sets the value for ContainerTemplates to be an explicit nil

### UnsetContainerTemplates
`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) UnsetContainerTemplates()`

UnsetContainerTemplates ensures that no value is present for ContainerTemplates, not even an explicit nil
### GetEnvironmentVariables

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetEnvironmentVariables() []map[string]interface{}`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetEnvironmentVariablesOk() (*[]map[string]interface{}, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetEnvironmentVariables(v []map[string]interface{})`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### SetEnvironmentVariablesNil

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetEnvironmentVariablesNil(b bool)`

 SetEnvironmentVariablesNil sets the value for EnvironmentVariables to be an explicit nil

### UnsetEnvironmentVariables
`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) UnsetEnvironmentVariables()`

UnsetEnvironmentVariables ensures that no value is present for EnvironmentVariables, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


