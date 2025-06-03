# AddServicePlans200ResponseAllOfServicePlan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**SortOrder** | Pointer to **int64** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MaxStorage** | Pointer to **float32** |  | [optional] 
**MaxMemory** | Pointer to **float32** |  | [optional] 
**MaxCpu** | Pointer to **float32** |  | [optional] 
**MaxCores** | Pointer to **float32** |  | [optional] 
**MaxDisks** | Pointer to **float32** |  | [optional] 
**CoresPerSocket** | Pointer to **float32** |  | [optional] 
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
**Config** | Pointer to [**ListServicePlans200ResponseAllOfServicePlansInnerConfig**](ListServicePlans200ResponseAllOfServicePlansInnerConfig.md) |  | [optional] 
**Zones** | Pointer to [**[]ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 
**Permissions** | Pointer to [**AddServicePlans200ResponseAllOfServicePlanPermissions**](AddServicePlans200ResponseAllOfServicePlanPermissions.md) |  | [optional] 

## Methods

### NewAddServicePlans200ResponseAllOfServicePlan

`func NewAddServicePlans200ResponseAllOfServicePlan() *AddServicePlans200ResponseAllOfServicePlan`

NewAddServicePlans200ResponseAllOfServicePlan instantiates a new AddServicePlans200ResponseAllOfServicePlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddServicePlans200ResponseAllOfServicePlanWithDefaults

`func NewAddServicePlans200ResponseAllOfServicePlanWithDefaults() *AddServicePlans200ResponseAllOfServicePlan`

NewAddServicePlans200ResponseAllOfServicePlanWithDefaults instantiates a new AddServicePlans200ResponseAllOfServicePlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddServicePlans200ResponseAllOfServicePlan) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddServicePlans200ResponseAllOfServicePlan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddServicePlans200ResponseAllOfServicePlan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddServicePlans200ResponseAllOfServicePlan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AddServicePlans200ResponseAllOfServicePlan) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AddServicePlans200ResponseAllOfServicePlan) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetActive

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *AddServicePlans200ResponseAllOfServicePlan) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *AddServicePlans200ResponseAllOfServicePlan) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetSortOrder

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetSortOrder() int64`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetSortOrderOk() (*int64, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *AddServicePlans200ResponseAllOfServicePlan) SetSortOrder(v int64)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *AddServicePlans200ResponseAllOfServicePlan) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetDescription

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddServicePlans200ResponseAllOfServicePlan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddServicePlans200ResponseAllOfServicePlan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMaxStorage

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetMaxStorage() float32`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetMaxStorageOk() (*float32, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *AddServicePlans200ResponseAllOfServicePlan) SetMaxStorage(v float32)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *AddServicePlans200ResponseAllOfServicePlan) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetMaxMemory

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetMaxMemory() float32`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetMaxMemoryOk() (*float32, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *AddServicePlans200ResponseAllOfServicePlan) SetMaxMemory(v float32)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *AddServicePlans200ResponseAllOfServicePlan) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxCpu

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetMaxCpu() float32`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetMaxCpuOk() (*float32, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *AddServicePlans200ResponseAllOfServicePlan) SetMaxCpu(v float32)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *AddServicePlans200ResponseAllOfServicePlan) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### GetMaxCores

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetMaxCores() float32`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetMaxCoresOk() (*float32, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *AddServicePlans200ResponseAllOfServicePlan) SetMaxCores(v float32)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *AddServicePlans200ResponseAllOfServicePlan) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetMaxDisks

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetMaxDisks() float32`

GetMaxDisks returns the MaxDisks field if non-nil, zero value otherwise.

### GetMaxDisksOk

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetMaxDisksOk() (*float32, bool)`

GetMaxDisksOk returns a tuple with the MaxDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDisks

`func (o *AddServicePlans200ResponseAllOfServicePlan) SetMaxDisks(v float32)`

SetMaxDisks sets MaxDisks field to given value.

### HasMaxDisks

`func (o *AddServicePlans200ResponseAllOfServicePlan) HasMaxDisks() bool`

HasMaxDisks returns a boolean if a field has been set.

### GetCoresPerSocket

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetCoresPerSocket() float32`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetCoresPerSocketOk() (*float32, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *AddServicePlans200ResponseAllOfServicePlan) SetCoresPerSocket(v float32)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *AddServicePlans200ResponseAllOfServicePlan) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### GetCustomCpu

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetCustomCpu() bool`

GetCustomCpu returns the CustomCpu field if non-nil, zero value otherwise.

### GetCustomCpuOk

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetCustomCpuOk() (*bool, bool)`

GetCustomCpuOk returns a tuple with the CustomCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCpu

`func (o *AddServicePlans200ResponseAllOfServicePlan) SetCustomCpu(v bool)`

SetCustomCpu sets CustomCpu field to given value.

### HasCustomCpu

`func (o *AddServicePlans200ResponseAllOfServicePlan) HasCustomCpu() bool`

HasCustomCpu returns a boolean if a field has been set.

### GetCustomCores

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetCustomCores() bool`

GetCustomCores returns the CustomCores field if non-nil, zero value otherwise.

### GetCustomCoresOk

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetCustomCoresOk() (*bool, bool)`

GetCustomCoresOk returns a tuple with the CustomCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCores

`func (o *AddServicePlans200ResponseAllOfServicePlan) SetCustomCores(v bool)`

SetCustomCores sets CustomCores field to given value.

### HasCustomCores

`func (o *AddServicePlans200ResponseAllOfServicePlan) HasCustomCores() bool`

HasCustomCores returns a boolean if a field has been set.

### GetCustomMaxStorage

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetCustomMaxStorage() bool`

GetCustomMaxStorage returns the CustomMaxStorage field if non-nil, zero value otherwise.

### GetCustomMaxStorageOk

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetCustomMaxStorageOk() (*bool, bool)`

GetCustomMaxStorageOk returns a tuple with the CustomMaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxStorage

`func (o *AddServicePlans200ResponseAllOfServicePlan) SetCustomMaxStorage(v bool)`

SetCustomMaxStorage sets CustomMaxStorage field to given value.

### HasCustomMaxStorage

`func (o *AddServicePlans200ResponseAllOfServicePlan) HasCustomMaxStorage() bool`

HasCustomMaxStorage returns a boolean if a field has been set.

### GetCustomMaxDataStorage

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetCustomMaxDataStorage() bool`

GetCustomMaxDataStorage returns the CustomMaxDataStorage field if non-nil, zero value otherwise.

### GetCustomMaxDataStorageOk

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetCustomMaxDataStorageOk() (*bool, bool)`

GetCustomMaxDataStorageOk returns a tuple with the CustomMaxDataStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxDataStorage

`func (o *AddServicePlans200ResponseAllOfServicePlan) SetCustomMaxDataStorage(v bool)`

SetCustomMaxDataStorage sets CustomMaxDataStorage field to given value.

### HasCustomMaxDataStorage

`func (o *AddServicePlans200ResponseAllOfServicePlan) HasCustomMaxDataStorage() bool`

HasCustomMaxDataStorage returns a boolean if a field has been set.

### GetCustomMaxMemory

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetCustomMaxMemory() bool`

GetCustomMaxMemory returns the CustomMaxMemory field if non-nil, zero value otherwise.

### GetCustomMaxMemoryOk

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetCustomMaxMemoryOk() (*bool, bool)`

GetCustomMaxMemoryOk returns a tuple with the CustomMaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxMemory

`func (o *AddServicePlans200ResponseAllOfServicePlan) SetCustomMaxMemory(v bool)`

SetCustomMaxMemory sets CustomMaxMemory field to given value.

### HasCustomMaxMemory

`func (o *AddServicePlans200ResponseAllOfServicePlan) HasCustomMaxMemory() bool`

HasCustomMaxMemory returns a boolean if a field has been set.

### GetAddVolumes

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetAddVolumes() bool`

GetAddVolumes returns the AddVolumes field if non-nil, zero value otherwise.

### GetAddVolumesOk

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetAddVolumesOk() (*bool, bool)`

GetAddVolumesOk returns a tuple with the AddVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddVolumes

`func (o *AddServicePlans200ResponseAllOfServicePlan) SetAddVolumes(v bool)`

SetAddVolumes sets AddVolumes field to given value.

### HasAddVolumes

`func (o *AddServicePlans200ResponseAllOfServicePlan) HasAddVolumes() bool`

HasAddVolumes returns a boolean if a field has been set.

### GetMemoryOptionSource

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetMemoryOptionSource() string`

GetMemoryOptionSource returns the MemoryOptionSource field if non-nil, zero value otherwise.

### GetMemoryOptionSourceOk

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetMemoryOptionSourceOk() (*string, bool)`

GetMemoryOptionSourceOk returns a tuple with the MemoryOptionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryOptionSource

`func (o *AddServicePlans200ResponseAllOfServicePlan) SetMemoryOptionSource(v string)`

SetMemoryOptionSource sets MemoryOptionSource field to given value.

### HasMemoryOptionSource

`func (o *AddServicePlans200ResponseAllOfServicePlan) HasMemoryOptionSource() bool`

HasMemoryOptionSource returns a boolean if a field has been set.

### GetCpuOptionSource

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetCpuOptionSource() string`

GetCpuOptionSource returns the CpuOptionSource field if non-nil, zero value otherwise.

### GetCpuOptionSourceOk

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetCpuOptionSourceOk() (*string, bool)`

GetCpuOptionSourceOk returns a tuple with the CpuOptionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuOptionSource

`func (o *AddServicePlans200ResponseAllOfServicePlan) SetCpuOptionSource(v string)`

SetCpuOptionSource sets CpuOptionSource field to given value.

### HasCpuOptionSource

`func (o *AddServicePlans200ResponseAllOfServicePlan) HasCpuOptionSource() bool`

HasCpuOptionSource returns a boolean if a field has been set.

### GetDateCreated

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AddServicePlans200ResponseAllOfServicePlan) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AddServicePlans200ResponseAllOfServicePlan) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AddServicePlans200ResponseAllOfServicePlan) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AddServicePlans200ResponseAllOfServicePlan) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetRegionCode

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *AddServicePlans200ResponseAllOfServicePlan) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *AddServicePlans200ResponseAllOfServicePlan) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetVisibility

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AddServicePlans200ResponseAllOfServicePlan) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AddServicePlans200ResponseAllOfServicePlan) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetEditable

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *AddServicePlans200ResponseAllOfServicePlan) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *AddServicePlans200ResponseAllOfServicePlan) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetProvisionType

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetProvisionType() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionProvisionType`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetProvisionTypeOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionProvisionType, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *AddServicePlans200ResponseAllOfServicePlan) SetProvisionType(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionProvisionType)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *AddServicePlans200ResponseAllOfServicePlan) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetTenants

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetTenants() string`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetTenantsOk() (*string, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *AddServicePlans200ResponseAllOfServicePlan) SetTenants(v string)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *AddServicePlans200ResponseAllOfServicePlan) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetPriceSets

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetPriceSets() []ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionPriceSetsInner`

GetPriceSets returns the PriceSets field if non-nil, zero value otherwise.

### GetPriceSetsOk

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetPriceSetsOk() (*[]ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionPriceSetsInner, bool)`

GetPriceSetsOk returns a tuple with the PriceSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSets

`func (o *AddServicePlans200ResponseAllOfServicePlan) SetPriceSets(v []ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionPriceSetsInner)`

SetPriceSets sets PriceSets field to given value.

### HasPriceSets

`func (o *AddServicePlans200ResponseAllOfServicePlan) HasPriceSets() bool`

HasPriceSets returns a boolean if a field has been set.

### GetConfig

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetConfig() ListServicePlans200ResponseAllOfServicePlansInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetConfigOk() (*ListServicePlans200ResponseAllOfServicePlansInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddServicePlans200ResponseAllOfServicePlan) SetConfig(v ListServicePlans200ResponseAllOfServicePlansInnerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddServicePlans200ResponseAllOfServicePlan) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetZones

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetZones() []ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetZonesOk() (*[]ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *AddServicePlans200ResponseAllOfServicePlan) SetZones(v []ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetZones sets Zones field to given value.

### HasZones

`func (o *AddServicePlans200ResponseAllOfServicePlan) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetPermissions

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetPermissions() AddServicePlans200ResponseAllOfServicePlanPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *AddServicePlans200ResponseAllOfServicePlan) GetPermissionsOk() (*AddServicePlans200ResponseAllOfServicePlanPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *AddServicePlans200ResponseAllOfServicePlan) SetPermissions(v AddServicePlans200ResponseAllOfServicePlanPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *AddServicePlans200ResponseAllOfServicePlan) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


