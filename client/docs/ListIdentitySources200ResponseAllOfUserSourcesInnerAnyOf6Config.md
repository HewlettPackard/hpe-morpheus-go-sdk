# ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Url** | Pointer to **string** |  | [optional] 
**LogoutUrl** | Pointer to **string** |  | [optional] 
**DoNotIncludeSAMLRequest** | Pointer to **bool** |  | [optional] 
**SAMLSignatureMode** | Pointer to **string** |  | [optional] 
**DoNotValidateSignature** | Pointer to **bool** |  | [optional] 
**DoNotValidateStatusCode** | Pointer to **bool** |  | [optional] 
**DoNotValidateDestination** | Pointer to **bool** |  | [optional] 
**DoNotValidateIssueInstants** | Pointer to **bool** |  | [optional] 
**DoNotValidateAssertions** | Pointer to **bool** |  | [optional] 
**GivenNameAttribute** | Pointer to **string** |  | [optional] 
**SurnameAttribute** | Pointer to **string** |  | [optional] 
**EmailAttribute** | Pointer to **string** |  | [optional] 
**RequiredAttributeValue** | Pointer to **string** |  | [optional] 
**RoleAttributeName** | Pointer to **string** |  | [optional] 
**AzureTenantId** | Pointer to **string** |  | [optional] 
**AzureAppId** | Pointer to **string** |  | [optional] 
**AzureAppSecret** | Pointer to **string** |  | [optional] 
**RoleLinkAttributeName** | Pointer to **string** |  | [optional] 
**PublicKey** | Pointer to **string** |  | [optional] 
**AzureAppSecretHash** | Pointer to **string** |  | [optional] 

## Methods

### NewListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config

`func NewListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config() *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config`

NewListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config instantiates a new ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6ConfigWithDefaults

`func NewListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6ConfigWithDefaults() *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config`

NewListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6ConfigWithDefaults instantiates a new ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUrl

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetUrl(v string)`

SetUrl sets Url field to given value.

### HasUrl

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) HasUrl() bool`

HasUrl returns a boolean if a field has been set.

### GetLogoutUrl

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetLogoutUrl() string`

GetLogoutUrl returns the LogoutUrl field if non-nil, zero value otherwise.

### GetLogoutUrlOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetLogoutUrlOk() (*string, bool)`

GetLogoutUrlOk returns a tuple with the LogoutUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogoutUrl

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetLogoutUrl(v string)`

SetLogoutUrl sets LogoutUrl field to given value.

### HasLogoutUrl

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) HasLogoutUrl() bool`

HasLogoutUrl returns a boolean if a field has been set.

### GetDoNotIncludeSAMLRequest

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetDoNotIncludeSAMLRequest() bool`

GetDoNotIncludeSAMLRequest returns the DoNotIncludeSAMLRequest field if non-nil, zero value otherwise.

### GetDoNotIncludeSAMLRequestOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetDoNotIncludeSAMLRequestOk() (*bool, bool)`

GetDoNotIncludeSAMLRequestOk returns a tuple with the DoNotIncludeSAMLRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotIncludeSAMLRequest

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetDoNotIncludeSAMLRequest(v bool)`

SetDoNotIncludeSAMLRequest sets DoNotIncludeSAMLRequest field to given value.

### HasDoNotIncludeSAMLRequest

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) HasDoNotIncludeSAMLRequest() bool`

HasDoNotIncludeSAMLRequest returns a boolean if a field has been set.

### GetSAMLSignatureMode

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetSAMLSignatureMode() string`

GetSAMLSignatureMode returns the SAMLSignatureMode field if non-nil, zero value otherwise.

### GetSAMLSignatureModeOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetSAMLSignatureModeOk() (*string, bool)`

GetSAMLSignatureModeOk returns a tuple with the SAMLSignatureMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSAMLSignatureMode

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetSAMLSignatureMode(v string)`

SetSAMLSignatureMode sets SAMLSignatureMode field to given value.

### HasSAMLSignatureMode

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) HasSAMLSignatureMode() bool`

HasSAMLSignatureMode returns a boolean if a field has been set.

### GetDoNotValidateSignature

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetDoNotValidateSignature() bool`

GetDoNotValidateSignature returns the DoNotValidateSignature field if non-nil, zero value otherwise.

### GetDoNotValidateSignatureOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetDoNotValidateSignatureOk() (*bool, bool)`

GetDoNotValidateSignatureOk returns a tuple with the DoNotValidateSignature field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotValidateSignature

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetDoNotValidateSignature(v bool)`

SetDoNotValidateSignature sets DoNotValidateSignature field to given value.

### HasDoNotValidateSignature

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) HasDoNotValidateSignature() bool`

HasDoNotValidateSignature returns a boolean if a field has been set.

### GetDoNotValidateStatusCode

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetDoNotValidateStatusCode() bool`

GetDoNotValidateStatusCode returns the DoNotValidateStatusCode field if non-nil, zero value otherwise.

### GetDoNotValidateStatusCodeOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetDoNotValidateStatusCodeOk() (*bool, bool)`

GetDoNotValidateStatusCodeOk returns a tuple with the DoNotValidateStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotValidateStatusCode

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetDoNotValidateStatusCode(v bool)`

SetDoNotValidateStatusCode sets DoNotValidateStatusCode field to given value.

### HasDoNotValidateStatusCode

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) HasDoNotValidateStatusCode() bool`

HasDoNotValidateStatusCode returns a boolean if a field has been set.

### GetDoNotValidateDestination

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetDoNotValidateDestination() bool`

GetDoNotValidateDestination returns the DoNotValidateDestination field if non-nil, zero value otherwise.

### GetDoNotValidateDestinationOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetDoNotValidateDestinationOk() (*bool, bool)`

GetDoNotValidateDestinationOk returns a tuple with the DoNotValidateDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotValidateDestination

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetDoNotValidateDestination(v bool)`

SetDoNotValidateDestination sets DoNotValidateDestination field to given value.

### HasDoNotValidateDestination

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) HasDoNotValidateDestination() bool`

HasDoNotValidateDestination returns a boolean if a field has been set.

### GetDoNotValidateIssueInstants

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetDoNotValidateIssueInstants() bool`

GetDoNotValidateIssueInstants returns the DoNotValidateIssueInstants field if non-nil, zero value otherwise.

### GetDoNotValidateIssueInstantsOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetDoNotValidateIssueInstantsOk() (*bool, bool)`

GetDoNotValidateIssueInstantsOk returns a tuple with the DoNotValidateIssueInstants field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotValidateIssueInstants

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetDoNotValidateIssueInstants(v bool)`

SetDoNotValidateIssueInstants sets DoNotValidateIssueInstants field to given value.

### HasDoNotValidateIssueInstants

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) HasDoNotValidateIssueInstants() bool`

HasDoNotValidateIssueInstants returns a boolean if a field has been set.

### GetDoNotValidateAssertions

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetDoNotValidateAssertions() bool`

GetDoNotValidateAssertions returns the DoNotValidateAssertions field if non-nil, zero value otherwise.

### GetDoNotValidateAssertionsOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetDoNotValidateAssertionsOk() (*bool, bool)`

GetDoNotValidateAssertionsOk returns a tuple with the DoNotValidateAssertions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotValidateAssertions

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetDoNotValidateAssertions(v bool)`

SetDoNotValidateAssertions sets DoNotValidateAssertions field to given value.

### HasDoNotValidateAssertions

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) HasDoNotValidateAssertions() bool`

HasDoNotValidateAssertions returns a boolean if a field has been set.

### GetGivenNameAttribute

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetGivenNameAttribute() string`

GetGivenNameAttribute returns the GivenNameAttribute field if non-nil, zero value otherwise.

### GetGivenNameAttributeOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetGivenNameAttributeOk() (*string, bool)`

GetGivenNameAttributeOk returns a tuple with the GivenNameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGivenNameAttribute

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetGivenNameAttribute(v string)`

SetGivenNameAttribute sets GivenNameAttribute field to given value.

### HasGivenNameAttribute

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) HasGivenNameAttribute() bool`

HasGivenNameAttribute returns a boolean if a field has been set.

### GetSurnameAttribute

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetSurnameAttribute() string`

GetSurnameAttribute returns the SurnameAttribute field if non-nil, zero value otherwise.

### GetSurnameAttributeOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetSurnameAttributeOk() (*string, bool)`

GetSurnameAttributeOk returns a tuple with the SurnameAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurnameAttribute

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetSurnameAttribute(v string)`

SetSurnameAttribute sets SurnameAttribute field to given value.

### HasSurnameAttribute

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) HasSurnameAttribute() bool`

HasSurnameAttribute returns a boolean if a field has been set.

### GetEmailAttribute

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetEmailAttribute() string`

GetEmailAttribute returns the EmailAttribute field if non-nil, zero value otherwise.

### GetEmailAttributeOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetEmailAttributeOk() (*string, bool)`

GetEmailAttributeOk returns a tuple with the EmailAttribute field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAttribute

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetEmailAttribute(v string)`

SetEmailAttribute sets EmailAttribute field to given value.

### HasEmailAttribute

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) HasEmailAttribute() bool`

HasEmailAttribute returns a boolean if a field has been set.

### GetRequiredAttributeValue

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetRequiredAttributeValue() string`

GetRequiredAttributeValue returns the RequiredAttributeValue field if non-nil, zero value otherwise.

### GetRequiredAttributeValueOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetRequiredAttributeValueOk() (*string, bool)`

GetRequiredAttributeValueOk returns a tuple with the RequiredAttributeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequiredAttributeValue

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetRequiredAttributeValue(v string)`

SetRequiredAttributeValue sets RequiredAttributeValue field to given value.

### HasRequiredAttributeValue

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) HasRequiredAttributeValue() bool`

HasRequiredAttributeValue returns a boolean if a field has been set.

### GetRoleAttributeName

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetRoleAttributeName() string`

GetRoleAttributeName returns the RoleAttributeName field if non-nil, zero value otherwise.

### GetRoleAttributeNameOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetRoleAttributeNameOk() (*string, bool)`

GetRoleAttributeNameOk returns a tuple with the RoleAttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleAttributeName

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetRoleAttributeName(v string)`

SetRoleAttributeName sets RoleAttributeName field to given value.

### HasRoleAttributeName

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) HasRoleAttributeName() bool`

HasRoleAttributeName returns a boolean if a field has been set.

### GetAzureTenantId

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetAzureTenantId() string`

GetAzureTenantId returns the AzureTenantId field if non-nil, zero value otherwise.

### GetAzureTenantIdOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetAzureTenantIdOk() (*string, bool)`

GetAzureTenantIdOk returns a tuple with the AzureTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureTenantId

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetAzureTenantId(v string)`

SetAzureTenantId sets AzureTenantId field to given value.

### HasAzureTenantId

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) HasAzureTenantId() bool`

HasAzureTenantId returns a boolean if a field has been set.

### GetAzureAppId

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetAzureAppId() string`

GetAzureAppId returns the AzureAppId field if non-nil, zero value otherwise.

### GetAzureAppIdOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetAzureAppIdOk() (*string, bool)`

GetAzureAppIdOk returns a tuple with the AzureAppId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAppId

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetAzureAppId(v string)`

SetAzureAppId sets AzureAppId field to given value.

### HasAzureAppId

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) HasAzureAppId() bool`

HasAzureAppId returns a boolean if a field has been set.

### GetAzureAppSecret

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetAzureAppSecret() string`

GetAzureAppSecret returns the AzureAppSecret field if non-nil, zero value otherwise.

### GetAzureAppSecretOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetAzureAppSecretOk() (*string, bool)`

GetAzureAppSecretOk returns a tuple with the AzureAppSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAppSecret

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetAzureAppSecret(v string)`

SetAzureAppSecret sets AzureAppSecret field to given value.

### HasAzureAppSecret

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) HasAzureAppSecret() bool`

HasAzureAppSecret returns a boolean if a field has been set.

### GetRoleLinkAttributeName

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetRoleLinkAttributeName() string`

GetRoleLinkAttributeName returns the RoleLinkAttributeName field if non-nil, zero value otherwise.

### GetRoleLinkAttributeNameOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetRoleLinkAttributeNameOk() (*string, bool)`

GetRoleLinkAttributeNameOk returns a tuple with the RoleLinkAttributeName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoleLinkAttributeName

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetRoleLinkAttributeName(v string)`

SetRoleLinkAttributeName sets RoleLinkAttributeName field to given value.

### HasRoleLinkAttributeName

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) HasRoleLinkAttributeName() bool`

HasRoleLinkAttributeName returns a boolean if a field has been set.

### GetPublicKey

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetAzureAppSecretHash

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetAzureAppSecretHash() string`

GetAzureAppSecretHash returns the AzureAppSecretHash field if non-nil, zero value otherwise.

### GetAzureAppSecretHashOk

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) GetAzureAppSecretHashOk() (*string, bool)`

GetAzureAppSecretHashOk returns a tuple with the AzureAppSecretHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureAppSecretHash

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) SetAzureAppSecretHash(v string)`

SetAzureAppSecretHash sets AzureAppSecretHash field to given value.

### HasAzureAppSecretHash

`func (o *ListIdentitySources200ResponseAllOfUserSourcesInnerAnyOf6Config) HasAzureAppSecretHash() bool`

HasAzureAppSecretHash returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


