# MaxMemoryPolicyTypeConfiguration2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxMemory** | Pointer to [**MaxMemoryPolicyTypeConfiguration1MaxMemory**](MaxMemoryPolicyTypeConfiguration1MaxMemory.md) |  | [optional] 
**ExcludeContainers** | Pointer to **string** | Set to on to exclude containers | [optional] [default to "off"]

## Methods

### NewMaxMemoryPolicyTypeConfiguration2

`func NewMaxMemoryPolicyTypeConfiguration2() *MaxMemoryPolicyTypeConfiguration2`

NewMaxMemoryPolicyTypeConfiguration2 instantiates a new MaxMemoryPolicyTypeConfiguration2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaxMemoryPolicyTypeConfiguration2WithDefaults

`func NewMaxMemoryPolicyTypeConfiguration2WithDefaults() *MaxMemoryPolicyTypeConfiguration2`

NewMaxMemoryPolicyTypeConfiguration2WithDefaults instantiates a new MaxMemoryPolicyTypeConfiguration2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxMemory

`func (o *MaxMemoryPolicyTypeConfiguration2) GetMaxMemory() MaxMemoryPolicyTypeConfiguration1MaxMemory`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *MaxMemoryPolicyTypeConfiguration2) GetMaxMemoryOk() (*MaxMemoryPolicyTypeConfiguration1MaxMemory, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *MaxMemoryPolicyTypeConfiguration2) SetMaxMemory(v MaxMemoryPolicyTypeConfiguration1MaxMemory)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *MaxMemoryPolicyTypeConfiguration2) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetExcludeContainers

`func (o *MaxMemoryPolicyTypeConfiguration2) GetExcludeContainers() string`

GetExcludeContainers returns the ExcludeContainers field if non-nil, zero value otherwise.

### GetExcludeContainersOk

`func (o *MaxMemoryPolicyTypeConfiguration2) GetExcludeContainersOk() (*string, bool)`

GetExcludeContainersOk returns a tuple with the ExcludeContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeContainers

`func (o *MaxMemoryPolicyTypeConfiguration2) SetExcludeContainers(v string)`

SetExcludeContainers sets ExcludeContainers field to given value.

### HasExcludeContainers

`func (o *MaxMemoryPolicyTypeConfiguration2) HasExcludeContainers() bool`

HasExcludeContainers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


