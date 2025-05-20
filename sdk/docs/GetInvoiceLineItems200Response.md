# GetInvoiceLineItems200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LineItem** | Pointer to [**ListInvoiceLineItems200ResponseAllOfLineItemsInner**](ListInvoiceLineItems200ResponseAllOfLineItemsInner.md) |  | [optional] 
**MasterAccount** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetInvoiceLineItems200Response

`func NewGetInvoiceLineItems200Response() *GetInvoiceLineItems200Response`

NewGetInvoiceLineItems200Response instantiates a new GetInvoiceLineItems200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInvoiceLineItems200ResponseWithDefaults

`func NewGetInvoiceLineItems200ResponseWithDefaults() *GetInvoiceLineItems200Response`

NewGetInvoiceLineItems200ResponseWithDefaults instantiates a new GetInvoiceLineItems200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLineItem

`func (o *GetInvoiceLineItems200Response) GetLineItem() ListInvoiceLineItems200ResponseAllOfLineItemsInner`

GetLineItem returns the LineItem field if non-nil, zero value otherwise.

### GetLineItemOk

`func (o *GetInvoiceLineItems200Response) GetLineItemOk() (*ListInvoiceLineItems200ResponseAllOfLineItemsInner, bool)`

GetLineItemOk returns a tuple with the LineItem field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineItem

`func (o *GetInvoiceLineItems200Response) SetLineItem(v ListInvoiceLineItems200ResponseAllOfLineItemsInner)`

SetLineItem sets LineItem field to given value.

### HasLineItem

`func (o *GetInvoiceLineItems200Response) HasLineItem() bool`

HasLineItem returns a boolean if a field has been set.

### GetMasterAccount

`func (o *GetInvoiceLineItems200Response) GetMasterAccount() bool`

GetMasterAccount returns the MasterAccount field if non-nil, zero value otherwise.

### GetMasterAccountOk

`func (o *GetInvoiceLineItems200Response) GetMasterAccountOk() (*bool, bool)`

GetMasterAccountOk returns a tuple with the MasterAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterAccount

`func (o *GetInvoiceLineItems200Response) SetMasterAccount(v bool)`

SetMasterAccount sets MasterAccount field to given value.

### HasMasterAccount

`func (o *GetInvoiceLineItems200Response) HasMasterAccount() bool`

HasMasterAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


