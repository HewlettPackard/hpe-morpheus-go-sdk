# CloudFoundryResourcePoolConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Managers** | Pointer to **[]string** | Array of manager usernames | [optional] 
**Developers** | Pointer to **[]string** | Array of developer usernames | [optional] 
**Auditors** | Pointer to **[]string** | Array of auditor usernames | [optional] 

## Methods

### NewCloudFoundryResourcePoolConfig

`func NewCloudFoundryResourcePoolConfig() *CloudFoundryResourcePoolConfig`

NewCloudFoundryResourcePoolConfig instantiates a new CloudFoundryResourcePoolConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCloudFoundryResourcePoolConfigWithDefaults

`func NewCloudFoundryResourcePoolConfigWithDefaults() *CloudFoundryResourcePoolConfig`

NewCloudFoundryResourcePoolConfigWithDefaults instantiates a new CloudFoundryResourcePoolConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManagers

`func (o *CloudFoundryResourcePoolConfig) GetManagers() []string`

GetManagers returns the Managers field if non-nil, zero value otherwise.

### GetManagersOk

`func (o *CloudFoundryResourcePoolConfig) GetManagersOk() (*[]string, bool)`

GetManagersOk returns a tuple with the Managers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManagers

`func (o *CloudFoundryResourcePoolConfig) SetManagers(v []string)`

SetManagers sets Managers field to given value.

### HasManagers

`func (o *CloudFoundryResourcePoolConfig) HasManagers() bool`

HasManagers returns a boolean if a field has been set.

### GetDevelopers

`func (o *CloudFoundryResourcePoolConfig) GetDevelopers() []string`

GetDevelopers returns the Developers field if non-nil, zero value otherwise.

### GetDevelopersOk

`func (o *CloudFoundryResourcePoolConfig) GetDevelopersOk() (*[]string, bool)`

GetDevelopersOk returns a tuple with the Developers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevelopers

`func (o *CloudFoundryResourcePoolConfig) SetDevelopers(v []string)`

SetDevelopers sets Developers field to given value.

### HasDevelopers

`func (o *CloudFoundryResourcePoolConfig) HasDevelopers() bool`

HasDevelopers returns a boolean if a field has been set.

### GetAuditors

`func (o *CloudFoundryResourcePoolConfig) GetAuditors() []string`

GetAuditors returns the Auditors field if non-nil, zero value otherwise.

### GetAuditorsOk

`func (o *CloudFoundryResourcePoolConfig) GetAuditorsOk() (*[]string, bool)`

GetAuditorsOk returns a tuple with the Auditors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuditors

`func (o *CloudFoundryResourcePoolConfig) SetAuditors(v []string)`

SetAuditors sets Auditors field to given value.

### HasAuditors

`func (o *CloudFoundryResourcePoolConfig) HasAuditors() bool`

HasAuditors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


