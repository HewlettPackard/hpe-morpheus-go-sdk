# ListCodeRepositories200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Success** | Pointer to **bool** |  | [optional] 
**Data** | Pointer to [**[]ListCodeRepositories200ResponseAllOfDataInner**](ListCodeRepositories200ResponseAllOfDataInner.md) |  | [optional] 

## Methods

### NewListCodeRepositories200Response

`func NewListCodeRepositories200Response() *ListCodeRepositories200Response`

NewListCodeRepositories200Response instantiates a new ListCodeRepositories200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListCodeRepositories200ResponseWithDefaults

`func NewListCodeRepositories200ResponseWithDefaults() *ListCodeRepositories200Response`

NewListCodeRepositories200ResponseWithDefaults instantiates a new ListCodeRepositories200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSuccess

`func (o *ListCodeRepositories200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ListCodeRepositories200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ListCodeRepositories200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ListCodeRepositories200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetData

`func (o *ListCodeRepositories200Response) GetData() []ListCodeRepositories200ResponseAllOfDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ListCodeRepositories200Response) GetDataOk() (*[]ListCodeRepositories200ResponseAllOfDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ListCodeRepositories200Response) SetData(v []ListCodeRepositories200ResponseAllOfDataInner)`

SetData sets Data field to given value.

### HasData

`func (o *ListCodeRepositories200Response) HasData() bool`

HasData returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


