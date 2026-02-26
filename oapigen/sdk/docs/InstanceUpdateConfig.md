# InstanceUpdateConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CustomOptions** | Pointer to **map[string]interface{}** | Custom Option Type settings object containing name value pairs. | [optional] 
**UserData** | Pointer to **NullableString** | User Data. Allows for override of cloud-init based user-data yaml or custom scripts | [optional] 

## Methods

### NewInstanceUpdateConfig

`func NewInstanceUpdateConfig() *InstanceUpdateConfig`

NewInstanceUpdateConfig instantiates a new InstanceUpdateConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstanceUpdateConfigWithDefaults

`func NewInstanceUpdateConfigWithDefaults() *InstanceUpdateConfig`

NewInstanceUpdateConfigWithDefaults instantiates a new InstanceUpdateConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCustomOptions

`func (o *InstanceUpdateConfig) GetCustomOptions() map[string]interface{}`

GetCustomOptions returns the CustomOptions field if non-nil, zero value otherwise.

### GetCustomOptionsOk

`func (o *InstanceUpdateConfig) GetCustomOptionsOk() (*map[string]interface{}, bool)`

GetCustomOptionsOk returns a tuple with the CustomOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomOptions

`func (o *InstanceUpdateConfig) SetCustomOptions(v map[string]interface{})`

SetCustomOptions sets CustomOptions field to given value.

### HasCustomOptions

`func (o *InstanceUpdateConfig) HasCustomOptions() bool`

HasCustomOptions returns a boolean if a field has been set.

### GetUserData

`func (o *InstanceUpdateConfig) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *InstanceUpdateConfig) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *InstanceUpdateConfig) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *InstanceUpdateConfig) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### SetUserDataNil

`func (o *InstanceUpdateConfig) SetUserDataNil(b bool)`

 SetUserDataNil sets the value for UserData to be an explicit nil

### UnsetUserData
`func (o *InstanceUpdateConfig) UnsetUserData()`

UnsetUserData ensures that no value is present for UserData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


