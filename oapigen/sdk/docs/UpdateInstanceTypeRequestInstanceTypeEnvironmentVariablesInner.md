# UpdateInstanceTypeRequestInstanceTypeEnvironmentVariablesInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | Environment variable name | 
**Value** | Pointer to **string** | Sets fixed value for variable | [optional] 
**Masked** | Pointer to **bool** | Can be used to enable / disable masking of variable | [optional] [default to false]
**Export** | Pointer to **bool** | Can be used to enable / disable export of variable | [optional] [default to false]

## Methods

### NewUpdateInstanceTypeRequestInstanceTypeEnvironmentVariablesInner

`func NewUpdateInstanceTypeRequestInstanceTypeEnvironmentVariablesInner(name string, ) *UpdateInstanceTypeRequestInstanceTypeEnvironmentVariablesInner`

NewUpdateInstanceTypeRequestInstanceTypeEnvironmentVariablesInner instantiates a new UpdateInstanceTypeRequestInstanceTypeEnvironmentVariablesInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInstanceTypeRequestInstanceTypeEnvironmentVariablesInnerWithDefaults

`func NewUpdateInstanceTypeRequestInstanceTypeEnvironmentVariablesInnerWithDefaults() *UpdateInstanceTypeRequestInstanceTypeEnvironmentVariablesInner`

NewUpdateInstanceTypeRequestInstanceTypeEnvironmentVariablesInnerWithDefaults instantiates a new UpdateInstanceTypeRequestInstanceTypeEnvironmentVariablesInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateInstanceTypeRequestInstanceTypeEnvironmentVariablesInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateInstanceTypeRequestInstanceTypeEnvironmentVariablesInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateInstanceTypeRequestInstanceTypeEnvironmentVariablesInner) SetName(v string)`

SetName sets Name field to given value.


### GetValue

`func (o *UpdateInstanceTypeRequestInstanceTypeEnvironmentVariablesInner) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *UpdateInstanceTypeRequestInstanceTypeEnvironmentVariablesInner) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *UpdateInstanceTypeRequestInstanceTypeEnvironmentVariablesInner) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *UpdateInstanceTypeRequestInstanceTypeEnvironmentVariablesInner) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetMasked

`func (o *UpdateInstanceTypeRequestInstanceTypeEnvironmentVariablesInner) GetMasked() bool`

GetMasked returns the Masked field if non-nil, zero value otherwise.

### GetMaskedOk

`func (o *UpdateInstanceTypeRequestInstanceTypeEnvironmentVariablesInner) GetMaskedOk() (*bool, bool)`

GetMaskedOk returns a tuple with the Masked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasked

`func (o *UpdateInstanceTypeRequestInstanceTypeEnvironmentVariablesInner) SetMasked(v bool)`

SetMasked sets Masked field to given value.

### HasMasked

`func (o *UpdateInstanceTypeRequestInstanceTypeEnvironmentVariablesInner) HasMasked() bool`

HasMasked returns a boolean if a field has been set.

### GetExport

`func (o *UpdateInstanceTypeRequestInstanceTypeEnvironmentVariablesInner) GetExport() bool`

GetExport returns the Export field if non-nil, zero value otherwise.

### GetExportOk

`func (o *UpdateInstanceTypeRequestInstanceTypeEnvironmentVariablesInner) GetExportOk() (*bool, bool)`

GetExportOk returns a tuple with the Export field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExport

`func (o *UpdateInstanceTypeRequestInstanceTypeEnvironmentVariablesInner) SetExport(v bool)`

SetExport sets Export field to given value.

### HasExport

`func (o *UpdateInstanceTypeRequestInstanceTypeEnvironmentVariablesInner) HasExport() bool`

HasExport returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


