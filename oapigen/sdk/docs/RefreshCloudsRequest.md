# RefreshCloudsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mode** | Pointer to **string** | Refresh Mode. Run the &#x60;daily&#x60; or &#x60;costing&#x60; job instead of the default &#x60;hourly&#x60; refresh job. | [optional] 
**Rebuild** | Pointer to **string** | Rebuild. Pass &#x60;true&#x60; to purge existing invoices for the period before refreshing. Only applies to mode &#x60;costing&#x60;. | [optional] 
**Period** | Pointer to **string** | Period. Invoice billing period to refresh in the format &#x60;YYYYMM&#x60;. The default period is the current month. Only applies to mode &#x60;costing&#x60;. | [optional] 

## Methods

### NewRefreshCloudsRequest

`func NewRefreshCloudsRequest() *RefreshCloudsRequest`

NewRefreshCloudsRequest instantiates a new RefreshCloudsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefreshCloudsRequestWithDefaults

`func NewRefreshCloudsRequestWithDefaults() *RefreshCloudsRequest`

NewRefreshCloudsRequestWithDefaults instantiates a new RefreshCloudsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMode

`func (o *RefreshCloudsRequest) GetMode() string`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *RefreshCloudsRequest) GetModeOk() (*string, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *RefreshCloudsRequest) SetMode(v string)`

SetMode sets Mode field to given value.

### HasMode

`func (o *RefreshCloudsRequest) HasMode() bool`

HasMode returns a boolean if a field has been set.

### GetRebuild

`func (o *RefreshCloudsRequest) GetRebuild() string`

GetRebuild returns the Rebuild field if non-nil, zero value otherwise.

### GetRebuildOk

`func (o *RefreshCloudsRequest) GetRebuildOk() (*string, bool)`

GetRebuildOk returns a tuple with the Rebuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRebuild

`func (o *RefreshCloudsRequest) SetRebuild(v string)`

SetRebuild sets Rebuild field to given value.

### HasRebuild

`func (o *RefreshCloudsRequest) HasRebuild() bool`

HasRebuild returns a boolean if a field has been set.

### GetPeriod

`func (o *RefreshCloudsRequest) GetPeriod() string`

GetPeriod returns the Period field if non-nil, zero value otherwise.

### GetPeriodOk

`func (o *RefreshCloudsRequest) GetPeriodOk() (*string, bool)`

GetPeriodOk returns a tuple with the Period field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeriod

`func (o *RefreshCloudsRequest) SetPeriod(v string)`

SetPeriod sets Period field to given value.

### HasPeriod

`func (o *RefreshCloudsRequest) HasPeriod() bool`

HasPeriod returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


