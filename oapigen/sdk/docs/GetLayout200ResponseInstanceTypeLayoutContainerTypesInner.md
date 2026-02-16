# GetLayout200ResponseInstanceTypeLayoutContainerTypesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] 
**Account** | Pointer to [**GetLayout200ResponseInstanceTypeLayoutContainerTypesInnerAccount**](GetLayout200ResponseInstanceTypeLayoutContainerTypesInnerAccount.md) |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**ShortName** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**ContainerVersion** | Pointer to **string** |  | [optional] 
**ProvisionType** | Pointer to [**GetLayout200ResponseInstanceTypeLayoutContainerTypesInnerProvisionType**](GetLayout200ResponseInstanceTypeLayoutContainerTypesInnerProvisionType.md) |  | [optional] 
**VirtualImage** | Pointer to [**GetLayout200ResponseInstanceTypeLayoutContainerTypesInnerVirtualImage**](GetLayout200ResponseInstanceTypeLayoutContainerTypesInnerVirtualImage.md) |  | [optional] 
**OsType** | Pointer to [**GetLayout200ResponseInstanceTypeLayoutContainerTypesInnerOsType**](GetLayout200ResponseInstanceTypeLayoutContainerTypesInnerOsType.md) |  | [optional] 
**Category** | Pointer to **NullableString** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**ContainerPorts** | Pointer to [**[]GetLayout200ResponseInstanceTypeLayoutContainerTypesInnerContainerPortsInner**](GetLayout200ResponseInstanceTypeLayoutContainerTypesInnerContainerPortsInner.md) |  | [optional] 
**ContainerScripts** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ContainerTemplates** | Pointer to **[]map[string]interface{}** |  | [optional] 
**EnvironmentVariables** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewGetLayout200ResponseInstanceTypeLayoutContainerTypesInner

`func NewGetLayout200ResponseInstanceTypeLayoutContainerTypesInner() *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner`

NewGetLayout200ResponseInstanceTypeLayoutContainerTypesInner instantiates a new GetLayout200ResponseInstanceTypeLayoutContainerTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLayout200ResponseInstanceTypeLayoutContainerTypesInnerWithDefaults

`func NewGetLayout200ResponseInstanceTypeLayoutContainerTypesInnerWithDefaults() *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner`

NewGetLayout200ResponseInstanceTypeLayoutContainerTypesInnerWithDefaults instantiates a new GetLayout200ResponseInstanceTypeLayoutContainerTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) GetAccount() GetLayout200ResponseInstanceTypeLayoutContainerTypesInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) GetAccountOk() (*GetLayout200ResponseInstanceTypeLayoutContainerTypesInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) SetAccount(v GetLayout200ResponseInstanceTypeLayoutContainerTypesInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetName

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetShortName

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetCode

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetContainerVersion

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) GetContainerVersion() string`

GetContainerVersion returns the ContainerVersion field if non-nil, zero value otherwise.

### GetContainerVersionOk

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) GetContainerVersionOk() (*string, bool)`

GetContainerVersionOk returns a tuple with the ContainerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerVersion

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) SetContainerVersion(v string)`

SetContainerVersion sets ContainerVersion field to given value.

### HasContainerVersion

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) HasContainerVersion() bool`

HasContainerVersion returns a boolean if a field has been set.

### GetProvisionType

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) GetProvisionType() GetLayout200ResponseInstanceTypeLayoutContainerTypesInnerProvisionType`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) GetProvisionTypeOk() (*GetLayout200ResponseInstanceTypeLayoutContainerTypesInnerProvisionType, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) SetProvisionType(v GetLayout200ResponseInstanceTypeLayoutContainerTypesInnerProvisionType)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetVirtualImage

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) GetVirtualImage() GetLayout200ResponseInstanceTypeLayoutContainerTypesInnerVirtualImage`

GetVirtualImage returns the VirtualImage field if non-nil, zero value otherwise.

### GetVirtualImageOk

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) GetVirtualImageOk() (*GetLayout200ResponseInstanceTypeLayoutContainerTypesInnerVirtualImage, bool)`

GetVirtualImageOk returns a tuple with the VirtualImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImage

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) SetVirtualImage(v GetLayout200ResponseInstanceTypeLayoutContainerTypesInnerVirtualImage)`

SetVirtualImage sets VirtualImage field to given value.

### HasVirtualImage

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) HasVirtualImage() bool`

HasVirtualImage returns a boolean if a field has been set.

### GetOsType

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) GetOsType() GetLayout200ResponseInstanceTypeLayoutContainerTypesInnerOsType`

GetOsType returns the OsType field if non-nil, zero value otherwise.

### GetOsTypeOk

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) GetOsTypeOk() (*GetLayout200ResponseInstanceTypeLayoutContainerTypesInnerOsType, bool)`

GetOsTypeOk returns a tuple with the OsType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOsType

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) SetOsType(v GetLayout200ResponseInstanceTypeLayoutContainerTypesInnerOsType)`

SetOsType sets OsType field to given value.

### HasOsType

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) HasOsType() bool`

HasOsType returns a boolean if a field has been set.

### GetCategory

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetConfig

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### SetConfigNil

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) SetConfigNil(b bool)`

 SetConfigNil sets the value for Config to be an explicit nil

### UnsetConfig
`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) UnsetConfig()`

UnsetConfig ensures that no value is present for Config, not even an explicit nil
### GetContainerPorts

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) GetContainerPorts() []GetLayout200ResponseInstanceTypeLayoutContainerTypesInnerContainerPortsInner`

GetContainerPorts returns the ContainerPorts field if non-nil, zero value otherwise.

### GetContainerPortsOk

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) GetContainerPortsOk() (*[]GetLayout200ResponseInstanceTypeLayoutContainerTypesInnerContainerPortsInner, bool)`

GetContainerPortsOk returns a tuple with the ContainerPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerPorts

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) SetContainerPorts(v []GetLayout200ResponseInstanceTypeLayoutContainerTypesInnerContainerPortsInner)`

SetContainerPorts sets ContainerPorts field to given value.

### HasContainerPorts

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) HasContainerPorts() bool`

HasContainerPorts returns a boolean if a field has been set.

### SetContainerPortsNil

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) SetContainerPortsNil(b bool)`

 SetContainerPortsNil sets the value for ContainerPorts to be an explicit nil

### UnsetContainerPorts
`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) UnsetContainerPorts()`

UnsetContainerPorts ensures that no value is present for ContainerPorts, not even an explicit nil
### GetContainerScripts

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) GetContainerScripts() []map[string]interface{}`

GetContainerScripts returns the ContainerScripts field if non-nil, zero value otherwise.

### GetContainerScriptsOk

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) GetContainerScriptsOk() (*[]map[string]interface{}, bool)`

GetContainerScriptsOk returns a tuple with the ContainerScripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerScripts

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) SetContainerScripts(v []map[string]interface{})`

SetContainerScripts sets ContainerScripts field to given value.

### HasContainerScripts

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) HasContainerScripts() bool`

HasContainerScripts returns a boolean if a field has been set.

### SetContainerScriptsNil

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) SetContainerScriptsNil(b bool)`

 SetContainerScriptsNil sets the value for ContainerScripts to be an explicit nil

### UnsetContainerScripts
`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) UnsetContainerScripts()`

UnsetContainerScripts ensures that no value is present for ContainerScripts, not even an explicit nil
### GetContainerTemplates

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) GetContainerTemplates() []map[string]interface{}`

GetContainerTemplates returns the ContainerTemplates field if non-nil, zero value otherwise.

### GetContainerTemplatesOk

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) GetContainerTemplatesOk() (*[]map[string]interface{}, bool)`

GetContainerTemplatesOk returns a tuple with the ContainerTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTemplates

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) SetContainerTemplates(v []map[string]interface{})`

SetContainerTemplates sets ContainerTemplates field to given value.

### HasContainerTemplates

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) HasContainerTemplates() bool`

HasContainerTemplates returns a boolean if a field has been set.

### SetContainerTemplatesNil

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) SetContainerTemplatesNil(b bool)`

 SetContainerTemplatesNil sets the value for ContainerTemplates to be an explicit nil

### UnsetContainerTemplates
`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) UnsetContainerTemplates()`

UnsetContainerTemplates ensures that no value is present for ContainerTemplates, not even an explicit nil
### GetEnvironmentVariables

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) GetEnvironmentVariables() []map[string]interface{}`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) GetEnvironmentVariablesOk() (*[]map[string]interface{}, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) SetEnvironmentVariables(v []map[string]interface{})`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### SetEnvironmentVariablesNil

`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) SetEnvironmentVariablesNil(b bool)`

 SetEnvironmentVariablesNil sets the value for EnvironmentVariables to be an explicit nil

### UnsetEnvironmentVariables
`func (o *GetLayout200ResponseInstanceTypeLayoutContainerTypesInner) UnsetEnvironmentVariables()`

UnsetEnvironmentVariables ensures that no value is present for EnvironmentVariables, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


