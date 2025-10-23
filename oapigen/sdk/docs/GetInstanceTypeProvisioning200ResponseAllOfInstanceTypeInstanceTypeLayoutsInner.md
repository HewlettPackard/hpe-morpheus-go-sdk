# GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**InstanceType** | Pointer to [**ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**Account** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInnerInstance**](GetAlerts200ResponseAllOfCheckGroupsInnerInstance.md) |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**InstanceVersion** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Creatable** | Pointer to **bool** |  | [optional] 
**MemoryRequirement** | Pointer to **NullableInt64** |  | [optional] 
**SortOrder** | Pointer to **int64** |  | [optional] 
**SupportsConvertToManaged** | Pointer to **NullableBool** |  | [optional] 
**ProvisionType** | Pointer to [**GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerProvisionType**](GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerProvisionType.md) |  | [optional] 
**TaskSets** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ContainerTypes** | Pointer to [**[]GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner**](GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner.md) |  | [optional] 
**Mounts** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Ports** | Pointer to **[]map[string]interface{}** |  | [optional] 
**OptionTypes** | Pointer to **[]map[string]interface{}** |  | [optional] 
**EnvironmentVariables** | Pointer to **[]map[string]interface{}** |  | [optional] 
**PriceSets** | Pointer to **[]map[string]interface{}** |  | [optional] 
**SpecTemplates** | Pointer to **[]map[string]interface{}** |  | [optional] 
**TfvarSecret** | Pointer to **NullableString** |  | [optional] 
**Permissions** | Pointer to [**GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerPermissions**](GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerPermissions.md) |  | [optional] 

## Methods

### NewGetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner

`func NewGetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner() *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner`

NewGetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner instantiates a new GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerWithDefaults

`func NewGetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerWithDefaults() *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner`

NewGetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerWithDefaults instantiates a new GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstanceType

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetInstanceType() ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetInstanceTypeOk() (*ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) SetInstanceType(v ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetAccount

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetAccount() GetAlerts200ResponseAllOfCheckGroupsInnerInstance`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetAccountOk() (*GetAlerts200ResponseAllOfCheckGroupsInnerInstance, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) SetAccount(v GetAlerts200ResponseAllOfCheckGroupsInnerInstance)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCode

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetInstanceVersion

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetInstanceVersion() string`

GetInstanceVersion returns the InstanceVersion field if non-nil, zero value otherwise.

### GetInstanceVersionOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetInstanceVersionOk() (*string, bool)`

GetInstanceVersionOk returns a tuple with the InstanceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceVersion

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) SetInstanceVersion(v string)`

SetInstanceVersion sets InstanceVersion field to given value.

### HasInstanceVersion

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) HasInstanceVersion() bool`

HasInstanceVersion returns a boolean if a field has been set.

### GetDescription

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCreatable

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetMemoryRequirement

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetMemoryRequirement() int64`

GetMemoryRequirement returns the MemoryRequirement field if non-nil, zero value otherwise.

### GetMemoryRequirementOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetMemoryRequirementOk() (*int64, bool)`

GetMemoryRequirementOk returns a tuple with the MemoryRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryRequirement

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) SetMemoryRequirement(v int64)`

SetMemoryRequirement sets MemoryRequirement field to given value.

### HasMemoryRequirement

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) HasMemoryRequirement() bool`

HasMemoryRequirement returns a boolean if a field has been set.

### SetMemoryRequirementNil

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) SetMemoryRequirementNil(b bool)`

 SetMemoryRequirementNil sets the value for MemoryRequirement to be an explicit nil

### UnsetMemoryRequirement
`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) UnsetMemoryRequirement()`

UnsetMemoryRequirement ensures that no value is present for MemoryRequirement, not even an explicit nil
### GetSortOrder

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetSortOrder() int64`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetSortOrderOk() (*int64, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) SetSortOrder(v int64)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetSupportsConvertToManaged

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetSupportsConvertToManaged() bool`

GetSupportsConvertToManaged returns the SupportsConvertToManaged field if non-nil, zero value otherwise.

### GetSupportsConvertToManagedOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetSupportsConvertToManagedOk() (*bool, bool)`

GetSupportsConvertToManagedOk returns a tuple with the SupportsConvertToManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsConvertToManaged

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) SetSupportsConvertToManaged(v bool)`

SetSupportsConvertToManaged sets SupportsConvertToManaged field to given value.

### HasSupportsConvertToManaged

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) HasSupportsConvertToManaged() bool`

HasSupportsConvertToManaged returns a boolean if a field has been set.

### SetSupportsConvertToManagedNil

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) SetSupportsConvertToManagedNil(b bool)`

 SetSupportsConvertToManagedNil sets the value for SupportsConvertToManaged to be an explicit nil

### UnsetSupportsConvertToManaged
`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) UnsetSupportsConvertToManaged()`

UnsetSupportsConvertToManaged ensures that no value is present for SupportsConvertToManaged, not even an explicit nil
### GetProvisionType

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetProvisionType() GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerProvisionType`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetProvisionTypeOk() (*GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerProvisionType, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) SetProvisionType(v GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerProvisionType)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetTaskSets

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetTaskSets() []map[string]interface{}`

GetTaskSets returns the TaskSets field if non-nil, zero value otherwise.

### GetTaskSetsOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetTaskSetsOk() (*[]map[string]interface{}, bool)`

GetTaskSetsOk returns a tuple with the TaskSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSets

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) SetTaskSets(v []map[string]interface{})`

SetTaskSets sets TaskSets field to given value.

### HasTaskSets

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) HasTaskSets() bool`

HasTaskSets returns a boolean if a field has been set.

### SetTaskSetsNil

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) SetTaskSetsNil(b bool)`

 SetTaskSetsNil sets the value for TaskSets to be an explicit nil

### UnsetTaskSets
`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) UnsetTaskSets()`

UnsetTaskSets ensures that no value is present for TaskSets, not even an explicit nil
### GetContainerTypes

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetContainerTypes() []GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner`

GetContainerTypes returns the ContainerTypes field if non-nil, zero value otherwise.

### GetContainerTypesOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetContainerTypesOk() (*[]GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner, bool)`

GetContainerTypesOk returns a tuple with the ContainerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTypes

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) SetContainerTypes(v []GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerContainerTypesInner)`

SetContainerTypes sets ContainerTypes field to given value.

### HasContainerTypes

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) HasContainerTypes() bool`

HasContainerTypes returns a boolean if a field has been set.

### GetMounts

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetMounts() []map[string]interface{}`

GetMounts returns the Mounts field if non-nil, zero value otherwise.

### GetMountsOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetMountsOk() (*[]map[string]interface{}, bool)`

GetMountsOk returns a tuple with the Mounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMounts

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) SetMounts(v []map[string]interface{})`

SetMounts sets Mounts field to given value.

### HasMounts

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) HasMounts() bool`

HasMounts returns a boolean if a field has been set.

### SetMountsNil

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) SetMountsNil(b bool)`

 SetMountsNil sets the value for Mounts to be an explicit nil

### UnsetMounts
`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) UnsetMounts()`

UnsetMounts ensures that no value is present for Mounts, not even an explicit nil
### GetPorts

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetPorts() []map[string]interface{}`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetPortsOk() (*[]map[string]interface{}, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) SetPorts(v []map[string]interface{})`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### SetPortsNil

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) SetPortsNil(b bool)`

 SetPortsNil sets the value for Ports to be an explicit nil

### UnsetPorts
`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) UnsetPorts()`

UnsetPorts ensures that no value is present for Ports, not even an explicit nil
### GetOptionTypes

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetOptionTypes() []map[string]interface{}`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetOptionTypesOk() (*[]map[string]interface{}, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) SetOptionTypes(v []map[string]interface{})`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### SetOptionTypesNil

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) SetOptionTypesNil(b bool)`

 SetOptionTypesNil sets the value for OptionTypes to be an explicit nil

### UnsetOptionTypes
`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) UnsetOptionTypes()`

UnsetOptionTypes ensures that no value is present for OptionTypes, not even an explicit nil
### GetEnvironmentVariables

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetEnvironmentVariables() []map[string]interface{}`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetEnvironmentVariablesOk() (*[]map[string]interface{}, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) SetEnvironmentVariables(v []map[string]interface{})`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### SetEnvironmentVariablesNil

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) SetEnvironmentVariablesNil(b bool)`

 SetEnvironmentVariablesNil sets the value for EnvironmentVariables to be an explicit nil

### UnsetEnvironmentVariables
`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) UnsetEnvironmentVariables()`

UnsetEnvironmentVariables ensures that no value is present for EnvironmentVariables, not even an explicit nil
### GetPriceSets

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetPriceSets() []map[string]interface{}`

GetPriceSets returns the PriceSets field if non-nil, zero value otherwise.

### GetPriceSetsOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetPriceSetsOk() (*[]map[string]interface{}, bool)`

GetPriceSetsOk returns a tuple with the PriceSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSets

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) SetPriceSets(v []map[string]interface{})`

SetPriceSets sets PriceSets field to given value.

### HasPriceSets

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) HasPriceSets() bool`

HasPriceSets returns a boolean if a field has been set.

### SetPriceSetsNil

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) SetPriceSetsNil(b bool)`

 SetPriceSetsNil sets the value for PriceSets to be an explicit nil

### UnsetPriceSets
`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) UnsetPriceSets()`

UnsetPriceSets ensures that no value is present for PriceSets, not even an explicit nil
### GetSpecTemplates

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetSpecTemplates() []map[string]interface{}`

GetSpecTemplates returns the SpecTemplates field if non-nil, zero value otherwise.

### GetSpecTemplatesOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetSpecTemplatesOk() (*[]map[string]interface{}, bool)`

GetSpecTemplatesOk returns a tuple with the SpecTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecTemplates

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) SetSpecTemplates(v []map[string]interface{})`

SetSpecTemplates sets SpecTemplates field to given value.

### HasSpecTemplates

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) HasSpecTemplates() bool`

HasSpecTemplates returns a boolean if a field has been set.

### SetSpecTemplatesNil

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) SetSpecTemplatesNil(b bool)`

 SetSpecTemplatesNil sets the value for SpecTemplates to be an explicit nil

### UnsetSpecTemplates
`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) UnsetSpecTemplates()`

UnsetSpecTemplates ensures that no value is present for SpecTemplates, not even an explicit nil
### GetTfvarSecret

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetTfvarSecret() string`

GetTfvarSecret returns the TfvarSecret field if non-nil, zero value otherwise.

### GetTfvarSecretOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetTfvarSecretOk() (*string, bool)`

GetTfvarSecretOk returns a tuple with the TfvarSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTfvarSecret

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) SetTfvarSecret(v string)`

SetTfvarSecret sets TfvarSecret field to given value.

### HasTfvarSecret

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) HasTfvarSecret() bool`

HasTfvarSecret returns a boolean if a field has been set.

### SetTfvarSecretNil

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) SetTfvarSecretNil(b bool)`

 SetTfvarSecretNil sets the value for TfvarSecret to be an explicit nil

### UnsetTfvarSecret
`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) UnsetTfvarSecret()`

UnsetTfvarSecret ensures that no value is present for TfvarSecret, not even an explicit nil
### GetPermissions

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetPermissions() GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) GetPermissionsOk() (*GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) SetPermissions(v GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInnerPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *GetInstanceTypeProvisioning200ResponseAllOfInstanceTypeInstanceTypeLayoutsInner) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


