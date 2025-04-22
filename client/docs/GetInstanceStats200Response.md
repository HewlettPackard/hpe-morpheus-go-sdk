# GetInstanceStats200Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstanceStats** | Pointer to [**GetInstanceStats200ResponseInstanceStats**](GetInstanceStats200ResponseInstanceStats.md) |  | [optional] 
**ZoneIds** | Pointer to **[]int64** | Array of Cloud IDs that are included in the stats. By default all the clouds the user has access to are returned. | [optional] 

## Methods

### NewGetInstanceStats200Response

`func NewGetInstanceStats200Response() *GetInstanceStats200Response`

NewGetInstanceStats200Response instantiates a new GetInstanceStats200Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetInstanceStats200ResponseWithDefaults

`func NewGetInstanceStats200ResponseWithDefaults() *GetInstanceStats200Response`

NewGetInstanceStats200ResponseWithDefaults instantiates a new GetInstanceStats200Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstanceStats

`func (o *GetInstanceStats200Response) GetInstanceStats() GetInstanceStats200ResponseInstanceStats`

GetInstanceStats returns the InstanceStats field if non-nil, zero value otherwise.

### GetInstanceStatsOk

`func (o *GetInstanceStats200Response) GetInstanceStatsOk() (*GetInstanceStats200ResponseInstanceStats, bool)`

GetInstanceStatsOk returns a tuple with the InstanceStats field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstanceStats

`func (o *GetInstanceStats200Response) SetInstanceStats(v GetInstanceStats200ResponseInstanceStats)`

SetInstanceStats sets InstanceStats field to given value.

### HasInstanceStats

`func (o *GetInstanceStats200Response) HasInstanceStats() bool`

HasInstanceStats returns a boolean if a field has been set.

### GetZoneIds

`func (o *GetInstanceStats200Response) GetZoneIds() []int64`

GetZoneIds returns the ZoneIds field if non-nil, zero value otherwise.

### GetZoneIdsOk

`func (o *GetInstanceStats200Response) GetZoneIdsOk() (*[]int64, bool)`

GetZoneIdsOk returns a tuple with the ZoneIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneIds

`func (o *GetInstanceStats200Response) SetZoneIds(v []int64)`

SetZoneIds sets ZoneIds field to given value.

### HasZoneIds

`func (o *GetInstanceStats200Response) HasZoneIds() bool`

HasZoneIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


