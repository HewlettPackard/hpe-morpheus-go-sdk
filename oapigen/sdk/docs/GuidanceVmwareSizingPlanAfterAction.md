# GuidanceVmwareSizingPlanAfterAction

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
**MaxDisks** | Pointer to **NullableString** |  | [optional] 
**CoresPerSocket** | Pointer to **int64** |  | [optional] 
**CustomCpu** | Pointer to **bool** |  | [optional] 
**CustomCores** | Pointer to **bool** |  | [optional] 
**CustomMaxStorage** | Pointer to **bool** |  | [optional] 
**CustomMaxDataStorage** | Pointer to **bool** |  | [optional] 
**CustomMaxMemory** | Pointer to **bool** |  | [optional] 
**AddVolumes** | Pointer to **bool** |  | [optional] 
**MemoryOptionSource** | Pointer to **NullableString** |  | [optional] 
**CpuOptionSource** | Pointer to **NullableString** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**RegionCode** | Pointer to **NullableString** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Editable** | Pointer to **bool** |  | [optional] 
**ProvisionType** | Pointer to [**GuidanceVmwareSizingPlanAfterActionProvisionType**](GuidanceVmwareSizingPlanAfterActionProvisionType.md) |  | [optional] 
**Tenants** | Pointer to **string** |  | [optional] 
**PriceSets** | Pointer to [**[]GuidanceVmwareSizingPlanAfterActionPriceSetsInner**](GuidanceVmwareSizingPlanAfterActionPriceSetsInner.md) |  | [optional] 
**Config** | Pointer to [**GuidanceVmwareSizingPlanAfterActionConfig**](GuidanceVmwareSizingPlanAfterActionConfig.md) |  | [optional] 

## Methods

### NewGuidanceVmwareSizingPlanAfterAction

`func NewGuidanceVmwareSizingPlanAfterAction() *GuidanceVmwareSizingPlanAfterAction`

NewGuidanceVmwareSizingPlanAfterAction instantiates a new GuidanceVmwareSizingPlanAfterAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GuidanceVmwareSizingPlanAfterAction) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GuidanceVmwareSizingPlanAfterAction) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GuidanceVmwareSizingPlanAfterAction) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GuidanceVmwareSizingPlanAfterAction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GuidanceVmwareSizingPlanAfterAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GuidanceVmwareSizingPlanAfterAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GuidanceVmwareSizingPlanAfterAction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GuidanceVmwareSizingPlanAfterAction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *GuidanceVmwareSizingPlanAfterAction) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GuidanceVmwareSizingPlanAfterAction) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GuidanceVmwareSizingPlanAfterAction) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GuidanceVmwareSizingPlanAfterAction) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetActive

`func (o *GuidanceVmwareSizingPlanAfterAction) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GuidanceVmwareSizingPlanAfterAction) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GuidanceVmwareSizingPlanAfterAction) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GuidanceVmwareSizingPlanAfterAction) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetSortOrder

`func (o *GuidanceVmwareSizingPlanAfterAction) GetSortOrder() int64`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *GuidanceVmwareSizingPlanAfterAction) GetSortOrderOk() (*int64, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *GuidanceVmwareSizingPlanAfterAction) SetSortOrder(v int64)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *GuidanceVmwareSizingPlanAfterAction) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetDescription

`func (o *GuidanceVmwareSizingPlanAfterAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GuidanceVmwareSizingPlanAfterAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GuidanceVmwareSizingPlanAfterAction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GuidanceVmwareSizingPlanAfterAction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMaxStorage

`func (o *GuidanceVmwareSizingPlanAfterAction) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *GuidanceVmwareSizingPlanAfterAction) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *GuidanceVmwareSizingPlanAfterAction) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *GuidanceVmwareSizingPlanAfterAction) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetMaxMemory

`func (o *GuidanceVmwareSizingPlanAfterAction) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *GuidanceVmwareSizingPlanAfterAction) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *GuidanceVmwareSizingPlanAfterAction) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *GuidanceVmwareSizingPlanAfterAction) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxCpu

`func (o *GuidanceVmwareSizingPlanAfterAction) GetMaxCpu() int64`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *GuidanceVmwareSizingPlanAfterAction) GetMaxCpuOk() (*int64, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *GuidanceVmwareSizingPlanAfterAction) SetMaxCpu(v int64)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *GuidanceVmwareSizingPlanAfterAction) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### GetMaxCores

`func (o *GuidanceVmwareSizingPlanAfterAction) GetMaxCores() int64`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *GuidanceVmwareSizingPlanAfterAction) GetMaxCoresOk() (*int64, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *GuidanceVmwareSizingPlanAfterAction) SetMaxCores(v int64)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *GuidanceVmwareSizingPlanAfterAction) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetMaxDisks

`func (o *GuidanceVmwareSizingPlanAfterAction) GetMaxDisks() string`

GetMaxDisks returns the MaxDisks field if non-nil, zero value otherwise.

### GetMaxDisksOk

`func (o *GuidanceVmwareSizingPlanAfterAction) GetMaxDisksOk() (*string, bool)`

GetMaxDisksOk returns a tuple with the MaxDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDisks

`func (o *GuidanceVmwareSizingPlanAfterAction) SetMaxDisks(v string)`

SetMaxDisks sets MaxDisks field to given value.

### HasMaxDisks

`func (o *GuidanceVmwareSizingPlanAfterAction) HasMaxDisks() bool`

HasMaxDisks returns a boolean if a field has been set.

### SetMaxDisksNil

`func (o *GuidanceVmwareSizingPlanAfterAction) SetMaxDisksNil(b bool)`

 SetMaxDisksNil sets the value for MaxDisks to be an explicit nil

### UnsetMaxDisks
`func (o *GuidanceVmwareSizingPlanAfterAction) UnsetMaxDisks()`

UnsetMaxDisks ensures that no value is present for MaxDisks, not even an explicit nil
### GetCoresPerSocket

`func (o *GuidanceVmwareSizingPlanAfterAction) GetCoresPerSocket() int64`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *GuidanceVmwareSizingPlanAfterAction) GetCoresPerSocketOk() (*int64, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *GuidanceVmwareSizingPlanAfterAction) SetCoresPerSocket(v int64)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *GuidanceVmwareSizingPlanAfterAction) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### GetCustomCpu

`func (o *GuidanceVmwareSizingPlanAfterAction) GetCustomCpu() bool`

GetCustomCpu returns the CustomCpu field if non-nil, zero value otherwise.

### GetCustomCpuOk

`func (o *GuidanceVmwareSizingPlanAfterAction) GetCustomCpuOk() (*bool, bool)`

GetCustomCpuOk returns a tuple with the CustomCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCpu

`func (o *GuidanceVmwareSizingPlanAfterAction) SetCustomCpu(v bool)`

SetCustomCpu sets CustomCpu field to given value.

### HasCustomCpu

`func (o *GuidanceVmwareSizingPlanAfterAction) HasCustomCpu() bool`

HasCustomCpu returns a boolean if a field has been set.

### GetCustomCores

`func (o *GuidanceVmwareSizingPlanAfterAction) GetCustomCores() bool`

GetCustomCores returns the CustomCores field if non-nil, zero value otherwise.

### GetCustomCoresOk

`func (o *GuidanceVmwareSizingPlanAfterAction) GetCustomCoresOk() (*bool, bool)`

GetCustomCoresOk returns a tuple with the CustomCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCores

`func (o *GuidanceVmwareSizingPlanAfterAction) SetCustomCores(v bool)`

SetCustomCores sets CustomCores field to given value.

### HasCustomCores

`func (o *GuidanceVmwareSizingPlanAfterAction) HasCustomCores() bool`

HasCustomCores returns a boolean if a field has been set.

### GetCustomMaxStorage

`func (o *GuidanceVmwareSizingPlanAfterAction) GetCustomMaxStorage() bool`

GetCustomMaxStorage returns the CustomMaxStorage field if non-nil, zero value otherwise.

### GetCustomMaxStorageOk

`func (o *GuidanceVmwareSizingPlanAfterAction) GetCustomMaxStorageOk() (*bool, bool)`

GetCustomMaxStorageOk returns a tuple with the CustomMaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxStorage

`func (o *GuidanceVmwareSizingPlanAfterAction) SetCustomMaxStorage(v bool)`

SetCustomMaxStorage sets CustomMaxStorage field to given value.

### HasCustomMaxStorage

`func (o *GuidanceVmwareSizingPlanAfterAction) HasCustomMaxStorage() bool`

HasCustomMaxStorage returns a boolean if a field has been set.

### GetCustomMaxDataStorage

`func (o *GuidanceVmwareSizingPlanAfterAction) GetCustomMaxDataStorage() bool`

GetCustomMaxDataStorage returns the CustomMaxDataStorage field if non-nil, zero value otherwise.

### GetCustomMaxDataStorageOk

`func (o *GuidanceVmwareSizingPlanAfterAction) GetCustomMaxDataStorageOk() (*bool, bool)`

GetCustomMaxDataStorageOk returns a tuple with the CustomMaxDataStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxDataStorage

`func (o *GuidanceVmwareSizingPlanAfterAction) SetCustomMaxDataStorage(v bool)`

SetCustomMaxDataStorage sets CustomMaxDataStorage field to given value.

### HasCustomMaxDataStorage

`func (o *GuidanceVmwareSizingPlanAfterAction) HasCustomMaxDataStorage() bool`

HasCustomMaxDataStorage returns a boolean if a field has been set.

### GetCustomMaxMemory

`func (o *GuidanceVmwareSizingPlanAfterAction) GetCustomMaxMemory() bool`

GetCustomMaxMemory returns the CustomMaxMemory field if non-nil, zero value otherwise.

### GetCustomMaxMemoryOk

`func (o *GuidanceVmwareSizingPlanAfterAction) GetCustomMaxMemoryOk() (*bool, bool)`

GetCustomMaxMemoryOk returns a tuple with the CustomMaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxMemory

`func (o *GuidanceVmwareSizingPlanAfterAction) SetCustomMaxMemory(v bool)`

SetCustomMaxMemory sets CustomMaxMemory field to given value.

### HasCustomMaxMemory

`func (o *GuidanceVmwareSizingPlanAfterAction) HasCustomMaxMemory() bool`

HasCustomMaxMemory returns a boolean if a field has been set.

### GetAddVolumes

`func (o *GuidanceVmwareSizingPlanAfterAction) GetAddVolumes() bool`

GetAddVolumes returns the AddVolumes field if non-nil, zero value otherwise.

### GetAddVolumesOk

`func (o *GuidanceVmwareSizingPlanAfterAction) GetAddVolumesOk() (*bool, bool)`

GetAddVolumesOk returns a tuple with the AddVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddVolumes

`func (o *GuidanceVmwareSizingPlanAfterAction) SetAddVolumes(v bool)`

SetAddVolumes sets AddVolumes field to given value.

### HasAddVolumes

`func (o *GuidanceVmwareSizingPlanAfterAction) HasAddVolumes() bool`

HasAddVolumes returns a boolean if a field has been set.

### GetMemoryOptionSource

`func (o *GuidanceVmwareSizingPlanAfterAction) GetMemoryOptionSource() string`

GetMemoryOptionSource returns the MemoryOptionSource field if non-nil, zero value otherwise.

### GetMemoryOptionSourceOk

`func (o *GuidanceVmwareSizingPlanAfterAction) GetMemoryOptionSourceOk() (*string, bool)`

GetMemoryOptionSourceOk returns a tuple with the MemoryOptionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryOptionSource

`func (o *GuidanceVmwareSizingPlanAfterAction) SetMemoryOptionSource(v string)`

SetMemoryOptionSource sets MemoryOptionSource field to given value.

### HasMemoryOptionSource

`func (o *GuidanceVmwareSizingPlanAfterAction) HasMemoryOptionSource() bool`

HasMemoryOptionSource returns a boolean if a field has been set.

### SetMemoryOptionSourceNil

`func (o *GuidanceVmwareSizingPlanAfterAction) SetMemoryOptionSourceNil(b bool)`

 SetMemoryOptionSourceNil sets the value for MemoryOptionSource to be an explicit nil

### UnsetMemoryOptionSource
`func (o *GuidanceVmwareSizingPlanAfterAction) UnsetMemoryOptionSource()`

UnsetMemoryOptionSource ensures that no value is present for MemoryOptionSource, not even an explicit nil
### GetCpuOptionSource

`func (o *GuidanceVmwareSizingPlanAfterAction) GetCpuOptionSource() string`

GetCpuOptionSource returns the CpuOptionSource field if non-nil, zero value otherwise.

### GetCpuOptionSourceOk

`func (o *GuidanceVmwareSizingPlanAfterAction) GetCpuOptionSourceOk() (*string, bool)`

GetCpuOptionSourceOk returns a tuple with the CpuOptionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuOptionSource

`func (o *GuidanceVmwareSizingPlanAfterAction) SetCpuOptionSource(v string)`

SetCpuOptionSource sets CpuOptionSource field to given value.

### HasCpuOptionSource

`func (o *GuidanceVmwareSizingPlanAfterAction) HasCpuOptionSource() bool`

HasCpuOptionSource returns a boolean if a field has been set.

### SetCpuOptionSourceNil

`func (o *GuidanceVmwareSizingPlanAfterAction) SetCpuOptionSourceNil(b bool)`

 SetCpuOptionSourceNil sets the value for CpuOptionSource to be an explicit nil

### UnsetCpuOptionSource
`func (o *GuidanceVmwareSizingPlanAfterAction) UnsetCpuOptionSource()`

UnsetCpuOptionSource ensures that no value is present for CpuOptionSource, not even an explicit nil
### GetDateCreated

`func (o *GuidanceVmwareSizingPlanAfterAction) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GuidanceVmwareSizingPlanAfterAction) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GuidanceVmwareSizingPlanAfterAction) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GuidanceVmwareSizingPlanAfterAction) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GuidanceVmwareSizingPlanAfterAction) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GuidanceVmwareSizingPlanAfterAction) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GuidanceVmwareSizingPlanAfterAction) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GuidanceVmwareSizingPlanAfterAction) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetRegionCode

`func (o *GuidanceVmwareSizingPlanAfterAction) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *GuidanceVmwareSizingPlanAfterAction) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *GuidanceVmwareSizingPlanAfterAction) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *GuidanceVmwareSizingPlanAfterAction) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### SetRegionCodeNil

`func (o *GuidanceVmwareSizingPlanAfterAction) SetRegionCodeNil(b bool)`

 SetRegionCodeNil sets the value for RegionCode to be an explicit nil

### UnsetRegionCode
`func (o *GuidanceVmwareSizingPlanAfterAction) UnsetRegionCode()`

UnsetRegionCode ensures that no value is present for RegionCode, not even an explicit nil
### GetVisibility

`func (o *GuidanceVmwareSizingPlanAfterAction) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GuidanceVmwareSizingPlanAfterAction) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GuidanceVmwareSizingPlanAfterAction) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GuidanceVmwareSizingPlanAfterAction) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetEditable

`func (o *GuidanceVmwareSizingPlanAfterAction) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *GuidanceVmwareSizingPlanAfterAction) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *GuidanceVmwareSizingPlanAfterAction) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *GuidanceVmwareSizingPlanAfterAction) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetProvisionType

`func (o *GuidanceVmwareSizingPlanAfterAction) GetProvisionType() GuidanceVmwareSizingPlanAfterActionProvisionType`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *GuidanceVmwareSizingPlanAfterAction) GetProvisionTypeOk() (*GuidanceVmwareSizingPlanAfterActionProvisionType, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *GuidanceVmwareSizingPlanAfterAction) SetProvisionType(v GuidanceVmwareSizingPlanAfterActionProvisionType)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *GuidanceVmwareSizingPlanAfterAction) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetTenants

`func (o *GuidanceVmwareSizingPlanAfterAction) GetTenants() string`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *GuidanceVmwareSizingPlanAfterAction) GetTenantsOk() (*string, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *GuidanceVmwareSizingPlanAfterAction) SetTenants(v string)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *GuidanceVmwareSizingPlanAfterAction) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetPriceSets

`func (o *GuidanceVmwareSizingPlanAfterAction) GetPriceSets() []GuidanceVmwareSizingPlanAfterActionPriceSetsInner`

GetPriceSets returns the PriceSets field if non-nil, zero value otherwise.

### GetPriceSetsOk

`func (o *GuidanceVmwareSizingPlanAfterAction) GetPriceSetsOk() (*[]GuidanceVmwareSizingPlanAfterActionPriceSetsInner, bool)`

GetPriceSetsOk returns a tuple with the PriceSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSets

`func (o *GuidanceVmwareSizingPlanAfterAction) SetPriceSets(v []GuidanceVmwareSizingPlanAfterActionPriceSetsInner)`

SetPriceSets sets PriceSets field to given value.

### HasPriceSets

`func (o *GuidanceVmwareSizingPlanAfterAction) HasPriceSets() bool`

HasPriceSets returns a boolean if a field has been set.

### GetConfig

`func (o *GuidanceVmwareSizingPlanAfterAction) GetConfig() GuidanceVmwareSizingPlanAfterActionConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GuidanceVmwareSizingPlanAfterAction) GetConfigOk() (*GuidanceVmwareSizingPlanAfterActionConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GuidanceVmwareSizingPlanAfterAction) SetConfig(v GuidanceVmwareSizingPlanAfterActionConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GuidanceVmwareSizingPlanAfterAction) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


