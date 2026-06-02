# AddSecurityGroupRules200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rule** | Pointer to [**AddSecurityGroupRules200ResponseAllOfRule**](AddSecurityGroupRules200ResponseAllOfRule.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddSecurityGroupRules200Response

`func NewAddSecurityGroupRules200Response() *AddSecurityGroupRules200Response`

NewAddSecurityGroupRules200Response instantiates a new AddSecurityGroupRules200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetRule

`func (o *AddSecurityGroupRules200Response) GetRule() AddSecurityGroupRules200ResponseAllOfRule`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *AddSecurityGroupRules200Response) GetRuleOk() (*AddSecurityGroupRules200ResponseAllOfRule, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *AddSecurityGroupRules200Response) SetRule(v AddSecurityGroupRules200ResponseAllOfRule)`

SetRule sets Rule field to given value.

### HasRule

`func (o *AddSecurityGroupRules200Response) HasRule() bool`

HasRule returns a boolean if a field has been set.

### GetSuccess

`func (o *AddSecurityGroupRules200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddSecurityGroupRules200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddSecurityGroupRules200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddSecurityGroupRules200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


