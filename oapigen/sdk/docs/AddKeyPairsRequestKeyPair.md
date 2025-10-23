# AddKeyPairsRequestKeyPair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**PublicKey** | **string** |  | 
**PrivateKey** | Pointer to **string** |  | [optional] 
**Passphrase** | Pointer to **string** |  | [optional] 

## Methods

### NewAddKeyPairsRequestKeyPair

`func NewAddKeyPairsRequestKeyPair(name string, publicKey string, ) *AddKeyPairsRequestKeyPair`

NewAddKeyPairsRequestKeyPair instantiates a new AddKeyPairsRequestKeyPair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddKeyPairsRequestKeyPairWithDefaults

`func NewAddKeyPairsRequestKeyPairWithDefaults() *AddKeyPairsRequestKeyPair`

NewAddKeyPairsRequestKeyPairWithDefaults instantiates a new AddKeyPairsRequestKeyPair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddKeyPairsRequestKeyPair) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddKeyPairsRequestKeyPair) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddKeyPairsRequestKeyPair) SetName(v string)`

SetName sets Name field to given value.


### GetPublicKey

`func (o *AddKeyPairsRequestKeyPair) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *AddKeyPairsRequestKeyPair) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *AddKeyPairsRequestKeyPair) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.


### GetPrivateKey

`func (o *AddKeyPairsRequestKeyPair) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *AddKeyPairsRequestKeyPair) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *AddKeyPairsRequestKeyPair) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *AddKeyPairsRequestKeyPair) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetPassphrase

`func (o *AddKeyPairsRequestKeyPair) GetPassphrase() string`

GetPassphrase returns the Passphrase field if non-nil, zero value otherwise.

### GetPassphraseOk

`func (o *AddKeyPairsRequestKeyPair) GetPassphraseOk() (*string, bool)`

GetPassphraseOk returns a tuple with the Passphrase field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassphrase

`func (o *AddKeyPairsRequestKeyPair) SetPassphrase(v string)`

SetPassphrase sets Passphrase field to given value.

### HasPassphrase

`func (o *AddKeyPairsRequestKeyPair) HasPassphrase() bool`

HasPassphrase returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


