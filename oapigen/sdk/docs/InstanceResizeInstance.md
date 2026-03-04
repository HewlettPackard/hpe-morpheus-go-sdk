# InstanceResizeInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The id of the instance to resize. | [optional] 
**Plan** | Pointer to [**InstanceResizeInstancePlan**](InstanceResizeInstancePlan.md) |  | [optional] 

## Methods

### NewInstanceResizeInstance

`func NewInstanceResizeInstance() *InstanceResizeInstance`

NewInstanceResizeInstance instantiates a new InstanceResizeInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceResizeInstanceWithDefaults

`func NewInstanceResizeInstanceWithDefaults() *InstanceResizeInstance`

NewInstanceResizeInstanceWithDefaults instantiates a new InstanceResizeInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *InstanceResizeInstance) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InstanceResizeInstance) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InstanceResizeInstance) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *InstanceResizeInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlan

`func (o *InstanceResizeInstance) GetPlan() InstanceResizeInstancePlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *InstanceResizeInstance) GetPlanOk() (*InstanceResizeInstancePlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *InstanceResizeInstance) SetPlan(v InstanceResizeInstancePlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *InstanceResizeInstance) HasPlan() bool`

HasPlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


