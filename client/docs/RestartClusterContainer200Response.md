# RestartClusterContainer200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Errors** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewRestartClusterContainer200Response

`func NewRestartClusterContainer200Response() *RestartClusterContainer200Response`

NewRestartClusterContainer200Response instantiates a new RestartClusterContainer200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRestartClusterContainer200ResponseWithDefaults

`func NewRestartClusterContainer200ResponseWithDefaults() *RestartClusterContainer200Response`

NewRestartClusterContainer200ResponseWithDefaults instantiates a new RestartClusterContainer200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *RestartClusterContainer200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *RestartClusterContainer200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *RestartClusterContainer200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *RestartClusterContainer200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetErrors

`func (o *RestartClusterContainer200Response) GetErrors() map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *RestartClusterContainer200Response) GetErrorsOk() (*map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *RestartClusterContainer200Response) SetErrors(v map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *RestartClusterContainer200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


