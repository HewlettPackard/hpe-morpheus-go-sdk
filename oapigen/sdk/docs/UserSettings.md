# UserSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**UserSettingsUser**](UserSettingsUser.md) |  | [optional] 
**AccessTokens** | Pointer to [**[]UserSettingsAccessTokensInner**](UserSettingsAccessTokensInner.md) |  | [optional] 

## Methods

### NewUserSettings

`func NewUserSettings() *UserSettings`

NewUserSettings instantiates a new UserSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetUser

`func (o *UserSettings) GetUser() UserSettingsUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UserSettings) GetUserOk() (*UserSettingsUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UserSettings) SetUser(v UserSettingsUser)`

SetUser sets User field to given value.

### HasUser

`func (o *UserSettings) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetAccessTokens

`func (o *UserSettings) GetAccessTokens() []UserSettingsAccessTokensInner`

GetAccessTokens returns the AccessTokens field if non-nil, zero value otherwise.

### GetAccessTokensOk

`func (o *UserSettings) GetAccessTokensOk() (*[]UserSettingsAccessTokensInner, bool)`

GetAccessTokensOk returns a tuple with the AccessTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokens

`func (o *UserSettings) SetAccessTokens(v []UserSettingsAccessTokensInner)`

SetAccessTokens sets AccessTokens field to given value.

### HasAccessTokens

`func (o *UserSettings) HasAccessTokens() bool`

HasAccessTokens returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


