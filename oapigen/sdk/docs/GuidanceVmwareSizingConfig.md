# GuidanceVmwareSizingConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exists** | Pointer to **bool** |  | [optional] 
**ObjectId** | Pointer to **int64** |  | [optional] 
**CpuTotalTimeCount** | Pointer to **NullableFloat32** |  | [optional] 
**CpuTotalTimeMin** | Pointer to **NullableFloat32** |  | [optional] 
**CpuTotalTimeMax** | Pointer to **NullableFloat32** |  | [optional] 
**CpuTotalTimeAvg** | Pointer to **NullableFloat32** |  | [optional] 
**CpuTotalTimeSum** | Pointer to **NullableFloat32** |  | [optional] 
**CpuIdleTimeCount** | Pointer to **NullableFloat32** |  | [optional] 
**CpuIdleTimeMin** | Pointer to **NullableFloat32** |  | [optional] 
**CpuIdleTimeMax** | Pointer to **NullableFloat32** |  | [optional] 
**CpuIdleTimeAvg** | Pointer to **NullableFloat32** |  | [optional] 
**CpuIdleTimeSum** | Pointer to **NullableFloat32** |  | [optional] 
**CpuUsageCount** | Pointer to **NullableFloat32** |  | [optional] 
**CpuUsageMin** | Pointer to **NullableFloat32** |  | [optional] 
**CpuUsageMax** | Pointer to **NullableFloat32** |  | [optional] 
**CpuUsageAvg** | Pointer to **NullableFloat32** |  | [optional] 
**CpuUsageSum** | Pointer to **NullableFloat32** |  | [optional] 
**MaxMemoryCount** | Pointer to **NullableFloat32** |  | [optional] 
**MaxMemoryMin** | Pointer to **NullableFloat32** |  | [optional] 
**MaxMemoryMax** | Pointer to **NullableFloat32** |  | [optional] 
**MaxMemoryAvg** | Pointer to **NullableFloat32** |  | [optional] 
**MaxMemorySum** | Pointer to **NullableFloat32** |  | [optional] 
**CpuUserTimeCount** | Pointer to **NullableFloat32** |  | [optional] 
**CpuUserTimeMin** | Pointer to **NullableFloat32** |  | [optional] 
**CpuUserTimeMax** | Pointer to **NullableFloat32** |  | [optional] 
**CpuUserTimeAvg** | Pointer to **NullableFloat32** |  | [optional] 
**CpuUserTimeSum** | Pointer to **NullableFloat32** |  | [optional] 
**CpuSystemTimeCount** | Pointer to **NullableFloat32** |  | [optional] 
**CpuSystemTimeMin** | Pointer to **NullableFloat32** |  | [optional] 
**CpuSystemTimeMax** | Pointer to **NullableFloat32** |  | [optional] 
**CpuSystemTimeAvg** | Pointer to **float32** |  | [optional] 
**CpuSystemTimeSum** | Pointer to **NullableFloat32** |  | [optional] 
**UsedMemoryCount** | Pointer to **NullableFloat32** |  | [optional] 
**UsedMemoryMin** | Pointer to **NullableFloat32** |  | [optional] 
**UsedMemoryMax** | Pointer to **NullableFloat32** |  | [optional] 
**UsedMemoryAvg** | Pointer to **NullableFloat32** |  | [optional] 
**UsedMemorySum** | Pointer to **NullableFloat32** |  | [optional] 
**FreeMemoryCount** | Pointer to **NullableFloat32** |  | [optional] 
**FreeMemoryMin** | Pointer to **NullableFloat32** |  | [optional] 
**FreeMemoryMax** | Pointer to **NullableFloat32** |  | [optional] 
**FreeMemoryAvg** | Pointer to **NullableFloat32** |  | [optional] 
**FreeMemorySum** | Pointer to **NullableFloat32** |  | [optional] 

## Methods

### NewGuidanceVmwareSizingConfig

`func NewGuidanceVmwareSizingConfig() *GuidanceVmwareSizingConfig`

NewGuidanceVmwareSizingConfig instantiates a new GuidanceVmwareSizingConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGuidanceVmwareSizingConfigWithDefaults

`func NewGuidanceVmwareSizingConfigWithDefaults() *GuidanceVmwareSizingConfig`

NewGuidanceVmwareSizingConfigWithDefaults instantiates a new GuidanceVmwareSizingConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExists

`func (o *GuidanceVmwareSizingConfig) GetExists() bool`

GetExists returns the Exists field if non-nil, zero value otherwise.

### GetExistsOk

`func (o *GuidanceVmwareSizingConfig) GetExistsOk() (*bool, bool)`

GetExistsOk returns a tuple with the Exists field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExists

`func (o *GuidanceVmwareSizingConfig) SetExists(v bool)`

SetExists sets Exists field to given value.

### HasExists

`func (o *GuidanceVmwareSizingConfig) HasExists() bool`

HasExists returns a boolean if a field has been set.

### GetObjectId

`func (o *GuidanceVmwareSizingConfig) GetObjectId() int64`

GetObjectId returns the ObjectId field if non-nil, zero value otherwise.

### GetObjectIdOk

`func (o *GuidanceVmwareSizingConfig) GetObjectIdOk() (*int64, bool)`

GetObjectIdOk returns a tuple with the ObjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetObjectId

`func (o *GuidanceVmwareSizingConfig) SetObjectId(v int64)`

SetObjectId sets ObjectId field to given value.

### HasObjectId

`func (o *GuidanceVmwareSizingConfig) HasObjectId() bool`

HasObjectId returns a boolean if a field has been set.

### GetCpuTotalTimeCount

`func (o *GuidanceVmwareSizingConfig) GetCpuTotalTimeCount() float32`

GetCpuTotalTimeCount returns the CpuTotalTimeCount field if non-nil, zero value otherwise.

### GetCpuTotalTimeCountOk

`func (o *GuidanceVmwareSizingConfig) GetCpuTotalTimeCountOk() (*float32, bool)`

GetCpuTotalTimeCountOk returns a tuple with the CpuTotalTimeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuTotalTimeCount

`func (o *GuidanceVmwareSizingConfig) SetCpuTotalTimeCount(v float32)`

SetCpuTotalTimeCount sets CpuTotalTimeCount field to given value.

### HasCpuTotalTimeCount

`func (o *GuidanceVmwareSizingConfig) HasCpuTotalTimeCount() bool`

HasCpuTotalTimeCount returns a boolean if a field has been set.

### SetCpuTotalTimeCountNil

`func (o *GuidanceVmwareSizingConfig) SetCpuTotalTimeCountNil(b bool)`

 SetCpuTotalTimeCountNil sets the value for CpuTotalTimeCount to be an explicit nil

### UnsetCpuTotalTimeCount
`func (o *GuidanceVmwareSizingConfig) UnsetCpuTotalTimeCount()`

UnsetCpuTotalTimeCount ensures that no value is present for CpuTotalTimeCount, not even an explicit nil
### GetCpuTotalTimeMin

`func (o *GuidanceVmwareSizingConfig) GetCpuTotalTimeMin() float32`

GetCpuTotalTimeMin returns the CpuTotalTimeMin field if non-nil, zero value otherwise.

### GetCpuTotalTimeMinOk

`func (o *GuidanceVmwareSizingConfig) GetCpuTotalTimeMinOk() (*float32, bool)`

GetCpuTotalTimeMinOk returns a tuple with the CpuTotalTimeMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuTotalTimeMin

`func (o *GuidanceVmwareSizingConfig) SetCpuTotalTimeMin(v float32)`

SetCpuTotalTimeMin sets CpuTotalTimeMin field to given value.

### HasCpuTotalTimeMin

`func (o *GuidanceVmwareSizingConfig) HasCpuTotalTimeMin() bool`

HasCpuTotalTimeMin returns a boolean if a field has been set.

### SetCpuTotalTimeMinNil

`func (o *GuidanceVmwareSizingConfig) SetCpuTotalTimeMinNil(b bool)`

 SetCpuTotalTimeMinNil sets the value for CpuTotalTimeMin to be an explicit nil

### UnsetCpuTotalTimeMin
`func (o *GuidanceVmwareSizingConfig) UnsetCpuTotalTimeMin()`

UnsetCpuTotalTimeMin ensures that no value is present for CpuTotalTimeMin, not even an explicit nil
### GetCpuTotalTimeMax

`func (o *GuidanceVmwareSizingConfig) GetCpuTotalTimeMax() float32`

GetCpuTotalTimeMax returns the CpuTotalTimeMax field if non-nil, zero value otherwise.

### GetCpuTotalTimeMaxOk

`func (o *GuidanceVmwareSizingConfig) GetCpuTotalTimeMaxOk() (*float32, bool)`

GetCpuTotalTimeMaxOk returns a tuple with the CpuTotalTimeMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuTotalTimeMax

`func (o *GuidanceVmwareSizingConfig) SetCpuTotalTimeMax(v float32)`

SetCpuTotalTimeMax sets CpuTotalTimeMax field to given value.

### HasCpuTotalTimeMax

`func (o *GuidanceVmwareSizingConfig) HasCpuTotalTimeMax() bool`

HasCpuTotalTimeMax returns a boolean if a field has been set.

### SetCpuTotalTimeMaxNil

`func (o *GuidanceVmwareSizingConfig) SetCpuTotalTimeMaxNil(b bool)`

 SetCpuTotalTimeMaxNil sets the value for CpuTotalTimeMax to be an explicit nil

### UnsetCpuTotalTimeMax
`func (o *GuidanceVmwareSizingConfig) UnsetCpuTotalTimeMax()`

UnsetCpuTotalTimeMax ensures that no value is present for CpuTotalTimeMax, not even an explicit nil
### GetCpuTotalTimeAvg

`func (o *GuidanceVmwareSizingConfig) GetCpuTotalTimeAvg() float32`

GetCpuTotalTimeAvg returns the CpuTotalTimeAvg field if non-nil, zero value otherwise.

### GetCpuTotalTimeAvgOk

`func (o *GuidanceVmwareSizingConfig) GetCpuTotalTimeAvgOk() (*float32, bool)`

GetCpuTotalTimeAvgOk returns a tuple with the CpuTotalTimeAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuTotalTimeAvg

`func (o *GuidanceVmwareSizingConfig) SetCpuTotalTimeAvg(v float32)`

SetCpuTotalTimeAvg sets CpuTotalTimeAvg field to given value.

### HasCpuTotalTimeAvg

`func (o *GuidanceVmwareSizingConfig) HasCpuTotalTimeAvg() bool`

HasCpuTotalTimeAvg returns a boolean if a field has been set.

### SetCpuTotalTimeAvgNil

`func (o *GuidanceVmwareSizingConfig) SetCpuTotalTimeAvgNil(b bool)`

 SetCpuTotalTimeAvgNil sets the value for CpuTotalTimeAvg to be an explicit nil

### UnsetCpuTotalTimeAvg
`func (o *GuidanceVmwareSizingConfig) UnsetCpuTotalTimeAvg()`

UnsetCpuTotalTimeAvg ensures that no value is present for CpuTotalTimeAvg, not even an explicit nil
### GetCpuTotalTimeSum

`func (o *GuidanceVmwareSizingConfig) GetCpuTotalTimeSum() float32`

GetCpuTotalTimeSum returns the CpuTotalTimeSum field if non-nil, zero value otherwise.

### GetCpuTotalTimeSumOk

`func (o *GuidanceVmwareSizingConfig) GetCpuTotalTimeSumOk() (*float32, bool)`

GetCpuTotalTimeSumOk returns a tuple with the CpuTotalTimeSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuTotalTimeSum

`func (o *GuidanceVmwareSizingConfig) SetCpuTotalTimeSum(v float32)`

SetCpuTotalTimeSum sets CpuTotalTimeSum field to given value.

### HasCpuTotalTimeSum

`func (o *GuidanceVmwareSizingConfig) HasCpuTotalTimeSum() bool`

HasCpuTotalTimeSum returns a boolean if a field has been set.

### SetCpuTotalTimeSumNil

`func (o *GuidanceVmwareSizingConfig) SetCpuTotalTimeSumNil(b bool)`

 SetCpuTotalTimeSumNil sets the value for CpuTotalTimeSum to be an explicit nil

### UnsetCpuTotalTimeSum
`func (o *GuidanceVmwareSizingConfig) UnsetCpuTotalTimeSum()`

UnsetCpuTotalTimeSum ensures that no value is present for CpuTotalTimeSum, not even an explicit nil
### GetCpuIdleTimeCount

`func (o *GuidanceVmwareSizingConfig) GetCpuIdleTimeCount() float32`

GetCpuIdleTimeCount returns the CpuIdleTimeCount field if non-nil, zero value otherwise.

### GetCpuIdleTimeCountOk

`func (o *GuidanceVmwareSizingConfig) GetCpuIdleTimeCountOk() (*float32, bool)`

GetCpuIdleTimeCountOk returns a tuple with the CpuIdleTimeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuIdleTimeCount

`func (o *GuidanceVmwareSizingConfig) SetCpuIdleTimeCount(v float32)`

SetCpuIdleTimeCount sets CpuIdleTimeCount field to given value.

### HasCpuIdleTimeCount

`func (o *GuidanceVmwareSizingConfig) HasCpuIdleTimeCount() bool`

HasCpuIdleTimeCount returns a boolean if a field has been set.

### SetCpuIdleTimeCountNil

`func (o *GuidanceVmwareSizingConfig) SetCpuIdleTimeCountNil(b bool)`

 SetCpuIdleTimeCountNil sets the value for CpuIdleTimeCount to be an explicit nil

### UnsetCpuIdleTimeCount
`func (o *GuidanceVmwareSizingConfig) UnsetCpuIdleTimeCount()`

UnsetCpuIdleTimeCount ensures that no value is present for CpuIdleTimeCount, not even an explicit nil
### GetCpuIdleTimeMin

`func (o *GuidanceVmwareSizingConfig) GetCpuIdleTimeMin() float32`

GetCpuIdleTimeMin returns the CpuIdleTimeMin field if non-nil, zero value otherwise.

### GetCpuIdleTimeMinOk

`func (o *GuidanceVmwareSizingConfig) GetCpuIdleTimeMinOk() (*float32, bool)`

GetCpuIdleTimeMinOk returns a tuple with the CpuIdleTimeMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuIdleTimeMin

`func (o *GuidanceVmwareSizingConfig) SetCpuIdleTimeMin(v float32)`

SetCpuIdleTimeMin sets CpuIdleTimeMin field to given value.

### HasCpuIdleTimeMin

`func (o *GuidanceVmwareSizingConfig) HasCpuIdleTimeMin() bool`

HasCpuIdleTimeMin returns a boolean if a field has been set.

### SetCpuIdleTimeMinNil

`func (o *GuidanceVmwareSizingConfig) SetCpuIdleTimeMinNil(b bool)`

 SetCpuIdleTimeMinNil sets the value for CpuIdleTimeMin to be an explicit nil

### UnsetCpuIdleTimeMin
`func (o *GuidanceVmwareSizingConfig) UnsetCpuIdleTimeMin()`

UnsetCpuIdleTimeMin ensures that no value is present for CpuIdleTimeMin, not even an explicit nil
### GetCpuIdleTimeMax

`func (o *GuidanceVmwareSizingConfig) GetCpuIdleTimeMax() float32`

GetCpuIdleTimeMax returns the CpuIdleTimeMax field if non-nil, zero value otherwise.

### GetCpuIdleTimeMaxOk

`func (o *GuidanceVmwareSizingConfig) GetCpuIdleTimeMaxOk() (*float32, bool)`

GetCpuIdleTimeMaxOk returns a tuple with the CpuIdleTimeMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuIdleTimeMax

`func (o *GuidanceVmwareSizingConfig) SetCpuIdleTimeMax(v float32)`

SetCpuIdleTimeMax sets CpuIdleTimeMax field to given value.

### HasCpuIdleTimeMax

`func (o *GuidanceVmwareSizingConfig) HasCpuIdleTimeMax() bool`

HasCpuIdleTimeMax returns a boolean if a field has been set.

### SetCpuIdleTimeMaxNil

`func (o *GuidanceVmwareSizingConfig) SetCpuIdleTimeMaxNil(b bool)`

 SetCpuIdleTimeMaxNil sets the value for CpuIdleTimeMax to be an explicit nil

### UnsetCpuIdleTimeMax
`func (o *GuidanceVmwareSizingConfig) UnsetCpuIdleTimeMax()`

UnsetCpuIdleTimeMax ensures that no value is present for CpuIdleTimeMax, not even an explicit nil
### GetCpuIdleTimeAvg

`func (o *GuidanceVmwareSizingConfig) GetCpuIdleTimeAvg() float32`

GetCpuIdleTimeAvg returns the CpuIdleTimeAvg field if non-nil, zero value otherwise.

### GetCpuIdleTimeAvgOk

`func (o *GuidanceVmwareSizingConfig) GetCpuIdleTimeAvgOk() (*float32, bool)`

GetCpuIdleTimeAvgOk returns a tuple with the CpuIdleTimeAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuIdleTimeAvg

`func (o *GuidanceVmwareSizingConfig) SetCpuIdleTimeAvg(v float32)`

SetCpuIdleTimeAvg sets CpuIdleTimeAvg field to given value.

### HasCpuIdleTimeAvg

`func (o *GuidanceVmwareSizingConfig) HasCpuIdleTimeAvg() bool`

HasCpuIdleTimeAvg returns a boolean if a field has been set.

### SetCpuIdleTimeAvgNil

`func (o *GuidanceVmwareSizingConfig) SetCpuIdleTimeAvgNil(b bool)`

 SetCpuIdleTimeAvgNil sets the value for CpuIdleTimeAvg to be an explicit nil

### UnsetCpuIdleTimeAvg
`func (o *GuidanceVmwareSizingConfig) UnsetCpuIdleTimeAvg()`

UnsetCpuIdleTimeAvg ensures that no value is present for CpuIdleTimeAvg, not even an explicit nil
### GetCpuIdleTimeSum

`func (o *GuidanceVmwareSizingConfig) GetCpuIdleTimeSum() float32`

GetCpuIdleTimeSum returns the CpuIdleTimeSum field if non-nil, zero value otherwise.

### GetCpuIdleTimeSumOk

`func (o *GuidanceVmwareSizingConfig) GetCpuIdleTimeSumOk() (*float32, bool)`

GetCpuIdleTimeSumOk returns a tuple with the CpuIdleTimeSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuIdleTimeSum

`func (o *GuidanceVmwareSizingConfig) SetCpuIdleTimeSum(v float32)`

SetCpuIdleTimeSum sets CpuIdleTimeSum field to given value.

### HasCpuIdleTimeSum

`func (o *GuidanceVmwareSizingConfig) HasCpuIdleTimeSum() bool`

HasCpuIdleTimeSum returns a boolean if a field has been set.

### SetCpuIdleTimeSumNil

`func (o *GuidanceVmwareSizingConfig) SetCpuIdleTimeSumNil(b bool)`

 SetCpuIdleTimeSumNil sets the value for CpuIdleTimeSum to be an explicit nil

### UnsetCpuIdleTimeSum
`func (o *GuidanceVmwareSizingConfig) UnsetCpuIdleTimeSum()`

UnsetCpuIdleTimeSum ensures that no value is present for CpuIdleTimeSum, not even an explicit nil
### GetCpuUsageCount

`func (o *GuidanceVmwareSizingConfig) GetCpuUsageCount() float32`

GetCpuUsageCount returns the CpuUsageCount field if non-nil, zero value otherwise.

### GetCpuUsageCountOk

`func (o *GuidanceVmwareSizingConfig) GetCpuUsageCountOk() (*float32, bool)`

GetCpuUsageCountOk returns a tuple with the CpuUsageCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsageCount

`func (o *GuidanceVmwareSizingConfig) SetCpuUsageCount(v float32)`

SetCpuUsageCount sets CpuUsageCount field to given value.

### HasCpuUsageCount

`func (o *GuidanceVmwareSizingConfig) HasCpuUsageCount() bool`

HasCpuUsageCount returns a boolean if a field has been set.

### SetCpuUsageCountNil

`func (o *GuidanceVmwareSizingConfig) SetCpuUsageCountNil(b bool)`

 SetCpuUsageCountNil sets the value for CpuUsageCount to be an explicit nil

### UnsetCpuUsageCount
`func (o *GuidanceVmwareSizingConfig) UnsetCpuUsageCount()`

UnsetCpuUsageCount ensures that no value is present for CpuUsageCount, not even an explicit nil
### GetCpuUsageMin

`func (o *GuidanceVmwareSizingConfig) GetCpuUsageMin() float32`

GetCpuUsageMin returns the CpuUsageMin field if non-nil, zero value otherwise.

### GetCpuUsageMinOk

`func (o *GuidanceVmwareSizingConfig) GetCpuUsageMinOk() (*float32, bool)`

GetCpuUsageMinOk returns a tuple with the CpuUsageMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsageMin

`func (o *GuidanceVmwareSizingConfig) SetCpuUsageMin(v float32)`

SetCpuUsageMin sets CpuUsageMin field to given value.

### HasCpuUsageMin

`func (o *GuidanceVmwareSizingConfig) HasCpuUsageMin() bool`

HasCpuUsageMin returns a boolean if a field has been set.

### SetCpuUsageMinNil

`func (o *GuidanceVmwareSizingConfig) SetCpuUsageMinNil(b bool)`

 SetCpuUsageMinNil sets the value for CpuUsageMin to be an explicit nil

### UnsetCpuUsageMin
`func (o *GuidanceVmwareSizingConfig) UnsetCpuUsageMin()`

UnsetCpuUsageMin ensures that no value is present for CpuUsageMin, not even an explicit nil
### GetCpuUsageMax

`func (o *GuidanceVmwareSizingConfig) GetCpuUsageMax() float32`

GetCpuUsageMax returns the CpuUsageMax field if non-nil, zero value otherwise.

### GetCpuUsageMaxOk

`func (o *GuidanceVmwareSizingConfig) GetCpuUsageMaxOk() (*float32, bool)`

GetCpuUsageMaxOk returns a tuple with the CpuUsageMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsageMax

`func (o *GuidanceVmwareSizingConfig) SetCpuUsageMax(v float32)`

SetCpuUsageMax sets CpuUsageMax field to given value.

### HasCpuUsageMax

`func (o *GuidanceVmwareSizingConfig) HasCpuUsageMax() bool`

HasCpuUsageMax returns a boolean if a field has been set.

### SetCpuUsageMaxNil

`func (o *GuidanceVmwareSizingConfig) SetCpuUsageMaxNil(b bool)`

 SetCpuUsageMaxNil sets the value for CpuUsageMax to be an explicit nil

### UnsetCpuUsageMax
`func (o *GuidanceVmwareSizingConfig) UnsetCpuUsageMax()`

UnsetCpuUsageMax ensures that no value is present for CpuUsageMax, not even an explicit nil
### GetCpuUsageAvg

`func (o *GuidanceVmwareSizingConfig) GetCpuUsageAvg() float32`

GetCpuUsageAvg returns the CpuUsageAvg field if non-nil, zero value otherwise.

### GetCpuUsageAvgOk

`func (o *GuidanceVmwareSizingConfig) GetCpuUsageAvgOk() (*float32, bool)`

GetCpuUsageAvgOk returns a tuple with the CpuUsageAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsageAvg

`func (o *GuidanceVmwareSizingConfig) SetCpuUsageAvg(v float32)`

SetCpuUsageAvg sets CpuUsageAvg field to given value.

### HasCpuUsageAvg

`func (o *GuidanceVmwareSizingConfig) HasCpuUsageAvg() bool`

HasCpuUsageAvg returns a boolean if a field has been set.

### SetCpuUsageAvgNil

`func (o *GuidanceVmwareSizingConfig) SetCpuUsageAvgNil(b bool)`

 SetCpuUsageAvgNil sets the value for CpuUsageAvg to be an explicit nil

### UnsetCpuUsageAvg
`func (o *GuidanceVmwareSizingConfig) UnsetCpuUsageAvg()`

UnsetCpuUsageAvg ensures that no value is present for CpuUsageAvg, not even an explicit nil
### GetCpuUsageSum

`func (o *GuidanceVmwareSizingConfig) GetCpuUsageSum() float32`

GetCpuUsageSum returns the CpuUsageSum field if non-nil, zero value otherwise.

### GetCpuUsageSumOk

`func (o *GuidanceVmwareSizingConfig) GetCpuUsageSumOk() (*float32, bool)`

GetCpuUsageSumOk returns a tuple with the CpuUsageSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUsageSum

`func (o *GuidanceVmwareSizingConfig) SetCpuUsageSum(v float32)`

SetCpuUsageSum sets CpuUsageSum field to given value.

### HasCpuUsageSum

`func (o *GuidanceVmwareSizingConfig) HasCpuUsageSum() bool`

HasCpuUsageSum returns a boolean if a field has been set.

### SetCpuUsageSumNil

`func (o *GuidanceVmwareSizingConfig) SetCpuUsageSumNil(b bool)`

 SetCpuUsageSumNil sets the value for CpuUsageSum to be an explicit nil

### UnsetCpuUsageSum
`func (o *GuidanceVmwareSizingConfig) UnsetCpuUsageSum()`

UnsetCpuUsageSum ensures that no value is present for CpuUsageSum, not even an explicit nil
### GetMaxMemoryCount

`func (o *GuidanceVmwareSizingConfig) GetMaxMemoryCount() float32`

GetMaxMemoryCount returns the MaxMemoryCount field if non-nil, zero value otherwise.

### GetMaxMemoryCountOk

`func (o *GuidanceVmwareSizingConfig) GetMaxMemoryCountOk() (*float32, bool)`

GetMaxMemoryCountOk returns a tuple with the MaxMemoryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemoryCount

`func (o *GuidanceVmwareSizingConfig) SetMaxMemoryCount(v float32)`

SetMaxMemoryCount sets MaxMemoryCount field to given value.

### HasMaxMemoryCount

`func (o *GuidanceVmwareSizingConfig) HasMaxMemoryCount() bool`

HasMaxMemoryCount returns a boolean if a field has been set.

### SetMaxMemoryCountNil

`func (o *GuidanceVmwareSizingConfig) SetMaxMemoryCountNil(b bool)`

 SetMaxMemoryCountNil sets the value for MaxMemoryCount to be an explicit nil

### UnsetMaxMemoryCount
`func (o *GuidanceVmwareSizingConfig) UnsetMaxMemoryCount()`

UnsetMaxMemoryCount ensures that no value is present for MaxMemoryCount, not even an explicit nil
### GetMaxMemoryMin

`func (o *GuidanceVmwareSizingConfig) GetMaxMemoryMin() float32`

GetMaxMemoryMin returns the MaxMemoryMin field if non-nil, zero value otherwise.

### GetMaxMemoryMinOk

`func (o *GuidanceVmwareSizingConfig) GetMaxMemoryMinOk() (*float32, bool)`

GetMaxMemoryMinOk returns a tuple with the MaxMemoryMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemoryMin

`func (o *GuidanceVmwareSizingConfig) SetMaxMemoryMin(v float32)`

SetMaxMemoryMin sets MaxMemoryMin field to given value.

### HasMaxMemoryMin

`func (o *GuidanceVmwareSizingConfig) HasMaxMemoryMin() bool`

HasMaxMemoryMin returns a boolean if a field has been set.

### SetMaxMemoryMinNil

`func (o *GuidanceVmwareSizingConfig) SetMaxMemoryMinNil(b bool)`

 SetMaxMemoryMinNil sets the value for MaxMemoryMin to be an explicit nil

### UnsetMaxMemoryMin
`func (o *GuidanceVmwareSizingConfig) UnsetMaxMemoryMin()`

UnsetMaxMemoryMin ensures that no value is present for MaxMemoryMin, not even an explicit nil
### GetMaxMemoryMax

`func (o *GuidanceVmwareSizingConfig) GetMaxMemoryMax() float32`

GetMaxMemoryMax returns the MaxMemoryMax field if non-nil, zero value otherwise.

### GetMaxMemoryMaxOk

`func (o *GuidanceVmwareSizingConfig) GetMaxMemoryMaxOk() (*float32, bool)`

GetMaxMemoryMaxOk returns a tuple with the MaxMemoryMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemoryMax

`func (o *GuidanceVmwareSizingConfig) SetMaxMemoryMax(v float32)`

SetMaxMemoryMax sets MaxMemoryMax field to given value.

### HasMaxMemoryMax

`func (o *GuidanceVmwareSizingConfig) HasMaxMemoryMax() bool`

HasMaxMemoryMax returns a boolean if a field has been set.

### SetMaxMemoryMaxNil

`func (o *GuidanceVmwareSizingConfig) SetMaxMemoryMaxNil(b bool)`

 SetMaxMemoryMaxNil sets the value for MaxMemoryMax to be an explicit nil

### UnsetMaxMemoryMax
`func (o *GuidanceVmwareSizingConfig) UnsetMaxMemoryMax()`

UnsetMaxMemoryMax ensures that no value is present for MaxMemoryMax, not even an explicit nil
### GetMaxMemoryAvg

`func (o *GuidanceVmwareSizingConfig) GetMaxMemoryAvg() float32`

GetMaxMemoryAvg returns the MaxMemoryAvg field if non-nil, zero value otherwise.

### GetMaxMemoryAvgOk

`func (o *GuidanceVmwareSizingConfig) GetMaxMemoryAvgOk() (*float32, bool)`

GetMaxMemoryAvgOk returns a tuple with the MaxMemoryAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemoryAvg

`func (o *GuidanceVmwareSizingConfig) SetMaxMemoryAvg(v float32)`

SetMaxMemoryAvg sets MaxMemoryAvg field to given value.

### HasMaxMemoryAvg

`func (o *GuidanceVmwareSizingConfig) HasMaxMemoryAvg() bool`

HasMaxMemoryAvg returns a boolean if a field has been set.

### SetMaxMemoryAvgNil

`func (o *GuidanceVmwareSizingConfig) SetMaxMemoryAvgNil(b bool)`

 SetMaxMemoryAvgNil sets the value for MaxMemoryAvg to be an explicit nil

### UnsetMaxMemoryAvg
`func (o *GuidanceVmwareSizingConfig) UnsetMaxMemoryAvg()`

UnsetMaxMemoryAvg ensures that no value is present for MaxMemoryAvg, not even an explicit nil
### GetMaxMemorySum

`func (o *GuidanceVmwareSizingConfig) GetMaxMemorySum() float32`

GetMaxMemorySum returns the MaxMemorySum field if non-nil, zero value otherwise.

### GetMaxMemorySumOk

`func (o *GuidanceVmwareSizingConfig) GetMaxMemorySumOk() (*float32, bool)`

GetMaxMemorySumOk returns a tuple with the MaxMemorySum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemorySum

`func (o *GuidanceVmwareSizingConfig) SetMaxMemorySum(v float32)`

SetMaxMemorySum sets MaxMemorySum field to given value.

### HasMaxMemorySum

`func (o *GuidanceVmwareSizingConfig) HasMaxMemorySum() bool`

HasMaxMemorySum returns a boolean if a field has been set.

### SetMaxMemorySumNil

`func (o *GuidanceVmwareSizingConfig) SetMaxMemorySumNil(b bool)`

 SetMaxMemorySumNil sets the value for MaxMemorySum to be an explicit nil

### UnsetMaxMemorySum
`func (o *GuidanceVmwareSizingConfig) UnsetMaxMemorySum()`

UnsetMaxMemorySum ensures that no value is present for MaxMemorySum, not even an explicit nil
### GetCpuUserTimeCount

`func (o *GuidanceVmwareSizingConfig) GetCpuUserTimeCount() float32`

GetCpuUserTimeCount returns the CpuUserTimeCount field if non-nil, zero value otherwise.

### GetCpuUserTimeCountOk

`func (o *GuidanceVmwareSizingConfig) GetCpuUserTimeCountOk() (*float32, bool)`

GetCpuUserTimeCountOk returns a tuple with the CpuUserTimeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUserTimeCount

`func (o *GuidanceVmwareSizingConfig) SetCpuUserTimeCount(v float32)`

SetCpuUserTimeCount sets CpuUserTimeCount field to given value.

### HasCpuUserTimeCount

`func (o *GuidanceVmwareSizingConfig) HasCpuUserTimeCount() bool`

HasCpuUserTimeCount returns a boolean if a field has been set.

### SetCpuUserTimeCountNil

`func (o *GuidanceVmwareSizingConfig) SetCpuUserTimeCountNil(b bool)`

 SetCpuUserTimeCountNil sets the value for CpuUserTimeCount to be an explicit nil

### UnsetCpuUserTimeCount
`func (o *GuidanceVmwareSizingConfig) UnsetCpuUserTimeCount()`

UnsetCpuUserTimeCount ensures that no value is present for CpuUserTimeCount, not even an explicit nil
### GetCpuUserTimeMin

`func (o *GuidanceVmwareSizingConfig) GetCpuUserTimeMin() float32`

GetCpuUserTimeMin returns the CpuUserTimeMin field if non-nil, zero value otherwise.

### GetCpuUserTimeMinOk

`func (o *GuidanceVmwareSizingConfig) GetCpuUserTimeMinOk() (*float32, bool)`

GetCpuUserTimeMinOk returns a tuple with the CpuUserTimeMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUserTimeMin

`func (o *GuidanceVmwareSizingConfig) SetCpuUserTimeMin(v float32)`

SetCpuUserTimeMin sets CpuUserTimeMin field to given value.

### HasCpuUserTimeMin

`func (o *GuidanceVmwareSizingConfig) HasCpuUserTimeMin() bool`

HasCpuUserTimeMin returns a boolean if a field has been set.

### SetCpuUserTimeMinNil

`func (o *GuidanceVmwareSizingConfig) SetCpuUserTimeMinNil(b bool)`

 SetCpuUserTimeMinNil sets the value for CpuUserTimeMin to be an explicit nil

### UnsetCpuUserTimeMin
`func (o *GuidanceVmwareSizingConfig) UnsetCpuUserTimeMin()`

UnsetCpuUserTimeMin ensures that no value is present for CpuUserTimeMin, not even an explicit nil
### GetCpuUserTimeMax

`func (o *GuidanceVmwareSizingConfig) GetCpuUserTimeMax() float32`

GetCpuUserTimeMax returns the CpuUserTimeMax field if non-nil, zero value otherwise.

### GetCpuUserTimeMaxOk

`func (o *GuidanceVmwareSizingConfig) GetCpuUserTimeMaxOk() (*float32, bool)`

GetCpuUserTimeMaxOk returns a tuple with the CpuUserTimeMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUserTimeMax

`func (o *GuidanceVmwareSizingConfig) SetCpuUserTimeMax(v float32)`

SetCpuUserTimeMax sets CpuUserTimeMax field to given value.

### HasCpuUserTimeMax

`func (o *GuidanceVmwareSizingConfig) HasCpuUserTimeMax() bool`

HasCpuUserTimeMax returns a boolean if a field has been set.

### SetCpuUserTimeMaxNil

`func (o *GuidanceVmwareSizingConfig) SetCpuUserTimeMaxNil(b bool)`

 SetCpuUserTimeMaxNil sets the value for CpuUserTimeMax to be an explicit nil

### UnsetCpuUserTimeMax
`func (o *GuidanceVmwareSizingConfig) UnsetCpuUserTimeMax()`

UnsetCpuUserTimeMax ensures that no value is present for CpuUserTimeMax, not even an explicit nil
### GetCpuUserTimeAvg

`func (o *GuidanceVmwareSizingConfig) GetCpuUserTimeAvg() float32`

GetCpuUserTimeAvg returns the CpuUserTimeAvg field if non-nil, zero value otherwise.

### GetCpuUserTimeAvgOk

`func (o *GuidanceVmwareSizingConfig) GetCpuUserTimeAvgOk() (*float32, bool)`

GetCpuUserTimeAvgOk returns a tuple with the CpuUserTimeAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUserTimeAvg

`func (o *GuidanceVmwareSizingConfig) SetCpuUserTimeAvg(v float32)`

SetCpuUserTimeAvg sets CpuUserTimeAvg field to given value.

### HasCpuUserTimeAvg

`func (o *GuidanceVmwareSizingConfig) HasCpuUserTimeAvg() bool`

HasCpuUserTimeAvg returns a boolean if a field has been set.

### SetCpuUserTimeAvgNil

`func (o *GuidanceVmwareSizingConfig) SetCpuUserTimeAvgNil(b bool)`

 SetCpuUserTimeAvgNil sets the value for CpuUserTimeAvg to be an explicit nil

### UnsetCpuUserTimeAvg
`func (o *GuidanceVmwareSizingConfig) UnsetCpuUserTimeAvg()`

UnsetCpuUserTimeAvg ensures that no value is present for CpuUserTimeAvg, not even an explicit nil
### GetCpuUserTimeSum

`func (o *GuidanceVmwareSizingConfig) GetCpuUserTimeSum() float32`

GetCpuUserTimeSum returns the CpuUserTimeSum field if non-nil, zero value otherwise.

### GetCpuUserTimeSumOk

`func (o *GuidanceVmwareSizingConfig) GetCpuUserTimeSumOk() (*float32, bool)`

GetCpuUserTimeSumOk returns a tuple with the CpuUserTimeSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuUserTimeSum

`func (o *GuidanceVmwareSizingConfig) SetCpuUserTimeSum(v float32)`

SetCpuUserTimeSum sets CpuUserTimeSum field to given value.

### HasCpuUserTimeSum

`func (o *GuidanceVmwareSizingConfig) HasCpuUserTimeSum() bool`

HasCpuUserTimeSum returns a boolean if a field has been set.

### SetCpuUserTimeSumNil

`func (o *GuidanceVmwareSizingConfig) SetCpuUserTimeSumNil(b bool)`

 SetCpuUserTimeSumNil sets the value for CpuUserTimeSum to be an explicit nil

### UnsetCpuUserTimeSum
`func (o *GuidanceVmwareSizingConfig) UnsetCpuUserTimeSum()`

UnsetCpuUserTimeSum ensures that no value is present for CpuUserTimeSum, not even an explicit nil
### GetCpuSystemTimeCount

`func (o *GuidanceVmwareSizingConfig) GetCpuSystemTimeCount() float32`

GetCpuSystemTimeCount returns the CpuSystemTimeCount field if non-nil, zero value otherwise.

### GetCpuSystemTimeCountOk

`func (o *GuidanceVmwareSizingConfig) GetCpuSystemTimeCountOk() (*float32, bool)`

GetCpuSystemTimeCountOk returns a tuple with the CpuSystemTimeCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuSystemTimeCount

`func (o *GuidanceVmwareSizingConfig) SetCpuSystemTimeCount(v float32)`

SetCpuSystemTimeCount sets CpuSystemTimeCount field to given value.

### HasCpuSystemTimeCount

`func (o *GuidanceVmwareSizingConfig) HasCpuSystemTimeCount() bool`

HasCpuSystemTimeCount returns a boolean if a field has been set.

### SetCpuSystemTimeCountNil

`func (o *GuidanceVmwareSizingConfig) SetCpuSystemTimeCountNil(b bool)`

 SetCpuSystemTimeCountNil sets the value for CpuSystemTimeCount to be an explicit nil

### UnsetCpuSystemTimeCount
`func (o *GuidanceVmwareSizingConfig) UnsetCpuSystemTimeCount()`

UnsetCpuSystemTimeCount ensures that no value is present for CpuSystemTimeCount, not even an explicit nil
### GetCpuSystemTimeMin

`func (o *GuidanceVmwareSizingConfig) GetCpuSystemTimeMin() float32`

GetCpuSystemTimeMin returns the CpuSystemTimeMin field if non-nil, zero value otherwise.

### GetCpuSystemTimeMinOk

`func (o *GuidanceVmwareSizingConfig) GetCpuSystemTimeMinOk() (*float32, bool)`

GetCpuSystemTimeMinOk returns a tuple with the CpuSystemTimeMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuSystemTimeMin

`func (o *GuidanceVmwareSizingConfig) SetCpuSystemTimeMin(v float32)`

SetCpuSystemTimeMin sets CpuSystemTimeMin field to given value.

### HasCpuSystemTimeMin

`func (o *GuidanceVmwareSizingConfig) HasCpuSystemTimeMin() bool`

HasCpuSystemTimeMin returns a boolean if a field has been set.

### SetCpuSystemTimeMinNil

`func (o *GuidanceVmwareSizingConfig) SetCpuSystemTimeMinNil(b bool)`

 SetCpuSystemTimeMinNil sets the value for CpuSystemTimeMin to be an explicit nil

### UnsetCpuSystemTimeMin
`func (o *GuidanceVmwareSizingConfig) UnsetCpuSystemTimeMin()`

UnsetCpuSystemTimeMin ensures that no value is present for CpuSystemTimeMin, not even an explicit nil
### GetCpuSystemTimeMax

`func (o *GuidanceVmwareSizingConfig) GetCpuSystemTimeMax() float32`

GetCpuSystemTimeMax returns the CpuSystemTimeMax field if non-nil, zero value otherwise.

### GetCpuSystemTimeMaxOk

`func (o *GuidanceVmwareSizingConfig) GetCpuSystemTimeMaxOk() (*float32, bool)`

GetCpuSystemTimeMaxOk returns a tuple with the CpuSystemTimeMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuSystemTimeMax

`func (o *GuidanceVmwareSizingConfig) SetCpuSystemTimeMax(v float32)`

SetCpuSystemTimeMax sets CpuSystemTimeMax field to given value.

### HasCpuSystemTimeMax

`func (o *GuidanceVmwareSizingConfig) HasCpuSystemTimeMax() bool`

HasCpuSystemTimeMax returns a boolean if a field has been set.

### SetCpuSystemTimeMaxNil

`func (o *GuidanceVmwareSizingConfig) SetCpuSystemTimeMaxNil(b bool)`

 SetCpuSystemTimeMaxNil sets the value for CpuSystemTimeMax to be an explicit nil

### UnsetCpuSystemTimeMax
`func (o *GuidanceVmwareSizingConfig) UnsetCpuSystemTimeMax()`

UnsetCpuSystemTimeMax ensures that no value is present for CpuSystemTimeMax, not even an explicit nil
### GetCpuSystemTimeAvg

`func (o *GuidanceVmwareSizingConfig) GetCpuSystemTimeAvg() float32`

GetCpuSystemTimeAvg returns the CpuSystemTimeAvg field if non-nil, zero value otherwise.

### GetCpuSystemTimeAvgOk

`func (o *GuidanceVmwareSizingConfig) GetCpuSystemTimeAvgOk() (*float32, bool)`

GetCpuSystemTimeAvgOk returns a tuple with the CpuSystemTimeAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuSystemTimeAvg

`func (o *GuidanceVmwareSizingConfig) SetCpuSystemTimeAvg(v float32)`

SetCpuSystemTimeAvg sets CpuSystemTimeAvg field to given value.

### HasCpuSystemTimeAvg

`func (o *GuidanceVmwareSizingConfig) HasCpuSystemTimeAvg() bool`

HasCpuSystemTimeAvg returns a boolean if a field has been set.

### GetCpuSystemTimeSum

`func (o *GuidanceVmwareSizingConfig) GetCpuSystemTimeSum() float32`

GetCpuSystemTimeSum returns the CpuSystemTimeSum field if non-nil, zero value otherwise.

### GetCpuSystemTimeSumOk

`func (o *GuidanceVmwareSizingConfig) GetCpuSystemTimeSumOk() (*float32, bool)`

GetCpuSystemTimeSumOk returns a tuple with the CpuSystemTimeSum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpuSystemTimeSum

`func (o *GuidanceVmwareSizingConfig) SetCpuSystemTimeSum(v float32)`

SetCpuSystemTimeSum sets CpuSystemTimeSum field to given value.

### HasCpuSystemTimeSum

`func (o *GuidanceVmwareSizingConfig) HasCpuSystemTimeSum() bool`

HasCpuSystemTimeSum returns a boolean if a field has been set.

### SetCpuSystemTimeSumNil

`func (o *GuidanceVmwareSizingConfig) SetCpuSystemTimeSumNil(b bool)`

 SetCpuSystemTimeSumNil sets the value for CpuSystemTimeSum to be an explicit nil

### UnsetCpuSystemTimeSum
`func (o *GuidanceVmwareSizingConfig) UnsetCpuSystemTimeSum()`

UnsetCpuSystemTimeSum ensures that no value is present for CpuSystemTimeSum, not even an explicit nil
### GetUsedMemoryCount

`func (o *GuidanceVmwareSizingConfig) GetUsedMemoryCount() float32`

GetUsedMemoryCount returns the UsedMemoryCount field if non-nil, zero value otherwise.

### GetUsedMemoryCountOk

`func (o *GuidanceVmwareSizingConfig) GetUsedMemoryCountOk() (*float32, bool)`

GetUsedMemoryCountOk returns a tuple with the UsedMemoryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedMemoryCount

`func (o *GuidanceVmwareSizingConfig) SetUsedMemoryCount(v float32)`

SetUsedMemoryCount sets UsedMemoryCount field to given value.

### HasUsedMemoryCount

`func (o *GuidanceVmwareSizingConfig) HasUsedMemoryCount() bool`

HasUsedMemoryCount returns a boolean if a field has been set.

### SetUsedMemoryCountNil

`func (o *GuidanceVmwareSizingConfig) SetUsedMemoryCountNil(b bool)`

 SetUsedMemoryCountNil sets the value for UsedMemoryCount to be an explicit nil

### UnsetUsedMemoryCount
`func (o *GuidanceVmwareSizingConfig) UnsetUsedMemoryCount()`

UnsetUsedMemoryCount ensures that no value is present for UsedMemoryCount, not even an explicit nil
### GetUsedMemoryMin

`func (o *GuidanceVmwareSizingConfig) GetUsedMemoryMin() float32`

GetUsedMemoryMin returns the UsedMemoryMin field if non-nil, zero value otherwise.

### GetUsedMemoryMinOk

`func (o *GuidanceVmwareSizingConfig) GetUsedMemoryMinOk() (*float32, bool)`

GetUsedMemoryMinOk returns a tuple with the UsedMemoryMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedMemoryMin

`func (o *GuidanceVmwareSizingConfig) SetUsedMemoryMin(v float32)`

SetUsedMemoryMin sets UsedMemoryMin field to given value.

### HasUsedMemoryMin

`func (o *GuidanceVmwareSizingConfig) HasUsedMemoryMin() bool`

HasUsedMemoryMin returns a boolean if a field has been set.

### SetUsedMemoryMinNil

`func (o *GuidanceVmwareSizingConfig) SetUsedMemoryMinNil(b bool)`

 SetUsedMemoryMinNil sets the value for UsedMemoryMin to be an explicit nil

### UnsetUsedMemoryMin
`func (o *GuidanceVmwareSizingConfig) UnsetUsedMemoryMin()`

UnsetUsedMemoryMin ensures that no value is present for UsedMemoryMin, not even an explicit nil
### GetUsedMemoryMax

`func (o *GuidanceVmwareSizingConfig) GetUsedMemoryMax() float32`

GetUsedMemoryMax returns the UsedMemoryMax field if non-nil, zero value otherwise.

### GetUsedMemoryMaxOk

`func (o *GuidanceVmwareSizingConfig) GetUsedMemoryMaxOk() (*float32, bool)`

GetUsedMemoryMaxOk returns a tuple with the UsedMemoryMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedMemoryMax

`func (o *GuidanceVmwareSizingConfig) SetUsedMemoryMax(v float32)`

SetUsedMemoryMax sets UsedMemoryMax field to given value.

### HasUsedMemoryMax

`func (o *GuidanceVmwareSizingConfig) HasUsedMemoryMax() bool`

HasUsedMemoryMax returns a boolean if a field has been set.

### SetUsedMemoryMaxNil

`func (o *GuidanceVmwareSizingConfig) SetUsedMemoryMaxNil(b bool)`

 SetUsedMemoryMaxNil sets the value for UsedMemoryMax to be an explicit nil

### UnsetUsedMemoryMax
`func (o *GuidanceVmwareSizingConfig) UnsetUsedMemoryMax()`

UnsetUsedMemoryMax ensures that no value is present for UsedMemoryMax, not even an explicit nil
### GetUsedMemoryAvg

`func (o *GuidanceVmwareSizingConfig) GetUsedMemoryAvg() float32`

GetUsedMemoryAvg returns the UsedMemoryAvg field if non-nil, zero value otherwise.

### GetUsedMemoryAvgOk

`func (o *GuidanceVmwareSizingConfig) GetUsedMemoryAvgOk() (*float32, bool)`

GetUsedMemoryAvgOk returns a tuple with the UsedMemoryAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedMemoryAvg

`func (o *GuidanceVmwareSizingConfig) SetUsedMemoryAvg(v float32)`

SetUsedMemoryAvg sets UsedMemoryAvg field to given value.

### HasUsedMemoryAvg

`func (o *GuidanceVmwareSizingConfig) HasUsedMemoryAvg() bool`

HasUsedMemoryAvg returns a boolean if a field has been set.

### SetUsedMemoryAvgNil

`func (o *GuidanceVmwareSizingConfig) SetUsedMemoryAvgNil(b bool)`

 SetUsedMemoryAvgNil sets the value for UsedMemoryAvg to be an explicit nil

### UnsetUsedMemoryAvg
`func (o *GuidanceVmwareSizingConfig) UnsetUsedMemoryAvg()`

UnsetUsedMemoryAvg ensures that no value is present for UsedMemoryAvg, not even an explicit nil
### GetUsedMemorySum

`func (o *GuidanceVmwareSizingConfig) GetUsedMemorySum() float32`

GetUsedMemorySum returns the UsedMemorySum field if non-nil, zero value otherwise.

### GetUsedMemorySumOk

`func (o *GuidanceVmwareSizingConfig) GetUsedMemorySumOk() (*float32, bool)`

GetUsedMemorySumOk returns a tuple with the UsedMemorySum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsedMemorySum

`func (o *GuidanceVmwareSizingConfig) SetUsedMemorySum(v float32)`

SetUsedMemorySum sets UsedMemorySum field to given value.

### HasUsedMemorySum

`func (o *GuidanceVmwareSizingConfig) HasUsedMemorySum() bool`

HasUsedMemorySum returns a boolean if a field has been set.

### SetUsedMemorySumNil

`func (o *GuidanceVmwareSizingConfig) SetUsedMemorySumNil(b bool)`

 SetUsedMemorySumNil sets the value for UsedMemorySum to be an explicit nil

### UnsetUsedMemorySum
`func (o *GuidanceVmwareSizingConfig) UnsetUsedMemorySum()`

UnsetUsedMemorySum ensures that no value is present for UsedMemorySum, not even an explicit nil
### GetFreeMemoryCount

`func (o *GuidanceVmwareSizingConfig) GetFreeMemoryCount() float32`

GetFreeMemoryCount returns the FreeMemoryCount field if non-nil, zero value otherwise.

### GetFreeMemoryCountOk

`func (o *GuidanceVmwareSizingConfig) GetFreeMemoryCountOk() (*float32, bool)`

GetFreeMemoryCountOk returns a tuple with the FreeMemoryCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeMemoryCount

`func (o *GuidanceVmwareSizingConfig) SetFreeMemoryCount(v float32)`

SetFreeMemoryCount sets FreeMemoryCount field to given value.

### HasFreeMemoryCount

`func (o *GuidanceVmwareSizingConfig) HasFreeMemoryCount() bool`

HasFreeMemoryCount returns a boolean if a field has been set.

### SetFreeMemoryCountNil

`func (o *GuidanceVmwareSizingConfig) SetFreeMemoryCountNil(b bool)`

 SetFreeMemoryCountNil sets the value for FreeMemoryCount to be an explicit nil

### UnsetFreeMemoryCount
`func (o *GuidanceVmwareSizingConfig) UnsetFreeMemoryCount()`

UnsetFreeMemoryCount ensures that no value is present for FreeMemoryCount, not even an explicit nil
### GetFreeMemoryMin

`func (o *GuidanceVmwareSizingConfig) GetFreeMemoryMin() float32`

GetFreeMemoryMin returns the FreeMemoryMin field if non-nil, zero value otherwise.

### GetFreeMemoryMinOk

`func (o *GuidanceVmwareSizingConfig) GetFreeMemoryMinOk() (*float32, bool)`

GetFreeMemoryMinOk returns a tuple with the FreeMemoryMin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeMemoryMin

`func (o *GuidanceVmwareSizingConfig) SetFreeMemoryMin(v float32)`

SetFreeMemoryMin sets FreeMemoryMin field to given value.

### HasFreeMemoryMin

`func (o *GuidanceVmwareSizingConfig) HasFreeMemoryMin() bool`

HasFreeMemoryMin returns a boolean if a field has been set.

### SetFreeMemoryMinNil

`func (o *GuidanceVmwareSizingConfig) SetFreeMemoryMinNil(b bool)`

 SetFreeMemoryMinNil sets the value for FreeMemoryMin to be an explicit nil

### UnsetFreeMemoryMin
`func (o *GuidanceVmwareSizingConfig) UnsetFreeMemoryMin()`

UnsetFreeMemoryMin ensures that no value is present for FreeMemoryMin, not even an explicit nil
### GetFreeMemoryMax

`func (o *GuidanceVmwareSizingConfig) GetFreeMemoryMax() float32`

GetFreeMemoryMax returns the FreeMemoryMax field if non-nil, zero value otherwise.

### GetFreeMemoryMaxOk

`func (o *GuidanceVmwareSizingConfig) GetFreeMemoryMaxOk() (*float32, bool)`

GetFreeMemoryMaxOk returns a tuple with the FreeMemoryMax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeMemoryMax

`func (o *GuidanceVmwareSizingConfig) SetFreeMemoryMax(v float32)`

SetFreeMemoryMax sets FreeMemoryMax field to given value.

### HasFreeMemoryMax

`func (o *GuidanceVmwareSizingConfig) HasFreeMemoryMax() bool`

HasFreeMemoryMax returns a boolean if a field has been set.

### SetFreeMemoryMaxNil

`func (o *GuidanceVmwareSizingConfig) SetFreeMemoryMaxNil(b bool)`

 SetFreeMemoryMaxNil sets the value for FreeMemoryMax to be an explicit nil

### UnsetFreeMemoryMax
`func (o *GuidanceVmwareSizingConfig) UnsetFreeMemoryMax()`

UnsetFreeMemoryMax ensures that no value is present for FreeMemoryMax, not even an explicit nil
### GetFreeMemoryAvg

`func (o *GuidanceVmwareSizingConfig) GetFreeMemoryAvg() float32`

GetFreeMemoryAvg returns the FreeMemoryAvg field if non-nil, zero value otherwise.

### GetFreeMemoryAvgOk

`func (o *GuidanceVmwareSizingConfig) GetFreeMemoryAvgOk() (*float32, bool)`

GetFreeMemoryAvgOk returns a tuple with the FreeMemoryAvg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeMemoryAvg

`func (o *GuidanceVmwareSizingConfig) SetFreeMemoryAvg(v float32)`

SetFreeMemoryAvg sets FreeMemoryAvg field to given value.

### HasFreeMemoryAvg

`func (o *GuidanceVmwareSizingConfig) HasFreeMemoryAvg() bool`

HasFreeMemoryAvg returns a boolean if a field has been set.

### SetFreeMemoryAvgNil

`func (o *GuidanceVmwareSizingConfig) SetFreeMemoryAvgNil(b bool)`

 SetFreeMemoryAvgNil sets the value for FreeMemoryAvg to be an explicit nil

### UnsetFreeMemoryAvg
`func (o *GuidanceVmwareSizingConfig) UnsetFreeMemoryAvg()`

UnsetFreeMemoryAvg ensures that no value is present for FreeMemoryAvg, not even an explicit nil
### GetFreeMemorySum

`func (o *GuidanceVmwareSizingConfig) GetFreeMemorySum() float32`

GetFreeMemorySum returns the FreeMemorySum field if non-nil, zero value otherwise.

### GetFreeMemorySumOk

`func (o *GuidanceVmwareSizingConfig) GetFreeMemorySumOk() (*float32, bool)`

GetFreeMemorySumOk returns a tuple with the FreeMemorySum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFreeMemorySum

`func (o *GuidanceVmwareSizingConfig) SetFreeMemorySum(v float32)`

SetFreeMemorySum sets FreeMemorySum field to given value.

### HasFreeMemorySum

`func (o *GuidanceVmwareSizingConfig) HasFreeMemorySum() bool`

HasFreeMemorySum returns a boolean if a field has been set.

### SetFreeMemorySumNil

`func (o *GuidanceVmwareSizingConfig) SetFreeMemorySumNil(b bool)`

 SetFreeMemorySumNil sets the value for FreeMemorySum to be an explicit nil

### UnsetFreeMemorySum
`func (o *GuidanceVmwareSizingConfig) UnsetFreeMemorySum()`

UnsetFreeMemorySum ensures that no value is present for FreeMemorySum, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


