# UpdateIntegrationInventory200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inventory** | Pointer to [**ListIntegrationInventory200ResponseAllOfInventoryInner**](ListIntegrationInventory200ResponseAllOfInventoryInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateIntegrationInventory200Response

`func NewUpdateIntegrationInventory200Response() *UpdateIntegrationInventory200Response`

NewUpdateIntegrationInventory200Response instantiates a new UpdateIntegrationInventory200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateIntegrationInventory200ResponseWithDefaults

`func NewUpdateIntegrationInventory200ResponseWithDefaults() *UpdateIntegrationInventory200Response`

NewUpdateIntegrationInventory200ResponseWithDefaults instantiates a new UpdateIntegrationInventory200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInventory

`func (o *UpdateIntegrationInventory200Response) GetInventory() ListIntegrationInventory200ResponseAllOfInventoryInner`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *UpdateIntegrationInventory200Response) GetInventoryOk() (*ListIntegrationInventory200ResponseAllOfInventoryInner, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *UpdateIntegrationInventory200Response) SetInventory(v ListIntegrationInventory200ResponseAllOfInventoryInner)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *UpdateIntegrationInventory200Response) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateIntegrationInventory200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateIntegrationInventory200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateIntegrationInventory200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateIntegrationInventory200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


