# AddTokenRequestToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | Client ID | 
**Name** | Pointer to **string** | Name is optional and can be used to help identify the token. It is not used for authentication. | [optional] 

## Methods

### NewAddTokenRequestToken

`func NewAddTokenRequestToken(clientId string, ) *AddTokenRequestToken`

NewAddTokenRequestToken instantiates a new AddTokenRequestToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetClientId

`func (o *AddTokenRequestToken) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *AddTokenRequestToken) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *AddTokenRequestToken) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetName

`func (o *AddTokenRequestToken) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddTokenRequestToken) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddTokenRequestToken) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddTokenRequestToken) HasName() bool`

HasName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


