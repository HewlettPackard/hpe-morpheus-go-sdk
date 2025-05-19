# ListTenants200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accounts** | Pointer to [**[]ListTenants200ResponseAllOfAccountsInner**](ListTenants200ResponseAllOfAccountsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListTenants200Response

`func NewListTenants200Response() *ListTenants200Response`

NewListTenants200Response instantiates a new ListTenants200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListTenants200ResponseWithDefaults

`func NewListTenants200ResponseWithDefaults() *ListTenants200Response`

NewListTenants200ResponseWithDefaults instantiates a new ListTenants200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccounts

`func (o *ListTenants200Response) GetAccounts() []ListTenants200ResponseAllOfAccountsInner`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *ListTenants200Response) GetAccountsOk() (*[]ListTenants200ResponseAllOfAccountsInner, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *ListTenants200Response) SetAccounts(v []ListTenants200ResponseAllOfAccountsInner)`

SetAccounts sets Accounts field to given value.

### HasAccounts

`func (o *ListTenants200Response) HasAccounts() bool`

HasAccounts returns a boolean if a field has been set.

### GetMeta

`func (o *ListTenants200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListTenants200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListTenants200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListTenants200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


