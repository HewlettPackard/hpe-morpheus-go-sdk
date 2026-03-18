# UpdateHostInstallAgent200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicKey** | Pointer to **string** | Public key to be put into &#x60;authorized_keys&#x60; on target VM | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateHostInstallAgent200Response

`func NewUpdateHostInstallAgent200Response() *UpdateHostInstallAgent200Response`

NewUpdateHostInstallAgent200Response instantiates a new UpdateHostInstallAgent200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateHostInstallAgent200ResponseWithDefaults

`func NewUpdateHostInstallAgent200ResponseWithDefaults() *UpdateHostInstallAgent200Response`

NewUpdateHostInstallAgent200ResponseWithDefaults instantiates a new UpdateHostInstallAgent200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicKey

`func (o *UpdateHostInstallAgent200Response) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *UpdateHostInstallAgent200Response) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *UpdateHostInstallAgent200Response) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *UpdateHostInstallAgent200Response) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateHostInstallAgent200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateHostInstallAgent200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateHostInstallAgent200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateHostInstallAgent200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


