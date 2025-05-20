# AddGroups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Group** | Pointer to [**ListGroups200ResponseAllOfGroupsInner**](ListGroups200ResponseAllOfGroupsInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddGroups200Response

`func NewAddGroups200Response() *AddGroups200Response`

NewAddGroups200Response instantiates a new AddGroups200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddGroups200ResponseWithDefaults

`func NewAddGroups200ResponseWithDefaults() *AddGroups200Response`

NewAddGroups200ResponseWithDefaults instantiates a new AddGroups200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroup

`func (o *AddGroups200Response) GetGroup() ListGroups200ResponseAllOfGroupsInner`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *AddGroups200Response) GetGroupOk() (*ListGroups200ResponseAllOfGroupsInner, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *AddGroups200Response) SetGroup(v ListGroups200ResponseAllOfGroupsInner)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *AddGroups200Response) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetSuccess

`func (o *AddGroups200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddGroups200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddGroups200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddGroups200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


