# GetImageBuild200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ImageBuild** | Pointer to [**AddImageBuild200ResponseAllOfImageBuild**](AddImageBuild200ResponseAllOfImageBuild.md) |  | [optional] 
**ImageBuildExecutions** | Pointer to [**[]GetImageBuild200ResponseImageBuildExecutionsInner**](GetImageBuild200ResponseImageBuildExecutionsInner.md) |  | [optional] 

## Methods

### NewGetImageBuild200Response

`func NewGetImageBuild200Response() *GetImageBuild200Response`

NewGetImageBuild200Response instantiates a new GetImageBuild200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetImageBuild200ResponseWithDefaults

`func NewGetImageBuild200ResponseWithDefaults() *GetImageBuild200Response`

NewGetImageBuild200ResponseWithDefaults instantiates a new GetImageBuild200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetImageBuild

`func (o *GetImageBuild200Response) GetImageBuild() AddImageBuild200ResponseAllOfImageBuild`

GetImageBuild returns the ImageBuild field if non-nil, zero value otherwise.

### GetImageBuildOk

`func (o *GetImageBuild200Response) GetImageBuildOk() (*AddImageBuild200ResponseAllOfImageBuild, bool)`

GetImageBuildOk returns a tuple with the ImageBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBuild

`func (o *GetImageBuild200Response) SetImageBuild(v AddImageBuild200ResponseAllOfImageBuild)`

SetImageBuild sets ImageBuild field to given value.

### HasImageBuild

`func (o *GetImageBuild200Response) HasImageBuild() bool`

HasImageBuild returns a boolean if a field has been set.

### GetImageBuildExecutions

`func (o *GetImageBuild200Response) GetImageBuildExecutions() []GetImageBuild200ResponseImageBuildExecutionsInner`

GetImageBuildExecutions returns the ImageBuildExecutions field if non-nil, zero value otherwise.

### GetImageBuildExecutionsOk

`func (o *GetImageBuild200Response) GetImageBuildExecutionsOk() (*[]GetImageBuild200ResponseImageBuildExecutionsInner, bool)`

GetImageBuildExecutionsOk returns a tuple with the ImageBuildExecutions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImageBuildExecutions

`func (o *GetImageBuild200Response) SetImageBuildExecutions(v []GetImageBuild200ResponseImageBuildExecutionsInner)`

SetImageBuildExecutions sets ImageBuildExecutions field to given value.

### HasImageBuildExecutions

`func (o *GetImageBuild200Response) HasImageBuildExecutions() bool`

HasImageBuildExecutions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


