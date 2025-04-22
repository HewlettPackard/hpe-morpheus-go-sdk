# ListBillingInstances200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BillingInfo** | Pointer to [**ListBillingInstances200ResponseAllOfBillingInfo**](ListBillingInstances200ResponseAllOfBillingInfo.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewListBillingInstances200Response

`func NewListBillingInstances200Response() *ListBillingInstances200Response`

NewListBillingInstances200Response instantiates a new ListBillingInstances200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListBillingInstances200ResponseWithDefaults

`func NewListBillingInstances200ResponseWithDefaults() *ListBillingInstances200Response`

NewListBillingInstances200ResponseWithDefaults instantiates a new ListBillingInstances200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBillingInfo

`func (o *ListBillingInstances200Response) GetBillingInfo() ListBillingInstances200ResponseAllOfBillingInfo`

GetBillingInfo returns the BillingInfo field if non-nil, zero value otherwise.

### GetBillingInfoOk

`func (o *ListBillingInstances200Response) GetBillingInfoOk() (*ListBillingInstances200ResponseAllOfBillingInfo, bool)`

GetBillingInfoOk returns a tuple with the BillingInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBillingInfo

`func (o *ListBillingInstances200Response) SetBillingInfo(v ListBillingInstances200ResponseAllOfBillingInfo)`

SetBillingInfo sets BillingInfo field to given value.

### HasBillingInfo

`func (o *ListBillingInstances200Response) HasBillingInfo() bool`

HasBillingInfo returns a boolean if a field has been set.

### GetSuccess

`func (o *ListBillingInstances200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *ListBillingInstances200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *ListBillingInstances200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *ListBillingInstances200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


