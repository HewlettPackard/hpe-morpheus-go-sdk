# KeyPair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**PublicKey** | Pointer to **string** |  | [optional] 
**HasPrivateKey** | Pointer to **bool** |  | [optional] 
**PrivateKeyHash** | Pointer to **string** |  | [optional] 
**PrivateKey** | Pointer to **string** | Only present in response to generate | [optional] 
**Fingerprint** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewKeyPair

`func NewKeyPair() *KeyPair`

NewKeyPair instantiates a new KeyPair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewKeyPairWithDefaults

`func NewKeyPairWithDefaults() *KeyPair`

NewKeyPairWithDefaults instantiates a new KeyPair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *KeyPair) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *KeyPair) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *KeyPair) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *KeyPair) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *KeyPair) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *KeyPair) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *KeyPair) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *KeyPair) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccountId

`func (o *KeyPair) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *KeyPair) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *KeyPair) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *KeyPair) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPublicKey

`func (o *KeyPair) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *KeyPair) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *KeyPair) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *KeyPair) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetHasPrivateKey

`func (o *KeyPair) GetHasPrivateKey() bool`

GetHasPrivateKey returns the HasPrivateKey field if non-nil, zero value otherwise.

### GetHasPrivateKeyOk

`func (o *KeyPair) GetHasPrivateKeyOk() (*bool, bool)`

GetHasPrivateKeyOk returns a tuple with the HasPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrivateKey

`func (o *KeyPair) SetHasPrivateKey(v bool)`

SetHasPrivateKey sets HasPrivateKey field to given value.

### HasHasPrivateKey

`func (o *KeyPair) HasHasPrivateKey() bool`

HasHasPrivateKey returns a boolean if a field has been set.

### GetPrivateKeyHash

`func (o *KeyPair) GetPrivateKeyHash() string`

GetPrivateKeyHash returns the PrivateKeyHash field if non-nil, zero value otherwise.

### GetPrivateKeyHashOk

`func (o *KeyPair) GetPrivateKeyHashOk() (*string, bool)`

GetPrivateKeyHashOk returns a tuple with the PrivateKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyHash

`func (o *KeyPair) SetPrivateKeyHash(v string)`

SetPrivateKeyHash sets PrivateKeyHash field to given value.

### HasPrivateKeyHash

`func (o *KeyPair) HasPrivateKeyHash() bool`

HasPrivateKeyHash returns a boolean if a field has been set.

### GetPrivateKey

`func (o *KeyPair) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *KeyPair) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *KeyPair) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *KeyPair) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetFingerprint

`func (o *KeyPair) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *KeyPair) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *KeyPair) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *KeyPair) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetDateCreated

`func (o *KeyPair) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *KeyPair) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *KeyPair) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *KeyPair) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *KeyPair) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *KeyPair) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *KeyPair) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *KeyPair) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


