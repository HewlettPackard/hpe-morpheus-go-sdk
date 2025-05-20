# ListUserSettings200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**ListUserSettings200ResponseAllOfUser**](ListUserSettings200ResponseAllOfUser.md) |  | [optional] 
**AccessTokens** | Pointer to [**[]ListUserSettings200ResponseAllOfAccessTokensInner**](ListUserSettings200ResponseAllOfAccessTokensInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewListUserSettings200Response

`func NewListUserSettings200Response() *ListUserSettings200Response`

NewListUserSettings200Response instantiates a new ListUserSettings200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListUserSettings200ResponseWithDefaults

`func NewListUserSettings200ResponseWithDefaults() *ListUserSettings200Response`

NewListUserSettings200ResponseWithDefaults instantiates a new ListUserSettings200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *ListUserSettings200Response) GetUser() ListUserSettings200ResponseAllOfUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ListUserSettings200Response) GetUserOk() (*ListUserSettings200ResponseAllOfUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ListUserSettings200Response) SetUser(v ListUserSettings200ResponseAllOfUser)`

SetUser sets User field to given value.

### HasUser

`func (o *ListUserSettings200Response) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetAccessTokens

`func (o *ListUserSettings200Response) GetAccessTokens() []ListUserSettings200ResponseAllOfAccessTokensInner`

GetAccessTokens returns the AccessTokens field if non-nil, zero value otherwise.

### GetAccessTokensOk

`func (o *ListUserSettings200Response) GetAccessTokensOk() (*[]ListUserSettings200ResponseAllOfAccessTokensInner, bool)`

GetAccessTokensOk returns a tuple with the AccessTokens field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessTokens

`func (o *ListUserSettings200Response) SetAccessTokens(v []ListUserSettings200ResponseAllOfAccessTokensInner)`

SetAccessTokens sets AccessTokens field to given value.

### HasAccessTokens

`func (o *ListUserSettings200Response) HasAccessTokens() bool`

HasAccessTokens returns a boolean if a field has been set.

### GetSuccess

`func (o *ListUserSettings200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ListUserSettings200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ListUserSettings200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ListUserSettings200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


