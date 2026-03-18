# GetLoadBalancerType200ResponseLoadBalancerType

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
**OptionTypes** | Pointer to [**[]GetLoadBalancerType200ResponseLoadBalancerTypeOptionTypesInner**](GetLoadBalancerType200ResponseLoadBalancerTypeOptionTypesInner.md) |  | [optional] 
**VipOptionTypes** | Pointer to [**[]GetLoadBalancerType200ResponseLoadBalancerTypeVipOptionTypesInner**](GetLoadBalancerType200ResponseLoadBalancerTypeVipOptionTypesInner.md) |  | [optional] 

## Methods

### NewGetLoadBalancerType200ResponseLoadBalancerType

`func NewGetLoadBalancerType200ResponseLoadBalancerType() *GetLoadBalancerType200ResponseLoadBalancerType`

NewGetLoadBalancerType200ResponseLoadBalancerType instantiates a new GetLoadBalancerType200ResponseLoadBalancerType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetLoadBalancerType200ResponseLoadBalancerTypeWithDefaults

`func NewGetLoadBalancerType200ResponseLoadBalancerTypeWithDefaults() *GetLoadBalancerType200ResponseLoadBalancerType`

NewGetLoadBalancerType200ResponseLoadBalancerTypeWithDefaults instantiates a new GetLoadBalancerType200ResponseLoadBalancerType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetLoadBalancerType200ResponseLoadBalancerType) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetLoadBalancerType200ResponseLoadBalancerType) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetLoadBalancerType200ResponseLoadBalancerType) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetLoadBalancerType200ResponseLoadBalancerType) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetLoadBalancerType200ResponseLoadBalancerType) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetLoadBalancerType200ResponseLoadBalancerType) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetLoadBalancerType200ResponseLoadBalancerType) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetLoadBalancerType200ResponseLoadBalancerType) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *GetLoadBalancerType200ResponseLoadBalancerType) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetLoadBalancerType200ResponseLoadBalancerType) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetLoadBalancerType200ResponseLoadBalancerType) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetLoadBalancerType200ResponseLoadBalancerType) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetEnabled

`func (o *GetLoadBalancerType200ResponseLoadBalancerType) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetLoadBalancerType200ResponseLoadBalancerType) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetLoadBalancerType200ResponseLoadBalancerType) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetLoadBalancerType200ResponseLoadBalancerType) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetInternal

`func (o *GetLoadBalancerType200ResponseLoadBalancerType) GetInternal() bool`

GetInternal returns the Internal field if non-nil, zero value otherwise.

### GetInternalOk

`func (o *GetLoadBalancerType200ResponseLoadBalancerType) GetInternalOk() (*bool, bool)`

GetInternalOk returns a tuple with the Internal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternal

`func (o *GetLoadBalancerType200ResponseLoadBalancerType) SetInternal(v bool)`

SetInternal sets Internal field to given value.

### HasInternal

`func (o *GetLoadBalancerType200ResponseLoadBalancerType) HasInternal() bool`

HasInternal returns a boolean if a field has been set.

### GetCreatable

`func (o *GetLoadBalancerType200ResponseLoadBalancerType) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *GetLoadBalancerType200ResponseLoadBalancerType) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *GetLoadBalancerType200ResponseLoadBalancerType) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *GetLoadBalancerType200ResponseLoadBalancerType) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetCreateType

`func (o *GetLoadBalancerType200ResponseLoadBalancerType) GetCreateType() string`

GetCreateType returns the CreateType field if non-nil, zero value otherwise.

### GetCreateTypeOk

`func (o *GetLoadBalancerType200ResponseLoadBalancerType) GetCreateTypeOk() (*string, bool)`

GetCreateTypeOk returns a tuple with the CreateType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreateType

`func (o *GetLoadBalancerType200ResponseLoadBalancerType) SetCreateType(v string)`

SetCreateType sets CreateType field to given value.

### HasCreateType

`func (o *GetLoadBalancerType200ResponseLoadBalancerType) HasCreateType() bool`

HasCreateType returns a boolean if a field has been set.

### GetOptionTypes

`func (o *GetLoadBalancerType200ResponseLoadBalancerType) GetOptionTypes() []GetLoadBalancerType200ResponseLoadBalancerTypeOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *GetLoadBalancerType200ResponseLoadBalancerType) GetOptionTypesOk() (*[]GetLoadBalancerType200ResponseLoadBalancerTypeOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *GetLoadBalancerType200ResponseLoadBalancerType) SetOptionTypes(v []GetLoadBalancerType200ResponseLoadBalancerTypeOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *GetLoadBalancerType200ResponseLoadBalancerType) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.

### GetVipOptionTypes

`func (o *GetLoadBalancerType200ResponseLoadBalancerType) GetVipOptionTypes() []GetLoadBalancerType200ResponseLoadBalancerTypeVipOptionTypesInner`

GetVipOptionTypes returns the VipOptionTypes field if non-nil, zero value otherwise.

### GetVipOptionTypesOk

`func (o *GetLoadBalancerType200ResponseLoadBalancerType) GetVipOptionTypesOk() (*[]GetLoadBalancerType200ResponseLoadBalancerTypeVipOptionTypesInner, bool)`

GetVipOptionTypesOk returns a tuple with the VipOptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVipOptionTypes

`func (o *GetLoadBalancerType200ResponseLoadBalancerType) SetVipOptionTypes(v []GetLoadBalancerType200ResponseLoadBalancerTypeVipOptionTypesInner)`

SetVipOptionTypes sets VipOptionTypes field to given value.

### HasVipOptionTypes

`func (o *GetLoadBalancerType200ResponseLoadBalancerType) HasVipOptionTypes() bool`

HasVipOptionTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


