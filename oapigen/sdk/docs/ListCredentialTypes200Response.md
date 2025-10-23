# ListCredentialTypes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CredentialTypes** | Pointer to [**[]ListCredentialTypes200ResponseAllOfCredentialTypesInner**](ListCredentialTypes200ResponseAllOfCredentialTypesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListCredentialTypes200Response

`func NewListCredentialTypes200Response() *ListCredentialTypes200Response`

NewListCredentialTypes200Response instantiates a new ListCredentialTypes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCredentialTypes200ResponseWithDefaults

`func NewListCredentialTypes200ResponseWithDefaults() *ListCredentialTypes200Response`

NewListCredentialTypes200ResponseWithDefaults instantiates a new ListCredentialTypes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCredentialTypes

`func (o *ListCredentialTypes200Response) GetCredentialTypes() []ListCredentialTypes200ResponseAllOfCredentialTypesInner`

GetCredentialTypes returns the CredentialTypes field if non-nil, zero value otherwise.

### GetCredentialTypesOk

`func (o *ListCredentialTypes200Response) GetCredentialTypesOk() (*[]ListCredentialTypes200ResponseAllOfCredentialTypesInner, bool)`

GetCredentialTypesOk returns a tuple with the CredentialTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCredentialTypes

`func (o *ListCredentialTypes200Response) SetCredentialTypes(v []ListCredentialTypes200ResponseAllOfCredentialTypesInner)`

SetCredentialTypes sets CredentialTypes field to given value.

### HasCredentialTypes

`func (o *ListCredentialTypes200Response) HasCredentialTypes() bool`

HasCredentialTypes returns a boolean if a field has been set.

### GetMeta

`func (o *ListCredentialTypes200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListCredentialTypes200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListCredentialTypes200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListCredentialTypes200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


