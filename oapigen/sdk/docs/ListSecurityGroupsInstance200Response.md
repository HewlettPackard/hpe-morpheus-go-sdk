# ListSecurityGroupsInstance200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirewallEnabled** | Pointer to **bool** |  | [optional] 
**SecurityGroups** | Pointer to [**[]ListSecurityGroupsInstance200ResponseAllOfSecurityGroupsInner**](ListSecurityGroupsInstance200ResponseAllOfSecurityGroupsInner.md) | Array of security group objects | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewListSecurityGroupsInstance200Response

`func NewListSecurityGroupsInstance200Response() *ListSecurityGroupsInstance200Response`

NewListSecurityGroupsInstance200Response instantiates a new ListSecurityGroupsInstance200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSecurityGroupsInstance200ResponseWithDefaults

`func NewListSecurityGroupsInstance200ResponseWithDefaults() *ListSecurityGroupsInstance200Response`

NewListSecurityGroupsInstance200ResponseWithDefaults instantiates a new ListSecurityGroupsInstance200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirewallEnabled

`func (o *ListSecurityGroupsInstance200Response) GetFirewallEnabled() bool`

GetFirewallEnabled returns the FirewallEnabled field if non-nil, zero value otherwise.

### GetFirewallEnabledOk

`func (o *ListSecurityGroupsInstance200Response) GetFirewallEnabledOk() (*bool, bool)`

GetFirewallEnabledOk returns a tuple with the FirewallEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallEnabled

`func (o *ListSecurityGroupsInstance200Response) SetFirewallEnabled(v bool)`

SetFirewallEnabled sets FirewallEnabled field to given value.

### HasFirewallEnabled

`func (o *ListSecurityGroupsInstance200Response) HasFirewallEnabled() bool`

HasFirewallEnabled returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *ListSecurityGroupsInstance200Response) GetSecurityGroups() []ListSecurityGroupsInstance200ResponseAllOfSecurityGroupsInner`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *ListSecurityGroupsInstance200Response) GetSecurityGroupsOk() (*[]ListSecurityGroupsInstance200ResponseAllOfSecurityGroupsInner, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *ListSecurityGroupsInstance200Response) SetSecurityGroups(v []ListSecurityGroupsInstance200ResponseAllOfSecurityGroupsInner)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *ListSecurityGroupsInstance200Response) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetSuccess

`func (o *ListSecurityGroupsInstance200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ListSecurityGroupsInstance200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ListSecurityGroupsInstance200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ListSecurityGroupsInstance200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


