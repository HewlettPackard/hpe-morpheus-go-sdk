# ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**InstanceType** | Pointer to [**ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerInstanceType**](ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerInstanceType.md) |  | [optional] 
**Account** | Pointer to [**ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerAccount**](ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerAccount.md) |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**InstanceVersion** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Creatable** | Pointer to **bool** |  | [optional] 
**MemoryRequirement** | Pointer to **NullableInt64** |  | [optional] 
**SortOrder** | Pointer to **int64** |  | [optional] 
**SupportsConvertToManaged** | Pointer to **NullableBool** |  | [optional] 
**ProvisionType** | Pointer to [**ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionType**](ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionType.md) |  | [optional] 
**TaskSets** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ContainerTypes** | Pointer to [**[]ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner**](ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner.md) |  | [optional] 
**Mounts** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Ports** | Pointer to **[]map[string]interface{}** |  | [optional] 
**OptionTypes** | Pointer to **[]map[string]interface{}** |  | [optional] 
**EnvironmentVariables** | Pointer to **[]map[string]interface{}** |  | [optional] 
**PriceSets** | Pointer to **[]map[string]interface{}** |  | [optional] 
**SpecTemplates** | Pointer to **[]map[string]interface{}** |  | [optional] 
**TfvarSecret** | Pointer to **NullableString** |  | [optional] 
**Permissions** | Pointer to [**ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerPermissions**](ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerPermissions.md) |  | [optional] 

## Methods

### NewListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner

`func NewListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner() *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner`

NewListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner instantiates a new ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerWithDefaults

`func NewListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerWithDefaults() *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner`

NewListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerWithDefaults instantiates a new ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstanceType

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetInstanceType() ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerInstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetInstanceTypeOk() (*ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerInstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) SetInstanceType(v ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerInstanceType)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetAccount

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetAccount() ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetAccountOk() (*ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) SetAccount(v ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCode

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetInstanceVersion

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetInstanceVersion() string`

GetInstanceVersion returns the InstanceVersion field if non-nil, zero value otherwise.

### GetInstanceVersionOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetInstanceVersionOk() (*string, bool)`

GetInstanceVersionOk returns a tuple with the InstanceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceVersion

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) SetInstanceVersion(v string)`

SetInstanceVersion sets InstanceVersion field to given value.

### HasInstanceVersion

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) HasInstanceVersion() bool`

HasInstanceVersion returns a boolean if a field has been set.

### GetDescription

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCreatable

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetMemoryRequirement

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetMemoryRequirement() int64`

GetMemoryRequirement returns the MemoryRequirement field if non-nil, zero value otherwise.

### GetMemoryRequirementOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetMemoryRequirementOk() (*int64, bool)`

GetMemoryRequirementOk returns a tuple with the MemoryRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryRequirement

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) SetMemoryRequirement(v int64)`

SetMemoryRequirement sets MemoryRequirement field to given value.

### HasMemoryRequirement

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) HasMemoryRequirement() bool`

HasMemoryRequirement returns a boolean if a field has been set.

### SetMemoryRequirementNil

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) SetMemoryRequirementNil(b bool)`

 SetMemoryRequirementNil sets the value for MemoryRequirement to be an explicit nil

### UnsetMemoryRequirement
`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) UnsetMemoryRequirement()`

UnsetMemoryRequirement ensures that no value is present for MemoryRequirement, not even an explicit nil
### GetSortOrder

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetSortOrder() int64`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetSortOrderOk() (*int64, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) SetSortOrder(v int64)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetSupportsConvertToManaged

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetSupportsConvertToManaged() bool`

GetSupportsConvertToManaged returns the SupportsConvertToManaged field if non-nil, zero value otherwise.

### GetSupportsConvertToManagedOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetSupportsConvertToManagedOk() (*bool, bool)`

GetSupportsConvertToManagedOk returns a tuple with the SupportsConvertToManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsConvertToManaged

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) SetSupportsConvertToManaged(v bool)`

SetSupportsConvertToManaged sets SupportsConvertToManaged field to given value.

### HasSupportsConvertToManaged

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) HasSupportsConvertToManaged() bool`

HasSupportsConvertToManaged returns a boolean if a field has been set.

### SetSupportsConvertToManagedNil

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) SetSupportsConvertToManagedNil(b bool)`

 SetSupportsConvertToManagedNil sets the value for SupportsConvertToManaged to be an explicit nil

### UnsetSupportsConvertToManaged
`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) UnsetSupportsConvertToManaged()`

UnsetSupportsConvertToManaged ensures that no value is present for SupportsConvertToManaged, not even an explicit nil
### GetProvisionType

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetProvisionType() ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionType`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetProvisionTypeOk() (*ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionType, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) SetProvisionType(v ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerProvisionType)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetTaskSets

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetTaskSets() []map[string]interface{}`

GetTaskSets returns the TaskSets field if non-nil, zero value otherwise.

### GetTaskSetsOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetTaskSetsOk() (*[]map[string]interface{}, bool)`

GetTaskSetsOk returns a tuple with the TaskSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSets

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) SetTaskSets(v []map[string]interface{})`

SetTaskSets sets TaskSets field to given value.

### HasTaskSets

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) HasTaskSets() bool`

HasTaskSets returns a boolean if a field has been set.

### SetTaskSetsNil

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) SetTaskSetsNil(b bool)`

 SetTaskSetsNil sets the value for TaskSets to be an explicit nil

### UnsetTaskSets
`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) UnsetTaskSets()`

UnsetTaskSets ensures that no value is present for TaskSets, not even an explicit nil
### GetContainerTypes

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetContainerTypes() []ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner`

GetContainerTypes returns the ContainerTypes field if non-nil, zero value otherwise.

### GetContainerTypesOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetContainerTypesOk() (*[]ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner, bool)`

GetContainerTypesOk returns a tuple with the ContainerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTypes

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) SetContainerTypes(v []ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerContainerTypesInner)`

SetContainerTypes sets ContainerTypes field to given value.

### HasContainerTypes

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) HasContainerTypes() bool`

HasContainerTypes returns a boolean if a field has been set.

### GetMounts

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetMounts() []map[string]interface{}`

GetMounts returns the Mounts field if non-nil, zero value otherwise.

### GetMountsOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetMountsOk() (*[]map[string]interface{}, bool)`

GetMountsOk returns a tuple with the Mounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMounts

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) SetMounts(v []map[string]interface{})`

SetMounts sets Mounts field to given value.

### HasMounts

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) HasMounts() bool`

HasMounts returns a boolean if a field has been set.

### SetMountsNil

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) SetMountsNil(b bool)`

 SetMountsNil sets the value for Mounts to be an explicit nil

### UnsetMounts
`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) UnsetMounts()`

UnsetMounts ensures that no value is present for Mounts, not even an explicit nil
### GetPorts

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetPorts() []map[string]interface{}`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetPortsOk() (*[]map[string]interface{}, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) SetPorts(v []map[string]interface{})`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### SetPortsNil

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) SetPortsNil(b bool)`

 SetPortsNil sets the value for Ports to be an explicit nil

### UnsetPorts
`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) UnsetPorts()`

UnsetPorts ensures that no value is present for Ports, not even an explicit nil
### GetOptionTypes

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetOptionTypes() []map[string]interface{}`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetOptionTypesOk() (*[]map[string]interface{}, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) SetOptionTypes(v []map[string]interface{})`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### SetOptionTypesNil

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) SetOptionTypesNil(b bool)`

 SetOptionTypesNil sets the value for OptionTypes to be an explicit nil

### UnsetOptionTypes
`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) UnsetOptionTypes()`

UnsetOptionTypes ensures that no value is present for OptionTypes, not even an explicit nil
### GetEnvironmentVariables

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetEnvironmentVariables() []map[string]interface{}`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetEnvironmentVariablesOk() (*[]map[string]interface{}, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) SetEnvironmentVariables(v []map[string]interface{})`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### SetEnvironmentVariablesNil

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) SetEnvironmentVariablesNil(b bool)`

 SetEnvironmentVariablesNil sets the value for EnvironmentVariables to be an explicit nil

### UnsetEnvironmentVariables
`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) UnsetEnvironmentVariables()`

UnsetEnvironmentVariables ensures that no value is present for EnvironmentVariables, not even an explicit nil
### GetPriceSets

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetPriceSets() []map[string]interface{}`

GetPriceSets returns the PriceSets field if non-nil, zero value otherwise.

### GetPriceSetsOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetPriceSetsOk() (*[]map[string]interface{}, bool)`

GetPriceSetsOk returns a tuple with the PriceSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSets

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) SetPriceSets(v []map[string]interface{})`

SetPriceSets sets PriceSets field to given value.

### HasPriceSets

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) HasPriceSets() bool`

HasPriceSets returns a boolean if a field has been set.

### SetPriceSetsNil

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) SetPriceSetsNil(b bool)`

 SetPriceSetsNil sets the value for PriceSets to be an explicit nil

### UnsetPriceSets
`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) UnsetPriceSets()`

UnsetPriceSets ensures that no value is present for PriceSets, not even an explicit nil
### GetSpecTemplates

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetSpecTemplates() []map[string]interface{}`

GetSpecTemplates returns the SpecTemplates field if non-nil, zero value otherwise.

### GetSpecTemplatesOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetSpecTemplatesOk() (*[]map[string]interface{}, bool)`

GetSpecTemplatesOk returns a tuple with the SpecTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecTemplates

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) SetSpecTemplates(v []map[string]interface{})`

SetSpecTemplates sets SpecTemplates field to given value.

### HasSpecTemplates

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) HasSpecTemplates() bool`

HasSpecTemplates returns a boolean if a field has been set.

### SetSpecTemplatesNil

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) SetSpecTemplatesNil(b bool)`

 SetSpecTemplatesNil sets the value for SpecTemplates to be an explicit nil

### UnsetSpecTemplates
`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) UnsetSpecTemplates()`

UnsetSpecTemplates ensures that no value is present for SpecTemplates, not even an explicit nil
### GetTfvarSecret

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetTfvarSecret() string`

GetTfvarSecret returns the TfvarSecret field if non-nil, zero value otherwise.

### GetTfvarSecretOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetTfvarSecretOk() (*string, bool)`

GetTfvarSecretOk returns a tuple with the TfvarSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTfvarSecret

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) SetTfvarSecret(v string)`

SetTfvarSecret sets TfvarSecret field to given value.

### HasTfvarSecret

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) HasTfvarSecret() bool`

HasTfvarSecret returns a boolean if a field has been set.

### SetTfvarSecretNil

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) SetTfvarSecretNil(b bool)`

 SetTfvarSecretNil sets the value for TfvarSecret to be an explicit nil

### UnsetTfvarSecret
`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) UnsetTfvarSecret()`

UnsetTfvarSecret ensures that no value is present for TfvarSecret, not even an explicit nil
### GetPermissions

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetPermissions() ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) GetPermissionsOk() (*ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) SetPermissions(v ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInnerPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ListLayoutsForInstanceType200ResponseAllOfInstanceTypeLayoutsInner) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


