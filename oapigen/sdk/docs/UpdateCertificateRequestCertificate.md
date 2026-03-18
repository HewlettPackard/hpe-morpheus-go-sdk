# UpdateCertificateRequestCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A unique name scoped to your account for the key | [optional] 
**Description** | Pointer to **string** | A description of the certificate | [optional] 
**CertFile** | Pointer to **string** | The contents of the certificate file | [optional] 
**KeyFile** | Pointer to **string** | The contents of the key file | [optional] 
**ChainFile** | Pointer to **string** | The contents of the root certificate file | [optional] 
**DomainName** | Pointer to **string** | The domain name this certificate is tied to | [optional] 
**Wildcard** | Pointer to **bool** | Whether or not this certificate is a wildcard cert | [optional] [default to false]
**Type** | Pointer to **string** | Certificate Type Code to create a certificate of a type other than the default &#39;internal&#39;. | [optional] [default to "internal"]
**IntegrationId** | Pointer to **int64** | ID of the Service (Trust Integration) to create the certificate with, if using a type other than &#39;internal&#39;. eg. Internal, NSXT or Venafi | [optional] 

## Methods

### NewUpdateCertificateRequestCertificate

`func NewUpdateCertificateRequestCertificate() *UpdateCertificateRequestCertificate`

NewUpdateCertificateRequestCertificate instantiates a new UpdateCertificateRequestCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCertificateRequestCertificateWithDefaults

`func NewUpdateCertificateRequestCertificateWithDefaults() *UpdateCertificateRequestCertificate`

NewUpdateCertificateRequestCertificateWithDefaults instantiates a new UpdateCertificateRequestCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateCertificateRequestCertificate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateCertificateRequestCertificate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateCertificateRequestCertificate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateCertificateRequestCertificate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateCertificateRequestCertificate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateCertificateRequestCertificate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateCertificateRequestCertificate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateCertificateRequestCertificate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCertFile

`func (o *UpdateCertificateRequestCertificate) GetCertFile() string`

GetCertFile returns the CertFile field if non-nil, zero value otherwise.

### GetCertFileOk

`func (o *UpdateCertificateRequestCertificate) GetCertFileOk() (*string, bool)`

GetCertFileOk returns a tuple with the CertFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertFile

`func (o *UpdateCertificateRequestCertificate) SetCertFile(v string)`

SetCertFile sets CertFile field to given value.

### HasCertFile

`func (o *UpdateCertificateRequestCertificate) HasCertFile() bool`

HasCertFile returns a boolean if a field has been set.

### GetKeyFile

`func (o *UpdateCertificateRequestCertificate) GetKeyFile() string`

GetKeyFile returns the KeyFile field if non-nil, zero value otherwise.

### GetKeyFileOk

`func (o *UpdateCertificateRequestCertificate) GetKeyFileOk() (*string, bool)`

GetKeyFileOk returns a tuple with the KeyFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyFile

`func (o *UpdateCertificateRequestCertificate) SetKeyFile(v string)`

SetKeyFile sets KeyFile field to given value.

### HasKeyFile

`func (o *UpdateCertificateRequestCertificate) HasKeyFile() bool`

HasKeyFile returns a boolean if a field has been set.

### GetChainFile

`func (o *UpdateCertificateRequestCertificate) GetChainFile() string`

GetChainFile returns the ChainFile field if non-nil, zero value otherwise.

### GetChainFileOk

`func (o *UpdateCertificateRequestCertificate) GetChainFileOk() (*string, bool)`

GetChainFileOk returns a tuple with the ChainFile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChainFile

`func (o *UpdateCertificateRequestCertificate) SetChainFile(v string)`

SetChainFile sets ChainFile field to given value.

### HasChainFile

`func (o *UpdateCertificateRequestCertificate) HasChainFile() bool`

HasChainFile returns a boolean if a field has been set.

### GetDomainName

`func (o *UpdateCertificateRequestCertificate) GetDomainName() string`

GetDomainName returns the DomainName field if non-nil, zero value otherwise.

### GetDomainNameOk

`func (o *UpdateCertificateRequestCertificate) GetDomainNameOk() (*string, bool)`

GetDomainNameOk returns a tuple with the DomainName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainName

`func (o *UpdateCertificateRequestCertificate) SetDomainName(v string)`

SetDomainName sets DomainName field to given value.

### HasDomainName

`func (o *UpdateCertificateRequestCertificate) HasDomainName() bool`

HasDomainName returns a boolean if a field has been set.

### GetWildcard

`func (o *UpdateCertificateRequestCertificate) GetWildcard() bool`

GetWildcard returns the Wildcard field if non-nil, zero value otherwise.

### GetWildcardOk

`func (o *UpdateCertificateRequestCertificate) GetWildcardOk() (*bool, bool)`

GetWildcardOk returns a tuple with the Wildcard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWildcard

`func (o *UpdateCertificateRequestCertificate) SetWildcard(v bool)`

SetWildcard sets Wildcard field to given value.

### HasWildcard

`func (o *UpdateCertificateRequestCertificate) HasWildcard() bool`

HasWildcard returns a boolean if a field has been set.

### GetType

`func (o *UpdateCertificateRequestCertificate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *UpdateCertificateRequestCertificate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *UpdateCertificateRequestCertificate) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *UpdateCertificateRequestCertificate) HasType() bool`

HasType returns a boolean if a field has been set.

### GetIntegrationId

`func (o *UpdateCertificateRequestCertificate) GetIntegrationId() int64`

GetIntegrationId returns the IntegrationId field if non-nil, zero value otherwise.

### GetIntegrationIdOk

`func (o *UpdateCertificateRequestCertificate) GetIntegrationIdOk() (*int64, bool)`

GetIntegrationIdOk returns a tuple with the IntegrationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIntegrationId

`func (o *UpdateCertificateRequestCertificate) SetIntegrationId(v int64)`

SetIntegrationId sets IntegrationId field to given value.

### HasIntegrationId

`func (o *UpdateCertificateRequestCertificate) HasIntegrationId() bool`

HasIntegrationId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


