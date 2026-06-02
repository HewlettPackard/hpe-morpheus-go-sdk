# ListDeploymentVersions200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Versions** | Pointer to [**[]ListDeploymentVersions200ResponseAllOfVersionsInner**](ListDeploymentVersions200ResponseAllOfVersionsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListApprovals200ResponseAllOfMeta**](ListApprovals200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListDeploymentVersions200Response

`func NewListDeploymentVersions200Response() *ListDeploymentVersions200Response`

NewListDeploymentVersions200Response instantiates a new ListDeploymentVersions200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetVersions

`func (o *ListDeploymentVersions200Response) GetVersions() []ListDeploymentVersions200ResponseAllOfVersionsInner`

GetVersions returns the Versions field if non-nil, zero value otherwise.

### GetVersionsOk

`func (o *ListDeploymentVersions200Response) GetVersionsOk() (*[]ListDeploymentVersions200ResponseAllOfVersionsInner, bool)`

GetVersionsOk returns a tuple with the Versions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersions

`func (o *ListDeploymentVersions200Response) SetVersions(v []ListDeploymentVersions200ResponseAllOfVersionsInner)`

SetVersions sets Versions field to given value.

### HasVersions

`func (o *ListDeploymentVersions200Response) HasVersions() bool`

HasVersions returns a boolean if a field has been set.

### GetMeta

`func (o *ListDeploymentVersions200Response) GetMeta() ListApprovals200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListDeploymentVersions200Response) GetMetaOk() (*ListApprovals200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListDeploymentVersions200Response) SetMeta(v ListApprovals200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListDeploymentVersions200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


