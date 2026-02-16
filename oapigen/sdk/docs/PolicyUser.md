# PolicyUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Username** | Pointer to **string** |  | [optional] 

## Methods

### NewPolicyUser

`func NewPolicyUser() *PolicyUser`

NewPolicyUser instantiates a new PolicyUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyUserWithDefaults

`func NewPolicyUserWithDefaults() *PolicyUser`

NewPolicyUserWithDefaults instantiates a new PolicyUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PolicyUser) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyUser) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyUser) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PolicyUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetUsername

`func (o *PolicyUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *PolicyUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *PolicyUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *PolicyUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


