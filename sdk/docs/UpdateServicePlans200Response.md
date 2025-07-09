# UpdateServicePlans200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServicePlan** | Pointer to [**GetServicePlans200ResponseServicePlan**](GetServicePlans200ResponseServicePlan.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateServicePlans200Response

`func NewUpdateServicePlans200Response() *UpdateServicePlans200Response`

NewUpdateServicePlans200Response instantiates a new UpdateServicePlans200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateServicePlans200ResponseWithDefaults

`func NewUpdateServicePlans200ResponseWithDefaults() *UpdateServicePlans200Response`

NewUpdateServicePlans200ResponseWithDefaults instantiates a new UpdateServicePlans200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServicePlan

`func (o *UpdateServicePlans200Response) GetServicePlan() GetServicePlans200ResponseServicePlan`

GetServicePlan returns the ServicePlan field if non-nil, zero value otherwise.

### GetServicePlanOk

`func (o *UpdateServicePlans200Response) GetServicePlanOk() (*GetServicePlans200ResponseServicePlan, bool)`

GetServicePlanOk returns a tuple with the ServicePlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlan

`func (o *UpdateServicePlans200Response) SetServicePlan(v GetServicePlans200ResponseServicePlan)`

SetServicePlan sets ServicePlan field to given value.

### HasServicePlan

`func (o *UpdateServicePlans200Response) HasServicePlan() bool`

HasServicePlan returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateServicePlans200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateServicePlans200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateServicePlans200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateServicePlans200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


