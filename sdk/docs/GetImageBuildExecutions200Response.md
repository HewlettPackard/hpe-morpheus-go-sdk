# GetImageBuildExecutions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageBuildExecutions** | Pointer to [**[]GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner**](GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner.md) |  | [optional] 
**ImageBuildExecutionCount** | Pointer to **int64** |  | [optional] 
**ImageBuild** | Pointer to [**GetImageBuild200ResponseImageBuild**](GetImageBuild200ResponseImageBuild.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewGetImageBuildExecutions200Response

`func NewGetImageBuildExecutions200Response() *GetImageBuildExecutions200Response`

NewGetImageBuildExecutions200Response instantiates a new GetImageBuildExecutions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetImageBuildExecutions200ResponseWithDefaults

`func NewGetImageBuildExecutions200ResponseWithDefaults() *GetImageBuildExecutions200Response`

NewGetImageBuildExecutions200ResponseWithDefaults instantiates a new GetImageBuildExecutions200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageBuildExecutions

`func (o *GetImageBuildExecutions200Response) GetImageBuildExecutions() []GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner`

GetImageBuildExecutions returns the ImageBuildExecutions field if non-nil, zero value otherwise.

### GetImageBuildExecutionsOk

`func (o *GetImageBuildExecutions200Response) GetImageBuildExecutionsOk() (*[]GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner, bool)`

GetImageBuildExecutionsOk returns a tuple with the ImageBuildExecutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBuildExecutions

`func (o *GetImageBuildExecutions200Response) SetImageBuildExecutions(v []GetImageBuildExecutions200ResponseAllOfImageBuildExecutionsInner)`

SetImageBuildExecutions sets ImageBuildExecutions field to given value.

### HasImageBuildExecutions

`func (o *GetImageBuildExecutions200Response) HasImageBuildExecutions() bool`

HasImageBuildExecutions returns a boolean if a field has been set.

### GetImageBuildExecutionCount

`func (o *GetImageBuildExecutions200Response) GetImageBuildExecutionCount() int64`

GetImageBuildExecutionCount returns the ImageBuildExecutionCount field if non-nil, zero value otherwise.

### GetImageBuildExecutionCountOk

`func (o *GetImageBuildExecutions200Response) GetImageBuildExecutionCountOk() (*int64, bool)`

GetImageBuildExecutionCountOk returns a tuple with the ImageBuildExecutionCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBuildExecutionCount

`func (o *GetImageBuildExecutions200Response) SetImageBuildExecutionCount(v int64)`

SetImageBuildExecutionCount sets ImageBuildExecutionCount field to given value.

### HasImageBuildExecutionCount

`func (o *GetImageBuildExecutions200Response) HasImageBuildExecutionCount() bool`

HasImageBuildExecutionCount returns a boolean if a field has been set.

### GetImageBuild

`func (o *GetImageBuildExecutions200Response) GetImageBuild() GetImageBuild200ResponseImageBuild`

GetImageBuild returns the ImageBuild field if non-nil, zero value otherwise.

### GetImageBuildOk

`func (o *GetImageBuildExecutions200Response) GetImageBuildOk() (*GetImageBuild200ResponseImageBuild, bool)`

GetImageBuildOk returns a tuple with the ImageBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBuild

`func (o *GetImageBuildExecutions200Response) SetImageBuild(v GetImageBuild200ResponseImageBuild)`

SetImageBuild sets ImageBuild field to given value.

### HasImageBuild

`func (o *GetImageBuildExecutions200Response) HasImageBuild() bool`

HasImageBuild returns a boolean if a field has been set.

### GetMeta

`func (o *GetImageBuildExecutions200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetImageBuildExecutions200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetImageBuildExecutions200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetImageBuildExecutions200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


