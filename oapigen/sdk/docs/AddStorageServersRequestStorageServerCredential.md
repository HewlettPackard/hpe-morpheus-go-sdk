# AddStorageServersRequestStorageServerCredential

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | Credential type code or &#x60;local&#x60; | [optional] 
**Id** | Pointer to **int64** | Credential ID (for non-local types) | [optional] 

## Methods

### NewAddStorageServersRequestStorageServerCredential

`func NewAddStorageServersRequestStorageServerCredential() *AddStorageServersRequestStorageServerCredential`

NewAddStorageServersRequestStorageServerCredential instantiates a new AddStorageServersRequestStorageServerCredential object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddStorageServersRequestStorageServerCredentialWithDefaults

`func NewAddStorageServersRequestStorageServerCredentialWithDefaults() *AddStorageServersRequestStorageServerCredential`

NewAddStorageServersRequestStorageServerCredentialWithDefaults instantiates a new AddStorageServersRequestStorageServerCredential object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AddStorageServersRequestStorageServerCredential) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AddStorageServersRequestStorageServerCredential) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AddStorageServersRequestStorageServerCredential) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *AddStorageServersRequestStorageServerCredential) HasType() bool`

HasType returns a boolean if a field has been set.

### GetId

`func (o *AddStorageServersRequestStorageServerCredential) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AddStorageServersRequestStorageServerCredential) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AddStorageServersRequestStorageServerCredential) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *AddStorageServersRequestStorageServerCredential) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


