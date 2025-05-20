# UserSourceCreateSaml

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** | Login Redirect URL | [optional] 
**DoNotIncludeSAMLRequest** | Pointer to **bool** | Exclude SAML Request Parameter | [optional] [default to false]
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

## Methods

### NewUserSourceCreateSaml

`func NewUserSourceCreateSaml() *UserSourceCreateSaml`

NewUserSourceCreateSaml instantiates a new UserSourceCreateSaml object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSourceCreateSamlWithDefaults

`func NewUserSourceCreateSamlWithDefaults() *UserSourceCreateSaml`

NewUserSourceCreateSamlWithDefaults instantiates a new UserSourceCreateSaml object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *UserSourceCreateSaml) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UserSourceCreateSaml) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UserSourceCreateSaml) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UserSourceCreateSaml) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDoNotIncludeSAMLRequest

`func (o *UserSourceCreateSaml) GetDoNotIncludeSAMLRequest() bool`

GetDoNotIncludeSAMLRequest returns the DoNotIncludeSAMLRequest field if non-nil, zero value otherwise.

### GetDoNotIncludeSAMLRequestOk

`func (o *UserSourceCreateSaml) GetDoNotIncludeSAMLRequestOk() (*bool, bool)`

GetDoNotIncludeSAMLRequestOk returns a tuple with the DoNotIncludeSAMLRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotIncludeSAMLRequest

`func (o *UserSourceCreateSaml) SetDoNotIncludeSAMLRequest(v bool)`

SetDoNotIncludeSAMLRequest sets DoNotIncludeSAMLRequest field to given value.

### HasDoNotIncludeSAMLRequest

`func (o *UserSourceCreateSaml) HasDoNotIncludeSAMLRequest() bool`

HasDoNotIncludeSAMLRequest returns a boolean if a field has been set.

### GetLogoutUrl

`func (o *UserSourceCreateSaml) GetLogoutUrl() string`

GetLogoutUrl returns the LogoutUrl field if non-nil, zero value otherwise.

### GetLogoutUrlOk

`func (o *UserSourceCreateSaml) GetLogoutUrlOk() (*string, bool)`

GetLogoutUrlOk returns a tuple with the LogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUrl

`func (o *UserSourceCreateSaml) SetLogoutUrl(v string)`

SetLogoutUrl sets LogoutUrl field to given value.

### HasLogoutUrl

`func (o *UserSourceCreateSaml) HasLogoutUrl() bool`

HasLogoutUrl returns a boolean if a field has been set.

### GetSAMLSignatureMode

`func (o *UserSourceCreateSaml) GetSAMLSignatureMode() string`

GetSAMLSignatureMode returns the SAMLSignatureMode field if non-nil, zero value otherwise.

### GetSAMLSignatureModeOk

`func (o *UserSourceCreateSaml) GetSAMLSignatureModeOk() (*string, bool)`

GetSAMLSignatureModeOk returns a tuple with the SAMLSignatureMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAMLSignatureMode

`func (o *UserSourceCreateSaml) SetSAMLSignatureMode(v string)`

SetSAMLSignatureMode sets SAMLSignatureMode field to given value.

### HasSAMLSignatureMode

`func (o *UserSourceCreateSaml) HasSAMLSignatureMode() bool`

HasSAMLSignatureMode returns a boolean if a field has been set.

### GetRequest509Certificate

`func (o *UserSourceCreateSaml) GetRequest509Certificate() string`

GetRequest509Certificate returns the Request509Certificate field if non-nil, zero value otherwise.

### GetRequest509CertificateOk

`func (o *UserSourceCreateSaml) GetRequest509CertificateOk() (*string, bool)`

GetRequest509CertificateOk returns a tuple with the Request509Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest509Certificate

`func (o *UserSourceCreateSaml) SetRequest509Certificate(v string)`

SetRequest509Certificate sets Request509Certificate field to given value.

### HasRequest509Certificate

`func (o *UserSourceCreateSaml) HasRequest509Certificate() bool`

HasRequest509Certificate returns a boolean if a field has been set.

### GetRequestPrivateKey

`func (o *UserSourceCreateSaml) GetRequestPrivateKey() string`

GetRequestPrivateKey returns the RequestPrivateKey field if non-nil, zero value otherwise.

### GetRequestPrivateKeyOk

`func (o *UserSourceCreateSaml) GetRequestPrivateKeyOk() (*string, bool)`

GetRequestPrivateKeyOk returns a tuple with the RequestPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestPrivateKey

`func (o *UserSourceCreateSaml) SetRequestPrivateKey(v string)`

SetRequestPrivateKey sets RequestPrivateKey field to given value.

### HasRequestPrivateKey

`func (o *UserSourceCreateSaml) HasRequestPrivateKey() bool`

HasRequestPrivateKey returns a boolean if a field has been set.

### GetDoNotValidateSignature

`func (o *UserSourceCreateSaml) GetDoNotValidateSignature() bool`

GetDoNotValidateSignature returns the DoNotValidateSignature field if non-nil, zero value otherwise.

### GetDoNotValidateSignatureOk

`func (o *UserSourceCreateSaml) GetDoNotValidateSignatureOk() (*bool, bool)`

GetDoNotValidateSignatureOk returns a tuple with the DoNotValidateSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotValidateSignature

`func (o *UserSourceCreateSaml) SetDoNotValidateSignature(v bool)`

SetDoNotValidateSignature sets DoNotValidateSignature field to given value.

### HasDoNotValidateSignature

`func (o *UserSourceCreateSaml) HasDoNotValidateSignature() bool`

HasDoNotValidateSignature returns a boolean if a field has been set.

### GetPublicKey

`func (o *UserSourceCreateSaml) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *UserSourceCreateSaml) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *UserSourceCreateSaml) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *UserSourceCreateSaml) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetPrivateKey

`func (o *UserSourceCreateSaml) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *UserSourceCreateSaml) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *UserSourceCreateSaml) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *UserSourceCreateSaml) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetGivenNameAttribute

`func (o *UserSourceCreateSaml) GetGivenNameAttribute() string`

GetGivenNameAttribute returns the GivenNameAttribute field if non-nil, zero value otherwise.

### GetGivenNameAttributeOk

`func (o *UserSourceCreateSaml) GetGivenNameAttributeOk() (*string, bool)`

GetGivenNameAttributeOk returns a tuple with the GivenNameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenNameAttribute

`func (o *UserSourceCreateSaml) SetGivenNameAttribute(v string)`

SetGivenNameAttribute sets GivenNameAttribute field to given value.

### HasGivenNameAttribute

`func (o *UserSourceCreateSaml) HasGivenNameAttribute() bool`

HasGivenNameAttribute returns a boolean if a field has been set.

### GetSurnameAttribute

`func (o *UserSourceCreateSaml) GetSurnameAttribute() string`

GetSurnameAttribute returns the SurnameAttribute field if non-nil, zero value otherwise.

### GetSurnameAttributeOk

`func (o *UserSourceCreateSaml) GetSurnameAttributeOk() (*string, bool)`

GetSurnameAttributeOk returns a tuple with the SurnameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurnameAttribute

`func (o *UserSourceCreateSaml) SetSurnameAttribute(v string)`

SetSurnameAttribute sets SurnameAttribute field to given value.

### HasSurnameAttribute

`func (o *UserSourceCreateSaml) HasSurnameAttribute() bool`

HasSurnameAttribute returns a boolean if a field has been set.

### GetRoleAttributeName

`func (o *UserSourceCreateSaml) GetRoleAttributeName() string`

GetRoleAttributeName returns the RoleAttributeName field if non-nil, zero value otherwise.

### GetRoleAttributeNameOk

`func (o *UserSourceCreateSaml) GetRoleAttributeNameOk() (*string, bool)`

GetRoleAttributeNameOk returns a tuple with the RoleAttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleAttributeName

`func (o *UserSourceCreateSaml) SetRoleAttributeName(v string)`

SetRoleAttributeName sets RoleAttributeName field to given value.

### HasRoleAttributeName

`func (o *UserSourceCreateSaml) HasRoleAttributeName() bool`

HasRoleAttributeName returns a boolean if a field has been set.

### GetRequiredAttributeValue

`func (o *UserSourceCreateSaml) GetRequiredAttributeValue() string`

GetRequiredAttributeValue returns the RequiredAttributeValue field if non-nil, zero value otherwise.

### GetRequiredAttributeValueOk

`func (o *UserSourceCreateSaml) GetRequiredAttributeValueOk() (*string, bool)`

GetRequiredAttributeValueOk returns a tuple with the RequiredAttributeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredAttributeValue

`func (o *UserSourceCreateSaml) SetRequiredAttributeValue(v string)`

SetRequiredAttributeValue sets RequiredAttributeValue field to given value.

### HasRequiredAttributeValue

`func (o *UserSourceCreateSaml) HasRequiredAttributeValue() bool`

HasRequiredAttributeValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


