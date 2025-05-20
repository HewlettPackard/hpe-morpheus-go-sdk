# ListHealth200ResponseAllOfHealthRabbit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**BusyQueues** | Pointer to **[]map[string]interface{}** |  | [optional] 
**ErrorQueues** | Pointer to **[]map[string]interface{}** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Queues** | Pointer to [**[]ListHealth200ResponseAllOfHealthRabbitQueuesInner**](ListHealth200ResponseAllOfHealthRabbitQueuesInner.md) |  | [optional] 

## Methods

### NewListHealth200ResponseAllOfHealthRabbit

`func NewListHealth200ResponseAllOfHealthRabbit() *ListHealth200ResponseAllOfHealthRabbit`

NewListHealth200ResponseAllOfHealthRabbit instantiates a new ListHealth200ResponseAllOfHealthRabbit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListHealth200ResponseAllOfHealthRabbitWithDefaults

`func NewListHealth200ResponseAllOfHealthRabbitWithDefaults() *ListHealth200ResponseAllOfHealthRabbit`

NewListHealth200ResponseAllOfHealthRabbitWithDefaults instantiates a new ListHealth200ResponseAllOfHealthRabbit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ListHealth200ResponseAllOfHealthRabbit) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ListHealth200ResponseAllOfHealthRabbit) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ListHealth200ResponseAllOfHealthRabbit) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ListHealth200ResponseAllOfHealthRabbit) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetBusyQueues

`func (o *ListHealth200ResponseAllOfHealthRabbit) GetBusyQueues() []map[string]interface{}`

GetBusyQueues returns the BusyQueues field if non-nil, zero value otherwise.

### GetBusyQueuesOk

`func (o *ListHealth200ResponseAllOfHealthRabbit) GetBusyQueuesOk() (*[]map[string]interface{}, bool)`

GetBusyQueuesOk returns a tuple with the BusyQueues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusyQueues

`func (o *ListHealth200ResponseAllOfHealthRabbit) SetBusyQueues(v []map[string]interface{})`

SetBusyQueues sets BusyQueues field to given value.

### HasBusyQueues

`func (o *ListHealth200ResponseAllOfHealthRabbit) HasBusyQueues() bool`

HasBusyQueues returns a boolean if a field has been set.

### GetErrorQueues

`func (o *ListHealth200ResponseAllOfHealthRabbit) GetErrorQueues() []map[string]interface{}`

GetErrorQueues returns the ErrorQueues field if non-nil, zero value otherwise.

### GetErrorQueuesOk

`func (o *ListHealth200ResponseAllOfHealthRabbit) GetErrorQueuesOk() (*[]map[string]interface{}, bool)`

GetErrorQueuesOk returns a tuple with the ErrorQueues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorQueues

`func (o *ListHealth200ResponseAllOfHealthRabbit) SetErrorQueues(v []map[string]interface{})`

SetErrorQueues sets ErrorQueues field to given value.

### HasErrorQueues

`func (o *ListHealth200ResponseAllOfHealthRabbit) HasErrorQueues() bool`

HasErrorQueues returns a boolean if a field has been set.

### GetStatus

`func (o *ListHealth200ResponseAllOfHealthRabbit) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListHealth200ResponseAllOfHealthRabbit) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListHealth200ResponseAllOfHealthRabbit) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListHealth200ResponseAllOfHealthRabbit) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetQueues

`func (o *ListHealth200ResponseAllOfHealthRabbit) GetQueues() []ListHealth200ResponseAllOfHealthRabbitQueuesInner`

GetQueues returns the Queues field if non-nil, zero value otherwise.

### GetQueuesOk

`func (o *ListHealth200ResponseAllOfHealthRabbit) GetQueuesOk() (*[]ListHealth200ResponseAllOfHealthRabbitQueuesInner, bool)`

GetQueuesOk returns a tuple with the Queues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueues

`func (o *ListHealth200ResponseAllOfHealthRabbit) SetQueues(v []ListHealth200ResponseAllOfHealthRabbitQueuesInner)`

SetQueues sets Queues field to given value.

### HasQueues

`func (o *ListHealth200ResponseAllOfHealthRabbit) HasQueues() bool`

HasQueues returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


