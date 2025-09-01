# AddCloudsRequestZoneConfigAnyOfOneOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Endpoint** | **string** | AWS endpoint | 
**AccessKey** | Pointer to **string** | AWS access key | [optional] 
**SecretKey** | Pointer to **string** | AWS secret key | [optional] 
**UseHostCredentials** | Pointer to **string** | Whether to use the IAM profile associated with the Morpheus server or not | [optional] [default to "on"]
**EbsEncryption** | Pointer to **string** | Determines whether to configure default EBS volume encryption or not | [optional] [default to "on"]
**StsAssumeRole** | Pointer to **string** | The AWS IAM role ARN to assume for authentication | [optional] 
**ConfigManagementId** | Pointer to **string** | The id of the configuration management integration associated with the AWS cloud | [optional] 
**Vpc** | Pointer to **string** | The VPC ID for a specific VPC | [optional] 

## Methods

### NewAddCloudsRequestZoneConfigAnyOfOneOf

`func NewAddCloudsRequestZoneConfigAnyOfOneOf(endpoint string, ) *AddCloudsRequestZoneConfigAnyOfOneOf`

NewAddCloudsRequestZoneConfigAnyOfOneOf instantiates a new AddCloudsRequestZoneConfigAnyOfOneOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCloudsRequestZoneConfigAnyOfOneOfWithDefaults

`func NewAddCloudsRequestZoneConfigAnyOfOneOfWithDefaults() *AddCloudsRequestZoneConfigAnyOfOneOf`

NewAddCloudsRequestZoneConfigAnyOfOneOfWithDefaults instantiates a new AddCloudsRequestZoneConfigAnyOfOneOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoint

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetAccessKey

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetUseHostCredentials

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf) GetUseHostCredentials() string`

GetUseHostCredentials returns the UseHostCredentials field if non-nil, zero value otherwise.

### GetUseHostCredentialsOk

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf) GetUseHostCredentialsOk() (*string, bool)`

GetUseHostCredentialsOk returns a tuple with the UseHostCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseHostCredentials

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf) SetUseHostCredentials(v string)`

SetUseHostCredentials sets UseHostCredentials field to given value.

### HasUseHostCredentials

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf) HasUseHostCredentials() bool`

HasUseHostCredentials returns a boolean if a field has been set.

### GetEbsEncryption

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf) GetEbsEncryption() string`

GetEbsEncryption returns the EbsEncryption field if non-nil, zero value otherwise.

### GetEbsEncryptionOk

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf) GetEbsEncryptionOk() (*string, bool)`

GetEbsEncryptionOk returns a tuple with the EbsEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbsEncryption

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf) SetEbsEncryption(v string)`

SetEbsEncryption sets EbsEncryption field to given value.

### HasEbsEncryption

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf) HasEbsEncryption() bool`

HasEbsEncryption returns a boolean if a field has been set.

### GetStsAssumeRole

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf) GetStsAssumeRole() string`

GetStsAssumeRole returns the StsAssumeRole field if non-nil, zero value otherwise.

### GetStsAssumeRoleOk

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf) GetStsAssumeRoleOk() (*string, bool)`

GetStsAssumeRoleOk returns a tuple with the StsAssumeRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsAssumeRole

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf) SetStsAssumeRole(v string)`

SetStsAssumeRole sets StsAssumeRole field to given value.

### HasStsAssumeRole

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf) HasStsAssumeRole() bool`

HasStsAssumeRole returns a boolean if a field has been set.

### GetConfigManagementId

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf) GetConfigManagementId() string`

GetConfigManagementId returns the ConfigManagementId field if non-nil, zero value otherwise.

### GetConfigManagementIdOk

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf) GetConfigManagementIdOk() (*string, bool)`

GetConfigManagementIdOk returns a tuple with the ConfigManagementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigManagementId

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf) SetConfigManagementId(v string)`

SetConfigManagementId sets ConfigManagementId field to given value.

### HasConfigManagementId

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf) HasConfigManagementId() bool`

HasConfigManagementId returns a boolean if a field has been set.

### GetVpc

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf) GetVpc() string`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf) GetVpcOk() (*string, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf) SetVpc(v string)`

SetVpc sets Vpc field to given value.

### HasVpc

`func (o *AddCloudsRequestZoneConfigAnyOfOneOf) HasVpc() bool`

HasVpc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


