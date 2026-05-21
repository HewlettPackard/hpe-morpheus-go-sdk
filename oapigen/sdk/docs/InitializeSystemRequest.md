# InitializeSystemRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**System** | Pointer to [**InitializeSystemRequestSystem**](InitializeSystemRequestSystem.md) |  | [optional] 

## Methods

### NewInitializeSystemRequest

`func NewInitializeSystemRequest() *InitializeSystemRequest`

NewInitializeSystemRequest instantiates a new InitializeSystemRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInitializeSystemRequestWithDefaults

`func NewInitializeSystemRequestWithDefaults() *InitializeSystemRequest`

NewInitializeSystemRequestWithDefaults instantiates a new InitializeSystemRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSystem

`func (o *InitializeSystemRequest) GetSystem() InitializeSystemRequestSystem`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *InitializeSystemRequest) GetSystemOk() (*InitializeSystemRequestSystem, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *InitializeSystemRequest) SetSystem(v InitializeSystemRequestSystem)`

SetSystem sets System field to given value.

### HasSystem

`func (o *InitializeSystemRequest) HasSystem() bool`

HasSystem returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


