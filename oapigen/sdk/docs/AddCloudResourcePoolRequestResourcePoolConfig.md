# AddCloudResourcePoolRequestResourcePoolConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CidrBlock** | Pointer to **string** | Provide the base CIDR Block to use for this VPC (must be between a /16 and /28 Block) | [optional] 
**Tenancy** | Pointer to **string** | default or dedicated | [optional] [default to "default"]
**Managers** | Pointer to **[]string** | Array of manager usernames | [optional] 
**Developers** | Pointer to **[]string** | Array of developer usernames | [optional] 
**Auditors** | Pointer to **[]string** | Array of auditor usernames | [optional] 

## Methods

### NewAddCloudResourcePoolRequestResourcePoolConfig

`func NewAddCloudResourcePoolRequestResourcePoolConfig() *AddCloudResourcePoolRequestResourcePoolConfig`

NewAddCloudResourcePoolRequestResourcePoolConfig instantiates a new AddCloudResourcePoolRequestResourcePoolConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCloudResourcePoolRequestResourcePoolConfigWithDefaults

`func NewAddCloudResourcePoolRequestResourcePoolConfigWithDefaults() *AddCloudResourcePoolRequestResourcePoolConfig`

NewAddCloudResourcePoolRequestResourcePoolConfigWithDefaults instantiates a new AddCloudResourcePoolRequestResourcePoolConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidrBlock

`func (o *AddCloudResourcePoolRequestResourcePoolConfig) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *AddCloudResourcePoolRequestResourcePoolConfig) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *AddCloudResourcePoolRequestResourcePoolConfig) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.

### HasCidrBlock

`func (o *AddCloudResourcePoolRequestResourcePoolConfig) HasCidrBlock() bool`

HasCidrBlock returns a boolean if a field has been set.

### GetTenancy

`func (o *AddCloudResourcePoolRequestResourcePoolConfig) GetTenancy() string`

GetTenancy returns the Tenancy field if non-nil, zero value otherwise.

### GetTenancyOk

`func (o *AddCloudResourcePoolRequestResourcePoolConfig) GetTenancyOk() (*string, bool)`

GetTenancyOk returns a tuple with the Tenancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenancy

`func (o *AddCloudResourcePoolRequestResourcePoolConfig) SetTenancy(v string)`

SetTenancy sets Tenancy field to given value.

### HasTenancy

`func (o *AddCloudResourcePoolRequestResourcePoolConfig) HasTenancy() bool`

HasTenancy returns a boolean if a field has been set.

### GetManagers

`func (o *AddCloudResourcePoolRequestResourcePoolConfig) GetManagers() []string`

GetManagers returns the Managers field if non-nil, zero value otherwise.

### GetManagersOk

`func (o *AddCloudResourcePoolRequestResourcePoolConfig) GetManagersOk() (*[]string, bool)`

GetManagersOk returns a tuple with the Managers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagers

`func (o *AddCloudResourcePoolRequestResourcePoolConfig) SetManagers(v []string)`

SetManagers sets Managers field to given value.

### HasManagers

`func (o *AddCloudResourcePoolRequestResourcePoolConfig) HasManagers() bool`

HasManagers returns a boolean if a field has been set.

### GetDevelopers

`func (o *AddCloudResourcePoolRequestResourcePoolConfig) GetDevelopers() []string`

GetDevelopers returns the Developers field if non-nil, zero value otherwise.

### GetDevelopersOk

`func (o *AddCloudResourcePoolRequestResourcePoolConfig) GetDevelopersOk() (*[]string, bool)`

GetDevelopersOk returns a tuple with the Developers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevelopers

`func (o *AddCloudResourcePoolRequestResourcePoolConfig) SetDevelopers(v []string)`

SetDevelopers sets Developers field to given value.

### HasDevelopers

`func (o *AddCloudResourcePoolRequestResourcePoolConfig) HasDevelopers() bool`

HasDevelopers returns a boolean if a field has been set.

### GetAuditors

`func (o *AddCloudResourcePoolRequestResourcePoolConfig) GetAuditors() []string`

GetAuditors returns the Auditors field if non-nil, zero value otherwise.

### GetAuditorsOk

`func (o *AddCloudResourcePoolRequestResourcePoolConfig) GetAuditorsOk() (*[]string, bool)`

GetAuditorsOk returns a tuple with the Auditors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditors

`func (o *AddCloudResourcePoolRequestResourcePoolConfig) SetAuditors(v []string)`

SetAuditors sets Auditors field to given value.

### HasAuditors

`func (o *AddCloudResourcePoolRequestResourcePoolConfig) HasAuditors() bool`

HasAuditors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


