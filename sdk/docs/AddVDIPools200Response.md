# AddVDIPools200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VdiPool** | Pointer to [**ListVDIPools200ResponseAllOfVdiPoolsInner**](ListVDIPools200ResponseAllOfVdiPoolsInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddVDIPools200Response

`func NewAddVDIPools200Response() *AddVDIPools200Response`

NewAddVDIPools200Response instantiates a new AddVDIPools200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddVDIPools200ResponseWithDefaults

`func NewAddVDIPools200ResponseWithDefaults() *AddVDIPools200Response`

NewAddVDIPools200ResponseWithDefaults instantiates a new AddVDIPools200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVdiPool

`func (o *AddVDIPools200Response) GetVdiPool() ListVDIPools200ResponseAllOfVdiPoolsInner`

GetVdiPool returns the VdiPool field if non-nil, zero value otherwise.

### GetVdiPoolOk

`func (o *AddVDIPools200Response) GetVdiPoolOk() (*ListVDIPools200ResponseAllOfVdiPoolsInner, bool)`

GetVdiPoolOk returns a tuple with the VdiPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdiPool

`func (o *AddVDIPools200Response) SetVdiPool(v ListVDIPools200ResponseAllOfVdiPoolsInner)`

SetVdiPool sets VdiPool field to given value.

### HasVdiPool

`func (o *AddVDIPools200Response) HasVdiPool() bool`

HasVdiPool returns a boolean if a field has been set.

### GetSuccess

`func (o *AddVDIPools200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddVDIPools200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddVDIPools200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddVDIPools200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


