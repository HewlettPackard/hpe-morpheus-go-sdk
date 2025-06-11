# InstanceTypeLayoutUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Layout name | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**InstanceVersion** | Pointer to **string** | Version of the layout | [optional] 
**Description** | Pointer to **string** | Layout description | [optional] 
**SortOrder** | Pointer to **int64** | Display order of the layout, higher to lower | [optional] 
**Creatable** | Pointer to **bool** | Can be used to enable / disable the creatability of the layout. | [optional] [default to true]
**ProvisionTypeCode** | Pointer to **string** | Provision type code | [optional] 
**MemoryRequirement** | Pointer to **string** | Memory requirement in megabytes | [optional] 
**HasAutoScale** | Pointer to **bool** | Can be used to enable / disable the horizontal scaling. | [optional] [default to false]
**SupportsConvertToManaged** | Pointer to **bool** | Can be used to enable / disable the supports convert to managed. | [optional] [default to false]
**ContainerTypes** | Pointer to **[]int64** | Array of layout node type IDs | [optional] 
**OptionTypes** | Pointer to **[]int64** | Array of layout option type IDs | [optional] 
**SpecTemplates** | Pointer to **[]int64** | Array of layout spec template IDs | [optional] 
**EnvironmentVariables** | Pointer to [**[]AddClusterLayoutsRequestLayoutEnvironmentVariablesInner**](AddClusterLayoutsRequestLayoutEnvironmentVariablesInner.md) | The environmentVariables parameter is array of env objects | [optional] 
**PriceSets** | Pointer to [**[]AddInstanceTypeRequestInstanceTypePriceSetsInner**](AddInstanceTypeRequestInstanceTypePriceSetsInner.md) | Array of price set objects | [optional] 
**Permissions** | Pointer to [**AddLayoutRequestInstanceTypeLayoutPermissions**](AddLayoutRequestInstanceTypeLayoutPermissions.md) |  | [optional] 

## Methods

### NewInstanceTypeLayoutUpdate

`func NewInstanceTypeLayoutUpdate() *InstanceTypeLayoutUpdate`

NewInstanceTypeLayoutUpdate instantiates a new InstanceTypeLayoutUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceTypeLayoutUpdateWithDefaults

`func NewInstanceTypeLayoutUpdateWithDefaults() *InstanceTypeLayoutUpdate`

NewInstanceTypeLayoutUpdateWithDefaults instantiates a new InstanceTypeLayoutUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InstanceTypeLayoutUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceTypeLayoutUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceTypeLayoutUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *InstanceTypeLayoutUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetLabels

`func (o *InstanceTypeLayoutUpdate) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *InstanceTypeLayoutUpdate) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *InstanceTypeLayoutUpdate) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *InstanceTypeLayoutUpdate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetInstanceVersion

`func (o *InstanceTypeLayoutUpdate) GetInstanceVersion() string`

GetInstanceVersion returns the InstanceVersion field if non-nil, zero value otherwise.

### GetInstanceVersionOk

`func (o *InstanceTypeLayoutUpdate) GetInstanceVersionOk() (*string, bool)`

GetInstanceVersionOk returns a tuple with the InstanceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceVersion

`func (o *InstanceTypeLayoutUpdate) SetInstanceVersion(v string)`

SetInstanceVersion sets InstanceVersion field to given value.

### HasInstanceVersion

`func (o *InstanceTypeLayoutUpdate) HasInstanceVersion() bool`

HasInstanceVersion returns a boolean if a field has been set.

### GetDescription

`func (o *InstanceTypeLayoutUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InstanceTypeLayoutUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InstanceTypeLayoutUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InstanceTypeLayoutUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSortOrder

`func (o *InstanceTypeLayoutUpdate) GetSortOrder() int64`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *InstanceTypeLayoutUpdate) GetSortOrderOk() (*int64, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *InstanceTypeLayoutUpdate) SetSortOrder(v int64)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *InstanceTypeLayoutUpdate) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetCreatable

`func (o *InstanceTypeLayoutUpdate) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *InstanceTypeLayoutUpdate) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *InstanceTypeLayoutUpdate) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *InstanceTypeLayoutUpdate) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetProvisionTypeCode

`func (o *InstanceTypeLayoutUpdate) GetProvisionTypeCode() string`

GetProvisionTypeCode returns the ProvisionTypeCode field if non-nil, zero value otherwise.

### GetProvisionTypeCodeOk

`func (o *InstanceTypeLayoutUpdate) GetProvisionTypeCodeOk() (*string, bool)`

GetProvisionTypeCodeOk returns a tuple with the ProvisionTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionTypeCode

`func (o *InstanceTypeLayoutUpdate) SetProvisionTypeCode(v string)`

SetProvisionTypeCode sets ProvisionTypeCode field to given value.

### HasProvisionTypeCode

`func (o *InstanceTypeLayoutUpdate) HasProvisionTypeCode() bool`

HasProvisionTypeCode returns a boolean if a field has been set.

### GetMemoryRequirement

`func (o *InstanceTypeLayoutUpdate) GetMemoryRequirement() string`

GetMemoryRequirement returns the MemoryRequirement field if non-nil, zero value otherwise.

### GetMemoryRequirementOk

`func (o *InstanceTypeLayoutUpdate) GetMemoryRequirementOk() (*string, bool)`

GetMemoryRequirementOk returns a tuple with the MemoryRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryRequirement

`func (o *InstanceTypeLayoutUpdate) SetMemoryRequirement(v string)`

SetMemoryRequirement sets MemoryRequirement field to given value.

### HasMemoryRequirement

`func (o *InstanceTypeLayoutUpdate) HasMemoryRequirement() bool`

HasMemoryRequirement returns a boolean if a field has been set.

### GetHasAutoScale

`func (o *InstanceTypeLayoutUpdate) GetHasAutoScale() bool`

GetHasAutoScale returns the HasAutoScale field if non-nil, zero value otherwise.

### GetHasAutoScaleOk

`func (o *InstanceTypeLayoutUpdate) GetHasAutoScaleOk() (*bool, bool)`

GetHasAutoScaleOk returns a tuple with the HasAutoScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAutoScale

`func (o *InstanceTypeLayoutUpdate) SetHasAutoScale(v bool)`

SetHasAutoScale sets HasAutoScale field to given value.

### HasHasAutoScale

`func (o *InstanceTypeLayoutUpdate) HasHasAutoScale() bool`

HasHasAutoScale returns a boolean if a field has been set.

### GetSupportsConvertToManaged

`func (o *InstanceTypeLayoutUpdate) GetSupportsConvertToManaged() bool`

GetSupportsConvertToManaged returns the SupportsConvertToManaged field if non-nil, zero value otherwise.

### GetSupportsConvertToManagedOk

`func (o *InstanceTypeLayoutUpdate) GetSupportsConvertToManagedOk() (*bool, bool)`

GetSupportsConvertToManagedOk returns a tuple with the SupportsConvertToManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsConvertToManaged

`func (o *InstanceTypeLayoutUpdate) SetSupportsConvertToManaged(v bool)`

SetSupportsConvertToManaged sets SupportsConvertToManaged field to given value.

### HasSupportsConvertToManaged

`func (o *InstanceTypeLayoutUpdate) HasSupportsConvertToManaged() bool`

HasSupportsConvertToManaged returns a boolean if a field has been set.

### GetContainerTypes

`func (o *InstanceTypeLayoutUpdate) GetContainerTypes() []int64`

GetContainerTypes returns the ContainerTypes field if non-nil, zero value otherwise.

### GetContainerTypesOk

`func (o *InstanceTypeLayoutUpdate) GetContainerTypesOk() (*[]int64, bool)`

GetContainerTypesOk returns a tuple with the ContainerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTypes

`func (o *InstanceTypeLayoutUpdate) SetContainerTypes(v []int64)`

SetContainerTypes sets ContainerTypes field to given value.

### HasContainerTypes

`func (o *InstanceTypeLayoutUpdate) HasContainerTypes() bool`

HasContainerTypes returns a boolean if a field has been set.

### GetOptionTypes

`func (o *InstanceTypeLayoutUpdate) GetOptionTypes() []int64`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *InstanceTypeLayoutUpdate) GetOptionTypesOk() (*[]int64, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *InstanceTypeLayoutUpdate) SetOptionTypes(v []int64)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *InstanceTypeLayoutUpdate) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetSpecTemplates

`func (o *InstanceTypeLayoutUpdate) GetSpecTemplates() []int64`

GetSpecTemplates returns the SpecTemplates field if non-nil, zero value otherwise.

### GetSpecTemplatesOk

`func (o *InstanceTypeLayoutUpdate) GetSpecTemplatesOk() (*[]int64, bool)`

GetSpecTemplatesOk returns a tuple with the SpecTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecTemplates

`func (o *InstanceTypeLayoutUpdate) SetSpecTemplates(v []int64)`

SetSpecTemplates sets SpecTemplates field to given value.

### HasSpecTemplates

`func (o *InstanceTypeLayoutUpdate) HasSpecTemplates() bool`

HasSpecTemplates returns a boolean if a field has been set.

### GetEnvironmentVariables

`func (o *InstanceTypeLayoutUpdate) GetEnvironmentVariables() []AddClusterLayoutsRequestLayoutEnvironmentVariablesInner`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *InstanceTypeLayoutUpdate) GetEnvironmentVariablesOk() (*[]AddClusterLayoutsRequestLayoutEnvironmentVariablesInner, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *InstanceTypeLayoutUpdate) SetEnvironmentVariables(v []AddClusterLayoutsRequestLayoutEnvironmentVariablesInner)`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *InstanceTypeLayoutUpdate) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetPriceSets

`func (o *InstanceTypeLayoutUpdate) GetPriceSets() []AddInstanceTypeRequestInstanceTypePriceSetsInner`

GetPriceSets returns the PriceSets field if non-nil, zero value otherwise.

### GetPriceSetsOk

`func (o *InstanceTypeLayoutUpdate) GetPriceSetsOk() (*[]AddInstanceTypeRequestInstanceTypePriceSetsInner, bool)`

GetPriceSetsOk returns a tuple with the PriceSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSets

`func (o *InstanceTypeLayoutUpdate) SetPriceSets(v []AddInstanceTypeRequestInstanceTypePriceSetsInner)`

SetPriceSets sets PriceSets field to given value.

### HasPriceSets

`func (o *InstanceTypeLayoutUpdate) HasPriceSets() bool`

HasPriceSets returns a boolean if a field has been set.

### GetPermissions

`func (o *InstanceTypeLayoutUpdate) GetPermissions() AddLayoutRequestInstanceTypeLayoutPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *InstanceTypeLayoutUpdate) GetPermissionsOk() (*AddLayoutRequestInstanceTypeLayoutPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *InstanceTypeLayoutUpdate) SetPermissions(v AddLayoutRequestInstanceTypeLayoutPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *InstanceTypeLayoutUpdate) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


