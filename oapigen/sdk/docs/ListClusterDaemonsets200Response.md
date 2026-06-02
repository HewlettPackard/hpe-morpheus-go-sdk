# ListClusterDaemonsets200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Daemonsets** | Pointer to [**[]ListClusterDaemonsets200ResponseAllOfDaemonsetsInner**](ListClusterDaemonsets200ResponseAllOfDaemonsetsInner.md) |  | [optional] 
**Meta** | Pointer to [**ListApprovals200ResponseAllOfMeta**](ListApprovals200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListClusterDaemonsets200Response

`func NewListClusterDaemonsets200Response() *ListClusterDaemonsets200Response`

NewListClusterDaemonsets200Response instantiates a new ListClusterDaemonsets200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetDaemonsets

`func (o *ListClusterDaemonsets200Response) GetDaemonsets() []ListClusterDaemonsets200ResponseAllOfDaemonsetsInner`

GetDaemonsets returns the Daemonsets field if non-nil, zero value otherwise.

### GetDaemonsetsOk

`func (o *ListClusterDaemonsets200Response) GetDaemonsetsOk() (*[]ListClusterDaemonsets200ResponseAllOfDaemonsetsInner, bool)`

GetDaemonsetsOk returns a tuple with the Daemonsets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDaemonsets

`func (o *ListClusterDaemonsets200Response) SetDaemonsets(v []ListClusterDaemonsets200ResponseAllOfDaemonsetsInner)`

SetDaemonsets sets Daemonsets field to given value.

### HasDaemonsets

`func (o *ListClusterDaemonsets200Response) HasDaemonsets() bool`

HasDaemonsets returns a boolean if a field has been set.

### GetMeta

`func (o *ListClusterDaemonsets200Response) GetMeta() ListApprovals200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListClusterDaemonsets200Response) GetMetaOk() (*ListApprovals200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListClusterDaemonsets200Response) SetMeta(v ListApprovals200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListClusterDaemonsets200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


