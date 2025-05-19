# ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**SortOrder** | Pointer to **int64** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MaxStorage** | Pointer to **int64** |  | [optional] 
**MaxMemory** | Pointer to **int64** |  | [optional] 
**MaxCpu** | Pointer to **int64** |  | [optional] 
**MaxCores** | Pointer to **int64** |  | [optional] 
**MaxDisks** | Pointer to **string** |  | [optional] 
**CoresPerSocket** | Pointer to **int64** |  | [optional] 
**CustomCpu** | Pointer to **bool** |  | [optional] 
**CustomCores** | Pointer to **bool** |  | [optional] 
**CustomMaxStorage** | Pointer to **bool** |  | [optional] 
**CustomMaxDataStorage** | Pointer to **bool** |  | [optional] 
**CustomMaxMemory** | Pointer to **bool** |  | [optional] 
**AddVolumes** | Pointer to **bool** |  | [optional] 
**MemoryOptionSource** | Pointer to **string** |  | [optional] 
**CpuOptionSource** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**RegionCode** | Pointer to **string** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Editable** | Pointer to **bool** |  | [optional] 
**ProvisionType** | Pointer to [**ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionProvisionType**](ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionProvisionType.md) |  | [optional] 
**Tenants** | Pointer to **string** |  | [optional] 
**PriceSets** | Pointer to [**[]ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionPriceSetsInner**](ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionPriceSetsInner.md) |  | [optional] 
**Config** | Pointer to [**ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfig**](ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfig.md) |  | [optional] 

## Methods

### NewListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction

`func NewListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction() *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction`

NewListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction instantiates a new ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterActionWithDefaults

`func NewListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterActionWithDefaults() *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction`

NewListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterActionWithDefaults instantiates a new ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetActive

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetSortOrder

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetSortOrder() int64`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetSortOrderOk() (*int64, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) SetSortOrder(v int64)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetDescription

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMaxStorage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetMaxMemory

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxCpu

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetMaxCpu() int64`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetMaxCpuOk() (*int64, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) SetMaxCpu(v int64)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### GetMaxCores

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetMaxCores() int64`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetMaxCoresOk() (*int64, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) SetMaxCores(v int64)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetMaxDisks

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetMaxDisks() string`

GetMaxDisks returns the MaxDisks field if non-nil, zero value otherwise.

### GetMaxDisksOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetMaxDisksOk() (*string, bool)`

GetMaxDisksOk returns a tuple with the MaxDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDisks

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) SetMaxDisks(v string)`

SetMaxDisks sets MaxDisks field to given value.

### HasMaxDisks

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) HasMaxDisks() bool`

HasMaxDisks returns a boolean if a field has been set.

### GetCoresPerSocket

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetCoresPerSocket() int64`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetCoresPerSocketOk() (*int64, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) SetCoresPerSocket(v int64)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### GetCustomCpu

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetCustomCpu() bool`

GetCustomCpu returns the CustomCpu field if non-nil, zero value otherwise.

### GetCustomCpuOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetCustomCpuOk() (*bool, bool)`

GetCustomCpuOk returns a tuple with the CustomCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCpu

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) SetCustomCpu(v bool)`

SetCustomCpu sets CustomCpu field to given value.

### HasCustomCpu

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) HasCustomCpu() bool`

HasCustomCpu returns a boolean if a field has been set.

### GetCustomCores

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetCustomCores() bool`

GetCustomCores returns the CustomCores field if non-nil, zero value otherwise.

### GetCustomCoresOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetCustomCoresOk() (*bool, bool)`

GetCustomCoresOk returns a tuple with the CustomCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCores

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) SetCustomCores(v bool)`

SetCustomCores sets CustomCores field to given value.

### HasCustomCores

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) HasCustomCores() bool`

HasCustomCores returns a boolean if a field has been set.

### GetCustomMaxStorage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetCustomMaxStorage() bool`

GetCustomMaxStorage returns the CustomMaxStorage field if non-nil, zero value otherwise.

### GetCustomMaxStorageOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetCustomMaxStorageOk() (*bool, bool)`

GetCustomMaxStorageOk returns a tuple with the CustomMaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxStorage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) SetCustomMaxStorage(v bool)`

SetCustomMaxStorage sets CustomMaxStorage field to given value.

### HasCustomMaxStorage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) HasCustomMaxStorage() bool`

HasCustomMaxStorage returns a boolean if a field has been set.

### GetCustomMaxDataStorage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetCustomMaxDataStorage() bool`

GetCustomMaxDataStorage returns the CustomMaxDataStorage field if non-nil, zero value otherwise.

### GetCustomMaxDataStorageOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetCustomMaxDataStorageOk() (*bool, bool)`

GetCustomMaxDataStorageOk returns a tuple with the CustomMaxDataStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxDataStorage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) SetCustomMaxDataStorage(v bool)`

SetCustomMaxDataStorage sets CustomMaxDataStorage field to given value.

### HasCustomMaxDataStorage

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) HasCustomMaxDataStorage() bool`

HasCustomMaxDataStorage returns a boolean if a field has been set.

### GetCustomMaxMemory

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetCustomMaxMemory() bool`

GetCustomMaxMemory returns the CustomMaxMemory field if non-nil, zero value otherwise.

### GetCustomMaxMemoryOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetCustomMaxMemoryOk() (*bool, bool)`

GetCustomMaxMemoryOk returns a tuple with the CustomMaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxMemory

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) SetCustomMaxMemory(v bool)`

SetCustomMaxMemory sets CustomMaxMemory field to given value.

### HasCustomMaxMemory

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) HasCustomMaxMemory() bool`

HasCustomMaxMemory returns a boolean if a field has been set.

### GetAddVolumes

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetAddVolumes() bool`

GetAddVolumes returns the AddVolumes field if non-nil, zero value otherwise.

### GetAddVolumesOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetAddVolumesOk() (*bool, bool)`

GetAddVolumesOk returns a tuple with the AddVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddVolumes

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) SetAddVolumes(v bool)`

SetAddVolumes sets AddVolumes field to given value.

### HasAddVolumes

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) HasAddVolumes() bool`

HasAddVolumes returns a boolean if a field has been set.

### GetMemoryOptionSource

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetMemoryOptionSource() string`

GetMemoryOptionSource returns the MemoryOptionSource field if non-nil, zero value otherwise.

### GetMemoryOptionSourceOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetMemoryOptionSourceOk() (*string, bool)`

GetMemoryOptionSourceOk returns a tuple with the MemoryOptionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryOptionSource

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) SetMemoryOptionSource(v string)`

SetMemoryOptionSource sets MemoryOptionSource field to given value.

### HasMemoryOptionSource

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) HasMemoryOptionSource() bool`

HasMemoryOptionSource returns a boolean if a field has been set.

### GetCpuOptionSource

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetCpuOptionSource() string`

GetCpuOptionSource returns the CpuOptionSource field if non-nil, zero value otherwise.

### GetCpuOptionSourceOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetCpuOptionSourceOk() (*string, bool)`

GetCpuOptionSourceOk returns a tuple with the CpuOptionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuOptionSource

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) SetCpuOptionSource(v string)`

SetCpuOptionSource sets CpuOptionSource field to given value.

### HasCpuOptionSource

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) HasCpuOptionSource() bool`

HasCpuOptionSource returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetRegionCode

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetVisibility

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetEditable

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetProvisionType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetProvisionType() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionProvisionType`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetProvisionTypeOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionProvisionType, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) SetProvisionType(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionProvisionType)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetTenants

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetTenants() string`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetTenantsOk() (*string, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) SetTenants(v string)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetPriceSets

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetPriceSets() []ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionPriceSetsInner`

GetPriceSets returns the PriceSets field if non-nil, zero value otherwise.

### GetPriceSetsOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetPriceSetsOk() (*[]ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionPriceSetsInner, bool)`

GetPriceSetsOk returns a tuple with the PriceSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSets

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) SetPriceSets(v []ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionPriceSetsInner)`

SetPriceSets sets PriceSets field to given value.

### HasPriceSets

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) HasPriceSets() bool`

HasPriceSets returns a boolean if a field has been set.

### GetConfig

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetConfig() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) GetConfigOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) SetConfig(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanAfterAction) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


