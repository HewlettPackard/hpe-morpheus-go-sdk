# ListDeployments200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deployments** | Pointer to [**[]ListDeployments200ResponseAllOfDeploymentsInner**](ListDeployments200ResponseAllOfDeploymentsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListDeployments200Response

`func NewListDeployments200Response() *ListDeployments200Response`

NewListDeployments200Response instantiates a new ListDeployments200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListDeployments200ResponseWithDefaults

`func NewListDeployments200ResponseWithDefaults() *ListDeployments200Response`

NewListDeployments200ResponseWithDefaults instantiates a new ListDeployments200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeployments

`func (o *ListDeployments200Response) GetDeployments() []ListDeployments200ResponseAllOfDeploymentsInner`

GetDeployments returns the Deployments field if non-nil, zero value otherwise.

### GetDeploymentsOk

`func (o *ListDeployments200Response) GetDeploymentsOk() (*[]ListDeployments200ResponseAllOfDeploymentsInner, bool)`

GetDeploymentsOk returns a tuple with the Deployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployments

`func (o *ListDeployments200Response) SetDeployments(v []ListDeployments200ResponseAllOfDeploymentsInner)`

SetDeployments sets Deployments field to given value.

### HasDeployments

`func (o *ListDeployments200Response) HasDeployments() bool`

HasDeployments returns a boolean if a field has been set.

### GetMeta

`func (o *ListDeployments200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListDeployments200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListDeployments200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListDeployments200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


