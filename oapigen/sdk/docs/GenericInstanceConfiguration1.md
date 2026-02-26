# GenericInstanceConfiguration1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**UserData** | Pointer to **NullableString** | User Data. Allows for override of cloud-init based user-data yaml or custom scripts | [optional] 

## Methods

### NewGenericInstanceConfiguration1

`func NewGenericInstanceConfiguration1() *GenericInstanceConfiguration1`

NewGenericInstanceConfiguration1 instantiates a new GenericInstanceConfiguration1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGenericInstanceConfiguration1WithDefaults

`func NewGenericInstanceConfiguration1WithDefaults() *GenericInstanceConfiguration1`

NewGenericInstanceConfiguration1WithDefaults instantiates a new GenericInstanceConfiguration1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUserData

`func (o *GenericInstanceConfiguration1) GetUserData() string`

GetUserData returns the UserData field if non-nil, zero value otherwise.

### GetUserDataOk

`func (o *GenericInstanceConfiguration1) GetUserDataOk() (*string, bool)`

GetUserDataOk returns a tuple with the UserData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserData

`func (o *GenericInstanceConfiguration1) SetUserData(v string)`

SetUserData sets UserData field to given value.

### HasUserData

`func (o *GenericInstanceConfiguration1) HasUserData() bool`

HasUserData returns a boolean if a field has been set.

### SetUserDataNil

`func (o *GenericInstanceConfiguration1) SetUserDataNil(b bool)`

 SetUserDataNil sets the value for UserData to be an explicit nil

### UnsetUserData
`func (o *GenericInstanceConfiguration1) UnsetUserData()`

UnsetUserData ensures that no value is present for UserData, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


