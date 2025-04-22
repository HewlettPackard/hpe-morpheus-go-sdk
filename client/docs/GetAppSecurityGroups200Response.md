# GetAppSecurityGroups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityGroups** | Pointer to [**[]GetAppSecurityGroups200ResponseAllOfSecurityGroupsInner**](GetAppSecurityGroups200ResponseAllOfSecurityGroupsInner.md) |  | [optional] 
**FirewallEnabled** | Pointer to **bool** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetAppSecurityGroups200Response

`func NewGetAppSecurityGroups200Response() *GetAppSecurityGroups200Response`

NewGetAppSecurityGroups200Response instantiates a new GetAppSecurityGroups200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetAppSecurityGroups200ResponseWithDefaults

`func NewGetAppSecurityGroups200ResponseWithDefaults() *GetAppSecurityGroups200Response`

NewGetAppSecurityGroups200ResponseWithDefaults instantiates a new GetAppSecurityGroups200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityGroups

`func (o *GetAppSecurityGroups200Response) GetSecurityGroups() []GetAppSecurityGroups200ResponseAllOfSecurityGroupsInner`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *GetAppSecurityGroups200Response) GetSecurityGroupsOk() (*[]GetAppSecurityGroups200ResponseAllOfSecurityGroupsInner, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *GetAppSecurityGroups200Response) SetSecurityGroups(v []GetAppSecurityGroups200ResponseAllOfSecurityGroupsInner)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *GetAppSecurityGroups200Response) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetFirewallEnabled

`func (o *GetAppSecurityGroups200Response) GetFirewallEnabled() bool`

GetFirewallEnabled returns the FirewallEnabled field if non-nil, zero value otherwise.

### GetFirewallEnabledOk

`func (o *GetAppSecurityGroups200Response) GetFirewallEnabledOk() (*bool, bool)`

GetFirewallEnabledOk returns a tuple with the FirewallEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallEnabled

`func (o *GetAppSecurityGroups200Response) SetFirewallEnabled(v bool)`

SetFirewallEnabled sets FirewallEnabled field to given value.

### HasFirewallEnabled

`func (o *GetAppSecurityGroups200Response) HasFirewallEnabled() bool`

HasFirewallEnabled returns a boolean if a field has been set.

### GetSuccess

`func (o *GetAppSecurityGroups200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetAppSecurityGroups200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetAppSecurityGroups200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetAppSecurityGroups200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


