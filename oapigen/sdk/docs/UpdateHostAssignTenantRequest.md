# UpdateHostAssignTenantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MoveAssociatedInstances** | Pointer to **bool** | Move associated instances to specified Tenant account. | [optional] [default to true]

## Methods

### NewUpdateHostAssignTenantRequest

`func NewUpdateHostAssignTenantRequest() *UpdateHostAssignTenantRequest`

NewUpdateHostAssignTenantRequest instantiates a new UpdateHostAssignTenantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateHostAssignTenantRequestWithDefaults

`func NewUpdateHostAssignTenantRequestWithDefaults() *UpdateHostAssignTenantRequest`

NewUpdateHostAssignTenantRequestWithDefaults instantiates a new UpdateHostAssignTenantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMoveAssociatedInstances

`func (o *UpdateHostAssignTenantRequest) GetMoveAssociatedInstances() bool`

GetMoveAssociatedInstances returns the MoveAssociatedInstances field if non-nil, zero value otherwise.

### GetMoveAssociatedInstancesOk

`func (o *UpdateHostAssignTenantRequest) GetMoveAssociatedInstancesOk() (*bool, bool)`

GetMoveAssociatedInstancesOk returns a tuple with the MoveAssociatedInstances field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMoveAssociatedInstances

`func (o *UpdateHostAssignTenantRequest) SetMoveAssociatedInstances(v bool)`

SetMoveAssociatedInstances sets MoveAssociatedInstances field to given value.

### HasMoveAssociatedInstances

`func (o *UpdateHostAssignTenantRequest) HasMoveAssociatedInstances() bool`

HasMoveAssociatedInstances returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


