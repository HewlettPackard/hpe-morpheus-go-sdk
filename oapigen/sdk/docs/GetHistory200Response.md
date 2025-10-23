# GetHistory200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Process** | Pointer to [**ListHistory200ResponseAllOfProcessesInner**](ListHistory200ResponseAllOfProcessesInner.md) |  | [optional] 

## Methods

### NewGetHistory200Response

`func NewGetHistory200Response() *GetHistory200Response`

NewGetHistory200Response instantiates a new GetHistory200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetHistory200ResponseWithDefaults

`func NewGetHistory200ResponseWithDefaults() *GetHistory200Response`

NewGetHistory200ResponseWithDefaults instantiates a new GetHistory200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProcess

`func (o *GetHistory200Response) GetProcess() ListHistory200ResponseAllOfProcessesInner`

GetProcess returns the Process field if non-nil, zero value otherwise.

### GetProcessOk

`func (o *GetHistory200Response) GetProcessOk() (*ListHistory200ResponseAllOfProcessesInner, bool)`

GetProcessOk returns a tuple with the Process field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcess

`func (o *GetHistory200Response) SetProcess(v ListHistory200ResponseAllOfProcessesInner)`

SetProcess sets Process field to given value.

### HasProcess

`func (o *GetHistory200Response) HasProcess() bool`

HasProcess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


