# ListHealth200ResponseAllOfHealth

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**StatusMessage** | Pointer to **string** |  | [optional] 
**ApplianceUrl** | Pointer to **string** |  | [optional] 
**BuildVersion** | Pointer to **string** |  | [optional] 
**Uuid** | Pointer to **string** |  | [optional] 
**SetupNeeded** | Pointer to **bool** |  | [optional] 
**Date** | Pointer to **time.Time** |  | [optional] 
**Cpu** | Pointer to [**ListHealth200ResponseAllOfHealthCpu**](ListHealth200ResponseAllOfHealthCpu.md) |  | [optional] 
**Memory** | Pointer to [**ListHealth200ResponseAllOfHealthMemory**](ListHealth200ResponseAllOfHealthMemory.md) |  | [optional] 
**Threads** | Pointer to [**ListHealth200ResponseAllOfHealthThreads**](ListHealth200ResponseAllOfHealthThreads.md) |  | [optional] 
**Database** | Pointer to [**ListHealth200ResponseAllOfHealthDatabase**](ListHealth200ResponseAllOfHealthDatabase.md) |  | [optional] 
**Elastic** | Pointer to [**ListHealth200ResponseAllOfHealthElastic**](ListHealth200ResponseAllOfHealthElastic.md) |  | [optional] 
**Rabbit** | Pointer to [**ListHealth200ResponseAllOfHealthRabbit**](ListHealth200ResponseAllOfHealthRabbit.md) |  | [optional] 

## Methods

### NewListHealth200ResponseAllOfHealth

`func NewListHealth200ResponseAllOfHealth() *ListHealth200ResponseAllOfHealth`

NewListHealth200ResponseAllOfHealth instantiates a new ListHealth200ResponseAllOfHealth object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListHealth200ResponseAllOfHealthWithDefaults

`func NewListHealth200ResponseAllOfHealthWithDefaults() *ListHealth200ResponseAllOfHealth`

NewListHealth200ResponseAllOfHealthWithDefaults instantiates a new ListHealth200ResponseAllOfHealth object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ListHealth200ResponseAllOfHealth) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ListHealth200ResponseAllOfHealth) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ListHealth200ResponseAllOfHealth) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ListHealth200ResponseAllOfHealth) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ListHealth200ResponseAllOfHealth) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ListHealth200ResponseAllOfHealth) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ListHealth200ResponseAllOfHealth) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ListHealth200ResponseAllOfHealth) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetApplianceUrl

`func (o *ListHealth200ResponseAllOfHealth) GetApplianceUrl() string`

GetApplianceUrl returns the ApplianceUrl field if non-nil, zero value otherwise.

### GetApplianceUrlOk

`func (o *ListHealth200ResponseAllOfHealth) GetApplianceUrlOk() (*string, bool)`

GetApplianceUrlOk returns a tuple with the ApplianceUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplianceUrl

`func (o *ListHealth200ResponseAllOfHealth) SetApplianceUrl(v string)`

SetApplianceUrl sets ApplianceUrl field to given value.

### HasApplianceUrl

`func (o *ListHealth200ResponseAllOfHealth) HasApplianceUrl() bool`

HasApplianceUrl returns a boolean if a field has been set.

### GetBuildVersion

`func (o *ListHealth200ResponseAllOfHealth) GetBuildVersion() string`

GetBuildVersion returns the BuildVersion field if non-nil, zero value otherwise.

### GetBuildVersionOk

`func (o *ListHealth200ResponseAllOfHealth) GetBuildVersionOk() (*string, bool)`

GetBuildVersionOk returns a tuple with the BuildVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBuildVersion

`func (o *ListHealth200ResponseAllOfHealth) SetBuildVersion(v string)`

SetBuildVersion sets BuildVersion field to given value.

### HasBuildVersion

`func (o *ListHealth200ResponseAllOfHealth) HasBuildVersion() bool`

HasBuildVersion returns a boolean if a field has been set.

### GetUuid

`func (o *ListHealth200ResponseAllOfHealth) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ListHealth200ResponseAllOfHealth) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ListHealth200ResponseAllOfHealth) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ListHealth200ResponseAllOfHealth) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetSetupNeeded

`func (o *ListHealth200ResponseAllOfHealth) GetSetupNeeded() bool`

GetSetupNeeded returns the SetupNeeded field if non-nil, zero value otherwise.

### GetSetupNeededOk

`func (o *ListHealth200ResponseAllOfHealth) GetSetupNeededOk() (*bool, bool)`

GetSetupNeededOk returns a tuple with the SetupNeeded field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSetupNeeded

`func (o *ListHealth200ResponseAllOfHealth) SetSetupNeeded(v bool)`

SetSetupNeeded sets SetupNeeded field to given value.

### HasSetupNeeded

`func (o *ListHealth200ResponseAllOfHealth) HasSetupNeeded() bool`

HasSetupNeeded returns a boolean if a field has been set.

### GetDate

`func (o *ListHealth200ResponseAllOfHealth) GetDate() time.Time`

GetDate returns the Date field if non-nil, zero value otherwise.

### GetDateOk

`func (o *ListHealth200ResponseAllOfHealth) GetDateOk() (*time.Time, bool)`

GetDateOk returns a tuple with the Date field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDate

`func (o *ListHealth200ResponseAllOfHealth) SetDate(v time.Time)`

SetDate sets Date field to given value.

### HasDate

`func (o *ListHealth200ResponseAllOfHealth) HasDate() bool`

HasDate returns a boolean if a field has been set.

### GetCpu

`func (o *ListHealth200ResponseAllOfHealth) GetCpu() ListHealth200ResponseAllOfHealthCpu`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *ListHealth200ResponseAllOfHealth) GetCpuOk() (*ListHealth200ResponseAllOfHealthCpu, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *ListHealth200ResponseAllOfHealth) SetCpu(v ListHealth200ResponseAllOfHealthCpu)`

SetCpu sets Cpu field to given value.

### HasCpu

`func (o *ListHealth200ResponseAllOfHealth) HasCpu() bool`

HasCpu returns a boolean if a field has been set.

### GetMemory

`func (o *ListHealth200ResponseAllOfHealth) GetMemory() ListHealth200ResponseAllOfHealthMemory`

GetMemory returns the Memory field if non-nil, zero value otherwise.

### GetMemoryOk

`func (o *ListHealth200ResponseAllOfHealth) GetMemoryOk() (*ListHealth200ResponseAllOfHealthMemory, bool)`

GetMemoryOk returns a tuple with the Memory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemory

`func (o *ListHealth200ResponseAllOfHealth) SetMemory(v ListHealth200ResponseAllOfHealthMemory)`

SetMemory sets Memory field to given value.

### HasMemory

`func (o *ListHealth200ResponseAllOfHealth) HasMemory() bool`

HasMemory returns a boolean if a field has been set.

### GetThreads

`func (o *ListHealth200ResponseAllOfHealth) GetThreads() ListHealth200ResponseAllOfHealthThreads`

GetThreads returns the Threads field if non-nil, zero value otherwise.

### GetThreadsOk

`func (o *ListHealth200ResponseAllOfHealth) GetThreadsOk() (*ListHealth200ResponseAllOfHealthThreads, bool)`

GetThreadsOk returns a tuple with the Threads field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetThreads

`func (o *ListHealth200ResponseAllOfHealth) SetThreads(v ListHealth200ResponseAllOfHealthThreads)`

SetThreads sets Threads field to given value.

### HasThreads

`func (o *ListHealth200ResponseAllOfHealth) HasThreads() bool`

HasThreads returns a boolean if a field has been set.

### GetDatabase

`func (o *ListHealth200ResponseAllOfHealth) GetDatabase() ListHealth200ResponseAllOfHealthDatabase`

GetDatabase returns the Database field if non-nil, zero value otherwise.

### GetDatabaseOk

`func (o *ListHealth200ResponseAllOfHealth) GetDatabaseOk() (*ListHealth200ResponseAllOfHealthDatabase, bool)`

GetDatabaseOk returns a tuple with the Database field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatabase

`func (o *ListHealth200ResponseAllOfHealth) SetDatabase(v ListHealth200ResponseAllOfHealthDatabase)`

SetDatabase sets Database field to given value.

### HasDatabase

`func (o *ListHealth200ResponseAllOfHealth) HasDatabase() bool`

HasDatabase returns a boolean if a field has been set.

### GetElastic

`func (o *ListHealth200ResponseAllOfHealth) GetElastic() ListHealth200ResponseAllOfHealthElastic`

GetElastic returns the Elastic field if non-nil, zero value otherwise.

### GetElasticOk

`func (o *ListHealth200ResponseAllOfHealth) GetElasticOk() (*ListHealth200ResponseAllOfHealthElastic, bool)`

GetElasticOk returns a tuple with the Elastic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetElastic

`func (o *ListHealth200ResponseAllOfHealth) SetElastic(v ListHealth200ResponseAllOfHealthElastic)`

SetElastic sets Elastic field to given value.

### HasElastic

`func (o *ListHealth200ResponseAllOfHealth) HasElastic() bool`

HasElastic returns a boolean if a field has been set.

### GetRabbit

`func (o *ListHealth200ResponseAllOfHealth) GetRabbit() ListHealth200ResponseAllOfHealthRabbit`

GetRabbit returns the Rabbit field if non-nil, zero value otherwise.

### GetRabbitOk

`func (o *ListHealth200ResponseAllOfHealth) GetRabbitOk() (*ListHealth200ResponseAllOfHealthRabbit, bool)`

GetRabbitOk returns a tuple with the Rabbit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRabbit

`func (o *ListHealth200ResponseAllOfHealth) SetRabbit(v ListHealth200ResponseAllOfHealthRabbit)`

SetRabbit sets Rabbit field to given value.

### HasRabbit

`func (o *ListHealth200ResponseAllOfHealth) HasRabbit() bool`

HasRabbit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


