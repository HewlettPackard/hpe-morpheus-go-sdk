# ListBillingZone200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingInfo** | Pointer to [**ListBillingZone200ResponseAllOfBillingInfo**](ListBillingZone200ResponseAllOfBillingInfo.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewListBillingZone200Response

`func NewListBillingZone200Response() *ListBillingZone200Response`

NewListBillingZone200Response instantiates a new ListBillingZone200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBillingZone200ResponseWithDefaults

`func NewListBillingZone200ResponseWithDefaults() *ListBillingZone200Response`

NewListBillingZone200ResponseWithDefaults instantiates a new ListBillingZone200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingInfo

`func (o *ListBillingZone200Response) GetBillingInfo() ListBillingZone200ResponseAllOfBillingInfo`

GetBillingInfo returns the BillingInfo field if non-nil, zero value otherwise.

### GetBillingInfoOk

`func (o *ListBillingZone200Response) GetBillingInfoOk() (*ListBillingZone200ResponseAllOfBillingInfo, bool)`

GetBillingInfoOk returns a tuple with the BillingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingInfo

`func (o *ListBillingZone200Response) SetBillingInfo(v ListBillingZone200ResponseAllOfBillingInfo)`

SetBillingInfo sets BillingInfo field to given value.

### HasBillingInfo

`func (o *ListBillingZone200Response) HasBillingInfo() bool`

HasBillingInfo returns a boolean if a field has been set.

### GetSuccess

`func (o *ListBillingZone200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ListBillingZone200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ListBillingZone200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ListBillingZone200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


