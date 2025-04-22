# AddClient200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Client** | Pointer to [**ListClients200ResponseAllOfClientsInner**](ListClients200ResponseAllOfClientsInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddClient200Response

`func NewAddClient200Response() *AddClient200Response`

NewAddClient200Response instantiates a new AddClient200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddClient200ResponseWithDefaults

`func NewAddClient200ResponseWithDefaults() *AddClient200Response`

NewAddClient200ResponseWithDefaults instantiates a new AddClient200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClient

`func (o *AddClient200Response) GetClient() ListClients200ResponseAllOfClientsInner`

GetClient returns the Client field if non-nil, zero value otherwise.

### GetClientOk

`func (o *AddClient200Response) GetClientOk() (*ListClients200ResponseAllOfClientsInner, bool)`

GetClientOk returns a tuple with the Client field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClient

`func (o *AddClient200Response) SetClient(v ListClients200ResponseAllOfClientsInner)`

SetClient sets Client field to given value.

### HasClient

`func (o *AddClient200Response) HasClient() bool`

HasClient returns a boolean if a field has been set.

### GetSuccess

`func (o *AddClient200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddClient200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddClient200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddClient200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


