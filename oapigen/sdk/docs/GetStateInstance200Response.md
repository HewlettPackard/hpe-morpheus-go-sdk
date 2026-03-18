# GetStateInstance200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Instance** | Pointer to [**GetStateInstance200ResponseAllOfInstance**](GetStateInstance200ResponseAllOfInstance.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetStateInstance200Response

`func NewGetStateInstance200Response() *GetStateInstance200Response`

NewGetStateInstance200Response instantiates a new GetStateInstance200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetStateInstance200ResponseWithDefaults

`func NewGetStateInstance200ResponseWithDefaults() *GetStateInstance200Response`

NewGetStateInstance200ResponseWithDefaults instantiates a new GetStateInstance200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstance

`func (o *GetStateInstance200Response) GetInstance() GetStateInstance200ResponseAllOfInstance`

GetInstance returns the Instance field if non-nil, zero value otherwise.

### GetInstanceOk

`func (o *GetStateInstance200Response) GetInstanceOk() (*GetStateInstance200ResponseAllOfInstance, bool)`

GetInstanceOk returns a tuple with the Instance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstance

`func (o *GetStateInstance200Response) SetInstance(v GetStateInstance200ResponseAllOfInstance)`

SetInstance sets Instance field to given value.

### HasInstance

`func (o *GetStateInstance200Response) HasInstance() bool`

HasInstance returns a boolean if a field has been set.

### GetSuccess

`func (o *GetStateInstance200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetStateInstance200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetStateInstance200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetStateInstance200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


