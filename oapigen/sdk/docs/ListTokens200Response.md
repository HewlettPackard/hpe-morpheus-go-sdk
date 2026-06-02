# ListTokens200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tokens** | Pointer to [**[]ListTokens200ResponseAllOfTokensInner**](ListTokens200ResponseAllOfTokensInner.md) |  | [optional] 
**Meta** | Pointer to [**ListApprovals200ResponseAllOfMeta**](ListApprovals200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListTokens200Response

`func NewListTokens200Response() *ListTokens200Response`

NewListTokens200Response instantiates a new ListTokens200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetTokens

`func (o *ListTokens200Response) GetTokens() []ListTokens200ResponseAllOfTokensInner`

GetTokens returns the Tokens field if non-nil, zero value otherwise.

### GetTokensOk

`func (o *ListTokens200Response) GetTokensOk() (*[]ListTokens200ResponseAllOfTokensInner, bool)`

GetTokensOk returns a tuple with the Tokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokens

`func (o *ListTokens200Response) SetTokens(v []ListTokens200ResponseAllOfTokensInner)`

SetTokens sets Tokens field to given value.

### HasTokens

`func (o *ListTokens200Response) HasTokens() bool`

HasTokens returns a boolean if a field has been set.

### GetMeta

`func (o *ListTokens200Response) GetMeta() ListApprovals200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListTokens200Response) GetMetaOk() (*ListApprovals200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListTokens200Response) SetMeta(v ListApprovals200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListTokens200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


