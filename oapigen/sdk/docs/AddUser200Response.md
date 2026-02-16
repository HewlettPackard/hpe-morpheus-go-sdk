# AddUser200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**User** | Pointer to [**AddUser200ResponseAllOfUser**](AddUser200ResponseAllOfUser.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddUser200Response

`func NewAddUser200Response() *AddUser200Response`

NewAddUser200Response instantiates a new AddUser200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddUser200ResponseWithDefaults

`func NewAddUser200ResponseWithDefaults() *AddUser200Response`

NewAddUser200ResponseWithDefaults instantiates a new AddUser200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUser

`func (o *AddUser200Response) GetUser() AddUser200ResponseAllOfUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *AddUser200Response) GetUserOk() (*AddUser200ResponseAllOfUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *AddUser200Response) SetUser(v AddUser200ResponseAllOfUser)`

SetUser sets User field to given value.

### HasUser

`func (o *AddUser200Response) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetSuccess

`func (o *AddUser200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddUser200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddUser200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddUser200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


