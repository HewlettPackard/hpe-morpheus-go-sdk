# RemoveInstancesFromControlRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Ids** | Pointer to **[]int64** | Array of Ids of brownfield Instances to be deleted | [optional] 

## Methods

### NewRemoveInstancesFromControlRequest

`func NewRemoveInstancesFromControlRequest() *RemoveInstancesFromControlRequest`

NewRemoveInstancesFromControlRequest instantiates a new RemoveInstancesFromControlRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRemoveInstancesFromControlRequestWithDefaults

`func NewRemoveInstancesFromControlRequestWithDefaults() *RemoveInstancesFromControlRequest`

NewRemoveInstancesFromControlRequestWithDefaults instantiates a new RemoveInstancesFromControlRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIds

`func (o *RemoveInstancesFromControlRequest) GetIds() []int64`

GetIds returns the Ids field if non-nil, zero value otherwise.

### GetIdsOk

`func (o *RemoveInstancesFromControlRequest) GetIdsOk() (*[]int64, bool)`

GetIdsOk returns a tuple with the Ids field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIds

`func (o *RemoveInstancesFromControlRequest) SetIds(v []int64)`

SetIds sets Ids field to given value.

### HasIds

`func (o *RemoveInstancesFromControlRequest) HasIds() bool`

HasIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


