# MaxMemoryPolicyTypeConfiguration8

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxMemory** | Pointer to [**MaxMemoryPolicyTypeConfiguration8MaxMemory**](MaxMemoryPolicyTypeConfiguration8MaxMemory.md) |  | [optional] 
**ExcludeContainers** | Pointer to **string** | Set to on to exclude containers | [optional] [default to "off"]

## Methods

### NewMaxMemoryPolicyTypeConfiguration8

`func NewMaxMemoryPolicyTypeConfiguration8() *MaxMemoryPolicyTypeConfiguration8`

NewMaxMemoryPolicyTypeConfiguration8 instantiates a new MaxMemoryPolicyTypeConfiguration8 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaxMemoryPolicyTypeConfiguration8WithDefaults

`func NewMaxMemoryPolicyTypeConfiguration8WithDefaults() *MaxMemoryPolicyTypeConfiguration8`

NewMaxMemoryPolicyTypeConfiguration8WithDefaults instantiates a new MaxMemoryPolicyTypeConfiguration8 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxMemory

`func (o *MaxMemoryPolicyTypeConfiguration8) GetMaxMemory() MaxMemoryPolicyTypeConfiguration8MaxMemory`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *MaxMemoryPolicyTypeConfiguration8) GetMaxMemoryOk() (*MaxMemoryPolicyTypeConfiguration8MaxMemory, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *MaxMemoryPolicyTypeConfiguration8) SetMaxMemory(v MaxMemoryPolicyTypeConfiguration8MaxMemory)`

SetMaxMemory sets MaxMemory field to given value.

### HasMaxMemory

`func (o *MaxMemoryPolicyTypeConfiguration8) HasMaxMemory() bool`

HasMaxMemory returns a boolean if a field has been set.

### GetExcludeContainers

`func (o *MaxMemoryPolicyTypeConfiguration8) GetExcludeContainers() string`

GetExcludeContainers returns the ExcludeContainers field if non-nil, zero value otherwise.

### GetExcludeContainersOk

`func (o *MaxMemoryPolicyTypeConfiguration8) GetExcludeContainersOk() (*string, bool)`

GetExcludeContainersOk returns a tuple with the ExcludeContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeContainers

`func (o *MaxMemoryPolicyTypeConfiguration8) SetExcludeContainers(v string)`

SetExcludeContainers sets ExcludeContainers field to given value.

### HasExcludeContainers

`func (o *MaxMemoryPolicyTypeConfiguration8) HasExcludeContainers() bool`

HasExcludeContainers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


