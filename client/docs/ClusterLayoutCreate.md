# ClusterLayoutCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Cluster layout name | 
**Description** | Pointer to **string** | Cluster layout description | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**ComputeVersion** | **string** | Version of the cluster layout | 
**Creatable** | Pointer to **bool** | Can be used to enable / disable the creatability of the cluster layout. | [optional] [default to true]
**HasAutoScale** | Pointer to **bool** | Can be used to enable / disable the horizontal scaling. | [optional] [default to false]
**InstallContainerRuntime** | Pointer to **bool** | Install Docker (container runtime). | [optional] [default to false]
**MemoryRequirement** | Pointer to **int64** | Memory requirement in bytes | [optional] 
**GroupType** | [**AddClusterLayoutsRequestLayoutGroupType**](AddClusterLayoutsRequestLayoutGroupType.md) |  | 
**ProvisionType** | [**AddClusterLayoutsRequestLayoutProvisionType**](AddClusterLayoutsRequestLayoutProvisionType.md) |  | 
**OptionTypes** | Pointer to [**[]GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) | Array of cluster layout option types | [optional] 
**TaskSets** | Pointer to [**[]GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) | Array of cluster layout task sets | [optional] 
**EnvironmentVariables** | Pointer to [**[]AddClusterLayoutsRequestLayoutEnvironmentVariablesInner**](AddClusterLayoutsRequestLayoutEnvironmentVariablesInner.md) | Array of cluster layout env variables | [optional] 
**Masters** | Pointer to [**[]AddClusterLayoutsRequestLayoutMastersInner**](AddClusterLayoutsRequestLayoutMastersInner.md) | Array of cluster layout master nodes | [optional] 
**Workers** | Pointer to [**[]AddClusterLayoutsRequestLayoutMastersInner**](AddClusterLayoutsRequestLayoutMastersInner.md) | Array of cluster layout worker nodes | [optional] 

## Methods

### NewClusterLayoutCreate

`func NewClusterLayoutCreate(name string, computeVersion string, groupType AddClusterLayoutsRequestLayoutGroupType, provisionType AddClusterLayoutsRequestLayoutProvisionType, ) *ClusterLayoutCreate`

NewClusterLayoutCreate instantiates a new ClusterLayoutCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClusterLayoutCreateWithDefaults

`func NewClusterLayoutCreateWithDefaults() *ClusterLayoutCreate`

NewClusterLayoutCreateWithDefaults instantiates a new ClusterLayoutCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *ClusterLayoutCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ClusterLayoutCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ClusterLayoutCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *ClusterLayoutCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ClusterLayoutCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ClusterLayoutCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ClusterLayoutCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabels

`func (o *ClusterLayoutCreate) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *ClusterLayoutCreate) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *ClusterLayoutCreate) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *ClusterLayoutCreate) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetComputeVersion

`func (o *ClusterLayoutCreate) GetComputeVersion() string`

GetComputeVersion returns the ComputeVersion field if non-nil, zero value otherwise.

### GetComputeVersionOk

`func (o *ClusterLayoutCreate) GetComputeVersionOk() (*string, bool)`

GetComputeVersionOk returns a tuple with the ComputeVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeVersion

`func (o *ClusterLayoutCreate) SetComputeVersion(v string)`

SetComputeVersion sets ComputeVersion field to given value.


### GetCreatable

`func (o *ClusterLayoutCreate) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *ClusterLayoutCreate) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *ClusterLayoutCreate) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *ClusterLayoutCreate) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetHasAutoScale

`func (o *ClusterLayoutCreate) GetHasAutoScale() bool`

GetHasAutoScale returns the HasAutoScale field if non-nil, zero value otherwise.

### GetHasAutoScaleOk

`func (o *ClusterLayoutCreate) GetHasAutoScaleOk() (*bool, bool)`

GetHasAutoScaleOk returns a tuple with the HasAutoScale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasAutoScale

`func (o *ClusterLayoutCreate) SetHasAutoScale(v bool)`

SetHasAutoScale sets HasAutoScale field to given value.

### HasHasAutoScale

`func (o *ClusterLayoutCreate) HasHasAutoScale() bool`

HasHasAutoScale returns a boolean if a field has been set.

### GetInstallContainerRuntime

`func (o *ClusterLayoutCreate) GetInstallContainerRuntime() bool`

GetInstallContainerRuntime returns the InstallContainerRuntime field if non-nil, zero value otherwise.

### GetInstallContainerRuntimeOk

`func (o *ClusterLayoutCreate) GetInstallContainerRuntimeOk() (*bool, bool)`

GetInstallContainerRuntimeOk returns a tuple with the InstallContainerRuntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallContainerRuntime

`func (o *ClusterLayoutCreate) SetInstallContainerRuntime(v bool)`

SetInstallContainerRuntime sets InstallContainerRuntime field to given value.

### HasInstallContainerRuntime

`func (o *ClusterLayoutCreate) HasInstallContainerRuntime() bool`

HasInstallContainerRuntime returns a boolean if a field has been set.

### GetMemoryRequirement

`func (o *ClusterLayoutCreate) GetMemoryRequirement() int64`

GetMemoryRequirement returns the MemoryRequirement field if non-nil, zero value otherwise.

### GetMemoryRequirementOk

`func (o *ClusterLayoutCreate) GetMemoryRequirementOk() (*int64, bool)`

GetMemoryRequirementOk returns a tuple with the MemoryRequirement field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryRequirement

`func (o *ClusterLayoutCreate) SetMemoryRequirement(v int64)`

SetMemoryRequirement sets MemoryRequirement field to given value.

### HasMemoryRequirement

`func (o *ClusterLayoutCreate) HasMemoryRequirement() bool`

HasMemoryRequirement returns a boolean if a field has been set.

### GetGroupType

`func (o *ClusterLayoutCreate) GetGroupType() AddClusterLayoutsRequestLayoutGroupType`

GetGroupType returns the GroupType field if non-nil, zero value otherwise.

### GetGroupTypeOk

`func (o *ClusterLayoutCreate) GetGroupTypeOk() (*AddClusterLayoutsRequestLayoutGroupType, bool)`

GetGroupTypeOk returns a tuple with the GroupType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupType

`func (o *ClusterLayoutCreate) SetGroupType(v AddClusterLayoutsRequestLayoutGroupType)`

SetGroupType sets GroupType field to given value.


### GetProvisionType

`func (o *ClusterLayoutCreate) GetProvisionType() AddClusterLayoutsRequestLayoutProvisionType`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *ClusterLayoutCreate) GetProvisionTypeOk() (*AddClusterLayoutsRequestLayoutProvisionType, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *ClusterLayoutCreate) SetProvisionType(v AddClusterLayoutsRequestLayoutProvisionType)`

SetProvisionType sets ProvisionType field to given value.


### GetOptionTypes

`func (o *ClusterLayoutCreate) GetOptionTypes() []GetAlerts200ResponseAllOfChecksInnerAccount`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *ClusterLayoutCreate) GetOptionTypesOk() (*[]GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *ClusterLayoutCreate) SetOptionTypes(v []GetAlerts200ResponseAllOfChecksInnerAccount)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *ClusterLayoutCreate) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetTaskSets

`func (o *ClusterLayoutCreate) GetTaskSets() []GetAlerts200ResponseAllOfChecksInnerAccount`

GetTaskSets returns the TaskSets field if non-nil, zero value otherwise.

### GetTaskSetsOk

`func (o *ClusterLayoutCreate) GetTaskSetsOk() (*[]GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetTaskSetsOk returns a tuple with the TaskSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaskSets

`func (o *ClusterLayoutCreate) SetTaskSets(v []GetAlerts200ResponseAllOfChecksInnerAccount)`

SetTaskSets sets TaskSets field to given value.

### HasTaskSets

`func (o *ClusterLayoutCreate) HasTaskSets() bool`

HasTaskSets returns a boolean if a field has been set.

### GetEnvironmentVariables

`func (o *ClusterLayoutCreate) GetEnvironmentVariables() []AddClusterLayoutsRequestLayoutEnvironmentVariablesInner`

GetEnvironmentVariables returns the EnvironmentVariables field if non-nil, zero value otherwise.

### GetEnvironmentVariablesOk

`func (o *ClusterLayoutCreate) GetEnvironmentVariablesOk() (*[]AddClusterLayoutsRequestLayoutEnvironmentVariablesInner, bool)`

GetEnvironmentVariablesOk returns a tuple with the EnvironmentVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvironmentVariables

`func (o *ClusterLayoutCreate) SetEnvironmentVariables(v []AddClusterLayoutsRequestLayoutEnvironmentVariablesInner)`

SetEnvironmentVariables sets EnvironmentVariables field to given value.

### HasEnvironmentVariables

`func (o *ClusterLayoutCreate) HasEnvironmentVariables() bool`

HasEnvironmentVariables returns a boolean if a field has been set.

### GetMasters

`func (o *ClusterLayoutCreate) GetMasters() []AddClusterLayoutsRequestLayoutMastersInner`

GetMasters returns the Masters field if non-nil, zero value otherwise.

### GetMastersOk

`func (o *ClusterLayoutCreate) GetMastersOk() (*[]AddClusterLayoutsRequestLayoutMastersInner, bool)`

GetMastersOk returns a tuple with the Masters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasters

`func (o *ClusterLayoutCreate) SetMasters(v []AddClusterLayoutsRequestLayoutMastersInner)`

SetMasters sets Masters field to given value.

### HasMasters

`func (o *ClusterLayoutCreate) HasMasters() bool`

HasMasters returns a boolean if a field has been set.

### GetWorkers

`func (o *ClusterLayoutCreate) GetWorkers() []AddClusterLayoutsRequestLayoutMastersInner`

GetWorkers returns the Workers field if non-nil, zero value otherwise.

### GetWorkersOk

`func (o *ClusterLayoutCreate) GetWorkersOk() (*[]AddClusterLayoutsRequestLayoutMastersInner, bool)`

GetWorkersOk returns a tuple with the Workers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkers

`func (o *ClusterLayoutCreate) SetWorkers(v []AddClusterLayoutsRequestLayoutMastersInner)`

SetWorkers sets Workers field to given value.

### HasWorkers

`func (o *ClusterLayoutCreate) HasWorkers() bool`

HasWorkers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


