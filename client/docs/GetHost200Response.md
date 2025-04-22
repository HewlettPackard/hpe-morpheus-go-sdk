# GetHost200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Server** | Pointer to [**ListHosts200ResponseAllOfServersInner**](ListHosts200ResponseAllOfServersInner.md) |  | [optional] 
**Stats** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewGetHost200Response

`func NewGetHost200Response() *GetHost200Response`

NewGetHost200Response instantiates a new GetHost200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetHost200ResponseWithDefaults

`func NewGetHost200ResponseWithDefaults() *GetHost200Response`

NewGetHost200ResponseWithDefaults instantiates a new GetHost200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetServer

`func (o *GetHost200Response) GetServer() ListHosts200ResponseAllOfServersInner`

GetServer returns the Server field if non-nil, zero value otherwise.

### GetServerOk

`func (o *GetHost200Response) GetServerOk() (*ListHosts200ResponseAllOfServersInner, bool)`

GetServerOk returns a tuple with the Server field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServer

`func (o *GetHost200Response) SetServer(v ListHosts200ResponseAllOfServersInner)`

SetServer sets Server field to given value.

### HasServer

`func (o *GetHost200Response) HasServer() bool`

HasServer returns a boolean if a field has been set.

### GetStats

`func (o *GetHost200Response) GetStats() map[string]interface{}`

GetStats returns the Stats field if non-nil, zero value otherwise.

### GetStatsOk

`func (o *GetHost200Response) GetStatsOk() (*map[string]interface{}, bool)`

GetStatsOk returns a tuple with the Stats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStats

`func (o *GetHost200Response) SetStats(v map[string]interface{})`

SetStats sets Stats field to given value.

### HasStats

`func (o *GetHost200Response) HasStats() bool`

HasStats returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


