# GetGuidanceStats200ResponseStats

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Total** | Pointer to **int64** |  | [optional] 
**Savings** | Pointer to [**ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfSavings**](ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfSavings.md) |  | [optional] 
**Severity** | Pointer to [**GetGuidanceStats200ResponseStatsSeverity**](GetGuidanceStats200ResponseStatsSeverity.md) |  | [optional] 
**Type** | Pointer to [**GetGuidanceStats200ResponseStatsType**](GetGuidanceStats200ResponseStatsType.md) |  | [optional] 

## Methods

### NewGetGuidanceStats200ResponseStats

`func NewGetGuidanceStats200ResponseStats() *GetGuidanceStats200ResponseStats`

NewGetGuidanceStats200ResponseStats instantiates a new GetGuidanceStats200ResponseStats object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGuidanceStats200ResponseStatsWithDefaults

`func NewGetGuidanceStats200ResponseStatsWithDefaults() *GetGuidanceStats200ResponseStats`

NewGetGuidanceStats200ResponseStatsWithDefaults instantiates a new GetGuidanceStats200ResponseStats object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTotal

`func (o *GetGuidanceStats200ResponseStats) GetTotal() int64`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetGuidanceStats200ResponseStats) GetTotalOk() (*int64, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetGuidanceStats200ResponseStats) SetTotal(v int64)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *GetGuidanceStats200ResponseStats) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

### GetSavings

`func (o *GetGuidanceStats200ResponseStats) GetSavings() ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfSavings`

GetSavings returns the Savings field if non-nil, zero value otherwise.

### GetSavingsOk

`func (o *GetGuidanceStats200ResponseStats) GetSavingsOk() (*ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfSavings, bool)`

GetSavingsOk returns a tuple with the Savings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSavings

`func (o *GetGuidanceStats200ResponseStats) SetSavings(v ListGuidances200ResponseAllOfDiscoveriesInnerAnyOfSavings)`

SetSavings sets Savings field to given value.

### HasSavings

`func (o *GetGuidanceStats200ResponseStats) HasSavings() bool`

HasSavings returns a boolean if a field has been set.

### GetSeverity

`func (o *GetGuidanceStats200ResponseStats) GetSeverity() GetGuidanceStats200ResponseStatsSeverity`

GetSeverity returns the Severity field if non-nil, zero value otherwise.

### GetSeverityOk

`func (o *GetGuidanceStats200ResponseStats) GetSeverityOk() (*GetGuidanceStats200ResponseStatsSeverity, bool)`

GetSeverityOk returns a tuple with the Severity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSeverity

`func (o *GetGuidanceStats200ResponseStats) SetSeverity(v GetGuidanceStats200ResponseStatsSeverity)`

SetSeverity sets Severity field to given value.

### HasSeverity

`func (o *GetGuidanceStats200ResponseStats) HasSeverity() bool`

HasSeverity returns a boolean if a field has been set.

### GetType

`func (o *GetGuidanceStats200ResponseStats) GetType() GetGuidanceStats200ResponseStatsType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GetGuidanceStats200ResponseStats) GetTypeOk() (*GetGuidanceStats200ResponseStatsType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GetGuidanceStats200ResponseStats) SetType(v GetGuidanceStats200ResponseStatsType)`

SetType sets Type field to given value.

### HasType

`func (o *GetGuidanceStats200ResponseStats) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


