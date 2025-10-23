# MaxMemoryPolicyTypeConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxMemory** | [**MaxMemoryPolicyTypeConfigurationMaxMemory**](MaxMemoryPolicyTypeConfigurationMaxMemory.md) |  | 
**ExcludeContainers** | Pointer to **string** |  | [optional] [default to "off"]

## Methods

### NewMaxMemoryPolicyTypeConfiguration

`func NewMaxMemoryPolicyTypeConfiguration(maxMemory MaxMemoryPolicyTypeConfigurationMaxMemory, ) *MaxMemoryPolicyTypeConfiguration`

NewMaxMemoryPolicyTypeConfiguration instantiates a new MaxMemoryPolicyTypeConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMaxMemoryPolicyTypeConfigurationWithDefaults

`func NewMaxMemoryPolicyTypeConfigurationWithDefaults() *MaxMemoryPolicyTypeConfiguration`

NewMaxMemoryPolicyTypeConfigurationWithDefaults instantiates a new MaxMemoryPolicyTypeConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxMemory

`func (o *MaxMemoryPolicyTypeConfiguration) GetMaxMemory() MaxMemoryPolicyTypeConfigurationMaxMemory`

GetMaxMemory returns the MaxMemory field if non-nil, zero value otherwise.

### GetMaxMemoryOk

`func (o *MaxMemoryPolicyTypeConfiguration) GetMaxMemoryOk() (*MaxMemoryPolicyTypeConfigurationMaxMemory, bool)`

GetMaxMemoryOk returns a tuple with the MaxMemory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMemory

`func (o *MaxMemoryPolicyTypeConfiguration) SetMaxMemory(v MaxMemoryPolicyTypeConfigurationMaxMemory)`

SetMaxMemory sets MaxMemory field to given value.


### GetExcludeContainers

`func (o *MaxMemoryPolicyTypeConfiguration) GetExcludeContainers() string`

GetExcludeContainers returns the ExcludeContainers field if non-nil, zero value otherwise.

### GetExcludeContainersOk

`func (o *MaxMemoryPolicyTypeConfiguration) GetExcludeContainersOk() (*string, bool)`

GetExcludeContainersOk returns a tuple with the ExcludeContainers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExcludeContainers

`func (o *MaxMemoryPolicyTypeConfiguration) SetExcludeContainers(v string)`

SetExcludeContainers sets ExcludeContainers field to given value.

### HasExcludeContainers

`func (o *MaxMemoryPolicyTypeConfiguration) HasExcludeContainers() bool`

HasExcludeContainers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


