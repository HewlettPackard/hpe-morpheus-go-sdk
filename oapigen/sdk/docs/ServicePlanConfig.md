# ServicePlanConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageSizeType** | Pointer to **NullableString** |  | [optional] 
**MemorySizeType** | Pointer to **NullableString** |  | [optional] 
**Ranges** | Pointer to [**ServicePlanConfigRanges**](ServicePlanConfigRanges.md) |  | [optional] 

## Methods

### NewServicePlanConfig

`func NewServicePlanConfig() *ServicePlanConfig`

NewServicePlanConfig instantiates a new ServicePlanConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePlanConfigWithDefaults

`func NewServicePlanConfigWithDefaults() *ServicePlanConfig`

NewServicePlanConfigWithDefaults instantiates a new ServicePlanConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageSizeType

`func (o *ServicePlanConfig) GetStorageSizeType() string`

GetStorageSizeType returns the StorageSizeType field if non-nil, zero value otherwise.

### GetStorageSizeTypeOk

`func (o *ServicePlanConfig) GetStorageSizeTypeOk() (*string, bool)`

GetStorageSizeTypeOk returns a tuple with the StorageSizeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSizeType

`func (o *ServicePlanConfig) SetStorageSizeType(v string)`

SetStorageSizeType sets StorageSizeType field to given value.

### HasStorageSizeType

`func (o *ServicePlanConfig) HasStorageSizeType() bool`

HasStorageSizeType returns a boolean if a field has been set.

### SetStorageSizeTypeNil

`func (o *ServicePlanConfig) SetStorageSizeTypeNil(b bool)`

 SetStorageSizeTypeNil sets the value for StorageSizeType to be an explicit nil

### UnsetStorageSizeType
`func (o *ServicePlanConfig) UnsetStorageSizeType()`

UnsetStorageSizeType ensures that no value is present for StorageSizeType, not even an explicit nil
### GetMemorySizeType

`func (o *ServicePlanConfig) GetMemorySizeType() string`

GetMemorySizeType returns the MemorySizeType field if non-nil, zero value otherwise.

### GetMemorySizeTypeOk

`func (o *ServicePlanConfig) GetMemorySizeTypeOk() (*string, bool)`

GetMemorySizeTypeOk returns a tuple with the MemorySizeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySizeType

`func (o *ServicePlanConfig) SetMemorySizeType(v string)`

SetMemorySizeType sets MemorySizeType field to given value.

### HasMemorySizeType

`func (o *ServicePlanConfig) HasMemorySizeType() bool`

HasMemorySizeType returns a boolean if a field has been set.

### SetMemorySizeTypeNil

`func (o *ServicePlanConfig) SetMemorySizeTypeNil(b bool)`

 SetMemorySizeTypeNil sets the value for MemorySizeType to be an explicit nil

### UnsetMemorySizeType
`func (o *ServicePlanConfig) UnsetMemorySizeType()`

UnsetMemorySizeType ensures that no value is present for MemorySizeType, not even an explicit nil
### GetRanges

`func (o *ServicePlanConfig) GetRanges() ServicePlanConfigRanges`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *ServicePlanConfig) GetRangesOk() (*ServicePlanConfigRanges, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *ServicePlanConfig) SetRanges(v ServicePlanConfigRanges)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *ServicePlanConfig) HasRanges() bool`

HasRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


