# ResizeInstanceRequestInstance

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | The id of the instance to resize. | [optional] 
**Plan** | Pointer to [**ResizeInstanceRequestInstancePlan**](ResizeInstanceRequestInstancePlan.md) |  | [optional] 

## Methods

### NewResizeInstanceRequestInstance

`func NewResizeInstanceRequestInstance() *ResizeInstanceRequestInstance`

NewResizeInstanceRequestInstance instantiates a new ResizeInstanceRequestInstance object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResizeInstanceRequestInstanceWithDefaults

`func NewResizeInstanceRequestInstanceWithDefaults() *ResizeInstanceRequestInstance`

NewResizeInstanceRequestInstanceWithDefaults instantiates a new ResizeInstanceRequestInstance object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ResizeInstanceRequestInstance) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResizeInstanceRequestInstance) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResizeInstanceRequestInstance) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ResizeInstanceRequestInstance) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPlan

`func (o *ResizeInstanceRequestInstance) GetPlan() ResizeInstanceRequestInstancePlan`

GetPlan returns the Plan field if non-nil, zero value otherwise.

### GetPlanOk

`func (o *ResizeInstanceRequestInstance) GetPlanOk() (*ResizeInstanceRequestInstancePlan, bool)`

GetPlanOk returns a tuple with the Plan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlan

`func (o *ResizeInstanceRequestInstance) SetPlan(v ResizeInstanceRequestInstancePlan)`

SetPlan sets Plan field to given value.

### HasPlan

`func (o *ResizeInstanceRequestInstance) HasPlan() bool`

HasPlan returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


