# SetInstanceSecurityGroups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityGroups** | Pointer to [**[]SetInstanceSecurityGroups200ResponseAllOfSecurityGroupsInner**](SetInstanceSecurityGroups200ResponseAllOfSecurityGroupsInner.md) | Array of security group objects | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewSetInstanceSecurityGroups200Response

`func NewSetInstanceSecurityGroups200Response() *SetInstanceSecurityGroups200Response`

NewSetInstanceSecurityGroups200Response instantiates a new SetInstanceSecurityGroups200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetInstanceSecurityGroups200ResponseWithDefaults

`func NewSetInstanceSecurityGroups200ResponseWithDefaults() *SetInstanceSecurityGroups200Response`

NewSetInstanceSecurityGroups200ResponseWithDefaults instantiates a new SetInstanceSecurityGroups200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityGroups

`func (o *SetInstanceSecurityGroups200Response) GetSecurityGroups() []SetInstanceSecurityGroups200ResponseAllOfSecurityGroupsInner`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *SetInstanceSecurityGroups200Response) GetSecurityGroupsOk() (*[]SetInstanceSecurityGroups200ResponseAllOfSecurityGroupsInner, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *SetInstanceSecurityGroups200Response) SetSecurityGroups(v []SetInstanceSecurityGroups200ResponseAllOfSecurityGroupsInner)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *SetInstanceSecurityGroups200Response) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.

### GetSuccess

`func (o *SetInstanceSecurityGroups200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *SetInstanceSecurityGroups200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *SetInstanceSecurityGroups200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *SetInstanceSecurityGroups200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


