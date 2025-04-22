# AddKeyPairs200ResponseAllOfAccount

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

### NewAddKeyPairs200ResponseAllOfAccount

`func NewAddKeyPairs200ResponseAllOfAccount() *AddKeyPairs200ResponseAllOfAccount`

NewAddKeyPairs200ResponseAllOfAccount instantiates a new AddKeyPairs200ResponseAllOfAccount object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddKeyPairs200ResponseAllOfAccountWithDefaults

`func NewAddKeyPairs200ResponseAllOfAccountWithDefaults() *AddKeyPairs200ResponseAllOfAccount`

NewAddKeyPairs200ResponseAllOfAccountWithDefaults instantiates a new AddKeyPairs200ResponseAllOfAccount object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AddKeyPairs200ResponseAllOfAccount) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddKeyPairs200ResponseAllOfAccount) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddKeyPairs200ResponseAllOfAccount) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddKeyPairs200ResponseAllOfAccount) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *AddKeyPairs200ResponseAllOfAccount) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddKeyPairs200ResponseAllOfAccount) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddKeyPairs200ResponseAllOfAccount) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AddKeyPairs200ResponseAllOfAccount) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAccountId

`func (o *AddKeyPairs200ResponseAllOfAccount) GetAccountId() int64`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *AddKeyPairs200ResponseAllOfAccount) GetAccountIdOk() (*int64, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *AddKeyPairs200ResponseAllOfAccount) SetAccountId(v int64)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *AddKeyPairs200ResponseAllOfAccount) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### GetPublicKey

`func (o *AddKeyPairs200ResponseAllOfAccount) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *AddKeyPairs200ResponseAllOfAccount) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *AddKeyPairs200ResponseAllOfAccount) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *AddKeyPairs200ResponseAllOfAccount) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetHasPrivateKey

`func (o *AddKeyPairs200ResponseAllOfAccount) GetHasPrivateKey() bool`

GetHasPrivateKey returns the HasPrivateKey field if non-nil, zero value otherwise.

### GetHasPrivateKeyOk

`func (o *AddKeyPairs200ResponseAllOfAccount) GetHasPrivateKeyOk() (*bool, bool)`

GetHasPrivateKeyOk returns a tuple with the HasPrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasPrivateKey

`func (o *AddKeyPairs200ResponseAllOfAccount) SetHasPrivateKey(v bool)`

SetHasPrivateKey sets HasPrivateKey field to given value.

### HasHasPrivateKey

`func (o *AddKeyPairs200ResponseAllOfAccount) HasHasPrivateKey() bool`

HasHasPrivateKey returns a boolean if a field has been set.

### GetPrivateKeyHash

`func (o *AddKeyPairs200ResponseAllOfAccount) GetPrivateKeyHash() string`

GetPrivateKeyHash returns the PrivateKeyHash field if non-nil, zero value otherwise.

### GetPrivateKeyHashOk

`func (o *AddKeyPairs200ResponseAllOfAccount) GetPrivateKeyHashOk() (*string, bool)`

GetPrivateKeyHashOk returns a tuple with the PrivateKeyHash field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKeyHash

`func (o *AddKeyPairs200ResponseAllOfAccount) SetPrivateKeyHash(v string)`

SetPrivateKeyHash sets PrivateKeyHash field to given value.

### HasPrivateKeyHash

`func (o *AddKeyPairs200ResponseAllOfAccount) HasPrivateKeyHash() bool`

HasPrivateKeyHash returns a boolean if a field has been set.

### GetPrivateKey

`func (o *AddKeyPairs200ResponseAllOfAccount) GetPrivateKey() string`

GetPrivateKey returns the PrivateKey field if non-nil, zero value otherwise.

### GetPrivateKeyOk

`func (o *AddKeyPairs200ResponseAllOfAccount) GetPrivateKeyOk() (*string, bool)`

GetPrivateKeyOk returns a tuple with the PrivateKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrivateKey

`func (o *AddKeyPairs200ResponseAllOfAccount) SetPrivateKey(v string)`

SetPrivateKey sets PrivateKey field to given value.

### HasPrivateKey

`func (o *AddKeyPairs200ResponseAllOfAccount) HasPrivateKey() bool`

HasPrivateKey returns a boolean if a field has been set.

### GetFingerprint

`func (o *AddKeyPairs200ResponseAllOfAccount) GetFingerprint() string`

GetFingerprint returns the Fingerprint field if non-nil, zero value otherwise.

### GetFingerprintOk

`func (o *AddKeyPairs200ResponseAllOfAccount) GetFingerprintOk() (*string, bool)`

GetFingerprintOk returns a tuple with the Fingerprint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFingerprint

`func (o *AddKeyPairs200ResponseAllOfAccount) SetFingerprint(v string)`

SetFingerprint sets Fingerprint field to given value.

### HasFingerprint

`func (o *AddKeyPairs200ResponseAllOfAccount) HasFingerprint() bool`

HasFingerprint returns a boolean if a field has been set.

### GetDateCreated

`func (o *AddKeyPairs200ResponseAllOfAccount) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *AddKeyPairs200ResponseAllOfAccount) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *AddKeyPairs200ResponseAllOfAccount) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *AddKeyPairs200ResponseAllOfAccount) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *AddKeyPairs200ResponseAllOfAccount) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *AddKeyPairs200ResponseAllOfAccount) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *AddKeyPairs200ResponseAllOfAccount) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *AddKeyPairs200ResponseAllOfAccount) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


