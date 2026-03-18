# GetClusterLayout200ResponseLayoutComputeServersInnerContainerType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to **NullableString** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**ShortName** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**ContainerVersion** | Pointer to **string** |  | [optional] 
**ProvisionType** | Pointer to [**GetClusterLayout200ResponseLayoutComputeServersInnerContainerTypeProvisionType**](GetClusterLayout200ResponseLayoutComputeServersInnerContainerTypeProvisionType.md) |  | [optional] 
**VirtualImage** | Pointer to **NullableString** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**ContainerPorts** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ContainerScripts** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ContainerTemplates** | Pointer to **[]map[string]interface{}** |  | [optional] 
**EnvironmentVariables** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewGetClusterLayout200ResponseLayoutComputeServersInnerContainerType

`func NewGetClusterLayout200ResponseLayoutComputeServersInnerContainerType() *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType`

NewGetClusterLayout200ResponseLayoutComputeServersInnerContainerType instantiates a new GetClusterLayout200ResponseLayoutComputeServersInnerContainerType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClusterLayout200ResponseLayoutComputeServersInnerContainerTypeWithDefaults

`func NewGetClusterLayout200ResponseLayoutComputeServersInnerContainerTypeWithDefaults() *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType`

NewGetClusterLayout200ResponseLayoutComputeServersInnerContainerTypeWithDefaults instantiates a new GetClusterLayout200ResponseLayoutComputeServersInnerContainerType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### SetAccountNil

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) SetAccountNil(b bool)`

 SetAccountNil sets the value for Account to be an explicit nil

### UnsetAccount
`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) UnsetAccount()`

UnsetAccount ensures that no value is present for Account, not even an explicit nil
### GetName

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetShortName

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetCode

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetContainerVersion

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) GetContainerVersion() string`

GetContainerVersion returns the ContainerVersion field if non-nil, zero value otherwise.

### GetContainerVersionOk

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) GetContainerVersionOk() (*string, bool)`

GetContainerVersionOk returns a tuple with the ContainerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerVersion

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) SetContainerVersion(v string)`

SetContainerVersion sets ContainerVersion field to given value.

### HasContainerVersion

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) HasContainerVersion() bool`

HasContainerVersion returns a boolean if a field has been set.

### GetProvisionType

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) GetProvisionType() GetClusterLayout200ResponseLayoutComputeServersInnerContainerTypeProvisionType`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) GetProvisionTypeOk() (*GetClusterLayout200ResponseLayoutComputeServersInnerContainerTypeProvisionType, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) SetProvisionType(v GetClusterLayout200ResponseLayoutComputeServersInnerContainerTypeProvisionType)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetVirtualImage

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) GetVirtualImage() string`

GetVirtualImage returns the VirtualImage field if non-nil, zero value otherwise.

### GetVirtualImageOk

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) GetVirtualImageOk() (*string, bool)`

GetVirtualImageOk returns a tuple with the VirtualImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImage

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) SetVirtualImage(v string)`

SetVirtualImage sets VirtualImage field to given value.

### HasVirtualImage

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) HasVirtualImage() bool`

HasVirtualImage returns a boolean if a field has been set.

### SetVirtualImageNil

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) SetVirtualImageNil(b bool)`

 SetVirtualImageNil sets the value for VirtualImage to be an explicit nil

### UnsetVirtualImage
`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) UnsetVirtualImage()`

UnsetVirtualImage ensures that no value is present for VirtualImage, not even an explicit nil
### GetCategory

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetConfig

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetContainerPorts

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) GetContainerPorts() []map[string]interface{}`

GetContainerPorts returns the ContainerPorts field if non-nil, zero value otherwise.

### GetContainerPortsOk

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) GetContainerPortsOk() (*[]map[string]interface{}, bool)`

GetContainerPortsOk returns a tuple with the ContainerPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerPorts

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) SetContainerPorts(v []map[string]interface{})`

SetContainerPorts sets ContainerPorts field to given value.

### HasContainerPorts

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) HasContainerPorts() bool`

HasContainerPorts returns a boolean if a field has been set.

### GetContainerScripts

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) GetContainerScripts() []map[string]interface{}`

GetContainerScripts returns the ContainerScripts field if non-nil, zero value otherwise.

### GetContainerScriptsOk

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) GetContainerScriptsOk() (*[]map[string]interface{}, bool)`

GetContainerScriptsOk returns a tuple with the ContainerScripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerScripts

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) SetContainerScripts(v []map[string]interface{})`

SetContainerScripts sets ContainerScripts field to given value.

### HasContainerScripts

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) HasContainerScripts() bool`

HasContainerScripts returns a boolean if a field has been set.

### GetContainerTemplates

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) GetContainerTemplates() []map[string]interface{}`

GetContainerTemplates returns the ContainerTemplates field if non-nil, zero value otherwise.

### GetContainerTemplatesOk

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) GetContainerTemplatesOk() (*[]map[string]interface{}, bool)`

GetContainerTemplatesOk returns a tuple with the ContainerTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTemplates

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) SetContainerTemplates(v []map[string]interface{})`

SetContainerTemplates sets ContainerTemplates field to given value.

### HasContainerTemplates

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) HasContainerTemplates() bool`

HasContainerTemplates returns a boolean if a field has been set.

### GetEnvironmentVariables

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) GetEnvironmentVariables() []map[string]interface{}`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) GetEnvironmentVariablesOk() (*[]map[string]interface{}, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) SetEnvironmentVariables(v []map[string]interface{})`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *GetClusterLayout200ResponseLayoutComputeServersInnerContainerType) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


