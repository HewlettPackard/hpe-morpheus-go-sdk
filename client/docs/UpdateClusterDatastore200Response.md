# UpdateClusterDatastore200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Datastore** | Pointer to [**SaveClusterDatastore200ResponseAllOfCluster**](SaveClusterDatastore200ResponseAllOfCluster.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateClusterDatastore200Response

`func NewUpdateClusterDatastore200Response() *UpdateClusterDatastore200Response`

NewUpdateClusterDatastore200Response instantiates a new UpdateClusterDatastore200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateClusterDatastore200ResponseWithDefaults

`func NewUpdateClusterDatastore200ResponseWithDefaults() *UpdateClusterDatastore200Response`

NewUpdateClusterDatastore200ResponseWithDefaults instantiates a new UpdateClusterDatastore200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDatastore

`func (o *UpdateClusterDatastore200Response) GetDatastore() SaveClusterDatastore200ResponseAllOfCluster`

GetDatastore returns the Datastore field if non-nil, zero value otherwise.

### GetDatastoreOk

`func (o *UpdateClusterDatastore200Response) GetDatastoreOk() (*SaveClusterDatastore200ResponseAllOfCluster, bool)`

GetDatastoreOk returns a tuple with the Datastore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDatastore

`func (o *UpdateClusterDatastore200Response) SetDatastore(v SaveClusterDatastore200ResponseAllOfCluster)`

SetDatastore sets Datastore field to given value.

### HasDatastore

`func (o *UpdateClusterDatastore200Response) HasDatastore() bool`

HasDatastore returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateClusterDatastore200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateClusterDatastore200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateClusterDatastore200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateClusterDatastore200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


