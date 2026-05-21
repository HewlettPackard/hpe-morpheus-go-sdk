# ListClusterSecrets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Secrets** | Pointer to [**[]ListClusterSecrets200ResponseAllOfSecretsInner**](ListClusterSecrets200ResponseAllOfSecretsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListApprovals200ResponseAllOfMeta**](ListApprovals200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListClusterSecrets200Response

`func NewListClusterSecrets200Response() *ListClusterSecrets200Response`

NewListClusterSecrets200Response instantiates a new ListClusterSecrets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterSecrets200ResponseWithDefaults

`func NewListClusterSecrets200ResponseWithDefaults() *ListClusterSecrets200Response`

NewListClusterSecrets200ResponseWithDefaults instantiates a new ListClusterSecrets200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecrets

`func (o *ListClusterSecrets200Response) GetSecrets() []ListClusterSecrets200ResponseAllOfSecretsInner`

GetSecrets returns the Secrets field if non-nil, zero value otherwise.

### GetSecretsOk

`func (o *ListClusterSecrets200Response) GetSecretsOk() (*[]ListClusterSecrets200ResponseAllOfSecretsInner, bool)`

GetSecretsOk returns a tuple with the Secrets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecrets

`func (o *ListClusterSecrets200Response) SetSecrets(v []ListClusterSecrets200ResponseAllOfSecretsInner)`

SetSecrets sets Secrets field to given value.

### HasSecrets

`func (o *ListClusterSecrets200Response) HasSecrets() bool`

HasSecrets returns a boolean if a field has been set.

### GetMeta

`func (o *ListClusterSecrets200Response) GetMeta() ListApprovals200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListClusterSecrets200Response) GetMetaOk() (*ListApprovals200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListClusterSecrets200Response) SetMeta(v ListApprovals200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListClusterSecrets200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


