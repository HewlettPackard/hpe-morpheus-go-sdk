# ListSecurityPackageTypes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityPackageTypes** | Pointer to [**[]ListSecurityPackageTypes200ResponseAllOfSecurityPackageTypesInner**](ListSecurityPackageTypes200ResponseAllOfSecurityPackageTypesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListSecurityPackageTypes200Response

`func NewListSecurityPackageTypes200Response() *ListSecurityPackageTypes200Response`

NewListSecurityPackageTypes200Response instantiates a new ListSecurityPackageTypes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSecurityPackageTypes200ResponseWithDefaults

`func NewListSecurityPackageTypes200ResponseWithDefaults() *ListSecurityPackageTypes200Response`

NewListSecurityPackageTypes200ResponseWithDefaults instantiates a new ListSecurityPackageTypes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityPackageTypes

`func (o *ListSecurityPackageTypes200Response) GetSecurityPackageTypes() []ListSecurityPackageTypes200ResponseAllOfSecurityPackageTypesInner`

GetSecurityPackageTypes returns the SecurityPackageTypes field if non-nil, zero value otherwise.

### GetSecurityPackageTypesOk

`func (o *ListSecurityPackageTypes200Response) GetSecurityPackageTypesOk() (*[]ListSecurityPackageTypes200ResponseAllOfSecurityPackageTypesInner, bool)`

GetSecurityPackageTypesOk returns a tuple with the SecurityPackageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPackageTypes

`func (o *ListSecurityPackageTypes200Response) SetSecurityPackageTypes(v []ListSecurityPackageTypes200ResponseAllOfSecurityPackageTypesInner)`

SetSecurityPackageTypes sets SecurityPackageTypes field to given value.

### HasSecurityPackageTypes

`func (o *ListSecurityPackageTypes200Response) HasSecurityPackageTypes() bool`

HasSecurityPackageTypes returns a boolean if a field has been set.

### GetMeta

`func (o *ListSecurityPackageTypes200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListSecurityPackageTypes200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListSecurityPackageTypes200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListSecurityPackageTypes200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


