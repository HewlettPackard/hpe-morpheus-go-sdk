# ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Account** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**ShortName** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**ContainerVersion** | Pointer to **string** |  | [optional] 
**ProvisionType** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**VirtualImage** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 
**ContainerPorts** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ContainerScripts** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ContainerTemplates** | Pointer to **[]map[string]interface{}** |  | [optional] 
**EnvironmentVariables** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType

`func NewListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType() *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType`

NewListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType instantiates a new ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerTypeWithDefaults

`func NewListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerTypeWithDefaults() *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType`

NewListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerTypeWithDefaults instantiates a new ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAccount

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetName

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetShortName

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) GetShortName() string`

GetShortName returns the ShortName field if non-nil, zero value otherwise.

### GetShortNameOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) GetShortNameOk() (*string, bool)`

GetShortNameOk returns a tuple with the ShortName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortName

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) SetShortName(v string)`

SetShortName sets ShortName field to given value.

### HasShortName

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) HasShortName() bool`

HasShortName returns a boolean if a field has been set.

### GetCode

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetContainerVersion

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) GetContainerVersion() string`

GetContainerVersion returns the ContainerVersion field if non-nil, zero value otherwise.

### GetContainerVersionOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) GetContainerVersionOk() (*string, bool)`

GetContainerVersionOk returns a tuple with the ContainerVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerVersion

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) SetContainerVersion(v string)`

SetContainerVersion sets ContainerVersion field to given value.

### HasContainerVersion

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) HasContainerVersion() bool`

HasContainerVersion returns a boolean if a field has been set.

### GetProvisionType

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) GetProvisionType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) GetProvisionTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) SetProvisionType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetVirtualImage

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) GetVirtualImage() string`

GetVirtualImage returns the VirtualImage field if non-nil, zero value otherwise.

### GetVirtualImageOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) GetVirtualImageOk() (*string, bool)`

GetVirtualImageOk returns a tuple with the VirtualImage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVirtualImage

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) SetVirtualImage(v string)`

SetVirtualImage sets VirtualImage field to given value.

### HasVirtualImage

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) HasVirtualImage() bool`

HasVirtualImage returns a boolean if a field has been set.

### GetCategory

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetConfig

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetContainerPorts

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) GetContainerPorts() []map[string]interface{}`

GetContainerPorts returns the ContainerPorts field if non-nil, zero value otherwise.

### GetContainerPortsOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) GetContainerPortsOk() (*[]map[string]interface{}, bool)`

GetContainerPortsOk returns a tuple with the ContainerPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerPorts

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) SetContainerPorts(v []map[string]interface{})`

SetContainerPorts sets ContainerPorts field to given value.

### HasContainerPorts

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) HasContainerPorts() bool`

HasContainerPorts returns a boolean if a field has been set.

### GetContainerScripts

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) GetContainerScripts() []map[string]interface{}`

GetContainerScripts returns the ContainerScripts field if non-nil, zero value otherwise.

### GetContainerScriptsOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) GetContainerScriptsOk() (*[]map[string]interface{}, bool)`

GetContainerScriptsOk returns a tuple with the ContainerScripts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerScripts

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) SetContainerScripts(v []map[string]interface{})`

SetContainerScripts sets ContainerScripts field to given value.

### HasContainerScripts

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) HasContainerScripts() bool`

HasContainerScripts returns a boolean if a field has been set.

### GetContainerTemplates

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) GetContainerTemplates() []map[string]interface{}`

GetContainerTemplates returns the ContainerTemplates field if non-nil, zero value otherwise.

### GetContainerTemplatesOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) GetContainerTemplatesOk() (*[]map[string]interface{}, bool)`

GetContainerTemplatesOk returns a tuple with the ContainerTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTemplates

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) SetContainerTemplates(v []map[string]interface{})`

SetContainerTemplates sets ContainerTemplates field to given value.

### HasContainerTemplates

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) HasContainerTemplates() bool`

HasContainerTemplates returns a boolean if a field has been set.

### GetEnvironmentVariables

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) GetEnvironmentVariables() []map[string]interface{}`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) GetEnvironmentVariablesOk() (*[]map[string]interface{}, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) SetEnvironmentVariables(v []map[string]interface{})`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


