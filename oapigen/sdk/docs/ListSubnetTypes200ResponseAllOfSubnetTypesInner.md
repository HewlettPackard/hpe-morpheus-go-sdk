# ListSubnetTypes200ResponseAllOfSubnetTypesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Creatable** | Pointer to **bool** |  | [optional] 
**Deletable** | Pointer to **bool** |  | [optional] 
**DhcpServerEditable** | Pointer to **bool** |  | [optional] 
**CanAssignPool** | Pointer to **bool** |  | [optional] 
**OptionTypes** | Pointer to [**[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner**](ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner.md) |  | [optional] 

## Methods

### NewListSubnetTypes200ResponseAllOfSubnetTypesInner

`func NewListSubnetTypes200ResponseAllOfSubnetTypesInner() *ListSubnetTypes200ResponseAllOfSubnetTypesInner`

NewListSubnetTypes200ResponseAllOfSubnetTypesInner instantiates a new ListSubnetTypes200ResponseAllOfSubnetTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSubnetTypes200ResponseAllOfSubnetTypesInnerWithDefaults

`func NewListSubnetTypes200ResponseAllOfSubnetTypesInnerWithDefaults() *ListSubnetTypes200ResponseAllOfSubnetTypesInner`

NewListSubnetTypes200ResponseAllOfSubnetTypesInnerWithDefaults instantiates a new ListSubnetTypes200ResponseAllOfSubnetTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListSubnetTypes200ResponseAllOfSubnetTypesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListSubnetTypes200ResponseAllOfSubnetTypesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListSubnetTypes200ResponseAllOfSubnetTypesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListSubnetTypes200ResponseAllOfSubnetTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *ListSubnetTypes200ResponseAllOfSubnetTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListSubnetTypes200ResponseAllOfSubnetTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListSubnetTypes200ResponseAllOfSubnetTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListSubnetTypes200ResponseAllOfSubnetTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *ListSubnetTypes200ResponseAllOfSubnetTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListSubnetTypes200ResponseAllOfSubnetTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListSubnetTypes200ResponseAllOfSubnetTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListSubnetTypes200ResponseAllOfSubnetTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListSubnetTypes200ResponseAllOfSubnetTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListSubnetTypes200ResponseAllOfSubnetTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListSubnetTypes200ResponseAllOfSubnetTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListSubnetTypes200ResponseAllOfSubnetTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetCreatable

`func (o *ListSubnetTypes200ResponseAllOfSubnetTypesInner) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *ListSubnetTypes200ResponseAllOfSubnetTypesInner) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *ListSubnetTypes200ResponseAllOfSubnetTypesInner) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *ListSubnetTypes200ResponseAllOfSubnetTypesInner) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetDeletable

`func (o *ListSubnetTypes200ResponseAllOfSubnetTypesInner) GetDeletable() bool`

GetDeletable returns the Deletable field if non-nil, zero value otherwise.

### GetDeletableOk

`func (o *ListSubnetTypes200ResponseAllOfSubnetTypesInner) GetDeletableOk() (*bool, bool)`

GetDeletableOk returns a tuple with the Deletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletable

`func (o *ListSubnetTypes200ResponseAllOfSubnetTypesInner) SetDeletable(v bool)`

SetDeletable sets Deletable field to given value.

### HasDeletable

`func (o *ListSubnetTypes200ResponseAllOfSubnetTypesInner) HasDeletable() bool`

HasDeletable returns a boolean if a field has been set.

### GetDhcpServerEditable

`func (o *ListSubnetTypes200ResponseAllOfSubnetTypesInner) GetDhcpServerEditable() bool`

GetDhcpServerEditable returns the DhcpServerEditable field if non-nil, zero value otherwise.

### GetDhcpServerEditableOk

`func (o *ListSubnetTypes200ResponseAllOfSubnetTypesInner) GetDhcpServerEditableOk() (*bool, bool)`

GetDhcpServerEditableOk returns a tuple with the DhcpServerEditable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDhcpServerEditable

`func (o *ListSubnetTypes200ResponseAllOfSubnetTypesInner) SetDhcpServerEditable(v bool)`

SetDhcpServerEditable sets DhcpServerEditable field to given value.

### HasDhcpServerEditable

`func (o *ListSubnetTypes200ResponseAllOfSubnetTypesInner) HasDhcpServerEditable() bool`

HasDhcpServerEditable returns a boolean if a field has been set.

### GetCanAssignPool

`func (o *ListSubnetTypes200ResponseAllOfSubnetTypesInner) GetCanAssignPool() bool`

GetCanAssignPool returns the CanAssignPool field if non-nil, zero value otherwise.

### GetCanAssignPoolOk

`func (o *ListSubnetTypes200ResponseAllOfSubnetTypesInner) GetCanAssignPoolOk() (*bool, bool)`

GetCanAssignPoolOk returns a tuple with the CanAssignPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCanAssignPool

`func (o *ListSubnetTypes200ResponseAllOfSubnetTypesInner) SetCanAssignPool(v bool)`

SetCanAssignPool sets CanAssignPool field to given value.

### HasCanAssignPool

`func (o *ListSubnetTypes200ResponseAllOfSubnetTypesInner) HasCanAssignPool() bool`

HasCanAssignPool returns a boolean if a field has been set.

### GetOptionTypes

`func (o *ListSubnetTypes200ResponseAllOfSubnetTypesInner) GetOptionTypes() []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *ListSubnetTypes200ResponseAllOfSubnetTypesInner) GetOptionTypesOk() (*[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *ListSubnetTypes200ResponseAllOfSubnetTypesInner) SetOptionTypes(v []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *ListSubnetTypes200ResponseAllOfSubnetTypesInner) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


