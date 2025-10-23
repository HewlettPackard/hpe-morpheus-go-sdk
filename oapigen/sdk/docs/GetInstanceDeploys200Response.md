# GetInstanceDeploys200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppDeploys** | Pointer to [**[]UpdateDeploy200ResponseAppDeploy**](UpdateDeploy200ResponseAppDeploy.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewGetInstanceDeploys200Response

`func NewGetInstanceDeploys200Response() *GetInstanceDeploys200Response`

NewGetInstanceDeploys200Response instantiates a new GetInstanceDeploys200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInstanceDeploys200ResponseWithDefaults

`func NewGetInstanceDeploys200ResponseWithDefaults() *GetInstanceDeploys200Response`

NewGetInstanceDeploys200ResponseWithDefaults instantiates a new GetInstanceDeploys200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppDeploys

`func (o *GetInstanceDeploys200Response) GetAppDeploys() []UpdateDeploy200ResponseAppDeploy`

GetAppDeploys returns the AppDeploys field if non-nil, zero value otherwise.

### GetAppDeploysOk

`func (o *GetInstanceDeploys200Response) GetAppDeploysOk() (*[]UpdateDeploy200ResponseAppDeploy, bool)`

GetAppDeploysOk returns a tuple with the AppDeploys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDeploys

`func (o *GetInstanceDeploys200Response) SetAppDeploys(v []UpdateDeploy200ResponseAppDeploy)`

SetAppDeploys sets AppDeploys field to given value.

### HasAppDeploys

`func (o *GetInstanceDeploys200Response) HasAppDeploys() bool`

HasAppDeploys returns a boolean if a field has been set.

### GetMeta

`func (o *GetInstanceDeploys200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetInstanceDeploys200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetInstanceDeploys200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetInstanceDeploys200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


