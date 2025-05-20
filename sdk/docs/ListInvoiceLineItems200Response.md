# ListInvoiceLineItems200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineItems** | Pointer to [**[]ListInvoiceLineItems200ResponseAllOfLineItemsInner**](ListInvoiceLineItems200ResponseAllOfLineItemsInner.md) |  | [optional] 
**MasterAccount** | Pointer to **bool** |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListInvoiceLineItems200Response

`func NewListInvoiceLineItems200Response() *ListInvoiceLineItems200Response`

NewListInvoiceLineItems200Response instantiates a new ListInvoiceLineItems200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInvoiceLineItems200ResponseWithDefaults

`func NewListInvoiceLineItems200ResponseWithDefaults() *ListInvoiceLineItems200Response`

NewListInvoiceLineItems200ResponseWithDefaults instantiates a new ListInvoiceLineItems200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLineItems

`func (o *ListInvoiceLineItems200Response) GetLineItems() []ListInvoiceLineItems200ResponseAllOfLineItemsInner`

GetLineItems returns the LineItems field if non-nil, zero value otherwise.

### GetLineItemsOk

`func (o *ListInvoiceLineItems200Response) GetLineItemsOk() (*[]ListInvoiceLineItems200ResponseAllOfLineItemsInner, bool)`

GetLineItemsOk returns a tuple with the LineItems field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItems

`func (o *ListInvoiceLineItems200Response) SetLineItems(v []ListInvoiceLineItems200ResponseAllOfLineItemsInner)`

SetLineItems sets LineItems field to given value.

### HasLineItems

`func (o *ListInvoiceLineItems200Response) HasLineItems() bool`

HasLineItems returns a boolean if a field has been set.

### GetMasterAccount

`func (o *ListInvoiceLineItems200Response) GetMasterAccount() bool`

GetMasterAccount returns the MasterAccount field if non-nil, zero value otherwise.

### GetMasterAccountOk

`func (o *ListInvoiceLineItems200Response) GetMasterAccountOk() (*bool, bool)`

GetMasterAccountOk returns a tuple with the MasterAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterAccount

`func (o *ListInvoiceLineItems200Response) SetMasterAccount(v bool)`

SetMasterAccount sets MasterAccount field to given value.

### HasMasterAccount

`func (o *ListInvoiceLineItems200Response) HasMasterAccount() bool`

HasMasterAccount returns a boolean if a field has been set.

### GetMeta

`func (o *ListInvoiceLineItems200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListInvoiceLineItems200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListInvoiceLineItems200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListInvoiceLineItems200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


