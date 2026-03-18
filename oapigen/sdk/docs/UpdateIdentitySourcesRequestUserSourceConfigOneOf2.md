# UpdateIdentitySourcesRequestUserSourceConfigOneOf2

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

### NewUpdateIdentitySourcesRequestUserSourceConfigOneOf2

`func NewUpdateIdentitySourcesRequestUserSourceConfigOneOf2() *UpdateIdentitySourcesRequestUserSourceConfigOneOf2`

NewUpdateIdentitySourcesRequestUserSourceConfigOneOf2 instantiates a new UpdateIdentitySourcesRequestUserSourceConfigOneOf2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIdentitySourcesRequestUserSourceConfigOneOf2WithDefaults

`func NewUpdateIdentitySourcesRequestUserSourceConfigOneOf2WithDefaults() *UpdateIdentitySourcesRequestUserSourceConfigOneOf2`

NewUpdateIdentitySourcesRequestUserSourceConfigOneOf2WithDefaults instantiates a new UpdateIdentitySourcesRequestUserSourceConfigOneOf2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *UpdateIdentitySourcesRequestUserSourceConfigOneOf2) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UpdateIdentitySourcesRequestUserSourceConfigOneOf2) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UpdateIdentitySourcesRequestUserSourceConfigOneOf2) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UpdateIdentitySourcesRequestUserSourceConfigOneOf2) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDomain

`func (o *UpdateIdentitySourcesRequestUserSourceConfigOneOf2) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *UpdateIdentitySourcesRequestUserSourceConfigOneOf2) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *UpdateIdentitySourcesRequestUserSourceConfigOneOf2) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *UpdateIdentitySourcesRequestUserSourceConfigOneOf2) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetUseSSL

`func (o *UpdateIdentitySourcesRequestUserSourceConfigOneOf2) GetUseSSL() string`

GetUseSSL returns the UseSSL field if non-nil, zero value otherwise.

### GetUseSSLOk

`func (o *UpdateIdentitySourcesRequestUserSourceConfigOneOf2) GetUseSSLOk() (*string, bool)`

GetUseSSLOk returns a tuple with the UseSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSSL

`func (o *UpdateIdentitySourcesRequestUserSourceConfigOneOf2) SetUseSSL(v string)`

SetUseSSL sets UseSSL field to given value.

### HasUseSSL

`func (o *UpdateIdentitySourcesRequestUserSourceConfigOneOf2) HasUseSSL() bool`

HasUseSSL returns a boolean if a field has been set.

### GetBindingUsername

`func (o *UpdateIdentitySourcesRequestUserSourceConfigOneOf2) GetBindingUsername() string`

GetBindingUsername returns the BindingUsername field if non-nil, zero value otherwise.

### GetBindingUsernameOk

`func (o *UpdateIdentitySourcesRequestUserSourceConfigOneOf2) GetBindingUsernameOk() (*string, bool)`

GetBindingUsernameOk returns a tuple with the BindingUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingUsername

`func (o *UpdateIdentitySourcesRequestUserSourceConfigOneOf2) SetBindingUsername(v string)`

SetBindingUsername sets BindingUsername field to given value.

### HasBindingUsername

`func (o *UpdateIdentitySourcesRequestUserSourceConfigOneOf2) HasBindingUsername() bool`

HasBindingUsername returns a boolean if a field has been set.

### GetBindingPassword

`func (o *UpdateIdentitySourcesRequestUserSourceConfigOneOf2) GetBindingPassword() string`

GetBindingPassword returns the BindingPassword field if non-nil, zero value otherwise.

### GetBindingPasswordOk

`func (o *UpdateIdentitySourcesRequestUserSourceConfigOneOf2) GetBindingPasswordOk() (*string, bool)`

GetBindingPasswordOk returns a tuple with the BindingPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingPassword

`func (o *UpdateIdentitySourcesRequestUserSourceConfigOneOf2) SetBindingPassword(v string)`

SetBindingPassword sets BindingPassword field to given value.

### HasBindingPassword

`func (o *UpdateIdentitySourcesRequestUserSourceConfigOneOf2) HasBindingPassword() bool`

HasBindingPassword returns a boolean if a field has been set.

### GetRequiredGroup

`func (o *UpdateIdentitySourcesRequestUserSourceConfigOneOf2) GetRequiredGroup() string`

GetRequiredGroup returns the RequiredGroup field if non-nil, zero value otherwise.

### GetRequiredGroupOk

`func (o *UpdateIdentitySourcesRequestUserSourceConfigOneOf2) GetRequiredGroupOk() (*string, bool)`

GetRequiredGroupOk returns a tuple with the RequiredGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredGroup

`func (o *UpdateIdentitySourcesRequestUserSourceConfigOneOf2) SetRequiredGroup(v string)`

SetRequiredGroup sets RequiredGroup field to given value.

### HasRequiredGroup

`func (o *UpdateIdentitySourcesRequestUserSourceConfigOneOf2) HasRequiredGroup() bool`

HasRequiredGroup returns a boolean if a field has been set.

### GetSearchMemberGroups

`func (o *UpdateIdentitySourcesRequestUserSourceConfigOneOf2) GetSearchMemberGroups() bool`

GetSearchMemberGroups returns the SearchMemberGroups field if non-nil, zero value otherwise.

### GetSearchMemberGroupsOk

`func (o *UpdateIdentitySourcesRequestUserSourceConfigOneOf2) GetSearchMemberGroupsOk() (*bool, bool)`

GetSearchMemberGroupsOk returns a tuple with the SearchMemberGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchMemberGroups

`func (o *UpdateIdentitySourcesRequestUserSourceConfigOneOf2) SetSearchMemberGroups(v bool)`

SetSearchMemberGroups sets SearchMemberGroups field to given value.

### HasSearchMemberGroups

`func (o *UpdateIdentitySourcesRequestUserSourceConfigOneOf2) HasSearchMemberGroups() bool`

HasSearchMemberGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


