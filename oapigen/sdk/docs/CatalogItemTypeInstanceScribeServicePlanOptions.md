# CatalogItemTypeInstanceScribeServicePlanOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxCores** | Pointer to **int64** | Core Count | [optional] 
**CoresPerSocket** | Pointer to **int64** | Cores Per Socket | [optional] 
**MaxMemory** | Pointer to **int64** | Memory in bytes For backwards compatability, values less than 1048576 are treated as being in MB and will be converted to bytes | [optional] 

## Methods

### NewCatalogItemTypeInstanceScribeServicePlanOptions

`func NewCatalogItemTypeInstanceScribeServicePlanOptions() *CatalogItemTypeInstanceScribeServicePlanOptions`

NewCatalogItemTypeInstanceScribeServicePlanOptions instantiates a new CatalogItemTypeInstanceScribeServicePlanOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCatalogItemTypeInstanceScribeServicePlanOptionsWithDefaults

`func NewCatalogItemTypeInstanceScribeServicePlanOptionsWithDefaults() *CatalogItemTypeInstanceScribeServicePlanOptions`

NewCatalogItemTypeInstanceScribeServicePlanOptionsWithDefaults instantiates a new CatalogItemTypeInstanceScribeServicePlanOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxCores

`func (o *CatalogItemTypeInstanceScribeServicePlanOptions) GetMaxCores() int64`

GetMaxCores returns the MaxCores field if non-nil, zero value otherwise.

### GetMaxCoresOk

`func (o *CatalogItemTypeInstanceScribeServicePlanOptions) GetMaxCoresOk() (*int64, bool)`

GetMaxCoresOk returns a tuple with the MaxCores field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxCores

`func (o *CatalogItemTypeInstanceScribeServicePlanOptions) SetMaxCores(v int64)`

SetMaxCores sets MaxCores field to given value.

### HasMaxCores

`func (o *CatalogItemTypeInstanceScribeServicePlanOptions) HasMaxCores() bool`

HasMaxCores returns a boolean if a field has been set.

### GetCoresPerSocket

`func (o *CatalogItemTypeInstanceScribeServicePlanOptions) GetCoresPerSocket() int64`

GetCoresPerSocket returns the CoresPerSocket field if non-nil, zero value otherwise.

### GetCoresPerSocketOk

`func (o *CatalogItemTypeInstanceScribeServicePlanOptions) GetCoresPerSocketOk() (*int64, bool)`

GetCoresPerSocketOk returns a tuple with the CoresPerSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCoresPerSocket

`func (o *CatalogItemTypeInstanceScribeServicePlanOptions) SetCoresPerSocket(v int64)`

SetCoresPerSocket sets CoresPerSocket field to given value.

### HasCoresPerSocket

`func (o *CatalogItemTypeInstanceScribeServicePlanOptions) HasCoresPerSocket() bool`

HasCoresPerSocket returns a boolean if a field has been set.

### GetMaxMemory

`func (o *CatalogItemTypeInstanceScribeServicePlanOptions) GetMaxMemory() int64`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *CatalogItemTypeInstanceScribeServicePlanOptions) GetMaxMemoryOk() (*int64, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *CatalogItemTypeInstanceScribeServicePlanOptions) SetMaxMemory(v int64)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *CatalogItemTypeInstanceScribeServicePlanOptions) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


