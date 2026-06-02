# AddCertificate200ResponseAllOfCertificate

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
**Type** | Pointer to [**AddCertificate200ResponseAllOfCertificateType**](AddCertificate200ResponseAllOfCertificateType.md) |  | [optional] 
**Category** | Pointer to **NullableString** |  | [optional] 
**CommonName** | Pointer to **NullableString** |  | [optional] 
**CertType** | Pointer to **NullableString** |  | [optional] 
**KeyFileMD5** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewAddCertificate200ResponseAllOfCertificate

`func NewAddCertificate200ResponseAllOfCertificate() *AddCertificate200ResponseAllOfCertificate`

NewAddCertificate200ResponseAllOfCertificate instantiates a new AddCertificate200ResponseAllOfCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *AddCertificate200ResponseAllOfCertificate) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddCertificate200ResponseAllOfCertificate) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddCertificate200ResponseAllOfCertificate) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddCertificate200ResponseAllOfCertificate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddCertificate200ResponseAllOfCertificate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddCertificate200ResponseAllOfCertificate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddCertificate200ResponseAllOfCertificate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddCertificate200ResponseAllOfCertificate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *AddCertificate200ResponseAllOfCertificate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddCertificate200ResponseAllOfCertificate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddCertificate200ResponseAllOfCertificate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddCertificate200ResponseAllOfCertificate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AddCertificate200ResponseAllOfCertificate) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AddCertificate200ResponseAllOfCertificate) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetDomainName

`func (o *AddCertificate200ResponseAllOfCertificate) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *AddCertificate200ResponseAllOfCertificate) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *AddCertificate200ResponseAllOfCertificate) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *AddCertificate200ResponseAllOfCertificate) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### SetDomainNameNil

`func (o *AddCertificate200ResponseAllOfCertificate) SetDomainNameNil(b bool)`

 SetDomainNameNil sets the value for DomainName to be an explicit nil

### UnsetDomainName
`func (o *AddCertificate200ResponseAllOfCertificate) UnsetDomainName()`

UnsetDomainName ensures that no value is present for DomainName, not even an explicit nil
### GetAccountId

`func (o *AddCertificate200ResponseAllOfCertificate) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AddCertificate200ResponseAllOfCertificate) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AddCertificate200ResponseAllOfCertificate) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AddCertificate200ResponseAllOfCertificate) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetIntegrationId

`func (o *AddCertificate200ResponseAllOfCertificate) GetIntegrationId() int64`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *AddCertificate200ResponseAllOfCertificate) GetIntegrationIdOk() (*int64, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *AddCertificate200ResponseAllOfCertificate) SetIntegrationId(v int64)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *AddCertificate200ResponseAllOfCertificate) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.

### SetIntegrationIdNil

`func (o *AddCertificate200ResponseAllOfCertificate) SetIntegrationIdNil(b bool)`

 SetIntegrationIdNil sets the value for IntegrationId to be an explicit nil

### UnsetIntegrationId
`func (o *AddCertificate200ResponseAllOfCertificate) UnsetIntegrationId()`

UnsetIntegrationId ensures that no value is present for IntegrationId, not even an explicit nil
### GetEnabled

`func (o *AddCertificate200ResponseAllOfCertificate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddCertificate200ResponseAllOfCertificate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddCertificate200ResponseAllOfCertificate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddCertificate200ResponseAllOfCertificate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGenerated

`func (o *AddCertificate200ResponseAllOfCertificate) GetGenerated() bool`

GetGenerated returns the Generated field if non-nil, zero value otherwise.

### GetGeneratedOk

`func (o *AddCertificate200ResponseAllOfCertificate) GetGeneratedOk() (*bool, bool)`

GetGeneratedOk returns a tuple with the Generated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGenerated

`func (o *AddCertificate200ResponseAllOfCertificate) SetGenerated(v bool)`

SetGenerated sets Generated field to given value.

### HasGenerated

`func (o *AddCertificate200ResponseAllOfCertificate) HasGenerated() bool`

HasGenerated returns a boolean if a field has been set.

### GetWildcard

`func (o *AddCertificate200ResponseAllOfCertificate) GetWildcard() bool`

GetWildcard returns the Wildcard field if non-nil, zero value otherwise.

### GetWildcardOk

`func (o *AddCertificate200ResponseAllOfCertificate) GetWildcardOk() (*bool, bool)`

GetWildcardOk returns a tuple with the Wildcard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWildcard

`func (o *AddCertificate200ResponseAllOfCertificate) SetWildcard(v bool)`

SetWildcard sets Wildcard field to given value.

### HasWildcard

`func (o *AddCertificate200ResponseAllOfCertificate) HasWildcard() bool`

HasWildcard returns a boolean if a field has been set.

### GetSelfSigned

`func (o *AddCertificate200ResponseAllOfCertificate) GetSelfSigned() bool`

GetSelfSigned returns the SelfSigned field if non-nil, zero value otherwise.

### GetSelfSignedOk

`func (o *AddCertificate200ResponseAllOfCertificate) GetSelfSignedOk() (*bool, bool)`

GetSelfSignedOk returns a tuple with the SelfSigned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelfSigned

`func (o *AddCertificate200ResponseAllOfCertificate) SetSelfSigned(v bool)`

SetSelfSigned sets SelfSigned field to given value.

### HasSelfSigned

`func (o *AddCertificate200ResponseAllOfCertificate) HasSelfSigned() bool`

HasSelfSigned returns a boolean if a field has been set.

### GetType

`func (o *AddCertificate200ResponseAllOfCertificate) GetType() AddCertificate200ResponseAllOfCertificateType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddCertificate200ResponseAllOfCertificate) GetTypeOk() (*AddCertificate200ResponseAllOfCertificateType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddCertificate200ResponseAllOfCertificate) SetType(v AddCertificate200ResponseAllOfCertificateType)`

SetType sets Type field to given value.

### HasType

`func (o *AddCertificate200ResponseAllOfCertificate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetCategory

`func (o *AddCertificate200ResponseAllOfCertificate) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AddCertificate200ResponseAllOfCertificate) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AddCertificate200ResponseAllOfCertificate) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *AddCertificate200ResponseAllOfCertificate) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### SetCategoryNil

`func (o *AddCertificate200ResponseAllOfCertificate) SetCategoryNil(b bool)`

 SetCategoryNil sets the value for Category to be an explicit nil

### UnsetCategory
`func (o *AddCertificate200ResponseAllOfCertificate) UnsetCategory()`

UnsetCategory ensures that no value is present for Category, not even an explicit nil
### GetCommonName

`func (o *AddCertificate200ResponseAllOfCertificate) GetCommonName() string`

GetCommonName returns the CommonName field if non-nil, zero value otherwise.

### GetCommonNameOk

`func (o *AddCertificate200ResponseAllOfCertificate) GetCommonNameOk() (*string, bool)`

GetCommonNameOk returns a tuple with the CommonName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommonName

`func (o *AddCertificate200ResponseAllOfCertificate) SetCommonName(v string)`

SetCommonName sets CommonName field to given value.

### HasCommonName

`func (o *AddCertificate200ResponseAllOfCertificate) HasCommonName() bool`

HasCommonName returns a boolean if a field has been set.

### SetCommonNameNil

`func (o *AddCertificate200ResponseAllOfCertificate) SetCommonNameNil(b bool)`

 SetCommonNameNil sets the value for CommonName to be an explicit nil

### UnsetCommonName
`func (o *AddCertificate200ResponseAllOfCertificate) UnsetCommonName()`

UnsetCommonName ensures that no value is present for CommonName, not even an explicit nil
### GetCertType

`func (o *AddCertificate200ResponseAllOfCertificate) GetCertType() string`

GetCertType returns the CertType field if non-nil, zero value otherwise.

### GetCertTypeOk

`func (o *AddCertificate200ResponseAllOfCertificate) GetCertTypeOk() (*string, bool)`

GetCertTypeOk returns a tuple with the CertType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertType

`func (o *AddCertificate200ResponseAllOfCertificate) SetCertType(v string)`

SetCertType sets CertType field to given value.

### HasCertType

`func (o *AddCertificate200ResponseAllOfCertificate) HasCertType() bool`

HasCertType returns a boolean if a field has been set.

### SetCertTypeNil

`func (o *AddCertificate200ResponseAllOfCertificate) SetCertTypeNil(b bool)`

 SetCertTypeNil sets the value for CertType to be an explicit nil

### UnsetCertType
`func (o *AddCertificate200ResponseAllOfCertificate) UnsetCertType()`

UnsetCertType ensures that no value is present for CertType, not even an explicit nil
### GetKeyFileMD5

`func (o *AddCertificate200ResponseAllOfCertificate) GetKeyFileMD5() string`

GetKeyFileMD5 returns the KeyFileMD5 field if non-nil, zero value otherwise.

### GetKeyFileMD5Ok

`func (o *AddCertificate200ResponseAllOfCertificate) GetKeyFileMD5Ok() (*string, bool)`

GetKeyFileMD5Ok returns a tuple with the KeyFileMD5 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyFileMD5

`func (o *AddCertificate200ResponseAllOfCertificate) SetKeyFileMD5(v string)`

SetKeyFileMD5 sets KeyFileMD5 field to given value.

### HasKeyFileMD5

`func (o *AddCertificate200ResponseAllOfCertificate) HasKeyFileMD5() bool`

HasKeyFileMD5 returns a boolean if a field has been set.

### SetKeyFileMD5Nil

`func (o *AddCertificate200ResponseAllOfCertificate) SetKeyFileMD5Nil(b bool)`

 SetKeyFileMD5Nil sets the value for KeyFileMD5 to be an explicit nil

### UnsetKeyFileMD5
`func (o *AddCertificate200ResponseAllOfCertificate) UnsetKeyFileMD5()`

UnsetKeyFileMD5 ensures that no value is present for KeyFileMD5, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


