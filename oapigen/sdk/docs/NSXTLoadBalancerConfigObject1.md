# NSXTLoadBalancerConfigObject1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminState** | Pointer to **bool** | If true, the admin state is enabled. | [optional] 
**Size** | Pointer to **string** | Load balancer size. | [optional] 
**Tier1** | Pointer to **string** | Tier 1 gateway provider ID. | [optional] 
**Loglevel** | Pointer to **string** | Log level. | [optional] 

## Methods

### NewNSXTLoadBalancerConfigObject1

`func NewNSXTLoadBalancerConfigObject1() *NSXTLoadBalancerConfigObject1`

NewNSXTLoadBalancerConfigObject1 instantiates a new NSXTLoadBalancerConfigObject1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNSXTLoadBalancerConfigObject1WithDefaults

`func NewNSXTLoadBalancerConfigObject1WithDefaults() *NSXTLoadBalancerConfigObject1`

NewNSXTLoadBalancerConfigObject1WithDefaults instantiates a new NSXTLoadBalancerConfigObject1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdminState

`func (o *NSXTLoadBalancerConfigObject1) GetAdminState() bool`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *NSXTLoadBalancerConfigObject1) GetAdminStateOk() (*bool, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *NSXTLoadBalancerConfigObject1) SetAdminState(v bool)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *NSXTLoadBalancerConfigObject1) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetSize

`func (o *NSXTLoadBalancerConfigObject1) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *NSXTLoadBalancerConfigObject1) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *NSXTLoadBalancerConfigObject1) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *NSXTLoadBalancerConfigObject1) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTier1

`func (o *NSXTLoadBalancerConfigObject1) GetTier1() string`

GetTier1 returns the Tier1 field if non-nil, zero value otherwise.

### GetTier1Ok

`func (o *NSXTLoadBalancerConfigObject1) GetTier1Ok() (*string, bool)`

GetTier1Ok returns a tuple with the Tier1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier1

`func (o *NSXTLoadBalancerConfigObject1) SetTier1(v string)`

SetTier1 sets Tier1 field to given value.

### HasTier1

`func (o *NSXTLoadBalancerConfigObject1) HasTier1() bool`

HasTier1 returns a boolean if a field has been set.

### GetLoglevel

`func (o *NSXTLoadBalancerConfigObject1) GetLoglevel() string`

GetLoglevel returns the Loglevel field if non-nil, zero value otherwise.

### GetLoglevelOk

`func (o *NSXTLoadBalancerConfigObject1) GetLoglevelOk() (*string, bool)`

GetLoglevelOk returns a tuple with the Loglevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoglevel

`func (o *NSXTLoadBalancerConfigObject1) SetLoglevel(v string)`

SetLoglevel sets Loglevel field to given value.

### HasLoglevel

`func (o *NSXTLoadBalancerConfigObject1) HasLoglevel() bool`

HasLoglevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


