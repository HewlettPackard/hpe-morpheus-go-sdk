# ListClusterPackages200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClusterPackages** | Pointer to [**[]ListClusterPackages200ResponseAllOfClusterPackagesInner**](ListClusterPackages200ResponseAllOfClusterPackagesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListClusterPackages200Response

`func NewListClusterPackages200Response() *ListClusterPackages200Response`

NewListClusterPackages200Response instantiates a new ListClusterPackages200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterPackages200ResponseWithDefaults

`func NewListClusterPackages200ResponseWithDefaults() *ListClusterPackages200Response`

NewListClusterPackages200ResponseWithDefaults instantiates a new ListClusterPackages200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClusterPackages

`func (o *ListClusterPackages200Response) GetClusterPackages() []ListClusterPackages200ResponseAllOfClusterPackagesInner`

GetClusterPackages returns the ClusterPackages field if non-nil, zero value otherwise.

### GetClusterPackagesOk

`func (o *ListClusterPackages200Response) GetClusterPackagesOk() (*[]ListClusterPackages200ResponseAllOfClusterPackagesInner, bool)`

GetClusterPackagesOk returns a tuple with the ClusterPackages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClusterPackages

`func (o *ListClusterPackages200Response) SetClusterPackages(v []ListClusterPackages200ResponseAllOfClusterPackagesInner)`

SetClusterPackages sets ClusterPackages field to given value.

### HasClusterPackages

`func (o *ListClusterPackages200Response) HasClusterPackages() bool`

HasClusterPackages returns a boolean if a field has been set.

### GetMeta

`func (o *ListClusterPackages200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListClusterPackages200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListClusterPackages200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListClusterPackages200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


