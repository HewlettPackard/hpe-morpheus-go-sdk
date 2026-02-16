# UpdateLayoutRequestInstanceTypeLayoutEnvironmentVariablesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Environment variable name | 
**Value** | Pointer to **string** | Sets fixed value for variable | [optional] 
**Masked** | Pointer to **bool** | Can be used to enable / disable masking of variable | [optional] [default to false]
**Export** | Pointer to **bool** | Can be used to enable / disable export of variable | [optional] [default to false]

## Methods

### NewUpdateLayoutRequestInstanceTypeLayoutEnvironmentVariablesInner

`func NewUpdateLayoutRequestInstanceTypeLayoutEnvironmentVariablesInner(name string, ) *UpdateLayoutRequestInstanceTypeLayoutEnvironmentVariablesInner`

NewUpdateLayoutRequestInstanceTypeLayoutEnvironmentVariablesInner instantiates a new UpdateLayoutRequestInstanceTypeLayoutEnvironmentVariablesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateLayoutRequestInstanceTypeLayoutEnvironmentVariablesInnerWithDefaults

`func NewUpdateLayoutRequestInstanceTypeLayoutEnvironmentVariablesInnerWithDefaults() *UpdateLayoutRequestInstanceTypeLayoutEnvironmentVariablesInner`

NewUpdateLayoutRequestInstanceTypeLayoutEnvironmentVariablesInnerWithDefaults instantiates a new UpdateLayoutRequestInstanceTypeLayoutEnvironmentVariablesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateLayoutRequestInstanceTypeLayoutEnvironmentVariablesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateLayoutRequestInstanceTypeLayoutEnvironmentVariablesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateLayoutRequestInstanceTypeLayoutEnvironmentVariablesInner) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *UpdateLayoutRequestInstanceTypeLayoutEnvironmentVariablesInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateLayoutRequestInstanceTypeLayoutEnvironmentVariablesInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateLayoutRequestInstanceTypeLayoutEnvironmentVariablesInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *UpdateLayoutRequestInstanceTypeLayoutEnvironmentVariablesInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetMasked

`func (o *UpdateLayoutRequestInstanceTypeLayoutEnvironmentVariablesInner) GetMasked() bool`

GetMasked returns the Masked field if non-nil, zero value otherwise.

### GetMaskedOk

`func (o *UpdateLayoutRequestInstanceTypeLayoutEnvironmentVariablesInner) GetMaskedOk() (*bool, bool)`

GetMaskedOk returns a tuple with the Masked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasked

`func (o *UpdateLayoutRequestInstanceTypeLayoutEnvironmentVariablesInner) SetMasked(v bool)`

SetMasked sets Masked field to given value.

### HasMasked

`func (o *UpdateLayoutRequestInstanceTypeLayoutEnvironmentVariablesInner) HasMasked() bool`

HasMasked returns a boolean if a field has been set.

### GetExport

`func (o *UpdateLayoutRequestInstanceTypeLayoutEnvironmentVariablesInner) GetExport() bool`

GetExport returns the Export field if non-nil, zero value otherwise.

### GetExportOk

`func (o *UpdateLayoutRequestInstanceTypeLayoutEnvironmentVariablesInner) GetExportOk() (*bool, bool)`

GetExportOk returns a tuple with the Export field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExport

`func (o *UpdateLayoutRequestInstanceTypeLayoutEnvironmentVariablesInner) SetExport(v bool)`

SetExport sets Export field to given value.

### HasExport

`func (o *UpdateLayoutRequestInstanceTypeLayoutEnvironmentVariablesInner) HasExport() bool`

HasExport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


