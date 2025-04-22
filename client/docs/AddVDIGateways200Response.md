# AddVDIGateways200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**VdiGateway** | Pointer to [**ListVDIGateways200ResponseAllOfVdiGatewaysInner**](ListVDIGateways200ResponseAllOfVdiGatewaysInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddVDIGateways200Response

`func NewAddVDIGateways200Response() *AddVDIGateways200Response`

NewAddVDIGateways200Response instantiates a new AddVDIGateways200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddVDIGateways200ResponseWithDefaults

`func NewAddVDIGateways200ResponseWithDefaults() *AddVDIGateways200Response`

NewAddVDIGateways200ResponseWithDefaults instantiates a new AddVDIGateways200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVdiGateway

`func (o *AddVDIGateways200Response) GetVdiGateway() ListVDIGateways200ResponseAllOfVdiGatewaysInner`

GetVdiGateway returns the VdiGateway field if non-nil, zero value otherwise.

### GetVdiGatewayOk

`func (o *AddVDIGateways200Response) GetVdiGatewayOk() (*ListVDIGateways200ResponseAllOfVdiGatewaysInner, bool)`

GetVdiGatewayOk returns a tuple with the VdiGateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVdiGateway

`func (o *AddVDIGateways200Response) SetVdiGateway(v ListVDIGateways200ResponseAllOfVdiGatewaysInner)`

SetVdiGateway sets VdiGateway field to given value.

### HasVdiGateway

`func (o *AddVDIGateways200Response) HasVdiGateway() bool`

HasVdiGateway returns a boolean if a field has been set.

### GetSuccess

`func (o *AddVDIGateways200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddVDIGateways200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddVDIGateways200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddVDIGateways200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


