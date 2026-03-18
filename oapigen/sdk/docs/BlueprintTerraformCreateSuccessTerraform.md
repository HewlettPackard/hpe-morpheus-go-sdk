# BlueprintTerraformCreateSuccessTerraform

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ConfigType** | **string** | Configuration Type | 
**Json** | Pointer to **string** | Terraform definition in JSON for &#x60;configType&#x60; &#x60;json&#x60; | [optional] 
**Tf** | Pointer to **string** | Terraform definition for &#x60;configType&#x60; &#x60;tf&#x60; | [optional] 
**Git** | Pointer to [**BlueprintTerraformCreateSuccessTerraformGit**](BlueprintTerraformCreateSuccessTerraformGit.md) |  | [optional] 
**TfvarSecret** | Pointer to **string** | tfvar secret from Morpheus Cypher | [optional] 

## Methods

### NewBlueprintTerraformCreateSuccessTerraform

`func NewBlueprintTerraformCreateSuccessTerraform(configType string, ) *BlueprintTerraformCreateSuccessTerraform`

NewBlueprintTerraformCreateSuccessTerraform instantiates a new BlueprintTerraformCreateSuccessTerraform object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBlueprintTerraformCreateSuccessTerraformWithDefaults

`func NewBlueprintTerraformCreateSuccessTerraformWithDefaults() *BlueprintTerraformCreateSuccessTerraform`

NewBlueprintTerraformCreateSuccessTerraformWithDefaults instantiates a new BlueprintTerraformCreateSuccessTerraform object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfigType

`func (o *BlueprintTerraformCreateSuccessTerraform) GetConfigType() string`

GetConfigType returns the ConfigType field if non-nil, zero value otherwise.

### GetConfigTypeOk

`func (o *BlueprintTerraformCreateSuccessTerraform) GetConfigTypeOk() (*string, bool)`

GetConfigTypeOk returns a tuple with the ConfigType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfigType

`func (o *BlueprintTerraformCreateSuccessTerraform) SetConfigType(v string)`

SetConfigType sets ConfigType field to given value.


### GetJson

`func (o *BlueprintTerraformCreateSuccessTerraform) GetJson() string`

GetJson returns the Json field if non-nil, zero value otherwise.

### GetJsonOk

`func (o *BlueprintTerraformCreateSuccessTerraform) GetJsonOk() (*string, bool)`

GetJsonOk returns a tuple with the Json field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJson

`func (o *BlueprintTerraformCreateSuccessTerraform) SetJson(v string)`

SetJson sets Json field to given value.

### HasJson

`func (o *BlueprintTerraformCreateSuccessTerraform) HasJson() bool`

HasJson returns a boolean if a field has been set.

### GetTf

`func (o *BlueprintTerraformCreateSuccessTerraform) GetTf() string`

GetTf returns the Tf field if non-nil, zero value otherwise.

### GetTfOk

`func (o *BlueprintTerraformCreateSuccessTerraform) GetTfOk() (*string, bool)`

GetTfOk returns a tuple with the Tf field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTf

`func (o *BlueprintTerraformCreateSuccessTerraform) SetTf(v string)`

SetTf sets Tf field to given value.

### HasTf

`func (o *BlueprintTerraformCreateSuccessTerraform) HasTf() bool`

HasTf returns a boolean if a field has been set.

### GetGit

`func (o *BlueprintTerraformCreateSuccessTerraform) GetGit() BlueprintTerraformCreateSuccessTerraformGit`

GetGit returns the Git field if non-nil, zero value otherwise.

### GetGitOk

`func (o *BlueprintTerraformCreateSuccessTerraform) GetGitOk() (*BlueprintTerraformCreateSuccessTerraformGit, bool)`

GetGitOk returns a tuple with the Git field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGit

`func (o *BlueprintTerraformCreateSuccessTerraform) SetGit(v BlueprintTerraformCreateSuccessTerraformGit)`

SetGit sets Git field to given value.

### HasGit

`func (o *BlueprintTerraformCreateSuccessTerraform) HasGit() bool`

HasGit returns a boolean if a field has been set.

### GetTfvarSecret

`func (o *BlueprintTerraformCreateSuccessTerraform) GetTfvarSecret() string`

GetTfvarSecret returns the TfvarSecret field if non-nil, zero value otherwise.

### GetTfvarSecretOk

`func (o *BlueprintTerraformCreateSuccessTerraform) GetTfvarSecretOk() (*string, bool)`

GetTfvarSecretOk returns a tuple with the TfvarSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTfvarSecret

`func (o *BlueprintTerraformCreateSuccessTerraform) SetTfvarSecret(v string)`

SetTfvarSecret sets TfvarSecret field to given value.

### HasTfvarSecret

`func (o *BlueprintTerraformCreateSuccessTerraform) HasTfvarSecret() bool`

HasTfvarSecret returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


