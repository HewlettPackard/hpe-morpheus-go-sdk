# AddLayout200ResponseInstanceTypeLayout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**InstanceType** | Pointer to [**AddLayout200ResponseInstanceTypeLayoutInstanceType**](AddLayout200ResponseInstanceTypeLayoutInstanceType.md) |  | [optional] 
**Account** | Pointer to [**AddLayout200ResponseInstanceTypeLayoutAccount**](AddLayout200ResponseInstanceTypeLayoutAccount.md) |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**InstanceVersion** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Creatable** | Pointer to **bool** |  | [optional] 
**MemoryRequirement** | Pointer to **NullableInt64** |  | [optional] 
**SortOrder** | Pointer to **int64** |  | [optional] 
**SupportsConvertToManaged** | Pointer to **NullableBool** |  | [optional] 
**ProvisionType** | Pointer to [**AddLayout200ResponseInstanceTypeLayoutProvisionType**](AddLayout200ResponseInstanceTypeLayoutProvisionType.md) |  | [optional] 
**TaskSets** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ContainerTypes** | Pointer to [**[]AddLayout200ResponseInstanceTypeLayoutContainerTypesInner**](AddLayout200ResponseInstanceTypeLayoutContainerTypesInner.md) |  | [optional] 
**Mounts** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Ports** | Pointer to **[]map[string]interface{}** |  | [optional] 
**OptionTypes** | Pointer to **[]map[string]interface{}** |  | [optional] 
**EnvironmentVariables** | Pointer to **[]map[string]interface{}** |  | [optional] 
**PriceSets** | Pointer to **[]map[string]interface{}** |  | [optional] 
**SpecTemplates** | Pointer to **[]map[string]interface{}** |  | [optional] 
**TfvarSecret** | Pointer to **NullableString** |  | [optional] 
**Permissions** | Pointer to [**AddLayout200ResponseInstanceTypeLayoutPermissions**](AddLayout200ResponseInstanceTypeLayoutPermissions.md) |  | [optional] 

## Methods

### NewAddLayout200ResponseInstanceTypeLayout

`func NewAddLayout200ResponseInstanceTypeLayout() *AddLayout200ResponseInstanceTypeLayout`

NewAddLayout200ResponseInstanceTypeLayout instantiates a new AddLayout200ResponseInstanceTypeLayout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddLayout200ResponseInstanceTypeLayoutWithDefaults

`func NewAddLayout200ResponseInstanceTypeLayoutWithDefaults() *AddLayout200ResponseInstanceTypeLayout`

NewAddLayout200ResponseInstanceTypeLayoutWithDefaults instantiates a new AddLayout200ResponseInstanceTypeLayout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddLayout200ResponseInstanceTypeLayout) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddLayout200ResponseInstanceTypeLayout) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddLayout200ResponseInstanceTypeLayout) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddLayout200ResponseInstanceTypeLayout) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstanceType

`func (o *AddLayout200ResponseInstanceTypeLayout) GetInstanceType() AddLayout200ResponseInstanceTypeLayoutInstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *AddLayout200ResponseInstanceTypeLayout) GetInstanceTypeOk() (*AddLayout200ResponseInstanceTypeLayoutInstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *AddLayout200ResponseInstanceTypeLayout) SetInstanceType(v AddLayout200ResponseInstanceTypeLayoutInstanceType)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *AddLayout200ResponseInstanceTypeLayout) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetAccount

`func (o *AddLayout200ResponseInstanceTypeLayout) GetAccount() AddLayout200ResponseInstanceTypeLayoutAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AddLayout200ResponseInstanceTypeLayout) GetAccountOk() (*AddLayout200ResponseInstanceTypeLayoutAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AddLayout200ResponseInstanceTypeLayout) SetAccount(v AddLayout200ResponseInstanceTypeLayoutAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AddLayout200ResponseInstanceTypeLayout) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCode

`func (o *AddLayout200ResponseInstanceTypeLayout) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AddLayout200ResponseInstanceTypeLayout) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AddLayout200ResponseInstanceTypeLayout) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AddLayout200ResponseInstanceTypeLayout) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *AddLayout200ResponseInstanceTypeLayout) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddLayout200ResponseInstanceTypeLayout) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddLayout200ResponseInstanceTypeLayout) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddLayout200ResponseInstanceTypeLayout) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *AddLayout200ResponseInstanceTypeLayout) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddLayout200ResponseInstanceTypeLayout) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddLayout200ResponseInstanceTypeLayout) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddLayout200ResponseInstanceTypeLayout) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *AddLayout200ResponseInstanceTypeLayout) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *AddLayout200ResponseInstanceTypeLayout) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetInstanceVersion

`func (o *AddLayout200ResponseInstanceTypeLayout) GetInstanceVersion() string`

GetInstanceVersion returns the InstanceVersion field if non-nil, zero value otherwise.

### GetInstanceVersionOk

`func (o *AddLayout200ResponseInstanceTypeLayout) GetInstanceVersionOk() (*string, bool)`

GetInstanceVersionOk returns a tuple with the InstanceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceVersion

`func (o *AddLayout200ResponseInstanceTypeLayout) SetInstanceVersion(v string)`

SetInstanceVersion sets InstanceVersion field to given value.

### HasInstanceVersion

`func (o *AddLayout200ResponseInstanceTypeLayout) HasInstanceVersion() bool`

HasInstanceVersion returns a boolean if a field has been set.

### GetDescription

`func (o *AddLayout200ResponseInstanceTypeLayout) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddLayout200ResponseInstanceTypeLayout) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddLayout200ResponseInstanceTypeLayout) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddLayout200ResponseInstanceTypeLayout) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AddLayout200ResponseInstanceTypeLayout) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AddLayout200ResponseInstanceTypeLayout) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCreatable

`func (o *AddLayout200ResponseInstanceTypeLayout) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *AddLayout200ResponseInstanceTypeLayout) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *AddLayout200ResponseInstanceTypeLayout) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *AddLayout200ResponseInstanceTypeLayout) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetMemoryRequirement

`func (o *AddLayout200ResponseInstanceTypeLayout) GetMemoryRequirement() int64`

GetMemoryRequirement returns the MemoryRequirement field if non-nil, zero value otherwise.

### GetMemoryRequirementOk

`func (o *AddLayout200ResponseInstanceTypeLayout) GetMemoryRequirementOk() (*int64, bool)`

GetMemoryRequirementOk returns a tuple with the MemoryRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryRequirement

`func (o *AddLayout200ResponseInstanceTypeLayout) SetMemoryRequirement(v int64)`

SetMemoryRequirement sets MemoryRequirement field to given value.

### HasMemoryRequirement

`func (o *AddLayout200ResponseInstanceTypeLayout) HasMemoryRequirement() bool`

HasMemoryRequirement returns a boolean if a field has been set.

### SetMemoryRequirementNil

`func (o *AddLayout200ResponseInstanceTypeLayout) SetMemoryRequirementNil(b bool)`

 SetMemoryRequirementNil sets the value for MemoryRequirement to be an explicit nil

### UnsetMemoryRequirement
`func (o *AddLayout200ResponseInstanceTypeLayout) UnsetMemoryRequirement()`

UnsetMemoryRequirement ensures that no value is present for MemoryRequirement, not even an explicit nil
### GetSortOrder

`func (o *AddLayout200ResponseInstanceTypeLayout) GetSortOrder() int64`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *AddLayout200ResponseInstanceTypeLayout) GetSortOrderOk() (*int64, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *AddLayout200ResponseInstanceTypeLayout) SetSortOrder(v int64)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *AddLayout200ResponseInstanceTypeLayout) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetSupportsConvertToManaged

`func (o *AddLayout200ResponseInstanceTypeLayout) GetSupportsConvertToManaged() bool`

GetSupportsConvertToManaged returns the SupportsConvertToManaged field if non-nil, zero value otherwise.

### GetSupportsConvertToManagedOk

`func (o *AddLayout200ResponseInstanceTypeLayout) GetSupportsConvertToManagedOk() (*bool, bool)`

GetSupportsConvertToManagedOk returns a tuple with the SupportsConvertToManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsConvertToManaged

`func (o *AddLayout200ResponseInstanceTypeLayout) SetSupportsConvertToManaged(v bool)`

SetSupportsConvertToManaged sets SupportsConvertToManaged field to given value.

### HasSupportsConvertToManaged

`func (o *AddLayout200ResponseInstanceTypeLayout) HasSupportsConvertToManaged() bool`

HasSupportsConvertToManaged returns a boolean if a field has been set.

### SetSupportsConvertToManagedNil

`func (o *AddLayout200ResponseInstanceTypeLayout) SetSupportsConvertToManagedNil(b bool)`

 SetSupportsConvertToManagedNil sets the value for SupportsConvertToManaged to be an explicit nil

### UnsetSupportsConvertToManaged
`func (o *AddLayout200ResponseInstanceTypeLayout) UnsetSupportsConvertToManaged()`

UnsetSupportsConvertToManaged ensures that no value is present for SupportsConvertToManaged, not even an explicit nil
### GetProvisionType

`func (o *AddLayout200ResponseInstanceTypeLayout) GetProvisionType() AddLayout200ResponseInstanceTypeLayoutProvisionType`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *AddLayout200ResponseInstanceTypeLayout) GetProvisionTypeOk() (*AddLayout200ResponseInstanceTypeLayoutProvisionType, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *AddLayout200ResponseInstanceTypeLayout) SetProvisionType(v AddLayout200ResponseInstanceTypeLayoutProvisionType)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *AddLayout200ResponseInstanceTypeLayout) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetTaskSets

`func (o *AddLayout200ResponseInstanceTypeLayout) GetTaskSets() []map[string]interface{}`

GetTaskSets returns the TaskSets field if non-nil, zero value otherwise.

### GetTaskSetsOk

`func (o *AddLayout200ResponseInstanceTypeLayout) GetTaskSetsOk() (*[]map[string]interface{}, bool)`

GetTaskSetsOk returns a tuple with the TaskSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSets

`func (o *AddLayout200ResponseInstanceTypeLayout) SetTaskSets(v []map[string]interface{})`

SetTaskSets sets TaskSets field to given value.

### HasTaskSets

`func (o *AddLayout200ResponseInstanceTypeLayout) HasTaskSets() bool`

HasTaskSets returns a boolean if a field has been set.

### SetTaskSetsNil

`func (o *AddLayout200ResponseInstanceTypeLayout) SetTaskSetsNil(b bool)`

 SetTaskSetsNil sets the value for TaskSets to be an explicit nil

### UnsetTaskSets
`func (o *AddLayout200ResponseInstanceTypeLayout) UnsetTaskSets()`

UnsetTaskSets ensures that no value is present for TaskSets, not even an explicit nil
### GetContainerTypes

`func (o *AddLayout200ResponseInstanceTypeLayout) GetContainerTypes() []AddLayout200ResponseInstanceTypeLayoutContainerTypesInner`

GetContainerTypes returns the ContainerTypes field if non-nil, zero value otherwise.

### GetContainerTypesOk

`func (o *AddLayout200ResponseInstanceTypeLayout) GetContainerTypesOk() (*[]AddLayout200ResponseInstanceTypeLayoutContainerTypesInner, bool)`

GetContainerTypesOk returns a tuple with the ContainerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTypes

`func (o *AddLayout200ResponseInstanceTypeLayout) SetContainerTypes(v []AddLayout200ResponseInstanceTypeLayoutContainerTypesInner)`

SetContainerTypes sets ContainerTypes field to given value.

### HasContainerTypes

`func (o *AddLayout200ResponseInstanceTypeLayout) HasContainerTypes() bool`

HasContainerTypes returns a boolean if a field has been set.

### GetMounts

`func (o *AddLayout200ResponseInstanceTypeLayout) GetMounts() []map[string]interface{}`

GetMounts returns the Mounts field if non-nil, zero value otherwise.

### GetMountsOk

`func (o *AddLayout200ResponseInstanceTypeLayout) GetMountsOk() (*[]map[string]interface{}, bool)`

GetMountsOk returns a tuple with the Mounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMounts

`func (o *AddLayout200ResponseInstanceTypeLayout) SetMounts(v []map[string]interface{})`

SetMounts sets Mounts field to given value.

### HasMounts

`func (o *AddLayout200ResponseInstanceTypeLayout) HasMounts() bool`

HasMounts returns a boolean if a field has been set.

### SetMountsNil

`func (o *AddLayout200ResponseInstanceTypeLayout) SetMountsNil(b bool)`

 SetMountsNil sets the value for Mounts to be an explicit nil

### UnsetMounts
`func (o *AddLayout200ResponseInstanceTypeLayout) UnsetMounts()`

UnsetMounts ensures that no value is present for Mounts, not even an explicit nil
### GetPorts

`func (o *AddLayout200ResponseInstanceTypeLayout) GetPorts() []map[string]interface{}`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *AddLayout200ResponseInstanceTypeLayout) GetPortsOk() (*[]map[string]interface{}, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *AddLayout200ResponseInstanceTypeLayout) SetPorts(v []map[string]interface{})`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *AddLayout200ResponseInstanceTypeLayout) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### SetPortsNil

`func (o *AddLayout200ResponseInstanceTypeLayout) SetPortsNil(b bool)`

 SetPortsNil sets the value for Ports to be an explicit nil

### UnsetPorts
`func (o *AddLayout200ResponseInstanceTypeLayout) UnsetPorts()`

UnsetPorts ensures that no value is present for Ports, not even an explicit nil
### GetOptionTypes

`func (o *AddLayout200ResponseInstanceTypeLayout) GetOptionTypes() []map[string]interface{}`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *AddLayout200ResponseInstanceTypeLayout) GetOptionTypesOk() (*[]map[string]interface{}, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *AddLayout200ResponseInstanceTypeLayout) SetOptionTypes(v []map[string]interface{})`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *AddLayout200ResponseInstanceTypeLayout) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### SetOptionTypesNil

`func (o *AddLayout200ResponseInstanceTypeLayout) SetOptionTypesNil(b bool)`

 SetOptionTypesNil sets the value for OptionTypes to be an explicit nil

### UnsetOptionTypes
`func (o *AddLayout200ResponseInstanceTypeLayout) UnsetOptionTypes()`

UnsetOptionTypes ensures that no value is present for OptionTypes, not even an explicit nil
### GetEnvironmentVariables

`func (o *AddLayout200ResponseInstanceTypeLayout) GetEnvironmentVariables() []map[string]interface{}`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *AddLayout200ResponseInstanceTypeLayout) GetEnvironmentVariablesOk() (*[]map[string]interface{}, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *AddLayout200ResponseInstanceTypeLayout) SetEnvironmentVariables(v []map[string]interface{})`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *AddLayout200ResponseInstanceTypeLayout) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### SetEnvironmentVariablesNil

`func (o *AddLayout200ResponseInstanceTypeLayout) SetEnvironmentVariablesNil(b bool)`

 SetEnvironmentVariablesNil sets the value for EnvironmentVariables to be an explicit nil

### UnsetEnvironmentVariables
`func (o *AddLayout200ResponseInstanceTypeLayout) UnsetEnvironmentVariables()`

UnsetEnvironmentVariables ensures that no value is present for EnvironmentVariables, not even an explicit nil
### GetPriceSets

`func (o *AddLayout200ResponseInstanceTypeLayout) GetPriceSets() []map[string]interface{}`

GetPriceSets returns the PriceSets field if non-nil, zero value otherwise.

### GetPriceSetsOk

`func (o *AddLayout200ResponseInstanceTypeLayout) GetPriceSetsOk() (*[]map[string]interface{}, bool)`

GetPriceSetsOk returns a tuple with the PriceSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSets

`func (o *AddLayout200ResponseInstanceTypeLayout) SetPriceSets(v []map[string]interface{})`

SetPriceSets sets PriceSets field to given value.

### HasPriceSets

`func (o *AddLayout200ResponseInstanceTypeLayout) HasPriceSets() bool`

HasPriceSets returns a boolean if a field has been set.

### SetPriceSetsNil

`func (o *AddLayout200ResponseInstanceTypeLayout) SetPriceSetsNil(b bool)`

 SetPriceSetsNil sets the value for PriceSets to be an explicit nil

### UnsetPriceSets
`func (o *AddLayout200ResponseInstanceTypeLayout) UnsetPriceSets()`

UnsetPriceSets ensures that no value is present for PriceSets, not even an explicit nil
### GetSpecTemplates

`func (o *AddLayout200ResponseInstanceTypeLayout) GetSpecTemplates() []map[string]interface{}`

GetSpecTemplates returns the SpecTemplates field if non-nil, zero value otherwise.

### GetSpecTemplatesOk

`func (o *AddLayout200ResponseInstanceTypeLayout) GetSpecTemplatesOk() (*[]map[string]interface{}, bool)`

GetSpecTemplatesOk returns a tuple with the SpecTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecTemplates

`func (o *AddLayout200ResponseInstanceTypeLayout) SetSpecTemplates(v []map[string]interface{})`

SetSpecTemplates sets SpecTemplates field to given value.

### HasSpecTemplates

`func (o *AddLayout200ResponseInstanceTypeLayout) HasSpecTemplates() bool`

HasSpecTemplates returns a boolean if a field has been set.

### SetSpecTemplatesNil

`func (o *AddLayout200ResponseInstanceTypeLayout) SetSpecTemplatesNil(b bool)`

 SetSpecTemplatesNil sets the value for SpecTemplates to be an explicit nil

### UnsetSpecTemplates
`func (o *AddLayout200ResponseInstanceTypeLayout) UnsetSpecTemplates()`

UnsetSpecTemplates ensures that no value is present for SpecTemplates, not even an explicit nil
### GetTfvarSecret

`func (o *AddLayout200ResponseInstanceTypeLayout) GetTfvarSecret() string`

GetTfvarSecret returns the TfvarSecret field if non-nil, zero value otherwise.

### GetTfvarSecretOk

`func (o *AddLayout200ResponseInstanceTypeLayout) GetTfvarSecretOk() (*string, bool)`

GetTfvarSecretOk returns a tuple with the TfvarSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTfvarSecret

`func (o *AddLayout200ResponseInstanceTypeLayout) SetTfvarSecret(v string)`

SetTfvarSecret sets TfvarSecret field to given value.

### HasTfvarSecret

`func (o *AddLayout200ResponseInstanceTypeLayout) HasTfvarSecret() bool`

HasTfvarSecret returns a boolean if a field has been set.

### SetTfvarSecretNil

`func (o *AddLayout200ResponseInstanceTypeLayout) SetTfvarSecretNil(b bool)`

 SetTfvarSecretNil sets the value for TfvarSecret to be an explicit nil

### UnsetTfvarSecret
`func (o *AddLayout200ResponseInstanceTypeLayout) UnsetTfvarSecret()`

UnsetTfvarSecret ensures that no value is present for TfvarSecret, not even an explicit nil
### GetPermissions

`func (o *AddLayout200ResponseInstanceTypeLayout) GetPermissions() AddLayout200ResponseInstanceTypeLayoutPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AddLayout200ResponseInstanceTypeLayout) GetPermissionsOk() (*AddLayout200ResponseInstanceTypeLayoutPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AddLayout200ResponseInstanceTypeLayout) SetPermissions(v AddLayout200ResponseInstanceTypeLayoutPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *AddLayout200ResponseInstanceTypeLayout) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


