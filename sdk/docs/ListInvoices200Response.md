# ListInvoices200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Invoices** | Pointer to [**[]ListInvoices200ResponseAllOfInvoicesInner**](ListInvoices200ResponseAllOfInvoicesInner.md) |  | [optional] 
**MasterAccount** | Pointer to **bool** |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListInvoices200Response

`func NewListInvoices200Response() *ListInvoices200Response`

NewListInvoices200Response instantiates a new ListInvoices200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListInvoices200ResponseWithDefaults

`func NewListInvoices200ResponseWithDefaults() *ListInvoices200Response`

NewListInvoices200ResponseWithDefaults instantiates a new ListInvoices200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInvoices

`func (o *ListInvoices200Response) GetInvoices() []ListInvoices200ResponseAllOfInvoicesInner`

GetInvoices returns the Invoices field if non-nil, zero value otherwise.

### GetInvoicesOk

`func (o *ListInvoices200Response) GetInvoicesOk() (*[]ListInvoices200ResponseAllOfInvoicesInner, bool)`

GetInvoicesOk returns a tuple with the Invoices field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInvoices

`func (o *ListInvoices200Response) SetInvoices(v []ListInvoices200ResponseAllOfInvoicesInner)`

SetInvoices sets Invoices field to given value.

### HasInvoices

`func (o *ListInvoices200Response) HasInvoices() bool`

HasInvoices returns a boolean if a field has been set.

### GetMasterAccount

`func (o *ListInvoices200Response) GetMasterAccount() bool`

GetMasterAccount returns the MasterAccount field if non-nil, zero value otherwise.

### GetMasterAccountOk

`func (o *ListInvoices200Response) GetMasterAccountOk() (*bool, bool)`

GetMasterAccountOk returns a tuple with the MasterAccount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterAccount

`func (o *ListInvoices200Response) SetMasterAccount(v bool)`

SetMasterAccount sets MasterAccount field to given value.

### HasMasterAccount

`func (o *ListInvoices200Response) HasMasterAccount() bool`

HasMasterAccount returns a boolean if a field has been set.

### GetMeta

`func (o *ListInvoices200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListInvoices200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListInvoices200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListInvoices200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


