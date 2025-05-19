# GetCheckGroups200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CheckGroup** | Pointer to [**GetAlerts200ResponseAllOfCheckGroupsInner**](GetAlerts200ResponseAllOfCheckGroupsInner.md) |  | [optional] 
**Checks** | Pointer to [**[]GetAlerts200ResponseAllOfChecksInner**](GetAlerts200ResponseAllOfChecksInner.md) |  | [optional] 

## Methods

### NewGetCheckGroups200Response

`func NewGetCheckGroups200Response() *GetCheckGroups200Response`

NewGetCheckGroups200Response instantiates a new GetCheckGroups200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetCheckGroups200ResponseWithDefaults

`func NewGetCheckGroups200ResponseWithDefaults() *GetCheckGroups200Response`

NewGetCheckGroups200ResponseWithDefaults instantiates a new GetCheckGroups200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheckGroup

`func (o *GetCheckGroups200Response) GetCheckGroup() GetAlerts200ResponseAllOfCheckGroupsInner`

GetCheckGroup returns the CheckGroup field if non-nil, zero value otherwise.

### GetCheckGroupOk

`func (o *GetCheckGroups200Response) GetCheckGroupOk() (*GetAlerts200ResponseAllOfCheckGroupsInner, bool)`

GetCheckGroupOk returns a tuple with the CheckGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheckGroup

`func (o *GetCheckGroups200Response) SetCheckGroup(v GetAlerts200ResponseAllOfCheckGroupsInner)`

SetCheckGroup sets CheckGroup field to given value.

### HasCheckGroup

`func (o *GetCheckGroups200Response) HasCheckGroup() bool`

HasCheckGroup returns a boolean if a field has been set.

### GetChecks

`func (o *GetCheckGroups200Response) GetChecks() []GetAlerts200ResponseAllOfChecksInner`

GetChecks returns the Checks field if non-nil, zero value otherwise.

### GetChecksOk

`func (o *GetCheckGroups200Response) GetChecksOk() (*[]GetAlerts200ResponseAllOfChecksInner, bool)`

GetChecksOk returns a tuple with the Checks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecks

`func (o *GetCheckGroups200Response) SetChecks(v []GetAlerts200ResponseAllOfChecksInner)`

SetChecks sets Checks field to given value.

### HasChecks

`func (o *GetCheckGroups200Response) HasChecks() bool`

HasChecks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


