# ListSecurityGroupRules200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to [**[]ListSecurityGroupRules200ResponseAllOfRulesInner**](ListSecurityGroupRules200ResponseAllOfRulesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListSecurityGroupRules200Response

`func NewListSecurityGroupRules200Response() *ListSecurityGroupRules200Response`

NewListSecurityGroupRules200Response instantiates a new ListSecurityGroupRules200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSecurityGroupRules200ResponseWithDefaults

`func NewListSecurityGroupRules200ResponseWithDefaults() *ListSecurityGroupRules200Response`

NewListSecurityGroupRules200ResponseWithDefaults instantiates a new ListSecurityGroupRules200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *ListSecurityGroupRules200Response) GetRules() []ListSecurityGroupRules200ResponseAllOfRulesInner`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *ListSecurityGroupRules200Response) GetRulesOk() (*[]ListSecurityGroupRules200ResponseAllOfRulesInner, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *ListSecurityGroupRules200Response) SetRules(v []ListSecurityGroupRules200ResponseAllOfRulesInner)`

SetRules sets Rules field to given value.

### HasRules

`func (o *ListSecurityGroupRules200Response) HasRules() bool`

HasRules returns a boolean if a field has been set.

### GetMeta

`func (o *ListSecurityGroupRules200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListSecurityGroupRules200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListSecurityGroupRules200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListSecurityGroupRules200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


