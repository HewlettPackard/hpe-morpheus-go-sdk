# GetApprovalsItem200ResponseApprovalItemDetailsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | Pointer to **string** | The category of the change (e.g. network, volume, plan, planMemory, planCores) | [optional] 
**Type** | Pointer to **string** | The type of change (e.g. increase, decrease, add, remove, change) | [optional] 
**Name** | Pointer to **string** | The name of the item being changed | [optional] 
**FromValue** | Pointer to **NullableString** | The original value before the change | [optional] 
**ToValue** | Pointer to **NullableString** | The new value after the change | [optional] 

## Methods

### NewGetApprovalsItem200ResponseApprovalItemDetailsInner

`func NewGetApprovalsItem200ResponseApprovalItemDetailsInner() *GetApprovalsItem200ResponseApprovalItemDetailsInner`

NewGetApprovalsItem200ResponseApprovalItemDetailsInner instantiates a new GetApprovalsItem200ResponseApprovalItemDetailsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetCategory

`func (o *GetApprovalsItem200ResponseApprovalItemDetailsInner) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *GetApprovalsItem200ResponseApprovalItemDetailsInner) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *GetApprovalsItem200ResponseApprovalItemDetailsInner) SetCategory(v string)`

SetCategory sets Category field to given value.

### HasCategory

`func (o *GetApprovalsItem200ResponseApprovalItemDetailsInner) HasCategory() bool`

HasCategory returns a boolean if a field has been set.

### GetType

`func (o *GetApprovalsItem200ResponseApprovalItemDetailsInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetApprovalsItem200ResponseApprovalItemDetailsInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetApprovalsItem200ResponseApprovalItemDetailsInner) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *GetApprovalsItem200ResponseApprovalItemDetailsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetName

`func (o *GetApprovalsItem200ResponseApprovalItemDetailsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetApprovalsItem200ResponseApprovalItemDetailsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetApprovalsItem200ResponseApprovalItemDetailsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetApprovalsItem200ResponseApprovalItemDetailsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetFromValue

`func (o *GetApprovalsItem200ResponseApprovalItemDetailsInner) GetFromValue() string`

GetFromValue returns the FromValue field if non-nil, zero value otherwise.

### GetFromValueOk

`func (o *GetApprovalsItem200ResponseApprovalItemDetailsInner) GetFromValueOk() (*string, bool)`

GetFromValueOk returns a tuple with the FromValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromValue

`func (o *GetApprovalsItem200ResponseApprovalItemDetailsInner) SetFromValue(v string)`

SetFromValue sets FromValue field to given value.

### HasFromValue

`func (o *GetApprovalsItem200ResponseApprovalItemDetailsInner) HasFromValue() bool`

HasFromValue returns a boolean if a field has been set.

### SetFromValueNil

`func (o *GetApprovalsItem200ResponseApprovalItemDetailsInner) SetFromValueNil(b bool)`

 SetFromValueNil sets the value for FromValue to be an explicit nil

### UnsetFromValue
`func (o *GetApprovalsItem200ResponseApprovalItemDetailsInner) UnsetFromValue()`

UnsetFromValue ensures that no value is present for FromValue, not even an explicit nil
### GetToValue

`func (o *GetApprovalsItem200ResponseApprovalItemDetailsInner) GetToValue() string`

GetToValue returns the ToValue field if non-nil, zero value otherwise.

### GetToValueOk

`func (o *GetApprovalsItem200ResponseApprovalItemDetailsInner) GetToValueOk() (*string, bool)`

GetToValueOk returns a tuple with the ToValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetToValue

`func (o *GetApprovalsItem200ResponseApprovalItemDetailsInner) SetToValue(v string)`

SetToValue sets ToValue field to given value.

### HasToValue

`func (o *GetApprovalsItem200ResponseApprovalItemDetailsInner) HasToValue() bool`

HasToValue returns a boolean if a field has been set.

### SetToValueNil

`func (o *GetApprovalsItem200ResponseApprovalItemDetailsInner) SetToValueNil(b bool)`

 SetToValueNil sets the value for ToValue to be an explicit nil

### UnsetToValue
`func (o *GetApprovalsItem200ResponseApprovalItemDetailsInner) UnsetToValue()`

UnsetToValue ensures that no value is present for ToValue, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


