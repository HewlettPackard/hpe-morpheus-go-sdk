# UserCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** | The user&#39;s first name (optional) | [optional] 
**LastName** | Pointer to **string** | The user&#39;s last name (optional) | [optional] 
**Username** | **string** | Username (unique per tenant). | 
**Email** | **string** | Email address | 
**Password** | **string** | Password to apply to the user | 
**Roles** | [**[]GetAlerts200ResponseAllOfChecksInnerAccount**](GetAlerts200ResponseAllOfChecksInnerAccount.md) | Array of objects with id of the role(s) to assign to the user. | 
**ReceiveNotifications** | Pointer to **bool** | Receive Notifications? | [optional] [default to true]
**LinuxUsername** | Pointer to **string** | Linux Username, user settings for provisioning | [optional] 
**LinuxPassword** | Pointer to **string** | Linux Password, user settings for provisioning | [optional] 
**LinuxKeyPairId** | Pointer to **int64** | Linux SSH Key, user settings for provisioning | [optional] 
**WindowsUsername** | Pointer to **string** | Windows Username, user settings for provisioning | [optional] 
**WindowsPassword** | Pointer to **string** | Windows Password, user settings for provisioning | [optional] 

## Methods

### NewUserCreate

`func NewUserCreate(username string, email string, password string, roles []GetAlerts200ResponseAllOfChecksInnerAccount, ) *UserCreate`

NewUserCreate instantiates a new UserCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCreateWithDefaults

`func NewUserCreateWithDefaults() *UserCreate`

NewUserCreateWithDefaults instantiates a new UserCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *UserCreate) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserCreate) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserCreate) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserCreate) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UserCreate) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserCreate) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserCreate) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserCreate) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetUsername

`func (o *UserCreate) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserCreate) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserCreate) SetUsername(v string)`

SetUsername sets Username field to given value.


### GetEmail

`func (o *UserCreate) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserCreate) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserCreate) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetPassword

`func (o *UserCreate) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserCreate) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserCreate) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetRoles

`func (o *UserCreate) GetRoles() []GetAlerts200ResponseAllOfChecksInnerAccount`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserCreate) GetRolesOk() (*[]GetAlerts200ResponseAllOfChecksInnerAccount, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserCreate) SetRoles(v []GetAlerts200ResponseAllOfChecksInnerAccount)`

SetRoles sets Roles field to given value.


### GetReceiveNotifications

`func (o *UserCreate) GetReceiveNotifications() bool`

GetReceiveNotifications returns the ReceiveNotifications field if non-nil, zero value otherwise.

### GetReceiveNotificationsOk

`func (o *UserCreate) GetReceiveNotificationsOk() (*bool, bool)`

GetReceiveNotificationsOk returns a tuple with the ReceiveNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveNotifications

`func (o *UserCreate) SetReceiveNotifications(v bool)`

SetReceiveNotifications sets ReceiveNotifications field to given value.

### HasReceiveNotifications

`func (o *UserCreate) HasReceiveNotifications() bool`

HasReceiveNotifications returns a boolean if a field has been set.

### GetLinuxUsername

`func (o *UserCreate) GetLinuxUsername() string`

GetLinuxUsername returns the LinuxUsername field if non-nil, zero value otherwise.

### GetLinuxUsernameOk

`func (o *UserCreate) GetLinuxUsernameOk() (*string, bool)`

GetLinuxUsernameOk returns a tuple with the LinuxUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxUsername

`func (o *UserCreate) SetLinuxUsername(v string)`

SetLinuxUsername sets LinuxUsername field to given value.

### HasLinuxUsername

`func (o *UserCreate) HasLinuxUsername() bool`

HasLinuxUsername returns a boolean if a field has been set.

### GetLinuxPassword

`func (o *UserCreate) GetLinuxPassword() string`

GetLinuxPassword returns the LinuxPassword field if non-nil, zero value otherwise.

### GetLinuxPasswordOk

`func (o *UserCreate) GetLinuxPasswordOk() (*string, bool)`

GetLinuxPasswordOk returns a tuple with the LinuxPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxPassword

`func (o *UserCreate) SetLinuxPassword(v string)`

SetLinuxPassword sets LinuxPassword field to given value.

### HasLinuxPassword

`func (o *UserCreate) HasLinuxPassword() bool`

HasLinuxPassword returns a boolean if a field has been set.

### GetLinuxKeyPairId

`func (o *UserCreate) GetLinuxKeyPairId() int64`

GetLinuxKeyPairId returns the LinuxKeyPairId field if non-nil, zero value otherwise.

### GetLinuxKeyPairIdOk

`func (o *UserCreate) GetLinuxKeyPairIdOk() (*int64, bool)`

GetLinuxKeyPairIdOk returns a tuple with the LinuxKeyPairId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxKeyPairId

`func (o *UserCreate) SetLinuxKeyPairId(v int64)`

SetLinuxKeyPairId sets LinuxKeyPairId field to given value.

### HasLinuxKeyPairId

`func (o *UserCreate) HasLinuxKeyPairId() bool`

HasLinuxKeyPairId returns a boolean if a field has been set.

### GetWindowsUsername

`func (o *UserCreate) GetWindowsUsername() string`

GetWindowsUsername returns the WindowsUsername field if non-nil, zero value otherwise.

### GetWindowsUsernameOk

`func (o *UserCreate) GetWindowsUsernameOk() (*string, bool)`

GetWindowsUsernameOk returns a tuple with the WindowsUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsUsername

`func (o *UserCreate) SetWindowsUsername(v string)`

SetWindowsUsername sets WindowsUsername field to given value.

### HasWindowsUsername

`func (o *UserCreate) HasWindowsUsername() bool`

HasWindowsUsername returns a boolean if a field has been set.

### GetWindowsPassword

`func (o *UserCreate) GetWindowsPassword() string`

GetWindowsPassword returns the WindowsPassword field if non-nil, zero value otherwise.

### GetWindowsPasswordOk

`func (o *UserCreate) GetWindowsPasswordOk() (*string, bool)`

GetWindowsPasswordOk returns a tuple with the WindowsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsPassword

`func (o *UserCreate) SetWindowsPassword(v string)`

SetWindowsPassword sets WindowsPassword field to given value.

### HasWindowsPassword

`func (o *UserCreate) HasWindowsPassword() bool`

HasWindowsPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


