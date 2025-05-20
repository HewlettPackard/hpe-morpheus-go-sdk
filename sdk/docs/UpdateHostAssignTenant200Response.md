# UpdateHostAssignTenant200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server** | Pointer to [**ListHosts200ResponseAllOfServersInner**](ListHosts200ResponseAllOfServersInner.md) |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 

## Methods

### NewUpdateHostAssignTenant200Response

`func NewUpdateHostAssignTenant200Response() *UpdateHostAssignTenant200Response`

NewUpdateHostAssignTenant200Response instantiates a new UpdateHostAssignTenant200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateHostAssignTenant200ResponseWithDefaults

`func NewUpdateHostAssignTenant200ResponseWithDefaults() *UpdateHostAssignTenant200Response`

NewUpdateHostAssignTenant200ResponseWithDefaults instantiates a new UpdateHostAssignTenant200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer

`func (o *UpdateHostAssignTenant200Response) GetServer() ListHosts200ResponseAllOfServersInner`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *UpdateHostAssignTenant200Response) GetServerOk() (*ListHosts200ResponseAllOfServersInner, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *UpdateHostAssignTenant200Response) SetServer(v ListHosts200ResponseAllOfServersInner)`

SetServer sets Server field to given value.

### HasServer

`func (o *UpdateHostAssignTenant200Response) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetSuccess

`func (o *UpdateHostAssignTenant200Response) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *UpdateHostAssignTenant200Response) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *UpdateHostAssignTenant200Response) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *UpdateHostAssignTenant200Response) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


