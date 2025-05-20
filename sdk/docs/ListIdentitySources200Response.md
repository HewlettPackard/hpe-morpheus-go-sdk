# ListIdentitySources200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserSources** | Pointer to [**[]ListIdentitySources200ResponseAllOfUserSourcesInner**](ListIdentitySources200ResponseAllOfUserSourcesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListIdentitySources200Response

`func NewListIdentitySources200Response() *ListIdentitySources200Response`

NewListIdentitySources200Response instantiates a new ListIdentitySources200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListIdentitySources200ResponseWithDefaults

`func NewListIdentitySources200ResponseWithDefaults() *ListIdentitySources200Response`

NewListIdentitySources200ResponseWithDefaults instantiates a new ListIdentitySources200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserSources

`func (o *ListIdentitySources200Response) GetUserSources() []ListIdentitySources200ResponseAllOfUserSourcesInner`

GetUserSources returns the UserSources field if non-nil, zero value otherwise.

### GetUserSourcesOk

`func (o *ListIdentitySources200Response) GetUserSourcesOk() (*[]ListIdentitySources200ResponseAllOfUserSourcesInner, bool)`

GetUserSourcesOk returns a tuple with the UserSources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSources

`func (o *ListIdentitySources200Response) SetUserSources(v []ListIdentitySources200ResponseAllOfUserSourcesInner)`

SetUserSources sets UserSources field to given value.

### HasUserSources

`func (o *ListIdentitySources200Response) HasUserSources() bool`

HasUserSources returns a boolean if a field has been set.

### GetMeta

`func (o *ListIdentitySources200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListIdentitySources200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListIdentitySources200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListIdentitySources200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


