# ApplyAppStateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateParameter** | Pointer to **map[string]interface{}** | Template Parameter object. A map of key-value pairs that correspond to the template variables i.e. tfvars | [optional] 

## Methods

### NewApplyAppStateRequest

`func NewApplyAppStateRequest() *ApplyAppStateRequest`

NewApplyAppStateRequest instantiates a new ApplyAppStateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplyAppStateRequestWithDefaults

`func NewApplyAppStateRequestWithDefaults() *ApplyAppStateRequest`

NewApplyAppStateRequestWithDefaults instantiates a new ApplyAppStateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTemplateParameter

`func (o *ApplyAppStateRequest) GetTemplateParameter() map[string]interface{}`

GetTemplateParameter returns the TemplateParameter field if non-nil, zero value otherwise.

### GetTemplateParameterOk

`func (o *ApplyAppStateRequest) GetTemplateParameterOk() (*map[string]interface{}, bool)`

GetTemplateParameterOk returns a tuple with the TemplateParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateParameter

`func (o *ApplyAppStateRequest) SetTemplateParameter(v map[string]interface{})`

SetTemplateParameter sets TemplateParameter field to given value.

### HasTemplateParameter

`func (o *ApplyAppStateRequest) HasTemplateParameter() bool`

HasTemplateParameter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


