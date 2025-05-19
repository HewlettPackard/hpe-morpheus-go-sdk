# ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**PriorityOrder** | Pointer to **int64** |  | [optional] 
**NodeCount** | Pointer to **int64** |  | [optional] 
**NodeType** | Pointer to **string** |  | [optional] 
**MinNodeCount** | Pointer to **int64** |  | [optional] 
**MaxNodeCount** | Pointer to **string** |  | [optional] 
**DynamicCount** | Pointer to **bool** |  | [optional] 
**InstallContainerRuntime** | Pointer to **bool** |  | [optional] 
**InstallStorageRuntime** | Pointer to **bool** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Category** | Pointer to **string** |  | [optional] 
**Config** | Pointer to **string** |  | [optional] 
**ContainerType** | Pointer to [**ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType**](ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType.md) |  | [optional] 
**ComputeServerType** | Pointer to [**ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerComputeServerType**](ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerComputeServerType.md) |  | [optional] 
**ProvisionService** | Pointer to **string** |  | [optional] 
**PlanCategory** | Pointer to **string** |  | [optional] 
**NamePrefix** | Pointer to **string** |  | [optional] 
**NameSuffix** | Pointer to **string** |  | [optional] 
**ForceNameIndex** | Pointer to **bool** |  | [optional] 
**LoadBalance** | Pointer to **bool** |  | [optional] 

## Methods

### NewListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner

`func NewListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner() *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner`

NewListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner instantiates a new ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerWithDefaults

`func NewListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerWithDefaults() *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner`

NewListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerWithDefaults instantiates a new ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPriorityOrder

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetPriorityOrder() int64`

GetPriorityOrder returns the PriorityOrder field if non-nil, zero value otherwise.

### GetPriorityOrderOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetPriorityOrderOk() (*int64, bool)`

GetPriorityOrderOk returns a tuple with the PriorityOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityOrder

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) SetPriorityOrder(v int64)`

SetPriorityOrder sets PriorityOrder field to given value.

### HasPriorityOrder

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) HasPriorityOrder() bool`

HasPriorityOrder returns a boolean if a field has been set.

### GetNodeCount

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetNodeCount() int64`

GetNodeCount returns the NodeCount field if non-nil, zero value otherwise.

### GetNodeCountOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetNodeCountOk() (*int64, bool)`

GetNodeCountOk returns a tuple with the NodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeCount

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) SetNodeCount(v int64)`

SetNodeCount sets NodeCount field to given value.

### HasNodeCount

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) HasNodeCount() bool`

HasNodeCount returns a boolean if a field has been set.

### GetNodeType

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.

### HasNodeType

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) HasNodeType() bool`

HasNodeType returns a boolean if a field has been set.

### GetMinNodeCount

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetMinNodeCount() int64`

GetMinNodeCount returns the MinNodeCount field if non-nil, zero value otherwise.

### GetMinNodeCountOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetMinNodeCountOk() (*int64, bool)`

GetMinNodeCountOk returns a tuple with the MinNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinNodeCount

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) SetMinNodeCount(v int64)`

SetMinNodeCount sets MinNodeCount field to given value.

### HasMinNodeCount

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) HasMinNodeCount() bool`

HasMinNodeCount returns a boolean if a field has been set.

### GetMaxNodeCount

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetMaxNodeCount() string`

GetMaxNodeCount returns the MaxNodeCount field if non-nil, zero value otherwise.

### GetMaxNodeCountOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetMaxNodeCountOk() (*string, bool)`

GetMaxNodeCountOk returns a tuple with the MaxNodeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxNodeCount

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) SetMaxNodeCount(v string)`

SetMaxNodeCount sets MaxNodeCount field to given value.

### HasMaxNodeCount

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) HasMaxNodeCount() bool`

HasMaxNodeCount returns a boolean if a field has been set.

### GetDynamicCount

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetDynamicCount() bool`

GetDynamicCount returns the DynamicCount field if non-nil, zero value otherwise.

### GetDynamicCountOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetDynamicCountOk() (*bool, bool)`

GetDynamicCountOk returns a tuple with the DynamicCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamicCount

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) SetDynamicCount(v bool)`

SetDynamicCount sets DynamicCount field to given value.

### HasDynamicCount

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) HasDynamicCount() bool`

HasDynamicCount returns a boolean if a field has been set.

### GetInstallContainerRuntime

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetInstallContainerRuntime() bool`

GetInstallContainerRuntime returns the InstallContainerRuntime field if non-nil, zero value otherwise.

### GetInstallContainerRuntimeOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetInstallContainerRuntimeOk() (*bool, bool)`

GetInstallContainerRuntimeOk returns a tuple with the InstallContainerRuntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallContainerRuntime

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) SetInstallContainerRuntime(v bool)`

SetInstallContainerRuntime sets InstallContainerRuntime field to given value.

### HasInstallContainerRuntime

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) HasInstallContainerRuntime() bool`

HasInstallContainerRuntime returns a boolean if a field has been set.

### GetInstallStorageRuntime

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetInstallStorageRuntime() bool`

GetInstallStorageRuntime returns the InstallStorageRuntime field if non-nil, zero value otherwise.

### GetInstallStorageRuntimeOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetInstallStorageRuntimeOk() (*bool, bool)`

GetInstallStorageRuntimeOk returns a tuple with the InstallStorageRuntime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstallStorageRuntime

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) SetInstallStorageRuntime(v bool)`

SetInstallStorageRuntime sets InstallStorageRuntime field to given value.

### HasInstallStorageRuntime

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) HasInstallStorageRuntime() bool`

HasInstallStorageRuntime returns a boolean if a field has been set.

### GetName

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetCategory

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetConfig

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetConfig() string`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetConfigOk() (*string, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) SetConfig(v string)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetContainerType

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetContainerType() ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType`

GetContainerType returns the ContainerType field if non-nil, zero value otherwise.

### GetContainerTypeOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetContainerTypeOk() (*ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType, bool)`

GetContainerTypeOk returns a tuple with the ContainerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContainerType

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) SetContainerType(v ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerContainerType)`

SetContainerType sets ContainerType field to given value.

### HasContainerType

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) HasContainerType() bool`

HasContainerType returns a boolean if a field has been set.

### GetComputeServerType

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetComputeServerType() ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerComputeServerType`

GetComputeServerType returns the ComputeServerType field if non-nil, zero value otherwise.

### GetComputeServerTypeOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetComputeServerTypeOk() (*ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerComputeServerType, bool)`

GetComputeServerTypeOk returns a tuple with the ComputeServerType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComputeServerType

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) SetComputeServerType(v ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInnerComputeServerType)`

SetComputeServerType sets ComputeServerType field to given value.

### HasComputeServerType

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) HasComputeServerType() bool`

HasComputeServerType returns a boolean if a field has been set.

### GetProvisionService

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetProvisionService() string`

GetProvisionService returns the ProvisionService field if non-nil, zero value otherwise.

### GetProvisionServiceOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetProvisionServiceOk() (*string, bool)`

GetProvisionServiceOk returns a tuple with the ProvisionService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionService

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) SetProvisionService(v string)`

SetProvisionService sets ProvisionService field to given value.

### HasProvisionService

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) HasProvisionService() bool`

HasProvisionService returns a boolean if a field has been set.

### GetPlanCategory

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetPlanCategory() string`

GetPlanCategory returns the PlanCategory field if non-nil, zero value otherwise.

### GetPlanCategoryOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetPlanCategoryOk() (*string, bool)`

GetPlanCategoryOk returns a tuple with the PlanCategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanCategory

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) SetPlanCategory(v string)`

SetPlanCategory sets PlanCategory field to given value.

### HasPlanCategory

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) HasPlanCategory() bool`

HasPlanCategory returns a boolean if a field has been set.

### GetNamePrefix

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetNamePrefix() string`

GetNamePrefix returns the NamePrefix field if non-nil, zero value otherwise.

### GetNamePrefixOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetNamePrefixOk() (*string, bool)`

GetNamePrefixOk returns a tuple with the NamePrefix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamePrefix

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) SetNamePrefix(v string)`

SetNamePrefix sets NamePrefix field to given value.

### HasNamePrefix

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) HasNamePrefix() bool`

HasNamePrefix returns a boolean if a field has been set.

### GetNameSuffix

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetNameSuffix() string`

GetNameSuffix returns the NameSuffix field if non-nil, zero value otherwise.

### GetNameSuffixOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetNameSuffixOk() (*string, bool)`

GetNameSuffixOk returns a tuple with the NameSuffix field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNameSuffix

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) SetNameSuffix(v string)`

SetNameSuffix sets NameSuffix field to given value.

### HasNameSuffix

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) HasNameSuffix() bool`

HasNameSuffix returns a boolean if a field has been set.

### GetForceNameIndex

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetForceNameIndex() bool`

GetForceNameIndex returns the ForceNameIndex field if non-nil, zero value otherwise.

### GetForceNameIndexOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetForceNameIndexOk() (*bool, bool)`

GetForceNameIndexOk returns a tuple with the ForceNameIndex field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForceNameIndex

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) SetForceNameIndex(v bool)`

SetForceNameIndex sets ForceNameIndex field to given value.

### HasForceNameIndex

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) HasForceNameIndex() bool`

HasForceNameIndex returns a boolean if a field has been set.

### GetLoadBalance

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetLoadBalance() bool`

GetLoadBalance returns the LoadBalance field if non-nil, zero value otherwise.

### GetLoadBalanceOk

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) GetLoadBalanceOk() (*bool, bool)`

GetLoadBalanceOk returns a tuple with the LoadBalance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoadBalance

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) SetLoadBalance(v bool)`

SetLoadBalance sets LoadBalance field to given value.

### HasLoadBalance

`func (o *ListClusterLayouts200ResponseAllOfLayoutsInnerComputeServersInner) HasLoadBalance() bool`

HasLoadBalance returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


