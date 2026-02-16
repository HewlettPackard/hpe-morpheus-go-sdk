# InstanceTypeInstanceTypeLayoutsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**InstanceType** | Pointer to [**InstanceTypeInstanceTypeLayoutsInnerInstanceType**](InstanceTypeInstanceTypeLayoutsInnerInstanceType.md) |  | [optional] 
**Account** | Pointer to [**InstanceTypeInstanceTypeLayoutsInnerAccount**](InstanceTypeInstanceTypeLayoutsInnerAccount.md) |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**InstanceVersion** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Creatable** | Pointer to **bool** |  | [optional] 
**MemoryRequirement** | Pointer to **NullableInt64** |  | [optional] 
**SortOrder** | Pointer to **int64** |  | [optional] 
**SupportsConvertToManaged** | Pointer to **NullableBool** |  | [optional] 
**ProvisionType** | Pointer to [**InstanceTypeInstanceTypeLayoutsInnerProvisionType**](InstanceTypeInstanceTypeLayoutsInnerProvisionType.md) |  | [optional] 
**TaskSets** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ContainerTypes** | Pointer to [**[]InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner**](InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner.md) |  | [optional] 
**Mounts** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Ports** | Pointer to **[]map[string]interface{}** |  | [optional] 
**OptionTypes** | Pointer to **[]map[string]interface{}** |  | [optional] 
**EnvironmentVariables** | Pointer to **[]map[string]interface{}** |  | [optional] 
**PriceSets** | Pointer to **[]map[string]interface{}** |  | [optional] 
**SpecTemplates** | Pointer to **[]map[string]interface{}** |  | [optional] 
**TfvarSecret** | Pointer to **NullableString** |  | [optional] 
**Permissions** | Pointer to [**InstanceTypeInstanceTypeLayoutsInnerPermissions**](InstanceTypeInstanceTypeLayoutsInnerPermissions.md) |  | [optional] 

## Methods

### NewInstanceTypeInstanceTypeLayoutsInner

`func NewInstanceTypeInstanceTypeLayoutsInner() *InstanceTypeInstanceTypeLayoutsInner`

NewInstanceTypeInstanceTypeLayoutsInner instantiates a new InstanceTypeInstanceTypeLayoutsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceTypeInstanceTypeLayoutsInnerWithDefaults

`func NewInstanceTypeInstanceTypeLayoutsInnerWithDefaults() *InstanceTypeInstanceTypeLayoutsInner`

NewInstanceTypeInstanceTypeLayoutsInnerWithDefaults instantiates a new InstanceTypeInstanceTypeLayoutsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceTypeInstanceTypeLayoutsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceTypeInstanceTypeLayoutsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstanceType

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetInstanceType() InstanceTypeInstanceTypeLayoutsInnerInstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetInstanceTypeOk() (*InstanceTypeInstanceTypeLayoutsInnerInstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *InstanceTypeInstanceTypeLayoutsInner) SetInstanceType(v InstanceTypeInstanceTypeLayoutsInnerInstanceType)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *InstanceTypeInstanceTypeLayoutsInner) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetAccount

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetAccount() InstanceTypeInstanceTypeLayoutsInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetAccountOk() (*InstanceTypeInstanceTypeLayoutsInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *InstanceTypeInstanceTypeLayoutsInner) SetAccount(v InstanceTypeInstanceTypeLayoutsInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *InstanceTypeInstanceTypeLayoutsInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCode

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *InstanceTypeInstanceTypeLayoutsInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *InstanceTypeInstanceTypeLayoutsInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceTypeInstanceTypeLayoutsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceTypeInstanceTypeLayoutsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *InstanceTypeInstanceTypeLayoutsInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *InstanceTypeInstanceTypeLayoutsInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *InstanceTypeInstanceTypeLayoutsInner) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *InstanceTypeInstanceTypeLayoutsInner) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetInstanceVersion

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetInstanceVersion() string`

GetInstanceVersion returns the InstanceVersion field if non-nil, zero value otherwise.

### GetInstanceVersionOk

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetInstanceVersionOk() (*string, bool)`

GetInstanceVersionOk returns a tuple with the InstanceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceVersion

`func (o *InstanceTypeInstanceTypeLayoutsInner) SetInstanceVersion(v string)`

SetInstanceVersion sets InstanceVersion field to given value.

### HasInstanceVersion

`func (o *InstanceTypeInstanceTypeLayoutsInner) HasInstanceVersion() bool`

HasInstanceVersion returns a boolean if a field has been set.

### GetDescription

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InstanceTypeInstanceTypeLayoutsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InstanceTypeInstanceTypeLayoutsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *InstanceTypeInstanceTypeLayoutsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *InstanceTypeInstanceTypeLayoutsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCreatable

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *InstanceTypeInstanceTypeLayoutsInner) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *InstanceTypeInstanceTypeLayoutsInner) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetMemoryRequirement

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetMemoryRequirement() int64`

GetMemoryRequirement returns the MemoryRequirement field if non-nil, zero value otherwise.

### GetMemoryRequirementOk

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetMemoryRequirementOk() (*int64, bool)`

GetMemoryRequirementOk returns a tuple with the MemoryRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryRequirement

`func (o *InstanceTypeInstanceTypeLayoutsInner) SetMemoryRequirement(v int64)`

SetMemoryRequirement sets MemoryRequirement field to given value.

### HasMemoryRequirement

`func (o *InstanceTypeInstanceTypeLayoutsInner) HasMemoryRequirement() bool`

HasMemoryRequirement returns a boolean if a field has been set.

### SetMemoryRequirementNil

`func (o *InstanceTypeInstanceTypeLayoutsInner) SetMemoryRequirementNil(b bool)`

 SetMemoryRequirementNil sets the value for MemoryRequirement to be an explicit nil

### UnsetMemoryRequirement
`func (o *InstanceTypeInstanceTypeLayoutsInner) UnsetMemoryRequirement()`

UnsetMemoryRequirement ensures that no value is present for MemoryRequirement, not even an explicit nil
### GetSortOrder

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetSortOrder() int64`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetSortOrderOk() (*int64, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *InstanceTypeInstanceTypeLayoutsInner) SetSortOrder(v int64)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *InstanceTypeInstanceTypeLayoutsInner) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetSupportsConvertToManaged

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetSupportsConvertToManaged() bool`

GetSupportsConvertToManaged returns the SupportsConvertToManaged field if non-nil, zero value otherwise.

### GetSupportsConvertToManagedOk

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetSupportsConvertToManagedOk() (*bool, bool)`

GetSupportsConvertToManagedOk returns a tuple with the SupportsConvertToManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsConvertToManaged

`func (o *InstanceTypeInstanceTypeLayoutsInner) SetSupportsConvertToManaged(v bool)`

SetSupportsConvertToManaged sets SupportsConvertToManaged field to given value.

### HasSupportsConvertToManaged

`func (o *InstanceTypeInstanceTypeLayoutsInner) HasSupportsConvertToManaged() bool`

HasSupportsConvertToManaged returns a boolean if a field has been set.

### SetSupportsConvertToManagedNil

`func (o *InstanceTypeInstanceTypeLayoutsInner) SetSupportsConvertToManagedNil(b bool)`

 SetSupportsConvertToManagedNil sets the value for SupportsConvertToManaged to be an explicit nil

### UnsetSupportsConvertToManaged
`func (o *InstanceTypeInstanceTypeLayoutsInner) UnsetSupportsConvertToManaged()`

UnsetSupportsConvertToManaged ensures that no value is present for SupportsConvertToManaged, not even an explicit nil
### GetProvisionType

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetProvisionType() InstanceTypeInstanceTypeLayoutsInnerProvisionType`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetProvisionTypeOk() (*InstanceTypeInstanceTypeLayoutsInnerProvisionType, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *InstanceTypeInstanceTypeLayoutsInner) SetProvisionType(v InstanceTypeInstanceTypeLayoutsInnerProvisionType)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *InstanceTypeInstanceTypeLayoutsInner) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetTaskSets

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetTaskSets() []map[string]interface{}`

GetTaskSets returns the TaskSets field if non-nil, zero value otherwise.

### GetTaskSetsOk

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetTaskSetsOk() (*[]map[string]interface{}, bool)`

GetTaskSetsOk returns a tuple with the TaskSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSets

`func (o *InstanceTypeInstanceTypeLayoutsInner) SetTaskSets(v []map[string]interface{})`

SetTaskSets sets TaskSets field to given value.

### HasTaskSets

`func (o *InstanceTypeInstanceTypeLayoutsInner) HasTaskSets() bool`

HasTaskSets returns a boolean if a field has been set.

### SetTaskSetsNil

`func (o *InstanceTypeInstanceTypeLayoutsInner) SetTaskSetsNil(b bool)`

 SetTaskSetsNil sets the value for TaskSets to be an explicit nil

### UnsetTaskSets
`func (o *InstanceTypeInstanceTypeLayoutsInner) UnsetTaskSets()`

UnsetTaskSets ensures that no value is present for TaskSets, not even an explicit nil
### GetContainerTypes

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetContainerTypes() []InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner`

GetContainerTypes returns the ContainerTypes field if non-nil, zero value otherwise.

### GetContainerTypesOk

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetContainerTypesOk() (*[]InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner, bool)`

GetContainerTypesOk returns a tuple with the ContainerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTypes

`func (o *InstanceTypeInstanceTypeLayoutsInner) SetContainerTypes(v []InstanceTypeInstanceTypeLayoutsInnerContainerTypesInner)`

SetContainerTypes sets ContainerTypes field to given value.

### HasContainerTypes

`func (o *InstanceTypeInstanceTypeLayoutsInner) HasContainerTypes() bool`

HasContainerTypes returns a boolean if a field has been set.

### GetMounts

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetMounts() []map[string]interface{}`

GetMounts returns the Mounts field if non-nil, zero value otherwise.

### GetMountsOk

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetMountsOk() (*[]map[string]interface{}, bool)`

GetMountsOk returns a tuple with the Mounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMounts

`func (o *InstanceTypeInstanceTypeLayoutsInner) SetMounts(v []map[string]interface{})`

SetMounts sets Mounts field to given value.

### HasMounts

`func (o *InstanceTypeInstanceTypeLayoutsInner) HasMounts() bool`

HasMounts returns a boolean if a field has been set.

### SetMountsNil

`func (o *InstanceTypeInstanceTypeLayoutsInner) SetMountsNil(b bool)`

 SetMountsNil sets the value for Mounts to be an explicit nil

### UnsetMounts
`func (o *InstanceTypeInstanceTypeLayoutsInner) UnsetMounts()`

UnsetMounts ensures that no value is present for Mounts, not even an explicit nil
### GetPorts

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetPorts() []map[string]interface{}`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetPortsOk() (*[]map[string]interface{}, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *InstanceTypeInstanceTypeLayoutsInner) SetPorts(v []map[string]interface{})`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *InstanceTypeInstanceTypeLayoutsInner) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### SetPortsNil

`func (o *InstanceTypeInstanceTypeLayoutsInner) SetPortsNil(b bool)`

 SetPortsNil sets the value for Ports to be an explicit nil

### UnsetPorts
`func (o *InstanceTypeInstanceTypeLayoutsInner) UnsetPorts()`

UnsetPorts ensures that no value is present for Ports, not even an explicit nil
### GetOptionTypes

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetOptionTypes() []map[string]interface{}`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetOptionTypesOk() (*[]map[string]interface{}, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *InstanceTypeInstanceTypeLayoutsInner) SetOptionTypes(v []map[string]interface{})`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *InstanceTypeInstanceTypeLayoutsInner) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### SetOptionTypesNil

`func (o *InstanceTypeInstanceTypeLayoutsInner) SetOptionTypesNil(b bool)`

 SetOptionTypesNil sets the value for OptionTypes to be an explicit nil

### UnsetOptionTypes
`func (o *InstanceTypeInstanceTypeLayoutsInner) UnsetOptionTypes()`

UnsetOptionTypes ensures that no value is present for OptionTypes, not even an explicit nil
### GetEnvironmentVariables

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetEnvironmentVariables() []map[string]interface{}`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetEnvironmentVariablesOk() (*[]map[string]interface{}, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *InstanceTypeInstanceTypeLayoutsInner) SetEnvironmentVariables(v []map[string]interface{})`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *InstanceTypeInstanceTypeLayoutsInner) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### SetEnvironmentVariablesNil

`func (o *InstanceTypeInstanceTypeLayoutsInner) SetEnvironmentVariablesNil(b bool)`

 SetEnvironmentVariablesNil sets the value for EnvironmentVariables to be an explicit nil

### UnsetEnvironmentVariables
`func (o *InstanceTypeInstanceTypeLayoutsInner) UnsetEnvironmentVariables()`

UnsetEnvironmentVariables ensures that no value is present for EnvironmentVariables, not even an explicit nil
### GetPriceSets

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetPriceSets() []map[string]interface{}`

GetPriceSets returns the PriceSets field if non-nil, zero value otherwise.

### GetPriceSetsOk

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetPriceSetsOk() (*[]map[string]interface{}, bool)`

GetPriceSetsOk returns a tuple with the PriceSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSets

`func (o *InstanceTypeInstanceTypeLayoutsInner) SetPriceSets(v []map[string]interface{})`

SetPriceSets sets PriceSets field to given value.

### HasPriceSets

`func (o *InstanceTypeInstanceTypeLayoutsInner) HasPriceSets() bool`

HasPriceSets returns a boolean if a field has been set.

### SetPriceSetsNil

`func (o *InstanceTypeInstanceTypeLayoutsInner) SetPriceSetsNil(b bool)`

 SetPriceSetsNil sets the value for PriceSets to be an explicit nil

### UnsetPriceSets
`func (o *InstanceTypeInstanceTypeLayoutsInner) UnsetPriceSets()`

UnsetPriceSets ensures that no value is present for PriceSets, not even an explicit nil
### GetSpecTemplates

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetSpecTemplates() []map[string]interface{}`

GetSpecTemplates returns the SpecTemplates field if non-nil, zero value otherwise.

### GetSpecTemplatesOk

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetSpecTemplatesOk() (*[]map[string]interface{}, bool)`

GetSpecTemplatesOk returns a tuple with the SpecTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecTemplates

`func (o *InstanceTypeInstanceTypeLayoutsInner) SetSpecTemplates(v []map[string]interface{})`

SetSpecTemplates sets SpecTemplates field to given value.

### HasSpecTemplates

`func (o *InstanceTypeInstanceTypeLayoutsInner) HasSpecTemplates() bool`

HasSpecTemplates returns a boolean if a field has been set.

### SetSpecTemplatesNil

`func (o *InstanceTypeInstanceTypeLayoutsInner) SetSpecTemplatesNil(b bool)`

 SetSpecTemplatesNil sets the value for SpecTemplates to be an explicit nil

### UnsetSpecTemplates
`func (o *InstanceTypeInstanceTypeLayoutsInner) UnsetSpecTemplates()`

UnsetSpecTemplates ensures that no value is present for SpecTemplates, not even an explicit nil
### GetTfvarSecret

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetTfvarSecret() string`

GetTfvarSecret returns the TfvarSecret field if non-nil, zero value otherwise.

### GetTfvarSecretOk

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetTfvarSecretOk() (*string, bool)`

GetTfvarSecretOk returns a tuple with the TfvarSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTfvarSecret

`func (o *InstanceTypeInstanceTypeLayoutsInner) SetTfvarSecret(v string)`

SetTfvarSecret sets TfvarSecret field to given value.

### HasTfvarSecret

`func (o *InstanceTypeInstanceTypeLayoutsInner) HasTfvarSecret() bool`

HasTfvarSecret returns a boolean if a field has been set.

### SetTfvarSecretNil

`func (o *InstanceTypeInstanceTypeLayoutsInner) SetTfvarSecretNil(b bool)`

 SetTfvarSecretNil sets the value for TfvarSecret to be an explicit nil

### UnsetTfvarSecret
`func (o *InstanceTypeInstanceTypeLayoutsInner) UnsetTfvarSecret()`

UnsetTfvarSecret ensures that no value is present for TfvarSecret, not even an explicit nil
### GetPermissions

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetPermissions() InstanceTypeInstanceTypeLayoutsInnerPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *InstanceTypeInstanceTypeLayoutsInner) GetPermissionsOk() (*InstanceTypeInstanceTypeLayoutsInnerPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *InstanceTypeInstanceTypeLayoutsInner) SetPermissions(v InstanceTypeInstanceTypeLayoutsInnerPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *InstanceTypeInstanceTypeLayoutsInner) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


