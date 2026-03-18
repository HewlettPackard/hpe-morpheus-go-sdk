# GetInstanceHistory200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Processes** | Pointer to [**[]GetInstanceHistory200ResponseAllOfProcessesInner**](GetInstanceHistory200ResponseAllOfProcessesInner.md) |  | [optional] 
**Meta** | Pointer to [**GetInstanceHistory200ResponseAllOfMeta**](GetInstanceHistory200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewGetInstanceHistory200Response

`func NewGetInstanceHistory200Response() *GetInstanceHistory200Response`

NewGetInstanceHistory200Response instantiates a new GetInstanceHistory200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInstanceHistory200ResponseWithDefaults

`func NewGetInstanceHistory200ResponseWithDefaults() *GetInstanceHistory200Response`

NewGetInstanceHistory200ResponseWithDefaults instantiates a new GetInstanceHistory200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcesses

`func (o *GetInstanceHistory200Response) GetProcesses() []GetInstanceHistory200ResponseAllOfProcessesInner`

GetProcesses returns the Processes field if non-nil, zero value otherwise.

### GetProcessesOk

`func (o *GetInstanceHistory200Response) GetProcessesOk() (*[]GetInstanceHistory200ResponseAllOfProcessesInner, bool)`

GetProcessesOk returns a tuple with the Processes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcesses

`func (o *GetInstanceHistory200Response) SetProcesses(v []GetInstanceHistory200ResponseAllOfProcessesInner)`

SetProcesses sets Processes field to given value.

### HasProcesses

`func (o *GetInstanceHistory200Response) HasProcesses() bool`

HasProcesses returns a boolean if a field has been set.

### GetMeta

`func (o *GetInstanceHistory200Response) GetMeta() GetInstanceHistory200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GetInstanceHistory200Response) GetMetaOk() (*GetInstanceHistory200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GetInstanceHistory200Response) SetMeta(v GetInstanceHistory200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GetInstanceHistory200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


