# ListClusterVolumes200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Volumes** | Pointer to [**[]ListClusterVolumes200ResponseAllOfVolumesInner**](ListClusterVolumes200ResponseAllOfVolumesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListClusterVolumes200Response

`func NewListClusterVolumes200Response() *ListClusterVolumes200Response`

NewListClusterVolumes200Response instantiates a new ListClusterVolumes200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterVolumes200ResponseWithDefaults

`func NewListClusterVolumes200ResponseWithDefaults() *ListClusterVolumes200Response`

NewListClusterVolumes200ResponseWithDefaults instantiates a new ListClusterVolumes200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVolumes

`func (o *ListClusterVolumes200Response) GetVolumes() []ListClusterVolumes200ResponseAllOfVolumesInner`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *ListClusterVolumes200Response) GetVolumesOk() (*[]ListClusterVolumes200ResponseAllOfVolumesInner, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *ListClusterVolumes200Response) SetVolumes(v []ListClusterVolumes200ResponseAllOfVolumesInner)`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *ListClusterVolumes200Response) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetMeta

`func (o *ListClusterVolumes200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListClusterVolumes200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListClusterVolumes200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListClusterVolumes200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


