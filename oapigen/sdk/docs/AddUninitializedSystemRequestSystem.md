# AddUninitializedSystemRequestSystem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the system. | 
**Description** | Pointer to **string** | Optional description for the system. | [optional] 
**Type** | [**AddUninitializedSystemRequestSystemType**](AddUninitializedSystemRequestSystemType.md) |  | 
**Layout** | [**AddUninitializedSystemRequestSystemLayout**](AddUninitializedSystemRequestSystemLayout.md) |  | 
**Enabled** | Pointer to **bool** | Whether the system is enabled. | [optional] 
**ExternalId** | Pointer to **string** | External identifier for the system. | [optional] 
**Config** | Pointer to **map[string]interface{}** | Arbitrary configuration data for the system. | [optional] 
**Components** | Pointer to [**[]AddUninitializedSystemRequestSystemComponentsInner**](AddUninitializedSystemRequestSystemComponentsInner.md) | Optional component payloads to enrich skeleton components at creation time. Each entry is matched to a layout component type by &#x60;typeCode&#x60;. Unmatched entries are ignored.  | [optional] 

## Methods

### NewAddUninitializedSystemRequestSystem

`func NewAddUninitializedSystemRequestSystem(name string, type_ AddUninitializedSystemRequestSystemType, layout AddUninitializedSystemRequestSystemLayout, ) *AddUninitializedSystemRequestSystem`

NewAddUninitializedSystemRequestSystem instantiates a new AddUninitializedSystemRequestSystem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetName

`func (o *AddUninitializedSystemRequestSystem) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddUninitializedSystemRequestSystem) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddUninitializedSystemRequestSystem) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *AddUninitializedSystemRequestSystem) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AddUninitializedSystemRequestSystem) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AddUninitializedSystemRequestSystem) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AddUninitializedSystemRequestSystem) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *AddUninitializedSystemRequestSystem) GetType() AddUninitializedSystemRequestSystemType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddUninitializedSystemRequestSystem) GetTypeOk() (*AddUninitializedSystemRequestSystemType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddUninitializedSystemRequestSystem) SetType(v AddUninitializedSystemRequestSystemType)`

SetType sets Type field to given value.


### GetLayout

`func (o *AddUninitializedSystemRequestSystem) GetLayout() AddUninitializedSystemRequestSystemLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *AddUninitializedSystemRequestSystem) GetLayoutOk() (*AddUninitializedSystemRequestSystemLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *AddUninitializedSystemRequestSystem) SetLayout(v AddUninitializedSystemRequestSystemLayout)`

SetLayout sets Layout field to given value.


### GetEnabled

`func (o *AddUninitializedSystemRequestSystem) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *AddUninitializedSystemRequestSystem) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *AddUninitializedSystemRequestSystem) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *AddUninitializedSystemRequestSystem) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExternalId

`func (o *AddUninitializedSystemRequestSystem) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AddUninitializedSystemRequestSystem) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AddUninitializedSystemRequestSystem) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AddUninitializedSystemRequestSystem) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetConfig

`func (o *AddUninitializedSystemRequestSystem) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddUninitializedSystemRequestSystem) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddUninitializedSystemRequestSystem) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddUninitializedSystemRequestSystem) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetComponents

`func (o *AddUninitializedSystemRequestSystem) GetComponents() []AddUninitializedSystemRequestSystemComponentsInner`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *AddUninitializedSystemRequestSystem) GetComponentsOk() (*[]AddUninitializedSystemRequestSystemComponentsInner, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *AddUninitializedSystemRequestSystem) SetComponents(v []AddUninitializedSystemRequestSystemComponentsInner)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *AddUninitializedSystemRequestSystem) HasComponents() bool`

HasComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


