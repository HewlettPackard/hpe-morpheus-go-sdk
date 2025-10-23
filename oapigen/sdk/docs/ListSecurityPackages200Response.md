# ListSecurityPackages200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityPackages** | Pointer to [**[]ListSecurityPackages200ResponseAllOfSecurityPackagesInner**](ListSecurityPackages200ResponseAllOfSecurityPackagesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListSecurityPackages200Response

`func NewListSecurityPackages200Response() *ListSecurityPackages200Response`

NewListSecurityPackages200Response instantiates a new ListSecurityPackages200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSecurityPackages200ResponseWithDefaults

`func NewListSecurityPackages200ResponseWithDefaults() *ListSecurityPackages200Response`

NewListSecurityPackages200ResponseWithDefaults instantiates a new ListSecurityPackages200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityPackages

`func (o *ListSecurityPackages200Response) GetSecurityPackages() []ListSecurityPackages200ResponseAllOfSecurityPackagesInner`

GetSecurityPackages returns the SecurityPackages field if non-nil, zero value otherwise.

### GetSecurityPackagesOk

`func (o *ListSecurityPackages200Response) GetSecurityPackagesOk() (*[]ListSecurityPackages200ResponseAllOfSecurityPackagesInner, bool)`

GetSecurityPackagesOk returns a tuple with the SecurityPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityPackages

`func (o *ListSecurityPackages200Response) SetSecurityPackages(v []ListSecurityPackages200ResponseAllOfSecurityPackagesInner)`

SetSecurityPackages sets SecurityPackages field to given value.

### HasSecurityPackages

`func (o *ListSecurityPackages200Response) HasSecurityPackages() bool`

HasSecurityPackages returns a boolean if a field has been set.

### GetMeta

`func (o *ListSecurityPackages200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListSecurityPackages200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListSecurityPackages200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListSecurityPackages200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


