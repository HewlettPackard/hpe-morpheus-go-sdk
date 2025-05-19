# ResetPasswordRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Token** | **string** | The secret Reset Password token that was included in the **Forgot Password Email**. | 
**Password** | **string** | User new password. This is the new password for your user. | 

## Methods

### NewResetPasswordRequest

`func NewResetPasswordRequest(token string, password string, ) *ResetPasswordRequest`

NewResetPasswordRequest instantiates a new ResetPasswordRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResetPasswordRequestWithDefaults

`func NewResetPasswordRequestWithDefaults() *ResetPasswordRequest`

NewResetPasswordRequestWithDefaults instantiates a new ResetPasswordRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetToken

`func (o *ResetPasswordRequest) GetToken() string`

GetToken returns the Token field if non-nil, zero value otherwise.

### GetTokenOk

`func (o *ResetPasswordRequest) GetTokenOk() (*string, bool)`

GetTokenOk returns a tuple with the Token field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToken

`func (o *ResetPasswordRequest) SetToken(v string)`

SetToken sets Token field to given value.


### GetPassword

`func (o *ResetPasswordRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *ResetPasswordRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *ResetPasswordRequest) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


