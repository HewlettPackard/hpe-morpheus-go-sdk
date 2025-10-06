# SaveDatastore200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datastore** | Pointer to [**ListDatastores200ResponseAllOfDatastoresInner**](ListDatastores200ResponseAllOfDatastoresInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**ExecutionId** | Pointer to **string** |  | [optional] 

## Methods

### NewSaveDatastore200Response

`func NewSaveDatastore200Response() *SaveDatastore200Response`

NewSaveDatastore200Response instantiates a new SaveDatastore200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSaveDatastore200ResponseWithDefaults

`func NewSaveDatastore200ResponseWithDefaults() *SaveDatastore200Response`

NewSaveDatastore200ResponseWithDefaults instantiates a new SaveDatastore200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastore

`func (o *SaveDatastore200Response) GetDatastore() ListDatastores200ResponseAllOfDatastoresInner`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *SaveDatastore200Response) GetDatastoreOk() (*ListDatastores200ResponseAllOfDatastoresInner, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *SaveDatastore200Response) SetDatastore(v ListDatastores200ResponseAllOfDatastoresInner)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *SaveDatastore200Response) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetSuccess

`func (o *SaveDatastore200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *SaveDatastore200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *SaveDatastore200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *SaveDatastore200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetExecutionId

`func (o *SaveDatastore200Response) GetExecutionId() string`

GetExecutionId returns the ExecutionId field if non-nil, zero value otherwise.

### GetExecutionIdOk

`func (o *SaveDatastore200Response) GetExecutionIdOk() (*string, bool)`

GetExecutionIdOk returns a tuple with the ExecutionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExecutionId

`func (o *SaveDatastore200Response) SetExecutionId(v string)`

SetExecutionId sets ExecutionId field to given value.

### HasExecutionId

`func (o *SaveDatastore200Response) HasExecutionId() bool`

HasExecutionId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


