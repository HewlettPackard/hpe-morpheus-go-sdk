# GenericInstanceConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserData** | Pointer to **string** | User Data. Allows for override of cloud-init based user-data yaml or custom scripts | [optional] 

## Methods

### NewGenericInstanceConfiguration

`func NewGenericInstanceConfiguration() *GenericInstanceConfiguration`

NewGenericInstanceConfiguration instantiates a new GenericInstanceConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericInstanceConfigurationWithDefaults

`func NewGenericInstanceConfigurationWithDefaults() *GenericInstanceConfiguration`

NewGenericInstanceConfigurationWithDefaults instantiates a new GenericInstanceConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserData

`func (o *GenericInstanceConfiguration) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *GenericInstanceConfiguration) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *GenericInstanceConfiguration) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *GenericInstanceConfiguration) HasUserData() bool`

HasUserData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


