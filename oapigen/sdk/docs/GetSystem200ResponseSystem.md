# GetSystem200ResponseSystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** | The name of the system. | [optional] 
**Description** | Pointer to **string** | Optional description for the system. | [optional] 
**Type** | Pointer to [**GetSystem200ResponseSystemType**](GetSystem200ResponseSystemType.md) |  | [optional] 
**Layout** | Pointer to [**GetSystem200ResponseSystemLayout**](GetSystem200ResponseSystemLayout.md) |  | [optional] 
**Status** | Pointer to **string** | The current status of the system. | [optional] 
**StatusMessage** | Pointer to **NullableString** | A message describing the current status. | [optional] 
**Enabled** | Pointer to **bool** | Whether the system is enabled. | [optional] 
**ExternalId** | Pointer to **string** | External ID of the system. | [optional] 
**Config** | Pointer to **map[string]interface{}** | Arbitrary configuration data for the system. | [optional] 
**Components** | Pointer to [**[]GetSystem200ResponseSystemComponentsInner**](GetSystem200ResponseSystemComponentsInner.md) | Component instances belonging to this system. | [optional] 
**DateCreated** | Pointer to **time.Time** | The date the system was created. | [optional] 
**LastUpdated** | Pointer to **time.Time** | The date the system was last updated. | [optional] 

## Methods

### NewGetSystem200ResponseSystem

`func NewGetSystem200ResponseSystem() *GetSystem200ResponseSystem`

NewGetSystem200ResponseSystem instantiates a new GetSystem200ResponseSystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetSystem200ResponseSystem) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSystem200ResponseSystem) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSystem200ResponseSystem) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetSystem200ResponseSystem) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetSystem200ResponseSystem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetSystem200ResponseSystem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetSystem200ResponseSystem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetSystem200ResponseSystem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *GetSystem200ResponseSystem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GetSystem200ResponseSystem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GetSystem200ResponseSystem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GetSystem200ResponseSystem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *GetSystem200ResponseSystem) GetType() GetSystem200ResponseSystemType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetSystem200ResponseSystem) GetTypeOk() (*GetSystem200ResponseSystemType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetSystem200ResponseSystem) SetType(v GetSystem200ResponseSystemType)`

SetType sets Type field to given value.

### HasType

`func (o *GetSystem200ResponseSystem) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLayout

`func (o *GetSystem200ResponseSystem) GetLayout() GetSystem200ResponseSystemLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *GetSystem200ResponseSystem) GetLayoutOk() (*GetSystem200ResponseSystemLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *GetSystem200ResponseSystem) SetLayout(v GetSystem200ResponseSystemLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *GetSystem200ResponseSystem) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetStatus

`func (o *GetSystem200ResponseSystem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSystem200ResponseSystem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSystem200ResponseSystem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetSystem200ResponseSystem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *GetSystem200ResponseSystem) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *GetSystem200ResponseSystem) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *GetSystem200ResponseSystem) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *GetSystem200ResponseSystem) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *GetSystem200ResponseSystem) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *GetSystem200ResponseSystem) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetEnabled

`func (o *GetSystem200ResponseSystem) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *GetSystem200ResponseSystem) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *GetSystem200ResponseSystem) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *GetSystem200ResponseSystem) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExternalId

`func (o *GetSystem200ResponseSystem) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetSystem200ResponseSystem) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetSystem200ResponseSystem) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetSystem200ResponseSystem) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetConfig

`func (o *GetSystem200ResponseSystem) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetSystem200ResponseSystem) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetSystem200ResponseSystem) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetSystem200ResponseSystem) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetComponents

`func (o *GetSystem200ResponseSystem) GetComponents() []GetSystem200ResponseSystemComponentsInner`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *GetSystem200ResponseSystem) GetComponentsOk() (*[]GetSystem200ResponseSystemComponentsInner, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *GetSystem200ResponseSystem) SetComponents(v []GetSystem200ResponseSystemComponentsInner)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *GetSystem200ResponseSystem) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetSystem200ResponseSystem) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetSystem200ResponseSystem) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetSystem200ResponseSystem) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetSystem200ResponseSystem) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetSystem200ResponseSystem) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetSystem200ResponseSystem) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetSystem200ResponseSystem) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetSystem200ResponseSystem) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


