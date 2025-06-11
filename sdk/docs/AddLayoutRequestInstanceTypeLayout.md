# AddLayoutRequestInstanceTypeLayout

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Layout name | 
**Labels** | Pointer to **[]string** |  | [optional] 
**InstanceVersion** | **string** | Version of the layout | 
**Description** | Pointer to **string** | Layout description | [optional] 
**SortOrder** | Pointer to **int64** | Display order of the layout, higher to lower | [optional] 
**Creatable** | Pointer to **bool** | Can be used to enable / disable the creatability of the layout. | [optional] [default to true]
**ProvisionTypeCode** | **string** | Provision type code | 
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

### NewAddLayoutRequestInstanceTypeLayout

`func NewAddLayoutRequestInstanceTypeLayout(name string, instanceVersion string, provisionTypeCode string, ) *AddLayoutRequestInstanceTypeLayout`

NewAddLayoutRequestInstanceTypeLayout instantiates a new AddLayoutRequestInstanceTypeLayout object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddLayoutRequestInstanceTypeLayoutWithDefaults

`func NewAddLayoutRequestInstanceTypeLayoutWithDefaults() *AddLayoutRequestInstanceTypeLayout`

NewAddLayoutRequestInstanceTypeLayoutWithDefaults instantiates a new AddLayoutRequestInstanceTypeLayout object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddLayoutRequestInstanceTypeLayout) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddLayoutRequestInstanceTypeLayout) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddLayoutRequestInstanceTypeLayout) SetName(v string)`

SetName sets Name field to given value.


### GetLabels

`func (o *AddLayoutRequestInstanceTypeLayout) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddLayoutRequestInstanceTypeLayout) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddLayoutRequestInstanceTypeLayout) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddLayoutRequestInstanceTypeLayout) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetInstanceVersion

`func (o *AddLayoutRequestInstanceTypeLayout) GetInstanceVersion() string`

GetInstanceVersion returns the InstanceVersion field if non-nil, zero value otherwise.

### GetInstanceVersionOk

`func (o *AddLayoutRequestInstanceTypeLayout) GetInstanceVersionOk() (*string, bool)`

GetInstanceVersionOk returns a tuple with the InstanceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceVersion

`func (o *AddLayoutRequestInstanceTypeLayout) SetInstanceVersion(v string)`

SetInstanceVersion sets InstanceVersion field to given value.


### GetDescription

`func (o *AddLayoutRequestInstanceTypeLayout) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddLayoutRequestInstanceTypeLayout) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddLayoutRequestInstanceTypeLayout) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddLayoutRequestInstanceTypeLayout) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetSortOrder

`func (o *AddLayoutRequestInstanceTypeLayout) GetSortOrder() int64`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *AddLayoutRequestInstanceTypeLayout) GetSortOrderOk() (*int64, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *AddLayoutRequestInstanceTypeLayout) SetSortOrder(v int64)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *AddLayoutRequestInstanceTypeLayout) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetCreatable

`func (o *AddLayoutRequestInstanceTypeLayout) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *AddLayoutRequestInstanceTypeLayout) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *AddLayoutRequestInstanceTypeLayout) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *AddLayoutRequestInstanceTypeLayout) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetProvisionTypeCode

`func (o *AddLayoutRequestInstanceTypeLayout) GetProvisionTypeCode() string`

GetProvisionTypeCode returns the ProvisionTypeCode field if non-nil, zero value otherwise.

### GetProvisionTypeCodeOk

`func (o *AddLayoutRequestInstanceTypeLayout) GetProvisionTypeCodeOk() (*string, bool)`

GetProvisionTypeCodeOk returns a tuple with the ProvisionTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionTypeCode

`func (o *AddLayoutRequestInstanceTypeLayout) SetProvisionTypeCode(v string)`

SetProvisionTypeCode sets ProvisionTypeCode field to given value.


### GetMemoryRequirement

`func (o *AddLayoutRequestInstanceTypeLayout) GetMemoryRequirement() string`

GetMemoryRequirement returns the MemoryRequirement field if non-nil, zero value otherwise.

### GetMemoryRequirementOk

`func (o *AddLayoutRequestInstanceTypeLayout) GetMemoryRequirementOk() (*string, bool)`

GetMemoryRequirementOk returns a tuple with the MemoryRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryRequirement

`func (o *AddLayoutRequestInstanceTypeLayout) SetMemoryRequirement(v string)`

SetMemoryRequirement sets MemoryRequirement field to given value.

### HasMemoryRequirement

`func (o *AddLayoutRequestInstanceTypeLayout) HasMemoryRequirement() bool`

HasMemoryRequirement returns a boolean if a field has been set.

### GetHasAutoScale

`func (o *AddLayoutRequestInstanceTypeLayout) GetHasAutoScale() bool`

GetHasAutoScale returns the HasAutoScale field if non-nil, zero value otherwise.

### GetHasAutoScaleOk

`func (o *AddLayoutRequestInstanceTypeLayout) GetHasAutoScaleOk() (*bool, bool)`

GetHasAutoScaleOk returns a tuple with the HasAutoScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAutoScale

`func (o *AddLayoutRequestInstanceTypeLayout) SetHasAutoScale(v bool)`

SetHasAutoScale sets HasAutoScale field to given value.

### HasHasAutoScale

`func (o *AddLayoutRequestInstanceTypeLayout) HasHasAutoScale() bool`

HasHasAutoScale returns a boolean if a field has been set.

### GetSupportsConvertToManaged

`func (o *AddLayoutRequestInstanceTypeLayout) GetSupportsConvertToManaged() bool`

GetSupportsConvertToManaged returns the SupportsConvertToManaged field if non-nil, zero value otherwise.

### GetSupportsConvertToManagedOk

`func (o *AddLayoutRequestInstanceTypeLayout) GetSupportsConvertToManagedOk() (*bool, bool)`

GetSupportsConvertToManagedOk returns a tuple with the SupportsConvertToManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsConvertToManaged

`func (o *AddLayoutRequestInstanceTypeLayout) SetSupportsConvertToManaged(v bool)`

SetSupportsConvertToManaged sets SupportsConvertToManaged field to given value.

### HasSupportsConvertToManaged

`func (o *AddLayoutRequestInstanceTypeLayout) HasSupportsConvertToManaged() bool`

HasSupportsConvertToManaged returns a boolean if a field has been set.

### GetContainerTypes

`func (o *AddLayoutRequestInstanceTypeLayout) GetContainerTypes() []int64`

GetContainerTypes returns the ContainerTypes field if non-nil, zero value otherwise.

### GetContainerTypesOk

`func (o *AddLayoutRequestInstanceTypeLayout) GetContainerTypesOk() (*[]int64, bool)`

GetContainerTypesOk returns a tuple with the ContainerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTypes

`func (o *AddLayoutRequestInstanceTypeLayout) SetContainerTypes(v []int64)`

SetContainerTypes sets ContainerTypes field to given value.

### HasContainerTypes

`func (o *AddLayoutRequestInstanceTypeLayout) HasContainerTypes() bool`

HasContainerTypes returns a boolean if a field has been set.

### GetOptionTypes

`func (o *AddLayoutRequestInstanceTypeLayout) GetOptionTypes() []int64`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *AddLayoutRequestInstanceTypeLayout) GetOptionTypesOk() (*[]int64, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *AddLayoutRequestInstanceTypeLayout) SetOptionTypes(v []int64)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *AddLayoutRequestInstanceTypeLayout) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetSpecTemplates

`func (o *AddLayoutRequestInstanceTypeLayout) GetSpecTemplates() []int64`

GetSpecTemplates returns the SpecTemplates field if non-nil, zero value otherwise.

### GetSpecTemplatesOk

`func (o *AddLayoutRequestInstanceTypeLayout) GetSpecTemplatesOk() (*[]int64, bool)`

GetSpecTemplatesOk returns a tuple with the SpecTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecTemplates

`func (o *AddLayoutRequestInstanceTypeLayout) SetSpecTemplates(v []int64)`

SetSpecTemplates sets SpecTemplates field to given value.

### HasSpecTemplates

`func (o *AddLayoutRequestInstanceTypeLayout) HasSpecTemplates() bool`

HasSpecTemplates returns a boolean if a field has been set.

### GetEnvironmentVariables

`func (o *AddLayoutRequestInstanceTypeLayout) GetEnvironmentVariables() []AddClusterLayoutsRequestLayoutEnvironmentVariablesInner`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *AddLayoutRequestInstanceTypeLayout) GetEnvironmentVariablesOk() (*[]AddClusterLayoutsRequestLayoutEnvironmentVariablesInner, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *AddLayoutRequestInstanceTypeLayout) SetEnvironmentVariables(v []AddClusterLayoutsRequestLayoutEnvironmentVariablesInner)`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *AddLayoutRequestInstanceTypeLayout) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetPriceSets

`func (o *AddLayoutRequestInstanceTypeLayout) GetPriceSets() []AddInstanceTypeRequestInstanceTypePriceSetsInner`

GetPriceSets returns the PriceSets field if non-nil, zero value otherwise.

### GetPriceSetsOk

`func (o *AddLayoutRequestInstanceTypeLayout) GetPriceSetsOk() (*[]AddInstanceTypeRequestInstanceTypePriceSetsInner, bool)`

GetPriceSetsOk returns a tuple with the PriceSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSets

`func (o *AddLayoutRequestInstanceTypeLayout) SetPriceSets(v []AddInstanceTypeRequestInstanceTypePriceSetsInner)`

SetPriceSets sets PriceSets field to given value.

### HasPriceSets

`func (o *AddLayoutRequestInstanceTypeLayout) HasPriceSets() bool`

HasPriceSets returns a boolean if a field has been set.

### GetPermissions

`func (o *AddLayoutRequestInstanceTypeLayout) GetPermissions() AddLayoutRequestInstanceTypeLayoutPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AddLayoutRequestInstanceTypeLayout) GetPermissionsOk() (*AddLayoutRequestInstanceTypeLayoutPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AddLayoutRequestInstanceTypeLayout) SetPermissions(v AddLayoutRequestInstanceTypeLayoutPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *AddLayoutRequestInstanceTypeLayout) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


