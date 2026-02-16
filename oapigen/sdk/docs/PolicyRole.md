# PolicyRole

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int64** |  | [optional] 
**Authority** | Pointer to **string** |  | [optional] 

## Methods

### NewPolicyRole

`func NewPolicyRole() *PolicyRole`

NewPolicyRole instantiates a new PolicyRole object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPolicyRoleWithDefaults

`func NewPolicyRoleWithDefaults() *PolicyRole`

NewPolicyRoleWithDefaults instantiates a new PolicyRole object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *PolicyRole) GetId() int64`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PolicyRole) GetIdOk() (*int64, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PolicyRole) SetId(v int64)`

SetId sets Id field to given value.

### HasId

`func (o *PolicyRole) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAuthority

`func (o *PolicyRole) GetAuthority() string`

GetAuthority returns the Authority field if non-nil, zero value otherwise.

### GetAuthorityOk

`func (o *PolicyRole) GetAuthorityOk() (*string, bool)`

GetAuthorityOk returns a tuple with the Authority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthority

`func (o *PolicyRole) SetAuthority(v string)`

SetAuthority sets Authority field to given value.

### HasAuthority

`func (o *PolicyRole) HasAuthority() bool`

HasAuthority returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


