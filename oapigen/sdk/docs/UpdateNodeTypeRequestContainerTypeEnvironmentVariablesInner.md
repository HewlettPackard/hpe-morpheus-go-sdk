# UpdateNodeTypeRequestContainerTypeEnvironmentVariablesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Environment variable name | 
**Value** | Pointer to **string** | Sets fixed value for variable | [optional] 
**Masked** | Pointer to **bool** | Can be used to enable / disable masking of variable | [optional] [default to false]
**Export** | Pointer to **bool** | Can be used to enable / disable export of variable | [optional] [default to false]

## Methods

### NewUpdateNodeTypeRequestContainerTypeEnvironmentVariablesInner

`func NewUpdateNodeTypeRequestContainerTypeEnvironmentVariablesInner(name string, ) *UpdateNodeTypeRequestContainerTypeEnvironmentVariablesInner`

NewUpdateNodeTypeRequestContainerTypeEnvironmentVariablesInner instantiates a new UpdateNodeTypeRequestContainerTypeEnvironmentVariablesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateNodeTypeRequestContainerTypeEnvironmentVariablesInnerWithDefaults

`func NewUpdateNodeTypeRequestContainerTypeEnvironmentVariablesInnerWithDefaults() *UpdateNodeTypeRequestContainerTypeEnvironmentVariablesInner`

NewUpdateNodeTypeRequestContainerTypeEnvironmentVariablesInnerWithDefaults instantiates a new UpdateNodeTypeRequestContainerTypeEnvironmentVariablesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateNodeTypeRequestContainerTypeEnvironmentVariablesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateNodeTypeRequestContainerTypeEnvironmentVariablesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateNodeTypeRequestContainerTypeEnvironmentVariablesInner) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *UpdateNodeTypeRequestContainerTypeEnvironmentVariablesInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateNodeTypeRequestContainerTypeEnvironmentVariablesInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateNodeTypeRequestContainerTypeEnvironmentVariablesInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *UpdateNodeTypeRequestContainerTypeEnvironmentVariablesInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetMasked

`func (o *UpdateNodeTypeRequestContainerTypeEnvironmentVariablesInner) GetMasked() bool`

GetMasked returns the Masked field if non-nil, zero value otherwise.

### GetMaskedOk

`func (o *UpdateNodeTypeRequestContainerTypeEnvironmentVariablesInner) GetMaskedOk() (*bool, bool)`

GetMaskedOk returns a tuple with the Masked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasked

`func (o *UpdateNodeTypeRequestContainerTypeEnvironmentVariablesInner) SetMasked(v bool)`

SetMasked sets Masked field to given value.

### HasMasked

`func (o *UpdateNodeTypeRequestContainerTypeEnvironmentVariablesInner) HasMasked() bool`

HasMasked returns a boolean if a field has been set.

### GetExport

`func (o *UpdateNodeTypeRequestContainerTypeEnvironmentVariablesInner) GetExport() bool`

GetExport returns the Export field if non-nil, zero value otherwise.

### GetExportOk

`func (o *UpdateNodeTypeRequestContainerTypeEnvironmentVariablesInner) GetExportOk() (*bool, bool)`

GetExportOk returns a tuple with the Export field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExport

`func (o *UpdateNodeTypeRequestContainerTypeEnvironmentVariablesInner) SetExport(v bool)`

SetExport sets Export field to given value.

### HasExport

`func (o *UpdateNodeTypeRequestContainerTypeEnvironmentVariablesInner) HasExport() bool`

HasExport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


