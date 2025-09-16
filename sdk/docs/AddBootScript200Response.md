# AddBootScript200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BootScript** | Pointer to [**AddBootScript200ResponseAllOfBootScript**](AddBootScript200ResponseAllOfBootScript.md) |  | [optional] 
**ErrorCode** | Pointer to **NullableString** |  | [optional] 
**InProgress** | Pointer to **bool** |  | [optional] 
**Success** | Pointer to **bool** | Success indicator, true when the request succeeded and false when an error occurred | [optional] [default to true]
**Msg** | Pointer to **NullableString** | Message containing a description of the result, usually a message about the error that occurred | [optional] 
**Errors** | Pointer to **map[string]interface{}** | Validation errors, with a key for Object containing error messages for each invalid parameter (key) | [optional] 

## Methods

### NewAddBootScript200Response

`func NewAddBootScript200Response() *AddBootScript200Response`

NewAddBootScript200Response instantiates a new AddBootScript200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddBootScript200ResponseWithDefaults

`func NewAddBootScript200ResponseWithDefaults() *AddBootScript200Response`

NewAddBootScript200ResponseWithDefaults instantiates a new AddBootScript200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBootScript

`func (o *AddBootScript200Response) GetBootScript() AddBootScript200ResponseAllOfBootScript`

GetBootScript returns the BootScript field if non-nil, zero value otherwise.

### GetBootScriptOk

`func (o *AddBootScript200Response) GetBootScriptOk() (*AddBootScript200ResponseAllOfBootScript, bool)`

GetBootScriptOk returns a tuple with the BootScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBootScript

`func (o *AddBootScript200Response) SetBootScript(v AddBootScript200ResponseAllOfBootScript)`

SetBootScript sets BootScript field to given value.

### HasBootScript

`func (o *AddBootScript200Response) HasBootScript() bool`

HasBootScript returns a boolean if a field has been set.

### GetErrorCode

`func (o *AddBootScript200Response) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *AddBootScript200Response) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *AddBootScript200Response) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *AddBootScript200Response) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### SetErrorCodeNil

`func (o *AddBootScript200Response) SetErrorCodeNil(b bool)`

 SetErrorCodeNil sets the value for ErrorCode to be an explicit nil

### UnsetErrorCode
`func (o *AddBootScript200Response) UnsetErrorCode()`

UnsetErrorCode ensures that no value is present for ErrorCode, not even an explicit nil
### GetInProgress

`func (o *AddBootScript200Response) GetInProgress() bool`

GetInProgress returns the InProgress field if non-nil, zero value otherwise.

### GetInProgressOk

`func (o *AddBootScript200Response) GetInProgressOk() (*bool, bool)`

GetInProgressOk returns a tuple with the InProgress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInProgress

`func (o *AddBootScript200Response) SetInProgress(v bool)`

SetInProgress sets InProgress field to given value.

### HasInProgress

`func (o *AddBootScript200Response) HasInProgress() bool`

HasInProgress returns a boolean if a field has been set.

### GetSuccess

`func (o *AddBootScript200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddBootScript200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddBootScript200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddBootScript200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMsg

`func (o *AddBootScript200Response) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *AddBootScript200Response) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *AddBootScript200Response) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *AddBootScript200Response) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### SetMsgNil

`func (o *AddBootScript200Response) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *AddBootScript200Response) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetErrors

`func (o *AddBootScript200Response) GetErrors() map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *AddBootScript200Response) GetErrorsOk() (*map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *AddBootScript200Response) SetErrors(v map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *AddBootScript200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *AddBootScript200Response) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *AddBootScript200Response) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


