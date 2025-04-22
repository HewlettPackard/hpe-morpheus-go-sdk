# AddImageBuild200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageBuild** | Pointer to [**AddImageBuild200ResponseAllOfImageBuild**](AddImageBuild200ResponseAllOfImageBuild.md) |  | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 
**Success** | Pointer to **bool** | Success indicator, true when the request succeeded and false when an error occurred | [optional] [default to true]
**Msg** | Pointer to **string** | Message containing a description of the result, usually a message about the error that occurred | [optional] 
**Errors** | Pointer to **map[string]interface{}** | Validation errors, with a key for Object containing error messages for each invalid parameter (key) | [optional] 

## Methods

### NewAddImageBuild200Response

`func NewAddImageBuild200Response() *AddImageBuild200Response`

NewAddImageBuild200Response instantiates a new AddImageBuild200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddImageBuild200ResponseWithDefaults

`func NewAddImageBuild200ResponseWithDefaults() *AddImageBuild200Response`

NewAddImageBuild200ResponseWithDefaults instantiates a new AddImageBuild200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageBuild

`func (o *AddImageBuild200Response) GetImageBuild() AddImageBuild200ResponseAllOfImageBuild`

GetImageBuild returns the ImageBuild field if non-nil, zero value otherwise.

### GetImageBuildOk

`func (o *AddImageBuild200Response) GetImageBuildOk() (*AddImageBuild200ResponseAllOfImageBuild, bool)`

GetImageBuildOk returns a tuple with the ImageBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBuild

`func (o *AddImageBuild200Response) SetImageBuild(v AddImageBuild200ResponseAllOfImageBuild)`

SetImageBuild sets ImageBuild field to given value.

### HasImageBuild

`func (o *AddImageBuild200Response) HasImageBuild() bool`

HasImageBuild returns a boolean if a field has been set.

### GetErrorCode

`func (o *AddImageBuild200Response) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *AddImageBuild200Response) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *AddImageBuild200Response) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *AddImageBuild200Response) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetSuccess

`func (o *AddImageBuild200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddImageBuild200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddImageBuild200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddImageBuild200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetMsg

`func (o *AddImageBuild200Response) GetMsg() string`

GetMsg returns the Msg field if non-nil, zero value otherwise.

### GetMsgOk

`func (o *AddImageBuild200Response) GetMsgOk() (*string, bool)`

GetMsgOk returns a tuple with the Msg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsg

`func (o *AddImageBuild200Response) SetMsg(v string)`

SetMsg sets Msg field to given value.

### HasMsg

`func (o *AddImageBuild200Response) HasMsg() bool`

HasMsg returns a boolean if a field has been set.

### GetErrors

`func (o *AddImageBuild200Response) GetErrors() map[string]interface{}`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *AddImageBuild200Response) GetErrorsOk() (*map[string]interface{}, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *AddImageBuild200Response) SetErrors(v map[string]interface{})`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *AddImageBuild200Response) HasErrors() bool`

HasErrors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


