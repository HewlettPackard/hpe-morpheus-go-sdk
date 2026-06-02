# ListHistory200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Processes** | Pointer to [**[]ListHistory200ResponseAllOfProcessesInner**](ListHistory200ResponseAllOfProcessesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListApprovals200ResponseAllOfMeta**](ListApprovals200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListHistory200Response

`func NewListHistory200Response() *ListHistory200Response`

NewListHistory200Response instantiates a new ListHistory200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetProcesses

`func (o *ListHistory200Response) GetProcesses() []ListHistory200ResponseAllOfProcessesInner`

GetProcesses returns the Processes field if non-nil, zero value otherwise.

### GetProcessesOk

`func (o *ListHistory200Response) GetProcessesOk() (*[]ListHistory200ResponseAllOfProcessesInner, bool)`

GetProcessesOk returns a tuple with the Processes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcesses

`func (o *ListHistory200Response) SetProcesses(v []ListHistory200ResponseAllOfProcessesInner)`

SetProcesses sets Processes field to given value.

### HasProcesses

`func (o *ListHistory200Response) HasProcesses() bool`

HasProcesses returns a boolean if a field has been set.

### GetMeta

`func (o *ListHistory200Response) GetMeta() ListApprovals200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListHistory200Response) GetMetaOk() (*ListApprovals200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListHistory200Response) SetMeta(v ListApprovals200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListHistory200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


