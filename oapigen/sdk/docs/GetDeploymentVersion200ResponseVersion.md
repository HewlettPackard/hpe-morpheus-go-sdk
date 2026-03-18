# GetDeploymentVersion200ResponseVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**DeployType** | Pointer to **string** |  | [optional] 
**DeploymentId** | Pointer to **int64** |  | [optional] 
**FetchUrl** | Pointer to **NullableString** |  | [optional] 
**GitUrl** | Pointer to **NullableString** |  | [optional] 
**GitRef** | Pointer to **NullableString** |  | [optional] 
**UserVersion** | Pointer to **string** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**DateCreated** | Pointer to **time.Time** |  | [optional] 
**LastUpdated** | Pointer to **time.Time** |  | [optional] 

## Methods

### NewGetDeploymentVersion200ResponseVersion

`func NewGetDeploymentVersion200ResponseVersion() *GetDeploymentVersion200ResponseVersion`

NewGetDeploymentVersion200ResponseVersion instantiates a new GetDeploymentVersion200ResponseVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetDeploymentVersion200ResponseVersionWithDefaults

`func NewGetDeploymentVersion200ResponseVersionWithDefaults() *GetDeploymentVersion200ResponseVersion`

NewGetDeploymentVersion200ResponseVersionWithDefaults instantiates a new GetDeploymentVersion200ResponseVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *GetDeploymentVersion200ResponseVersion) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GetDeploymentVersion200ResponseVersion) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GetDeploymentVersion200ResponseVersion) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *GetDeploymentVersion200ResponseVersion) HasId() bool`

HasId returns a boolean if a field has been set.

### GetDeployType

`func (o *GetDeploymentVersion200ResponseVersion) GetDeployType() string`

GetDeployType returns the DeployType field if non-nil, zero value otherwise.

### GetDeployTypeOk

`func (o *GetDeploymentVersion200ResponseVersion) GetDeployTypeOk() (*string, bool)`

GetDeployTypeOk returns a tuple with the DeployType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeployType

`func (o *GetDeploymentVersion200ResponseVersion) SetDeployType(v string)`

SetDeployType sets DeployType field to given value.

### HasDeployType

`func (o *GetDeploymentVersion200ResponseVersion) HasDeployType() bool`

HasDeployType returns a boolean if a field has been set.

### GetDeploymentId

`func (o *GetDeploymentVersion200ResponseVersion) GetDeploymentId() int64`

GetDeploymentId returns the DeploymentId field if non-nil, zero value otherwise.

### GetDeploymentIdOk

`func (o *GetDeploymentVersion200ResponseVersion) GetDeploymentIdOk() (*int64, bool)`

GetDeploymentIdOk returns a tuple with the DeploymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentId

`func (o *GetDeploymentVersion200ResponseVersion) SetDeploymentId(v int64)`

SetDeploymentId sets DeploymentId field to given value.

### HasDeploymentId

`func (o *GetDeploymentVersion200ResponseVersion) HasDeploymentId() bool`

HasDeploymentId returns a boolean if a field has been set.

### GetFetchUrl

`func (o *GetDeploymentVersion200ResponseVersion) GetFetchUrl() string`

GetFetchUrl returns the FetchUrl field if non-nil, zero value otherwise.

### GetFetchUrlOk

`func (o *GetDeploymentVersion200ResponseVersion) GetFetchUrlOk() (*string, bool)`

GetFetchUrlOk returns a tuple with the FetchUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFetchUrl

`func (o *GetDeploymentVersion200ResponseVersion) SetFetchUrl(v string)`

SetFetchUrl sets FetchUrl field to given value.

### HasFetchUrl

`func (o *GetDeploymentVersion200ResponseVersion) HasFetchUrl() bool`

HasFetchUrl returns a boolean if a field has been set.

### SetFetchUrlNil

`func (o *GetDeploymentVersion200ResponseVersion) SetFetchUrlNil(b bool)`

 SetFetchUrlNil sets the value for FetchUrl to be an explicit nil

### UnsetFetchUrl
`func (o *GetDeploymentVersion200ResponseVersion) UnsetFetchUrl()`

UnsetFetchUrl ensures that no value is present for FetchUrl, not even an explicit nil
### GetGitUrl

`func (o *GetDeploymentVersion200ResponseVersion) GetGitUrl() string`

GetGitUrl returns the GitUrl field if non-nil, zero value otherwise.

### GetGitUrlOk

`func (o *GetDeploymentVersion200ResponseVersion) GetGitUrlOk() (*string, bool)`

GetGitUrlOk returns a tuple with the GitUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitUrl

`func (o *GetDeploymentVersion200ResponseVersion) SetGitUrl(v string)`

SetGitUrl sets GitUrl field to given value.

### HasGitUrl

`func (o *GetDeploymentVersion200ResponseVersion) HasGitUrl() bool`

HasGitUrl returns a boolean if a field has been set.

### SetGitUrlNil

`func (o *GetDeploymentVersion200ResponseVersion) SetGitUrlNil(b bool)`

 SetGitUrlNil sets the value for GitUrl to be an explicit nil

### UnsetGitUrl
`func (o *GetDeploymentVersion200ResponseVersion) UnsetGitUrl()`

UnsetGitUrl ensures that no value is present for GitUrl, not even an explicit nil
### GetGitRef

`func (o *GetDeploymentVersion200ResponseVersion) GetGitRef() string`

GetGitRef returns the GitRef field if non-nil, zero value otherwise.

### GetGitRefOk

`func (o *GetDeploymentVersion200ResponseVersion) GetGitRefOk() (*string, bool)`

GetGitRefOk returns a tuple with the GitRef field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitRef

`func (o *GetDeploymentVersion200ResponseVersion) SetGitRef(v string)`

SetGitRef sets GitRef field to given value.

### HasGitRef

`func (o *GetDeploymentVersion200ResponseVersion) HasGitRef() bool`

HasGitRef returns a boolean if a field has been set.

### SetGitRefNil

`func (o *GetDeploymentVersion200ResponseVersion) SetGitRefNil(b bool)`

 SetGitRefNil sets the value for GitRef to be an explicit nil

### UnsetGitRef
`func (o *GetDeploymentVersion200ResponseVersion) UnsetGitRef()`

UnsetGitRef ensures that no value is present for GitRef, not even an explicit nil
### GetUserVersion

`func (o *GetDeploymentVersion200ResponseVersion) GetUserVersion() string`

GetUserVersion returns the UserVersion field if non-nil, zero value otherwise.

### GetUserVersionOk

`func (o *GetDeploymentVersion200ResponseVersion) GetUserVersionOk() (*string, bool)`

GetUserVersionOk returns a tuple with the UserVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserVersion

`func (o *GetDeploymentVersion200ResponseVersion) SetUserVersion(v string)`

SetUserVersion sets UserVersion field to given value.

### HasUserVersion

`func (o *GetDeploymentVersion200ResponseVersion) HasUserVersion() bool`

HasUserVersion returns a boolean if a field has been set.

### GetVersion

`func (o *GetDeploymentVersion200ResponseVersion) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *GetDeploymentVersion200ResponseVersion) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *GetDeploymentVersion200ResponseVersion) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *GetDeploymentVersion200ResponseVersion) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetStatus

`func (o *GetDeploymentVersion200ResponseVersion) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *GetDeploymentVersion200ResponseVersion) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *GetDeploymentVersion200ResponseVersion) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *GetDeploymentVersion200ResponseVersion) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDateCreated

`func (o *GetDeploymentVersion200ResponseVersion) GetDateCreated() time.Time`

GetDateCreated returns the DateCreated field if non-nil, zero value otherwise.

### GetDateCreatedOk

`func (o *GetDeploymentVersion200ResponseVersion) GetDateCreatedOk() (*time.Time, bool)`

GetDateCreatedOk returns a tuple with the DateCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDateCreated

`func (o *GetDeploymentVersion200ResponseVersion) SetDateCreated(v time.Time)`

SetDateCreated sets DateCreated field to given value.

### HasDateCreated

`func (o *GetDeploymentVersion200ResponseVersion) HasDateCreated() bool`

HasDateCreated returns a boolean if a field has been set.

### GetLastUpdated

`func (o *GetDeploymentVersion200ResponseVersion) GetLastUpdated() time.Time`

GetLastUpdated returns the LastUpdated field if non-nil, zero value otherwise.

### GetLastUpdatedOk

`func (o *GetDeploymentVersion200ResponseVersion) GetLastUpdatedOk() (*time.Time, bool)`

GetLastUpdatedOk returns a tuple with the LastUpdated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdated

`func (o *GetDeploymentVersion200ResponseVersion) SetLastUpdated(v time.Time)`

SetLastUpdated sets LastUpdated field to given value.

### HasLastUpdated

`func (o *GetDeploymentVersion200ResponseVersion) HasLastUpdated() bool`

HasLastUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


