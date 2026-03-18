# UserSourceCreateLDAP

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | URL of Endpoint | [optional] 
**BindingUsername** | Pointer to **string** | Binding Username | [optional] 
**BindingPassword** | Pointer to **string** | Binding Password | [optional] 
**RequiredGroup** | Pointer to **string** | User DN Expression | [optional] 

## Methods

### NewUserSourceCreateLDAP

`func NewUserSourceCreateLDAP() *UserSourceCreateLDAP`

NewUserSourceCreateLDAP instantiates a new UserSourceCreateLDAP object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSourceCreateLDAPWithDefaults

`func NewUserSourceCreateLDAPWithDefaults() *UserSourceCreateLDAP`

NewUserSourceCreateLDAPWithDefaults instantiates a new UserSourceCreateLDAP object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *UserSourceCreateLDAP) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UserSourceCreateLDAP) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UserSourceCreateLDAP) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UserSourceCreateLDAP) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetBindingUsername

`func (o *UserSourceCreateLDAP) GetBindingUsername() string`

GetBindingUsername returns the BindingUsername field if non-nil, zero value otherwise.

### GetBindingUsernameOk

`func (o *UserSourceCreateLDAP) GetBindingUsernameOk() (*string, bool)`

GetBindingUsernameOk returns a tuple with the BindingUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingUsername

`func (o *UserSourceCreateLDAP) SetBindingUsername(v string)`

SetBindingUsername sets BindingUsername field to given value.

### HasBindingUsername

`func (o *UserSourceCreateLDAP) HasBindingUsername() bool`

HasBindingUsername returns a boolean if a field has been set.

### GetBindingPassword

`func (o *UserSourceCreateLDAP) GetBindingPassword() string`

GetBindingPassword returns the BindingPassword field if non-nil, zero value otherwise.

### GetBindingPasswordOk

`func (o *UserSourceCreateLDAP) GetBindingPasswordOk() (*string, bool)`

GetBindingPasswordOk returns a tuple with the BindingPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingPassword

`func (o *UserSourceCreateLDAP) SetBindingPassword(v string)`

SetBindingPassword sets BindingPassword field to given value.

### HasBindingPassword

`func (o *UserSourceCreateLDAP) HasBindingPassword() bool`

HasBindingPassword returns a boolean if a field has been set.

### GetRequiredGroup

`func (o *UserSourceCreateLDAP) GetRequiredGroup() string`

GetRequiredGroup returns the RequiredGroup field if non-nil, zero value otherwise.

### GetRequiredGroupOk

`func (o *UserSourceCreateLDAP) GetRequiredGroupOk() (*string, bool)`

GetRequiredGroupOk returns a tuple with the RequiredGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredGroup

`func (o *UserSourceCreateLDAP) SetRequiredGroup(v string)`

SetRequiredGroup sets RequiredGroup field to given value.

### HasRequiredGroup

`func (o *UserSourceCreateLDAP) HasRequiredGroup() bool`

HasRequiredGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


