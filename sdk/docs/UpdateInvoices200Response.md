# UpdateInvoices200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**ListInvoices200ResponseAllOfInvoicesInner**](ListInvoices200ResponseAllOfInvoicesInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateInvoices200Response

`func NewUpdateInvoices200Response() *UpdateInvoices200Response`

NewUpdateInvoices200Response instantiates a new UpdateInvoices200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInvoices200ResponseWithDefaults

`func NewUpdateInvoices200ResponseWithDefaults() *UpdateInvoices200Response`

NewUpdateInvoices200ResponseWithDefaults instantiates a new UpdateInvoices200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *UpdateInvoices200Response) GetUser() ListInvoices200ResponseAllOfInvoicesInner`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *UpdateInvoices200Response) GetUserOk() (*ListInvoices200ResponseAllOfInvoicesInner, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *UpdateInvoices200Response) SetUser(v ListInvoices200ResponseAllOfInvoicesInner)`

SetUser sets User field to given value.

### HasUser

`func (o *UpdateInvoices200Response) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateInvoices200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateInvoices200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateInvoices200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateInvoices200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


