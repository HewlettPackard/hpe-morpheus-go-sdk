# HealthRabbit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**BusyQueues** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ErrorQueues** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Queues** | Pointer to [**[]HealthRabbitQueuesInner**](HealthRabbitQueuesInner.md) |  | [optional] 

## Methods

### NewHealthRabbit

`func NewHealthRabbit() *HealthRabbit`

NewHealthRabbit instantiates a new HealthRabbit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHealthRabbitWithDefaults

`func NewHealthRabbitWithDefaults() *HealthRabbit`

NewHealthRabbitWithDefaults instantiates a new HealthRabbit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *HealthRabbit) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *HealthRabbit) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *HealthRabbit) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *HealthRabbit) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetBusyQueues

`func (o *HealthRabbit) GetBusyQueues() []map[string]interface{}`

GetBusyQueues returns the BusyQueues field if non-nil, zero value otherwise.

### GetBusyQueuesOk

`func (o *HealthRabbit) GetBusyQueuesOk() (*[]map[string]interface{}, bool)`

GetBusyQueuesOk returns a tuple with the BusyQueues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusyQueues

`func (o *HealthRabbit) SetBusyQueues(v []map[string]interface{})`

SetBusyQueues sets BusyQueues field to given value.

### HasBusyQueues

`func (o *HealthRabbit) HasBusyQueues() bool`

HasBusyQueues returns a boolean if a field has been set.

### SetBusyQueuesNil

`func (o *HealthRabbit) SetBusyQueuesNil(b bool)`

 SetBusyQueuesNil sets the value for BusyQueues to be an explicit nil

### UnsetBusyQueues
`func (o *HealthRabbit) UnsetBusyQueues()`

UnsetBusyQueues ensures that no value is present for BusyQueues, not even an explicit nil
### GetErrorQueues

`func (o *HealthRabbit) GetErrorQueues() []map[string]interface{}`

GetErrorQueues returns the ErrorQueues field if non-nil, zero value otherwise.

### GetErrorQueuesOk

`func (o *HealthRabbit) GetErrorQueuesOk() (*[]map[string]interface{}, bool)`

GetErrorQueuesOk returns a tuple with the ErrorQueues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorQueues

`func (o *HealthRabbit) SetErrorQueues(v []map[string]interface{})`

SetErrorQueues sets ErrorQueues field to given value.

### HasErrorQueues

`func (o *HealthRabbit) HasErrorQueues() bool`

HasErrorQueues returns a boolean if a field has been set.

### SetErrorQueuesNil

`func (o *HealthRabbit) SetErrorQueuesNil(b bool)`

 SetErrorQueuesNil sets the value for ErrorQueues to be an explicit nil

### UnsetErrorQueues
`func (o *HealthRabbit) UnsetErrorQueues()`

UnsetErrorQueues ensures that no value is present for ErrorQueues, not even an explicit nil
### GetStatus

`func (o *HealthRabbit) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *HealthRabbit) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *HealthRabbit) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *HealthRabbit) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetQueues

`func (o *HealthRabbit) GetQueues() []HealthRabbitQueuesInner`

GetQueues returns the Queues field if non-nil, zero value otherwise.

### GetQueuesOk

`func (o *HealthRabbit) GetQueuesOk() (*[]HealthRabbitQueuesInner, bool)`

GetQueuesOk returns a tuple with the Queues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueues

`func (o *HealthRabbit) SetQueues(v []HealthRabbitQueuesInner)`

SetQueues sets Queues field to given value.

### HasQueues

`func (o *HealthRabbit) HasQueues() bool`

HasQueues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


