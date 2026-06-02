# SystemCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the system. | 
**Description** | Pointer to **string** | Optional description for the system. | [optional] 
**Type** | [**SystemCreateType**](SystemCreateType.md) |  | 
**Layout** | [**SystemCreateLayout**](SystemCreateLayout.md) |  | 
**Enabled** | Pointer to **bool** | Whether the system is enabled. | [optional] 
**ExternalId** | Pointer to **string** | External identifier for the system. | [optional] 
**Config** | Pointer to **map[string]interface{}** | Arbitrary configuration data for the system. | [optional] 
**Components** | Pointer to [**[]SystemCreateComponentsInner**](SystemCreateComponentsInner.md) | Optional component payloads for create. Existing components are matched by &#x60;id&#x60; first, then by &#x60;typeCode&#x60; when exactly one existing component matches. New components can be created by supplying a valid &#x60;typeCode&#x60;.  | [optional] 

## Methods

### NewSystemCreate

`func NewSystemCreate(name string, type_ SystemCreateType, layout SystemCreateLayout, ) *SystemCreate`

NewSystemCreate instantiates a new SystemCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetName

`func (o *SystemCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SystemCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SystemCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SystemCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SystemCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SystemCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SystemCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *SystemCreate) GetType() SystemCreateType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SystemCreate) GetTypeOk() (*SystemCreateType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SystemCreate) SetType(v SystemCreateType)`

SetType sets Type field to given value.


### GetLayout

`func (o *SystemCreate) GetLayout() SystemCreateLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *SystemCreate) GetLayoutOk() (*SystemCreateLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *SystemCreate) SetLayout(v SystemCreateLayout)`

SetLayout sets Layout field to given value.


### GetEnabled

`func (o *SystemCreate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SystemCreate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SystemCreate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SystemCreate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExternalId

`func (o *SystemCreate) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SystemCreate) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SystemCreate) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *SystemCreate) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetConfig

`func (o *SystemCreate) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SystemCreate) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SystemCreate) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *SystemCreate) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetComponents

`func (o *SystemCreate) GetComponents() []SystemCreateComponentsInner`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *SystemCreate) GetComponentsOk() (*[]SystemCreateComponentsInner, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *SystemCreate) SetComponents(v []SystemCreateComponentsInner)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *SystemCreate) HasComponents() bool`

HasComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


