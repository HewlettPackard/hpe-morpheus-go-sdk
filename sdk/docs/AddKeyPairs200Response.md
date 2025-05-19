# AddKeyPairs200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Account** | Pointer to [**AddKeyPairs200ResponseAllOfAccount**](AddKeyPairs200ResponseAllOfAccount.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddKeyPairs200Response

`func NewAddKeyPairs200Response() *AddKeyPairs200Response`

NewAddKeyPairs200Response instantiates a new AddKeyPairs200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddKeyPairs200ResponseWithDefaults

`func NewAddKeyPairs200ResponseWithDefaults() *AddKeyPairs200Response`

NewAddKeyPairs200ResponseWithDefaults instantiates a new AddKeyPairs200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccount

`func (o *AddKeyPairs200Response) GetAccount() AddKeyPairs200ResponseAllOfAccount`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *AddKeyPairs200Response) GetAccountOk() (*AddKeyPairs200ResponseAllOfAccount, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *AddKeyPairs200Response) SetAccount(v AddKeyPairs200ResponseAllOfAccount)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *AddKeyPairs200Response) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetSuccess

`func (o *AddKeyPairs200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddKeyPairs200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddKeyPairs200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddKeyPairs200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


