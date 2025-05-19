# AddServicePlansRequestServicePlanConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StorageSizeType** | Pointer to **string** | Specifies range min / max storage multiplier | [optional] [default to "gb"]
**MemorySizeType** | Pointer to **string** | Specifies range min / max memory multiplier | [optional] [default to "mb"]
**Ranges** | Pointer to [**AddServicePlansRequestServicePlanConfigRanges**](AddServicePlansRequestServicePlanConfigRanges.md) |  | [optional] 

## Methods

### NewAddServicePlansRequestServicePlanConfig

`func NewAddServicePlansRequestServicePlanConfig() *AddServicePlansRequestServicePlanConfig`

NewAddServicePlansRequestServicePlanConfig instantiates a new AddServicePlansRequestServicePlanConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddServicePlansRequestServicePlanConfigWithDefaults

`func NewAddServicePlansRequestServicePlanConfigWithDefaults() *AddServicePlansRequestServicePlanConfig`

NewAddServicePlansRequestServicePlanConfigWithDefaults instantiates a new AddServicePlansRequestServicePlanConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStorageSizeType

`func (o *AddServicePlansRequestServicePlanConfig) GetStorageSizeType() string`

GetStorageSizeType returns the StorageSizeType field if non-nil, zero value otherwise.

### GetStorageSizeTypeOk

`func (o *AddServicePlansRequestServicePlanConfig) GetStorageSizeTypeOk() (*string, bool)`

GetStorageSizeTypeOk returns a tuple with the StorageSizeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStorageSizeType

`func (o *AddServicePlansRequestServicePlanConfig) SetStorageSizeType(v string)`

SetStorageSizeType sets StorageSizeType field to given value.

### HasStorageSizeType

`func (o *AddServicePlansRequestServicePlanConfig) HasStorageSizeType() bool`

HasStorageSizeType returns a boolean if a field has been set.

### GetMemorySizeType

`func (o *AddServicePlansRequestServicePlanConfig) GetMemorySizeType() string`

GetMemorySizeType returns the MemorySizeType field if non-nil, zero value otherwise.

### GetMemorySizeTypeOk

`func (o *AddServicePlansRequestServicePlanConfig) GetMemorySizeTypeOk() (*string, bool)`

GetMemorySizeTypeOk returns a tuple with the MemorySizeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemorySizeType

`func (o *AddServicePlansRequestServicePlanConfig) SetMemorySizeType(v string)`

SetMemorySizeType sets MemorySizeType field to given value.

### HasMemorySizeType

`func (o *AddServicePlansRequestServicePlanConfig) HasMemorySizeType() bool`

HasMemorySizeType returns a boolean if a field has been set.

### GetRanges

`func (o *AddServicePlansRequestServicePlanConfig) GetRanges() AddServicePlansRequestServicePlanConfigRanges`

GetRanges returns the Ranges field if non-nil, zero value otherwise.

### GetRangesOk

`func (o *AddServicePlansRequestServicePlanConfig) GetRangesOk() (*AddServicePlansRequestServicePlanConfigRanges, bool)`

GetRangesOk returns a tuple with the Ranges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRanges

`func (o *AddServicePlansRequestServicePlanConfig) SetRanges(v AddServicePlansRequestServicePlanConfigRanges)`

SetRanges sets Ranges field to given value.

### HasRanges

`func (o *AddServicePlansRequestServicePlanConfig) HasRanges() bool`

HasRanges returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


