# ApprovalItemApprovalItemDetailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | The category of the change (e.g. network, volume, plan, planMemory, planCores) | [optional] 
**Type** | Pointer to **string** | The type of change (e.g. increase, decrease, add, remove, change) | [optional] 
**Name** | Pointer to **string** | The name of the item being changed | [optional] 
**FromValue** | Pointer to **NullableString** | The original value before the change | [optional] 
**ToValue** | Pointer to **NullableString** | The new value after the change | [optional] 

## Methods

### NewApprovalItemApprovalItemDetailsInner

`func NewApprovalItemApprovalItemDetailsInner() *ApprovalItemApprovalItemDetailsInner`

NewApprovalItemApprovalItemDetailsInner instantiates a new ApprovalItemApprovalItemDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApprovalItemApprovalItemDetailsInnerWithDefaults

`func NewApprovalItemApprovalItemDetailsInnerWithDefaults() *ApprovalItemApprovalItemDetailsInner`

NewApprovalItemApprovalItemDetailsInnerWithDefaults instantiates a new ApprovalItemApprovalItemDetailsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *ApprovalItemApprovalItemDetailsInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *ApprovalItemApprovalItemDetailsInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *ApprovalItemApprovalItemDetailsInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *ApprovalItemApprovalItemDetailsInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetType

`func (o *ApprovalItemApprovalItemDetailsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ApprovalItemApprovalItemDetailsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ApprovalItemApprovalItemDetailsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *ApprovalItemApprovalItemDetailsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *ApprovalItemApprovalItemDetailsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ApprovalItemApprovalItemDetailsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ApprovalItemApprovalItemDetailsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ApprovalItemApprovalItemDetailsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFromValue

`func (o *ApprovalItemApprovalItemDetailsInner) GetFromValue() string`

GetFromValue returns the FromValue field if non-nil, zero value otherwise.

### GetFromValueOk

`func (o *ApprovalItemApprovalItemDetailsInner) GetFromValueOk() (*string, bool)`

GetFromValueOk returns a tuple with the FromValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromValue

`func (o *ApprovalItemApprovalItemDetailsInner) SetFromValue(v string)`

SetFromValue sets FromValue field to given value.

### HasFromValue

`func (o *ApprovalItemApprovalItemDetailsInner) HasFromValue() bool`

HasFromValue returns a boolean if a field has been set.

### SetFromValueNil

`func (o *ApprovalItemApprovalItemDetailsInner) SetFromValueNil(b bool)`

 SetFromValueNil sets the value for FromValue to be an explicit nil

### UnsetFromValue
`func (o *ApprovalItemApprovalItemDetailsInner) UnsetFromValue()`

UnsetFromValue ensures that no value is present for FromValue, not even an explicit nil
### GetToValue

`func (o *ApprovalItemApprovalItemDetailsInner) GetToValue() string`

GetToValue returns the ToValue field if non-nil, zero value otherwise.

### GetToValueOk

`func (o *ApprovalItemApprovalItemDetailsInner) GetToValueOk() (*string, bool)`

GetToValueOk returns a tuple with the ToValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToValue

`func (o *ApprovalItemApprovalItemDetailsInner) SetToValue(v string)`

SetToValue sets ToValue field to given value.

### HasToValue

`func (o *ApprovalItemApprovalItemDetailsInner) HasToValue() bool`

HasToValue returns a boolean if a field has been set.

### SetToValueNil

`func (o *ApprovalItemApprovalItemDetailsInner) SetToValueNil(b bool)`

 SetToValueNil sets the value for ToValue to be an explicit nil

### UnsetToValue
`func (o *ApprovalItemApprovalItemDetailsInner) UnsetToValue()`

UnsetToValue ensures that no value is present for ToValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


