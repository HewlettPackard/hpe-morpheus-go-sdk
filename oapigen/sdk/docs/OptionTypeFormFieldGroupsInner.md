# OptionTypeFormFieldGroupsInner

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
**Options** | Pointer to [**[]OptionTypeFormFieldGroupsInnerOptionsInner**](OptionTypeFormFieldGroupsInnerOptionsInner.md) |  | [optional] 

## Methods

### NewOptionTypeFormFieldGroupsInner

`func NewOptionTypeFormFieldGroupsInner() *OptionTypeFormFieldGroupsInner`

NewOptionTypeFormFieldGroupsInner instantiates a new OptionTypeFormFieldGroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *OptionTypeFormFieldGroupsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OptionTypeFormFieldGroupsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OptionTypeFormFieldGroupsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OptionTypeFormFieldGroupsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OptionTypeFormFieldGroupsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OptionTypeFormFieldGroupsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OptionTypeFormFieldGroupsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OptionTypeFormFieldGroupsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *OptionTypeFormFieldGroupsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OptionTypeFormFieldGroupsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OptionTypeFormFieldGroupsInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *OptionTypeFormFieldGroupsInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *OptionTypeFormFieldGroupsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OptionTypeFormFieldGroupsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OptionTypeFormFieldGroupsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OptionTypeFormFieldGroupsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *OptionTypeFormFieldGroupsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *OptionTypeFormFieldGroupsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLocalizedName

`func (o *OptionTypeFormFieldGroupsInner) GetLocalizedName() string`

GetLocalizedName returns the LocalizedName field if non-nil, zero value otherwise.

### GetLocalizedNameOk

`func (o *OptionTypeFormFieldGroupsInner) GetLocalizedNameOk() (*string, bool)`

GetLocalizedNameOk returns a tuple with the LocalizedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedName

`func (o *OptionTypeFormFieldGroupsInner) SetLocalizedName(v string)`

SetLocalizedName sets LocalizedName field to given value.

### HasLocalizedName

`func (o *OptionTypeFormFieldGroupsInner) HasLocalizedName() bool`

HasLocalizedName returns a boolean if a field has been set.

### SetLocalizedNameNil

`func (o *OptionTypeFormFieldGroupsInner) SetLocalizedNameNil(b bool)`

 SetLocalizedNameNil sets the value for LocalizedName to be an explicit nil

### UnsetLocalizedName
`func (o *OptionTypeFormFieldGroupsInner) UnsetLocalizedName()`

UnsetLocalizedName ensures that no value is present for LocalizedName, not even an explicit nil
### GetCollapsible

`func (o *OptionTypeFormFieldGroupsInner) GetCollapsible() bool`

GetCollapsible returns the Collapsible field if non-nil, zero value otherwise.

### GetCollapsibleOk

`func (o *OptionTypeFormFieldGroupsInner) GetCollapsibleOk() (*bool, bool)`

GetCollapsibleOk returns a tuple with the Collapsible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollapsible

`func (o *OptionTypeFormFieldGroupsInner) SetCollapsible(v bool)`

SetCollapsible sets Collapsible field to given value.

### HasCollapsible

`func (o *OptionTypeFormFieldGroupsInner) HasCollapsible() bool`

HasCollapsible returns a boolean if a field has been set.

### GetDefaultCollapsed

`func (o *OptionTypeFormFieldGroupsInner) GetDefaultCollapsed() bool`

GetDefaultCollapsed returns the DefaultCollapsed field if non-nil, zero value otherwise.

### GetDefaultCollapsedOk

`func (o *OptionTypeFormFieldGroupsInner) GetDefaultCollapsedOk() (*bool, bool)`

GetDefaultCollapsedOk returns a tuple with the DefaultCollapsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCollapsed

`func (o *OptionTypeFormFieldGroupsInner) SetDefaultCollapsed(v bool)`

SetDefaultCollapsed sets DefaultCollapsed field to given value.

### HasDefaultCollapsed

`func (o *OptionTypeFormFieldGroupsInner) HasDefaultCollapsed() bool`

HasDefaultCollapsed returns a boolean if a field has been set.

### GetVisibleOnCode

`func (o *OptionTypeFormFieldGroupsInner) GetVisibleOnCode() string`

GetVisibleOnCode returns the VisibleOnCode field if non-nil, zero value otherwise.

### GetVisibleOnCodeOk

`func (o *OptionTypeFormFieldGroupsInner) GetVisibleOnCodeOk() (*string, bool)`

GetVisibleOnCodeOk returns a tuple with the VisibleOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleOnCode

`func (o *OptionTypeFormFieldGroupsInner) SetVisibleOnCode(v string)`

SetVisibleOnCode sets VisibleOnCode field to given value.

### HasVisibleOnCode

`func (o *OptionTypeFormFieldGroupsInner) HasVisibleOnCode() bool`

HasVisibleOnCode returns a boolean if a field has been set.

### SetVisibleOnCodeNil

`func (o *OptionTypeFormFieldGroupsInner) SetVisibleOnCodeNil(b bool)`

 SetVisibleOnCodeNil sets the value for VisibleOnCode to be an explicit nil

### UnsetVisibleOnCode
`func (o *OptionTypeFormFieldGroupsInner) UnsetVisibleOnCode()`

UnsetVisibleOnCode ensures that no value is present for VisibleOnCode, not even an explicit nil
### GetOptions

`func (o *OptionTypeFormFieldGroupsInner) GetOptions() []OptionTypeFormFieldGroupsInnerOptionsInner`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *OptionTypeFormFieldGroupsInner) GetOptionsOk() (*[]OptionTypeFormFieldGroupsInnerOptionsInner, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *OptionTypeFormFieldGroupsInner) SetOptions(v []OptionTypeFormFieldGroupsInnerOptionsInner)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *OptionTypeFormFieldGroupsInner) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


