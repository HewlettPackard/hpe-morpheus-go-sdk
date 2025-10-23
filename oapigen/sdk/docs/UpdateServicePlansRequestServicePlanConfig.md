# UpdateServicePlansRequestServicePlanConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageSizeType** | Pointer to **string** | Specifies range min / max storage multiplier | [optional] [default to "gb"]
**MemorySizeType** | Pointer to **string** | Specifies range min / max memory multiplier | [optional] [default to "mb"]
**Ranges** | Pointer to [**UpdateServicePlansRequestServicePlanConfigRanges**](UpdateServicePlansRequestServicePlanConfigRanges.md) |  | [optional] 

## Methods

### NewUpdateServicePlansRequestServicePlanConfig

`func NewUpdateServicePlansRequestServicePlanConfig() *UpdateServicePlansRequestServicePlanConfig`

NewUpdateServicePlansRequestServicePlanConfig instantiates a new UpdateServicePlansRequestServicePlanConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServicePlansRequestServicePlanConfigWithDefaults

`func NewUpdateServicePlansRequestServicePlanConfigWithDefaults() *UpdateServicePlansRequestServicePlanConfig`

NewUpdateServicePlansRequestServicePlanConfigWithDefaults instantiates a new UpdateServicePlansRequestServicePlanConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageSizeType

`func (o *UpdateServicePlansRequestServicePlanConfig) GetStorageSizeType() string`

GetStorageSizeType returns the StorageSizeType field if non-nil, zero value otherwise.

### GetStorageSizeTypeOk

`func (o *UpdateServicePlansRequestServicePlanConfig) GetStorageSizeTypeOk() (*string, bool)`

GetStorageSizeTypeOk returns a tuple with the StorageSizeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSizeType

`func (o *UpdateServicePlansRequestServicePlanConfig) SetStorageSizeType(v string)`

SetStorageSizeType sets StorageSizeType field to given value.

### HasStorageSizeType

`func (o *UpdateServicePlansRequestServicePlanConfig) HasStorageSizeType() bool`

HasStorageSizeType returns a boolean if a field has been set.

### GetMemorySizeType

`func (o *UpdateServicePlansRequestServicePlanConfig) GetMemorySizeType() string`

GetMemorySizeType returns the MemorySizeType field if non-nil, zero value otherwise.

### GetMemorySizeTypeOk

`func (o *UpdateServicePlansRequestServicePlanConfig) GetMemorySizeTypeOk() (*string, bool)`

GetMemorySizeTypeOk returns a tuple with the MemorySizeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySizeType

`func (o *UpdateServicePlansRequestServicePlanConfig) SetMemorySizeType(v string)`

SetMemorySizeType sets MemorySizeType field to given value.

### HasMemorySizeType

`func (o *UpdateServicePlansRequestServicePlanConfig) HasMemorySizeType() bool`

HasMemorySizeType returns a boolean if a field has been set.

### GetRanges

`func (o *UpdateServicePlansRequestServicePlanConfig) GetRanges() UpdateServicePlansRequestServicePlanConfigRanges`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *UpdateServicePlansRequestServicePlanConfig) GetRangesOk() (*UpdateServicePlansRequestServicePlanConfigRanges, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *UpdateServicePlansRequestServicePlanConfig) SetRanges(v UpdateServicePlansRequestServicePlanConfigRanges)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *UpdateServicePlansRequestServicePlanConfig) HasRanges() bool`

HasRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


