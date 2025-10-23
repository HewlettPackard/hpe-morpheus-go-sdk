# AddCatalogOrder200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Order** | Pointer to [**AddCatalogOrder200ResponseAllOfOrder**](AddCatalogOrder200ResponseAllOfOrder.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddCatalogOrder200Response

`func NewAddCatalogOrder200Response() *AddCatalogOrder200Response`

NewAddCatalogOrder200Response instantiates a new AddCatalogOrder200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCatalogOrder200ResponseWithDefaults

`func NewAddCatalogOrder200ResponseWithDefaults() *AddCatalogOrder200Response`

NewAddCatalogOrder200ResponseWithDefaults instantiates a new AddCatalogOrder200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOrder

`func (o *AddCatalogOrder200Response) GetOrder() AddCatalogOrder200ResponseAllOfOrder`

GetOrder returns the Order field if non-nil, zero value otherwise.

### GetOrderOk

`func (o *AddCatalogOrder200Response) GetOrderOk() (*AddCatalogOrder200ResponseAllOfOrder, bool)`

GetOrderOk returns a tuple with the Order field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrder

`func (o *AddCatalogOrder200Response) SetOrder(v AddCatalogOrder200ResponseAllOfOrder)`

SetOrder sets Order field to given value.

### HasOrder

`func (o *AddCatalogOrder200Response) HasOrder() bool`

HasOrder returns a boolean if a field has been set.

### GetSuccess

`func (o *AddCatalogOrder200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddCatalogOrder200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddCatalogOrder200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddCatalogOrder200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


