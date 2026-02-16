# AddUserRequestUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** | The user&#39;s first name (optional) | [optional] 
**LastName** | Pointer to **string** | The user&#39;s last name (optional) | [optional] 
**Username** | **string** | Username (unique per tenant). | 
**Email** | **string** | Email address | 
**Password** | **string** | Password to apply to the user | 
**Roles** | [**[]AddUserRequestUserRolesInner**](AddUserRequestUserRolesInner.md) | Array of objects with id of the role(s) to assign to the user. | 
**ReceiveNotifications** | Pointer to **bool** | Receive Notifications? | [optional] [default to true]
**LinuxUsername** | Pointer to **string** | Linux Username, user settings for provisioning | [optional] 
**LinuxPassword** | Pointer to **string** | Linux Password, user settings for provisioning | [optional] 
**LinuxKeyPairId** | Pointer to **int64** | Linux SSH Key, user settings for provisioning | [optional] 
**WindowsUsername** | Pointer to **string** | Windows Username, user settings for provisioning | [optional] 
**WindowsPassword** | Pointer to **string** | Windows Password, user settings for provisioning | [optional] 

## Methods

### NewAddUserRequestUser

`func NewAddUserRequestUser(username string, email string, password string, roles []AddUserRequestUserRolesInner, ) *AddUserRequestUser`

NewAddUserRequestUser instantiates a new AddUserRequestUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddUserRequestUserWithDefaults

`func NewAddUserRequestUserWithDefaults() *AddUserRequestUser`

NewAddUserRequestUserWithDefaults instantiates a new AddUserRequestUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *AddUserRequestUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *AddUserRequestUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *AddUserRequestUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *AddUserRequestUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *AddUserRequestUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *AddUserRequestUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *AddUserRequestUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *AddUserRequestUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetUsername

`func (o *AddUserRequestUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *AddUserRequestUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *AddUserRequestUser) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetEmail

`func (o *AddUserRequestUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AddUserRequestUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AddUserRequestUser) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *AddUserRequestUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AddUserRequestUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AddUserRequestUser) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetRoles

`func (o *AddUserRequestUser) GetRoles() []AddUserRequestUserRolesInner`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *AddUserRequestUser) GetRolesOk() (*[]AddUserRequestUserRolesInner, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *AddUserRequestUser) SetRoles(v []AddUserRequestUserRolesInner)`

SetRoles sets Roles field to given value.


### GetReceiveNotifications

`func (o *AddUserRequestUser) GetReceiveNotifications() bool`

GetReceiveNotifications returns the ReceiveNotifications field if non-nil, zero value otherwise.

### GetReceiveNotificationsOk

`func (o *AddUserRequestUser) GetReceiveNotificationsOk() (*bool, bool)`

GetReceiveNotificationsOk returns a tuple with the ReceiveNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveNotifications

`func (o *AddUserRequestUser) SetReceiveNotifications(v bool)`

SetReceiveNotifications sets ReceiveNotifications field to given value.

### HasReceiveNotifications

`func (o *AddUserRequestUser) HasReceiveNotifications() bool`

HasReceiveNotifications returns a boolean if a field has been set.

### GetLinuxUsername

`func (o *AddUserRequestUser) GetLinuxUsername() string`

GetLinuxUsername returns the LinuxUsername field if non-nil, zero value otherwise.

### GetLinuxUsernameOk

`func (o *AddUserRequestUser) GetLinuxUsernameOk() (*string, bool)`

GetLinuxUsernameOk returns a tuple with the LinuxUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxUsername

`func (o *AddUserRequestUser) SetLinuxUsername(v string)`

SetLinuxUsername sets LinuxUsername field to given value.

### HasLinuxUsername

`func (o *AddUserRequestUser) HasLinuxUsername() bool`

HasLinuxUsername returns a boolean if a field has been set.

### GetLinuxPassword

`func (o *AddUserRequestUser) GetLinuxPassword() string`

GetLinuxPassword returns the LinuxPassword field if non-nil, zero value otherwise.

### GetLinuxPasswordOk

`func (o *AddUserRequestUser) GetLinuxPasswordOk() (*string, bool)`

GetLinuxPasswordOk returns a tuple with the LinuxPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxPassword

`func (o *AddUserRequestUser) SetLinuxPassword(v string)`

SetLinuxPassword sets LinuxPassword field to given value.

### HasLinuxPassword

`func (o *AddUserRequestUser) HasLinuxPassword() bool`

HasLinuxPassword returns a boolean if a field has been set.

### GetLinuxKeyPairId

`func (o *AddUserRequestUser) GetLinuxKeyPairId() int64`

GetLinuxKeyPairId returns the LinuxKeyPairId field if non-nil, zero value otherwise.

### GetLinuxKeyPairIdOk

`func (o *AddUserRequestUser) GetLinuxKeyPairIdOk() (*int64, bool)`

GetLinuxKeyPairIdOk returns a tuple with the LinuxKeyPairId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxKeyPairId

`func (o *AddUserRequestUser) SetLinuxKeyPairId(v int64)`

SetLinuxKeyPairId sets LinuxKeyPairId field to given value.

### HasLinuxKeyPairId

`func (o *AddUserRequestUser) HasLinuxKeyPairId() bool`

HasLinuxKeyPairId returns a boolean if a field has been set.

### GetWindowsUsername

`func (o *AddUserRequestUser) GetWindowsUsername() string`

GetWindowsUsername returns the WindowsUsername field if non-nil, zero value otherwise.

### GetWindowsUsernameOk

`func (o *AddUserRequestUser) GetWindowsUsernameOk() (*string, bool)`

GetWindowsUsernameOk returns a tuple with the WindowsUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsUsername

`func (o *AddUserRequestUser) SetWindowsUsername(v string)`

SetWindowsUsername sets WindowsUsername field to given value.

### HasWindowsUsername

`func (o *AddUserRequestUser) HasWindowsUsername() bool`

HasWindowsUsername returns a boolean if a field has been set.

### GetWindowsPassword

`func (o *AddUserRequestUser) GetWindowsPassword() string`

GetWindowsPassword returns the WindowsPassword field if non-nil, zero value otherwise.

### GetWindowsPasswordOk

`func (o *AddUserRequestUser) GetWindowsPasswordOk() (*string, bool)`

GetWindowsPasswordOk returns a tuple with the WindowsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsPassword

`func (o *AddUserRequestUser) SetWindowsPassword(v string)`

SetWindowsPassword sets WindowsPassword field to given value.

### HasWindowsPassword

`func (o *AddUserRequestUser) HasWindowsPassword() bool`

HasWindowsPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


