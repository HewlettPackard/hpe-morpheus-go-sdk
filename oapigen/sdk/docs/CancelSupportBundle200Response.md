# CancelSupportBundle200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportBundle** | Pointer to [**CancelSupportBundle200ResponseAllOfSupportBundle**](CancelSupportBundle200ResponseAllOfSupportBundle.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewCancelSupportBundle200Response

`func NewCancelSupportBundle200Response() *CancelSupportBundle200Response`

NewCancelSupportBundle200Response instantiates a new CancelSupportBundle200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCancelSupportBundle200ResponseWithDefaults

`func NewCancelSupportBundle200ResponseWithDefaults() *CancelSupportBundle200Response`

NewCancelSupportBundle200ResponseWithDefaults instantiates a new CancelSupportBundle200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportBundle

`func (o *CancelSupportBundle200Response) GetSupportBundle() CancelSupportBundle200ResponseAllOfSupportBundle`

GetSupportBundle returns the SupportBundle field if non-nil, zero value otherwise.

### GetSupportBundleOk

`func (o *CancelSupportBundle200Response) GetSupportBundleOk() (*CancelSupportBundle200ResponseAllOfSupportBundle, bool)`

GetSupportBundleOk returns a tuple with the SupportBundle field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportBundle

`func (o *CancelSupportBundle200Response) SetSupportBundle(v CancelSupportBundle200ResponseAllOfSupportBundle)`

SetSupportBundle sets SupportBundle field to given value.

### HasSupportBundle

`func (o *CancelSupportBundle200Response) HasSupportBundle() bool`

HasSupportBundle returns a boolean if a field has been set.

### GetSuccess

`func (o *CancelSupportBundle200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *CancelSupportBundle200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *CancelSupportBundle200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *CancelSupportBundle200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


