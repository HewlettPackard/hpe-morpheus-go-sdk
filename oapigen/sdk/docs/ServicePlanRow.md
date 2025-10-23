# ServicePlanRow

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Active** | Pointer to **bool** |  | [optional] 
**SortOrder** | Pointer to **int64** |  | [optional] 
**RegionCode** | Pointer to **NullableString** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**MaxStorage** | Pointer to **int64** |  | [optional] 
**MaxMemory** | Pointer to **int64** |  | [optional] 
**MaxCpu** | Pointer to **NullableInt64** |  | [optional] 
**MaxGpus** | Pointer to **NullableInt64** |  | [optional] 
**MaxCores** | Pointer to **NullableInt64** |  | [optional] 
**MaxDisks** | Pointer to **NullableInt64** |  | [optional] 
**CoresPerSocket** | Pointer to **NullableInt64** |  | [optional] 
**CustomCpu** | Pointer to **bool** |  | [optional] 
**CustomCores** | Pointer to **bool** |  | [optional] 
**CustomMaxStorage** | Pointer to **NullableBool** |  | [optional] 
**CustomMaxDataStorage** | Pointer to **NullableBool** |  | [optional] 
**CustomMaxMemory** | Pointer to **NullableBool** |  | [optional] 
**AddVolumes** | Pointer to **NullableBool** |  | [optional] 
**MemoryOptionSource** | Pointer to **NullableString** |  | [optional] 
**CpuOptionSource** | Pointer to **NullableString** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Editable** | Pointer to **bool** |  | [optional] 
**ProvisionType** | Pointer to [**ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionProvisionType**](ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionProvisionType.md) |  | [optional] 
**Tenants** | Pointer to **string** |  | [optional] 
**PriceSets** | Pointer to [**[]ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionPriceSetsInner**](ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionPriceSetsInner.md) |  | [optional] 
**Config** | Pointer to [**ListServicePlans200ResponseAllOfServicePlansInnerConfig**](ListServicePlans200ResponseAllOfServicePlansInnerConfig.md) |  | [optional] 
**Zones** | Pointer to [**[]ListBackupSettings200ResponseBackupSettingsDefaultSchedule**](ListBackupSettings200ResponseBackupSettingsDefaultSchedule.md) |  | [optional] 

## Methods

### NewServicePlanRow

`func NewServicePlanRow() *ServicePlanRow`

NewServicePlanRow instantiates a new ServicePlanRow object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePlanRowWithDefaults

`func NewServicePlanRowWithDefaults() *ServicePlanRow`

NewServicePlanRowWithDefaults instantiates a new ServicePlanRow object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ServicePlanRow) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ServicePlanRow) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ServicePlanRow) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ServicePlanRow) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ServicePlanRow) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ServicePlanRow) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ServicePlanRow) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ServicePlanRow) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *ServicePlanRow) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ServicePlanRow) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ServicePlanRow) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ServicePlanRow) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetActive

`func (o *ServicePlanRow) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ServicePlanRow) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ServicePlanRow) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ServicePlanRow) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetSortOrder

`func (o *ServicePlanRow) GetSortOrder() int64`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *ServicePlanRow) GetSortOrderOk() (*int64, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *ServicePlanRow) SetSortOrder(v int64)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *ServicePlanRow) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetRegionCode

`func (o *ServicePlanRow) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *ServicePlanRow) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *ServicePlanRow) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *ServicePlanRow) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### SetRegionCodeNil

`func (o *ServicePlanRow) SetRegionCodeNil(b bool)`

 SetRegionCodeNil sets the value for RegionCode to be an explicit nil

### UnsetRegionCode
`func (o *ServicePlanRow) UnsetRegionCode()`

UnsetRegionCode ensures that no value is present for RegionCode, not even an explicit nil
### GetDescription

`func (o *ServicePlanRow) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ServicePlanRow) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ServicePlanRow) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ServicePlanRow) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMaxStorage

`func (o *ServicePlanRow) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *ServicePlanRow) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *ServicePlanRow) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *ServicePlanRow) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetMaxMemory

`func (o *ServicePlanRow) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *ServicePlanRow) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *ServicePlanRow) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *ServicePlanRow) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxCpu

`func (o *ServicePlanRow) GetMaxCpu() int64`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *ServicePlanRow) GetMaxCpuOk() (*int64, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *ServicePlanRow) SetMaxCpu(v int64)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *ServicePlanRow) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### SetMaxCpuNil

`func (o *ServicePlanRow) SetMaxCpuNil(b bool)`

 SetMaxCpuNil sets the value for MaxCpu to be an explicit nil

### UnsetMaxCpu
`func (o *ServicePlanRow) UnsetMaxCpu()`

UnsetMaxCpu ensures that no value is present for MaxCpu, not even an explicit nil
### GetMaxGpus

`func (o *ServicePlanRow) GetMaxGpus() int64`

GetMaxGpus returns the MaxGpus field if non-nil, zero value otherwise.

### GetMaxGpusOk

`func (o *ServicePlanRow) GetMaxGpusOk() (*int64, bool)`

GetMaxGpusOk returns a tuple with the MaxGpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGpus

`func (o *ServicePlanRow) SetMaxGpus(v int64)`

SetMaxGpus sets MaxGpus field to given value.

### HasMaxGpus

`func (o *ServicePlanRow) HasMaxGpus() bool`

HasMaxGpus returns a boolean if a field has been set.

### SetMaxGpusNil

`func (o *ServicePlanRow) SetMaxGpusNil(b bool)`

 SetMaxGpusNil sets the value for MaxGpus to be an explicit nil

### UnsetMaxGpus
`func (o *ServicePlanRow) UnsetMaxGpus()`

UnsetMaxGpus ensures that no value is present for MaxGpus, not even an explicit nil
### GetMaxCores

`func (o *ServicePlanRow) GetMaxCores() int64`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *ServicePlanRow) GetMaxCoresOk() (*int64, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *ServicePlanRow) SetMaxCores(v int64)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *ServicePlanRow) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### SetMaxCoresNil

`func (o *ServicePlanRow) SetMaxCoresNil(b bool)`

 SetMaxCoresNil sets the value for MaxCores to be an explicit nil

### UnsetMaxCores
`func (o *ServicePlanRow) UnsetMaxCores()`

UnsetMaxCores ensures that no value is present for MaxCores, not even an explicit nil
### GetMaxDisks

`func (o *ServicePlanRow) GetMaxDisks() int64`

GetMaxDisks returns the MaxDisks field if non-nil, zero value otherwise.

### GetMaxDisksOk

`func (o *ServicePlanRow) GetMaxDisksOk() (*int64, bool)`

GetMaxDisksOk returns a tuple with the MaxDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDisks

`func (o *ServicePlanRow) SetMaxDisks(v int64)`

SetMaxDisks sets MaxDisks field to given value.

### HasMaxDisks

`func (o *ServicePlanRow) HasMaxDisks() bool`

HasMaxDisks returns a boolean if a field has been set.

### SetMaxDisksNil

`func (o *ServicePlanRow) SetMaxDisksNil(b bool)`

 SetMaxDisksNil sets the value for MaxDisks to be an explicit nil

### UnsetMaxDisks
`func (o *ServicePlanRow) UnsetMaxDisks()`

UnsetMaxDisks ensures that no value is present for MaxDisks, not even an explicit nil
### GetCoresPerSocket

`func (o *ServicePlanRow) GetCoresPerSocket() int64`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *ServicePlanRow) GetCoresPerSocketOk() (*int64, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *ServicePlanRow) SetCoresPerSocket(v int64)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *ServicePlanRow) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### SetCoresPerSocketNil

`func (o *ServicePlanRow) SetCoresPerSocketNil(b bool)`

 SetCoresPerSocketNil sets the value for CoresPerSocket to be an explicit nil

### UnsetCoresPerSocket
`func (o *ServicePlanRow) UnsetCoresPerSocket()`

UnsetCoresPerSocket ensures that no value is present for CoresPerSocket, not even an explicit nil
### GetCustomCpu

`func (o *ServicePlanRow) GetCustomCpu() bool`

GetCustomCpu returns the CustomCpu field if non-nil, zero value otherwise.

### GetCustomCpuOk

`func (o *ServicePlanRow) GetCustomCpuOk() (*bool, bool)`

GetCustomCpuOk returns a tuple with the CustomCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCpu

`func (o *ServicePlanRow) SetCustomCpu(v bool)`

SetCustomCpu sets CustomCpu field to given value.

### HasCustomCpu

`func (o *ServicePlanRow) HasCustomCpu() bool`

HasCustomCpu returns a boolean if a field has been set.

### GetCustomCores

`func (o *ServicePlanRow) GetCustomCores() bool`

GetCustomCores returns the CustomCores field if non-nil, zero value otherwise.

### GetCustomCoresOk

`func (o *ServicePlanRow) GetCustomCoresOk() (*bool, bool)`

GetCustomCoresOk returns a tuple with the CustomCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCores

`func (o *ServicePlanRow) SetCustomCores(v bool)`

SetCustomCores sets CustomCores field to given value.

### HasCustomCores

`func (o *ServicePlanRow) HasCustomCores() bool`

HasCustomCores returns a boolean if a field has been set.

### GetCustomMaxStorage

`func (o *ServicePlanRow) GetCustomMaxStorage() bool`

GetCustomMaxStorage returns the CustomMaxStorage field if non-nil, zero value otherwise.

### GetCustomMaxStorageOk

`func (o *ServicePlanRow) GetCustomMaxStorageOk() (*bool, bool)`

GetCustomMaxStorageOk returns a tuple with the CustomMaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxStorage

`func (o *ServicePlanRow) SetCustomMaxStorage(v bool)`

SetCustomMaxStorage sets CustomMaxStorage field to given value.

### HasCustomMaxStorage

`func (o *ServicePlanRow) HasCustomMaxStorage() bool`

HasCustomMaxStorage returns a boolean if a field has been set.

### SetCustomMaxStorageNil

`func (o *ServicePlanRow) SetCustomMaxStorageNil(b bool)`

 SetCustomMaxStorageNil sets the value for CustomMaxStorage to be an explicit nil

### UnsetCustomMaxStorage
`func (o *ServicePlanRow) UnsetCustomMaxStorage()`

UnsetCustomMaxStorage ensures that no value is present for CustomMaxStorage, not even an explicit nil
### GetCustomMaxDataStorage

`func (o *ServicePlanRow) GetCustomMaxDataStorage() bool`

GetCustomMaxDataStorage returns the CustomMaxDataStorage field if non-nil, zero value otherwise.

### GetCustomMaxDataStorageOk

`func (o *ServicePlanRow) GetCustomMaxDataStorageOk() (*bool, bool)`

GetCustomMaxDataStorageOk returns a tuple with the CustomMaxDataStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxDataStorage

`func (o *ServicePlanRow) SetCustomMaxDataStorage(v bool)`

SetCustomMaxDataStorage sets CustomMaxDataStorage field to given value.

### HasCustomMaxDataStorage

`func (o *ServicePlanRow) HasCustomMaxDataStorage() bool`

HasCustomMaxDataStorage returns a boolean if a field has been set.

### SetCustomMaxDataStorageNil

`func (o *ServicePlanRow) SetCustomMaxDataStorageNil(b bool)`

 SetCustomMaxDataStorageNil sets the value for CustomMaxDataStorage to be an explicit nil

### UnsetCustomMaxDataStorage
`func (o *ServicePlanRow) UnsetCustomMaxDataStorage()`

UnsetCustomMaxDataStorage ensures that no value is present for CustomMaxDataStorage, not even an explicit nil
### GetCustomMaxMemory

`func (o *ServicePlanRow) GetCustomMaxMemory() bool`

GetCustomMaxMemory returns the CustomMaxMemory field if non-nil, zero value otherwise.

### GetCustomMaxMemoryOk

`func (o *ServicePlanRow) GetCustomMaxMemoryOk() (*bool, bool)`

GetCustomMaxMemoryOk returns a tuple with the CustomMaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxMemory

`func (o *ServicePlanRow) SetCustomMaxMemory(v bool)`

SetCustomMaxMemory sets CustomMaxMemory field to given value.

### HasCustomMaxMemory

`func (o *ServicePlanRow) HasCustomMaxMemory() bool`

HasCustomMaxMemory returns a boolean if a field has been set.

### SetCustomMaxMemoryNil

`func (o *ServicePlanRow) SetCustomMaxMemoryNil(b bool)`

 SetCustomMaxMemoryNil sets the value for CustomMaxMemory to be an explicit nil

### UnsetCustomMaxMemory
`func (o *ServicePlanRow) UnsetCustomMaxMemory()`

UnsetCustomMaxMemory ensures that no value is present for CustomMaxMemory, not even an explicit nil
### GetAddVolumes

`func (o *ServicePlanRow) GetAddVolumes() bool`

GetAddVolumes returns the AddVolumes field if non-nil, zero value otherwise.

### GetAddVolumesOk

`func (o *ServicePlanRow) GetAddVolumesOk() (*bool, bool)`

GetAddVolumesOk returns a tuple with the AddVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddVolumes

`func (o *ServicePlanRow) SetAddVolumes(v bool)`

SetAddVolumes sets AddVolumes field to given value.

### HasAddVolumes

`func (o *ServicePlanRow) HasAddVolumes() bool`

HasAddVolumes returns a boolean if a field has been set.

### SetAddVolumesNil

`func (o *ServicePlanRow) SetAddVolumesNil(b bool)`

 SetAddVolumesNil sets the value for AddVolumes to be an explicit nil

### UnsetAddVolumes
`func (o *ServicePlanRow) UnsetAddVolumes()`

UnsetAddVolumes ensures that no value is present for AddVolumes, not even an explicit nil
### GetMemoryOptionSource

`func (o *ServicePlanRow) GetMemoryOptionSource() string`

GetMemoryOptionSource returns the MemoryOptionSource field if non-nil, zero value otherwise.

### GetMemoryOptionSourceOk

`func (o *ServicePlanRow) GetMemoryOptionSourceOk() (*string, bool)`

GetMemoryOptionSourceOk returns a tuple with the MemoryOptionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryOptionSource

`func (o *ServicePlanRow) SetMemoryOptionSource(v string)`

SetMemoryOptionSource sets MemoryOptionSource field to given value.

### HasMemoryOptionSource

`func (o *ServicePlanRow) HasMemoryOptionSource() bool`

HasMemoryOptionSource returns a boolean if a field has been set.

### SetMemoryOptionSourceNil

`func (o *ServicePlanRow) SetMemoryOptionSourceNil(b bool)`

 SetMemoryOptionSourceNil sets the value for MemoryOptionSource to be an explicit nil

### UnsetMemoryOptionSource
`func (o *ServicePlanRow) UnsetMemoryOptionSource()`

UnsetMemoryOptionSource ensures that no value is present for MemoryOptionSource, not even an explicit nil
### GetCpuOptionSource

`func (o *ServicePlanRow) GetCpuOptionSource() string`

GetCpuOptionSource returns the CpuOptionSource field if non-nil, zero value otherwise.

### GetCpuOptionSourceOk

`func (o *ServicePlanRow) GetCpuOptionSourceOk() (*string, bool)`

GetCpuOptionSourceOk returns a tuple with the CpuOptionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuOptionSource

`func (o *ServicePlanRow) SetCpuOptionSource(v string)`

SetCpuOptionSource sets CpuOptionSource field to given value.

### HasCpuOptionSource

`func (o *ServicePlanRow) HasCpuOptionSource() bool`

HasCpuOptionSource returns a boolean if a field has been set.

### SetCpuOptionSourceNil

`func (o *ServicePlanRow) SetCpuOptionSourceNil(b bool)`

 SetCpuOptionSourceNil sets the value for CpuOptionSource to be an explicit nil

### UnsetCpuOptionSource
`func (o *ServicePlanRow) UnsetCpuOptionSource()`

UnsetCpuOptionSource ensures that no value is present for CpuOptionSource, not even an explicit nil
### GetDateCreated

`func (o *ServicePlanRow) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ServicePlanRow) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ServicePlanRow) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ServicePlanRow) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ServicePlanRow) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ServicePlanRow) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ServicePlanRow) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ServicePlanRow) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetVisibility

`func (o *ServicePlanRow) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *ServicePlanRow) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *ServicePlanRow) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *ServicePlanRow) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetEditable

`func (o *ServicePlanRow) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *ServicePlanRow) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *ServicePlanRow) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *ServicePlanRow) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetProvisionType

`func (o *ServicePlanRow) GetProvisionType() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionProvisionType`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *ServicePlanRow) GetProvisionTypeOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionProvisionType, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *ServicePlanRow) SetProvisionType(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionProvisionType)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *ServicePlanRow) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetTenants

`func (o *ServicePlanRow) GetTenants() string`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *ServicePlanRow) GetTenantsOk() (*string, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *ServicePlanRow) SetTenants(v string)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *ServicePlanRow) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetPriceSets

`func (o *ServicePlanRow) GetPriceSets() []ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionPriceSetsInner`

GetPriceSets returns the PriceSets field if non-nil, zero value otherwise.

### GetPriceSetsOk

`func (o *ServicePlanRow) GetPriceSetsOk() (*[]ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionPriceSetsInner, bool)`

GetPriceSetsOk returns a tuple with the PriceSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSets

`func (o *ServicePlanRow) SetPriceSets(v []ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfPlanBeforeActionPriceSetsInner)`

SetPriceSets sets PriceSets field to given value.

### HasPriceSets

`func (o *ServicePlanRow) HasPriceSets() bool`

HasPriceSets returns a boolean if a field has been set.

### SetPriceSetsNil

`func (o *ServicePlanRow) SetPriceSetsNil(b bool)`

 SetPriceSetsNil sets the value for PriceSets to be an explicit nil

### UnsetPriceSets
`func (o *ServicePlanRow) UnsetPriceSets()`

UnsetPriceSets ensures that no value is present for PriceSets, not even an explicit nil
### GetConfig

`func (o *ServicePlanRow) GetConfig() ListServicePlans200ResponseAllOfServicePlansInnerConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ServicePlanRow) GetConfigOk() (*ListServicePlans200ResponseAllOfServicePlansInnerConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ServicePlanRow) SetConfig(v ListServicePlans200ResponseAllOfServicePlansInnerConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ServicePlanRow) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetZones

`func (o *ServicePlanRow) GetZones() []ListBackupSettings200ResponseBackupSettingsDefaultSchedule`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *ServicePlanRow) GetZonesOk() (*[]ListBackupSettings200ResponseBackupSettingsDefaultSchedule, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *ServicePlanRow) SetZones(v []ListBackupSettings200ResponseBackupSettingsDefaultSchedule)`

SetZones sets Zones field to given value.

### HasZones

`func (o *ServicePlanRow) HasZones() bool`

HasZones returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


