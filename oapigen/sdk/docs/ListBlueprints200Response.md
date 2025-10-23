# ListBlueprints200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Blueprints** | Pointer to [**[]ListBlueprints200ResponseAllOfBlueprintsInner**](ListBlueprints200ResponseAllOfBlueprintsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListBlueprints200Response

`func NewListBlueprints200Response() *ListBlueprints200Response`

NewListBlueprints200Response instantiates a new ListBlueprints200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBlueprints200ResponseWithDefaults

`func NewListBlueprints200ResponseWithDefaults() *ListBlueprints200Response`

NewListBlueprints200ResponseWithDefaults instantiates a new ListBlueprints200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBlueprints

`func (o *ListBlueprints200Response) GetBlueprints() []ListBlueprints200ResponseAllOfBlueprintsInner`

GetBlueprints returns the Blueprints field if non-nil, zero value otherwise.

### GetBlueprintsOk

`func (o *ListBlueprints200Response) GetBlueprintsOk() (*[]ListBlueprints200ResponseAllOfBlueprintsInner, bool)`

GetBlueprintsOk returns a tuple with the Blueprints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlueprints

`func (o *ListBlueprints200Response) SetBlueprints(v []ListBlueprints200ResponseAllOfBlueprintsInner)`

SetBlueprints sets Blueprints field to given value.

### HasBlueprints

`func (o *ListBlueprints200Response) HasBlueprints() bool`

HasBlueprints returns a boolean if a field has been set.

### GetMeta

`func (o *ListBlueprints200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListBlueprints200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListBlueprints200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListBlueprints200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


