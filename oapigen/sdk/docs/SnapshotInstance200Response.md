# SnapshotInstance200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**ProcessIds** | Pointer to **[]int64** | Process ID(s) to track execution results with, use &#x60;/api/processes/$processId&#x60;. Instances with more than one server will result in multiple processes, one for each snapshot. | [optional] 

## Methods

### NewSnapshotInstance200Response

`func NewSnapshotInstance200Response() *SnapshotInstance200Response`

NewSnapshotInstance200Response instantiates a new SnapshotInstance200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSnapshotInstance200ResponseWithDefaults

`func NewSnapshotInstance200ResponseWithDefaults() *SnapshotInstance200Response`

NewSnapshotInstance200ResponseWithDefaults instantiates a new SnapshotInstance200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *SnapshotInstance200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *SnapshotInstance200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *SnapshotInstance200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *SnapshotInstance200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetProcessIds

`func (o *SnapshotInstance200Response) GetProcessIds() []int64`

GetProcessIds returns the ProcessIds field if non-nil, zero value otherwise.

### GetProcessIdsOk

`func (o *SnapshotInstance200Response) GetProcessIdsOk() (*[]int64, bool)`

GetProcessIdsOk returns a tuple with the ProcessIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessIds

`func (o *SnapshotInstance200Response) SetProcessIds(v []int64)`

SetProcessIds sets ProcessIds field to given value.

### HasProcessIds

`func (o *SnapshotInstance200Response) HasProcessIds() bool`

HasProcessIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


