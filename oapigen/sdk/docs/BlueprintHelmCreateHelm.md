# BlueprintHelmCreateHelm

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigType** | **string** | Configuration Type | 
**Git** | Pointer to [**BlueprintHelmCreateHelmGit**](BlueprintHelmCreateHelmGit.md) |  | [optional] 

## Methods

### NewBlueprintHelmCreateHelm

`func NewBlueprintHelmCreateHelm(configType string, ) *BlueprintHelmCreateHelm`

NewBlueprintHelmCreateHelm instantiates a new BlueprintHelmCreateHelm object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintHelmCreateHelmWithDefaults

`func NewBlueprintHelmCreateHelmWithDefaults() *BlueprintHelmCreateHelm`

NewBlueprintHelmCreateHelmWithDefaults instantiates a new BlueprintHelmCreateHelm object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigType

`func (o *BlueprintHelmCreateHelm) GetConfigType() string`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *BlueprintHelmCreateHelm) GetConfigTypeOk() (*string, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *BlueprintHelmCreateHelm) SetConfigType(v string)`

SetConfigType sets ConfigType field to given value.


### GetGit

`func (o *BlueprintHelmCreateHelm) GetGit() BlueprintHelmCreateHelmGit`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *BlueprintHelmCreateHelm) GetGitOk() (*BlueprintHelmCreateHelmGit, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *BlueprintHelmCreateHelm) SetGit(v BlueprintHelmCreateHelmGit)`

SetGit sets Git field to given value.

### HasGit

`func (o *BlueprintHelmCreateHelm) HasGit() bool`

HasGit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


