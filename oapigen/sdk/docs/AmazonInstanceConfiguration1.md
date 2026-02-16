# AmazonInstanceConfiguration1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NoAgent** | Pointer to **NullableBool** | Skipping Agent installation will result in a lack of logging and guest operating system statistics. Automation scripts may also be adversely affected. | [optional] [default to false]
**IsEC2** | Pointer to **string** | Amazon Cloud Type | [optional] [default to "false"]
**AvailabilityId** | Pointer to **string** | Amazon Zone | [optional] 
**SecurityId** | Pointer to **string** | Security Group | [optional] 
**PublicIpType** | Pointer to **string** | Public IP | [optional] 
**InstanceProfile** | Pointer to **string** | IAM Profile | [optional] 
**KmsKeyId** | Pointer to **string** | KMS Key ID | [optional] 

## Methods

### NewAmazonInstanceConfiguration1

`func NewAmazonInstanceConfiguration1() *AmazonInstanceConfiguration1`

NewAmazonInstanceConfiguration1 instantiates a new AmazonInstanceConfiguration1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAmazonInstanceConfiguration1WithDefaults

`func NewAmazonInstanceConfiguration1WithDefaults() *AmazonInstanceConfiguration1`

NewAmazonInstanceConfiguration1WithDefaults instantiates a new AmazonInstanceConfiguration1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNoAgent

`func (o *AmazonInstanceConfiguration1) GetNoAgent() bool`

GetNoAgent returns the NoAgent field if non-nil, zero value otherwise.

### GetNoAgentOk

`func (o *AmazonInstanceConfiguration1) GetNoAgentOk() (*bool, bool)`

GetNoAgentOk returns a tuple with the NoAgent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoAgent

`func (o *AmazonInstanceConfiguration1) SetNoAgent(v bool)`

SetNoAgent sets NoAgent field to given value.

### HasNoAgent

`func (o *AmazonInstanceConfiguration1) HasNoAgent() bool`

HasNoAgent returns a boolean if a field has been set.

### SetNoAgentNil

`func (o *AmazonInstanceConfiguration1) SetNoAgentNil(b bool)`

 SetNoAgentNil sets the value for NoAgent to be an explicit nil

### UnsetNoAgent
`func (o *AmazonInstanceConfiguration1) UnsetNoAgent()`

UnsetNoAgent ensures that no value is present for NoAgent, not even an explicit nil
### GetIsEC2

`func (o *AmazonInstanceConfiguration1) GetIsEC2() string`

GetIsEC2 returns the IsEC2 field if non-nil, zero value otherwise.

### GetIsEC2Ok

`func (o *AmazonInstanceConfiguration1) GetIsEC2Ok() (*string, bool)`

GetIsEC2Ok returns a tuple with the IsEC2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsEC2

`func (o *AmazonInstanceConfiguration1) SetIsEC2(v string)`

SetIsEC2 sets IsEC2 field to given value.

### HasIsEC2

`func (o *AmazonInstanceConfiguration1) HasIsEC2() bool`

HasIsEC2 returns a boolean if a field has been set.

### GetAvailabilityId

`func (o *AmazonInstanceConfiguration1) GetAvailabilityId() string`

GetAvailabilityId returns the AvailabilityId field if non-nil, zero value otherwise.

### GetAvailabilityIdOk

`func (o *AmazonInstanceConfiguration1) GetAvailabilityIdOk() (*string, bool)`

GetAvailabilityIdOk returns a tuple with the AvailabilityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailabilityId

`func (o *AmazonInstanceConfiguration1) SetAvailabilityId(v string)`

SetAvailabilityId sets AvailabilityId field to given value.

### HasAvailabilityId

`func (o *AmazonInstanceConfiguration1) HasAvailabilityId() bool`

HasAvailabilityId returns a boolean if a field has been set.

### GetSecurityId

`func (o *AmazonInstanceConfiguration1) GetSecurityId() string`

GetSecurityId returns the SecurityId field if non-nil, zero value otherwise.

### GetSecurityIdOk

`func (o *AmazonInstanceConfiguration1) GetSecurityIdOk() (*string, bool)`

GetSecurityIdOk returns a tuple with the SecurityId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityId

`func (o *AmazonInstanceConfiguration1) SetSecurityId(v string)`

SetSecurityId sets SecurityId field to given value.

### HasSecurityId

`func (o *AmazonInstanceConfiguration1) HasSecurityId() bool`

HasSecurityId returns a boolean if a field has been set.

### GetPublicIpType

`func (o *AmazonInstanceConfiguration1) GetPublicIpType() string`

GetPublicIpType returns the PublicIpType field if non-nil, zero value otherwise.

### GetPublicIpTypeOk

`func (o *AmazonInstanceConfiguration1) GetPublicIpTypeOk() (*string, bool)`

GetPublicIpTypeOk returns a tuple with the PublicIpType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicIpType

`func (o *AmazonInstanceConfiguration1) SetPublicIpType(v string)`

SetPublicIpType sets PublicIpType field to given value.

### HasPublicIpType

`func (o *AmazonInstanceConfiguration1) HasPublicIpType() bool`

HasPublicIpType returns a boolean if a field has been set.

### GetInstanceProfile

`func (o *AmazonInstanceConfiguration1) GetInstanceProfile() string`

GetInstanceProfile returns the InstanceProfile field if non-nil, zero value otherwise.

### GetInstanceProfileOk

`func (o *AmazonInstanceConfiguration1) GetInstanceProfileOk() (*string, bool)`

GetInstanceProfileOk returns a tuple with the InstanceProfile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceProfile

`func (o *AmazonInstanceConfiguration1) SetInstanceProfile(v string)`

SetInstanceProfile sets InstanceProfile field to given value.

### HasInstanceProfile

`func (o *AmazonInstanceConfiguration1) HasInstanceProfile() bool`

HasInstanceProfile returns a boolean if a field has been set.

### GetKmsKeyId

`func (o *AmazonInstanceConfiguration1) GetKmsKeyId() string`

GetKmsKeyId returns the KmsKeyId field if non-nil, zero value otherwise.

### GetKmsKeyIdOk

`func (o *AmazonInstanceConfiguration1) GetKmsKeyIdOk() (*string, bool)`

GetKmsKeyIdOk returns a tuple with the KmsKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKmsKeyId

`func (o *AmazonInstanceConfiguration1) SetKmsKeyId(v string)`

SetKmsKeyId sets KmsKeyId field to given value.

### HasKmsKeyId

`func (o *AmazonInstanceConfiguration1) HasKmsKeyId() bool`

HasKmsKeyId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


