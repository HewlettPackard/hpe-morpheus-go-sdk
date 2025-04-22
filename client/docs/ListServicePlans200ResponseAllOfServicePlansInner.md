# ListServicePlans200ResponseAllOfServicePlansInner

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
**Permissions** | Pointer to [**ListServicePlans200ResponseAllOfServicePlansInnerPermissions**](ListServicePlans200ResponseAllOfServicePlansInnerPermissions.md) |  | [optional] 

## Methods

### NewListServicePlans200ResponseAllOfServicePlansInner

`func NewListServicePlans200ResponseAllOfServicePlansInner() *ListServicePlans200ResponseAllOfServicePlansInner`

NewListServicePlans200ResponseAllOfServicePlansInner instantiates a new ListServicePlans200ResponseAllOfServicePlansInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListServicePlans200ResponseAllOfServicePlansInnerWithDefaults

`func NewListServicePlans200ResponseAllOfServicePlansInnerWithDefaults() *ListServicePlans200ResponseAllOfServicePlansInner`

NewListServicePlans200ResponseAllOfServicePlansInnerWithDefaults instantiates a new ListServicePlans200ResponseAllOfServicePlansInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetActive

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetSortOrder

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetSortOrder() int64`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetSortOrderOk() (*int64, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) SetSortOrder(v int64)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetDescription

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMaxStorage

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetMaxStorage() float32`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetMaxStorageOk() (*float32, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) SetMaxStorage(v float32)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetMaxMemory

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetMaxMemory() float32`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetMaxMemoryOk() (*float32, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) SetMaxMemory(v float32)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxCpu

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetMaxCpu() float32`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetMaxCpuOk() (*float32, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) SetMaxCpu(v float32)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### GetMaxCores

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetMaxCores() float32`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetMaxCoresOk() (*float32, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) SetMaxCores(v float32)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetMaxDisks

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetMaxDisks() float32`

GetMaxDisks returns the MaxDisks field if non-nil, zero value otherwise.

### GetMaxDisksOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetMaxDisksOk() (*float32, bool)`

GetMaxDisksOk returns a tuple with the MaxDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDisks

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) SetMaxDisks(v float32)`

SetMaxDisks sets MaxDisks field to given value.

### HasMaxDisks

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) HasMaxDisks() bool`

HasMaxDisks returns a boolean if a field has been set.

### GetCoresPerSocket

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetCoresPerSocket() float32`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetCoresPerSocketOk() (*float32, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) SetCoresPerSocket(v float32)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### GetCustomCpu

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetCustomCpu() bool`

GetCustomCpu returns the CustomCpu field if non-nil, zero value otherwise.

### GetCustomCpuOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetCustomCpuOk() (*bool, bool)`

GetCustomCpuOk returns a tuple with the CustomCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCpu

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) SetCustomCpu(v bool)`

SetCustomCpu sets CustomCpu field to given value.

### HasCustomCpu

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) HasCustomCpu() bool`

HasCustomCpu returns a boolean if a field has been set.

### GetCustomCores

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetCustomCores() bool`

GetCustomCores returns the CustomCores field if non-nil, zero value otherwise.

### GetCustomCoresOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetCustomCoresOk() (*bool, bool)`

GetCustomCoresOk returns a tuple with the CustomCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCores

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) SetCustomCores(v bool)`

SetCustomCores sets CustomCores field to given value.

### HasCustomCores

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) HasCustomCores() bool`

HasCustomCores returns a boolean if a field has been set.

### GetCustomMaxStorage

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetCustomMaxStorage() bool`

GetCustomMaxStorage returns the CustomMaxStorage field if non-nil, zero value otherwise.

### GetCustomMaxStorageOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetCustomMaxStorageOk() (*bool, bool)`

GetCustomMaxStorageOk returns a tuple with the CustomMaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxStorage

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) SetCustomMaxStorage(v bool)`

SetCustomMaxStorage sets CustomMaxStorage field to given value.

### HasCustomMaxStorage

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) HasCustomMaxStorage() bool`

HasCustomMaxStorage returns a boolean if a field has been set.

### GetCustomMaxDataStorage

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetCustomMaxDataStorage() bool`

GetCustomMaxDataStorage returns the CustomMaxDataStorage field if non-nil, zero value otherwise.

### GetCustomMaxDataStorageOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetCustomMaxDataStorageOk() (*bool, bool)`

GetCustomMaxDataStorageOk returns a tuple with the CustomMaxDataStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxDataStorage

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) SetCustomMaxDataStorage(v bool)`

SetCustomMaxDataStorage sets CustomMaxDataStorage field to given value.

### HasCustomMaxDataStorage

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) HasCustomMaxDataStorage() bool`

HasCustomMaxDataStorage returns a boolean if a field has been set.

### GetCustomMaxMemory

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetCustomMaxMemory() bool`

GetCustomMaxMemory returns the CustomMaxMemory field if non-nil, zero value otherwise.

### GetCustomMaxMemoryOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetCustomMaxMemoryOk() (*bool, bool)`

GetCustomMaxMemoryOk returns a tuple with the CustomMaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxMemory

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) SetCustomMaxMemory(v bool)`

SetCustomMaxMemory sets CustomMaxMemory field to given value.

### HasCustomMaxMemory

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) HasCustomMaxMemory() bool`

HasCustomMaxMemory returns a boolean if a field has been set.

### GetAddVolumes

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetAddVolumes() bool`

GetAddVolumes returns the AddVolumes field if non-nil, zero value otherwise.

### GetAddVolumesOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetAddVolumesOk() (*bool, bool)`

GetAddVolumesOk returns a tuple with the AddVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddVolumes

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) SetAddVolumes(v bool)`

SetAddVolumes sets AddVolumes field to given value.

### HasAddVolumes

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) HasAddVolumes() bool`

HasAddVolumes returns a boolean if a field has been set.

### GetMemoryOptionSource

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetMemoryOptionSource() string`

GetMemoryOptionSource returns the MemoryOptionSource field if non-nil, zero value otherwise.

### GetMemoryOptionSourceOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetMemoryOptionSourceOk() (*string, bool)`

GetMemoryOptionSourceOk returns a tuple with the MemoryOptionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryOptionSource

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) SetMemoryOptionSource(v string)`

SetMemoryOptionSource sets MemoryOptionSource field to given value.

### HasMemoryOptionSource

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) HasMemoryOptionSource() bool`

HasMemoryOptionSource returns a boolean if a field has been set.

### GetCpuOptionSource

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetCpuOptionSource() string`

GetCpuOptionSource returns the CpuOptionSource field if non-nil, zero value otherwise.

### GetCpuOptionSourceOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetCpuOptionSourceOk() (*string, bool)`

GetCpuOptionSourceOk returns a tuple with the CpuOptionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuOptionSource

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) SetCpuOptionSource(v string)`

SetCpuOptionSource sets CpuOptionSource field to given value.

### HasCpuOptionSource

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) HasCpuOptionSource() bool`

HasCpuOptionSource returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetRegionCode

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### GetVisibility

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetEditable

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetProvisionType

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetProvisionType() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionProvisionType`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetProvisionTypeOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionProvisionType, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) SetProvisionType(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionProvisionType)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetTenants

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetTenants() string`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetTenantsOk() (*string, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) SetTenants(v string)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetPriceSets

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetPriceSets() []ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionPriceSetsInner`

GetPriceSets returns the PriceSets field if non-nil, zero value otherwise.

### GetPriceSetsOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetPriceSetsOk() (*[]ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionPriceSetsInner, bool)`

GetPriceSetsOk returns a tuple with the PriceSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSets

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) SetPriceSets(v []ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionPriceSetsInner)`

SetPriceSets sets PriceSets field to given value.

### HasPriceSets

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) HasPriceSets() bool`

HasPriceSets returns a boolean if a field has been set.

### GetConfig

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetConfig() ListServicePlans200ResponseAllOfServicePlansInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetConfigOk() (*ListServicePlans200ResponseAllOfServicePlansInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) SetConfig(v ListServicePlans200ResponseAllOfServicePlansInnerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetZones

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetZones() []ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetZonesOk() (*[]ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) SetZones(v []ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetZones sets Zones field to given value.

### HasZones

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetPermissions

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetPermissions() ListServicePlans200ResponseAllOfServicePlansInnerPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) GetPermissionsOk() (*ListServicePlans200ResponseAllOfServicePlansInnerPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) SetPermissions(v ListServicePlans200ResponseAllOfServicePlansInnerPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *ListServicePlans200ResponseAllOfServicePlansInner) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


