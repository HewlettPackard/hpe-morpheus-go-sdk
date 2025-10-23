# ListClusterDeployments200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deployments** | Pointer to [**[]ListClusterDeployments200ResponseAllOfDeploymentsInner**](ListClusterDeployments200ResponseAllOfDeploymentsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListClusterDeployments200Response

`func NewListClusterDeployments200Response() *ListClusterDeployments200Response`

NewListClusterDeployments200Response instantiates a new ListClusterDeployments200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListClusterDeployments200ResponseWithDefaults

`func NewListClusterDeployments200ResponseWithDefaults() *ListClusterDeployments200Response`

NewListClusterDeployments200ResponseWithDefaults instantiates a new ListClusterDeployments200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeployments

`func (o *ListClusterDeployments200Response) GetDeployments() []ListClusterDeployments200ResponseAllOfDeploymentsInner`

GetDeployments returns the Deployments field if non-nil, zero value otherwise.

### GetDeploymentsOk

`func (o *ListClusterDeployments200Response) GetDeploymentsOk() (*[]ListClusterDeployments200ResponseAllOfDeploymentsInner, bool)`

GetDeploymentsOk returns a tuple with the Deployments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployments

`func (o *ListClusterDeployments200Response) SetDeployments(v []ListClusterDeployments200ResponseAllOfDeploymentsInner)`

SetDeployments sets Deployments field to given value.

### HasDeployments

`func (o *ListClusterDeployments200Response) HasDeployments() bool`

HasDeployments returns a boolean if a field has been set.

### GetMeta

`func (o *ListClusterDeployments200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListClusterDeployments200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListClusterDeployments200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListClusterDeployments200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


