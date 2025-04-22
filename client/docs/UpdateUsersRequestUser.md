# UpdateUsersRequestUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | Pointer to **string** | First Name | [optional] 
**LastName** | Pointer to **string** | Last Name | [optional] 
**Username** | Pointer to **string** | Username (unique per tenant). | [optional] 
**Email** | Pointer to **string** | Email address | [optional] 
**Password** | Pointer to **string** | Password | [optional] 
**Roles** | Pointer to [**[]UpdateUsersRequestUserRolesInner**](UpdateUsersRequestUserRolesInner.md) | List of Roles | [optional] 

## Methods

### NewUpdateUsersRequestUser

`func NewUpdateUsersRequestUser() *UpdateUsersRequestUser`

NewUpdateUsersRequestUser instantiates a new UpdateUsersRequestUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateUsersRequestUserWithDefaults

`func NewUpdateUsersRequestUserWithDefaults() *UpdateUsersRequestUser`

NewUpdateUsersRequestUserWithDefaults instantiates a new UpdateUsersRequestUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *UpdateUsersRequestUser) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UpdateUsersRequestUser) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UpdateUsersRequestUser) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.

### HasFirstName

`func (o *UpdateUsersRequestUser) HasFirstName() bool`

HasFirstName returns a boolean if a field has been set.

### GetLastName

`func (o *UpdateUsersRequestUser) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UpdateUsersRequestUser) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UpdateUsersRequestUser) SetLastName(v string)`

SetLastName sets LastName field to given value.

### HasLastName

`func (o *UpdateUsersRequestUser) HasLastName() bool`

HasLastName returns a boolean if a field has been set.

### GetUsername

`func (o *UpdateUsersRequestUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UpdateUsersRequestUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UpdateUsersRequestUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UpdateUsersRequestUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetEmail

`func (o *UpdateUsersRequestUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UpdateUsersRequestUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UpdateUsersRequestUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UpdateUsersRequestUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetPassword

`func (o *UpdateUsersRequestUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdateUsersRequestUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdateUsersRequestUser) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UpdateUsersRequestUser) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetRoles

`func (o *UpdateUsersRequestUser) GetRoles() []UpdateUsersRequestUserRolesInner`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UpdateUsersRequestUser) GetRolesOk() (*[]UpdateUsersRequestUserRolesInner, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UpdateUsersRequestUser) SetRoles(v []UpdateUsersRequestUserRolesInner)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UpdateUsersRequestUser) HasRoles() bool`

HasRoles returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


