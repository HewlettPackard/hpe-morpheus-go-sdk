# GetNetworkFirewallRules200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Rules** | Pointer to **interface{}** |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewGetNetworkFirewallRules200Response

`func NewGetNetworkFirewallRules200Response() *GetNetworkFirewallRules200Response`

NewGetNetworkFirewallRules200Response instantiates a new GetNetworkFirewallRules200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkFirewallRules200ResponseWithDefaults

`func NewGetNetworkFirewallRules200ResponseWithDefaults() *GetNetworkFirewallRules200Response`

NewGetNetworkFirewallRules200ResponseWithDefaults instantiates a new GetNetworkFirewallRules200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRules

`func (o *GetNetworkFirewallRules200Response) GetRules() interface{}`

GetRules returns the Rules field if non-nil, zero value otherwise.

### GetRulesOk

`func (o *GetNetworkFirewallRules200Response) GetRulesOk() (*interface{}, bool)`

GetRulesOk returns a tuple with the Rules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRules

`func (o *GetNetworkFirewallRules200Response) SetRules(v interface{})`

SetRules sets Rules field to given value.

### HasRules

`func (o *GetNetworkFirewallRules200Response) HasRules() bool`

HasRules returns a boolean if a field has been set.

### SetRulesNil

`func (o *GetNetworkFirewallRules200Response) SetRulesNil(b bool)`

 SetRulesNil sets the value for Rules to be an explicit nil

### UnsetRules
`func (o *GetNetworkFirewallRules200Response) UnsetRules()`

UnsetRules ensures that no value is present for Rules, not even an explicit nil
### GetMeta

`func (o *GetNetworkFirewallRules200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetNetworkFirewallRules200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetNetworkFirewallRules200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetNetworkFirewallRules200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


