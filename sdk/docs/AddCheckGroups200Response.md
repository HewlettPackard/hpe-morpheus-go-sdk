# AddCheckGroups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckGroup** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInner**](GetAlerts200ResponseAllOfCheckGroupsInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddCheckGroups200Response

`func NewAddCheckGroups200Response() *AddCheckGroups200Response`

NewAddCheckGroups200Response instantiates a new AddCheckGroups200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddCheckGroups200ResponseWithDefaults

`func NewAddCheckGroups200ResponseWithDefaults() *AddCheckGroups200Response`

NewAddCheckGroups200ResponseWithDefaults instantiates a new AddCheckGroups200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckGroup

`func (o *AddCheckGroups200Response) GetCheckGroup() GetAlerts200ResponseAllOfCheckGroupsInner`

GetCheckGroup returns the CheckGroup field if non-nil, zero value otherwise.

### GetCheckGroupOk

`func (o *AddCheckGroups200Response) GetCheckGroupOk() (*GetAlerts200ResponseAllOfCheckGroupsInner, bool)`

GetCheckGroupOk returns a tuple with the CheckGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckGroup

`func (o *AddCheckGroups200Response) SetCheckGroup(v GetAlerts200ResponseAllOfCheckGroupsInner)`

SetCheckGroup sets CheckGroup field to given value.

### HasCheckGroup

`func (o *AddCheckGroups200Response) HasCheckGroup() bool`

HasCheckGroup returns a boolean if a field has been set.

### GetSuccess

`func (o *AddCheckGroups200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddCheckGroups200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddCheckGroups200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddCheckGroups200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


