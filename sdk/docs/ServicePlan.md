# ServicePlan

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

### NewServicePlan

`func NewServicePlan() *ServicePlan`

NewServicePlan instantiates a new ServicePlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePlanWithDefaults

`func NewServicePlanWithDefaults() *ServicePlan`

NewServicePlanWithDefaults instantiates a new ServicePlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServicePlan) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServicePlan) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServicePlan) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ServicePlan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ServicePlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServicePlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServicePlan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServicePlan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *ServicePlan) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ServicePlan) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ServicePlan) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ServicePlan) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetActive

`func (o *ServicePlan) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ServicePlan) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ServicePlan) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ServicePlan) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetSortOrder

`func (o *ServicePlan) GetSortOrder() int64`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *ServicePlan) GetSortOrderOk() (*int64, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *ServicePlan) SetSortOrder(v int64)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *ServicePlan) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetDescription

`func (o *ServicePlan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServicePlan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServicePlan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServicePlan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMaxStorage

`func (o *ServicePlan) GetMaxStorage() float32`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *ServicePlan) GetMaxStorageOk() (*float32, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *ServicePlan) SetMaxStorage(v float32)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *ServicePlan) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetMaxMemory

`func (o *ServicePlan) GetMaxMemory() float32`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *ServicePlan) GetMaxMemoryOk() (*float32, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *ServicePlan) SetMaxMemory(v float32)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *ServicePlan) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxCpu

`func (o *ServicePlan) GetMaxCpu() float32`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *ServicePlan) GetMaxCpuOk() (*float32, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *ServicePlan) SetMaxCpu(v float32)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *ServicePlan) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### GetMaxCores

`func (o *ServicePlan) GetMaxCores() float32`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *ServicePlan) GetMaxCoresOk() (*float32, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *ServicePlan) SetMaxCores(v float32)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *ServicePlan) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetMaxDisks

`func (o *ServicePlan) GetMaxDisks() float32`

GetMaxDisks returns the MaxDisks field if non-nil, zero value otherwise.

### GetMaxDisksOk

`func (o *ServicePlan) GetMaxDisksOk() (*float32, bool)`

GetMaxDisksOk returns a tuple with the MaxDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDisks

`func (o *ServicePlan) SetMaxDisks(v float32)`

SetMaxDisks sets MaxDisks field to given value.

### HasMaxDisks

`func (o *ServicePlan) HasMaxDisks() bool`

HasMaxDisks returns a boolean if a field has been set.

### GetCoresPerSocket

`func (o *ServicePlan) GetCoresPerSocket() float32`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *ServicePlan) GetCoresPerSocketOk() (*float32, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *ServicePlan) SetCoresPerSocket(v float32)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *ServicePlan) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### GetCustomCpu

`func (o *ServicePlan) GetCustomCpu() bool`

GetCustomCpu returns the CustomCpu field if non-nil, zero value otherwise.

### GetCustomCpuOk

`func (o *ServicePlan) GetCustomCpuOk() (*bool, bool)`

GetCustomCpuOk returns a tuple with the CustomCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCpu

`func (o *ServicePlan) SetCustomCpu(v bool)`

SetCustomCpu sets CustomCpu field to given value.

### HasCustomCpu

`func (o *ServicePlan) HasCustomCpu() bool`

HasCustomCpu returns a boolean if a field has been set.

### GetCustomCores

`func (o *ServicePlan) GetCustomCores() bool`

GetCustomCores returns the CustomCores field if non-nil, zero value otherwise.

### GetCustomCoresOk

`func (o *ServicePlan) GetCustomCoresOk() (*bool, bool)`

GetCustomCoresOk returns a tuple with the CustomCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCores

`func (o *ServicePlan) SetCustomCores(v bool)`

SetCustomCores sets CustomCores field to given value.

### HasCustomCores

`func (o *ServicePlan) HasCustomCores() bool`

HasCustomCores returns a boolean if a field has been set.

### GetCustomMaxStorage

`func (o *ServicePlan) GetCustomMaxStorage() bool`

GetCustomMaxStorage returns the CustomMaxStorage field if non-nil, zero value otherwise.

### GetCustomMaxStorageOk

`func (o *ServicePlan) GetCustomMaxStorageOk() (*bool, bool)`

GetCustomMaxStorageOk returns a tuple with the CustomMaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxStorage

`func (o *ServicePlan) SetCustomMaxStorage(v bool)`

SetCustomMaxStorage sets CustomMaxStorage field to given value.

### HasCustomMaxStorage

`func (o *ServicePlan) HasCustomMaxStorage() bool`

HasCustomMaxStorage returns a boolean if a field has been set.

### GetCustomMaxDataStorage

`func (o *ServicePlan) GetCustomMaxDataStorage() bool`

GetCustomMaxDataStorage returns the CustomMaxDataStorage field if non-nil, zero value otherwise.

### GetCustomMaxDataStorageOk

`func (o *ServicePlan) GetCustomMaxDataStorageOk() (*bool, bool)`

GetCustomMaxDataStorageOk returns a tuple with the CustomMaxDataStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxDataStorage

`func (o *ServicePlan) SetCustomMaxDataStorage(v bool)`

SetCustomMaxDataStorage sets CustomMaxDataStorage field to given value.

### HasCustomMaxDataStorage

`func (o *ServicePlan) HasCustomMaxDataStorage() bool`

HasCustomMaxDataStorage returns a boolean if a field has been set.

### GetCustomMaxMemory

`func (o *ServicePlan) GetCustomMaxMemory() bool`

GetCustomMaxMemory returns the CustomMaxMemory field if non-nil, zero value otherwise.

### GetCustomMaxMemoryOk

`func (o *ServicePlan) GetCustomMaxMemoryOk() (*bool, bool)`

GetCustomMaxMemoryOk returns a tuple with the CustomMaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxMemory

`func (o *ServicePlan) SetCustomMaxMemory(v bool)`

SetCustomMaxMemory sets CustomMaxMemory field to given value.

### HasCustomMaxMemory

`func (o *ServicePlan) HasCustomMaxMemory() bool`

HasCustomMaxMemory returns a boolean if a field has been set.

### GetAddVolumes

`func (o *ServicePlan) GetAddVolumes() bool`

GetAddVolumes returns the AddVolumes field if non-nil, zero value otherwise.

### GetAddVolumesOk

`func (o *ServicePlan) GetAddVolumesOk() (*bool, bool)`

GetAddVolumesOk returns a tuple with the AddVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddVolumes

`func (o *ServicePlan) SetAddVolumes(v bool)`

SetAddVolumes sets AddVolumes field to given value.

### HasAddVolumes

`func (o *ServicePlan) HasAddVolumes() bool`

HasAddVolumes returns a boolean if a field has been set.

### GetMemoryOptionSource

`func (o *ServicePlan) GetMemoryOptionSource() string`

GetMemoryOptionSource returns the MemoryOptionSource field if non-nil, zero value otherwise.

### GetMemoryOptionSourceOk

`func (o *ServicePlan) GetMemoryOptionSourceOk() (*string, bool)`

GetMemoryOptionSourceOk returns a tuple with the MemoryOptionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryOptionSource

`func (o *ServicePlan) SetMemoryOptionSource(v string)`

SetMemoryOptionSource sets MemoryOptionSource field to given value.

### HasMemoryOptionSource

`func (o *ServicePlan) HasMemoryOptionSource() bool`

HasMemoryOptionSource returns a boolean if a field has been set.

### GetCpuOptionSource

`func (o *ServicePlan) GetCpuOptionSource() string`

GetCpuOptionSource returns the CpuOptionSource field if non-nil, zero value otherwise.

### GetCpuOptionSourceOk

`func (o *ServicePlan) GetCpuOptionSourceOk() (*string, bool)`

GetCpuOptionSourceOk returns a tuple with the CpuOptionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuOptionSource

`func (o *ServicePlan) SetCpuOptionSource(v string)`

SetCpuOptionSource sets CpuOptionSource field to given value.

### HasCpuOptionSource

`func (o *ServicePlan) HasCpuOptionSource() bool`

HasCpuOptionSource returns a boolean if a field has been set.

### GetDateCreated

`func (o *ServicePlan) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ServicePlan) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ServicePlan) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ServicePlan) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ServicePlan) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ServicePlan) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ServicePlan) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ServicePlan) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetRegionCode

`func (o *ServicePlan) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *ServicePlan) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *ServicePlan) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *ServicePlan) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetVisibility

`func (o *ServicePlan) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ServicePlan) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ServicePlan) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ServicePlan) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetEditable

`func (o *ServicePlan) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *ServicePlan) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *ServicePlan) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *ServicePlan) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetProvisionType

`func (o *ServicePlan) GetProvisionType() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionProvisionType`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *ServicePlan) GetProvisionTypeOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionProvisionType, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *ServicePlan) SetProvisionType(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionProvisionType)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *ServicePlan) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetTenants

`func (o *ServicePlan) GetTenants() string`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ServicePlan) GetTenantsOk() (*string, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ServicePlan) SetTenants(v string)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ServicePlan) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetPriceSets

`func (o *ServicePlan) GetPriceSets() []ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionPriceSetsInner`

GetPriceSets returns the PriceSets field if non-nil, zero value otherwise.

### GetPriceSetsOk

`func (o *ServicePlan) GetPriceSetsOk() (*[]ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionPriceSetsInner, bool)`

GetPriceSetsOk returns a tuple with the PriceSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSets

`func (o *ServicePlan) SetPriceSets(v []ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionPriceSetsInner)`

SetPriceSets sets PriceSets field to given value.

### HasPriceSets

`func (o *ServicePlan) HasPriceSets() bool`

HasPriceSets returns a boolean if a field has been set.

### GetConfig

`func (o *ServicePlan) GetConfig() ListServicePlans200ResponseAllOfServicePlansInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ServicePlan) GetConfigOk() (*ListServicePlans200ResponseAllOfServicePlansInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ServicePlan) SetConfig(v ListServicePlans200ResponseAllOfServicePlansInnerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ServicePlan) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetZones

`func (o *ServicePlan) GetZones() []ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *ServicePlan) GetZonesOk() (*[]ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *ServicePlan) SetZones(v []ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetZones sets Zones field to given value.

### HasZones

`func (o *ServicePlan) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetPermissions

`func (o *ServicePlan) GetPermissions() AddServicePlans200ResponseAllOfServicePlanPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ServicePlan) GetPermissionsOk() (*AddServicePlans200ResponseAllOfServicePlanPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ServicePlan) SetPermissions(v AddServicePlans200ResponseAllOfServicePlanPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ServicePlan) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


