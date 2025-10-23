# GetClusterHistory200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Processes** | Pointer to [**[]GetClusterHistory200ResponseAllOfProcessesInner**](GetClusterHistory200ResponseAllOfProcessesInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewGetClusterHistory200Response

`func NewGetClusterHistory200Response() *GetClusterHistory200Response`

NewGetClusterHistory200Response instantiates a new GetClusterHistory200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetClusterHistory200ResponseWithDefaults

`func NewGetClusterHistory200ResponseWithDefaults() *GetClusterHistory200Response`

NewGetClusterHistory200ResponseWithDefaults instantiates a new GetClusterHistory200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcesses

`func (o *GetClusterHistory200Response) GetProcesses() []GetClusterHistory200ResponseAllOfProcessesInner`

GetProcesses returns the Processes field if non-nil, zero value otherwise.

### GetProcessesOk

`func (o *GetClusterHistory200Response) GetProcessesOk() (*[]GetClusterHistory200ResponseAllOfProcessesInner, bool)`

GetProcessesOk returns a tuple with the Processes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcesses

`func (o *GetClusterHistory200Response) SetProcesses(v []GetClusterHistory200ResponseAllOfProcessesInner)`

SetProcesses sets Processes field to given value.

### HasProcesses

`func (o *GetClusterHistory200Response) HasProcesses() bool`

HasProcesses returns a boolean if a field has been set.

### GetMeta

`func (o *GetClusterHistory200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetClusterHistory200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetClusterHistory200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetClusterHistory200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


