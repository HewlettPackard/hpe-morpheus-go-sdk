# UploadPlugin200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Plugin** | Pointer to [**ListPlugins200ResponseAllOfPluginsInner**](ListPlugins200ResponseAllOfPluginsInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUploadPlugin200Response

`func NewUploadPlugin200Response() *UploadPlugin200Response`

NewUploadPlugin200Response instantiates a new UploadPlugin200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUploadPlugin200ResponseWithDefaults

`func NewUploadPlugin200ResponseWithDefaults() *UploadPlugin200Response`

NewUploadPlugin200ResponseWithDefaults instantiates a new UploadPlugin200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPlugin

`func (o *UploadPlugin200Response) GetPlugin() ListPlugins200ResponseAllOfPluginsInner`

GetPlugin returns the Plugin field if non-nil, zero value otherwise.

### GetPluginOk

`func (o *UploadPlugin200Response) GetPluginOk() (*ListPlugins200ResponseAllOfPluginsInner, bool)`

GetPluginOk returns a tuple with the Plugin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlugin

`func (o *UploadPlugin200Response) SetPlugin(v ListPlugins200ResponseAllOfPluginsInner)`

SetPlugin sets Plugin field to given value.

### HasPlugin

`func (o *UploadPlugin200Response) HasPlugin() bool`

HasPlugin returns a boolean if a field has been set.

### GetSuccess

`func (o *UploadPlugin200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UploadPlugin200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UploadPlugin200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UploadPlugin200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


