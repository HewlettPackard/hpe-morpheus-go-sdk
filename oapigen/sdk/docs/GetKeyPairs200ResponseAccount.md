# GetKeyPairs200ResponseAccount

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

### NewGetKeyPairs200ResponseAccount

`func NewGetKeyPairs200ResponseAccount() *GetKeyPairs200ResponseAccount`

NewGetKeyPairs200ResponseAccount instantiates a new GetKeyPairs200ResponseAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetId

`func (o *GetKeyPairs200ResponseAccount) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetKeyPairs200ResponseAccount) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetKeyPairs200ResponseAccount) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetKeyPairs200ResponseAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *GetKeyPairs200ResponseAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *GetKeyPairs200ResponseAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *GetKeyPairs200ResponseAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *GetKeyPairs200ResponseAccount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccountId

`func (o *GetKeyPairs200ResponseAccount) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *GetKeyPairs200ResponseAccount) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *GetKeyPairs200ResponseAccount) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *GetKeyPairs200ResponseAccount) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPublicKey

`func (o *GetKeyPairs200ResponseAccount) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *GetKeyPairs200ResponseAccount) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *GetKeyPairs200ResponseAccount) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *GetKeyPairs200ResponseAccount) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### SetPublicKeyNil

`func (o *GetKeyPairs200ResponseAccount) SetPublicKeyNil(b bool)`

 SetPublicKeyNil sets the value for PublicKey to be an explicit nil

### UnsetPublicKey
`func (o *GetKeyPairs200ResponseAccount) UnsetPublicKey()`

UnsetPublicKey ensures that no value is present for PublicKey, not even an explicit nil
### GetHasPrivateKey

`func (o *GetKeyPairs200ResponseAccount) GetHasPrivateKey() bool`

GetHasPrivateKey returns the HasPrivateKey field if non-nil, zero value otherwise.

### GetHasPrivateKeyOk

`func (o *GetKeyPairs200ResponseAccount) GetHasPrivateKeyOk() (*bool, bool)`

GetHasPrivateKeyOk returns a tuple with the HasPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrivateKey

`func (o *GetKeyPairs200ResponseAccount) SetHasPrivateKey(v bool)`

SetHasPrivateKey sets HasPrivateKey field to given value.

### HasHasPrivateKey

`func (o *GetKeyPairs200ResponseAccount) HasHasPrivateKey() bool`

HasHasPrivateKey returns a boolean if a field has been set.

### GetPrivateKeyHash

`func (o *GetKeyPairs200ResponseAccount) GetPrivateKeyHash() string`

GetPrivateKeyHash returns the PrivateKeyHash field if non-nil, zero value otherwise.

### GetPrivateKeyHashOk

`func (o *GetKeyPairs200ResponseAccount) GetPrivateKeyHashOk() (*string, bool)`

GetPrivateKeyHashOk returns a tuple with the PrivateKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyHash

`func (o *GetKeyPairs200ResponseAccount) SetPrivateKeyHash(v string)`

SetPrivateKeyHash sets PrivateKeyHash field to given value.

### HasPrivateKeyHash

`func (o *GetKeyPairs200ResponseAccount) HasPrivateKeyHash() bool`

HasPrivateKeyHash returns a boolean if a field has been set.

### SetPrivateKeyHashNil

`func (o *GetKeyPairs200ResponseAccount) SetPrivateKeyHashNil(b bool)`

 SetPrivateKeyHashNil sets the value for PrivateKeyHash to be an explicit nil

### UnsetPrivateKeyHash
`func (o *GetKeyPairs200ResponseAccount) UnsetPrivateKeyHash()`

UnsetPrivateKeyHash ensures that no value is present for PrivateKeyHash, not even an explicit nil
### GetPrivateKey

`func (o *GetKeyPairs200ResponseAccount) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *GetKeyPairs200ResponseAccount) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *GetKeyPairs200ResponseAccount) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *GetKeyPairs200ResponseAccount) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### SetPrivateKeyNil

`func (o *GetKeyPairs200ResponseAccount) SetPrivateKeyNil(b bool)`

 SetPrivateKeyNil sets the value for PrivateKey to be an explicit nil

### UnsetPrivateKey
`func (o *GetKeyPairs200ResponseAccount) UnsetPrivateKey()`

UnsetPrivateKey ensures that no value is present for PrivateKey, not even an explicit nil
### GetFingerprint

`func (o *GetKeyPairs200ResponseAccount) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *GetKeyPairs200ResponseAccount) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *GetKeyPairs200ResponseAccount) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *GetKeyPairs200ResponseAccount) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### SetFingerprintNil

`func (o *GetKeyPairs200ResponseAccount) SetFingerprintNil(b bool)`

 SetFingerprintNil sets the value for Fingerprint to be an explicit nil

### UnsetFingerprint
`func (o *GetKeyPairs200ResponseAccount) UnsetFingerprint()`

UnsetFingerprint ensures that no value is present for Fingerprint, not even an explicit nil
### GetDateCreated

`func (o *GetKeyPairs200ResponseAccount) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetKeyPairs200ResponseAccount) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetKeyPairs200ResponseAccount) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetKeyPairs200ResponseAccount) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetKeyPairs200ResponseAccount) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetKeyPairs200ResponseAccount) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetKeyPairs200ResponseAccount) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetKeyPairs200ResponseAccount) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


