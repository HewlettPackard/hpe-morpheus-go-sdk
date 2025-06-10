# UpdateUserRequestUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** | First Name | [optional] 
**LastName** | Pointer to **string** | Last Name | [optional] 
**Username** | Pointer to **string** | Username (unique per tenant). | [optional] 
**LinuxUsername** | Pointer to **string** |  | [optional] 
**LinuxPassword** | Pointer to **string** |  | [optional] 
**LinuxKeyPairId** | Pointer to **string** |  | [optional] 
**WindowsUsername** | Pointer to **string** |  | [optional] 
**WindowsPassword** | Pointer to **string** |  | [optional] 
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

### GetLinuxKeyPairId

`func (o *UpdateUserRequestUser) GetLinuxKeyPairId() string`

GetLinuxKeyPairId returns the LinuxKeyPairId field if non-nil, zero value otherwise.

### GetLinuxKeyPairIdOk

`func (o *UpdateUserRequestUser) GetLinuxKeyPairIdOk() (*string, bool)`

GetLinuxKeyPairIdOk returns a tuple with the LinuxKeyPairId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinuxKeyPairId

`func (o *UpdateUserRequestUser) SetLinuxKeyPairId(v string)`

SetLinuxKeyPairId sets LinuxKeyPairId field to given value.

### HasLinuxKeyPairId

`func (o *UpdateUserRequestUser) HasLinuxKeyPairId() bool`

HasLinuxKeyPairId returns a boolean if a field has been set.

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


