# AddCatalogCartItem200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Item** | Pointer to [**ListCatalogCart200ResponseCartItemsInner**](ListCatalogCart200ResponseCartItemsInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddCatalogCartItem200Response

`func NewAddCatalogCartItem200Response() *AddCatalogCartItem200Response`

NewAddCatalogCartItem200Response instantiates a new AddCatalogCartItem200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCatalogCartItem200ResponseWithDefaults

`func NewAddCatalogCartItem200ResponseWithDefaults() *AddCatalogCartItem200Response`

NewAddCatalogCartItem200ResponseWithDefaults instantiates a new AddCatalogCartItem200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItem

`func (o *AddCatalogCartItem200Response) GetItem() ListCatalogCart200ResponseCartItemsInner`

GetItem returns the Item field if non-nil, zero value otherwise.

### GetItemOk

`func (o *AddCatalogCartItem200Response) GetItemOk() (*ListCatalogCart200ResponseCartItemsInner, bool)`

GetItemOk returns a tuple with the Item field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItem

`func (o *AddCatalogCartItem200Response) SetItem(v ListCatalogCart200ResponseCartItemsInner)`

SetItem sets Item field to given value.

### HasItem

`func (o *AddCatalogCartItem200Response) HasItem() bool`

HasItem returns a boolean if a field has been set.

### GetSuccess

`func (o *AddCatalogCartItem200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddCatalogCartItem200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddCatalogCartItem200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddCatalogCartItem200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


