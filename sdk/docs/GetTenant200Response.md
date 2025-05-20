# GetTenant200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**ListTenants200ResponseAllOfAccountsInner**](ListTenants200ResponseAllOfAccountsInner.md) |  | [optional] 

## Methods

### NewGetTenant200Response

`func NewGetTenant200Response() *GetTenant200Response`

NewGetTenant200Response instantiates a new GetTenant200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetTenant200ResponseWithDefaults

`func NewGetTenant200ResponseWithDefaults() *GetTenant200Response`

NewGetTenant200ResponseWithDefaults instantiates a new GetTenant200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *GetTenant200Response) GetAccount() ListTenants200ResponseAllOfAccountsInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *GetTenant200Response) GetAccountOk() (*ListTenants200ResponseAllOfAccountsInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *GetTenant200Response) SetAccount(v ListTenants200ResponseAllOfAccountsInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *GetTenant200Response) HasAccount() bool`

HasAccount returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


