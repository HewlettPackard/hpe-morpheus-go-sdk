# AddUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | [**AddUserRequestUser**](AddUserRequestUser.md) |  | 

## Methods

### NewAddUserRequest

`func NewAddUserRequest(user AddUserRequestUser, ) *AddUserRequest`

NewAddUserRequest instantiates a new AddUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddUserRequestWithDefaults

`func NewAddUserRequestWithDefaults() *AddUserRequest`

NewAddUserRequestWithDefaults instantiates a new AddUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *AddUserRequest) GetUser() AddUserRequestUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AddUserRequest) GetUserOk() (*AddUserRequestUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AddUserRequest) SetUser(v AddUserRequestUser)`

SetUser sets User field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


