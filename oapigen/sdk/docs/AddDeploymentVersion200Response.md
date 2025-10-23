# AddDeploymentVersion200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to [**ListDeploymentVersions200ResponseAllOfVersionsInner**](ListDeploymentVersions200ResponseAllOfVersionsInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddDeploymentVersion200Response

`func NewAddDeploymentVersion200Response() *AddDeploymentVersion200Response`

NewAddDeploymentVersion200Response instantiates a new AddDeploymentVersion200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddDeploymentVersion200ResponseWithDefaults

`func NewAddDeploymentVersion200ResponseWithDefaults() *AddDeploymentVersion200Response`

NewAddDeploymentVersion200ResponseWithDefaults instantiates a new AddDeploymentVersion200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *AddDeploymentVersion200Response) GetVersion() ListDeploymentVersions200ResponseAllOfVersionsInner`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *AddDeploymentVersion200Response) GetVersionOk() (*ListDeploymentVersions200ResponseAllOfVersionsInner, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *AddDeploymentVersion200Response) SetVersion(v ListDeploymentVersions200ResponseAllOfVersionsInner)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *AddDeploymentVersion200Response) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetSuccess

`func (o *AddDeploymentVersion200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddDeploymentVersion200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddDeploymentVersion200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddDeploymentVersion200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


