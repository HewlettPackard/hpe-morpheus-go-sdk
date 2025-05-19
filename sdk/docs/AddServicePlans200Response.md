# AddServicePlans200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ServicePlan** | Pointer to [**ListServicePlans200ResponseAllOfServicePlansInner**](ListServicePlans200ResponseAllOfServicePlansInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddServicePlans200Response

`func NewAddServicePlans200Response() *AddServicePlans200Response`

NewAddServicePlans200Response instantiates a new AddServicePlans200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddServicePlans200ResponseWithDefaults

`func NewAddServicePlans200ResponseWithDefaults() *AddServicePlans200Response`

NewAddServicePlans200ResponseWithDefaults instantiates a new AddServicePlans200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServicePlan

`func (o *AddServicePlans200Response) GetServicePlan() ListServicePlans200ResponseAllOfServicePlansInner`

GetServicePlan returns the ServicePlan field if non-nil, zero value otherwise.

### GetServicePlanOk

`func (o *AddServicePlans200Response) GetServicePlanOk() (*ListServicePlans200ResponseAllOfServicePlansInner, bool)`

GetServicePlanOk returns a tuple with the ServicePlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServicePlan

`func (o *AddServicePlans200Response) SetServicePlan(v ListServicePlans200ResponseAllOfServicePlansInner)`

SetServicePlan sets ServicePlan field to given value.

### HasServicePlan

`func (o *AddServicePlans200Response) HasServicePlan() bool`

HasServicePlan returns a boolean if a field has been set.

### GetSuccess

`func (o *AddServicePlans200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddServicePlans200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddServicePlans200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddServicePlans200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


