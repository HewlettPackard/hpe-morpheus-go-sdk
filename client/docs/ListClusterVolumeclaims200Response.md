# ListClusterVolumeclaims200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Volumeclaims** | Pointer to [**[]ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner**](ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListClusterVolumeclaims200Response

`func NewListClusterVolumeclaims200Response() *ListClusterVolumeclaims200Response`

NewListClusterVolumeclaims200Response instantiates a new ListClusterVolumeclaims200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterVolumeclaims200ResponseWithDefaults

`func NewListClusterVolumeclaims200ResponseWithDefaults() *ListClusterVolumeclaims200Response`

NewListClusterVolumeclaims200ResponseWithDefaults instantiates a new ListClusterVolumeclaims200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolumeclaims

`func (o *ListClusterVolumeclaims200Response) GetVolumeclaims() []ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner`

GetVolumeclaims returns the Volumeclaims field if non-nil, zero value otherwise.

### GetVolumeclaimsOk

`func (o *ListClusterVolumeclaims200Response) GetVolumeclaimsOk() (*[]ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner, bool)`

GetVolumeclaimsOk returns a tuple with the Volumeclaims field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumeclaims

`func (o *ListClusterVolumeclaims200Response) SetVolumeclaims(v []ListClusterVolumeclaims200ResponseAllOfVolumeclaimsInner)`

SetVolumeclaims sets Volumeclaims field to given value.

### HasVolumeclaims

`func (o *ListClusterVolumeclaims200Response) HasVolumeclaims() bool`

HasVolumeclaims returns a boolean if a field has been set.

### GetMeta

`func (o *ListClusterVolumeclaims200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListClusterVolumeclaims200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListClusterVolumeclaims200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListClusterVolumeclaims200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


