# ListLayouts200ResponseAllOfInstanceTypeLayoutsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**InstanceType** | Pointer to [**ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerInstanceType**](ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerInstanceType.md) |  | [optional] 
**Account** | Pointer to [**ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerAccount**](ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerAccount.md) |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**InstanceVersion** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Creatable** | Pointer to **bool** |  | [optional] 
**MemoryRequirement** | Pointer to **NullableInt64** |  | [optional] 
**SortOrder** | Pointer to **int64** |  | [optional] 
**SupportsConvertToManaged** | Pointer to **NullableBool** |  | [optional] 
**ProvisionType** | Pointer to [**ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType**](ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType.md) |  | [optional] 
**TaskSets** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ContainerTypes** | Pointer to [**[]ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner**](ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner.md) |  | [optional] 
**Mounts** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Ports** | Pointer to **[]map[string]interface{}** |  | [optional] 
**OptionTypes** | Pointer to **[]map[string]interface{}** |  | [optional] 
**EnvironmentVariables** | Pointer to **[]map[string]interface{}** |  | [optional] 
**PriceSets** | Pointer to **[]map[string]interface{}** |  | [optional] 
**SpecTemplates** | Pointer to **[]map[string]interface{}** |  | [optional] 
**TfvarSecret** | Pointer to **NullableString** |  | [optional] 
**Permissions** | Pointer to [**ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerPermissions**](ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerPermissions.md) |  | [optional] 

## Methods

### NewListLayouts200ResponseAllOfInstanceTypeLayoutsInner

`func NewListLayouts200ResponseAllOfInstanceTypeLayoutsInner() *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner`

NewListLayouts200ResponseAllOfInstanceTypeLayoutsInner instantiates a new ListLayouts200ResponseAllOfInstanceTypeLayoutsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLayouts200ResponseAllOfInstanceTypeLayoutsInnerWithDefaults

`func NewListLayouts200ResponseAllOfInstanceTypeLayoutsInnerWithDefaults() *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner`

NewListLayouts200ResponseAllOfInstanceTypeLayoutsInnerWithDefaults instantiates a new ListLayouts200ResponseAllOfInstanceTypeLayoutsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstanceType

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetInstanceType() ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerInstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetInstanceTypeOk() (*ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerInstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) SetInstanceType(v ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerInstanceType)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetAccount

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetAccount() ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetAccountOk() (*ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) SetAccount(v ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCode

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetInstanceVersion

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetInstanceVersion() string`

GetInstanceVersion returns the InstanceVersion field if non-nil, zero value otherwise.

### GetInstanceVersionOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetInstanceVersionOk() (*string, bool)`

GetInstanceVersionOk returns a tuple with the InstanceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceVersion

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) SetInstanceVersion(v string)`

SetInstanceVersion sets InstanceVersion field to given value.

### HasInstanceVersion

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) HasInstanceVersion() bool`

HasInstanceVersion returns a boolean if a field has been set.

### GetDescription

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCreatable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetMemoryRequirement

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetMemoryRequirement() int64`

GetMemoryRequirement returns the MemoryRequirement field if non-nil, zero value otherwise.

### GetMemoryRequirementOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetMemoryRequirementOk() (*int64, bool)`

GetMemoryRequirementOk returns a tuple with the MemoryRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryRequirement

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) SetMemoryRequirement(v int64)`

SetMemoryRequirement sets MemoryRequirement field to given value.

### HasMemoryRequirement

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) HasMemoryRequirement() bool`

HasMemoryRequirement returns a boolean if a field has been set.

### SetMemoryRequirementNil

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) SetMemoryRequirementNil(b bool)`

 SetMemoryRequirementNil sets the value for MemoryRequirement to be an explicit nil

### UnsetMemoryRequirement
`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) UnsetMemoryRequirement()`

UnsetMemoryRequirement ensures that no value is present for MemoryRequirement, not even an explicit nil
### GetSortOrder

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetSortOrder() int64`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetSortOrderOk() (*int64, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) SetSortOrder(v int64)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetSupportsConvertToManaged

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetSupportsConvertToManaged() bool`

GetSupportsConvertToManaged returns the SupportsConvertToManaged field if non-nil, zero value otherwise.

### GetSupportsConvertToManagedOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetSupportsConvertToManagedOk() (*bool, bool)`

GetSupportsConvertToManagedOk returns a tuple with the SupportsConvertToManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsConvertToManaged

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) SetSupportsConvertToManaged(v bool)`

SetSupportsConvertToManaged sets SupportsConvertToManaged field to given value.

### HasSupportsConvertToManaged

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) HasSupportsConvertToManaged() bool`

HasSupportsConvertToManaged returns a boolean if a field has been set.

### SetSupportsConvertToManagedNil

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) SetSupportsConvertToManagedNil(b bool)`

 SetSupportsConvertToManagedNil sets the value for SupportsConvertToManaged to be an explicit nil

### UnsetSupportsConvertToManaged
`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) UnsetSupportsConvertToManaged()`

UnsetSupportsConvertToManaged ensures that no value is present for SupportsConvertToManaged, not even an explicit nil
### GetProvisionType

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetProvisionType() ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetProvisionTypeOk() (*ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) SetProvisionType(v ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerProvisionType)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetTaskSets

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetTaskSets() []map[string]interface{}`

GetTaskSets returns the TaskSets field if non-nil, zero value otherwise.

### GetTaskSetsOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetTaskSetsOk() (*[]map[string]interface{}, bool)`

GetTaskSetsOk returns a tuple with the TaskSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSets

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) SetTaskSets(v []map[string]interface{})`

SetTaskSets sets TaskSets field to given value.

### HasTaskSets

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) HasTaskSets() bool`

HasTaskSets returns a boolean if a field has been set.

### SetTaskSetsNil

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) SetTaskSetsNil(b bool)`

 SetTaskSetsNil sets the value for TaskSets to be an explicit nil

### UnsetTaskSets
`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) UnsetTaskSets()`

UnsetTaskSets ensures that no value is present for TaskSets, not even an explicit nil
### GetContainerTypes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetContainerTypes() []ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner`

GetContainerTypes returns the ContainerTypes field if non-nil, zero value otherwise.

### GetContainerTypesOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetContainerTypesOk() (*[]ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner, bool)`

GetContainerTypesOk returns a tuple with the ContainerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTypes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) SetContainerTypes(v []ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner)`

SetContainerTypes sets ContainerTypes field to given value.

### HasContainerTypes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) HasContainerTypes() bool`

HasContainerTypes returns a boolean if a field has been set.

### GetMounts

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetMounts() []map[string]interface{}`

GetMounts returns the Mounts field if non-nil, zero value otherwise.

### GetMountsOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetMountsOk() (*[]map[string]interface{}, bool)`

GetMountsOk returns a tuple with the Mounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMounts

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) SetMounts(v []map[string]interface{})`

SetMounts sets Mounts field to given value.

### HasMounts

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) HasMounts() bool`

HasMounts returns a boolean if a field has been set.

### SetMountsNil

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) SetMountsNil(b bool)`

 SetMountsNil sets the value for Mounts to be an explicit nil

### UnsetMounts
`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) UnsetMounts()`

UnsetMounts ensures that no value is present for Mounts, not even an explicit nil
### GetPorts

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetPorts() []map[string]interface{}`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetPortsOk() (*[]map[string]interface{}, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) SetPorts(v []map[string]interface{})`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### SetPortsNil

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) SetPortsNil(b bool)`

 SetPortsNil sets the value for Ports to be an explicit nil

### UnsetPorts
`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) UnsetPorts()`

UnsetPorts ensures that no value is present for Ports, not even an explicit nil
### GetOptionTypes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetOptionTypes() []map[string]interface{}`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetOptionTypesOk() (*[]map[string]interface{}, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) SetOptionTypes(v []map[string]interface{})`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### SetOptionTypesNil

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) SetOptionTypesNil(b bool)`

 SetOptionTypesNil sets the value for OptionTypes to be an explicit nil

### UnsetOptionTypes
`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) UnsetOptionTypes()`

UnsetOptionTypes ensures that no value is present for OptionTypes, not even an explicit nil
### GetEnvironmentVariables

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetEnvironmentVariables() []map[string]interface{}`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetEnvironmentVariablesOk() (*[]map[string]interface{}, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) SetEnvironmentVariables(v []map[string]interface{})`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### SetEnvironmentVariablesNil

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) SetEnvironmentVariablesNil(b bool)`

 SetEnvironmentVariablesNil sets the value for EnvironmentVariables to be an explicit nil

### UnsetEnvironmentVariables
`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) UnsetEnvironmentVariables()`

UnsetEnvironmentVariables ensures that no value is present for EnvironmentVariables, not even an explicit nil
### GetPriceSets

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetPriceSets() []map[string]interface{}`

GetPriceSets returns the PriceSets field if non-nil, zero value otherwise.

### GetPriceSetsOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetPriceSetsOk() (*[]map[string]interface{}, bool)`

GetPriceSetsOk returns a tuple with the PriceSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSets

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) SetPriceSets(v []map[string]interface{})`

SetPriceSets sets PriceSets field to given value.

### HasPriceSets

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) HasPriceSets() bool`

HasPriceSets returns a boolean if a field has been set.

### SetPriceSetsNil

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) SetPriceSetsNil(b bool)`

 SetPriceSetsNil sets the value for PriceSets to be an explicit nil

### UnsetPriceSets
`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) UnsetPriceSets()`

UnsetPriceSets ensures that no value is present for PriceSets, not even an explicit nil
### GetSpecTemplates

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetSpecTemplates() []map[string]interface{}`

GetSpecTemplates returns the SpecTemplates field if non-nil, zero value otherwise.

### GetSpecTemplatesOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetSpecTemplatesOk() (*[]map[string]interface{}, bool)`

GetSpecTemplatesOk returns a tuple with the SpecTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecTemplates

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) SetSpecTemplates(v []map[string]interface{})`

SetSpecTemplates sets SpecTemplates field to given value.

### HasSpecTemplates

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) HasSpecTemplates() bool`

HasSpecTemplates returns a boolean if a field has been set.

### SetSpecTemplatesNil

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) SetSpecTemplatesNil(b bool)`

 SetSpecTemplatesNil sets the value for SpecTemplates to be an explicit nil

### UnsetSpecTemplates
`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) UnsetSpecTemplates()`

UnsetSpecTemplates ensures that no value is present for SpecTemplates, not even an explicit nil
### GetTfvarSecret

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetTfvarSecret() string`

GetTfvarSecret returns the TfvarSecret field if non-nil, zero value otherwise.

### GetTfvarSecretOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetTfvarSecretOk() (*string, bool)`

GetTfvarSecretOk returns a tuple with the TfvarSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTfvarSecret

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) SetTfvarSecret(v string)`

SetTfvarSecret sets TfvarSecret field to given value.

### HasTfvarSecret

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) HasTfvarSecret() bool`

HasTfvarSecret returns a boolean if a field has been set.

### SetTfvarSecretNil

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) SetTfvarSecretNil(b bool)`

 SetTfvarSecretNil sets the value for TfvarSecret to be an explicit nil

### UnsetTfvarSecret
`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) UnsetTfvarSecret()`

UnsetTfvarSecret ensures that no value is present for TfvarSecret, not even an explicit nil
### GetPermissions

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetPermissions() ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) GetPermissionsOk() (*ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) SetPermissions(v ListLayouts200ResponseAllOfInstanceTypeLayoutsInnerPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ListLayouts200ResponseAllOfInstanceTypeLayoutsInner) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


