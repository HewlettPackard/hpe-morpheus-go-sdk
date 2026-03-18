# UpdateOptionFormRequestOptionTypeFormFieldGroupsInner

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
**Options** | Pointer to [**[]UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner**](UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner.md) |  | [optional] 

## Methods

### NewUpdateOptionFormRequestOptionTypeFormFieldGroupsInner

`func NewUpdateOptionFormRequestOptionTypeFormFieldGroupsInner() *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner`

NewUpdateOptionFormRequestOptionTypeFormFieldGroupsInner instantiates a new UpdateOptionFormRequestOptionTypeFormFieldGroupsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateOptionFormRequestOptionTypeFormFieldGroupsInnerWithDefaults

`func NewUpdateOptionFormRequestOptionTypeFormFieldGroupsInnerWithDefaults() *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner`

NewUpdateOptionFormRequestOptionTypeFormFieldGroupsInnerWithDefaults instantiates a new UpdateOptionFormRequestOptionTypeFormFieldGroupsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLocalizedName

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) GetLocalizedName() string`

GetLocalizedName returns the LocalizedName field if non-nil, zero value otherwise.

### GetLocalizedNameOk

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) GetLocalizedNameOk() (*string, bool)`

GetLocalizedNameOk returns a tuple with the LocalizedName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizedName

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) SetLocalizedName(v string)`

SetLocalizedName sets LocalizedName field to given value.

### HasLocalizedName

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) HasLocalizedName() bool`

HasLocalizedName returns a boolean if a field has been set.

### SetLocalizedNameNil

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) SetLocalizedNameNil(b bool)`

 SetLocalizedNameNil sets the value for LocalizedName to be an explicit nil

### UnsetLocalizedName
`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) UnsetLocalizedName()`

UnsetLocalizedName ensures that no value is present for LocalizedName, not even an explicit nil
### GetCollapsible

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) GetCollapsible() bool`

GetCollapsible returns the Collapsible field if non-nil, zero value otherwise.

### GetCollapsibleOk

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) GetCollapsibleOk() (*bool, bool)`

GetCollapsibleOk returns a tuple with the Collapsible field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollapsible

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) SetCollapsible(v bool)`

SetCollapsible sets Collapsible field to given value.

### HasCollapsible

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) HasCollapsible() bool`

HasCollapsible returns a boolean if a field has been set.

### GetDefaultCollapsed

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) GetDefaultCollapsed() bool`

GetDefaultCollapsed returns the DefaultCollapsed field if non-nil, zero value otherwise.

### GetDefaultCollapsedOk

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) GetDefaultCollapsedOk() (*bool, bool)`

GetDefaultCollapsedOk returns a tuple with the DefaultCollapsed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultCollapsed

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) SetDefaultCollapsed(v bool)`

SetDefaultCollapsed sets DefaultCollapsed field to given value.

### HasDefaultCollapsed

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) HasDefaultCollapsed() bool`

HasDefaultCollapsed returns a boolean if a field has been set.

### GetVisibleOnCode

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) GetVisibleOnCode() string`

GetVisibleOnCode returns the VisibleOnCode field if non-nil, zero value otherwise.

### GetVisibleOnCodeOk

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) GetVisibleOnCodeOk() (*string, bool)`

GetVisibleOnCodeOk returns a tuple with the VisibleOnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibleOnCode

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) SetVisibleOnCode(v string)`

SetVisibleOnCode sets VisibleOnCode field to given value.

### HasVisibleOnCode

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) HasVisibleOnCode() bool`

HasVisibleOnCode returns a boolean if a field has been set.

### SetVisibleOnCodeNil

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) SetVisibleOnCodeNil(b bool)`

 SetVisibleOnCodeNil sets the value for VisibleOnCode to be an explicit nil

### UnsetVisibleOnCode
`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) UnsetVisibleOnCode()`

UnsetVisibleOnCode ensures that no value is present for VisibleOnCode, not even an explicit nil
### GetOptions

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) GetOptions() []UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) GetOptionsOk() (*[]UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) SetOptions(v []UpdateOptionFormRequestOptionTypeFormFieldGroupsInnerOptionsInner)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *UpdateOptionFormRequestOptionTypeFormFieldGroupsInner) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


