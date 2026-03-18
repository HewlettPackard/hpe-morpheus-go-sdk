# OptionTypeFormCreateFieldGroupsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**LocalizedName** | Pointer to **NullableString** |  | [optional] 
**Collapsible** | Pointer to **bool** |  | [optional] 
**DefaultCollapsed** | Pointer to **bool** |  | [optional] 
**VisibleOnCode** | Pointer to **NullableString** |  | [optional] 
**Options** | Pointer to [**[]OptionTypeFormCreateFieldGroupsInnerOptionsInner**](OptionTypeFormCreateFieldGroupsInnerOptionsInner.md) |  | [optional] 

## Methods

### NewOptionTypeFormCreateFieldGroupsInner

`func NewOptionTypeFormCreateFieldGroupsInner() *OptionTypeFormCreateFieldGroupsInner`

NewOptionTypeFormCreateFieldGroupsInner instantiates a new OptionTypeFormCreateFieldGroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionTypeFormCreateFieldGroupsInnerWithDefaults

`func NewOptionTypeFormCreateFieldGroupsInnerWithDefaults() *OptionTypeFormCreateFieldGroupsInner`

NewOptionTypeFormCreateFieldGroupsInnerWithDefaults instantiates a new OptionTypeFormCreateFieldGroupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OptionTypeFormCreateFieldGroupsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OptionTypeFormCreateFieldGroupsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OptionTypeFormCreateFieldGroupsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OptionTypeFormCreateFieldGroupsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OptionTypeFormCreateFieldGroupsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OptionTypeFormCreateFieldGroupsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OptionTypeFormCreateFieldGroupsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OptionTypeFormCreateFieldGroupsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *OptionTypeFormCreateFieldGroupsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OptionTypeFormCreateFieldGroupsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OptionTypeFormCreateFieldGroupsInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *OptionTypeFormCreateFieldGroupsInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *OptionTypeFormCreateFieldGroupsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OptionTypeFormCreateFieldGroupsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OptionTypeFormCreateFieldGroupsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OptionTypeFormCreateFieldGroupsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *OptionTypeFormCreateFieldGroupsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *OptionTypeFormCreateFieldGroupsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLocalizedName

`func (o *OptionTypeFormCreateFieldGroupsInner) GetLocalizedName() string`

GetLocalizedName returns the LocalizedName field if non-nil, zero value otherwise.

### GetLocalizedNameOk

`func (o *OptionTypeFormCreateFieldGroupsInner) GetLocalizedNameOk() (*string, bool)`

GetLocalizedNameOk returns a tuple with the LocalizedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedName

`func (o *OptionTypeFormCreateFieldGroupsInner) SetLocalizedName(v string)`

SetLocalizedName sets LocalizedName field to given value.

### HasLocalizedName

`func (o *OptionTypeFormCreateFieldGroupsInner) HasLocalizedName() bool`

HasLocalizedName returns a boolean if a field has been set.

### SetLocalizedNameNil

`func (o *OptionTypeFormCreateFieldGroupsInner) SetLocalizedNameNil(b bool)`

 SetLocalizedNameNil sets the value for LocalizedName to be an explicit nil

### UnsetLocalizedName
`func (o *OptionTypeFormCreateFieldGroupsInner) UnsetLocalizedName()`

UnsetLocalizedName ensures that no value is present for LocalizedName, not even an explicit nil
### GetCollapsible

`func (o *OptionTypeFormCreateFieldGroupsInner) GetCollapsible() bool`

GetCollapsible returns the Collapsible field if non-nil, zero value otherwise.

### GetCollapsibleOk

`func (o *OptionTypeFormCreateFieldGroupsInner) GetCollapsibleOk() (*bool, bool)`

GetCollapsibleOk returns a tuple with the Collapsible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollapsible

`func (o *OptionTypeFormCreateFieldGroupsInner) SetCollapsible(v bool)`

SetCollapsible sets Collapsible field to given value.

### HasCollapsible

`func (o *OptionTypeFormCreateFieldGroupsInner) HasCollapsible() bool`

HasCollapsible returns a boolean if a field has been set.

### GetDefaultCollapsed

`func (o *OptionTypeFormCreateFieldGroupsInner) GetDefaultCollapsed() bool`

GetDefaultCollapsed returns the DefaultCollapsed field if non-nil, zero value otherwise.

### GetDefaultCollapsedOk

`func (o *OptionTypeFormCreateFieldGroupsInner) GetDefaultCollapsedOk() (*bool, bool)`

GetDefaultCollapsedOk returns a tuple with the DefaultCollapsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCollapsed

`func (o *OptionTypeFormCreateFieldGroupsInner) SetDefaultCollapsed(v bool)`

SetDefaultCollapsed sets DefaultCollapsed field to given value.

### HasDefaultCollapsed

`func (o *OptionTypeFormCreateFieldGroupsInner) HasDefaultCollapsed() bool`

HasDefaultCollapsed returns a boolean if a field has been set.

### GetVisibleOnCode

`func (o *OptionTypeFormCreateFieldGroupsInner) GetVisibleOnCode() string`

GetVisibleOnCode returns the VisibleOnCode field if non-nil, zero value otherwise.

### GetVisibleOnCodeOk

`func (o *OptionTypeFormCreateFieldGroupsInner) GetVisibleOnCodeOk() (*string, bool)`

GetVisibleOnCodeOk returns a tuple with the VisibleOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleOnCode

`func (o *OptionTypeFormCreateFieldGroupsInner) SetVisibleOnCode(v string)`

SetVisibleOnCode sets VisibleOnCode field to given value.

### HasVisibleOnCode

`func (o *OptionTypeFormCreateFieldGroupsInner) HasVisibleOnCode() bool`

HasVisibleOnCode returns a boolean if a field has been set.

### SetVisibleOnCodeNil

`func (o *OptionTypeFormCreateFieldGroupsInner) SetVisibleOnCodeNil(b bool)`

 SetVisibleOnCodeNil sets the value for VisibleOnCode to be an explicit nil

### UnsetVisibleOnCode
`func (o *OptionTypeFormCreateFieldGroupsInner) UnsetVisibleOnCode()`

UnsetVisibleOnCode ensures that no value is present for VisibleOnCode, not even an explicit nil
### GetOptions

`func (o *OptionTypeFormCreateFieldGroupsInner) GetOptions() []OptionTypeFormCreateFieldGroupsInnerOptionsInner`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *OptionTypeFormCreateFieldGroupsInner) GetOptionsOk() (*[]OptionTypeFormCreateFieldGroupsInnerOptionsInner, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *OptionTypeFormCreateFieldGroupsInner) SetOptions(v []OptionTypeFormCreateFieldGroupsInnerOptionsInner)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *OptionTypeFormCreateFieldGroupsInner) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


