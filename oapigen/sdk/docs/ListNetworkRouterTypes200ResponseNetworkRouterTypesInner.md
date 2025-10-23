# ListNetworkRouterTypes200ResponseNetworkRouterTypesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Creatable** | Pointer to **bool** |  | [optional] 
**Selectable** | Pointer to **bool** |  | [optional] 
**HasFirewall** | Pointer to **bool** |  | [optional] 
**HasDhcp** | Pointer to **bool** |  | [optional] 
**HasRouting** | Pointer to **bool** |  | [optional] 
**HasNetworkServer** | Pointer to **bool** |  | [optional] 
**OptionTypes** | Pointer to [**[]ListNetworkRouterTypes200ResponseNetworkRouterTypesInnerOptionTypesInner**](ListNetworkRouterTypes200ResponseNetworkRouterTypesInnerOptionTypesInner.md) |  | [optional] 
**RuleOptionTypes** | Pointer to **[]map[string]interface{}** |  | [optional] 
**NatOptionTypes** | Pointer to **[]map[string]interface{}** |  | [optional] 
**RuleGroupOptionTypes** | Pointer to **[]map[string]interface{}** |  | [optional] 

## Methods

### NewListNetworkRouterTypes200ResponseNetworkRouterTypesInner

`func NewListNetworkRouterTypes200ResponseNetworkRouterTypesInner() *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner`

NewListNetworkRouterTypes200ResponseNetworkRouterTypesInner instantiates a new ListNetworkRouterTypes200ResponseNetworkRouterTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListNetworkRouterTypes200ResponseNetworkRouterTypesInnerWithDefaults

`func NewListNetworkRouterTypes200ResponseNetworkRouterTypesInnerWithDefaults() *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner`

NewListNetworkRouterTypes200ResponseNetworkRouterTypesInnerWithDefaults instantiates a new ListNetworkRouterTypes200ResponseNetworkRouterTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCreatable

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetSelectable

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) GetSelectable() bool`

GetSelectable returns the Selectable field if non-nil, zero value otherwise.

### GetSelectableOk

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) GetSelectableOk() (*bool, bool)`

GetSelectableOk returns a tuple with the Selectable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectable

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) SetSelectable(v bool)`

SetSelectable sets Selectable field to given value.

### HasSelectable

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) HasSelectable() bool`

HasSelectable returns a boolean if a field has been set.

### GetHasFirewall

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) GetHasFirewall() bool`

GetHasFirewall returns the HasFirewall field if non-nil, zero value otherwise.

### GetHasFirewallOk

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) GetHasFirewallOk() (*bool, bool)`

GetHasFirewallOk returns a tuple with the HasFirewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFirewall

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) SetHasFirewall(v bool)`

SetHasFirewall sets HasFirewall field to given value.

### HasHasFirewall

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) HasHasFirewall() bool`

HasHasFirewall returns a boolean if a field has been set.

### GetHasDhcp

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) GetHasDhcp() bool`

GetHasDhcp returns the HasDhcp field if non-nil, zero value otherwise.

### GetHasDhcpOk

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) GetHasDhcpOk() (*bool, bool)`

GetHasDhcpOk returns a tuple with the HasDhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDhcp

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) SetHasDhcp(v bool)`

SetHasDhcp sets HasDhcp field to given value.

### HasHasDhcp

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) HasHasDhcp() bool`

HasHasDhcp returns a boolean if a field has been set.

### GetHasRouting

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) GetHasRouting() bool`

GetHasRouting returns the HasRouting field if non-nil, zero value otherwise.

### GetHasRoutingOk

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) GetHasRoutingOk() (*bool, bool)`

GetHasRoutingOk returns a tuple with the HasRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRouting

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) SetHasRouting(v bool)`

SetHasRouting sets HasRouting field to given value.

### HasHasRouting

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) HasHasRouting() bool`

HasHasRouting returns a boolean if a field has been set.

### GetHasNetworkServer

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) GetHasNetworkServer() bool`

GetHasNetworkServer returns the HasNetworkServer field if non-nil, zero value otherwise.

### GetHasNetworkServerOk

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) GetHasNetworkServerOk() (*bool, bool)`

GetHasNetworkServerOk returns a tuple with the HasNetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNetworkServer

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) SetHasNetworkServer(v bool)`

SetHasNetworkServer sets HasNetworkServer field to given value.

### HasHasNetworkServer

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) HasHasNetworkServer() bool`

HasHasNetworkServer returns a boolean if a field has been set.

### GetOptionTypes

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) GetOptionTypes() []ListNetworkRouterTypes200ResponseNetworkRouterTypesInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) GetOptionTypesOk() (*[]ListNetworkRouterTypes200ResponseNetworkRouterTypesInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) SetOptionTypes(v []ListNetworkRouterTypes200ResponseNetworkRouterTypesInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetRuleOptionTypes

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) GetRuleOptionTypes() []map[string]interface{}`

GetRuleOptionTypes returns the RuleOptionTypes field if non-nil, zero value otherwise.

### GetRuleOptionTypesOk

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) GetRuleOptionTypesOk() (*[]map[string]interface{}, bool)`

GetRuleOptionTypesOk returns a tuple with the RuleOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleOptionTypes

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) SetRuleOptionTypes(v []map[string]interface{})`

SetRuleOptionTypes sets RuleOptionTypes field to given value.

### HasRuleOptionTypes

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) HasRuleOptionTypes() bool`

HasRuleOptionTypes returns a boolean if a field has been set.

### GetNatOptionTypes

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) GetNatOptionTypes() []map[string]interface{}`

GetNatOptionTypes returns the NatOptionTypes field if non-nil, zero value otherwise.

### GetNatOptionTypesOk

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) GetNatOptionTypesOk() (*[]map[string]interface{}, bool)`

GetNatOptionTypesOk returns a tuple with the NatOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatOptionTypes

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) SetNatOptionTypes(v []map[string]interface{})`

SetNatOptionTypes sets NatOptionTypes field to given value.

### HasNatOptionTypes

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) HasNatOptionTypes() bool`

HasNatOptionTypes returns a boolean if a field has been set.

### GetRuleGroupOptionTypes

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) GetRuleGroupOptionTypes() []map[string]interface{}`

GetRuleGroupOptionTypes returns the RuleGroupOptionTypes field if non-nil, zero value otherwise.

### GetRuleGroupOptionTypesOk

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) GetRuleGroupOptionTypesOk() (*[]map[string]interface{}, bool)`

GetRuleGroupOptionTypesOk returns a tuple with the RuleGroupOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleGroupOptionTypes

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) SetRuleGroupOptionTypes(v []map[string]interface{})`

SetRuleGroupOptionTypes sets RuleGroupOptionTypes field to given value.

### HasRuleGroupOptionTypes

`func (o *ListNetworkRouterTypes200ResponseNetworkRouterTypesInner) HasRuleGroupOptionTypes() bool`

HasRuleGroupOptionTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


