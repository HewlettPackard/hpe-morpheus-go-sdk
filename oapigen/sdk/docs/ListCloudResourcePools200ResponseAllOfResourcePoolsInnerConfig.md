# ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CidrBlock** | **string** | Provide the base CIDR Block to use for this VPC (must be between a /16 and /28 Block) | 
**Tenancy** | **string** |  | [default to "default"]
**ProjectId** | Pointer to **string** | Project ID can have lowercase letters, digits, or hyphens. It must start with a lowercase letter and end with a letter or number.  | [optional] 
**Parent** | **interface{}** |  | 
**BillingAccount** | **string** |  | 

## Methods

### NewListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfig

`func NewListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfig(cidrBlock string, tenancy string, parent interface{}, billingAccount string, ) *ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfig`

NewListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfig instantiates a new ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigWithDefaults

`func NewListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigWithDefaults() *ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfig`

NewListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfigWithDefaults instantiates a new ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCidrBlock

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfig) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfig) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfig) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.


### GetTenancy

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfig) GetTenancy() string`

GetTenancy returns the Tenancy field if non-nil, zero value otherwise.

### GetTenancyOk

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfig) GetTenancyOk() (*string, bool)`

GetTenancyOk returns a tuple with the Tenancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenancy

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfig) SetTenancy(v string)`

SetTenancy sets Tenancy field to given value.


### GetProjectId

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfig) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfig) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfig) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfig) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetParent

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfig) GetParent() interface{}`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfig) GetParentOk() (*interface{}, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfig) SetParent(v interface{})`

SetParent sets Parent field to given value.


### SetParentNil

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfig) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfig) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetBillingAccount

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfig) GetBillingAccount() string`

GetBillingAccount returns the BillingAccount field if non-nil, zero value otherwise.

### GetBillingAccountOk

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfig) GetBillingAccountOk() (*string, bool)`

GetBillingAccountOk returns a tuple with the BillingAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAccount

`func (o *ListCloudResourcePools200ResponseAllOfResourcePoolsInnerConfig) SetBillingAccount(v string)`

SetBillingAccount sets BillingAccount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


