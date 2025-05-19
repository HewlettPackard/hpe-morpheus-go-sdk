# UpdateEnvironmentsRequestEnvironment

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | A unique name for the environment | [optional] 
**Description** | Pointer to **string** | A description of the environment | [optional] 
**Visibility** | Pointer to **string** | private or public | [optional] [default to "private"]
**SortOrder** | Pointer to **int64** | Sort order | [optional] [default to 0]
**Active** | Pointer to **bool** | Set to false to deactivate the environment | [optional] 

## Methods

### NewUpdateEnvironmentsRequestEnvironment

`func NewUpdateEnvironmentsRequestEnvironment() *UpdateEnvironmentsRequestEnvironment`

NewUpdateEnvironmentsRequestEnvironment instantiates a new UpdateEnvironmentsRequestEnvironment object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEnvironmentsRequestEnvironmentWithDefaults

`func NewUpdateEnvironmentsRequestEnvironmentWithDefaults() *UpdateEnvironmentsRequestEnvironment`

NewUpdateEnvironmentsRequestEnvironmentWithDefaults instantiates a new UpdateEnvironmentsRequestEnvironment object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateEnvironmentsRequestEnvironment) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateEnvironmentsRequestEnvironment) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateEnvironmentsRequestEnvironment) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateEnvironmentsRequestEnvironment) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateEnvironmentsRequestEnvironment) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateEnvironmentsRequestEnvironment) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateEnvironmentsRequestEnvironment) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateEnvironmentsRequestEnvironment) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVisibility

`func (o *UpdateEnvironmentsRequestEnvironment) GetVisibility() string`

GetVisibility returns the Visibility field if non-nil, zero value otherwise.

### GetVisibilityOk

`func (o *UpdateEnvironmentsRequestEnvironment) GetVisibilityOk() (*string, bool)`

GetVisibilityOk returns a tuple with the Visibility field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVisibility

`func (o *UpdateEnvironmentsRequestEnvironment) SetVisibility(v string)`

SetVisibility sets Visibility field to given value.

### HasVisibility

`func (o *UpdateEnvironmentsRequestEnvironment) HasVisibility() bool`

HasVisibility returns a boolean if a field has been set.

### GetSortOrder

`func (o *UpdateEnvironmentsRequestEnvironment) GetSortOrder() int64`

GetSortOrder returns the SortOrder field if non-nil, zero value otherwise.

### GetSortOrderOk

`func (o *UpdateEnvironmentsRequestEnvironment) GetSortOrderOk() (*int64, bool)`

GetSortOrderOk returns a tuple with the SortOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSortOrder

`func (o *UpdateEnvironmentsRequestEnvironment) SetSortOrder(v int64)`

SetSortOrder sets SortOrder field to given value.

### HasSortOrder

`func (o *UpdateEnvironmentsRequestEnvironment) HasSortOrder() bool`

HasSortOrder returns a boolean if a field has been set.

### GetActive

`func (o *UpdateEnvironmentsRequestEnvironment) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *UpdateEnvironmentsRequestEnvironment) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *UpdateEnvironmentsRequestEnvironment) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *UpdateEnvironmentsRequestEnvironment) HasActive() bool`

HasActive returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


