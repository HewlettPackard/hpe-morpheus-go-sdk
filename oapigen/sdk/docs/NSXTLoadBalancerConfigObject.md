# NSXTLoadBalancerConfigObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AdminState** | Pointer to **bool** | If true, the admin state is enabled. | [optional] 
**Size** | Pointer to **string** | Load balancer size. | [optional] 
**Tier1** | Pointer to **string** | Tier 1 gateway provider ID. | [optional] 
**Loglevel** | Pointer to **string** | Log level. | [optional] 

## Methods

### NewNSXTLoadBalancerConfigObject

`func NewNSXTLoadBalancerConfigObject() *NSXTLoadBalancerConfigObject`

NewNSXTLoadBalancerConfigObject instantiates a new NSXTLoadBalancerConfigObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetAdminState

`func (o *NSXTLoadBalancerConfigObject) GetAdminState() bool`

GetAdminState returns the AdminState field if non-nil, zero value otherwise.

### GetAdminStateOk

`func (o *NSXTLoadBalancerConfigObject) GetAdminStateOk() (*bool, bool)`

GetAdminStateOk returns a tuple with the AdminState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminState

`func (o *NSXTLoadBalancerConfigObject) SetAdminState(v bool)`

SetAdminState sets AdminState field to given value.

### HasAdminState

`func (o *NSXTLoadBalancerConfigObject) HasAdminState() bool`

HasAdminState returns a boolean if a field has been set.

### GetSize

`func (o *NSXTLoadBalancerConfigObject) GetSize() string`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *NSXTLoadBalancerConfigObject) GetSizeOk() (*string, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *NSXTLoadBalancerConfigObject) SetSize(v string)`

SetSize sets Size field to given value.

### HasSize

`func (o *NSXTLoadBalancerConfigObject) HasSize() bool`

HasSize returns a boolean if a field has been set.

### GetTier1

`func (o *NSXTLoadBalancerConfigObject) GetTier1() string`

GetTier1 returns the Tier1 field if non-nil, zero value otherwise.

### GetTier1Ok

`func (o *NSXTLoadBalancerConfigObject) GetTier1Ok() (*string, bool)`

GetTier1Ok returns a tuple with the Tier1 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTier1

`func (o *NSXTLoadBalancerConfigObject) SetTier1(v string)`

SetTier1 sets Tier1 field to given value.

### HasTier1

`func (o *NSXTLoadBalancerConfigObject) HasTier1() bool`

HasTier1 returns a boolean if a field has been set.

### GetLoglevel

`func (o *NSXTLoadBalancerConfigObject) GetLoglevel() string`

GetLoglevel returns the Loglevel field if non-nil, zero value otherwise.

### GetLoglevelOk

`func (o *NSXTLoadBalancerConfigObject) GetLoglevelOk() (*string, bool)`

GetLoglevelOk returns a tuple with the Loglevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoglevel

`func (o *NSXTLoadBalancerConfigObject) SetLoglevel(v string)`

SetLoglevel sets Loglevel field to given value.

### HasLoglevel

`func (o *NSXTLoadBalancerConfigObject) HasLoglevel() bool`

HasLoglevel returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


