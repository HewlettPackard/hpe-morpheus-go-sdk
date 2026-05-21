# ListApplianceSettings4XXResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** | Success indicator, true when the request succeeded and false when an error occurred | [optional] [default to true]
**Msg** | Pointer to **NullableString** | Message containing a description of the result, usually a message about the error that occurred | [optional] 
**Errors** | Pointer to **map[string]interface{}** | Validation errors, with a key for Object containing error messages for each invalid parameter (key) | [optional] 

## Methods

### NewListApplianceSettings4XXResponse

`func NewListApplianceSettings4XXResponse() *ListApplianceSettings4XXResponse`

NewListApplianceSettings4XXResponse instantiates a new ListApplianceSettings4XXResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListApplianceSettings4XXResponseWithDefaults

`func NewListApplianceSettings4XXResponseWithDefaults() *ListApplianceSettings4XXResponse`

NewListApplianceSettings4XXResponseWithDefaults instantiates a new ListApplianceSettings4XXResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ListApplianceSettings4XXResponse) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ListApplianceSettings4XXResponse) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ListApplianceSettings4XXResponse) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ListApplianceSettings4XXResponse) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMsg

`func (o *ListApplianceSettings4XXResponse) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *ListApplianceSettings4XXResponse) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *ListApplianceSettings4XXResponse) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *ListApplianceSettings4XXResponse) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### SetMsgNil

`func (o *ListApplianceSettings4XXResponse) SetMsgNil(b bool)`

 SetMsgNil sets the value for Msg to be an explicit nil

### UnsetMsg
`func (o *ListApplianceSettings4XXResponse) UnsetMsg()`

UnsetMsg ensures that no value is present for Msg, not even an explicit nil
### GetErrors

`func (o *ListApplianceSettings4XXResponse) GetErrors() map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *ListApplianceSettings4XXResponse) GetErrorsOk() (*map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *ListApplianceSettings4XXResponse) SetErrors(v map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *ListApplianceSettings4XXResponse) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

### SetErrorsNil

`func (o *ListApplianceSettings4XXResponse) SetErrorsNil(b bool)`

 SetErrorsNil sets the value for Errors to be an explicit nil

### UnsetErrors
`func (o *ListApplianceSettings4XXResponse) UnsetErrors()`

UnsetErrors ensures that no value is present for Errors, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


