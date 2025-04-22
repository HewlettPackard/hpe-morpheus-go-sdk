# ListCredentialTypes200ResponseAllOfCredentialTypesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Code** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Description** | Pointer to **string** |  | [optional] 
**Enabled** | Pointer to **bool** |  | [optional] 
**Creatable** | Pointer to **bool** |  | [optional] 
**Editable** | Pointer to **bool** |  | [optional] 
**OptionTypes** | Pointer to [**[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner**](ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner.md) |  | [optional] 

## Methods

### NewListCredentialTypes200ResponseAllOfCredentialTypesInner

`func NewListCredentialTypes200ResponseAllOfCredentialTypesInner() *ListCredentialTypes200ResponseAllOfCredentialTypesInner`

NewListCredentialTypes200ResponseAllOfCredentialTypesInner instantiates a new ListCredentialTypes200ResponseAllOfCredentialTypesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCredentialTypes200ResponseAllOfCredentialTypesInnerWithDefaults

`func NewListCredentialTypes200ResponseAllOfCredentialTypesInnerWithDefaults() *ListCredentialTypes200ResponseAllOfCredentialTypesInner`

NewListCredentialTypes200ResponseAllOfCredentialTypesInnerWithDefaults instantiates a new ListCredentialTypes200ResponseAllOfCredentialTypesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ListCredentialTypes200ResponseAllOfCredentialTypesInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListCredentialTypes200ResponseAllOfCredentialTypesInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListCredentialTypes200ResponseAllOfCredentialTypesInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListCredentialTypes200ResponseAllOfCredentialTypesInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetCode

`func (o *ListCredentialTypes200ResponseAllOfCredentialTypesInner) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ListCredentialTypes200ResponseAllOfCredentialTypesInner) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ListCredentialTypes200ResponseAllOfCredentialTypesInner) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *ListCredentialTypes200ResponseAllOfCredentialTypesInner) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetName

`func (o *ListCredentialTypes200ResponseAllOfCredentialTypesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListCredentialTypes200ResponseAllOfCredentialTypesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListCredentialTypes200ResponseAllOfCredentialTypesInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListCredentialTypes200ResponseAllOfCredentialTypesInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListCredentialTypes200ResponseAllOfCredentialTypesInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListCredentialTypes200ResponseAllOfCredentialTypesInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListCredentialTypes200ResponseAllOfCredentialTypesInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListCredentialTypes200ResponseAllOfCredentialTypesInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *ListCredentialTypes200ResponseAllOfCredentialTypesInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListCredentialTypes200ResponseAllOfCredentialTypesInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListCredentialTypes200ResponseAllOfCredentialTypesInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListCredentialTypes200ResponseAllOfCredentialTypesInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetCreatable

`func (o *ListCredentialTypes200ResponseAllOfCredentialTypesInner) GetCreatable() bool`

GetCreatable returns the Creatable field if non-nil, zero value otherwise.

### GetCreatableOk

`func (o *ListCredentialTypes200ResponseAllOfCredentialTypesInner) GetCreatableOk() (*bool, bool)`

GetCreatableOk returns a tuple with the Creatable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatable

`func (o *ListCredentialTypes200ResponseAllOfCredentialTypesInner) SetCreatable(v bool)`

SetCreatable sets Creatable field to given value.

### HasCreatable

`func (o *ListCredentialTypes200ResponseAllOfCredentialTypesInner) HasCreatable() bool`

HasCreatable returns a boolean if a field has been set.

### GetEditable

`func (o *ListCredentialTypes200ResponseAllOfCredentialTypesInner) GetEditable() bool`

GetEditable returns the Editable field if non-nil, zero value otherwise.

### GetEditableOk

`func (o *ListCredentialTypes200ResponseAllOfCredentialTypesInner) GetEditableOk() (*bool, bool)`

GetEditableOk returns a tuple with the Editable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEditable

`func (o *ListCredentialTypes200ResponseAllOfCredentialTypesInner) SetEditable(v bool)`

SetEditable sets Editable field to given value.

### HasEditable

`func (o *ListCredentialTypes200ResponseAllOfCredentialTypesInner) HasEditable() bool`

HasEditable returns a boolean if a field has been set.

### GetOptionTypes

`func (o *ListCredentialTypes200ResponseAllOfCredentialTypesInner) GetOptionTypes() []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner`

GetOptionTypes returns the OptionTypes field if non-nil, zero value otherwise.

### GetOptionTypesOk

`func (o *ListCredentialTypes200ResponseAllOfCredentialTypesInner) GetOptionTypesOk() (*[]ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner, bool)`

GetOptionTypesOk returns a tuple with the OptionTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptionTypes

`func (o *ListCredentialTypes200ResponseAllOfCredentialTypesInner) SetOptionTypes(v []ListCatalogItemTypes200ResponseAllOfCatalogItemTypesInnerOptionTypesInner)`

SetOptionTypes sets OptionTypes field to given value.

### HasOptionTypes

`func (o *ListCredentialTypes200ResponseAllOfCredentialTypesInner) HasOptionTypes() bool`

HasOptionTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


