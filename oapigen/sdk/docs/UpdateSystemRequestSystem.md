# UpdateSystemRequestSystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | New name for the system. | [optional] 
**Description** | Pointer to **string** | New description for the system. | [optional] 
**Enabled** | Pointer to **bool** | Enable or disable the system. | [optional] 
**ExternalId** | Pointer to **string** | Override the external identifier. | [optional] 
**Config** | Pointer to **map[string]interface{}** | Override the system configuration data. | [optional] 
**Status** | Pointer to **string** | Override the system status. | [optional] 
**StatusMessage** | Pointer to **string** | Override the status message. | [optional] 
**Layout** | Pointer to [**UpdateSystemRequestSystemLayout**](UpdateSystemRequestSystemLayout.md) |  | [optional] 
**Components** | Pointer to [**[]UpdateSystemRequestSystemComponentsInner**](UpdateSystemRequestSystemComponentsInner.md) | Optional authoritative component payloads. Components are matched by &#x60;id&#x60; first, then by &#x60;typeCode&#x60; when exactly one existing component matches. New components can be created by supplying a valid &#x60;typeCode&#x60;. When this field is present, omitted existing components are removed from the system.  | [optional] 

## Methods

### NewUpdateSystemRequestSystem

`func NewUpdateSystemRequestSystem() *UpdateSystemRequestSystem`

NewUpdateSystemRequestSystem instantiates a new UpdateSystemRequestSystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSystemRequestSystemWithDefaults

`func NewUpdateSystemRequestSystemWithDefaults() *UpdateSystemRequestSystem`

NewUpdateSystemRequestSystemWithDefaults instantiates a new UpdateSystemRequestSystem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *UpdateSystemRequestSystem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSystemRequestSystem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSystemRequestSystem) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateSystemRequestSystem) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *UpdateSystemRequestSystem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UpdateSystemRequestSystem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UpdateSystemRequestSystem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UpdateSystemRequestSystem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateSystemRequestSystem) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateSystemRequestSystem) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateSystemRequestSystem) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateSystemRequestSystem) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExternalId

`func (o *UpdateSystemRequestSystem) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdateSystemRequestSystem) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdateSystemRequestSystem) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdateSystemRequestSystem) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateSystemRequestSystem) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateSystemRequestSystem) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateSystemRequestSystem) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateSystemRequestSystem) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetStatus

`func (o *UpdateSystemRequestSystem) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *UpdateSystemRequestSystem) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *UpdateSystemRequestSystem) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *UpdateSystemRequestSystem) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetStatusMessage

`func (o *UpdateSystemRequestSystem) GetStatusMessage() string`

GetStatusMessage returns the StatusMessage field if non-nil, zero value otherwise.

### GetStatusMessageOk

`func (o *UpdateSystemRequestSystem) GetStatusMessageOk() (*string, bool)`

GetStatusMessageOk returns a tuple with the StatusMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusMessage

`func (o *UpdateSystemRequestSystem) SetStatusMessage(v string)`

SetStatusMessage sets StatusMessage field to given value.

### HasStatusMessage

`func (o *UpdateSystemRequestSystem) HasStatusMessage() bool`

HasStatusMessage returns a boolean if a field has been set.

### GetLayout

`func (o *UpdateSystemRequestSystem) GetLayout() UpdateSystemRequestSystemLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *UpdateSystemRequestSystem) GetLayoutOk() (*UpdateSystemRequestSystemLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *UpdateSystemRequestSystem) SetLayout(v UpdateSystemRequestSystemLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *UpdateSystemRequestSystem) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetComponents

`func (o *UpdateSystemRequestSystem) GetComponents() []UpdateSystemRequestSystemComponentsInner`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *UpdateSystemRequestSystem) GetComponentsOk() (*[]UpdateSystemRequestSystemComponentsInner, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *UpdateSystemRequestSystem) SetComponents(v []UpdateSystemRequestSystemComponentsInner)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *UpdateSystemRequestSystem) HasComponents() bool`

HasComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


