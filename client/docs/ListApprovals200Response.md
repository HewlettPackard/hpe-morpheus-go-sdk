# ListApprovals200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Approvals** | Pointer to [**[]ListApprovals200ResponseAllOfApprovalsInner**](ListApprovals200ResponseAllOfApprovalsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListActivity200ResponseAllOfMeta**](ListActivity200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListApprovals200Response

`func NewListApprovals200Response() *ListApprovals200Response`

NewListApprovals200Response instantiates a new ListApprovals200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewListApprovals200ResponseWithDefaults

`func NewListApprovals200ResponseWithDefaults() *ListApprovals200Response`

NewListApprovals200ResponseWithDefaults instantiates a new ListApprovals200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApprovals

`func (o *ListApprovals200Response) GetApprovals() []ListApprovals200ResponseAllOfApprovalsInner`

GetApprovals returns the Approvals field if non-nil, zero value otherwise.

### GetApprovalsOk

`func (o *ListApprovals200Response) GetApprovalsOk() (*[]ListApprovals200ResponseAllOfApprovalsInner, bool)`

GetApprovalsOk returns a tuple with the Approvals field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApprovals

`func (o *ListApprovals200Response) SetApprovals(v []ListApprovals200ResponseAllOfApprovalsInner)`

SetApprovals sets Approvals field to given value.

### HasApprovals

`func (o *ListApprovals200Response) HasApprovals() bool`

HasApprovals returns a boolean if a field has been set.

### GetMeta

`func (o *ListApprovals200Response) GetMeta() ListActivity200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListApprovals200Response) GetMetaOk() (*ListActivity200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListApprovals200Response) SetMeta(v ListActivity200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListApprovals200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


