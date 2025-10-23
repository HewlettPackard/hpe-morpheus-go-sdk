# UpdateSecurityGroupsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SecurityGroup** | [**UpdateSecurityGroupsRequestSecurityGroup**](UpdateSecurityGroupsRequestSecurityGroup.md) |  | 

## Methods

### NewUpdateSecurityGroupsRequest

`func NewUpdateSecurityGroupsRequest(securityGroup UpdateSecurityGroupsRequestSecurityGroup, ) *UpdateSecurityGroupsRequest`

NewUpdateSecurityGroupsRequest instantiates a new UpdateSecurityGroupsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateSecurityGroupsRequestWithDefaults

`func NewUpdateSecurityGroupsRequestWithDefaults() *UpdateSecurityGroupsRequest`

NewUpdateSecurityGroupsRequestWithDefaults instantiates a new UpdateSecurityGroupsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSecurityGroup

`func (o *UpdateSecurityGroupsRequest) GetSecurityGroup() UpdateSecurityGroupsRequestSecurityGroup`

GetSecurityGroup returns the SecurityGroup field if non-nil, zero value otherwise.

### GetSecurityGroupOk

`func (o *UpdateSecurityGroupsRequest) GetSecurityGroupOk() (*UpdateSecurityGroupsRequestSecurityGroup, bool)`

GetSecurityGroupOk returns a tuple with the SecurityGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecurityGroup

`func (o *UpdateSecurityGroupsRequest) SetSecurityGroup(v UpdateSecurityGroupsRequestSecurityGroup)`

SetSecurityGroup sets SecurityGroup field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


