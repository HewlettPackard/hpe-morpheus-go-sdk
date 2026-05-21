# SystemInitialize

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Override the system name. | [optional] 
**Description** | Pointer to **string** | Override the system description. | [optional] 
**Enabled** | Pointer to **bool** | Enable or disable the system. | [optional] 
**ExternalId** | Pointer to **string** | Override the external identifier. | [optional] 
**Config** | Pointer to **map[string]interface{}** | Override the system configuration data. | [optional] 
**Layout** | Pointer to [**SystemInitializeLayout**](SystemInitializeLayout.md) |  | [optional] 
**Components** | Pointer to [**[]SystemInitializeComponentsInner**](SystemInitializeComponentsInner.md) | Optional component overrides. Components are matched by &#x60;id&#x60; first, then by &#x60;typeCode&#x60; when exactly one existing component matches. Updates name, externalId, and config on matched components, and can create a new component when a valid &#x60;typeCode&#x60; is supplied with no existing match.  | [optional] 

## Methods

### NewSystemInitialize

`func NewSystemInitialize() *SystemInitialize`

NewSystemInitialize instantiates a new SystemInitialize object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemInitializeWithDefaults

`func NewSystemInitializeWithDefaults() *SystemInitialize`

NewSystemInitializeWithDefaults instantiates a new SystemInitialize object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SystemInitialize) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SystemInitialize) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SystemInitialize) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SystemInitialize) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *SystemInitialize) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SystemInitialize) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SystemInitialize) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SystemInitialize) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *SystemInitialize) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SystemInitialize) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SystemInitialize) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SystemInitialize) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExternalId

`func (o *SystemInitialize) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SystemInitialize) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SystemInitialize) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *SystemInitialize) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetConfig

`func (o *SystemInitialize) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SystemInitialize) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SystemInitialize) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *SystemInitialize) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetLayout

`func (o *SystemInitialize) GetLayout() SystemInitializeLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *SystemInitialize) GetLayoutOk() (*SystemInitializeLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *SystemInitialize) SetLayout(v SystemInitializeLayout)`

SetLayout sets Layout field to given value.

### HasLayout

`func (o *SystemInitialize) HasLayout() bool`

HasLayout returns a boolean if a field has been set.

### GetComponents

`func (o *SystemInitialize) GetComponents() []SystemInitializeComponentsInner`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *SystemInitialize) GetComponentsOk() (*[]SystemInitializeComponentsInner, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *SystemInitialize) SetComponents(v []SystemInitializeComponentsInner)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *SystemInitialize) HasComponents() bool`

HasComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


