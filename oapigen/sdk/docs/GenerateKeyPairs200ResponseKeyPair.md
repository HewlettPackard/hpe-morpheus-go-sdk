# GenerateKeyPairs200ResponseKeyPair

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**AccountId** | Pointer to **int64** |  | [optional] 
**PublicKey** | Pointer to **NullableString** |  | [optional] 
**HasPrivateKey** | Pointer to **bool** |  | [optional] 
**PrivateKeyHash** | Pointer to **NullableString** |  | [optional] 
**PrivateKey** | Pointer to **NullableString** | Only present in response to generate | [optional] 
**Fingerprint** | Pointer to **NullableString** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGenerateKeyPairs200ResponseKeyPair

`func NewGenerateKeyPairs200ResponseKeyPair() *GenerateKeyPairs200ResponseKeyPair`

NewGenerateKeyPairs200ResponseKeyPair instantiates a new GenerateKeyPairs200ResponseKeyPair object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenerateKeyPairs200ResponseKeyPairWithDefaults

`func NewGenerateKeyPairs200ResponseKeyPairWithDefaults() *GenerateKeyPairs200ResponseKeyPair`

NewGenerateKeyPairs200ResponseKeyPairWithDefaults instantiates a new GenerateKeyPairs200ResponseKeyPair object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GenerateKeyPairs200ResponseKeyPair) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GenerateKeyPairs200ResponseKeyPair) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GenerateKeyPairs200ResponseKeyPair) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GenerateKeyPairs200ResponseKeyPair) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GenerateKeyPairs200ResponseKeyPair) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GenerateKeyPairs200ResponseKeyPair) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GenerateKeyPairs200ResponseKeyPair) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GenerateKeyPairs200ResponseKeyPair) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccountId

`func (o *GenerateKeyPairs200ResponseKeyPair) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GenerateKeyPairs200ResponseKeyPair) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GenerateKeyPairs200ResponseKeyPair) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GenerateKeyPairs200ResponseKeyPair) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPublicKey

`func (o *GenerateKeyPairs200ResponseKeyPair) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *GenerateKeyPairs200ResponseKeyPair) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *GenerateKeyPairs200ResponseKeyPair) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *GenerateKeyPairs200ResponseKeyPair) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKeyNil

`func (o *GenerateKeyPairs200ResponseKeyPair) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *GenerateKeyPairs200ResponseKeyPair) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetHasPrivateKey

`func (o *GenerateKeyPairs200ResponseKeyPair) GetHasPrivateKey() bool`

GetHasPrivateKey returns the HasPrivateKey field if non-nil, zero value otherwise.

### GetHasPrivateKeyOk

`func (o *GenerateKeyPairs200ResponseKeyPair) GetHasPrivateKeyOk() (*bool, bool)`

GetHasPrivateKeyOk returns a tuple with the HasPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrivateKey

`func (o *GenerateKeyPairs200ResponseKeyPair) SetHasPrivateKey(v bool)`

SetHasPrivateKey sets HasPrivateKey field to given value.

### HasHasPrivateKey

`func (o *GenerateKeyPairs200ResponseKeyPair) HasHasPrivateKey() bool`

HasHasPrivateKey returns a boolean if a field has been set.

### GetPrivateKeyHash

`func (o *GenerateKeyPairs200ResponseKeyPair) GetPrivateKeyHash() string`

GetPrivateKeyHash returns the PrivateKeyHash field if non-nil, zero value otherwise.

### GetPrivateKeyHashOk

`func (o *GenerateKeyPairs200ResponseKeyPair) GetPrivateKeyHashOk() (*string, bool)`

GetPrivateKeyHashOk returns a tuple with the PrivateKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyHash

`func (o *GenerateKeyPairs200ResponseKeyPair) SetPrivateKeyHash(v string)`

SetPrivateKeyHash sets PrivateKeyHash field to given value.

### HasPrivateKeyHash

`func (o *GenerateKeyPairs200ResponseKeyPair) HasPrivateKeyHash() bool`

HasPrivateKeyHash returns a boolean if a field has been set.

### SetPrivateKeyHashNil

`func (o *GenerateKeyPairs200ResponseKeyPair) SetPrivateKeyHashNil(b bool)`

 SetPrivateKeyHashNil sets the value for PrivateKeyHash to be an explicit nil

### UnsetPrivateKeyHash
`func (o *GenerateKeyPairs200ResponseKeyPair) UnsetPrivateKeyHash()`

UnsetPrivateKeyHash ensures that no value is present for PrivateKeyHash, not even an explicit nil
### GetPrivateKey

`func (o *GenerateKeyPairs200ResponseKeyPair) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *GenerateKeyPairs200ResponseKeyPair) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *GenerateKeyPairs200ResponseKeyPair) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *GenerateKeyPairs200ResponseKeyPair) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### SetPrivateKeyNil

`func (o *GenerateKeyPairs200ResponseKeyPair) SetPrivateKeyNil(b bool)`

 SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil

### UnsetPrivateKey
`func (o *GenerateKeyPairs200ResponseKeyPair) UnsetPrivateKey()`

UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil
### GetFingerprint

`func (o *GenerateKeyPairs200ResponseKeyPair) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *GenerateKeyPairs200ResponseKeyPair) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *GenerateKeyPairs200ResponseKeyPair) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *GenerateKeyPairs200ResponseKeyPair) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### SetFingerprintNil

`func (o *GenerateKeyPairs200ResponseKeyPair) SetFingerprintNil(b bool)`

 SetFingerprintNil sets the value for Fingerprint to be an explicit nil

### UnsetFingerprint
`func (o *GenerateKeyPairs200ResponseKeyPair) UnsetFingerprint()`

UnsetFingerprint ensures that no value is present for Fingerprint, not even an explicit nil
### GetDateCreated

`func (o *GenerateKeyPairs200ResponseKeyPair) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GenerateKeyPairs200ResponseKeyPair) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GenerateKeyPairs200ResponseKeyPair) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GenerateKeyPairs200ResponseKeyPair) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GenerateKeyPairs200ResponseKeyPair) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GenerateKeyPairs200ResponseKeyPair) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GenerateKeyPairs200ResponseKeyPair) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GenerateKeyPairs200ResponseKeyPair) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


