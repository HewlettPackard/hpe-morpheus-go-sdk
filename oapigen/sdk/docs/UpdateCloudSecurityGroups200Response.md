# UpdateCloudSecurityGroups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**SecurityGroups** | Pointer to [**[]ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner**](ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner.md) |  | [optional] 

## Methods

### NewUpdateCloudSecurityGroups200Response

`func NewUpdateCloudSecurityGroups200Response() *UpdateCloudSecurityGroups200Response`

NewUpdateCloudSecurityGroups200Response instantiates a new UpdateCloudSecurityGroups200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCloudSecurityGroups200ResponseWithDefaults

`func NewUpdateCloudSecurityGroups200ResponseWithDefaults() *UpdateCloudSecurityGroups200Response`

NewUpdateCloudSecurityGroups200ResponseWithDefaults instantiates a new UpdateCloudSecurityGroups200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *UpdateCloudSecurityGroups200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateCloudSecurityGroups200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateCloudSecurityGroups200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateCloudSecurityGroups200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetSecurityGroups

`func (o *UpdateCloudSecurityGroups200Response) GetSecurityGroups() []ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner`

GetSecurityGroups returns the SecurityGroups field if non-nil, zero value otherwise.

### GetSecurityGroupsOk

`func (o *UpdateCloudSecurityGroups200Response) GetSecurityGroupsOk() (*[]ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner, bool)`

GetSecurityGroupsOk returns a tuple with the SecurityGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroups

`func (o *UpdateCloudSecurityGroups200Response) SetSecurityGroups(v []ListCloudSecurityGroups200ResponseAllOfSecurityGroupsInner)`

SetSecurityGroups sets SecurityGroups field to given value.

### HasSecurityGroups

`func (o *UpdateCloudSecurityGroups200Response) HasSecurityGroups() bool`

HasSecurityGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


