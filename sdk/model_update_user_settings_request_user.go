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

// checks if the UpdateUserSettingsRequestUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateUserSettingsRequestUser{}

// UpdateUserSettingsRequestUser struct for UpdateUserSettingsRequestUser
type UpdateUserSettingsRequestUser struct {
	// Username
	Username *string `json:"username,omitempty"`
	// Email
	Email *string `json:"email,omitempty"`
	// First Name
	FirstName *string `json:"firstName,omitempty"`
	// Last Name
	LastName *string `json:"lastName,omitempty"`
	// Change your password
	Password *string `json:"password,omitempty"`
	// Linux Username
	LinuxUsername *string `json:"linuxUsername,omitempty"`
	// Linux Password
	LinuxPassword *string `json:"linuxPassword,omitempty"`
	// Linux Key Pair ID
	LinuxKeyPairId *int64 `json:"linuxKeyPairId,omitempty"`
	// Windows Username
	WindowsUsername *string `json:"windowsUsername,omitempty"`
	// Windows Password
	WindowsPassword *string `json:"windowsPassword,omitempty"`
	// Receive Notifications (true or false)
	ReceiveNotifications *bool                                        `json:"receiveNotifications,omitempty"`
	DefaultGroup         *UpdateUserSettingsRequestUserDefaultGroup   `json:"defaultGroup,omitempty"`
	DefaultCloud         *UpdateUserSettingsRequestUserDefaultCloud   `json:"defaultCloud,omitempty"`
	DefaultPersona       *UpdateUserSettingsRequestUserDefaultPersona `json:"defaultPersona,omitempty"`
	AdditionalProperties map[string]interface{}                       `json:",remain"`
}

type _UpdateUserSettingsRequestUser UpdateUserSettingsRequestUser

// NewUpdateUserSettingsRequestUser instantiates a new UpdateUserSettingsRequestUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateUserSettingsRequestUser() *UpdateUserSettingsRequestUser {
	this := UpdateUserSettingsRequestUser{}
	return &this
}

// NewUpdateUserSettingsRequestUserWithDefaults instantiates a new UpdateUserSettingsRequestUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateUserSettingsRequestUserWithDefaults() *UpdateUserSettingsRequestUser {
	this := UpdateUserSettingsRequestUser{}
	return &this
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UpdateUserSettingsRequestUser) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserSettingsRequestUser) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// IsSetUsername returns a boolean if a field has been set.
func (o *UpdateUserSettingsRequestUser) IsSetUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UpdateUserSettingsRequestUser) SetUsername(v string) {
	o.Username = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UpdateUserSettingsRequestUser) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserSettingsRequestUser) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// IsSetEmail returns a boolean if a field has been set.
func (o *UpdateUserSettingsRequestUser) IsSetEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UpdateUserSettingsRequestUser) SetEmail(v string) {
	o.Email = &v
}

// GetFirstName returns the FirstName field value if set, zero value otherwise.
func (o *UpdateUserSettingsRequestUser) GetFirstName() string {
	if o == nil || IsNil(o.FirstName) {
		var ret string
		return ret
	}
	return *o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserSettingsRequestUser) GetFirstNameOk() (*string, bool) {
	if o == nil || IsNil(o.FirstName) {
		return nil, false
	}
	return o.FirstName, true
}

// IsSetFirstName returns a boolean if a field has been set.
func (o *UpdateUserSettingsRequestUser) IsSetFirstName() bool {
	if o != nil && !IsNil(o.FirstName) {
		return true
	}

	return false
}

// SetFirstName gets a reference to the given string and assigns it to the FirstName field.
func (o *UpdateUserSettingsRequestUser) SetFirstName(v string) {
	o.FirstName = &v
}

// GetLastName returns the LastName field value if set, zero value otherwise.
func (o *UpdateUserSettingsRequestUser) GetLastName() string {
	if o == nil || IsNil(o.LastName) {
		var ret string
		return ret
	}
	return *o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserSettingsRequestUser) GetLastNameOk() (*string, bool) {
	if o == nil || IsNil(o.LastName) {
		return nil, false
	}
	return o.LastName, true
}

// IsSetLastName returns a boolean if a field has been set.
func (o *UpdateUserSettingsRequestUser) IsSetLastName() bool {
	if o != nil && !IsNil(o.LastName) {
		return true
	}

	return false
}

// SetLastName gets a reference to the given string and assigns it to the LastName field.
func (o *UpdateUserSettingsRequestUser) SetLastName(v string) {
	o.LastName = &v
}

// GetPassword returns the Password field value if set, zero value otherwise.
func (o *UpdateUserSettingsRequestUser) GetPassword() string {
	if o == nil || IsNil(o.Password) {
		var ret string
		return ret
	}
	return *o.Password
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserSettingsRequestUser) GetPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.Password) {
		return nil, false
	}
	return o.Password, true
}

// IsSetPassword returns a boolean if a field has been set.
func (o *UpdateUserSettingsRequestUser) IsSetPassword() bool {
	if o != nil && !IsNil(o.Password) {
		return true
	}

	return false
}

// SetPassword gets a reference to the given string and assigns it to the Password field.
func (o *UpdateUserSettingsRequestUser) SetPassword(v string) {
	o.Password = &v
}

// GetLinuxUsername returns the LinuxUsername field value if set, zero value otherwise.
func (o *UpdateUserSettingsRequestUser) GetLinuxUsername() string {
	if o == nil || IsNil(o.LinuxUsername) {
		var ret string
		return ret
	}
	return *o.LinuxUsername
}

// GetLinuxUsernameOk returns a tuple with the LinuxUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserSettingsRequestUser) GetLinuxUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.LinuxUsername) {
		return nil, false
	}
	return o.LinuxUsername, true
}

// IsSetLinuxUsername returns a boolean if a field has been set.
func (o *UpdateUserSettingsRequestUser) IsSetLinuxUsername() bool {
	if o != nil && !IsNil(o.LinuxUsername) {
		return true
	}

	return false
}

// SetLinuxUsername gets a reference to the given string and assigns it to the LinuxUsername field.
func (o *UpdateUserSettingsRequestUser) SetLinuxUsername(v string) {
	o.LinuxUsername = &v
}

// GetLinuxPassword returns the LinuxPassword field value if set, zero value otherwise.
func (o *UpdateUserSettingsRequestUser) GetLinuxPassword() string {
	if o == nil || IsNil(o.LinuxPassword) {
		var ret string
		return ret
	}
	return *o.LinuxPassword
}

// GetLinuxPasswordOk returns a tuple with the LinuxPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserSettingsRequestUser) GetLinuxPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.LinuxPassword) {
		return nil, false
	}
	return o.LinuxPassword, true
}

// IsSetLinuxPassword returns a boolean if a field has been set.
func (o *UpdateUserSettingsRequestUser) IsSetLinuxPassword() bool {
	if o != nil && !IsNil(o.LinuxPassword) {
		return true
	}

	return false
}

// SetLinuxPassword gets a reference to the given string and assigns it to the LinuxPassword field.
func (o *UpdateUserSettingsRequestUser) SetLinuxPassword(v string) {
	o.LinuxPassword = &v
}

// GetLinuxKeyPairId returns the LinuxKeyPairId field value if set, zero value otherwise.
func (o *UpdateUserSettingsRequestUser) GetLinuxKeyPairId() int64 {
	if o == nil || IsNil(o.LinuxKeyPairId) {
		var ret int64
		return ret
	}
	return *o.LinuxKeyPairId
}

// GetLinuxKeyPairIdOk returns a tuple with the LinuxKeyPairId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserSettingsRequestUser) GetLinuxKeyPairIdOk() (*int64, bool) {
	if o == nil || IsNil(o.LinuxKeyPairId) {
		return nil, false
	}
	return o.LinuxKeyPairId, true
}

// IsSetLinuxKeyPairId returns a boolean if a field has been set.
func (o *UpdateUserSettingsRequestUser) IsSetLinuxKeyPairId() bool {
	if o != nil && !IsNil(o.LinuxKeyPairId) {
		return true
	}

	return false
}

// SetLinuxKeyPairId gets a reference to the given int64 and assigns it to the LinuxKeyPairId field.
func (o *UpdateUserSettingsRequestUser) SetLinuxKeyPairId(v int64) {
	o.LinuxKeyPairId = &v
}

// GetWindowsUsername returns the WindowsUsername field value if set, zero value otherwise.
func (o *UpdateUserSettingsRequestUser) GetWindowsUsername() string {
	if o == nil || IsNil(o.WindowsUsername) {
		var ret string
		return ret
	}
	return *o.WindowsUsername
}

// GetWindowsUsernameOk returns a tuple with the WindowsUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserSettingsRequestUser) GetWindowsUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.WindowsUsername) {
		return nil, false
	}
	return o.WindowsUsername, true
}

// IsSetWindowsUsername returns a boolean if a field has been set.
func (o *UpdateUserSettingsRequestUser) IsSetWindowsUsername() bool {
	if o != nil && !IsNil(o.WindowsUsername) {
		return true
	}

	return false
}

// SetWindowsUsername gets a reference to the given string and assigns it to the WindowsUsername field.
func (o *UpdateUserSettingsRequestUser) SetWindowsUsername(v string) {
	o.WindowsUsername = &v
}

// GetWindowsPassword returns the WindowsPassword field value if set, zero value otherwise.
func (o *UpdateUserSettingsRequestUser) GetWindowsPassword() string {
	if o == nil || IsNil(o.WindowsPassword) {
		var ret string
		return ret
	}
	return *o.WindowsPassword
}

// GetWindowsPasswordOk returns a tuple with the WindowsPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserSettingsRequestUser) GetWindowsPasswordOk() (*string, bool) {
	if o == nil || IsNil(o.WindowsPassword) {
		return nil, false
	}
	return o.WindowsPassword, true
}

// IsSetWindowsPassword returns a boolean if a field has been set.
func (o *UpdateUserSettingsRequestUser) IsSetWindowsPassword() bool {
	if o != nil && !IsNil(o.WindowsPassword) {
		return true
	}

	return false
}

// SetWindowsPassword gets a reference to the given string and assigns it to the WindowsPassword field.
func (o *UpdateUserSettingsRequestUser) SetWindowsPassword(v string) {
	o.WindowsPassword = &v
}

// GetReceiveNotifications returns the ReceiveNotifications field value if set, zero value otherwise.
func (o *UpdateUserSettingsRequestUser) GetReceiveNotifications() bool {
	if o == nil || IsNil(o.ReceiveNotifications) {
		var ret bool
		return ret
	}
	return *o.ReceiveNotifications
}

// GetReceiveNotificationsOk returns a tuple with the ReceiveNotifications field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserSettingsRequestUser) GetReceiveNotificationsOk() (*bool, bool) {
	if o == nil || IsNil(o.ReceiveNotifications) {
		return nil, false
	}
	return o.ReceiveNotifications, true
}

// IsSetReceiveNotifications returns a boolean if a field has been set.
func (o *UpdateUserSettingsRequestUser) IsSetReceiveNotifications() bool {
	if o != nil && !IsNil(o.ReceiveNotifications) {
		return true
	}

	return false
}

// SetReceiveNotifications gets a reference to the given bool and assigns it to the ReceiveNotifications field.
func (o *UpdateUserSettingsRequestUser) SetReceiveNotifications(v bool) {
	o.ReceiveNotifications = &v
}

// GetDefaultGroup returns the DefaultGroup field value if set, zero value otherwise.
func (o *UpdateUserSettingsRequestUser) GetDefaultGroup() UpdateUserSettingsRequestUserDefaultGroup {
	if o == nil || IsNil(o.DefaultGroup) {
		var ret UpdateUserSettingsRequestUserDefaultGroup
		return ret
	}
	return *o.DefaultGroup
}

// GetDefaultGroupOk returns a tuple with the DefaultGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserSettingsRequestUser) GetDefaultGroupOk() (*UpdateUserSettingsRequestUserDefaultGroup, bool) {
	if o == nil || IsNil(o.DefaultGroup) {
		return nil, false
	}
	return o.DefaultGroup, true
}

// IsSetDefaultGroup returns a boolean if a field has been set.
func (o *UpdateUserSettingsRequestUser) IsSetDefaultGroup() bool {
	if o != nil && !IsNil(o.DefaultGroup) {
		return true
	}

	return false
}

// SetDefaultGroup gets a reference to the given UpdateUserSettingsRequestUserDefaultGroup and assigns it to the DefaultGroup field.
func (o *UpdateUserSettingsRequestUser) SetDefaultGroup(v UpdateUserSettingsRequestUserDefaultGroup) {
	o.DefaultGroup = &v
}

// GetDefaultCloud returns the DefaultCloud field value if set, zero value otherwise.
func (o *UpdateUserSettingsRequestUser) GetDefaultCloud() UpdateUserSettingsRequestUserDefaultCloud {
	if o == nil || IsNil(o.DefaultCloud) {
		var ret UpdateUserSettingsRequestUserDefaultCloud
		return ret
	}
	return *o.DefaultCloud
}

// GetDefaultCloudOk returns a tuple with the DefaultCloud field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserSettingsRequestUser) GetDefaultCloudOk() (*UpdateUserSettingsRequestUserDefaultCloud, bool) {
	if o == nil || IsNil(o.DefaultCloud) {
		return nil, false
	}
	return o.DefaultCloud, true
}

// IsSetDefaultCloud returns a boolean if a field has been set.
func (o *UpdateUserSettingsRequestUser) IsSetDefaultCloud() bool {
	if o != nil && !IsNil(o.DefaultCloud) {
		return true
	}

	return false
}

// SetDefaultCloud gets a reference to the given UpdateUserSettingsRequestUserDefaultCloud and assigns it to the DefaultCloud field.
func (o *UpdateUserSettingsRequestUser) SetDefaultCloud(v UpdateUserSettingsRequestUserDefaultCloud) {
	o.DefaultCloud = &v
}

// GetDefaultPersona returns the DefaultPersona field value if set, zero value otherwise.
func (o *UpdateUserSettingsRequestUser) GetDefaultPersona() UpdateUserSettingsRequestUserDefaultPersona {
	if o == nil || IsNil(o.DefaultPersona) {
		var ret UpdateUserSettingsRequestUserDefaultPersona
		return ret
	}
	return *o.DefaultPersona
}

// GetDefaultPersonaOk returns a tuple with the DefaultPersona field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateUserSettingsRequestUser) GetDefaultPersonaOk() (*UpdateUserSettingsRequestUserDefaultPersona, bool) {
	if o == nil || IsNil(o.DefaultPersona) {
		return nil, false
	}
	return o.DefaultPersona, true
}

// IsSetDefaultPersona returns a boolean if a field has been set.
func (o *UpdateUserSettingsRequestUser) IsSetDefaultPersona() bool {
	if o != nil && !IsNil(o.DefaultPersona) {
		return true
	}

	return false
}

// SetDefaultPersona gets a reference to the given UpdateUserSettingsRequestUserDefaultPersona and assigns it to the DefaultPersona field.
func (o *UpdateUserSettingsRequestUser) SetDefaultPersona(v UpdateUserSettingsRequestUserDefaultPersona) {
	o.DefaultPersona = &v
}

func (o UpdateUserSettingsRequestUser) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateUserSettingsRequestUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	if !IsNil(o.FirstName) {
		toSerialize["firstName"] = o.FirstName
	}
	if !IsNil(o.LastName) {
		toSerialize["lastName"] = o.LastName
	}
	if !IsNil(o.Password) {
		toSerialize["password"] = o.Password
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
	if !IsNil(o.ReceiveNotifications) {
		toSerialize["receiveNotifications"] = o.ReceiveNotifications
	}
	if !IsNil(o.DefaultGroup) {
		toSerialize["defaultGroup"] = o.DefaultGroup
	}
	if !IsNil(o.DefaultCloud) {
		toSerialize["defaultCloud"] = o.DefaultCloud
	}
	if !IsNil(o.DefaultPersona) {
		toSerialize["defaultPersona"] = o.DefaultPersona
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}
func (o *UpdateUserSettingsRequestUser) UnmarshalJSON(data []byte) (err error) {
	return decode(data, &o)
}

// - model_simple.mustache
