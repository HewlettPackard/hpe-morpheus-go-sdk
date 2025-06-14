/*
Morpheus API

Morpheus is a powerful cloud management tool that provides provisioning, monitoring, logging, backups, and application deployment strategies.  This document describes the Morpheus API protocol and the available endpoints. Sections are organized in the same manner as they appear in the Morpheus UI.

API version: 8.0.7
Contact: dev@morpheusdata.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sdk

import (
	"encoding/json"
)

// checks if the AddUserTenantRequestUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AddUserTenantRequestUser{}

// AddUserTenantRequestUser Payload for creating a new user
type AddUserTenantRequestUser struct {
	// The user's first name (optional)
	FirstName *string `json:"firstName,omitempty"`
	// The user's last name (optional)
	LastName *string `json:"lastName,omitempty"`
	// Username (unique per tenant).
	Username string `json:"username"`
	// Email address
	Email string `json:"email"`
	// Password to apply to the user
	Password string `json:"password"`
	// Array of objects with id of the role(s) to assign to the user.
	Roles []GetAlerts200ResponseAllOfChecksInnerAccount `json:"roles"`
	// Receive Notifications?
	ReceiveNotifications *bool `json:"receiveNotifications,omitempty"`
	// Linux Username, user settings for provisioning
	LinuxUsername *string `json:"linuxUsername,omitempty"`
	// Linux Password, user settings for provisioning
	LinuxPassword *string `json:"linuxPassword,omitempty"`
	// Linux SSH Key, user settings for provisioning
	LinuxKeyPairId *int64 `json:"linuxKeyPairId,omitempty"`
	// Windows Username, user settings for provisioning
	WindowsUsername *string `json:"windowsUsername,omitempty"`
	// Windows Password, user settings for provisioning
	WindowsPassword      *string                `json:"windowsPassword,omitempty"`
	AdditionalProperties map[string]interface{} `json:",remain"`
}

type _AddUserTenantRequestUser AddUserTenantRequestUser

// NewAddUserTenantRequestUser instantiates a new AddUserTenantRequestUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAddUserTenantRequestUser(username string, email string, password string, roles []GetAlerts200ResponseAllOfChecksInnerAccount) *AddUserTenantRequestUser {
	this := AddUserTenantRequestUser{}
	this.Username = username
	this.Email = email
	this.Password = password
	this.Roles = roles
	var receiveNotifications bool = true
	this.ReceiveNotifications = &receiveNotifications
	return &this
}

// NewAddUserTenantRequestUserWithDefaults instantiates a new AddUserTenantRequestUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAddUserTenantRequestUserWithDefaults() *AddUserTenantRequestUser {
	this := AddUserTenantRequestUser{}
	var receiveNotifications bool = true
	this.ReceiveNotifications = &receiveNotifications
	return &this
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *AddUserTenantRequestUser) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddUserTenantRequestUser) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// IsSetFirstName returns a boolean if a field has been set.
func (o *AddUserTenantRequestUser) IsSetFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *AddUserTenantRequestUser) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *AddUserTenantRequestUser) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddUserTenantRequestUser) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// IsSetLastName returns a boolean if a field has been set.
func (o *AddUserTenantRequestUser) IsSetLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *AddUserTenantRequestUser) SetLastName(v string) {
	o.LastName = &v
}

// GetUsername returns the Username field value
func (o *AddUserTenantRequestUser) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *AddUserTenantRequestUser) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *AddUserTenantRequestUser) SetUsername(v string) {
	o.Username = v
}

// GetEmail returns the Email field value
func (o *AddUserTenantRequestUser) GetEmail() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Email
}

// GetEmailOk returns a tuple with the Email field value
// and a boolean to check if the value has been set.
func (o *AddUserTenantRequestUser) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Email, true
}

// SetEmail sets field value
func (o *AddUserTenantRequestUser) SetEmail(v string) {
	o.Email = v
}

// GetPassword returns the Password field value
func (o *AddUserTenantRequestUser) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *AddUserTenantRequestUser) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *AddUserTenantRequestUser) SetPassword(v string) {
	o.Password = v
}

// GetRoles returns the Roles field value
func (o *AddUserTenantRequestUser) GetRoles() []GetAlerts200ResponseAllOfChecksInnerAccount {
	if o == nil {
		var ret []GetAlerts200ResponseAllOfChecksInnerAccount
		return ret
	}

	return o.Roles
}

// GetRolesOk returns a tuple with the Roles field value
// and a boolean to check if the value has been set.
func (o *AddUserTenantRequestUser) GetRolesOk() ([]GetAlerts200ResponseAllOfChecksInnerAccount, bool) {
	if o == nil {
		return nil, false
	}
	return o.Roles, true
}

// SetRoles sets field value
func (o *AddUserTenantRequestUser) SetRoles(v []GetAlerts200ResponseAllOfChecksInnerAccount) {
	o.Roles = v
}

// GetReceiveNotifications returns the ReceiveNotifications field value if set, zero value otherwise.
func (o *AddUserTenantRequestUser) GetReceiveNotifications() bool {
	if o == nil || IsNil(o.ReceiveNotifications) {
		var ret bool
		return ret
	}
	return *o.ReceiveNotifications
}

// GetReceiveNotificationsOk returns a tuple with the ReceiveNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddUserTenantRequestUser) GetReceiveNotificationsOk() (*bool, bool) {
	if o == nil || IsNil(o.ReceiveNotifications) {
		return nil, false
	}
	return o.ReceiveNotifications, true
}

// IsSetReceiveNotifications returns a boolean if a field has been set.
func (o *AddUserTenantRequestUser) IsSetReceiveNotifications() bool {
	if o != nil && !IsNil(o.ReceiveNotifications) {
		return true
	}

	return false
}

// SetReceiveNotifications gets a reference to the given bool and assigns it to the ReceiveNotifications field.
func (o *AddUserTenantRequestUser) SetReceiveNotifications(v bool) {
	o.ReceiveNotifications = &v
}

// GetLinuxUsername returns the LinuxUsername field value if set, zero value otherwise.
func (o *AddUserTenantRequestUser) GetLinuxUsername() string {
	if o == nil || IsNil(o.LinuxUsername) {
		var ret string
		return ret
	}
	return *o.LinuxUsername
}

// GetLinuxUsernameOk returns a tuple with the LinuxUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddUserTenantRequestUser) GetLinuxUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.LinuxUsername) {
		return nil, false
	}
	return o.LinuxUsername, true
}

// IsSetLinuxUsername returns a boolean if a field has been set.
func (o *AddUserTenantRequestUser) IsSetLinuxUsername() bool {
	if o != nil && !IsNil(o.LinuxUsername) {
		return true
	}

	return false
}

// SetLinuxUsername gets a reference to the given string and assigns it to the LinuxUsername field.
func (o *AddUserTenantRequestUser) SetLinuxUsername(v string) {
	o.LinuxUsername = &v
}

// GetLinuxPassword returns the LinuxPassword field value if set, zero value otherwise.
func (o *AddUserTenantRequestUser) GetLinuxPassword() string {
	if o == nil || IsNil(o.LinuxPassword) {
		var ret string
		return ret
	}
	return *o.LinuxPassword
}

// GetLinuxPasswordOk returns a tuple with the LinuxPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddUserTenantRequestUser) GetLinuxPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.LinuxPassword) {
		return nil, false
	}
	return o.LinuxPassword, true
}

// IsSetLinuxPassword returns a boolean if a field has been set.
func (o *AddUserTenantRequestUser) IsSetLinuxPassword() bool {
	if o != nil && !IsNil(o.LinuxPassword) {
		return true
	}

	return false
}

// SetLinuxPassword gets a reference to the given string and assigns it to the LinuxPassword field.
func (o *AddUserTenantRequestUser) SetLinuxPassword(v string) {
	o.LinuxPassword = &v
}

// GetLinuxKeyPairId returns the LinuxKeyPairId field value if set, zero value otherwise.
func (o *AddUserTenantRequestUser) GetLinuxKeyPairId() int64 {
	if o == nil || IsNil(o.LinuxKeyPairId) {
		var ret int64
		return ret
	}
	return *o.LinuxKeyPairId
}

// GetLinuxKeyPairIdOk returns a tuple with the LinuxKeyPairId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddUserTenantRequestUser) GetLinuxKeyPairIdOk() (*int64, bool) {
	if o == nil || IsNil(o.LinuxKeyPairId) {
		return nil, false
	}
	return o.LinuxKeyPairId, true
}

// IsSetLinuxKeyPairId returns a boolean if a field has been set.
func (o *AddUserTenantRequestUser) IsSetLinuxKeyPairId() bool {
	if o != nil && !IsNil(o.LinuxKeyPairId) {
		return true
	}

	return false
}

// SetLinuxKeyPairId gets a reference to the given int64 and assigns it to the LinuxKeyPairId field.
func (o *AddUserTenantRequestUser) SetLinuxKeyPairId(v int64) {
	o.LinuxKeyPairId = &v
}

// GetWindowsUsername returns the WindowsUsername field value if set, zero value otherwise.
func (o *AddUserTenantRequestUser) GetWindowsUsername() string {
	if o == nil || IsNil(o.WindowsUsername) {
		var ret string
		return ret
	}
	return *o.WindowsUsername
}

// GetWindowsUsernameOk returns a tuple with the WindowsUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddUserTenantRequestUser) GetWindowsUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.WindowsUsername) {
		return nil, false
	}
	return o.WindowsUsername, true
}

// IsSetWindowsUsername returns a boolean if a field has been set.
func (o *AddUserTenantRequestUser) IsSetWindowsUsername() bool {
	if o != nil && !IsNil(o.WindowsUsername) {
		return true
	}

	return false
}

// SetWindowsUsername gets a reference to the given string and assigns it to the WindowsUsername field.
func (o *AddUserTenantRequestUser) SetWindowsUsername(v string) {
	o.WindowsUsername = &v
}

// GetWindowsPassword returns the WindowsPassword field value if set, zero value otherwise.
func (o *AddUserTenantRequestUser) GetWindowsPassword() string {
	if o == nil || IsNil(o.WindowsPassword) {
		var ret string
		return ret
	}
	return *o.WindowsPassword
}

// GetWindowsPasswordOk returns a tuple with the WindowsPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AddUserTenantRequestUser) GetWindowsPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.WindowsPassword) {
		return nil, false
	}
	return o.WindowsPassword, true
}

// IsSetWindowsPassword returns a boolean if a field has been set.
func (o *AddUserTenantRequestUser) IsSetWindowsPassword() bool {
	if o != nil && !IsNil(o.WindowsPassword) {
		return true
	}

	return false
}

// SetWindowsPassword gets a reference to the given string and assigns it to the WindowsPassword field.
func (o *AddUserTenantRequestUser) SetWindowsPassword(v string) {
	o.WindowsPassword = &v
}

func (o AddUserTenantRequestUser) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AddUserTenantRequestUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	toSerialize["username"] = o.Username
	toSerialize["email"] = o.Email
	toSerialize["password"] = o.Password
	toSerialize["roles"] = o.Roles
	if !IsNil(o.ReceiveNotifications) {
		toSerialize["receiveNotifications"] = o.ReceiveNotifications
	}
	if !IsNil(o.LinuxUsername) {
		toSerialize["linuxUsername"] = o.LinuxUsername
	}
	if !IsNil(o.LinuxPassword) {
		toSerialize["linuxPassword"] = o.LinuxPassword
	}
	if !IsNil(o.LinuxKeyPairId) {
		toSerialize["linuxKeyPairId"] = o.LinuxKeyPairId
	}
	if !IsNil(o.WindowsUsername) {
		toSerialize["windowsUsername"] = o.WindowsUsername
	}
	if !IsNil(o.WindowsPassword) {
		toSerialize["windowsPassword"] = o.WindowsPassword
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *AddUserTenantRequestUser) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache
