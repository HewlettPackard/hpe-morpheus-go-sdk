# DeploymentVersionCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | Version number (userVersion), a unique version identifier for the deployment version. | [optional] 
**UserVersion** | Pointer to **string** | Alias for version | [optional] 
**DeployType** | Pointer to **string** | Deploy Type, eg. file, git, fetch | [optional] 
**GitUrl** | Pointer to **string** |  | [optional] 
**GitRef** | Pointer to **string** |  | [optional] 
**FetchUrl** | Pointer to **string** |  | [optional] 

## Methods

### NewDeploymentVersionCreate

`func NewDeploymentVersionCreate() *DeploymentVersionCreate`

NewDeploymentVersionCreate instantiates a new DeploymentVersionCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentVersionCreateWithDefaults

`func NewDeploymentVersionCreateWithDefaults() *DeploymentVersionCreate`

NewDeploymentVersionCreateWithDefaults instantiates a new DeploymentVersionCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *DeploymentVersionCreate) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DeploymentVersionCreate) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DeploymentVersionCreate) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DeploymentVersionCreate) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetUserVersion

`func (o *DeploymentVersionCreate) GetUserVersion() string`

GetUserVersion returns the UserVersion field if non-nil, zero value otherwise.

### GetUserVersionOk

`func (o *DeploymentVersionCreate) GetUserVersionOk() (*string, bool)`

GetUserVersionOk returns a tuple with the UserVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVersion

`func (o *DeploymentVersionCreate) SetUserVersion(v string)`

SetUserVersion sets UserVersion field to given value.

### HasUserVersion

`func (o *DeploymentVersionCreate) HasUserVersion() bool`

HasUserVersion returns a boolean if a field has been set.

### GetDeployType

`func (o *DeploymentVersionCreate) GetDeployType() string`

GetDeployType returns the DeployType field if non-nil, zero value otherwise.

### GetDeployTypeOk

`func (o *DeploymentVersionCreate) GetDeployTypeOk() (*string, bool)`

GetDeployTypeOk returns a tuple with the DeployType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployType

`func (o *DeploymentVersionCreate) SetDeployType(v string)`

SetDeployType sets DeployType field to given value.

### HasDeployType

`func (o *DeploymentVersionCreate) HasDeployType() bool`

HasDeployType returns a boolean if a field has been set.

### GetGitUrl

`func (o *DeploymentVersionCreate) GetGitUrl() string`

GetGitUrl returns the GitUrl field if non-nil, zero value otherwise.

### GetGitUrlOk

`func (o *DeploymentVersionCreate) GetGitUrlOk() (*string, bool)`

GetGitUrlOk returns a tuple with the GitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitUrl

`func (o *DeploymentVersionCreate) SetGitUrl(v string)`

SetGitUrl sets GitUrl field to given value.

### HasGitUrl

`func (o *DeploymentVersionCreate) HasGitUrl() bool`

HasGitUrl returns a boolean if a field has been set.

### GetGitRef

`func (o *DeploymentVersionCreate) GetGitRef() string`

GetGitRef returns the GitRef field if non-nil, zero value otherwise.

### GetGitRefOk

`func (o *DeploymentVersionCreate) GetGitRefOk() (*string, bool)`

GetGitRefOk returns a tuple with the GitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRef

`func (o *DeploymentVersionCreate) SetGitRef(v string)`

SetGitRef sets GitRef field to given value.

### HasGitRef

`func (o *DeploymentVersionCreate) HasGitRef() bool`

HasGitRef returns a boolean if a field has been set.

### GetFetchUrl

`func (o *DeploymentVersionCreate) GetFetchUrl() string`

GetFetchUrl returns the FetchUrl field if non-nil, zero value otherwise.

### GetFetchUrlOk

`func (o *DeploymentVersionCreate) GetFetchUrlOk() (*string, bool)`

GetFetchUrlOk returns a tuple with the FetchUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchUrl

`func (o *DeploymentVersionCreate) SetFetchUrl(v string)`

SetFetchUrl sets FetchUrl field to given value.

### HasFetchUrl

`func (o *DeploymentVersionCreate) HasFetchUrl() bool`

HasFetchUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


