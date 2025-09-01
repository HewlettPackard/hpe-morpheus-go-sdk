# CloudCreateConfigAWS

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

### NewCloudCreateConfigAWS

`func NewCloudCreateConfigAWS(endpoint string, ) *CloudCreateConfigAWS`

NewCloudCreateConfigAWS instantiates a new CloudCreateConfigAWS object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudCreateConfigAWSWithDefaults

`func NewCloudCreateConfigAWSWithDefaults() *CloudCreateConfigAWS`

NewCloudCreateConfigAWSWithDefaults instantiates a new CloudCreateConfigAWS object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEndpoint

`func (o *CloudCreateConfigAWS) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *CloudCreateConfigAWS) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *CloudCreateConfigAWS) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.


### GetAccessKey

`func (o *CloudCreateConfigAWS) GetAccessKey() string`

GetAccessKey returns the AccessKey field if non-nil, zero value otherwise.

### GetAccessKeyOk

`func (o *CloudCreateConfigAWS) GetAccessKeyOk() (*string, bool)`

GetAccessKeyOk returns a tuple with the AccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessKey

`func (o *CloudCreateConfigAWS) SetAccessKey(v string)`

SetAccessKey sets AccessKey field to given value.

### HasAccessKey

`func (o *CloudCreateConfigAWS) HasAccessKey() bool`

HasAccessKey returns a boolean if a field has been set.

### GetSecretKey

`func (o *CloudCreateConfigAWS) GetSecretKey() string`

GetSecretKey returns the SecretKey field if non-nil, zero value otherwise.

### GetSecretKeyOk

`func (o *CloudCreateConfigAWS) GetSecretKeyOk() (*string, bool)`

GetSecretKeyOk returns a tuple with the SecretKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecretKey

`func (o *CloudCreateConfigAWS) SetSecretKey(v string)`

SetSecretKey sets SecretKey field to given value.

### HasSecretKey

`func (o *CloudCreateConfigAWS) HasSecretKey() bool`

HasSecretKey returns a boolean if a field has been set.

### GetUseHostCredentials

`func (o *CloudCreateConfigAWS) GetUseHostCredentials() string`

GetUseHostCredentials returns the UseHostCredentials field if non-nil, zero value otherwise.

### GetUseHostCredentialsOk

`func (o *CloudCreateConfigAWS) GetUseHostCredentialsOk() (*string, bool)`

GetUseHostCredentialsOk returns a tuple with the UseHostCredentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUseHostCredentials

`func (o *CloudCreateConfigAWS) SetUseHostCredentials(v string)`

SetUseHostCredentials sets UseHostCredentials field to given value.

### HasUseHostCredentials

`func (o *CloudCreateConfigAWS) HasUseHostCredentials() bool`

HasUseHostCredentials returns a boolean if a field has been set.

### GetEbsEncryption

`func (o *CloudCreateConfigAWS) GetEbsEncryption() string`

GetEbsEncryption returns the EbsEncryption field if non-nil, zero value otherwise.

### GetEbsEncryptionOk

`func (o *CloudCreateConfigAWS) GetEbsEncryptionOk() (*string, bool)`

GetEbsEncryptionOk returns a tuple with the EbsEncryption field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEbsEncryption

`func (o *CloudCreateConfigAWS) SetEbsEncryption(v string)`

SetEbsEncryption sets EbsEncryption field to given value.

### HasEbsEncryption

`func (o *CloudCreateConfigAWS) HasEbsEncryption() bool`

HasEbsEncryption returns a boolean if a field has been set.

### GetStsAssumeRole

`func (o *CloudCreateConfigAWS) GetStsAssumeRole() string`

GetStsAssumeRole returns the StsAssumeRole field if non-nil, zero value otherwise.

### GetStsAssumeRoleOk

`func (o *CloudCreateConfigAWS) GetStsAssumeRoleOk() (*string, bool)`

GetStsAssumeRoleOk returns a tuple with the StsAssumeRole field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStsAssumeRole

`func (o *CloudCreateConfigAWS) SetStsAssumeRole(v string)`

SetStsAssumeRole sets StsAssumeRole field to given value.

### HasStsAssumeRole

`func (o *CloudCreateConfigAWS) HasStsAssumeRole() bool`

HasStsAssumeRole returns a boolean if a field has been set.

### GetConfigManagementId

`func (o *CloudCreateConfigAWS) GetConfigManagementId() string`

GetConfigManagementId returns the ConfigManagementId field if non-nil, zero value otherwise.

### GetConfigManagementIdOk

`func (o *CloudCreateConfigAWS) GetConfigManagementIdOk() (*string, bool)`

GetConfigManagementIdOk returns a tuple with the ConfigManagementId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigManagementId

`func (o *CloudCreateConfigAWS) SetConfigManagementId(v string)`

SetConfigManagementId sets ConfigManagementId field to given value.

### HasConfigManagementId

`func (o *CloudCreateConfigAWS) HasConfigManagementId() bool`

HasConfigManagementId returns a boolean if a field has been set.

### GetVpc

`func (o *CloudCreateConfigAWS) GetVpc() string`

GetVpc returns the Vpc field if non-nil, zero value otherwise.

### GetVpcOk

`func (o *CloudCreateConfigAWS) GetVpcOk() (*string, bool)`

GetVpcOk returns a tuple with the Vpc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVpc

`func (o *CloudCreateConfigAWS) SetVpc(v string)`

SetVpc sets Vpc field to given value.

### HasVpc

`func (o *CloudCreateConfigAWS) HasVpc() bool`

HasVpc returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


