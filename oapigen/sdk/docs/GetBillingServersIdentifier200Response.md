# GetBillingServersIdentifier200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingInfo** | Pointer to [**GetBillingServersIdentifier200ResponseAllOfBillingInfo**](GetBillingServersIdentifier200ResponseAllOfBillingInfo.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewGetBillingServersIdentifier200Response

`func NewGetBillingServersIdentifier200Response() *GetBillingServersIdentifier200Response`

NewGetBillingServersIdentifier200Response instantiates a new GetBillingServersIdentifier200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetBillingServersIdentifier200ResponseWithDefaults

`func NewGetBillingServersIdentifier200ResponseWithDefaults() *GetBillingServersIdentifier200Response`

NewGetBillingServersIdentifier200ResponseWithDefaults instantiates a new GetBillingServersIdentifier200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingInfo

`func (o *GetBillingServersIdentifier200Response) GetBillingInfo() GetBillingServersIdentifier200ResponseAllOfBillingInfo`

GetBillingInfo returns the BillingInfo field if non-nil, zero value otherwise.

### GetBillingInfoOk

`func (o *GetBillingServersIdentifier200Response) GetBillingInfoOk() (*GetBillingServersIdentifier200ResponseAllOfBillingInfo, bool)`

GetBillingInfoOk returns a tuple with the BillingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingInfo

`func (o *GetBillingServersIdentifier200Response) SetBillingInfo(v GetBillingServersIdentifier200ResponseAllOfBillingInfo)`

SetBillingInfo sets BillingInfo field to given value.

### HasBillingInfo

`func (o *GetBillingServersIdentifier200Response) HasBillingInfo() bool`

HasBillingInfo returns a boolean if a field has been set.

### GetSuccess

`func (o *GetBillingServersIdentifier200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *GetBillingServersIdentifier200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *GetBillingServersIdentifier200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *GetBillingServersIdentifier200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


