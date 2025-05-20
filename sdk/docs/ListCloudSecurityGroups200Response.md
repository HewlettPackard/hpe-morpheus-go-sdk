# ListCloudSecurityGroups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**FirewallEnabled** | Pointer to **bool** |  | [optional] 
**SecurityGroups** | Pointer to [**[]ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner**](ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner.md) |  | [optional] 

## Methods

### NewListCloudSecurityGroups200Response

`func NewListCloudSecurityGroups200Response() *ListCloudSecurityGroups200Response`

NewListCloudSecurityGroups200Response instantiates a new ListCloudSecurityGroups200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCloudSecurityGroups200ResponseWithDefaults

`func NewListCloudSecurityGroups200ResponseWithDefaults() *ListCloudSecurityGroups200Response`

NewListCloudSecurityGroups200ResponseWithDefaults instantiates a new ListCloudSecurityGroups200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ListCloudSecurityGroups200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ListCloudSecurityGroups200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ListCloudSecurityGroups200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ListCloudSecurityGroups200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetFirewallEnabled

`func (o *ListCloudSecurityGroups200Response) GetFirewallEnabled() bool`

GetFirewallEnabled returns the FirewallEnabled field if non-nil, zero value otherwise.

### GetFirewallEnabledOk

`func (o *ListCloudSecurityGroups200Response) GetFirewallEnabledOk() (*bool, bool)`

GetFirewallEnabledOk returns a tuple with the FirewallEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallEnabled

`func (o *ListCloudSecurityGroups200Response) SetFirewallEnabled(v bool)`

SetFirewallEnabled sets FirewallEnabled field to given value.

### HasFirewallEnabled

`func (o *ListCloudSecurityGroups200Response) HasFirewallEnabled() bool`

HasFirewallEnabled returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *ListCloudSecurityGroups200Response) GetSecurityGroups() []ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *ListCloudSecurityGroups200Response) GetSecurityGroupsOk() (*[]ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *ListCloudSecurityGroups200Response) SetSecurityGroups(v []ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *ListCloudSecurityGroups200Response) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


