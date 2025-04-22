# SetInstanceSecurityGroupsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityGroupIds** | Pointer to **[]int64** | List of all security groups ids which should be applied. If no security groups should apply, pass &#39;[]&#39; | [optional] 

## Methods

### NewSetInstanceSecurityGroupsRequest

`func NewSetInstanceSecurityGroupsRequest() *SetInstanceSecurityGroupsRequest`

NewSetInstanceSecurityGroupsRequest instantiates a new SetInstanceSecurityGroupsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSetInstanceSecurityGroupsRequestWithDefaults

`func NewSetInstanceSecurityGroupsRequestWithDefaults() *SetInstanceSecurityGroupsRequest`

NewSetInstanceSecurityGroupsRequestWithDefaults instantiates a new SetInstanceSecurityGroupsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityGroupIds

`func (o *SetInstanceSecurityGroupsRequest) GetSecurityGroupIds() []int64`

GetSecurityGroupIds returns the SecurityGroupIds field if non-nil, zero value otherwise.

### GetSecurityGroupIdsOk

`func (o *SetInstanceSecurityGroupsRequest) GetSecurityGroupIdsOk() (*[]int64, bool)`

GetSecurityGroupIdsOk returns a tuple with the SecurityGroupIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroupIds

`func (o *SetInstanceSecurityGroupsRequest) SetSecurityGroupIds(v []int64)`

SetSecurityGroupIds sets SecurityGroupIds field to given value.

### HasSecurityGroupIds

`func (o *SetInstanceSecurityGroupsRequest) HasSecurityGroupIds() bool`

HasSecurityGroupIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


