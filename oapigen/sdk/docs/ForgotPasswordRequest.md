# ForgotPasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | **string** | Specified as \&quot;username\&quot; or \&quot;tenantId\\username\&quot; with proper HTML encoding and used in conjunction with &#x60;password&#x60;.  | 

## Methods

### NewForgotPasswordRequest

`func NewForgotPasswordRequest(username string, ) *ForgotPasswordRequest`

NewForgotPasswordRequest instantiates a new ForgotPasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewForgotPasswordRequestWithDefaults

`func NewForgotPasswordRequestWithDefaults() *ForgotPasswordRequest`

NewForgotPasswordRequestWithDefaults instantiates a new ForgotPasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *ForgotPasswordRequest) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *ForgotPasswordRequest) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *ForgotPasswordRequest) SetUsername(v string)`

SetUsername sets Username field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


