# AddUninitializedSystemRequestSystemComponentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** | Existing system component id. When supplied, endpoints that support component updates match by id first. | [optional] 
**TypeCode** | Pointer to **string** | The code of the component type this entry applies to. | [optional] 
**Name** | Pointer to **string** | Optional override for the component name. Defaults to the component type name. | [optional] 
**ExternalId** | Pointer to **string** | External identifier for the component. | [optional] 
**Config** | Pointer to **map[string]interface{}** | Arbitrary configuration data for the component. | [optional] 

## Methods

### NewAddUninitializedSystemRequestSystemComponentsInner

`func NewAddUninitializedSystemRequestSystemComponentsInner() *AddUninitializedSystemRequestSystemComponentsInner`

NewAddUninitializedSystemRequestSystemComponentsInner instantiates a new AddUninitializedSystemRequestSystemComponentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddUninitializedSystemRequestSystemComponentsInnerWithDefaults

`func NewAddUninitializedSystemRequestSystemComponentsInnerWithDefaults() *AddUninitializedSystemRequestSystemComponentsInner`

NewAddUninitializedSystemRequestSystemComponentsInnerWithDefaults instantiates a new AddUninitializedSystemRequestSystemComponentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddUninitializedSystemRequestSystemComponentsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddUninitializedSystemRequestSystemComponentsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddUninitializedSystemRequestSystemComponentsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddUninitializedSystemRequestSystemComponentsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTypeCode

`func (o *AddUninitializedSystemRequestSystemComponentsInner) GetTypeCode() string`

GetTypeCode returns the TypeCode field if non-nil, zero value otherwise.

### GetTypeCodeOk

`func (o *AddUninitializedSystemRequestSystemComponentsInner) GetTypeCodeOk() (*string, bool)`

GetTypeCodeOk returns a tuple with the TypeCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTypeCode

`func (o *AddUninitializedSystemRequestSystemComponentsInner) SetTypeCode(v string)`

SetTypeCode sets TypeCode field to given value.

### HasTypeCode

`func (o *AddUninitializedSystemRequestSystemComponentsInner) HasTypeCode() bool`

HasTypeCode returns a boolean if a field has been set.

### GetName

`func (o *AddUninitializedSystemRequestSystemComponentsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddUninitializedSystemRequestSystemComponentsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddUninitializedSystemRequestSystemComponentsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddUninitializedSystemRequestSystemComponentsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExternalId

`func (o *AddUninitializedSystemRequestSystemComponentsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *AddUninitializedSystemRequestSystemComponentsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *AddUninitializedSystemRequestSystemComponentsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *AddUninitializedSystemRequestSystemComponentsInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### GetConfig

`func (o *AddUninitializedSystemRequestSystemComponentsInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddUninitializedSystemRequestSystemComponentsInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddUninitializedSystemRequestSystemComponentsInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddUninitializedSystemRequestSystemComponentsInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


