# SystemUninitializedCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | The name of the system. | 
**Description** | Pointer to **string** | Optional description for the system. | [optional] 
**Type** | [**SystemUninitializedCreateType**](SystemUninitializedCreateType.md) |  | 
**Layout** | [**SystemUninitializedCreateLayout**](SystemUninitializedCreateLayout.md) |  | 
**Enabled** | Pointer to **bool** | Whether the system is enabled. | [optional] 
**ExternalId** | Pointer to **string** | External identifier for the system. | [optional] 
**Config** | Pointer to **map[string]interface{}** | Arbitrary configuration data for the system. | [optional] 
**Components** | Pointer to [**[]SystemUninitializedCreateComponentsInner**](SystemUninitializedCreateComponentsInner.md) | Optional component payloads to enrich skeleton components at creation time. Each entry is matched to a layout component type by &#x60;typeCode&#x60;. Unmatched entries are ignored.  | [optional] 

## Methods

### NewSystemUninitializedCreate

`func NewSystemUninitializedCreate(name string, type_ SystemUninitializedCreateType, layout SystemUninitializedCreateLayout, ) *SystemUninitializedCreate`

NewSystemUninitializedCreate instantiates a new SystemUninitializedCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemUninitializedCreateWithDefaults

`func NewSystemUninitializedCreateWithDefaults() *SystemUninitializedCreate`

NewSystemUninitializedCreateWithDefaults instantiates a new SystemUninitializedCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SystemUninitializedCreate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SystemUninitializedCreate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SystemUninitializedCreate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SystemUninitializedCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SystemUninitializedCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SystemUninitializedCreate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SystemUninitializedCreate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetType

`func (o *SystemUninitializedCreate) GetType() SystemUninitializedCreateType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SystemUninitializedCreate) GetTypeOk() (*SystemUninitializedCreateType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SystemUninitializedCreate) SetType(v SystemUninitializedCreateType)`

SetType sets Type field to given value.


### GetLayout

`func (o *SystemUninitializedCreate) GetLayout() SystemUninitializedCreateLayout`

GetLayout returns the Layout field if non-nil, zero value otherwise.

### GetLayoutOk

`func (o *SystemUninitializedCreate) GetLayoutOk() (*SystemUninitializedCreateLayout, bool)`

GetLayoutOk returns a tuple with the Layout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLayout

`func (o *SystemUninitializedCreate) SetLayout(v SystemUninitializedCreateLayout)`

SetLayout sets Layout field to given value.


### GetEnabled

`func (o *SystemUninitializedCreate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SystemUninitializedCreate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SystemUninitializedCreate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *SystemUninitializedCreate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExternalId

`func (o *SystemUninitializedCreate) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SystemUninitializedCreate) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SystemUninitializedCreate) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *SystemUninitializedCreate) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetConfig

`func (o *SystemUninitializedCreate) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SystemUninitializedCreate) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SystemUninitializedCreate) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *SystemUninitializedCreate) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetComponents

`func (o *SystemUninitializedCreate) GetComponents() []SystemUninitializedCreateComponentsInner`

GetComponents returns the Components field if non-nil, zero value otherwise.

### GetComponentsOk

`func (o *SystemUninitializedCreate) GetComponentsOk() (*[]SystemUninitializedCreateComponentsInner, bool)`

GetComponentsOk returns a tuple with the Components field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComponents

`func (o *SystemUninitializedCreate) SetComponents(v []SystemUninitializedCreateComponentsInner)`

SetComponents sets Components field to given value.

### HasComponents

`func (o *SystemUninitializedCreate) HasComponents() bool`

HasComponents returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


