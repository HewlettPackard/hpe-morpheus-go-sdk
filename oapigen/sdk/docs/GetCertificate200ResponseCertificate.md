# GetCertificate200ResponseCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**DomainName** | Pointer to **NullableString** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**IntegrationId** | Pointer to **NullableInt64** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Generated** | Pointer to **bool** |  | [optional] 
**Wildcard** | Pointer to **bool** |  | [optional] 
**SelfSigned** | Pointer to **bool** |  | [optional] 
**Type** | Pointer to [**GetCertificate200ResponseCertificateType**](GetCertificate200ResponseCertificateType.md) |  | [optional] 
**Category** | Pointer to **NullableString** |  | [optional] 
**CommonName** | Pointer to **NullableString** |  | [optional] 
**CertType** | Pointer to **NullableString** |  | [optional] 
**KeyFileMD5** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewGetCertificate200ResponseCertificate

`func NewGetCertificate200ResponseCertificate() *GetCertificate200ResponseCertificate`

NewGetCertificate200ResponseCertificate instantiates a new GetCertificate200ResponseCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCertificate200ResponseCertificateWithDefaults

`func NewGetCertificate200ResponseCertificateWithDefaults() *GetCertificate200ResponseCertificate`

NewGetCertificate200ResponseCertificateWithDefaults instantiates a new GetCertificate200ResponseCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetCertificate200ResponseCertificate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetCertificate200ResponseCertificate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetCertificate200ResponseCertificate) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetCertificate200ResponseCertificate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetCertificate200ResponseCertificate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetCertificate200ResponseCertificate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetCertificate200ResponseCertificate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetCertificate200ResponseCertificate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetCertificate200ResponseCertificate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetCertificate200ResponseCertificate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetCertificate200ResponseCertificate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetCertificate200ResponseCertificate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetCertificate200ResponseCertificate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetCertificate200ResponseCertificate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDomainName

`func (o *GetCertificate200ResponseCertificate) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *GetCertificate200ResponseCertificate) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *GetCertificate200ResponseCertificate) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *GetCertificate200ResponseCertificate) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### SetDomainNameNil

`func (o *GetCertificate200ResponseCertificate) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *GetCertificate200ResponseCertificate) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
### GetAccountId

`func (o *GetCertificate200ResponseCertificate) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetCertificate200ResponseCertificate) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetCertificate200ResponseCertificate) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetCertificate200ResponseCertificate) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetIntegrationId

`func (o *GetCertificate200ResponseCertificate) GetIntegrationId() int64`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *GetCertificate200ResponseCertificate) GetIntegrationIdOk() (*int64, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *GetCertificate200ResponseCertificate) SetIntegrationId(v int64)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *GetCertificate200ResponseCertificate) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### SetIntegrationIdNil

`func (o *GetCertificate200ResponseCertificate) SetIntegrationIdNil(b bool)`

 SetIntegrationIdNil sets the value for IntegrationId to be an explicit nil

### UnsetIntegrationId
`func (o *GetCertificate200ResponseCertificate) UnsetIntegrationId()`

UnsetIntegrationId ensures that no value is present for IntegrationId, not even an explicit nil
### GetEnabled

`func (o *GetCertificate200ResponseCertificate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetCertificate200ResponseCertificate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetCertificate200ResponseCertificate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetCertificate200ResponseCertificate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGenerated

`func (o *GetCertificate200ResponseCertificate) GetGenerated() bool`

GetGenerated returns the Generated field if non-nil, zero value otherwise.

### GetGeneratedOk

`func (o *GetCertificate200ResponseCertificate) GetGeneratedOk() (*bool, bool)`

GetGeneratedOk returns a tuple with the Generated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerated

`func (o *GetCertificate200ResponseCertificate) SetGenerated(v bool)`

SetGenerated sets Generated field to given value.

### HasGenerated

`func (o *GetCertificate200ResponseCertificate) HasGenerated() bool`

HasGenerated returns a boolean if a field has been set.

### GetWildcard

`func (o *GetCertificate200ResponseCertificate) GetWildcard() bool`

GetWildcard returns the Wildcard field if non-nil, zero value otherwise.

### GetWildcardOk

`func (o *GetCertificate200ResponseCertificate) GetWildcardOk() (*bool, bool)`

GetWildcardOk returns a tuple with the Wildcard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWildcard

`func (o *GetCertificate200ResponseCertificate) SetWildcard(v bool)`

SetWildcard sets Wildcard field to given value.

### HasWildcard

`func (o *GetCertificate200ResponseCertificate) HasWildcard() bool`

HasWildcard returns a boolean if a field has been set.

### GetSelfSigned

`func (o *GetCertificate200ResponseCertificate) GetSelfSigned() bool`

GetSelfSigned returns the SelfSigned field if non-nil, zero value otherwise.

### GetSelfSignedOk

`func (o *GetCertificate200ResponseCertificate) GetSelfSignedOk() (*bool, bool)`

GetSelfSignedOk returns a tuple with the SelfSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfSigned

`func (o *GetCertificate200ResponseCertificate) SetSelfSigned(v bool)`

SetSelfSigned sets SelfSigned field to given value.

### HasSelfSigned

`func (o *GetCertificate200ResponseCertificate) HasSelfSigned() bool`

HasSelfSigned returns a boolean if a field has been set.

### GetType

`func (o *GetCertificate200ResponseCertificate) GetType() GetCertificate200ResponseCertificateType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetCertificate200ResponseCertificate) GetTypeOk() (*GetCertificate200ResponseCertificateType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetCertificate200ResponseCertificate) SetType(v GetCertificate200ResponseCertificateType)`

SetType sets Type field to given value.

### HasType

`func (o *GetCertificate200ResponseCertificate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCategory

`func (o *GetCertificate200ResponseCertificate) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetCertificate200ResponseCertificate) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetCertificate200ResponseCertificate) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetCertificate200ResponseCertificate) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *GetCertificate200ResponseCertificate) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *GetCertificate200ResponseCertificate) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetCommonName

`func (o *GetCertificate200ResponseCertificate) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *GetCertificate200ResponseCertificate) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *GetCertificate200ResponseCertificate) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *GetCertificate200ResponseCertificate) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### SetCommonNameNil

`func (o *GetCertificate200ResponseCertificate) SetCommonNameNil(b bool)`

 SetCommonNameNil sets the value for CommonName to be an explicit nil

### UnsetCommonName
`func (o *GetCertificate200ResponseCertificate) UnsetCommonName()`

UnsetCommonName ensures that no value is present for CommonName, not even an explicit nil
### GetCertType

`func (o *GetCertificate200ResponseCertificate) GetCertType() string`

GetCertType returns the CertType field if non-nil, zero value otherwise.

### GetCertTypeOk

`func (o *GetCertificate200ResponseCertificate) GetCertTypeOk() (*string, bool)`

GetCertTypeOk returns a tuple with the CertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertType

`func (o *GetCertificate200ResponseCertificate) SetCertType(v string)`

SetCertType sets CertType field to given value.

### HasCertType

`func (o *GetCertificate200ResponseCertificate) HasCertType() bool`

HasCertType returns a boolean if a field has been set.

### SetCertTypeNil

`func (o *GetCertificate200ResponseCertificate) SetCertTypeNil(b bool)`

 SetCertTypeNil sets the value for CertType to be an explicit nil

### UnsetCertType
`func (o *GetCertificate200ResponseCertificate) UnsetCertType()`

UnsetCertType ensures that no value is present for CertType, not even an explicit nil
### GetKeyFileMD5

`func (o *GetCertificate200ResponseCertificate) GetKeyFileMD5() string`

GetKeyFileMD5 returns the KeyFileMD5 field if non-nil, zero value otherwise.

### GetKeyFileMD5Ok

`func (o *GetCertificate200ResponseCertificate) GetKeyFileMD5Ok() (*string, bool)`

GetKeyFileMD5Ok returns a tuple with the KeyFileMD5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyFileMD5

`func (o *GetCertificate200ResponseCertificate) SetKeyFileMD5(v string)`

SetKeyFileMD5 sets KeyFileMD5 field to given value.

### HasKeyFileMD5

`func (o *GetCertificate200ResponseCertificate) HasKeyFileMD5() bool`

HasKeyFileMD5 returns a boolean if a field has been set.

### SetKeyFileMD5Nil

`func (o *GetCertificate200ResponseCertificate) SetKeyFileMD5Nil(b bool)`

 SetKeyFileMD5Nil sets the value for KeyFileMD5 to be an explicit nil

### UnsetKeyFileMD5
`func (o *GetCertificate200ResponseCertificate) UnsetKeyFileMD5()`

UnsetKeyFileMD5 ensures that no value is present for KeyFileMD5, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


