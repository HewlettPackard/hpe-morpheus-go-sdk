# UpdateSystemRequestSystemComponentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Existing system component id. When supplied, endpoints that support component updates match by id first. | [optional] 
**TypeCode** | Pointer to **string** | The code of the component type this entry applies to. | [optional] 
**Name** | Pointer to **string** | Optional override for the component name. Defaults to the component type name. | [optional] 
**ExternalId** | Pointer to **string** | External identifier for the component. | [optional] 
**Config** | Pointer to **map[string]interface{}** | Arbitrary configuration data for the component. | [optional] 

## Methods

### NewUpdateSystemRequestSystemComponentsInner

`func NewUpdateSystemRequestSystemComponentsInner() *UpdateSystemRequestSystemComponentsInner`

NewUpdateSystemRequestSystemComponentsInner instantiates a new UpdateSystemRequestSystemComponentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *UpdateSystemRequestSystemComponentsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UpdateSystemRequestSystemComponentsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UpdateSystemRequestSystemComponentsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *UpdateSystemRequestSystemComponentsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTypeCode

`func (o *UpdateSystemRequestSystemComponentsInner) GetTypeCode() string`

GetTypeCode returns the TypeCode field if non-nil, zero value otherwise.

### GetTypeCodeOk

`func (o *UpdateSystemRequestSystemComponentsInner) GetTypeCodeOk() (*string, bool)`

GetTypeCodeOk returns a tuple with the TypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCode

`func (o *UpdateSystemRequestSystemComponentsInner) SetTypeCode(v string)`

SetTypeCode sets TypeCode field to given value.

### HasTypeCode

`func (o *UpdateSystemRequestSystemComponentsInner) HasTypeCode() bool`

HasTypeCode returns a boolean if a field has been set.

### GetName

`func (o *UpdateSystemRequestSystemComponentsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *UpdateSystemRequestSystemComponentsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *UpdateSystemRequestSystemComponentsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *UpdateSystemRequestSystemComponentsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExternalId

`func (o *UpdateSystemRequestSystemComponentsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *UpdateSystemRequestSystemComponentsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *UpdateSystemRequestSystemComponentsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *UpdateSystemRequestSystemComponentsInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetConfig

`func (o *UpdateSystemRequestSystemComponentsInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *UpdateSystemRequestSystemComponentsInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *UpdateSystemRequestSystemComponentsInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *UpdateSystemRequestSystemComponentsInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


