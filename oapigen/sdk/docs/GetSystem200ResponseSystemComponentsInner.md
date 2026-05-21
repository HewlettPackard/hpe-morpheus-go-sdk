# GetSystem200ResponseSystemComponentsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**Type** | Pointer to [**GetSystem200ResponseSystemComponentsInnerType**](GetSystem200ResponseSystemComponentsInnerType.md) |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**ExternalId** | Pointer to **NullableString** |  | [optional] 
**Config** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewGetSystem200ResponseSystemComponentsInner

`func NewGetSystem200ResponseSystemComponentsInner() *GetSystem200ResponseSystemComponentsInner`

NewGetSystem200ResponseSystemComponentsInner instantiates a new GetSystem200ResponseSystemComponentsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSystem200ResponseSystemComponentsInnerWithDefaults

`func NewGetSystem200ResponseSystemComponentsInnerWithDefaults() *GetSystem200ResponseSystemComponentsInner`

NewGetSystem200ResponseSystemComponentsInnerWithDefaults instantiates a new GetSystem200ResponseSystemComponentsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetSystem200ResponseSystemComponentsInner) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetSystem200ResponseSystemComponentsInner) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetSystem200ResponseSystemComponentsInner) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetSystem200ResponseSystemComponentsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetSystem200ResponseSystemComponentsInner) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetSystem200ResponseSystemComponentsInner) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetSystem200ResponseSystemComponentsInner) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetSystem200ResponseSystemComponentsInner) HasName() bool`

HasName returns a boolean if a field has been set.

### GetType

`func (o *GetSystem200ResponseSystemComponentsInner) GetType() GetSystem200ResponseSystemComponentsInnerType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetSystem200ResponseSystemComponentsInner) GetTypeOk() (*GetSystem200ResponseSystemComponentsInnerType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetSystem200ResponseSystemComponentsInner) SetType(v GetSystem200ResponseSystemComponentsInnerType)`

SetType sets Type field to given value.

### HasType

`func (o *GetSystem200ResponseSystemComponentsInner) HasType() bool`

HasType returns a boolean if a field has been set.

### GetStatus

`func (o *GetSystem200ResponseSystemComponentsInner) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetSystem200ResponseSystemComponentsInner) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetSystem200ResponseSystemComponentsInner) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetSystem200ResponseSystemComponentsInner) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetExternalId

`func (o *GetSystem200ResponseSystemComponentsInner) GetExternalId() string`

GetExternalId returns the ExternalId field if non-nil, zero value otherwise.

### GetExternalIdOk

`func (o *GetSystem200ResponseSystemComponentsInner) GetExternalIdOk() (*string, bool)`

GetExternalIdOk returns a tuple with the ExternalId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalId

`func (o *GetSystem200ResponseSystemComponentsInner) SetExternalId(v string)`

SetExternalId sets ExternalId field to given value.

### HasExternalId

`func (o *GetSystem200ResponseSystemComponentsInner) HasExternalId() bool`

HasExternalId returns a boolean if a field has been set.

### SetExternalIdNil

`func (o *GetSystem200ResponseSystemComponentsInner) SetExternalIdNil(b bool)`

 SetExternalIdNil sets the value for ExternalId to be an explicit nil

### UnsetExternalId
`func (o *GetSystem200ResponseSystemComponentsInner) UnsetExternalId()`

UnsetExternalId ensures that no value is present for ExternalId, not even an explicit nil
### GetConfig

`func (o *GetSystem200ResponseSystemComponentsInner) GetConfig() map[string]interface{}`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *GetSystem200ResponseSystemComponentsInner) GetConfigOk() (*map[string]interface{}, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *GetSystem200ResponseSystemComponentsInner) SetConfig(v map[string]interface{})`

SetConfig sets Config field to given value.

### HasConfig

`func (o *GetSystem200ResponseSystemComponentsInner) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


