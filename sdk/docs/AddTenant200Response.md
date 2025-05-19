# AddTenant200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**ListTenants200ResponseAllOfAccountsInner**](ListTenants200ResponseAllOfAccountsInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddTenant200Response

`func NewAddTenant200Response() *AddTenant200Response`

NewAddTenant200Response instantiates a new AddTenant200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddTenant200ResponseWithDefaults

`func NewAddTenant200ResponseWithDefaults() *AddTenant200Response`

NewAddTenant200ResponseWithDefaults instantiates a new AddTenant200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *AddTenant200Response) GetAccount() ListTenants200ResponseAllOfAccountsInner`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AddTenant200Response) GetAccountOk() (*ListTenants200ResponseAllOfAccountsInner, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AddTenant200Response) SetAccount(v ListTenants200ResponseAllOfAccountsInner)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AddTenant200Response) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSuccess

`func (o *AddTenant200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddTenant200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddTenant200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddTenant200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


