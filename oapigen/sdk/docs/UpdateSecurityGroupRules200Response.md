# UpdateSecurityGroupRules200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rule** | Pointer to [**UpdateSecurityGroupRules200ResponseAllOfRule**](UpdateSecurityGroupRules200ResponseAllOfRule.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateSecurityGroupRules200Response

`func NewUpdateSecurityGroupRules200Response() *UpdateSecurityGroupRules200Response`

NewUpdateSecurityGroupRules200Response instantiates a new UpdateSecurityGroupRules200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetRule

`func (o *UpdateSecurityGroupRules200Response) GetRule() UpdateSecurityGroupRules200ResponseAllOfRule`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *UpdateSecurityGroupRules200Response) GetRuleOk() (*UpdateSecurityGroupRules200ResponseAllOfRule, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *UpdateSecurityGroupRules200Response) SetRule(v UpdateSecurityGroupRules200ResponseAllOfRule)`

SetRule sets Rule field to given value.

### HasRule

`func (o *UpdateSecurityGroupRules200Response) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateSecurityGroupRules200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateSecurityGroupRules200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateSecurityGroupRules200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateSecurityGroupRules200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


