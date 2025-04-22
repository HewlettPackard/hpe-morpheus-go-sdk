# UserSettingsUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Username** | Pointer to **string** | Username | [optional] 
**Email** | Pointer to **string** | Email | [optional] 
**FirstName** | Pointer to **string** | First Name | [optional] 
**LastName** | Pointer to **string** | Last Name | [optional] 
**Password** | Pointer to **string** | Change your password | [optional] 
**LinuxUsername** | Pointer to **string** | Linux Username | [optional] 
**LinuxPassword** | Pointer to **string** | Linux Password | [optional] 
**LinuxKeyPairId** | Pointer to **int64** | Linux Key Pair ID | [optional] 
**WindowsUsername** | Pointer to **string** | Windows Username | [optional] 
**WindowsPassword** | Pointer to **string** | Windows Password | [optional] 
**ReceiveNotifications** | Pointer to **bool** | Receive Notifications (true or false) | [optional] 
**DefaultGroup** | Pointer to [**UpdateUserSettingsRequestUserDefaultGroup**](UpdateUserSettingsRequestUserDefaultGroup.md) |  | [optional] 
**DefaultCloud** | Pointer to [**UpdateUserSettingsRequestUserDefaultCloud**](UpdateUserSettingsRequestUserDefaultCloud.md) |  | [optional] 
**DefaultPersona** | Pointer to [**UpdateUserSettingsRequestUserDefaultPersona**](UpdateUserSettingsRequestUserDefaultPersona.md) |  | [optional] 

## Methods

### NewUserSettingsUpdate

`func NewUserSettingsUpdate() *UserSettingsUpdate`

NewUserSettingsUpdate instantiates a new UserSettingsUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSettingsUpdateWithDefaults

`func NewUserSettingsUpdateWithDefaults() *UserSettingsUpdate`

NewUserSettingsUpdateWithDefaults instantiates a new UserSettingsUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUsername

`func (o *UserSettingsUpdate) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserSettingsUpdate) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserSettingsUpdate) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserSettingsUpdate) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetEmail

`func (o *UserSettingsUpdate) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserSettingsUpdate) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserSettingsUpdate) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserSettingsUpdate) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetFirstName

`func (o *UserSettingsUpdate) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UserSettingsUpdate) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UserSettingsUpdate) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UserSettingsUpdate) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UserSettingsUpdate) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UserSettingsUpdate) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UserSettingsUpdate) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UserSettingsUpdate) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetPassword

`func (o *UserSettingsUpdate) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserSettingsUpdate) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserSettingsUpdate) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserSettingsUpdate) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetLinuxUsername

`func (o *UserSettingsUpdate) GetLinuxUsername() string`

GetLinuxUsername returns the LinuxUsername field if non-nil, zero value otherwise.

### GetLinuxUsernameOk

`func (o *UserSettingsUpdate) GetLinuxUsernameOk() (*string, bool)`

GetLinuxUsernameOk returns a tuple with the LinuxUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxUsername

`func (o *UserSettingsUpdate) SetLinuxUsername(v string)`

SetLinuxUsername sets LinuxUsername field to given value.

### HasLinuxUsername

`func (o *UserSettingsUpdate) HasLinuxUsername() bool`

HasLinuxUsername returns a boolean if a field has been set.

### GetLinuxPassword

`func (o *UserSettingsUpdate) GetLinuxPassword() string`

GetLinuxPassword returns the LinuxPassword field if non-nil, zero value otherwise.

### GetLinuxPasswordOk

`func (o *UserSettingsUpdate) GetLinuxPasswordOk() (*string, bool)`

GetLinuxPasswordOk returns a tuple with the LinuxPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxPassword

`func (o *UserSettingsUpdate) SetLinuxPassword(v string)`

SetLinuxPassword sets LinuxPassword field to given value.

### HasLinuxPassword

`func (o *UserSettingsUpdate) HasLinuxPassword() bool`

HasLinuxPassword returns a boolean if a field has been set.

### GetLinuxKeyPairId

`func (o *UserSettingsUpdate) GetLinuxKeyPairId() int64`

GetLinuxKeyPairId returns the LinuxKeyPairId field if non-nil, zero value otherwise.

### GetLinuxKeyPairIdOk

`func (o *UserSettingsUpdate) GetLinuxKeyPairIdOk() (*int64, bool)`

GetLinuxKeyPairIdOk returns a tuple with the LinuxKeyPairId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxKeyPairId

`func (o *UserSettingsUpdate) SetLinuxKeyPairId(v int64)`

SetLinuxKeyPairId sets LinuxKeyPairId field to given value.

### HasLinuxKeyPairId

`func (o *UserSettingsUpdate) HasLinuxKeyPairId() bool`

HasLinuxKeyPairId returns a boolean if a field has been set.

### GetWindowsUsername

`func (o *UserSettingsUpdate) GetWindowsUsername() string`

GetWindowsUsername returns the WindowsUsername field if non-nil, zero value otherwise.

### GetWindowsUsernameOk

`func (o *UserSettingsUpdate) GetWindowsUsernameOk() (*string, bool)`

GetWindowsUsernameOk returns a tuple with the WindowsUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsUsername

`func (o *UserSettingsUpdate) SetWindowsUsername(v string)`

SetWindowsUsername sets WindowsUsername field to given value.

### HasWindowsUsername

`func (o *UserSettingsUpdate) HasWindowsUsername() bool`

HasWindowsUsername returns a boolean if a field has been set.

### GetWindowsPassword

`func (o *UserSettingsUpdate) GetWindowsPassword() string`

GetWindowsPassword returns the WindowsPassword field if non-nil, zero value otherwise.

### GetWindowsPasswordOk

`func (o *UserSettingsUpdate) GetWindowsPasswordOk() (*string, bool)`

GetWindowsPasswordOk returns a tuple with the WindowsPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWindowsPassword

`func (o *UserSettingsUpdate) SetWindowsPassword(v string)`

SetWindowsPassword sets WindowsPassword field to given value.

### HasWindowsPassword

`func (o *UserSettingsUpdate) HasWindowsPassword() bool`

HasWindowsPassword returns a boolean if a field has been set.

### GetReceiveNotifications

`func (o *UserSettingsUpdate) GetReceiveNotifications() bool`

GetReceiveNotifications returns the ReceiveNotifications field if non-nil, zero value otherwise.

### GetReceiveNotificationsOk

`func (o *UserSettingsUpdate) GetReceiveNotificationsOk() (*bool, bool)`

GetReceiveNotificationsOk returns a tuple with the ReceiveNotifications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceiveNotifications

`func (o *UserSettingsUpdate) SetReceiveNotifications(v bool)`

SetReceiveNotifications sets ReceiveNotifications field to given value.

### HasReceiveNotifications

`func (o *UserSettingsUpdate) HasReceiveNotifications() bool`

HasReceiveNotifications returns a boolean if a field has been set.

### GetDefaultGroup

`func (o *UserSettingsUpdate) GetDefaultGroup() UpdateUserSettingsRequestUserDefaultGroup`

GetDefaultGroup returns the DefaultGroup field if non-nil, zero value otherwise.

### GetDefaultGroupOk

`func (o *UserSettingsUpdate) GetDefaultGroupOk() (*UpdateUserSettingsRequestUserDefaultGroup, bool)`

GetDefaultGroupOk returns a tuple with the DefaultGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGroup

`func (o *UserSettingsUpdate) SetDefaultGroup(v UpdateUserSettingsRequestUserDefaultGroup)`

SetDefaultGroup sets DefaultGroup field to given value.

### HasDefaultGroup

`func (o *UserSettingsUpdate) HasDefaultGroup() bool`

HasDefaultGroup returns a boolean if a field has been set.

### GetDefaultCloud

`func (o *UserSettingsUpdate) GetDefaultCloud() UpdateUserSettingsRequestUserDefaultCloud`

GetDefaultCloud returns the DefaultCloud field if non-nil, zero value otherwise.

### GetDefaultCloudOk

`func (o *UserSettingsUpdate) GetDefaultCloudOk() (*UpdateUserSettingsRequestUserDefaultCloud, bool)`

GetDefaultCloudOk returns a tuple with the DefaultCloud field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCloud

`func (o *UserSettingsUpdate) SetDefaultCloud(v UpdateUserSettingsRequestUserDefaultCloud)`

SetDefaultCloud sets DefaultCloud field to given value.

### HasDefaultCloud

`func (o *UserSettingsUpdate) HasDefaultCloud() bool`

HasDefaultCloud returns a boolean if a field has been set.

### GetDefaultPersona

`func (o *UserSettingsUpdate) GetDefaultPersona() UpdateUserSettingsRequestUserDefaultPersona`

GetDefaultPersona returns the DefaultPersona field if non-nil, zero value otherwise.

### GetDefaultPersonaOk

`func (o *UserSettingsUpdate) GetDefaultPersonaOk() (*UpdateUserSettingsRequestUserDefaultPersona, bool)`

GetDefaultPersonaOk returns a tuple with the DefaultPersona field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultPersona

`func (o *UserSettingsUpdate) SetDefaultPersona(v UpdateUserSettingsRequestUserDefaultPersona)`

SetDefaultPersona sets DefaultPersona field to given value.

### HasDefaultPersona

`func (o *UserSettingsUpdate) HasDefaultPersona() bool`

HasDefaultPersona returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


