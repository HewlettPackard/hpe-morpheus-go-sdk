# GetGuidanceSettings200ResponseGuidanceSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CpuAvgCutoffPower** | Pointer to **int32** | Power Shutdown Average CPU (%). Lower limit for average CPU usage | [optional] 
**CpuMaxCutoffPower** | Pointer to **int32** | Power Shutdown Maximum CPU (%). Lower limit for peak CPU usage | [optional] 
**NetworkCutoffPower** | Pointer to **int32** | Power Shutdown Network threshold (bytes). Lower limit for average network bandwidth | [optional] 
**CpuUpAvgStandardCutoffRightSize** | Pointer to **int32** | CPU Up-size Average CPU (%). Upper limit for CPU usage | [optional] 
**CpuUpMaxStandardCutoffRightSize** | Pointer to **int32** | CPU Up-size Maximum CPU (%). Upper limit for peak CPU usage | [optional] 
**MemoryUpAvgStandardCutoffRightSize** | Pointer to **int32** | Memory Up-size Minimum Free Memory (%). Lower limit for average free memory usage | [optional] 
**MemoryDownAvgStandardCutoffRightSize** | Pointer to **int32** | Memory Down-size Maximum Free Memory (%). Upper limit for average free memory | [optional] 
**MemoryDownMaxStandardCutoffRightSize** | Pointer to **int32** | Memory Down-size Maximum Free Memory (%). Upper limit for peak memory usage | [optional] 

## Methods

### NewGetGuidanceSettings200ResponseGuidanceSettings

`func NewGetGuidanceSettings200ResponseGuidanceSettings() *GetGuidanceSettings200ResponseGuidanceSettings`

NewGetGuidanceSettings200ResponseGuidanceSettings instantiates a new GetGuidanceSettings200ResponseGuidanceSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGuidanceSettings200ResponseGuidanceSettingsWithDefaults

`func NewGetGuidanceSettings200ResponseGuidanceSettingsWithDefaults() *GetGuidanceSettings200ResponseGuidanceSettings`

NewGetGuidanceSettings200ResponseGuidanceSettingsWithDefaults instantiates a new GetGuidanceSettings200ResponseGuidanceSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpuAvgCutoffPower

`func (o *GetGuidanceSettings200ResponseGuidanceSettings) GetCpuAvgCutoffPower() int32`

GetCpuAvgCutoffPower returns the CpuAvgCutoffPower field if non-nil, zero value otherwise.

### GetCpuAvgCutoffPowerOk

`func (o *GetGuidanceSettings200ResponseGuidanceSettings) GetCpuAvgCutoffPowerOk() (*int32, bool)`

GetCpuAvgCutoffPowerOk returns a tuple with the CpuAvgCutoffPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuAvgCutoffPower

`func (o *GetGuidanceSettings200ResponseGuidanceSettings) SetCpuAvgCutoffPower(v int32)`

SetCpuAvgCutoffPower sets CpuAvgCutoffPower field to given value.

### HasCpuAvgCutoffPower

`func (o *GetGuidanceSettings200ResponseGuidanceSettings) HasCpuAvgCutoffPower() bool`

HasCpuAvgCutoffPower returns a boolean if a field has been set.

### GetCpuMaxCutoffPower

`func (o *GetGuidanceSettings200ResponseGuidanceSettings) GetCpuMaxCutoffPower() int32`

GetCpuMaxCutoffPower returns the CpuMaxCutoffPower field if non-nil, zero value otherwise.

### GetCpuMaxCutoffPowerOk

`func (o *GetGuidanceSettings200ResponseGuidanceSettings) GetCpuMaxCutoffPowerOk() (*int32, bool)`

GetCpuMaxCutoffPowerOk returns a tuple with the CpuMaxCutoffPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuMaxCutoffPower

`func (o *GetGuidanceSettings200ResponseGuidanceSettings) SetCpuMaxCutoffPower(v int32)`

SetCpuMaxCutoffPower sets CpuMaxCutoffPower field to given value.

### HasCpuMaxCutoffPower

`func (o *GetGuidanceSettings200ResponseGuidanceSettings) HasCpuMaxCutoffPower() bool`

HasCpuMaxCutoffPower returns a boolean if a field has been set.

### GetNetworkCutoffPower

`func (o *GetGuidanceSettings200ResponseGuidanceSettings) GetNetworkCutoffPower() int32`

GetNetworkCutoffPower returns the NetworkCutoffPower field if non-nil, zero value otherwise.

### GetNetworkCutoffPowerOk

`func (o *GetGuidanceSettings200ResponseGuidanceSettings) GetNetworkCutoffPowerOk() (*int32, bool)`

GetNetworkCutoffPowerOk returns a tuple with the NetworkCutoffPower field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkCutoffPower

`func (o *GetGuidanceSettings200ResponseGuidanceSettings) SetNetworkCutoffPower(v int32)`

SetNetworkCutoffPower sets NetworkCutoffPower field to given value.

### HasNetworkCutoffPower

`func (o *GetGuidanceSettings200ResponseGuidanceSettings) HasNetworkCutoffPower() bool`

HasNetworkCutoffPower returns a boolean if a field has been set.

### GetCpuUpAvgStandardCutoffRightSize

`func (o *GetGuidanceSettings200ResponseGuidanceSettings) GetCpuUpAvgStandardCutoffRightSize() int32`

GetCpuUpAvgStandardCutoffRightSize returns the CpuUpAvgStandardCutoffRightSize field if non-nil, zero value otherwise.

### GetCpuUpAvgStandardCutoffRightSizeOk

`func (o *GetGuidanceSettings200ResponseGuidanceSettings) GetCpuUpAvgStandardCutoffRightSizeOk() (*int32, bool)`

GetCpuUpAvgStandardCutoffRightSizeOk returns a tuple with the CpuUpAvgStandardCutoffRightSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUpAvgStandardCutoffRightSize

`func (o *GetGuidanceSettings200ResponseGuidanceSettings) SetCpuUpAvgStandardCutoffRightSize(v int32)`

SetCpuUpAvgStandardCutoffRightSize sets CpuUpAvgStandardCutoffRightSize field to given value.

### HasCpuUpAvgStandardCutoffRightSize

`func (o *GetGuidanceSettings200ResponseGuidanceSettings) HasCpuUpAvgStandardCutoffRightSize() bool`

HasCpuUpAvgStandardCutoffRightSize returns a boolean if a field has been set.

### GetCpuUpMaxStandardCutoffRightSize

`func (o *GetGuidanceSettings200ResponseGuidanceSettings) GetCpuUpMaxStandardCutoffRightSize() int32`

GetCpuUpMaxStandardCutoffRightSize returns the CpuUpMaxStandardCutoffRightSize field if non-nil, zero value otherwise.

### GetCpuUpMaxStandardCutoffRightSizeOk

`func (o *GetGuidanceSettings200ResponseGuidanceSettings) GetCpuUpMaxStandardCutoffRightSizeOk() (*int32, bool)`

GetCpuUpMaxStandardCutoffRightSizeOk returns a tuple with the CpuUpMaxStandardCutoffRightSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUpMaxStandardCutoffRightSize

`func (o *GetGuidanceSettings200ResponseGuidanceSettings) SetCpuUpMaxStandardCutoffRightSize(v int32)`

SetCpuUpMaxStandardCutoffRightSize sets CpuUpMaxStandardCutoffRightSize field to given value.

### HasCpuUpMaxStandardCutoffRightSize

`func (o *GetGuidanceSettings200ResponseGuidanceSettings) HasCpuUpMaxStandardCutoffRightSize() bool`

HasCpuUpMaxStandardCutoffRightSize returns a boolean if a field has been set.

### GetMemoryUpAvgStandardCutoffRightSize

`func (o *GetGuidanceSettings200ResponseGuidanceSettings) GetMemoryUpAvgStandardCutoffRightSize() int32`

GetMemoryUpAvgStandardCutoffRightSize returns the MemoryUpAvgStandardCutoffRightSize field if non-nil, zero value otherwise.

### GetMemoryUpAvgStandardCutoffRightSizeOk

`func (o *GetGuidanceSettings200ResponseGuidanceSettings) GetMemoryUpAvgStandardCutoffRightSizeOk() (*int32, bool)`

GetMemoryUpAvgStandardCutoffRightSizeOk returns a tuple with the MemoryUpAvgStandardCutoffRightSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryUpAvgStandardCutoffRightSize

`func (o *GetGuidanceSettings200ResponseGuidanceSettings) SetMemoryUpAvgStandardCutoffRightSize(v int32)`

SetMemoryUpAvgStandardCutoffRightSize sets MemoryUpAvgStandardCutoffRightSize field to given value.

### HasMemoryUpAvgStandardCutoffRightSize

`func (o *GetGuidanceSettings200ResponseGuidanceSettings) HasMemoryUpAvgStandardCutoffRightSize() bool`

HasMemoryUpAvgStandardCutoffRightSize returns a boolean if a field has been set.

### GetMemoryDownAvgStandardCutoffRightSize

`func (o *GetGuidanceSettings200ResponseGuidanceSettings) GetMemoryDownAvgStandardCutoffRightSize() int32`

GetMemoryDownAvgStandardCutoffRightSize returns the MemoryDownAvgStandardCutoffRightSize field if non-nil, zero value otherwise.

### GetMemoryDownAvgStandardCutoffRightSizeOk

`func (o *GetGuidanceSettings200ResponseGuidanceSettings) GetMemoryDownAvgStandardCutoffRightSizeOk() (*int32, bool)`

GetMemoryDownAvgStandardCutoffRightSizeOk returns a tuple with the MemoryDownAvgStandardCutoffRightSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryDownAvgStandardCutoffRightSize

`func (o *GetGuidanceSettings200ResponseGuidanceSettings) SetMemoryDownAvgStandardCutoffRightSize(v int32)`

SetMemoryDownAvgStandardCutoffRightSize sets MemoryDownAvgStandardCutoffRightSize field to given value.

### HasMemoryDownAvgStandardCutoffRightSize

`func (o *GetGuidanceSettings200ResponseGuidanceSettings) HasMemoryDownAvgStandardCutoffRightSize() bool`

HasMemoryDownAvgStandardCutoffRightSize returns a boolean if a field has been set.

### GetMemoryDownMaxStandardCutoffRightSize

`func (o *GetGuidanceSettings200ResponseGuidanceSettings) GetMemoryDownMaxStandardCutoffRightSize() int32`

GetMemoryDownMaxStandardCutoffRightSize returns the MemoryDownMaxStandardCutoffRightSize field if non-nil, zero value otherwise.

### GetMemoryDownMaxStandardCutoffRightSizeOk

`func (o *GetGuidanceSettings200ResponseGuidanceSettings) GetMemoryDownMaxStandardCutoffRightSizeOk() (*int32, bool)`

GetMemoryDownMaxStandardCutoffRightSizeOk returns a tuple with the MemoryDownMaxStandardCutoffRightSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemoryDownMaxStandardCutoffRightSize

`func (o *GetGuidanceSettings200ResponseGuidanceSettings) SetMemoryDownMaxStandardCutoffRightSize(v int32)`

SetMemoryDownMaxStandardCutoffRightSize sets MemoryDownMaxStandardCutoffRightSize field to given value.

### HasMemoryDownMaxStandardCutoffRightSize

`func (o *GetGuidanceSettings200ResponseGuidanceSettings) HasMemoryDownMaxStandardCutoffRightSize() bool`

HasMemoryDownMaxStandardCutoffRightSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


