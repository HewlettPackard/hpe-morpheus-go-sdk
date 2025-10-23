# GetNetworkRouter200ResponseNetworkRouterType

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
**HasNat** | Pointer to **bool** |  | [optional] 
**HasNetworkServer** | Pointer to **bool** |  | [optional] 
**HasFirewallGroups** | Pointer to **bool** |  | [optional] 
**HasSecurityGroupPriority** | Pointer to **bool** |  | [optional] 
**OptionTypes** | Pointer to [**[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner**](ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner.md) |  | [optional] 
**RuleOptionTypes** | Pointer to [**[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner**](ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner.md) |  | [optional] 
**FirewallGroupOptionTypes** | Pointer to [**[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner**](ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner.md) |  | [optional] 
**NatOptionTypes** | Pointer to [**[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner**](ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner.md) |  | [optional] 

## Methods

### NewGetNetworkRouter200ResponseNetworkRouterType

`func NewGetNetworkRouter200ResponseNetworkRouterType() *GetNetworkRouter200ResponseNetworkRouterType`

NewGetNetworkRouter200ResponseNetworkRouterType instantiates a new GetNetworkRouter200ResponseNetworkRouterType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkRouter200ResponseNetworkRouterTypeWithDefaults

`func NewGetNetworkRouter200ResponseNetworkRouterTypeWithDefaults() *GetNetworkRouter200ResponseNetworkRouterType`

NewGetNetworkRouter200ResponseNetworkRouterTypeWithDefaults instantiates a new GetNetworkRouter200ResponseNetworkRouterType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkRouter200ResponseNetworkRouterType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkRouter200ResponseNetworkRouterType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkRouter200ResponseNetworkRouterType) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkRouter200ResponseNetworkRouterType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *GetNetworkRouter200ResponseNetworkRouterType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetNetworkRouter200ResponseNetworkRouterType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetNetworkRouter200ResponseNetworkRouterType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetNetworkRouter200ResponseNetworkRouterType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkRouter200ResponseNetworkRouterType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkRouter200ResponseNetworkRouterType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkRouter200ResponseNetworkRouterType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkRouter200ResponseNetworkRouterType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetNetworkRouter200ResponseNetworkRouterType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetNetworkRouter200ResponseNetworkRouterType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetNetworkRouter200ResponseNetworkRouterType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetNetworkRouter200ResponseNetworkRouterType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *GetNetworkRouter200ResponseNetworkRouterType) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetNetworkRouter200ResponseNetworkRouterType) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetNetworkRouter200ResponseNetworkRouterType) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetNetworkRouter200ResponseNetworkRouterType) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCreatable

`func (o *GetNetworkRouter200ResponseNetworkRouterType) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *GetNetworkRouter200ResponseNetworkRouterType) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *GetNetworkRouter200ResponseNetworkRouterType) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *GetNetworkRouter200ResponseNetworkRouterType) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetSelectable

`func (o *GetNetworkRouter200ResponseNetworkRouterType) GetSelectable() bool`

GetSelectable returns the Selectable field if non-nil, zero value otherwise.

### GetSelectableOk

`func (o *GetNetworkRouter200ResponseNetworkRouterType) GetSelectableOk() (*bool, bool)`

GetSelectableOk returns a tuple with the Selectable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectable

`func (o *GetNetworkRouter200ResponseNetworkRouterType) SetSelectable(v bool)`

SetSelectable sets Selectable field to given value.

### HasSelectable

`func (o *GetNetworkRouter200ResponseNetworkRouterType) HasSelectable() bool`

HasSelectable returns a boolean if a field has been set.

### GetHasFirewall

`func (o *GetNetworkRouter200ResponseNetworkRouterType) GetHasFirewall() bool`

GetHasFirewall returns the HasFirewall field if non-nil, zero value otherwise.

### GetHasFirewallOk

`func (o *GetNetworkRouter200ResponseNetworkRouterType) GetHasFirewallOk() (*bool, bool)`

GetHasFirewallOk returns a tuple with the HasFirewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFirewall

`func (o *GetNetworkRouter200ResponseNetworkRouterType) SetHasFirewall(v bool)`

SetHasFirewall sets HasFirewall field to given value.

### HasHasFirewall

`func (o *GetNetworkRouter200ResponseNetworkRouterType) HasHasFirewall() bool`

HasHasFirewall returns a boolean if a field has been set.

### GetHasDhcp

`func (o *GetNetworkRouter200ResponseNetworkRouterType) GetHasDhcp() bool`

GetHasDhcp returns the HasDhcp field if non-nil, zero value otherwise.

### GetHasDhcpOk

`func (o *GetNetworkRouter200ResponseNetworkRouterType) GetHasDhcpOk() (*bool, bool)`

GetHasDhcpOk returns a tuple with the HasDhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDhcp

`func (o *GetNetworkRouter200ResponseNetworkRouterType) SetHasDhcp(v bool)`

SetHasDhcp sets HasDhcp field to given value.

### HasHasDhcp

`func (o *GetNetworkRouter200ResponseNetworkRouterType) HasHasDhcp() bool`

HasHasDhcp returns a boolean if a field has been set.

### GetHasRouting

`func (o *GetNetworkRouter200ResponseNetworkRouterType) GetHasRouting() bool`

GetHasRouting returns the HasRouting field if non-nil, zero value otherwise.

### GetHasRoutingOk

`func (o *GetNetworkRouter200ResponseNetworkRouterType) GetHasRoutingOk() (*bool, bool)`

GetHasRoutingOk returns a tuple with the HasRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRouting

`func (o *GetNetworkRouter200ResponseNetworkRouterType) SetHasRouting(v bool)`

SetHasRouting sets HasRouting field to given value.

### HasHasRouting

`func (o *GetNetworkRouter200ResponseNetworkRouterType) HasHasRouting() bool`

HasHasRouting returns a boolean if a field has been set.

### GetHasNat

`func (o *GetNetworkRouter200ResponseNetworkRouterType) GetHasNat() bool`

GetHasNat returns the HasNat field if non-nil, zero value otherwise.

### GetHasNatOk

`func (o *GetNetworkRouter200ResponseNetworkRouterType) GetHasNatOk() (*bool, bool)`

GetHasNatOk returns a tuple with the HasNat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNat

`func (o *GetNetworkRouter200ResponseNetworkRouterType) SetHasNat(v bool)`

SetHasNat sets HasNat field to given value.

### HasHasNat

`func (o *GetNetworkRouter200ResponseNetworkRouterType) HasHasNat() bool`

HasHasNat returns a boolean if a field has been set.

### GetHasNetworkServer

`func (o *GetNetworkRouter200ResponseNetworkRouterType) GetHasNetworkServer() bool`

GetHasNetworkServer returns the HasNetworkServer field if non-nil, zero value otherwise.

### GetHasNetworkServerOk

`func (o *GetNetworkRouter200ResponseNetworkRouterType) GetHasNetworkServerOk() (*bool, bool)`

GetHasNetworkServerOk returns a tuple with the HasNetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNetworkServer

`func (o *GetNetworkRouter200ResponseNetworkRouterType) SetHasNetworkServer(v bool)`

SetHasNetworkServer sets HasNetworkServer field to given value.

### HasHasNetworkServer

`func (o *GetNetworkRouter200ResponseNetworkRouterType) HasHasNetworkServer() bool`

HasHasNetworkServer returns a boolean if a field has been set.

### GetHasFirewallGroups

`func (o *GetNetworkRouter200ResponseNetworkRouterType) GetHasFirewallGroups() bool`

GetHasFirewallGroups returns the HasFirewallGroups field if non-nil, zero value otherwise.

### GetHasFirewallGroupsOk

`func (o *GetNetworkRouter200ResponseNetworkRouterType) GetHasFirewallGroupsOk() (*bool, bool)`

GetHasFirewallGroupsOk returns a tuple with the HasFirewallGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFirewallGroups

`func (o *GetNetworkRouter200ResponseNetworkRouterType) SetHasFirewallGroups(v bool)`

SetHasFirewallGroups sets HasFirewallGroups field to given value.

### HasHasFirewallGroups

`func (o *GetNetworkRouter200ResponseNetworkRouterType) HasHasFirewallGroups() bool`

HasHasFirewallGroups returns a boolean if a field has been set.

### GetHasSecurityGroupPriority

`func (o *GetNetworkRouter200ResponseNetworkRouterType) GetHasSecurityGroupPriority() bool`

GetHasSecurityGroupPriority returns the HasSecurityGroupPriority field if non-nil, zero value otherwise.

### GetHasSecurityGroupPriorityOk

`func (o *GetNetworkRouter200ResponseNetworkRouterType) GetHasSecurityGroupPriorityOk() (*bool, bool)`

GetHasSecurityGroupPriorityOk returns a tuple with the HasSecurityGroupPriority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasSecurityGroupPriority

`func (o *GetNetworkRouter200ResponseNetworkRouterType) SetHasSecurityGroupPriority(v bool)`

SetHasSecurityGroupPriority sets HasSecurityGroupPriority field to given value.

### HasHasSecurityGroupPriority

`func (o *GetNetworkRouter200ResponseNetworkRouterType) HasHasSecurityGroupPriority() bool`

HasHasSecurityGroupPriority returns a boolean if a field has been set.

### GetOptionTypes

`func (o *GetNetworkRouter200ResponseNetworkRouterType) GetOptionTypes() []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *GetNetworkRouter200ResponseNetworkRouterType) GetOptionTypesOk() (*[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *GetNetworkRouter200ResponseNetworkRouterType) SetOptionTypes(v []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *GetNetworkRouter200ResponseNetworkRouterType) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetRuleOptionTypes

`func (o *GetNetworkRouter200ResponseNetworkRouterType) GetRuleOptionTypes() []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner`

GetRuleOptionTypes returns the RuleOptionTypes field if non-nil, zero value otherwise.

### GetRuleOptionTypesOk

`func (o *GetNetworkRouter200ResponseNetworkRouterType) GetRuleOptionTypesOk() (*[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner, bool)`

GetRuleOptionTypesOk returns a tuple with the RuleOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleOptionTypes

`func (o *GetNetworkRouter200ResponseNetworkRouterType) SetRuleOptionTypes(v []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner)`

SetRuleOptionTypes sets RuleOptionTypes field to given value.

### HasRuleOptionTypes

`func (o *GetNetworkRouter200ResponseNetworkRouterType) HasRuleOptionTypes() bool`

HasRuleOptionTypes returns a boolean if a field has been set.

### GetFirewallGroupOptionTypes

`func (o *GetNetworkRouter200ResponseNetworkRouterType) GetFirewallGroupOptionTypes() []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner`

GetFirewallGroupOptionTypes returns the FirewallGroupOptionTypes field if non-nil, zero value otherwise.

### GetFirewallGroupOptionTypesOk

`func (o *GetNetworkRouter200ResponseNetworkRouterType) GetFirewallGroupOptionTypesOk() (*[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner, bool)`

GetFirewallGroupOptionTypesOk returns a tuple with the FirewallGroupOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirewallGroupOptionTypes

`func (o *GetNetworkRouter200ResponseNetworkRouterType) SetFirewallGroupOptionTypes(v []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner)`

SetFirewallGroupOptionTypes sets FirewallGroupOptionTypes field to given value.

### HasFirewallGroupOptionTypes

`func (o *GetNetworkRouter200ResponseNetworkRouterType) HasFirewallGroupOptionTypes() bool`

HasFirewallGroupOptionTypes returns a boolean if a field has been set.

### GetNatOptionTypes

`func (o *GetNetworkRouter200ResponseNetworkRouterType) GetNatOptionTypes() []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner`

GetNatOptionTypes returns the NatOptionTypes field if non-nil, zero value otherwise.

### GetNatOptionTypesOk

`func (o *GetNetworkRouter200ResponseNetworkRouterType) GetNatOptionTypesOk() (*[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner, bool)`

GetNatOptionTypesOk returns a tuple with the NatOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatOptionTypes

`func (o *GetNetworkRouter200ResponseNetworkRouterType) SetNatOptionTypes(v []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner)`

SetNatOptionTypes sets NatOptionTypes field to given value.

### HasNatOptionTypes

`func (o *GetNetworkRouter200ResponseNetworkRouterType) HasNatOptionTypes() bool`

HasNatOptionTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


