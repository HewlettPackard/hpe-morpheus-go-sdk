# OptionTypeForm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Context** | Pointer to **string** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Options** | Pointer to [**[]ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner**](ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner.md) |  | [optional] 
**FieldGroups** | Pointer to [**[]ListOptionForms200ResponseAllOfOptionTypesInnerFieldGroupsInner**](ListOptionForms200ResponseAllOfOptionTypesInnerFieldGroupsInner.md) |  | [optional] 

## Methods

### NewOptionTypeForm

`func NewOptionTypeForm() *OptionTypeForm`

NewOptionTypeForm instantiates a new OptionTypeForm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionTypeFormWithDefaults

`func NewOptionTypeFormWithDefaults() *OptionTypeForm`

NewOptionTypeFormWithDefaults instantiates a new OptionTypeForm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OptionTypeForm) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OptionTypeForm) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OptionTypeForm) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *OptionTypeForm) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *OptionTypeForm) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *OptionTypeForm) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *OptionTypeForm) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *OptionTypeForm) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *OptionTypeForm) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *OptionTypeForm) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *OptionTypeForm) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *OptionTypeForm) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *OptionTypeForm) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *OptionTypeForm) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *OptionTypeForm) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *OptionTypeForm) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetContext

`func (o *OptionTypeForm) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *OptionTypeForm) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *OptionTypeForm) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *OptionTypeForm) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetLocked

`func (o *OptionTypeForm) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *OptionTypeForm) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *OptionTypeForm) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *OptionTypeForm) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetLabels

`func (o *OptionTypeForm) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *OptionTypeForm) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *OptionTypeForm) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *OptionTypeForm) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetOptions

`func (o *OptionTypeForm) GetOptions() []ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *OptionTypeForm) GetOptionsOk() (*[]ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *OptionTypeForm) SetOptions(v []ListOptionForms200ResponseAllOfOptionTypesInnerOptionsInner)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *OptionTypeForm) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetFieldGroups

`func (o *OptionTypeForm) GetFieldGroups() []ListOptionForms200ResponseAllOfOptionTypesInnerFieldGroupsInner`

GetFieldGroups returns the FieldGroups field if non-nil, zero value otherwise.

### GetFieldGroupsOk

`func (o *OptionTypeForm) GetFieldGroupsOk() (*[]ListOptionForms200ResponseAllOfOptionTypesInnerFieldGroupsInner, bool)`

GetFieldGroupsOk returns a tuple with the FieldGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldGroups

`func (o *OptionTypeForm) SetFieldGroups(v []ListOptionForms200ResponseAllOfOptionTypesInnerFieldGroupsInner)`

SetFieldGroups sets FieldGroups field to given value.

### HasFieldGroups

`func (o *OptionTypeForm) HasFieldGroups() bool`

HasFieldGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


