# UpdateDeploymentVersionRequestVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | Pointer to **string** | Version number (userVersion), a unique version identifier for the deployment version. | [optional] 
**UserVersion** | Pointer to **string** | Alias for version | [optional] 
**DeployType** | Pointer to **string** | Deploy Type, eg. file, git, fetch | [optional] 
**GitUrl** | Pointer to **NullableString** |  | [optional] 
**GitRef** | Pointer to **NullableString** |  | [optional] 
**FetchUrl** | Pointer to **NullableString** |  | [optional] 

## Methods

### NewUpdateDeploymentVersionRequestVersion

`func NewUpdateDeploymentVersionRequestVersion() *UpdateDeploymentVersionRequestVersion`

NewUpdateDeploymentVersionRequestVersion instantiates a new UpdateDeploymentVersionRequestVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeploymentVersionRequestVersionWithDefaults

`func NewUpdateDeploymentVersionRequestVersionWithDefaults() *UpdateDeploymentVersionRequestVersion`

NewUpdateDeploymentVersionRequestVersionWithDefaults instantiates a new UpdateDeploymentVersionRequestVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *UpdateDeploymentVersionRequestVersion) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpdateDeploymentVersionRequestVersion) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpdateDeploymentVersionRequestVersion) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UpdateDeploymentVersionRequestVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetUserVersion

`func (o *UpdateDeploymentVersionRequestVersion) GetUserVersion() string`

GetUserVersion returns the UserVersion field if non-nil, zero value otherwise.

### GetUserVersionOk

`func (o *UpdateDeploymentVersionRequestVersion) GetUserVersionOk() (*string, bool)`

GetUserVersionOk returns a tuple with the UserVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVersion

`func (o *UpdateDeploymentVersionRequestVersion) SetUserVersion(v string)`

SetUserVersion sets UserVersion field to given value.

### HasUserVersion

`func (o *UpdateDeploymentVersionRequestVersion) HasUserVersion() bool`

HasUserVersion returns a boolean if a field has been set.

### GetDeployType

`func (o *UpdateDeploymentVersionRequestVersion) GetDeployType() string`

GetDeployType returns the DeployType field if non-nil, zero value otherwise.

### GetDeployTypeOk

`func (o *UpdateDeploymentVersionRequestVersion) GetDeployTypeOk() (*string, bool)`

GetDeployTypeOk returns a tuple with the DeployType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployType

`func (o *UpdateDeploymentVersionRequestVersion) SetDeployType(v string)`

SetDeployType sets DeployType field to given value.

### HasDeployType

`func (o *UpdateDeploymentVersionRequestVersion) HasDeployType() bool`

HasDeployType returns a boolean if a field has been set.

### GetGitUrl

`func (o *UpdateDeploymentVersionRequestVersion) GetGitUrl() string`

GetGitUrl returns the GitUrl field if non-nil, zero value otherwise.

### GetGitUrlOk

`func (o *UpdateDeploymentVersionRequestVersion) GetGitUrlOk() (*string, bool)`

GetGitUrlOk returns a tuple with the GitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitUrl

`func (o *UpdateDeploymentVersionRequestVersion) SetGitUrl(v string)`

SetGitUrl sets GitUrl field to given value.

### HasGitUrl

`func (o *UpdateDeploymentVersionRequestVersion) HasGitUrl() bool`

HasGitUrl returns a boolean if a field has been set.

### SetGitUrlNil

`func (o *UpdateDeploymentVersionRequestVersion) SetGitUrlNil(b bool)`

 SetGitUrlNil sets the value for GitUrl to be an explicit nil

### UnsetGitUrl
`func (o *UpdateDeploymentVersionRequestVersion) UnsetGitUrl()`

UnsetGitUrl ensures that no value is present for GitUrl, not even an explicit nil
### GetGitRef

`func (o *UpdateDeploymentVersionRequestVersion) GetGitRef() string`

GetGitRef returns the GitRef field if non-nil, zero value otherwise.

### GetGitRefOk

`func (o *UpdateDeploymentVersionRequestVersion) GetGitRefOk() (*string, bool)`

GetGitRefOk returns a tuple with the GitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRef

`func (o *UpdateDeploymentVersionRequestVersion) SetGitRef(v string)`

SetGitRef sets GitRef field to given value.

### HasGitRef

`func (o *UpdateDeploymentVersionRequestVersion) HasGitRef() bool`

HasGitRef returns a boolean if a field has been set.

### SetGitRefNil

`func (o *UpdateDeploymentVersionRequestVersion) SetGitRefNil(b bool)`

 SetGitRefNil sets the value for GitRef to be an explicit nil

### UnsetGitRef
`func (o *UpdateDeploymentVersionRequestVersion) UnsetGitRef()`

UnsetGitRef ensures that no value is present for GitRef, not even an explicit nil
### GetFetchUrl

`func (o *UpdateDeploymentVersionRequestVersion) GetFetchUrl() string`

GetFetchUrl returns the FetchUrl field if non-nil, zero value otherwise.

### GetFetchUrlOk

`func (o *UpdateDeploymentVersionRequestVersion) GetFetchUrlOk() (*string, bool)`

GetFetchUrlOk returns a tuple with the FetchUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchUrl

`func (o *UpdateDeploymentVersionRequestVersion) SetFetchUrl(v string)`

SetFetchUrl sets FetchUrl field to given value.

### HasFetchUrl

`func (o *UpdateDeploymentVersionRequestVersion) HasFetchUrl() bool`

HasFetchUrl returns a boolean if a field has been set.

### SetFetchUrlNil

`func (o *UpdateDeploymentVersionRequestVersion) SetFetchUrlNil(b bool)`

 SetFetchUrlNil sets the value for FetchUrl to be an explicit nil

### UnsetFetchUrl
`func (o *UpdateDeploymentVersionRequestVersion) UnsetFetchUrl()`

UnsetFetchUrl ensures that no value is present for FetchUrl, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


