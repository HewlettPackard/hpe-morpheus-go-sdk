# ListIntegrationInventory200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inventory** | Pointer to [**[]ListIntegrationInventory200ResponseAllOfInventoryInner**](ListIntegrationInventory200ResponseAllOfInventoryInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListIntegrationInventory200Response

`func NewListIntegrationInventory200Response() *ListIntegrationInventory200Response`

NewListIntegrationInventory200Response instantiates a new ListIntegrationInventory200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListIntegrationInventory200ResponseWithDefaults

`func NewListIntegrationInventory200ResponseWithDefaults() *ListIntegrationInventory200Response`

NewListIntegrationInventory200ResponseWithDefaults instantiates a new ListIntegrationInventory200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInventory

`func (o *ListIntegrationInventory200Response) GetInventory() []ListIntegrationInventory200ResponseAllOfInventoryInner`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *ListIntegrationInventory200Response) GetInventoryOk() (*[]ListIntegrationInventory200ResponseAllOfInventoryInner, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *ListIntegrationInventory200Response) SetInventory(v []ListIntegrationInventory200ResponseAllOfInventoryInner)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *ListIntegrationInventory200Response) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetMeta

`func (o *ListIntegrationInventory200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListIntegrationInventory200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListIntegrationInventory200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListIntegrationInventory200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


