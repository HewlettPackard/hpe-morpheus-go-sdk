# ExecuteExecutionRequestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Script** | **string** | A script or command to be executed | 
**SendKeys** | Pointer to **bool** | Pass true to send key mappings to the hypervisor console so commands such as &lt;LEFT&gt;, &lt;RIGHT&gt; and &lt;WAIT&gt; can be used. | [optional] [default to false]

## Methods

### NewExecuteExecutionRequestRequest

`func NewExecuteExecutionRequestRequest(script string, ) *ExecuteExecutionRequestRequest`

NewExecuteExecutionRequestRequest instantiates a new ExecuteExecutionRequestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExecuteExecutionRequestRequestWithDefaults

`func NewExecuteExecutionRequestRequestWithDefaults() *ExecuteExecutionRequestRequest`

NewExecuteExecutionRequestRequestWithDefaults instantiates a new ExecuteExecutionRequestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetScript

`func (o *ExecuteExecutionRequestRequest) GetScript() string`

GetScript returns the Script field if non-nil, zero value otherwise.

### GetScriptOk

`func (o *ExecuteExecutionRequestRequest) GetScriptOk() (*string, bool)`

GetScriptOk returns a tuple with the Script field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScript

`func (o *ExecuteExecutionRequestRequest) SetScript(v string)`

SetScript sets Script field to given value.


### GetSendKeys

`func (o *ExecuteExecutionRequestRequest) GetSendKeys() bool`

GetSendKeys returns the SendKeys field if non-nil, zero value otherwise.

### GetSendKeysOk

`func (o *ExecuteExecutionRequestRequest) GetSendKeysOk() (*bool, bool)`

GetSendKeysOk returns a tuple with the SendKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSendKeys

`func (o *ExecuteExecutionRequestRequest) SetSendKeys(v bool)`

SetSendKeys sets SendKeys field to given value.

### HasSendKeys

`func (o *ExecuteExecutionRequestRequest) HasSendKeys() bool`

HasSendKeys returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


