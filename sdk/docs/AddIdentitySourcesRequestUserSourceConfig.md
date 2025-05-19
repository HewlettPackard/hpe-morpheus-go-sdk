# AddIdentitySourcesRequestUserSourceConfig

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

### NewAddIdentitySourcesRequestUserSourceConfig

`func NewAddIdentitySourcesRequestUserSourceConfig() *AddIdentitySourcesRequestUserSourceConfig`

NewAddIdentitySourcesRequestUserSourceConfig instantiates a new AddIdentitySourcesRequestUserSourceConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddIdentitySourcesRequestUserSourceConfigWithDefaults

`func NewAddIdentitySourcesRequestUserSourceConfigWithDefaults() *AddIdentitySourcesRequestUserSourceConfig`

NewAddIdentitySourcesRequestUserSourceConfigWithDefaults instantiates a new AddIdentitySourcesRequestUserSourceConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *AddIdentitySourcesRequestUserSourceConfig) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *AddIdentitySourcesRequestUserSourceConfig) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetBindingUsername

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetBindingUsername() string`

GetBindingUsername returns the BindingUsername field if non-nil, zero value otherwise.

### GetBindingUsernameOk

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetBindingUsernameOk() (*string, bool)`

GetBindingUsernameOk returns a tuple with the BindingUsername field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingUsername

`func (o *AddIdentitySourcesRequestUserSourceConfig) SetBindingUsername(v string)`

SetBindingUsername sets BindingUsername field to given value.

### HasBindingUsername

`func (o *AddIdentitySourcesRequestUserSourceConfig) HasBindingUsername() bool`

HasBindingUsername returns a boolean if a field has been set.

### GetBindingPassword

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetBindingPassword() string`

GetBindingPassword returns the BindingPassword field if non-nil, zero value otherwise.

### GetBindingPasswordOk

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetBindingPasswordOk() (*string, bool)`

GetBindingPasswordOk returns a tuple with the BindingPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBindingPassword

`func (o *AddIdentitySourcesRequestUserSourceConfig) SetBindingPassword(v string)`

SetBindingPassword sets BindingPassword field to given value.

### HasBindingPassword

`func (o *AddIdentitySourcesRequestUserSourceConfig) HasBindingPassword() bool`

HasBindingPassword returns a boolean if a field has been set.

### GetRequiredGroup

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetRequiredGroup() string`

GetRequiredGroup returns the RequiredGroup field if non-nil, zero value otherwise.

### GetRequiredGroupOk

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetRequiredGroupOk() (*string, bool)`

GetRequiredGroupOk returns a tuple with the RequiredGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredGroup

`func (o *AddIdentitySourcesRequestUserSourceConfig) SetRequiredGroup(v string)`

SetRequiredGroup sets RequiredGroup field to given value.

### HasRequiredGroup

`func (o *AddIdentitySourcesRequestUserSourceConfig) HasRequiredGroup() bool`

HasRequiredGroup returns a boolean if a field has been set.

### GetOrganizationId

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetOrganizationId() bool`

GetOrganizationId returns the OrganizationId field if non-nil, zero value otherwise.

### GetOrganizationIdOk

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetOrganizationIdOk() (*bool, bool)`

GetOrganizationIdOk returns a tuple with the OrganizationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganizationId

`func (o *AddIdentitySourcesRequestUserSourceConfig) SetOrganizationId(v bool)`

SetOrganizationId sets OrganizationId field to given value.

### HasOrganizationId

`func (o *AddIdentitySourcesRequestUserSourceConfig) HasOrganizationId() bool`

HasOrganizationId returns a boolean if a field has been set.

### GetRequiredRole

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetRequiredRole() string`

GetRequiredRole returns the RequiredRole field if non-nil, zero value otherwise.

### GetRequiredRoleOk

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetRequiredRoleOk() (*string, bool)`

GetRequiredRoleOk returns a tuple with the RequiredRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredRole

`func (o *AddIdentitySourcesRequestUserSourceConfig) SetRequiredRole(v string)`

SetRequiredRole sets RequiredRole field to given value.

### HasRequiredRole

`func (o *AddIdentitySourcesRequestUserSourceConfig) HasRequiredRole() bool`

HasRequiredRole returns a boolean if a field has been set.

### GetDomain

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetDomain() string`

GetDomain returns the Domain field if non-nil, zero value otherwise.

### GetDomainOk

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetDomainOk() (*string, bool)`

GetDomainOk returns a tuple with the Domain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomain

`func (o *AddIdentitySourcesRequestUserSourceConfig) SetDomain(v string)`

SetDomain sets Domain field to given value.

### HasDomain

`func (o *AddIdentitySourcesRequestUserSourceConfig) HasDomain() bool`

HasDomain returns a boolean if a field has been set.

### GetUseSSL

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetUseSSL() string`

GetUseSSL returns the UseSSL field if non-nil, zero value otherwise.

### GetUseSSLOk

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetUseSSLOk() (*string, bool)`

GetUseSSLOk returns a tuple with the UseSSL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseSSL

`func (o *AddIdentitySourcesRequestUserSourceConfig) SetUseSSL(v string)`

SetUseSSL sets UseSSL field to given value.

### HasUseSSL

`func (o *AddIdentitySourcesRequestUserSourceConfig) HasUseSSL() bool`

HasUseSSL returns a boolean if a field has been set.

### GetSearchMemberGroups

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetSearchMemberGroups() bool`

GetSearchMemberGroups returns the SearchMemberGroups field if non-nil, zero value otherwise.

### GetSearchMemberGroupsOk

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetSearchMemberGroupsOk() (*bool, bool)`

GetSearchMemberGroupsOk returns a tuple with the SearchMemberGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSearchMemberGroups

`func (o *AddIdentitySourcesRequestUserSourceConfig) SetSearchMemberGroups(v bool)`

SetSearchMemberGroups sets SearchMemberGroups field to given value.

### HasSearchMemberGroups

`func (o *AddIdentitySourcesRequestUserSourceConfig) HasSearchMemberGroups() bool`

HasSearchMemberGroups returns a boolean if a field has been set.

### GetAdministratorAPIToken

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetAdministratorAPIToken() string`

GetAdministratorAPIToken returns the AdministratorAPIToken field if non-nil, zero value otherwise.

### GetAdministratorAPITokenOk

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetAdministratorAPITokenOk() (*string, bool)`

GetAdministratorAPITokenOk returns a tuple with the AdministratorAPIToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdministratorAPIToken

`func (o *AddIdentitySourcesRequestUserSourceConfig) SetAdministratorAPIToken(v string)`

SetAdministratorAPIToken sets AdministratorAPIToken field to given value.

### HasAdministratorAPIToken

`func (o *AddIdentitySourcesRequestUserSourceConfig) HasAdministratorAPIToken() bool`

HasAdministratorAPIToken returns a boolean if a field has been set.

### GetSubdomain

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetSubdomain() string`

GetSubdomain returns the Subdomain field if non-nil, zero value otherwise.

### GetSubdomainOk

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetSubdomainOk() (*string, bool)`

GetSubdomainOk returns a tuple with the Subdomain field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubdomain

`func (o *AddIdentitySourcesRequestUserSourceConfig) SetSubdomain(v string)`

SetSubdomain sets Subdomain field to given value.

### HasSubdomain

`func (o *AddIdentitySourcesRequestUserSourceConfig) HasSubdomain() bool`

HasSubdomain returns a boolean if a field has been set.

### GetRegion

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AddIdentitySourcesRequestUserSourceConfig) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *AddIdentitySourcesRequestUserSourceConfig) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetClientSecret

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetClientSecret() string`

GetClientSecret returns the ClientSecret field if non-nil, zero value otherwise.

### GetClientSecretOk

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetClientSecretOk() (*string, bool)`

GetClientSecretOk returns a tuple with the ClientSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientSecret

`func (o *AddIdentitySourcesRequestUserSourceConfig) SetClientSecret(v string)`

SetClientSecret sets ClientSecret field to given value.

### HasClientSecret

`func (o *AddIdentitySourcesRequestUserSourceConfig) HasClientSecret() bool`

HasClientSecret returns a boolean if a field has been set.

### GetClientId

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AddIdentitySourcesRequestUserSourceConfig) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *AddIdentitySourcesRequestUserSourceConfig) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetDoNotIncludeSAMLRequest

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetDoNotIncludeSAMLRequest() bool`

GetDoNotIncludeSAMLRequest returns the DoNotIncludeSAMLRequest field if non-nil, zero value otherwise.

### GetDoNotIncludeSAMLRequestOk

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetDoNotIncludeSAMLRequestOk() (*bool, bool)`

GetDoNotIncludeSAMLRequestOk returns a tuple with the DoNotIncludeSAMLRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotIncludeSAMLRequest

`func (o *AddIdentitySourcesRequestUserSourceConfig) SetDoNotIncludeSAMLRequest(v bool)`

SetDoNotIncludeSAMLRequest sets DoNotIncludeSAMLRequest field to given value.

### HasDoNotIncludeSAMLRequest

`func (o *AddIdentitySourcesRequestUserSourceConfig) HasDoNotIncludeSAMLRequest() bool`

HasDoNotIncludeSAMLRequest returns a boolean if a field has been set.

### GetLogoutUrl

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetLogoutUrl() string`

GetLogoutUrl returns the LogoutUrl field if non-nil, zero value otherwise.

### GetLogoutUrlOk

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetLogoutUrlOk() (*string, bool)`

GetLogoutUrlOk returns a tuple with the LogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUrl

`func (o *AddIdentitySourcesRequestUserSourceConfig) SetLogoutUrl(v string)`

SetLogoutUrl sets LogoutUrl field to given value.

### HasLogoutUrl

`func (o *AddIdentitySourcesRequestUserSourceConfig) HasLogoutUrl() bool`

HasLogoutUrl returns a boolean if a field has been set.

### GetSAMLSignatureMode

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetSAMLSignatureMode() string`

GetSAMLSignatureMode returns the SAMLSignatureMode field if non-nil, zero value otherwise.

### GetSAMLSignatureModeOk

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetSAMLSignatureModeOk() (*string, bool)`

GetSAMLSignatureModeOk returns a tuple with the SAMLSignatureMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAMLSignatureMode

`func (o *AddIdentitySourcesRequestUserSourceConfig) SetSAMLSignatureMode(v string)`

SetSAMLSignatureMode sets SAMLSignatureMode field to given value.

### HasSAMLSignatureMode

`func (o *AddIdentitySourcesRequestUserSourceConfig) HasSAMLSignatureMode() bool`

HasSAMLSignatureMode returns a boolean if a field has been set.

### GetRequest509Certificate

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetRequest509Certificate() string`

GetRequest509Certificate returns the Request509Certificate field if non-nil, zero value otherwise.

### GetRequest509CertificateOk

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetRequest509CertificateOk() (*string, bool)`

GetRequest509CertificateOk returns a tuple with the Request509Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest509Certificate

`func (o *AddIdentitySourcesRequestUserSourceConfig) SetRequest509Certificate(v string)`

SetRequest509Certificate sets Request509Certificate field to given value.

### HasRequest509Certificate

`func (o *AddIdentitySourcesRequestUserSourceConfig) HasRequest509Certificate() bool`

HasRequest509Certificate returns a boolean if a field has been set.

### GetRequestPrivateKey

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetRequestPrivateKey() string`

GetRequestPrivateKey returns the RequestPrivateKey field if non-nil, zero value otherwise.

### GetRequestPrivateKeyOk

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetRequestPrivateKeyOk() (*string, bool)`

GetRequestPrivateKeyOk returns a tuple with the RequestPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestPrivateKey

`func (o *AddIdentitySourcesRequestUserSourceConfig) SetRequestPrivateKey(v string)`

SetRequestPrivateKey sets RequestPrivateKey field to given value.

### HasRequestPrivateKey

`func (o *AddIdentitySourcesRequestUserSourceConfig) HasRequestPrivateKey() bool`

HasRequestPrivateKey returns a boolean if a field has been set.

### GetDoNotValidateSignature

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetDoNotValidateSignature() bool`

GetDoNotValidateSignature returns the DoNotValidateSignature field if non-nil, zero value otherwise.

### GetDoNotValidateSignatureOk

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetDoNotValidateSignatureOk() (*bool, bool)`

GetDoNotValidateSignatureOk returns a tuple with the DoNotValidateSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotValidateSignature

`func (o *AddIdentitySourcesRequestUserSourceConfig) SetDoNotValidateSignature(v bool)`

SetDoNotValidateSignature sets DoNotValidateSignature field to given value.

### HasDoNotValidateSignature

`func (o *AddIdentitySourcesRequestUserSourceConfig) HasDoNotValidateSignature() bool`

HasDoNotValidateSignature returns a boolean if a field has been set.

### GetPublicKey

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *AddIdentitySourcesRequestUserSourceConfig) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *AddIdentitySourcesRequestUserSourceConfig) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetPrivateKey

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *AddIdentitySourcesRequestUserSourceConfig) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *AddIdentitySourcesRequestUserSourceConfig) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetGivenNameAttribute

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetGivenNameAttribute() string`

GetGivenNameAttribute returns the GivenNameAttribute field if non-nil, zero value otherwise.

### GetGivenNameAttributeOk

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetGivenNameAttributeOk() (*string, bool)`

GetGivenNameAttributeOk returns a tuple with the GivenNameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenNameAttribute

`func (o *AddIdentitySourcesRequestUserSourceConfig) SetGivenNameAttribute(v string)`

SetGivenNameAttribute sets GivenNameAttribute field to given value.

### HasGivenNameAttribute

`func (o *AddIdentitySourcesRequestUserSourceConfig) HasGivenNameAttribute() bool`

HasGivenNameAttribute returns a boolean if a field has been set.

### GetSurnameAttribute

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetSurnameAttribute() string`

GetSurnameAttribute returns the SurnameAttribute field if non-nil, zero value otherwise.

### GetSurnameAttributeOk

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetSurnameAttributeOk() (*string, bool)`

GetSurnameAttributeOk returns a tuple with the SurnameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurnameAttribute

`func (o *AddIdentitySourcesRequestUserSourceConfig) SetSurnameAttribute(v string)`

SetSurnameAttribute sets SurnameAttribute field to given value.

### HasSurnameAttribute

`func (o *AddIdentitySourcesRequestUserSourceConfig) HasSurnameAttribute() bool`

HasSurnameAttribute returns a boolean if a field has been set.

### GetRoleAttributeName

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetRoleAttributeName() string`

GetRoleAttributeName returns the RoleAttributeName field if non-nil, zero value otherwise.

### GetRoleAttributeNameOk

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetRoleAttributeNameOk() (*string, bool)`

GetRoleAttributeNameOk returns a tuple with the RoleAttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleAttributeName

`func (o *AddIdentitySourcesRequestUserSourceConfig) SetRoleAttributeName(v string)`

SetRoleAttributeName sets RoleAttributeName field to given value.

### HasRoleAttributeName

`func (o *AddIdentitySourcesRequestUserSourceConfig) HasRoleAttributeName() bool`

HasRoleAttributeName returns a boolean if a field has been set.

### GetRequiredAttributeValue

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetRequiredAttributeValue() string`

GetRequiredAttributeValue returns the RequiredAttributeValue field if non-nil, zero value otherwise.

### GetRequiredAttributeValueOk

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetRequiredAttributeValueOk() (*string, bool)`

GetRequiredAttributeValueOk returns a tuple with the RequiredAttributeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredAttributeValue

`func (o *AddIdentitySourcesRequestUserSourceConfig) SetRequiredAttributeValue(v string)`

SetRequiredAttributeValue sets RequiredAttributeValue field to given value.

### HasRequiredAttributeValue

`func (o *AddIdentitySourcesRequestUserSourceConfig) HasRequiredAttributeValue() bool`

HasRequiredAttributeValue returns a boolean if a field has been set.

### GetLoginUrl

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetLoginUrl() string`

GetLoginUrl returns the LoginUrl field if non-nil, zero value otherwise.

### GetLoginUrlOk

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetLoginUrlOk() (*string, bool)`

GetLoginUrlOk returns a tuple with the LoginUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginUrl

`func (o *AddIdentitySourcesRequestUserSourceConfig) SetLoginUrl(v string)`

SetLoginUrl sets LoginUrl field to given value.

### HasLoginUrl

`func (o *AddIdentitySourcesRequestUserSourceConfig) HasLoginUrl() bool`

HasLoginUrl returns a boolean if a field has been set.

### GetLogout

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetLogout() string`

GetLogout returns the Logout field if non-nil, zero value otherwise.

### GetLogoutOk

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetLogoutOk() (*string, bool)`

GetLogoutOk returns a tuple with the Logout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogout

`func (o *AddIdentitySourcesRequestUserSourceConfig) SetLogout(v string)`

SetLogout sets Logout field to given value.

### HasLogout

`func (o *AddIdentitySourcesRequestUserSourceConfig) HasLogout() bool`

HasLogout returns a boolean if a field has been set.

### GetEncryptionAlgo

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetEncryptionAlgo() string`

GetEncryptionAlgo returns the EncryptionAlgo field if non-nil, zero value otherwise.

### GetEncryptionAlgoOk

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetEncryptionAlgoOk() (*string, bool)`

GetEncryptionAlgoOk returns a tuple with the EncryptionAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAlgo

`func (o *AddIdentitySourcesRequestUserSourceConfig) SetEncryptionAlgo(v string)`

SetEncryptionAlgo sets EncryptionAlgo field to given value.

### HasEncryptionAlgo

`func (o *AddIdentitySourcesRequestUserSourceConfig) HasEncryptionAlgo() bool`

HasEncryptionAlgo returns a boolean if a field has been set.

### GetEncryptionKey

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *AddIdentitySourcesRequestUserSourceConfig) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.

### HasEncryptionKey

`func (o *AddIdentitySourcesRequestUserSourceConfig) HasEncryptionKey() bool`

HasEncryptionKey returns a boolean if a field has been set.

### GetEndpoint

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *AddIdentitySourcesRequestUserSourceConfig) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *AddIdentitySourcesRequestUserSourceConfig) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetApiStyle

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetApiStyle() string`

GetApiStyle returns the ApiStyle field if non-nil, zero value otherwise.

### GetApiStyleOk

`func (o *AddIdentitySourcesRequestUserSourceConfig) GetApiStyleOk() (*string, bool)`

GetApiStyleOk returns a tuple with the ApiStyle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiStyle

`func (o *AddIdentitySourcesRequestUserSourceConfig) SetApiStyle(v string)`

SetApiStyle sets ApiStyle field to given value.

### HasApiStyle

`func (o *AddIdentitySourcesRequestUserSourceConfig) HasApiStyle() bool`

HasApiStyle returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


