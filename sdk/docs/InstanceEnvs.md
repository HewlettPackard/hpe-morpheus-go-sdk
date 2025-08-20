# InstanceEnvs

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Envs** | Pointer to [**[]GetEnvVariables200ResponseEnvsInner**](GetEnvVariables200ResponseEnvsInner.md) |  | [optional] 
**ReadOnlyEnvs** | Pointer to [**map[string]GetEnvVariables200ResponseReadOnlyEnvsValue**](GetEnvVariables200ResponseReadOnlyEnvsValue.md) |  | [optional] 
**ImportedEnvs** | Pointer to [**map[string]GetEnvVariables200ResponseReadOnlyEnvsValue**](GetEnvVariables200ResponseReadOnlyEnvsValue.md) |  | [optional] 

## Methods

### NewInstanceEnvs

`func NewInstanceEnvs() *InstanceEnvs`

NewInstanceEnvs instantiates a new InstanceEnvs object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceEnvsWithDefaults

`func NewInstanceEnvsWithDefaults() *InstanceEnvs`

NewInstanceEnvsWithDefaults instantiates a new InstanceEnvs object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnvs

`func (o *InstanceEnvs) GetEnvs() []GetEnvVariables200ResponseEnvsInner`

GetEnvs returns the Envs field if non-nil, zero value otherwise.

### GetEnvsOk

`func (o *InstanceEnvs) GetEnvsOk() (*[]GetEnvVariables200ResponseEnvsInner, bool)`

GetEnvsOk returns a tuple with the Envs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnvs

`func (o *InstanceEnvs) SetEnvs(v []GetEnvVariables200ResponseEnvsInner)`

SetEnvs sets Envs field to given value.

### HasEnvs

`func (o *InstanceEnvs) HasEnvs() bool`

HasEnvs returns a boolean if a field has been set.

### SetEnvsNil

`func (o *InstanceEnvs) SetEnvsNil(b bool)`

 SetEnvsNil sets the value for Envs to be an explicit nil

### UnsetEnvs
`func (o *InstanceEnvs) UnsetEnvs()`

UnsetEnvs ensures that no value is present for Envs, not even an explicit nil
### GetReadOnlyEnvs

`func (o *InstanceEnvs) GetReadOnlyEnvs() map[string]GetEnvVariables200ResponseReadOnlyEnvsValue`

GetReadOnlyEnvs returns the ReadOnlyEnvs field if non-nil, zero value otherwise.

### GetReadOnlyEnvsOk

`func (o *InstanceEnvs) GetReadOnlyEnvsOk() (*map[string]GetEnvVariables200ResponseReadOnlyEnvsValue, bool)`

GetReadOnlyEnvsOk returns a tuple with the ReadOnlyEnvs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReadOnlyEnvs

`func (o *InstanceEnvs) SetReadOnlyEnvs(v map[string]GetEnvVariables200ResponseReadOnlyEnvsValue)`

SetReadOnlyEnvs sets ReadOnlyEnvs field to given value.

### HasReadOnlyEnvs

`func (o *InstanceEnvs) HasReadOnlyEnvs() bool`

HasReadOnlyEnvs returns a boolean if a field has been set.

### GetImportedEnvs

`func (o *InstanceEnvs) GetImportedEnvs() map[string]GetEnvVariables200ResponseReadOnlyEnvsValue`

GetImportedEnvs returns the ImportedEnvs field if non-nil, zero value otherwise.

### GetImportedEnvsOk

`func (o *InstanceEnvs) GetImportedEnvsOk() (*map[string]GetEnvVariables200ResponseReadOnlyEnvsValue, bool)`

GetImportedEnvsOk returns a tuple with the ImportedEnvs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImportedEnvs

`func (o *InstanceEnvs) SetImportedEnvs(v map[string]GetEnvVariables200ResponseReadOnlyEnvsValue)`

SetImportedEnvs sets ImportedEnvs field to given value.

### HasImportedEnvs

`func (o *InstanceEnvs) HasImportedEnvs() bool`

HasImportedEnvs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


