# UpdateVDIPools200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VdiPool** | Pointer to [**UpdateVDIPools200ResponseAnyOfVdiPool**](UpdateVDIPools200ResponseAnyOfVdiPool.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateVDIPools200Response

`func NewUpdateVDIPools200Response() *UpdateVDIPools200Response`

NewUpdateVDIPools200Response instantiates a new UpdateVDIPools200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateVDIPools200ResponseWithDefaults

`func NewUpdateVDIPools200ResponseWithDefaults() *UpdateVDIPools200Response`

NewUpdateVDIPools200ResponseWithDefaults instantiates a new UpdateVDIPools200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVdiPool

`func (o *UpdateVDIPools200Response) GetVdiPool() UpdateVDIPools200ResponseAnyOfVdiPool`

GetVdiPool returns the VdiPool field if non-nil, zero value otherwise.

### GetVdiPoolOk

`func (o *UpdateVDIPools200Response) GetVdiPoolOk() (*UpdateVDIPools200ResponseAnyOfVdiPool, bool)`

GetVdiPoolOk returns a tuple with the VdiPool field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdiPool

`func (o *UpdateVDIPools200Response) SetVdiPool(v UpdateVDIPools200ResponseAnyOfVdiPool)`

SetVdiPool sets VdiPool field to given value.

### HasVdiPool

`func (o *UpdateVDIPools200Response) HasVdiPool() bool`

HasVdiPool returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateVDIPools200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateVDIPools200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateVDIPools200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateVDIPools200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


