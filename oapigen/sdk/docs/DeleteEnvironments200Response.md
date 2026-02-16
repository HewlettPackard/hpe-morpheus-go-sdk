# DeleteEnvironments200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Errors** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewDeleteEnvironments200Response

`func NewDeleteEnvironments200Response() *DeleteEnvironments200Response`

NewDeleteEnvironments200Response instantiates a new DeleteEnvironments200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeleteEnvironments200ResponseWithDefaults

`func NewDeleteEnvironments200ResponseWithDefaults() *DeleteEnvironments200Response`

NewDeleteEnvironments200ResponseWithDefaults instantiates a new DeleteEnvironments200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *DeleteEnvironments200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *DeleteEnvironments200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *DeleteEnvironments200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *DeleteEnvironments200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetErrors

`func (o *DeleteEnvironments200Response) GetErrors() map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *DeleteEnvironments200Response) GetErrorsOk() (*map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *DeleteEnvironments200Response) SetErrors(v map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *DeleteEnvironments200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *DeleteEnvironments200Response) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *DeleteEnvironments200Response) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


