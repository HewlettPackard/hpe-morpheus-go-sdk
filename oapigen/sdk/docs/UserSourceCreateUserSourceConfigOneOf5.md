# UserSourceCreateUserSourceConfigOneOf5

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

### NewUserSourceCreateUserSourceConfigOneOf5

`func NewUserSourceCreateUserSourceConfigOneOf5() *UserSourceCreateUserSourceConfigOneOf5`

NewUserSourceCreateUserSourceConfigOneOf5 instantiates a new UserSourceCreateUserSourceConfigOneOf5 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSourceCreateUserSourceConfigOneOf5WithDefaults

`func NewUserSourceCreateUserSourceConfigOneOf5WithDefaults() *UserSourceCreateUserSourceConfigOneOf5`

NewUserSourceCreateUserSourceConfigOneOf5WithDefaults instantiates a new UserSourceCreateUserSourceConfigOneOf5 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *UserSourceCreateUserSourceConfigOneOf5) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *UserSourceCreateUserSourceConfigOneOf5) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *UserSourceCreateUserSourceConfigOneOf5) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *UserSourceCreateUserSourceConfigOneOf5) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetDoNotIncludeSAMLRequest

`func (o *UserSourceCreateUserSourceConfigOneOf5) GetDoNotIncludeSAMLRequest() bool`

GetDoNotIncludeSAMLRequest returns the DoNotIncludeSAMLRequest field if non-nil, zero value otherwise.

### GetDoNotIncludeSAMLRequestOk

`func (o *UserSourceCreateUserSourceConfigOneOf5) GetDoNotIncludeSAMLRequestOk() (*bool, bool)`

GetDoNotIncludeSAMLRequestOk returns a tuple with the DoNotIncludeSAMLRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotIncludeSAMLRequest

`func (o *UserSourceCreateUserSourceConfigOneOf5) SetDoNotIncludeSAMLRequest(v bool)`

SetDoNotIncludeSAMLRequest sets DoNotIncludeSAMLRequest field to given value.

### HasDoNotIncludeSAMLRequest

`func (o *UserSourceCreateUserSourceConfigOneOf5) HasDoNotIncludeSAMLRequest() bool`

HasDoNotIncludeSAMLRequest returns a boolean if a field has been set.

### GetLogoutUrl

`func (o *UserSourceCreateUserSourceConfigOneOf5) GetLogoutUrl() string`

GetLogoutUrl returns the LogoutUrl field if non-nil, zero value otherwise.

### GetLogoutUrlOk

`func (o *UserSourceCreateUserSourceConfigOneOf5) GetLogoutUrlOk() (*string, bool)`

GetLogoutUrlOk returns a tuple with the LogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUrl

`func (o *UserSourceCreateUserSourceConfigOneOf5) SetLogoutUrl(v string)`

SetLogoutUrl sets LogoutUrl field to given value.

### HasLogoutUrl

`func (o *UserSourceCreateUserSourceConfigOneOf5) HasLogoutUrl() bool`

HasLogoutUrl returns a boolean if a field has been set.

### GetSAMLSignatureMode

`func (o *UserSourceCreateUserSourceConfigOneOf5) GetSAMLSignatureMode() string`

GetSAMLSignatureMode returns the SAMLSignatureMode field if non-nil, zero value otherwise.

### GetSAMLSignatureModeOk

`func (o *UserSourceCreateUserSourceConfigOneOf5) GetSAMLSignatureModeOk() (*string, bool)`

GetSAMLSignatureModeOk returns a tuple with the SAMLSignatureMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAMLSignatureMode

`func (o *UserSourceCreateUserSourceConfigOneOf5) SetSAMLSignatureMode(v string)`

SetSAMLSignatureMode sets SAMLSignatureMode field to given value.

### HasSAMLSignatureMode

`func (o *UserSourceCreateUserSourceConfigOneOf5) HasSAMLSignatureMode() bool`

HasSAMLSignatureMode returns a boolean if a field has been set.

### GetRequest509Certificate

`func (o *UserSourceCreateUserSourceConfigOneOf5) GetRequest509Certificate() string`

GetRequest509Certificate returns the Request509Certificate field if non-nil, zero value otherwise.

### GetRequest509CertificateOk

`func (o *UserSourceCreateUserSourceConfigOneOf5) GetRequest509CertificateOk() (*string, bool)`

GetRequest509CertificateOk returns a tuple with the Request509Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest509Certificate

`func (o *UserSourceCreateUserSourceConfigOneOf5) SetRequest509Certificate(v string)`

SetRequest509Certificate sets Request509Certificate field to given value.

### HasRequest509Certificate

`func (o *UserSourceCreateUserSourceConfigOneOf5) HasRequest509Certificate() bool`

HasRequest509Certificate returns a boolean if a field has been set.

### GetRequestPrivateKey

`func (o *UserSourceCreateUserSourceConfigOneOf5) GetRequestPrivateKey() string`

GetRequestPrivateKey returns the RequestPrivateKey field if non-nil, zero value otherwise.

### GetRequestPrivateKeyOk

`func (o *UserSourceCreateUserSourceConfigOneOf5) GetRequestPrivateKeyOk() (*string, bool)`

GetRequestPrivateKeyOk returns a tuple with the RequestPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestPrivateKey

`func (o *UserSourceCreateUserSourceConfigOneOf5) SetRequestPrivateKey(v string)`

SetRequestPrivateKey sets RequestPrivateKey field to given value.

### HasRequestPrivateKey

`func (o *UserSourceCreateUserSourceConfigOneOf5) HasRequestPrivateKey() bool`

HasRequestPrivateKey returns a boolean if a field has been set.

### GetDoNotValidateSignature

`func (o *UserSourceCreateUserSourceConfigOneOf5) GetDoNotValidateSignature() bool`

GetDoNotValidateSignature returns the DoNotValidateSignature field if non-nil, zero value otherwise.

### GetDoNotValidateSignatureOk

`func (o *UserSourceCreateUserSourceConfigOneOf5) GetDoNotValidateSignatureOk() (*bool, bool)`

GetDoNotValidateSignatureOk returns a tuple with the DoNotValidateSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotValidateSignature

`func (o *UserSourceCreateUserSourceConfigOneOf5) SetDoNotValidateSignature(v bool)`

SetDoNotValidateSignature sets DoNotValidateSignature field to given value.

### HasDoNotValidateSignature

`func (o *UserSourceCreateUserSourceConfigOneOf5) HasDoNotValidateSignature() bool`

HasDoNotValidateSignature returns a boolean if a field has been set.

### GetPublicKey

`func (o *UserSourceCreateUserSourceConfigOneOf5) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *UserSourceCreateUserSourceConfigOneOf5) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *UserSourceCreateUserSourceConfigOneOf5) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *UserSourceCreateUserSourceConfigOneOf5) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetPrivateKey

`func (o *UserSourceCreateUserSourceConfigOneOf5) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *UserSourceCreateUserSourceConfigOneOf5) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *UserSourceCreateUserSourceConfigOneOf5) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *UserSourceCreateUserSourceConfigOneOf5) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetGivenNameAttribute

`func (o *UserSourceCreateUserSourceConfigOneOf5) GetGivenNameAttribute() string`

GetGivenNameAttribute returns the GivenNameAttribute field if non-nil, zero value otherwise.

### GetGivenNameAttributeOk

`func (o *UserSourceCreateUserSourceConfigOneOf5) GetGivenNameAttributeOk() (*string, bool)`

GetGivenNameAttributeOk returns a tuple with the GivenNameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenNameAttribute

`func (o *UserSourceCreateUserSourceConfigOneOf5) SetGivenNameAttribute(v string)`

SetGivenNameAttribute sets GivenNameAttribute field to given value.

### HasGivenNameAttribute

`func (o *UserSourceCreateUserSourceConfigOneOf5) HasGivenNameAttribute() bool`

HasGivenNameAttribute returns a boolean if a field has been set.

### GetSurnameAttribute

`func (o *UserSourceCreateUserSourceConfigOneOf5) GetSurnameAttribute() string`

GetSurnameAttribute returns the SurnameAttribute field if non-nil, zero value otherwise.

### GetSurnameAttributeOk

`func (o *UserSourceCreateUserSourceConfigOneOf5) GetSurnameAttributeOk() (*string, bool)`

GetSurnameAttributeOk returns a tuple with the SurnameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurnameAttribute

`func (o *UserSourceCreateUserSourceConfigOneOf5) SetSurnameAttribute(v string)`

SetSurnameAttribute sets SurnameAttribute field to given value.

### HasSurnameAttribute

`func (o *UserSourceCreateUserSourceConfigOneOf5) HasSurnameAttribute() bool`

HasSurnameAttribute returns a boolean if a field has been set.

### GetRoleAttributeName

`func (o *UserSourceCreateUserSourceConfigOneOf5) GetRoleAttributeName() string`

GetRoleAttributeName returns the RoleAttributeName field if non-nil, zero value otherwise.

### GetRoleAttributeNameOk

`func (o *UserSourceCreateUserSourceConfigOneOf5) GetRoleAttributeNameOk() (*string, bool)`

GetRoleAttributeNameOk returns a tuple with the RoleAttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleAttributeName

`func (o *UserSourceCreateUserSourceConfigOneOf5) SetRoleAttributeName(v string)`

SetRoleAttributeName sets RoleAttributeName field to given value.

### HasRoleAttributeName

`func (o *UserSourceCreateUserSourceConfigOneOf5) HasRoleAttributeName() bool`

HasRoleAttributeName returns a boolean if a field has been set.

### GetRequiredAttributeValue

`func (o *UserSourceCreateUserSourceConfigOneOf5) GetRequiredAttributeValue() string`

GetRequiredAttributeValue returns the RequiredAttributeValue field if non-nil, zero value otherwise.

### GetRequiredAttributeValueOk

`func (o *UserSourceCreateUserSourceConfigOneOf5) GetRequiredAttributeValueOk() (*string, bool)`

GetRequiredAttributeValueOk returns a tuple with the RequiredAttributeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredAttributeValue

`func (o *UserSourceCreateUserSourceConfigOneOf5) SetRequiredAttributeValue(v string)`

SetRequiredAttributeValue sets RequiredAttributeValue field to given value.

### HasRequiredAttributeValue

`func (o *UserSourceCreateUserSourceConfigOneOf5) HasRequiredAttributeValue() bool`

HasRequiredAttributeValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


