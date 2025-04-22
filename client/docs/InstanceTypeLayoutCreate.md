# InstanceTypeLayoutCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Layout name | 
**Labels** | Pointer to **[]string** |  | [optional] 
**InstanceVersion** | **string** | Version of the layout | 
**Description** | Pointer to **string** | Layout description | [optional] 
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

### NewInstanceTypeLayoutCreate

`func NewInstanceTypeLayoutCreate(name string, instanceVersion string, provisionTypeCode string, ) *InstanceTypeLayoutCreate`

NewInstanceTypeLayoutCreate instantiates a new InstanceTypeLayoutCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceTypeLayoutCreateWithDefaults

`func NewInstanceTypeLayoutCreateWithDefaults() *InstanceTypeLayoutCreate`

NewInstanceTypeLayoutCreateWithDefaults instantiates a new InstanceTypeLayoutCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *InstanceTypeLayoutCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *InstanceTypeLayoutCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *InstanceTypeLayoutCreate) SetName(v string)`

SetName sets Name field to given value.


### GetLabels

`func (o *InstanceTypeLayoutCreate) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *InstanceTypeLayoutCreate) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *InstanceTypeLayoutCreate) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *InstanceTypeLayoutCreate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetInstanceVersion

`func (o *InstanceTypeLayoutCreate) GetInstanceVersion() string`

GetInstanceVersion returns the InstanceVersion field if non-nil, zero value otherwise.

### GetInstanceVersionOk

`func (o *InstanceTypeLayoutCreate) GetInstanceVersionOk() (*string, bool)`

GetInstanceVersionOk returns a tuple with the InstanceVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceVersion

`func (o *InstanceTypeLayoutCreate) SetInstanceVersion(v string)`

SetInstanceVersion sets InstanceVersion field to given value.


### GetDescription

`func (o *InstanceTypeLayoutCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *InstanceTypeLayoutCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *InstanceTypeLayoutCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *InstanceTypeLayoutCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatable

`func (o *InstanceTypeLayoutCreate) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *InstanceTypeLayoutCreate) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *InstanceTypeLayoutCreate) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *InstanceTypeLayoutCreate) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetProvisionTypeCode

`func (o *InstanceTypeLayoutCreate) GetProvisionTypeCode() string`

GetProvisionTypeCode returns the ProvisionTypeCode field if non-nil, zero value otherwise.

### GetProvisionTypeCodeOk

`func (o *InstanceTypeLayoutCreate) GetProvisionTypeCodeOk() (*string, bool)`

GetProvisionTypeCodeOk returns a tuple with the ProvisionTypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionTypeCode

`func (o *InstanceTypeLayoutCreate) SetProvisionTypeCode(v string)`

SetProvisionTypeCode sets ProvisionTypeCode field to given value.


### GetMemoryRequirement

`func (o *InstanceTypeLayoutCreate) GetMemoryRequirement() string`

GetMemoryRequirement returns the MemoryRequirement field if non-nil, zero value otherwise.

### GetMemoryRequirementOk

`func (o *InstanceTypeLayoutCreate) GetMemoryRequirementOk() (*string, bool)`

GetMemoryRequirementOk returns a tuple with the MemoryRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryRequirement

`func (o *InstanceTypeLayoutCreate) SetMemoryRequirement(v string)`

SetMemoryRequirement sets MemoryRequirement field to given value.

### HasMemoryRequirement

`func (o *InstanceTypeLayoutCreate) HasMemoryRequirement() bool`

HasMemoryRequirement returns a boolean if a field has been set.

### GetHasAutoScale

`func (o *InstanceTypeLayoutCreate) GetHasAutoScale() bool`

GetHasAutoScale returns the HasAutoScale field if non-nil, zero value otherwise.

### GetHasAutoScaleOk

`func (o *InstanceTypeLayoutCreate) GetHasAutoScaleOk() (*bool, bool)`

GetHasAutoScaleOk returns a tuple with the HasAutoScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAutoScale

`func (o *InstanceTypeLayoutCreate) SetHasAutoScale(v bool)`

SetHasAutoScale sets HasAutoScale field to given value.

### HasHasAutoScale

`func (o *InstanceTypeLayoutCreate) HasHasAutoScale() bool`

HasHasAutoScale returns a boolean if a field has been set.

### GetSupportsConvertToManaged

`func (o *InstanceTypeLayoutCreate) GetSupportsConvertToManaged() bool`

GetSupportsConvertToManaged returns the SupportsConvertToManaged field if non-nil, zero value otherwise.

### GetSupportsConvertToManagedOk

`func (o *InstanceTypeLayoutCreate) GetSupportsConvertToManagedOk() (*bool, bool)`

GetSupportsConvertToManagedOk returns a tuple with the SupportsConvertToManaged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportsConvertToManaged

`func (o *InstanceTypeLayoutCreate) SetSupportsConvertToManaged(v bool)`

SetSupportsConvertToManaged sets SupportsConvertToManaged field to given value.

### HasSupportsConvertToManaged

`func (o *InstanceTypeLayoutCreate) HasSupportsConvertToManaged() bool`

HasSupportsConvertToManaged returns a boolean if a field has been set.

### GetContainerTypes

`func (o *InstanceTypeLayoutCreate) GetContainerTypes() []int64`

GetContainerTypes returns the ContainerTypes field if non-nil, zero value otherwise.

### GetContainerTypesOk

`func (o *InstanceTypeLayoutCreate) GetContainerTypesOk() (*[]int64, bool)`

GetContainerTypesOk returns a tuple with the ContainerTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerTypes

`func (o *InstanceTypeLayoutCreate) SetContainerTypes(v []int64)`

SetContainerTypes sets ContainerTypes field to given value.

### HasContainerTypes

`func (o *InstanceTypeLayoutCreate) HasContainerTypes() bool`

HasContainerTypes returns a boolean if a field has been set.

### GetOptionTypes

`func (o *InstanceTypeLayoutCreate) GetOptionTypes() []int64`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *InstanceTypeLayoutCreate) GetOptionTypesOk() (*[]int64, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *InstanceTypeLayoutCreate) SetOptionTypes(v []int64)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *InstanceTypeLayoutCreate) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetSpecTemplates

`func (o *InstanceTypeLayoutCreate) GetSpecTemplates() []int64`

GetSpecTemplates returns the SpecTemplates field if non-nil, zero value otherwise.

### GetSpecTemplatesOk

`func (o *InstanceTypeLayoutCreate) GetSpecTemplatesOk() (*[]int64, bool)`

GetSpecTemplatesOk returns a tuple with the SpecTemplates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSpecTemplates

`func (o *InstanceTypeLayoutCreate) SetSpecTemplates(v []int64)`

SetSpecTemplates sets SpecTemplates field to given value.

### HasSpecTemplates

`func (o *InstanceTypeLayoutCreate) HasSpecTemplates() bool`

HasSpecTemplates returns a boolean if a field has been set.

### GetEnvironmentVariables

`func (o *InstanceTypeLayoutCreate) GetEnvironmentVariables() []AddClusterLayoutsRequestLayoutEnvironmentVariablesInner`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *InstanceTypeLayoutCreate) GetEnvironmentVariablesOk() (*[]AddClusterLayoutsRequestLayoutEnvironmentVariablesInner, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *InstanceTypeLayoutCreate) SetEnvironmentVariables(v []AddClusterLayoutsRequestLayoutEnvironmentVariablesInner)`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *InstanceTypeLayoutCreate) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetPriceSets

`func (o *InstanceTypeLayoutCreate) GetPriceSets() []AddInstanceTypeRequestInstanceTypePriceSetsInner`

GetPriceSets returns the PriceSets field if non-nil, zero value otherwise.

### GetPriceSetsOk

`func (o *InstanceTypeLayoutCreate) GetPriceSetsOk() (*[]AddInstanceTypeRequestInstanceTypePriceSetsInner, bool)`

GetPriceSetsOk returns a tuple with the PriceSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSets

`func (o *InstanceTypeLayoutCreate) SetPriceSets(v []AddInstanceTypeRequestInstanceTypePriceSetsInner)`

SetPriceSets sets PriceSets field to given value.

### HasPriceSets

`func (o *InstanceTypeLayoutCreate) HasPriceSets() bool`

HasPriceSets returns a boolean if a field has been set.

### GetPermissions

`func (o *InstanceTypeLayoutCreate) GetPermissions() AddLayoutRequestInstanceTypeLayoutPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *InstanceTypeLayoutCreate) GetPermissionsOk() (*AddLayoutRequestInstanceTypeLayoutPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *InstanceTypeLayoutCreate) SetPermissions(v AddLayoutRequestInstanceTypeLayoutPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *InstanceTypeLayoutCreate) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


