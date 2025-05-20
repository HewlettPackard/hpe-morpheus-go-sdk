# AddPreseedScript200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PreseedScript** | Pointer to [**ListPreseedScripts200ResponseAllOfPreseedScriptsInner**](ListPreseedScripts200ResponseAllOfPreseedScriptsInner.md) |  | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 
**Success** | Pointer to **bool** | Success indicator, true when the request succeeded and false when an error occurred | [optional] [default to true]
**Msg** | Pointer to **string** | Message containing a description of the result, usually a message about the error that occurred | [optional] 
**Errors** | Pointer to **map[string]interface{}** | Validation errors, with a key for Object containing error messages for each invalid parameter (key) | [optional] 

## Methods

### NewAddPreseedScript200Response

`func NewAddPreseedScript200Response() *AddPreseedScript200Response`

NewAddPreseedScript200Response instantiates a new AddPreseedScript200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPreseedScript200ResponseWithDefaults

`func NewAddPreseedScript200ResponseWithDefaults() *AddPreseedScript200Response`

NewAddPreseedScript200ResponseWithDefaults instantiates a new AddPreseedScript200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPreseedScript

`func (o *AddPreseedScript200Response) GetPreseedScript() ListPreseedScripts200ResponseAllOfPreseedScriptsInner`

GetPreseedScript returns the PreseedScript field if non-nil, zero value otherwise.

### GetPreseedScriptOk

`func (o *AddPreseedScript200Response) GetPreseedScriptOk() (*ListPreseedScripts200ResponseAllOfPreseedScriptsInner, bool)`

GetPreseedScriptOk returns a tuple with the PreseedScript field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreseedScript

`func (o *AddPreseedScript200Response) SetPreseedScript(v ListPreseedScripts200ResponseAllOfPreseedScriptsInner)`

SetPreseedScript sets PreseedScript field to given value.

### HasPreseedScript

`func (o *AddPreseedScript200Response) HasPreseedScript() bool`

HasPreseedScript returns a boolean if a field has been set.

### GetErrorCode

`func (o *AddPreseedScript200Response) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *AddPreseedScript200Response) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *AddPreseedScript200Response) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *AddPreseedScript200Response) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetSuccess

`func (o *AddPreseedScript200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddPreseedScript200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddPreseedScript200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddPreseedScript200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMsg

`func (o *AddPreseedScript200Response) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *AddPreseedScript200Response) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *AddPreseedScript200Response) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *AddPreseedScript200Response) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetErrors

`func (o *AddPreseedScript200Response) GetErrors() map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *AddPreseedScript200Response) GetErrorsOk() (*map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *AddPreseedScript200Response) SetErrors(v map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *AddPreseedScript200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


