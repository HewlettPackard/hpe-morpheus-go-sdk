# UserSourceCreateActiveDirectory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | AD Server IP/FQDN | [optional] 
**Domain** | Pointer to **string** | Domain | [optional] 
**UseSSL** | Pointer to **string** | Use SSL (\&quot;on\&quot; or \&quot;off\&quot;) | [optional] 
**BindingUsername** | Pointer to **string** | Binding Username | [optional] 
**BindingPassword** | Pointer to **string** | Binding Password | [optional] 
**RequiredGroup** | Pointer to **string** | Required Group | [optional] 
**SearchMemberGroups** | Pointer to **bool** | Include Member Groups | [optional] [default to false]

## Methods

### NewUserSourceCreateActiveDirectory

`func NewUserSourceCreateActiveDirectory() *UserSourceCreateActiveDirectory`

NewUserSourceCreateActiveDirectory instantiates a new UserSourceCreateActiveDirectory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSourceCreateActiveDirectoryWithDefaults

`func NewUserSourceCreateActiveDirectoryWithDefaults() *UserSourceCreateActiveDirectory`

NewUserSourceCreateActiveDirectoryWithDefaults instantiates a new UserSourceCreateActiveDirectory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *UserSourceCreateActiveDirectory) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UserSourceCreateActiveDirectory) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UserSourceCreateActiveDirectory) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UserSourceCreateActiveDirectory) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDomain

`func (o *UserSourceCreateActiveDirectory) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *UserSourceCreateActiveDirectory) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *UserSourceCreateActiveDirectory) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *UserSourceCreateActiveDirectory) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetUseSSL

`func (o *UserSourceCreateActiveDirectory) GetUseSSL() string`

GetUseSSL returns the UseSSL field if non-nil, zero value otherwise.

### GetUseSSLOk

`func (o *UserSourceCreateActiveDirectory) GetUseSSLOk() (*string, bool)`

GetUseSSLOk returns a tuple with the UseSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSSL

`func (o *UserSourceCreateActiveDirectory) SetUseSSL(v string)`

SetUseSSL sets UseSSL field to given value.

### HasUseSSL

`func (o *UserSourceCreateActiveDirectory) HasUseSSL() bool`

HasUseSSL returns a boolean if a field has been set.

### GetBindingUsername

`func (o *UserSourceCreateActiveDirectory) GetBindingUsername() string`

GetBindingUsername returns the BindingUsername field if non-nil, zero value otherwise.

### GetBindingUsernameOk

`func (o *UserSourceCreateActiveDirectory) GetBindingUsernameOk() (*string, bool)`

GetBindingUsernameOk returns a tuple with the BindingUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingUsername

`func (o *UserSourceCreateActiveDirectory) SetBindingUsername(v string)`

SetBindingUsername sets BindingUsername field to given value.

### HasBindingUsername

`func (o *UserSourceCreateActiveDirectory) HasBindingUsername() bool`

HasBindingUsername returns a boolean if a field has been set.

### GetBindingPassword

`func (o *UserSourceCreateActiveDirectory) GetBindingPassword() string`

GetBindingPassword returns the BindingPassword field if non-nil, zero value otherwise.

### GetBindingPasswordOk

`func (o *UserSourceCreateActiveDirectory) GetBindingPasswordOk() (*string, bool)`

GetBindingPasswordOk returns a tuple with the BindingPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingPassword

`func (o *UserSourceCreateActiveDirectory) SetBindingPassword(v string)`

SetBindingPassword sets BindingPassword field to given value.

### HasBindingPassword

`func (o *UserSourceCreateActiveDirectory) HasBindingPassword() bool`

HasBindingPassword returns a boolean if a field has been set.

### GetRequiredGroup

`func (o *UserSourceCreateActiveDirectory) GetRequiredGroup() string`

GetRequiredGroup returns the RequiredGroup field if non-nil, zero value otherwise.

### GetRequiredGroupOk

`func (o *UserSourceCreateActiveDirectory) GetRequiredGroupOk() (*string, bool)`

GetRequiredGroupOk returns a tuple with the RequiredGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredGroup

`func (o *UserSourceCreateActiveDirectory) SetRequiredGroup(v string)`

SetRequiredGroup sets RequiredGroup field to given value.

### HasRequiredGroup

`func (o *UserSourceCreateActiveDirectory) HasRequiredGroup() bool`

HasRequiredGroup returns a boolean if a field has been set.

### GetSearchMemberGroups

`func (o *UserSourceCreateActiveDirectory) GetSearchMemberGroups() bool`

GetSearchMemberGroups returns the SearchMemberGroups field if non-nil, zero value otherwise.

### GetSearchMemberGroupsOk

`func (o *UserSourceCreateActiveDirectory) GetSearchMemberGroupsOk() (*bool, bool)`

GetSearchMemberGroupsOk returns a tuple with the SearchMemberGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchMemberGroups

`func (o *UserSourceCreateActiveDirectory) SetSearchMemberGroups(v bool)`

SetSearchMemberGroups sets SearchMemberGroups field to given value.

### HasSearchMemberGroups

`func (o *UserSourceCreateActiveDirectory) HasSearchMemberGroups() bool`

HasSearchMemberGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


