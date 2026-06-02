# GetInstanceDeploys200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppDeploys** | Pointer to [**[]GetInstanceDeploys200ResponseAllOfAppDeploysInner**](GetInstanceDeploys200ResponseAllOfAppDeploysInner.md) |  | [optional] 
**Meta** | Pointer to [**GetInstanceDeploys200ResponseAllOfMeta**](GetInstanceDeploys200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewGetInstanceDeploys200Response

`func NewGetInstanceDeploys200Response() *GetInstanceDeploys200Response`

NewGetInstanceDeploys200Response instantiates a new GetInstanceDeploys200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetAppDeploys

`func (o *GetInstanceDeploys200Response) GetAppDeploys() []GetInstanceDeploys200ResponseAllOfAppDeploysInner`

GetAppDeploys returns the AppDeploys field if non-nil, zero value otherwise.

### GetAppDeploysOk

`func (o *GetInstanceDeploys200Response) GetAppDeploysOk() (*[]GetInstanceDeploys200ResponseAllOfAppDeploysInner, bool)`

GetAppDeploysOk returns a tuple with the AppDeploys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDeploys

`func (o *GetInstanceDeploys200Response) SetAppDeploys(v []GetInstanceDeploys200ResponseAllOfAppDeploysInner)`

SetAppDeploys sets AppDeploys field to given value.

### HasAppDeploys

`func (o *GetInstanceDeploys200Response) HasAppDeploys() bool`

HasAppDeploys returns a boolean if a field has been set.

### GetMeta

`func (o *GetInstanceDeploys200Response) GetMeta() GetInstanceDeploys200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetInstanceDeploys200Response) GetMetaOk() (*GetInstanceDeploys200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetInstanceDeploys200Response) SetMeta(v GetInstanceDeploys200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetInstanceDeploys200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


