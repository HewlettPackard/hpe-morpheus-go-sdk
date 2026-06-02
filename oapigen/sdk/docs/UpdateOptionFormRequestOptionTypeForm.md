# UpdateOptionFormRequestOptionTypeForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Form name | [optional] 
**Code** | Pointer to **string** | Unique form code | [optional] 
**Description** | Pointer to **NullableString** | A short description of the form | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Options** | Pointer to [**[]UpdateOptionFormRequestOptionTypeFormOptionsInner**](UpdateOptionFormRequestOptionTypeFormOptionsInner.md) | Inputs | [optional] 
**FieldGroups** | Pointer to [**[]UpdateOptionFormRequestOptionTypeFormFieldGroupsInner**](UpdateOptionFormRequestOptionTypeFormFieldGroupsInner.md) | Field Groups | [optional] 

## Methods

### NewUpdateOptionFormRequestOptionTypeForm

`func NewUpdateOptionFormRequestOptionTypeForm() *UpdateOptionFormRequestOptionTypeForm`

NewUpdateOptionFormRequestOptionTypeForm instantiates a new UpdateOptionFormRequestOptionTypeForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetName

`func (o *UpdateOptionFormRequestOptionTypeForm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateOptionFormRequestOptionTypeForm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateOptionFormRequestOptionTypeForm) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateOptionFormRequestOptionTypeForm) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *UpdateOptionFormRequestOptionTypeForm) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *UpdateOptionFormRequestOptionTypeForm) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *UpdateOptionFormRequestOptionTypeForm) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *UpdateOptionFormRequestOptionTypeForm) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateOptionFormRequestOptionTypeForm) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateOptionFormRequestOptionTypeForm) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateOptionFormRequestOptionTypeForm) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateOptionFormRequestOptionTypeForm) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *UpdateOptionFormRequestOptionTypeForm) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *UpdateOptionFormRequestOptionTypeForm) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetLabels

`func (o *UpdateOptionFormRequestOptionTypeForm) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *UpdateOptionFormRequestOptionTypeForm) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *UpdateOptionFormRequestOptionTypeForm) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *UpdateOptionFormRequestOptionTypeForm) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *UpdateOptionFormRequestOptionTypeForm) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *UpdateOptionFormRequestOptionTypeForm) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetOptions

`func (o *UpdateOptionFormRequestOptionTypeForm) GetOptions() []UpdateOptionFormRequestOptionTypeFormOptionsInner`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *UpdateOptionFormRequestOptionTypeForm) GetOptionsOk() (*[]UpdateOptionFormRequestOptionTypeFormOptionsInner, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *UpdateOptionFormRequestOptionTypeForm) SetOptions(v []UpdateOptionFormRequestOptionTypeFormOptionsInner)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *UpdateOptionFormRequestOptionTypeForm) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetFieldGroups

`func (o *UpdateOptionFormRequestOptionTypeForm) GetFieldGroups() []UpdateOptionFormRequestOptionTypeFormFieldGroupsInner`

GetFieldGroups returns the FieldGroups field if non-nil, zero value otherwise.

### GetFieldGroupsOk

`func (o *UpdateOptionFormRequestOptionTypeForm) GetFieldGroupsOk() (*[]UpdateOptionFormRequestOptionTypeFormFieldGroupsInner, bool)`

GetFieldGroupsOk returns a tuple with the FieldGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldGroups

`func (o *UpdateOptionFormRequestOptionTypeForm) SetFieldGroups(v []UpdateOptionFormRequestOptionTypeFormFieldGroupsInner)`

SetFieldGroups sets FieldGroups field to given value.

### HasFieldGroups

`func (o *UpdateOptionFormRequestOptionTypeForm) HasFieldGroups() bool`

HasFieldGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


