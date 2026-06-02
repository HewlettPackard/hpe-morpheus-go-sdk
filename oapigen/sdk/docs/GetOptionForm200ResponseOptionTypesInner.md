# GetOptionForm200ResponseOptionTypesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **NullableString** |  | [optional] 
**Context** | Pointer to **NullableString** |  | [optional] 
**Locked** | Pointer to **bool** |  | [optional] 
**Labels** | Pointer to **[]string** |  | [optional] 
**Options** | Pointer to [**[]GetOptionForm200ResponseOptionTypesInnerOptionsInner**](GetOptionForm200ResponseOptionTypesInnerOptionsInner.md) |  | [optional] 
**FieldGroups** | Pointer to [**[]GetOptionForm200ResponseOptionTypesInnerFieldGroupsInner**](GetOptionForm200ResponseOptionTypesInnerFieldGroupsInner.md) |  | [optional] 

## Methods

### NewGetOptionForm200ResponseOptionTypesInner

`func NewGetOptionForm200ResponseOptionTypesInner() *GetOptionForm200ResponseOptionTypesInner`

NewGetOptionForm200ResponseOptionTypesInner instantiates a new GetOptionForm200ResponseOptionTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetOptionForm200ResponseOptionTypesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetOptionForm200ResponseOptionTypesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetOptionForm200ResponseOptionTypesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetOptionForm200ResponseOptionTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetOptionForm200ResponseOptionTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetOptionForm200ResponseOptionTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetOptionForm200ResponseOptionTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetOptionForm200ResponseOptionTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *GetOptionForm200ResponseOptionTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *GetOptionForm200ResponseOptionTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *GetOptionForm200ResponseOptionTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *GetOptionForm200ResponseOptionTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDescription

`func (o *GetOptionForm200ResponseOptionTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetOptionForm200ResponseOptionTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetOptionForm200ResponseOptionTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetOptionForm200ResponseOptionTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### SetDescriptionNil

`func (o *GetOptionForm200ResponseOptionTypesInner) SetDescriptionNil(b bool)`

 SetDescriptionNil sets the value for Description to be an explicit nil

### UnsetDescription
`func (o *GetOptionForm200ResponseOptionTypesInner) UnsetDescription()`

UnsetDescription ensures that no value is present for Description, not even an explicit nil
### GetContext

`func (o *GetOptionForm200ResponseOptionTypesInner) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *GetOptionForm200ResponseOptionTypesInner) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *GetOptionForm200ResponseOptionTypesInner) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *GetOptionForm200ResponseOptionTypesInner) HasContext() bool`

HasContext returns a boolean if a field has been set.

### SetContextNil

`func (o *GetOptionForm200ResponseOptionTypesInner) SetContextNil(b bool)`

 SetContextNil sets the value for Context to be an explicit nil

### UnsetContext
`func (o *GetOptionForm200ResponseOptionTypesInner) UnsetContext()`

UnsetContext ensures that no value is present for Context, not even an explicit nil
### GetLocked

`func (o *GetOptionForm200ResponseOptionTypesInner) GetLocked() bool`

GetLocked returns the Locked field if non-nil, zero value otherwise.

### GetLockedOk

`func (o *GetOptionForm200ResponseOptionTypesInner) GetLockedOk() (*bool, bool)`

GetLockedOk returns a tuple with the Locked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocked

`func (o *GetOptionForm200ResponseOptionTypesInner) SetLocked(v bool)`

SetLocked sets Locked field to given value.

### HasLocked

`func (o *GetOptionForm200ResponseOptionTypesInner) HasLocked() bool`

HasLocked returns a boolean if a field has been set.

### GetLabels

`func (o *GetOptionForm200ResponseOptionTypesInner) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *GetOptionForm200ResponseOptionTypesInner) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *GetOptionForm200ResponseOptionTypesInner) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *GetOptionForm200ResponseOptionTypesInner) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### SetLabelsNil

`func (o *GetOptionForm200ResponseOptionTypesInner) SetLabelsNil(b bool)`

 SetLabelsNil sets the value for Labels to be an explicit nil

### UnsetLabels
`func (o *GetOptionForm200ResponseOptionTypesInner) UnsetLabels()`

UnsetLabels ensures that no value is present for Labels, not even an explicit nil
### GetOptions

`func (o *GetOptionForm200ResponseOptionTypesInner) GetOptions() []GetOptionForm200ResponseOptionTypesInnerOptionsInner`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *GetOptionForm200ResponseOptionTypesInner) GetOptionsOk() (*[]GetOptionForm200ResponseOptionTypesInnerOptionsInner, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *GetOptionForm200ResponseOptionTypesInner) SetOptions(v []GetOptionForm200ResponseOptionTypesInnerOptionsInner)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *GetOptionForm200ResponseOptionTypesInner) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetFieldGroups

`func (o *GetOptionForm200ResponseOptionTypesInner) GetFieldGroups() []GetOptionForm200ResponseOptionTypesInnerFieldGroupsInner`

GetFieldGroups returns the FieldGroups field if non-nil, zero value otherwise.

### GetFieldGroupsOk

`func (o *GetOptionForm200ResponseOptionTypesInner) GetFieldGroupsOk() (*[]GetOptionForm200ResponseOptionTypesInnerFieldGroupsInner, bool)`

GetFieldGroupsOk returns a tuple with the FieldGroups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFieldGroups

`func (o *GetOptionForm200ResponseOptionTypesInner) SetFieldGroups(v []GetOptionForm200ResponseOptionTypesInnerFieldGroupsInner)`

SetFieldGroups sets FieldGroups field to given value.

### HasFieldGroups

`func (o *GetOptionForm200ResponseOptionTypesInner) HasFieldGroups() bool`

HasFieldGroups returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


