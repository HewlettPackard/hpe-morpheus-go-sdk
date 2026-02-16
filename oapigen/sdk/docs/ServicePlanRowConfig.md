# ServicePlanRowConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageSizeType** | Pointer to **NullableString** |  | [optional] 
**MemorySizeType** | Pointer to **NullableString** |  | [optional] 
**Ranges** | Pointer to [**ServicePlanRowConfigRanges**](ServicePlanRowConfigRanges.md) |  | [optional] 

## Methods

### NewServicePlanRowConfig

`func NewServicePlanRowConfig() *ServicePlanRowConfig`

NewServicePlanRowConfig instantiates a new ServicePlanRowConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewServicePlanRowConfigWithDefaults

`func NewServicePlanRowConfigWithDefaults() *ServicePlanRowConfig`

NewServicePlanRowConfigWithDefaults instantiates a new ServicePlanRowConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageSizeType

`func (o *ServicePlanRowConfig) GetStorageSizeType() string`

GetStorageSizeType returns the StorageSizeType field if non-nil, zero value otherwise.

### GetStorageSizeTypeOk

`func (o *ServicePlanRowConfig) GetStorageSizeTypeOk() (*string, bool)`

GetStorageSizeTypeOk returns a tuple with the StorageSizeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSizeType

`func (o *ServicePlanRowConfig) SetStorageSizeType(v string)`

SetStorageSizeType sets StorageSizeType field to given value.

### HasStorageSizeType

`func (o *ServicePlanRowConfig) HasStorageSizeType() bool`

HasStorageSizeType returns a boolean if a field has been set.

### SetStorageSizeTypeNil

`func (o *ServicePlanRowConfig) SetStorageSizeTypeNil(b bool)`

 SetStorageSizeTypeNil sets the value for StorageSizeType to be an explicit nil

### UnsetStorageSizeType
`func (o *ServicePlanRowConfig) UnsetStorageSizeType()`

UnsetStorageSizeType ensures that no value is present for StorageSizeType, not even an explicit nil
### GetMemorySizeType

`func (o *ServicePlanRowConfig) GetMemorySizeType() string`

GetMemorySizeType returns the MemorySizeType field if non-nil, zero value otherwise.

### GetMemorySizeTypeOk

`func (o *ServicePlanRowConfig) GetMemorySizeTypeOk() (*string, bool)`

GetMemorySizeTypeOk returns a tuple with the MemorySizeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySizeType

`func (o *ServicePlanRowConfig) SetMemorySizeType(v string)`

SetMemorySizeType sets MemorySizeType field to given value.

### HasMemorySizeType

`func (o *ServicePlanRowConfig) HasMemorySizeType() bool`

HasMemorySizeType returns a boolean if a field has been set.

### SetMemorySizeTypeNil

`func (o *ServicePlanRowConfig) SetMemorySizeTypeNil(b bool)`

 SetMemorySizeTypeNil sets the value for MemorySizeType to be an explicit nil

### UnsetMemorySizeType
`func (o *ServicePlanRowConfig) UnsetMemorySizeType()`

UnsetMemorySizeType ensures that no value is present for MemorySizeType, not even an explicit nil
### GetRanges

`func (o *ServicePlanRowConfig) GetRanges() ServicePlanRowConfigRanges`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *ServicePlanRowConfig) GetRangesOk() (*ServicePlanRowConfigRanges, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *ServicePlanRowConfig) SetRanges(v ServicePlanRowConfigRanges)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *ServicePlanRowConfig) HasRanges() bool`

HasRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


