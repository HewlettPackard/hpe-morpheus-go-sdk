# SystemUpdateComponentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Existing system component id. When supplied, endpoints that support component updates match by id first. | [optional] 
**TypeCode** | Pointer to **string** | The code of the component type this entry applies to. | [optional] 
**Name** | Pointer to **string** | Optional override for the component name. Defaults to the component type name. | [optional] 
**ExternalId** | Pointer to **string** | External identifier for the component. | [optional] 
**Config** | Pointer to **map[string]interface{}** | Arbitrary configuration data for the component. | [optional] 

## Methods

### NewSystemUpdateComponentsInner

`func NewSystemUpdateComponentsInner() *SystemUpdateComponentsInner`

NewSystemUpdateComponentsInner instantiates a new SystemUpdateComponentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSystemUpdateComponentsInnerWithDefaults

`func NewSystemUpdateComponentsInnerWithDefaults() *SystemUpdateComponentsInner`

NewSystemUpdateComponentsInnerWithDefaults instantiates a new SystemUpdateComponentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SystemUpdateComponentsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SystemUpdateComponentsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SystemUpdateComponentsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *SystemUpdateComponentsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTypeCode

`func (o *SystemUpdateComponentsInner) GetTypeCode() string`

GetTypeCode returns the TypeCode field if non-nil, zero value otherwise.

### GetTypeCodeOk

`func (o *SystemUpdateComponentsInner) GetTypeCodeOk() (*string, bool)`

GetTypeCodeOk returns a tuple with the TypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCode

`func (o *SystemUpdateComponentsInner) SetTypeCode(v string)`

SetTypeCode sets TypeCode field to given value.

### HasTypeCode

`func (o *SystemUpdateComponentsInner) HasTypeCode() bool`

HasTypeCode returns a boolean if a field has been set.

### GetName

`func (o *SystemUpdateComponentsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SystemUpdateComponentsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SystemUpdateComponentsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *SystemUpdateComponentsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExternalId

`func (o *SystemUpdateComponentsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *SystemUpdateComponentsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *SystemUpdateComponentsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *SystemUpdateComponentsInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetConfig

`func (o *SystemUpdateComponentsInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *SystemUpdateComponentsInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *SystemUpdateComponentsInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *SystemUpdateComponentsInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


