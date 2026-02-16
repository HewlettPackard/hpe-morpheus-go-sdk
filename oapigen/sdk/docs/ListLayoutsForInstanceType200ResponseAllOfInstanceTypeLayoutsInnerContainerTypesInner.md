# ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Account** | Pointer to [**ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerAccount**](ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerAccount.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**ShortName** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**ContainerVersion** | Pointer to **string** |  | [optional] 
**ProvisionType** | Pointer to [**ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerProvisionType**](ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerProvisionType.md) |  | [optional] 
**VirtualImage** | Pointer to [**ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerVirtualImage**](ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerVirtualImage.md) |  | [optional] 
**OsType** | Pointer to [**ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerOsType**](ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerOsType.md) |  | [optional] 
**Category** | Pointer to **NullableString** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**ContainerPorts** | Pointer to [**[]ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerContainerPortsInner**](ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerContainerPortsInner.md) |  | [optional] 
**ContainerScripts** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ContainerTemplates** | Pointer to **[]map[string]interface{}** |  | [optional] 
**EnvironmentVariables** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner

`func NewListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner() *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner`

NewListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner instantiates a new ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerWithDefaults

`func NewListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerWithDefaults() *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner`

NewListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerWithDefaults instantiates a new ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetAccount() ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetAccountOk() (*ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetAccount(v ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetName

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetShortName

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetCode

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetContainerVersion

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetContainerVersion() string`

GetContainerVersion returns the ContainerVersion field if non-nil, zero value otherwise.

### GetContainerVersionOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetContainerVersionOk() (*string, bool)`

GetContainerVersionOk returns a tuple with the ContainerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerVersion

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetContainerVersion(v string)`

SetContainerVersion sets ContainerVersion field to given value.

### HasContainerVersion

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) HasContainerVersion() bool`

HasContainerVersion returns a boolean if a field has been set.

### GetProvisionType

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetProvisionType() ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerProvisionType`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetProvisionTypeOk() (*ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerProvisionType, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetProvisionType(v ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerProvisionType)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetVirtualImage

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetVirtualImage() ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerVirtualImage`

GetVirtualImage returns the VirtualImage field if non-nil, zero value otherwise.

### GetVirtualImageOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetVirtualImageOk() (*ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerVirtualImage, bool)`

GetVirtualImageOk returns a tuple with the VirtualImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImage

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetVirtualImage(v ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerVirtualImage)`

SetVirtualImage sets VirtualImage field to given value.

### HasVirtualImage

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) HasVirtualImage() bool`

HasVirtualImage returns a boolean if a field has been set.

### GetOsType

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetOsType() ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerOsType`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetOsTypeOk() (*ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerOsType, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetOsType(v ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerOsType)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetCategory

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetConfig

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetContainerPorts

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetContainerPorts() []ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerContainerPortsInner`

GetContainerPorts returns the ContainerPorts field if non-nil, zero value otherwise.

### GetContainerPortsOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetContainerPortsOk() (*[]ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerContainerPortsInner, bool)`

GetContainerPortsOk returns a tuple with the ContainerPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerPorts

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetContainerPorts(v []ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInnerContainerPortsInner)`

SetContainerPorts sets ContainerPorts field to given value.

### HasContainerPorts

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) HasContainerPorts() bool`

HasContainerPorts returns a boolean if a field has been set.

### SetContainerPortsNil

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetContainerPortsNil(b bool)`

 SetContainerPortsNil sets the value for ContainerPorts to be an explicit nil

### UnsetContainerPorts
`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) UnsetContainerPorts()`

UnsetContainerPorts ensures that no value is present for ContainerPorts, not even an explicit nil
### GetContainerScripts

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetContainerScripts() []map[string]interface{}`

GetContainerScripts returns the ContainerScripts field if non-nil, zero value otherwise.

### GetContainerScriptsOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetContainerScriptsOk() (*[]map[string]interface{}, bool)`

GetContainerScriptsOk returns a tuple with the ContainerScripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerScripts

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetContainerScripts(v []map[string]interface{})`

SetContainerScripts sets ContainerScripts field to given value.

### HasContainerScripts

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) HasContainerScripts() bool`

HasContainerScripts returns a boolean if a field has been set.

### SetContainerScriptsNil

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetContainerScriptsNil(b bool)`

 SetContainerScriptsNil sets the value for ContainerScripts to be an explicit nil

### UnsetContainerScripts
`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) UnsetContainerScripts()`

UnsetContainerScripts ensures that no value is present for ContainerScripts, not even an explicit nil
### GetContainerTemplates

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetContainerTemplates() []map[string]interface{}`

GetContainerTemplates returns the ContainerTemplates field if non-nil, zero value otherwise.

### GetContainerTemplatesOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetContainerTemplatesOk() (*[]map[string]interface{}, bool)`

GetContainerTemplatesOk returns a tuple with the ContainerTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTemplates

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetContainerTemplates(v []map[string]interface{})`

SetContainerTemplates sets ContainerTemplates field to given value.

### HasContainerTemplates

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) HasContainerTemplates() bool`

HasContainerTemplates returns a boolean if a field has been set.

### SetContainerTemplatesNil

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetContainerTemplatesNil(b bool)`

 SetContainerTemplatesNil sets the value for ContainerTemplates to be an explicit nil

### UnsetContainerTemplates
`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) UnsetContainerTemplates()`

UnsetContainerTemplates ensures that no value is present for ContainerTemplates, not even an explicit nil
### GetEnvironmentVariables

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetEnvironmentVariables() []map[string]interface{}`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) GetEnvironmentVariablesOk() (*[]map[string]interface{}, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetEnvironmentVariables(v []map[string]interface{})`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### SetEnvironmentVariablesNil

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) SetEnvironmentVariablesNil(b bool)`

 SetEnvironmentVariablesNil sets the value for EnvironmentVariables to be an explicit nil

### UnsetEnvironmentVariables
`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner) UnsetEnvironmentVariables()`

UnsetEnvironmentVariables ensures that no value is present for EnvironmentVariables, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


