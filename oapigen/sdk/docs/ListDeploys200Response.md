# ListDeploys200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppDeploys** | Pointer to [**[]ListDeploys200ResponseAllOfAppDeploysInner**](ListDeploys200ResponseAllOfAppDeploysInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListDeploys200Response

`func NewListDeploys200Response() *ListDeploys200Response`

NewListDeploys200Response instantiates a new ListDeploys200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDeploys200ResponseWithDefaults

`func NewListDeploys200ResponseWithDefaults() *ListDeploys200Response`

NewListDeploys200ResponseWithDefaults instantiates a new ListDeploys200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppDeploys

`func (o *ListDeploys200Response) GetAppDeploys() []ListDeploys200ResponseAllOfAppDeploysInner`

GetAppDeploys returns the AppDeploys field if non-nil, zero value otherwise.

### GetAppDeploysOk

`func (o *ListDeploys200Response) GetAppDeploysOk() (*[]ListDeploys200ResponseAllOfAppDeploysInner, bool)`

GetAppDeploysOk returns a tuple with the AppDeploys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppDeploys

`func (o *ListDeploys200Response) SetAppDeploys(v []ListDeploys200ResponseAllOfAppDeploysInner)`

SetAppDeploys sets AppDeploys field to given value.

### HasAppDeploys

`func (o *ListDeploys200Response) HasAppDeploys() bool`

HasAppDeploys returns a boolean if a field has been set.

### GetMeta

`func (o *ListDeploys200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListDeploys200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListDeploys200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListDeploys200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


