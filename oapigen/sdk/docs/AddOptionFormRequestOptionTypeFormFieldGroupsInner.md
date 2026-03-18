# AddOptionFormRequestOptionTypeFormFieldGroupsInner

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
**Options** | Pointer to [**[]AddOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner**](AddOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner.md) |  | [optional] 

## Methods

### NewAddOptionFormRequestOptionTypeFormFieldGroupsInner

`func NewAddOptionFormRequestOptionTypeFormFieldGroupsInner() *AddOptionFormRequestOptionTypeFormFieldGroupsInner`

NewAddOptionFormRequestOptionTypeFormFieldGroupsInner instantiates a new AddOptionFormRequestOptionTypeFormFieldGroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddOptionFormRequestOptionTypeFormFieldGroupsInnerWithDefaults

`func NewAddOptionFormRequestOptionTypeFormFieldGroupsInnerWithDefaults() *AddOptionFormRequestOptionTypeFormFieldGroupsInner`

NewAddOptionFormRequestOptionTypeFormFieldGroupsInnerWithDefaults instantiates a new AddOptionFormRequestOptionTypeFormFieldGroupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLocalizedName

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) GetLocalizedName() string`

GetLocalizedName returns the LocalizedName field if non-nil, zero value otherwise.

### GetLocalizedNameOk

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) GetLocalizedNameOk() (*string, bool)`

GetLocalizedNameOk returns a tuple with the LocalizedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedName

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) SetLocalizedName(v string)`

SetLocalizedName sets LocalizedName field to given value.

### HasLocalizedName

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) HasLocalizedName() bool`

HasLocalizedName returns a boolean if a field has been set.

### SetLocalizedNameNil

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) SetLocalizedNameNil(b bool)`

 SetLocalizedNameNil sets the value for LocalizedName to be an explicit nil

### UnsetLocalizedName
`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) UnsetLocalizedName()`

UnsetLocalizedName ensures that no value is present for LocalizedName, not even an explicit nil
### GetCollapsible

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) GetCollapsible() bool`

GetCollapsible returns the Collapsible field if non-nil, zero value otherwise.

### GetCollapsibleOk

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) GetCollapsibleOk() (*bool, bool)`

GetCollapsibleOk returns a tuple with the Collapsible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollapsible

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) SetCollapsible(v bool)`

SetCollapsible sets Collapsible field to given value.

### HasCollapsible

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) HasCollapsible() bool`

HasCollapsible returns a boolean if a field has been set.

### GetDefaultCollapsed

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) GetDefaultCollapsed() bool`

GetDefaultCollapsed returns the DefaultCollapsed field if non-nil, zero value otherwise.

### GetDefaultCollapsedOk

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) GetDefaultCollapsedOk() (*bool, bool)`

GetDefaultCollapsedOk returns a tuple with the DefaultCollapsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCollapsed

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) SetDefaultCollapsed(v bool)`

SetDefaultCollapsed sets DefaultCollapsed field to given value.

### HasDefaultCollapsed

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) HasDefaultCollapsed() bool`

HasDefaultCollapsed returns a boolean if a field has been set.

### GetVisibleOnCode

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) GetVisibleOnCode() string`

GetVisibleOnCode returns the VisibleOnCode field if non-nil, zero value otherwise.

### GetVisibleOnCodeOk

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) GetVisibleOnCodeOk() (*string, bool)`

GetVisibleOnCodeOk returns a tuple with the VisibleOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleOnCode

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) SetVisibleOnCode(v string)`

SetVisibleOnCode sets VisibleOnCode field to given value.

### HasVisibleOnCode

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) HasVisibleOnCode() bool`

HasVisibleOnCode returns a boolean if a field has been set.

### SetVisibleOnCodeNil

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) SetVisibleOnCodeNil(b bool)`

 SetVisibleOnCodeNil sets the value for VisibleOnCode to be an explicit nil

### UnsetVisibleOnCode
`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) UnsetVisibleOnCode()`

UnsetVisibleOnCode ensures that no value is present for VisibleOnCode, not even an explicit nil
### GetOptions

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) GetOptions() []AddOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) GetOptionsOk() (*[]AddOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) SetOptions(v []AddOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AddOptionFormRequestOptionTypeFormFieldGroupsInner) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


