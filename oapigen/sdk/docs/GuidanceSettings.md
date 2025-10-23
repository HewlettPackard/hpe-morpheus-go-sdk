# GuidanceSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuAvgCutoffPower** | Pointer to **NullableInt32** | Power Shutdown Average CPU (%). Lower limit for average CPU usage | [optional] 
**CpuMaxCutoffPower** | Pointer to **NullableInt32** | Power Shutdown Maximum CPU (%). Lower limit for peak CPU usage | [optional] 
**NetworkCutoffPower** | Pointer to **NullableInt32** | Power Shutdown Network threshold (bytes). Lower limit for average network bandwidth | [optional] 
**CpuUpAvgStandardCutoffRightSize** | Pointer to **NullableInt32** | CPU Up-size Average CPU (%). Upper limit for CPU usage | [optional] 
**CpuUpMaxStandardCutoffRightSize** | Pointer to **NullableInt32** | CPU Up-size Maximum CPU (%). Upper limit for peak CPU usage | [optional] 
**MemoryUpAvgStandardCutoffRightSize** | Pointer to **NullableInt32** | Memory Up-size Minimum Free Memory (%). Lower limit for average free memory usage | [optional] 
**MemoryDownAvgStandardCutoffRightSize** | Pointer to **NullableInt32** | Memory Down-size Maximum Free Memory (%). Upper limit for average free memory | [optional] 
**MemoryDownMaxStandardCutoffRightSize** | Pointer to **NullableInt32** | Memory Down-size Maximum Free Memory (%). Upper limit for peak memory usage | [optional] 

## Methods

### NewGuidanceSettings

`func NewGuidanceSettings() *GuidanceSettings`

NewGuidanceSettings instantiates a new GuidanceSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuidanceSettingsWithDefaults

`func NewGuidanceSettingsWithDefaults() *GuidanceSettings`

NewGuidanceSettingsWithDefaults instantiates a new GuidanceSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuAvgCutoffPower

`func (o *GuidanceSettings) GetCpuAvgCutoffPower() int32`

GetCpuAvgCutoffPower returns the CpuAvgCutoffPower field if non-nil, zero value otherwise.

### GetCpuAvgCutoffPowerOk

`func (o *GuidanceSettings) GetCpuAvgCutoffPowerOk() (*int32, bool)`

GetCpuAvgCutoffPowerOk returns a tuple with the CpuAvgCutoffPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuAvgCutoffPower

`func (o *GuidanceSettings) SetCpuAvgCutoffPower(v int32)`

SetCpuAvgCutoffPower sets CpuAvgCutoffPower field to given value.

### HasCpuAvgCutoffPower

`func (o *GuidanceSettings) HasCpuAvgCutoffPower() bool`

HasCpuAvgCutoffPower returns a boolean if a field has been set.

### SetCpuAvgCutoffPowerNil

`func (o *GuidanceSettings) SetCpuAvgCutoffPowerNil(b bool)`

 SetCpuAvgCutoffPowerNil sets the value for CpuAvgCutoffPower to be an explicit nil

### UnsetCpuAvgCutoffPower
`func (o *GuidanceSettings) UnsetCpuAvgCutoffPower()`

UnsetCpuAvgCutoffPower ensures that no value is present for CpuAvgCutoffPower, not even an explicit nil
### GetCpuMaxCutoffPower

`func (o *GuidanceSettings) GetCpuMaxCutoffPower() int32`

GetCpuMaxCutoffPower returns the CpuMaxCutoffPower field if non-nil, zero value otherwise.

### GetCpuMaxCutoffPowerOk

`func (o *GuidanceSettings) GetCpuMaxCutoffPowerOk() (*int32, bool)`

GetCpuMaxCutoffPowerOk returns a tuple with the CpuMaxCutoffPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuMaxCutoffPower

`func (o *GuidanceSettings) SetCpuMaxCutoffPower(v int32)`

SetCpuMaxCutoffPower sets CpuMaxCutoffPower field to given value.

### HasCpuMaxCutoffPower

`func (o *GuidanceSettings) HasCpuMaxCutoffPower() bool`

HasCpuMaxCutoffPower returns a boolean if a field has been set.

### SetCpuMaxCutoffPowerNil

`func (o *GuidanceSettings) SetCpuMaxCutoffPowerNil(b bool)`

 SetCpuMaxCutoffPowerNil sets the value for CpuMaxCutoffPower to be an explicit nil

### UnsetCpuMaxCutoffPower
`func (o *GuidanceSettings) UnsetCpuMaxCutoffPower()`

UnsetCpuMaxCutoffPower ensures that no value is present for CpuMaxCutoffPower, not even an explicit nil
### GetNetworkCutoffPower

`func (o *GuidanceSettings) GetNetworkCutoffPower() int32`

GetNetworkCutoffPower returns the NetworkCutoffPower field if non-nil, zero value otherwise.

### GetNetworkCutoffPowerOk

`func (o *GuidanceSettings) GetNetworkCutoffPowerOk() (*int32, bool)`

GetNetworkCutoffPowerOk returns a tuple with the NetworkCutoffPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkCutoffPower

`func (o *GuidanceSettings) SetNetworkCutoffPower(v int32)`

SetNetworkCutoffPower sets NetworkCutoffPower field to given value.

### HasNetworkCutoffPower

`func (o *GuidanceSettings) HasNetworkCutoffPower() bool`

HasNetworkCutoffPower returns a boolean if a field has been set.

### SetNetworkCutoffPowerNil

`func (o *GuidanceSettings) SetNetworkCutoffPowerNil(b bool)`

 SetNetworkCutoffPowerNil sets the value for NetworkCutoffPower to be an explicit nil

### UnsetNetworkCutoffPower
`func (o *GuidanceSettings) UnsetNetworkCutoffPower()`

UnsetNetworkCutoffPower ensures that no value is present for NetworkCutoffPower, not even an explicit nil
### GetCpuUpAvgStandardCutoffRightSize

`func (o *GuidanceSettings) GetCpuUpAvgStandardCutoffRightSize() int32`

GetCpuUpAvgStandardCutoffRightSize returns the CpuUpAvgStandardCutoffRightSize field if non-nil, zero value otherwise.

### GetCpuUpAvgStandardCutoffRightSizeOk

`func (o *GuidanceSettings) GetCpuUpAvgStandardCutoffRightSizeOk() (*int32, bool)`

GetCpuUpAvgStandardCutoffRightSizeOk returns a tuple with the CpuUpAvgStandardCutoffRightSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUpAvgStandardCutoffRightSize

`func (o *GuidanceSettings) SetCpuUpAvgStandardCutoffRightSize(v int32)`

SetCpuUpAvgStandardCutoffRightSize sets CpuUpAvgStandardCutoffRightSize field to given value.

### HasCpuUpAvgStandardCutoffRightSize

`func (o *GuidanceSettings) HasCpuUpAvgStandardCutoffRightSize() bool`

HasCpuUpAvgStandardCutoffRightSize returns a boolean if a field has been set.

### SetCpuUpAvgStandardCutoffRightSizeNil

`func (o *GuidanceSettings) SetCpuUpAvgStandardCutoffRightSizeNil(b bool)`

 SetCpuUpAvgStandardCutoffRightSizeNil sets the value for CpuUpAvgStandardCutoffRightSize to be an explicit nil

### UnsetCpuUpAvgStandardCutoffRightSize
`func (o *GuidanceSettings) UnsetCpuUpAvgStandardCutoffRightSize()`

UnsetCpuUpAvgStandardCutoffRightSize ensures that no value is present for CpuUpAvgStandardCutoffRightSize, not even an explicit nil
### GetCpuUpMaxStandardCutoffRightSize

`func (o *GuidanceSettings) GetCpuUpMaxStandardCutoffRightSize() int32`

GetCpuUpMaxStandardCutoffRightSize returns the CpuUpMaxStandardCutoffRightSize field if non-nil, zero value otherwise.

### GetCpuUpMaxStandardCutoffRightSizeOk

`func (o *GuidanceSettings) GetCpuUpMaxStandardCutoffRightSizeOk() (*int32, bool)`

GetCpuUpMaxStandardCutoffRightSizeOk returns a tuple with the CpuUpMaxStandardCutoffRightSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUpMaxStandardCutoffRightSize

`func (o *GuidanceSettings) SetCpuUpMaxStandardCutoffRightSize(v int32)`

SetCpuUpMaxStandardCutoffRightSize sets CpuUpMaxStandardCutoffRightSize field to given value.

### HasCpuUpMaxStandardCutoffRightSize

`func (o *GuidanceSettings) HasCpuUpMaxStandardCutoffRightSize() bool`

HasCpuUpMaxStandardCutoffRightSize returns a boolean if a field has been set.

### SetCpuUpMaxStandardCutoffRightSizeNil

`func (o *GuidanceSettings) SetCpuUpMaxStandardCutoffRightSizeNil(b bool)`

 SetCpuUpMaxStandardCutoffRightSizeNil sets the value for CpuUpMaxStandardCutoffRightSize to be an explicit nil

### UnsetCpuUpMaxStandardCutoffRightSize
`func (o *GuidanceSettings) UnsetCpuUpMaxStandardCutoffRightSize()`

UnsetCpuUpMaxStandardCutoffRightSize ensures that no value is present for CpuUpMaxStandardCutoffRightSize, not even an explicit nil
### GetMemoryUpAvgStandardCutoffRightSize

`func (o *GuidanceSettings) GetMemoryUpAvgStandardCutoffRightSize() int32`

GetMemoryUpAvgStandardCutoffRightSize returns the MemoryUpAvgStandardCutoffRightSize field if non-nil, zero value otherwise.

### GetMemoryUpAvgStandardCutoffRightSizeOk

`func (o *GuidanceSettings) GetMemoryUpAvgStandardCutoffRightSizeOk() (*int32, bool)`

GetMemoryUpAvgStandardCutoffRightSizeOk returns a tuple with the MemoryUpAvgStandardCutoffRightSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUpAvgStandardCutoffRightSize

`func (o *GuidanceSettings) SetMemoryUpAvgStandardCutoffRightSize(v int32)`

SetMemoryUpAvgStandardCutoffRightSize sets MemoryUpAvgStandardCutoffRightSize field to given value.

### HasMemoryUpAvgStandardCutoffRightSize

`func (o *GuidanceSettings) HasMemoryUpAvgStandardCutoffRightSize() bool`

HasMemoryUpAvgStandardCutoffRightSize returns a boolean if a field has been set.

### SetMemoryUpAvgStandardCutoffRightSizeNil

`func (o *GuidanceSettings) SetMemoryUpAvgStandardCutoffRightSizeNil(b bool)`

 SetMemoryUpAvgStandardCutoffRightSizeNil sets the value for MemoryUpAvgStandardCutoffRightSize to be an explicit nil

### UnsetMemoryUpAvgStandardCutoffRightSize
`func (o *GuidanceSettings) UnsetMemoryUpAvgStandardCutoffRightSize()`

UnsetMemoryUpAvgStandardCutoffRightSize ensures that no value is present for MemoryUpAvgStandardCutoffRightSize, not even an explicit nil
### GetMemoryDownAvgStandardCutoffRightSize

`func (o *GuidanceSettings) GetMemoryDownAvgStandardCutoffRightSize() int32`

GetMemoryDownAvgStandardCutoffRightSize returns the MemoryDownAvgStandardCutoffRightSize field if non-nil, zero value otherwise.

### GetMemoryDownAvgStandardCutoffRightSizeOk

`func (o *GuidanceSettings) GetMemoryDownAvgStandardCutoffRightSizeOk() (*int32, bool)`

GetMemoryDownAvgStandardCutoffRightSizeOk returns a tuple with the MemoryDownAvgStandardCutoffRightSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryDownAvgStandardCutoffRightSize

`func (o *GuidanceSettings) SetMemoryDownAvgStandardCutoffRightSize(v int32)`

SetMemoryDownAvgStandardCutoffRightSize sets MemoryDownAvgStandardCutoffRightSize field to given value.

### HasMemoryDownAvgStandardCutoffRightSize

`func (o *GuidanceSettings) HasMemoryDownAvgStandardCutoffRightSize() bool`

HasMemoryDownAvgStandardCutoffRightSize returns a boolean if a field has been set.

### SetMemoryDownAvgStandardCutoffRightSizeNil

`func (o *GuidanceSettings) SetMemoryDownAvgStandardCutoffRightSizeNil(b bool)`

 SetMemoryDownAvgStandardCutoffRightSizeNil sets the value for MemoryDownAvgStandardCutoffRightSize to be an explicit nil

### UnsetMemoryDownAvgStandardCutoffRightSize
`func (o *GuidanceSettings) UnsetMemoryDownAvgStandardCutoffRightSize()`

UnsetMemoryDownAvgStandardCutoffRightSize ensures that no value is present for MemoryDownAvgStandardCutoffRightSize, not even an explicit nil
### GetMemoryDownMaxStandardCutoffRightSize

`func (o *GuidanceSettings) GetMemoryDownMaxStandardCutoffRightSize() int32`

GetMemoryDownMaxStandardCutoffRightSize returns the MemoryDownMaxStandardCutoffRightSize field if non-nil, zero value otherwise.

### GetMemoryDownMaxStandardCutoffRightSizeOk

`func (o *GuidanceSettings) GetMemoryDownMaxStandardCutoffRightSizeOk() (*int32, bool)`

GetMemoryDownMaxStandardCutoffRightSizeOk returns a tuple with the MemoryDownMaxStandardCutoffRightSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryDownMaxStandardCutoffRightSize

`func (o *GuidanceSettings) SetMemoryDownMaxStandardCutoffRightSize(v int32)`

SetMemoryDownMaxStandardCutoffRightSize sets MemoryDownMaxStandardCutoffRightSize field to given value.

### HasMemoryDownMaxStandardCutoffRightSize

`func (o *GuidanceSettings) HasMemoryDownMaxStandardCutoffRightSize() bool`

HasMemoryDownMaxStandardCutoffRightSize returns a boolean if a field has been set.

### SetMemoryDownMaxStandardCutoffRightSizeNil

`func (o *GuidanceSettings) SetMemoryDownMaxStandardCutoffRightSizeNil(b bool)`

 SetMemoryDownMaxStandardCutoffRightSizeNil sets the value for MemoryDownMaxStandardCutoffRightSize to be an explicit nil

### UnsetMemoryDownMaxStandardCutoffRightSize
`func (o *GuidanceSettings) UnsetMemoryDownMaxStandardCutoffRightSize()`

UnsetMemoryDownMaxStandardCutoffRightSize ensures that no value is present for MemoryDownMaxStandardCutoffRightSize, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


