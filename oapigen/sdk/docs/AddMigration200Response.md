# AddMigration200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Migration** | Pointer to [**AddMigration200ResponseAnyOfMigration**](AddMigration200ResponseAnyOfMigration.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddMigration200Response

`func NewAddMigration200Response() *AddMigration200Response`

NewAddMigration200Response instantiates a new AddMigration200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetMigration

`func (o *AddMigration200Response) GetMigration() AddMigration200ResponseAnyOfMigration`

GetMigration returns the Migration field if non-nil, zero value otherwise.

### GetMigrationOk

`func (o *AddMigration200Response) GetMigrationOk() (*AddMigration200ResponseAnyOfMigration, bool)`

GetMigrationOk returns a tuple with the Migration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMigration

`func (o *AddMigration200Response) SetMigration(v AddMigration200ResponseAnyOfMigration)`

SetMigration sets Migration field to given value.

### HasMigration

`func (o *AddMigration200Response) HasMigration() bool`

HasMigration returns a boolean if a field has been set.

### GetSuccess

`func (o *AddMigration200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddMigration200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddMigration200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddMigration200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


