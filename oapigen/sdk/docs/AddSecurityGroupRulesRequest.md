# AddSecurityGroupRulesRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rule** | [**AddSecurityGroupRulesRequestRule**](AddSecurityGroupRulesRequestRule.md) |  | 

## Methods

### NewAddSecurityGroupRulesRequest

`func NewAddSecurityGroupRulesRequest(rule AddSecurityGroupRulesRequestRule, ) *AddSecurityGroupRulesRequest`

NewAddSecurityGroupRulesRequest instantiates a new AddSecurityGroupRulesRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddSecurityGroupRulesRequestWithDefaults

`func NewAddSecurityGroupRulesRequestWithDefaults() *AddSecurityGroupRulesRequest`

NewAddSecurityGroupRulesRequestWithDefaults instantiates a new AddSecurityGroupRulesRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRule

`func (o *AddSecurityGroupRulesRequest) GetRule() AddSecurityGroupRulesRequestRule`

GetRule returns the Rule field if non-nil, zero value otherwise.

### GetRuleOk

`func (o *AddSecurityGroupRulesRequest) GetRuleOk() (*AddSecurityGroupRulesRequestRule, bool)`

GetRuleOk returns a tuple with the Rule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRule

`func (o *AddSecurityGroupRulesRequest) SetRule(v AddSecurityGroupRulesRequestRule)`

SetRule sets Rule field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


