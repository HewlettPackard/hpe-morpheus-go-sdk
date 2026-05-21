# ListSupportBundles200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportBundles** | Pointer to [**[]ListSupportBundles200ResponseAllOfSupportBundlesInner**](ListSupportBundles200ResponseAllOfSupportBundlesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListApprovals200ResponseAllOfMeta**](ListApprovals200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListSupportBundles200Response

`func NewListSupportBundles200Response() *ListSupportBundles200Response`

NewListSupportBundles200Response instantiates a new ListSupportBundles200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListSupportBundles200ResponseWithDefaults

`func NewListSupportBundles200ResponseWithDefaults() *ListSupportBundles200Response`

NewListSupportBundles200ResponseWithDefaults instantiates a new ListSupportBundles200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportBundles

`func (o *ListSupportBundles200Response) GetSupportBundles() []ListSupportBundles200ResponseAllOfSupportBundlesInner`

GetSupportBundles returns the SupportBundles field if non-nil, zero value otherwise.

### GetSupportBundlesOk

`func (o *ListSupportBundles200Response) GetSupportBundlesOk() (*[]ListSupportBundles200ResponseAllOfSupportBundlesInner, bool)`

GetSupportBundlesOk returns a tuple with the SupportBundles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportBundles

`func (o *ListSupportBundles200Response) SetSupportBundles(v []ListSupportBundles200ResponseAllOfSupportBundlesInner)`

SetSupportBundles sets SupportBundles field to given value.

### HasSupportBundles

`func (o *ListSupportBundles200Response) HasSupportBundles() bool`

HasSupportBundles returns a boolean if a field has been set.

### GetMeta

`func (o *ListSupportBundles200Response) GetMeta() ListApprovals200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListSupportBundles200Response) GetMetaOk() (*ListApprovals200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListSupportBundles200Response) SetMeta(v ListApprovals200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListSupportBundles200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


