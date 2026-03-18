# RunWorkflowInstance200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**ProcessId** | Pointer to **int64** | Process ID to track execution results with, use &#x60;/api/processes/$processId&#x60; | [optional] 

## Methods

### NewRunWorkflowInstance200Response

`func NewRunWorkflowInstance200Response() *RunWorkflowInstance200Response`

NewRunWorkflowInstance200Response instantiates a new RunWorkflowInstance200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRunWorkflowInstance200ResponseWithDefaults

`func NewRunWorkflowInstance200ResponseWithDefaults() *RunWorkflowInstance200Response`

NewRunWorkflowInstance200ResponseWithDefaults instantiates a new RunWorkflowInstance200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *RunWorkflowInstance200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *RunWorkflowInstance200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *RunWorkflowInstance200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *RunWorkflowInstance200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetProcessId

`func (o *RunWorkflowInstance200Response) GetProcessId() int64`

GetProcessId returns the ProcessId field if non-nil, zero value otherwise.

### GetProcessIdOk

`func (o *RunWorkflowInstance200Response) GetProcessIdOk() (*int64, bool)`

GetProcessIdOk returns a tuple with the ProcessId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessId

`func (o *RunWorkflowInstance200Response) SetProcessId(v int64)`

SetProcessId sets ProcessId field to given value.

### HasProcessId

`func (o *RunWorkflowInstance200Response) HasProcessId() bool`

HasProcessId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


