# UpdateServicePlans200ResponseAllOfServicePlan

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
**RegionCode** | Pointer to **NullableString** |  | [optional] 
**Visibility** | Pointer to **string** |  | [optional] 
**Editable** | Pointer to **bool** |  | [optional] 
**ProvisionType** | Pointer to [**UpdateServicePlans200ResponseAllOfServicePlanProvisionType**](UpdateServicePlans200ResponseAllOfServicePlanProvisionType.md) |  | [optional] 
**Tenants** | Pointer to **string** |  | [optional] 
**PriceSets** | Pointer to [**[]UpdateServicePlans200ResponseAllOfServicePlanPriceSetsInner**](UpdateServicePlans200ResponseAllOfServicePlanPriceSetsInner.md) |  | [optional] 
**Config** | Pointer to [**UpdateServicePlans200ResponseAllOfServicePlanConfig**](UpdateServicePlans200ResponseAllOfServicePlanConfig.md) |  | [optional] 
**Zones** | Pointer to [**[]UpdateServicePlans200ResponseAllOfServicePlanZonesInner**](UpdateServicePlans200ResponseAllOfServicePlanZonesInner.md) |  | [optional] 
**Permissions** | Pointer to [**UpdateServicePlans200ResponseAllOfServicePlanPermissions**](UpdateServicePlans200ResponseAllOfServicePlanPermissions.md) |  | [optional] 

## Methods

### NewUpdateServicePlans200ResponseAllOfServicePlan

`func NewUpdateServicePlans200ResponseAllOfServicePlan() *UpdateServicePlans200ResponseAllOfServicePlan`

NewUpdateServicePlans200ResponseAllOfServicePlan instantiates a new UpdateServicePlans200ResponseAllOfServicePlan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServicePlans200ResponseAllOfServicePlanWithDefaults

`func NewUpdateServicePlans200ResponseAllOfServicePlanWithDefaults() *UpdateServicePlans200ResponseAllOfServicePlan`

NewUpdateServicePlans200ResponseAllOfServicePlanWithDefaults instantiates a new UpdateServicePlans200ResponseAllOfServicePlan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetActive

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetSortOrder

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetSortOrder() int64`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetSortOrderOk() (*int64, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetSortOrder(v int64)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetMaxStorage

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetMaxMemory

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxCpu

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetMaxCpu() int64`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetMaxCpuOk() (*int64, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetMaxCpu(v int64)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### SetMaxCpuNil

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetMaxCpuNil(b bool)`

 SetMaxCpuNil sets the value for MaxCpu to be an explicit nil

### UnsetMaxCpu
`func (o *UpdateServicePlans200ResponseAllOfServicePlan) UnsetMaxCpu()`

UnsetMaxCpu ensures that no value is present for MaxCpu, not even an explicit nil
### GetMaxGpus

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetMaxGpus() int64`

GetMaxGpus returns the MaxGpus field if non-nil, zero value otherwise.

### GetMaxGpusOk

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetMaxGpusOk() (*int64, bool)`

GetMaxGpusOk returns a tuple with the MaxGpus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxGpus

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetMaxGpus(v int64)`

SetMaxGpus sets MaxGpus field to given value.

### HasMaxGpus

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) HasMaxGpus() bool`

HasMaxGpus returns a boolean if a field has been set.

### SetMaxGpusNil

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetMaxGpusNil(b bool)`

 SetMaxGpusNil sets the value for MaxGpus to be an explicit nil

### UnsetMaxGpus
`func (o *UpdateServicePlans200ResponseAllOfServicePlan) UnsetMaxGpus()`

UnsetMaxGpus ensures that no value is present for MaxGpus, not even an explicit nil
### GetMaxCores

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetMaxCores() int64`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetMaxCoresOk() (*int64, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetMaxCores(v int64)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### SetMaxCoresNil

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetMaxCoresNil(b bool)`

 SetMaxCoresNil sets the value for MaxCores to be an explicit nil

### UnsetMaxCores
`func (o *UpdateServicePlans200ResponseAllOfServicePlan) UnsetMaxCores()`

UnsetMaxCores ensures that no value is present for MaxCores, not even an explicit nil
### GetMaxDisks

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetMaxDisks() int64`

GetMaxDisks returns the MaxDisks field if non-nil, zero value otherwise.

### GetMaxDisksOk

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetMaxDisksOk() (*int64, bool)`

GetMaxDisksOk returns a tuple with the MaxDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDisks

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetMaxDisks(v int64)`

SetMaxDisks sets MaxDisks field to given value.

### HasMaxDisks

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) HasMaxDisks() bool`

HasMaxDisks returns a boolean if a field has been set.

### SetMaxDisksNil

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetMaxDisksNil(b bool)`

 SetMaxDisksNil sets the value for MaxDisks to be an explicit nil

### UnsetMaxDisks
`func (o *UpdateServicePlans200ResponseAllOfServicePlan) UnsetMaxDisks()`

UnsetMaxDisks ensures that no value is present for MaxDisks, not even an explicit nil
### GetCoresPerSocket

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetCoresPerSocket() int64`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetCoresPerSocketOk() (*int64, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetCoresPerSocket(v int64)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### SetCoresPerSocketNil

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetCoresPerSocketNil(b bool)`

 SetCoresPerSocketNil sets the value for CoresPerSocket to be an explicit nil

### UnsetCoresPerSocket
`func (o *UpdateServicePlans200ResponseAllOfServicePlan) UnsetCoresPerSocket()`

UnsetCoresPerSocket ensures that no value is present for CoresPerSocket, not even an explicit nil
### GetCustomCpu

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetCustomCpu() bool`

GetCustomCpu returns the CustomCpu field if non-nil, zero value otherwise.

### GetCustomCpuOk

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetCustomCpuOk() (*bool, bool)`

GetCustomCpuOk returns a tuple with the CustomCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCpu

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetCustomCpu(v bool)`

SetCustomCpu sets CustomCpu field to given value.

### HasCustomCpu

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) HasCustomCpu() bool`

HasCustomCpu returns a boolean if a field has been set.

### GetCustomCores

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetCustomCores() bool`

GetCustomCores returns the CustomCores field if non-nil, zero value otherwise.

### GetCustomCoresOk

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetCustomCoresOk() (*bool, bool)`

GetCustomCoresOk returns a tuple with the CustomCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCores

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetCustomCores(v bool)`

SetCustomCores sets CustomCores field to given value.

### HasCustomCores

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) HasCustomCores() bool`

HasCustomCores returns a boolean if a field has been set.

### GetCustomMaxStorage

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetCustomMaxStorage() bool`

GetCustomMaxStorage returns the CustomMaxStorage field if non-nil, zero value otherwise.

### GetCustomMaxStorageOk

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetCustomMaxStorageOk() (*bool, bool)`

GetCustomMaxStorageOk returns a tuple with the CustomMaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxStorage

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetCustomMaxStorage(v bool)`

SetCustomMaxStorage sets CustomMaxStorage field to given value.

### HasCustomMaxStorage

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) HasCustomMaxStorage() bool`

HasCustomMaxStorage returns a boolean if a field has been set.

### SetCustomMaxStorageNil

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetCustomMaxStorageNil(b bool)`

 SetCustomMaxStorageNil sets the value for CustomMaxStorage to be an explicit nil

### UnsetCustomMaxStorage
`func (o *UpdateServicePlans200ResponseAllOfServicePlan) UnsetCustomMaxStorage()`

UnsetCustomMaxStorage ensures that no value is present for CustomMaxStorage, not even an explicit nil
### GetCustomMaxDataStorage

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetCustomMaxDataStorage() bool`

GetCustomMaxDataStorage returns the CustomMaxDataStorage field if non-nil, zero value otherwise.

### GetCustomMaxDataStorageOk

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetCustomMaxDataStorageOk() (*bool, bool)`

GetCustomMaxDataStorageOk returns a tuple with the CustomMaxDataStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxDataStorage

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetCustomMaxDataStorage(v bool)`

SetCustomMaxDataStorage sets CustomMaxDataStorage field to given value.

### HasCustomMaxDataStorage

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) HasCustomMaxDataStorage() bool`

HasCustomMaxDataStorage returns a boolean if a field has been set.

### SetCustomMaxDataStorageNil

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetCustomMaxDataStorageNil(b bool)`

 SetCustomMaxDataStorageNil sets the value for CustomMaxDataStorage to be an explicit nil

### UnsetCustomMaxDataStorage
`func (o *UpdateServicePlans200ResponseAllOfServicePlan) UnsetCustomMaxDataStorage()`

UnsetCustomMaxDataStorage ensures that no value is present for CustomMaxDataStorage, not even an explicit nil
### GetCustomMaxMemory

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetCustomMaxMemory() bool`

GetCustomMaxMemory returns the CustomMaxMemory field if non-nil, zero value otherwise.

### GetCustomMaxMemoryOk

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetCustomMaxMemoryOk() (*bool, bool)`

GetCustomMaxMemoryOk returns a tuple with the CustomMaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxMemory

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetCustomMaxMemory(v bool)`

SetCustomMaxMemory sets CustomMaxMemory field to given value.

### HasCustomMaxMemory

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) HasCustomMaxMemory() bool`

HasCustomMaxMemory returns a boolean if a field has been set.

### SetCustomMaxMemoryNil

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetCustomMaxMemoryNil(b bool)`

 SetCustomMaxMemoryNil sets the value for CustomMaxMemory to be an explicit nil

### UnsetCustomMaxMemory
`func (o *UpdateServicePlans200ResponseAllOfServicePlan) UnsetCustomMaxMemory()`

UnsetCustomMaxMemory ensures that no value is present for CustomMaxMemory, not even an explicit nil
### GetAddVolumes

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetAddVolumes() bool`

GetAddVolumes returns the AddVolumes field if non-nil, zero value otherwise.

### GetAddVolumesOk

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetAddVolumesOk() (*bool, bool)`

GetAddVolumesOk returns a tuple with the AddVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddVolumes

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetAddVolumes(v bool)`

SetAddVolumes sets AddVolumes field to given value.

### HasAddVolumes

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) HasAddVolumes() bool`

HasAddVolumes returns a boolean if a field has been set.

### SetAddVolumesNil

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetAddVolumesNil(b bool)`

 SetAddVolumesNil sets the value for AddVolumes to be an explicit nil

### UnsetAddVolumes
`func (o *UpdateServicePlans200ResponseAllOfServicePlan) UnsetAddVolumes()`

UnsetAddVolumes ensures that no value is present for AddVolumes, not even an explicit nil
### GetMemoryOptionSource

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetMemoryOptionSource() string`

GetMemoryOptionSource returns the MemoryOptionSource field if non-nil, zero value otherwise.

### GetMemoryOptionSourceOk

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetMemoryOptionSourceOk() (*string, bool)`

GetMemoryOptionSourceOk returns a tuple with the MemoryOptionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryOptionSource

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetMemoryOptionSource(v string)`

SetMemoryOptionSource sets MemoryOptionSource field to given value.

### HasMemoryOptionSource

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) HasMemoryOptionSource() bool`

HasMemoryOptionSource returns a boolean if a field has been set.

### SetMemoryOptionSourceNil

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetMemoryOptionSourceNil(b bool)`

 SetMemoryOptionSourceNil sets the value for MemoryOptionSource to be an explicit nil

### UnsetMemoryOptionSource
`func (o *UpdateServicePlans200ResponseAllOfServicePlan) UnsetMemoryOptionSource()`

UnsetMemoryOptionSource ensures that no value is present for MemoryOptionSource, not even an explicit nil
### GetCpuOptionSource

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetCpuOptionSource() string`

GetCpuOptionSource returns the CpuOptionSource field if non-nil, zero value otherwise.

### GetCpuOptionSourceOk

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetCpuOptionSourceOk() (*string, bool)`

GetCpuOptionSourceOk returns a tuple with the CpuOptionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuOptionSource

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetCpuOptionSource(v string)`

SetCpuOptionSource sets CpuOptionSource field to given value.

### HasCpuOptionSource

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) HasCpuOptionSource() bool`

HasCpuOptionSource returns a boolean if a field has been set.

### SetCpuOptionSourceNil

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetCpuOptionSourceNil(b bool)`

 SetCpuOptionSourceNil sets the value for CpuOptionSource to be an explicit nil

### UnsetCpuOptionSource
`func (o *UpdateServicePlans200ResponseAllOfServicePlan) UnsetCpuOptionSource()`

UnsetCpuOptionSource ensures that no value is present for CpuOptionSource, not even an explicit nil
### GetDateCreated

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetRegionCode

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### SetRegionCodeNil

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetRegionCodeNil(b bool)`

 SetRegionCodeNil sets the value for RegionCode to be an explicit nil

### UnsetRegionCode
`func (o *UpdateServicePlans200ResponseAllOfServicePlan) UnsetRegionCode()`

UnsetRegionCode ensures that no value is present for RegionCode, not even an explicit nil
### GetVisibility

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetEditable

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetProvisionType

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetProvisionType() UpdateServicePlans200ResponseAllOfServicePlanProvisionType`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetProvisionTypeOk() (*UpdateServicePlans200ResponseAllOfServicePlanProvisionType, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetProvisionType(v UpdateServicePlans200ResponseAllOfServicePlanProvisionType)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetTenants

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetTenants() string`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetTenantsOk() (*string, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetTenants(v string)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetPriceSets

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetPriceSets() []UpdateServicePlans200ResponseAllOfServicePlanPriceSetsInner`

GetPriceSets returns the PriceSets field if non-nil, zero value otherwise.

### GetPriceSetsOk

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetPriceSetsOk() (*[]UpdateServicePlans200ResponseAllOfServicePlanPriceSetsInner, bool)`

GetPriceSetsOk returns a tuple with the PriceSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSets

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetPriceSets(v []UpdateServicePlans200ResponseAllOfServicePlanPriceSetsInner)`

SetPriceSets sets PriceSets field to given value.

### HasPriceSets

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) HasPriceSets() bool`

HasPriceSets returns a boolean if a field has been set.

### SetPriceSetsNil

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetPriceSetsNil(b bool)`

 SetPriceSetsNil sets the value for PriceSets to be an explicit nil

### UnsetPriceSets
`func (o *UpdateServicePlans200ResponseAllOfServicePlan) UnsetPriceSets()`

UnsetPriceSets ensures that no value is present for PriceSets, not even an explicit nil
### GetConfig

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetConfig() UpdateServicePlans200ResponseAllOfServicePlanConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetConfigOk() (*UpdateServicePlans200ResponseAllOfServicePlanConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetConfig(v UpdateServicePlans200ResponseAllOfServicePlanConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetZones

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetZones() []UpdateServicePlans200ResponseAllOfServicePlanZonesInner`

GetZones returns the Zones field if non-nil, zero value otherwise.

### GetZonesOk

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetZonesOk() (*[]UpdateServicePlans200ResponseAllOfServicePlanZonesInner, bool)`

GetZonesOk returns a tuple with the Zones field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZones

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetZones(v []UpdateServicePlans200ResponseAllOfServicePlanZonesInner)`

SetZones sets Zones field to given value.

### HasZones

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) HasZones() bool`

HasZones returns a boolean if a field has been set.

### GetPermissions

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetPermissions() UpdateServicePlans200ResponseAllOfServicePlanPermissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) GetPermissionsOk() (*UpdateServicePlans200ResponseAllOfServicePlanPermissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) SetPermissions(v UpdateServicePlans200ResponseAllOfServicePlanPermissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *UpdateServicePlans200ResponseAllOfServicePlan) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


