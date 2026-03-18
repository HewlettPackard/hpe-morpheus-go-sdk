# GuidanceVmwareSizingPlanBeforeAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**SortOrder** | Pointer to **int64** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**MaxStorage** | Pointer to **int64** |  | [optional] 
**MaxMemory** | Pointer to **int64** |  | [optional] 
**MaxCpu** | Pointer to **NullableString** |  | [optional] 
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
**ProvisionType** | Pointer to [**GuidanceVmwareSizingPlanBeforeActionProvisionType**](GuidanceVmwareSizingPlanBeforeActionProvisionType.md) |  | [optional] 
**Tenants** | Pointer to **string** |  | [optional] 
**PriceSets** | Pointer to [**[]GuidanceVmwareSizingPlanBeforeActionPriceSetsInner**](GuidanceVmwareSizingPlanBeforeActionPriceSetsInner.md) |  | [optional] 
**Config** | Pointer to [**GuidanceVmwareSizingPlanBeforeActionConfig**](GuidanceVmwareSizingPlanBeforeActionConfig.md) |  | [optional] 

## Methods

### NewGuidanceVmwareSizingPlanBeforeAction

`func NewGuidanceVmwareSizingPlanBeforeAction() *GuidanceVmwareSizingPlanBeforeAction`

NewGuidanceVmwareSizingPlanBeforeAction instantiates a new GuidanceVmwareSizingPlanBeforeAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuidanceVmwareSizingPlanBeforeActionWithDefaults

`func NewGuidanceVmwareSizingPlanBeforeActionWithDefaults() *GuidanceVmwareSizingPlanBeforeAction`

NewGuidanceVmwareSizingPlanBeforeActionWithDefaults instantiates a new GuidanceVmwareSizingPlanBeforeAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GuidanceVmwareSizingPlanBeforeAction) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GuidanceVmwareSizingPlanBeforeAction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GuidanceVmwareSizingPlanBeforeAction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GuidanceVmwareSizingPlanBeforeAction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GuidanceVmwareSizingPlanBeforeAction) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GuidanceVmwareSizingPlanBeforeAction) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetActive

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GuidanceVmwareSizingPlanBeforeAction) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GuidanceVmwareSizingPlanBeforeAction) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetSortOrder

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetSortOrder() int64`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetSortOrderOk() (*int64, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *GuidanceVmwareSizingPlanBeforeAction) SetSortOrder(v int64)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *GuidanceVmwareSizingPlanBeforeAction) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetDescription

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GuidanceVmwareSizingPlanBeforeAction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GuidanceVmwareSizingPlanBeforeAction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GuidanceVmwareSizingPlanBeforeAction) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GuidanceVmwareSizingPlanBeforeAction) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMaxStorage

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *GuidanceVmwareSizingPlanBeforeAction) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *GuidanceVmwareSizingPlanBeforeAction) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetMaxMemory

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *GuidanceVmwareSizingPlanBeforeAction) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *GuidanceVmwareSizingPlanBeforeAction) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxCpu

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetMaxCpu() string`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetMaxCpuOk() (*string, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *GuidanceVmwareSizingPlanBeforeAction) SetMaxCpu(v string)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *GuidanceVmwareSizingPlanBeforeAction) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### SetMaxCpuNil

`func (o *GuidanceVmwareSizingPlanBeforeAction) SetMaxCpuNil(b bool)`

 SetMaxCpuNil sets the value for MaxCpu to be an explicit nil

### UnsetMaxCpu
`func (o *GuidanceVmwareSizingPlanBeforeAction) UnsetMaxCpu()`

UnsetMaxCpu ensures that no value is present for MaxCpu, not even an explicit nil
### GetMaxCores

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetMaxCores() int64`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetMaxCoresOk() (*int64, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *GuidanceVmwareSizingPlanBeforeAction) SetMaxCores(v int64)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *GuidanceVmwareSizingPlanBeforeAction) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetMaxDisks

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetMaxDisks() string`

GetMaxDisks returns the MaxDisks field if non-nil, zero value otherwise.

### GetMaxDisksOk

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetMaxDisksOk() (*string, bool)`

GetMaxDisksOk returns a tuple with the MaxDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDisks

`func (o *GuidanceVmwareSizingPlanBeforeAction) SetMaxDisks(v string)`

SetMaxDisks sets MaxDisks field to given value.

### HasMaxDisks

`func (o *GuidanceVmwareSizingPlanBeforeAction) HasMaxDisks() bool`

HasMaxDisks returns a boolean if a field has been set.

### SetMaxDisksNil

`func (o *GuidanceVmwareSizingPlanBeforeAction) SetMaxDisksNil(b bool)`

 SetMaxDisksNil sets the value for MaxDisks to be an explicit nil

### UnsetMaxDisks
`func (o *GuidanceVmwareSizingPlanBeforeAction) UnsetMaxDisks()`

UnsetMaxDisks ensures that no value is present for MaxDisks, not even an explicit nil
### GetCoresPerSocket

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetCoresPerSocket() int64`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetCoresPerSocketOk() (*int64, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *GuidanceVmwareSizingPlanBeforeAction) SetCoresPerSocket(v int64)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *GuidanceVmwareSizingPlanBeforeAction) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### GetCustomCpu

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetCustomCpu() bool`

GetCustomCpu returns the CustomCpu field if non-nil, zero value otherwise.

### GetCustomCpuOk

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetCustomCpuOk() (*bool, bool)`

GetCustomCpuOk returns a tuple with the CustomCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCpu

`func (o *GuidanceVmwareSizingPlanBeforeAction) SetCustomCpu(v bool)`

SetCustomCpu sets CustomCpu field to given value.

### HasCustomCpu

`func (o *GuidanceVmwareSizingPlanBeforeAction) HasCustomCpu() bool`

HasCustomCpu returns a boolean if a field has been set.

### GetCustomCores

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetCustomCores() bool`

GetCustomCores returns the CustomCores field if non-nil, zero value otherwise.

### GetCustomCoresOk

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetCustomCoresOk() (*bool, bool)`

GetCustomCoresOk returns a tuple with the CustomCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCores

`func (o *GuidanceVmwareSizingPlanBeforeAction) SetCustomCores(v bool)`

SetCustomCores sets CustomCores field to given value.

### HasCustomCores

`func (o *GuidanceVmwareSizingPlanBeforeAction) HasCustomCores() bool`

HasCustomCores returns a boolean if a field has been set.

### GetCustomMaxStorage

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetCustomMaxStorage() bool`

GetCustomMaxStorage returns the CustomMaxStorage field if non-nil, zero value otherwise.

### GetCustomMaxStorageOk

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetCustomMaxStorageOk() (*bool, bool)`

GetCustomMaxStorageOk returns a tuple with the CustomMaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxStorage

`func (o *GuidanceVmwareSizingPlanBeforeAction) SetCustomMaxStorage(v bool)`

SetCustomMaxStorage sets CustomMaxStorage field to given value.

### HasCustomMaxStorage

`func (o *GuidanceVmwareSizingPlanBeforeAction) HasCustomMaxStorage() bool`

HasCustomMaxStorage returns a boolean if a field has been set.

### GetCustomMaxDataStorage

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetCustomMaxDataStorage() bool`

GetCustomMaxDataStorage returns the CustomMaxDataStorage field if non-nil, zero value otherwise.

### GetCustomMaxDataStorageOk

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetCustomMaxDataStorageOk() (*bool, bool)`

GetCustomMaxDataStorageOk returns a tuple with the CustomMaxDataStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxDataStorage

`func (o *GuidanceVmwareSizingPlanBeforeAction) SetCustomMaxDataStorage(v bool)`

SetCustomMaxDataStorage sets CustomMaxDataStorage field to given value.

### HasCustomMaxDataStorage

`func (o *GuidanceVmwareSizingPlanBeforeAction) HasCustomMaxDataStorage() bool`

HasCustomMaxDataStorage returns a boolean if a field has been set.

### GetCustomMaxMemory

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetCustomMaxMemory() bool`

GetCustomMaxMemory returns the CustomMaxMemory field if non-nil, zero value otherwise.

### GetCustomMaxMemoryOk

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetCustomMaxMemoryOk() (*bool, bool)`

GetCustomMaxMemoryOk returns a tuple with the CustomMaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxMemory

`func (o *GuidanceVmwareSizingPlanBeforeAction) SetCustomMaxMemory(v bool)`

SetCustomMaxMemory sets CustomMaxMemory field to given value.

### HasCustomMaxMemory

`func (o *GuidanceVmwareSizingPlanBeforeAction) HasCustomMaxMemory() bool`

HasCustomMaxMemory returns a boolean if a field has been set.

### GetAddVolumes

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetAddVolumes() bool`

GetAddVolumes returns the AddVolumes field if non-nil, zero value otherwise.

### GetAddVolumesOk

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetAddVolumesOk() (*bool, bool)`

GetAddVolumesOk returns a tuple with the AddVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddVolumes

`func (o *GuidanceVmwareSizingPlanBeforeAction) SetAddVolumes(v bool)`

SetAddVolumes sets AddVolumes field to given value.

### HasAddVolumes

`func (o *GuidanceVmwareSizingPlanBeforeAction) HasAddVolumes() bool`

HasAddVolumes returns a boolean if a field has been set.

### GetMemoryOptionSource

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetMemoryOptionSource() string`

GetMemoryOptionSource returns the MemoryOptionSource field if non-nil, zero value otherwise.

### GetMemoryOptionSourceOk

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetMemoryOptionSourceOk() (*string, bool)`

GetMemoryOptionSourceOk returns a tuple with the MemoryOptionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryOptionSource

`func (o *GuidanceVmwareSizingPlanBeforeAction) SetMemoryOptionSource(v string)`

SetMemoryOptionSource sets MemoryOptionSource field to given value.

### HasMemoryOptionSource

`func (o *GuidanceVmwareSizingPlanBeforeAction) HasMemoryOptionSource() bool`

HasMemoryOptionSource returns a boolean if a field has been set.

### SetMemoryOptionSourceNil

`func (o *GuidanceVmwareSizingPlanBeforeAction) SetMemoryOptionSourceNil(b bool)`

 SetMemoryOptionSourceNil sets the value for MemoryOptionSource to be an explicit nil

### UnsetMemoryOptionSource
`func (o *GuidanceVmwareSizingPlanBeforeAction) UnsetMemoryOptionSource()`

UnsetMemoryOptionSource ensures that no value is present for MemoryOptionSource, not even an explicit nil
### GetCpuOptionSource

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetCpuOptionSource() string`

GetCpuOptionSource returns the CpuOptionSource field if non-nil, zero value otherwise.

### GetCpuOptionSourceOk

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetCpuOptionSourceOk() (*string, bool)`

GetCpuOptionSourceOk returns a tuple with the CpuOptionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuOptionSource

`func (o *GuidanceVmwareSizingPlanBeforeAction) SetCpuOptionSource(v string)`

SetCpuOptionSource sets CpuOptionSource field to given value.

### HasCpuOptionSource

`func (o *GuidanceVmwareSizingPlanBeforeAction) HasCpuOptionSource() bool`

HasCpuOptionSource returns a boolean if a field has been set.

### SetCpuOptionSourceNil

`func (o *GuidanceVmwareSizingPlanBeforeAction) SetCpuOptionSourceNil(b bool)`

 SetCpuOptionSourceNil sets the value for CpuOptionSource to be an explicit nil

### UnsetCpuOptionSource
`func (o *GuidanceVmwareSizingPlanBeforeAction) UnsetCpuOptionSource()`

UnsetCpuOptionSource ensures that no value is present for CpuOptionSource, not even an explicit nil
### GetDateCreated

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GuidanceVmwareSizingPlanBeforeAction) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GuidanceVmwareSizingPlanBeforeAction) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GuidanceVmwareSizingPlanBeforeAction) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GuidanceVmwareSizingPlanBeforeAction) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetRegionCode

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *GuidanceVmwareSizingPlanBeforeAction) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *GuidanceVmwareSizingPlanBeforeAction) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### SetRegionCodeNil

`func (o *GuidanceVmwareSizingPlanBeforeAction) SetRegionCodeNil(b bool)`

 SetRegionCodeNil sets the value for RegionCode to be an explicit nil

### UnsetRegionCode
`func (o *GuidanceVmwareSizingPlanBeforeAction) UnsetRegionCode()`

UnsetRegionCode ensures that no value is present for RegionCode, not even an explicit nil
### GetVisibility

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GuidanceVmwareSizingPlanBeforeAction) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GuidanceVmwareSizingPlanBeforeAction) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetEditable

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *GuidanceVmwareSizingPlanBeforeAction) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *GuidanceVmwareSizingPlanBeforeAction) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetProvisionType

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetProvisionType() GuidanceVmwareSizingPlanBeforeActionProvisionType`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetProvisionTypeOk() (*GuidanceVmwareSizingPlanBeforeActionProvisionType, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *GuidanceVmwareSizingPlanBeforeAction) SetProvisionType(v GuidanceVmwareSizingPlanBeforeActionProvisionType)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *GuidanceVmwareSizingPlanBeforeAction) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetTenants

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetTenants() string`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetTenantsOk() (*string, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *GuidanceVmwareSizingPlanBeforeAction) SetTenants(v string)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *GuidanceVmwareSizingPlanBeforeAction) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetPriceSets

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetPriceSets() []GuidanceVmwareSizingPlanBeforeActionPriceSetsInner`

GetPriceSets returns the PriceSets field if non-nil, zero value otherwise.

### GetPriceSetsOk

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetPriceSetsOk() (*[]GuidanceVmwareSizingPlanBeforeActionPriceSetsInner, bool)`

GetPriceSetsOk returns a tuple with the PriceSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSets

`func (o *GuidanceVmwareSizingPlanBeforeAction) SetPriceSets(v []GuidanceVmwareSizingPlanBeforeActionPriceSetsInner)`

SetPriceSets sets PriceSets field to given value.

### HasPriceSets

`func (o *GuidanceVmwareSizingPlanBeforeAction) HasPriceSets() bool`

HasPriceSets returns a boolean if a field has been set.

### GetConfig

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetConfig() GuidanceVmwareSizingPlanBeforeActionConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GuidanceVmwareSizingPlanBeforeAction) GetConfigOk() (*GuidanceVmwareSizingPlanBeforeActionConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GuidanceVmwareSizingPlanBeforeAction) SetConfig(v GuidanceVmwareSizingPlanBeforeActionConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GuidanceVmwareSizingPlanBeforeAction) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


