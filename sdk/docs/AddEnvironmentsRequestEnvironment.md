# AddEnvironmentsRequestEnvironment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A unique name for the environment | 
**Code** | **string** | A unique code for the environment | 
**Description** | Pointer to **string** | A description of the environment | [optional] 
**Visibility** | Pointer to **string** | private or public | [optional] [default to "private"]
**SortOrder** | Pointer to **int64** | Sort order | [optional] [default to 0]

## Methods

### NewAddEnvironmentsRequestEnvironment

`func NewAddEnvironmentsRequestEnvironment(name string, code string, ) *AddEnvironmentsRequestEnvironment`

NewAddEnvironmentsRequestEnvironment instantiates a new AddEnvironmentsRequestEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddEnvironmentsRequestEnvironmentWithDefaults

`func NewAddEnvironmentsRequestEnvironmentWithDefaults() *AddEnvironmentsRequestEnvironment`

NewAddEnvironmentsRequestEnvironmentWithDefaults instantiates a new AddEnvironmentsRequestEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddEnvironmentsRequestEnvironment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddEnvironmentsRequestEnvironment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddEnvironmentsRequestEnvironment) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *AddEnvironmentsRequestEnvironment) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AddEnvironmentsRequestEnvironment) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AddEnvironmentsRequestEnvironment) SetCode(v string)`

SetCode sets Code field to given value.


### GetDescription

`func (o *AddEnvironmentsRequestEnvironment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddEnvironmentsRequestEnvironment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddEnvironmentsRequestEnvironment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddEnvironmentsRequestEnvironment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVisibility

`func (o *AddEnvironmentsRequestEnvironment) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *AddEnvironmentsRequestEnvironment) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *AddEnvironmentsRequestEnvironment) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *AddEnvironmentsRequestEnvironment) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetSortOrder

`func (o *AddEnvironmentsRequestEnvironment) GetSortOrder() int64`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *AddEnvironmentsRequestEnvironment) GetSortOrderOk() (*int64, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *AddEnvironmentsRequestEnvironment) SetSortOrder(v int64)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *AddEnvironmentsRequestEnvironment) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


