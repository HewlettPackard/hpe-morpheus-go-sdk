# UpdateCloudResourcePool200ResponseResourcePoolAllOfConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CidrBlock** | **string** | Provide the base CIDR Block to use for this VPC (must be between a /16 and /28 Block) | 
**Tenancy** | **string** |  | [default to "default"]
**ProjectId** | Pointer to **string** | Project ID can have lowercase letters, digits, or hyphens. It must start with a lowercase letter and end with a letter or number.  | [optional] 
**Parent** | **interface{}** |  | 
**BillingAccount** | **string** |  | 

## Methods

### NewUpdateCloudResourcePool200ResponseResourcePoolAllOfConfig

`func NewUpdateCloudResourcePool200ResponseResourcePoolAllOfConfig(cidrBlock string, tenancy string, parent interface{}, billingAccount string, ) *UpdateCloudResourcePool200ResponseResourcePoolAllOfConfig`

NewUpdateCloudResourcePool200ResponseResourcePoolAllOfConfig instantiates a new UpdateCloudResourcePool200ResponseResourcePoolAllOfConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetCidrBlock

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfConfig) GetCidrBlock() string`

GetCidrBlock returns the CidrBlock field if non-nil, zero value otherwise.

### GetCidrBlockOk

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfConfig) GetCidrBlockOk() (*string, bool)`

GetCidrBlockOk returns a tuple with the CidrBlock field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCidrBlock

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfConfig) SetCidrBlock(v string)`

SetCidrBlock sets CidrBlock field to given value.


### GetTenancy

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfConfig) GetTenancy() string`

GetTenancy returns the Tenancy field if non-nil, zero value otherwise.

### GetTenancyOk

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfConfig) GetTenancyOk() (*string, bool)`

GetTenancyOk returns a tuple with the Tenancy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenancy

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfConfig) SetTenancy(v string)`

SetTenancy sets Tenancy field to given value.


### GetProjectId

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfConfig) GetProjectId() string`

GetProjectId returns the ProjectId field if non-nil, zero value otherwise.

### GetProjectIdOk

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfConfig) GetProjectIdOk() (*string, bool)`

GetProjectIdOk returns a tuple with the ProjectId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProjectId

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfConfig) SetProjectId(v string)`

SetProjectId sets ProjectId field to given value.

### HasProjectId

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfConfig) HasProjectId() bool`

HasProjectId returns a boolean if a field has been set.

### GetParent

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfConfig) GetParent() interface{}`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfConfig) GetParentOk() (*interface{}, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfConfig) SetParent(v interface{})`

SetParent sets Parent field to given value.


### SetParentNil

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfConfig) SetParentNil(b bool)`

 SetParentNil sets the value for Parent to be an explicit nil

### UnsetParent
`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfConfig) UnsetParent()`

UnsetParent ensures that no value is present for Parent, not even an explicit nil
### GetBillingAccount

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfConfig) GetBillingAccount() string`

GetBillingAccount returns the BillingAccount field if non-nil, zero value otherwise.

### GetBillingAccountOk

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfConfig) GetBillingAccountOk() (*string, bool)`

GetBillingAccountOk returns a tuple with the BillingAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingAccount

`func (o *UpdateCloudResourcePool200ResponseResourcePoolAllOfConfig) SetBillingAccount(v string)`

SetBillingAccount sets BillingAccount field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


