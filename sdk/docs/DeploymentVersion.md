# DeploymentVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**DeployType** | Pointer to **string** |  | [optional] 
**DeploymentId** | Pointer to **int64** |  | [optional] 
**FetchUrl** | Pointer to **string** |  | [optional] 
**GitUrl** | Pointer to **string** |  | [optional] 
**GitRef** | Pointer to **string** |  | [optional] 
**UserVersion** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewDeploymentVersion

`func NewDeploymentVersion() *DeploymentVersion`

NewDeploymentVersion instantiates a new DeploymentVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeploymentVersionWithDefaults

`func NewDeploymentVersionWithDefaults() *DeploymentVersion`

NewDeploymentVersionWithDefaults instantiates a new DeploymentVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeploymentVersion) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeploymentVersion) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeploymentVersion) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *DeploymentVersion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeployType

`func (o *DeploymentVersion) GetDeployType() string`

GetDeployType returns the DeployType field if non-nil, zero value otherwise.

### GetDeployTypeOk

`func (o *DeploymentVersion) GetDeployTypeOk() (*string, bool)`

GetDeployTypeOk returns a tuple with the DeployType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployType

`func (o *DeploymentVersion) SetDeployType(v string)`

SetDeployType sets DeployType field to given value.

### HasDeployType

`func (o *DeploymentVersion) HasDeployType() bool`

HasDeployType returns a boolean if a field has been set.

### GetDeploymentId

`func (o *DeploymentVersion) GetDeploymentId() int64`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *DeploymentVersion) GetDeploymentIdOk() (*int64, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *DeploymentVersion) SetDeploymentId(v int64)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *DeploymentVersion) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### GetFetchUrl

`func (o *DeploymentVersion) GetFetchUrl() string`

GetFetchUrl returns the FetchUrl field if non-nil, zero value otherwise.

### GetFetchUrlOk

`func (o *DeploymentVersion) GetFetchUrlOk() (*string, bool)`

GetFetchUrlOk returns a tuple with the FetchUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchUrl

`func (o *DeploymentVersion) SetFetchUrl(v string)`

SetFetchUrl sets FetchUrl field to given value.

### HasFetchUrl

`func (o *DeploymentVersion) HasFetchUrl() bool`

HasFetchUrl returns a boolean if a field has been set.

### GetGitUrl

`func (o *DeploymentVersion) GetGitUrl() string`

GetGitUrl returns the GitUrl field if non-nil, zero value otherwise.

### GetGitUrlOk

`func (o *DeploymentVersion) GetGitUrlOk() (*string, bool)`

GetGitUrlOk returns a tuple with the GitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitUrl

`func (o *DeploymentVersion) SetGitUrl(v string)`

SetGitUrl sets GitUrl field to given value.

### HasGitUrl

`func (o *DeploymentVersion) HasGitUrl() bool`

HasGitUrl returns a boolean if a field has been set.

### GetGitRef

`func (o *DeploymentVersion) GetGitRef() string`

GetGitRef returns the GitRef field if non-nil, zero value otherwise.

### GetGitRefOk

`func (o *DeploymentVersion) GetGitRefOk() (*string, bool)`

GetGitRefOk returns a tuple with the GitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRef

`func (o *DeploymentVersion) SetGitRef(v string)`

SetGitRef sets GitRef field to given value.

### HasGitRef

`func (o *DeploymentVersion) HasGitRef() bool`

HasGitRef returns a boolean if a field has been set.

### GetUserVersion

`func (o *DeploymentVersion) GetUserVersion() string`

GetUserVersion returns the UserVersion field if non-nil, zero value otherwise.

### GetUserVersionOk

`func (o *DeploymentVersion) GetUserVersionOk() (*string, bool)`

GetUserVersionOk returns a tuple with the UserVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVersion

`func (o *DeploymentVersion) SetUserVersion(v string)`

SetUserVersion sets UserVersion field to given value.

### HasUserVersion

`func (o *DeploymentVersion) HasUserVersion() bool`

HasUserVersion returns a boolean if a field has been set.

### GetVersion

`func (o *DeploymentVersion) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DeploymentVersion) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DeploymentVersion) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *DeploymentVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetStatus

`func (o *DeploymentVersion) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeploymentVersion) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeploymentVersion) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeploymentVersion) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDateCreated

`func (o *DeploymentVersion) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *DeploymentVersion) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *DeploymentVersion) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *DeploymentVersion) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *DeploymentVersion) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *DeploymentVersion) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *DeploymentVersion) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *DeploymentVersion) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


