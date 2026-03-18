# UserSourceCreateCustomExternal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LoginUrl** | Pointer to **string** | External Login URL | [optional] 
**DoNotIncludeSAMLRequest** | Pointer to **bool** | Do not include SAMLRequest | [optional] [default to false]
**Logout** | Pointer to **string** | External Logout URL | [optional] 
**EncryptionAlgo** | Pointer to **string** | Encryption Algorithm | [optional] 
**EncryptionKey** | Pointer to **string** | Encryption Key | [optional] 

## Methods

### NewUserSourceCreateCustomExternal

`func NewUserSourceCreateCustomExternal() *UserSourceCreateCustomExternal`

NewUserSourceCreateCustomExternal instantiates a new UserSourceCreateCustomExternal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserSourceCreateCustomExternalWithDefaults

`func NewUserSourceCreateCustomExternalWithDefaults() *UserSourceCreateCustomExternal`

NewUserSourceCreateCustomExternalWithDefaults instantiates a new UserSourceCreateCustomExternal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLoginUrl

`func (o *UserSourceCreateCustomExternal) GetLoginUrl() string`

GetLoginUrl returns the LoginUrl field if non-nil, zero value otherwise.

### GetLoginUrlOk

`func (o *UserSourceCreateCustomExternal) GetLoginUrlOk() (*string, bool)`

GetLoginUrlOk returns a tuple with the LoginUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginUrl

`func (o *UserSourceCreateCustomExternal) SetLoginUrl(v string)`

SetLoginUrl sets LoginUrl field to given value.

### HasLoginUrl

`func (o *UserSourceCreateCustomExternal) HasLoginUrl() bool`

HasLoginUrl returns a boolean if a field has been set.

### GetDoNotIncludeSAMLRequest

`func (o *UserSourceCreateCustomExternal) GetDoNotIncludeSAMLRequest() bool`

GetDoNotIncludeSAMLRequest returns the DoNotIncludeSAMLRequest field if non-nil, zero value otherwise.

### GetDoNotIncludeSAMLRequestOk

`func (o *UserSourceCreateCustomExternal) GetDoNotIncludeSAMLRequestOk() (*bool, bool)`

GetDoNotIncludeSAMLRequestOk returns a tuple with the DoNotIncludeSAMLRequest field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDoNotIncludeSAMLRequest

`func (o *UserSourceCreateCustomExternal) SetDoNotIncludeSAMLRequest(v bool)`

SetDoNotIncludeSAMLRequest sets DoNotIncludeSAMLRequest field to given value.

### HasDoNotIncludeSAMLRequest

`func (o *UserSourceCreateCustomExternal) HasDoNotIncludeSAMLRequest() bool`

HasDoNotIncludeSAMLRequest returns a boolean if a field has been set.

### GetLogout

`func (o *UserSourceCreateCustomExternal) GetLogout() string`

GetLogout returns the Logout field if non-nil, zero value otherwise.

### GetLogoutOk

`func (o *UserSourceCreateCustomExternal) GetLogoutOk() (*string, bool)`

GetLogoutOk returns a tuple with the Logout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogout

`func (o *UserSourceCreateCustomExternal) SetLogout(v string)`

SetLogout sets Logout field to given value.

### HasLogout

`func (o *UserSourceCreateCustomExternal) HasLogout() bool`

HasLogout returns a boolean if a field has been set.

### GetEncryptionAlgo

`func (o *UserSourceCreateCustomExternal) GetEncryptionAlgo() string`

GetEncryptionAlgo returns the EncryptionAlgo field if non-nil, zero value otherwise.

### GetEncryptionAlgoOk

`func (o *UserSourceCreateCustomExternal) GetEncryptionAlgoOk() (*string, bool)`

GetEncryptionAlgoOk returns a tuple with the EncryptionAlgo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionAlgo

`func (o *UserSourceCreateCustomExternal) SetEncryptionAlgo(v string)`

SetEncryptionAlgo sets EncryptionAlgo field to given value.

### HasEncryptionAlgo

`func (o *UserSourceCreateCustomExternal) HasEncryptionAlgo() bool`

HasEncryptionAlgo returns a boolean if a field has been set.

### GetEncryptionKey

`func (o *UserSourceCreateCustomExternal) GetEncryptionKey() string`

GetEncryptionKey returns the EncryptionKey field if non-nil, zero value otherwise.

### GetEncryptionKeyOk

`func (o *UserSourceCreateCustomExternal) GetEncryptionKeyOk() (*string, bool)`

GetEncryptionKeyOk returns a tuple with the EncryptionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncryptionKey

`func (o *UserSourceCreateCustomExternal) SetEncryptionKey(v string)`

SetEncryptionKey sets EncryptionKey field to given value.

### HasEncryptionKey

`func (o *UserSourceCreateCustomExternal) HasEncryptionKey() bool`

HasEncryptionKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


