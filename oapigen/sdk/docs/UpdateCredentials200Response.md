# UpdateCredentials200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credential** | Pointer to [**UpdateCredentials200ResponseAllOfCredential**](UpdateCredentials200ResponseAllOfCredential.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateCredentials200Response

`func NewUpdateCredentials200Response() *UpdateCredentials200Response`

NewUpdateCredentials200Response instantiates a new UpdateCredentials200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCredentials200ResponseWithDefaults

`func NewUpdateCredentials200ResponseWithDefaults() *UpdateCredentials200Response`

NewUpdateCredentials200ResponseWithDefaults instantiates a new UpdateCredentials200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredential

`func (o *UpdateCredentials200Response) GetCredential() UpdateCredentials200ResponseAllOfCredential`

GetCredential returns the Credential field if non-nil, zero value otherwise.

### GetCredentialOk

`func (o *UpdateCredentials200Response) GetCredentialOk() (*UpdateCredentials200ResponseAllOfCredential, bool)`

GetCredentialOk returns a tuple with the Credential field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredential

`func (o *UpdateCredentials200Response) SetCredential(v UpdateCredentials200ResponseAllOfCredential)`

SetCredential sets Credential field to given value.

### HasCredential

`func (o *UpdateCredentials200Response) HasCredential() bool`

HasCredential returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateCredentials200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateCredentials200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateCredentials200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateCredentials200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


