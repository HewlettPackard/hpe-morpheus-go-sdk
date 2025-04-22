# GetNetworkRouterType200ResponseNetworkRouterType

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
**OptionTypes** | Pointer to [**[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner**](ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner.md) |  | [optional] 
**RuleOptionTypes** | Pointer to [**[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner**](ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner.md) |  | [optional] 
**RuleGroupOptionTypes** | Pointer to [**[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner**](ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner.md) |  | [optional] 
**NatOptionTypes** | Pointer to [**[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner**](ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner.md) |  | [optional] 
**BgpOptionTypes** | Pointer to [**[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner**](ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner.md) |  | [optional] 

## Methods

### NewGetNetworkRouterType200ResponseNetworkRouterType

`func NewGetNetworkRouterType200ResponseNetworkRouterType() *GetNetworkRouterType200ResponseNetworkRouterType`

NewGetNetworkRouterType200ResponseNetworkRouterType instantiates a new GetNetworkRouterType200ResponseNetworkRouterType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetNetworkRouterType200ResponseNetworkRouterTypeWithDefaults

`func NewGetNetworkRouterType200ResponseNetworkRouterTypeWithDefaults() *GetNetworkRouterType200ResponseNetworkRouterType`

NewGetNetworkRouterType200ResponseNetworkRouterTypeWithDefaults instantiates a new GetNetworkRouterType200ResponseNetworkRouterType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCreatable

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetSelectable

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) GetSelectable() bool`

GetSelectable returns the Selectable field if non-nil, zero value otherwise.

### GetSelectableOk

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) GetSelectableOk() (*bool, bool)`

GetSelectableOk returns a tuple with the Selectable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectable

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) SetSelectable(v bool)`

SetSelectable sets Selectable field to given value.

### HasSelectable

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) HasSelectable() bool`

HasSelectable returns a boolean if a field has been set.

### GetHasFirewall

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) GetHasFirewall() bool`

GetHasFirewall returns the HasFirewall field if non-nil, zero value otherwise.

### GetHasFirewallOk

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) GetHasFirewallOk() (*bool, bool)`

GetHasFirewallOk returns a tuple with the HasFirewall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasFirewall

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) SetHasFirewall(v bool)`

SetHasFirewall sets HasFirewall field to given value.

### HasHasFirewall

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) HasHasFirewall() bool`

HasHasFirewall returns a boolean if a field has been set.

### GetHasDhcp

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) GetHasDhcp() bool`

GetHasDhcp returns the HasDhcp field if non-nil, zero value otherwise.

### GetHasDhcpOk

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) GetHasDhcpOk() (*bool, bool)`

GetHasDhcpOk returns a tuple with the HasDhcp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasDhcp

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) SetHasDhcp(v bool)`

SetHasDhcp sets HasDhcp field to given value.

### HasHasDhcp

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) HasHasDhcp() bool`

HasHasDhcp returns a boolean if a field has been set.

### GetHasRouting

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) GetHasRouting() bool`

GetHasRouting returns the HasRouting field if non-nil, zero value otherwise.

### GetHasRoutingOk

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) GetHasRoutingOk() (*bool, bool)`

GetHasRoutingOk returns a tuple with the HasRouting field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasRouting

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) SetHasRouting(v bool)`

SetHasRouting sets HasRouting field to given value.

### HasHasRouting

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) HasHasRouting() bool`

HasHasRouting returns a boolean if a field has been set.

### GetHasNetworkServer

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) GetHasNetworkServer() bool`

GetHasNetworkServer returns the HasNetworkServer field if non-nil, zero value otherwise.

### GetHasNetworkServerOk

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) GetHasNetworkServerOk() (*bool, bool)`

GetHasNetworkServerOk returns a tuple with the HasNetworkServer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasNetworkServer

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) SetHasNetworkServer(v bool)`

SetHasNetworkServer sets HasNetworkServer field to given value.

### HasHasNetworkServer

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) HasHasNetworkServer() bool`

HasHasNetworkServer returns a boolean if a field has been set.

### GetOptionTypes

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) GetOptionTypes() []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) GetOptionTypesOk() (*[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) SetOptionTypes(v []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetRuleOptionTypes

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) GetRuleOptionTypes() []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner`

GetRuleOptionTypes returns the RuleOptionTypes field if non-nil, zero value otherwise.

### GetRuleOptionTypesOk

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) GetRuleOptionTypesOk() (*[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner, bool)`

GetRuleOptionTypesOk returns a tuple with the RuleOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleOptionTypes

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) SetRuleOptionTypes(v []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner)`

SetRuleOptionTypes sets RuleOptionTypes field to given value.

### HasRuleOptionTypes

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) HasRuleOptionTypes() bool`

HasRuleOptionTypes returns a boolean if a field has been set.

### GetRuleGroupOptionTypes

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) GetRuleGroupOptionTypes() []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner`

GetRuleGroupOptionTypes returns the RuleGroupOptionTypes field if non-nil, zero value otherwise.

### GetRuleGroupOptionTypesOk

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) GetRuleGroupOptionTypesOk() (*[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner, bool)`

GetRuleGroupOptionTypesOk returns a tuple with the RuleGroupOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleGroupOptionTypes

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) SetRuleGroupOptionTypes(v []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner)`

SetRuleGroupOptionTypes sets RuleGroupOptionTypes field to given value.

### HasRuleGroupOptionTypes

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) HasRuleGroupOptionTypes() bool`

HasRuleGroupOptionTypes returns a boolean if a field has been set.

### GetNatOptionTypes

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) GetNatOptionTypes() []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner`

GetNatOptionTypes returns the NatOptionTypes field if non-nil, zero value otherwise.

### GetNatOptionTypesOk

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) GetNatOptionTypesOk() (*[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner, bool)`

GetNatOptionTypesOk returns a tuple with the NatOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNatOptionTypes

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) SetNatOptionTypes(v []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner)`

SetNatOptionTypes sets NatOptionTypes field to given value.

### HasNatOptionTypes

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) HasNatOptionTypes() bool`

HasNatOptionTypes returns a boolean if a field has been set.

### GetBgpOptionTypes

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) GetBgpOptionTypes() []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner`

GetBgpOptionTypes returns the BgpOptionTypes field if non-nil, zero value otherwise.

### GetBgpOptionTypesOk

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) GetBgpOptionTypesOk() (*[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner, bool)`

GetBgpOptionTypesOk returns a tuple with the BgpOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBgpOptionTypes

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) SetBgpOptionTypes(v []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner)`

SetBgpOptionTypes sets BgpOptionTypes field to given value.

### HasBgpOptionTypes

`func (o *GetNetworkRouterType200ResponseNetworkRouterType) HasBgpOptionTypes() bool`

HasBgpOptionTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


