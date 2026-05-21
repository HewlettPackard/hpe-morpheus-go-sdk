# SystemTypeLayoutSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Unique identifier for the layout. | [optional] 
**Name** | Pointer to **string** | Display name of the layout. | [optional] 
**Code** | Pointer to **string** | Unique code identifier for the layout. | [optional] 
**ComponentTypes** | Pointer to [**[]SystemTypeLayoutSummaryComponentTypesInner**](SystemTypeLayoutSummaryComponentTypesInner.md) | The component type definitions associated with this layout. | [optional] 

## Methods

### NewSystemTypeLayoutSummary

`func NewSystemTypeLayoutSummary() *SystemTypeLayoutSummary`

NewSystemTypeLayoutSummary instantiates a new SystemTypeLayoutSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemTypeLayoutSummaryWithDefaults

`func NewSystemTypeLayoutSummaryWithDefaults() *SystemTypeLayoutSummary`

NewSystemTypeLayoutSummaryWithDefaults instantiates a new SystemTypeLayoutSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SystemTypeLayoutSummary) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SystemTypeLayoutSummary) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SystemTypeLayoutSummary) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SystemTypeLayoutSummary) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *SystemTypeLayoutSummary) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SystemTypeLayoutSummary) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SystemTypeLayoutSummary) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SystemTypeLayoutSummary) HasName() bool`

HasName returns a boolean if a field has been set.

### GetCode

`func (o *SystemTypeLayoutSummary) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *SystemTypeLayoutSummary) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *SystemTypeLayoutSummary) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *SystemTypeLayoutSummary) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetComponentTypes

`func (o *SystemTypeLayoutSummary) GetComponentTypes() []SystemTypeLayoutSummaryComponentTypesInner`

GetComponentTypes returns the ComponentTypes field if non-nil, zero value otherwise.

### GetComponentTypesOk

`func (o *SystemTypeLayoutSummary) GetComponentTypesOk() (*[]SystemTypeLayoutSummaryComponentTypesInner, bool)`

GetComponentTypesOk returns a tuple with the ComponentTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponentTypes

`func (o *SystemTypeLayoutSummary) SetComponentTypes(v []SystemTypeLayoutSummaryComponentTypesInner)`

SetComponentTypes sets ComponentTypes field to given value.

### HasComponentTypes

`func (o *SystemTypeLayoutSummary) HasComponentTypes() bool`

HasComponentTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


