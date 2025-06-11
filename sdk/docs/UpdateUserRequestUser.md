# UpdateUserRequestUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **NullableString** | First Name | [optional] 
**LastName** | Pointer to **NullableString** | Last Name | [optional] 
**Username** | Pointer to **string** | Username (unique per tenant). | [optional] 
**LinuxUsername** | Pointer to **NullableString** |  | [optional] 
**LinuxPassword** | Pointer to **NullableString** |  | [optional] 
**LinuxKeyPairId** | Pointer to **NullableInt64** |  | [optional] 
**WindowsUsername** | Pointer to **NullableString** |  | [optional] 
**WindowsPassword** | Pointer to **NullableString** |  | [optional] 
**Email** | Pointer to **string** | Email address | [optional] 
**Password** | Pointer to **string** | Password | [optional] 
**Roles** | Pointer to [**[]UpdateUserRequestUserRolesInner**](UpdateUserRequestUserRolesInner.md) | List of Roles | [optional] 

## Methods

### NewUpdateUserRequestUser

`func NewUpdateUserRequestUser() *UpdateUserRequestUser`

NewUpdateUserRequestUser instantiates a new UpdateUserRequestUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUserRequestUserWithDefaults

`func NewUpdateUserRequestUserWithDefaults() *UpdateUserRequestUser`

NewUpdateUserRequestUserWithDefaults instantiates a new UpdateUserRequestUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *UpdateUserRequestUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UpdateUserRequestUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UpdateUserRequestUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UpdateUserRequestUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### SetFirstNameNil

`func (o *UpdateUserRequestUser) SetFirstNameNil(b bool)`

 SetFirstNameNil sets the value for FirstName to be an explicit nil

### UnsetFirstName
`func (o *UpdateUserRequestUser) UnsetFirstName()`

UnsetFirstName ensures that no value is present for FirstName, not even an explicit nil
### GetLastName

`func (o *UpdateUserRequestUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UpdateUserRequestUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UpdateUserRequestUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UpdateUserRequestUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### SetLastNameNil

`func (o *UpdateUserRequestUser) SetLastNameNil(b bool)`

 SetLastNameNil sets the value for LastName to be an explicit nil

### UnsetLastName
`func (o *UpdateUserRequestUser) UnsetLastName()`

UnsetLastName ensures that no value is present for LastName, not even an explicit nil
### GetUsername

`func (o *UpdateUserRequestUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateUserRequestUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateUserRequestUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateUserRequestUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetLinuxUsername

`func (o *UpdateUserRequestUser) GetLinuxUsername() string`

GetLinuxUsername returns the LinuxUsername field if non-nil, zero value otherwise.

### GetLinuxUsernameOk

`func (o *UpdateUserRequestUser) GetLinuxUsernameOk() (*string, bool)`

GetLinuxUsernameOk returns a tuple with the LinuxUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxUsername

`func (o *UpdateUserRequestUser) SetLinuxUsername(v string)`

SetLinuxUsername sets LinuxUsername field to given value.

### HasLinuxUsername

`func (o *UpdateUserRequestUser) HasLinuxUsername() bool`

HasLinuxUsername returns a boolean if a field has been set.

### SetLinuxUsernameNil

`func (o *UpdateUserRequestUser) SetLinuxUsernameNil(b bool)`

 SetLinuxUsernameNil sets the value for LinuxUsername to be an explicit nil

### UnsetLinuxUsername
`func (o *UpdateUserRequestUser) UnsetLinuxUsername()`

UnsetLinuxUsername ensures that no value is present for LinuxUsername, not even an explicit nil
### GetLinuxPassword

`func (o *UpdateUserRequestUser) GetLinuxPassword() string`

GetLinuxPassword returns the LinuxPassword field if non-nil, zero value otherwise.

### GetLinuxPasswordOk

`func (o *UpdateUserRequestUser) GetLinuxPasswordOk() (*string, bool)`

GetLinuxPasswordOk returns a tuple with the LinuxPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxPassword

`func (o *UpdateUserRequestUser) SetLinuxPassword(v string)`

SetLinuxPassword sets LinuxPassword field to given value.

### HasLinuxPassword

`func (o *UpdateUserRequestUser) HasLinuxPassword() bool`

HasLinuxPassword returns a boolean if a field has been set.

### SetLinuxPasswordNil

`func (o *UpdateUserRequestUser) SetLinuxPasswordNil(b bool)`

 SetLinuxPasswordNil sets the value for LinuxPassword to be an explicit nil

### UnsetLinuxPassword
`func (o *UpdateUserRequestUser) UnsetLinuxPassword()`

UnsetLinuxPassword ensures that no value is present for LinuxPassword, not even an explicit nil
### GetLinuxKeyPairId

`func (o *UpdateUserRequestUser) GetLinuxKeyPairId() int64`

GetLinuxKeyPairId returns the LinuxKeyPairId field if non-nil, zero value otherwise.

### GetLinuxKeyPairIdOk

`func (o *UpdateUserRequestUser) GetLinuxKeyPairIdOk() (*int64, bool)`

GetLinuxKeyPairIdOk returns a tuple with the LinuxKeyPairId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxKeyPairId

`func (o *UpdateUserRequestUser) SetLinuxKeyPairId(v int64)`

SetLinuxKeyPairId sets LinuxKeyPairId field to given value.

### HasLinuxKeyPairId

`func (o *UpdateUserRequestUser) HasLinuxKeyPairId() bool`

HasLinuxKeyPairId returns a boolean if a field has been set.

### SetLinuxKeyPairIdNil

`func (o *UpdateUserRequestUser) SetLinuxKeyPairIdNil(b bool)`

 SetLinuxKeyPairIdNil sets the value for LinuxKeyPairId to be an explicit nil

### UnsetLinuxKeyPairId
`func (o *UpdateUserRequestUser) UnsetLinuxKeyPairId()`

UnsetLinuxKeyPairId ensures that no value is present for LinuxKeyPairId, not even an explicit nil
### GetWindowsUsername

`func (o *UpdateUserRequestUser) GetWindowsUsername() string`

GetWindowsUsername returns the WindowsUsername field if non-nil, zero value otherwise.

### GetWindowsUsernameOk

`func (o *UpdateUserRequestUser) GetWindowsUsernameOk() (*string, bool)`

GetWindowsUsernameOk returns a tuple with the WindowsUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsUsername

`func (o *UpdateUserRequestUser) SetWindowsUsername(v string)`

SetWindowsUsername sets WindowsUsername field to given value.

### HasWindowsUsername

`func (o *UpdateUserRequestUser) HasWindowsUsername() bool`

HasWindowsUsername returns a boolean if a field has been set.

### SetWindowsUsernameNil

`func (o *UpdateUserRequestUser) SetWindowsUsernameNil(b bool)`

 SetWindowsUsernameNil sets the value for WindowsUsername to be an explicit nil

### UnsetWindowsUsername
`func (o *UpdateUserRequestUser) UnsetWindowsUsername()`

UnsetWindowsUsername ensures that no value is present for WindowsUsername, not even an explicit nil
### GetWindowsPassword

`func (o *UpdateUserRequestUser) GetWindowsPassword() string`

GetWindowsPassword returns the WindowsPassword field if non-nil, zero value otherwise.

### GetWindowsPasswordOk

`func (o *UpdateUserRequestUser) GetWindowsPasswordOk() (*string, bool)`

GetWindowsPasswordOk returns a tuple with the WindowsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsPassword

`func (o *UpdateUserRequestUser) SetWindowsPassword(v string)`

SetWindowsPassword sets WindowsPassword field to given value.

### HasWindowsPassword

`func (o *UpdateUserRequestUser) HasWindowsPassword() bool`

HasWindowsPassword returns a boolean if a field has been set.

### SetWindowsPasswordNil

`func (o *UpdateUserRequestUser) SetWindowsPasswordNil(b bool)`

 SetWindowsPasswordNil sets the value for WindowsPassword to be an explicit nil

### UnsetWindowsPassword
`func (o *UpdateUserRequestUser) UnsetWindowsPassword()`

UnsetWindowsPassword ensures that no value is present for WindowsPassword, not even an explicit nil
### GetEmail

`func (o *UpdateUserRequestUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateUserRequestUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateUserRequestUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateUserRequestUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPassword

`func (o *UpdateUserRequestUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateUserRequestUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateUserRequestUser) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateUserRequestUser) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRoles

`func (o *UpdateUserRequestUser) GetRoles() []UpdateUserRequestUserRolesInner`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UpdateUserRequestUser) GetRolesOk() (*[]UpdateUserRequestUserRolesInner, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UpdateUserRequestUser) SetRoles(v []UpdateUserRequestUserRolesInner)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UpdateUserRequestUser) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


