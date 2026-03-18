# UserSourceCreateUserSourceConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | Login Redirect URL | [optional] 
**BindingUsername** | Pointer to **string** | Binding Username | [optional] 
**BindingPassword** | Pointer to **string** | Binding Password | [optional] 
**RequiredGroup** | Pointer to **string** | Required Group | [optional] 
**OrganizationId** | Pointer to **bool** | Organization ID | [optional] [default to false]
**RequiredRole** | Pointer to **string** | Required Role | [optional] 
**Domain** | Pointer to **string** | Domain | [optional] 
**UseSSL** | Pointer to **string** | Use SSL (\&quot;on\&quot; or \&quot;off\&quot;) | [optional] 
**SearchMemberGroups** | Pointer to **bool** | Include Member Groups | [optional] [default to false]
**AdministratorAPIToken** | Pointer to **string** | Administrator API Token | [optional] 
**Subdomain** | Pointer to **string** | OneLogin Subdomain | [optional] 
**Region** | Pointer to **string** | OneLogin Region | [optional] 
**ClientSecret** | Pointer to **string** | API Client Secret | [optional] 
**ClientId** | Pointer to **string** | API Client ID | [optional] 
**DoNotIncludeSAMLRequest** | Pointer to **bool** | Do not include SAMLRequest | [optional] [default to false]
**LogoutUrl** | Pointer to **string** | Lougout Post URL | [optional] 
**SAMLSignatureMode** | Pointer to **string** | SAML Request signing mode | [optional] [default to "NoSignature"]
**Request509Certificate** | Pointer to **string** | Only applies when &#x60;SAMLSignatureMode&#x3D;CustomSignature&#x60; | [optional] 
**RequestPrivateKey** | Pointer to **string** | RSA Private Key. Only applies when &#x60;SAMLSignatureMode&#x3D;CustomSignature&#x60; | [optional] 
**DoNotValidateSignature** | Pointer to **bool** | SAML Response Signing Flag | [optional] [default to true]
**PublicKey** | Pointer to **string** | Signing Public Key. Only applies when &#x60;doNotValidateSignature&#x3D;true&#x60; | [optional] 
**PrivateKey** | Pointer to **string** | Encryption RSA Private Key | [optional] 
**GivenNameAttribute** | Pointer to **string** | Given Name Attribute Name | [optional] 
**SurnameAttribute** | Pointer to **string** | Surname Attribute Name | [optional] 
**RoleAttributeName** | Pointer to **string** | Role Attribute Name | [optional] 
**RequiredAttributeValue** | Pointer to **string** | Role Attibute Required Value | [optional] 
**LoginUrl** | Pointer to **string** | External Login URL | [optional] 
**Logout** | Pointer to **string** | External Logout URL | [optional] 
**EncryptionAlgo** | Pointer to **string** | Encryption Algorithm | [optional] 
**EncryptionKey** | Pointer to **string** | Encryption Key | [optional] 
**Endpoint** | Pointer to **string** | API Endpoint | [optional] 
**ApiStyle** | Pointer to **string** | API Style | [optional] 

## Methods

### NewUserSourceCreateUserSourceConfig

`func NewUserSourceCreateUserSourceConfig() *UserSourceCreateUserSourceConfig`

NewUserSourceCreateUserSourceConfig instantiates a new UserSourceCreateUserSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSourceCreateUserSourceConfigWithDefaults

`func NewUserSourceCreateUserSourceConfigWithDefaults() *UserSourceCreateUserSourceConfig`

NewUserSourceCreateUserSourceConfigWithDefaults instantiates a new UserSourceCreateUserSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *UserSourceCreateUserSourceConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UserSourceCreateUserSourceConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UserSourceCreateUserSourceConfig) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UserSourceCreateUserSourceConfig) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetBindingUsername

`func (o *UserSourceCreateUserSourceConfig) GetBindingUsername() string`

GetBindingUsername returns the BindingUsername field if non-nil, zero value otherwise.

### GetBindingUsernameOk

`func (o *UserSourceCreateUserSourceConfig) GetBindingUsernameOk() (*string, bool)`

GetBindingUsernameOk returns a tuple with the BindingUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingUsername

`func (o *UserSourceCreateUserSourceConfig) SetBindingUsername(v string)`

SetBindingUsername sets BindingUsername field to given value.

### HasBindingUsername

`func (o *UserSourceCreateUserSourceConfig) HasBindingUsername() bool`

HasBindingUsername returns a boolean if a field has been set.

### GetBindingPassword

`func (o *UserSourceCreateUserSourceConfig) GetBindingPassword() string`

GetBindingPassword returns the BindingPassword field if non-nil, zero value otherwise.

### GetBindingPasswordOk

`func (o *UserSourceCreateUserSourceConfig) GetBindingPasswordOk() (*string, bool)`

GetBindingPasswordOk returns a tuple with the BindingPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingPassword

`func (o *UserSourceCreateUserSourceConfig) SetBindingPassword(v string)`

SetBindingPassword sets BindingPassword field to given value.

### HasBindingPassword

`func (o *UserSourceCreateUserSourceConfig) HasBindingPassword() bool`

HasBindingPassword returns a boolean if a field has been set.

### GetRequiredGroup

`func (o *UserSourceCreateUserSourceConfig) GetRequiredGroup() string`

GetRequiredGroup returns the RequiredGroup field if non-nil, zero value otherwise.

### GetRequiredGroupOk

`func (o *UserSourceCreateUserSourceConfig) GetRequiredGroupOk() (*string, bool)`

GetRequiredGroupOk returns a tuple with the RequiredGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredGroup

`func (o *UserSourceCreateUserSourceConfig) SetRequiredGroup(v string)`

SetRequiredGroup sets RequiredGroup field to given value.

### HasRequiredGroup

`func (o *UserSourceCreateUserSourceConfig) HasRequiredGroup() bool`

HasRequiredGroup returns a boolean if a field has been set.

### GetOrganizationId

`func (o *UserSourceCreateUserSourceConfig) GetOrganizationId() bool`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *UserSourceCreateUserSourceConfig) GetOrganizationIdOk() (*bool, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *UserSourceCreateUserSourceConfig) SetOrganizationId(v bool)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *UserSourceCreateUserSourceConfig) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetRequiredRole

`func (o *UserSourceCreateUserSourceConfig) GetRequiredRole() string`

GetRequiredRole returns the RequiredRole field if non-nil, zero value otherwise.

### GetRequiredRoleOk

`func (o *UserSourceCreateUserSourceConfig) GetRequiredRoleOk() (*string, bool)`

GetRequiredRoleOk returns a tuple with the RequiredRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredRole

`func (o *UserSourceCreateUserSourceConfig) SetRequiredRole(v string)`

SetRequiredRole sets RequiredRole field to given value.

### HasRequiredRole

`func (o *UserSourceCreateUserSourceConfig) HasRequiredRole() bool`

HasRequiredRole returns a boolean if a field has been set.

### GetDomain

`func (o *UserSourceCreateUserSourceConfig) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *UserSourceCreateUserSourceConfig) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *UserSourceCreateUserSourceConfig) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *UserSourceCreateUserSourceConfig) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetUseSSL

`func (o *UserSourceCreateUserSourceConfig) GetUseSSL() string`

GetUseSSL returns the UseSSL field if non-nil, zero value otherwise.

### GetUseSSLOk

`func (o *UserSourceCreateUserSourceConfig) GetUseSSLOk() (*string, bool)`

GetUseSSLOk returns a tuple with the UseSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSSL

`func (o *UserSourceCreateUserSourceConfig) SetUseSSL(v string)`

SetUseSSL sets UseSSL field to given value.

### HasUseSSL

`func (o *UserSourceCreateUserSourceConfig) HasUseSSL() bool`

HasUseSSL returns a boolean if a field has been set.

### GetSearchMemberGroups

`func (o *UserSourceCreateUserSourceConfig) GetSearchMemberGroups() bool`

GetSearchMemberGroups returns the SearchMemberGroups field if non-nil, zero value otherwise.

### GetSearchMemberGroupsOk

`func (o *UserSourceCreateUserSourceConfig) GetSearchMemberGroupsOk() (*bool, bool)`

GetSearchMemberGroupsOk returns a tuple with the SearchMemberGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchMemberGroups

`func (o *UserSourceCreateUserSourceConfig) SetSearchMemberGroups(v bool)`

SetSearchMemberGroups sets SearchMemberGroups field to given value.

### HasSearchMemberGroups

`func (o *UserSourceCreateUserSourceConfig) HasSearchMemberGroups() bool`

HasSearchMemberGroups returns a boolean if a field has been set.

### GetAdministratorAPIToken

`func (o *UserSourceCreateUserSourceConfig) GetAdministratorAPIToken() string`

GetAdministratorAPIToken returns the AdministratorAPIToken field if non-nil, zero value otherwise.

### GetAdministratorAPITokenOk

`func (o *UserSourceCreateUserSourceConfig) GetAdministratorAPITokenOk() (*string, bool)`

GetAdministratorAPITokenOk returns a tuple with the AdministratorAPIToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministratorAPIToken

`func (o *UserSourceCreateUserSourceConfig) SetAdministratorAPIToken(v string)`

SetAdministratorAPIToken sets AdministratorAPIToken field to given value.

### HasAdministratorAPIToken

`func (o *UserSourceCreateUserSourceConfig) HasAdministratorAPIToken() bool`

HasAdministratorAPIToken returns a boolean if a field has been set.

### GetSubdomain

`func (o *UserSourceCreateUserSourceConfig) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *UserSourceCreateUserSourceConfig) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *UserSourceCreateUserSourceConfig) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *UserSourceCreateUserSourceConfig) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetRegion

`func (o *UserSourceCreateUserSourceConfig) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *UserSourceCreateUserSourceConfig) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *UserSourceCreateUserSourceConfig) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *UserSourceCreateUserSourceConfig) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetClientSecret

`func (o *UserSourceCreateUserSourceConfig) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *UserSourceCreateUserSourceConfig) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *UserSourceCreateUserSourceConfig) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *UserSourceCreateUserSourceConfig) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetClientId

`func (o *UserSourceCreateUserSourceConfig) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *UserSourceCreateUserSourceConfig) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *UserSourceCreateUserSourceConfig) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *UserSourceCreateUserSourceConfig) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetDoNotIncludeSAMLRequest

`func (o *UserSourceCreateUserSourceConfig) GetDoNotIncludeSAMLRequest() bool`

GetDoNotIncludeSAMLRequest returns the DoNotIncludeSAMLRequest field if non-nil, zero value otherwise.

### GetDoNotIncludeSAMLRequestOk

`func (o *UserSourceCreateUserSourceConfig) GetDoNotIncludeSAMLRequestOk() (*bool, bool)`

GetDoNotIncludeSAMLRequestOk returns a tuple with the DoNotIncludeSAMLRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotIncludeSAMLRequest

`func (o *UserSourceCreateUserSourceConfig) SetDoNotIncludeSAMLRequest(v bool)`

SetDoNotIncludeSAMLRequest sets DoNotIncludeSAMLRequest field to given value.

### HasDoNotIncludeSAMLRequest

`func (o *UserSourceCreateUserSourceConfig) HasDoNotIncludeSAMLRequest() bool`

HasDoNotIncludeSAMLRequest returns a boolean if a field has been set.

### GetLogoutUrl

`func (o *UserSourceCreateUserSourceConfig) GetLogoutUrl() string`

GetLogoutUrl returns the LogoutUrl field if non-nil, zero value otherwise.

### GetLogoutUrlOk

`func (o *UserSourceCreateUserSourceConfig) GetLogoutUrlOk() (*string, bool)`

GetLogoutUrlOk returns a tuple with the LogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUrl

`func (o *UserSourceCreateUserSourceConfig) SetLogoutUrl(v string)`

SetLogoutUrl sets LogoutUrl field to given value.

### HasLogoutUrl

`func (o *UserSourceCreateUserSourceConfig) HasLogoutUrl() bool`

HasLogoutUrl returns a boolean if a field has been set.

### GetSAMLSignatureMode

`func (o *UserSourceCreateUserSourceConfig) GetSAMLSignatureMode() string`

GetSAMLSignatureMode returns the SAMLSignatureMode field if non-nil, zero value otherwise.

### GetSAMLSignatureModeOk

`func (o *UserSourceCreateUserSourceConfig) GetSAMLSignatureModeOk() (*string, bool)`

GetSAMLSignatureModeOk returns a tuple with the SAMLSignatureMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAMLSignatureMode

`func (o *UserSourceCreateUserSourceConfig) SetSAMLSignatureMode(v string)`

SetSAMLSignatureMode sets SAMLSignatureMode field to given value.

### HasSAMLSignatureMode

`func (o *UserSourceCreateUserSourceConfig) HasSAMLSignatureMode() bool`

HasSAMLSignatureMode returns a boolean if a field has been set.

### GetRequest509Certificate

`func (o *UserSourceCreateUserSourceConfig) GetRequest509Certificate() string`

GetRequest509Certificate returns the Request509Certificate field if non-nil, zero value otherwise.

### GetRequest509CertificateOk

`func (o *UserSourceCreateUserSourceConfig) GetRequest509CertificateOk() (*string, bool)`

GetRequest509CertificateOk returns a tuple with the Request509Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest509Certificate

`func (o *UserSourceCreateUserSourceConfig) SetRequest509Certificate(v string)`

SetRequest509Certificate sets Request509Certificate field to given value.

### HasRequest509Certificate

`func (o *UserSourceCreateUserSourceConfig) HasRequest509Certificate() bool`

HasRequest509Certificate returns a boolean if a field has been set.

### GetRequestPrivateKey

`func (o *UserSourceCreateUserSourceConfig) GetRequestPrivateKey() string`

GetRequestPrivateKey returns the RequestPrivateKey field if non-nil, zero value otherwise.

### GetRequestPrivateKeyOk

`func (o *UserSourceCreateUserSourceConfig) GetRequestPrivateKeyOk() (*string, bool)`

GetRequestPrivateKeyOk returns a tuple with the RequestPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestPrivateKey

`func (o *UserSourceCreateUserSourceConfig) SetRequestPrivateKey(v string)`

SetRequestPrivateKey sets RequestPrivateKey field to given value.

### HasRequestPrivateKey

`func (o *UserSourceCreateUserSourceConfig) HasRequestPrivateKey() bool`

HasRequestPrivateKey returns a boolean if a field has been set.

### GetDoNotValidateSignature

`func (o *UserSourceCreateUserSourceConfig) GetDoNotValidateSignature() bool`

GetDoNotValidateSignature returns the DoNotValidateSignature field if non-nil, zero value otherwise.

### GetDoNotValidateSignatureOk

`func (o *UserSourceCreateUserSourceConfig) GetDoNotValidateSignatureOk() (*bool, bool)`

GetDoNotValidateSignatureOk returns a tuple with the DoNotValidateSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotValidateSignature

`func (o *UserSourceCreateUserSourceConfig) SetDoNotValidateSignature(v bool)`

SetDoNotValidateSignature sets DoNotValidateSignature field to given value.

### HasDoNotValidateSignature

`func (o *UserSourceCreateUserSourceConfig) HasDoNotValidateSignature() bool`

HasDoNotValidateSignature returns a boolean if a field has been set.

### GetPublicKey

`func (o *UserSourceCreateUserSourceConfig) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *UserSourceCreateUserSourceConfig) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *UserSourceCreateUserSourceConfig) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *UserSourceCreateUserSourceConfig) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetPrivateKey

`func (o *UserSourceCreateUserSourceConfig) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *UserSourceCreateUserSourceConfig) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *UserSourceCreateUserSourceConfig) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *UserSourceCreateUserSourceConfig) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetGivenNameAttribute

`func (o *UserSourceCreateUserSourceConfig) GetGivenNameAttribute() string`

GetGivenNameAttribute returns the GivenNameAttribute field if non-nil, zero value otherwise.

### GetGivenNameAttributeOk

`func (o *UserSourceCreateUserSourceConfig) GetGivenNameAttributeOk() (*string, bool)`

GetGivenNameAttributeOk returns a tuple with the GivenNameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenNameAttribute

`func (o *UserSourceCreateUserSourceConfig) SetGivenNameAttribute(v string)`

SetGivenNameAttribute sets GivenNameAttribute field to given value.

### HasGivenNameAttribute

`func (o *UserSourceCreateUserSourceConfig) HasGivenNameAttribute() bool`

HasGivenNameAttribute returns a boolean if a field has been set.

### GetSurnameAttribute

`func (o *UserSourceCreateUserSourceConfig) GetSurnameAttribute() string`

GetSurnameAttribute returns the SurnameAttribute field if non-nil, zero value otherwise.

### GetSurnameAttributeOk

`func (o *UserSourceCreateUserSourceConfig) GetSurnameAttributeOk() (*string, bool)`

GetSurnameAttributeOk returns a tuple with the SurnameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurnameAttribute

`func (o *UserSourceCreateUserSourceConfig) SetSurnameAttribute(v string)`

SetSurnameAttribute sets SurnameAttribute field to given value.

### HasSurnameAttribute

`func (o *UserSourceCreateUserSourceConfig) HasSurnameAttribute() bool`

HasSurnameAttribute returns a boolean if a field has been set.

### GetRoleAttributeName

`func (o *UserSourceCreateUserSourceConfig) GetRoleAttributeName() string`

GetRoleAttributeName returns the RoleAttributeName field if non-nil, zero value otherwise.

### GetRoleAttributeNameOk

`func (o *UserSourceCreateUserSourceConfig) GetRoleAttributeNameOk() (*string, bool)`

GetRoleAttributeNameOk returns a tuple with the RoleAttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleAttributeName

`func (o *UserSourceCreateUserSourceConfig) SetRoleAttributeName(v string)`

SetRoleAttributeName sets RoleAttributeName field to given value.

### HasRoleAttributeName

`func (o *UserSourceCreateUserSourceConfig) HasRoleAttributeName() bool`

HasRoleAttributeName returns a boolean if a field has been set.

### GetRequiredAttributeValue

`func (o *UserSourceCreateUserSourceConfig) GetRequiredAttributeValue() string`

GetRequiredAttributeValue returns the RequiredAttributeValue field if non-nil, zero value otherwise.

### GetRequiredAttributeValueOk

`func (o *UserSourceCreateUserSourceConfig) GetRequiredAttributeValueOk() (*string, bool)`

GetRequiredAttributeValueOk returns a tuple with the RequiredAttributeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredAttributeValue

`func (o *UserSourceCreateUserSourceConfig) SetRequiredAttributeValue(v string)`

SetRequiredAttributeValue sets RequiredAttributeValue field to given value.

### HasRequiredAttributeValue

`func (o *UserSourceCreateUserSourceConfig) HasRequiredAttributeValue() bool`

HasRequiredAttributeValue returns a boolean if a field has been set.

### GetLoginUrl

`func (o *UserSourceCreateUserSourceConfig) GetLoginUrl() string`

GetLoginUrl returns the LoginUrl field if non-nil, zero value otherwise.

### GetLoginUrlOk

`func (o *UserSourceCreateUserSourceConfig) GetLoginUrlOk() (*string, bool)`

GetLoginUrlOk returns a tuple with the LoginUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginUrl

`func (o *UserSourceCreateUserSourceConfig) SetLoginUrl(v string)`

SetLoginUrl sets LoginUrl field to given value.

### HasLoginUrl

`func (o *UserSourceCreateUserSourceConfig) HasLoginUrl() bool`

HasLoginUrl returns a boolean if a field has been set.

### GetLogout

`func (o *UserSourceCreateUserSourceConfig) GetLogout() string`

GetLogout returns the Logout field if non-nil, zero value otherwise.

### GetLogoutOk

`func (o *UserSourceCreateUserSourceConfig) GetLogoutOk() (*string, bool)`

GetLogoutOk returns a tuple with the Logout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogout

`func (o *UserSourceCreateUserSourceConfig) SetLogout(v string)`

SetLogout sets Logout field to given value.

### HasLogout

`func (o *UserSourceCreateUserSourceConfig) HasLogout() bool`

HasLogout returns a boolean if a field has been set.

### GetEncryptionAlgo

`func (o *UserSourceCreateUserSourceConfig) GetEncryptionAlgo() string`

GetEncryptionAlgo returns the EncryptionAlgo field if non-nil, zero value otherwise.

### GetEncryptionAlgoOk

`func (o *UserSourceCreateUserSourceConfig) GetEncryptionAlgoOk() (*string, bool)`

GetEncryptionAlgoOk returns a tuple with the EncryptionAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAlgo

`func (o *UserSourceCreateUserSourceConfig) SetEncryptionAlgo(v string)`

SetEncryptionAlgo sets EncryptionAlgo field to given value.

### HasEncryptionAlgo

`func (o *UserSourceCreateUserSourceConfig) HasEncryptionAlgo() bool`

HasEncryptionAlgo returns a boolean if a field has been set.

### GetEncryptionKey

`func (o *UserSourceCreateUserSourceConfig) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *UserSourceCreateUserSourceConfig) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *UserSourceCreateUserSourceConfig) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.

### HasEncryptionKey

`func (o *UserSourceCreateUserSourceConfig) HasEncryptionKey() bool`

HasEncryptionKey returns a boolean if a field has been set.

### GetEndpoint

`func (o *UserSourceCreateUserSourceConfig) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *UserSourceCreateUserSourceConfig) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *UserSourceCreateUserSourceConfig) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *UserSourceCreateUserSourceConfig) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetApiStyle

`func (o *UserSourceCreateUserSourceConfig) GetApiStyle() string`

GetApiStyle returns the ApiStyle field if non-nil, zero value otherwise.

### GetApiStyleOk

`func (o *UserSourceCreateUserSourceConfig) GetApiStyleOk() (*string, bool)`

GetApiStyleOk returns a tuple with the ApiStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiStyle

`func (o *UserSourceCreateUserSourceConfig) SetApiStyle(v string)`

SetApiStyle sets ApiStyle field to given value.

### HasApiStyle

`func (o *UserSourceCreateUserSourceConfig) HasApiStyle() bool`

HasApiStyle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


