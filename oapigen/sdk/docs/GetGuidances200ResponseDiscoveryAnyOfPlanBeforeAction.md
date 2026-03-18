# GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction

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
**ProvisionType** | Pointer to [**GetGuidances200ResponseDiscoveryAnyOfPlanBeforeActionProvisionType**](GetGuidances200ResponseDiscoveryAnyOfPlanBeforeActionProvisionType.md) |  | [optional] 
**Tenants** | Pointer to **string** |  | [optional] 
**PriceSets** | Pointer to [**[]GetGuidances200ResponseDiscoveryAnyOfPlanBeforeActionPriceSetsInner**](GetGuidances200ResponseDiscoveryAnyOfPlanBeforeActionPriceSetsInner.md) |  | [optional] 
**Config** | Pointer to [**GetGuidances200ResponseDiscoveryAnyOfPlanBeforeActionConfig**](GetGuidances200ResponseDiscoveryAnyOfPlanBeforeActionConfig.md) |  | [optional] 

## Methods

### NewGetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction

`func NewGetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction() *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction`

NewGetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction instantiates a new GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGuidances200ResponseDiscoveryAnyOfPlanBeforeActionWithDefaults

`func NewGetGuidances200ResponseDiscoveryAnyOfPlanBeforeActionWithDefaults() *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction`

NewGetGuidances200ResponseDiscoveryAnyOfPlanBeforeActionWithDefaults instantiates a new GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetActive

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetSortOrder

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetSortOrder() int64`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetSortOrderOk() (*int64, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) SetSortOrder(v int64)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetDescription

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetMaxStorage

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetMaxStorage() int64`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetMaxStorageOk() (*int64, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) SetMaxStorage(v int64)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetMaxMemory

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMaxCpu

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetMaxCpu() string`

GetMaxCpu returns the MaxCpu field if non-nil, zero value otherwise.

### GetMaxCpuOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetMaxCpuOk() (*string, bool)`

GetMaxCpuOk returns a tuple with the MaxCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCpu

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) SetMaxCpu(v string)`

SetMaxCpu sets MaxCpu field to given value.

### HasMaxCpu

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) HasMaxCpu() bool`

HasMaxCpu returns a boolean if a field has been set.

### SetMaxCpuNil

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) SetMaxCpuNil(b bool)`

 SetMaxCpuNil sets the value for MaxCpu to be an explicit nil

### UnsetMaxCpu
`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) UnsetMaxCpu()`

UnsetMaxCpu ensures that no value is present for MaxCpu, not even an explicit nil
### GetMaxCores

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetMaxCores() int64`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetMaxCoresOk() (*int64, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) SetMaxCores(v int64)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetMaxDisks

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetMaxDisks() string`

GetMaxDisks returns the MaxDisks field if non-nil, zero value otherwise.

### GetMaxDisksOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetMaxDisksOk() (*string, bool)`

GetMaxDisksOk returns a tuple with the MaxDisks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxDisks

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) SetMaxDisks(v string)`

SetMaxDisks sets MaxDisks field to given value.

### HasMaxDisks

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) HasMaxDisks() bool`

HasMaxDisks returns a boolean if a field has been set.

### SetMaxDisksNil

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) SetMaxDisksNil(b bool)`

 SetMaxDisksNil sets the value for MaxDisks to be an explicit nil

### UnsetMaxDisks
`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) UnsetMaxDisks()`

UnsetMaxDisks ensures that no value is present for MaxDisks, not even an explicit nil
### GetCoresPerSocket

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetCoresPerSocket() int64`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetCoresPerSocketOk() (*int64, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) SetCoresPerSocket(v int64)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### GetCustomCpu

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetCustomCpu() bool`

GetCustomCpu returns the CustomCpu field if non-nil, zero value otherwise.

### GetCustomCpuOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetCustomCpuOk() (*bool, bool)`

GetCustomCpuOk returns a tuple with the CustomCpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCpu

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) SetCustomCpu(v bool)`

SetCustomCpu sets CustomCpu field to given value.

### HasCustomCpu

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) HasCustomCpu() bool`

HasCustomCpu returns a boolean if a field has been set.

### GetCustomCores

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetCustomCores() bool`

GetCustomCores returns the CustomCores field if non-nil, zero value otherwise.

### GetCustomCoresOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetCustomCoresOk() (*bool, bool)`

GetCustomCoresOk returns a tuple with the CustomCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomCores

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) SetCustomCores(v bool)`

SetCustomCores sets CustomCores field to given value.

### HasCustomCores

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) HasCustomCores() bool`

HasCustomCores returns a boolean if a field has been set.

### GetCustomMaxStorage

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetCustomMaxStorage() bool`

GetCustomMaxStorage returns the CustomMaxStorage field if non-nil, zero value otherwise.

### GetCustomMaxStorageOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetCustomMaxStorageOk() (*bool, bool)`

GetCustomMaxStorageOk returns a tuple with the CustomMaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxStorage

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) SetCustomMaxStorage(v bool)`

SetCustomMaxStorage sets CustomMaxStorage field to given value.

### HasCustomMaxStorage

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) HasCustomMaxStorage() bool`

HasCustomMaxStorage returns a boolean if a field has been set.

### GetCustomMaxDataStorage

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetCustomMaxDataStorage() bool`

GetCustomMaxDataStorage returns the CustomMaxDataStorage field if non-nil, zero value otherwise.

### GetCustomMaxDataStorageOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetCustomMaxDataStorageOk() (*bool, bool)`

GetCustomMaxDataStorageOk returns a tuple with the CustomMaxDataStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxDataStorage

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) SetCustomMaxDataStorage(v bool)`

SetCustomMaxDataStorage sets CustomMaxDataStorage field to given value.

### HasCustomMaxDataStorage

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) HasCustomMaxDataStorage() bool`

HasCustomMaxDataStorage returns a boolean if a field has been set.

### GetCustomMaxMemory

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetCustomMaxMemory() bool`

GetCustomMaxMemory returns the CustomMaxMemory field if non-nil, zero value otherwise.

### GetCustomMaxMemoryOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetCustomMaxMemoryOk() (*bool, bool)`

GetCustomMaxMemoryOk returns a tuple with the CustomMaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomMaxMemory

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) SetCustomMaxMemory(v bool)`

SetCustomMaxMemory sets CustomMaxMemory field to given value.

### HasCustomMaxMemory

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) HasCustomMaxMemory() bool`

HasCustomMaxMemory returns a boolean if a field has been set.

### GetAddVolumes

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetAddVolumes() bool`

GetAddVolumes returns the AddVolumes field if non-nil, zero value otherwise.

### GetAddVolumesOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetAddVolumesOk() (*bool, bool)`

GetAddVolumesOk returns a tuple with the AddVolumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddVolumes

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) SetAddVolumes(v bool)`

SetAddVolumes sets AddVolumes field to given value.

### HasAddVolumes

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) HasAddVolumes() bool`

HasAddVolumes returns a boolean if a field has been set.

### GetMemoryOptionSource

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetMemoryOptionSource() string`

GetMemoryOptionSource returns the MemoryOptionSource field if non-nil, zero value otherwise.

### GetMemoryOptionSourceOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetMemoryOptionSourceOk() (*string, bool)`

GetMemoryOptionSourceOk returns a tuple with the MemoryOptionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryOptionSource

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) SetMemoryOptionSource(v string)`

SetMemoryOptionSource sets MemoryOptionSource field to given value.

### HasMemoryOptionSource

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) HasMemoryOptionSource() bool`

HasMemoryOptionSource returns a boolean if a field has been set.

### SetMemoryOptionSourceNil

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) SetMemoryOptionSourceNil(b bool)`

 SetMemoryOptionSourceNil sets the value for MemoryOptionSource to be an explicit nil

### UnsetMemoryOptionSource
`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) UnsetMemoryOptionSource()`

UnsetMemoryOptionSource ensures that no value is present for MemoryOptionSource, not even an explicit nil
### GetCpuOptionSource

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetCpuOptionSource() string`

GetCpuOptionSource returns the CpuOptionSource field if non-nil, zero value otherwise.

### GetCpuOptionSourceOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetCpuOptionSourceOk() (*string, bool)`

GetCpuOptionSourceOk returns a tuple with the CpuOptionSource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuOptionSource

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) SetCpuOptionSource(v string)`

SetCpuOptionSource sets CpuOptionSource field to given value.

### HasCpuOptionSource

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) HasCpuOptionSource() bool`

HasCpuOptionSource returns a boolean if a field has been set.

### SetCpuOptionSourceNil

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) SetCpuOptionSourceNil(b bool)`

 SetCpuOptionSourceNil sets the value for CpuOptionSource to be an explicit nil

### UnsetCpuOptionSource
`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) UnsetCpuOptionSource()`

UnsetCpuOptionSource ensures that no value is present for CpuOptionSource, not even an explicit nil
### GetDateCreated

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.

### GetRegionCode

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetRegionCode() string`

GetRegionCode returns the RegionCode field if non-nil, zero value otherwise.

### GetRegionCodeOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetRegionCodeOk() (*string, bool)`

GetRegionCodeOk returns a tuple with the RegionCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegionCode

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) SetRegionCode(v string)`

SetRegionCode sets RegionCode field to given value.

### HasRegionCode

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) HasRegionCode() bool`

HasRegionCode returns a boolean if a field has been set.

### SetRegionCodeNil

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) SetRegionCodeNil(b bool)`

 SetRegionCodeNil sets the value for RegionCode to be an explicit nil

### UnsetRegionCode
`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) UnsetRegionCode()`

UnsetRegionCode ensures that no value is present for RegionCode, not even an explicit nil
### GetVisibility

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetEditable

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetProvisionType

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetProvisionType() GetGuidances200ResponseDiscoveryAnyOfPlanBeforeActionProvisionType`

GetProvisionType returns the ProvisionType field if non-nil, zero value otherwise.

### GetProvisionTypeOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetProvisionTypeOk() (*GetGuidances200ResponseDiscoveryAnyOfPlanBeforeActionProvisionType, bool)`

GetProvisionTypeOk returns a tuple with the ProvisionType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisionType

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) SetProvisionType(v GetGuidances200ResponseDiscoveryAnyOfPlanBeforeActionProvisionType)`

SetProvisionType sets ProvisionType field to given value.

### HasProvisionType

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) HasProvisionType() bool`

HasProvisionType returns a boolean if a field has been set.

### GetTenants

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetTenants() string`

GetTenants returns the Tenants field if non-nil, zero value otherwise.

### GetTenantsOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetTenantsOk() (*string, bool)`

GetTenantsOk returns a tuple with the Tenants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenants

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) SetTenants(v string)`

SetTenants sets Tenants field to given value.

### HasTenants

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) HasTenants() bool`

HasTenants returns a boolean if a field has been set.

### GetPriceSets

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetPriceSets() []GetGuidances200ResponseDiscoveryAnyOfPlanBeforeActionPriceSetsInner`

GetPriceSets returns the PriceSets field if non-nil, zero value otherwise.

### GetPriceSetsOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetPriceSetsOk() (*[]GetGuidances200ResponseDiscoveryAnyOfPlanBeforeActionPriceSetsInner, bool)`

GetPriceSetsOk returns a tuple with the PriceSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriceSets

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) SetPriceSets(v []GetGuidances200ResponseDiscoveryAnyOfPlanBeforeActionPriceSetsInner)`

SetPriceSets sets PriceSets field to given value.

### HasPriceSets

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) HasPriceSets() bool`

HasPriceSets returns a boolean if a field has been set.

### GetConfig

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetConfig() GetGuidances200ResponseDiscoveryAnyOfPlanBeforeActionConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) GetConfigOk() (*GetGuidances200ResponseDiscoveryAnyOfPlanBeforeActionConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) SetConfig(v GetGuidances200ResponseDiscoveryAnyOfPlanBeforeActionConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetGuidances200ResponseDiscoveryAnyOfPlanBeforeAction) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


