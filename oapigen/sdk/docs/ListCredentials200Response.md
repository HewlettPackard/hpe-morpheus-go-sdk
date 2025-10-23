# ListCredentials200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Credentials** | Pointer to [**[]ListCredentials200ResponseAllOfCredentialsInner**](ListCredentials200ResponseAllOfCredentialsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListCredentials200Response

`func NewListCredentials200Response() *ListCredentials200Response`

NewListCredentials200Response instantiates a new ListCredentials200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCredentials200ResponseWithDefaults

`func NewListCredentials200ResponseWithDefaults() *ListCredentials200Response`

NewListCredentials200ResponseWithDefaults instantiates a new ListCredentials200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentials

`func (o *ListCredentials200Response) GetCredentials() []ListCredentials200ResponseAllOfCredentialsInner`

GetCredentials returns the Credentials field if non-nil, zero value otherwise.

### GetCredentialsOk

`func (o *ListCredentials200Response) GetCredentialsOk() (*[]ListCredentials200ResponseAllOfCredentialsInner, bool)`

GetCredentialsOk returns a tuple with the Credentials field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentials

`func (o *ListCredentials200Response) SetCredentials(v []ListCredentials200ResponseAllOfCredentialsInner)`

SetCredentials sets Credentials field to given value.

### HasCredentials

`func (o *ListCredentials200Response) HasCredentials() bool`

HasCredentials returns a boolean if a field has been set.

### GetMeta

`func (o *ListCredentials200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListCredentials200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListCredentials200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListCredentials200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


