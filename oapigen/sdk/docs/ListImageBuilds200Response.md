# ListImageBuilds200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageBuilds** | Pointer to [**[]ListImageBuilds200ResponseAllOfImageBuildsInner**](ListImageBuilds200ResponseAllOfImageBuildsInner.md) |  | [optional] 
**ImageBuildCount** | Pointer to **int64** |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListImageBuilds200Response

`func NewListImageBuilds200Response() *ListImageBuilds200Response`

NewListImageBuilds200Response instantiates a new ListImageBuilds200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListImageBuilds200ResponseWithDefaults

`func NewListImageBuilds200ResponseWithDefaults() *ListImageBuilds200Response`

NewListImageBuilds200ResponseWithDefaults instantiates a new ListImageBuilds200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageBuilds

`func (o *ListImageBuilds200Response) GetImageBuilds() []ListImageBuilds200ResponseAllOfImageBuildsInner`

GetImageBuilds returns the ImageBuilds field if non-nil, zero value otherwise.

### GetImageBuildsOk

`func (o *ListImageBuilds200Response) GetImageBuildsOk() (*[]ListImageBuilds200ResponseAllOfImageBuildsInner, bool)`

GetImageBuildsOk returns a tuple with the ImageBuilds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBuilds

`func (o *ListImageBuilds200Response) SetImageBuilds(v []ListImageBuilds200ResponseAllOfImageBuildsInner)`

SetImageBuilds sets ImageBuilds field to given value.

### HasImageBuilds

`func (o *ListImageBuilds200Response) HasImageBuilds() bool`

HasImageBuilds returns a boolean if a field has been set.

### GetImageBuildCount

`func (o *ListImageBuilds200Response) GetImageBuildCount() int64`

GetImageBuildCount returns the ImageBuildCount field if non-nil, zero value otherwise.

### GetImageBuildCountOk

`func (o *ListImageBuilds200Response) GetImageBuildCountOk() (*int64, bool)`

GetImageBuildCountOk returns a tuple with the ImageBuildCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBuildCount

`func (o *ListImageBuilds200Response) SetImageBuildCount(v int64)`

SetImageBuildCount sets ImageBuildCount field to given value.

### HasImageBuildCount

`func (o *ListImageBuilds200Response) HasImageBuildCount() bool`

HasImageBuildCount returns a boolean if a field has been set.

### GetMeta

`func (o *ListImageBuilds200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListImageBuilds200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListImageBuilds200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListImageBuilds200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


