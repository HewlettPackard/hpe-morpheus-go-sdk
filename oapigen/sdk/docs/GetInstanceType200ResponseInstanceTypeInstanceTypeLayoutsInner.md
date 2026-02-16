# GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**InstanceType** | Pointer to [**GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerInstanceType**](GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerInstanceType.md) |  | [optional] 
**Account** | Pointer to [**GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerAccount**](GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerAccount.md) |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**InstanceVersion** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Creatable** | Pointer to **bool** |  | [optional] 
**MemoryRequirement** | Pointer to **NullableInt64** |  | [optional] 
**SortOrder** | Pointer to **int64** |  | [optional] 
**SupportsConvertToManaged** | Pointer to **NullableBool** |  | [optional] 
**ProvisionType** | Pointer to [**GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType**](GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType.md) |  | [optional] 
**TaskSets** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ContainerTypes** | Pointer to [**[]GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner**](GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner.md) |  | [optional] 
**Mounts** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Ports** | Pointer to **[]map[string]interface{}** |  | [optional] 
**OptionTypes** | Pointer to **[]map[string]interface{}** |  | [optional] 
**EnvironmentVariables** | Pointer to **[]map[string]interface{}** |  | [optional] 
**PriceSets** | Pointer to **[]map[string]interface{}** |  | [optional] 
**SpecTemplates** | Pointer to **[]map[string]interface{}** |  | [optional] 
**TfvarSecret** | Pointer to **NullableString** |  | [optional] 
**Permissions** | Pointer to [**GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerPermissions**](GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerPermissions.md) |  | [optional] 

## Methods

### NewGetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner

`func NewGetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner() *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner`

NewGetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner instantiates a new GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerWithDefaults

`func NewGetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerWithDefaults() *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner`

NewGetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerWithDefaults instantiates a new GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstanceType

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetInstanceType() GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerInstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetInstanceTypeOk() (*GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerInstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetInstanceType(v GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerInstanceType)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetAccount

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetAccount() GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetAccountOk() (*GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetAccount(v GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCode

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetInstanceVersion

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetInstanceVersion() string`

GetInstanceVersion returns the InstanceVersion field if non-nil, zero value otherwise.

### GetInstanceVersionOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetInstanceVersionOk() (*string, bool)`

GetInstanceVersionOk returns a tuple with the InstanceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceVersion

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetInstanceVersion(v string)`

SetInstanceVersion sets InstanceVersion field to given value.

### HasInstanceVersion

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) HasInstanceVersion() bool`

HasInstanceVersion returns a boolean if a field has been set.

### GetDescription

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCreatable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetMemoryRequirement

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetMemoryRequirement() int64`

GetMemoryRequirement returns the MemoryRequirement field if non-nil, zero value otherwise.

### GetMemoryRequirementOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetMemoryRequirementOk() (*int64, bool)`

GetMemoryRequirementOk returns a tuple with the MemoryRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryRequirement

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetMemoryRequirement(v int64)`

SetMemoryRequirement sets MemoryRequirement field to given value.

### HasMemoryRequirement

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) HasMemoryRequirement() bool`

HasMemoryRequirement returns a boolean if a field has been set.

### SetMemoryRequirementNil

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetMemoryRequirementNil(b bool)`

 SetMemoryRequirementNil sets the value for MemoryRequirement to be an explicit nil

### UnsetMemoryRequirement
`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) UnsetMemoryRequirement()`

UnsetMemoryRequirement ensures that no value is present for MemoryRequirement, not even an explicit nil
### GetSortOrder

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetSortOrder() int64`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetSortOrderOk() (*int64, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetSortOrder(v int64)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetSupportsConvertToManaged

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetSupportsConvertToManaged() bool`

GetSupportsConvertToManaged returns the SupportsConvertToManaged field if non-nil, zero value otherwise.

### GetSupportsConvertToManagedOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetSupportsConvertToManagedOk() (*bool, bool)`

GetSupportsConvertToManagedOk returns a tuple with the SupportsConvertToManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsConvertToManaged

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetSupportsConvertToManaged(v bool)`

SetSupportsConvertToManaged sets SupportsConvertToManaged field to given value.

### HasSupportsConvertToManaged

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) HasSupportsConvertToManaged() bool`

HasSupportsConvertToManaged returns a boolean if a field has been set.

### SetSupportsConvertToManagedNil

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetSupportsConvertToManagedNil(b bool)`

 SetSupportsConvertToManagedNil sets the value for SupportsConvertToManaged to be an explicit nil

### UnsetSupportsConvertToManaged
`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) UnsetSupportsConvertToManaged()`

UnsetSupportsConvertToManaged ensures that no value is present for SupportsConvertToManaged, not even an explicit nil
### GetProvisionType

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetProvisionType() GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetProvisionTypeOk() (*GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetProvisionType(v GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerProvisionType)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetTaskSets

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetTaskSets() []map[string]interface{}`

GetTaskSets returns the TaskSets field if non-nil, zero value otherwise.

### GetTaskSetsOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetTaskSetsOk() (*[]map[string]interface{}, bool)`

GetTaskSetsOk returns a tuple with the TaskSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSets

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetTaskSets(v []map[string]interface{})`

SetTaskSets sets TaskSets field to given value.

### HasTaskSets

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) HasTaskSets() bool`

HasTaskSets returns a boolean if a field has been set.

### SetTaskSetsNil

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetTaskSetsNil(b bool)`

 SetTaskSetsNil sets the value for TaskSets to be an explicit nil

### UnsetTaskSets
`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) UnsetTaskSets()`

UnsetTaskSets ensures that no value is present for TaskSets, not even an explicit nil
### GetContainerTypes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetContainerTypes() []GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner`

GetContainerTypes returns the ContainerTypes field if non-nil, zero value otherwise.

### GetContainerTypesOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetContainerTypesOk() (*[]GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner, bool)`

GetContainerTypesOk returns a tuple with the ContainerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTypes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetContainerTypes(v []GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner)`

SetContainerTypes sets ContainerTypes field to given value.

### HasContainerTypes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) HasContainerTypes() bool`

HasContainerTypes returns a boolean if a field has been set.

### GetMounts

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetMounts() []map[string]interface{}`

GetMounts returns the Mounts field if non-nil, zero value otherwise.

### GetMountsOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetMountsOk() (*[]map[string]interface{}, bool)`

GetMountsOk returns a tuple with the Mounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMounts

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetMounts(v []map[string]interface{})`

SetMounts sets Mounts field to given value.

### HasMounts

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) HasMounts() bool`

HasMounts returns a boolean if a field has been set.

### SetMountsNil

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetMountsNil(b bool)`

 SetMountsNil sets the value for Mounts to be an explicit nil

### UnsetMounts
`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) UnsetMounts()`

UnsetMounts ensures that no value is present for Mounts, not even an explicit nil
### GetPorts

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetPorts() []map[string]interface{}`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetPortsOk() (*[]map[string]interface{}, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetPorts(v []map[string]interface{})`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### SetPortsNil

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetPortsNil(b bool)`

 SetPortsNil sets the value for Ports to be an explicit nil

### UnsetPorts
`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) UnsetPorts()`

UnsetPorts ensures that no value is present for Ports, not even an explicit nil
### GetOptionTypes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetOptionTypes() []map[string]interface{}`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetOptionTypesOk() (*[]map[string]interface{}, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetOptionTypes(v []map[string]interface{})`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### SetOptionTypesNil

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetOptionTypesNil(b bool)`

 SetOptionTypesNil sets the value for OptionTypes to be an explicit nil

### UnsetOptionTypes
`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) UnsetOptionTypes()`

UnsetOptionTypes ensures that no value is present for OptionTypes, not even an explicit nil
### GetEnvironmentVariables

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetEnvironmentVariables() []map[string]interface{}`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetEnvironmentVariablesOk() (*[]map[string]interface{}, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetEnvironmentVariables(v []map[string]interface{})`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### SetEnvironmentVariablesNil

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetEnvironmentVariablesNil(b bool)`

 SetEnvironmentVariablesNil sets the value for EnvironmentVariables to be an explicit nil

### UnsetEnvironmentVariables
`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) UnsetEnvironmentVariables()`

UnsetEnvironmentVariables ensures that no value is present for EnvironmentVariables, not even an explicit nil
### GetPriceSets

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetPriceSets() []map[string]interface{}`

GetPriceSets returns the PriceSets field if non-nil, zero value otherwise.

### GetPriceSetsOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetPriceSetsOk() (*[]map[string]interface{}, bool)`

GetPriceSetsOk returns a tuple with the PriceSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSets

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetPriceSets(v []map[string]interface{})`

SetPriceSets sets PriceSets field to given value.

### HasPriceSets

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) HasPriceSets() bool`

HasPriceSets returns a boolean if a field has been set.

### SetPriceSetsNil

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetPriceSetsNil(b bool)`

 SetPriceSetsNil sets the value for PriceSets to be an explicit nil

### UnsetPriceSets
`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) UnsetPriceSets()`

UnsetPriceSets ensures that no value is present for PriceSets, not even an explicit nil
### GetSpecTemplates

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetSpecTemplates() []map[string]interface{}`

GetSpecTemplates returns the SpecTemplates field if non-nil, zero value otherwise.

### GetSpecTemplatesOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetSpecTemplatesOk() (*[]map[string]interface{}, bool)`

GetSpecTemplatesOk returns a tuple with the SpecTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecTemplates

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetSpecTemplates(v []map[string]interface{})`

SetSpecTemplates sets SpecTemplates field to given value.

### HasSpecTemplates

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) HasSpecTemplates() bool`

HasSpecTemplates returns a boolean if a field has been set.

### SetSpecTemplatesNil

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetSpecTemplatesNil(b bool)`

 SetSpecTemplatesNil sets the value for SpecTemplates to be an explicit nil

### UnsetSpecTemplates
`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) UnsetSpecTemplates()`

UnsetSpecTemplates ensures that no value is present for SpecTemplates, not even an explicit nil
### GetTfvarSecret

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetTfvarSecret() string`

GetTfvarSecret returns the TfvarSecret field if non-nil, zero value otherwise.

### GetTfvarSecretOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetTfvarSecretOk() (*string, bool)`

GetTfvarSecretOk returns a tuple with the TfvarSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTfvarSecret

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetTfvarSecret(v string)`

SetTfvarSecret sets TfvarSecret field to given value.

### HasTfvarSecret

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) HasTfvarSecret() bool`

HasTfvarSecret returns a boolean if a field has been set.

### SetTfvarSecretNil

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetTfvarSecretNil(b bool)`

 SetTfvarSecretNil sets the value for TfvarSecret to be an explicit nil

### UnsetTfvarSecret
`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) UnsetTfvarSecret()`

UnsetTfvarSecret ensures that no value is present for TfvarSecret, not even an explicit nil
### GetPermissions

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetPermissions() GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) GetPermissionsOk() (*GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) SetPermissions(v GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInnerPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *GetInstanceType200ResponseInstanceTypeInstanceTypeLayoutsInner) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


