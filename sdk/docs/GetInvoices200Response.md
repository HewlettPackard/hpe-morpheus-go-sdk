# GetInvoices200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invoice** | Pointer to [**ListInvoices200ResponseAllOfInvoicesInner**](ListInvoices200ResponseAllOfInvoicesInner.md) |  | [optional] 
**MasterAccount** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetInvoices200Response

`func NewGetInvoices200Response() *GetInvoices200Response`

NewGetInvoices200Response instantiates a new GetInvoices200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInvoices200ResponseWithDefaults

`func NewGetInvoices200ResponseWithDefaults() *GetInvoices200Response`

NewGetInvoices200ResponseWithDefaults instantiates a new GetInvoices200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoice

`func (o *GetInvoices200Response) GetInvoice() ListInvoices200ResponseAllOfInvoicesInner`

GetInvoice returns the Invoice field if non-nil, zero value otherwise.

### GetInvoiceOk

`func (o *GetInvoices200Response) GetInvoiceOk() (*ListInvoices200ResponseAllOfInvoicesInner, bool)`

GetInvoiceOk returns a tuple with the Invoice field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoice

`func (o *GetInvoices200Response) SetInvoice(v ListInvoices200ResponseAllOfInvoicesInner)`

SetInvoice sets Invoice field to given value.

### HasInvoice

`func (o *GetInvoices200Response) HasInvoice() bool`

HasInvoice returns a boolean if a field has been set.

### GetMasterAccount

`func (o *GetInvoices200Response) GetMasterAccount() bool`

GetMasterAccount returns the MasterAccount field if non-nil, zero value otherwise.

### GetMasterAccountOk

`func (o *GetInvoices200Response) GetMasterAccountOk() (*bool, bool)`

GetMasterAccountOk returns a tuple with the MasterAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterAccount

`func (o *GetInvoices200Response) SetMasterAccount(v bool)`

SetMasterAccount sets MasterAccount field to given value.

### HasMasterAccount

`func (o *GetInvoices200Response) HasMasterAccount() bool`

HasMasterAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


