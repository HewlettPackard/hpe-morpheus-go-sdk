# UpdateCheckGroups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckGroup** | Pointer to [**UpdateCheckGroups200ResponseAllOfCheckGroup**](UpdateCheckGroups200ResponseAllOfCheckGroup.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateCheckGroups200Response

`func NewUpdateCheckGroups200Response() *UpdateCheckGroups200Response`

NewUpdateCheckGroups200Response instantiates a new UpdateCheckGroups200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateCheckGroups200ResponseWithDefaults

`func NewUpdateCheckGroups200ResponseWithDefaults() *UpdateCheckGroups200Response`

NewUpdateCheckGroups200ResponseWithDefaults instantiates a new UpdateCheckGroups200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckGroup

`func (o *UpdateCheckGroups200Response) GetCheckGroup() UpdateCheckGroups200ResponseAllOfCheckGroup`

GetCheckGroup returns the CheckGroup field if non-nil, zero value otherwise.

### GetCheckGroupOk

`func (o *UpdateCheckGroups200Response) GetCheckGroupOk() (*UpdateCheckGroups200ResponseAllOfCheckGroup, bool)`

GetCheckGroupOk returns a tuple with the CheckGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckGroup

`func (o *UpdateCheckGroups200Response) SetCheckGroup(v UpdateCheckGroups200ResponseAllOfCheckGroup)`

SetCheckGroup sets CheckGroup field to given value.

### HasCheckGroup

`func (o *UpdateCheckGroups200Response) HasCheckGroup() bool`

HasCheckGroup returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateCheckGroups200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateCheckGroups200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateCheckGroups200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateCheckGroups200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


