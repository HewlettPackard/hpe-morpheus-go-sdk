# AddChecks200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Check** | Pointer to [**GetAlerts200ResponseAllOfChecksInner**](GetAlerts200ResponseAllOfChecksInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewAddChecks200Response

`func NewAddChecks200Response() *AddChecks200Response`

NewAddChecks200Response instantiates a new AddChecks200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddChecks200ResponseWithDefaults

`func NewAddChecks200ResponseWithDefaults() *AddChecks200Response`

NewAddChecks200ResponseWithDefaults instantiates a new AddChecks200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCheck

`func (o *AddChecks200Response) GetCheck() GetAlerts200ResponseAllOfChecksInner`

GetCheck returns the Check field if non-nil, zero value otherwise.

### GetCheckOk

`func (o *AddChecks200Response) GetCheckOk() (*GetAlerts200ResponseAllOfChecksInner, bool)`

GetCheckOk returns a tuple with the Check field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCheck

`func (o *AddChecks200Response) SetCheck(v GetAlerts200ResponseAllOfChecksInner)`

SetCheck sets Check field to given value.

### HasCheck

`func (o *AddChecks200Response) HasCheck() bool`

HasCheck returns a boolean if a field has been set.

### GetSuccess

`func (o *AddChecks200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *AddChecks200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *AddChecks200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *AddChecks200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


