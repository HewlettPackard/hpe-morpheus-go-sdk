# Model200SuccessExpanded

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** | Success indicator, true when the request succeeded and false when an error occurred | [optional] [default to true]
**Msg** | Pointer to **string** | Message containing a description of the result, usually a message about the error that occurred | [optional] 
**Errors** | Pointer to **map[string]interface{}** | Validation errors, with a key for Object containing error messages for each invalid parameter (key) | [optional] 

## Methods

### NewModel200SuccessExpanded

`func NewModel200SuccessExpanded() *Model200SuccessExpanded`

NewModel200SuccessExpanded instantiates a new Model200SuccessExpanded object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewModel200SuccessExpandedWithDefaults

`func NewModel200SuccessExpandedWithDefaults() *Model200SuccessExpanded`

NewModel200SuccessExpandedWithDefaults instantiates a new Model200SuccessExpanded object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *Model200SuccessExpanded) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *Model200SuccessExpanded) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *Model200SuccessExpanded) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *Model200SuccessExpanded) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMsg

`func (o *Model200SuccessExpanded) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *Model200SuccessExpanded) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *Model200SuccessExpanded) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *Model200SuccessExpanded) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetErrors

`func (o *Model200SuccessExpanded) GetErrors() map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *Model200SuccessExpanded) GetErrorsOk() (*map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *Model200SuccessExpanded) SetErrors(v map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *Model200SuccessExpanded) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


