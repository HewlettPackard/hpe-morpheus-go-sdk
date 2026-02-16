# GetLayout200ResponseInstanceTypeLayout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**InstanceType** | Pointer to [**GetLayout200ResponseInstanceTypeLayoutInstanceType**](GetLayout200ResponseInstanceTypeLayoutInstanceType.md) |  | [optional] 
**Account** | Pointer to [**GetLayout200ResponseInstanceTypeLayoutAccount**](GetLayout200ResponseInstanceTypeLayoutAccount.md) |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**InstanceVersion** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Creatable** | Pointer to **bool** |  | [optional] 
**MemoryRequirement** | Pointer to **NullableInt64** |  | [optional] 
**SortOrder** | Pointer to **int64** |  | [optional] 
**SupportsConvertToManaged** | Pointer to **NullableBool** |  | [optional] 
**ProvisionType** | Pointer to [**GetLayout200ResponseInstanceTypeLayoutProvisionType**](GetLayout200ResponseInstanceTypeLayoutProvisionType.md) |  | [optional] 
**TaskSets** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ContainerTypes** | Pointer to [**[]GetLayout200ResponseInstanceTypeLayoutContainerTypesInner**](GetLayout200ResponseInstanceTypeLayoutContainerTypesInner.md) |  | [optional] 
**Mounts** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Ports** | Pointer to **[]map[string]interface{}** |  | [optional] 
**OptionTypes** | Pointer to **[]map[string]interface{}** |  | [optional] 
**EnvironmentVariables** | Pointer to **[]map[string]interface{}** |  | [optional] 
**PriceSets** | Pointer to **[]map[string]interface{}** |  | [optional] 
**SpecTemplates** | Pointer to **[]map[string]interface{}** |  | [optional] 
**TfvarSecret** | Pointer to **NullableString** |  | [optional] 
**Permissions** | Pointer to [**GetLayout200ResponseInstanceTypeLayoutPermissions**](GetLayout200ResponseInstanceTypeLayoutPermissions.md) |  | [optional] 

## Methods

### NewGetLayout200ResponseInstanceTypeLayout

`func NewGetLayout200ResponseInstanceTypeLayout() *GetLayout200ResponseInstanceTypeLayout`

NewGetLayout200ResponseInstanceTypeLayout instantiates a new GetLayout200ResponseInstanceTypeLayout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLayout200ResponseInstanceTypeLayoutWithDefaults

`func NewGetLayout200ResponseInstanceTypeLayoutWithDefaults() *GetLayout200ResponseInstanceTypeLayout`

NewGetLayout200ResponseInstanceTypeLayoutWithDefaults instantiates a new GetLayout200ResponseInstanceTypeLayout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetLayout200ResponseInstanceTypeLayout) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetLayout200ResponseInstanceTypeLayout) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetLayout200ResponseInstanceTypeLayout) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetLayout200ResponseInstanceTypeLayout) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInstanceType

`func (o *GetLayout200ResponseInstanceTypeLayout) GetInstanceType() GetLayout200ResponseInstanceTypeLayoutInstanceType`

GetInstanceType returns the InstanceType field if non-nil, zero value otherwise.

### GetInstanceTypeOk

`func (o *GetLayout200ResponseInstanceTypeLayout) GetInstanceTypeOk() (*GetLayout200ResponseInstanceTypeLayoutInstanceType, bool)`

GetInstanceTypeOk returns a tuple with the InstanceType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceType

`func (o *GetLayout200ResponseInstanceTypeLayout) SetInstanceType(v GetLayout200ResponseInstanceTypeLayoutInstanceType)`

SetInstanceType sets InstanceType field to given value.

### HasInstanceType

`func (o *GetLayout200ResponseInstanceTypeLayout) HasInstanceType() bool`

HasInstanceType returns a boolean if a field has been set.

### GetAccount

`func (o *GetLayout200ResponseInstanceTypeLayout) GetAccount() GetLayout200ResponseInstanceTypeLayoutAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetLayout200ResponseInstanceTypeLayout) GetAccountOk() (*GetLayout200ResponseInstanceTypeLayoutAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetLayout200ResponseInstanceTypeLayout) SetAccount(v GetLayout200ResponseInstanceTypeLayoutAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetLayout200ResponseInstanceTypeLayout) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetCode

`func (o *GetLayout200ResponseInstanceTypeLayout) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetLayout200ResponseInstanceTypeLayout) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetLayout200ResponseInstanceTypeLayout) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetLayout200ResponseInstanceTypeLayout) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *GetLayout200ResponseInstanceTypeLayout) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetLayout200ResponseInstanceTypeLayout) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetLayout200ResponseInstanceTypeLayout) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetLayout200ResponseInstanceTypeLayout) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *GetLayout200ResponseInstanceTypeLayout) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetLayout200ResponseInstanceTypeLayout) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetLayout200ResponseInstanceTypeLayout) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetLayout200ResponseInstanceTypeLayout) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *GetLayout200ResponseInstanceTypeLayout) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *GetLayout200ResponseInstanceTypeLayout) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetInstanceVersion

`func (o *GetLayout200ResponseInstanceTypeLayout) GetInstanceVersion() string`

GetInstanceVersion returns the InstanceVersion field if non-nil, zero value otherwise.

### GetInstanceVersionOk

`func (o *GetLayout200ResponseInstanceTypeLayout) GetInstanceVersionOk() (*string, bool)`

GetInstanceVersionOk returns a tuple with the InstanceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceVersion

`func (o *GetLayout200ResponseInstanceTypeLayout) SetInstanceVersion(v string)`

SetInstanceVersion sets InstanceVersion field to given value.

### HasInstanceVersion

`func (o *GetLayout200ResponseInstanceTypeLayout) HasInstanceVersion() bool`

HasInstanceVersion returns a boolean if a field has been set.

### GetDescription

`func (o *GetLayout200ResponseInstanceTypeLayout) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetLayout200ResponseInstanceTypeLayout) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetLayout200ResponseInstanceTypeLayout) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetLayout200ResponseInstanceTypeLayout) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetLayout200ResponseInstanceTypeLayout) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetLayout200ResponseInstanceTypeLayout) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetCreatable

`func (o *GetLayout200ResponseInstanceTypeLayout) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *GetLayout200ResponseInstanceTypeLayout) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *GetLayout200ResponseInstanceTypeLayout) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *GetLayout200ResponseInstanceTypeLayout) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetMemoryRequirement

`func (o *GetLayout200ResponseInstanceTypeLayout) GetMemoryRequirement() int64`

GetMemoryRequirement returns the MemoryRequirement field if non-nil, zero value otherwise.

### GetMemoryRequirementOk

`func (o *GetLayout200ResponseInstanceTypeLayout) GetMemoryRequirementOk() (*int64, bool)`

GetMemoryRequirementOk returns a tuple with the MemoryRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryRequirement

`func (o *GetLayout200ResponseInstanceTypeLayout) SetMemoryRequirement(v int64)`

SetMemoryRequirement sets MemoryRequirement field to given value.

### HasMemoryRequirement

`func (o *GetLayout200ResponseInstanceTypeLayout) HasMemoryRequirement() bool`

HasMemoryRequirement returns a boolean if a field has been set.

### SetMemoryRequirementNil

`func (o *GetLayout200ResponseInstanceTypeLayout) SetMemoryRequirementNil(b bool)`

 SetMemoryRequirementNil sets the value for MemoryRequirement to be an explicit nil

### UnsetMemoryRequirement
`func (o *GetLayout200ResponseInstanceTypeLayout) UnsetMemoryRequirement()`

UnsetMemoryRequirement ensures that no value is present for MemoryRequirement, not even an explicit nil
### GetSortOrder

`func (o *GetLayout200ResponseInstanceTypeLayout) GetSortOrder() int64`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *GetLayout200ResponseInstanceTypeLayout) GetSortOrderOk() (*int64, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *GetLayout200ResponseInstanceTypeLayout) SetSortOrder(v int64)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *GetLayout200ResponseInstanceTypeLayout) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetSupportsConvertToManaged

`func (o *GetLayout200ResponseInstanceTypeLayout) GetSupportsConvertToManaged() bool`

GetSupportsConvertToManaged returns the SupportsConvertToManaged field if non-nil, zero value otherwise.

### GetSupportsConvertToManagedOk

`func (o *GetLayout200ResponseInstanceTypeLayout) GetSupportsConvertToManagedOk() (*bool, bool)`

GetSupportsConvertToManagedOk returns a tuple with the SupportsConvertToManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsConvertToManaged

`func (o *GetLayout200ResponseInstanceTypeLayout) SetSupportsConvertToManaged(v bool)`

SetSupportsConvertToManaged sets SupportsConvertToManaged field to given value.

### HasSupportsConvertToManaged

`func (o *GetLayout200ResponseInstanceTypeLayout) HasSupportsConvertToManaged() bool`

HasSupportsConvertToManaged returns a boolean if a field has been set.

### SetSupportsConvertToManagedNil

`func (o *GetLayout200ResponseInstanceTypeLayout) SetSupportsConvertToManagedNil(b bool)`

 SetSupportsConvertToManagedNil sets the value for SupportsConvertToManaged to be an explicit nil

### UnsetSupportsConvertToManaged
`func (o *GetLayout200ResponseInstanceTypeLayout) UnsetSupportsConvertToManaged()`

UnsetSupportsConvertToManaged ensures that no value is present for SupportsConvertToManaged, not even an explicit nil
### GetProvisionType

`func (o *GetLayout200ResponseInstanceTypeLayout) GetProvisionType() GetLayout200ResponseInstanceTypeLayoutProvisionType`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *GetLayout200ResponseInstanceTypeLayout) GetProvisionTypeOk() (*GetLayout200ResponseInstanceTypeLayoutProvisionType, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *GetLayout200ResponseInstanceTypeLayout) SetProvisionType(v GetLayout200ResponseInstanceTypeLayoutProvisionType)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *GetLayout200ResponseInstanceTypeLayout) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetTaskSets

`func (o *GetLayout200ResponseInstanceTypeLayout) GetTaskSets() []map[string]interface{}`

GetTaskSets returns the TaskSets field if non-nil, zero value otherwise.

### GetTaskSetsOk

`func (o *GetLayout200ResponseInstanceTypeLayout) GetTaskSetsOk() (*[]map[string]interface{}, bool)`

GetTaskSetsOk returns a tuple with the TaskSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSets

`func (o *GetLayout200ResponseInstanceTypeLayout) SetTaskSets(v []map[string]interface{})`

SetTaskSets sets TaskSets field to given value.

### HasTaskSets

`func (o *GetLayout200ResponseInstanceTypeLayout) HasTaskSets() bool`

HasTaskSets returns a boolean if a field has been set.

### SetTaskSetsNil

`func (o *GetLayout200ResponseInstanceTypeLayout) SetTaskSetsNil(b bool)`

 SetTaskSetsNil sets the value for TaskSets to be an explicit nil

### UnsetTaskSets
`func (o *GetLayout200ResponseInstanceTypeLayout) UnsetTaskSets()`

UnsetTaskSets ensures that no value is present for TaskSets, not even an explicit nil
### GetContainerTypes

`func (o *GetLayout200ResponseInstanceTypeLayout) GetContainerTypes() []GetLayout200ResponseInstanceTypeLayoutContainerTypesInner`

GetContainerTypes returns the ContainerTypes field if non-nil, zero value otherwise.

### GetContainerTypesOk

`func (o *GetLayout200ResponseInstanceTypeLayout) GetContainerTypesOk() (*[]GetLayout200ResponseInstanceTypeLayoutContainerTypesInner, bool)`

GetContainerTypesOk returns a tuple with the ContainerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTypes

`func (o *GetLayout200ResponseInstanceTypeLayout) SetContainerTypes(v []GetLayout200ResponseInstanceTypeLayoutContainerTypesInner)`

SetContainerTypes sets ContainerTypes field to given value.

### HasContainerTypes

`func (o *GetLayout200ResponseInstanceTypeLayout) HasContainerTypes() bool`

HasContainerTypes returns a boolean if a field has been set.

### GetMounts

`func (o *GetLayout200ResponseInstanceTypeLayout) GetMounts() []map[string]interface{}`

GetMounts returns the Mounts field if non-nil, zero value otherwise.

### GetMountsOk

`func (o *GetLayout200ResponseInstanceTypeLayout) GetMountsOk() (*[]map[string]interface{}, bool)`

GetMountsOk returns a tuple with the Mounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMounts

`func (o *GetLayout200ResponseInstanceTypeLayout) SetMounts(v []map[string]interface{})`

SetMounts sets Mounts field to given value.

### HasMounts

`func (o *GetLayout200ResponseInstanceTypeLayout) HasMounts() bool`

HasMounts returns a boolean if a field has been set.

### SetMountsNil

`func (o *GetLayout200ResponseInstanceTypeLayout) SetMountsNil(b bool)`

 SetMountsNil sets the value for Mounts to be an explicit nil

### UnsetMounts
`func (o *GetLayout200ResponseInstanceTypeLayout) UnsetMounts()`

UnsetMounts ensures that no value is present for Mounts, not even an explicit nil
### GetPorts

`func (o *GetLayout200ResponseInstanceTypeLayout) GetPorts() []map[string]interface{}`

GetPorts returns the Ports field if non-nil, zero value otherwise.

### GetPortsOk

`func (o *GetLayout200ResponseInstanceTypeLayout) GetPortsOk() (*[]map[string]interface{}, bool)`

GetPortsOk returns a tuple with the Ports field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPorts

`func (o *GetLayout200ResponseInstanceTypeLayout) SetPorts(v []map[string]interface{})`

SetPorts sets Ports field to given value.

### HasPorts

`func (o *GetLayout200ResponseInstanceTypeLayout) HasPorts() bool`

HasPorts returns a boolean if a field has been set.

### SetPortsNil

`func (o *GetLayout200ResponseInstanceTypeLayout) SetPortsNil(b bool)`

 SetPortsNil sets the value for Ports to be an explicit nil

### UnsetPorts
`func (o *GetLayout200ResponseInstanceTypeLayout) UnsetPorts()`

UnsetPorts ensures that no value is present for Ports, not even an explicit nil
### GetOptionTypes

`func (o *GetLayout200ResponseInstanceTypeLayout) GetOptionTypes() []map[string]interface{}`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *GetLayout200ResponseInstanceTypeLayout) GetOptionTypesOk() (*[]map[string]interface{}, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *GetLayout200ResponseInstanceTypeLayout) SetOptionTypes(v []map[string]interface{})`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *GetLayout200ResponseInstanceTypeLayout) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### SetOptionTypesNil

`func (o *GetLayout200ResponseInstanceTypeLayout) SetOptionTypesNil(b bool)`

 SetOptionTypesNil sets the value for OptionTypes to be an explicit nil

### UnsetOptionTypes
`func (o *GetLayout200ResponseInstanceTypeLayout) UnsetOptionTypes()`

UnsetOptionTypes ensures that no value is present for OptionTypes, not even an explicit nil
### GetEnvironmentVariables

`func (o *GetLayout200ResponseInstanceTypeLayout) GetEnvironmentVariables() []map[string]interface{}`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *GetLayout200ResponseInstanceTypeLayout) GetEnvironmentVariablesOk() (*[]map[string]interface{}, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *GetLayout200ResponseInstanceTypeLayout) SetEnvironmentVariables(v []map[string]interface{})`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *GetLayout200ResponseInstanceTypeLayout) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### SetEnvironmentVariablesNil

`func (o *GetLayout200ResponseInstanceTypeLayout) SetEnvironmentVariablesNil(b bool)`

 SetEnvironmentVariablesNil sets the value for EnvironmentVariables to be an explicit nil

### UnsetEnvironmentVariables
`func (o *GetLayout200ResponseInstanceTypeLayout) UnsetEnvironmentVariables()`

UnsetEnvironmentVariables ensures that no value is present for EnvironmentVariables, not even an explicit nil
### GetPriceSets

`func (o *GetLayout200ResponseInstanceTypeLayout) GetPriceSets() []map[string]interface{}`

GetPriceSets returns the PriceSets field if non-nil, zero value otherwise.

### GetPriceSetsOk

`func (o *GetLayout200ResponseInstanceTypeLayout) GetPriceSetsOk() (*[]map[string]interface{}, bool)`

GetPriceSetsOk returns a tuple with the PriceSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSets

`func (o *GetLayout200ResponseInstanceTypeLayout) SetPriceSets(v []map[string]interface{})`

SetPriceSets sets PriceSets field to given value.

### HasPriceSets

`func (o *GetLayout200ResponseInstanceTypeLayout) HasPriceSets() bool`

HasPriceSets returns a boolean if a field has been set.

### SetPriceSetsNil

`func (o *GetLayout200ResponseInstanceTypeLayout) SetPriceSetsNil(b bool)`

 SetPriceSetsNil sets the value for PriceSets to be an explicit nil

### UnsetPriceSets
`func (o *GetLayout200ResponseInstanceTypeLayout) UnsetPriceSets()`

UnsetPriceSets ensures that no value is present for PriceSets, not even an explicit nil
### GetSpecTemplates

`func (o *GetLayout200ResponseInstanceTypeLayout) GetSpecTemplates() []map[string]interface{}`

GetSpecTemplates returns the SpecTemplates field if non-nil, zero value otherwise.

### GetSpecTemplatesOk

`func (o *GetLayout200ResponseInstanceTypeLayout) GetSpecTemplatesOk() (*[]map[string]interface{}, bool)`

GetSpecTemplatesOk returns a tuple with the SpecTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecTemplates

`func (o *GetLayout200ResponseInstanceTypeLayout) SetSpecTemplates(v []map[string]interface{})`

SetSpecTemplates sets SpecTemplates field to given value.

### HasSpecTemplates

`func (o *GetLayout200ResponseInstanceTypeLayout) HasSpecTemplates() bool`

HasSpecTemplates returns a boolean if a field has been set.

### SetSpecTemplatesNil

`func (o *GetLayout200ResponseInstanceTypeLayout) SetSpecTemplatesNil(b bool)`

 SetSpecTemplatesNil sets the value for SpecTemplates to be an explicit nil

### UnsetSpecTemplates
`func (o *GetLayout200ResponseInstanceTypeLayout) UnsetSpecTemplates()`

UnsetSpecTemplates ensures that no value is present for SpecTemplates, not even an explicit nil
### GetTfvarSecret

`func (o *GetLayout200ResponseInstanceTypeLayout) GetTfvarSecret() string`

GetTfvarSecret returns the TfvarSecret field if non-nil, zero value otherwise.

### GetTfvarSecretOk

`func (o *GetLayout200ResponseInstanceTypeLayout) GetTfvarSecretOk() (*string, bool)`

GetTfvarSecretOk returns a tuple with the TfvarSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTfvarSecret

`func (o *GetLayout200ResponseInstanceTypeLayout) SetTfvarSecret(v string)`

SetTfvarSecret sets TfvarSecret field to given value.

### HasTfvarSecret

`func (o *GetLayout200ResponseInstanceTypeLayout) HasTfvarSecret() bool`

HasTfvarSecret returns a boolean if a field has been set.

### SetTfvarSecretNil

`func (o *GetLayout200ResponseInstanceTypeLayout) SetTfvarSecretNil(b bool)`

 SetTfvarSecretNil sets the value for TfvarSecret to be an explicit nil

### UnsetTfvarSecret
`func (o *GetLayout200ResponseInstanceTypeLayout) UnsetTfvarSecret()`

UnsetTfvarSecret ensures that no value is present for TfvarSecret, not even an explicit nil
### GetPermissions

`func (o *GetLayout200ResponseInstanceTypeLayout) GetPermissions() GetLayout200ResponseInstanceTypeLayoutPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *GetLayout200ResponseInstanceTypeLayout) GetPermissionsOk() (*GetLayout200ResponseInstanceTypeLayoutPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *GetLayout200ResponseInstanceTypeLayout) SetPermissions(v GetLayout200ResponseInstanceTypeLayoutPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *GetLayout200ResponseInstanceTypeLayout) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


