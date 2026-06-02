# ValidateAppStateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TemplateParameter** | Pointer to **map[string]interface{}** | Template Parameter object. A map of key-value pairs that correspond to the template variables i.e. tfvars | [optional] 

## Methods

### NewValidateAppStateRequest

`func NewValidateAppStateRequest() *ValidateAppStateRequest`

NewValidateAppStateRequest instantiates a new ValidateAppStateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetTemplateParameter

`func (o *ValidateAppStateRequest) GetTemplateParameter() map[string]interface{}`

GetTemplateParameter returns the TemplateParameter field if non-nil, zero value otherwise.

### GetTemplateParameterOk

`func (o *ValidateAppStateRequest) GetTemplateParameterOk() (*map[string]interface{}, bool)`

GetTemplateParameterOk returns a tuple with the TemplateParameter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplateParameter

`func (o *ValidateAppStateRequest) SetTemplateParameter(v map[string]interface{})`

SetTemplateParameter sets TemplateParameter field to given value.

### HasTemplateParameter

`func (o *ValidateAppStateRequest) HasTemplateParameter() bool`

HasTemplateParameter returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


