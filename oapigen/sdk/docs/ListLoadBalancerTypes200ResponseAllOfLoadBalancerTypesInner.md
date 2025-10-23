# ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Internal** | Pointer to **bool** |  | [optional] 
**Creatable** | Pointer to **bool** |  | [optional] 
**CreateType** | Pointer to **string** |  | [optional] 
**OptionTypes** | Pointer to [**[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner**](ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner.md) |  | [optional] 
**VipOptionTypes** | Pointer to [**[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner**](ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner.md) |  | [optional] 

## Methods

### NewListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner

`func NewListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner() *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner`

NewListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner instantiates a new ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerWithDefaults

`func NewListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerWithDefaults() *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner`

NewListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInnerWithDefaults instantiates a new ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetEnabled

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetInternal

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetCreatable

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetCreateType

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner) GetCreateType() string`

GetCreateType returns the CreateType field if non-nil, zero value otherwise.

### GetCreateTypeOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner) GetCreateTypeOk() (*string, bool)`

GetCreateTypeOk returns a tuple with the CreateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateType

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner) SetCreateType(v string)`

SetCreateType sets CreateType field to given value.

### HasCreateType

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner) HasCreateType() bool`

HasCreateType returns a boolean if a field has been set.

### GetOptionTypes

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner) GetOptionTypes() []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner) GetOptionTypesOk() (*[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner) SetOptionTypes(v []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetVipOptionTypes

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner) GetVipOptionTypes() []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner`

GetVipOptionTypes returns the VipOptionTypes field if non-nil, zero value otherwise.

### GetVipOptionTypesOk

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner) GetVipOptionTypesOk() (*[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner, bool)`

GetVipOptionTypesOk returns a tuple with the VipOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipOptionTypes

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner) SetVipOptionTypes(v []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner)`

SetVipOptionTypes sets VipOptionTypes field to given value.

### HasVipOptionTypes

`func (o *ListLoadBalancerTypes200ResponseAllOfLoadBalancerTypesInner) HasVipOptionTypes() bool`

HasVipOptionTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


