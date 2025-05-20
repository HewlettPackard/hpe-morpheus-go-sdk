# AwsResourcePoolConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CidrBlock** | Pointer to **string** | Provide the base CIDR Block to use for this VPC (must be between a /16 and /28 Block) | [optional] 
**Tenancy** | Pointer to **string** | default or dedicated | [optional] [default to "default"]

## Methods

### NewAwsResourcePoolConfig

`func NewAwsResourcePoolConfig() *AwsResourcePoolConfig`

NewAwsResourcePoolConfig instantiates a new AwsResourcePoolConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAwsResourcePoolConfigWithDefaults

`func NewAwsResourcePoolConfigWithDefaults() *AwsResourcePoolConfig`

NewAwsResourcePoolConfigWithDefaults instantiates a new AwsResourcePoolConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidrBlock

`func (o *AwsResourcePoolConfig) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *AwsResourcePoolConfig) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *AwsResourcePoolConfig) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.

### HasCidrBlock

`func (o *AwsResourcePoolConfig) HasCidrBlock() bool`

HasCidrBlock returns a boolean if a field has been set.

### GetTenancy

`func (o *AwsResourcePoolConfig) GetTenancy() string`

GetTenancy returns the Tenancy field if non-nil, zero value otherwise.

### GetTenancyOk

`func (o *AwsResourcePoolConfig) GetTenancyOk() (*string, bool)`

GetTenancyOk returns a tuple with the Tenancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenancy

`func (o *AwsResourcePoolConfig) SetTenancy(v string)`

SetTenancy sets Tenancy field to given value.

### HasTenancy

`func (o *AwsResourcePoolConfig) HasTenancy() bool`

HasTenancy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


