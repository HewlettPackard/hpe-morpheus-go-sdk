# ListApps200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Apps** | Pointer to [**[]ListApps200ResponseAllOfAppsInner**](ListApps200ResponseAllOfAppsInner.md) |  | [optional] 
**Stats** | Pointer to [**ListApps200ResponseAllOfStats**](ListApps200ResponseAllOfStats.md) |  | [optional] 
**Meta** | Pointer to [**ListApprovals200ResponseAllOfMeta**](ListApprovals200ResponseAllOfMeta.md) |  | [optional] 

## Methods

### NewListApps200Response

`func NewListApps200Response() *ListApps200Response`

NewListApps200Response instantiates a new ListApps200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### GetApps

`func (o *ListApps200Response) GetApps() []ListApps200ResponseAllOfAppsInner`

GetApps returns the Apps field if non-nil, zero value otherwise.

### GetAppsOk

`func (o *ListApps200Response) GetAppsOk() (*[]ListApps200ResponseAllOfAppsInner, bool)`

GetAppsOk returns a tuple with the Apps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApps

`func (o *ListApps200Response) SetApps(v []ListApps200ResponseAllOfAppsInner)`

SetApps sets Apps field to given value.

### HasApps

`func (o *ListApps200Response) HasApps() bool`

HasApps returns a boolean if a field has been set.

### GetStats

`func (o *ListApps200Response) GetStats() ListApps200ResponseAllOfStats`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *ListApps200Response) GetStatsOk() (*ListApps200ResponseAllOfStats, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *ListApps200Response) SetStats(v ListApps200ResponseAllOfStats)`

SetStats sets Stats field to given value.

### HasStats

`func (o *ListApps200Response) HasStats() bool`

HasStats returns a boolean if a field has been set.

### GetMeta

`func (o *ListApps200Response) GetMeta() ListApprovals200ResponseAllOfMeta`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *ListApps200Response) GetMetaOk() (*ListApprovals200ResponseAllOfMeta, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *ListApps200Response) SetMeta(v ListApprovals200ResponseAllOfMeta)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *ListApps200Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


