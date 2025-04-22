# AddGroupsRequestGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | A unique name scoped to your account for the group | 
**Code** | Pointer to **string** | Optional code for use with policies | [optional] 
**Labels** | Pointer to **[]string** | Array of label strings, can be used for filtering. | [optional] 
**Location** | Pointer to **string** | Optional location argument for your group | [optional] 
**Config** | Pointer to [**AddGroupsRequestGroupConfig**](AddGroupsRequestGroupConfig.md) |  | [optional] 

## Methods

### NewAddGroupsRequestGroup

`func NewAddGroupsRequestGroup(name string, ) *AddGroupsRequestGroup`

NewAddGroupsRequestGroup instantiates a new AddGroupsRequestGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddGroupsRequestGroupWithDefaults

`func NewAddGroupsRequestGroupWithDefaults() *AddGroupsRequestGroup`

NewAddGroupsRequestGroupWithDefaults instantiates a new AddGroupsRequestGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *AddGroupsRequestGroup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddGroupsRequestGroup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddGroupsRequestGroup) SetName(v string)`

SetName sets Name field to given value.


### GetCode

`func (o *AddGroupsRequestGroup) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *AddGroupsRequestGroup) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *AddGroupsRequestGroup) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *AddGroupsRequestGroup) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetLabels

`func (o *AddGroupsRequestGroup) GetLabels() []string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *AddGroupsRequestGroup) GetLabelsOk() (*[]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *AddGroupsRequestGroup) SetLabels(v []string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *AddGroupsRequestGroup) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetLocation

`func (o *AddGroupsRequestGroup) GetLocation() string`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *AddGroupsRequestGroup) GetLocationOk() (*string, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *AddGroupsRequestGroup) SetLocation(v string)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *AddGroupsRequestGroup) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetConfig

`func (o *AddGroupsRequestGroup) GetConfig() AddGroupsRequestGroupConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *AddGroupsRequestGroup) GetConfigOk() (*AddGroupsRequestGroupConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *AddGroupsRequestGroup) SetConfig(v AddGroupsRequestGroupConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *AddGroupsRequestGroup) HasConfig() bool`

HasConfig returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


