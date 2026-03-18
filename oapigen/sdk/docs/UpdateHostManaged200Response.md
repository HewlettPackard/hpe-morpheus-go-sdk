# UpdateHostManaged200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicKey** | Pointer to **string** | Public key to be put into &#x60;authorized_keys&#x60; on target VM | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateHostManaged200Response

`func NewUpdateHostManaged200Response() *UpdateHostManaged200Response`

NewUpdateHostManaged200Response instantiates a new UpdateHostManaged200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateHostManaged200ResponseWithDefaults

`func NewUpdateHostManaged200ResponseWithDefaults() *UpdateHostManaged200Response`

NewUpdateHostManaged200ResponseWithDefaults instantiates a new UpdateHostManaged200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicKey

`func (o *UpdateHostManaged200Response) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *UpdateHostManaged200Response) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *UpdateHostManaged200Response) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *UpdateHostManaged200Response) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateHostManaged200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateHostManaged200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateHostManaged200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateHostManaged200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


