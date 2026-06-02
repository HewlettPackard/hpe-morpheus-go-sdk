# ListSystems200ResponseSystemsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** | The name of the system. | [optional] 
**Description** | Pointer to **string** | Optional description for the system. | [optional] 
**Type** | Pointer to [**ListSystems200ResponseSystemsInnerType**](ListSystems200ResponseSystemsInnerType.md) |  | [optional] 
**Layout** | Pointer to [**ListSystems200ResponseSystemsInnerLayout**](ListSystems200ResponseSystemsInnerLayout.md) |  | [optional] 
**Status** | Pointer to **string** | The current status of the system. | [optional] 
**StatusMessage** | Pointer to **NullableString** | A message describing the current status. | [optional] 
**Enabled** | Pointer to **bool** | Whether the system is enabled. | [optional] 
**ExternalId** | Pointer to **string** | External ID of the system. | [optional] 
**Config** | Pointer to **map[string]interface{}** | Arbitrary configuration data for the system. | [optional] 
**Components** | Pointer to [**[]ListSystems200ResponseSystemsInnerComponentsInner**](ListSystems200ResponseSystemsInnerComponentsInner.md) | Component instances belonging to this system. | [optional] 
**DateCreated** | Pointer to **time.Time** | The date the system was created. | [optional] 
**LastUpdated** | Pointer to **time.Time** | The date the system was last updated. | [optional] 

## Methods

### NewListSystems200ResponseSystemsInner

`func NewListSystems200ResponseSystemsInner() *ListSystems200ResponseSystemsInner`

NewListSystems200ResponseSystemsInner instantiates a new ListSystems200ResponseSystemsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *ListSystems200ResponseSystemsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ListSystems200ResponseSystemsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ListSystems200ResponseSystemsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *ListSystems200ResponseSystemsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ListSystems200ResponseSystemsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ListSystems200ResponseSystemsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ListSystems200ResponseSystemsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ListSystems200ResponseSystemsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *ListSystems200ResponseSystemsInner) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ListSystems200ResponseSystemsInner) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ListSystems200ResponseSystemsInner) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ListSystems200ResponseSystemsInner) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *ListSystems200ResponseSystemsInner) GetType() ListSystems200ResponseSystemsInnerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ListSystems200ResponseSystemsInner) GetTypeOk() (*ListSystems200ResponseSystemsInnerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ListSystems200ResponseSystemsInner) SetType(v ListSystems200ResponseSystemsInnerType)`

SetType sets Type field to given value.

### HasType

`func (o *ListSystems200ResponseSystemsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLayout

`func (o *ListSystems200ResponseSystemsInner) GetLayout() ListSystems200ResponseSystemsInnerLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *ListSystems200ResponseSystemsInner) GetLayoutOk() (*ListSystems200ResponseSystemsInnerLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *ListSystems200ResponseSystemsInner) SetLayout(v ListSystems200ResponseSystemsInnerLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *ListSystems200ResponseSystemsInner) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetStatus

`func (o *ListSystems200ResponseSystemsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ListSystems200ResponseSystemsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ListSystems200ResponseSystemsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ListSystems200ResponseSystemsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *ListSystems200ResponseSystemsInner) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *ListSystems200ResponseSystemsInner) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *ListSystems200ResponseSystemsInner) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *ListSystems200ResponseSystemsInner) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### SetStatusMessageNil

`func (o *ListSystems200ResponseSystemsInner) SetStatusMessageNil(b bool)`

 SetStatusMessageNil sets the value for StatusMessage to be an explicit nil

### UnsetStatusMessage
`func (o *ListSystems200ResponseSystemsInner) UnsetStatusMessage()`

UnsetStatusMessage ensures that no value is present for StatusMessage, not even an explicit nil
### GetEnabled

`func (o *ListSystems200ResponseSystemsInner) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ListSystems200ResponseSystemsInner) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ListSystems200ResponseSystemsInner) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ListSystems200ResponseSystemsInner) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExternalId

`func (o *ListSystems200ResponseSystemsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *ListSystems200ResponseSystemsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *ListSystems200ResponseSystemsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *ListSystems200ResponseSystemsInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetConfig

`func (o *ListSystems200ResponseSystemsInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *ListSystems200ResponseSystemsInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *ListSystems200ResponseSystemsInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *ListSystems200ResponseSystemsInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetComponents

`func (o *ListSystems200ResponseSystemsInner) GetComponents() []ListSystems200ResponseSystemsInnerComponentsInner`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *ListSystems200ResponseSystemsInner) GetComponentsOk() (*[]ListSystems200ResponseSystemsInnerComponentsInner, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *ListSystems200ResponseSystemsInner) SetComponents(v []ListSystems200ResponseSystemsInnerComponentsInner)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *ListSystems200ResponseSystemsInner) HasComponents() bool`

HasComponents returns a boolean if a field has been set.

### GetDateCreated

`func (o *ListSystems200ResponseSystemsInner) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *ListSystems200ResponseSystemsInner) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *ListSystems200ResponseSystemsInner) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *ListSystems200ResponseSystemsInner) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *ListSystems200ResponseSystemsInner) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *ListSystems200ResponseSystemsInner) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *ListSystems200ResponseSystemsInner) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *ListSystems200ResponseSystemsInner) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


