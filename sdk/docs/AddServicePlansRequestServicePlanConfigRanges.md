# AddServicePlansRequestServicePlanConfigRanges

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinStorage** | Pointer to **string** | Custom min storage in GB | [optional] 
**MaxStorage** | Pointer to **string** | Custom max storage in GB | [optional] 
**MinPerDiskSize** | Pointer to **string** | Custom min per disk size in GB | [optional] 
**MaxPerDiskSize** | Pointer to **string** | Custom max per disk size in GB | [optional] 
**MinMemory** | Pointer to **int64** | Custom min memory in bytes | [optional] 
**MaxMemory** | Pointer to **int64** | Custom max memory in bytes | [optional] 
**MinCores** | Pointer to **string** | Custom min cores | [optional] 
**MaxCores** | Pointer to **string** | Custom max cores | [optional] 
**MinSockets** | Pointer to **string** | Custom min sockets | [optional] 
**MaxSockets** | Pointer to **string** | Custom max sockets | [optional] 
**MinCoresPerSocket** | Pointer to **string** | Custom min cores allowed per socket | [optional] 
**MaxCoresPerSocket** | Pointer to **string** | Custom max cores allowed per socket | [optional] 

## Methods

### NewAddServicePlansRequestServicePlanConfigRanges

`func NewAddServicePlansRequestServicePlanConfigRanges() *AddServicePlansRequestServicePlanConfigRanges`

NewAddServicePlansRequestServicePlanConfigRanges instantiates a new AddServicePlansRequestServicePlanConfigRanges object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddServicePlansRequestServicePlanConfigRangesWithDefaults

`func NewAddServicePlansRequestServicePlanConfigRangesWithDefaults() *AddServicePlansRequestServicePlanConfigRanges`

NewAddServicePlansRequestServicePlanConfigRangesWithDefaults instantiates a new AddServicePlansRequestServicePlanConfigRanges object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinStorage

`func (o *AddServicePlansRequestServicePlanConfigRanges) GetMinStorage() string`

GetMinStorage returns the MinStorage field if non-nil, zero value otherwise.

### GetMinStorageOk

`func (o *AddServicePlansRequestServicePlanConfigRanges) GetMinStorageOk() (*string, bool)`

GetMinStorageOk returns a tuple with the MinStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinStorage

`func (o *AddServicePlansRequestServicePlanConfigRanges) SetMinStorage(v string)`

SetMinStorage sets MinStorage field to given value.

### HasMinStorage

`func (o *AddServicePlansRequestServicePlanConfigRanges) HasMinStorage() bool`

HasMinStorage returns a boolean if a field has been set.

### GetMaxStorage

`func (o *AddServicePlansRequestServicePlanConfigRanges) GetMaxStorage() string`

GetMaxStorage returns the MaxStorage field if non-nil, zero value otherwise.

### GetMaxStorageOk

`func (o *AddServicePlansRequestServicePlanConfigRanges) GetMaxStorageOk() (*string, bool)`

GetMaxStorageOk returns a tuple with the MaxStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxStorage

`func (o *AddServicePlansRequestServicePlanConfigRanges) SetMaxStorage(v string)`

SetMaxStorage sets MaxStorage field to given value.

### HasMaxStorage

`func (o *AddServicePlansRequestServicePlanConfigRanges) HasMaxStorage() bool`

HasMaxStorage returns a boolean if a field has been set.

### GetMinPerDiskSize

`func (o *AddServicePlansRequestServicePlanConfigRanges) GetMinPerDiskSize() string`

GetMinPerDiskSize returns the MinPerDiskSize field if non-nil, zero value otherwise.

### GetMinPerDiskSizeOk

`func (o *AddServicePlansRequestServicePlanConfigRanges) GetMinPerDiskSizeOk() (*string, bool)`

GetMinPerDiskSizeOk returns a tuple with the MinPerDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPerDiskSize

`func (o *AddServicePlansRequestServicePlanConfigRanges) SetMinPerDiskSize(v string)`

SetMinPerDiskSize sets MinPerDiskSize field to given value.

### HasMinPerDiskSize

`func (o *AddServicePlansRequestServicePlanConfigRanges) HasMinPerDiskSize() bool`

HasMinPerDiskSize returns a boolean if a field has been set.

### GetMaxPerDiskSize

`func (o *AddServicePlansRequestServicePlanConfigRanges) GetMaxPerDiskSize() string`

GetMaxPerDiskSize returns the MaxPerDiskSize field if non-nil, zero value otherwise.

### GetMaxPerDiskSizeOk

`func (o *AddServicePlansRequestServicePlanConfigRanges) GetMaxPerDiskSizeOk() (*string, bool)`

GetMaxPerDiskSizeOk returns a tuple with the MaxPerDiskSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPerDiskSize

`func (o *AddServicePlansRequestServicePlanConfigRanges) SetMaxPerDiskSize(v string)`

SetMaxPerDiskSize sets MaxPerDiskSize field to given value.

### HasMaxPerDiskSize

`func (o *AddServicePlansRequestServicePlanConfigRanges) HasMaxPerDiskSize() bool`

HasMaxPerDiskSize returns a boolean if a field has been set.

### GetMinMemory

`func (o *AddServicePlansRequestServicePlanConfigRanges) GetMinMemory() int64`

GetMinMemory returns the MinMemory field if non-nil, zero value otherwise.

### GetMinMemoryOk

`func (o *AddServicePlansRequestServicePlanConfigRanges) GetMinMemoryOk() (*int64, bool)`

GetMinMemoryOk returns a tuple with the MinMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinMemory

`func (o *AddServicePlansRequestServicePlanConfigRanges) SetMinMemory(v int64)`

SetMinMemory sets MinMemory field to given value.

### HasMinMemory

`func (o *AddServicePlansRequestServicePlanConfigRanges) HasMinMemory() bool`

HasMinMemory returns a boolean if a field has been set.

### GetMaxMemory

`func (o *AddServicePlansRequestServicePlanConfigRanges) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *AddServicePlansRequestServicePlanConfigRanges) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *AddServicePlansRequestServicePlanConfigRanges) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *AddServicePlansRequestServicePlanConfigRanges) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetMinCores

`func (o *AddServicePlansRequestServicePlanConfigRanges) GetMinCores() string`

GetMinCores returns the MinCores field if non-nil, zero value otherwise.

### GetMinCoresOk

`func (o *AddServicePlansRequestServicePlanConfigRanges) GetMinCoresOk() (*string, bool)`

GetMinCoresOk returns a tuple with the MinCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCores

`func (o *AddServicePlansRequestServicePlanConfigRanges) SetMinCores(v string)`

SetMinCores sets MinCores field to given value.

### HasMinCores

`func (o *AddServicePlansRequestServicePlanConfigRanges) HasMinCores() bool`

HasMinCores returns a boolean if a field has been set.

### GetMaxCores

`func (o *AddServicePlansRequestServicePlanConfigRanges) GetMaxCores() string`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *AddServicePlansRequestServicePlanConfigRanges) GetMaxCoresOk() (*string, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *AddServicePlansRequestServicePlanConfigRanges) SetMaxCores(v string)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *AddServicePlansRequestServicePlanConfigRanges) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetMinSockets

`func (o *AddServicePlansRequestServicePlanConfigRanges) GetMinSockets() string`

GetMinSockets returns the MinSockets field if non-nil, zero value otherwise.

### GetMinSocketsOk

`func (o *AddServicePlansRequestServicePlanConfigRanges) GetMinSocketsOk() (*string, bool)`

GetMinSocketsOk returns a tuple with the MinSockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinSockets

`func (o *AddServicePlansRequestServicePlanConfigRanges) SetMinSockets(v string)`

SetMinSockets sets MinSockets field to given value.

### HasMinSockets

`func (o *AddServicePlansRequestServicePlanConfigRanges) HasMinSockets() bool`

HasMinSockets returns a boolean if a field has been set.

### GetMaxSockets

`func (o *AddServicePlansRequestServicePlanConfigRanges) GetMaxSockets() string`

GetMaxSockets returns the MaxSockets field if non-nil, zero value otherwise.

### GetMaxSocketsOk

`func (o *AddServicePlansRequestServicePlanConfigRanges) GetMaxSocketsOk() (*string, bool)`

GetMaxSocketsOk returns a tuple with the MaxSockets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSockets

`func (o *AddServicePlansRequestServicePlanConfigRanges) SetMaxSockets(v string)`

SetMaxSockets sets MaxSockets field to given value.

### HasMaxSockets

`func (o *AddServicePlansRequestServicePlanConfigRanges) HasMaxSockets() bool`

HasMaxSockets returns a boolean if a field has been set.

### GetMinCoresPerSocket

`func (o *AddServicePlansRequestServicePlanConfigRanges) GetMinCoresPerSocket() string`

GetMinCoresPerSocket returns the MinCoresPerSocket field if non-nil, zero value otherwise.

### GetMinCoresPerSocketOk

`func (o *AddServicePlansRequestServicePlanConfigRanges) GetMinCoresPerSocketOk() (*string, bool)`

GetMinCoresPerSocketOk returns a tuple with the MinCoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinCoresPerSocket

`func (o *AddServicePlansRequestServicePlanConfigRanges) SetMinCoresPerSocket(v string)`

SetMinCoresPerSocket sets MinCoresPerSocket field to given value.

### HasMinCoresPerSocket

`func (o *AddServicePlansRequestServicePlanConfigRanges) HasMinCoresPerSocket() bool`

HasMinCoresPerSocket returns a boolean if a field has been set.

### GetMaxCoresPerSocket

`func (o *AddServicePlansRequestServicePlanConfigRanges) GetMaxCoresPerSocket() string`

GetMaxCoresPerSocket returns the MaxCoresPerSocket field if non-nil, zero value otherwise.

### GetMaxCoresPerSocketOk

`func (o *AddServicePlansRequestServicePlanConfigRanges) GetMaxCoresPerSocketOk() (*string, bool)`

GetMaxCoresPerSocketOk returns a tuple with the MaxCoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCoresPerSocket

`func (o *AddServicePlansRequestServicePlanConfigRanges) SetMaxCoresPerSocket(v string)`

SetMaxCoresPerSocket sets MaxCoresPerSocket field to given value.

### HasMaxCoresPerSocket

`func (o *AddServicePlansRequestServicePlanConfigRanges) HasMaxCoresPerSocket() bool`

HasMaxCoresPerSocket returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


