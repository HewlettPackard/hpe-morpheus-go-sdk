# UpdateUserSettingsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**UpdateUserSettingsRequestUser**](UpdateUserSettingsRequestUser.md) |  | [optional] 

## Methods

### NewUpdateUserSettingsRequest

`func NewUpdateUserSettingsRequest() *UpdateUserSettingsRequest`

NewUpdateUserSettingsRequest instantiates a new UpdateUserSettingsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserSettingsRequestWithDefaults

`func NewUpdateUserSettingsRequestWithDefaults() *UpdateUserSettingsRequest`

NewUpdateUserSettingsRequestWithDefaults instantiates a new UpdateUserSettingsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *UpdateUserSettingsRequest) GetUser() UpdateUserSettingsRequestUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UpdateUserSettingsRequest) GetUserOk() (*UpdateUserSettingsRequestUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UpdateUserSettingsRequest) SetUser(v UpdateUserSettingsRequestUser)`

SetUser sets User field to given value.

### HasUser

`func (o *UpdateUserSettingsRequest) HasUser() bool`

HasUser returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


